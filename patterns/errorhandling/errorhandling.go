package errorhandling

import (
	"errors"
	"fmt"
)

type VeritabaniHatasi struct {
	Mesaj string
	Kod   int
}

func (h *VeritabaniHatasi) Error() string {
	return fmt.Sprintf("Veritabanı Hatası [%d]: %s", h.Kod, h.Mesaj)
}

func VeriGetir(id int) error {
	if id <= 0 {
		return &VeritabaniHatasi{
			Mesaj: "gecersiz id",
			Kod:   404,
		}
	}
	return nil
}

func IslemYap(id int) error {
	err := VeriGetir(id)
	if err != nil {
		return fmt.Errorf("islem basarisiz: %w", err)
	}
	return nil
}

func HataKontrolu(err error) string {
	var vh *VeritabaniHatasi
	if errors.As(err, &vh) {
		return fmt.Sprintf("Ozel Hata Yakalandi: %s", vh.Mesaj)
	}
	return "Genel Hata: " + err.Error()
}
