package datamodels

//AppConfig хранит настройки из конфигурационного файла приложения
// VersionApp - версия приложения
// RootDir - корневая директория приложения
// ConnectionsDataBase - настройки доступа к системам управления базами данных
// ModuleApiRequestProcessingSettings - настройки сетевых соеденений обеспечивающие доступ к подсистеме isems-mrsict из внешнего программного обеспечения
// ModuleAPIInteractionExternalSoftware - настройки доступа к внешнему программному обеспечению являющемуся источником информации о компьютерных угрозах
// CryptographySettings - настройки связанные с криптографией и защите каналов связи
// PathStorageDownloadedFiles - место для хранения загруженных файлов
// PathLogFiles - место расположение лог-файла приложения
type AppConfig struct {
	VersionApp                                   string
	RootDir                                      string
	ConnectionsDataBase                          ConnectionsDataBase
	ModuleAPIRequestProcessingSettings           []ModuleAPIRequestProcessingSetting
	ModuleAPIInteractionExternalSoftwareSettings ModuleAPIInteractionExternalSoftwareSettings
	CryptographySettings                         CryptographySettings
	PathStorageDownloadedFiles                   string
	PathLogFiles                                 string
}

//ConnectionsDataBase хранит настройки доступа к системам управления базами данных
type ConnectionsDataBase struct {
	MongoDBSettings MongoDBSettings
}

//MongoDBSettings хранит настройки доступа к системе управления базами данных MongoDB
// IsUseSocket - будет ли использоваться соединение через Unix socket
// Host - ip адрес сервера СУБД
// Port - порт сервера СУБД
// User - имя пользователя, используемое для соединения с СУБД
// Password - пароль, используемый для соединения с СУБД
// NameDB - название базы данных
// UnixSocketPath - месторасположение дескриптора Unix socket
type MongoDBSettings struct {
	IsUseSocket    bool
	Host           string
	Port           int
	User           string
	Password       string
	NameDB         string
	UnixSocketPath string
}

//ModuleAPIRequestProcessingSetting хранит настройки сетевых соеденений обеспечивающие доступ к подсистеме isems-mrsict из внешнего программного обеспечения
// ClientName - имя клиента
// Token - уникальный идентификационный токен клиента
type ModuleAPIRequestProcessingSetting struct {
	ClientName, Token string
}

//ModuleAPIInteractionExternalSoftwareSettings хранит настройки доступа к внешнему программному обеспечению являющемуся источником информации о компьютерных угрозах
type ModuleAPIInteractionExternalSoftwareSettings struct {
}

//CryptographySettings хранит настройки связанные с криптографией и защите каналов связи
// PathRootCAModuleAPIRequestProcessingSettings - месторасположение корневого сертификата для сервера модуля 'moduleApiRequestProcessingSettings'
type CryptographySettings struct {
	PathRootCAModuleAPIRequestProcessingSettings string `json:"pathRootCA_of_moduleApiRequestProcessingSettings"`
}
