package datamodels

import (
	mstixo "github.com/av-belyakov/methodstixobjects"
)

/********** 			Cyber-observable Objects STIX (ТИПЫ)			**********/

// OptionalCommonPropertiesCyberObservableObjectSTIX содержит опциональные общие свойства для Cyber-observable Objects STIX
// SpecVersion - версия STIX спецификации.
// ObjectMarkingRefs - определяет список ID ссылающиеся на объект "marking-definition", по терминалогии STIX, в котором содержатся
// значения применяющиеся к этому объекту
// GranularMarkings - определяет список "гранулярных меток" (granular_markings) относящихся к этому объекту
// Defanged - определяет были ли определены данные содержащиеся в объекте
// // Extensions - может содержать дополнительную информацию, относящуюся к объекту
type OptionalCommonPropertiesCyberObservableObjectSTIX struct {
	mstixo.OptionalCommonPropertiesCyberObservableObjectSTIX
}

// ArtifactCyberObservableObjectSTIX объект "Artifact", по терминалогии STIX, позволяет захватывать массив байтов (8 бит) в виде строки в кодировке base64
//
//	или связывать его с полезной нагрузкой, подобной файлу. Обязательно должен быть заполнено одно из полей PayloadBin или URL
//
// MimeType - по возможности это значение ДОЛЖНО быть одним из значений, определенных в реестре типов носителей IANA. В универсальном каталоге
//
//	всех существующих типов файлов.
//
// PayloadBin - бинарные данные в base64
// URL - унифицированный указатель ресурса (URL)
// Hashes - словарь хешей для URL или PayloadBin
// EncryptionAlgorithm - тип алгоритма шифрования для бинарных данных
// DecryptionKey - определяет ключ для дешифрования зашифрованных данных
type ArtifactCyberObservableObjectSTIX struct {
	mstixo.ArtifactCyberObservableObjectSTIX
}

// AutonomousSystemCyberObservableObjectSTIX объект "Autonomous System", по терминалогии STIX, содержит параметры Автономной системы
// Number - содержит номер присвоенный Автономной системе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Name - название Автономной системы
// RIR - содержит название регионального Интернет-реестра (Regional Internet Registry) которым было дано имя Автономной системы
type AutonomousSystemCyberObservableObjectSTIX struct {
	mstixo.AutonomousSystemCyberObservableObjectSTIX
}

// DirectoryCyberObservableObjectSTIX объект "Directory", по терминалогии STIX, содержит свойства, общие для каталога файловой системы
// Path - указывает путь, как было первоначально замечено, к каталогу в файловой системе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PathEnc - указывает наблюдаемую кодировку для пути. Значение ДОЛЖНО быть указано, если путь хранится в кодировке, отличной от Unicode.
// Ctime - время, в формате "2016-05-12T08:17:27.000Z", создания директории
// Mtime - время, в формате "2016-05-12T08:17:27.000Z", модификации или записи в директорию
// Atime - время, в формате "2016-05-12T08:17:27.000Z", последнего обращения к директории
// ContainsRefs - содержит список файловых объектов или директорий содержащихся внутри директории
type DirectoryCyberObservableObjectSTIX struct {
	mstixo.DirectoryCyberObservableObjectSTIX
}

// DomainNameCyberObservableObjectSTIX объект "Domain Name", по терминалогии STIX, содержит сетевое доменное имя
// Value - сетевое доменное имя (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ResolvesToRefs - список ссылок на один или несколько IP-адресов или доменных имен, на которые разрешается доменное имя
type DomainNameCyberObservableObjectSTIX struct {
	mstixo.DomainNameCyberObservableObjectSTIX
}

// EmailAddressCyberObservableObjectSTIX объект "Email Address", по терминалогии STIX, содержит представление единственного email адреса
// Value - email адрес (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// DisplayName - содержит единственное почтовое имя которое видит человек при просмотре письма
// BelongsToRef - учетная запись пользователя, которой принадлежит адрес электронной почты, в качестве ссылки на объект учетной записи пользователя
type EmailAddressCyberObservableObjectSTIX struct {
	mstixo.EmailAddressCyberObservableObjectSTIX
}

