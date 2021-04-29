package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	ecode "billing/constants"

	log "github.com/sirupsen/logrus"

	guuid "github.com/google/uuid"
)

const (
	minPageSize     = 1
	defaultPageSize = 10
	maxPageSize     = 99
)

var (
	// reUUID                = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	reMultiSpace          = regexp.MustCompile(`(\s)+`)
	reMoreThan2Linebreaks = regexp.MustCompile(`(\n){2,}`)
	// reMentions            = regexp.MustCompile(`\B@([a-zA-Z][a-zA-Z0-9_-]{0,17})`)
)

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

// MakeHTTPRequest ...
// Support method GET/POST/PUT/DELETE
// Support all content type
// Retry max 3 times for remote server issue (http response code 5xx)
func MakeHTTPRequest(method, url, contentType string, data interface{}, retry bool) (int, []byte, error) {
	const MAX_RETRY_TIMES = 3
	const RETRY_SLEEP_TIME = 50 // millisecond
	var resHTTPStatusCode int = 400
	var resHTTPBody []byte
	var req *http.Request
	var err error

	if data != nil {
		dataByte, _ := json.Marshal(data)
		req, err = http.NewRequest(method, url, bytes.NewBuffer(dataByte))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		Logger("", "").Errorln(err)
		return resHTTPStatusCode, resHTTPBody, err
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for i := 1; i <= MAX_RETRY_TIMES; i++ {
		var body []byte
		res, err := client.Do(req)
		if err == nil {
			defer res.Body.Close()
			body, _ = ioutil.ReadAll(res.Body)

			if res.StatusCode < 500 {
				return res.StatusCode, body, nil
			}
			// Only Retry when remote server error
		} else {
			Logger("", "").Errorln(err)
		}

		if !retry || i+1 > MAX_RETRY_TIMES {
			return res.StatusCode, body, errors.New(res.Status)
		}

		time.Sleep(RETRY_SLEEP_TIME * time.Millisecond)
		fmt.Println("RETRY")
	}

	return resHTTPStatusCode, resHTTPBody, errors.New("Max times retry exceeded")
}

func PostAPIJson(url string, data interface{}) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	dataByte, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataByte))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	req = req.WithContext(ctx)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}

func PostAPI(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}

func SmartTrim(s string) string {
	oldLines := strings.Split(s, "\n")
	newLines := []string{}
	for _, line := range oldLines {
		line = strings.TrimSpace(reMultiSpace.ReplaceAllString(line, "$1"))
		newLines = append(newLines, line)
	}
	s = strings.Join(newLines, "\n")
	s = reMoreThan2Linebreaks.ReplaceAllString(s, "$1$1")
	return strings.TrimSpace(s)
}

func NormalizePageSize(i int) int {
	if i == 0 {
		return defaultPageSize
	}
	if i < minPageSize {
		return minPageSize
	}
	if i > maxPageSize {
		return maxPageSize
	}
	return i
}

/************************************************************/
// PRIVATE Functions                                        */
/************************************************************/

func _randomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
