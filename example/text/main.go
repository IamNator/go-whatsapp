//
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
//	if errResponse != nil {
//		fmt.Println("ErrorResponse: ", errResponse.Message)
//		return
//	}
//
//	fmt.Println("Response: ", response)
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

	if errResponse != nil {
		fmt.Println("ErrorResponse: ", errResponse.Message)
		return
	}

	fmt.Println("Response: ", response)
}