// EmailMessageCyberObservableObjectSTIX объект "Email Message", по терминалогии STIX, содержит экземпляр email сообщения
// IsMultipart - информирует содержит ли 'тело' email множественные MIME части (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Date - время, в формате "2016-05-12T08:17:27.000Z", когда email сообщение было отправлено
// ContentType - содержит содержимое 'Content-Type' заголовка email сообщения
// FromRef - содержит содержимое 'From:' заголовка email сообщения
// SenderRef - содержит содержимое поля 'Sender:' email сообщения
// ToRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'To:'
// CcRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'CC:'
// BccRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'BCC:'
// MessageID - содержимое Message-ID email сообщения
// Subject - содержит тему сообщения
// ReceivedLines - содержит одно или несколько полей заголовка 'Received', которые могут быть включены в заголовки email
// AdditionalHeaderFields - содержит любые другие поля заголовка (за исключением date, received_lines, content_type, from_ref,
//
//	sender_ref, to_ref, cc_ref, bcc_refs и subject), найденные в сообщении электронной почты в виде словаря
//
// Body - содержит тело сообщения
// BodyMultipart - содержит адает список MIME-части, которые составляют тело email. Это свойство НЕ ДОЛЖНО использоваться, если
//
//	is_multipart имеет значение false
//
// RawEmailRef - содержит 'сырое' бинарное содержимое email сообщения
type EmailMessageCyberObservableObjectSTIX struct {
	mstixo.EmailMessageCyberObservableObjectSTIX
}

// CommonFileCyberObservableObjectSTIX общий объект "File Object", по терминалогии STIX, содержит объект со свойствами файла
// Extensions - определяет следующие расширения pdf-ext, archive-ext, ntfs-ext, raster-image-ext, windows-pebinary-ext. В дополнении к ним пользователь может создавать
//
//	свои расширения. При этом ключ словаря должен однозначно идентифицировать тип расширения.
//
// Hashes - определяет словарь хешей для файла. При этом ДОЛЖНЫ использоватся ключи из открытого словаря hash-algorithm- ov.
// Size - содержит размер файла в байтах
// Name - содержит имя файла
// NameEnc - определяет кодировку имени файла. Содержимое должно соответствовать ревизии IANA от 2013-12-20.
// MagicNumberHex - указывает шестнадцатеричную константу (“магическое число”), связанную с определенным форматом файла, который соответствует этому файлу, если это применимо.
// MimeType - определяет MIME имени файла, например, application/msword.
// Ctime - время, в формате "2016-05-12T08:17:27.000Z", создания файла
// Mtime - время, в формате "2016-05-12T08:17:27.000Z", модификации файла
// Atime - время, в формате "2016-05-12T08:17:27.000Z", обращения к файлу
// ParentDirectoryRef - определяет родительскую директорию для файла. Объект ссылающийся на это свойство ДОЛЖЕН быть типом directory
// ContainsRefs - содержит ссылки на другие Cyber-observable Objects STIX, содержащиеся в файле, например другой файл, добавленный в конец файла, или IP-адрес, содержащийся где-то в файле.
// ContentRef - определяет контент файла. Данное значение ДОЛЖНО иметь тип artifact, то есть ссылатся на ArtifactCyberObservableObjectSTIX
type CommonFileCyberObservableObjectSTIX struct {
	mstixo.CommonFileCyberObservableObjectSTIX
}

