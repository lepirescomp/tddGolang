package main

import "fmt"

type logMessageKey interface {
	~int | ~string
}

func getLogKeys[T logMessageKey, V any](logBody map[T]V) []T {
	var bodyKeys []T

	for k := range logBody {
		bodyKeys = append(bodyKeys, k)
	}

	return bodyKeys
}

func main() {
	logRecord1 := map[string]string{
		"1": "User1",
		"2": "paymentAmount1",
		"3": "data1",
	}
	logRecord2 := map[int]string{
		1: "User2",
		2: "paymentAmount2",
		3: "data2",
	}

	recordKeys1 := getLogKeys(logRecord1)

	fmt.Printf("Record1 has %v \n", recordKeys1)
	recordKeys2 := getLogKeys(logRecord2)

	fmt.Printf("Record2 has %v \n", recordKeys2)
}
