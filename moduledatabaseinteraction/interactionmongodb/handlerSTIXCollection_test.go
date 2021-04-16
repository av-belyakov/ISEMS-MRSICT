package interactionmongodb_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

var _ = Describe("HandlerSTIXCollection", func() {
	var (
		connectError, errReadFile, errUnmarchalReq, errUnmarchalToSTIX error
		docJSON                                                        []byte
		cdmdb                                                          interactionmongodb.ConnectionDescriptorMongoDB
		l                                                              []*datamodels.ElementSTIXObject
		qp                                                             interactionmongodb.QueryParameters
		modAPIRequestProcessingReqJSON                                 datamodels.ModAPIRequestProcessingReqJSON
	)

	var _ = BeforeSuite(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

		cdmdb = interactionmongodb.ConnectionDescriptorMongoDB{
			Ctx:       ctx,
			CtxCancel: cancel,
		}

		//подключаемся к базе данных MongoDB
		connectError = cdmdb.CreateConnection(&datamodels.MongoDBSettings{
			Host:     "192.168.13.200",
			Port:     27017,
			User:     "module-isems-mrsict",
			Password: "vkL6Zn$jPmt1e1",
			NameDB:   "isems-mrsict",
		})

		docJSON, errReadFile = ioutil.ReadFile("../../mytest/jsonSTIXExample.json")
		errUnmarchalReq = json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON)
		l, errUnmarchalToSTIX = routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectSTIXReq(modAPIRequestProcessingReqJSON)

		qp = interactionmongodb.QueryParameters{
			NameDB:         "isems-mrsict",
			CollectionName: "stix_object_collection",
			ConnectDB:      cdmdb.Connection,
		}
	})

	var _ = AfterSuite(func() {
		cdmdb.CtxCancel()
	})

	oldCommonPropertiesDomain := datamodels.CommonPropertiesDomainObjectSTIX{
		SpecVersion:  "12.3.4",
		Created:      time.Now(),
		Modified:     time.Now(),
		CreatedByRef: datamodels.IdentifierTypeSTIX("ref_cubu8c3gf8g8g3f83"),
		Revoked:      false,
		Labels:       []string{"lable_1", "lable_5", "lable_3"},
		Lang:         "RU",
		Сonfidence:   12,
		ExternalReferences: datamodels.ExternalReferencesTypeSTIX{
			datamodels.ExternalReferenceTypeElementSTIX{
				SourceName:  "source_name_1",
				Description: "just any descripton",
				URL:         "http://any-site.org/example_one",
				Hashes: datamodels.HashesTypeSTIX{
					"SHA-1":   "dcubdub883g3838fgc83f",
					"SHA-128": "cb8b38b8c38f83f888f844",
				},
				ExternalID: "12444_gddgdg",
			},
			datamodels.ExternalReferenceTypeElementSTIX{
				SourceName:  "source_name_2",
				Description: "just any descripton two",
				URL:         "http://any-site.org/example_two",
				Hashes: datamodels.HashesTypeSTIX{
					"SHA-1":   "dcubdub883g3838fgc83f",
					"SHA-128": "cb8b38b8c38f83f888f844",
					"SHA-256": "bccw62f626fd63fd63f6f36",
				},
				ExternalID: "12444_gddgdg",
			},
		},
		ObjectMarkingRefs: []datamodels.IdentifierTypeSTIX{
			datamodels.IdentifierTypeSTIX("obj_mark_ref_odcfc22y2vr34u41udv"),
			datamodels.IdentifierTypeSTIX("obj_mark_ref_ahvas223h23h2dh3h1h"),
			datamodels.IdentifierTypeSTIX("obj_mark_ref_nivnbufuf74gf74gf7s"),
		},
		Defanged: false,
		Extensions: map[string]string{
			"Key_1": "Value_1",
			"Key_2": "Value_2",
			"Key_3": "Value_3",
		},
	}

	newCommonPropertiesDomain := datamodels.CommonPropertiesDomainObjectSTIX{
		SpecVersion:  "12.3.5", //change
		Created:      time.Now(),
		Modified:     time.Now(),
		CreatedByRef: datamodels.IdentifierTypeSTIX("ref_cubu8c3gf8g8g3f83"),
		Revoked:      false,
		Labels:       []string{"lable_1", "lable_2", "lable_3"}, //change
		Lang:         "RU",
		Сonfidence:   111, //change
		ExternalReferences: datamodels.ExternalReferencesTypeSTIX{
			datamodels.ExternalReferenceTypeElementSTIX{
				SourceName:  "source_name_1",
				Description: "just any descripton",
				URL:         "http://any-site.org/example_one",
				Hashes: datamodels.HashesTypeSTIX{
					"SHA-1":   "dcubdub883g3838fgc83f",
					"SHA-128": "cb8b38b8c38f83f888f844",
					"SHA-256": "cbudbcuf737fdv7f7ve7fv3f", //change
				},
				ExternalID: "12444_gddgdg",
			},
			datamodels.ExternalReferenceTypeElementSTIX{ //change
				SourceName:  "source_name_2",
				Description: "just any descripton two",
				URL:         "http://any-site.org/example_two",
				Hashes: datamodels.HashesTypeSTIX{
					"SHA-1":   "dcubdub883g3838fgc83f",
					"SHA-128": "cb8b38b8c38f83fas88f844", //change
					"SHA-256": "bccw62f626fd63fd63f6f36",
				},
				ExternalID: "12444_gddgdg",
			},
		},
		ObjectMarkingRefs: []datamodels.IdentifierTypeSTIX{
			datamodels.IdentifierTypeSTIX("obj_mark_ref_hcudcud-odcfc22y2vr34u41udv"), //change
			datamodels.IdentifierTypeSTIX("obj_mark_ref_ahvas223h23h2dh3h1h"),
			datamodels.IdentifierTypeSTIX("obj_mark_ref_nivnbufuf74gf74gf7s"),
		},
		Defanged: false,
		Extensions: map[string]string{
			"Key_1": "Value_1",
			"Key_2": "Value_2",
			"Key_3": "Value_3",
		},
	}

	Context("Тест 1. Проверка наличия установленного соединения с БД", func() {
		It("При установления соединения с БД ошибки быть не должно", func() {
			Expect(connectError).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Читаем тестовый файл содержащий STIX объекты", func() {
		It("При чтении тестового файла не должно быть ошибок", func() {
			Expect(errReadFile).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 3. Проверка на наличие ошибок при предварительном преобразовании из JSON", func() {
		It("Ошибок при предварительном преобразовании из JSON быть не должно", func() {
			Expect(errReadFile).ShouldNot(HaveOccurred())
			Expect(errUnmarchalReq).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 4. Проверяем функцию 'UnmarshalJSONObjectSTIXReq'", func() {
		It("должен быть получен список из 65 STIX объектов, ошибок быть не должно", func() {
			Expect(errUnmarchalToSTIX).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(65))
		})
	})

	/*Context("Тест 5. Взаимодействие с коллекцией STIX объектов", func() {
		It("При добавлении STIX объектов не должно быть ошибок", func() {
			ok, err := SetListElementSTIXObject(qp, l)

			Expect(ok).Should(BeTrue())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})*/

	Context("Тест 6. Получаем информацию о STIX объект с ID 'indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f'", func() {
		It("должен быть получен список из 1 STIX объекта, ошибок быть не должно", func() {
			cur, err := qp.Find(bson.D{bson.E{Key: "commonpropertiesobjectstix.id", Value: "indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f"}})

			l := []*datamodels.IndicatorDomainObjectsSTIX{}
			for cur.Next(context.Background()) {
				var model datamodels.IndicatorDomainObjectsSTIX
				_ = cur.Decode(&model)

				l = append(l, &model)
			}

			/*for _, v := range l {
				fmt.Printf("Found STIX object with ID:'indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f' - '%v'\n", *v)
			}*/

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(1))
		})
	})

	Context("Тест 7. Получаем информацию обо всех STIX объектах имеющих тип 'location'", func() {
		It("должен быть получен список из 3 STIX объектов, ошибок быть не должно", func() {
			cur, err := qp.Find(bson.D{bson.E{Key: "commonpropertiesobjectstix.type", Value: "location"}})

			l := []*datamodels.LocationDomainObjectsSTIX{}
			for cur.Next(context.Background()) {
				var model datamodels.LocationDomainObjectsSTIX
				_ = cur.Decode(&model)

				l = append(l, &model)
			}

			/*for _, v := range l {
				fmt.Printf("Found STIX object with Type: ID:'%s'\n'location' - '%v'\n", v.ID, *v)
			}*/

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(3))
		})
	})

	Context("Тест 8. Получаем информацию обо всех STIX объектах", func() {
		It("должен быть получен список из 65 STIX объектов, ошибок быть не должно", func() {
			cur, err := qp.Find(bson.D{})

			l := GetListElementSTIXObject(cur)

			/*
				fmt.Println("Test 8. Check result func'GetListElementSTIXObject'")
				for _, v := range l {
					fmt.Printf("ID STIX element: '%s'\n", v.Data.GetID())
				}
			*/

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(65))
		})
	})

	Context("Тест 9. Проверяем возможность перебора значений содержащихся в пользовательском типе", func() {
		It("При обработки тестовых значений не должно быть ошибок", func() {
			bomr := []datamodels.IdentifierTypeSTIX{
				datamodels.IdentifierTypeSTIX("obj_mark_ref_odcfc22y2vr34u41udv"),
				datamodels.IdentifierTypeSTIX("obj_mark_ref_ahvas223h23h2dh3h1h"),
				datamodels.IdentifierTypeSTIX("obj_mark_ref_nivnbufuf74gf74gf7s"),
			}

			before := datamodels.AttackPatternDomainObjectsSTIX{
				CommonPropertiesObjectSTIX: datamodels.CommonPropertiesObjectSTIX{
					Type: "attack pattern",
					ID:   "attack_pattern_vyw7d27dffd3ffd6f6fd6fw",
				},
				CommonPropertiesDomainObjectSTIX: datamodels.CommonPropertiesDomainObjectSTIX{
					SpecVersion:  "12.3.4",
					Created:      time.Now(),
					Modified:     time.Now(),
					CreatedByRef: datamodels.IdentifierTypeSTIX("ref_cubu8c3gf8g8g3f83"),
					Revoked:      false,
					Labels:       []string{"lable_1", "lable_5", "lable_3"},
					Lang:         "RU",
					Сonfidence:   12,
					ExternalReferences: datamodels.ExternalReferencesTypeSTIX{
						datamodels.ExternalReferenceTypeElementSTIX{
							SourceName:  "source_name_1",
							Description: "just any descripton",
							URL:         "http://any-site.org/example_one",
							Hashes: datamodels.HashesTypeSTIX{
								"SHA-1":   "dcubdub883g3838fgc83f",
								"SHA-128": "cb8b38b8c38f83f888f844",
							},
							ExternalID: "12444_gddgdg",
						},
						datamodels.ExternalReferenceTypeElementSTIX{
							SourceName:  "source_name_2",
							Description: "just any descripton two",
							URL:         "http://any-site.org/example_two",
							Hashes: datamodels.HashesTypeSTIX{
								"SHA-1":   "dcubdub883g3838fgc83f",
								"SHA-128": "cb8b38b8c38f83f888f844",
								"SHA-256": "bccw62f626fd63fd63f6f36",
							},
							ExternalID: "12444_gddgdg",
						},
					},
					ObjectMarkingRefs: bomr,
					Defanged:          false,
					Extensions: map[string]string{
						"Key_1": "Value_1",
						"Key_2": "Value_2",
						"Key_3": "Value_3",
					},
				},
				Name:            "attack pattern name 1",
				Description:     "test dscription name 443",
				Aliases:         []string{"Alias attack pattern", "Alias only attack pattern"},
				KillChainPhases: datamodels.KillChainPhasesTypeSTIX{},
			}

			after := datamodels.AttackPatternDomainObjectsSTIX{
				CommonPropertiesObjectSTIX: datamodels.CommonPropertiesObjectSTIX{
					Type: "attack pattern",
					ID:   "attack_pattern_vyw7d27dffd3ffd6f6fd6fw",
				},
				CommonPropertiesDomainObjectSTIX: datamodels.CommonPropertiesDomainObjectSTIX{
					SpecVersion:  "", //change
					Created:      time.Now(),
					Modified:     time.Now(),
					CreatedByRef: datamodels.IdentifierTypeSTIX("ref_cubu8c3gf8g8g3f83"),
					Revoked:      false,
					Labels:       []string{"lable_1", "lable_2", "lable_3"}, //change
					Lang:         "RU",
					Сonfidence:   3, //change
					ExternalReferences: datamodels.ExternalReferencesTypeSTIX{
						datamodels.ExternalReferenceTypeElementSTIX{
							SourceName:  "source_name_1",
							Description: "just any descripton",
							URL:         "http://any-site.org/example_one",
							Hashes: datamodels.HashesTypeSTIX{
								"SHA-1":   "dcubdub883g3838fgc83f",
								"SHA-128": "cb8b38b8c38f83f888f844",
							},
							ExternalID: "12444_gddgdg",
						},
						datamodels.ExternalReferenceTypeElementSTIX{
							SourceName:  "source_name_2",
							Description: "just any descripton two",
							URL:         "http://any-site.org/example_two",
							Hashes: datamodels.HashesTypeSTIX{
								"SHA-1":   "dcubdub883g3838fgc83f",
								"SHA-128": "cb8b38b8c38f83f888f844",
								"SHA-256": "bccw62f626fd63fd63f6f36",
							},
							ExternalID: "12444_gddgdg",
						},
					},
					ObjectMarkingRefs: bomr,
					Defanged:          false,
					Extensions: map[string]string{
						"Key_1": "Value_1",
						"Key_2": "Value_2",
						"Key_3": "Value_3",
					},
				},
				Name:            "attack pattern name 1 diferent",
				Description:     "test dscription name 443-124",
				Aliases:         []string{"Alias attack pattern"},
				KillChainPhases: datamodels.KillChainPhasesTypeSTIX{},
			}

			isEqual, _, err := ComparisonTypeCommonFields(before.CommonPropertiesDomainObjectSTIX, after.CommonPropertiesDomainObjectSTIX)

			/*fmt.Println("_______ func 'ComparisonTypeCommonFields', contrastResult ________")
			fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", contrastResult.CollectionName, contrastResult.DocumentID, contrastResult.ModifiedTime)
			for k, v := range contrastResult.FieldList {
				fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
			}
			fmt.Println("------------------------------------------------------------------")*/

			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 10. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Attack-pattern'", func() {
		It("При сравнении двух объектов типа 'Attack-pattern' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "attack pattern",
				ID:   "attack-pattern--vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.AttackPatternDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "attack pattern name 1",
				Description:                      "test dscription name 443",
				Aliases:                          []string{"Alias attack pattern", "Alias only attack pattern"},
				KillChainPhases:                  datamodels.KillChainPhasesTypeSTIX{},
			}

			isEqual, _ /*differentResult*/, err := apOld.ComparisonTypeCommonFields(&datamodels.AttackPatternDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "attack pattern name 1 //change", //change
				Description:                      "test dscription name ",          //change
				Aliases:                          []string{"Alias attack pattern", "Alias only attack pattern"},
				KillChainPhases:                  datamodels.KillChainPhasesTypeSTIX{},
			}, "test source")

			/*
				fmt.Println("_______ func 'AttackPatternDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())

			/*
					Запись в БД, информации об объектах одинакого типа, меющих различные значения полей
					пока отключим!!!
								//попробуем положить список изменений в БД
				qp := interactionmongodb.QueryParameters{
					NameDB:         "isems-mrsict",
					CollectionName: "accounting_differences_objects_collection",
					ConnectDB:      cdmdb.Connection,
				}

					_, errReqDB := qp.InsertData([]interface{}{differentResult})

					Expect(errReqDB).ShouldNot(HaveOccurred())
			*/
		})
	})

	Context("Тест 11. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Campaign'", func() {
		It("При сравнении двух объектов типа 'Campaign' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "campaign",
				ID:   "campaign--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.CampaignDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "campaing name 1",
				Description:                      "test dscription name 443",
				Aliases:                          []string{"Alias campaing name", "Alias only campaing"},
				FirstSeen:                        time.Now(),
				Objective:                        "1",
			}

			isEqual, _ /*differentResult*/, err := apOld.ComparisonTypeCommonFields(&datamodels.CampaignDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "attack pattern name 1 //change",                              //change
				Description:                      "test dscription name ",                                       //change
				Aliases:                          []string{"Alias attack pattern", "Alias only attack pattern"}, //change
				FirstSeen:                        time.Now(),                                                    //change
				Objective:                        "1 ndini cisicisivv",                                          //change
			}, "test source")

			/*
				fmt.Println("_______ func 'CampaignDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 12. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Course-of-action'", func() {
		It("При сравнении двух объектов типа 'Course-of-action' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "course-of-action",
				ID:   "course-of-action--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.CourseOfActionDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "campaing name 1",
				Description:                      "This is how to add a filter rule to block inbound access to TCP port 80 tothe existing UDP 1434 filter ...",
			}

			isEqual, _ /*differentResult*/, err := apOld.ComparisonTypeCommonFields(&datamodels.CourseOfActionDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "attack pattern name 1 //change", //change
				Description:                      "test dscription name ",          //change
			}, "test source")

			/*
				fmt.Println("_______ func 'CourseOfActionDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 13. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Grouping'", func() {
		It("При сравнении двух объектов типа 'Grouping' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "grouping",
				ID:   "grouping--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.GroupingDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "campaing name 1",
				Description:                      "This is how to add a filter rule to block inbound access to TCP port 80 tothe existing UDP 1434 filter ...",
				ObjectRefs:                       []datamodels.IdentifierTypeSTIX{"test string"},
			}

			isEqual, _ /*differentResult*/, err := apOld.ComparisonTypeCommonFields(&datamodels.GroupingDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "attack pattern name 1 //change",                      //change
				Description:                      "test dscription name ",                               //change
				ObjectRefs:                       []datamodels.IdentifierTypeSTIX{"newwww test string"}, //change
				Context:                          datamodels.OpenVocabTypeSTIX("open vocab string"),     //change
			}, "test source")

			/*
				fmt.Println("_______ func 'GroupingDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 13. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Identity'", func() {
		It("При сравнении двух объектов типа 'Identity' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "identity",
				ID:   "identity--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.IdentityDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "campaing name 1",
				Description:                      "This is how to add a filter rule to block inbound access to TCP port 80 tothe existing UDP 1434 filter ...",
				Roles:                            []string{"role one", "role two"},
				Sectors: []datamodels.OpenVocabTypeSTIX{
					datamodels.OpenVocabTypeSTIX("1111 only"),
					datamodels.OpenVocabTypeSTIX("two only"),
				},
				ContactInformation: "City and Country",
			}

			isEqual, _ /*differentResult*/, err := apOld.ComparisonTypeCommonFields(&datamodels.IdentityDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "attack pattern name 1 //change",               //change
				Description:                      "test dscription name ",                        //change
				Roles:                            []string{"role one", "role two", "role three"}, //change
				IdentityClass:                    datamodels.OpenVocabTypeSTIX("new value"),      //change
				ContactInformation:               "104234 City and Country",                      //change
			}, "test source")

			/*
				fmt.Println("_______ func 'IdentityDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			//Expect(len(differentResult.FieldList)).Should(Equal(7 + 6))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 14. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Indicator'", func() {
		It("При сравнении двух объектов типа 'Indicator' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "identity",
				ID:   "identity--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.IndicatorDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "campaing name 1",
				Description:                      "This is how to add a filter rule to block inbound access to TCP port 80 tothe existing UDP 1434 filter ...",
				IndicatorTypes:                   []datamodels.OpenVocabTypeSTIX{datamodels.OpenVocabTypeSTIX("hcscsbuc c")},
				ValidFrom:                        time.Now(),
				KillChainPhases: datamodels.KillChainPhasesTypeSTIX{datamodels.KillChainPhasesTypeElementSTIX{
					KillChainName: "kill chain 111",
					PhaseName:     "very good phase",
				}},
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.IndicatorDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "attack pattern name 1 //change", //change
				Description:                      "test dscription name ",          //change
				IndicatorTypes: []datamodels.OpenVocabTypeSTIX{
					datamodels.OpenVocabTypeSTIX("hcscsbuc c"),
					datamodels.OpenVocabTypeSTIX("hcsc")}, //change
				Pattern:        "pattern",                             //change
				PatternType:    datamodels.OpenVocabTypeSTIX("ggggg"), //change
				PatternVersion: "3.1",                                 //change
				ValidFrom:      time.Now(),                            //change
				//change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 8))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 15. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Infrastructure'", func() {
		It("При сравнении двух объектов типа 'Infrastructure' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "infrastructure",
				ID:   "infrastructure--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.InfrastructureDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "campaing name 1",
				Description:                      "This is how to add a filter rule to block inbound access to TCP port 80 tothe existing UDP 1434 filter ...",
				InfrastructureTypes:              []datamodels.OpenVocabTypeSTIX{"gssds dcc"},
				Aliases:                          []string{"aliase_11", "aliase_22"},
				KillChainPhases: datamodels.KillChainPhasesTypeSTIX{datamodels.KillChainPhasesTypeElementSTIX{
					KillChainName: "kill chain 111",
					PhaseName:     "very good phase",
				}},
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.InfrastructureDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "attack pattern name 1 //change", //change
				Description:                      "test dscription name ",          //change
				//change
				//change
				//change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 5))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 15. Проверяем возможность перебора значений содержащихся в пользовательском типе 'IntrusionSet'", func() {
		It("При сравнении двух объектов типа 'IntrusionSet' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "intrusion-set",
				ID:   "intrusion-set--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.IntrusionSetDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "campaing name 1",
				Description:                      "This is how to add a filter rule to block inbound access to TCP port 80 tothe existing UDP 1434 filter ...",
				Aliases:                          []string{"aliase_11", "aliase_22"},
				FirstSeen:                        time.Now(),
				LastSeen:                         time.Now(),
				Goals:                            []string{"cbcd", "dfffbf"},
				ResourceLevel:                    datamodels.OpenVocabTypeSTIX("asssss"),
				SecondaryMotivations:             []datamodels.OpenVocabTypeSTIX{"xxsdsc", "cvc2e2"},
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.IntrusionSetDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "attack pattern name 1 //change",   //change
				Description:                      "test dscription name ",            //change
				Aliases:                          []string{"aliase_11", "aliase_23"}, //change
				FirstSeen:                        time.Now(),                         //change
				LastSeen:                         time.Now(),                         //change
				//change
				//change
				PrimaryMotivation: datamodels.OpenVocabTypeSTIX("ccvf"), //change
				//change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 9))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 16. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Location'", func() {
		It("При сравнении двух объектов типа 'Location' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "location",
				ID:   "location--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.LocationDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "campaing name 1",
				Description:                      "This is how to add a filter rule to block inbound access to TCP port 80 tothe existing UDP 1434 filter ...",
				Latitude:                         8.23,
				Longitude:                        23.44,
				Precision:                        12.34,
				Region:                           datamodels.OpenVocabTypeSTIX("cbhdbhd"),
				Country:                          "EU",
				AdministrativeArea:               "contry name",
				City:                             "St. Petersburg",
				StreetAddress:                    "123322 M, dsfd",
				PostalCode:                       "234343",
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.LocationDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "attack pattern name 1 //change",       //change
				Description:                      "test dscription name ",                //change
				Latitude:                         21.2,                                   //change
				Longitude:                        3.14,                                   //change
				Precision:                        32.1224,                                //change
				Region:                           datamodels.OpenVocabTypeSTIX("ccvdvd"), //change
				Country:                          "RU",                                   //change
				AdministrativeArea:               "contry name and city",                 //change
				City:                             "Moscow",                               //change
				StreetAddress:                    "244311 Moscow, dsfd",                  //change
				PostalCode:                       "125235",                               //change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 11))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 17. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Malware'", func() {
		It("При сравнении двух объектов типа 'Malware' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "malware",
				ID:   "malware--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.MalwareDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "campaing name 1",
				Description:                      "This is how to add a filter rule to block inbound access to TCP port 80 tothe existing UDP 1434 filter ...",
				MalwareTypes:                     []datamodels.OpenVocabTypeSTIX{datamodels.OpenVocabTypeSTIX("bbcc777bbcx")},
				Aliases:                          []string{"1", "2", "3"},
				FirstSeen:                        time.Now(),
				OperatingSystemRefs:              []datamodels.IdentifierTypeSTIX{datamodels.IdentifierTypeSTIX("qwe")},
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.MalwareDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "attack pattern name 1 //change", //change
				Description:                      "test dscription name ",          //change
				//change
				IsFamily:  true,                         //change
				Aliases:   []string{"1", "2", "4", "5"}, //change
				FirstSeen: time.Now(),                   //change
				OperatingSystemRefs: []datamodels.IdentifierTypeSTIX{
					datamodels.IdentifierTypeSTIX("qwe"),
					datamodels.IdentifierTypeSTIX("zxvbbb"), //change
				},
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 7))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 18. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Malware-analysis'", func() {
		It("При сравнении двух объектов типа 'Malware-analysis' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "malware-analysis",
				ID:   "malware-analysis--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.MalwareAnalysisDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Product:                          "Application Soft",
				Version:                          "1.23.1",
				HostVMRef:                        datamodels.IdentifierTypeSTIX("2sss 33"),
				OperatingSystemRef:               datamodels.IdentifierTypeSTIX("bbv nccn hh"),
				ConfigurationVersion:             "v12.1",
				Modules:                          []string{"x1", "x2"},
				AnalysisEngineVersion:            "eng version 12.3",
				AnalysisDefinitionVersion:        "def version 65.1",
				Submitted:                        time.Now(),
				AnalysisStarted:                  time.Now(),
				AnalysisEnded:                    time.Now(),
				ResultName:                       "mgjas",
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.MalwareAnalysisDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Product:                          "App Soft",                           //change
				Version:                          "1.24.56",                            //change
				HostVMRef:                        datamodels.IdentifierTypeSTIX("cv5"), //change
				//change
				ConfigurationVersion:      "v12.61",                          //change
				Modules:                   []string{"x1", "x2", "v3", "v12"}, //change
				AnalysisEngineVersion:     "rus version 12.3",                //change
				AnalysisDefinitionVersion: "def version 65.111",              //change
				Submitted:                 time.Now(),                        //change
				AnalysisStarted:           time.Now(),                        //change
				AnalysisEnded:             time.Now(),                        //change
				//change
				AnalysisScoRefs: []datamodels.IdentifierTypeSTIX{datamodels.IdentifierTypeSTIX("bbbb")}, //change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 13))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 19. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Note'", func() {
		It("При сравнении двух объектов типа 'Note' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "note",
				ID:   "note--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.NoteDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Abstract:                         "ever never talk",
				Content:                          "nnb vvn",
				Authors:                          []string{"nnm", "cvb"},
				ObjectRefs:                       []datamodels.IdentifierTypeSTIX{datamodels.IdentifierTypeSTIX("mongodb")},
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.NoteDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Abstract:                         "only close",                  //change
				Content:                          "c",                           //change
				Authors:                          []string{"nnm", "cvb", "yyy"}, //change
				ObjectRefs: []datamodels.IdentifierTypeSTIX{
					datamodels.IdentifierTypeSTIX("mongodb"),
					datamodels.IdentifierTypeSTIX("react"),
					datamodels.IdentifierTypeSTIX("angular"),
				}, //change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 4))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 20. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Observed-data'", func() {
		It("При сравнении двух объектов типа 'Observed-data' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "observed-data",
				ID:   "observed-data--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.ObservedDataDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				NumberObserved:                   90,
				ObjectRefs:                       []datamodels.IdentifierTypeSTIX{datamodels.IdentifierTypeSTIX("mongodb")},
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.ObservedDataDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				FirstObserved:                    time.Now(), //change
				LastObserved:                     time.Now(), //change
				NumberObserved:                   92,         //change
				ObjectRefs: []datamodels.IdentifierTypeSTIX{
					datamodels.IdentifierTypeSTIX("mongodb"),
					datamodels.IdentifierTypeSTIX("react"),
					datamodels.IdentifierTypeSTIX("angular"),
				}, //change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 4))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 21. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Opinion'", func() {
		It("При сравнении двух объектов типа 'Opinion' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "opinion",
				ID:   "opinion--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.OpinionDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Explanation:                      "notes",
				Authors:                          []string{"authors_5", "authors_3"},
				Opinion:                          datamodels.EnumTypeSTIX("vbbb"),
				ObjectRefs:                       []datamodels.IdentifierTypeSTIX{datamodels.IdentifierTypeSTIX("mongodb")},
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.OpinionDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Explanation:                      "notes ddd",                        //change
				Authors:                          []string{"authors_3"},              //change
				Opinion:                          datamodels.EnumTypeSTIX("11 vbbb"), //change
				//change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 4))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 22. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Report'", func() {
		It("При сравнении двух объектов типа 'Report' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "report",
				ID:   "report--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.ReportDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "new name 1",
				Description:                      "new description 1",
				Published:                        time.Now(),
				ObjectRefs:                       []datamodels.IdentifierTypeSTIX{datamodels.IdentifierTypeSTIX("mongodb")},
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.ReportDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "new name 122",                                                                 //change
				Description:                      "new description 1333",                                                         //change
				Published:                        time.Now(),                                                                     //change
				ObjectRefs:                       []datamodels.IdentifierTypeSTIX{datamodels.IdentifierTypeSTIX("mongodb 1111")}, //change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 4))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 23. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Threat-actor'", func() {
		It("При сравнении двух объектов типа 'Threat-actor' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "threat-actor",
				ID:   "threat-actor--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.ThreatActorDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "new name 1",
				Description:                      "new description 1",
				ThreatActorTypes: []datamodels.OpenVocabTypeSTIX{
					datamodels.OpenVocabTypeSTIX("bbnd d"),
					datamodels.OpenVocabTypeSTIX("mnbfdd"),
				},
				Aliases:   []string{"cvvv", "vbbb"},
				FirstSeen: time.Now(),
				Roles:     []datamodels.OpenVocabTypeSTIX{datamodels.OpenVocabTypeSTIX("mmm")},
				Goals:     []string{"1cc1", "22sc"},
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.ThreatActorDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "new name 122",                                                         //change
				Description:                      "new description 1333",                                                 //change
				ThreatActorTypes:                 []datamodels.OpenVocabTypeSTIX{datamodels.OpenVocabTypeSTIX("mnbfdd")}, //change
				Aliases:                          []string{"cvvv", "vbbb", "vvvv"},                                       //change
				FirstSeen:                        time.Now(),                                                             //change
				LastSeen:                         time.Now(),                                                             //change
				Roles:                            []datamodels.OpenVocabTypeSTIX{datamodels.OpenVocabTypeSTIX("m000m")},  //change
				Goals:                            []string{"1cc1"},                                                       //change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 8))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 23. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Tool'", func() {
		It("При сравнении двух объектов типа 'Tool' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "tool",
				ID:   "tool--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.ToolDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "new name 1",
				Description:                      "new description 1",
				Aliases:                          []string{"cvvv", "vbbb"},
				ToolVersion:                      "v23.1",
				ToolTypes:                        []datamodels.OpenVocabTypeSTIX{datamodels.OpenVocabTypeSTIX("ikk")},
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.ToolDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "new name 122",                   //change
				Description:                      "new description 1333",           //change
				Aliases:                          []string{"cvvv", "vbbb", "vvvv"}, //change
				ToolVersion:                      "v30.23",                         //change
				ToolTypes: []datamodels.OpenVocabTypeSTIX{
					datamodels.OpenVocabTypeSTIX("ikk"),
					datamodels.OpenVocabTypeSTIX("iopp"),
					datamodels.OpenVocabTypeSTIX("yyu"),
				}, //change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 5))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 24. Проверяем возможность перебора значений содержащихся в пользовательском типе 'Vulnerability'", func() {
		It("При сравнении двух объектов типа 'Vulnerability' содержащих некоторые разные данные переменная isEqual должна быть FALSE", func() {
			cp := datamodels.CommonPropertiesObjectSTIX{
				Type: "vulnerability",
				ID:   "vulnerability--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
			}

			apOld := datamodels.VulnerabilityDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
				Name:                             "new name 1",
				Description:                      "new description 1",
			}

			isEqual, differentResult, err := apOld.ComparisonTypeCommonFields(&datamodels.VulnerabilityDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       cp,
				CommonPropertiesDomainObjectSTIX: newCommonPropertiesDomain,
				Name:                             "new name 122",         //change
				Description:                      "new description 1333", //change
			}, "test source")

			/*
				fmt.Println("_______ func 'IndicatorDomainObjectsSTIX', contrastResult ________")
				fmt.Printf("Collection name: '%s'\nDocument ID: '%s'\nModified time: '%v'\n", differentResult.CollectionName, differentResult.DocumentID, differentResult.ModifiedTime)
				for k, v := range differentResult.FieldList {
					fmt.Printf("Key: %d\n\tFeildType: '%s'\n\tPath: '%s'\n\tValue: '%v'\n", k, v.FeildType, v.Path, v.Value)
				}
				fmt.Println("------------------------------------------------------------------")
			*/

			Expect(len(differentResult.FieldList)).Should(Equal(7 + 2))
			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	/*
		Написал функции apOld.ComparisonTypeCommonFields() для ВСЕХ объектов типа DomainObject и протестировал их, теперь
		могу выполнять сравнения двух объектов подобного типа.

		Кроме того есть функция GetListElementSTIXObject() возвращающая список STIX объектов в виде []*datamodels.ElementSTIXObject

		!!! Теперь надо написать функцию ReplacementElementsSTIXObject() которая должна выполнять замену, в БД, STIX объекта
		если он существует или добавлять новый. По идее она должна сначала удалить все объекты которые планируются заменить,
		а потом добавить новые. !!!
		Проверка, требует ли объект замены выполняется, в том числе, через метод ComparisonTypeCommonFields
	*/

	/*
			Context("", func() {
			It("", func(){

			})
		})
	*/
})