// FileCyberObservableObjectSTIX объект "File Object", по терминалогии STIX, последекодирования из JSON (основной, рабочий объект)
// Extensions - определяет следующие расширения pdf-ext, archive-ext, ntfs-ext, raster-image-ext, windows-pebinary-ext. В дополнении к ним пользователь может создавать
//
//	свои расширения. При этом ключ словаря должен однозначно идентифицировать тип расширения.
//
// Hashes - определяет словарь хешей для файла. При этом ДОЛЖНЫ использоватся ключи из открытого словаря hash-algorithm- ov.
// Size - содержит размер файла в байтах
// Name - содержит имя файла
// NameEnc - определяет кодировку имени файла. Содержимое должно соответствовать ревизии IANA от 2013-12-20.
// MagicNumberHex - указывает шестнадцатеричную константу (“магическое число”), связанную с определенным форматом файла, который соответствует этому файлу, если это применимо.
// MimeType - определяет MIME имени файла, например, application/msword.
// Ctime - время, в формате "2016-05-12T08:17:27.000Z", создания файла
// Mtime - время, в формате "2016-05-12T08:17:27.000Z", модификации файла
// Atime - время, в формате "2016-05-12T08:17:27.000Z", обращения к файлу
// ParentDirectoryRef - определяет родительскую директорию для файла. Объект ссылающийся на это свойство ДОЛЖЕН быть типом directory
// ContainsRefs - содержит ссылки на другие Cyber-observable Objects STIX, содержащиеся в файле, например другой файл, добавленный в конец файла, или IP-адрес, содержащийся где-то в файле.
// ContentRef - определяет контент файла. Данное значение ДОЛЖНО иметь тип artifact, то есть ссылатся на ArtifactCyberObservableObjectSTIX
type FileCyberObservableObjectSTIX struct {
	mstixo.FileCyberObservableObjectSTIX
}

// IPv4AddressCyberObservableObjectSTIX объект "IPv4 Address Object", по терминалогии STIX, содержит один или более IPv4 адресов, выраженных с помощью нотации CIDR.
// Value - указывает значения одного или нескольких IPv4-адресов, выраженные с помощью нотации CIDR. Если данный объект IPv4-адреса представляет собой один IPv4-адрес,
//
//	суффикс CIDR /32 МОЖЕТ быть опущен. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
//
// ResolvesToRefs - указывает список ссылок на один или несколько MAC-адресов управления доступом к носителям уровня 2, на которые разрешается IPv6-адрес. Объекты,
//
//	на которые ссылается этот список, ДОЛЖНЫ иметь тип macaddr.
//
// BelongsToRefs - указывает список ссылок на одну или несколько автономных систем (AS), к которым принадлежит IPv6-адрес. Объекты, на которые ссылается этот список,
//
//	ДОЛЖНЫ быть типа autonomous-system.
type IPv4AddressCyberObservableObjectSTIX struct {
	mstixo.IPv4AddressCyberObservableObjectSTIX
}

// IPv6AddressCyberObservableObjectSTIX объект "IPv6 Address Object", по терминалогии STIX, содержит один или более IPv6 адресов, выраженных с помощью нотации CIDR.
// Value - указывает значения одного или нескольких IPv6-адресов, выраженные с помощью нотации CIDR. Если данный объект IPv6-адреса представляет собой один IPv6-адрес,
//
//	суффикс CIDR /128 МОЖЕТ быть опущен. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
//
// ResolvesToRefs - указывает список ссылок на один или несколько MAC-адресов управления доступом к носителям уровня 2, на которые разрешается IPv6-адрес. Объекты,
//
//	на которые ссылается этот список, ДОЛЖНЫ иметь тип macaddr.
//
// BelongsToRefs - указывает список ссылок на одну или несколько автономных систем (AS), к которым принадлежит IPv4-адрес. Объекты, на которые ссылается этот список,
//
//	ДОЛЖНЫ быть типа autonomous-system.
type IPv6AddressCyberObservableObjectSTIX struct {
	mstixo.IPv6AddressCyberObservableObjectSTIX
}

// MACAddressCyberObservableObjectSTIX объект "MAC Address Object", по терминалогии STIX, содержит объект MAC-адрес, представляющий собой один адрес управления доступом к среде (MAC).
// Value - Задает значение одного MAC-адреса. Значение MAC - адреса ДОЛЖНО быть представлено в виде одного строчного MAC-48 address, разделенного двоеточием,
//
//	который ДОЛЖЕН включать начальные нули для каждого октета. Пример: 00:00:ab:cd:ef:01. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type MACAddressCyberObservableObjectSTIX struct {
	mstixo.MACAddressCyberObservableObjectSTIX
}

