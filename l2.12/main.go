package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Создаем reader для чтения из стандартного ввода
	reader := bufio.NewReader(os.Stdin)

	// Читаем до символа новой строки
	input, _ := reader.ReadString('\n')

	// Удаляем символ новой строки
	str := strings.TrimSpace(input)

	a := strings.Split(str, " ")
	mapa := make(map[string]int)
	regular := a[len(a)-1]
	index := len(a)
	for i := 0; i < len(a); i++ {
		if a[i] == "-A" {
			j, _ := strconv.Atoi(a[i+1])
			mapa["-A"] = j
			if index > i {
				index = i
			}
			i += 1
		}
		if a[i] == "-B" {
			j, _ := strconv.Atoi(a[i+1])
			mapa["-B"] = j
			if index > i {
				index = i
			}
			i += 1
		}
		if a[i] == "-C" {
			j, _ := strconv.Atoi(a[i+1])
			mapa["-C"] = j
			if index > i {
				index = i
			}
			i += 1
		}
		if a[i] == "-c" {
			mapa["-c"] = 0
			if index > i {
				index = i
			}
		}
		if a[i] == "-i" {
			mapa["-i"] = 0
			if index > i {
				index = i
			}
		}
		if a[i] == "-v" {
			mapa["-v"] = 0
			if index > i {
				index = i
			}
		}
		if a[i] == "-F" {
			mapa["-F"] = 0
			if index > i {
				index = i
			}
		}
		if a[i] == "-n" {
			mapa["-n"] = 0
			if index > i {
				index = i
			}
		}
	}
	fmt.Println(mapa)
	fmt.Print(index)

	a = a[:index]
	b := a
	startIndex := 0
	lastIndex := 1

	if v, e := mapa["-A"]; e {
		startIndex = v
	}
	if v, e := mapa["-B"]; e {
		lastIndex = v
	}
	if v, e := mapa["-C"]; e {
		startIndex = v
		lastIndex = v
	}
	if _, e := mapa["-i"]; e {
		regular = strings.ToLower(regular)
		for i := 0; i < len(b); i++ {
			b[i] = strings.ToLower(b[i])
		}
	}
	resultNotFinal := make([]string, 0)

	if _, e := mapa["-F"]; e {
		for j, v := range b {
			if strings.Contains(v, regular) {
				c := a[j+startIndex : j+lastIndex]
				for _, value := range c {
					resultNotFinal = append(resultNotFinal, value)
				}
			}
		}
	} else {
		for j, v := range b {
			matched, _ := regexp.MatchString(regular, v)
			if matched {
				c := a[j+startIndex : j+lastIndex]
				for _, value := range c {
					resultNotFinal = append(resultNotFinal, value)
				}
			}
		}
	}

	result := make([]string, 0)

	if _, e := mapa["-v"]; e {
		for _, v1 := range b {
			flag := false
			for _, v2 := range resultNotFinal {
				if v2 == v1 {
					flag = true
				}
			}
			if !flag {
				result = append(result, v1)
			}
		}
	} else {
		result = resultNotFinal
	}
	resultFinal := make([]string, 0)

	if _, e := mapa["-c"]; e {
		for i1, v1 := range b {
			for _, v2 := range result {
				if v1 == v2 {
					resultFinal = append(resultFinal, strconv.Itoa(i1))
				}
			}
		}
	} else {
		resultFinal = result
	}

	if _, e := mapa["-n"]; e {
		for i := 0; i < len(result); i++ {
			for i1, v1 := range b {
				if v1 == result[i] {
					result[i] = strconv.Itoa(i1) + " " + result[i]
				}
			}
		}
	}
	fmt.Println(resultFinal)
}

// removeRegexMeta - удаление регулярных символов
func removeRegexMeta(s string) string {
	const meta = `\.+*?()|[]{}^$`
	var builder strings.Builder
	for _, r := range s {
		if !strings.ContainsRune(meta, r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
