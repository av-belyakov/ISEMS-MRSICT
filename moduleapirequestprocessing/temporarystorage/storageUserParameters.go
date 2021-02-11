package temporarystorage

import (
	"fmt"

	"github.com/gorilla/websocket"
)

//NewRepository создание нового репозитория
func NewRepository() *ModReqProcTemporaryStorage {
	mrpts := ModReqProcTemporaryStorage{
		clientSettings:    map[string]*ModReqProcTemporaryStorageClientSettings{},
		chanReqParameters: make(chan typeChanReqParam),
	}

	go func() {
		for msg := range mrpts.chanReqParameters {
			switch msg.actionType {
			case "add new client":
				/*				smapi.clientSettings[msg.clientID] = &ClientSettings{
									IP:         msg.clientIP,
									Token:      msg.token,
									ClientName: msg.clientName,
									IsAllowed:  true,
								}

								msg.chanRes <- typeChanResSetting{}*/

			case "get client list":
				/*				msg.chanRes <- typeChanResSetting{
								clientList: smapi.clientSettings,
							}*/

			case "get client setting":
				/*				if err := smapi.searchID(msg.clientID); err != nil {
									msg.chanRes <- typeChanResSetting{
										msgErr: err,
									}

									continue
								}

								res := typeChanResSetting{}
								res.clientSetting = smapi.clientSettings[msg.clientID]

								msg.chanRes <- res*/

			case "search client for ip":
				/*				if len(smapi.clientSettings) == 0 {
									msg.chanRes <- typeChanResSetting{
										msgErr: fmt.Errorf("the client list is empty"),
									}

									continue
								}

								res := typeChanResSetting{}
								for clientID, setting := range smapi.clientSettings {
									if (msg.clientIP == setting.IP) && (msg.token == setting.Token) {
										res.clientID = clientID
										res.clientSetting = setting
									}
								}

								msg.chanRes <- res*/

			case "save client connection":
				/*				if err := smapi.searchID(msg.clientID); err != nil {
									msg.chanRes <- typeChanResSetting{
										msgErr: err,
									}

									continue
								}

								smapi.clientSettings[msg.clientID].Connection = msg.connect

								msg.chanRes <- typeChanResSetting{}*/

			case "get client connection":
				/*				if err := smapi.searchID(msg.clientID); err != nil {
									msg.chanRes <- typeChanResSetting{
										msgErr: err,
									}

									continue
								}
								res := typeChanResSetting{}
								res.connect = smapi.clientSettings[msg.clientID].Connection

								msg.chanRes <- res*/

			case "del client":
				/*				delete(smapi.clientSettings, msg.clientID)

								msg.chanRes <- typeChanResSetting{}*/

			}
		}
	}()

	return &mrpts
}

//AddNewClient добавляет нового клиента
func (mrpts *ModReqProcTemporaryStorage) AddNewClient(clientIP, port, clientName, token string) string {

	fmt.Println("func 'AddNewClient', START...")
	return ""

	/*	hsum := commonlibs.GetUniqIDFormatMD5(clientIP + "_" + port + "_client API")

		cr := make(chan typeChanResSetting)
		defer close(cr)

		mrpts.chanReqParameters <- typeChanReqParam{
			actionType: "add new client",
			clientIP:   clientIP,
			token:      token,
			clientID:   hsum,
			clientName: clientName,
			chanRes:    cr,
		}

		<-cr

		return hsum*/
}

//SearchClientForIP поиск информации о клиенте по его ip адресу
func (mrpts *ModReqProcTemporaryStorage) SearchClientForIP(clientIP, token string) /*(string, *ClientSettings, bool)*/ {

	fmt.Println("func 'SearchClientForIP', START...")

	/*cr := make(chan typeChanResSetting)
	defer close(cr)

	mrpts.chanReqParameters <- typeChanReqParam{
		actionType: "search client for ip",
		clientIP:   clientIP,
		token:      token,
		chanRes:    cr,
	}

	res := <-cr

	if (res.clientID == "") || (res.msgErr != nil) {
		return "", nil, false
	}

	return res.clientID, res.clientSetting, true*/
}

//GetClientSettings получить все настройки клиента
func (mrpts *ModReqProcTemporaryStorage) GetClientSettings(clientID string) /*(*ClientSettings, error)*/ {

	fmt.Println("func 'SearchClientForIP', START...")

	/*	cr := make(chan typeChanResSetting)
		defer close(cr)

		mrpts.chanReqParameters <- typeChanReqParam{
			actionType: "get client setting",
			clientID:   clientID,
			chanRes:    cr,
		}

		res := <-cr

		return res.clientSetting, res.msgErr*/
}

//GetClientList получить весь список клиентов
func (mrpts *ModReqProcTemporaryStorage) GetClientList() /*map[string]*ClientSettings*/ {

	fmt.Println("func 'GetClientList', START...")

	/*	cr := make(chan typeChanResSetting)
		defer close(cr)

		mrpts.chanReqParameters <- typeChanReqParam{
			actionType: "get client list",
			chanRes:    cr,
		}

		return (<-cr).clientList*/
}

//SaveWssClientConnection сохранить линк соединения с клиентом
func (mrpts *ModReqProcTemporaryStorage) SaveWssClientConnection(clientID string, conn *websocket.Conn) /*error*/ {

	fmt.Println("func 'SaveWssClientConnection', START...")

	/*cr := make(chan typeChanResSetting)
	defer close(cr)

	req := typeChanReqSetting{
		actionType: "save client connection",
		clientID:   clientID,
		chanRes:    cr,
	}
	req.connect = conn

	smapi.chanReqSetting <- req

	return (<-cr).msgErr*/
}

//GetWssClientConnection получить линк wss соединения
func (mrpts *ModReqProcTemporaryStorage) GetWssClientConnection(clientID string) /*(*websocket.Conn, error)*/ {

	fmt.Println("func 'SaveWssClientConnection', START...")

	/*cr := make(chan typeChanResSetting)
	defer close(cr)

	smapi.chanReqSetting <- typeChanReqSetting{
		actionType: "get client connection",
		clientID:   clientID,
		chanRes:    cr,
	}

	res := <-cr

	return res.connect, res.msgErr*/
}

//DelClientAPI удалить всю информацию о клиенте
func (mrpts *ModReqProcTemporaryStorage) DelClientAPI(clientID string) {

	fmt.Println("func 'DelClientAPI', START...")

	/*cr := make(chan typeChanResSetting)
	defer close(cr)

	smapi.chanReqSetting <- typeChanReqSetting{
		actionType: "del client",
		clientID:   clientID,
		chanRes:    cr,
	}

	<-cr*/
}
