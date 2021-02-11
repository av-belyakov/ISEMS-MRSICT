package temporarystorage

import (
	"sync"

	"github.com/gorilla/websocket"
)

//ModReqProcTemporaryStorage используется для хранения параметров клиентов
// clientSettings - настройки пользователя, КЛЮЧ уникальный идентификатор клиента
// chanReqParameters - канал для передачи параметров пользователя
type ModReqProcTemporaryStorage struct {
	clientSettings    map[string]*ModReqProcTemporaryStorageClientSettings
	chanReqParameters chan typeChanReqParam
}

//ModReqProcTemporaryStorageClientSettings параметры подключения клиента
// IP - IP адрес клиента
// Token - идентификационный токен клиента
// ClientName - имя клиента из config.json
// IsAllowed: разрешен ли доступ
// Connection - дескриптор соединения через websocket
type ModReqProcTemporaryStorageClientSettings struct {
	IP         string
	Token      string
	ClientName string
	IsAllowed  bool
	Connection *websocket.Conn
	mu         sync.Mutex
}

type typeChanReqParam struct {
	actionType string
	clientIP   string
	token      string
	clientID   string
	clientName string
	chanRes    chan typeChanResSetting
	typeReqResCommonParam
}

type typeChanResSetting struct {
	msgErr     error
	clientID   string
	clientList map[string]*ModReqProcTemporaryStorageClientSettings
	typeReqResCommonParam
}

type typeReqResCommonParam struct {
	clientSetting *ModReqProcTemporaryStorageClientSettings
	connect       *websocket.Conn
}
