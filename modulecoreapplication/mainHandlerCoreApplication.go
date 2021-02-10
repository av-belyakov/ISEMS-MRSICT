package modulecoreapplication

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/moduledatabaseinteraction"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

//ChannelsListInteractingModules содержит список каналов для межмодульного взаимодействия
type ChannelsListInteractingModules struct {
	ChansDataBaseInteraction moduledatabaseinteraction.ChansDataBaseInteraction
}

//MainHandlerCoreApplication основной обработчик ядра приложения
func MainHandlerCoreApplication(chanSaveLog chan modulelogginginformationerrors.LogMessageType, appConfig *datamodels.AppConfig) {
	funcName := "MainHandlerCoreApplication"
	clim := ChannelsListInteractingModules{}

	fmt.Println("func 'MainHandlerCoreApplication', START...")

	//инициализируем модули взаимодействия с БД
	chansDataBaseInteraction, err := moduledatabaseinteraction.MainHandlerDataBaseInteraction(&appConfig.ConnectionsDataBase)
	if err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    funcName,
		}

		fmt.Println("An error occurred while initializing the database connection module.")

		return
	}
	clim.ChansDataBaseInteraction = chansDataBaseInteraction

	//инициализируем модуль

	chanSaveLog <- modulelogginginformationerrors.LogMessageType{
		TypeMessage: "info",
		Description: "the application initialization was completed successfully",
		FuncName:    funcName,
	}

	Routing(&clim)
}
