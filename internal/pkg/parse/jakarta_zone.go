package parse

import "time"

func ParseToJakartaZone(timestamp string) (*time.Time, error) {
	jakartaZone := "Asia/Jakarta"
	jakartaOffset := 7 * 60 * 60

	parseToJakarta, err := time.ParseInLocation("2006-01-02 15:04:05", timestamp, time.FixedZone(jakartaZone, jakartaOffset))
	if err != nil {
		return nil, err
	}

	return &parseToJakarta, nil
}