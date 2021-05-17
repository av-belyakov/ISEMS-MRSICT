package memorytemporarystoragecommoninformation

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MainTemporaryStorageCommonInformation", func() {
	var _ = Describe("taskStorage", func() {
		var (
			appTaskIDOne    string
			errorAddTaskOne error
			dateTime        time.Time
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

			_, taskInfo, _ := tempStorage.GetTaskByID(appTaskIDOne)
			dateTime = taskInfo.DateTaskModification
		})

		Context("Тест 1. Проверяем наличие задач", func() {
			It("В taskStorage должна быть успешно добавлена задача", func(done Done) {
				Expect(errorAddTaskOne).ShouldNot(HaveOccurred())

				close(done)
			}, 1)
			It("В taskStorage должна быть найдена задача с заданным UUID", func(done Done) {
				_, taskInfo, err := tempStorage.GetTaskByID(appTaskIDOne)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(taskInfo.ClientID).To(Equal("gfydgf737gf7gf73g7gf7g37f38fg838345"))

				close(done)
			}, 1)
			It("Хранилище taskStorage должно содержать 2 задачи полученные от клиента с ID '19293hdh8883g827g7dg7373747'", func(done Done) {
				listTaskID := tempStorage.GetTasksByClientID("19293hdh8883g827g7dg7373747")

				Expect(len(listTaskID)).To(Equal(2))

				close(done)
			}, 1)
		})

		Context("Тест 2. Проверяем возможность управления задачами", func() {
			It("Должна быть успешная модификация статуса задачи", func(done Done) {
				err := tempStorage.ChangeTaskStatus(appTaskIDOne, "in progress")

				Expect(err).ShouldNot(HaveOccurred())

				_, taskInfo, err := tempStorage.GetTaskByID(appTaskIDOne)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(taskInfo.TaskStatus).To(Equal("in progress"))

				close(done)
			}, 1)
			It("Должна быть успешная модификация параметра RemovalRequired задачи, при модификации параметра ошибки быть не должно", func(done Done) {
				err := tempStorage.ChangeRemovalRequiredParameter(appTaskIDOne)

				Expect(err).ShouldNot(HaveOccurred())

				_, taskInfo, err := tempStorage.GetTaskByID(appTaskIDOne)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(taskInfo.RemovalRequired).Should(BeTrue())

				close(done)
			}, 1)

			Context("Тест 2.1. Проверяем время модификации задачи", func() {
				It("Время модификации задачи полученной ранее и текущей должно отличаться", func() {
					err := tempStorage.ChangeDateTaskModification(appTaskIDOne)

					Expect(err).ShouldNot(HaveOccurred())

					_, taskInfo, err := tempStorage.GetTaskByID(appTaskIDOne)

					fmt.Printf("date time defore: '%v', after: '%v'\n", dateTime, taskInfo.DateTaskModification)

					Expect(err).ShouldNot(HaveOccurred())
					Expect(taskInfo.DateTaskModification).ShouldNot(Equal(dateTime))
				})
			})
		})

		Context("Тест 3. Удаление задачи по заданному ID приложения", func() {
			It("После удаления задачи по ее ID, задачи с таким ID найдено быть не должно", func() {
				tempStorage.DeletingTaskByID(appTaskIDOne)

				_, _, err := tempStorage.GetTaskByID(appTaskIDOne)

				Expect(err).Should(HaveOccurred())
			})
		})

		foundInfoID := "5r3g7f7fg7gf7efe"

		Context("Тест 4.1. Добавление в хранилище, новой, найденной информации", func() {
			It("При добавлении в хранилище новой, найденной информации ошибки быть не должно", func(done Done) {
				err := tempStorage.AddNewFoundInformation(foundInfoID, &TemporaryStorageFoundInformation{
					Collection:  "stix_object_collection",
					ResultType:  "only_count",
					Information: 12,
				})

				Expect(err).ShouldNot(HaveOccurred())

				close(done)
			}, 1)
		})

		Context("Тест 4.2. Получение новой, найденной информации по ее ID", func() {
			It("Должна быть успешно полученна новая информация по ID", func(done Done) {
				info, err := tempStorage.GetFoundInformationByID(foundInfoID)

				fmt.Printf("result func 'GetFoundInformationByID', info: '%v'\n", info)

				Expect(err).ShouldNot(HaveOccurred())

				close(done)
			}, 1)
		})

		Context("Тест 4.3. Удаление из хранилищя, новой, найденной информации, по ее ID", func() {
			It("Должна быть успешно удалена информация по ее ID", func(done Done) {
				tempStorage.DeletingFoundInformationByID(foundInfoID)
				_, err := tempStorage.GetFoundInformationByID(foundInfoID)

				Expect(err).Should(HaveOccurred())

				close(done)
			}, 0.5)
		})

		/*
			Тесты по управлению хранилищем foundInformationStorage прошли успешно,
			однако надо почитать про Ginkgo 2.0, миграцию на эту новую версию и асинхронное тестирование
			в новой версии 2.0
		*/
	})
})
