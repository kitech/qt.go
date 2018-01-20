//  header block begin
// /usr/include/qt/QtWidgets/qboxlayout.h
// #include <qboxlayout.h>
// #include <QtWidgets>
package qtwidgets

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
type QVBoxLayout struct {
	*QBoxLayout
}

func (this *QVBoxLayout) GetCthis() unsafe.Pointer {
	return this.QBoxLayout.GetCthis()
}

// /usr/include/qt/QtWidgets/qboxlayout.h:128
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QVBoxLayout) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QVBoxLayout10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:130
// index:0
// void QVBoxLayout()
func NewQVBoxLayout() *QVBoxLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QVBoxLayoutC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQVBoxLayoutFromPointer(cthis)
	return gothis
}
func NewQVBoxLayoutFromPointer(cthis unsafe.Pointer) *QVBoxLayout {
	bcthis0 := NewQBoxLayoutFromPointer(cthis)
	return &QVBoxLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qboxlayout.h:131
// index:1
// void QVBoxLayout(class QWidget *)
func NewQVBoxLayout_1(parent unsafe.Pointer) *QVBoxLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QVBoxLayoutC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQVBoxLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:132
// index:0
// virtual
// void ~QVBoxLayout()
func DeleteQVBoxLayout(*QVBoxLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QVBoxLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
