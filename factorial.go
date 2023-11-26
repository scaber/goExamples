// Sayının faktöriyelini bulan go programı
package main

import "fmt"

/* Variable Declaration */
var factVal uint64 = 1 // uint64, imzasız tüm 64 bit tamsayıların kümesidir.
// Aralık: 0 ila 18446744073709551615.
var i int = 1
var n int

/*    fonksiyon bildirimi      */
func factorial(n int) uint64 {
	if n < 0 {
		fmt.Print("Factoriyel sayı negatif olmaz.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i) // uyumsuz tipler int64 ve int
		}
	}
	return factVal /* fonksiyon değeri döndürür*/
}

// func main() {
// 	fmt.Print("0 - 50 arasında bir sayı girin: ")
// 	fmt.Scan(&n)
// 	fmt.Print("Factoriyel sonuç : ", factorial(n))
// }
