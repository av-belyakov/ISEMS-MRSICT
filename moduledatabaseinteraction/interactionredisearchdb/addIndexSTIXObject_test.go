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

	Context("Test 3. Поиск индексов по различным полям", func() {
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
			docList, docNum, err := cdrdb.Connection.Search(redisearch.NewQuery("*исание*").
				AddFilter(
					redisearch.Filter{
						Field: "description",
					},
				))

			fmt.Printf("______ docNum1: %d\n docList1: %v\n", docNum, docList)

			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
