package main

import (
	"log"
	"strconv"
)

func joinInts(slice []int) string {
	var s string
	for i, v := range slice {
		s += strconv.Itoa(v)
		if i != len(slice) {
			s += ","
		}
	}
	return s
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
