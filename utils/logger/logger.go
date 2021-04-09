package logger

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	ecode "billing/constants"

	log "github.com/sirupsen/logrus"

	guuid "github.com/google/uuid"
)

// InitLogger ...
func InitLogger() {
	// More configs at https://godoc.org/github.com/sirupsen/logrus
	/* LOG level
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	log.Fatal("Bye.") // Calls os.Exit(1) after logging
	log.Panic("I'm bailing.") // Calls panic() after logging
	*/

	runmode := os.Getenv("PROD_MODE")
	if runmode == "1" {
		// log to json format for friendly with logstash/graylog/etc...
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	log.SetOutput(os.Stdout)
}

// PUBLIC Functions

// UUIDv4 ... Format c01d7cf6-ec3f-47f0-9556-a5d6e9009a43
func UUIDv4() string {
	uuid := guuid.New()
	return fmt.Sprintf("%s", uuid)
}

// LogTraceID ...
func LogTraceID() string {
	curTime := time.Now().Format(ecode.DATETIME_FORMAT_DEFAULT)
	return "TID" + ":" + curTime + ":" + fmt.Sprint(_randomNumber(1, 1e6))
}

// Logger ...
func Logger(requestID, extraData string) *log.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	if extraData == "" {
		return log.WithField("file", filename).WithField("func", fn).WithField("txn", requestID)
	}
	return log.WithField("file", filename).WithField("func", fn).WithField("txn", requestID).WithField("ext", extraData)
}

/************************************************************/
// PRIVATE Functions                                        */
/************************************************************/

func _randomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
