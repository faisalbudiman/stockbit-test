package main

import (
	"context"
	"fmt"
	"log"
	pb "omdb-api/gen/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := pb.NewSearchAPIClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{Page: "1"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp)
	fmt.Println(resp.Body)
}
