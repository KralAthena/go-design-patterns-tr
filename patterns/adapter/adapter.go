package adapter

import (
	"strings"
)

// OzelVeri (Legacy System) bizim değiştiremediğimiz eski bir yapı.
type OzelVeri struct {
	Govde string
}

func (v *OzelVeri) IcerigiGetir() string {
	return v.Govde
}

// Uyarlayici (Adapter) bu eski yapıyı Go'nun io.Reader arayüzüne uyarlar.
// Go kütüphaneleri (json.Decoder, http vb.) io.Reader bekler.
type Uyarlayici struct {
	veri   *OzelVeri
	okuyucu *strings.Reader
}

func YeniUyarlayici(v *OzelVeri) *Uyarlayici {
	return &Uyarlayici{
		veri:   v,
		okuyucu: strings.NewReader(v.IcerigiGetir()),
	}
}

// Read (io.Reader) arayüzünü uygular.
func (u *Uyarlayici) Read(p []byte) (n int, err error) {
	return u.okuyucu.Read(p)
}
