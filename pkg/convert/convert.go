package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) UInt() (uint, error) {
	v, err := strconv.Atoi(s.String())
	return uint(v), err
}

func (s StrTo) MustUInt() uint {
	v, _ := s.UInt()
	return v
}

func (s StrTo) Bool() (bool, error) {
	v, err := strconv.ParseBool(s.String())
	return v, err
}

func (s StrTo) MustBool() bool {
	v, _ := s.Bool()
	return v
}
