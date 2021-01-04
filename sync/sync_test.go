package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 timees leaves it at 3", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup

		//先add数量
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				//done 代表完成一个
				w.Done()
			}(&wg)
		}

		//所有数量都已经done
		//使用wait 就可以肯定所有都已经完成
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

//互斥对象在第一次使用后不能复制。所以这里的got需要使用引用
func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
