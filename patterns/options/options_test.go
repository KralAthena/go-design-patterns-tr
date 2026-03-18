package options

import (
	"testing"
	"time"
)

func TestYeniSunucu(t *testing.T) {
	sVarsayilan := YeniSunucu()
	if sVarsayilan.Adres != "localhost:8080" {
		t.Errorf("varsayilan adres localhost:8080 beklendi, %s alindi", sVarsayilan.Adres)
	}

	adres := "0.0.0.0:9000"
	sure := 60 * time.Second
	maks := 100

	s := YeniSunucu(
		Adresle(adres),
		ZamanAsimiylao(sure),
		MaksBaglantiyla(maks),
	)

	if s.Adres != adres {
		t.Errorf("%s beklendi, %s alindi", adres, s.Adres)
	}
	if s.ZamanAsimi != sure {
		t.Errorf("%v beklendi, %v alindi", sure, s.ZamanAsimi)
	}
	if s.MaksBaglanti != maks {
		t.Errorf("%d beklendi, %d alindi", maks, s.MaksBaglanti)
	}
}
