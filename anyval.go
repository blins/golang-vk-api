package vkapi

import (
	"fmt"
	"strconv"
)

/*
Чтобы проще было работать со значениями которые передаются вот так неопределенно
*/
type AnyVal struct{
	Value interface{}
}

/*
Возвращает строковое представление.
Работает если значение имеет типы:
 - int
 - string
 - fmt.Stringer
*/
func (v AnyVal) String() string {
	switch tv := v.Value.(type) {
	case int:
		return strconv.Itoa(tv)
	case uint64:
		return strconv.FormatUint(tv, 10)
	case int64:
		return strconv.FormatInt(tv, 10)
	case uint32:
		return strconv.FormatUint(uint64(tv), 10)
	case int32:
		return strconv.FormatInt(int64(tv), 10)
	case uint:
		return strconv.FormatUint(uint64(tv), 10)
	case []byte:
		return string(tv)
	case string:
		return tv
	case fmt.Stringer:
		return tv.String()
	}
	return ""
}

/*
Возвращает целочисленное представление.
Работает если значение имеет типы:
 - int
 - string
 - fmt.Stringer
*/
func (v AnyVal) Int() int {
	switch tv := v.Value.(type) {
	case int:
		return tv
	case uint64:
		return int(tv)
	case int64:
		return int(tv)
	case uint32:
		return int(tv)
	case int32:
		return int(tv)
	case uint:
		return int(tv)
	case string:
		i, _ := strconv.Atoi(tv)
		return i
	case fmt.Stringer:
		i,_ := strconv.Atoi(tv.String())
		return i
	}
	return 0
}

func (v AnyVal) Int64() int64 {
	switch tv := v.Value.(type) {
	case int:
		return int64(tv)
	case uint64:
		return int64(tv)
	case int64:
		return tv
	case uint32:
		return int64(tv)
	case int32:
		return int64(tv)
	case uint:
		return int64(tv)
	case string:
		i, _ := strconv.ParseInt(tv, 10, 64)
		return i
	case fmt.Stringer:
		i, _ := strconv.ParseInt(tv.String(), 10, 64)
		return i
	}
	return 0
}

