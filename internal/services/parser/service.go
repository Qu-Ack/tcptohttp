package parser

import (
	"fmt"
	"strings"
)

type config struct {
	protocol string
}

type parser struct {
	config config
}

type ParserInterface interface {
	ParseRequest(request []byte) *Request
	PrintRequest(r *Request)
}

func NewParser(protocol string) ParserInterface {
	var config config
	config.protocol = protocol
	return &parser{
		config: config,
	}
}

func (p *parser) ParseRequest(request []byte) *Request {
	lines := strings.Split(string(request), "\r\n")

	request_line := lines[0]

	headers := make([]string, 0)

	HeaderMap := make(map[HEADER]string)

	i := 1
	for i = 1; i < len(lines)-1; i++ {
		headers = append(headers, lines[i])
	}

	for _, header := range headers {
		arr := strings.Split(header, ": ")
		if len(arr) == 2 {
			field := arr[0]
			value := arr[1]
			HeaderMap[HEADER(field)] = value
		}
	}

	body := []byte(lines[i])

	requestLineFeatures := strings.Split(request_line, " ")

	method := requestLineFeatures[0]
	URI := requestLineFeatures[1]
	version := requestLineFeatures[2]

	return &Request{
		Method:          METHOD(method),
		URI:             URI,
		ProtocolVersion: version,

		Headers: HeaderMap,
		Body:    body,
	}
}
func (p *parser) PrintRequest(r *Request) {
	fmt.Printf("Method is %s\n", r.Method)
	fmt.Printf("Version is %s\n", r.ProtocolVersion)
	fmt.Printf("Requested Resource %s\n", r.URI)

	fmt.Println("Headers are { ")
	for key, value := range r.Headers {
		fmt.Printf("\t%s : %s\n", key, value)
	}

	fmt.Println("}")
	fmt.Print("Body: ")
	fmt.Println(string(r.Body))
}