// MutexCyberObservableObjectSTIX объект "Mutex Object", по терминалогии STIX, содержит свойства объекта взаимного исключения (mutex).
// Name - указывает имя объекта мьютекса (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
type MutexCyberObservableObjectSTIX struct {
	mstixo.MutexCyberObservableObjectSTIX
}

// CommonNetworkTrafficCyberObservableObjectSTIX общий объект "Network Traffic Object", по терминалогии STIX, содержит объект Сетевого трафика представляющий собой произвольный сетевой трафик,
//
//	который исходит из источника и адресуется адресату.
//
// Extensions - объект Сетевого трафика определяет следующие расширения. В дополнение к ним производители МОГУТ создавать свои собственные. ключи словаря http-request-ext, cp-ext,
//
//	icmp-ext, socket-ext ДОЛЖНЫ идентифицировать тип расширения по имени. Соответствующие значения словаря ДОЛЖНЫ содержать содержимое экземпляра расширения.
//
// Start - время, в формате "2016-05-12T08:17:27.000Z", инициирования сетевого трафика, если он известен.
// End - время, в формате "2016-05-12T08:17:27.000Z", окончания сетевого трафика, если он известен.
// IsActive - указывает, продолжается ли сетевой трафик. Если задано свойство end, то это свойство ДОЛЖНО быть false.
// SrcRef - указывает источник сетевого трафика в качестве ссылки на кибернаблюдаемый объект. Объект, на который ссылается ссылка, ДОЛЖЕН быть типа ipv4-addr, ipv6 - addr, mac-addr
//
//	или domain-name (для случаев, когда IP-адрес для доменного имени неизвестен).
//
// DstRef - указывает место назначения сетевого трафика в качестве ссылки на кибернаблюдаемый объект. Объект, на который ссылается ссылка, ДОЛЖЕН быть типа ipv4-addr, ipv6 - addr,
//
//	mac-addr или domain-name (для случаев, когда IP-адрес для доменного имени неизвестен).
//
// SrcPort - задает исходный порт, используемый в сетевом трафике, в виде целого числа. Значение порта ДОЛЖНО находиться в диапазоне от 0 до 65535.
// DstPort - задает порт назначения, используемый в сетевом трафике, в виде целого числа. Значение порта ДОЛЖНО находиться в диапазоне от 0 до 65535.
// Protocols - указывает протоколы, наблюдаемые в сетевом трафике, а также их соответствующее состояние.
// SrcByteCount - задает число байтов в виде положительного целого числа, отправленных от источника к месту назначения.
// DstByteCount - задает число байтов в виде положительного целого числа, отправленных из пункта назначения в источник.
// SrcPackets - задает количество пакетов в виде положительного целого числа, отправленных от источника к месту назначения.
// DstPackets - задает количество пакетов в виде положительного целого числа, отправленных от пункта назначения к источнику
// IPFix - указывает любые данные Экспорта информации IP-потока [IPFIX] для трафика в виде словаря. Каждая пара ключ/значение в словаре представляет имя/значение одного элемента IPFIX.
//
//	Соответственно, каждый ключ словаря ДОЛЖЕН быть сохраненной в регистре версией имени элемента IPFIX.
//
// SrcPayloadRef - указывает байты, отправленные из источника в пункт назначения. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
// DstPayloadRef - указывает байты, отправленные из пункта назначения в источник. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
// EncapsulatesRefs - ссылки на другие объекты, инкапсулированные этим объектом. Объекты, на которые ссылается это свойство, ДОЛЖНЫ иметь тип network-traffic.
// EncapsulatedByRef - ссылки на другой объект сетевого трафика, который инкапсулирует этот объект. Объекты, на которые ссылается это свойство, ДОЛЖНЫ иметь тип network-traffic.
type CommonNetworkTrafficCyberObservableObjectSTIX struct {
	mstixo.CommonNetworkTrafficCyberObservableObjectSTIX
}

