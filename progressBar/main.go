package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

func progressBar() {
	barElements := []string{`-`, `\`, `|`, `/`, ""}

	counter := 0
	stop := false

	cCommand := make(chan bool, 1)
	cProgress := make(chan int, 1)

	if keyboard.Open() != nil {
		panic("Err!")
	}
	defer keyboard.Close()

	go func() {
		for {
			counter = <-cProgress / 10
			barElements = barElements[:len(barElements)-1]
			barElements = append(barElements, strings.Repeat("=", counter))
		}
	}()

	go func() {
		for counter <= 100 {
			if counter == 100 {
				fmt.Printf("\r[%s] Загрузка завершена! Нажмите q чтобы продолжить.", strings.Repeat("=", counter))
				break
			}
			for i := 0; i != len(barElements)-1; i++ {
				fmt.Printf("\r[%s%s] %d/100", barElements[len(barElements)-1], barElements[i], counter)
				time.Sleep(20 * time.Microsecond)
			}
		}
	}()

	go func() {
		for {
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			if char == 113 {
				cCommand <- true
				break
			}
		}
	}()

	go func() {
		stop = <-cCommand
	}()

	for i := 0; !stop; i++ {
		if i%10 == 0 && i != 0 {
			cProgress <- i
		}
		time.Sleep(time.Duration(time.Second / 100))
	}
}

func main() {
	progressBar()
}
