package memo

type Func func(key string, done <-chan struct{}) (interface{}, error, bool)

type result struct {
	value     interface{}
	err       error
	cancelled bool
}

type entry struct {
	res    result
	ready  chan struct{}
	cancel chan struct{}
}

type request struct {
	key      string
	response chan<- result
	done     <-chan struct{}
}

type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error, bool) {
	response := make(chan result)
	memo.requests <- request{key, response, done}

	res := <-response
	return res.value, res.err, res.cancelled
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	removeKey := make(chan string)
	active := true

	for active {
		select {
		case req, ok := <-memo.requests:
			if !ok {
				active = false
				break
			}

			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{}), cancel:make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done)
			}
			go e.deliver(req.response, req.done, removeKey, req.key)
		case key := <-removeKey:
			delete(cache, key)
		}
	}

	for range removeKey {
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	e.res.value, e.res.err, e.res.cancelled = f(key, done)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result, done <-chan struct{}, removeKey chan<- string, key string) {

	select {
	case <-e.ready:
		response <- e.res
	case <-done:
		e.res = result{nil, nil, true}
		removeKey <- key
		response <- e.res
		close(e.cancel)
	case <-e.cancel:
		response <- e.res
	}
}
