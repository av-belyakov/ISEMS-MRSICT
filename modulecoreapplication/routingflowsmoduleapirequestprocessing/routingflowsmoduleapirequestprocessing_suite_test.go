package routingflowsmoduleapirequestprocessing_test

import (
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	moddatamodels "ISEMS-MRSICT/modulecoreapplication/datamodels"
	"ISEMS-MRSICT/moduledatabaseinteraction"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
	"ISEMS-MRSICT/modulelogginginformationerrors"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	tst         *memorytemporarystoragecommoninformation.TemporaryStorageType
	clim        *moddatamodels.ChannelsListInteractingModules
	chanSaveLog chan modulelogginginformationerrors.LogMessageType
	chdbi       moduledatabaseinteraction.ChannelsModuleDataBaseInteraction
	modulePath  string
)

func TestRoutingflowsmoduleapirequestprocessing(t *testing.T) {
	RegisterFailHandler(Fail)

	modulePath, _ = os.Getwd()
	//Глобально инициализируем необходимые для тестирования каналы
	tst = memorytemporarystoragecommoninformation.NewTemporaryStorage()
	chanSaveLog = make(chan modulelogginginformationerrors.LogMessageType)
	chdbi = moduledatabaseinteraction.ChannelsModuleDataBaseInteraction{
		ChannelsMongoDB: interactionmongodb.ChannelsMongoDBInteraction{
			InputModule:  make(chan datamodels.ModuleDataBaseInteractionChannel),
			OutputModule: make(chan datamodels.ModuleDataBaseInteractionChannel),
		},
	}
	clim = &moddatamodels.ChannelsListInteractingModules{}
	clim.ChannelsModuleDataBaseInteraction = chdbi

	RunSpecs(t, "Routingflowsmoduleapirequestprocessing Suite")
}

var _ = BeforeSuite(func() {
	By("Проверка  глобальных для Routingflowsmoduleapirequestprocessing Suite объектов")
	checkChanel()
})

//checkChanel - Утверждения проверки создания каналов для тестирования
func checkChanel() {
	//	It("Проверка  каналов", func() {
	Eventually(clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.InputModule).ShouldNot(BeClosed())
	Eventually(clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.OutputModule).ShouldNot(BeClosed())
	Eventually(chanSaveLog).ShouldNot(BeClosed())
	//	})
}
