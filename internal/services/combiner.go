package services

import (
	"fmt"
	"sync"
)

type Combiner struct {
	Names []string
}

func (c Combiner) CombineFor() (list []string, err error) {
	var last string
	for _, name := range c.Names {
		list = append(list, fmt.Sprintf("%s %s", name, last))
		last = name
	}
	return
}

func (c Combiner) CombineRoutine() (list []string, err error) {
	var wg sync.WaitGroup
	var last string
	namesChan := make(chan string, len(c.Names))
	fmt.Printf("here1")
	for _, name := range c.Names {
		wg.Add(1)
		fmt.Printf("here2")

		go func(first string, second string) {
			namesChan <- fmt.Sprintf("%s %s", first, second)
			wg.Done()
		}(name, last)
		last = name
	}
	fmt.Printf("here3")

	wg.Wait()
	close(namesChan)
	fmt.Printf("here4")

	for doneName := range namesChan {
		list = append(list, doneName)
	}
	return
}
