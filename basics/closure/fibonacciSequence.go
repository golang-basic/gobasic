package main

var f0, f1 int = 0, 1

func fibonacciNumber(n int) (value int) {
	switch {
	case n==0:
		value = f0
	case n==1:
		value = f1
	case n!=0 && n!=1:
		value = fibonacciNumber(n-1)+fibonacciNumber(n-2)
	}
	return
}

func GenerateFibonacciSeries(n int) (ser []int) {
	for i := 0; i <= n; i++ {
		ser = append(ser, fibonacciNumber(i))
	}
	return
}
