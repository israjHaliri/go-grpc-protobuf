package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	"github.com/israjHaliri/go-grpc-protobuf/common/config"
	"github.com/israjHaliri/go-grpc-protobuf/common/model"
)

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	emailAndPassword := model.EmailAndPassword{
		Email:       "israj.haliri@gmail.com",
		Password: "PasswordTralala",
	}

	service := serviceUser()

	userList, err := service.FindAllUsers(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}

	userListString, _ := json.Marshal(userList.List)
	log.Println(string(userListString))

	user, err := service.FindUserByEmailAndPAssword(context.Background(), &emailAndPassword)

	if err != nil {
		log.Fatal(err.Error())
	}

	userString, _ := json.Marshal(user)
	log.Println(string(userString))

}
