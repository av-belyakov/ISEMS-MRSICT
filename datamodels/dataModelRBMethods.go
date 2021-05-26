package datamodels

import (
	"ISEMS-MRSICT/commonlibs"
	"fmt"
	"reflect"
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

//IsValid - Метод валидатор API-запроса к объекту стправичной информации ReferersBook
func (rbp *RBookReqParameter) IsValid() (bool, error) {
	var (
		ok  bool = false
		err error
	)
	for _, elem := range Commands {
		if elem == rbp.OP {
			ok = true
			break
		}
	}
	if !ok {
		val := reflect.ValueOf(rbp)
		errMSG := fmt.Sprintf("The field %s for type %s contains is not valid value %s\n", val.FieldByName("OP"), val.Type().Name(), rbp.OP)
		err = fmt.Errorf(errMSG)
	}
	return ok, err
}

//IsValid - Метод валидатор среза запросов от API к объектам стправочной информации ReferersBook
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
