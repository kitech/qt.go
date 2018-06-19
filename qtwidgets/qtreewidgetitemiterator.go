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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QTreeWidgetItemIterator struct {
	*qtrt.CObject
}
type QTreeWidgetItemIterator_ITF interface {
	QTreeWidgetItemIterator_PTR() *QTreeWidgetItemIterator
}

func (ptr *QTreeWidgetItemIterator) QTreeWidgetItemIterator_PTR() *QTreeWidgetItemIterator { return ptr }

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

/*
Constructs an iterator for the same QTreeWidget as it. The current iterator item is set to point on the current item of it.
*/
func NewQTreeWidgetItemIterator(widget QTreeWidget_ITF /*777 QTreeWidget **/, flags int) *QTreeWidgetItemIterator {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QTreeWidget_PTR() != nil {
		convArg0 = widget.QTreeWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorC2EP11QTreeWidget6QFlagsINS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItemIterator)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItemIterator(QTreeWidget *, QTreeWidgetItemIterator::IteratorFlags)

/*
Constructs an iterator for the same QTreeWidget as it. The current iterator item is set to point on the current item of it.
*/
func NewQTreeWidgetItemIterator__(widget QTreeWidget_ITF /*777 QTreeWidget **/) *QTreeWidgetItemIterator {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QTreeWidget_PTR() != nil {
		convArg0 = widget.QTreeWidget_PTR().GetCthis()
	}
	// arg: 1, QTreeWidgetItemIterator::IteratorFlags=Typedef, QTreeWidgetItemIterator::IteratorFlags=Typedef, QFlags<QTreeWidgetItemIterator::IteratorFlag>, Unexposed
	flags := 0
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

/*
Constructs an iterator for the same QTreeWidget as it. The current iterator item is set to point on the current item of it.
*/
func NewQTreeWidgetItemIterator_1(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, flags int) *QTreeWidgetItemIterator {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorC2EP15QTreeWidgetItem6QFlagsINS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTreeWidgetItemIterator)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:86
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidgetItemIterator(QTreeWidgetItem *, QTreeWidgetItemIterator::IteratorFlags)

/*
Constructs an iterator for the same QTreeWidget as it. The current iterator item is set to point on the current item of it.
*/
func NewQTreeWidgetItemIterator_1_(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) *QTreeWidgetItemIterator {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, QTreeWidgetItemIterator::IteratorFlags=Typedef, QTreeWidgetItemIterator::IteratorFlags=Typedef, QFlags<QTreeWidgetItemIterator::IteratorFlag>, Unexposed
	flags := 0
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

/*

 */
func DeleteQTreeWidgetItemIterator(this *QTreeWidgetItemIterator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:89
// index:0
// Public Visibility=Default Availability=Available
// [24] QTreeWidgetItemIterator & operator=(const QTreeWidgetItemIterator &)

/*

 */
func (this *QTreeWidgetItemIterator) Operator_equal(it QTreeWidgetItemIterator_ITF) *QTreeWidgetItemIterator {
	var convArg0 unsafe.Pointer
	if it != nil && it.QTreeWidgetItemIterator_PTR() != nil {
		convArg0 = it.QTreeWidgetItemIterator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratoraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTreeWidgetItemIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTreeWidgetItemIterator)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:91
// index:0
// Public Visibility=Default Availability=Available
// [24] QTreeWidgetItemIterator & operator++()

/*

 */
func (this *QTreeWidgetItemIterator) Operator_add_add() *QTreeWidgetItemIterator {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorppEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTreeWidgetItemIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTreeWidgetItemIterator)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:92
// index:1
// Public inline Visibility=Default Availability=Available
// [24] const QTreeWidgetItemIterator operator++(int)

/*

 */
func (this *QTreeWidgetItemIterator) Operator_add_add_1(arg0 int) *QTreeWidgetItemIterator /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorppEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTreeWidgetItemIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTreeWidgetItemIterator)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QTreeWidgetItemIterator & operator+=(int)

/*

 */
func (this *QTreeWidgetItemIterator) Operator_add_equal(n int) *QTreeWidgetItemIterator {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratorpLEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTreeWidgetItemIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTreeWidgetItemIterator)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:95
// index:0
// Public Visibility=Default Availability=Available
// [24] QTreeWidgetItemIterator & operator--()

/*
The prefix -- operator (--it) advances the iterator to the previous matching item and returns a reference to the resulting iterator. Sets the current pointer to 0 if the current item is the first matching item.
*/
func (this *QTreeWidgetItemIterator) Operator_minus_minus() *QTreeWidgetItemIterator {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratormmEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTreeWidgetItemIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTreeWidgetItemIterator)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:96
// index:1
// Public inline Visibility=Default Availability=Available
// [24] const QTreeWidgetItemIterator operator--(int)

/*
The prefix -- operator (--it) advances the iterator to the previous matching item and returns a reference to the resulting iterator. Sets the current pointer to 0 if the current item is the first matching item.
*/
func (this *QTreeWidgetItemIterator) Operator_minus_minus_1(arg0 int) *QTreeWidgetItemIterator /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratormmEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTreeWidgetItemIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTreeWidgetItemIterator)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QTreeWidgetItemIterator & operator-=(int)

/*

 */
func (this *QTreeWidgetItemIterator) Operator_minus_equal(n int) *QTreeWidgetItemIterator {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QTreeWidgetItemIteratormIEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTreeWidgetItemIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTreeWidgetItemIterator)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidgetitemiterator.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QTreeWidgetItem * operator*() const

/*

 */
func (this *QTreeWidgetItemIterator) Operator_mul() *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QTreeWidgetItemIteratordeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

/*


 */
type QTreeWidgetItemIterator__IteratorFlag = int

//
const QTreeWidgetItemIterator__All QTreeWidgetItemIterator__IteratorFlag = 0

//
const QTreeWidgetItemIterator__Hidden QTreeWidgetItemIterator__IteratorFlag = 1

//
const QTreeWidgetItemIterator__NotHidden QTreeWidgetItemIterator__IteratorFlag = 2

//
const QTreeWidgetItemIterator__Selected QTreeWidgetItemIterator__IteratorFlag = 4

//
const QTreeWidgetItemIterator__Unselected QTreeWidgetItemIterator__IteratorFlag = 8

//
const QTreeWidgetItemIterator__Selectable QTreeWidgetItemIterator__IteratorFlag = 16

//
const QTreeWidgetItemIterator__NotSelectable QTreeWidgetItemIterator__IteratorFlag = 32

//
const QTreeWidgetItemIterator__DragEnabled QTreeWidgetItemIterator__IteratorFlag = 64

//
const QTreeWidgetItemIterator__DragDisabled QTreeWidgetItemIterator__IteratorFlag = 128

//
const QTreeWidgetItemIterator__DropEnabled QTreeWidgetItemIterator__IteratorFlag = 256

//
const QTreeWidgetItemIterator__DropDisabled QTreeWidgetItemIterator__IteratorFlag = 512

//
const QTreeWidgetItemIterator__HasChildren QTreeWidgetItemIterator__IteratorFlag = 1024

//
const QTreeWidgetItemIterator__NoChildren QTreeWidgetItemIterator__IteratorFlag = 2048

//
const QTreeWidgetItemIterator__Checked QTreeWidgetItemIterator__IteratorFlag = 4096

//
const QTreeWidgetItemIterator__NotChecked QTreeWidgetItemIterator__IteratorFlag = 8192

//
const QTreeWidgetItemIterator__Enabled QTreeWidgetItemIterator__IteratorFlag = 16384

//
const QTreeWidgetItemIterator__Disabled QTreeWidgetItemIterator__IteratorFlag = 32768

//
const QTreeWidgetItemIterator__Editable QTreeWidgetItemIterator__IteratorFlag = 65536

//
const QTreeWidgetItemIterator__NotEditable QTreeWidgetItemIterator__IteratorFlag = 131072

//
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
		log.Println(123)
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
