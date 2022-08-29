package meta

import (
	"encoding/json"
	"strings"
)

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

	ComponentType string

	ParameterType string
) //

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
	}

	panic("please help me")
}

const (
	TypeText     PayloadType = "text"
	TypeTemplate PayloadType = "template"

	ComponentTypeHeader ComponentType = "header"
	ComponentTypeBody   ComponentType = "body"
	ComponentTypeButton ComponentType = "button"
	ComponentTypeFooter ComponentType = "footer"

	ParameterTypeText  ParameterType = "text"
	ParameterTypeImage ParameterType = "image" //not sure
)

func New(templateName, to, langCode string) *metaPayload {
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

func NewFromBytes(b []byte) (*metaPayload, error) {

	m := new(metaPayload)
	if er := json.Unmarshal(b, m); er != nil {
		return nil, er
	}

	return m, nil
}

func FromByteToMap(b []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{}, 0)
	if er := json.Unmarshal(b, &m); er != nil {
		return nil, er
	}

	return m, nil
}

func (m *metaPayload) AddHeader(text string) *metaPayload {
	return m.addComponent(ComponentTypeHeader, ParameterTypeText, text)
}

func (m *metaPayload) AddImage(imageLink string) *metaPayload {
	return m.addComponent(ComponentTypeHeader, ParameterTypeImage, imageLink)
}

func (m *metaPayload) AddBody(text string) *metaPayload {
	return m.addComponent(ComponentTypeBody, ParameterTypeText, text)
}

func (m *metaPayload) AddButton(text string) *metaPayload {
	return m.addComponent(ComponentTypeButton, ParameterTypeText, text)
}

func (m *metaPayload) AddFooter(text string) *metaPayload {
	return m.addComponent(ComponentTypeFooter, ParameterTypeText, text)
}

func (m *metaPayload) Byte() ([]byte, error) {
	return json.Marshal(m)
}

//addComponenet ...
func (m *metaPayload) addComponent(componentType ComponentType, parameterType ParameterType, value string) *metaPayload {
	if len(m.Template.Components) < 1 {
		m.Template.Components = make([]Component, 0)
		m.Template.Components = append(m.Template.Components, Component{
			Type: componentType,
		})
	}

	var exist bool
	for index, component := range m.Template.Components { //runs true just once
		if component.Type == componentType {
			exist = true
			m.Template.Components[index].Parameters = append(m.Template.Components[index].Parameters,
				BuildParameter(parameterType, value))
		}
	}

	if !exist { //is the component does not exist
		m.Template.Components = append(m.Template.Components, Component{
			Type: componentType,
			Parameters: []ParameterInterface{
				BuildParameter(parameterType, value),
			},
		})
	}

	return m
}
