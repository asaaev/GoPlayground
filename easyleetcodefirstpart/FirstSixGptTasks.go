package easyleetcodefirstpart

import "fmt"

const (
	t = 21
	s = "dfsf"
)

var (
	g float64 = 12.3
	x int     = 43
)

func task1() {
	var b bool
	fmt.Println(x, t, g, s, b)
	fmt.Println(x + t)
}
func task2and3() {

	var r float64 = float64(t) + g
	fmt.Println(r)

	var h string = "rock"
	o := s + h
	fmt.Println(o)
}
func task4and5() {
	const con = 8
	y := con + t
	fmt.Println(y)

	if t >= x {
		fmt.Println("t true")
	} else {
		fmt.Println("x true")
	}
}
func task6() {

	var arr = [3]int{2, 4}
	arr[2] = 6
	sar := arr[0] + arr[1] + arr[2]
	fmt.Println(arr, sar)
}

//func main() {
//
//}
