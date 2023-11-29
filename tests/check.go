package main

import (
	"context"
	"fmt"

	"github.com/palp1tate/brevinect/proto/user"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := userProto.NewUserServiceClient(cc)

	res, err := client.CheckMobile(context.Background(), &userProto.CheckMobileRequest{
		Mobile: "17383835083",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Exist)
}
