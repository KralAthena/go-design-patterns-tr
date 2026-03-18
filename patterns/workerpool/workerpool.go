package workerpool

import (
	"context"
	"sync"
)

type Gorev[T any, R any] func(context.Context, T) (R, error)

type Sonuc[R any] struct {
	Deger R
	Hata  error
}

type Havuz[T any, R any] struct {
	is_paralel_sayisi int
	gorevler          chan T
	sonuclar          chan Sonuc[R]
	grup              sync.WaitGroup
	isleyici          Gorev[T, R]
}

func YeniHavuz[T any, R any](is_paralel_sayisi int, isleyici Gorev[T, R]) *Havuz[T, R] {
	return &Havuz[T, R]{
		is_paralel_sayisi: is_paralel_sayisi,
		gorevler:          make(chan T),
		sonuclar:          make(chan Sonuc[R]),
		isleyici:          isleyici,
	}
}

func (h *Havuz[T, R]) Baslat(ctx context.Context) {
	for i := 0; i < h.is_paralel_sayisi; i++ {
		h.grup.Add(1)
		go h.isci(ctx)
	}

	go func() {
		h.grup.Wait()
		close(h.sonuclar)
	}()
}

func (h *Havuz[T, R]) isci(ctx context.Context) {
	defer h.grup.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case gorev, acik := <-h.gorevler:
			if !acik {
				return
			}
			val, err := h.isleyici(ctx, gorev)
			h.sonuclar <- Sonuc[R]{Deger: val, Hata: err}
		}
	}
}

func (h *Havuz[T, R]) GorevEkle(gorev T) {
	h.gorevler <- gorev
}

func (h *Havuz[T, R]) Sonuclar() <-chan Sonuc[R] {
	return h.sonuclar
}

func (h *Havuz[T, R]) Durdur() {
	close(h.gorevler)
}
