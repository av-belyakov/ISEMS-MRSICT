package datamodels

import (
	"encoding/json"
	"time"

	"ISEMS-MRSICT/commonlibs"
)

/********** 			Некоторые 'сложные' типы относящиеся к Cyber-observable Objects STIX 			**********/

//AlternateDataStreamTypeSTIX тип "alternate-data-stream-type", по терминалогии STIX, содержит альтернативные потоки данных для NTFS
// Name - определяет имя альтернативного потока данных (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Hashes - задает словарь хэшей для данных, содержащихся в альтернативном потоке данных. Ключи для данного свойства ДОЛЖНЫ братся из открытого словоря hash-algorithm-ov.
// Size - определяет размер альтернативного потока данных, в байтах
type AlternateDataStreamTypeSTIX struct {
	Name   string         `json:"name" bson:"name"`
	Hashes HashesTypeSTIX `json:"hashes" bson:"hashes"`
	Size   uint64         `json:"size" bson:"size"`
}

//WindowsPEOptionalHeaderTypeSTIX тип "windows-pe-optional-header-type", по терминалогии STIX, содержит свойства опциональные для PE файла
// MagicHex - хеш значение информирующее о типе PE файла
// MajorLinkerVersion - указывает основной номер версии компоновщика.
// MinorLinkerVersion - указывает номер младшей версии компоновщика.
// SizeOfCode - указывает размер кода (текста) секции.
// SizeOfInitializedData - задает размер инициализированного раздела данных.
// SizeOfUninitializedData - задает размер раздела неинициализированных данных.
// AddressOfEntryPoint - указывает адрес точки входа относительно базы изображения при загрузке исполняемого файла в память.
// BaseOfCode - указывает адрес, который находится относительно базы изображений раздела начала кода при его загрузке в память.
// BaseOfData - указывает адрес, который находится относительно базы изображений раздела beginning-of-data при его загрузке в память.
// ImageBase - задает предпочтительный адрес первого байта изображения при загрузке в память.
// SectionAlignment - задает выравнивание (в байтах) PE-секций при их загрузке в память.
// FileAlignment - задает коэффициент (в байтах), который используется для выравнивания необработанных данных разделов в файле изображения.
// MajorOSVersion - указывает основной номер версии требуемой операционной системы.
// MinorOSVersion - указывает младший номер версии требуемой операционной системы.
// MajorImageVersion - указывает основной номер версии образа.
// MinorImageVersion - Указывает младший номер версии образа.
// MajorSubsystemVersion - указывает основной номер версии подсистемы.
// MinorSubsystemVersion - указывает младший номер версии подсистемы.
// Win32VersionValueHex - указывает зарезервированное значение версии win32.
// SizeOfImage - задает размер изображения в байтах, включая все заголовки, по мере загрузки изображения в память.
// SizeOfHeaders - задает комбинированный размер заголовков MS-DOS, PE и разделов, округленный до кратного значения, указанного в заголовке file_alignment.
// ChecksumHex - задает контрольную сумму двоичного файла PE.
// SubsystemHex - указывает подсистему (например, графический интерфейс, драйвер устройства и т.д.), необходимую для запуска этого образа.
// DllCharacteristicsHex - задает флаги, характеризующие двоичный файл PE.
// SizeOfStackReserve -указывает размер стека для резервирования в байтах.
// SizeOfStackCommit - указывает размер стека для фиксации в байтах.
// SizeOfHeapReserve - задает размер локального пространства кучи для резервирования в байтах.
// SizeOfHeapCommit - указывает размер локального пространства кучи для фиксации в байтах.
// LoaderFlagsHex - указывает зарезервированные флаги загрузчика.
// NumberOfRvaAndSizes - указывает количество записей каталога данных в оставшейся части необязательного заголовка.
// Hashes - указывает все хэши, которые были вычислены для необязательного заголовка.
type WindowsPEOptionalHeaderTypeSTIX struct {
	MagicHex                string            `json:"magic_hex" bson:"magic_hex"`
	MajorLinkerVersion      int               `json:"major_linker_version" bson:"major_linker_version"`
	MinorLinkerVersion      int               `json:"minor_linker_version" bson:"minor_linker_version"`
	SizeOfCode              int               `json:"size_of_code" bson:"size_of_code"`
	SizeOfInitializedData   int               `json:"size_of_initialized_data" bson:"size_of_initialized_data"`
	SizeOfUninitializedData int               `json:"size_of_uninitialized_data" bson:"size_of_uninitialized_data"`
	AddressOfEntryPoint     int               `json:"address_of_entry_point" bson:"address_of_entry_point"`
	BaseOfCode              int               `json:"base_of_code" bson:"base_of_code"`
	BaseOfData              int               `json:"base_of_data" bson:"base_of_data"`
	ImageBase               int               `json:"image_base" bson:"image_base"`
	SectionAlignment        int               `json:"section_alignment" bson:"section_alignment"`
	FileAlignment           int               `json:"file_alignment" bson:"file_alignment"`
	MajorOSVersion          int               `json:"major_os_version" bson:"major_os_version"`
	MinorOSVersion          int               `json:"minor_os_version" bson:"minor_os_version"`
	MajorImageVersion       int               `json:"major_image_version" bson:"major_image_version"`
	MinorImageVersion       int               `json:"minor_image_version" bson:"minor_image_version"`
	MajorSubsystemVersion   int               `json:"major_subsystem_version" bson:"major_subsystem_version"`
	MinorSubsystemVersion   int               `json:"minor_subsystem_version" bson:"minor_subsystem_version"`
	Win32VersionValueHex    string            `json:"win32_version_value_hex" bson:"win32_version_value_hex"`
	SizeOfImage             int               `json:"size_of_image" bson:"size_of_image"`
	SizeOfHeaders           int               `json:"size_of_headers" bson:"size_of_headers"`
	ChecksumHex             string            `json:"checksum_hex" bson:"checksum_hex"`
	SubsystemHex            string            `json:"subsystem_hex" bson:"subsystem_hex"`
	DllCharacteristicsHex   string            `json:"dll_characteristics_hex" bson:"dll_characteristics_hex"`
	SizeOfStackReserve      int               `json:"size_of_stack_reserve" bson:"size_of_stack_reserve"`
	SizeOfStackCommit       int               `json:"size_of_stack_commit" bson:"size_of_stack_commit"`
	SizeOfHeapReserve       int               `json:"size_of_heap_reserve" bson:"size_of_heap_reserve"`
	SizeOfHeapCommit        int               `json:"size_of_heap_commit" bson:"size_of_heap_commit"`
	LoaderFlagsHex          string            `json:"loader_flags_hex" bson:"loader_flags_hex"`
	NumberOfRvaAndSizes     int               `json:"number_of_rva_and_sizes" bson:"number_of_rva_and_sizes"`
	Hashes                  []*HashesTypeSTIX `json:"hashes" bson:"hashes"`
}

//SanitizeStructWindowsPEOptionalHeaderTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (wpeoh WindowsPEOptionalHeaderTypeSTIX) SanitizeStructWindowsPEOptionalHeaderTypeSTIX() WindowsPEOptionalHeaderTypeSTIX {
	wpeoh.MagicHex = commonlibs.StringSanitize(wpeoh.MagicHex)
	wpeoh.Win32VersionValueHex = commonlibs.StringSanitize(wpeoh.Win32VersionValueHex)
	wpeoh.ChecksumHex = commonlibs.StringSanitize(wpeoh.ChecksumHex)
	wpeoh.SubsystemHex = commonlibs.StringSanitize(wpeoh.SubsystemHex)
	wpeoh.DllCharacteristicsHex = commonlibs.StringSanitize(wpeoh.DllCharacteristicsHex)

	hsize := len(wpeoh.Hashes)
	if hsize == 0 {
		return wpeoh
	}

	nhashex := make([]*HashesTypeSTIX, 0, hsize)
	for _, value := range wpeoh.Hashes {
		ht := make(HashesTypeSTIX, len(*value))
		for k, v := range *value {
			ht[k] = v
		}
		nhashex = append(nhashex, &ht)
	}
	wpeoh.Hashes = nhashex

	return wpeoh
}