func ComparisonTypeCommonFields(before, after datamodels.CommonPropertiesDomainObjectSTIX) (bool, datamodels.DifferentObjectType, error) {
	var isEqual bool = true
	tmpOne := reflect.ValueOf(before)
	typeOfSOne := tmpOne.Type()

	tmpTwo := reflect.ValueOf(after)
	typeOfSTwo := tmpTwo.Type()

	fmt.Println("---=== func 'ComparisonTypeCommonFields' ===---")

	contrast := datamodels.DifferentObjectType{
		ModifiedTime: time.Now(),
		FieldList:    []datamodels.OldFieldValueObjectType{},
	}

	for i := 0; i < tmpOne.NumField(); i++ {
		for j := 0; j < tmpTwo.NumField(); j++ {
			resultDeepEqual := reflect.DeepEqual(tmpOne.Field(i).Interface(), tmpTwo.Field(j).Interface())

			if typeOfSOne.Field(i).Name == typeOfSTwo.Field(j).Name {

				fmt.Printf("Field: %s\tValue BEFORE: %v, AFTER: %v, Equal: %v\n", typeOfSOne.Field(i).Name, tmpOne.Field(i).Interface(), tmpTwo.Field(j).Interface(), resultDeepEqual)

				if !resultDeepEqual {
					contrast.FieldList = append(contrast.FieldList, datamodels.OldFieldValueObjectType{
						FeildType: typeOfSOne.Field(i).Type.Name(),
						Path:      typeOfSOne.Field(i).Name,
						Value:     tmpOne.Field(i).Interface(),
					})

					isEqual = false
				}
			}
		}
	}

	return isEqual, contrast, nil
}

