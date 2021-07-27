package moduleapirequestprocessing

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gorilla/websocket"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/moduleapirequestprocessing/temporarystorage"
	"ISEMS-MRSICT/modulelogginginformationerrors"
	//"ISEMS-MRSICT/moduletemporarymemorycommon"
)

//ChannelsModuleAPIRequestProcessing описание каналов передачи данных между ядром приложения и модулем обрабатывающем запросы с внешних источников
type ChannelsModuleAPIRequestProcessing struct {
	InputModule, OutputModule chan datamodels.ModuleReguestProcessingChannel
}

type settingsServerAPI struct {
	host, port  string
	users       []datamodels.ModuleAPIRequestProcessingUsersSetting
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType
}

var cmapirp ChannelsModuleAPIRequestProcessing
var repositoryStorageUserParameters *temporarystorage.RepositoryStorageUserParametersType

func init() {
	cmapirp = ChannelsModuleAPIRequestProcessing{
		InputModule:  make(chan datamodels.ModuleReguestProcessingChannel),
		OutputModule: make(chan datamodels.ModuleReguestProcessingChannel),
	}

	repositoryStorageUserParameters = temporarystorage.NewRepositoryStorageUserParameters()
}

//MainHandlerAPIReguestProcessing модуль инициализации обработчика запросов с внешних источников
func MainHandlerAPIReguestProcessing(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	mcs *datamodels.ModuleAPIRequestProcessingSetting,
	criptoSet *datamodels.CryptographySettings) ChannelsModuleAPIRequestProcessing {

	funcName := "MainHandlerAPIReguestProcessing"

	ssapi := settingsServerAPI{
		host:        mcs.Host,
		port:        strconv.Itoa(mcs.Port),
		users:       mcs.Users,
		chanSaveLog: chanSaveLog,
	}

	pathCertFile := criptoSet.PathCertFileModuleAPIRequestProcessingSettings
	pathPrivateKeyFile := criptoSet.PathPrivateKeyFileModuleAPIRequestProcessingSettings

	//сервер WSS для подключения клиентов
	go func() {
		http.HandleFunc("/api", ssapi.HandlerRequest)
		http.HandleFunc("/api_wss", ssapi.serverWss)

		err := http.ListenAndServeTLS(ssapi.host+":"+ssapi.port, pathCertFile, pathPrivateKeyFile, nil)
		if err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    funcName,
			}

			log.Println(err)
			os.Exit(1)
		}
	}()

	//маршрутизатор ответов поступающих от Ядра приложения
	go func() {
		for msg := range cmapirp.InputModule {
			if msg.ModuleReceiverMessage != "module api request processing" {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					Description: "the name of the source module or the destination module does not correspond to certain values",
					FuncName:    funcName,
				}

				continue
			}

			if msg.DataType > 2 {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					Description: "the type of data transmitted over the websocket connection is not defined",
					FuncName:    funcName,
				}

				continue
			}

			userParameters, err := repositoryStorageUserParameters.GetClientParametersForID(msg.ClientID)
			if err != nil {
				//если клиент с заданным ID не найден, отправляем широковещательное сообщение
				listUserParameters := repositoryStorageUserParameters.GetClientList()
				for _, up := range listUserParameters {
					if up.Connection == nil {
						continue
					}

					if err := up.SendWsMessage(msg.DataType, *msg.Data); err != nil {
						chanSaveLog <- modulelogginginformationerrors.LogMessageType{
							Description: fmt.Sprint(err),
							FuncName:    funcName,
						}
					}
				}

				continue
			}

			if err := userParameters.SendWsMessage(msg.DataType, *msg.Data); err != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					Description: fmt.Sprint(err),
					FuncName:    funcName,
				}
			}
		}
	}()

	log.Printf("\tThe module's API for processing requests from external sources has been successfully activated. The following API parameters are set ip address:%v, port:%v\n", ssapi.host, ssapi.port)

	return cmapirp
}

