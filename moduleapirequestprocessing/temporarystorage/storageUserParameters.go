package temporarystorage

import (
	"fmt"
	"time"

	"ISEMS-MRSICT/commonlibs"

	"github.com/gorilla/websocket"
)

//NewRepositoryStorageUserParameters создание нового репозитория для хранения параметров пользователя
func NewRepositoryStorageUserParameters() *RepositoryStorageUserParametersType {
	rsupt := RepositoryStorageUserParametersType{
		clientSettings:    map[string]*StorageUserParameters{},
		chanReqParameters: make(chan typeChanReqParameters),
	}

	go func() {
		for msg := range rsupt.chanReqParameters {
			switch msg.actionType {
			case "add new client":
				rsupt.clientSettings[msg.clientID] = &StorageUserParameters{
					IP:         msg.clientIP,
					Token:      msg.token,
					ClientName: msg.clientName,
					IsAllowed:  true,
				}

				msg.chanRes <- typeChanResParameters{}

			case "get client list":
				msg.chanRes <- typeChanResParameters{
					clientList: rsupt.clientSettings,
				}

			case "get client parameters":
				if err := rsupt.searchID(msg.clientID); err != nil {
					msg.chanRes <- typeChanResParameters{
						msgErr: err,
					}

					continue
				}

				res := typeChanResParameters{}
				res.clientSetting = rsupt.clientSettings[msg.clientID]

				msg.chanRes <- res

			case "search client for ip":
				if len(rsupt.clientSettings) == 0 {
					msg.chanRes <- typeChanResParameters{
						msgErr: fmt.Errorf("the client list is empty"),
					}

					continue
				}

				res := typeChanResParameters{}
				for clientID, setting := range rsupt.clientSettings {
					if (msg.clientIP == setting.IP) && (msg.token == setting.Token) {
						res.clientID = clientID
						res.clientSetting = setting
					}
				}

				msg.chanRes <- res

			case "save client connection":
				if err := rsupt.searchID(msg.clientID); err != nil {
					msg.chanRes <- typeChanResParameters{
						msgErr: err,
					}

					continue
				}

				rsupt.clientSettings[msg.clientID].Connection = msg.connect

				msg.chanRes <- typeChanResParameters{}

			case "get client connection":
				if err := rsupt.searchID(msg.clientID); err != nil {
					msg.chanRes <- typeChanResParameters{
						msgErr: err,
					}

					continue
				}
				res := typeChanResParameters{}
				res.connect = rsupt.clientSettings[msg.clientID].Connection

				msg.chanRes <- res

			case "delete client":
				delete(rsupt.clientSettings, msg.clientID)

				msg.chanRes <- typeChanResParameters{}

			}
		}
	}()

	return &rsupt
}

//AddNewClient добавляет нового клиента
func (rsupt *RepositoryStorageUserParametersType) AddNewClient(clientIP, clientPort, clientName, clientToken string) string {

	fmt.Println("func 'AddNewClient', START...")

	hsum := commonlibs.GetUniqIDFormatMD5(clientIP + "_" + clientPort + "_" + fmt.Sprint(time.Now().Unix()) + "_client API")

	cr := make(chan typeChanResParameters)
	defer close(cr)

	rsupt.chanReqParameters <- typeChanReqParameters{
		actionType: "add new client",
		clientIP:   clientIP,
		token:      clientToken,
		clientID:   hsum,
		clientName: clientName,
		chanRes:    cr,
	}

	<-cr

	return hsum
}

//SearchClientForIP поиск информации о клиенте по его ip адресу
func (rsupt *RepositoryStorageUserParametersType) SearchClientForIP(clientIP, clientToken string) (string, *StorageUserParameters, bool) {

	fmt.Println("func 'SearchClientForIP', START...")

	cr := make(chan typeChanResParameters)
	defer close(cr)

	rsupt.chanReqParameters <- typeChanReqParameters{
		actionType: "search client for ip",
		clientIP:   clientIP,
		token:      clientToken,
		chanRes:    cr,
	}

	res := <-cr

	if (res.clientID == "") || (res.msgErr != nil) {
		return "", nil, false
	}

	return res.clientID, res.clientSetting, true
}

//GetClientParametersForID получить все настройки клиента
func (rsupt *RepositoryStorageUserParametersType) GetClientParametersForID(clientID string) (*StorageUserParameters, error) {

	fmt.Println("func 'SearchClientForIP', START...")

	cr := make(chan typeChanResParameters)
	defer close(cr)

	rsupt.chanReqParameters <- typeChanReqParameters{
		actionType: "get client parameters",
		clientID:   clientID,
		chanRes:    cr,
	}

	res := <-cr

	return res.clientSetting, res.msgErr
}

//GetClientList получить весь список клиентов
func (rsupt *RepositoryStorageUserParametersType) GetClientList() map[string]*StorageUserParameters {

	fmt.Println("func 'GetClientList', START...")

	cr := make(chan typeChanResParameters)
	defer close(cr)

	rsupt.chanReqParameters <- typeChanReqParameters{
		actionType: "get client list",
		chanRes:    cr,
	}

	return (<-cr).clientList
}

//SaveWssClientConnection сохранить линк соединения с клиентом
func (rsupt *RepositoryStorageUserParametersType) SaveWssClientConnection(clientID string, conn *websocket.Conn) error {

	fmt.Println("func 'SaveWssClientConnection', START...")

	cr := make(chan typeChanResParameters)
	defer close(cr)

	req := typeChanReqParameters{
		actionType: "save client connection",
		clientID:   clientID,
		chanRes:    cr,
	}
	req.connect = conn

	rsupt.chanReqParameters <- req

	return (<-cr).msgErr
}

//GetWssClientConnection получить линк wss соединения
func (rsupt *RepositoryStorageUserParametersType) GetWssClientConnection(clientID string) (*websocket.Conn, error) {

	fmt.Println("func 'SaveWssClientConnection', START...")

	cr := make(chan typeChanResParameters)
	defer close(cr)

	rsupt.chanReqParameters <- typeChanReqParameters{
		actionType: "get client connection",
		clientID:   clientID,
		chanRes:    cr,
	}

	res := <-cr

	return res.connect, res.msgErr
}

//DeleteClient удалить всю информацию о клиенте
func (rsupt *RepositoryStorageUserParametersType) DeleteClient(clientID string) {

	fmt.Println("func 'DelClientAPI', START...")

	cr := make(chan typeChanResParameters)
	defer close(cr)

	rsupt.chanReqParameters <- typeChanReqParameters{
		actionType: "del client",
		clientID:   clientID,
		chanRes:    cr,
	}

	<-cr
}

func (rsupt *RepositoryStorageUserParametersType) searchID(clientID string) error {
	if _, ok := rsupt.clientSettings[clientID]; !ok {
		return fmt.Errorf("client with specified ID %v not found", clientID)
	}

	return nil
}
