package options

import (
	"time"
)

type Sunucu struct {
	Adres    string
	ZamanAsimi time.Duration
	MaksBaglanti int
}

type Ayar func(*Sunucu)

func Adresle(adres string) Ayar {
	return func(s *Sunucu) {
		s.Adres = adres
	}
}

func ZamanAsimiylao(sure time.Duration) Ayar {
	return func(s *Sunucu) {
		s.ZamanAsimi = sure
	}
}

func MaksBaglantiyla(sayi int) Ayar {
	return func(s *Sunucu) {
		s.MaksBaglanti = sayi
	}
}

func YeniSunucu(ayarlar ...Ayar) *Sunucu {
	s := &Sunucu{
		Adres:    "localhost:8080",
		ZamanAsimi: 30 * time.Second,
		MaksBaglanti: 10,
	}

	for _, ayar := range ayarlar {
		ayar(s)
	}

	return s
}
