package exceptions

import "log"

func CheckMainException(err error) {
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
