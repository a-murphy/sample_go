package sample

import "github.com/a-murphy/sample_go/newmath"

var result float64

func Setup() {
	result = newmath.Sqrt(4)
}

func GetResult() float64 {
	return result
}