//WindowsPESectionTypeSTIX тип "windows-pe-section-type", по терминалогии STIX, определяет методанные PE-секции файла
// Name - определяет имя секции (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Size - размер секции в байтах.
// Entropy - задает вычисленную энтропию для сечения, рассчитанную с использованием алгоритма Шеннона.
// Hashes - задает любые хэши, вычисляемые по разделу. Свойство ДОЛЖНО содержать ключи из открытого словаря hash-algorithm-ov
type WindowsPESectionTypeSTIX struct {
	Name    string         `json:"name" bson:"name"`
	Size    uint64         `json:"size" bson:"size"`
	Entropy float64        `json:"entropy" bson:"entropy"`
	Hashes  HashesTypeSTIX `json:"hashes" bson:"hashes"`
}

//SanitizeStructWindowsPESectionTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (wpes WindowsPESectionTypeSTIX) SanitizeStructWindowsPESectionTypeSTIX() WindowsPESectionTypeSTIX {
	wpes.Name = commonlibs.StringSanitize(wpes.Name)

	hsize := len(wpes.Hashes)
	if hsize == 0 {
		return wpes
	}

	nh := make(HashesTypeSTIX, hsize)
	for k, v := range wpes.Hashes {
		nh[k] = commonlibs.StringSanitize(v)
	}
	wpes.Hashes = nh

	return wpes
}

//EmailMIMEPartTypeSTIX тип "email-mime-part-type", по терминалогии STIX, содержит один компонент тела email из нескольких частей
// Body - содержит содержимое части MIME, если content_type не указан или начинается с text/ (например, в случае обычного текста или HTML-письма)
// BodyRawRef - содержит содержимое нетекстовых частей MIME, то есть тех, чей content_type не начинается с text, в качестве
//  ссылки на объект артефакта или Файловый объект
// ContentType - содержимое поля 'Content-Type' заголовка MIME части email
// ContentDisposition - содержимое поля 'Content-Disposition' заголовка MIME части email
type EmailMIMEPartTypeSTIX struct {
	Body               string             `json:"body" bson:"body"`
	BodyRawRef         IdentifierTypeSTIX `json:"body_raw_ref" bson:"body_raw_ref"`
	ContentType        string             `json:"content_type" bson:"content_type"`
	ContentDisposition string             `json:"content_disposition" bson:"content_disposition"`
}

//WindowsRegistryValueTypeSTIX объект "Windows Registry Value Type", по терминалогии STIX. Данный тип фиксирует
//  значения свойств находящихся в разделе реестра Windows. Поскольку все свойства этого типа являются необязательными,
// по крайней мере одно из свойств, определенных ниже, должно быть инициализировано при использовании этого типа.
// Name - содержит название параметра реестра. Для указания значения ключа реестра по умолчанию необходимо использовать пустую строку.
// Data - содержит данные, содержащиеся в значении реестра.
// DataType - содержит тип данных реестра (REG_*), используемый в значении реестра. Значения этого свойства должны быть получены из перечисления windows-registry-datatype enum.
type WindowsRegistryValueTypeSTIX struct {
	Name     string       `json:"name" bson:"name"`
	Data     string       `json:"data" bson:"data"`
	DataType EnumTypeSTIX `json:"data_type" bson:"data_type"`
}

//SanitizeStructWindowsRegistryValueTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (wrv WindowsRegistryValueTypeSTIX) SanitizeStructWindowsRegistryValueTypeSTIX() WindowsRegistryValueTypeSTIX {
	wrv.Name = commonlibs.StringSanitize(wrv.Name)
	wrv.Data = commonlibs.StringSanitize(wrv.Data)
	wrv.DataType = EnumTypeSTIX(commonlibs.StringSanitize(string(wrv.DataType)))

	return wrv
}

//X509V3ExtensionsTypeSTIX тип "X.509 v3 Extensions Type", по терминалогии STIX. Описывает поля расширения X.509 v3, фиксрует дополнительную информацию,
//  такую как альтернативные имена субъектов. Объект, использующий тип "x509-v3-extensions-type", должен определить хотя бы одно из этих полей в нем.
//  Данный тип расширяет только объекты "X.509 Certificate Object".
// BasicConstraints - задает многозначное расширение, которое указывает, является ли сертификат сертификатом Удостоверяющего центра (CA). Первое (обязательное)
//  название CA сопровождается истинным или ложным. Если CA имеет значение TRUE, то может быть включено необязательное имя pathlen, за которым следует
//  неотрицательное значение. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.19.
// NameConstraints - указывает пространство имен, в котором должны располагаться все имена  применяемые в сертификатах указанных в пути сертификации. Также
//  эквивалентно значению идентификатора объекта (OID) 2.5.29.30.
// PolicyConstraints - содержит любые ограничения на проверку сертификатов, выданных Удостоверяющим центром.  Также эквивалентно значению идентификатора
//  объекта (OID) 2.5.29.36.
// KeyUsage - задает многозначное расширение, состоящее из списка имен разрешенных для использования ключей. Также эквивалентно значению идентификатора
//  объекта (OID) 2.5.29.15.
// ExtendedKeyUsage - содержит список целей, для которых может использоваться открытый ключ сертификата. Также эквивалентно значению идентификатора объекта
//  (OID) 2.5.29.37.
// SubjectKeyIdentifier - указывает идентификатор, который обеспечивает средство идентификации сертификатов, содержащих определенный открытый ключ. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.14.
// AuthorityKeyIdentifier - содержит идентификатор, который обеспечивает средство идентификации открытого ключа, соответствующего закрытому ключу, используемому
//  для подписи сертификата. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.35.
// SubjectAlternativeName - указывает дополнительные идентификаторы, которые должны быть привязаны к субъекту сертификата. Также эквивалентно значению
//  идентификатора объекта (OID) 2.5.29.17.
// IssuerAlternativeName - указывает дополнительные идентификаторы, которые должны быть привязаны к эмитенту сертификата. Также эквивалентно значению
//  идентификатора объекта (OID) 2.5.29.18.
// SubjectDirectoryAttributes - указывает идентификационные признаки (например, национальность) субъекта. Также эквивалентно значению идентификатора
//  объекта (OID) 2.5.29.9.
// CrlDistributionPoints - указывает способ получения информации CRL. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.31.
// InhibitAnyPolicy - содержит количество дополнительных сертификатов, которые могут появиться в пути до того, как любая Политика больше не будет разрешена.
//  Также эквивалентно значению идентификатора объекта (OID) 2.5.29.54.
// PrivateKeyUsagePeriodNotBefore - время, в формате "2016-05-12T08:17:27.000Z", начала срока действия закрытого ключа, если он отличается от срока действия сертификата.
// PrivateKeyUsagePeriodNotAfter -  время, в формате "2016-05-12T08:17:27.000Z", окончания срока действия закрытого ключа, если он отличается от срока действия сертификата.
// CertificatePolicies - содержит последовательность из одного или нескольких терминов информации о политике, каждый из которых состоит из идентификатора
//  объекта (OID) и необязательных квалификаторов. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.32.
// PolicyMappings - содержит одну или несколько пар OID; каждая пара включает issuerDomainPolicy и subjectDomainPolicy. Пары индикаторов указывают на то,
//  считает ли выдающий УЦ свою issuerDomainPolicy эквивалентной subjectDomainPolicy субъекта УЦ. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.33.
type X509V3ExtensionsTypeSTIX struct {
	BasicConstraints               string    `json:"basic_constraints" bson:"basic_constraints"`
	NameConstraints                string    `json:"name_constraints" bson:"name_constraints"`
	PolicyContraints               string    `json:"policy_contraints" bson:"policy_contraints"`
	KeyUsage                       string    `json:"key_usage" bson:"key_usage"`
	ExtendedKeyUsage               string    `json:"extended_key_usage" bson:"extended_key_usage"`
	SubjectKeyIdentifier           string    `json:"subject_key_identifier" bson:"subject_key_identifier"`
	AuthorityKeyIdentifier         string    `json:"authority_key_identifier" bson:"authority_key_identifier"`
	SubjectAlternativeName         string    `json:"subject_alternative_name" bson:"subject_alternative_name"`
	IssuerAlternativeName          string    `json:"issuer_alternative_name" bson:"issuer_alternative_name"`
	SubjectDirectoryAttributes     string    `json:"subject_directory_attributes" bson:"subject_directory_attributes"`
	CrlDistributionPoints          string    `json:"crl_distribution_points" bson:"crl_distribution_points"`
	InhibitAnyPolicy               string    `json:"inhibit_any_policy" bson:"inhibit_any_policy"`
	PrivateKeyUsagePeriodNotBefore time.Time `json:"private_key_usage_period_not_before" bson:"private_key_usage_period_not_before"`
	PrivateKeyUsagePeriodNotAfter  time.Time `json:"private_key_usage_period_not_after" bson:"private_key_usage_period_not_after"`
	CertificatePolicies            string    `json:"certificate_policies" bson:"certificate_policies"`
	PolicyMappings                 string    `json:"policy_mappings" bson:"policy_mappings"`
}

