package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	a = iota + 1
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	l
	m
	n
	o
	p
	q
	r
	s
	t
	u
	v
	w
	x
	y
	z
	A
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
)

func main() {

	puzzle, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ruckSacks := strings.Split(string(puzzle), "\n")

	for _, rs := range ruckSacks {
		//Easiest way I could find to "split" a string in half
		//The : seems to indicate 'take the first item' or 'take the last item'
		//If I had done [2:len(rs)] it would take the
		compartmentOne := rs[:len(rs)/2]
		compartmentTwo := rs[len(rs)/2:]

		fmt.Println(compartmentOne)
		fmt.Println(compartmentTwo)
	}
}
