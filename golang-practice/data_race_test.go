package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestDataRace(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var idr_balance int64 = 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mutex.Lock()
				idr_balance += 1
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(idr_balance)
}

type Human struct {
	Name string
}

func (h Human) Speak() {
	fmt.Println("Hello")
}

func Speak() {
	fmt.Println("Hello")
}

type Speaker interface {
	Speak()
}

func Talk(s Speaker) {
	s.Speak()
}

func TestFunc(t *testing.T) {
	siregar := Human{"Siregar"}

	Talk(siregar)
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func TestShapeInterface(t *testing.T) {
	meja_belajar := Rectangle{5, 10}
	meja_bundar := Circle{5}

	PrintArea(meja_belajar)
	PrintArea(meja_bundar)
}

type Payment interface {
	Pay(amount float64)
}

type Dana struct {
	saldo float64
}
