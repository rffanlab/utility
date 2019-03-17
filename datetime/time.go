package datetime

import "time"

// "2006-01-02 15:04:05"
// "20060102150405"
func TimeNowForSecondFormated(format string) string {
	return time.Now().Format(format)
}
