package json2duration

import (
	"encoding/json"
	"time"
)

type Seconds struct {
	time.Duration
}

func (d *Seconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Seconds())
}

func (d *Seconds) UnmarshalJSON(b []byte) error {
	var Seconds int64
	if err := json.Unmarshal(b, &Seconds); err != nil {
		return err
	}

	d.Duration = time.Duration(Seconds * int64(time.Second))

	return nil
}
