package cmap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

var m *CMap

func TestCMap_Get(t *testing.T) {
	fmt.Println("testing get")
	m = NewCMap()
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func(i int) {
			m.Set(strconv.Itoa(i), i*2)
			wg.Done()
		}(i)
	}

	wg.Wait()

	a := assert.New(t)
	r, _ := m.Get("20")
	a.Equal(40, r)
}

func TestCMap_Size(t *testing.T) {
	fmt.Println("testing size")
	if m.Size() != 100 {
		t.Error("size failed")
	}
}
