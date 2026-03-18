package singleton

import (
	"sync"
	"testing"
)

func TestOrnekAl(t *testing.T) {
	metin1 := "host=localhost port=5432 user=postgres"
	metin2 := "host=remote-host port=5432 user=admin"

	v1 := OrnekAl(metin1)
	v2 := OrnekAl(metin2)

	if v1 != v2 {
		t.Error("ayni ornek bekleniyordu ancak farkli ornekler alindi")
	}

	if v1.BaglantiMetniAl() != metin1 {
		t.Errorf("%s beklendi, %s alindi", metin1, v1.BaglantiMetniAl())
	}
}

func TestOrnekAlEsZamanli(t *testing.T) {
	const kanalSayisi = 100
	var grup sync.WaitGroup
	ornekler := make([]*Veritabani, kanalSayisi)

	grup.Add(kanalSayisi)
	for i := 0; i < kanalSayisi; i++ {
		go func(indeks int) {
			defer grup.Done()
			ornekler[indeks] = OrnekAl("es-zamanli-baslatma")
		}(i)
	}
	grup.Wait()

	ilkOrnek := ornekler[0]
	for i := 1; i < kanalSayisi; i++ {
		if ornekler[i] != ilkOrnek {
			t.Errorf("%d indeksindeki ornek farkli", i)
		}
	}
}
