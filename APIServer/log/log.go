/*
Package log implements the relevant logging functions.
*/
package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

//variables for the different log types
var (
	Info    *log.Logger //Important information
	Warning *log.Logger //Be concerned
	Error   *log.Logger //Critical problem
)

//initializes the logging functions
func InitLog() {
	//opens the file for writing logs
	logName := fmt.Sprintf("%s-%s.log", "server_log", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile("./log/"+logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file", err)
	}
	//formats the different log types
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(io.MultiWriter(file, os.Stderr), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}
