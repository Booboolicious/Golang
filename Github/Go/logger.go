package mylog

import (
	"fmt"
)

type logger func(...any)

func (l logger) err(format string, a ...any) error {
	err := fmt.Errorf(format, a...)
	l("ERROR:", err)
	return err
}

func (l logger) typeof(v any) {
	l(fmt.Sprintf("TYPE: %T", v ))
	// l(fmt.Sprintf("TYPE: %v", reflect.TypeOf(v)))
}

var log logger = func(a ...any) {
	fmt.Println(a...)
}