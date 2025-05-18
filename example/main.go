package main

import (
	"connectrpc.com/connect"
	"golang.org/x/net/http2/h2c"
	"net/http"

	greetv1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"
)

type (
	GreetReq  = connect.Request[greetv1.GreetRequest]
	GreetRes  = connect.Response[greetv1.GreetResponse]
)

type GreetServer struct{}

func (s *GreetServer) Greet(_ context.Context, req *GreetReq) (*GreetRes, error) {
	return &GreetRes{
		Msg: &greetv1.GreetResponse{
			Greeting: "Hello, " + req.Msg.Name + "!",
		},
	}, nil
}

func main() {
	mux := http.NewServeMux()
	mux.Handle(greetv1connect.NewGreetServiceHandler(&GreetServer{}))
	http.ListenAndServe(":8080", h2c.NewHandler(mux, new(http2.Server)))
}
