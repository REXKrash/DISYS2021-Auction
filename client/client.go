package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "auction/routeguide"
	sh "auction/shared"

	"google.golang.org/grpc"
)

const (
	address = "localhost:5001"
)

var timestamp sh.SafeTimestamp
var serverPorts = [3]int{5001, 5002, 5003}
var currentLeader = -1
var userName string
var sc *bufio.Scanner

func main() {
	sc = bufio.NewScanner(os.Stdin)

	if len(os.Args) > 0 {
		userName = os.Args[1]
	} else {
		log.Println("Please enter your username:")
		sc.Scan()
		userName = sc.Text()
	}

	setCurrentLeader()
	startAuction()
}

func startAuction() {
	address := "localhost:" + strconv.Itoa(currentLeader)
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println(address, "not active -", err)
		return
	}
	defer connection.Close()
	client := pb.NewAuctionServiceClient(connection)
	ctx := context.Background()
	for {
		sc.Scan()
		var msg = sc.Text()

		if len(msg) > 0 && len(msg) <= 128 {
			if strings.HasPrefix(strings.ToLower(msg), "/bid") {
				words := strings.Fields(msg)
				if len(words) >= 2 {
					amount, err := strconv.Atoi(words[1])
					if err != nil {
						log.Println("Failed to parse integer: %v", err)
						return
					}
					response, err := client.Bid(ctx, &pb.BidRequest{Sender: userName, Amount: int32(amount)})

					if err != nil {
						log.Println("Failed to bid: %v", err)

						//Restarting...
						setCurrentLeader()
						startAuction()
						return
					}

					if response.Status == 1 {
						log.Println("You successfully placed a bid with the amount", response.CurrentHighestBid)
					} else if response.Status == 2 {
						log.Println("Bid too low - Current highest bid:", response.CurrentHighestBid)
					} else if response.Status == 3 {
						log.Println("No active auction")
					}

				} else {
					log.Println("Command usage: /bid <amount>")
				}
			} else if strings.HasPrefix(strings.ToLower(msg), "/result") {
				response, err := client.Result(ctx, &pb.ResultRequest{Status: 1})
				if err != nil {
					log.Println("Could not retreive the result")
					setCurrentLeader()
					startAuction()
					return
				} else {
					log.Println("Highest bid is", response.GetHighestBid(), "Winner is", response.GetHighestBidder(), "Is auction running:", response.GetIsAuctionRunning())
				}
			} else {
				log.Println("Unknown command")
			}
		} else {
			log.Println("Message must be between 1-128 characters")
		}
	}
}

func setCurrentLeader() {
	for {
		log.Println("Looking for leader...")
		currentLeader = -1
		for _, port := range serverPorts {

			address := "localhost:" + strconv.Itoa(port)
			connection, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Println("Could not connect :(")
				continue
			}
			defer connection.Close()
			_client := pb.NewAuctionServiceClient(connection)
			ctx := context.Background()

			log.Println("Sending request to find current leader...")

			response, err := _client.AskForLeader(ctx, &pb.AskRequest{Status: 1})
			if err != nil {
				log.Println("Could not connect :(")
				continue
			} else {
				currentLeader = int(response.LeaderPort)
				if currentLeader == -1 {
					continue
				}
				log.Println("Found leader", response.LeaderPort)
				log.Println("Available commands:")
				log.Println("- /bid <amount>")
				log.Println("- /result")
				break
			}
		}
		if currentLeader != -1 {
			break
		}
		time.Sleep(5 * time.Second)
	}
}
