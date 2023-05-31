package routingflowsmoduleapirequestprocessing_test

/*
import (
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"
	"fmt"
	"io/ioutil"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//Набор поведенческийх тестов прохождения JSON сообщений содержащих объекты
// ReferenceBook, SDO, Статистических запросов и запросов управления
// через обработчик HandlerAssignmentsModuleAPIRequestProcessing

//ReadTestJSONFile - чтение тестового набора данных из файла и первичное их преобразование
// к виду данных циркулирующих между модулем обрабатывающем запросы с внешних источников
// и ядром приложения
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

var _ = Describe("Тесты объектов ReferenceBook, SDO, Статистических запросов и запросов управления через обработчик HandlerAssignmentsModuleAPIRequestProcessing", func() {

	TestData := make(map[string]datamodels.ModuleReguestProcessingChannel)

	// Загрузка из файла тестовых данных для STIXObjects
	TestFilePath := filepath.Join(modulePath, "..", "..", "mytest/test_resources/jsonSTIXExample.json")
	TestData["STIXObject"] = *ReadTestJSONFile(TestFilePath)

	// Загрузка из файла правильных тестовых данных для RBook
	TestFilePath = filepath.Join(modulePath, "..", "..", "mytest/test_resources/RBookAPIHNotationReq_good.json")
	TestData["GoodRBook"] = *ReadTestJSONFile(TestFilePath)

	// Загрузка из файла сломанных тестовых данных для RBook
	TestFilePath = filepath.Join(modulePath, "..", "..", "mytest/test_resources/RBookAPIHNotationReq_bad.json")
	TestData["BadRBook"] = *ReadTestJSONFile(TestFilePath)

	//GoodRouteBehavior - Утверждения проверки обработки валидных запросов
	GoodRouteBehavior := func() {

		It("Сравнение по TaskByID сообщений во временном хранилище и отправленного в канал БД", func() {
			DBMsg := <-clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.InputModule
			AppTaskID, _, err := tst.GetTaskByID(DBMsg.AppTaskID)
			Expect(err).Should(BeNil())
			Expect(DBMsg.AppTaskID).Should(Equal(AppTaskID))
			//close(done)
		})
	}

	Describe("Тестирование прохождения API запросов через секцию handling reference book", func() {
		Context("Прохождение валидных RB запросов", func() {
			go routingflowsmoduleapirequestprocessing.HandlerAssignmentsModuleAPIRequestProcessing(chanSaveLog, TestData["GoodRBook"], tst, clim)
			GoodRouteBehavior()
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
	})
})*/
