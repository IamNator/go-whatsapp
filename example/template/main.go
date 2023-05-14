/*
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
*/

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
