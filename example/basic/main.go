package main

import (
	"context"
	"fmt"

	whatsapp "github.com/iamNator/go-whatsapp"
	"github.com/iamNator/go-whatsapp/template"
)

func main() {
	whatsApp := whatsapp.New("metaAppId", "metaAppAccessToken", whatsapp.V15)

	data := template.New("templateName", "+2349045057268", template.EN_US)
	data.AddHeader("header").
		AddBody("body").
		AddBody("body").
		AddBody("body")

	obj, er := data.Byte()
	if er != nil {
		fmt.Printf("error: %s", er.Error())
		return
	}

	response, errResponse, er := whatsApp.Send(context.Background(), whatsapp.Message{Data: obj})
	if er != nil {
		fmt.Println("error: ", er.Error())
		return
	}

	fmt.Println("Err: ", errResponse, "\nResponse: ", response)
}
