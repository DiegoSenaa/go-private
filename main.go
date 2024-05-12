package main

import (
	"fmt"

	events "github.com/DiegoSenaa/eventos-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
