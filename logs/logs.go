package logs

import (
	"log"
	"os"
	"fmt"
)

type Logger struct {
	loggerHandler *log.Logger
}

var logger *Logger

func NewLogger(logPath string) *Logger {
	fmt.Printf("LoggerFile : %s%c%s \n",logPath,os.PathSeparator,"rpc.log")
	fileHandler, err := os.OpenFile(fmt.Sprintf("%s%c%s",logPath,os.PathSeparator,"rpc.log"), os.O_CREATE|os.O_APPEND , os.ModePerm)
	if err != nil {
		panic(err)
	}
	return &Logger{
		loggerHandler:log.New(fileHandler, "LOGGER:", log.LstdFlags),
	}
}

func Init(logPath string)  {
	logger = NewLogger(logPath)
}

//INFO
func Info(message string, context interface{})  {
	logger.loggerHandler.Printf("%s \t %v", message, context)
}
//ERROR
func Error(message string, context interface{})  {
	logger.loggerHandler.Fatalf("%s \t %v", message, context)
}