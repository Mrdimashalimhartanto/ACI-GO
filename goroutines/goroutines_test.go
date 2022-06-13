package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("hello world")
}

func TotalNomorpolis() {
	fmt.Println("NOMOR POLIS")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func TestNomorPolis(t *testing.T) {
	go TotalNomorpolis()
	fmt.Println("INI NOMOR POLIS")

	time.Sleep(2 * time.Second)

}
