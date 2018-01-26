package qtcore

// /usr/include/qt/QtCore/qhash.h
// #include <qhash.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QHashDummyValue struct {
	*qtrt.CObject
}

func (this *QHashDummyValue) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHashDummyValue) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQHashDummyValueFromPointer(cthis unsafe.Pointer) *QHashDummyValue {
	return &QHashDummyValue{&qtrt.CObject{cthis}}
}
func (*QHashDummyValue) NewFromPointer(cthis unsafe.Pointer) *QHashDummyValue {
	return NewQHashDummyValueFromPointer(cthis)
}

//  body block end
