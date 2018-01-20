//  header block begin
// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h
// #include <qtreewidgetitemiterator.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 99
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
type QTreeWidgetItemIterator struct {
	*qtrt.CObject
}

func (this *QTreeWidgetItemIterator) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTreeWidgetItemIteratorFromPointer(cthis unsafe.Pointer) *QTreeWidgetItemIterator {
	return &QTreeWidgetItemIterator{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:87
// index:0
// Public
// void ~QTreeWidgetItemIterator()
func DeleteQTreeWidgetItemIterator(*QTreeWidgetItemIterator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
