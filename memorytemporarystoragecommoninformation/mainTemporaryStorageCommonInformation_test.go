package memorytemporarystoragecommoninformation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MainTemporaryStorageCommonInformation", func() {
	var _ = Describe("taskStorage", func() {
		var (
			appTaskIDOne    string
			errorAddTaskOne error
		)
		tempStorage := NewTemporaryStorage()

		var _ = BeforeSuite(func() {
			appTaskIDOne, errorAddTaskOne = tempStorage.AddNewTask(&TemporaryStorageTaskType{
				TaskGenerator:  "task generated test module",
				ClientID:       "gfydgf737gf7gf73g7gf7g37f38fg838345",
				ClientName:     "client_name_1",
				ClientTaskID:   "3399hhd82wqa222h388ddp",
				Section:        "stix object test",
				Command:        "none",
				TaskParameters: 23,
			})

			_, _ = tempStorage.AddNewTask(&TemporaryStorageTaskType{
				TaskGenerator:  "task generated test module",
				ClientID:       "19293hdh8883g827g7dg7373747",
				ClientName:     "client_name_2",
				ClientTaskID:   "iififieif99f939fjjf3jjr",
				Section:        "stix object test",
				Command:        "none",
				TaskParameters: 42,
			})
		})

		Context("Тест 1. Проверяем наличие задач", func() {
			It("В taskStorage должна быть успешно добавлена задача", func(done Done) {
				Expect(errorAddTaskOne).ShouldNot(HaveOccurred())

				close(done)
			})
			It("В taskStorage должна быть найдена задача с заданным UUID", func(done Done) {
				_, taskInfo, err := tempStorage.GetTaskByID(appTaskIDOne)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(taskInfo.ClientID).To(Equal("gfydgf737gf7gf73g7gf7g37f38fg838345"))

				close(done)
			})
			/*It("Хранилище taskStorage должно содержать не менее 2 задач", func() {

			})*/
		})
	})
})
