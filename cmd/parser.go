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

	for i, line := range req {
		fmt.Println("%d line", i+1)
		feats := strings.Split(line, " ")

		for _, feat := range feats {
			fmt.Printf("%s feat", feat)
		}
		fmt.Println("")
	}
}
