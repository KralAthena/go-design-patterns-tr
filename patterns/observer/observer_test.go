package observer

import (
	"fmt"
	"testing"
)

type GercekGozlemci struct {
	Ad   string
	SonMesaj string
}

func (g *GercekGozlemci) Guncelle(o Olay) {
	g.SonMesaj = o.Veri
	fmt.Printf("%s Gozlemcisi mesajı aldı: %s\n", g.Ad, o.Veri)
}

func TestObserver(t *testing.T) {
	konu := YeniKonu()

	g1 := &GercekGozlemci{Ad: "SmsServisi"}
	g2 := &GercekGozlemci{Ad: "EmailServisi"}

	konu.Kaydet(g1)
	konu.Kaydet(g2)

	olay := Olay{Veri: "Sipariş Hazır"}
	konu.Bildir(olay)

	if g1.SonMesaj != "Sipariş Hazır" {
		t.Errorf("Beklenen: 'Sipariş Hazır', Alınan: %s", g1.SonMesaj)
	}

	if g2.SonMesaj != "Sipariş Hazır" {
		t.Errorf("Beklenen: 'Sipariş Hazır', Alınan: %s", g2.SonMesaj)
	}

	konu.Ayril(g1)
	konu.Bildir(Olay{Veri: "Sistem Kapaniyor"})

	if g1.SonMesaj == "Sistem Kapaniyor" {
		t.Error("Ayrılan gözlemciye hala mesaj gidiyor")
	}
}
