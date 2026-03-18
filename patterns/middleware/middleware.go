package middleware

import (
	"fmt"
	"time"
)

type Islem func(string) string

func ZamanOlcer(islem Islem) Islem {
	return func(s string) string {
		baslangic := time.Now()
		sonuc := islem(s)
		fmt.Printf("Islem suresi: %v\n", time.Since(baslangic))
		return sonuc
	}
}

func Loglayici(islem Islem) Islem {
	return func(s string) string {
		fmt.Printf("Giris parametresi: %s\n", s)
		sonuc := islem(s)
		fmt.Printf("Sonuc: %s\n", sonuc)
		return sonuc
	}
}

func OrnekIslem(mesaj string) string {
	time.Sleep(100 * time.Millisecond)
	return "Islendi: " + mesaj
}
