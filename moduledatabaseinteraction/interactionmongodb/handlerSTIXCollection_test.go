package interactionmongodb_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"ISEMS-MRSICT/commonhandlers"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

func getListSettings(f string, appConfig *datamodels.AppConfig) (map[string]datamodels.StorageApplicationCommonListType, error) {
	tmp := map[string]string{}
	configFileSettings := map[string]datamodels.StorageApplicationCommonListType{}

	//проверяем наличие файлов с дефолтными настройками приложения
	row, err := ioutil.ReadFile(path.Join(appConfig.RootDir, f))
	if err != nil {
		return configFileSettings, fmt.Errorf("Error! The file '%s' with default settings not found.", f)
	}

	err = json.Unmarshal(row, &tmp)
	if err != nil {
		return configFileSettings, err
	}

	for k, v := range tmp {
		configFileSettings[k] = datamodels.StorageApplicationCommonListType{Description: v}
	}

	return configFileSettings, err
}

var _ = Describe("HandlerSTIXCollection", func() {
	var (
		connectError, errReadFile, errUnmarchalReq, errUnmarchalToSTIX error
		docJSON                                                        []byte
		cdmdb                                                          interactionmongodb.ConnectionDescriptorMongoDB
		l                                                              []*datamodels.ElementSTIXObject
		qp                                                             interactionmongodb.QueryParameters
		modAPIRequestProcessingReqJSON                                 datamodels.ModAPIRequestProcessingReqJSON
		tst                                                            *memorytemporarystoragecommoninformation.TemporaryStorageType
		appConfig                                                      datamodels.AppConfig
	)

	var _ = BeforeSuite(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

		cdmdb = interactionmongodb.ConnectionDescriptorMongoDB{
			Ctx:       ctx,
			CtxCancel: cancel,
		}

		//подключаемся к базе данных MongoDB
		connectError = cdmdb.CreateConnection(&datamodels.MongoDBSettings{
			Host:     "192.168.13.201",
			Port:     27017,
			User:     "module-isems-mrsict",
			Password: "vkL6Zn$jPmt1e1",
			NameDB:   "isems-mrsict",
		})

		docJSON, errReadFile = ioutil.ReadFile("../../mytest/test_resources/jsonSTIXExample.json")
		errUnmarchalReq = json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON)
		l, errUnmarchalToSTIX = routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectSTIXReq(modAPIRequestProcessingReqJSON)

		qp = interactionmongodb.QueryParameters{
			NameDB:         "isems-mrsict",
			CollectionName: "stix_object_collection",
			ConnectDB:      cdmdb.Connection,
		}

		appConfig.RootDir = "/Users/user/go/src/ISEMS-MRSICT"

		tst = memorytemporarystoragecommoninformation.NewTemporaryStorage()
		/* получаем и сохраняем во временном хранилище дефолтные настройки приложения */
		ssdmct, _ := getListSettings("defaultsettingsfiles/settingsStatusesDecisionsMadeComputerThreats.json", &appConfig)
		tst.SetListDecisionsMade(ssdmct)

		sctt, _ := getListSettings("defaultsettingsfiles/settingsComputerThreatTypes.json", &appConfig)
		tst.SetListComputerThreat(sctt)

		fmt.Println("============")
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		fmt.Printf("DIR path: '%s'\n", dir)
		fmt.Println(tst.GetListDecisionsMade())
		//fmt.Println(tst.GetListComputerThreat())
		fmt.Println("============")
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
		It("должен быть получен список из 66 STIX объектов, ошибок быть не должно", func() {
			Expect(errUnmarchalToSTIX).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(66))
		})
	})

	Context("Тест 5. Взаимодействие с коллекцией STIX объектов.", func() {
		It("При добавлении STIX объектов не должно быть ошибок. STIX объекты идентификаторы которых уже есть в БД добавлятся не должны.", func() {
			//получаем список ID STIX объектов предназначенных для добавление в БД
			listID := commonhandlers.GetListIDFromListSTIXObjects(l)

			fmt.Println("-----------=====================================================----------")

			countBefore := len(l)

			//выполняем запрос к БД, для получения полной информации об STIX объектах по их ID
			listElemetSTIXObject, err := interactionmongodb.FindSTIXObjectByID(qp, listID)
			Expect(err).ShouldNot(HaveOccurred())

			routingflowsmoduleapirequestprocessing.VerifyOutsideSpecificationFields(l, tst, "client-test")
			l := interactionmongodb.SavingAdditionalNameListSTIXObject(listElemetSTIXObject, l)

			countAfter := len(l)

			err = interactionmongodb.ReplacementElementsSTIXObject(qp, l)

			fmt.Println("-----------=====================================================----------")

			Expect(err).ShouldNot(HaveOccurred())
			Expect(countBefore).Should(Equal(countAfter))
		})
	})

	Context("Тест 6. Получаем информацию о STIX объект с ID 'indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f'", func() {
		It("должен быть получен список из 1 STIX объекта, ошибок быть не должно", func() {
			cur, err := qp.Find(bson.D{bson.E{Key: "commonpropertiesobjectstix.id", Value: "indicator--d38a99ae-c5ee-4542-bc12-dfe68b48cc08"}})

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
		It("должен быть получен список из 66 STIX объектов, ошибок быть не должно", func() {
			var objID primitive.A

			for _, v := range l {
				objID = append(objID, v.Data.GetID())
			}

			cur, err := qp.Find((bson.D{{Key: "commonpropertiesobjectstix.id", Value: bson.D{{Key: "$in", Value: objID}}}}))
			lr := interactionmongodb.GetListElementSTIXObject(cur)

			fmt.Println("Test 8. Check result func'GetListElementSTIXObject'")
			for _, v := range l {
				if v.Data.GetType() == "ipv4-addr" {
					if ipv4, ok := v.Data.(datamodels.IPv4AddressCyberObservableObjectSTIX); ok {
						fmt.Printf("\t\t---ID STIX element ipv4-addr: '%v'\n", ipv4)
					}
				}
			}

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(lr)).Should(Equal(66))
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

	Context("Тест 25. Проверяем возможность перебора значений для списка STIX объектов", func() {
		It("При переборе значений ошибок быть не должно и при выполнении сравнения должны быть найдены одинаковые объекты содержащие разные значения", func() {
			listOldSTIXObj := []datamodels.ElementSTIXObject{
				{
					DataType: "tool",
					Data: datamodels.ToolDomainObjectsSTIX{
						CommonPropertiesObjectSTIX: datamodels.CommonPropertiesObjectSTIX{
							Type: "tool",
							ID:   "tool--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
						},
						CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
						Name:                             "new name 1",
						Description:                      "new description 1",
						Aliases:                          []string{"cvvv", "vbbb"},
						ToolVersion:                      "v23.1",
						ToolTypes:                        []datamodels.OpenVocabTypeSTIX{datamodels.OpenVocabTypeSTIX("ikk")},
					},
				}, {
					DataType: "threat-actor",
					Data: datamodels.ThreatActorDomainObjectsSTIX{
						CommonPropertiesObjectSTIX: datamodels.CommonPropertiesObjectSTIX{
							Type: "threat-actor",
							ID:   "threat-actor--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
						},
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
					},
				}, {
					DataType: "opinion",
					Data: datamodels.OpinionDomainObjectsSTIX{
						CommonPropertiesObjectSTIX: datamodels.CommonPropertiesObjectSTIX{
							Type: "opinion",
							ID:   "opinion--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
						},
						CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
						Explanation:                      "notes",
						Authors:                          []string{"authors_5", "authors_3"},
						Opinion:                          datamodels.EnumTypeSTIX("vbbb"),
						ObjectRefs:                       []datamodels.IdentifierTypeSTIX{datamodels.IdentifierTypeSTIX("mongodb")},
					},
				},
			}

			listNewSTIXObj := []datamodels.ElementSTIXObject{
				{
					DataType: "opinion",
					Data: datamodels.OpinionDomainObjectsSTIX{
						CommonPropertiesObjectSTIX: datamodels.CommonPropertiesObjectSTIX{
							Type: "opinion",
							ID:   "opinion--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
						},
						CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
						Explanation:                      "notes ddd",                        //change
						Authors:                          []string{"authors_3"},              //change
						Opinion:                          datamodels.EnumTypeSTIX("11 vbbb"), //change
					},
				}, {
					DataType: "threat-actor",
					Data: datamodels.ThreatActorDomainObjectsSTIX{
						CommonPropertiesObjectSTIX: datamodels.CommonPropertiesObjectSTIX{
							Type: "threat-actor",
							ID:   "threat-actor--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
						},
						CommonPropertiesDomainObjectSTIX: oldCommonPropertiesDomain,
						Name:                             "new name 122",                                                         //change
						Description:                      "new description 1333",                                                 //change
						ThreatActorTypes:                 []datamodels.OpenVocabTypeSTIX{datamodels.OpenVocabTypeSTIX("mnbfdd")}, //change
						Aliases:                          []string{"cvvv", "vbbb", "vvvv"},                                       //change
						FirstSeen:                        time.Now(),                                                             //change
						LastSeen:                         time.Now(),                                                             //change
						Roles:                            []datamodels.OpenVocabTypeSTIX{datamodels.OpenVocabTypeSTIX("m000m")},  //change
						Goals:                            []string{"1cc1"},                                                       //change
					},
				}, {
					DataType: "tool",
					Data: datamodels.ToolDomainObjectsSTIX{
						CommonPropertiesObjectSTIX: datamodels.CommonPropertiesObjectSTIX{
							Type: "tool",
							ID:   "tool--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
						},
						Name:        "new name 122",                   //change
						Description: "new description 1333",           //change
						Aliases:     []string{"cvvv", "vbbb", "vvvv"}, //change
						ToolVersion: "v30.23",                         //change
						ToolTypes: []datamodels.OpenVocabTypeSTIX{
							datamodels.OpenVocabTypeSTIX("ikk"),
							datamodels.OpenVocabTypeSTIX("iopp"),
							datamodels.OpenVocabTypeSTIX("yyu"),
						}, //change
					},
				}, {
					DataType: "malware-analysis",
					Data: datamodels.MalwareAnalysisDomainObjectsSTIX{
						CommonPropertiesObjectSTIX: datamodels.CommonPropertiesObjectSTIX{
							Type: "malware-analysis",
							ID:   "malware-analysis--bcbdd-vyw7d27dffd3ffd6f6fd6fw",
						},
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
					},
				},
			}

			listDifferentResult := ComparasionListSTIXObject(ComparasionListTypeSTIXObject{
				OldList: listOldSTIXObj,
				NewList: listNewSTIXObj,
			})

			Expect(len(listDifferentResult)).Should(Equal(3))
		})
	})

	Context("Тест 26. Получаем информацию о ПЕРВЫХ 10 STIX объектах при выполнении поиска БЕЗ поисковых параметров", func() {
		It("Должен быть получен список из ПЕРВЫХ 10 объектов, ошибки при этом быть не должно", func() {
			cur, err := qp.FindAllWithLimit(bson.D{}, &interactionmongodb.FindAllWithLimitOptions{
				Offset:        1,
				LimitMaxSize:  10,
				SortAscending: false,
			})

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			fmt.Println("____ first 10 elements ___")
			for _, v := range elemSTIXObj {
				fmt.Printf("\tid: '%s'\n", v.Data.GetID())
			}

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(elemSTIXObj)).Should(Equal(10))
		})
	})

	Context("Тест 27. Получаем информацию о ВТОРЫХ 10 STIX объектах при выполнении поиска БЕЗ поисковых параметров", func() {
		It("Должен быть получен список из ВТОРЫХ 10 объектов, ошибки при этом быть не должно", func() {
			cur, err := qp.FindAllWithLimit(bson.D{}, &interactionmongodb.FindAllWithLimitOptions{
				Offset:        2,
				LimitMaxSize:  10,
				SortAscending: false,
			})

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			fmt.Println("____ second 10 elements ___")
			for _, v := range elemSTIXObj {
				fmt.Printf("\tid: '%s'\n", v.Data.GetID())
			}

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(elemSTIXObj)).Should(Equal(10))
		})
	})

	Context("Тест 28. Получаем информацию о ПОСЛЕДНИЕ 10 STIX объектах при выполнении поиска БЕЗ поисковых параметров", func() {
		It("Должен быть получен список из ПОСЛЕДНИХ 4 объектов, ошибки при этом быть не должно", func() {
			cur, err := qp.FindAllWithLimit(bson.D{}, &interactionmongodb.FindAllWithLimitOptions{
				Offset:        11,
				LimitMaxSize:  10,
				SortAscending: false,
			})

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			fmt.Println("____ last 10 elements ___")
			for _, v := range elemSTIXObj {
				fmt.Printf("\tid: '%s'\n", v.Data.GetID())
			}

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(elemSTIXObj)).Should(Equal(4))
		})
	})

	Context("Тест 29. Получаем ВСЮ информацию STIX объектах при выполнении поиска БЕЗ поисковых параметров", func() {
		It("Должен быть получен ВЕСЬ список ОТСОРТИРОВАННЫХ объектов, ошибки при этом быть не должно", func() {
			sizeElem, err := qp.CountDocuments(bson.D{})

			Expect(err).ShouldNot(HaveOccurred())
			Expect(sizeElem).Should(Equal(int64(104)))
		})
	})

	/*
			Context("", func() {
			It("", func(){

			})
		})
	*/
})

