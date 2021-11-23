package main

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"time"

	"log"
	"net"
	"os"

	pb "auction/routeguide"

	"google.golang.org/grpc"
)

var highestBid = 0
var highestBidder string
var isAuctionRunning bool
var requiredAlive = 2
var targetPort int32
var listenPort int32
var lowestPort int32
var currentLeader int32 = -1
var tokenServiceClient pb.TokenServiceClient
var serverPorts = [3]int{5001, 5002, 5003}
var serverID int32
var clientMap map[int]client
var connection *grpc.ClientConn
var mux = &sync.Mutex{}

type tokenServer struct {
	pb.UnimplementedTokenServiceServer
}
type auctionServer struct {
	pb.UnimplementedAuctionServiceServer
}

type client struct {
	service pb.TokenServiceClient
	alive   bool
}

func (s *auctionServer) AskForLeader(ctx context.Context, request *pb.AskRequest) (*pb.LeaderPort, error) {
	return &pb.LeaderPort{LeaderPort: currentLeader}, nil
}

func (s *auctionServer) Bid(ctx context.Context, request *pb.BidRequest) (*pb.BidResponse, error) {
	if currentLeader == listenPort { //Is this node the current leader?
		if !isAuctionRunning {
			return &pb.BidResponse{Status: 3, CurrentHighestBid: int32(highestBid)}, nil
		}
		amount := request.Amount
		if highestBid >= int(amount) {
			return &pb.BidResponse{Status: 2, CurrentHighestBid: amount}, nil
		}
		highestBid = int(amount)
		highestBidder = request.Sender
		ShareDataWithBackups()
		return &pb.BidResponse{Status: 1, CurrentHighestBid: amount}, nil
	}
	return nil, errors.New("Not current leader")
}

func (s *auctionServer) Result(ctx context.Context, request *pb.ResultRequest) (*pb.AuctionData, error) {
	if currentLeader == listenPort { //Is this node the current leader?
		return &pb.AuctionData{HighestBid: int32(highestBid), HighestBidder: highestBidder, IsAuctionRunning: isAuctionRunning}, nil
	}
	return nil, errors.New("Not current leader")
}

func (s *tokenServer) SendHeartbeat(ctx context.Context, request *pb.Empty) (*pb.HeartbeatResponse, error) {
	return &pb.HeartbeatResponse{Status: 1}, nil
}

func (s *tokenServer) FindLeaderRequest(ctx context.Context, request *pb.LeaderRequest) (*pb.Empty, error) {
	log.Println("Received LeaderRequest to", listenPort, "currentLeader:", request.CurrentLeader, "lowestPort:", request.LowestPort)
	if request.CurrentLeader != -1 && request.CurrentLeader < currentLeader || currentLeader == -1 {
		currentLeader = request.CurrentLeader
	}

	if request.LowestPort <= lowestPort && currentLeader == -1 {
		lowestPort = request.LowestPort
		if lowestPort == listenPort {
			currentLeader = listenPort
			log.Println("Found leader, it is me! Starting auction...")
			startAuction()
		}
	}
	return &pb.Empty{}, nil
}

func (s *tokenServer) ShareData(ctx context.Context, request *pb.AuctionData) (*pb.Empty, error) {
	isAuctionRunning = request.GetIsAuctionRunning()
	highestBid = int(request.GetHighestBid())
	highestBidder = request.GetHighestBidder()
	log.Println("Updated data:", isAuctionRunning, highestBid, highestBidder)
	return &pb.Empty{}, nil
}

func ShareDataWithBackups() {
	if listenPort == currentLeader {
		log.Println("Sharing data with backups")
		ctx := context.Background()
		for _, cli := range clientMap {
			if cli.alive {
				cli.service.ShareData(ctx, &pb.AuctionData{HighestBid: int32(highestBid), HighestBidder: highestBidder, IsAuctionRunning: isAuctionRunning})
			}
		}
		log.Println("Data shared")
	}
}

func startAuction() {
	if !isAuctionRunning {
		isAuctionRunning = true
		highestBid = 0
		highestBidder = "None"
	}
	go startTimer()
	log.Println("Started auction")
	ShareDataWithBackups()
}

func startTimer() {
	time.Sleep(60 * time.Second)
	stopAuction()
}
func stopAuction() {
	isAuctionRunning = false
	log.Println("Stopped auction. Having a 30 seconds break")
	ShareDataWithBackups()
	time.Sleep(30 * time.Second)
	startAuction()
}

