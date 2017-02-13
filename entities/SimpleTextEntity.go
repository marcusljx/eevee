package entities

type SimpleTextEntity struct {
	data []rune
}

func NewSimpleTextEntity(s string) *SimpleTextEntity {
	return &SimpleTextEntity{
		data: []rune(s),
	}
}

func (s *SimpleTextEntity) RuneArray() []rune {
	return s.data
}

func (s *SimpleTextEntity) Parse(r []rune) {
	s.data = r
}

func (s *SimpleTextEntity) String() string {
	return string(s.data)
}

func (s *SimpleTextEntity) Score() (sum float64) {
	for _, r := range s.data {
		sum += float64(r)
	}
	return
}
