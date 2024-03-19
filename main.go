package main

import (
	"fmt"
	"time"
)

func main() {
	var horas, minutos, segundos int

	for horas = 0; horas < 24; horas++ {
		for minutos = 0; minutos < 60; minutos++ {
			for segundos = 0; segundos < 60; segundos++ {
				fmt.Printf("%02d:%02d:%02d\n", horas, minutos, segundos)
				time.Sleep(time.Second * 2)
			}

		}
	}
}
