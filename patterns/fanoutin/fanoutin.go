package fanoutin

import (
	"context"
	"sync"
)

func Dagil[T any, R any](ctx context.Context, giris <-chan T, isciSayisi int, islem func(T) R) <-chan R {
	cikis := make(chan R)
	var grup sync.WaitGroup

	for i := 0; i < isciSayisi; i++ {
		grup.Add(1)
		go func() {
			defer grup.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case deger, acik := <-giris:
					if !acik {
						return
					}
					select {
					case cikis <- islem(deger):
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}

	go func() {
		grup.Wait()
		close(cikis)
	}()

	return cikis
}

func Topla[T any](ctx context.Context, kanallar ...<-chan T) <-chan T {
	var grup sync.WaitGroup
	cikis := make(chan T)

	cokla := func(c <-chan T) {
		defer grup.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case deger, acik := <-c:
				if !acik {
					return
				}
				select {
				case cikis <- deger:
				case <-ctx.Done():
					return
				}
			}
		}
	}

	grup.Add(len(kanallar))
	for _, c := range kanallar {
		go cokla(c)
	}

	go func() {
		grup.Wait()
		close(cikis)
	}()

	return cikis
}
