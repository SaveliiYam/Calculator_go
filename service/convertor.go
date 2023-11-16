package service

import (
	"errors"
	"strings"
)

type convertor struct {
	first  string
	second string
}

func NewConvertor(first, second string) *convertor {
	return &convertor{
		first:  first,
		second: second,
	}
}

func (c *convertor) CheckNumbers() (bool, error) {
	arabValues := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	romeValues := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	firstCheck, secondCheck := false, false
	for _, value := range arabValues {
		if strings.Contains(c.first, value) {
			firstCheck = true
		}
		if strings.Contains(c.second, value) {
			secondCheck = true
		}
	}
	for _, value := range romeValues {
		if c.first == value {
			firstCheck = false
		}
		if c.second == value {
			secondCheck = false
		}
	}
	if firstCheck == secondCheck {
		return firstCheck, nil
	} else {
		return false, errors.New("invalid input values")
	}
}

func (c *convertor) FromRome(string) (string, error) {
	return "", nil
}
func (c *convertor) FromArabian(string) (string, error) {
	return "", nil
}
