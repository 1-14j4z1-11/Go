package memo

import (
	"testing"
)

func TestCancel(t *testing.T) {
	const try = 4

	f, _, counter := makeWaitableFunc()
	memo := New(f)
	doneCh := make(chan chan struct{}, try)
	result := make(chan bool)

	for i := 0; i < try; i++ {
		done := make(chan struct{})
		doneCh <- done
		go func() {
			_, _, cancelled := memo.Get("key", done)
			result <- cancelled
		}()
	}

	for i := 0; i < try; i++ {
		close(<-doneCh)
	}


	if *counter != 1 {
		t.Errorf("Unexpected call count %d != %d", *counter , 1)
	}

	for i := 0; i < try; i++ {
		if !<-result {
			t.Error("Unexpected result")
		}
	}
}

func TestConcurrent(t *testing.T) {
	const try = 2

	f, ready, counter := makeWaitableFunc()
	memo := New(f)
	done := make(chan struct{})
	results := make(chan string)

	tryGet := func(key string) {
		r, _, _ := memo.Get(key, done)
		results <- r.(string)
	}

	for i := 0; i < try; i++ {
		go tryGet("key1")
	}
	for i := 0; i < try; i++ {
		go tryGet("key2")
	}
	for i := 0; i < try; i++ {
		go tryGet("key3")
	}

	close(ready)

	resultMap := make(map[string]int)
	for i := 0; i < try*3; i++ {
		resultMap[<-results]++
	}

	count, ok := resultMap["key1"]
	if count != try || !ok {
		t.Errorf("key1 count = %d, want = %d", count, try)
	}

	count, ok = resultMap["key2"]
	if count != try || !ok {
		t.Errorf("key2 count = %d, want = %d", count, try)
	}

	count, ok = resultMap["key3"]
	if count != try || !ok {
		t.Errorf("key3 count = %d, want = %d", count, try)
	}

	if *counter != 3 {
		t.Error("failed")
	}
}

func makeWaitableFunc() (func(key string, done <-chan struct{}) (interface{}, error, bool), chan<- struct{}, *int) {
	ready := make(chan struct{})
	counter := 0

	f := func(key string, done <-chan struct{}) (interface{}, error, bool) {
		counter++
		select {
		case <-ready:
			return key, nil, false
		case <-done:
			return "", nil, true
		}
	}

	return f, ready, &counter
}
