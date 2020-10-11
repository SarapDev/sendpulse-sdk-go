package types

import (
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		d.Time = time.Time{}
		return nil
	}
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}
