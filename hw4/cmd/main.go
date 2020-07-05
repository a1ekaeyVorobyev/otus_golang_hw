package main

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_golang_hw/hw4/pkg/lru"
)

func main() {
	lru := lru.NewLru(10)
	lru.Set("12", 12)
	fmt.Print("ddd")
}
