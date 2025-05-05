package cmd

import "time"

func EpochToTime(epoch int64) time.Time {
	return time.Unix(epoch, 0).UTC()
}

func Forward(basis *time.Time, duration_str string) (*time.Time, error) {

	if d, d_err := time.ParseDuration(duration_str); d_err == nil {

		r := basis.Add(d * -1)
		return &r, nil
	} else {
		return nil, d_err
	}
}
