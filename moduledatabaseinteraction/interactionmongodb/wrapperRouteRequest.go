package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/commonhandlers"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

//WrapperFuncTypeHandlingSTIXObject набор обработчиков для работы с запросами связанными со STIX объектами
func (ws *wrappersSetting) wrapperFuncTypeHandlingSTIXObject(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	data datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {

	fn := "wrapperFuncTypeHandlingSTIXObject"

	fmt.Println("func 'wrapperFuncTypeHandlingSTIXObject', START...")
	fmt.Printf("func 'wrapperFuncTypeHandlingSTIXObject', received message: '%v'\n", ws)

	//получаем всю информацию о выполняемой задаче
	_, taskInfo, err := tst.GetTaskByID(data.AppTaskID)
	if err != nil {
		ws.ChanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    fn,
		}

		chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module database interaction",
				ModuleReceiverMessage:  "module core application",
				ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
					FuncName:                                fn,
					ModuleAPIRequestProcessingSettingSendTo: true,
					Error:                                   fmt.Errorf("no information about the task by its id was found in the temporary storage"),
				},
			},
			Section:   "handling stix object",
			AppTaskID: data.AppTaskID,
		}

		return
	}

	fmt.Printf("func 'wrapperFuncTypeHandlingSTIXObject', task info: '%v'\n", taskInfo)

	ti, ok := taskInfo.TaskParameters.([]*datamodels.ElementSTIXObject)
	if !ok {
		msg := "type conversion error"

		ws.ChanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: msg,
			FuncName:    fn,
		}

		chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module database interaction",
				ModuleReceiverMessage:  "module core application",
				ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
					FuncName:                                fn,
					ModuleAPIRequestProcessingSettingSendTo: true,
					Error:                                   fmt.Errorf(msg),
				},
			},
			Section:   "handling stix object",
			AppTaskID: data.AppTaskID,
		}

		return
	}

	//получаем список ID STIX объектов предназначенных для добавление в БД
	listID := commonhandlers.GetListIDFromListSTIXObjects(ti)

	fmt.Printf("func 'wrapperFuncTypeHandlingSTIXObject', list ID: '%v'\n", listID)

	//делаем запрос к БД для получения полной информации об STIX объектах по их ID

	//выполняем сравнение объектов и ищем внесенные изменения для каждого из STIX объектов

	//логируем изменения в STIX объектах в отдельную коллекцию построенную на основе JSON patch (однако в JSON patch нет времени изменения
	// и кто внес изменения), по этому JSON patch нужно будет немного модифицировать

	//добавляем или обновляем STIX объекты в БД

	/* последние ДВА пункта можно делать параллельно */

	/*
		   Так как это обработка STIX объектов то определение и выбор действий по команде не нужен
		   - получаем информацию о задаче по ее ID в memorytemporarystoragecommoninformation.TemporaryStorageType
		   - надо выполнить проверку наличия в БД каждого STIX объекта из полученного среза и если он есть в БД
		    проверить чем они отличаются
		   - нужен какой то промежуточный обработчик для логирования изменений в STIX объекте (возможно на основе JSON patch),
		    для этого необходимо получить все поля со всеми изменениями. Фактически есть три случая:
			1. Объект в БД отсутствует. Проверку на наличие изменений не производим.
			2. Объект в БД есть, но он полностью совпадает с полученным объектом (тогда его не нужно загружать в БД)
			3. Объект в БД есть и он отличается от полученного (тогда нужно найти внесенные изменения)
	*/
}

//wrapperFuncTypeHandlingSearchRequests набор обработчиков для работы с запросами направленными на обработку поисковой машине
func (ws *wrappersSetting) wrapperFuncTypeHandlingSearchRequests(tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {
	/*switch wt.command {
	case "find_all":

	case "find_all_for_client_API":

	case "":

	}*/
}

//wrapperFuncTypeHandlingReferenceBook набор обработчиков для работы с запросами к справочнику
func (ws *wrappersSetting) wrapperFuncTypeHandlingReferenceBook(tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {
	/*switch wt.command {
	case "find_all":

	case "find_all_for_client_API":

	case "":

	}*/
}
