package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// parseFields -  преобразует строку в отсортированный список номеров полей.
func parseFields(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	fieldSet := make(map[int]bool)

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if strings.Contains(part, "-") {
			bounds := strings.Split(part, "-")
			if len(bounds) != 2 {
				return nil, fmt.Errorf("invalid range format: %s", part)
			}
			start, err1 := strconv.Atoi(strings.TrimSpace(bounds[0]))
			end, err2 := strconv.Atoi(strings.TrimSpace(bounds[1]))
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("non-numeric range: %s", part)
			}
			if start > end {
				return nil, fmt.Errorf("range start > end: %s", part)
			}
			for i := start; i <= end; i++ {
				fieldSet[i] = true
			}
		} else {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", part)
			}
			fieldSet[num] = true
		}
	}

	if len(fieldSet) == 0 {
		return nil, fmt.Errorf("no fields specified")
	}

	fields := make([]int, 0, len(fieldSet))
	for f := range fieldSet {
		fields = append(fields, f)
	}
	sort.Ints(fields)
	return fields, nil
}

func main() {
	var fieldsStr string
	var delimiter string
	var separated bool

	flag.StringVar(&fieldsStr, "f", "", "Fields to select (e.g., 1,3-5)")
	flag.StringVar(&delimiter, "d", "\t", "Delimiter (default tab)")
	flag.BoolVar(&separated, "s", false, "Only print lines containing delimiter")
	flag.Parse()

	if fieldsStr == "" {
		fmt.Fprintln(os.Stderr, "Error: -f flag is required")
		os.Exit(1)
	}

	fields, err := parseFields(fieldsStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing fields: %v\n", err)
		os.Exit(1)
	}
	for _, f := range fields {
		if f < 1 {
			fmt.Fprintln(os.Stderr, "Error: field numbers must be >= 1")
			os.Exit(1)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, delimiter)

		if separated && len(parts) == 1 {
			continue
		}

		selected := make([]string, 0, len(fields))
		for _, f := range fields {
			if f <= len(parts) {
				selected = append(selected, parts[f-1])
			}
		}

		if len(selected) > 0 {
			fmt.Println(strings.Join(selected, delimiter))
		} else {
			fmt.Println()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
