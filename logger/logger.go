package logger

import "log"

func LogHttpMethodErr(r string, m string) {
	log.Println("http request error:", r, m, "Method Not Found.")
}

func LogHttpWriteErr(e error) {
	log.Println(e)
}

func LogJsonSerialize(e error) {
	log.Println(e)
}

func LogRun(s string) {
	log.Println(s)
}