//SanitizeStructX509V3ExtensionsTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (x509v3e X509V3ExtensionsTypeSTIX) SanitizeStructX509V3ExtensionsTypeSTIX() X509V3ExtensionsTypeSTIX {
	x509v3e.BasicConstraints = commonlibs.StringSanitize(x509v3e.BasicConstraints)
	x509v3e.NameConstraints = commonlibs.StringSanitize(x509v3e.NameConstraints)
	x509v3e.PolicyContraints = commonlibs.StringSanitize(x509v3e.PolicyContraints)
	x509v3e.KeyUsage = commonlibs.StringSanitize(x509v3e.KeyUsage)
	x509v3e.ExtendedKeyUsage = commonlibs.StringSanitize(x509v3e.ExtendedKeyUsage)
	x509v3e.SubjectKeyIdentifier = commonlibs.StringSanitize(x509v3e.SubjectKeyIdentifier)
	x509v3e.AuthorityKeyIdentifier = commonlibs.StringSanitize(x509v3e.AuthorityKeyIdentifier)
	x509v3e.SubjectAlternativeName = commonlibs.StringSanitize(x509v3e.SubjectAlternativeName)
	x509v3e.IssuerAlternativeName = commonlibs.StringSanitize(x509v3e.IssuerAlternativeName)
	x509v3e.SubjectDirectoryAttributes = commonlibs.StringSanitize(x509v3e.SubjectDirectoryAttributes)
	x509v3e.CrlDistributionPoints = commonlibs.StringSanitize(x509v3e.CrlDistributionPoints)
	x509v3e.InhibitAnyPolicy = commonlibs.StringSanitize(x509v3e.InhibitAnyPolicy)
	x509v3e.CertificatePolicies = commonlibs.StringSanitize(x509v3e.CertificatePolicies)
	x509v3e.PolicyMappings = commonlibs.StringSanitize(x509v3e.PolicyMappings)

	return x509v3e
}

/********** 			Некоторые расширения, относящиеся к Cyber-observable Objects STIX 			**********/

//ArchiveFileExtensionSTIX тип "archive-ext", по терминалогии STIX, содержит рассширение архивного файла. В ней задается расширение по умолчанию для захвата свойств,
//  специфичных для архивных файлов
// ContainsRefs - данное свойство определяет файлы содержащиеся в архиве. ДОЛЖНО содержать список типа file или directory (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Comment - определяет комментарий включенный как часть архивного файла
type ArchiveFileExtensionSTIX struct {
	ContainsRefs []*IdentifierTypeSTIX `json:"contains_refs" bson:"contains_refs"`
	Comment      string                `json:"comment" bson:"comment"`
}

//NTFSFileExtensionSTIX  тип "ntfs-ext", по терминалогии STIX, содержит расширение определяющее значения свойств, характерные для файловой системы NTFS
// SID - опреляет безопасный идентификатор, связанный с файлом
// AlternateDataStreams - определяет список альтернативных NTFS потоков данных связанных с файлом
type NTFSFileExtensionSTIX struct {
	SID                  string                         `json:"sid" bson:"sid"`
	AlternateDataStreams []*AlternateDataStreamTypeSTIX `json:"alternate_data_streams" bson:"alternate_data_streams"`
}

//PDFFileExtensionSTIX тип "pdf-ext", по терминалогии STIX, содержит свойства специфичные для файлов в формате PDF
// Version - номер версии, берется из заголовка PDF документа
// IsOptimized - определяетс был ли PDF файл оптимизирован
// DocumentInfoDict - содержит информацию полученную из document information dictionary (DID) PDF документа
// Pdfid0 - определяет первый файловый идентификатор найденный в PDF файле
// Pdfid1 - определяет второй файловый идентификатор найденный в PDF файле
type PDFFileExtensionSTIX struct {
	Version          string                         `json:"version" bson:"version"`
	IsOptimized      bool                           `json:"is_optimized" bson:"is_optimized"`
	DocumentInfoDict map[string]*DictionaryTypeSTIX `json:"document_info_dict" bson:"document_info_dict"`
	Pdfid0           string                         `json:"pdfid0" bson:"pdfid0"`
	Pdfid1           string                         `json:"pdfid1" bson:"pdfid1"`
}

//RasterImageFileExtensionSTIX тип "raster-image-ext", по терминалогии STIX, определяет специфичные расширения для растровых, графических файлов
// ImageHeight - определяет высоту изображения в файле, в пикселях
// ImageWidth - определяет ширину изображения в файле, в пикселях
// BitsPerPixel - определяет сумму бит для каждого цветового туннеля в файлес изображением, а также общее количество пикселей, используемых для выражения глубины цвета изображения.
// ExifTags - задает набор тегов EXIF, найденных в файле изображения, в виде словаря. Каждая пара ключ/значение в словаре представляет имя/значение одного тега EXIF.
//  Соответственно, каждый ключ словаря ДОЛЖЕН быть сохраненной в регистре версией имени тега EXIF, например . Каждое значение словаря ДОЛЖНО быть либо (для типов данных int* EXIF),
//  либо строкой (для всех других типов данных EXIF).
type RasterImageFileExtensionSTIX struct {
	ImageHeight  int                            `json:"image_height" bson:"image_height"`
	ImageWidth   int                            `json:"image_width" bson:"image_width"`
	BitsPerPixel int                            `json:"bits_per_pixel" bson:"bits_per_pixel"`
	ExifTags     map[string]*DictionaryTypeSTIX `json:"exif_tags" bson:"exif_tags"`
}

//HTTPRequestExtensionSTIX тип "http-request-ext", по терминалогии STIX, определяет специфичное расширение HTTP-запроса, задает расширение по умолчанию для захвата свойств
//  сетевого трафика, специфичных для HTTP-запросов. Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть http-request-ext.
// RequestMethod - указывает часть HTTP-метода строки HTTP - запроса в виде строчной строки. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// RequestValue - указывает часть значения (обычно путь к ресурсу) строки HTTP-запроса. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// RequestVersion - указывает часть HTTP-версии строки HTTP - запроса в виде строчной строки.
// RequestHeader - задает все поля заголовка HTTP, которые могут быть найдены в запросе клиента HTTP, в виде словаря.
// MessageBodyLength - указывает длину тела HTTP-сообщения, если оно включено, в байтах.
// MessageBodyDataRef - указывает данные, содержащиеся в теле HTTP-сообщения, если они включены. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
type HTTPRequestExtensionSTIX struct {
	RequestMethod      string                         `json:"request_method" bson:"request_method"`
	RequestValue       string                         `json:"request_value" bson:"request_value"`
	RequestVersion     string                         `json:"request_version" bson:"request_version"`
	RequestHeader      map[string]*DictionaryTypeSTIX `json:"request_header" bson:"request_header"`
	MessageBodyLength  int                            `json:"message_body_length" bson:"message_body_length"`
	MessageBodyDataRef IdentifierTypeSTIX             `json:"message_body_data_ref" bson:"message_body_data_ref"`
}

//ICMPExtensionSTIX тип "icmp-ext", по терминалогии STIX, определяет специфичное расширение по умолчанию для захвата свойств сетевого трафика, специфичных для ICMP. Ключ для этого
//  расширения при использовании в словаре расширений ДОЛЖЕН быть icmp-ext.
// ICMPTypeHex - указывает тип ICMP байт.
// ICMPCodeHex - задает байт кода ICMP.
type ICMPExtensionSTIX struct {
	ICMPTypeHex string `json:"icmp_type_hex" bson:"icmp_type_hex"`
	ICMPCodeHex string `json:"icmp_code_hex" bson:"icmp_code_hex"`
}

