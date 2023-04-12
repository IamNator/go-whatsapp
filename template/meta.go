package template

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

//https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages/#template-messages

type (
	metaPayload struct {
		MessagingProduct string               `json:"messaging_product"`
		To               string               `json:"to"`
		Type             PayloadType          `json:"type"`
		Text             *Text                `json:"text,omitempty"`
		Template         *MetaPayloadTemplate `json:"template"`
	}

	PayloadType string

	Text struct {
		Body string `json:"body"`
	}

	MetaPayloadTemplate struct {
		Name       string      `json:"name"`
		Language   *Language   `json:"language"`
		Components []Component `json:"components"`
	}

	Language struct {
		Code string `json:"code"`
	}

	Component struct {
		Type       ComponentType        `json:"type"`
		Parameters []ParameterInterface `json:"parameters"`

		SubType SubType `json:"sub_type,omitempty"` // e.g quick_reply, url etc
		Index   string  `json:"index,omitempty"`    // e.g 0, 1, 2, 3 etc
	}

	Parameter struct {
		Type ParameterType `json:"type"`
		Text string        `json:"text"`
	}

	ImageParameter struct {
		Type  ParameterType `json:"type"`
		Image ImageLink     `json:"image"`
	}

	ImageLink struct {
		Link string `json:"link"`
	}

	//ButtonPayloadParameter ...
	// for component_type = button
	//
	//
	// {
	// 	"type": "payload",
	// 	"payload": "<payload>"
	// }
	ButtonPayloadParameter struct {
		Type    ParameterType `json:"type"`
		Payload string        `json:"payload"`
	}

	ComponentType string

	SubType string // e.g quick_reply, url etc

	// subtypes of ComponentType e.g text, image, quick_reply, url etc
	ParameterType string

	LangCode string
)

const (
	USLangCode LangCode = "en_US"
	UKLangCode LangCode = "en_GB"
)

func (l LangCode) String() string {
	return string(l)
}

const (
	TypeText     PayloadType = "text"
	TypeTemplate PayloadType = "template"

	ComponentTypeHeader ComponentType = "header"
	ComponentTypeBody   ComponentType = "body"
	ComponentTypeButton ComponentType = "button"

	SubTypeQuickReply SubType = "quick_reply"
	SubTypeTypeUrl    SubType = "url"
	SubTypeNone       SubType = ""

	ParameterTypeText  ParameterType = "text"
	ParameterTypeImage ParameterType = "image" //not sure

	//for button
	ParameterTypeButtonPayload ParameterType = "payload"
)

type (
	ParameterInterface interface {
		Build(value string)
	}
)

func (p *Parameter) Build(text string) {
	p.Type = ParameterTypeText
	p.Text = text
}

func (p *ImageParameter) Build(imageLink string) {
	p.Type = ParameterTypeImage
	p.Image.Link = imageLink
}

func (p *ButtonPayloadParameter) Build(payload string) {
	p.Type = ParameterTypeButtonPayload
	p.Payload = payload
}

// BuildParameter returns a ParameterInterface based on the parameterType
func BuildParameter(parameterType ParameterType, value string) ParameterInterface {
	switch parameterType {
	case ParameterTypeText:
		p := new(Parameter)
		p.Build(value)
		return p
	case ParameterTypeImage:
		p := new(ImageParameter)
		p.Build(value)
		return p
	case ParameterTypeButtonPayload:
		p := new(ButtonPayloadParameter)
		p.Build(value)
		return p

	}

	panic("please help me")
}

func New(templateName, to, langCode string) *metaPayload {
	to = strings.Trim(to, "+")
	return &metaPayload{
		MessagingProduct: "whatsapp",
		To:               strings.ReplaceAll(to, "+", ""),
		Type:             TypeTemplate,
		Template: &MetaPayloadTemplate{
			Name: templateName,
			Language: &Language{
				Code: langCode,
			},
		},
	}
}

// CleanText removes all double spaces, new lines and tabs from a string
//
// according to whatsapp specifications.
func CleanText(s string) string {

	space := regexp.MustCompile(`\s+`)
	s = space.ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")

	return s
}
func (m *metaPayload) AddHeader(text string) *metaPayload {
	return m.addComponent(ComponentTypeHeader, SubTypeNone, ParameterTypeText, CleanText(text))
}

func (m *metaPayload) AddHeaderImage(imageLink string) *metaPayload {
	return m.addComponent(ComponentTypeHeader, SubTypeNone, ParameterTypeImage, CleanText(imageLink))
}

func (m *metaPayload) AddImage(imageLink string) *metaPayload {
	return m.addComponent(ComponentTypeHeader, SubTypeNone, ParameterTypeImage, CleanText(imageLink))
}

