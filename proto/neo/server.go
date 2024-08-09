package neo

import (
	"context"
	"google.golang.org/grpc/reflection"

	"net"

	"google.golang.org/grpc"

	neo "github.com/D0K-ich/JinJi-Service/proto/neo/genned"
)

type serverAPI struct {
	neo.UnimplementedAuthServer // Хитрая штука, о ней ниже
	neo Auth
}

type Auth interface {
	Login(ctx context.Context, email string, password string, appID int) (token string, err error)
	RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error)
	//neo.UnimplementedAuthServer
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	neo.RegisterAuthServer(gRPCServer, &serverAPI{neo: auth})
}

func (s *serverAPI) Login(ctx context.Context, in *neo.LoginRequest) (*neo.LoginResponse, error) {
	// TODO
	return &neo.LoginResponse{Token: "token"}, nil
}
//
//func (s *serverAPI) Register(ctx context.Context, in *neo.RegisterRequest) (*neo.RegisterResponse, error) {
//	// TODO
//	return &neo.RegisterResponse{}, nil
//}

//type server struct {
//	auth neo.AuthServer
//	neo.UnimplementedAuthServer
//}

func StartgRPC() (err error) {
	var listener net.Listener
	if listener, err = net.Listen("tcp", ":11223"); err != nil {return}

	var s = grpc.NewServer()
	neo.RegisterAuthServer(s, &serverAPI{})

	reflection.Register(s)

	go func() {
		if err = s.Serve(listener); err != nil {return}
	}()

	return
}
