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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlPropertyValueSourceFromPointer(cthis unsafe.Pointer) *QQmlPropertyValueSource {
	return &QQmlPropertyValueSource{&qtrt.CObject{cthis}}
}
func (*QQmlPropertyValueSource) NewFromPointer(cthis unsafe.Pointer) *QQmlPropertyValueSource {
	return NewQQmlPropertyValueSourceFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlpropertyvaluesource.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlPropertyValueSource()
func NewQQmlPropertyValueSource() *QQmlPropertyValueSource {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQmlPropertyValueSourceC1Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlPropertyValueSourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlPropertyValueSource)
	return gothis
}

// /usr/include/qt/QtQml/qqmlpropertyvaluesource.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlPropertyValueSource()
func DeleteQQmlPropertyValueSource(this *QQmlPropertyValueSource) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQmlPropertyValueSourceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlpropertyvaluesource.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setTarget(const QQmlProperty &)
func (this *QQmlPropertyValueSource) SetTarget(arg0 *QQmlProperty) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQmlPropertyValueSource9setTargetERK12QQmlProperty", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