/*
		1. id: 'x509-certificate--b595eaf0-0b28-5dad-9e8e-0fab9c1facc9'
        2. id: 'x509-certificate--463d7b2a-8516-5a50-a3d7-6f801465d5de'
        3. id: 'windows-registry-key--2ba37ae7-2745-5082-9dfd-9486dad41016'
        4. id: 'user-account--0d5b424b-93b8-5cd8-ac36-306e1789d63c'
        5. id: 'user-account--0d5b424b-93b8-5cd8-ac36-306e1789d63c'
        6. id: 'url--c1477287-23ac-5971-a010-5c287877fa60'
        7. id: 'software--a1827f6d-ca53-5605-9e93-4316cd22a00a'
        8. id: 'process--99ab297d-4c39-48ea-9d64-052d596864df'
        9. id: 'process--07bc30cad-ebc2-4579-881d-b9cdc7f2b33c'
        10. id: 'process--f52a906a-0dfc-40bd-92f1-e7778ead38a9'
        11. id: 'network-traffic--09ca55c3-97e5-5966-bad0-1d41d557ae13'
        12. id: 'network-traffic--15a157a8-26e3-56e0-820b-0c2a8e553a2c'
        13. id: 'network-traffic--c95e972a-20a4-5307-b00d-b8393faf02c5'
        14. id: 'network-traffic--e7a939ca-78c6-5f27-8ae0-4ad112454626'
        15. id: 'network-traffic--f8ae967a-3dc3-5cdf-8f94-8505abff00c2'
        16. id: 'network-traffic--ac267abc-1a41-536d-8e8d-98458d9bf491'
        17. id: 'network-traffic--630d7bb1-0bbc-53a6-a6d4-f3c2d35c2734'
        18. id: 'network-traffic--15a157a8-26e3-56e0-820b-0c2a1e553a2c'
        19. id: 'mutex--eba44954-d4e4-5d3b-814c-2b17dd8de300'
        20. id: 'mac-addr--65cfcf98-8a6e-5a1b-8f61-379ac4f92d00'
        21. id: 'ipv6-addr--5daf7456-8863-5481-9d42-237d477697f4'
        22. id: 'ipv6-addr--1e61d36c-a16c-53b7-a80f-2a00161c96b1'
        23. id: 'ipv4-addr--5853f6a4-638f-5b4e-9b0f-ded361ae3812'
        24. id: 'ipv4-addr--ff26c055-6336-5bc5-b98d-13d6226742dd'
        25. id: 'file--fb0419a8-f09c-57f8-be64-71a80417591c'
        26. id: 'file--e277603e-1060-5ad4-9937-c26c97f1ca68'
        27. id: 'file--ec3415cc-5f4f-5ec8-bdb1-6f86996ae66d'
        28. id: 'file--c7d1e135-8b34-549a-bb47-302f5cf998ed'
        29. id: 'file--73c4cd13-7206-5100-88ef-822c42d3f02c'
        30. id: 'file--9a1f834d-2506-5367-baec-7aa63996ac43'
        31. id: 'file--e277603e-1060-5ad4-9937-c26c97f1ca68'
        32. id: 'file--9a1f834d-2506-5367-baec-7aa63996ac43'
        33. id: 'email-message--cf9b4b7f-14c8-5955-8065-020e0316b559'
        34. id: 'email-addr--d1b3bf0c-f02a-51a1-8102-11aba7959868'
        35. id: 'email-addr--9b7e29b3-fd8d-562e-b3f0-8fc8134f5dda'
        36. id: 'email-message--0c57a381-2a17-5e61-8754-5ef96efb286c'
        37. id: 'email-addr--2d77a846-6264-5d51-b586-e43822ea1ea3'
        38. id: 'domain-name--3c10e93f-798e-5a26-a0c1-08156efab7f5'
        39. id: 'directory--93c0a9b0-520d-545d-9094-1a08ddf46b05'
        40. id: 'autonomous-system--f720c34b-98ae-597f-ade5-27dc241e8c74'
        41. id: 'artifact--6f437177-6e48-5cf8-9d9e-872a2bddd641'
        42. id: 'vulnerability--0c7b5b88-8ff7-4a4d-aa9d-feb398cd0061'
        43. id: 'tool--a80c07ac-45f7-4d16-ab17-406d3a50b726'
        44. id: 'threat-actor--5b1fdb52-bdab-4a05-8daa-08d6a68f10b1'
        45. id: 'report--84e4d88f-44ea-4bcd-bbf3-b2c1c320bcb3'
        46. id: 'opinion--b01efc25-77b4-4003-b18b-f6e24b5cd9f7'
        47. id: 'observed-data--b67d30ff-02ac-498a-92f9-32f845f448cf'
        48. id: 'note--eebf5811-e34a-48ef-8831-b8d6dbe4cf3f'
        49. id: 'malware-analysis--d25167b7-fed0-4068-9ccd-a73dd2c5b07c'
        50. id: 'malware--8bcf14e9-2ba2-44ef-9e32-fbbc9d2608b2'
        51. id: 'malware--e82e93f6-7911-40d9-8b4a-5abc9dfc1efa'
        52. id: 'location--9dc0370c-c6ec-4a9f-b1dd-44ddbd9bd5e9'
        53. id: 'location--a6e9345f-5a15-4c29-8bb3-7dcc5d234d64'
        54. id: 'location--a6e9345f-5a15-4c29-8bb3-8dac5d168d64'
        55. id: 'intrusion-set--4e78f46f-a023-4e5f-bc24-71b3ca22ec29'
        56. id: 'infrastructure--38c47d93-d984-4fd9-b87b-d69d0841628d'
        57. id: 'indicator--d38a99ae-c5ee-4542-bc12-dfe68b48cc08'
        58. id: 'identity--023d105b-752e-4e3c-941c-7d3f3cb15e9e'
        59. id: 'grouping--911f0f23-2e7c-4f07-9436-d6e7bdd3b236'
        60. id: 'course-of-action--fdcb81ce-b5d3-4ed2-b962-c8287fba2d6a'
        61. id: 'campaign--ce88a5a8-69ff-4349-86da-ac59b35c5672'
        62. id: 'intrusion-set--0c7e22ad-b099-4dc3-b0df-2ea3f49ae2e6'
        63. id: 'sighting--ee20065d-2555-424f-ad9e-0f8428623c75'
        64. id: 'relationship--57b56a43-b8b0-4cba-9deb-34e3e1faed9e'
        65. id: 'attack-pattern--7e33a43e-e34b-40ec-89da-36c9bb2cacd5'
*/

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

