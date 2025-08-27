package main

import "fmt"

func main() {
	mas := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	t10 := []float32{}
	t20 := []float32{}
	t30 := []float32{}

	for _, v := range mas {
		del := int(v / 10)
		if del < 0 {
			del *= -1
		}
		fmt.Println(del)
		if del == 1 {
			t10 = append(t10, v)
		} else if del == 2 {
			t20 = append(t20, v)
		} else {
			t30 = append(t30, v)
		}
	}
	fmt.Println(t10)
	fmt.Println(t20)
	fmt.Println(t30)
}
