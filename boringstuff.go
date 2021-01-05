// Package boringstuff wraps some boring tasks.
package boringstuff

import (
	"strconv"
	"time"
)

// EpochToUnix converts epoch time to Unix time.
func EpochToUnix(epoch string) (time.Time, error) {
	var t time.Time
	i, err := strconv.ParseInt(epoch, 10, 64)
	if err != nil {
		return t, err
	}
	t = time.Unix(i, 0)

	return t, nil
}

// ReturnError checks for multiple errors at once.
func ReturnError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
