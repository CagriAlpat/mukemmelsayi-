package main

 import "fmt"

 /*MUKEMMEL SAYİ
 kendi bölimlerinin toplamı yine kendine eşit olan sayidir
 28=14+7+4+2+1
 6=3+2+1
 */

 func main() {
 	var sayi int

 	fmt.Print("sayi giriniz ")
 	fmt.Scanf("%d", &sayi)

 	toplam := 0

 	for i := 1; i < sayi; i++ {
 		if sayi%i == 0 {
 			toplam += i
 		}
 	}
 	if toplam == sayi {
 		fmt.Printf("%d sayisi mukemmel sayidir ", sayi)
 	} else {
 		fmt.Printf("%d sayisi mükemmel sayi değildir ", sayi)
 	}
 }
