package main

//go:generate oapi-codegen -package=main -generate=client,types -o ./americancloud.gen.go https://app.americancloud.com/docs/api-docs.json

import (
	"context"
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

func WithDebug(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		return nil
	}
}

func main() {
	bearerTokenProvider, bearerTokenProviderErr := securityprovider.NewSecurityProviderBearerToken(os.Getenv("AC_TOKEN"))
	if bearerTokenProviderErr != nil {
		panic(bearerTokenProviderErr)
	}

	c, err := NewClientWithResponses("https://app.americancloud.com/api",
		WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		panic(err)
	}

	resp, err := c.ListProjectsWithResponse(context.Background())
	if err != nil {
		panic(err)
	}
	if !(resp.StatusCode() >= 200 && resp.StatusCode() < 300) {
		panic(resp.Status())
	}

	fmt.Println("Projects:")
	projects := *resp.JSON200.Data
	for _, project := range projects {
		fmt.Printf("Id: %v\n", *project.Id)
		fmt.Printf("AccountID: %v\n", *project.AccountId)
		fmt.Printf("Description: %v\n", *project.Description)
		fmt.Printf("Created At: %v\n", *project.CreatedAt)
	}
	fmt.Println("")

	resp2, err := c.InstancesListWithResponse(context.Background(), *projects[0].Name)
	if err != nil {
		panic(err)
	}
	if !(resp2.StatusCode() >= 200 && resp2.StatusCode() < 300) {
		panic(resp2.Status())
	}

	fmt.Println("Instances:")
	instances := *resp2.JSON200.Data.Data
	for _, instance := range instances {
		fmt.Printf("Id: %v\n", *instance.Id)
		fmt.Printf("Name: %v\n", *instance.Name)
		fmt.Printf("Created At: %#v\n", *instance.CreatedAt)
	}
	fmt.Println("")
}
