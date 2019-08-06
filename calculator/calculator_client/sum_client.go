package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, this is a gRPC client")

	client_connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer client_connection.Close()

	client := calculatorpb.NewCalculatorServiceClient(client_connection)

	request := &calculatorpb.SumRequest{
		FirstNumber:  2,
		SecondNumber: 3,
	}

	response, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("error while calling Sum() gRPC: %v", err)
	}
	//	log.Printf("2 + 3 = %v", response.Result)
	fmt.Printf("2 + 3 = %d", response.Result)
}