func (m *metaPayload) AddBody(text string) *metaPayload {
	return m.addComponent(ComponentTypeBody, SubTypeNone, ParameterTypeText, CleanText(text))
}

func (m *metaPayload) AddButton(text string) *metaPayload {
	return m.addComponent(ComponentTypeButton, SubTypeNone, ParameterTypeText, CleanText(text))
}

func (m *metaPayload) AddButtonURL(link string) *metaPayload {
	return m.addComponent(ComponentTypeButton, SubTypeTypeUrl, ParameterTypeText, CleanText(link))
}

func (m *metaPayload) AddButtonQuickReply(text string) *metaPayload {
	return m.addComponent(ComponentTypeButton, SubTypeQuickReply, ParameterTypeButtonPayload, CleanText(text))
}

func (m *metaPayload) Byte() ([]byte, error) {
	return json.Marshal(m)
}

func NewFromBytes(b []byte) (*metaPayload, error) {
	var m metaPayload
	if er := json.Unmarshal(b, &m); er != nil {
		return nil, er
	}

	return &m, nil
}

// addComponent ...
//
//	takes in a component_type, parameter_type and value
//	and creates an array of components and parameters
//
// [
//
//	{
//	  "type": "<component_type>",
//	  "parameters": [
//		{
//		  "type": "<parameter_type>",
//		  "text": "<text>"
//		}
//	  ]
//	}
//	{
//	  "type": "<component_type>",
//	  "parameters": [
//		{
//		  "type": "<parameter_type>",
//		  "text": "<text>"
//		}
//	  ]
//	}
//
// ]
func (m *metaPayload) addComponent(componentType ComponentType, subType SubType, parameterType ParameterType, value string) *metaPayload {

	// 1.  if component does not exist, create it
	// 2.  append the parameter to the component if it exists
	// 3.  if component does not exist, create it and append the parameter to it

	// if the component does not exist, it creates the component
	//
	// {
	//	"type": "<adds component_type>",  //creates the component e.g header, body, button
	//  "sub_type": "<adds sub_type>",   // e.g quick_reply, url
	//  "index": "0",                    //
	//	"parameters": [{}]
	// }
	//
	if len(m.Template.Components) < 1 {
		m.Template.Components = make([]Component, 0)
		component := Component{
			Type:    componentType,
			SubType: subType,
		}

		if subType == SubTypeQuickReply || subType == SubTypeTypeUrl {
			component.Index = "0"
		}

		m.Template.Components = append(m.Template.Components, component)
	}

	var exist bool //this is to check if the component exist

	//  if the component exist, it appends the parameter to the component
	//
	//	{
	//	  "type": "<finds component_type>",
	//    "sub_type": "<adds sub_type>",   // e.g quick_reply, url
	//    "index": "<add index>",                    // e.g 0, 1, 2, 3
	//	  "parameters": [
	//		{
	//		  "type": "<parameter_type>",
	//		  "text": "<text>"
	//		}
	// 	   ...
	//      {  //appends the parameter to the component
	//		  "type": "<adds parameter_type>",
	//		  "text": "<adds text>"
	//		}
	//	  ]
	//	}
	for index, component := range m.Template.Components { // this is to check if the component exist, it iterates through the components to find the component
		if component.Type == componentType { //if the component exists

			if component.Type == ComponentTypeButton { //if the component is a button, it creates a new component with an index = (index + 1)
				newComponent := Component{
					Type:    componentType,
					SubType: subType,
				}
				num, _ := strconv.Atoi(component.Index)
				num++
				newComponent.Index = strconv.Itoa(num)
				newComponent.Parameters = append(newComponent.Parameters, BuildParameter(parameterType, value))
			}

			exist = true //set exist to true
			m.Template.Components[index].Parameters = append(m.Template.Components[index].Parameters,
				BuildParameter(parameterType, value)) //append the parameter to the component
		}
	}

	//   if the component does not exist, it creates the component and adds the parameter to it
	//	{
	//	  "type": "<creates component_type>",
	//    "sub_type": "<adds sub_type>",
	//	  "parameters": [  //creates first parameter
	//		{
	//		  "type": "<adds parameter_type>",
	//		  "text": "<adds text>"
	//		}
	//	  ]
	//	}
	if !exist { //if the component does not exist
		component := Component{
			Type:    componentType, //create the component
			SubType: subType,
			Parameters: []ParameterInterface{
				BuildParameter(parameterType, value), //add the first parameter to the component
			},
		}

		if subType == SubTypeQuickReply || subType == SubTypeTypeUrl {
			component.Index = "0"
		}

		m.Template.Components = append(m.Template.Components, component)
	}

	return m
}

func FromByteToMap(b []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{}, 0)
	if er := json.Unmarshal(b, &m); er != nil {
		return nil, er
	}

	return m, nil
}
