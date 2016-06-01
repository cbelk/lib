package loggers

import (
    "fmt"
    "log"
    "os"
    "time"
)

// This function writes the passed message to the passed file and returns a boolean indicating success/failure.
func Logger(application, logfile, message string) (success bool) {
	success = true
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		success = false
	} else {
		if _, err = file.WriteString(message + "\n"); err != nil {
			success = false
		}
		err = file.Close()
		if err != nil {
			fmt.Printf("[-] %s: Error (%v) closing the logfile. Not worthy of termination.", application, err)
		}
	}
	return
}

// This function will write the passed message to the logfile and then call Fatal to terminate execution.
func Freakout(msg, logfile, application string, err error) {
	for i := 0; !Logger(application, logfile, msg); i++ {
		if i == 3 {
			break
		}
	}
	log.Fatal(err)
}

// This function writes the start/end time to the logfile.
func TimeLogger(application, startOrEnd, logfile string) {
	for i := 0; !Logger(application, logfile, fmt.Sprintf("[+] %s: Run %s at %v.", application, startOrEnd, time.Now())); i++ {
		if i == 3 {
			break
		}
	}
}
