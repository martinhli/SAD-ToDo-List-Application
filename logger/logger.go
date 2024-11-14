package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Initialize() {
	// Initialize the logger
	Log = logrus.New()

	//Set logger format and level
	Log.SetFormatter(&logrus.JSONFormatter{}) // Using the JSON format for logging
	Log.SetLevel(logrus.InfoLevel)            // Set the log level to info
	Log.SetReportCaller(true)                 // Include call info

	//Log to a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Warn("There was an error with logging to the file, using stderr")
	}
}