//HandlerRequest обработчик HTTPS запроса к "/"
func (settingsServerAPI *settingsServerAPI) HandlerRequest(w http.ResponseWriter, req *http.Request) {
	funcName := "HandlerRequest"

	bodyHTTPResponseError := []byte(`<!DOCTYPE html>
		<html lang="en"
		<head><meta charset="utf-8"><title>Server Nginx</title></head>
		<body><h1>Access denied! For additional information, please contact the webmaster.</h1></body>
		</html>`)

	stringToken := ""
	for headerName := range req.Header {
		if headerName == "Token" {
			stringToken = req.Header[headerName][0]
			continue
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Language", "en")

	if req.Method != "GET" {
		http.Error(w, "Method not allowed", 405)

		return
	}

	if len(stringToken) == 0 {
		w.Header().Set("Content-Length", strconv.Itoa(utf8.RuneCount(bodyHTTPResponseError)))

		w.WriteHeader(400)
		w.Write(bodyHTTPResponseError)

		settingsServerAPI.chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprintf("missing or incorrect identification token (сlient ipaddress %v)", req.RemoteAddr),
			FuncName:    funcName,
		}
	}

	for _, user := range settingsServerAPI.users {
		if stringToken == user.Token {
			remoteIPAndPort := strings.Split(req.RemoteAddr, ":")
			remoteAddr := remoteIPAndPort[0]
			remotePort := remoteIPAndPort[1]

			//добавляем нового пользователя которому разрешен доступ
			_ = repositoryStorageUserParameters.AddNewClient(remoteAddr, remotePort, user.ClientName, stringToken)

			http.Redirect(w, req, "https://"+settingsServerAPI.host+":"+settingsServerAPI.port+"/api_wss", 301)

			return
		}
	}

	w.Header().Set("Content-Length", strconv.Itoa(utf8.RuneCount(bodyHTTPResponseError)))

	w.WriteHeader(400)
	w.Write(bodyHTTPResponseError)

	settingsServerAPI.chanSaveLog <- modulelogginginformationerrors.LogMessageType{
		TypeMessage: "error",
		Description: fmt.Sprintf("missing or incorrect identification token (сlient ipaddress %v) bodyHTTPResponseError", req.RemoteAddr),
		FuncName:    funcName,
	}
}

func (settingsServerAPI *settingsServerAPI) serverWss(w http.ResponseWriter, req *http.Request) {
	funcName := "serverWss"

	remoteIPAndPort := strings.Split(req.RemoteAddr, ":")
	remoteIP := remoteIPAndPort[0]
	remotePort := remoteIPAndPort[1]

	//проверяем прошел ли клиент аутентификацию
	clientID, _, ok := repositoryStorageUserParameters.SearchClientForIP(remoteIP, req.Header["Token"][0])
	if !ok {
		w.WriteHeader(401)
		settingsServerAPI.chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			Description: fmt.Sprintf("access for the user with ipaddress %v is prohibited", req.RemoteAddr),
			FuncName:    funcName,
		}

		return
	}

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: false,
		//ReadBufferSize:    1024,
		//WriteBufferSize:   100000000,
		HandshakeTimeout: (time.Duration(1) * time.Second),
	}

	c, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		c.Close()

		//удаляем информацию о клиенте
		repositoryStorageUserParameters.DeleteClient(clientID)

		settingsServerAPI.chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			Description: fmt.Sprint(err),
			FuncName:    funcName,
		}

		log.Printf("Client API module 'moduleapirequestprocessing' (ID %v) whis IP %v is disconnect!\n", clientID, remoteIP)
	}

	//получаем настройки клиента
	cp, err := repositoryStorageUserParameters.GetClientParametersForID(clientID)
	if err != nil {
		settingsServerAPI.chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			Description: fmt.Sprintf("client setup with ID %v not found", clientID),
			FuncName:    funcName,
		}

		return
	}

	repositoryStorageUserParameters.SaveWssClientConnection(clientID, c)

	log.Printf("Client API (ID %v) whis IP %v:%v is connect", clientID, remoteIP, remotePort)

	//маршрутизация сообщений приходящих от клиентов API
	go func() {
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				c.Close()

				//удаляем информацию о клиенте
				repositoryStorageUserParameters.DeleteClient(clientID)

				settingsServerAPI.chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					Description: fmt.Sprintf("client setup with ID %v not found", clientID),
					FuncName:    funcName,
				}

				log.Printf("Client API (ID %v) whis IP %v is disconnect", clientID, remoteIP)

				break
			}

			cmapirp.OutputModule <- datamodels.ModuleReguestProcessingChannel{
				CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
					ModuleGeneratorMessage: "module api request processing",
					ModuleReceiverMessage:  "module core application",
				},
				ClientID:   clientID,
				ClientName: cp.ClientName,
				Data:       &msg,
			}
		}
	}()
}