//NetworkSocketExtensionSTIX тип "socket-ext", по терминалогии STIX, определяет специфичное расширение по умолчанию для захвата свойств сетевого трафика, связанных с сетевыми сокетами.
//  Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть socket-ext.
// AddressFamily -  указывает семейство адресов (AF_*), для которого настроен сокет. Значения этого свойства ДОЛЖНЫ быть получены из network-socket-address-family-enum. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// IsBlocking - указывает, находится ли сокет в режиме блокировки.
// IsListening - указывает, находится ли сокет в режиме прослушивания.
// Options - указывает любые параметры (например, SO_*), которые могут использоваться сокетом в качестве словаря. Каждый ключ в словаре ДОЛЖЕН быть сохраненной в регистре версией имени
//  параметра, например SO_ACCEPTCONN. Каждое значение ключа в словаре ДОЛЖНО быть значением для соответствующего ключа опций. Каждое значение словаря ДОЛЖНО быть целым числом.
// SocketType - указывает тип сокета. Значения этого свойства ДОЛЖНЫ быть получены из network-socket-type-enum.
// SocketDescriptor - указывается значение дескриптора файла сокета, связанное с сокетом, как неотрицательное целое число.
// SocketHandle - указывает значение дескриптора или индекса, связанное с сокетом.
type NetworkSocketExtensionSTIX struct {
	AddressFamily    EnumTypeSTIX                   `json:"address_family" bson:"address_family"`
	IsBlocking       bool                           `json:"is_blocking" bson:"is_blocking"`
	IsListening      bool                           `json:"is_listening" bson:"is_listening"`
	Options          map[string]*DictionaryTypeSTIX `json:"options" bson:"options"`
	SocketType       EnumTypeSTIX                   `json:"socket_type" bson:"socket_type"`
	SocketDescriptor int                            `json:"socket_descriptor" bson:"socket_descriptor"`
	SocketHandle     int                            `json:"socket_handle" bson:"socket_handle"`
}

//TCPExtensionSTIX тип "tcp-ext", по терминалогии STIX, определяет специфичное расширение задает расширение по умолчанию для захвата свойств сетевого трафика, специфичных для TCP.
//  Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть tcp-ext.
// SrcFlagsHex - указывает исходные флаги TCP как объединение всех флагов TCP, наблюдаемых между началом трафика (как определено свойством start) и концом трафика (как определено свойством end).
// DstFlagsHex - указывает флаги TCP назначения как объединение всех флагов TCP, наблюдаемых между началом трафика (как определено свойством start) и концом трафика (как определено свойством end).
type TCPExtensionSTIX struct {
	SrcFlagsHex string `json:"src_flags_hex" bson:"src_flags_hex"`
	DstFlagsHex string `json:"dst_flags_hex" bson:"dst_flags_hex"`
}

//WindowsPEBinaryFileExtensionSTIX тип "windows-pebinary-ext", по терминалогии STIX, определяет специфичные свойства для Windows portable executable (PE) файлов
// PeType - определяет тип PE binary, в данном поле СЛЕДУЕТ использовать значение из открытого словаря windows-pebinary-type-ov (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// Imphash - определяет специальный 'импортируемый' хеш или imphash подсчитываемый для основанных на PE Binary импортируемых библиотек и функций.
// MachineHex - определяет специфичный, машинный хеш.
// NumberOfSections - определяет количество секций в PE Binary
// TimeDateStamp - время, в формате "2016-05-12T08:17:27.000Z", когда PE Binary был создан.
// PointerToSymbolTableHex - задает смещение файла таблицы символов COFF.
// NumberOfSymbols - задает номер записи в таблице символов PE binary.
// SizeOfOptionalHeader - задает размер заголовка в PE binary.
// CharacteristicsHex - задает флаги, определяющие характеристики файла.
// FileHeaderHashes - определяет любой, подсчитанный для этого файла, хеш.
// OptionalHeader - задает заголовок PE binary файла. Должно использоватся свойство из windows-pe-optional-header-type.
// Sections - задает список методанных касательно секций PE файла.
type WindowsPEBinaryFileExtensionSTIX struct {
	PeType                  OpenVocabTypeSTIX               `json:"pe_type" bson:"pe_type"`
	Imphash                 string                          `json:"imphash" bson:"imphash"`
	MachineHex              string                          `json:"machine_hex" bson:"machine_hex"`
	NumberOfSections        int                             `json:"number_of_sections" bson:"number_of_sections"`
	TimeDateStamp           time.Time                       `json:"time_date_stamp" bson:"time_date_stamp"`
	PointerToSymbolTableHex string                          `json:"pointer_to_symbol_table_hex" bson:"pointer_to_symbol_table_hex"`
	NumberOfSymbols         int                             `json:"number_of_symbols" bson:"number_of_symbols"`
	SizeOfOptionalHeader    int                             `json:"size_of_optional_header" bson:"size_of_optional_header"`
	CharacteristicsHex      string                          `json:"characteristics_hex" bson:"characteristics_hex"`
	FileHeaderHashes        HashesTypeSTIX                  `json:"file_header_hashes" bson:"file_header_hashes"`
	OptionalHeader          WindowsPEOptionalHeaderTypeSTIX `json:"optional_header" bson:"optional_header"`
	Sections                []*WindowsPESectionTypeSTIX     `json:"sections" bson:"sections"`
}

//WindowsProcessExtensionSTIX тип "windows-process-ext", по терминалогии STIX, содержит рассширения расширение по умолчанию для захвата свойств,
//  специфичных для процессов Windows. Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть windows-process-ext.
// ASLREnabled - указывает, включена ли для процесса рандомизация компоновки адресного пространства (ASLR).
// DEPEnabled - указывает, включено ли для процесса Предотвращение выполнения данных (DEP).
// Priority - задает текущий приоритет процесса в Windows. Это значение ДОЛЖНО быть строкой, которая заканчивается на _CLASS.
// OwnerSID - указывает значение идентификатора безопасности (SID) владельца процесса.
// WindowTitle - задает заголовок главного окна процесса.
// StartupInfo - указывает структуру STARTUP_INFO, используемую процессом в качестве словаря.
// IntegrityLevel - указывает уровень целостности Windows или надежность процесса. Значения этого свойства ДОЛЖНЫ быть получены из перечисления windows-integrity-level-enum.
type WindowsProcessExtensionSTIX struct {
	ASLREnabled    bool                           `json:"aslr_enabled" bson:"aslr_enabled"`
	DEPEnabled     bool                           `json:"dep_enabled" bson:"dep_enabled"`
	Priority       string                         `json:"priority" bson:"priority"`
	OwnerSID       string                         `json:"owner_sid" bson:"owner_sid"`
	WindowTitle    string                         `json:"window_title" bson:"window_title"`
	StartupInfo    map[string]*DictionaryTypeSTIX `json:"startup_info" bson:"startup_info"`
	IntegrityLevel EnumTypeSTIX                   `json:"integrity_level" bson:"integrity_level"`
}

//WindowsServiceExtensionSTIX тип "windows-service-ext", по терминалогии STIX, содержит рассширения службы Windows задает расширение по умолчанию для
//  захвата свойств, специфичных для служб Windows. Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть windows-service-ext.
// ServiceName - определяет название сервиса.
// Descriptions - определяет детальное описание сервиса.
// DisplayName - указывает отображаемое имя службы в элементах управления графическим интерфейсом Windows.
// GroupName - указывает имя группы заказа загрузки, членом которой является служба.
// StartType - указывает параметры запуска, определенные для службы. Значения этого свойства ДОЛЖНЫ быть получены из перечисления windows-service-start-type-enum.
// ServiceDllRefs - указывает библиотеки DLL, загруженные службой, в качестве ссылки на один или несколько файловых объектов. Объекты, на которые ссылается
//  это свойство, ДОЛЖНЫ иметь тип file.
// ServiceType - определяет тип сервиса. Значение данного поля ДОЛЖНО иметь тип windows-service-type-enum.
// ServiceStatus - определяет статус сервиса. Значение данного поля ДОЛЖНО иметь тип windows-service-status-enum.
type WindowsServiceExtensionSTIX struct {
	ServiceName    string                `json:"service_name" bson:"service_name"`
	Descriptions   []string              `json:"descriptions" bson:"descriptions"`
	DisplayName    string                `json:"display_name" bson:"display_name"`
	GroupName      string                `json:"group_name" bson:"group_name"`
	StartType      EnumTypeSTIX          `json:"start_type" bson:"start_type"`
	ServiceDllRefs []*IdentifierTypeSTIX `json:"service_dll_refs" bson:"service_dll_refs"`
	ServiceType    EnumTypeSTIX          `json:"service_type" bson:"service_type"`
	ServiceStatus  EnumTypeSTIX          `json:"service_status" bson:"service_status"`
}

