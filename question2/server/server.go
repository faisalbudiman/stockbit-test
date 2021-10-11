package main

import (
	"context"
	"log"
	"net"
	"net/http"
	pb "omdb-api/gen/proto"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type searchAPIServer struct {
	pb.UnimplementedSearchAPIServer
}

func (s *searchAPIServer) Search(ctx context.Context, param *pb.SearchRequest) (*pb.ResponseSearchAPI, error) {
	omdbKey := os.Getenv("OMDBKEY")
	req, err := http.NewRequest("GET", "http://www.omdbapi.com", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	keyWord := strings.ReplaceAll(param.Keyword, " ", "+")
	// if you appending to existing query this works fine
	q := req.URL.Query()
	q.Add("apikey", omdbKey)
	q.Add("page", param.Page)
	q.Add("s", keyWord)

	req.URL.RawQuery = q.Encode()

	res := pb.ResponseSearchAPI{
		
	}

	return &res, nil
}

func Init() {
	godotenv.Load(".env")
}

func main() {
	// mux listener for rest
	go func() {
		mux := runtime.NewServeMux()

		pb.RegisterSearchAPIHandlerServer(context.Background(), mux, &searchAPIServer{})

		log.Fatalln(http.ListenAndServe("localhost:8081", mux))
	}()

	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterSearchAPIServer(grpcServer, &searchAPIServer{})

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Listening on localhost:8080")
}
