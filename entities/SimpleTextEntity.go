package entities

type SimpleTextEntity struct {
	data string
}

func NewSimpleTextEntity(s string) *SimpleTextEntity {
	return &SimpleTextEntity{
		data: s,
	}
}

func (s *SimpleTextEntity) ParseRuneArray(rArr []rune) {
	s.data = string(rArr)
}

func (s *SimpleTextEntity) AsRuneArray() []rune {
	return []rune(s.data)
}

func (s *SimpleTextEntity) String() string {
	return s.data
}

func (s *SimpleTextEntity) Score() (sum float64) {
	for _, r := range s.AsRuneArray() {
		sum += float64(r)
	}
	return
}
