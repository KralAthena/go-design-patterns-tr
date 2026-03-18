package pipeline

import (
	"context"
	"testing"
)

func TestVeriHatti(t *testing.T) {
	ctx, iptal := context.WithCancel(context.Background())
	defer iptal()

	uretici := Uret(ctx, 1, 2, 3, 4, 5)
	kareleme := KareAl(ctx, uretici)
	filtreleme := Filtrele(ctx, kareleme, 10)

	sonuclar := make([]int, 0)
	for s := range filtreleme {
		sonuclar = append(sonuclar, s)
	}

	if len(sonuclar) != 2 {
		t.Errorf("2 sonuc beklendi, %d alindı", len(sonuclar))
	}

	if sonuclar[0] != 16 || sonuclar[1] != 25 {
		t.Errorf("16 ve 25 beklendi, %v alindı", sonuclar)
	}
}
