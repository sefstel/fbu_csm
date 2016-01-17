package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
func readCSV(filename string) *[]string {
	bs, err := ioutil.ReadFile(filename)
	check(err)
	str := strings.Split(string(bs), "\n")
	return &str
}
