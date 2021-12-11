package common

import "os"

var WORK_DIR string

func InitPaths() {
	var err error
	WORK_DIR, err = os.Getwd()

	if err != nil {
		panic(err)
	}
}