// NetworkTrafficCyberObservableObjectSTIX объект "Network Traffic Object", по терминалогии STIX, содержит объект Сетевого трафика представляющий собой произвольный сетевой трафик,
//
//	который исходит из источника и адресуется адресату.
//
// Extensions - объект Сетевого трафика определяет следующие расширения. В дополнение к ним производители МОГУТ создавать свои собственные. ключи словаря http-request-ext, cp-ext,
// Start - время, в формате "2016-05-12T08:17:27.000Z", инициирования сетевого трафика, если он известен.
// End - время, в формате "2016-05-12T08:17:27.000Z", окончания сетевого трафика, если он известен.
// IsActive - указывает, продолжается ли сетевой трафик. Если задано свойство end, то это свойство ДОЛЖНО быть false.
// SrcRef - указывает источник сетевого трафика в качестве ссылки на кибернаблюдаемый объект. Объект, на который ссылается ссылка, ДОЛЖЕН быть типа ipv4-addr, ipv6 - addr, mac-addr
//
//	или domain-name (для случаев, когда IP-адрес для доменного имени неизвестен).
//
// DstRef - указывает место назначения сетевого трафика в качестве ссылки на кибернаблюдаемый объект. Объект, на который ссылается ссылка, ДОЛЖЕН быть типа ipv4-addr, ipv6 - addr,
//
//	mac-addr или domain-name (для случаев, когда IP-адрес для доменного имени неизвестен).
//
// SrcPort - задает исходный порт, используемый в сетевом трафике, в виде целого числа. Значение порта ДОЛЖНО находиться в диапазоне от 0 до 65535.
// DstPort - задает порт назначения, используемый в сетевом трафике, в виде целого числа. Значение порта ДОЛЖНО находиться в диапазоне от 0 до 65535.
// Protocols - указывает протоколы, наблюдаемые в сетевом трафике, а также их соответствующее состояние.
// SrcByteCount - задает число байтов в виде положительного целого числа, отправленных от источника к месту назначения.
// DstByteCount - задает число байтов в виде положительного целого числа, отправленных из пункта назначения в источник.
// SrcPackets - задает количество пакетов в виде положительного целого числа, отправленных от источника к месту назначения.
// DstPackets - задает количество пакетов в виде положительного целого числа, отправленных от пункта назначения к источнику
// IPFix - указывает любые данные Экспорта информации IP-потока [IPFIX] для трафика в виде словаря. Каждая пара ключ/значение в словаре представляет имя/значение одного элемента IPFIX.
//
//	Соответственно, каждый ключ словаря ДОЛЖЕН быть сохраненной в регистре версией имени элемента IPFIX.
//
// SrcPayloadRef - указывает байты, отправленные из источника в пункт назначения. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
// DstPayloadRef - указывает байты, отправленные из пункта назначения в источник. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
// EncapsulatesRefs - ссылки на другие объекты, инкапсулированные этим объектом. Объекты, на которые ссылается это свойство, ДОЛЖНЫ иметь тип network-traffic.
// EncapsulatedByRef - ссылки на другой объект сетевого трафика, который инкапсулирует этот объект. Объекты, на которые ссылается это свойство, ДОЛЖНЫ иметь тип network-traffic.
type NetworkTrafficCyberObservableObjectSTIX struct {
	mstixo.NetworkTrafficCyberObservableObjectSTIX
}

