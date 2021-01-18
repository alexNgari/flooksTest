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

	wg := &sync.WaitGroup{}

	queue := list.New()

	// Simulate the queue of borrowers to be processed
	for _, person := range dataObject.Borrowers {
		queue.PushBack(person)
	}

	mutex := &sync.Mutex{}
	for queue.Len() > 0 {
		wg.Add(1)
		go worker(wg, queue, mutex)
	}

	wg.Wait()
}

func worker(wg *sync.WaitGroup, queue *list.List, mutex *sync.Mutex) {
	defer wg.Done()
	element := queue.Front()
	if element == nil {
		return
	}
	borrower := element.Value.(models.Borrower)
	queue.Remove(element)
	err := utils.WriteToJSONFile("./test_data/resuts.json", &borrower, mutex)
	if err != nil {
		panic(err)
	}
}

