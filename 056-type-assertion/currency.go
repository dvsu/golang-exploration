package main

import "fmt"

type currency float64

func (c currency) string() string {
	return fmt.Sprintf("$%.2f", c)
}
