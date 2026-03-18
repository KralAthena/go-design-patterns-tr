package fanoutin

import (
	"context"
	"sort"
	"testing"
)

func TestDagil(t *testing.T) {
	ctx, iptal := context.WithCancel(context.Background())
	defer iptal()

	giris := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			giris <- i
		}
		close(giris)
	}()

	isçiSayisi := 3
	cikis := Dagil(ctx, giris, isçiSayisi, func(i int) int {
		return i * 2
	})

	sonuclar := make([]int, 0)
	for val := range cikis {
		sonuclar = append(sonuclar, val)
	}

	if len(sonuclar) != 10 {
		t.Errorf("10 sonuc beklendi, %d alindı", len(sonuclar))
	}

	sort.Ints(sonuclar)
	for i, v := range sonuclar {
		if v != (i+1)*2 {
			t.Errorf("%d beklendi, %d alindı (indeks %d)", (i+1)*2, v, i)
		}
	}
}

func TestTopla(t *testing.T) {
	ctx, iptal := context.WithCancel(context.Background())
	defer iptal()

	k1 := make(chan string)
	k2 := make(chan string)

	go func() {
		k1 <- "A"
		k1 <- "B"
		close(k1)
	}()

	go func() {
		k2 <- "1"
		k2 <- "2"
		close(k2)
	}()

	birlesik := Topla(ctx, k1, k2)

	sonuclar := make([]string, 0)
	for val := range birlesik {
		sonuclar = append(sonuclar, val)
	}

	if len(sonuclar) != 4 {
		t.Errorf("4 sonuc beklendi, %d alindı", len(sonuclar))
	}
}
