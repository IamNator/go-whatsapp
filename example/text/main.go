/*
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

	fmt.Println("Response: ", response)
}
*/

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