//UNIXAccountExtensionSTIX тип "unix-account-ext", по терминалогии STIX, содержит рассширения 'по умолчанию' захваченной дополнительной информации
// предназначенной для аккаунтов UNIX систем.
// GID - содержит первичный групповой ID аккаунта
// Groups - содержит список имен групп которые являются членами аккаунта
// HomeDir - содержит домашную директорию аккаунта
// Shell - содержит командную оболочку аккаунта
type UNIXAccountExtensionSTIX struct {
	GID     int      `json:"gid" bson:"gid"`
	Groups  []string `json:"group" bson:"group"`
	HomeDir string   `json:"home_dir" bson:"home_dir"`
	Shell   string   `json:"shell" bson:"shell"`
}

/********** 			Cyber-observable Objects STIX (ТИПЫ)			**********/

//OptionalCommonPropertiesCyberObservableObjectSTIX содержит опциональные общие свойства для Cyber-observable Objects STIX
// SpecVersion - версия STIX спецификации.
// ObjectMarkingRefs - определяет список ID ссылающиеся на объект "marking-definition", по терминалогии STIX, в котором содержатся значения применяющиеся к этому объекту
// GranularMarkings - определяет список "гранулярных меток" (granular_markings) относящихся к этому объекту
// Defanged - определяет были ли определены данные содержащиеся в объекте
// Extensions - может содержать дополнительную информацию, относящуюся к объекту
type OptionalCommonPropertiesCyberObservableObjectSTIX struct {
	SpecVersion       string                         `json:"spec_version" bson:"spec_version"`
	ObjectMarkingRefs []*IdentifierTypeSTIX          `json:"object_marking_refs" bson:"object_marking_refs"`
	GranularMarkings  GranularMarkingsTypeSTIX       `json:"granular_markings" bson:"granular_markings"`
	Defanged          bool                           `json:"defanged" bson:"defanged"`
	Extensions        map[string]*DictionaryTypeSTIX `json:"extensions" bson:"extensions"`
}

//ArtifactCyberObservableObjectSTIX объект "Artifact", по терминалогии STIX, позволяет захватывать массив байтов (8 бит) в виде строки в кодировке base64
//  или связывать его с полезной нагрузкой, подобной файлу. Обязательно должен быть заполнено одно из полей PayloadBin или URL
// MimeType - по возможности это значение ДОЛЖНО быть одним из значений, определенных в реестре типов носителей IANA. В универсальном каталоге
//  всех существующих типов файлов.
// PayloadBin - бинарные данные в base64
// URL - унифицированный указатель ресурса (URL)
// Hashes - словарь хешей для URL или PayloadBin
// EncryptionAlgorithm - тип алгоритма шифрования для бинарных данных
// DecryptionKey - определяет ключ для дешифрования зашифрованных данных
type ArtifactCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	MimeType            string         `json:"mime_type" bson:"mime_type"`
	PayloadBin          string         `json:"payload_bin" bson:"payload_bin"`
	URL                 string         `json:"url" bson:"url"`
	Hashes              HashesTypeSTIX `json:"hashes" bson:"hashes"`
	EncryptionAlgorithm EnumTypeSTIX   `json:"encryption_algorithm" bson:"encryption_algorithm"`
	DecryptionKey       string         `json:"decryption_key" bson:"decryption_key"`
}

//AutonomousSystemCyberObservableObjectSTIX объект "Autonomous System", по терминалогии STIX, содержит параметры Автономной системы
// Number - содержит номер присвоенный Автономной системе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Name - название Автономной системы
// RIR - содержит название регионального Интернет-реестра (Regional Internet Registry) которым было дано имя Автономной системы
type AutonomousSystemCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Number int    `json:"number" bson:"number" required:"true"`
	Name   string `json:"name" bson:"name"`
	RIR    string `json:"rir" bson:"rir"`
}

//DirectoryCyberObservableObjectSTIX объект "Directory", по терминалогии STIX, содержит свойства, общие для каталога файловой системы
// Path - указывает путь, как было первоначально замечено, к каталогу в файловой системе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PathEnc - указывает наблюдаемую кодировку для пути. Значение ДОЛЖНО быть указано, если путь хранится в кодировке, отличной от Unicode.
// Ctime - время, в формате "2016-05-12T08:17:27.000Z", создания директории
// Mtime - время, в формате "2016-05-12T08:17:27.000Z", модификации или записи в директорию
// Atime - время, в формате "2016-05-12T08:17:27.000Z", последнего обращения к директории
// ContainsRefs - содержит список файловых объектов или директорий содержащихся внутри директории
type DirectoryCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Path         string                `json:"path" bson:"path" required:"true"`
	PathEnc      string                `json:"path_enc" bson:"path_enc"`
	Ctime        time.Time             `json:"ctime" bson:"ctime"`
	Mtime        time.Time             `json:"mtime" bson:"mtime"`
	Atime        time.Time             `json:"atime" bson:"atime"`
	ContainsRefs []*IdentifierTypeSTIX `json:"contains_refs" bson:"contains_refs"`
}

//DomainNameCyberObservableObjectSTIX объект "Domain Name", по терминалогии STIX, содержит сетевое доменное имя
// Value - сетевое доменное имя (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ResolvesToRefs - список ссылок на один или несколько IP-адресов или доменных имен, на которые разрешается доменное имя
type DomainNameCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Value          string                `json:"value" bson:"value" required:"true"`
	ResolvesToRefs []*IdentifierTypeSTIX `json:"resolves_to_refs" bson:"resolves_to_refs"`
}

//EmailAddressCyberObservableObjectSTIX объект "Email Address", по терминалогии STIX, содержит представление единственного email адреса
// Value - email адрес (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// DisplayName - содержит единственное почтовое имя которое видит человек при просмотре письма
// BelongsToRef - учетная запись пользователя, которой принадлежит адрес электронной почты, в качестве ссылки на объект учетной записи пользователя
type EmailAddressCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Value        string             `json:"value" bson:"value"`
	DisplayName  string             `json:"display_name" bson:"display_name"`
	BelongsToRef IdentifierTypeSTIX `json:"belongs_to_ref" bson:"belongs_to_ref"`
}

//EmailMessageCyberObservableObjectSTIX объект "Email Message", по терминалогии STIX, содержит экземпляр email сообщения
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
//  sender_ref, to_ref, cc_ref, bcc_refs и subject), найденные в сообщении электронной почты в виде словаря
// Body - содержит тело сообщения
// BodyMultipart - содержит адает список MIME-части, которые составляют тело email. Это свойство НЕ ДОЛЖНО использоваться, если
//  is_multipart имеет значение false
// RawEmailRef - содержит 'сырое' бинарное содержимое email сообщения
type EmailMessageCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	IsMultipart            bool                           `json:"is_multipart" bson:"is_multipart" required:"true"`
	Date                   time.Time                      `json:"date" bson:"date"`
	ContentType            string                         `json:"content_type" bson:"content_type"`
	FromRef                IdentifierTypeSTIX             `json:"from_ref" bson:"from_ref"`
	SenderRef              IdentifierTypeSTIX             `json:"sender_ref" bson:"sender_ref"`
	ToRefs                 []*IdentifierTypeSTIX          `json:"to_refs" bson:"to_refs"`
	CcRefs                 []*IdentifierTypeSTIX          `json:"cc_refs" bson:"cc_refs"`
	BccRefs                []*IdentifierTypeSTIX          `json:"bcc_refs" bson:"bcc_refs"`
	MessageID              string                         `json:"message_id" bson:"message_id"`
	Subject                string                         `json:"subject" bson:"subject"`
	ReceivedLines          []string                       `json:"received_lines" bson:"received_lines"`
	AdditionalHeaderFields map[string]*DictionaryTypeSTIX `json:"additional_header_fields" bson:"additional_header_fields"`
	Body                   string                         `json:"body" bson:"body"`
	BodyMultipart          []*EmailMIMEPartTypeSTIX       `json:"body_multipart" bson:"body_multipart"`
	RawEmailRef            IdentifierTypeSTIX             `json:"raw_email_ref" bson:"raw_email_ref"`
}

