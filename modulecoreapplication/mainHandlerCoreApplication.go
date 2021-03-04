package modulecoreapplication

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/moduleapirequestprocessing"
	moddatamodels "ISEMS-MRSICT/modulecoreapplication/datamodels"
	"ISEMS-MRSICT/moduledatabaseinteraction"
	"ISEMS-MRSICT/modulelogginginformationerrors"
	"ISEMS-MRSICT/moduletemporarymemorycommon"
)

var clim moddatamodels.ChannelsListInteractingModules
var stmc *moduletemporarymemorycommon.StorageTemporaryMemoryCommonType

func init() {
	clim = moddatamodels.ChannelsListInteractingModules{}
	stmc = moduletemporarymemorycommon.NewStorageTemporaryMemoryCommon()
}

//MainHandlerCoreApplication основной обработчик ядра приложения
func MainHandlerCoreApplication(chanSaveLog chan<- modulelogginginformationerrors.LogMessageType, appConfig *datamodels.AppConfig) {
	funcName := "MainHandlerCoreApplication"

	fmt.Println("func 'MainHandlerCoreApplication', START...")

	//добавляем доступ к методам модуля ModuleLoggingInformationOrErrors
	stmc.SetChanModuleLoggingInformationOrError(chanSaveLog)

	//инициализируем модули взаимодействия с БД
	cdbi, err := moduledatabaseinteraction.MainHandlerDataBaseInteraction(&appConfig.ConnectionsDataBase)
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

	RoutingCoreApp(chanSaveLog, appConfig, &clim)
}