type definingTypeSTIXObject struct {
	datamodels.CommonPropertiesObjectSTIX
}

func SetListElementSTIXObject(qp interactionmongodb.QueryParameters, l []*datamodels.ElementSTIXObject) (bool, error) {
	sizeListSTIXObject := len(l)

	/*
		ПОКА ПРИДЕТСЯ ЗАКОМЕНТИТЬ, надо сделать функцию GetListElementSTIXObject
		//получаем список ID STIX объектов для их поиска в БД
		reqSearchID := primitive.A{}
		for _, v := range l {
			reqSearchID = append(reqSearchID, bson.D{{Key: "commonpropertiesobjectstix.id", Value: v.Data.GetID()}})
		}

		//делаем запрос к БД на ниличие STIX объектов которые хотим добавить
		cur, err := qp.Find(bson.D{{Key: "$or", Value: reqSearchID}})
		if err != nil {
			return false, err
		}

		listFromDB := GetListElementSTIXObject(cur)

		fmt.Printf("func 'SetListElementSTIXObject', count list from DB: '%d'\n", len(listFromDB))

		//сравниваем все параметры STIX объектов полученных из БД и STIX объектов которые мы хотим добавить
	*/

	//добавляем только те элементы которых нет в БД и делаем обновление тех которые изменились
	addInfo := make([]interface{}, 0, sizeListSTIXObject)
	for _, v := range l {

		//				fmt.Printf("___ STIX object ID:'%s'\n", v.Data.GetID())

		addInfo = append(addInfo, v.Data)
	}

	return qp.InsertData(addInfo)
}

