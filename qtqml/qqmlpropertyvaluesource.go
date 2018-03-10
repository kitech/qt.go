package qtqml

// /usr/include/qt/QtQml/qqmlpropertyvaluesource.h
// #include <qqmlpropertyvaluesource.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QQmlPropertyValueSource struct {
	*qtrt.CObject
}
type QQmlPropertyValueSource_ITF interface {
	QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource
}

func (ptr *QQmlPropertyValueSource) QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource { return ptr }

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

/*
Constructs a QQmlPropertyValueSource.
*/
func NewQQmlPropertyValueSource() *QQmlPropertyValueSource {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQmlPropertyValueSourceC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlPropertyValueSourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlPropertyValueSource)
	return gothis
}

// /usr/include/qt/QtQml/qqmlpropertyvaluesource.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlPropertyValueSource()

/*

 */
func DeleteQQmlPropertyValueSource(this *QQmlPropertyValueSource) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQmlPropertyValueSourceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlpropertyvaluesource.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setTarget(const QQmlProperty &)

/*
Set the target property for the value source. This method will be called by the QML engine when assigning a value source.
*/
func (this *QQmlPropertyValueSource) SetTarget(arg0 QQmlProperty_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlProperty_PTR() != nil {
		convArg0 = arg0.QQmlProperty_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQmlPropertyValueSource9setTargetERK12QQmlProperty", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
