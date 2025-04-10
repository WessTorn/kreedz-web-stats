package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

type logWriter struct {
	logPath string
}

func (w logWriter) Write(b []byte) (int, error) {
	dt := fmt.Sprintf(time.Now().UTC().Format("2006-01-02"))
	file, err := os.OpenFile(w.logPath+dt+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	wrt := io.MultiWriter(os.Stdout, file)
	defer file.Close()
	b = bytes.Replace(b, []byte(".go:"), []byte{':'}, -1)
	return wrt.Write(b)
}

func InitLogger(basePath string) error {
	wrt := new(logWriter)
	wrt.logPath = basePath
	err := os.MkdirAll(wrt.logPath, os.ModePerm)
	if err != nil {
		return err
	}
	InfoLogger = log.New(wrt, "INFO: ", log.Ltime|log.Lshortfile)
	WarningLogger = log.New(wrt, "WARNING: ", log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(wrt, "ERROR: ", log.Ltime|log.Lshortfile)
	return nil
}
