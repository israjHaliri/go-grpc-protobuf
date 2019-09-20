package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/israjHaliri/go-grpc-protobuf/common/config"
	"github.com/israjHaliri/go-grpc-protobuf/common/model"
	"google.golang.org/grpc"
)

type UserServer struct{}

func main() {
	server := grpc.NewServer()

	var userServer UserServer

	model.RegisterUsersServer(server, userServer)

	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	listener, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}

	log.Fatal(server.Serve(listener))
}


func (UserServer) FindAllUsers(ctx context.Context, req *empty.Empty) (*model.UserList, error) {
	dummyUser := model.User{Id: 1, Email: "israj.haliri@gmail.com", Password: "-", Fullname: "israj haliri", Token: "-"}

	userList := []*model.User{}
	userList = append(userList, &dummyUser)

	var list *model.UserList

	list = new(model.UserList)
	list.List = userList

	return list, nil
}

func (UserServer) FindUserByEmailAndPAssword(ctx context.Context, req *model.EmailAndPassword) (*model.User, error) {
	dummyUser := model.User{Id: 1, Email: "israj.haliri@gmail.com", Password: req.GetPassword(), Fullname: "israj haliri", Token: "qwerty"}

	return &dummyUser, nil
}
