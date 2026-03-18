package observer

import (
	"sync"
)

// Olay (Event) bir bildirimi temsil eder.
type Olay struct {
	Veri string
}

// Gozlemci (Observer) olayları dinleyen taraf.
type Gozlemci interface {
	Guncelle(Olay)
}

// Konu (Subject) olayları yayınlayan taraf.
type Konu struct {
	gozlemciler map[Gozlemci]struct{}
	mu          sync.RWMutex
}

func YeniKonu() *Konu {
	return &Konu{
		gozlemciler: make(map[Gozlemci]struct{}),
	}
}

// Kaydet (Subscribe) yeni bir gözlemciyi listeye ekler.
func (k *Konu) Kaydet(g Gozlemci) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.gozlemciler[g] = struct{}{}
}

// Ayril (Unsubscribe) bir gözlemciyi listeden çıkarır.
func (k *Konu) Ayril(g Gozlemci) {
	k.mu.Lock()
	defer k.mu.Unlock()
	delete(k.gozlemciler, g)
}

// Bildir (Notify) tüm gözlemcilere olayı iletir.
func (k *Konu) Bildir(o Olay) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	for g := range k.gozlemciler {
		g.Guncelle(o)
	}
}
