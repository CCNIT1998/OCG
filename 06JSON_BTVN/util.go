package main

import (
	"fmt"
	// "sort"
)

type KeyValue struct {
	Key   string
	Value int
}

func PrintSliceKeyValue(input []KeyValue) {
	for index, item := range input {
		fmt.Printf("%d - %s - %d \n", index+1, item.Key, item.Value)
	}
}

func jobWorkMaxInEachCity(input map[string]int) (result map[string]int) {
	result = make(map[string]int)
	max := 0
	for _, value := range input {
		if value > max {
			max = value
		}
	}
	for key, value := range input {
		if value == max {
			result[key] = value
		}
	}
	return result
}
