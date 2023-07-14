package web_server

import (
	"container/heap"
	"fmt"
	"net/rpc"
	"sort"
	"testing"
)

func TestName(t *testing.T) {
	a := complex(2.0, 3.0)
	t.Log(a * a)
}

func TestRpc(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Log(err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		t.Log(err)
	}

	t.Log(reply)
}

func TestString(t *testing.T) {
	s := " "
	m := map[rune]struct{}{}
	arr := []rune{}
	tail := 0
	windowLen := 0
	maxLen := 0
	for _, v := range s {
		if _, ok := m[v]; ok {
			windowLen = len(arr) - tail
			if windowLen > maxLen {
				maxLen = windowLen
			}
			for arr[tail] != v {
				delete(m, arr[tail])
				tail += 1
			}
			delete(m, arr[tail])
			tail += 1
		}
		arr = append(arr, v)
		m[v] = struct{}{}
	}
	t.Log(maxLen)
}

func TestOJ(t *testing.T) {
	n := 0
	arr := []int{}
	fmt.Scan(&n)
	for n >= 0 {
		a, b := 0, 0
		fmt.Scan(&a, &b)
		arr = append(arr, a, b)
	}
	t.Log(arr)
	sort.Slice(arr, func(i, j int) bool {
		return true
	})
}

type Heap struct{ sort.IntSlice }

func (h Heap) Push(any) { return }

func (h Heap) Pop() (_ any) { return }

func (h Heap) replace(v int) int {
	top := h.IntSlice[0]
	h.IntSlice[0] = v
	heap.Fix(h, 0)
	return top
}

func TestCopy(t *testing.T) {
	arr := []int{2, 3, 4, 5}
	arr2 := make([]int, len(arr), len(arr))
	copy(arr2, arr)
	arr2[0] = 1
	t.Log(arr, arr2)
	h := Heap{arr}
	heap.Init(h)
	heap.Push(&h, 20)
}

func TestSlice(t *testing.T) {
	for i := 4; i <= 10; i++ {
		arr := make([]int, 1<<i, 1<<i)
		t.Logf("before:%d, %p", cap(arr), arr)
		arr = append(arr, 0)
		t.Logf("after:%d, %p", cap(arr), arr)
	}
}

func TestTmp(t *testing.T) {
	t.Log(848.0/512.0, 1536.0/1024.0)
}

func TestChan(t *testing.T) {
	var ch chan int
	ch = make(chan int)
	go func() {
		ch <- 5
	}()
	v := <-ch
	t.Log("success", v)
}

func TestArrCopy(t *testing.T) {
	arr := []int{1, 2, 3}
	t.Logf("before: %p,%p,cap: %d\n", arr, &arr[0], cap(arr))
	findArr(arr)
	t.Logf("after value:%d len:%d\n", arr[0], len(arr))
}

func findArr(arr []int) {
	fmt.Printf("after:%p\n", arr)
	for i := 0; i < 20; i++ {
		arr = append(arr, 10)
	}
	arr[0] = 10
	fmt.Printf("append: %p\n", arr)
}
