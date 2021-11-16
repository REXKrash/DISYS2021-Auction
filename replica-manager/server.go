package main

import (
	"context"
	"strconv"
	"time"

	"log"
	"net"
	"os"

	pb "auction/routeguide"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTokenServiceServer
}

type client struct {
	service pb.TokenServiceClient
	alive   bool
}

func (s *server) SendHeartbeat(ctx context.Context, request *pb.Empty) (*pb.HeartbeatResponse, error) {
	return &pb.HeartbeatResponse{Status: 1}, nil
}

func (s *server) FindLeaderRequest(ctx context.Context, request *pb.LeaderRequest) (*pb.Empty, error) {
	log.Println("Received LeaderRequest to", listenPort, "currentLeader:", request.CurrentLeader, "lowestPort:", request.LowestPort)
	currentLeader = request.CurrentLeader

	if request.LowestPort <= lowestPort && currentLeader == -1 {
		lowestPort = request.LowestPort
		if lowestPort == listenPort {
			currentLeader = listenPort
			log.Println("Found leader, it is me!")
		}
	}
	return &pb.Empty{}, nil
}

var requiredAlive = 2
var targetPort int32
var listenPort int32
var lowestPort int32
var currentLeader int32 = -1
var tokenServiceClient pb.TokenServiceClient
var serverPorts = [3]int{5001, 5002, 5003}
var serverID int32
var clientMap map[int]client

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
	if _serverID == len(serverPorts)-1 {
		targetPort = int32(serverPorts[0])
	} else {
		targetPort = int32(serverPorts[_serverID+1])
	}
	lowestPort = listenPort
	serverID = int32(_serverID)

	time.Sleep(10 * time.Second)
	go runServer()
	time.Sleep(1 * time.Second)
	go runClient(targetPort)
	for {
		time.Sleep(2 * time.Second)
	}
}

func runClient(target int32) {
	targetPort = target
	address := "localhost:" + strconv.Itoa(int(targetPort))
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer connection.Close()
	tokenServiceClient = pb.NewTokenServiceClient(connection)
	ctx := context.Background()

	findLeader(ctx)
	sendHeartbeat()
}

func sendHeartbeat() {
	clientMap = make(map[int]client)

	log.Println("Starting heartbeat check...")
	for _, port := range serverPorts {
		if port != int(listenPort) { //Make sure you dont target yourself

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
			clientMap[port] = cli

			for {
				response, err := cli.service.SendHeartbeat(context.Background(), &pb.Empty{})
				if err != nil {
					log.Println("Failed to connect to", port, "likely dead!")
					cli.alive = false
					clientMap[port] = cli

					if port == int(currentLeader) { //Is the leader dead?
						aliveCheck()
					}

				} else {
					log.Println("Reponse:", response)
					cli.alive = true
					clientMap[port] = cli
				}
				time.Sleep(10 * time.Second)
			}
		}
	}
}

func aliveCheck() { //Make sure that enough nodes are alive
	alive := 1 //Start
	for _, cli := range clientMap {
		if cli.alive {
			alive++
		}
	}
	if alive >= requiredAlive {
		if currentLeader == targetPort {
			currentLeader = -1

		}
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

	pb.RegisterTokenServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
