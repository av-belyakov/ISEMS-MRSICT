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

			_, _ = tempStorage.AddNewTask(&TemporaryStorageTaskType{
				TaskGenerator:  "task generated test module",
				ClientID:       "19293hdh8883g827g7dg7373747",
				ClientName:     "client_name_2",
				ClientTaskID:   "nduncuuhf4fh84fh8h48fh48f",
				Section:        "stix object test",
				Command:        "create",
				TaskParameters: 4100,
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
			It("Хранилище taskStorage должно содержать 2 задачи полученные от клиента с ID '19293hdh8883g827g7dg7373747'", func(done Done) {
				listTaskID := tempStorage.GetTasksByClientID("19293hdh8883g827g7dg7373747")

				Expect(len(listTaskID)).To(Equal(2))

				close(done)
			})
		})

		Context("Тест 2. Проверяем возможность управления задачами", func() {
			It("Должна быть успешная модификация статуса задачи", func(done Done) {
				err := tempStorage.ChangeTaskStatus(appTaskIDOne, "ex")

				Expect(err).ShouldNot(HaveOccurred())

				close(done)
			})
			It("Должна быть успешная модификация параметра RemovalRequired задачи, при модификации параметра ошибки быть не должно", func(done Done) {
				err := tempStorage.ChangeRemovalRequiredParameter(appTaskIDOne)

				Expect(err).ShouldNot(HaveOccurred())

				_, taskInfo, err := tempStorage.GetTaskByID(appTaskIDOne)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(taskInfo.RemovalRequired).Should(BeTrue())

				close(done)
			})
			It("Должно успешно изменятся время модификации информации о задачи", func() {
			})
		})
	})
})