// CommonProcessCyberObservableObjectSTIX общий объект "Process Object", по терминологии STIX, содержит общие свойства экземпляра компьютерной программы,
//
//	выполняемой в операционной системе. Объект процесса ДОЛЖЕН содержать хотя бы одно свойство (отличное от типа) этого объекта (или одного из его расширений).
//
// Extensions - определяет расширения windows-process-exit или windows-service-ext. В дополнение к ним производители МОГУТ создавать свои собственные. ключи словаря windows-process-exit,
//
//	windows-service-ext ДОЛЖНЫ идентифицировать тип расширения по имени. Соответствующие значения словаря ДОЛЖНЫ содержать содержимое экземпляра расширения.
//
// IsHidden - определяет является ли процесс скрытым.
// PID - униальный идентификатор процесса.
// CreatedTime - время, в формате "2016-05-12T08:17:27.000Z", создания процесса.
// Cwd - текущая рабочая директория процесса.
// CommandLine - поределяет полный перечень команд используемых для запуска процесса, включая имя процесса и аргументы.
// EnvironmentVariables - определяет список переменных окружения, в виде словаря, ассоциируемых с приложением.
// OpenedConnectionRefs - определяет список открытых, процессом, сетевых соединений ка одну или более ссылку на объект типа network-traffic.
// CreatorUserRef - определяет что за пользователь создал объект, ссылка ДОЛЖНА ссылатся на объект типа user-account.
// ImageRef - указывает исполняемый двоичный файл, который был выполнен как образ процесса, как ссылка на файловый объект. Объект, на который ссылается
//
//	это свойство, ДОЛЖЕН иметь тип file.
//
// ParentRef - указывает другой процесс, который породил (т. е. является родителем) этот процесс, как ссылку на объект процесса. Объект, на который
//
//	ссылается это свойство, ДОЛЖЕН иметь тип process.
//
// ChildRefs - указывает другие процессы, которые были порождены (т. е. дочерние) этим процессом, в качестве ссылки на один или несколько других
//
//	объектов процесса. Объекты, на которые ссылается этот список, ДОЛЖНЫ иметь тип process.
type CommonProcessCyberObservableObjectSTIX struct {
	mstixo.CommonProcessCyberObservableObjectSTIX
}

// ProcessCyberObservableObjectSTIX объект "Process Object", по терминологии STIX, содержит общие свойства экземпляра компьютерной программы,
//
//	выполняемой в операционной системе. Объект процесса ДОЛЖЕН содержать хотя бы одно свойство (отличное от типа) этого объекта (или одного из его расширений).
//
// Extensions - определяет расширения windows-process-exit или windows-service-ext. В дополнение к ним производители МОГУТ создавать свои собственные. ключи словаря windows-process-exit,
//
//	windows-service-ext ДОЛЖНЫ идентифицировать тип расширения по имени. Соответствующие значения словаря ДОЛЖНЫ содержать содержимое экземпляра расширения.
//
// IsHidden - определяет является ли процесс скрытым.
// PID - униальный идентификатор процесса.
// CreatedTime - время, в формате "2016-05-12T08:17:27.000Z", создания процесса.
// Cwd - текущая рабочая директория процесса.
// CommandLine - поределяет полный перечень команд используемых для запуска процесса, включая имя процесса и аргументы.
// EnvironmentVariables - определяет список переменных окружения, в виде словаря, ассоциируемых с приложением.
// OpenedConnectionRefs - определяет список открытых, процессом, сетевых соединений ка одну или более ссылку на объект типа network-traffic.
// CreatorUserRef - определяет что за пользователь создал объект, ссылка ДОЛЖНА ссылатся на объект типа user-account.
// ImageRef - указывает исполняемый двоичный файл, который был выполнен как образ процесса, как ссылка на файловый объект. Объект, на который ссылается
//
//	это свойство, ДОЛЖЕН иметь тип file.
//
// ParentRef - указывает другой процесс, который породил (т. е. является родителем) этот процесс, как ссылку на объект процесса. Объект, на который
//
//	ссылается это свойство, ДОЛЖЕН иметь тип process.
//
// ChildRefs - указывает другие процессы, которые были порождены (т. е. дочерние) этим процессом, в качестве ссылки на один или несколько других
//
//	объектов процесса. Объекты, на которые ссылается этот список, ДОЛЖНЫ иметь тип process.
type ProcessCyberObservableObjectSTIX struct {
	mstixo.ProcessCyberObservableObjectSTIX
}

