package util

import (
	"os"
	"time"
)

const dateFormt = "2006-01-02"

func FileExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func DateFormat(layout, date string) (string, error) {
	t, err := time.Parse(layout, date)
	if err != nil {
		return "", err
	}
	return t.Format(dateFormt), nil
}

func Today() string {
	return time.Now().Format(dateFormt)
}
