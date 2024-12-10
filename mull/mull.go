package mull

import (
	"bytes"
	"errors"
)

type Command struct {
	operation int
	op1       int
	op2       int
}

const (
	Mul = iota
	Do
	Dont
)

var UnsupportedOperationError = errors.New("unsupported operation")
var nullCommand = Command{-1, 0, 0}

func isDigit(input byte) bool {
	return input >= '0' && input <= '9'
}

func tryParseInt(input []byte) (int, []byte, bool) {
	if !isDigit(input[0]) {
		return -1, input, false
	}

	output := 0
	for isDigit(input[0]) {
		output *= 10
		output += int(input[0] - '0')
		input = input[1:]
	}

	return output, input, true
}

func tryParse(input []byte) (Command, []byte, bool) {
	if bytes.HasPrefix(input, []byte("mul(")) {
		op1, input, ok := tryParseInt(input[4:])
		if !ok {
			return nullCommand, input, false
		}

		if input[0] != ',' {
			return nullCommand, input, false
		}

		op2, input, ok := tryParseInt(input[1:])
		if !ok {
			return nullCommand, input, false
		}

		if input[0] != ')' {
			return nullCommand, input, false
		}

		return Command{Mul, op1, op2}, input[1:], true
	} else if bytes.HasPrefix(input, []byte("do()")) {
		return Command{Do, 0, 0}, input[4:], true
	} else if bytes.HasPrefix(input, []byte("don't()")) {
		return Command{Dont, 0, 0}, input[7:], true
	}

	return nullCommand, input, false
}

func Parse(input []byte) []Command {
	output := make([]Command, 0)

	for len(input) > 0 {
		if input[0] == 'm' || input[0] == 'd' {
			command, rest, parsed := tryParse(input)
			if parsed {
				input = rest
				output = append(output, command)
			} else {
				input = input[1:]
			}
		} else {
			input = input[1:]
		}
	}

	return output
}

func RunCommands(commands []Command, enableConditionals bool) ([]int, error) {
	output := make([]int, 0, len(commands))

	enabled := true
	for _, c := range commands {
		switch c.operation {
		case Mul:
			if enabled {
				output = append(output, c.op1*c.op2)
			}
		case Do:
			if enableConditionals {
				enabled = true
			}
		case Dont:
			if enableConditionals {
				enabled = false
			}
		default:
			return nil, UnsupportedOperationError
		}
	}

	return output, nil
}
