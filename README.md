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


	//build the template
	tmpl := template.New("templateName", template.EN_US).
		AddHeader("header").
		AddBody("body").
		AddBody("body").
		AddBody("body").
		Done()

	recipient := "2349045057268"

	// create a new client
	client := whatsapp.New("phoneNumberID", "appAccessToken", whatsapp.WithApiVersion(whatsapp.V15))
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

### Template Builder [example]

Whatsapp cloud api allows businesses to send messages to their customers using pre-defined templates, 
in order to use the template, one would have to send the request in a particular data structure.

<small> The whatsapp/template package in this library simplify's the process of building/contructing that request payload.</small>

<small> Below is a step by step description of the example above.</small>


1. <b>Create a new template</b> using the `template.New()` function, providing a name and language for the template.
```go
tmpl := template.New("templateName", template.EN_US)
```

2. <b>Add a header component</b> to the template using the `AddHeader()` method, specifying the header text as an argument.
```go
tmpl = tmpl.AddHeader("Daniel")
```

3. <b>Add body components</b>  to the template using the `AddBody()` method, providing the body texts as arguments.
```go
tmpl = tmpl.AddBody("Daniel")
tmpl = tmpl.AddBody("4523")
tmpl = tmpl.AddBody("30")
```

4. <b>Finalize the template construction</b> by calling the `Done()` method. It returns the constructed data structure.
```go
constructedTemplate := tmpl.Done()
```
   
   
## Data Structure.  

The resulting data structure:

```json
[
    {
        "type": "header",
        "parameters": [
            {
                "type": "text",
                "text": "Daniel"
            }
        ]
    },
    {
        "type": "body",
        "parameters": [
            {
                "type": "text",
                "text": "Daniel"
            },
            {
                "type": "text",
                "text": "4523"
            },
            {
                "type": "text",
                "text": "30"
            }
        ]
    }
]
```


