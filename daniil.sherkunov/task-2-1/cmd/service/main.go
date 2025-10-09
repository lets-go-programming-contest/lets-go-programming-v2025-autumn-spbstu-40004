package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	defaultMinTemperature   = 15
	defaultMaxTemperature   = 30
	scannerInitBufSizeBytes = 1 << 16
	scannerMaxBufSizeBytes  = 1 << 20
	expectedFieldsPerLine   = 2
)

func readInt(scanner *bufio.Scanner) (int, bool) {
	if !scanner.Scan() {
		return 0, false
	}

	text := strings.TrimSpace(scanner.Text())
	value, convErr := strconv.Atoi(text)

	if convErr != nil {
		return 0, false
	}

	return value, true
}

func readConstraint(scanner *bufio.Scanner) (string, int, bool) {
	if !scanner.Scan() {
		return "", 0, false
	}

	line := strings.TrimSpace(scanner.Text())
	fields := strings.Fields(line)

	if len(fields) != expectedFieldsPerLine {
		return "", 0, false
	}

	operatorToken := fields[0]
	operatorValue, convErr := strconv.Atoi(fields[1])

	if convErr != nil {
		return "", 0, false
	}

	return operatorToken, operatorValue, true
}

func applyConstraint(lowerBound, upperBound int, operatorToken string, operatorValue int) (int, int, bool) {
	switch operatorToken {
	case ">=", "≥":
		if operatorValue > lowerBound {
			lowerBound = operatorValue
		}
	case "<=", "≤":
		if operatorValue < upperBound {
			upperBound = operatorValue
		}
	default:
		return lowerBound, upperBound, false
	}

	return lowerBound, upperBound, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, scannerInitBufSizeBytes), scannerMaxBufSizeBytes)

	departmentsCount, ok := readInt(scanner)
	if !ok || departmentsCount < 1 {
		return
	}

	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		if flushErr := writer.Flush(); flushErr != nil {
			_ = flushErr
		}
	}()

	for departmentIndex := range departmentsCount {
		_ = departmentIndex

		employeesCount, readOK := readInt(scanner)
		if !readOK || employeesCount < 1 {
			return
		}

		lowerBound := defaultMinTemperature
		upperBound := defaultMaxTemperature

		for employeeIndex := range employeesCount {
			_ = employeeIndex

			operatorToken, operatorValue, okConstraint := readConstraint(scanner)
			if !okConstraint {
				return
			}

			var okApply bool
			lowerBound, upperBound, okApply = applyConstraint(lowerBound, upperBound, operatorToken, operatorValue)

			if !okApply {
				return
			}

			var out int
			if lowerBound > upperBound {
				out = -1
			} else {
				out = lowerBound
			}

			if _, err := fmt.Fprintln(writer, out); err != nil {
				return
			}
		}
	}
}
