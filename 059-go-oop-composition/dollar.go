package main

import (
	"fmt"
)

type dollar float64

func (d dollar) string() string {
	return fmt.Sprintf("$%.2f", d)
}
