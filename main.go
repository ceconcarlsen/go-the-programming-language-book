package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type Point struct {
	X int
	Y int
}
type Circle struct {
	Point
	Radius int
}
type Shape interface {
	Area() float64
	Perimeter() float64
}

func sayHello() {
	fmt.Println("Olá de uma goroutine!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá, mundo!")
}

func main() {

	var pointer *string

	fmt.Println("Pointer:", &pointer) // Imprime nil

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	go sayHello()
	time.Sleep(1 * time.Second) // Espera para a goroutine terminar
	fmt.Println("Fim do main")

	fmt.Println("Hello, World!")

	arr := []int{3, 1, 4, 1, 5}
	sort.Ints(arr)

	fmt.Println("Sorted array:", arr)

	//slices
	var s, sep string //empty string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("Command line arguments:", s)

	for j, _ := range os.Args {
		fmt.Println("Index", j)
	}
	//bufio
	file, err := os.Open("index.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

}
