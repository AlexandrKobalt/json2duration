package json2duration

import (
	"encoding/json"
	"time"
)

type Milliseconds struct {
	time.Duration
}

func (d *Milliseconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Milliseconds())
}

func (d *Milliseconds) UnmarshalJSON(b []byte) error {
	var Milliseconds int64
	if err := json.Unmarshal(b, &Milliseconds); err != nil {
		return err
	}

	d.Duration = time.Duration(Milliseconds * int64(time.Millisecond))

	return nil
}
