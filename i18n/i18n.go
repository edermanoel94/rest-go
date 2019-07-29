package i18n

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

// Langs available
const (
	PtBR = "pt-BR"
	EnUS = "en-US"
)

var messages map[string]string

func Start(path string, lang string) error {

	filePath := path + string(os.PathSeparator) + lang + ".json"
	if _, err := os.Stat(filePath); err == nil {

		file, err := os.Open(filePath)

		if err != nil {
			return err
		}

		bytes, err := ioutil.ReadAll(file)

		if err != nil {
			return err
		}

		err = json.Unmarshal(bytes, &messages)

		if err != nil {
			return err
		}

		return nil
	} else if os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("this %s language file not exist", lang))
	} else {
		return err
	}
}

func Message(key string) string {
	if message, ok := messages[key]; ok {
		return message
	}
	return ""
}

func Error(key string) error {
	if message, ok := messages[key]; ok {
		return errors.New(message)
	}
	return nil
}
