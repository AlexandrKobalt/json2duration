package json2duration

import (
	"encoding/json"
	"time"
)

type Minutes struct {
	time.Duration
}

func (d *Minutes) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Minutes())
}

func (d *Minutes) UnmarshalJSON(b []byte) error {
	var Minutes int64
	if err := json.Unmarshal(b, &Minutes); err != nil {
		return err
	}

	d.Duration = time.Duration(Minutes * int64(time.Minute))

	return nil
}
