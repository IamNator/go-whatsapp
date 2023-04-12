//package main
//
//import (
//	"context"
//	"fmt"
//
//	whatsapp "github.com/iamNator/go-whatsapp"
//	"github.com/iamNator/go-whatsapp/template"
//)
//
//func main() {
//	client := whatsapp.New("metaAppId", "metaAppAccessToken", whatsapp.V15)
//
//	// prepare the payload
//
//	//build the template
//	tmpl := template.New("templateName", template.EN_US).AddHeader("header").
//		AddBody("body").
//		AddBody("body").
//		AddBody("body").
//		Done()
//
//	// send the request
//	response, errResponse, er := client.SendTemplate(context.Background(), "2349045057268", tmpl)
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
	"github.com/iamNator/go-whatsapp/template"
)

func main() {
	client := whatsapp.New("metaAppId", "metaAppAccessToken", whatsapp.V15)

	// prepare the payload

	//build the template
	tmpl := template.New("templateName", template.EN_US).AddHeader("header").
		AddBody("body").
		AddBody("body").
		AddBody("body").
		Done()

	// send the request
	response, errResponse, er := client.SendTemplate(context.Background(), "2349045057268", tmpl)
	if er != nil {
		fmt.Println("error: ", er.Error())
		return
	}

	fmt.Println("Err: ", errResponse, "\nResponse: ", response)
}