// SoftwareCyberObservableObjectSTIX объект "Software Object", по терминологии STIX, содержит свойства, связанные с программным обеспечением, включая программные продукты.
// Name - назвыание программного обеспечения
// CPE - содержит запись Common Platform Enumeration (CPE) для программного обеспечения, если она доступна. Значение этого свойства должно быть значением
// CPE v2.3 из официального словаря NVD CPE [NVD]
// SwID - содержит запись Тегов Software Identification ID (SWID) [SWID] для программного обеспечения, если таковая имеется. SwID помеченный tagId, является
//
//	глобально уникальным идентификатором и ДОЛЖЕН использоваться как полномочие для идентификации помеченного продукта
//
// Languages -содержит языки, поддерживаемые программным обеспечением. Значение каждого елемента списка ДОЛЖНО быть кодом языка ISO 639-2
// Vendor - содержит название производителя программного обеспечения
// Version - содержит версию ПО
type SoftwareCyberObservableObjectSTIX struct {
	mstixo.SoftwareCyberObservableObjectSTIX
}

// URLCyberObservableObjectSTIX объект "URL Object", по терминологии STIX, содержит унифицированный указатель информационного ресурса (URL).
// Value - содержит унифицированный указатель информационного ресурса (URL).
type URLCyberObservableObjectSTIX struct {
	mstixo.URLCyberObservableObjectSTIX
}

// UserAccountCyberObservableObjectSTIX объект "User Account Object", по терминалогии STIX, содержит экземпляр любого типа учетной записи пользователя, включая,
// учетные записи операционной системы, устройства, службы обмена сообщениями и платформы социальных сетей и других прочих учетных записей
// Поскольку все свойства этого объекта являются необязательными, по крайней мере одно из свойств, определенных ниже, ДОЛЖНО быть инициализировано
// при использовании этого объекта
// Extensions - содержит словарь расширяющий тип "User Account Object" одно из расширений "unix-account-ext", реализуется описанным ниже типом, UNIXAccountExtensionSTIX
//
//	кроме этого производитель может созавать свои собственные типы расширений
//	Ключи данного словаря идентифицируют тип расширения по имени, значения являются содержимым экземпляра расширения
//
// UserID - содержит идентификатор учетной записи. Формат идентификатора зависит от системы в которой находится данная учетная запись пользователя,
//
//	и может быть числовым идентификатором, идентификатором GUID, именем учетной записи, адресом электронной почты и т.д. Свойство  UserId должно
//	быть заполнено любым значанием, являющимся уникальным идентификатором системы, членом которой является учетная запись. Например, в системах UNIX он
//	будет заполнено значением UID
//
// Credential - содержит учетные данные пользователя в открытом виде. Предназначено только для закрытого применения при изучении метаданных вредоносных программ
//
//	 при их исследовании (например, жестко закодированный пароль администратора домена, который вредоносная программа пытается использовать реализации тактики для
//		бокового (латерального) перемещения) и не должно применяться для совместного пользования PII
//
// AccountLogin - содержит логин пользователя. Используется в тех случаях,когда свойство UserId указывает другие данные, чем то, что пользователь вводит
//
//	при входе в систему
//
// AccountType - содержит одно, из заранее определенных (предложенных) значений. Является типом аккаунта. Значения этого свойства берутся из множества
//
//	закрепленного в открытом словаре, account-type-ov
//
// DisplayName - содержит отображаемое имя учетной записи, которое будет отображаться в пользовательских интерфейсах. В Unix, это равносильно полю gecos
//
//	(gecos это поле учётной записи пользователя в файле /etc/passwd )
//
// IsServiceAccount - содержит индикатор, сигнализирующий что, учетная запись связана с сетевой службой или системным процессом (демоном), а не с конкретным человеком. (системный пользователь)
// IsPrivileged - содержит индикатор, сигнализирующий что, учетная запись имеет повышенные привилегии (например, в случае root в Unix или учетной записи администратора
//
//	Windows)
//
// CanEscalatePrivs  - содержит индикатор, сигнализирующий что, учетная запись имеет возможность повышать привилегии (например, в случае sudo в Unix или учетной
//
//	записи администратора домена Windows)
//
// IsDisabled  - содержит индикатор, сигнализирующий что, учетная запись отключена
// AccountCreated - время, в формате "2016-05-12T08:17:27.000Z", создания аккаунта
// AccountExpires - время, в формате "2016-05-12T08:17:27.000Z", истечения срока действия учетной записи.
// CredentialLastChanged - время, в формате "2016-05-12T08:17:27.000Z", когда учетные данные учетной записи были изменены в последний раз.
// AccountFirstLogin - время, в формате "2016-05-12T08:17:27.000Z", первого доступа к учетной записи
// AccountLastLogin - время, в формате "2016-05-12T08:17:27.000Z", когда к учетной записи был последний доступ.
type UserAccountCyberObservableObjectSTIX struct {
	mstixo.UserAccountCyberObservableObjectSTIX
}

