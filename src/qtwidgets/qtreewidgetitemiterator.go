package qtwidgets

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h
// #include <qtreewidgetitemiterator.h>
// #include <QtWidgets>

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTreeWidgetItemIterator) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTreeWidgetItemIteratorFromPointer(cthis unsafe.Pointer) *QTreeWidgetItemIterator {
	return &QTreeWidgetItemIterator{&qtrt.CObject{cthis}}
}
func (*QTreeWidgetItemIterator) NewFromPointer(cthis unsafe.Pointer) *QTreeWidgetItemIterator {
	return NewQTreeWidgetItemIteratorFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:85
// index:0
// Public
// void QTreeWidgetItemIterator(class QTreeWidget *, QTreeWidgetItemIterator::IteratorFlags)
func NewQTreeWidgetItemIterator(widget *QTreeWidget /*777 QTreeWidget **/, flags int) *QTreeWidgetItemIterator {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorC2EP11QTreeWidget6QFlagsINS_12IteratorFlagEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemIteratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:86
// index:1
// Public
// void QTreeWidgetItemIterator(class QTreeWidgetItem *, QTreeWidgetItemIterator::IteratorFlags)
func NewQTreeWidgetItemIterator_1(item *QTreeWidgetItem /*777 QTreeWidgetItem **/, flags int) *QTreeWidgetItemIterator {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorC2EP15QTreeWidgetItem6QFlagsINS_12IteratorFlagEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemIteratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:87
// index:0
// Public
// void ~QTreeWidgetItemIterator()
func DeleteQTreeWidgetItemIterator(*QTreeWidgetItemIterator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

type QTreeWidgetItemIterator__IteratorFlag = int

const QTreeWidgetItemIterator__All QTreeWidgetItemIterator__IteratorFlag = 0
const QTreeWidgetItemIterator__Hidden QTreeWidgetItemIterator__IteratorFlag = 1
const QTreeWidgetItemIterator__NotHidden QTreeWidgetItemIterator__IteratorFlag = 2
const QTreeWidgetItemIterator__Selected QTreeWidgetItemIterator__IteratorFlag = 4
const QTreeWidgetItemIterator__Unselected QTreeWidgetItemIterator__IteratorFlag = 8
const QTreeWidgetItemIterator__Selectable QTreeWidgetItemIterator__IteratorFlag = 16
const QTreeWidgetItemIterator__NotSelectable QTreeWidgetItemIterator__IteratorFlag = 32
const QTreeWidgetItemIterator__DragEnabled QTreeWidgetItemIterator__IteratorFlag = 64
const QTreeWidgetItemIterator__DragDisabled QTreeWidgetItemIterator__IteratorFlag = 128
const QTreeWidgetItemIterator__DropEnabled QTreeWidgetItemIterator__IteratorFlag = 256
const QTreeWidgetItemIterator__DropDisabled QTreeWidgetItemIterator__IteratorFlag = 512
const QTreeWidgetItemIterator__HasChildren QTreeWidgetItemIterator__IteratorFlag = 1024
const QTreeWidgetItemIterator__NoChildren QTreeWidgetItemIterator__IteratorFlag = 2048
const QTreeWidgetItemIterator__Checked QTreeWidgetItemIterator__IteratorFlag = 4096
const QTreeWidgetItemIterator__NotChecked QTreeWidgetItemIterator__IteratorFlag = 8192
const QTreeWidgetItemIterator__Enabled QTreeWidgetItemIterator__IteratorFlag = 16384
const QTreeWidgetItemIterator__Disabled QTreeWidgetItemIterator__IteratorFlag = 32768
const QTreeWidgetItemIterator__Editable QTreeWidgetItemIterator__IteratorFlag = 65536
const QTreeWidgetItemIterator__NotEditable QTreeWidgetItemIterator__IteratorFlag = 131072
const QTreeWidgetItemIterator__UserFlag QTreeWidgetItemIterator__IteratorFlag = 16777216

//  body block end
