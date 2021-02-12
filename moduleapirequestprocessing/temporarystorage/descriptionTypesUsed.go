package temporarystorage

import (
	"sync"

	"github.com/gorilla/websocket"
)

//RepositoryStorageUserParametersType используется для хранения параметров клиентов
// clientSettings - настройки пользователя, КЛЮЧ уникальный идентификатор клиента
// chanReqParameters - канал для передачи параметров пользователя
type RepositoryStorageUserParametersType struct {
	clientSettings    map[string]*StorageUserParameters
	chanReqParameters chan typeChanReqParameters
}

//StorageUserParameters параметры подключения клиента
// IP - IP адрес клиента
// Token - идентификационный токен клиента
// ClientName - имя клиента из config.json
// IsAllowed: разрешен ли доступ
// Connection - дескриптор соединения через websocket
type StorageUserParameters struct {
	IP         string
	Token      string
	ClientName string
	IsAllowed  bool
	Connection *websocket.Conn
	mu         sync.Mutex
}

type typeChanReqParameters struct {
	actionType string
	clientIP   string
	token      string
	clientID   string
	clientName string
	chanRes    chan typeChanResParameters
	typeReqResCommonParameters
}

type typeChanResParameters struct {
	msgErr     error
	clientID   string
	clientList map[string]*StorageUserParameters
	typeReqResCommonParameters
}

type typeReqResCommonParameters struct {
	clientSetting *StorageUserParameters
	connect       *websocket.Conn
}

//SendWsMessage используется для отправки сообщений через протокол websocket (применяется Mutex)
func (sup *StorageUserParameters) SendWsMessage(t int, v []byte) error {
	sup.mu.Lock()
	defer sup.mu.Unlock()

	return sup.Connection.WriteMessage(t, v)
}
