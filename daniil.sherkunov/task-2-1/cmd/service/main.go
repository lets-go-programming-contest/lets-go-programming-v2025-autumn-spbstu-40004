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

var errInvalidOperator = errors.New("invalid operator")

type TemperatureController struct {
	lower int
	upper int
}

func NewTemperatureController(minT, maxT int) *TemperatureController {
	return &TemperatureController{lower: minT, upper: maxT}
}

func (t *TemperatureController) Apply(operatorToken string, operatorValue int) error {
	switch operatorToken {
	case ">=", "≥":
		if operatorValue > t.lower {
			t.lower = operatorValue
		}
	case "<=", "≤":
		if operatorValue < t.upper {
			t.upper = operatorValue
		}
	default:
		return errInvalidOperator
	}
	return nil
}

func (t *TemperatureController) Current() (int, bool) {
	if t.lower > t.upper {
		return 0, false
	}
	return t.lower, true
}

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

		ctrl := NewTemperatureController(defaultMinTemperature, defaultMaxTemperature)

		for employeeIndex := range employeesCount {
			_ = employeeIndex

			operatorToken, operatorValue, okConstraint := readConstraint(scanner)
			if !okConstraint {
				return
			}

			applyErr := ctrl.Apply(operatorToken, operatorValue)
			if applyErr != nil {
				return
			}

			value, okCurrent := ctrl.Current()
			out := -1
			if okCurrent {
				out = value
			}

			if _, err := fmt.Fprintln(writer, out); err != nil {
				return
			}
		}
	}
}
