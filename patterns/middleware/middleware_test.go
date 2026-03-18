package middleware

import (
	"testing"
)

func TestAraKatman(t *testing.T) {
	islem := OrnekIslem
	
	// Sarmalama
	islem = Loglayici(islem)
	islem = ZamanOlcer(islem)

	sonuc := islem("Merhaba Dunya")

	beklenen := "Islendi: Merhaba Dunya"
	if sonuc != beklenen {
		t.Errorf("'%s' beklendi, '%s' alindi", beklenen, sonuc)
	}
}
