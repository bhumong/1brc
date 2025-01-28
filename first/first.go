package first

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CalculateAll(path string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic("os.OpenFile:" + err.Error())
	}
	scanner := bufio.NewScanner(file)
	min := 9_999_999.0
	max := -9_999_999.0
	sum := 0.0
	total := 0.0
	for scanner.Scan() {
		total++
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) != 2 {
			continue
		}
		temp, _ := strconv.ParseFloat(parts[1], 32)
		sum += temp
		if temp < min {
			min = temp
		}
		if temp > max {
			max = temp
		}
	}
	mean := sum / float64(total)

	println("Min: " + strconv.FormatFloat(min, 'f', 2, 32))
	println("Max: " + strconv.FormatFloat(max, 'f', 2, 32))
	println("Mean: " + strconv.FormatFloat(mean, 'f', 2, 32))
}
