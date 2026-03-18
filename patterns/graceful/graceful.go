package graceful

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// IsYuku (Workload) örneği
type IsYuku struct {
	wg sync.WaitGroup
}

// Calis (Run) metodu bir context ile çalışır ve durdurulana kadar iş yapar.
func (i *IsYuku) Calis(ctx context.Context, id int) {
	i.wg.Add(1)
	defer i.wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Durduruluyor...\n", id)
			return
		default:
			// İş yapılıyormuş gibi simüle ediyoruz
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Durdur (Wait) tüm işlerin bitmesini bekler.
func (i *IsYuku) Bekle() {
	i.wg.Wait()
	fmt.Println("Tüm işçiler güvenli şekilde durduruldu.")
}
