package enviroment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fsvxavier/golang-worker-skeleton/pkg/enviroment/interfaces"
)

//Config is the struct of configurations
type ConfigEnviroment struct {
	Env        string
	StatusCode int
	FileConfig string
}

func NewEnviroment() interfaces.ConfigEnviroment {
	c := &ConfigEnviroment{}
	return c
}

func (c *ConfigEnviroment) SetFileConfig(file string) {
	c.FileConfig = file
}

//GetTag representa a consulta de uma tag no arquivo de configuração JSON
func (c *ConfigEnviroment) GetTag(tag string) (string, error) {

	file, err := ioutil.ReadFile(c.FileConfig)
	if err != nil {
		tagValue := os.Getenv(tag)
		return tagValue, nil
	}
	jsonMap := make(map[string]interface{})
	json.Unmarshal(file, &jsonMap)

	env := os.Getenv("ENV")
	database := jsonMap[env].(map[string]interface{})
	tagValue := fmt.Sprintf("%v", database[tag])

	for key, value := range database {

		switch value.(type) {
		case string:
			os.Setenv(key, value.(string))
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			os.Setenv(key, fmt.Sprintf("%d", value.(int)))
		case float32, float64:
			val := fmt.Sprintf("%.2f", value.(float64))
			strings := strings.Split(val, ".")
			if strings[1] != "00" {
				os.Setenv(key, val)
			} else {
				os.Setenv(key, strings[0])
			}
		case bool:
			os.Setenv(key, fmt.Sprintf("%v", value.(bool)))
		}
	}

	return tagValue, nil
}
