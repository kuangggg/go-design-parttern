package singleton

import (
	"sync"
	"testing"
)

func TestGetSingleton(t *testing.T) {
	ins := GetSingleton()
	ins1 := GetSingleton()

	if ins != ins1 {
		t.Fatal("ins is not eq")
	}
}

const count = 100

func TestMulSingleton(t *testing.T)  {

	wg := sync.WaitGroup{}

	wg.Add(count)

	insArr := [count]*Singleton{}

	for i:= 0; i < count; i++ {
		go func(idx int) {
			insArr[idx] = GetSingleton()
			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 1; i < count; i++ {
		if insArr[i] != insArr[i-1] {
			t.Fatal("ins is not eq")
		}
	}

}