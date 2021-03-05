package modulecoreapplication

import (
	"fmt"

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

//MainHandlerCoreApplication основной обработчик ядра приложения
func MainHandlerCoreApplication(chanSaveLog chan<- modulelogginginformationerrors.LogMessageType, appConfig *datamodels.AppConfig) {
	funcName := "MainHandlerCoreApplication"

	fmt.Println("func 'MainHandlerCoreApplication', START...")

	//инициализируем модули взаимодействия с БД
	cdbi, err := moduledatabaseinteraction.MainHandlerDataBaseInteraction(&appConfig.ConnectionsDataBase, tst)
	if err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    funcName,
		}

		fmt.Println("An error occurred while initializing the database connection module.")

		return
	}
	clim.ChannelsModuleDataBaseInteraction = cdbi

	//инициализируем модуль обработки запросов с внешних источников
	capirp := moduleapirequestprocessing.MainHandlerAPIReguestProcessing(stmc, &appConfig.ModuleAPIRequestProcessingSettings, &appConfig.CryptographySettings)
	clim.ChannelsModuleAPIRequestProcessing = capirp

	chanSaveLog <- modulelogginginformationerrors.LogMessageType{
		TypeMessage: "info",
		Description: "the application initialization was completed successfully",
		FuncName:    funcName,
	}

	RoutingCoreApp(chanSaveLog, appConfig, tst, &clim)
}
