package main

import (
	"fmt"
	"log"
)

func Print(v ...interface{}) {
	for _, val := range v {
		fmt.Printf("%v , %T \t", val, val)
	}
	fmt.Printf("\n")
}

func Validate(item string, err string) {
	if item == "" {
		log.Panicln(err)
	}
}
