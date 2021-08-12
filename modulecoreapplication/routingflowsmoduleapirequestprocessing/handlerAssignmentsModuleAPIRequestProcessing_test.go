package routingflowsmoduleapirequestprocessing_test

import (
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	moddatamodels "ISEMS-MRSICT/modulecoreapplication/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"
	"ISEMS-MRSICT/moduledatabaseinteraction"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
	"ISEMS-MRSICT/modulelogginginformationerrors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//ReadTestJSONFile - чтение тестового набора данных из файла и первичное их преобразование
func ReadTestJSONFile(path string) *datamodels.ModuleReguestProcessingChannel {
	var (
		TestModReq datamodels.ModuleReguestProcessingChannel
	)
	reqF, err := ioutil.ReadFile(path)
	Expect(err).ShouldNot(HaveOccurred(), fmt.Sprintf("Неудалось прочитать файл %s хранящий тестовые данные", path))
	TestModReq = datamodels.ModuleReguestProcessingChannel{ClientID: "TEST",
		ClientName: "TestName",
		DataType:   1,
		Data:       &reqF}
	return &TestModReq
}

var _ = Describe("Тестирование HandlerAssigmentsModuleAPIRequestProcessing", func() {

	var (
		//инициализируем необходимые каналы
		tst         *memorytemporarystoragecommoninformation.TemporaryStorageType = memorytemporarystoragecommoninformation.NewTemporaryStorage()
		clim        *moddatamodels.ChannelsListInteractingModules                 = &moddatamodels.ChannelsListInteractingModules{}
		chanSaveLog chan modulelogginginformationerrors.LogMessageType            = make(chan modulelogginginformationerrors.LogMessageType)
		chdbi       moduledatabaseinteraction.ChannelsModuleDataBaseInteraction   = moduledatabaseinteraction.ChannelsModuleDataBaseInteraction{
			ChannelsMongoDB: interactionmongodb.ChannelsMongoDBInteraction{
				InputModule:  make(chan datamodels.ModuleDataBaseInteractionChannel),
				OutputModule: make(chan datamodels.ModuleDataBaseInteractionChannel),
			},
		}
	)
	clim.ChannelsModuleDataBaseInteraction = chdbi
	TestData := make(map[string]*datamodels.ModuleReguestProcessingChannel)
	dir, _ := os.Getwd()

	// Тестовые данные для STIXObjects
	TestFilePath := filepath.Join(dir, "..", "..", "mytest/test_resources/jsonSTIXExample.json")
	TestData["TestSTIXObject"] = ReadTestJSONFile(TestFilePath)

	// Тестовые данные для RBook
	TestFilePath = filepath.Join(dir, "..", "..", "mytest/test_resources/RBookAPIHNotationReq_good.json")
	TestData["TestGoodRBook"] = ReadTestJSONFile(TestFilePath)
	TestFilePath = filepath.Join(dir, "..", "..", "mytest/test_resources/RBookAPIHNotationReq_bad.json")
	TestData["TestBadRBook"] = ReadTestJSONFile(TestFilePath)

	//assertionCheckChanel - Утверждения проверки создания каналов для тестирования
	assertionCheckChanel := func() {
		It("Проверка  каналов", func() {
			Eventually(clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.InputModule).ShouldNot(BeClosed())
			Eventually(clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.OutputModule).ShouldNot(BeClosed())
			Eventually(chanSaveLog).ShouldNot(BeClosed())
		})
	}

	//assertionGoodBehavior - Утверждения проверки обработки валидных запросов
	assertionGoodBehavior := func() {
		It("Сравнение по TaskByID сообщений во временном хранилище и отправленного в канал БД", func(done Done) {
			DBMsg := <-clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.InputModule
			AppTaskID, _, err := tst.GetTaskByID(DBMsg.AppTaskID)
			Expect(err).Should(BeNil())
			Expect(DBMsg.AppTaskID).Should(Equal(AppTaskID))
			close(done)
		}, 5)
	}
	//assertionGoodBehavior - Утверждения проверки обработки валидных запросов
	assertionCheckChanel()
	Describe("Тестирование прохождения API запросов через секцию handling reference book", func() {
		Context("Тест: обработка валидных запросов", func() {
			go routingflowsmoduleapirequestprocessing.HandlerAssignmentsModuleAPIRequestProcessing(chanSaveLog, TestData["TestGoodRBook"], tst, clim)
			assertionGoodBehavior()
		})
		//Context("Тест: обработка не валидных запросов", func() {
		//	go routingflowsmoduleapirequestprocessing.HandlerAssigmentsModuleAPIRequestProcessing(chanSaveLog, TestData["TestBadRBook"], tst, clim)
		//	assertionBadBehavior()
		//})
	})
	/*	Describe("Тестирование прохождения API запросов через секцию handling stix object", func() {
		Context("Тест: обработка валидных запросов", func() {
			Skip("")
			go routingflowsmoduleapirequestprocessing.HandlerAssigmentsModuleAPIRequestProcessing(chanSaveLog, TestData["TestSTIXObject"], tst, clim)
			assertionGoodBehavior()
		})
		Context("Тест: обработка не валидных запросов", func() {
		})
	})*/

})
