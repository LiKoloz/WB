package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {

	// Создаем reader для чтения из стандартного ввода
	reader := bufio.NewReader(os.Stdin)

	// Читаем до символа новой строки
	input, _ := reader.ReadString('\n')

	// Удаляем символ новой строки
	str := strings.TrimSpace(input)

	a := strings.Split(str, " ")
	flags := a[1]

	var isK, isN, isR, isU, isM, isC, isH bool
	for _, f := range flags {
		switch f {
		case 'k':
			isK = true
		case 'n':
			isN = true
		case 'r':
			isR = true
		case 'u':
			isU = true
		case 'm':
			isM = true
		case 'c':
			isC = true
		case 'h':
			isH = true
		}
	}
	if isC {
		b := IsSorted(a)
		if b != nil {
			panic(b)
		} else {
			fmt.Println("sorted")
		}
	} else {
		if isK {
			a = a[3:]
			k := a[2]
			b, _ := strconv.Atoi(k)
			QuickSortByColumn(a, b)
		} else {
			a = a[2:]
		}

		if isN {
			QuickSortByHash(a)
		}
		if isU {
			a = TakeUnique(a)
		}
		if isM {
			a = SortByMonth(a)
		}
		if isH {
			a = SortByBites(a)
		}
		if isR {
			Reverse(a)
		}
	}
	fmt.Println(a)
}

// IsSorted - проверка отсортирован ли массив
func IsSorted(strs []string) error {
	var i = 0
	for i < len(strs)-1 {
		if !reflect.DeepEqual(strs[i], strs[i+1]) {
			return errors.New("Is not sorted")
		}
		i++
	}
	return nil
}

