package fib

var cache = make(map[uint64]uint64)

func Fib(n uint64) uint64 {
	if n < 2 {
		return n
	}

	if _, ok := cache[n]; !ok {
		cache[n] = Fib(n-1) + Fib(n-2)
	}
	return cache[n]
}
