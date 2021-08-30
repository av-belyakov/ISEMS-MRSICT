package datamodels

import (
	"ISEMS-MRSICT/commonlibs"
	"fmt"
	"reflect"
	"regexp"
)

//Sanitize выполняет замену некоторых специальных символов на их HTML код в полях типа VocabularyElement
func (ve *VocElement) Sanitize() {
	ve.ShortDescription = commonlibs.StringSanitize(ve.ShortDescription)
	ve.FullDescription = commonlibs.StringSanitize(ve.FullDescription)
	ve.Name = commonlibs.StringSanitize(ve.Name)
}

//Sanitize выполняет замену некоторых специальных символов на их HTML код в значениях полей в типе VocElements
func (ve *VocElements) Sanitize() {
	for _, elem := range *ve {
		elem.Sanitize()
	}
}

//Sanitize выполняет замену некоторых специальных символов на их HTML код в полях типа Vocabulary
func (voc *Vocabulary) Sanitize() {
	if voc.Elements != nil {
		voc.Elements.Sanitize()
	}
}

//Sanitize выполняет замену некоторых специальных символов на их HTML код в полях типа RBookReqParameter
func (rbp *RBookReqParameter) Sanitize() {
	rbp.OP = commonlibs.StringSanitize(rbp.OP)
	rbp.Vocabulary.Sanitize()
}

//Sanitize выполняет замену некоторых специальных символов на их HTML код в полях типа  RBookReqParameters
func (rbps *RBookReqParameters) Sanitize() {
	for _, elem := range *rbps {
		elem.Sanitize()
	}
}

// Contains - проверка того что значение "op" содержится в списке команд для ReferenceBook
func (commandsList *VocOperations) Contains(op string) bool {
	for _, elem := range *commandsList {
		if op == elem {
			return true
		}
	}
	return false
}

//IsValid - Метод валидатор полей API-запроса к объекту справочной информации ReferersBook
func (rbp *RBookReqParameter) IsValid() (bool, error) {
	var (
		ok  bool = true
		err error
	)
	val := reflect.ValueOf(rbp)
	errMSG := fmt.Sprintf("The next fields of type %s:\n", val.Type().Name())

	//проверка значения поля OP - операции над объектом
	ok = CommandSet.Contains(rbp.OP)
	if !ok {
		errMSG = fmt.Sprintf("%s %s contain not valid value %s\n", errMSG, val.FieldByName("OP"), rbp.OP)
	}
	//Проверка обязательного поля Name
	if rbp.Name == "" {
		ok = false
		errMSG = fmt.Sprintf("%s %s contain not valid value %s\n", errMSG, val.FieldByName("Name"), rbp.Name)
	}
	if !ok {
		err = fmt.Errorf(errMSG)
	}
	return ok, err
}

//IsValid - Метод валидатор среза запросов от API к объектам стправочной информации  - тип ReferersBook
func (rbps *RBookReqParameters) IsValid() (bool, error) {
	var (
		ok     bool   = false
		errMSG string = ""
		err    error
	)
	if len(*rbps) != 0 {
		for _, elem := range *rbps {
			if ok, err = elem.IsValid(); !ok {
				errMSG = fmt.Sprintf("%s%s", errMSG, err.Error())
				err = fmt.Errorf(errMSG)
			}
		}
	} else {
		err = fmt.Errorf(fmt.Sprintf("RB Request is Empty"))
	}
	return ok, err
}

//IsValid- метод валидатор полей полей словаря - тип Vocabularys
func (vocs Vocabularys) IsValid() (bool, error) {
	var (
		ok     bool = true
		err    error
		errMSG string
	)

	for i, v := range vocs {
		if _, err = v.IsValid(); err != nil {
			ok = false
			errMSG = fmt.Sprintf("%s\n[%d].%s", errMSG, i, err.Error())
		}
	}
	if errMSG != "" {
		err = fmt.Errorf(errMSG)
	}

	return ok, err
}

//IsValid- метод валидатор полей полей словаря - тип Vocabulary
func (voc *Vocabulary) IsValid() (bool, error) {
	var (
		ok     bool = true
		err    error
		errMSG string
	)

	//валидация содержимого поля Name
	if voc.Name == "" || !(regexp.MustCompile(`-(enum|ov)$`).MatchString(voc.Name)) {
		errMSG = fmt.Sprintf("Name contain not valid value %s", voc.Name)
		ok = false
	}

	if _, err = voc.Elements.IsValid(); err != nil {
		ok = false
		errMSG = fmt.Sprintf("%s\nElements:\n  %s", errMSG, err.Error())
	}

	if errMSG != "" {
		err = fmt.Errorf(errMSG)
	}

	return ok, err
}

//IsValid - метод валидатор елементов списка - тип VocElements
func (l VocElements) IsValid() (bool, error) {
	var (
		ok     bool = true
		err    error
		errMSG string
	)
	for i, v := range l {
		if _, err = v.IsValid(); err != nil {
			ok = false
			errMSG = fmt.Sprintf("%s\n[%d].%s", errMSG, i, err.Error())
		}
	}
	if errMSG != "" {
		err = fmt.Errorf(errMSG)
	}
	return ok, err
}

//IsValid - метод валидатор полей элемента словаря - тип VocElement
func (vocEl *VocElement) IsValid() (bool, error) {
	var (
		ok     bool = true
		err    error
		errMSG string
	)

	//валидация содержимого поля Name
	if vocEl.Name == "" {
		errMSG = fmt.Sprintf(" Name contain not valid value %s", vocEl.Name)
		ok = false
		err = fmt.Errorf(errMSG)
	}
	return ok, err
}

//GetListID - метод типа Vocabularys, возвращает список имен объектов ввиде среза строк
func (l Vocabularys) GetListID() []string {
	list := make([]string, 0, len(l))
	for _, v := range l {
		list = append(list, v.Name)
	}
	return list
}

//GetListIDtoStr - метод типа Vocabularys, возвращает список имен объектов ввиде строки
func (l Vocabularys) GetListIDtoStr() string {
	var strID string
	for _, v := range l {
		strID = fmt.Sprintf("%s %s", strID, v.Name)
	}
	return strID
}
