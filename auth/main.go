package main

import (
	"context"
	"fmt"

	proto "github.com/micro/examples/helloworld/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/auth"
	"github.com/micro/go-micro/auth/service"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	a := service.NewAuth()

	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Auth(a),
	)
	service.Init()

	// Generate a service account
	sa, err := a.Generate(&auth.Account{
		Parent: &auth.Resource{
			Type: "user", Name: "asim@micro.mu",
		},
		Roles: []*auth.Role{
			&auth.Role{
				Name: "admin",
			},
			&auth.Role{
				Name: "contributor",
				Resource: &auth.Resource{
					Type: "service",
					Name: "go.micro.srv.foobar",
				},
			},
		},
		Metadata: map[string]string{
			"email":  "asim@micro.mu",
			"github": "asim",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("✅ Generated a service account")

	// Validate the service account
	sa, err = a.Validate(sa.Token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("✅ Validated the service account")

	// Revoke the service account
	err = a.Revoke(sa.Token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("✅ Revoked the service account")

}
