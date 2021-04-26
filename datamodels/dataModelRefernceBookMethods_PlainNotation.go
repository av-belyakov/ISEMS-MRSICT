package datamodels

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//Методы для NestedSliceType реализующие интерфейс NestedWorker
func (nsl NestedSliceType) GetValue() interface{} {
	return nsl
}

//Методы для NestedMapType реализующие интерфейс NestedWorker
func (nm NestedMapType) GetValue() interface{} {
	return nm
}

func (nm NestedMapType) AddValue(v NestedWorker) error {
	return nil
}

//Реализация Unmarshaler для json.Unmarshal типа NestedMapType
/*func (nmp *NestedMapType) UnmarshalJSON(b []byte) error {
	var (
		bufferMap map[string]json.RawMessage
		err       error
	)
	if err = json.Unmarshal(b, &bufferMap); err != nil {
		return err
	}
	*nmp = make(NestedMapType)
	for k, v := range bufferMap {
		var lastNode NestedWorker //Последний узел в иерархии
		lastNode, err = (*nmp).BuildNestedHierarchicalBranch(k)

		if nw, err = JSONSwitchCustomType(v); err != nil {
			return err
		}
		err = parentWorker.AddValue(nw)
	}

	return err
}*/

//Реализация Marshaler  для json.Marshal типа NestedMapType
func (nmp NestedMapType) MarshalJSON() ([]byte, error) {
	var (
		b   map[string]json.RawMessage
		raw json.RawMessage
		err error
	)

	b = make(map[string]json.RawMessage)
	for k, v := range nmp {
		b[k], err = json.Marshal(v)
	}
	raw, err = json.Marshal(b)
	return raw, err
}

// SubStrContainArrayNotation - функция проверки стороки на ее соответствие нотации массива "[]"
func SubStrIsArrayNotation(s string) (int, bool, error) {
	var (
		index   int
		isArray bool = false
		err     error
		matched bool
	)
	if matched, err = regexp.MatchString(s, "\\]\\["); matched {
		if !strings.HasPrefix(s, "[") {
			err = fmt.Errorf(`В подстроке %s  не обнаружена открывающая скобка '['`, s)
		}
		if !strings.HasSuffix(s, "]") {
			err = fmt.Errorf(`В подстроке %s не обнаружена закрывающая скобка ']'`, s)
		}
		if err == nil {
			if num := strings.Trim(s, "[]"); num != "" && reflect.ValueOf(num).IsNil() { // Если после Trim-а остаток строки не пустой
				if index, err = strconv.Atoi(num); err != nil { // Пытаемся выделить числовое значение
					err = fmt.Errorf("Подстрока %s не содержит числовое значение", s)
				}
			}
		}
	}
	return index, isArray, err
}

//BuildNestedHierarchicalBranch - функция построения вложенной древовидной структуры узлами которой явлется
// тип данных NestedMapType, данная древовидная структура порождается на основе строкового
// вх. параметра вида "elem1/elem2/.../elemn", "elem1/[number]/.../elemn", применяется для преобразования и размещения в памяти
// ввиде дерава данных приходящих в плоском текстовом JSON виде, например {"elem1/elem2/.../elemn":"value"}
/*func BuildNestedHierarchicalBranch(NWroot NestedWorker, path string) (NestedMapType, error) {
	var (
		lastNode NestedWorker //узел дерева для которого нужно приять решение о продолжении ветви вложенности
		err      error
		index    int  //индекс среза если подключ в нотации массива
		isArray  bool // результат проверки подключа на массив
	)
	if NWroot == nil {
		err = fmt.Errorf("Попытка создать дерево с несуществующим корнем, в BuildNestedHierarchicalBranch в качестве NWroot предан nil")
	}
	lastNode = NWroot // Сначало устанавливаем в качестве последнего узла корень в данном случае сам тип NestedMapType
	keys := strings.Split(path, "/")
	for i, k := range keys {
		index, isArray, err = SubStrIsArrayNotation(k)
		if err != nil {
			if isArray {
				if i == 0 {

				}
				if !reflect.ValueOf(index).IsNil() {

				}
			} else {
				if i != 0 {
					lastNode = nm
				}
				var ok bool
				lastNode, ok = lastNode.(NestedMapType)[k]
				if !ok { // Если ключ не существует
					lastNode = make(NestedMapType)
				} else if lastNode != nil {
				}
			}
		}

	}
	return innerNm
	return nil, nil
}*/

//Методы для NestedStringType реализующие интерфейс NestedWorker
func (ns NestedStringType) GetValue() interface{} {
	return ns
}

//(nsp *NestedStringType) UnmarshalJSON(b []byte) (err error) -
// реализация интерфейса Unmarshaler для Custom-ного типа NestedStringType
func (nsp *NestedStringType) UnmarshalJSON(b []byte) (err error) {
	var (
		bufferString string
	)
	err = json.Unmarshal(b, &bufferString)
	*nsp = NestedStringType(bufferString)
	return err
}

//Методы для NestedDigitType
func (nd NestedDigitType) GetValue() interface{} {
	return nd
}

//JSONSwitchCustomType -
// функция осуществляющая распознавание декодируемых данных в формате JSON
// на основе встроенных типов язака golang и приводящая распознаваемые данные
// к определенным пользователем Custom-ным типам реализующим интерфейс NestedWorker
func JSONSwitchCustomType(b json.RawMessage) (NestedWorker, error) {
	var (
		bufI interface{}
		//nw   NestedWorker
		err error
	)
	if err = json.Unmarshal(b, &bufI); err != nil {
		return nil, err
	}
	bufIType := reflect.TypeOf(bufI)

	switch bufIType.Kind() {
	case reflect.String:
		var Nested NestedStringType
		err = json.Unmarshal(b, &Nested)
		return Nested, err
	case reflect.Slice:
		var Nested NestedSliceType
		err = json.Unmarshal(b, &Nested)
		return Nested, err
	case reflect.Map:
		var Nested NestedMapType
		err = json.Unmarshal(b, &Nested)
		return Nested, err
	default:
		var Nested NestedWorker
		return Nested, err

	}
	//return nil, err
}
