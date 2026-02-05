package main

import (
	"fmt"
	"sync"
	"time"
)

func AfficheChiffre(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 6; i++ {
		fmt.Println("Nombre : ", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func AfficheLettre(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'A'; i < 'F'; i++ {
		fmt.Println("Lettre : ", string(i))
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	fmt.Println("Début")
	go AfficheChiffre(&wg)
	go AfficheLettre(&wg)

	fmt.Println("On attend la fin")
	wg.Wait()
	fmt.Println("Fin")

}
