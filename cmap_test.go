package cmap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

var m *CMap

func TestCMap_Size(t *testing.T) {
	total := 10
	m = NewCMap()
	wg := sync.WaitGroup{}
	wg.Add(total)
	for i := 1; i <= total; i++ {
		go func(iv int) {
			m.Set(strconv.Itoa(iv), iv*2)
			wg.Done()
		}(i)
	}

	wg.Wait()

	a := assert.New(t)
	a.Equal(total, m.Size())
}

func TestCMap_Dump(t *testing.T) {
	out := m.Dump()
	for item := range out {
		fmt.Println(item.Key, "=>", item.Value)
	}
}
