package ratelimiter

import (
	"context"
	"testing"
	"time"
)

func TestHizSinirlayici(t *testing.T) {
	// Saniyede 1 token, kova kapasitesi 1
	rl := YeniLimitleyici(1, 100*time.Millisecond)

	// Başlangıçta 1-2 ms bekle kova dolsun
	time.Sleep(150 * time.Millisecond)

	if !rl.IzinVarMi() {
		t.Error("İlk istekte izin verilmeliydi")
	}

	if rl.IzinVarMi() {
		t.Error("Kova boş olmalı, izin verilmemeliydi")
	}

	// İzin bekleme testi
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	err := rl.IzinBekle(ctx)
	if err != nil {
		t.Errorf("Error beklemiyorduk, token dolmalıydı: %v", err)
	}

	// Timeout testi
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel2()
	err = rl.IzinBekle(ctx2)
	if err == nil {
		t.Error("Kısa sürede timeout hatası bekliyorduk")
	}
}
