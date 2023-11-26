package main

// sumNumbers, bir dizi içindeki sayıları toplayan bir fonksiyondur.
func toplama(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// func main() {
// 	// Bir sayı dizisi oluşturun
// 	numbers := []int{1, 2, 3, 4, 5}

// 	// sumNumbers fonksiyonunu kullanarak dizideki sayıları toplayın
// 	result := toplama(numbers)

// 	// Sonucu ekrana yazdırın
// 	fmt.Printf("Dizi içindeki sayıların toplamı: %d\n", result)
// }
