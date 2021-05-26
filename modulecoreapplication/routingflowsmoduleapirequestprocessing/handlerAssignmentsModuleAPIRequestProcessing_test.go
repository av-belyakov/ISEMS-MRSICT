package routingflowsmoduleapirequestprocessing_test

import (
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	moddatamodels "ISEMS-MRSICT/modulecoreapplication/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"
	"ISEMS-MRSICT/modulelogginginformationerrors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Тестирование секции обработки reference book в handlerAssignmentsModuleAPIRequestProcessing", func() {
	var (
		testData    *datamodels.ModuleReguestProcessingChannel
		tst         *memorytemporarystoragecommoninformation.TemporaryStorageType
		clim        *moddatamodels.ChannelsListInteractingModules
		chanSaveLog chan<- modulelogginginformationerrors.LogMessageType
	)
	var_ = BeforeSuite(func() {
		dir, _ := os.Getwd()
	TestFilePath := filepath.Join(dir, "..", "..", "mytest/test_resources/RBookAPIHNotationResp_good.json")
	TestData["TestGoodRBook"] = ReadTestJSONFile(TestFilePath)

	assertionCheckChanel = func() {

	}
		json.Unmarshal(reqF)
		respTestFilePath := filepath.Join(dir, "..", "mytest/test_resources/ReferersBookAPIHierarchicalNotationResponseExample.json")
		if respF, err := ioutil.ReadFile(respTestFilePath); err != nil {
			fmt.Errorf(fmt.Sprintf("Неудалось прочитать файл %s хранящий с тестовые ответы от API к справочникам", respTestFilePath))
		}
		go routingflowsmoduleapirequestprocessing.HandlerAssigmentsModuleAPIRequestProcessing(chanSaveLog, testData, tst, clim)
	})
})
