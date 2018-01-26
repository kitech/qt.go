package qtqml

// /usr/include/qt/QtQml/qqmlpropertyvaluesource.h
// #include <qqmlpropertyvaluesource.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQmlPropertyValueSource struct {
	*qtrt.CObject
}

func (this *QQmlPropertyValueSource) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlPropertyValueSource) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQQmlPropertyValueSourceFromPointer(cthis unsafe.Pointer) *QQmlPropertyValueSource {
	return &QQmlPropertyValueSource{&qtrt.CObject{cthis}}
}
func (*QQmlPropertyValueSource) NewFromPointer(cthis unsafe.Pointer) *QQmlPropertyValueSource {
	return NewQQmlPropertyValueSourceFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlpropertyvaluesource.h:53
// index:0
// Public
// void QQmlPropertyValueSource()
func NewQQmlPropertyValueSource() *QQmlPropertyValueSource {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QQmlPropertyValueSourceC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlPropertyValueSourceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlpropertyvaluesource.h:54
// index:0
// Public virtual
// void ~QQmlPropertyValueSource()
func DeleteQQmlPropertyValueSource(*QQmlPropertyValueSource) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QQmlPropertyValueSourceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlpropertyvaluesource.h:55
// index:0
// Public pure virtual
// void setTarget(const class QQmlProperty &)
func (this *QQmlPropertyValueSource) SetTarget(arg0 *QQmlProperty) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QQmlPropertyValueSource9setTargetERK12QQmlProperty", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
