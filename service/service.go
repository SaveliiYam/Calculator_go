package service

type Calculator interface {
	GetAnswer() (string, error)
	UpdateValues(first, second string)
}

type Convertor interface {
	CheckNumbers() (bool, error)
	FromRome(first string) (string, error)
	FromArabian() (string, string, error)
}

type Service struct {
	Calculator
	Convertor
}

func NewService(first, second, operator string) *Service {
	return &Service{
		Calculator: NewCalculator(first, second, operator),
		Convertor:  NewConvertor(first, second),
	}
}
