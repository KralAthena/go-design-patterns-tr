package singleton

import (
	"sync"
)

type Veritabani struct {
	baglantiMetni string
}

var (
	ornek *Veritabani
	birkez sync.Once
)

func OrnekAl(metin string) *Veritabani {
	birkez.Do(func() {
		ornek = &Veritabani{
			baglantiMetni: metin,
		}
	})
	return ornek
}

func (v *Veritabani) BaglantiMetniAl() string {
	return v.baglantiMetni
}
