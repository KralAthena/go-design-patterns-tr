package genericscache

import (
	"sync"
)

// Onbellek (Cache) generic türlerle (K key, V value) çalışabilen,
// eşzamanlı erişime uygun bir veri yapısıdır.
type Onbellek[K comparable, V any] struct {
	veriler map[K]V
	mu      sync.RWMutex
}

func YeniOnbellek[K comparable, V any]() *Onbellek[K, V] {
	return &Onbellek[K, V]{
		veriler: make(map[K]V),
	}
}

// Ekle (Set) bir değeri anahtar ile saklar.
func (o *Onbellek[K, V]) Ekle(anahtar K, deger V) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.veriler[anahtar] = deger
}

// Getir (Get) bir değeri anahtara göre döndürür.
func (o *Onbellek[K, V]) Getir(anahtar K) (V, bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	deger, ok := o.veriler[anahtar]
	return deger, ok
}

// Sil (Delete) veriyi temizler.
func (o *Onbellek[K, V]) Sil(anahtar K) {
	o.mu.Lock()
	defer o.mu.Unlock()
	delete(o.veriler, anahtar)
}
