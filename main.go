package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/modulecoreapplication"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

var appConfig datamodels.AppConfig
var chanSaveLog chan modulelogginginformationerrors.LogMessageType

//ReadConfig читает конфигурационный файл и сохраняет данные в appConfig
func readConfigApp(fileName string, appc *datamodels.AppConfig) error {
	var err error
	row, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(row, &appc)
	if err != nil {
		return err
	}

	return err
}

//getVersionApp получает версию приложения из файла README.md
func getVersionApp(appc *datamodels.AppConfig) error {
	failureMessage := "version not found"
	content, err := ioutil.ReadFile(appc.RootDir + "README.md")
	if err != nil {
		return err
	}

	//Application ISEMS-NIH master, v0.1
	pattern := `^Application\sISEMS-MRSICT,\sv\d+\.\d+\.\d+`
	rx := regexp.MustCompile(pattern)
	numVersion := rx.FindString(string(content))

	if len(numVersion) == 0 {
		appc.VersionApp = failureMessage

		return nil
	}

	s := strings.Split(numVersion, " ")

	fmt.Println(s)

	if len(s) < 3 {
		appc.VersionApp = failureMessage

		return nil
	}

	appc.VersionApp = s[2]

	return nil
}

func init() {
	fmt.Println("func 'init', START...")

	var err error

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	//читаем конфигурационный файл приложения
	err = readConfigApp(path.Join(dir, "/config.json"), &appConfig)
	if err != nil {
		log.Fatal("Error! The configuration file cannot be read.")
	}

	appConfig.RootDir = dir + "/"

	//получаем номер версии приложения
	if err = getVersionApp(&appConfig); err != nil {
		fmt.Println(err)
	}

	chanSaveLog, err = modulelogginginformationerrors.New(&modulelogginginformationerrors.MainHandlerLoggingParameters{
		LocationLogDirectory: appConfig.LocationLogDirectory,
		NameLogDirectory:     appConfig.NameLogDirectory,
		MaxSizeLogFile:       appConfig.MaxSizeLogFile,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("func 'main', START...")
	fmt.Println(appConfig)

	defer func() {
		if err := recover(); err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprintf("STOP 'main' function, Error:'%v'", err),
				FuncName:    "main",
			}
		}
	}()

	modulecoreapplication.MainHandlerCoreApplication(chanSaveLog, &appConfig)
}