//CommonFileCyberObservableObjectSTIX общий объект "File Object", по терминалогии STIX, содержит объект со свойствами файла
// Extensions - определяет следующие расширения pdf-ext, archive-ext, ntfs-ext, raster-image-ext, windows-pebinary-ext. В дополнении к ним пользователь может создавать
//  свои расширения. При этом ключ словаря должен однозначно идентифицировать тип расширения.
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
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Extensions         map[string]*json.RawMessage `json:"extensions" bson:"extensions"`
	Hashes             HashesTypeSTIX              `json:"hashes" bson:"hashes"`
	Size               uint64                      `json:"size" bson:"size"`
	Name               string                      `json:"name" bson:"name"`
	NameEnc            string                      `json:"name_enc" bson:"name_enc"`
	MagicNumberHex     string                      `json:"magic_number_hex" bson:"magic_number_hex"`
	MimeType           string                      `json:"mime_type" bson:"mime_type"`
	Ctime              time.Time                   `json:"ctime" bson:"ctime"`
	Mtime              time.Time                   `json:"mtime" bson:"mtime"`
	Atime              time.Time                   `json:"atime" bson:"atime"`
	ParentDirectoryRef IdentifierTypeSTIX          `json:"parent_directory_ref" bson:"parent_directory_ref"`
	ContainsRefs       []*IdentifierTypeSTIX       `json:"contains_refs" bson:"contains_refs"`
	ContentRef         IdentifierTypeSTIX          `json:"content_ref" bson:"content_ref"`
}

//FileCyberObservableObjectSTIX объект "File Object", по терминалогии STIX, последекодирования из JSON (основной, рабочий объект)
// Extensions - определяет следующие расширения pdf-ext, archive-ext, ntfs-ext, raster-image-ext, windows-pebinary-ext. В дополнении к ним пользователь может создавать
//  свои расширения. При этом ключ словаря должен однозначно идентифицировать тип расширения.
type FileCyberObservableObjectSTIX struct {
	CommonFileCyberObservableObjectSTIX
	Extensions map[string]interface{} `json:"extensions" bson:"extensions"`
}

//IPv4AddressCyberObservableObjectSTIX объект "IPv4 Address Object", по терминалогии STIX, содержит один или более IPv4 адресов, выраженных с помощью нотации CIDR.
// Value - указывает значения одного или нескольких IPv4-адресов, выраженные с помощью нотации CIDR. Если данный объект IPv4-адреса представляет собой один IPv4-адрес,
//  суффикс CIDR /32 МОЖЕТ быть опущен. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ResolvesToRefs - указывает список ссылок на один или несколько MAC-адресов управления доступом к носителям уровня 2, на которые разрешается IPv6-адрес. Объекты,
//  на которые ссылается этот список, ДОЛЖНЫ иметь тип macaddr.
// BelongsToRefs - указывает список ссылок на одну или несколько автономных систем (AS), к которым принадлежит IPv6-адрес. Объекты, на которые ссылается этот список,
//  ДОЛЖНЫ быть типа autonomous-system.
type IPv4AddressCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Value          string                `json:"value" bson:"value"`
	ResolvesToRefs []*IdentifierTypeSTIX `json:"resolves_to_refs" bson:"resolves_to_refs"`
	BelongsToRefs  []*IdentifierTypeSTIX `json:"belongs_to_refs" bson:"belongs_to_refs"`
}

//IPv6AddressCyberObservableObjectSTIX объект "IPv6 Address Object", по терминалогии STIX, содержит один или более IPv6 адресов, выраженных с помощью нотации CIDR.
// Value - указывает значения одного или нескольких IPv6-адресов, выраженные с помощью нотации CIDR. Если данный объект IPv6-адреса представляет собой один IPv6-адрес,
//  суффикс CIDR /128 МОЖЕТ быть опущен. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ResolvesToRefs - указывает список ссылок на один или несколько MAC-адресов управления доступом к носителям уровня 2, на которые разрешается IPv6-адрес. Объекты,
//  на которые ссылается этот список, ДОЛЖНЫ иметь тип macaddr.
// BelongsToRefs - указывает список ссылок на одну или несколько автономных систем (AS), к которым принадлежит IPv4-адрес. Объекты, на которые ссылается этот список,
//  ДОЛЖНЫ быть типа autonomous-system.
type IPv6AddressCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Value          string                `json:"value" bson:"value"`
	ResolvesToRefs []*IdentifierTypeSTIX `json:"resolves_to_refs" bson:"resolves_to_refs"`
	BelongsToRefs  []*IdentifierTypeSTIX `json:"belongs_to_refs" bson:"belongs_to_refs"`
}

//MACAddressCyberObservableObjectSTIX объект "MAC Address Object", по терминалогии STIX, содержит объект MAC-адрес, представляющий собой один адрес управления доступом к среде (MAC).
// Value - Задает значение одного MAC-адреса. Значение MAC - адреса ДОЛЖНО быть представлено в виде одного строчного MAC-48 address, разделенного двоеточием,
//  который ДОЛЖЕН включать начальные нули для каждого октета. Пример: 00:00:ab:cd:ef:01. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type MACAddressCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Value string `json:"value" bson:"value"`
}

//MutexCyberObservableObjectSTIX объект "Mutex Object", по терминалогии STIX, содержит свойства объекта взаимного исключения (mutex).
// Name - указывает имя объекта мьютекса (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
type MutexCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Name string `json:"name" bson:"name"`
}

//CommonNetworkTrafficCyberObservableObjectSTIX общий объект "Network Traffic Object", по терминалогии STIX, содержит объект Сетевого трафика представляющий собой произвольный сетевой трафик,
//  который исходит из источника и адресуется адресату.
// Extensions - объект Сетевого трафика определяет следующие расширения. В дополнение к ним производители МОГУТ создавать свои собственные. ключи словаря http-request-ext, cp-ext,
//  icmp-ext, socket-ext ДОЛЖНЫ идентифицировать тип расширения по имени. Соответствующие значения словаря ДОЛЖНЫ содержать содержимое экземпляра расширения.
// Start - время, в формате "2016-05-12T08:17:27.000Z", инициирования сетевого трафика, если он известен.
// End - время, в формате "2016-05-12T08:17:27.000Z", окончания сетевого трафика, если он известен.
// IsActive - указывает, продолжается ли сетевой трафик. Если задано свойство end, то это свойство ДОЛЖНО быть false.
// SrcRef - указывает источник сетевого трафика в качестве ссылки на кибернаблюдаемый объект. Объект, на который ссылается ссылка, ДОЛЖЕН быть типа ipv4-addr, ipv6 - addr, mac-addr
//  или domain-name (для случаев, когда IP-адрес для доменного имени неизвестен).
// DstRef - указывает место назначения сетевого трафика в качестве ссылки на кибернаблюдаемый объект. Объект, на который ссылается ссылка, ДОЛЖЕН быть типа ipv4-addr, ipv6 - addr,
//  mac-addr или domain-name (для случаев, когда IP-адрес для доменного имени неизвестен).
// SrcPort - задает исходный порт, используемый в сетевом трафике, в виде целого числа. Значение порта ДОЛЖНО находиться в диапазоне от 0 до 65535.
// DstPort - задает порт назначения, используемый в сетевом трафике, в виде целого числа. Значение порта ДОЛЖНО находиться в диапазоне от 0 до 65535.
// Protocols - указывает протоколы, наблюдаемые в сетевом трафике, а также их соответствующее состояние.
// SrcByteCount - задает число байтов в виде положительного целого числа, отправленных от источника к месту назначения.
// DstByteCount - задает число байтов в виде положительного целого числа, отправленных из пункта назначения в источник.
// SrcPackets - задает количество пакетов в виде положительного целого числа, отправленных от источника к месту назначения.
// DstPackets - задает количество пакетов в виде положительного целого числа, отправленных от пункта назначения к источнику
// IPFix - указывает любые данные Экспорта информации IP-потока [IPFIX] для трафика в виде словаря. Каждая пара ключ/значение в словаре представляет имя/значение одного элемента IPFIX.
//  Соответственно, каждый ключ словаря ДОЛЖЕН быть сохраненной в регистре версией имени элемента IPFIX.
// SrcPayloadRef - указывает байты, отправленные из источника в пункт назначения. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
// DstPayloadRef - указывает байты, отправленные из пункта назначения в источник. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
// EncapsulatesRefs - ссылки на другие объекты, инкапсулированные этим объектом. Объекты, на которые ссылается это свойство, ДОЛЖНЫ иметь тип network-traffic.
// EncapsulatedByRef - ссылки на другой объект сетевого трафика, который инкапсулирует этот объект. Объекты, на которые ссылается это свойство, ДОЛЖНЫ иметь тип network-traffic.
type CommonNetworkTrafficCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Extensions        map[string]*json.RawMessage    `json:"extensions" bson:"extensions"`
	Start             time.Time                      `json:"start" bson:"start"`
	End               time.Time                      `json:"end" bson:"end"`
	IsActive          bool                           `json:"is_active" bson:"is_active"`
	SrcRef            IdentifierTypeSTIX             `json:"src_ref" bson:"src_ref"`
	DstRef            IdentifierTypeSTIX             `json:"dst_ref" bson:"dst_ref"`
	SrcPort           int                            `json:"src_port" bson:"src_port"`
	DstPort           int                            `json:"dst_port" bson:"dst_port"`
	Protocols         []string                       `json:"protocols" bson:"protocols"`
	SrcByteCount      uint64                         `json:"src_byte_count" bson:"src_byte_count"`
	DstByteCount      uint64                         `json:"dst_byte_count" bson:"dst_byte_count"`
	SrcPackets        int                            `json:"src_packets" bson:"src_packets"`
	DstPackets        int                            `json:"dst_packets" bson:"dst_packets"`
	IPFix             map[string]*DictionaryTypeSTIX `json:"ipfix" bson:"ipfix"`
	SrcPayloadRef     IdentifierTypeSTIX             `json:"src_payload_ref" bson:"src_payload_ref"`
	DstPayloadRef     IdentifierTypeSTIX             `json:"dst_payload_ref" bson:"dst_payload_ref"`
	EncapsulatesRefs  []*IdentifierTypeSTIX          `json:"encapsulates_refs" bson:"encapsulates_refs"`
	EncapsulatedByRef IdentifierTypeSTIX             `json:"encapsulated_by_ref" bson:"encapsulated_by_ref"`
}

