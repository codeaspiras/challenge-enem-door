package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	usualTR, err := NewTimeRule("10:00", "12:00")
	if err != nil {
		fmt.Printf(" # não foi possível inicializar as regras padrões de entrada, sistema abortado. %s\n", err)
		return
	}

	daylightSavingTR, err := NewTimeRule("09:00", "11:00")
	if err != nil {
		fmt.Printf(" # não foi possível inicializar as regras de entrada para horário de verão, sistema abortado. %s\n", err)
		return
	}

	door := &Door{
		usualTR:          usualTR,
		daylightSavingTR: daylightSavingTR,
	}

	fmt.Println(" Essa catraca está em horário de verão? S/N")
	fmt.Print(" > ")
	var answer string
	_, err = fmt.Scanln(&answer)
	if err != nil {
		fmt.Printf(" # não foi possível ler a entrada de configuração, sistema abortado. %s\n", err)
		return
	}

	fmt.Println()
	fmt.Println(" Perfeito. Agora o sistema da catraca será iniciado.")
	fmt.Println(" O sistema somente se encerrará quando receber um")
	fmt.Println(" horário após o limite de entrada.")
	fmt.Println(" Boa sorte.")
	fmt.Println("=====================================================")

	isDaylightSaving := strings.ToUpper(strings.TrimSpace(answer)) == "S"
	for {
		fmt.Print(" Horário (hh:mm): ")
		var t string
		_, err := fmt.Scanln(&t)
		if err != nil {
			fmt.Printf(" # não foi possível ler essa entrada. tente novamente. %s\n #########\n", err)
			continue
		}

		err = door.IsOpenToEnter(t, isDaylightSaving)
		if errors.Is(err, ErrNotOpenedYet) {
			fmt.Printf(" # %s\n", err)
			continue
		}

		if errors.Is(err, ErrClosedAlready) {
			fmt.Printf(" # %s - encerrando sistema\n", err)
			break
		}

		if err != nil {
			fmt.Printf(" # não foi possível verificar essa entrada. tente novamente. %s\n #########\n", err)
			continue
		}

		fmt.Println(" # o portão está liberado, pode passar")
	}
}
