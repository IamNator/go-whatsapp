package main

import (
	"context"
	"fmt"
	whatsapp "github.com/iamNator/go-whatsapp"
	"github.com/iamNator/go-whatsapp/meta"
)

func main() {
	whatsApp := whatsapp.New("metaAppId", "metaAppAccessToken")

	data := template.New("templateName", "+2349045057268", "en_US")
	data.AddHeader("header").
		AddBody("body").
		AddBody("body").
		AddBody("body")

	obj, er := data.Byte()
	if er != nil {
		fmt.Printf("error: %s", er.Error())
		return
	}

	response, er := whatsApp.Send(context.Background(), whatsapp.Message{Data: obj})
	if er != nil {
		fmt.Println("error: ", er.Error())
		return
	}

	fmt.Println(response)
}
