package circuitbreaker

import (
	"errors"
	"testing"
	"time"
)

func TestDevreKesici(t *testing.T) {
	dk := YeniDevreKesici(2, 50*time.Millisecond)

	err := dk.Calistir(func() error {
		return errors.New("hata")
	})
	if err == nil {
		t.Error("hata bekliyordum")
	}

	err = dk.Calistir(func() error {
		return errors.New("hata")
	})
	if err == nil {
		t.Error("ikinci hatayi bekliyordum")
	}

	err = dk.Calistir(func() error {
		return nil
	})
	if err == nil || err.Error() != "devre acik: islem reddedildi" {
		t.Errorf("devre acik olmaliydi: %v", err)
	}

	time.Sleep(60 * time.Millisecond)

	err = dk.Calistir(func() error {
		return nil
	})
	if err != nil {
		t.Errorf("devre kapali olmaliydi: %v", err)
	}
}
