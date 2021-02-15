package moduletemporarymemorycommon

import (
	"fmt"
)

//handlerModuleLoggingInformationOrErrors обработчик параметров модуля moduleLoggingInformationOrErrors
func (stmc *StorageTemporaryMemoryCommonType) handlerModuleLoggingInformationOrErrors(
	channelResponse chan<- parametersResponseModuleLoggingInformationOrErrors,
	actionType string,
	parameters *parametersType) {

	switch actionType {
	case "set chan module logging information or error":
		fmt.Printf("func 'handlerModuleLoggingInformationOrErrors', action type: 'SET chan module logging information or error'\n")

		stmc.parametersModuleLoggingInformationOrErrors = StorageModuleLoggingInformationOrErrors{
			chanLogMessage: parameters.parametersStorageModuleLoggingInformationOrErrors.chanLogMessage,
		}

		channelResponse <- parametersResponseModuleLoggingInformationOrErrors{}

	case "get chan module logging information or error":
		fmt.Printf("func 'handlerModuleLoggingInformationOrErrors', action type: 'GET chan module logging information or error'\n")

		channelResponse <- parametersResponseModuleLoggingInformationOrErrors{
			chanLogMessage: stmc.parametersModuleLoggingInformationOrErrors.chanLogMessage,
		}

	}
}
