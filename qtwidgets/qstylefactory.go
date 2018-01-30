package qtwidgets

// /usr/include/qt/QtWidgets/qstylefactory.h
// #include <qstylefactory.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStyleFactory struct {
	*qtrt.CObject
}

func (this *QStyleFactory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStyleFactory) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQStyleFactoryFromPointer(cthis unsafe.Pointer) *QStyleFactory {
	return &QStyleFactory{&qtrt.CObject{cthis}}
}
func (*QStyleFactory) NewFromPointer(cthis unsafe.Pointer) *QStyleFactory {
	return NewQStyleFactoryFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstylefactory.h:55
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStyle * create(const QString &)
func (this *QStyleFactory) Create(arg0 *qtcore.QString) *QStyle /*777 QStyle **/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStyleFactory6createERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QStyleFactory_Create(arg0 *qtcore.QString) *QStyle /*777 QStyle **/ {
	var nilthis *QStyleFactory
	rv := nilthis.Create(arg0)
	return rv
}

//  body block end
