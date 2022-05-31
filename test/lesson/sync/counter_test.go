package sync

import (
	"github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/sync"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		testing_tools.AssertEquals(t, counter.Val(), 3)
	})

	t.Run("concurrently test", func(t *testing.T) {
		want := 1000

		counter := Counter{}
		wg := sync.WaitGroup{}
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		testing_tools.AssertEquals(t, counter.Val(), want)
	})
}
