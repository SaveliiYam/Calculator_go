package service

import (
	"errors"
	"strconv"
)

type calculator struct {
	first    string
	second   string
	operator string
}

func NewCalculator(first, second, operator string) *calculator {
	return &calculator{
		first:    first,
		second:   second,
		operator: operator,
	}
}

func (c *calculator) UpdateValues(first, second string) {
	c.first = first
	c.second = second
}

func (c *calculator) GetAnswer() (string, error) {
	switch c.operator {
	case "+":
		f, _ := strconv.Atoi(c.first)
		s, _ := strconv.Atoi(c.second)
		return strconv.Itoa(f + s), nil
	case "-":
		f, _ := strconv.Atoi(c.first)
		s, _ := strconv.Atoi(c.second)
		return strconv.Itoa(f - s), nil
	case "*":
		f, _ := strconv.Atoi(c.first)
		s, _ := strconv.Atoi(c.second)
		return strconv.Itoa(f * s), nil
	case "/":
		f, _ := strconv.Atoi(c.first)
		s, _ := strconv.Atoi(c.second)
		if s == 0 {
			return "", errors.New("try 0")
		}
		return strconv.Itoa(f / s), nil
	}
	return "", errors.New("invalid operator")
}
