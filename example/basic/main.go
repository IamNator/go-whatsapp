package main

import (
	"context"
	"fmt"

	whatsapp "github.com/iamNator/go-whatsapp"
	"github.com/iamNator/go-whatsapp/template"
)

func main() {
	client := whatsapp.New("metaAppId", "metaAppAccessToken", whatsapp.V15)

	// prepare the payload
	data := whatsapp.NewPayload("templateName", "+2349045057268", template.EN_US)

	//build the template
	data.Template.AddHeader("header").
		AddBody("body").
		AddBody("body").
		AddBody("body")

	// send the request
	response, errResponse, er := client.Send(context.Background(), *data)
	if er != nil {
		fmt.Println("error: ", er.Error())
		return
	}

	fmt.Println("Err: ", errResponse, "\nResponse: ", response)
}
