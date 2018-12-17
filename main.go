package main

// Unbound Buffer Channel
// Concurrency in go build on three interdependent features: gorountine, channel and select statement
// channel come in two flavors: buffered and unbuffered
// How do I create a writer that never blocks on a write to a channel?

// you can not read from write channel
// you can not write to read channel
// you can not close write channel

// MakeInfinite

func MakeInfinite() (chan<- interface{}, <-chan interface{}) {
	in := make(chan interface{})
	out := make(chan interface{})

	go func() {
		var inQueue []interface{}
		curVal := func() interface{} {
			if len(inQueue) == 0 {
				return nil
			}
			return inQueue[0]
		}
		outCh := func() chan interface{} {
			if len(inQueue) == 0 {
				return nil
			}
			return out
		}
		for in != nil || len(inQueue) > 0 {
			select {
			case i, ok := <-in:
				if !ok {
					in = nil
				} else {
					inQueue = append(inQueue, i)
				}
			case outCh() <- curVal():
				inQueue = inQueue[1:]
			}
		}
		close(out)
	}()
	return in, out
}
func main() {

}
