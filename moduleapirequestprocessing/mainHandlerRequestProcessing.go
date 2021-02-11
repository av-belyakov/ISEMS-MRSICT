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
	"ISEMS-MRSICT/modulelogginginformationerrors"
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

func init() {
	cmapirp = ChannelsModuleAPIRequestProcessing{
		InputModule:  make(chan datamodels.ModuleReguestProcessingChannel),
		OutputModule: make(chan datamodels.ModuleReguestProcessingChannel),
	}
}

//MainHandlerAPIReguestProcessing модуль инициализации обработчика запросов с внешних источников
func MainHandlerAPIReguestProcessing(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	mcs *datamodels.ModuleAPIRequestProcessingSetting,
	criptoSet *datamodels.CryptographySettings) ChannelsModuleAPIRequestProcessing {

	funcName := "MainHandlerAPIReguestProcessing"

	fmt.Println("func 'MainHandlerReguestProcessing', START...")
	fmt.Printf("func 'MainHandlerReguestProcessing', module connection settings: '%v'\n", mcs)

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

			//			fmt.Println("An error when initializing the api module of the request handler from external sources.")

			log.Println(err)
			os.Exit(1)
		}
	}()

	return cmapirp
}

//HandlerRequest обработчик HTTPS запроса к "/"
func (settingsServerAPI *settingsServerAPI) HandlerRequest(w http.ResponseWriter, req *http.Request) {
	funcName := "HandlerRequest"

	/*defer func() {
		if err := recover(); err != nil {
			settingsServerAPI.SaveMessageApp.LogMessage(savemessageapp.TypeLogMessage{
				Description: fmt.Sprint(err),
				FuncName:    funcName,
			})
		}
	}()*/

	bodyHTTPResponseError := []byte(`<!DOCTYPE html>
		<html lang="en"
		<head><meta charset="utf-8"><title>Server Nginx</title></head>
		<body><h1>Access denied. For additional information, please contact the webmaster.</h1></body>
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
			Description: fmt.Sprintf("missing or incorrect identification token (сlient ipaddress %v), module 'mainAPIApp'", req.RemoteAddr),
			FuncName:    funcName,
		}
	}

	for _, user := range settingsServerAPI.users {
		if stringToken == user.Token {
			remoteIPAndPort := strings.Split(req.RemoteAddr, ":")
			remoteAddr := remoteIPAndPort[0]
			remotePort := remoteIPAndPort[1]

			/*
			   Вот здесь нужно ВРЕМЕННОЕ хранилище
			   Кроме того я не сделал сертификат и приватный ключ
			*/

			//добавляем нового пользователя которому разрешен доступ
			_ = storingMemoryAPI.AddNewClient(remoteAddr, remotePort, user.ClientName, stringToken)

			http.Redirect(w, req, "https://"+settingsServerAPI.host+":"+settingsServerAPI.port+"/api_wss", 301)

			return
		}
	}

	w.Header().Set("Content-Length", strconv.Itoa(utf8.RuneCount(bodyHTTPResponseError)))

	w.WriteHeader(400)
	w.Write(bodyHTTPResponseError)

	settingsServerAPI.chanSaveLog <- modulelogginginformationerrors.LogMessageType{
		TypeMessage: "error",
		Description: fmt.Sprintf("missing or incorrect identification token (сlient ipaddress %v), module 'mainAPIApp' bodyHTTPResponseError", req.RemoteAddr),
		FuncName:    funcName,
	}
}

func (settingsServerAPI *settingsServerAPI) serverWss(w http.ResponseWriter, req *http.Request) {
	funcName := "serverWss"

	/*defer func() {
		if err := recover(); err != nil {
			settingsServerAPI.SaveMessageApp.LogMessage(savemessageapp.TypeLogMessage{
				Description: fmt.Sprint(err),
				FuncName:    funcName,
			})
		}
	}()*/

	remoteIPAndPort := strings.Split(req.RemoteAddr, ":")
	remoteIP := remoteIPAndPort[0]
	remotePort := remoteIPAndPort[1]

	//проверяем прошел ли клиент аутентификацию
	clientID, _, ok := storingMemoryAPI.SearchClientForIP(remoteIP, req.Header["Token"][0])
	if !ok {

		fmt.Println("Client is Unauthorized")

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
		storingMemoryAPI.DelClientAPI(clientID)

		settingsServerAPI.chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			Description: fmt.Sprint(err),
			FuncName:    funcName,
		}

		log.Printf("Client API (ID %v) whis IP %v is disconnect!\n", clientID, remoteIP)
	}

	//получаем настройки клиента
	clientSettings, err := storingMemoryAPI.GetClientSettings(clientID)
	if err != nil {
		settingsServerAPI.chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			Description: fmt.Sprintf("client setup with ID %v not found", clientID),
			FuncName:    funcName,
		}

		return
	}

	storingMemoryAPI.SaveWssClientConnection(clientID, c)

	log.Printf("Client API (ID %v) whis IP %v:%v is connect", clientID, remoteIP, remotePort)

	//при подключении клиента отправляем запрос на получение списка источников
	//	sendMsgGetSourceList(clientID)

	//маршрутизация сообщений приходящих от клиентов API
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				c.Close()

				//удаляем информацию о клиенте
				storingMemoryAPI.DelClientAPI(clientID)

				settingsServerAPI.chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					Description: fmt.Sprintf("client setup with ID %v not found", clientID),
					FuncName:    funcName,
				}

				log.Printf("Client API (ID %v) whis IP %v is disconnect", clientID, remoteIP)

				break
			}

			/*			chn.ChanIn <- &configure.MsgBetweenCoreAndAPI{
						MsgGenerator: "API module",
						MsgRecipient: "Core module",
						IDClientAPI:  clientID,
						ClientName:   clientSettings.ClientName,
						ClientIP:     remoteIP,
						MsgJSON:      message,
					}*/
		}
	}()
}

/*
settingsServerAPI := settingsServerAPI{
		IP:             appConfig.ServerAPI.Host,
		Port:           strconv.Itoa(appConfig.ServerAPI.Port),
		Tokens:         appConfig.AuthenticationTokenClientsAPI,
		SaveMessageApp: saveMessageApp,
	}
	funcName := "MainAPIApp"

	//сервер WSS для подключения клиентов
	go func() {
		http.HandleFunc("/api", settingsServerAPI.HandlerRequest)
		http.HandleFunc("/api_wss", settingsServerAPI.serverWss)

		err := http.ListenAndServeTLS(settingsServerAPI.IP+":"+settingsServerAPI.Port, appConfig.ServerAPI.PathCertFile, appConfig.ServerAPI.PathPrivateKeyFile, nil)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}()

	//маршрутизатор ответов от Core module
	go func() {
		for msg := range chn.ChanOut {
			if msg.MsgGenerator == "Core module" && msg.MsgRecipient == "API module" {
				msgjson, ok := msg.MsgJSON.([]byte)
				if !ok {
					settingsServerAPI.SaveMessageApp.LogMessage(savemessageapp.TypeLogMessage{
						Description: "failed to send json message, error while casting type",
						FuncName:    funcName,
					})

					continue
				}

				clientSettings, err := storingMemoryAPI.GetClientSettings(msg.IDClientAPI)
				//если клиент с таким ID не найден, отправляем широковещательное сообщение
				if err != nil {
					cl := storingMemoryAPI.GetClientList()
					for _, cs := range cl {
						if cs.Connection == nil {
							continue
						}
						if err := cs.SendWsMessage(1, msgjson); err != nil {
							settingsServerAPI.SaveMessageApp.LogMessage(savemessageapp.TypeLogMessage{
								Description: fmt.Sprint(err),
								FuncName:    funcName,
							})
						}
					}

					continue
				}

				if clientSettings.Connection == nil {
					continue
				}

				if err := clientSettings.SendWsMessage(1, msgjson); err != nil {
					settingsServerAPI.SaveMessageApp.LogMessage(savemessageapp.TypeLogMessage{
						Description: fmt.Sprint(err),
						FuncName:    funcName,
					})
				}
			}
		}
	}()
*/
