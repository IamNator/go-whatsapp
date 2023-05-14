# go-whatsapp


[![GoDoc](https://godoc.org/github.com/iamNator/go-whatsapp?status.svg)](https://godoc.org/github.com/iamNator/go-whatsapp)
[![Go Report Card](https://goreportcard.com/badge/github.com/iamNator/go-whatsapp)](https://goreportcard.com/report/github.com/iamNator/go-whatsapp)
[![Coverage Status](https://coveralls.io/repos/github/iamNator/go-whatsapp/badge.svg?branch=main)](https://coveralls.io/github/iamNator/go-whatsapp?branch=main)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)]



This is a Go-Lang library that helps you send messages using pre-made templates on WhatsApp's business cloud platform.



### Send message using pre-made templates
````

package main

import (
	"context"
	"fmt"

	whatsapp "github.com/IamNator/go-whatsapp/v3"
	"github.com/IamNator/go-whatsapp/v3/template"
)

func main() {

	// create a new client
	client := whatsapp.New("phoneNumberID", "appAccessToken", whatsapp.WithApiVersion(whatsapp.V15))

	//build the template
	tmpl := template.New("templateName", template.EN_US).AddHeader("header").
		AddBody("body").
		AddBody("body").
		AddBody("body").
		Done()

	recipient := "2349045057268"

	// send the request
	response, er := client.SendTemplate(context.Background(), recipient, tmpl)
	if er != nil {
		fmt.Println("error: ", er.Error())
		return
	}

	fmt.Println("Response: ", response)
}



````


### Send raw text

```
package main

import (
	"context"
	"fmt"

	whatsapp "github.com/IamNator/go-whatsapp/v3"
)

func main() {
	client := whatsapp.New("phoneNumberID", "appAccessToken", whatsapp.WithApiVersion(whatsapp.V15))

	recipients := "2349045057268"
	text := "Hello World"

	// send the request
	response, er := client.SendText(context.Background(), recipients, text)
	if er != nil {
		fmt.Println("error: ", er.Error())
		return
	}

	// check the response for errors
	if response.Error != nil {
		fmt.Println("error: ", response.Error.Error())
		return
	}

	fmt.Println("Response: ", response)
}


```
