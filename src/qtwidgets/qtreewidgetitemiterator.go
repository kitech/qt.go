//  header block begin
// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h
// #include <qtreewidgetitemiterator.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 102
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

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:85
// index:0
// void QTreeWidgetItemIterator(class QTreeWidget *, QTreeWidgetItemIterator::IteratorFlags)
func NewQTreeWidgetItemIterator(widget unsafe.Pointer, flags int) *QTreeWidgetItemIterator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorC2EP11QTreeWidget6QFlagsINS_12IteratorFlagEE", ffiqt.FFI_TYPE_VOID, cthis, widget, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemIteratorFromPointer(cthis)
	return gothis
}
func NewQTreeWidgetItemIteratorFromPointer(cthis unsafe.Pointer) *QTreeWidgetItemIterator {
	return &QTreeWidgetItemIterator{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:86
// index:1
// void QTreeWidgetItemIterator(class QTreeWidgetItem *, QTreeWidgetItemIterator::IteratorFlags)
func NewQTreeWidgetItemIterator_1(item unsafe.Pointer, flags int) *QTreeWidgetItemIterator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorC2EP15QTreeWidgetItem6QFlagsINS_12IteratorFlagEE", ffiqt.FFI_TYPE_VOID, cthis, item, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemIteratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:87
// index:0
// void ~QTreeWidgetItemIterator()
func DeleteQTreeWidgetItemIterator(*QTreeWidgetItemIterator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
