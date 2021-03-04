package moduletemporarymemorycommon

import (
	"fmt"

	"ISEMS-MRSICT/modulelogginginformationerrors"
)

//StorageTemporaryMemoryCommonType репозиторий для хранения общей информации, необходимой для работы приложения
// parametersModuleLoggingInformationOrErrors - хранилище параметров модуля ModuleLoggingInformationOrErrors
// channelRequest - канал для передачи запросов
type StorageTemporaryMemoryCommonType struct {
	parametersModuleLoggingInformationOrErrors StorageModuleLoggingInformationOrErrors
	channelRequest                             chan channelStorageTemporaryMemoryCommonRequest
}

//StorageModuleLoggingInformationOrErrors хранилище модуля ModuleLoggingInformationOrErrors
// chanLogMessage - канал для взаимодействия с модулем ModuleLoggingInformationOrErrors
type StorageModuleLoggingInformationOrErrors struct {
	chanLogMessage chan<- modulelogginginformationerrors.LogMessageType
}

//channelStorageTemporaryMemoryCommonRequest канал для передачи запросов, управляющих хранилищем временной информации
// moduleType - тип модуля
// actionType - тип действия
// parameters - параметры различных модулей
// channels - каналы для взаимодействия с хранилищами различных модулей
type channelStorageTemporaryMemoryCommonRequest struct {
	moduleType string
	actionType string
	parameters *parametersType
	channels   *channelsType
}

//parametersType
// customParameters - общие, настраиваемые параметры
// parametersStorageModuleLoggingInformationOrErrors - параметры модуля ModuleLoggingInformationOrErrors
type parametersType struct {
	customParameters                                  interface{}
	parametersStorageModuleLoggingInformationOrErrors StorageModuleLoggingInformationOrErrors
}

//channelsRequestType
// channelResponseModuleLoggingInformationOrErrors - канал для приема ответов касающихся взаимодействия с модулем ModuleLoggingInformationOrErrors
// channelResponseStorageTemporaryMemoryCommonResponse - канал для приема общих ответов, полученных от хранилища временной информации
type channelsType struct {
	channelResponseModuleLoggingInformationOrErrors     chan parametersResponseModuleLoggingInformationOrErrors
	channelResponseStorageTemporaryMemoryCommonResponse chan channelStorageTemporaryMemoryCommonResponse
}

//channelStorageTemporaryMemoryCommonResponse канал для приема общих ответов, полученных от хранилища временной информации
type channelStorageTemporaryMemoryCommonResponse struct {
}

type parametersResponseModuleLoggingInformationOrErrors struct {
	chanLogMessage chan<- modulelogginginformationerrors.LogMessageType
}

//NewStorageTemporaryMemoryCommon создание нового репозитория для хранения общей информации, необходимой для работы приложения
func NewStorageTemporaryMemoryCommon() *StorageTemporaryMemoryCommonType {
	fmt.Println("fun 'NewStorageTemporaryMemoryCommon', START...")

	chanReq := make(chan channelStorageTemporaryMemoryCommonRequest)
	stmc := StorageTemporaryMemoryCommonType{
		channelRequest: chanReq,
	}

	go func() {
		for msg := range chanReq {
			switch msg.moduleType {
			case "module logging information or errors":
				chanRes := msg.channels.channelResponseModuleLoggingInformationOrErrors

				stmc.handlerModuleLoggingInformationOrErrors(chanRes, msg.actionType, msg.parameters)

			case "module internal timer":

			case "module notification message":

			}
		}
	}()

	return &stmc
}
