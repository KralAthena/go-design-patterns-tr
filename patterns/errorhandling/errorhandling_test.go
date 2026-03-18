package errorhandling

import (
	"errors"
	"testing"
)

func TestHataKontrolu(t *testing.T) {
	err := IslemYap(-1)
	if err == nil {
		t.Fatal("hata bekliyordum")
	}

	var vh *VeritabaniHatasi
	if !errors.As(err, &vh) {
		t.Fatal("veritabani hatasi beklendi")
	}

	if vh.Kod != 404 {
		t.Errorf("404 beklendi, %d alindı", vh.Kod)
	}

	mesaj := HataKontrolu(err)
	expected := "Ozel Hata Yakalandi: gecersiz id"
	if mesaj != expected {
		t.Errorf("'%s' beklendi, %s alindı", expected, mesaj)
	}
}
