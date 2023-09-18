package modulecoreapplication

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/moduleapirequestprocessing"
	moddatamodels "ISEMS-MRSICT/modulecoreapplication/datamodels"
	"ISEMS-MRSICT/moduledatabaseinteraction"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

var clim moddatamodels.ChannelsListInteractingModules
var tst *memorytemporarystoragecommoninformation.TemporaryStorageType

func init() {
	clim = moddatamodels.ChannelsListInteractingModules{}

	//инициализируем временное хранилище
	tst = memorytemporarystoragecommoninformation.NewTemporaryStorage()
}

// MainHandlerCoreApplication основной обработчик ядра приложения
func MainHandlerCoreApplication(chanSaveLog chan<- modulelogginginformationerrors.LogMessageType, appConfig *datamodels.AppConfig) {
	funcName := "MainHandlerCoreApplication"

	/* получаем и сохраняем во временном хранилище дефолтные настройки приложения */
	ssdmct, err := getListSettings(appConfig.DefaultSettingsFiles.SettingsStatusesDecisionsMadeComputerThreats, appConfig)
	if err != nil {
		log.Fatal("Error! The file 'settingsStatusesDecisionsMadeComputerThreats.json' with default settings not found.")
	}
	tst.SetListDecisionsMade(ssdmct)

	sctt, err := getListSettings(appConfig.DefaultSettingsFiles.SettingsComputerThreatTypes, appConfig)
	if err != nil {
		log.Fatal("Error! The file 'settingsComputerThreatTypes.json' with default settings not found.")
	}
	tst.SetListComputerThreat(sctt)

	/* инициализируем модули взаимодействия с БД */
	cdbi, err := moduledatabaseinteraction.MainHandlerDataBaseInteraction(chanSaveLog, &appConfig.ConnectionsDataBase, tst)
	if err != nil {
		fmt.Println("An error occurred while initializing the database connection module.")
		fmt.Println(err)

		return
	}
	clim.ChannelsModuleDataBaseInteraction = cdbi

	/* инициализируем модуль обработки запросов с внешних источников */
	capirp := moduleapirequestprocessing.MainHandlerAPIReguestProcessing(chanSaveLog, &appConfig.ModuleAPIRequestProcessingSettings, &appConfig.CryptographySettings)
	clim.ChannelsModuleAPIRequestProcessing = capirp

	chanSaveLog <- modulelogginginformationerrors.LogMessageType{
		TypeMessage: "info",
		Description: "the application initialization was completed successfully",
		FuncName:    funcName,
	}

	//делаем запрос к БД для инициализации хранилища DO STIX типа 'Grouping'
	clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.InputModule <- datamodels.ModuleDataBaseInteractionChannel{
		CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
			ModuleGeneratorMessage: "module core application",
			ModuleReceiverMessage:  "module database interaction",
		},
		Section: "handling technical part",
		Command: "create STIX DO type 'grouping'",
	}

	RoutingCoreApp(chanSaveLog, appConfig, tst, &clim)
}

func getListSettings(f string, appConfig *datamodels.AppConfig) (map[string]datamodels.StorageApplicationCommonListType, error) {
	tmp := map[string]string{}
	configFileSettings := map[string]datamodels.StorageApplicationCommonListType{}

	//проверяем наличие файлов с дефолтными настройками приложения
	row, err := os.ReadFile(path.Join(appConfig.RootDir, f))
	if err != nil {
		return configFileSettings, fmt.Errorf("the file '%s' with default settings not found", f)
	}

	err = json.Unmarshal(row, &tmp)
	if err != nil {
		return configFileSettings, err
	}

	for k, v := range tmp {
		configFileSettings[k] = datamodels.StorageApplicationCommonListType{Description: v}
	}

	return configFileSettings, err
}