//NetworkTrafficCyberObservableObjectSTIX объект "Network Traffic Object", по терминалогии STIX, содержит объект Сетевого трафика представляющий собой произвольный сетевой трафик,
//  который исходит из источника и адресуется адресату.
// Extensions - объект Сетевого трафика определяет следующие расширения. В дополнение к ним производители МОГУТ создавать свои собственные. ключи словаря http-request-ext, cp-ext,
type NetworkTrafficCyberObservableObjectSTIX struct {
	CommonNetworkTrafficCyberObservableObjectSTIX
	Extensions map[string]interface{} `json:"extensions" bson:"extensions"`
}

//CommonProcessCyberObservableObjectSTIX общий объект "Process Object", по терминологии STIX, содержит общие свойства экземпляра компьютерной программы,
//  выполняемой в операционной системе. Объект процесса ДОЛЖЕН содержать хотя бы одно свойство (отличное от типа) этого объекта (или одного из его расширений).
// Extensions - определяет расширения windows-process-exit или windows-service-ext. В дополнение к ним производители МОГУТ создавать свои собственные. ключи словаря windows-process-exit,
//  windows-service-ext ДОЛЖНЫ идентифицировать тип расширения по имени. Соответствующие значения словаря ДОЛЖНЫ содержать содержимое экземпляра расширения.
// IsHidden - определяет является ли процесс скрытым.
// PID - униальный идентификатор процесса.
// CreatedTime - время, в формате "2016-05-12T08:17:27.000Z", создания процесса.
// Cwd - текущая рабочая директория процесса.
// CommandLine - поределяет полный перечень команд используемых для запуска процесса, включая имя процесса и аргументы.
// EnvironmentVariables - определяет список переменных окружения, в виде словаря, ассоциируемых с приложением.
// OpenedConnectionRefs - определяет список открытых, процессом, сетевых соединений ка одну или более ссылку на объект типа network-traffic.
// CreatorUserRef - определяет что за пользователь создал объект, ссылка ДОЛЖНА ссылатся на объект типа user-account.
// ImageRef - указывает исполняемый двоичный файл, который был выполнен как образ процесса, как ссылка на файловый объект. Объект, на который ссылается
//  это свойство, ДОЛЖЕН иметь тип file.
// ParentRef - указывает другой процесс, который породил (т. е. является родителем) этот процесс, как ссылку на объект процесса. Объект, на который
//  ссылается это свойство, ДОЛЖЕН иметь тип process.
// ChildRefs - указывает другие процессы, которые были порождены (т. е. дочерние) этим процессом, в качестве ссылки на один или несколько других
//  объектов процесса. Объекты, на которые ссылается этот список, ДОЛЖНЫ иметь тип process.
type CommonProcessCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Extensions           map[string]*json.RawMessage    `json:"extensions" bson:"extensions"`
	IsHidden             bool                           `json:"is_hidden" bson:"is_hidden"`
	PID                  int                            `json:"pid" bson:"pid"`
	CreatedTime          time.Time                      `json:"created_time" bson:"created_time"`
	Cwd                  string                         `json:"cwd" bson:"cwd"`
	CommandLine          string                         `json:"command_line" bson:"command_line"`
	EnvironmentVariables map[string]*DictionaryTypeSTIX `json:"environment_variables" bson:"environment_variables"`
	OpenedConnectionRefs []*IdentifierTypeSTIX          `json:"opened_connection_refs" bson:"opened_connection_refs"`
	CreatorUserRef       IdentifierTypeSTIX             `json:"creator_user_ref" bson:"creator_user_ref"`
	ImageRef             IdentifierTypeSTIX             `json:"image_ref" bson:"image_ref"`
	ParentRef            IdentifierTypeSTIX             `json:"parent_ref" bson:"parent_ref"`
	ChildRefs            []*IdentifierTypeSTIX          `json:"child_refs" bson:"child_refs"`
}

//ProcessCyberObservableObjectSTIX объект "Process Object", по терминологии STIX, содержит общие свойства экземпляра компьютерной программы,
//  выполняемой в операционной системе. Объект процесса ДОЛЖЕН содержать хотя бы одно свойство (отличное от типа) этого объекта (или одного из его расширений).
// Extensions - определяет расширения windows-process-exit или windows-service-ext. В дополнение к ним производители МОГУТ создавать свои собственные. ключи словаря windows-process-exit,
//  windows-service-ext ДОЛЖНЫ идентифицировать тип расширения по имени. Соответствующие значения словаря ДОЛЖНЫ содержать содержимое экземпляра расширения.
type ProcessCyberObservableObjectSTIX struct {
	CommonProcessCyberObservableObjectSTIX
	Extensions map[string]interface{} `json:"extensions" bson:"extensions"`
}

//SoftwareCyberObservableObjectSTIX объект "Software Object", по терминологии STIX, содержит свойства, связанные с программным обеспечением, включая программные продукты.
// Name - назвыание программного обеспечения
// CPE - содержит запись Common Platform Enumeration (CPE) для программного обеспечения, если она доступна. Значение этого свойства должно быть значением
// CPE v2.3 из официального словаря NVD CPE [NVD]
// SwID - содержит запись Тегов Software Identification ID (SWID) [SWID] для программного обеспечения, если таковая имеется. SwID помеченный tagId, является
//  глобально уникальным идентификатором и ДОЛЖЕН использоваться как полномочие для идентификации помеченного продукта
// Languages -содержит языки, поддерживаемые программным обеспечением. Значение каждого елемента списка ДОЛЖНО быть кодом языка ISO 639-2
// Vendor - содержит название производителя программного обеспечения
// Version - содержит версию ПО
type SoftwareCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Name      string   `json:"name" bson:"name"`
	CPE       string   `json:"cpe" bson:"cpe"`
	SwID      string   `json:"swid" bson:"swid"`
	Languages []string `json:"languages" bson:"languages"`
	Vendor    string   `json:"vendor" bson:"vendor"`
	Version   string   `json:"version" bson:"version"`
}