func GetListElementSTIXObject(cur *mongo.Cursor) []*datamodels.ElementSTIXObject {
	elements := []*datamodels.ElementSTIXObject{}

	for cur.Next(context.Background()) {
		var modelType definingTypeSTIXObject
		if err := cur.Decode(&modelType); err != nil {
			continue
		}

		//fmt.Printf("func 'GetListElementSTIXObject', type STIX object: '%s', ID: '%s'\n", modelType.Type, modelType.ID)

		switch modelType.Type {
		/* *** Domain Objects STIX *** */
		case "attack-pattern":
			tmpObj := datamodels.AttackPatternDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})
		case "campaign":
			tmpObj := datamodels.CampaignDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "course-of-action":
			tmpObj := datamodels.CourseOfActionDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "grouping":
			tmpObj := datamodels.GroupingDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "identity":
			tmpObj := datamodels.IdentityDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "indicator":
			tmpObj := datamodels.IndicatorDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "infrastructure":
			tmpObj := datamodels.InfrastructureDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "intrusion-set":
			tmpObj := datamodels.IntrusionSetDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "location":
			tmpObj := datamodels.LocationDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "malware":
			tmpObj := datamodels.MalwareDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "malware-analysis":
			tmpObj := datamodels.MalwareAnalysisDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "note":
			tmpObj := datamodels.NoteDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "observed-data":
			tmpObj := datamodels.ObservedDataDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "opinion":
			tmpObj := datamodels.OpinionDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "report":
			tmpObj := datamodels.ReportDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "threat-actor":
			tmpObj := datamodels.ThreatActorDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "tool":
			tmpObj := datamodels.ToolDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "vulnerability":
			tmpObj := datamodels.VulnerabilityDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		/* *** Relationship Objects *** */
		case "relationship":
			tmpObj := datamodels.RelationshipObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "sighting":
			tmpObj := datamodels.SightingObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		/* *** Cyber-observable Objects STIX *** */
		case "artifact":
			tmpObj := datamodels.ArtifactCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "autonomous-system":
			tmpObj := datamodels.AutonomousSystemCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "directory":
			tmpObj := datamodels.DirectoryCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "domain-name":
			tmpObj := datamodels.DomainNameCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "email-addr":
			tmpObj := datamodels.EmailAddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "email-message":
			tmpObj := datamodels.EmailMessageCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "file":
			tmpObj := datamodels.FileCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "ipv4-addr":
			tmpObj := datamodels.IPv4AddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "ipv6-addr":
			tmpObj := datamodels.IPv6AddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "mac-addr":
			tmpObj := datamodels.MACAddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "mutex":
			tmpObj := datamodels.MutexCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "network-traffic":
			tmpObj := datamodels.NetworkTrafficCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "process":
			tmpObj := datamodels.ProcessCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "software":
			tmpObj := datamodels.SoftwareCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "url":
			tmpObj := datamodels.URLCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "user-account":
			tmpObj := datamodels.UserAccountCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "windows-registry-key":
			tmpObj := datamodels.WindowsRegistryKeyCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "x509-certificate":
			tmpObj := datamodels.X509CertificateCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})
		}
	}

	return elements
}
