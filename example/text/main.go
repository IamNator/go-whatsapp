//package main
//
//import (
//	"context"
//	"fmt"
//
//	whatsapp "github.com/iamNator/go-whatsapp"
//)
//
//func main() {
//	client := whatsapp.New("metaAppId", "metaAppAccessToken", whatsapp.V15)
//
//	// send the request
//	response, errResponse, er := client.SendText(context.Background(), "2349045057268", "Hello World")
//	if er != nil {
//		fmt.Println("error: ", er.Error())
//		return
//	}
//
//	fmt.Println("Err: ", errResponse, "\nResponse: ", response)
//}

package main

import (
	"context"
	"fmt"

	whatsapp "github.com/iamNator/go-whatsapp"
)

func main() {
	client := whatsapp.New("metaAppId", "metaAppAccessToken", whatsapp.V15)

	// send the request
	response, errResponse, er := client.SendText(context.Background(), "2349045057268", "Hello World")
	if er != nil {
		fmt.Println("error: ", er.Error())
		return
	}

	fmt.Println("Err: ", errResponse, "\nResponse: ", response)
}
