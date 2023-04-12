# go-whatsapp


[![GoDoc](https://godoc.org/github.com/iamNator/go-whatsapp?status.svg)](https://godoc.org/github.com/iamNator/go-whatsapp)
[![Go Report Card](https://goreportcard.com/badge/github.com/iamNator/go-whatsapp)](https://goreportcard.com/report/github.com/iamNator/go-whatsapp)
[![Build Status](https://travis-ci.org/iamNator/go-whatsapp.svg?branch=main)](https://travis-ci.org/iamNator/go-whatsapp)
[![Coverage Status](https://coveralls.io/repos/github/iamNator/go-whatsapp/badge.svg?branch=main)](https://coveralls.io/github/iamNator/go-whatsapp?branch=main)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)]



A Library for Sending out messages via templates on whatsapp business cloud API.



````
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

````
