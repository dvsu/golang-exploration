package main

import "fmt"

type character struct {
	hp, mp, level int
	name          string
}

func (c *character) status() {
	fmt.Printf("Name: %s Level: %d HP: %d MP: %d\n", c.name, c.level, c.hp, c.mp)
}

func (c *character) levelup() {
	c.level++
	c.hp = int(float64(c.hp) * 1.1)
	c.mp = int(float64(c.mp) * 1.05)
}