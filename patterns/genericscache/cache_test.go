package genericscache

import (
	"testing"
)

func TestJenerikOnbellek(t *testing.T) {
	// String anahtar, Int deger ile test ediyoruz
	cache := YeniOnbellek[string, int]()

	cache.Ekle("puan", 100)
	cache.Ekle("seviye", 5)

	val, ok := cache.Getir("puan")
	if !ok || val != 100 {
		t.Errorf("Beklenen: 100, Alınan: %v", val)
	}

	cache.Sil("seviye")
	_, ok = cache.Getir("seviye")
	if ok {
		t.Error("Silinen veri hala getiriliyor")
	}

	// Bir de yapılar (Struct) ile test edelim
	type Kullanici struct {
		ID   int
		Isim string
	}

	userCache := YeniOnbellek[int, Kullanici]()
	userCache.Ekle(1, Kullanici{ID: 1, Isim: "Can"})

	user, ok := userCache.Getir(1)
	if !ok || user.Isim != "Can" {
		t.Errorf("Beklenen: Can, Alınan: %s", user.Isim)
	}
}
