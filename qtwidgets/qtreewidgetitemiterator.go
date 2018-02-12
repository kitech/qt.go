package qtwidgets

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h
// #include <qtreewidgetitemiterator.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 101
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTreeWidgetItemIteratorFromPointer(cthis unsafe.Pointer) *QTreeWidgetItemIterator {
	return &QTreeWidgetItemIterator{&qtrt.CObject{cthis}}
}
func (*QTreeWidgetItemIterator) NewFromPointer(cthis unsafe.Pointer) *QTreeWidgetItemIterator {
	return NewQTreeWidgetItemIteratorFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItemIterator(QTreeWidget *, QTreeWidgetItemIterator::IteratorFlags)
func NewQTreeWidgetItemIterator(widget *QTreeWidget /*777 QTreeWidget **/, flags int) *QTreeWidgetItemIterator {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorC2EP11QTreeWidget6QFlagsINS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItemIterator)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:86
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItemIterator(QTreeWidgetItem *, QTreeWidgetItemIterator::IteratorFlags)
func NewQTreeWidgetItemIterator_1(item *QTreeWidgetItem /*777 QTreeWidgetItem **/, flags int) *QTreeWidgetItemIterator {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorC2EP15QTreeWidgetItem6QFlagsINS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItemIterator)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTreeWidgetItemIterator()
func DeleteQTreeWidgetItemIterator(this *QTreeWidgetItemIterator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