// WindowsRegistryKeyCyberObservableObjectSTIX объект "Windows Registry Key Object", по терминалогии STIX. Содержит описание значений полей раздела реестра Windows.
//
//	Поскольку все свойства этого объекта являются необязательными, по крайней мере одно из свойств,определенных ниже, должно быть инициализировано при
//	использовании этого объекта.
//
// Key - содержит полный путь к разделу реестра. Значение ключа,должно быть сохранено в регистре. В название ключа все сокращения должны быть раскрыты.
//
//	Например, вместо HKLM следует использовать HKEY_LOCAL_MACHINE.
//
// Values - содержит значения, найденные в разделе реестра.
// ModifiedTime - время, в формате "2016-05-12T08:17:27.000Z", последнего изменения раздела реестра.
// CreatorUserRef - содержит ссылку на учетную запись пользователя, из под которой создан раздел реестра. Объект, на который ссылается это свойство, должен иметь тип user-account.
// NumberOfSubkeys - Указывает количество подразделов, содержащихся в разделе реестра.
type WindowsRegistryKeyCyberObservableObjectSTIX struct {
	mstixo.WindowsRegistryKeyCyberObservableObjectSTIX
}

// X509CertificateCyberObservableObjectSTIX объект "X.509 Certificate Object", по терминологии STIX, представлет свойства сертификата X.509, определенные в рекомендациях
//
//	ITU X.509 [X.509]. X.509  Certificate объект должен содержать по крайней мере одно cвойство специфичное для этого объекта (помимо type).
//
// IsSelfSigned - содержит индикатор, является ли сертификат самоподписным, то есть подписан ли он тем же субъектом, личность которого он удостоверяет.
// Hashes - содержит любые хэши, которые были вычислены для всего содержимого сертификата. Является типом данных словар, значения ключей которого должны
//
//	быть из открытого словаря hash-algorithm-ov.
//
// Version- содержит версию закодированного сертификата
// SerialNumber - содержит уникальный идентификатор сертификата, выданного конкретным Центром сертификации.
// SignatureAlgorithm - содержит имя алгоритма, используемого для подписи сертификата.
// Issuer - содержит название удостоверяющего центра выдавшего сертификат
// ValidityNotBefore - время, в формате "2016-05-12T08:17:27.000Z", начала действия сертификата.
// ValidityNotAfter - время, в формате "2016-05-12T08:17:27.000Z", окончания действия сертификата.
// Subject - содержит имя сущности, связанной с открытым ключом, хранящимся в поле "subject public key" открого ключа сертификата.
// SubjectPublicKeyAlgorithm - содержит название алгоритма применяемого для шифрования данных, отправляемых субъекту.
// SubjectPublicKeyModulus - указывает модульную часть открытого ключа RSA.
// SubjectPublicKeyExponent - указывает экспоненциальную часть открытого ключа RSA субъекта в виде целого числа.
// X509V3Extension - указывает любые стандартные расширения X.509 v3, которые могут использоваться в сертификате.
type X509CertificateCyberObservableObjectSTIX struct {
	mstixo.X509CertificateCyberObservableObjectSTIX
}