//URLCyberObservableObjectSTIX объект "URL Object", по терминологии STIX, содержит унифицированный указатель информационного ресурса (URL).
// Value - содержит унифицированный указатель информационного ресурса (URL).
type URLCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Value string `json:"value" bson:"value"`
}

//UserAccountCyberObservableObjectSTIX объект "User Account Object", по терминалогии STIX, содержит экземпляр любого типа учетной записи пользователя, включая,
// учетные записи операционной системы, устройства, службы обмена сообщениями и платформы социальных сетей и других прочих учетных записей
// Поскольку все свойства этого объекта являются необязательными, по крайней мере одно из свойств, определенных ниже, ДОЛЖНО быть инициализировано
// при использовании этого объекта
// Extensions - содержит словарь расширяющий тип "User Account Object" одно из расширений "unix-account-ext", реализуется описанным ниже типом, UNIXAccountExtensionSTIX
//  кроме этого производитель может созавать свои собственные типы расширений
//  Ключи данного словаря идентифицируют тип расширения по имени, значения являются содержимым экземпляра расширения
// UserID - содержит идентификатор учетной записи. Формат идентификатора зависит от системы в которой находится данная учетная запись пользователя,
//  и может быть числовым идентификатором, идентификатором GUID, именем учетной записи, адресом электронной почты и т.д. Свойство  UserId должно
//  быть заполнено любым значанием, являющимся уникальным идентификатором системы, членом которой является учетная запись. Например, в системах UNIX он
//  будет заполнено значением UID
// Credential - содержит учетные данные пользователя в открытом виде. Предназначено только для закрытого применения при изучении метаданных вредоносных программ
//  при их исследовании (например, жестко закодированный пароль администратора домена, который вредоносная программа пытается использовать реализации тактики для
//	бокового (латерального) перемещения) и не должно применяться для совместного пользования PII
// AccountLogin - содержит логин пользователя. Используется в тех случаях,когда свойство UserId указывает другие данные, чем то, что пользователь вводит
//  при входе в систему
// AccountType - содержит одно, из заранее определенных (предложенных) значений. Является типом аккаунта. Значения этого свойства берутся из множества
//  закрепленного в открытом словаре, account-type-ov
// DisplayName - содержит отображаемое имя учетной записи, которое будет отображаться в пользовательских интерфейсах. В Unix, это равносильно полю gecos
//  (gecos это поле учётной записи пользователя в файле /etc/passwd )
// IsServiceAccount - содержит индикатор, сигнализирующий что, учетная запись связана с сетевой службой или системным процессом (демоном), а не с конкретным человеком. (системный пользователь)
// IsPrivileged - содержит индикатор, сигнализирующий что, учетная запись имеет повышенные привилегии (например, в случае root в Unix или учетной записи администратора
//  Windows)
// CanEscalatePrivs  - содержит индикатор, сигнализирующий что, учетная запись имеет возможность повышать привилегии (например, в случае sudo в Unix или учетной
//  записи администратора домена Windows)
// IsDisabled  - содержит индикатор, сигнализирующий что, учетная запись отключена
// AccountCreated - время, в формате "2016-05-12T08:17:27.000Z", создания аккаунта
// AccountExpires - время, в формате "2016-05-12T08:17:27.000Z", истечения срока действия учетной записи.
// CredentialLastChanged - время, в формате "2016-05-12T08:17:27.000Z", когда учетные данные учетной записи были изменены в последний раз.
// AccountFirstLogin - время, в формате "2016-05-12T08:17:27.000Z", первого доступа к учетной записи
// AccountLastLogin - время, в формате "2016-05-12T08:17:27.000Z", когда к учетной записи был последний доступ.
type UserAccountCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Extensions            map[string]UNIXAccountExtensionSTIX `json:"extensions" bson:"extensions"`
	UserID                string                              `json:"user_id" bson:"user_id"`
	Credential            string                              `json:"credential" bson:"credential"`
	AccountLogin          string                              `json:"account_login" bson:"account_login"`
	AccountType           OpenVocabTypeSTIX                   `json:"account_type" bson:"account_type"`
	DisplayName           string                              `json:"display_name" bson:"display_name"`
	IsServiceAccount      bool                                `json:"is_service_account" bson:"is_service_account"`
	IsPrivileged          bool                                `json:"is_privileged" bson:"is_privileged"`
	CanEscalatePrivs      bool                                `json:"can_escalate_privs" bson:"can_escalate_privs"`
	IsDisabled            bool                                `json:"is_disabled" bson:"is_disabled"`
	AccountCreated        time.Time                           `json:"account_created" bson:"account_created"`
	AccountExpires        time.Time                           `json:"account_expires" bson:"account_expires"`
	CredentialLastChanged time.Time                           `json:"credential_last_changed" bson:"credential_last_changed"`
	AccountFirstLogin     time.Time                           `json:"account_first_login" bson:"account_first_login"`
	AccountLastLogin      time.Time                           `json:"account_last_login" bson:"account_last_login"`
}

//WindowsRegistryKeyCyberObservableObjectSTIX объект "Windows Registry Key Object", по терминалогии STIX. Содержит описание значений полей раздела реестра Windows.
//  Поскольку все свойства этого объекта являются необязательными, по крайней мере одно из свойств,определенных ниже, должно быть инициализировано при
//  использовании этого объекта.
// Key - содержит полный путь к разделу реестра. Значение ключа,должно быть сохранено в регистре. В название ключа все сокращения должны быть раскрыты.
//  Например, вместо HKLM следует использовать HKEY_LOCAL_MACHINE.
// Values - содержит значения, найденные в разделе реестра.
// ModifiedTime - время, в формате "2016-05-12T08:17:27.000Z", последнего изменения раздела реестра.
// CreatorUserRef - содержит ссылку на учетную запись пользователя, из под которой создан раздел реестра. Объект, на который ссылается это свойство, должен иметь тип user-account.
// NumberOfSubkeys - Указывает количество подразделов, содержащихся в разделе реестра.
type WindowsRegistryKeyCyberObservableObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	Key             string                         `json:"key" bson:"key"`
	Values          []WindowsRegistryValueTypeSTIX `json:"values" bson:"values"`
	ModifiedTime    time.Time                      `json:"modified_time" bson:"modified_time"`
	CreatorUserRef  IdentifierTypeSTIX             `json:"creator_user_ref" bson:"creator_user_ref"`
	NumberOfSubkeys int                            `json:"number_of_subkeys" bson:"number_of_subkeys"`
}

//X509CertificateCyberObservableObjectSTIX объект "X.509 Certificate Object", по терминологии STIX, представлет свойства сертификата X.509, определенные в рекомендациях
//  ITU X.509 [X.509]. X.509  Certificate объект должен содержать по крайней мере одно cвойство специфичное для этого объекта (помимо type).
// IsSelfSigned - содержит индикатор, является ли сертификат самоподписным, то есть подписан ли он тем же субъектом, личность которого он удостоверяет.
// Hashes - содержит любые хэши, которые были вычислены для всего содержимого сертификата. Является типом данных словар, значения ключей которого должны
//  быть из открытого словаря hash-algorithm-ov.
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
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesCyberObservableObjectSTIX
	IsSelfSigned              bool                     `json:"is_self_signed" bson:"is_self_signed"`
	Hashes                    HashesTypeSTIX           `json:"hashes" bson:"hashes"`
	Version                   string                   `json:"version" bson:"version"`
	SerialNumber              string                   `json:"serial_number" bson:"serial_number"`
	SignatureAlgorithm        string                   `json:"signature_algorithm" bson:"signature_algorithm"`
	Issuer                    string                   `json:"issuer" bson:"issuer"`
	ValidityNotBefore         time.Time                `json:"validity_not_before" bson:"validity_not_before"`
	ValidityNotAfter          time.Time                `json:"validity_not_after" bson:"validity_not_after"`
	Subject                   string                   `json:"subject" bson:"subject"`
	SubjectPublicKeyAlgorithm string                   `json:"subject_public_key_algorithm" bson:"subject_public_key_algorithm"`
	SubjectPublicKeyModulus   string                   `json:"subject_public_key_modulus" bson:"subject_public_key_modulus"`
	SubjectPublicKeyExponent  int                      `json:"subject_public_key_exponent" bson:"subject_public_key_exponent"`
	X509V3Extensions          X509V3ExtensionsTypeSTIX `json:"x509_v3_extensions" bson:"x509_v3_extensions"`
}
