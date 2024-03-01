package main

import (
	"enemdoor/parsers"
	"time"
)

type TimeRule struct {
	from time.Duration
	to   time.Duration
}

func (tr *TimeRule) IsOpenToEnter(t string) error {
	tT, err := parsers.ParseHour(t)
	if err != nil {
		return err
	}

	if tT < tr.from {
		return ErrNotOpenedYet
	}

	if tT >= tr.to {
		return ErrClosedAlready
	}

	return nil
}

func NewTimeRule(from string, to string) (*TimeRule, error) {
	fromT, err := parsers.ParseHour(from)
	if err != nil {
		return nil, err
	}

	toT, err := parsers.ParseHour(to)
	if err != nil {
		return nil, err
	}

	return &TimeRule{
		from: fromT,
		to:   toT,
	}, nil
}
