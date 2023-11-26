package main

import (
	"fmt"
	"os"
)

func main() {
	var bolunen, bolen float64

	fmt.Print("Bölünen sayıyı girin: ")
	_, err := fmt.Scan(&bolunen)
	if err != nil {
		fmt.Println("Hata: Bir sayı girmelisiniz.")
		os.Exit(1)
	}

	fmt.Print("Bölen sayıyı girin: ")
	_, err = fmt.Scan(&bolen)
	if err != nil {
		fmt.Println("Hata: Bir sayı girmelisiniz.")
		os.Exit(1)
	}

	if bolen == 0 {
		fmt.Println("Hata: Bölen sıfır olamaz.")
		os.Exit(1)
	}

	sonuc := bolunen / bolen
	fmt.Printf("%.2f / %.2f = %.2f\n", bolunen, bolen, sonuc)
}
