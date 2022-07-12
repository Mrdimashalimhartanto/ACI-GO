package main

import (
	"fmt"
	"testing"
	"time"
)

// TODO: membuat channel golang
// channel dalam golang adalah CHAN

// channel harus di close agar memory tidak licking & go run time sudah mengerti karna sudah di close dari memory

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	//
	//channel <- "dimas"
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Dimas Halim Hartanto"
		fmt.Println("Selesai Mengirim Data")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

// test channel sebagai parameter

func KasihSayaResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Dimas Halim Hartanto"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	//
	//channel <- "dimas"
	defer close(channel)

	go KasihSayaResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

//========= TODO: materi channel in & out parameter

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Dimas Halim Hartanto"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

//========= TODO: buffered channel buffered adalah hitungan / jumlah yang bisa di berikan
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	//channel <- "API CMS EPAY"

	// TODO: PENGGUNAAN ANONYMOUS FUNCTION
	go func() {
		channel <- "Dimas"
		channel <- "Halim"
		channel <- "Hartanto"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("API CMS BERHASIL DI AMBIL")
}
