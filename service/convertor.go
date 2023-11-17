package service

import (
	"errors"
	"strconv"
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
	arabValues := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	romeValues := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	var exist1 bool = false
	var exist2 bool = false
	firstCheck, secondCheck := false, false
	for _, value := range arabValues {
		if strings.Contains(c.first, value) {
			firstCheck = true
			exist1 = true
		}
		if strings.Contains(c.second, value) {
			secondCheck = true
			exist2 = true
		}
	}
	for _, value := range romeValues {
		if c.first == value {
			firstCheck = false
			exist1 = true
		}
		if c.second == value {
			secondCheck = false
			exist2 = true
		}
	}
	if !exist1 || !exist2 {
		return false, errors.New("invalid input values")
	}
	if firstCheck == secondCheck {
		return firstCheck, nil
	} else {
		return false, errors.New("invalid input values")
	}
}

func (c *convertor) FromRome(first string) (string, error) {
	var table = [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	}
	firstVal, _ := strconv.Atoi(first)
	if firstVal < 0 {
		return "", errors.New("negative value")
	}
	var (
		answer = ""
		digit  = 100
	)
	for i := 2; i >= 0; i-- {
		d := firstVal / digit
		answer += table[i][d]
		firstVal %= digit
		digit /= 10
	}
	return answer, nil
}
func (c *convertor) FromArabian() (string, string, error) {
	arabValues := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	romeValues := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := range romeValues {
		if c.first == arabValues[i] {
			return "", "", errors.New("already arabian")
		}
		if c.second == arabValues[i] {
			return "", "", errors.New("already arabian")
		}
		if c.first == romeValues[i] {
			c.first = arabValues[i]
		}
		if c.second == romeValues[i] {
			c.second = arabValues[i]
		}
	}
	return c.first, c.second, nil
}
