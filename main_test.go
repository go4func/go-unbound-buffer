package main

import (
	"log"
	"sync"
	"testing"
)

func TestUnboundChannel(t *testing.T) {
	in, out := MakeInfinite()
	lastVal := -1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := range out {
			v := i.(int)
			// log.Println("reading", v)
			if lastVal+1 != v {
				t.Errorf("mismatch value, expected %d, got %d", lastVal+1, i)
			}
			lastVal = v
		}
		log.Println("finish reading")
		wg.Done()
	}()

	for i := 0; i < 100; i++ {
		// log.Println("writing", i)
		in <- i
	}
	close(in)
	log.Println("finish writing")
	wg.Wait()

	if lastVal != 99 {
		t.Errorf("missing values, expected %d, got %d", 99, lastVal)
	}
}
