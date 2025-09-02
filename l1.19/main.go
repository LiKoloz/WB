package main

func main() {
	str := "главрыба"
	println(str)
	println(reverseStr(str))

}

func reverseStr(str string) string {
	m := []rune(str)

	for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {

		m[i], m[j] = m[j], m[i]
	}

	return string(m)
}
