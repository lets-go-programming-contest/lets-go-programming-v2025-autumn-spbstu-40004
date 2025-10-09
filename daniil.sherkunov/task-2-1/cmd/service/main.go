package main

import (
	"bufio"
	"errors"
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

var (
	errInvalidLineFormat = errors.New("invalid constraint line format")
	errInvalidOperator   = errors.New("invalid operator")
)

// readInt reads one line and parses int; returns (value, ok).
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

func applyConstraint(lowerBound, upperBound int, operatorToken string, operatorValue int) (int, int, error) {
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
		return lowerBound, upperBound, errInvalidOperator
	}

	return lowerBound, upperBound, nil
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
		employeesCount, readOK := readInt(scanner)
		if !readOK || employeesCount < 1 {
			return
		}

		lowerBound := defaultMinTemperature
		upperBound := defaultMaxTemperature

		for employeeIndex := range employeesCount {
			if !scanner.Scan() {
				return
			}

			line := strings.TrimSpace(scanner.Text())
			fields := strings.Fields(line)
			if len(fields) != expectedFieldsPerLine {
				return
			}

			operatorToken := fields[0]
			operatorValue, convErr := strconv.Atoi(fields[1])
			if convErr != nil {
				return
			}

			var applyErr error

			lowerBound, upperBound, applyErr = applyConstraint(lowerBound, upperBound, operatorToken, operatorValue)
			if applyErr != nil {
				return
			}

			var printErr error

			if lowerBound > upperBound {
				_, printErr = fmt.Fprintln(writer, -1)
			} else {
				_, printErr = fmt.Fprintln(writer, lowerBound)
			}

			if printErr != nil {
				return
			}

			_ = employeeIndex
		}

		_ = departmentIndex
	}
