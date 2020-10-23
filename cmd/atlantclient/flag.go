package main

import (
	"regexp"

	"github.com/pkg/errors"
)

var (
	ErrEmptyHostFlag    = errors.New("empty host flag")
	ErrInvalidHostValue = errors.New("invalid host value")

	ErrEmptyURLFlag    = errors.New("empty URL flag")
	ErrInvalidURLValue = errors.New("invalid URL value")

	ErrInvalidStartParameterValue = errors.New(`"start" value should be greater or equal zero`)

	ErrInvalidLimitParameterMinValue = errors.New(`"limit" value should be greater or equal 1`)
	ErrInvalidLimitParameterMaxValue = errors.New(`"limit" value should be less or equal 100`)

	ErrInvalidSortingParamter = errors.New("invalid sorting parameter")
)

type Flag interface {
	Validate() (err error)
}

type HostFlag string

func (f HostFlag) String() (s string) {
	return (string)(f)
}

func (f *HostFlag) Pointer() (p *string) {
	return (*string)(f)
}

func (f HostFlag) Validate() (err error) {
	if f.String() == "" {
		return ErrEmptyHostFlag
	}

	r := regexp.MustCompile(`^(((\d{1,3}\.){3}(\d{1,3}))|((\w+\.){2}(\w+))):(\d{4,5})$`)

	if !r.MatchString(f.String()) {
		return errors.New(ErrInvalidHostValue.Error() + ": " + f.String())
	}

	return nil
}

type URLFlag string

func (f URLFlag) String() (s string) {
	return (string)(f)
}

func (f *URLFlag) Pointer() (p *string) {
	return (*string)(f)
}

func (f URLFlag) Validate() (err error) {
	if f.String() == "" {
		return ErrEmptyURLFlag
	}

	r := regexp.MustCompile(`(http|https)://[\w\-_]+(\.[\w\-_]+)+([\w\-.,@?^=%&amp;:/~+#]*[\w\-@?^=%&amp;/~+#])?`)

	if !r.MatchString(f.String()) {
		return errors.New(ErrInvalidURLValue.Error() + ": " + f.String())
	}

	return nil
}

const MinStartParameterValue int64 = 0

type StartFlag int64

func (f StartFlag) Int64() (val int64) {
	return int64(f)
}

func (f *StartFlag) Pointer() (p *int64) {
	return (*int64)(f)
}

func (f StartFlag) Validate() (err error) {
	if f.Int64() < MinStartParameterValue {
		return ErrInvalidStartParameterValue
	}

	return nil
}

const (
	MinLimitParameterValue int64 = 1
	MaxLimitParameterValue int64 = 100
)

type LimitFlag int64

func (f LimitFlag) Int64() (val int64) {
	return int64(f)
}

func (f *LimitFlag) Pointer() (p *int64) {
	return (*int64)(f)
}

func (f LimitFlag) Validate() (err error) {
	if f.Int64() < MinLimitParameterValue {
		return ErrInvalidLimitParameterMinValue
	}

	if f.Int64() > MaxLimitParameterValue {
		return ErrInvalidLimitParameterMaxValue
	}

	return nil
}

type SortFlag []string

func (f SortFlag) StringArray() (ss []string) {
	return f
}

func (f *SortFlag) Pointer() (ss *[]string) {
	return (*[]string)(f)
}

func (f SortFlag) Validate() (err error) {
	r := regexp.MustCompile(`^([\w_]+):(desc|asc)$`)

	for _, s := range f.StringArray() {
		if !r.MatchString(s) {
			return errors.WithMessage(ErrInvalidSortingParamter, s)
		}
	}

	return nil
}
