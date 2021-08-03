package main

import (
	"fmt"
	"time"
)

func main() {
	// use append to add new elements to these slices
	var (
		ingredients   []string
		switches      []bool
		historicYears []int
		appointments  []time.Time
	)

	ingredients = append(ingredients, "tomato", "cheese", "basil", "shrimp", "black olives")
	switches = append(switches, true, false, false, true)
	historicYears = append(historicYears, 1918, 1939, 1945)

	now := time.Now()

	appointments = append(appointments, now.Add(time.Hour*18), now.Add(time.Hour*36))

	fmt.Printf("ingredients: %s\n", ingredients)
	fmt.Printf("switches: %t\n", switches)
	fmt.Printf("historic years: %d\n", historicYears)
	fmt.Printf("appointments: %s\n", appointments)
}
