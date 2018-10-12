package main

import (
	"log"

	"github.com/Siryu/interview/internal/services"
)

func main() {
	names := []string{
		"corey",
		"massey",
		"bob",
		"henry",
	}
	combiner := services.Combiner{
		Names: names,
	}

	result, _ := combiner.CombineRoutine()
	for _, r := range result {
		log.Print(r)
	}
}
