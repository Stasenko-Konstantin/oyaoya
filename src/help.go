package src

import (
	"errors"
	"log"
	"os"
)

type myLogger struct {
	file   *os.File
	stdout *os.File
}

func (ml myLogger) Write(b []byte) (n int, err error) {
	var myerr string
	n, err = ml.file.Write(b)
	if err != nil {
		myerr += err.Error() + "\n"
	}
	_, err = ml.stdout.Write(b)
	if err != nil {
		myerr += err.Error()
	}
	if myerr != "" {
		err = errors.New(myerr)
	}
	return n, err
}

func newLogger() *log.Logger {
	if _, err := os.Stat("log.txt"); os.IsNotExist(err) {
		f, err := os.Create("log.txt")
		if err != nil {
			panic("Не удалось создать логер!")
		}
		return log.New(myLogger{f, os.Stdout}, "INFO\t", log.Ltime)
	}
	f, err := os.Open("log.txt")
	if err != nil {
		panic("Не удалось открыть логер!")
	}
	return log.New(myLogger{f, os.Stdout}, "INFO\t", log.Ltime)
}
