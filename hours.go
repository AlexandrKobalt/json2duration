package json2duration

import (
	"encoding/json"
	"time"
)

type Hours struct {
	time.Duration
}

func (d *Hours) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Hours())
}

func (d *Hours) UnmarshalJSON(b []byte) error {
	var Hours int64
	if err := json.Unmarshal(b, &Hours); err != nil {
		return err
	}

	d.Duration = time.Duration(Hours * int64(time.Hour))

	return nil
}
