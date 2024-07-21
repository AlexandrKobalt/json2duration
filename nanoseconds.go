package json2duration

import (
	"encoding/json"
	"time"
)

type Nanoseconds struct {
	time.Duration
}

func (d *Nanoseconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Nanoseconds())
}

func (d *Nanoseconds) UnmarshalJSON(b []byte) error {
	var Nanoseconds int64
	if err := json.Unmarshal(b, &Nanoseconds); err != nil {
		return err
	}

	d.Duration = time.Duration(Nanoseconds * int64(time.Nanosecond))

	return nil
}
