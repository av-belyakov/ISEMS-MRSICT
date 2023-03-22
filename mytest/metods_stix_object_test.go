package mytest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

var (
	connectError error
	qp           interactionmongodb.QueryParameters
	dm           datamodels.MongoDBSettings
	cdmd         interactionmongodb.ConnectionDescriptorMongoDB
)

var _ = BeforeSuite(func() {
	dm = datamodels.MongoDBSettings{
		Host:     "192.168.9.208",
		Port:     37017,
		User:     "module-isems-mrsict",
		Password: "vkL6Znj$Pmt1e1",
		NameDB:   "isems-mrsict",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	cdmd = interactionmongodb.ConnectionDescriptorMongoDB{
		Ctx:       ctx,
		CtxCancel: cancel,
	}

	connectError = cdmd.CreateConnection(&dm)

	qp = interactionmongodb.QueryParameters{
		NameDB:         "isems-mrsict",
		CollectionName: "stix_object_collection",
		ConnectDB:      cdmd.Connection,
	}
})

var _ = AfterSuite(func() {
	//ctxCancel()
	cdmd.CtxCancel()
})

var _ = Describe("MetodsStixObject", func() {
	Context("Тест 1. Проверка наличия установленного соединения с БД", func() {
		It("При установления соединения с БД ошибки быть не должно", func() {
			Expect(connectError).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Получаем информацию о STIX объект с ID 'network-traffic--68ead79b-28db-4811-85f6-971f46548343'", func() {
		It("должен быть получен список из 1 STIX объекта, ошибок быть не должно", func() {
			cur, err := qp.Find(bson.D{bson.E{Key: "commonpropertiesobjectstix.id", Value: "network-traffic--68ead79b-28db-4811-85f6-971f46548343"}})

			//fmt.Printf("___ ERROR find ID:'network-traffic--68ead79b-28db-4811-85f6-971f46548343' - '%v'\n", err)
			/*l := []*datamodels.NetworkTrafficCyberObservableObjectSTIX{}
			for cur.Next(context.Background()) {
				var model datamodels.NetworkTrafficCyberObservableObjectSTIX
				_ = cur.Decode(&model)

				l = append(l, &model)
			}*/

			listelm := interactionmongodb.GetListElementSTIXObject(cur)

			//for _, v := range listelm {
			//fmt.Printf("Found STIX object with ID:'network-traffic--68ead79b-28db-4811-85f6-971f46548343' - '%v'\n\tIPFix = %v\n", *v, &v.IPFix)
			//fmt.Printf("Found STIX object with ID:'network-traffic--68ead79b-28db-4811-85f6-971f46548343' - '%v'\n", *v)
			//}

			Expect(err).ShouldNot(HaveOccurred())
			// Expect(len(l)).Should(Equal(1))
			Expect(len(listelm)).Should(Equal(1))
		})
	})

	Context("Тест 3. Выполняем проверку типа 'network-traffic' методом ValidateStruct", func() {
		It("На валидное содержимое типа NetworkTrafficCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
	   				"type": "network-traffic",
	   				"spec_version": "2.1",
	   				"id": "network-traffic--630d7bb1-0bbc-53a6-a6d4-f3c2d35c2734",
	   				"src_ref": "ipv4-addr--e42c19c8-f9fe-5ae9-9fc8-22c398f78fb",
	   				"dst_ref": "ipv4-addr--03b708d9-7761-5523-ab75-5ea096294a68",
	   				"src_port": 2487,
	   				"dst_port": 1723,
	   				"protocols": ["ipv4", "tcp"],
	   				"src_byte_count": 147600,
	   				"src_packets": 100,
	   				"dst_byte_count": 935750,
	   				"encapsulates_refs": ["network-traffic--53e0bf48-2eee-5c03-8bde-ed7049d2c0a3"],
	   				"ipfix": {
	   					"minimumIpTotalLength": 32,
	   					"maximumIpTotalLength": 2556
	   				},
	   				"extensions": {
	   					"http-request-ext": {
	   						"request_method": "get",
	   						"request_value": "/download.html",
	   						"request_version": "http/1.1",
	   						"request_header": {
	   							"Accept-Encoding": "gzip,deflate",
	   							"User-Agent": "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.6)Gecko/20040113",
	   							"Host": "www.example.com"
	   						}
	   					},
	   					"socket-ext": {
	   						"is_listening": true,
	   						"address_family": "AF_INET",
	   						"socket_type": "SOCK_STREAM"
	   					}
	   				}
	   }`))
			var md datamodels.NetworkTrafficCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.NetworkTrafficCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.ValidateStruct()).Should(BeTrue())
		})
	})

	Context("Тест 4.1. Выполняем проверку типа 'file' методом ValidateStruct с расширением 'Extensions' типа 'raster-image-ext'", func() {
		It("На валидное содержимое типа FileCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
	   				"type": "file",
	   				"spec_version": "2.1",
	   				"id": "file--c7d1e135-8b34-549a-bb47-302f5cf998ed",
					"spec_version": "2.1",
					"granular_markings": [],
					"defanged": false,
					"size": 324495943,
					"name": "picture.jpg",
					"name_enc": "",
					"magic_number_hex": "nh8939r39rgh939rt4899239rg83gr99ed2e",
					"mime_type": "",
					"ctime": "2021-05-11T20:14:45.000+00:00",
					"mtime": "0001-01-01T00:00:00.000+00:00",
					"atime": "0001-01-01T00:00:00.000+00:00",
					"parent_directory_ref": "",
					"content_ref": "artifact--4cce66f8-6eaa-53cb-85d5-3a85fca3a6c5",
					"extensions": {
						"raster-image-ext": {
							"image_height": 150,
							"image_width": 340,
							"bits_per_pixel": 90,
							"exif_tags": {
								"Make": "Nikon",
								"Model": "D7000",
								"XResolution": 4928,
								"YResolution": 3264
							}
						}
					}
	   }`))
			var md datamodels.FileCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.FileCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.ValidateStruct()).Should(BeTrue())

			for _, v := range newmd.Extensions {
				rie, ok := v.(datamodels.RasterImageFileExtensionSTIX)

				Expect(ok).Should(BeTrue())
				Expect(rie.BitsPerPixel).Should(Equal(90))
				Expect(rie.ExifTags.Make).Should(Equal("Nikon"))
			}
		})
	})

	Context("Тест 4.2. Выполняем проверку типа 'file' методом ValidateStruct с расширением 'Extensions' типа 'pdf-ext'", func() {
		It("На валидное содержимое типа FileCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
			   				"type": "file",
			   				"spec_version": "2.1",
			   				"id": "file--68fa1c84-1217-48e7-bc62-02d53e51b270",
							"spec_version": "2.1",
							"granular_markings": [],
							"defanged": false,
							"size": 783894,
							"name": "my_movies.torrent",
							"name_enc": "",
							"magic_number_hex": "8ef88388yg8y83yh38d8fy84yty48ytr84",
							"mime_type": "",
							"ctime": "2022-03-21T12:25:45.000+00:00",
							"mtime": "2022-03-21T13:01:03.000+00:00",
							"atime": "0001-01-01T00:00:00.000+00:00",
							"parent_directory_ref": "",
							"content_ref": "artifact--4cce66f8-6eaa-53cb-85d5-3a85fca3a6c5",
							"extensions": {
								"pdf-ext": {
									"version": "1.7",
									"document_info_dict": {
										"Title": "Sample document",
										"Author": "Adobe Systems Incorporated",
										"Creator": "Adobe FrameMaker 5.5.3 for Power Macintosh",
										"Producer": "Acrobat Distiller 3.01 for Power Macintosh",
										"CreationDate": "20070412090123-02"
										},
									"pdfid0": "DFCE52BD827ECF765649852119D",
									"pdfid1": "57A1E0F9ED2AE523E313C"
								}
							}
			   }`))

			var md datamodels.FileCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.FileCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.ValidateStruct()).Should(BeTrue())

			for _, v := range newmd.Extensions {
				pdffe, ok := v.(datamodels.PDFFileExtensionSTIX)

				Expect(ok).Should(BeTrue())
				Expect(len(pdffe.DocumentInfoDict)).Should(Equal(5))
			}
		})
	})

	Context("Тест 4.3. Выполняем проверку типа 'network-traffic' методом ValidateStruct с расширением 'Extensions' типа 'socket-ext'", func() {
		It("На валидное содержимое типа FileCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "network-traffic",
				"spec_version": "2.1",
				"id": "network-traffic--630d7bb1-0bbc-53a6-a6d4-f3c2d35c2734",
				"src_ref": "ipv4-addr--e42c19c8-f9fe-5ae9-9fc8-22c398f78fb",
				"dst_ref": "ipv4-addr--03b708d9-7761-5523-ab75-5ea096294a68",
				"protocols": [
					"ipv4",
					"tcp"
				],
				"src_byte_count": 147600,
				"src_packets": 100,
				"ipfix": {
					"minimumIpTotalLength": 32,
					"maximumIpTotalLength": 2556
				},
				"extensions": {
					"socket-ext": {
						"address_family": "AF_INET",
						"is_blocking": false,
						"is_listening": true,
						"socket_type": "SOCK_STREAM",
						"socket_descriptor": 12,
						"socket_handle": 10,
						"options": {
							"SO_RCVTIMEO": 14325,
							"SO_LINGER": 7533
						}
					}
				}
			}`))

			var md datamodels.NetworkTrafficCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.NetworkTrafficCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.ValidateStruct()).Should(BeTrue())

			for _, v := range newmd.Extensions {
				pdnse, ok := v.(datamodels.NetworkSocketExtensionSTIX)

				Expect(ok).Should(BeTrue())
				Expect(len(pdnse.Options)).Should(Equal(2))
			}
		})
	})

	Context("Тест 4.4. Выполняем проверку типа 'process' методом ValidateStruct с расширением 'Extensions' типа 'windows-process-ext'", func() {
		It("На валидное содержимое типа ProcessCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "process",
				"spec_version": "2.1",
				"id": "process--f52a906a-0dfc-40bd-92f1-e7778ead38a9",
				"pid": 1221,
				"created": "2016-01-20T14:11:25.55Z",
				"command_line": "./gedit-bin --new-window",
				"image_ref": "file--e04f22d1-be2c-59de-add8-10f61d15fe20",
				"extensions": {
					"windows-process-ext": {
						"aslr_enabled": true,
						"dep_enabled": true,
						"priority": "HIGH_PRIORITY_CLASS",
						"owner_sid": "S-1-5-21-186985262-1144665072-74031268-1309",
						"window_title": "any window process",
						"startup_info": {
							"keyIp": "valueIp"
						},
						"integrity_level": "levelOne"
					}
				}
			}`))

			var md datamodels.ProcessCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.ProcessCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.ValidateStruct()).Should(BeTrue())

			for _, v := range newmd.Extensions {
				pdnse, ok := v.(datamodels.WindowsProcessExtensionSTIX)

				Expect(ok).Should(BeTrue())
				Expect(pdnse.Priority).Should(Equal("HIGH_PRIORITY_CLASS"))
				Expect(pdnse.WindowTitle).Should(Equal("any window process"))
			}
		})
	})

	Context("Тест 4.5. Выполняем проверку типа 'process' методом SanitizeStruct с расширением 'EnvironmentVariables'", func() {
		It("На валидное содержимое типа ProcessCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "process",
				"spec_version": "2.1",
				"id": "process--f52a906a-0dfc-40bd-92f1-e7778ead38a9",
				"pid": 1221,
				"created": "2016-01-20T14:11:25.55Z",
				"command_line": "./gedit-bin --new-window",
				"image_ref": "file--e04f22d1-be2c-59de-add8-10f61d15fe20",
				"extensions": {
					"windows-process-ext": {
						"aslr_enabled": true,
						"dep_enabled": true,
						"priority": "HIGH_PRIORITY_CLASS",
						"owner_sid": "S-1-5-21-186985262-1144665072-74031268-1309",
						"window_title": "any window process",
						"startup_info": {
							"keyIp": "valueIp"
						},
						"integrity_level": "levelOne"
					}
				},
				"environment_variables": {
					"env_key_1": "variable_1",
					"env_key_2": "variable_2",
					"env_key_3": "variable_3"
				}
			}`))

			var md datamodels.ProcessCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.ProcessCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.ValidateStruct()).Should(BeTrue())
			Expect(len(newmd.EnvironmentVariables)).Should(Equal(3))
		})
	})
})
