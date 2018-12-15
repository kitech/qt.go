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
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
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

// /usr/include/qt/QtCore/qitemselectionmodel.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QItemSelectionRange()

/*

 */
func (*QItemSelectionRange) NewForInherit() *QItemSelectionRange {
	return NewQItemSelectionRange()
}
func NewQItemSelectionRange() *QItemSelectionRange {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemSelectionRange)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:72
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QItemSelectionRange(const QModelIndex &, const QModelIndex &)

/*

 */
func (*QItemSelectionRange) NewForInherit1(topL QModelIndex_ITF, bottomR QModelIndex_ITF) *QItemSelectionRange {
	return NewQItemSelectionRange1(topL, bottomR)
}
func NewQItemSelectionRange1(topL QModelIndex_ITF, bottomR QModelIndex_ITF) *QItemSelectionRange {
	var convArg0 unsafe.Pointer
	if topL != nil && topL.QModelIndex_PTR() != nil {
		convArg0 = topL.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if bottomR != nil && bottomR.QModelIndex_PTR() != nil {
		convArg1 = bottomR.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2ERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemSelectionRange)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:73
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QItemSelectionRange(const QModelIndex &)

/*

 */
func (*QItemSelectionRange) NewForInherit2(index QModelIndex_ITF) *QItemSelectionRange {
	return NewQItemSelectionRange2(index)
}
func NewQItemSelectionRange2(index QModelIndex_ITF) *QItemSelectionRange {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2ERK11QModelIndex", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemSelectionRange)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QItemSelectionRange & operator=(QItemSelectionRange &&)

/*

 */
func (this *QItemSelectionRange) Operator_equal(other unsafe.Pointer /*333*/) *QItemSelectionRange {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionRangeaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelectionRange)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:69
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QItemSelectionRange & operator=(const QItemSelectionRange &)

/*

 */
func (this *QItemSelectionRange) Operator_equal1(other QItemSelectionRange_ITF) *QItemSelectionRange {
	var convArg0 unsafe.Pointer
	if other != nil && other.QItemSelectionRange_PTR() != nil {
		convArg0 = other.QItemSelectionRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionRangeaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelectionRange)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QItemSelectionRange &)

/*

 */
func (this *QItemSelectionRange) Swap(other QItemSelectionRange_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QItemSelectionRange_PTR() != nil {
		convArg0 = other.QItemSelectionRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionRange4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int top() const

/*

 */
func (this *QItemSelectionRange) Top() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange3topEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int left() const

/*

 */
func (this *QItemSelectionRange) Left() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange4leftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int bottom() const

/*

 */
func (this *QItemSelectionRange) Bottom() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange6bottomEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int right() const

/*

 */
func (this *QItemSelectionRange) Right() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange5rightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int width() const

/*

 */
func (this *QItemSelectionRange) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int height() const

/*

 */
func (this *QItemSelectionRange) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPersistentModelIndex & topLeft() const

/*

 */
func (this *QItemSelectionRange) TopLeft() *QPersistentModelIndex {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange7topLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPersistentModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPersistentModelIndex & bottomRight() const

/*

 */
func (this *QItemSelectionRange) BottomRight() *QPersistentModelIndex {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange11bottomRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPersistentModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex parent() const

/*

 */
func (this *QItemSelectionRange) Parent() *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QAbstractItemModel * model() const

/*
Returns the item model operated on by the selection model.

See also setModel().
*/
func (this *QItemSelectionRange) Model() *QAbstractItemModel /*777 const QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QModelIndex &) const

/*

 */
func (this *QItemSelectionRange) Contains(index QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange8containsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:100
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(int, int, const QModelIndex &) const

/*

 */
func (this *QItemSelectionRange) Contains1(row int, column int, parentIndex QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parentIndex != nil && parentIndex.QModelIndex_PTR() != nil {
		convArg2 = parentIndex.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange8containsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QItemSelectionRange &) const

/*

 */
func (this *QItemSelectionRange) Intersects(other QItemSelectionRange_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QItemSelectionRange_PTR() != nil {
		convArg0 = other.QItemSelectionRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange10intersectsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:112
// index:0
// Public Visibility=Default Availability=Available
// [16] QItemSelectionRange intersected(const QItemSelectionRange &) const

/*

 */
func (this *QItemSelectionRange) Intersected(other QItemSelectionRange_ITF) *QItemSelectionRange /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QItemSelectionRange_PTR() != nil {
		convArg0 = other.QItemSelectionRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange11intersectedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelectionRange)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QItemSelectionRange &) const

/*

 */
func (this *QItemSelectionRange) Operator_equal_equal(other QItemSelectionRange_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QItemSelectionRange_PTR() != nil {
		convArg0 = other.QItemSelectionRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRangeeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QItemSelectionRange &) const

/*

 */
func (this *QItemSelectionRange) Operator_not_equal(other QItemSelectionRange_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QItemSelectionRange_PTR() != nil {
		convArg0 = other.QItemSelectionRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRangeneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:119
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator<(const QItemSelectionRange &) const

/*

 */
func (this *QItemSelectionRange) Operator_less_than(other QItemSelectionRange_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QItemSelectionRange_PTR() != nil {
		convArg0 = other.QItemSelectionRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRangeltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QItemSelectionRange) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*

 */
func (this *QItemSelectionRange) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionRange7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QModelIndexList indexes() const

/*

 */
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
