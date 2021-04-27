package fluent

import (
	"os"
	"strconv"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/labstack/gommon/log"
)

var F *fluent.Fluent

func Init() {
	Host := os.Getenv("FLUENT_HOST")
	Port, _ := strconv.Atoi(os.Getenv("FLUENT_PORT"))
	logger, err := fluent.New(fluent.Config{
		FluentPort: Port, FluentHost: Host,
	})
	if err != nil {
		log.Error(err)
	}
	F = logger
}
func Push(app string) {
	tag := "myapp.access"
	var data = map[string]string{
		"advertiser_id":     "",
		"android_id":        "",
		"app_code":          "",
		"app_id":            "",
		"app_key":           "",
		"app_version":       "",
		"brand":             "",
		"bundle_identifier": "",
		"carrier":           "",
		"country_code":      "VN",
		"cpu_abi":           "",
		"cpu_abi2":          "",
		"device":            "",
		"device_model":      "",
		"device_type":       "user",
		"display":           "",
		"event_value":       "",
		"fcm":               "",
		"finger_print":      "",
		"install_time":      "",
		"language":          "Tiếng Việt",
		"last_update_time":  "",
		"operator":          "",
		"os_version":        "",
		"platform":          "",
		"product":           "",
		"sdk":               "23",
		"sdk_version":       "1.0.0",
		"server_timestamp":  "",
		"time_zone":         "UTC",
		"timestamp":         "",
	}
	err := F.Post(tag, data)

	if err != nil {
		log.Error(err)
	}
}
