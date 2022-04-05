package interactionredisearchdb_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RediSearch/redisearch-go/redisearch"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionredisearchdb"
)

var cdrdb interactionredisearchdb.ConnectionDescriptorRedisearchDB

var _ = Describe("AddIndexSTIXObject", func() {
	var errConnect error

	var _ = BeforeSuite(func() {
		errConnect = cdrdb.CreateConnection(&datamodels.RedisearchDBSettings{
			Host: "test-uchet-db.cloud.gcm",
			Port: "6379",
		})
	})

	Context("Test 1. Проверка успешного соединения с БД через функцию CreateConnection", func() {
		It("Должно быть успешно установленно соединение с БД", func() {
			Expect(errConnect).ShouldNot(HaveOccurred())
		})
	})

	Context("Test 2. Проверяем добавление новых документов для создания индексов", func() {
		It("Должно быть добавленно 2 новых документа и поним построенны тндексы", func() {
			doc_1 := redisearch.NewDocument("campaign--0bd1475b-02df-4f51-99db-e061b16a6956", 1.0).
				Set("type", "campaign").
				Set("name", "Blue group Attacks Against Finance").
				Set("description", "Campaign by Blue Group against a series of targets in the financial services sector! My new description. &#010;Returns an object that only contains the whitelisted attributes. It will remove all attributes that have a falsy value in the whitelist.&#010;&#010;It also accepts a constraints object used for the validation but to make it keep attributes that doesn&apos;t have any constraints you can simply set the constraints for that attribute to {}. Joido, ihd jieieof. GGhhhhh 122.")

			errDoc_1 := cdrdb.Connection.IndexOptions(
				redisearch.IndexingOptions{
					Replace: true,
					Partial: true,
				}, []redisearch.Document{doc_1}...)
			Expect(errDoc_1).ShouldNot(HaveOccurred())

			doc_2 := redisearch.NewDocument("identity--5f4cad3c-b271-4aec-8159-5b9a09ff2b80", 1.0).
				Set("type", "identity").
				Set("name", "Alexx Ivanov").
				Set("description", "Мое новое описание идентичности. Теперь добавим немного текста.").
				Set("url", "http.examle-many-domain-name.net/home.js?name=gun&news=1")
			errDoc_2 := cdrdb.Connection.IndexOptions(
				redisearch.IndexingOptions{
					Replace: true,
					Partial: true,
				}, []redisearch.Document{doc_2}...)
			Expect(errDoc_2).ShouldNot(HaveOccurred())

			list, err := cdrdb.Connection.List()

			fmt.Printf("___ list documents: %v\n", list)

			Expect(len(list)).Should(Equal(1))
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Test 3. Проверяем добавление ГРУППЫ новых документов или обнавления старых", func() {
		It("Должна быть добавлена новая группа документов состоящая из 3 документов", func() {
			var newDocumentList = make([]redisearch.Document, 0, 3)

			newDocumentList = append(newDocumentList,
				redisearch.NewDocument("indicator--d38a99ae-c5ee-4542-bc12-dfe68b48cc08", 1.0).
					Set("type", "indicator").
					Set("name", "Poison Ivy Malware").
					Set("description", "This file is part of Poison Ivy"))

			newDocumentList = append(newDocumentList,
				redisearch.NewDocument("location--a6e9345f-5a15-4c29-8bb3-7dcc5d234d64", 1.0).
					Set("type", "location").
					Set("street_address", "г. Пермь, ул. Старославянская, д.46, к.2"))

			newDocumentList = append(newDocumentList,
				redisearch.NewDocument("malware--e82e93f6-7911-40d9-8b4a-5abc9dfc1efa", 1.0).
					Set("type", "malware").
					Set("name", "Cryptolocker").
					Set("description", "A variant of the cryptolocker family"))

			errAddDoc := cdrdb.Connection.IndexOptions(
				redisearch.IndexingOptions{
					Replace: true,
					Partial: true,
				}, newDocumentList...)
			Expect(errAddDoc).ShouldNot(HaveOccurred())

			docList, docNum, err := cdrdb.Connection.Search(redisearch.NewQuery("*").
				AddFilter(
					redisearch.Filter{
						Field: "name",
					},
				).
				SetReturnFields("id"))

			fmt.Printf("______FULL SEARCH DOCUMENTS docNum: %d\n docList: %v\n", docNum, docList)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(docNum).Should(Equal(5))
		})
	})

	Context("Test 4. Поиск индексов по различным полям", func() {
		It("Должен быть найден индекс по полю name и значению 'Attacks Against'", func() {
			docList, docNum, err := cdrdb.Connection.Search(redisearch.NewQuery("Attacks Against").
				AddFilter(
					redisearch.Filter{
						Field: "name",
					},
				))

			fmt.Printf("______ docNum: %d\n docList: %v\n", docNum, docList)

			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Должен быть найден индекс по полю description и значению 'новое описание'", func() {
			docList, docNum, err := cdrdb.Connection.Search(redisearch.NewQuery("новое описание").
				AddFilter(
					redisearch.Filter{
						Field: "description",
					},
				))

			fmt.Printf("______ docNum1: %d\n docList1: %v\n", docNum, docList)

			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Должен быть найден индекс по полю description и значению 'исание'", func() {
			docList, docNum, err := cdrdb.Connection.Search(redisearch.NewQuery("%anov%").
				AddFilter(
					redisearch.Filter{
						Field: "name",
					},
				))

			fmt.Printf("______ docNum1: %d\n docList1: %v\n", docNum, docList)

			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
