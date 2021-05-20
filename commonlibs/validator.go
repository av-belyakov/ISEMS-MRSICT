package commonlibs

import (
	"regexp"
)

//StringSanitize преобразовавает некоторые символы, такие как '$', '"', '<', '\n' и т.д., в HTML код
func StringSanitize(strIn string) string {
	specialCharacters := [][2]string{
		{`\$`, "&#36;"},
		{`\"`, "&quot;"},
		{`'`, "&apos;"},
		{`<`, "&lt;"},
		{`>`, "&gt;"},
		{`\n`, "&#010;"},
		{`\\n`, "&#010;"},
		{`\t`, "&#009;"},
		{`\\t`, "&#009;"},
		{`\r`, "&#013;"},
		{`\\r`, "&#013;"},
	}

	for ch := range specialCharacters {
		strIn = regexp.MustCompile(specialCharacters[ch][0]).ReplaceAllString(strIn, specialCharacters[ch][1])
	}

	return strIn
}

const (
	ipv4                     string = `^((25[0-5]|2[0-4]\d|[01]?\d\d?)[.]){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`
	computerNetAddrIPv4Range string = `^((25[0-5]|2[0-4]\d|[01]?\d\d?)[.]){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)/[0-9]{1,2}$`
	tcpDumpFileName          string = `^(\w|_)+\.(tdp|pcap)$`
	uuid3                    string = `^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$`
	uuid4                    string = `^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
	uuid5                    string = `^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
	uuid                     string = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`
	alpha                    string = `^[a-zA-Z]+$`
	alphanumeric             string = `^[a-zA-Z0-9]+$`
	literaryline             string = `^[a-zA-Z0-9.,-_()!]+$`
	numeric                  string = `^[0-9]+$`
	integer                  string = `^(?:[-+]?(?:0|[1-9][0-9]*))$`
	float                    string = `^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$`
	hexadecimal              string = `^[0-9a-fA-F]+$`
	hexcolor                 string = `^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`
	ascii                    string = `^[\x00-\x7F]+$`
	multibyte                string = `[^\x00-\x7F]`
	fullWidth                string = `[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]`
	halfWidth                string = `[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]`
	base64                   string = `^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`
	printableASCII           string = `^[\x20-\x7E]+$`
	dataURI                  string = `^data:.+\\/(.+);base64$`
	magnetURI                string = `^magnet:\\?xt=urn:[a-zA-Z0-9]+:[a-zA-Z0-9]{32,40}&dn=.+&tr=.+$`
	latitude                 string = `^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$`
	longitude                string = `^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$`
	dnsName                  string = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	urlSchema                string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	urlUsername              string = `(\S+(:\S*)?@)`
	urlPath                  string = `((\/|\?|#)[^\s]*)`
	urlPort                  string = `(:(\d{1,5}))`
	urlIP                    string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3]|24\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-5]))`
	urlSubdomain             string = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	url                      string = `^` + urlSchema + `?` + urlUsername + `?` + `((` + urlIP + `|(\[` + ipv4 + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + urlSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + urlPort + `?` + urlPath + `?$`
	ssn                      string = `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
	winPath                  string = `^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	unixPath                 string = `^(/[^/\x00]*)+/?$`
)

//IsIPAddress проверяет значение на соответствие его ip адресу сети Интернет
func IsIPv4Address(value string) bool {
	patterCheckFileName := regexp.MustCompile(ipv4)

	return patterCheckFileName.MatchString(value)
}

//IsComputerNetAddrRange проверяет значение на соответствие его диапазону адресов компьютерной сети. Например, 192.168.0.1/24
func IsComputerNetAddrIPv4Range(value string) bool {
	patterCheckFileName := regexp.MustCompile(computerNetAddrIPv4Range)

	return patterCheckFileName.MatchString(value)
}

//IsTCPDumpFileName проверяет значение на соответствие имени файла, содержащего данные в формате tcpdump
func IsTCPDumpFileName(value string) bool {
	patterCheckFileName := regexp.MustCompile(tcpDumpFileName)

	return patterCheckFileName.MatchString(value)
}

//IsUUID3 проверяет значение на соответствие его UUID версии 3
func IsUUID3(value string) bool {
	patterCheckFileName := regexp.MustCompile(uuid3)

	return patterCheckFileName.MatchString(value)
}

//IsUUID4 проверяет значение на соответствие его UUID версии 4
func IsUUID4(value string) bool {
	patterCheckFileName := regexp.MustCompile(uuid4)

	return patterCheckFileName.MatchString(value)
}

//IsUUID5 проверяет значение на соответствие его UUID версии 5
func IsUUID5(value string) bool {
	patterCheckFileName := regexp.MustCompile(uuid5)

	return patterCheckFileName.MatchString(value)
}

//IsUUID проверяет значение на соответствие его UUID
func IsUUID(value string) bool {
	patterCheckFileName := regexp.MustCompile(uuid)

	return patterCheckFileName.MatchString(value)
}

//IsAlpha проверяет, содержит ли значение только буквенные символы
func IsAlpha(value string) bool {
	patterCheckFileName := regexp.MustCompile(alpha)

	return patterCheckFileName.MatchString(value)
}

//IsAlphaNumeric проверяет, содержит ли значение только буквенные или цифровые символы
func IsAlphaNumeric(value string) bool {
	patterCheckFileName := regexp.MustCompile(alphanumeric)

	return patterCheckFileName.MatchString(value)
}

//IsAlphaNumeric проверяет, содержит ли значение только буквенные или цифровые символы
func IsLiteraryLine(value string) bool {
	patterCheckFileName := regexp.MustCompile(literaryline)

	return patterCheckFileName.MatchString(value)
}

//IsInteger проверяет, содержит ли значение отрицательное или положительное, целое число
func IsInteger(value string) bool {
	patterCheckFileName := regexp.MustCompile(integer)

	return patterCheckFileName.MatchString(value)
}

//IsFloat проверяет, содержит ли значение отрицательное или положительное, число с плавающей точкой
func IsFloat(value string) bool {
	patterCheckFileName := regexp.MustCompile(float)

	return patterCheckFileName.MatchString(value)
}

//IsHexadecimal проверяет, содержит ли значение строку в шестнадцатеричном формате
func IsHexadecimal(value string) bool {
	patterCheckFileName := regexp.MustCompile(hexadecimal)

	return patterCheckFileName.MatchString(value)
}

//IsHexcolor проверяет, содержит ли значение строку, которая соответствует цвету, закодированному в шестнадцатеричном формате
func IsHexcolor(value string) bool {
	patterCheckFileName := regexp.MustCompile(hexcolor)

	return patterCheckFileName.MatchString(value)
}

//IsASCII проверяет значение на соответствие формату ASCII
func IsASCII(value string) bool {
	patterCheckFileName := regexp.MustCompile(ascii)

	return patterCheckFileName.MatchString(value)
}

//IsBase64 проверяет значение на соответствие формату Base64
func IsBase64(value string) bool {
	patterCheckFileName := regexp.MustCompile(base64)

	return patterCheckFileName.MatchString(value)
}

//IsPrintableASCII проверяет значение на соответствие формату PrintableASCII
func IsPrintableASCII(value string) bool {
	patterCheckFileName := regexp.MustCompile(printableASCII)

	return patterCheckFileName.MatchString(value)
}

//IsDataURI проверяет значение на соответствие URL содержащего данные
func IsDataURI(value string) bool {
	patterCheckFileName := regexp.MustCompile(dataURI)

	return patterCheckFileName.MatchString(value)
}

//IsMagnetURI проверяет значение на наличие в URL ссылки типа magnet
func IsMagnetURI(value string) bool {
	patterCheckFileName := regexp.MustCompile(magnetURI)

	return patterCheckFileName.MatchString(value)
}

//IsLatitude проверяет значение на соответствие географической широте
func IsLatitude(value string) bool {
	patterCheckFileName := regexp.MustCompile(latitude)

	return patterCheckFileName.MatchString(value)
}

//IsLongitude проверяет значение на соответствие географической долготе
func IsLongitude(value string) bool {
	patterCheckFileName := regexp.MustCompile(longitude)

	return patterCheckFileName.MatchString(value)
}

//IsDNSName проверяет значение на соответствие доменному имени
func IsDNSName(value string) bool {
	patterCheckFileName := regexp.MustCompile(dnsName)

	return patterCheckFileName.MatchString(value)
}

//IsURL проверяет строку на её соответствие URL
func IsURL(value string) bool {
	patterCheckFileName := regexp.MustCompile(url)

	return patterCheckFileName.MatchString(value)
}

//IsURLSchema проверяет строку на её соответствие схеме URL
func IsURLSchema(value string) bool {
	patterCheckFileName := regexp.MustCompile(urlSchema)

	return patterCheckFileName.MatchString(value)
}

//IsURLUsername проверяет, содержит ли сторка URL имя пользователя
func IsURLUsername(value string) bool {
	patterCheckFileName := regexp.MustCompile(urlUsername)

	return patterCheckFileName.MatchString(value)
}

//IsURLPath проверяет, содержит ли сторка URL путь
func IsURLPath(value string) bool {
	patterCheckFileName := regexp.MustCompile(urlPath)

	return patterCheckFileName.MatchString(value)
}

//IsURLPort проверяет, содержит ли сторка URL сетевой порт
func IsURLPort(value string) bool {
	patterCheckFileName := regexp.MustCompile(urlPort)

	return patterCheckFileName.MatchString(value)
}

//IsURLIP проверяет, содержит ли сторка URL ip адрес сети Интернет
func IsURLIP(value string) bool {
	patterCheckFileName := regexp.MustCompile(urlIP)

	return patterCheckFileName.MatchString(value)
}

//IsURLSubdomain проверяет, содержит ли сторка URL поддомены
func IsURLSubdomain(value string) bool {
	patterCheckFileName := regexp.MustCompile(urlSubdomain)

	return patterCheckFileName.MatchString(value)
}

/*// проверяет значение на соответствие
func (value string) bool {
	patterCheckFileName := regexp.MustCompile()

	return patterCheckFileName.MatchString(value)
}*/
