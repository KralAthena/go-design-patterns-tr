package workerpool

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestHavuz(t *testing.T) {
	isleyici := func(ctx context.Context, i int) (int, error) {
		if i == -1 {
			return 0, errors.New("gecersiz girdi")
		}
		time.Sleep(10 * time.Millisecond)
		return i * 2, nil
	}

	havuz := YeniHavuz(3, isleyici)
	ctx, iptal := context.WithCancel(context.Background())
	defer iptal()

	havuz.Baslat(ctx)

	go func() {
		for i := 1; i <= 5; i++ {
			havuz.GorevEkle(i)
		}
		havuz.GorevEkle(-1)
		havuz.Durdur()
	}()

	sonuclar := make([]int, 0)
	var hataSayisi int

	for res := range havuz.Sonuclar() {
		if res.Hata != nil {
			hataSayisi++
			continue
		}
		sonuclar = append(sonuclar, res.Deger)
	}

	if len(sonuclar) != 5 {
		t.Errorf("5 basarili sonuc beklendi, %d alindı", len(sonuclar))
	}
	if hataSayisi != 1 {
		t.Errorf("1 hata beklendi, %d alindı", hataSayisi)
	}
}
