package main

import (
	"fmt"
	"strings"
)

type Request struct {
	method   string
	resource string
	params   string
}

func ParseFirstLine(body string) {

	req := strings.Split(body, "\r\n")

	for _, line := range req {
		feats := strings.Split(line, " ")

		for _, feat := range feats {
			fmt.Println(feat)
		}

	}
}
