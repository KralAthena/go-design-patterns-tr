package pipeline

import (
	"context"
)

func Uret(ctx context.Context, sayilar ...int) <-chan int {
	cikis := make(chan int)
	go func() {
		defer close(cikis)
		for _, s := range sayilar {
			select {
			case <-ctx.Done():
				return
			case cikis <- s:
			}
		}
	}()
	return cikis
}

func KareAl(ctx context.Context, giris <-chan int) <-chan int {
	cikis := make(chan int)
	go func() {
		defer close(cikis)
		for s := range giris {
			select {
			case <-ctx.Done():
				return
			case cikis <- s * s:
			}
		}
	}()
	return cikis
}

func Filtrele(ctx context.Context, giris <-chan int, esik int) <-chan int {
	cikis := make(chan int)
	go func() {
		defer close(cikis)
		for s := range giris {
			if s > esik {
				select {
				case <-ctx.Done():
					return
				case cikis <- s:
				}
			}
		}
	}()
	return cikis
}