// QuickSortByColumn - сортировка по столбцу
func QuickSortByColumn(arr []string, n int) {
	if len(arr) <= 1 {
		return
	}
	rand.Seed(time.Now().UnixNano())

	stack := make([]int, 0, len(arr))
	stack = append(stack, 0, len(arr)-1)

	for len(stack) > 0 {
		high := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		low := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		pivotIdx := low + rand.Intn(high-low+1)
		arr[pivotIdx], arr[high] = arr[high], arr[pivotIdx] // перемещаем в конец
		pivot := arr[high]

		i := low
		for j := low; j < high; j++ {
			if arr[j][n] < pivot[n] {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
		}
		arr[i], arr[high] = arr[high], arr[i]
		pi := i

		if pi-1 > low {
			stack = append(stack, low, pi-1)
		}
		if pi+1 < high {
			stack = append(stack, pi+1, high)
		}
	}
}

// GetHashCode - получение хеша строки
func GetHashCode(str string) int {
	return len(str) % 3
}

// QuickSortByHash - быстрая сортировка по хэшу
func QuickSortByHash(arr []string) {
	rand.Seed(time.Now().UnixNano())

	stack := make([]int, 0, len(arr))
	stack = append(stack, 0, len(arr)-1)

	for len(stack) > 0 {
		high := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		low := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		pivotIdx := low + rand.Intn(high-low+1)
		arr[pivotIdx], arr[high] = arr[high], arr[pivotIdx] // перемещаем в конец
		pivot := arr[high]

		i := low
		for j := low; j < high; j++ {
			if GetHashCode(arr[j]) < GetHashCode(pivot) {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
		}
		arr[i], arr[high] = arr[high], arr[i]
		pi := i

		if pi-1 > low {
			stack = append(stack, low, pi-1)
		}
		if pi+1 < high {
			stack = append(stack, pi+1, high)
		}
	}
}

// TakeUnique - создание уникального сета
func TakeUnique(strs []string) []string {
	mapa := make(map[string]struct{})
	result := make([]string, 0)
	for _, v := range strs {
		if _, e := mapa[v]; !e {
			mapa[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Reverse - переворот слайса
func Reverse(strs []string) {
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
}

// SortByMonth - сортировка по месяцу
func SortByMonth(strs []string) []string {
	var (
		janMas    = make([]string, 0)
		febMas    = make([]string, 0)
		marchMas  = make([]string, 0)
		aprilMas  = make([]string, 0)
		mayMas    = make([]string, 0)
		juneMas   = make([]string, 0)
		julyMas   = make([]string, 0)
		augustMas = make([]string, 0)
		sepMas    = make([]string, 0)
		octMas    = make([]string, 0)
		novMas    = make([]string, 0)
		decMas    = make([]string, 0)

		janMasInt    = make([]int, 0)
		febMasInt    = make([]int, 0)
		marchMasInt  = make([]int, 0)
		aprilMasInt  = make([]int, 0)
		mayMasInt    = make([]int, 0)
		juneMasInt   = make([]int, 0)
		julyMasInt   = make([]int, 0)
		augustMasInt = make([]int, 0)
		sepMasInt    = make([]int, 0)
		octMasInt    = make([]int, 0)
		novMasInt    = make([]int, 0)
		decMasInt    = make([]int, 0)

		wg sync.WaitGroup
	)

	for _, v := range strs {
		if strings.Contains(v, "Jan") {
			janMas = append(janMas, strings.TrimRight(v, "Jan"))
		} else if strings.Contains(v, "Feb") {
			febMas = append(febMas, strings.TrimRight(v, "Feb"))
		} else if strings.Contains(v, "March") {
			marchMas = append(marchMas, strings.TrimRight(v, "March"))
		} else if strings.Contains(v, "April") {
			aprilMas = append(aprilMas, strings.TrimRight(v, "April"))
		} else if strings.Contains(v, "May") {
			mayMas = append(mayMas, strings.TrimRight(v, "May"))
		} else if strings.Contains(v, "June") {
			juneMas = append(juneMas, strings.TrimRight(v, "June"))
		} else if strings.Contains(v, "July") {
			julyMas = append(julyMas, strings.TrimRight(v, "July"))
		} else if strings.Contains(v, "August") {
			augustMas = append(augustMas, strings.TrimRight(v, "August"))
		} else if strings.Contains(v, "Sep") {
			sepMas = append(sepMas, strings.TrimRight(v, "Sep"))
		} else if strings.Contains(v, "Oct") {
			octMas = append(octMas, strings.TrimRight(v, "Oct"))
		} else if strings.Contains(v, "Nov") {
			novMas = append(novMas, strings.TrimRight(v, "Nov"))
		} else if strings.Contains(v, "Dec") {
			decMas = append(decMas, strings.TrimRight(v, "Dec"))
		}
	}
	wg.Add(12)

	go func() {
		for _, v := range janMas {
			a, _ := strconv.Atoi(v)
			janMasInt = append(janMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range febMas {
			a, _ := strconv.Atoi(v)
			febMasInt = append(febMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range marchMas {
			a, _ := strconv.Atoi(v)
			marchMasInt = append(marchMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range mayMas {
			a, _ := strconv.Atoi(v)
			mayMasInt = append(mayMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range juneMas {
			a, _ := strconv.Atoi(v)
			juneMasInt = append(juneMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range julyMas {
			a, _ := strconv.Atoi(v)
			julyMasInt = append(julyMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range augustMas {
			a, _ := strconv.Atoi(v)
			augustMasInt = append(augustMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range sepMas {
			a, _ := strconv.Atoi(v)
			sepMasInt = append(sepMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range octMas {
			a, _ := strconv.Atoi(v)
			octMasInt = append(octMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range novMas {
			a, _ := strconv.Atoi(v)
			novMasInt = append(novMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range decMas {
			a, _ := strconv.Atoi(v)
			decMasInt = append(decMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range aprilMas {
			a, _ := strconv.Atoi(v)
			aprilMasInt = append(aprilMasInt, a)
		}
		wg.Done()
	}()

	wg.Wait()

	wg.Add(12)

	go CountingSort(aprilMasInt, &wg)
	go CountingSort(augustMasInt, &wg)
	go CountingSort(decMasInt, &wg)
	go CountingSort(febMasInt, &wg)
	go CountingSort(janMasInt, &wg)
	go CountingSort(juneMasInt, &wg)
	go CountingSort(marchMasInt, &wg)
	go CountingSort(mayMasInt, &wg)
	go CountingSort(novMasInt, &wg)
	go CountingSort(octMasInt, &wg)
	go CountingSort(sepMasInt, &wg)

	wg.Wait()

	janMas = make([]string, 0)
	febMas = make([]string, 0)
	marchMas = make([]string, 0)
	aprilMas = make([]string, 0)
	mayMas = make([]string, 0)
	juneMas = make([]string, 0)
	julyMas = make([]string, 0)
	augustMas = make([]string, 0)
	sepMas = make([]string, 0)
	octMas = make([]string, 0)
	novMas = make([]string, 0)
	decMas = make([]string, 0)

	wg.Add(12)
	go func() {
		for _, v := range janMasInt {
			janMas = append(janMas, strconv.Itoa(v)+"Jan")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range febMasInt {
			febMas = append(febMas, strconv.Itoa(v)+"Feb")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range marchMasInt {
			marchMas = append(marchMas, strconv.Itoa(v)+"March")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range aprilMasInt {
			aprilMas = append(aprilMas, strconv.Itoa(v)+"April")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range mayMasInt {
			mayMas = append(mayMas, strconv.Itoa(v)+"May")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range juneMasInt {
			juneMas = append(juneMas, strconv.Itoa(v)+"June")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range julyMasInt {
			julyMas = append(julyMas, strconv.Itoa(v)+"July")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range augustMasInt {
			augustMas = append(augustMas, strconv.Itoa(v)+"August")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range sepMasInt {
			sepMas = append(sepMas, strconv.Itoa(v)+"Sep")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range octMasInt {
			octMas = append(octMas, strconv.Itoa(v)+"Oct")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range novMasInt {
			novMas = append(novMas, strconv.Itoa(v)+"Nov")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range decMasInt {
			decMas = append(decMas, strconv.Itoa(v)+"Dec")
		}
		wg.Done()
	}()

	wg.Wait()
	var result = make([]string, 0)

	for _, v := range janMas {
		result = append(result, v)
	}

	for _, v := range febMas {
		result = append(result, v)
	}

	for _, v := range marchMas {
		result = append(result, v)
	}

	for _, v := range aprilMas {
		result = append(result, v)
	}

	for _, v := range mayMas {
		result = append(result, v)
	}

	for _, v := range juneMas {
		result = append(result, v)
	}

	for _, v := range julyMas {
		result = append(result, v)
	}

	for _, v := range augustMas {
		result = append(result, v)
	}

	for _, v := range sepMas {
		result = append(result, v)
	}

	for _, v := range octMas {
		result = append(result, v)
	}

	for _, v := range novMas {
		result = append(result, v)
	}

	for _, v := range decMas {
		result = append(result, v)
	}

	return result
}

// CountingSort - сортировка подсчетом
func CountingSort(nums []int, wg *sync.WaitGroup) {
	var (
		mas   = make([]int, 32)
		slice = make([]int, 0)
	)

	for a := range mas {
		mas[a] = 0
	}

	for _, i := range nums {
		mas[i]++
	}

	for _, i := range mas {
		for range i {
			slice = append(slice, i)
		}
	}
	index := 0
	for value := 0; value < len(mas); value++ {
		for i := 0; i < mas[value]; i++ {
			nums[index] = value
			index++
		}
	}
	wg.Done()
}

// SortByBites - сортировка по битам, байтам и т.д.
func SortByBites(strs []string) []string {
	var (
		gbMas = make([]string, 0)
		mbMas = make([]string, 0)
		kbMas = make([]string, 0)
		BMas  = make([]string, 0)
		bMas  = make([]string, 0)

		gbMasInt = make([]int, 0)
		mbMasInt = make([]int, 0)
		kbMasInt = make([]int, 0)
		BMasInt  = make([]int, 0)
		bMasInt  = make([]int, 0)

		wg sync.WaitGroup
	)

	for _, v := range strs {
		if strings.Contains(v, "GB") {
			gbMas = append(gbMas, strings.TrimRight(v, "GB"))
		} else if strings.Contains(v, "M") {
			mbMas = append(mbMas, strings.TrimRight(v, "M"))
		} else if strings.Contains(v, "K") {
			kbMas = append(kbMas, strings.TrimRight(v, "K"))
		} else if strings.Contains(v, "B") {
			BMas = append(BMas, strings.TrimRight(v, "B"))
		} else if strings.Contains(v, "b") {
			bMas = append(bMas, strings.TrimRight(v, "b"))
		}
	}
	wg.Add(5)

	go func() {
		for _, v := range gbMas {
			a, _ := strconv.Atoi(v)
			gbMasInt = append(gbMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range mbMas {
			a, _ := strconv.Atoi(v)
			mbMasInt = append(mbMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range kbMas {
			a, _ := strconv.Atoi(v)
			kbMasInt = append(kbMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range BMas {
			a, _ := strconv.Atoi(v)
			BMasInt = append(BMasInt, a)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range bMas {
			a, _ := strconv.Atoi(v)
			bMasInt = append(bMasInt, a)
		}
		wg.Done()
	}()

	wg.Wait()

	wg.Add(5)

	go BubbleSort(gbMasInt, &wg)
	go BubbleSort(mbMasInt, &wg)
	go BubbleSort(kbMasInt, &wg)
	go BubbleSort(BMasInt, &wg)
	go BubbleSort(bMasInt, &wg)

	wg.Wait()

	gbMas = make([]string, 0)
	mbMas = make([]string, 0)
	kbMas = make([]string, 0)
	BMas = make([]string, 0)
	bMas = make([]string, 0)

	wg.Add(5)

	go func() {
		for _, v := range gbMasInt {
			gbMas = append(gbMas, strconv.Itoa(v)+"Gb")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range mbMasInt {
			mbMas = append(mbMas, strconv.Itoa(v)+"M")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range kbMasInt {
			kbMas = append(kbMas, strconv.Itoa(v)+"K")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range BMasInt {
			BMas = append(BMas, strconv.Itoa(v)+"B")
		}
		wg.Done()
	}()

	go func() {
		for _, v := range bMasInt {
			bMas = append(bMas, strconv.Itoa(v)+"b")
		}
		wg.Done()
	}()

	wg.Wait()

	var result = make([]string, 0)

	for _, v := range gbMas {
		result = append(result, v)
	}

	for _, v := range mbMas {
		result = append(result, v)
	}

	for _, v := range kbMas {
		result = append(result, v)
	}

	for _, v := range BMas {
		result = append(result, v)
	}

	for _, v := range bMas {
		result = append(result, v)
	}

	return result
}

// BubbleSort сортирует срез целых чисел методом пузырька.
func BubbleSort(nums []int, wg *sync.WaitGroup) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				// меняем местами
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		// если обменов не было — массив уже отсортирован

	}
	wg.Done()
}
