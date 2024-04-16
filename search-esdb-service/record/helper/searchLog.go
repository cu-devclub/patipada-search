package helper

func SearchLogsMessage(timeStamp string, query string) string {
	return timeStamp + "Q=" + query
}
