package main

func someFunc() string {
	v := createHugeString(1 << 10)
	if len(v) < 100 {
		return ""
	}
	runes := []rune{v}

	return string(runes[:100])
}
