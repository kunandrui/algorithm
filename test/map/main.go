package main

import "fmt"

func main() {
	hash := make(map[string]string)
	hash["key1"] = "value1"
	hash["key2"] = "value2"
	fmt.Printf("%v", hash)
}
