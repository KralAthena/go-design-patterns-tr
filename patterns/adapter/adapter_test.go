package adapter

import (
	"io"
	"testing"
)

func TestUyarlayici(t *testing.T) {
	eskiVeri := &OzelVeri{Govde: "Merhaba Dunya!"}
	
	// Eski yapı bir Reader değil, ama onu uyarlıyoruz.
	uyarlayici := YeniUyarlayici(eskiVeri)

	// Artık standart io kütüphaneleriyle çalışabilir.
	icerik, err := io.ReadAll(uyarlayici)
	if err != nil {
		t.Fatalf("Veri okunamadı: %v", err)
	}

	if string(icerik) != "Merhaba Dunya!" {
		t.Errorf("Beklenen: 'Merhaba Dunya!', Alınan: %s", string(icerik))
	}
}
