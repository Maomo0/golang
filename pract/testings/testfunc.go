package main

func loopfunc(n, m int) int {
	if n == 1 {
		return 0
	}
	x := loopfunc(n-1, m)
	return (m+x) % n
}

func forfunc(n, m int) int {
	res := 0
	for i := 2; i <= n; i++ {
		res = (res+m) % i
	}
	return res
}