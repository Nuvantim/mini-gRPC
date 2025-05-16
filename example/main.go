package main 

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	greetv1 "example/greet/v1"
	"example/gen/greet/v1/greetv1connect"
)

type GreetServer struct{}

func (s *GreetServer ) Greet(ctx context.Context,req *connect.Request[greetv1.GreetRequest],) (*connect.Response[greetv1.GreetResponse], error) {
	
}

func main(){

}