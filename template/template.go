package template

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

//https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages/#template-messages

type (
	// subtypes of ComponentType e.g text, image, quick_reply, url etc
	ParameterType string

	Parameter struct {
		Type ParameterType `json:"type"`
		Text string        `json:"text"`
	}

	ImageLink struct {
		Link string `json:"link"`
	}

	ImageParameter struct {
		Type  ParameterType `json:"type"`
		Image ImageLink     `json:"image"`
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

	ComponentType string // e.g header, body, button etc

	SubType string // e.g quick_reply, url etc

	// Component ...
	Component struct {
		Type       ComponentType        `json:"type"` // e.g header, body, button etc
		Parameters []ParameterInterface `json:"parameters"`
		SubType    SubType              `json:"sub_type,omitempty"` // e.g quick_reply, url etc
		Index      string               `json:"index,omitempty"`    // e.g 0, 1, 2, 3 etc
	}

	// Language ...
	Language struct {
		Code string `json:"code"` // e.g en_US, en_GB
	}

	// Template ...
	Template struct {
		Name       string      `json:"name"` // The name of the whatsapp cloup api messaging template e.g signup_otp
		Language   *Language   `json:"language"`
		Components []Component `json:"components"`
	}
)

const (
	ComponentTypeHeader ComponentType = "header"
	ComponentTypeBody   ComponentType = "body"
	ComponentTypeButton ComponentType = "button"

	SubTypeQuickReply SubType = "quick_reply"
	SubTypeTypeUrl    SubType = "url"
	SubTypeNone       SubType = "" //do not delete

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

	panic("please help me") // TODO: return error
}

func New(templateName string, langCode LanguageCode) *Template {

	return &Template{

		Name: templateName,
		Language: &Language{
			Code: langCode.String(),
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

// AddHeader specifies a header text component for the WhatsApp template content.
// Note: The order of adding components is important as it determines their position within the template.
func (m *Template) AddHeader(text string) *Template {
	return m.addComponent(ComponentTypeHeader, SubTypeNone, ParameterTypeText, CleanText(text))
}

// AddHeaderImage specifies a header image component for the WhatsApp template content.
// Note: The order of adding components is important as it determines their position within the template.
func (m *Template) AddHeaderImage(imageLink string) *Template {
	return m.addComponent(ComponentTypeHeader, SubTypeNone, ParameterTypeImage, CleanText(imageLink))
}

// AddBody specifies a body text component for the WhatsApp template content.
// Note: The order of adding components is important as it determines their position within the template.
func (m *Template) AddBody(text string) *Template {
	return m.addComponent(ComponentTypeBody, SubTypeNone, ParameterTypeText, CleanText(text))
}

// AddButton specifies a button text component for the WhatsApp template content.
// Note: The order of adding components is important as it determines their position within the template.
func (m *Template) AddButton(text string) *Template {
	return m.addComponent(ComponentTypeButton, SubTypeNone, ParameterTypeText, CleanText(text))
}

// AddButtonURL specifies a button URL component for the WhatsApp template content.
// Note: The order of adding components is important as it determines their position within the template.
func (m *Template) AddButtonURL(link string) *Template {
	return m.addComponent(ComponentTypeButton, SubTypeTypeUrl, ParameterTypeText, CleanText(link))
}

// AddButtonQuickReply specifies a button quick reply component for the WhatsApp template content.
// Note: The order of adding components is important as it determines their position within the template.
func (m *Template) AddButtonQuickReply(text string) *Template {
	return m.addComponent(ComponentTypeButton, SubTypeQuickReply, ParameterTypeButtonPayload, CleanText(text))
}

// Done returns the final WhatsApp template content.
func (m Template) Done() Template {
	return m
}

// addComponent ...
//
//	takes in a component_type, parameter_type and value
//	and creates an array of components and parameters
//
//	-->in other words, it adds parameters
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
func (m *Template) addComponent(componentType ComponentType, subType SubType, parameterType ParameterType, value string) *Template {

	// 1.  if there are no components, create one i.e if len(m.Template.Components) < 1
	// 2.  append the parameter to the component if it exists
	// 3.  if component does not exist among the already existing ones, create it and append the parameter to it

	// if there are no components, create one i.e if len(m.Template.Components) < 1
	//
	// {
	//	"type": "<adds component_type>",  //creates the component e.g header, body, button
	//  "sub_type": "<adds sub_type>",   // e.g quick_reply, url
	//  "index": "0",                    //
	//	"parameters": [{}]
	// }
	//
	if len(m.Components) < 1 {
		m.Components = make([]Component, 0)
		component := Component{
			Type:    componentType,
			SubType: subType,
		}

		if subType == SubTypeQuickReply || subType == SubTypeTypeUrl {
			component.Index = "0"
		}

		m.Components = append(m.Components, component)
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
	for index, component := range m.Components { // this is to check if the component exist, it iterates through the components to find the component
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
			m.Components[index].Parameters = append(m.Components[index].Parameters,
				BuildParameter(parameterType, value)) //append the parameter to the component
		}
	}

	//   if component does not exist among the already existing ones, create it and append the parameter to it
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

		m.Components = append(m.Components, component)
	}

	return m
}

func (m Template) Byte() ([]byte, error) {
	return json.Marshal(m)
}

func (m Template) String() (string, error) {
	b, er := m.Byte()
	if er != nil {
		return "", er
	}

	return string(b), nil
}

func FromByteToMap(b []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{}, 0)
	if er := json.Unmarshal(b, &m); er != nil {
		return nil, er
	}

	return m, nil
}
