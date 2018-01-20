//  header block begin
// /usr/include/qt/QtWidgets/qboxlayout.h
// #include <qboxlayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 36
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
type QHBoxLayout struct {
	*QBoxLayout
}

func (this *QHBoxLayout) GetCthis() unsafe.Pointer {
	return this.QBoxLayout.GetCthis()
}

// /usr/include/qt/QtWidgets/qboxlayout.h:115
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QHBoxLayout) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHBoxLayout10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:117
// index:0
// void QHBoxLayout()
func NewQHBoxLayout() *QHBoxLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHBoxLayoutC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQHBoxLayoutFromPointer(cthis)
	return gothis
}
func NewQHBoxLayoutFromPointer(cthis unsafe.Pointer) *QHBoxLayout {
	bcthis0 := NewQBoxLayoutFromPointer(cthis)
	return &QHBoxLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qboxlayout.h:118
// index:1
// void QHBoxLayout(class QWidget *)
func NewQHBoxLayout_1(parent unsafe.Pointer) *QHBoxLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHBoxLayoutC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQHBoxLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:119
// index:0
// virtual
// void ~QHBoxLayout()
func DeleteQHBoxLayout(*QHBoxLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHBoxLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
