package qtcore

// /usr/include/qt/QtCore/qitemselectionmodel.h
// #include <qitemselectionmodel.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QItemSelectionRange struct {
	*qtrt.CObject
}
type QItemSelectionRange_ITF interface {
	QItemSelectionRange_PTR() *QItemSelectionRange
}

func (ptr *QItemSelectionRange) QItemSelectionRange_PTR() *QItemSelectionRange { return ptr }

func (this *QItemSelectionRange) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QItemSelectionRange) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQItemSelectionRangeFromPointer(cthis unsafe.Pointer) *QItemSelectionRange {
	return &QItemSelectionRange{&qtrt.CObject{cthis}}
}
func (*QItemSelectionRange) NewFromPointer(cthis unsafe.Pointer) *QItemSelectionRange {
	return NewQItemSelectionRangeFromPointer(cthis)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:56
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QItemSelectionRange()
func NewQItemSelectionRange() *QItemSelectionRange {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemSelectionRange)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:70
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QItemSelectionRange(const QModelIndex &, const QModelIndex &)
func NewQItemSelectionRange_1(topL *QModelIndex, bottomR *QModelIndex) *QItemSelectionRange {
	var convArg0 = topL.GetCthis()
	var convArg1 = bottomR.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2ERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemSelectionRange)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:71
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QItemSelectionRange(const QModelIndex &)
func NewQItemSelectionRange_2(index *QModelIndex) *QItemSelectionRange {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2ERK11QModelIndex", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemSelectionRange)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QItemSelectionRange &)
func (this *QItemSelectionRange) Swap(other *QItemSelectionRange) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionRange4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int top()
func (this *QItemSelectionRange) Top() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange3topEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int left()
func (this *QItemSelectionRange) Left() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange4leftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int bottom()
func (this *QItemSelectionRange) Bottom() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange6bottomEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int right()
func (this *QItemSelectionRange) Right() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange5rightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int width()
func (this *QItemSelectionRange) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int height()
func (this *QItemSelectionRange) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPersistentModelIndex & topLeft()
func (this *QItemSelectionRange) TopLeft() *QPersistentModelIndex {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange7topLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPersistentModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:87
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPersistentModelIndex & bottomRight()
func (this *QItemSelectionRange) BottomRight() *QPersistentModelIndex {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange11bottomRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPersistentModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex parent()
func (this *QItemSelectionRange) Parent() *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QAbstractItemModel * model()
func (this *QItemSelectionRange) Model() *QAbstractItemModel /*777 const QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QModelIndex &)
func (this *QItemSelectionRange) Contains(index *QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange8containsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:98
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(int, int, const QModelIndex &)
func (this *QItemSelectionRange) Contains_1(row int, column int, parentIndex *QModelIndex) bool {
	var convArg2 = parentIndex.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange8containsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QItemSelectionRange &)
func (this *QItemSelectionRange) Intersects(other *QItemSelectionRange) bool {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange10intersectsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:110
// index:0
// Public Visibility=Default Availability=Available
// [16] QItemSelectionRange intersected(const QItemSelectionRange &)
func (this *QItemSelectionRange) Intersected(other *QItemSelectionRange) *QItemSelectionRange /*123*/ {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange11intersectedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelectionRange)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:119
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QItemSelectionRange) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:125
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QItemSelectionRange) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QModelIndexList indexes()
func (this *QItemSelectionRange) Indexes() *QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange7indexesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

func DeleteQItemSelectionRange(this *QItemSelectionRange) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionRangeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
		qtrt.KeepMe()
	}
}

//  keep block end
