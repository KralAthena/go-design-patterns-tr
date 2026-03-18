package ratelimiter

import (
	"context"
	"time"
)

// HizSinirlayici (RateLimiter) gelen istekleri belli bir hızda limitlemek için kullanılır.
type HizSinirlayici struct {
	tokenKovasi chan struct{}
}

// YeniLimitleyici (New) belirli bir kapasite ve hız ile limitleyici oluşturur.
func YeniLimitleyici(kapasite int, hiz time.Duration) *HizSinirlayici {
	rl := &HizSinirlayici{
		tokenKovasi: make(chan struct{}, kapasite),
	}

	// Kovayı token ile dolduran "üretici"
	go func() {
		ticker := time.NewTicker(hiz)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case rl.tokenKovasi <- struct{}{}:
			default:
				// Kova doluysa token eklemiyoruz
			}
		}
	}()

	return rl
}

// IzinBekle (Wait) bir token alana kadar bloklayarak bekler (Context desteği ile).
func (rl *HizSinirlayici) IzinBekle(ctx context.Context) error {
	select {
	case <-rl.tokenKovasi:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// IzinVarMi (Allow) eğer kova boş değilse anında true döner, yoksa false döner.
func (rl *HizSinirlayici) IzinVarMi() bool {
	select {
	case <-rl.tokenKovasi:
		return true
	default:
		return false
	}
}
