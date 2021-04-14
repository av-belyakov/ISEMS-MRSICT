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
			var bomr []*datamodels.IdentifierTypeSTIX
			bonrTmp := []datamodels.IdentifierTypeSTIX{
				datamodels.IdentifierTypeSTIX("obj_mark_ref_odcfc22y2vr34u41udv"),
				datamodels.IdentifierTypeSTIX("obj_mark_ref_ahvas223h23h2dh3h1h"),
				datamodels.IdentifierTypeSTIX("obj_mark_ref_nivnbufuf74gf74gf7s"),
			}

			for _, v := range bonrTmp {
				bomr = append(bomr, &v)
			}

			before := datamodels.AttackPatternDomainObjectsSTIX{
				CommonPropertiesObjectSTIX: datamodels.CommonPropertiesObjectSTIX{
					Type: "attack pattern",
					ID:   "attack_pattern_vyw7d27dffd3ffd6f6fd6fw",
				},
				CommonPropertiesDomainObjectSTIX: datamodels.CommonPropertiesDomainObjectSTIX{
					SpecVersion:  "",
					Created:      time.Now(),
					Modified:     time.Now(),
					CreatedByRef: datamodels.IdentifierTypeSTIX("ref_cubu8c3gf8g8g3f83"),
					Revoked:      false,
					Labels:       []string{"lable_1", "lable_2", "lable_3"},
					Lang:         "RU",
					Сonfidence:   12,
					ExternalReferences: datamodels.ExternalReferencesTypeSTIX{
						&datamodels.ExternalReferenceTypeElementSTIX{
							SourceName:  "source_name_1",
							Description: "just any descripton",
							URL:         "http://any-site.org/example_one",
							Hashes: datamodels.HashesTypeSTIX{
								"SHA-1":   "dcubdub883g3838fgc83f",
								"SHA-128": "cb8b38b8c38f83f888f844",
							},
							ExternalID: "12444_gddgdg",
						},
						&datamodels.ExternalReferenceTypeElementSTIX{
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
					SpecVersion:  "",
					Created:      time.Now(),
					Modified:     time.Now(),
					CreatedByRef: datamodels.IdentifierTypeSTIX("ref_cubu8c3gf8g8g3f83"),
					Revoked:      false,
					Labels:       []string{"lable_1", "lable_2", "lable_3"},
					Lang:         "RU",
					Сonfidence:   12,
					ExternalReferences: datamodels.ExternalReferencesTypeSTIX{
						&datamodels.ExternalReferenceTypeElementSTIX{
							SourceName:  "source_name_1",
							Description: "just any descripton",
							URL:         "http://any-site.org/example_one",
							Hashes: datamodels.HashesTypeSTIX{
								"SHA-1":   "dcubdub883g3838fgc83f",
								"SHA-128": "cb8b38b8c38f83f888f844",
							},
							ExternalID: "12444_gddgdg",
						},
						&datamodels.ExternalReferenceTypeElementSTIX{
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

			isEqual, contrastResult, err := ComparisonTypeCommonFields(before.CommonPropertiesDomainObjectSTIX, after.CommonPropertiesDomainObjectSTIX)

			fmt.Println("_______ func 'ComparisonTypeCommonFields', contrastResult ________")
			fmt.Println(contrastResult)
			fmt.Println("------------------------------------------------------------------")

			/*
			   Контроль за изменением STIX объектов будет выполнятся только для STIX объектов типа Domain Object (STIX DO)
			   остальные типы STIX объектов (по спецификации вроде) не подлежат модификации, соответственно их мы только
			   добавляем или удаляем (как например объект Relationship)
			   Значит метод сравнения нужен только для CommonPropertiesDomainObjectSTIX и всех типов STIX DO
			*/

			Expect(isEqual).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	/*
			Context("", func() {
			It("", func(){

			})
		})
	*/
})

func ComparisonTypeCommonFields(before, after datamodels.CommonPropertiesDomainObjectSTIX) (bool, datamodels.ContrastObjectType, error) {
	var isEqual bool = true
	tmpOne := reflect.ValueOf(before)
	typeOfSOne := tmpOne.Type()

	tmpTwo := reflect.ValueOf(after)
	typeOfSTwo := tmpTwo.Type()

	fmt.Println("---=== func 'ComparisonTypeCommonFields' ===---")

	contrast := datamodels.ContrastObjectType{
		ModifiedTime: time.Now(),
		FieldList:    []datamodels.OldValuesObjectType{},
	}

	for i := 0; i < tmpOne.NumField(); i++ {
		for j := 0; j < tmpTwo.NumField(); j++ {
			resultDeepEqual := reflect.DeepEqual(tmpOne.Field(i).Interface(), tmpTwo.Field(j).Interface())

			if typeOfSOne.Field(i).Name == typeOfSTwo.Field(j).Name {

				fmt.Printf("Field: %s\tValue BEFORE: %v, AFTER: %v, Equal: %v\n", typeOfSOne.Field(i).Name, tmpOne.Field(i).Interface(), tmpTwo.Field(j).Interface(), resultDeepEqual)

				if !resultDeepEqual {
					contrast.FieldList = append(contrast.FieldList, datamodels.OldValuesObjectType{
						Path:  typeOfSOne.Field(i).Name,
						Value: tmpOne.Field(i).Interface(),
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

		fmt.Printf("func 'GetListElementSTIXObject', type STIX object: '%s', ID: '%s'\n", modelType.Type, modelType.ID)

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
