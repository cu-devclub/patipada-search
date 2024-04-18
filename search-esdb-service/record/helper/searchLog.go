package helper

import "strconv"

func SearchLogsMessage(timeStamp string, query string, offset int, amount int) string {
	return timeStamp + "&Q=" + query + "&Offset=" + strconv.Itoa(offset) + "&Amount=" + strconv.Itoa(amount)
}
