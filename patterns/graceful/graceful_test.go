package graceful

import (
	"context"
	"testing"
	"time"
)

func TestZarifKapatma(t *testing.T) {
	isYuku := &IsYuku{}
	ctx, cancel := context.WithCancel(context.Background())

	// 3 adet işçi başlatıyoruz
	for i := 1; i <= 3; i++ {
		go isYuku.Calis(ctx, i)
	}

	// İşçilere biraz zaman tanıyalım
	time.Sleep(300 * time.Millisecond)

	// Durdurma işlemine başlıyoruz (Zarif Kapatma Başlıyor)
	cancel()

	// İşlerin güvenli şekilde ve zamanında bitmesini bekliyoruz
	isYuku.Bekle()

	t.Log("Test başarıyla tamamlandı: Tüm işçiler temizlendi.")
}
