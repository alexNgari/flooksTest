package main

import (
	"github.com/alexNgari/flooksTest/models"
	"sync"
	"container/list"

	"github.com/alexNgari/flooksTest/utils"
)


func main() {
	dataObject, err := utils.ReadJSONFile("./test_data/data.json")
	if err != nil {
		panic(err)
	}

	

	queue := list.New()

	// Simulate the queue of borrowers to be processed
	for _, person := range dataObject.Borrowers {
		queue.PushBack(person)
	}

	mutex := &sync.Mutex{}
	for queue.Len() > 0 {
		go func() {
			element := queue.Front()
			borrower := element.Value.(models.Borrower)
			utils.WriteToJSONFile("./test_data/resuts.json", &borrower, mutex)
			queue.Remove(element)
		} ()
	}
}

