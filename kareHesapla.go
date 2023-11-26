package main

import (
	"sync"
)

// kareHesapla, bir dizi içindeki sayıların karesini hesaplayan fonksiyondur.
func kareHesapla(sayilar []int, sonucChan chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Gorutine tamamlandığında WaitGroup'i azalt

	for _, sayi := range sayilar {
		kare := sayi * sayi
		sonucChan <- kare
	}
}

// func main() {
// 	// Kullanılacak sayı dizisi
// 	sayilar := []int{11, 22, 33, 44, 55}

// 	// Gorutineler arasında sonuçları iletmek için bir kanal oluşturun
// 	sonucChan := make(chan int, len(sayilar))

// 	// WaitGroup, tüm gorutinelerin tamamlanmasını beklemek için kullanılır
// 	var wg sync.WaitGroup

// 	// Her bir sayının karesini hesaplamak için bir gorutine başlatın
// 	for _, sayi := range sayilar {
// 		wg.Add(1) // WaitGroup'e bir gorutine eklendi
// 		go kareHesapla([]int{sayi}, sonucChan, &wg)
// 	}

// 	// Tüm gorutinelerin tamamlanmasını bekleyin
// 	wg.Wait()

// 	// Kapanan kanalı kontrol etmek için close() kullanın
// 	close(sonucChan)

// 	// Kanaldan sonuçları okuyun ve ekrana yazdırın
// 	for kare := range sonucChan {
// 		fmt.Printf("Karesi: %d\n", kare)
// 	}
// }
