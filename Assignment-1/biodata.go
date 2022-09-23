package main

import (
	"Assigment-1/peserta"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please enter Args")
		return
	}
	noUrut, err := strconv.Atoi(os.Args[1])
	_ = err

	if noUrut > len(peserta.Murid) || noUrut < 1 {
		fmt.Println("Invalid noUrut")
		return
	}

	peserta.Murid[noUrut-1].Cetak()
}
