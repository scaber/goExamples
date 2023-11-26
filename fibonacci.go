package main

// fibonacci, Fibonacci dizisinin n. elemanını hesaplayan bir fonksiyondur.
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// func main() {
// 	// İlk 10 Fibonacci sayısını yazdıralım.
// 	for i := 0; i < 10; i++ {
// 		result := fibonacci(i)
// 		fmt.Printf("Fibonacci(%d) = %d\n", i, result)
// 	}
// }