func main() {
	log.Println("Starting Token-ring service by the FijiAuction team")
	if len(os.Args) == 1 {
		log.Println("Please choose a server between 0 and", len(serverPorts)-1)
		return
	}
	_serverID, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		log.Fatalf("Bad serverId")
	}
	listenPort = int32(serverPorts[_serverID])
	lowestPort = listenPort
	serverID = int32(_serverID)
	targetPort = getTargetPort()

	time.Sleep(10 * time.Second)
	go runServer()
	time.Sleep(1 * time.Second)
	go runClient(targetPort)
	for {
		time.Sleep(2 * time.Second)
	}
}

func getTargetPort() int32 {
	if int(serverID) == len(serverPorts)-1 {
		return int32(serverPorts[0])
	} else {
		return int32(serverPorts[serverID+1])
	}
}

func runClient(target int32) {
	targetPort = target
	address := "localhost:" + strconv.Itoa(int(targetPort))
	_connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	if connection != nil {
		connection.Close()
	}
	connection = _connection
	defer connection.Close()
	tokenServiceClient = pb.NewTokenServiceClient(connection)
	ctx := context.Background()

	findLeader(ctx)
	sendHeartbeatToAll()
}

func sendHeartbeatToAll() {
	clientMap = make(map[int]client)

	log.Println("Starting heartbeat check...")
	for _, port := range serverPorts {
		if port != int(listenPort) { //Make sure you dont target yourself
			go sendHeartbeatTo(port)
		}
	}
}

func sendHeartbeatTo(port int) {
	log.Println("Send heartbeat check to", port)
	address := "localhost:" + strconv.Itoa(port)
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer connection.Close()
	tokenServiceClient := pb.NewTokenServiceClient(connection)
	cli := client{
		service: tokenServiceClient,
		alive:   true,
	}
	mux.Lock()
	clientMap[port] = cli
	mux.Unlock()

	for {
		response, err := cli.service.SendHeartbeat(context.Background(), &pb.Empty{})
		if err != nil {
			log.Println("Failed to connect to", port, "likely dead!")
			cli.alive = false
			mux.Lock()
			clientMap[port] = cli
			mux.Unlock()

			if port == int(currentLeader) || currentLeader == -1 { //Is the leader dead?
				aliveCheck()
			}
		} else {
			log.Println("Reponse:", response)
			cli.alive = true
			mux.Lock()
			clientMap[port] = cli
			mux.Unlock()
		}
		time.Sleep(10 * time.Second)
	}
}

func aliveCheck() { //Make sure that enough nodes are alive
	alive := 1 //Start at 1 because node includes itself as alive
	for _, cli := range clientMap {
		if cli.alive {
			alive++
		}
	}
	log.Println("Checking if new leader should be found:", alive, requiredAlive)
	if alive >= requiredAlive {
		_currentLeader := currentLeader
		currentLeader = -1
		lowestPort = listenPort //Reset
		if _currentLeader == targetPort {

			pos := 0
			for _, port := range serverPorts {
				pos++
				if port == int(_currentLeader) {
					pos %= len(serverPorts)
					break
				}
			}
			log.Println("Pos:", pos)

			nextPort := int32(serverPorts[pos])
			log.Println("Found next alive port:", nextPort)
			runClient(nextPort)
		} else {
			runClient(targetPort)
		}
	} else {
		log.Println("Can't connect to other servers. I must be the problem.")
	}
}

func findLeader(ctx context.Context) {
	for {
		if currentLeader == -1 {
			sendLeaderRequest(ctx)
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	sendLeaderRequest(ctx)
	time.Sleep(5 * time.Second)
}

func sendLeaderRequest(ctx context.Context) {
	log.Println("Sending LeaderRequest from", listenPort, "to", targetPort, "currentLeader", currentLeader, "lowestPort", lowestPort)
	tokenServiceClient.FindLeaderRequest(ctx, &pb.LeaderRequest{CurrentLeader: currentLeader, LowestPort: lowestPort})
}
func runServer() {
	log.Println("--- SERVER APP ---")

	address := "localhost:" + strconv.Itoa(int(listenPort))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterTokenServiceServer(s, &tokenServer{})
	pb.RegisterAuctionServiceServer(s, &auctionServer{})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
