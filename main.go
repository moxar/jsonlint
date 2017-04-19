package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {

	var data map[string]interface{}
	err := json.NewDecoder(os.Stdin).Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "   ")

	err = encoder.Encode(data)
	if err != nil {
		log.Fatalln(err)
	}
}
