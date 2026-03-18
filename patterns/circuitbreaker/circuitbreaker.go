package circuitbreaker

import (
	"errors"
	"sync"
	"time"
)

type Durum int

const (
	Kapali Durum = iota
	Acik
	YariAcik
)

type DevreKesici struct {
	durum            Durum
	hataEsigi        int
	hataSayisi       int
	beklemeSuresi    time.Duration
	sonHataZamani    time.Time
	mu               sync.Mutex
}

func YeniDevreKesici(esik int, bekleme time.Duration) *DevreKesici {
	return &DevreKesici{
		durum:         Kapali,
		hataEsigi:     esik,
		beklemeSuresi: bekleme,
	}
}

func (dk *DevreKesici) Calistir(islem func() error) error {
	dk.mu.Lock()
	
	if dk.durum == Acik {
		if time.Since(dk.sonHataZamani) > dk.beklemeSuresi {
			dk.durum = YariAcik
		} else {
			dk.mu.Unlock()
			return errors.New("devre acik: islem reddedildi")
		}
	}
	dk.mu.Unlock()

	err := islem()

	dk.mu.Lock()
	defer dk.mu.Unlock()

	if err != nil {
		dk.hataSayisi++
		dk.sonHataZamani = time.Now()
		if dk.hataSayisi >= dk.hataEsigi {
			dk.durum = Acik
		}
		return err
	}

	if dk.durum == YariAcik {
		dk.durum = Kapali
		dk.hataSayisi = 0
	}
	
	dk.hataSayisi = 0
	return nil
}
