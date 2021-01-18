package main

import (
	"sync"

	"github.com/alexNgari/flooksTest/utils"
)


func main() {
	dataObject, err := utils.ReadJSONFile("./test_data/data.json")
	if err != nil {
		panic(err)
	}

	mutex := &sync.Mutex{}

	for _, person := range dataObject.Borrowers {
		go func() {
			utils.WriteToJSONFile("./test_data/resuts.json", &person, mutex)
		} ()
	}
}