//ComparasionListTypeSTIXObject содержит два списка STIX объектов, предназначенных для сравнения
type ComparasionListTypeSTIXObject struct {
	OldList, NewList []datamodels.ElementSTIXObject
}

//ComparasionListSTIXObject выполняет сравнение двух списков STIX объектов, cписка STIX объектов, полученных из БД и принятых от клиента API
func ComparasionListSTIXObject(clt ComparasionListTypeSTIXObject) []datamodels.DifferentObjectType {
	var (
		listDifferentResult []datamodels.DifferentObjectType
		dot                 datamodels.DifferentObjectType
		err                 error
		isEqual             = true
	)

	for _, vo := range clt.OldList {
		for _, vn := range clt.NewList {
			if vo.DataType != vn.DataType {
				continue
			}

			switch vn.DataType {
			//только для Domain Objects STIX
			case "attack-pattern":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetAttackPatternDomainObjectsSTIX(), "test source")
			case "campaign":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetCampaignDomainObjectsSTIX(), "test source")
			case "course-of-action":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetCourseOfActionDomainObjectsSTIX(), "test source")
			case "grouping":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetGroupingDomainObjectsSTIX(), "test source")
			case "identity":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetIdentityDomainObjectsSTIX(), "test source")
			case "indicator":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetIndicatorDomainObjectsSTIX(), "test source")
			case "infrastructure":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetInfrastructureDomainObjectsSTIX(), "test source")
			case "intrusion-set":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetIntrusionSetDomainObjectsSTIX(), "test source")
			case "location":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetLocationDomainObjectsSTIX(), "test source")
			case "malware":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetMalwareDomainObjectsSTIX(), "test source")
			case "malware-analysis":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetMalwareAnalysisDomainObjectsSTIX(), "test source")
			case "note":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetNoteDomainObjectsSTIX(), "test source")
			case "observed-data":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetObservedDataDomainObjectsSTIX(), "test source")
			case "opinion":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetOpinionDomainObjectsSTIX(), "test source")
			case "report":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetReportDomainObjectsSTIX(), "test source")
			case "threat-actor":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetThreatActorDomainObjectsSTIX(), "test source")
			case "tool":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetToolDomainObjectsSTIX(), "test source")
			case "vulnerability":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetVulnerabilityDomainObjectsSTIX(), "test source")
			}

			if err != nil {
				continue
			}

			if isEqual {
				continue
			}

			listDifferentResult = append(listDifferentResult, dot)
		}
	}

	return listDifferentResult
}
