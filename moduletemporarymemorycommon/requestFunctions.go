package moduletemporarymemorycommon

import (
	"fmt"

	"ISEMS-MRSICT/modulelogginginformationerrors"
)

/*** методы относящиеся к хранилищу модуля moduleLoggingInformationOrErrors ***/

//SetChanModuleLoggingInformationOrError при инициализации приложения сохраняет канал доступа к модулю moduleLoggingInformationOrErrors
func (stmc *StorageTemporaryMemoryCommonType) SetChanModuleLoggingInformationOrError(logMessage chan modulelogginginformationerrors.LogMessageType) {
	fmt.Println("func 'SetChanModuleLoggingInformationOrError', START...")

	chanRes := make(chan parametersResponseModuleLoggingInformationOrErrors)
	defer func() {
		close(chanRes)
	}()

	stmc.channelRequest <- channelStorageTemporaryMemoryCommonRequest{
		moduleType: "module logging information or errors",
		actionType: "set chan module logging information or error",
		parameters: &parametersType{
			parametersStorageModuleLoggingInformationOrErrors: StorageModuleLoggingInformationOrErrors{
				chanLogMessage: logMessage,
			},
		},
		channels: &channelsType{
			channelResponseModuleLoggingInformationOrErrors: chanRes,
		},
	}

	<-chanRes
}

//GetChanModuleLoggingInformationOrError при инициализации приложения возвращает канал доступа к модулю moduleLoggingInformationOrErrors
func (stmc *StorageTemporaryMemoryCommonType) GetChanModuleLoggingInformationOrError() (logMessage chan modulelogginginformationerrors.LogMessageType) {
	fmt.Println("func 'GetChanModuleLoggingInformationOrError', START...")

	chanRes := make(chan parametersResponseModuleLoggingInformationOrErrors)
	defer func() {
		close(chanRes)
	}()

	stmc.channelRequest <- channelStorageTemporaryMemoryCommonRequest{
		moduleType: "module logging information or errors",
		actionType: "get chan module logging information or error",
		channels: &channelsType{
			channelResponseModuleLoggingInformationOrErrors: chanRes,
		},
	}

	return (<-chanRes).chanLogMessage
}
