package error

import "log"

func CheckError(err error) bool {
	if err != nil {
		log.Println("error occurs:", err)
		return true
	}
	return false
}
