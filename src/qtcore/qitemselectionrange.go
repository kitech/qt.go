//  header block begin
// /usr/include/qt/QtCore/qitemselectionmodel.h
// #include <qitemselectionmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QItemSelectionRange struct {
	*qtrt.CObject
}

func (this *QItemSelectionRange) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:56
// index:0
// inline
// void QItemSelectionRange()
func NewQItemSelectionRange() *QItemSelectionRange {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(cthis)
	return gothis
}
func NewQItemSelectionRangeFromPointer(cthis unsafe.Pointer) *QItemSelectionRange {
	return &QItemSelectionRange{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:70
// index:1
// inline
// void QItemSelectionRange(const class QModelIndex &, const class QModelIndex &)
func NewQItemSelectionRange_1(topL unsafe.Pointer, bottomR unsafe.Pointer) *QItemSelectionRange {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2ERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, cthis, topL, bottomR)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:71
// index:2
// inline
// void QItemSelectionRange(const class QModelIndex &)
func NewQItemSelectionRange_2(index unsafe.Pointer) *QItemSelectionRange {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2ERK11QModelIndex", ffiqt.FFI_TYPE_VOID, cthis, index)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:73
// index:0
// inline
// void swap(class QItemSelectionRange &)
func (this *QItemSelectionRange) Swap(other unsafe.Pointer) {
	// 0: (, other QItemSelectionRange &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionRange4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:79
// index:0
// inline
// int top()
func (this *QItemSelectionRange) Top() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange3topEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:80
// index:0
// inline
// int left()
func (this *QItemSelectionRange) Left() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange4leftEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:81
// index:0
// inline
// int bottom()
func (this *QItemSelectionRange) Bottom() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange6bottomEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:82
// index:0
// inline
// int right()
func (this *QItemSelectionRange) Right() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange5rightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:83
// index:0
// inline
// int width()
func (this *QItemSelectionRange) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange5widthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:84
// index:0
// inline
// int height()
func (this *QItemSelectionRange) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange6heightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:86
// index:0
// inline
// const QPersistentModelIndex & topLeft()
func (this *QItemSelectionRange) TopLeft() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange7topLeftEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:87
// index:0
// inline
// const QPersistentModelIndex & bottomRight()
func (this *QItemSelectionRange) BottomRight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange11bottomRightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:88
// index:0
// inline
// QModelIndex parent()
func (this *QItemSelectionRange) Parent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange6parentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:89
// index:0
// inline
// const QAbstractItemModel * model()
func (this *QItemSelectionRange) Model() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange5modelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:91
// index:0
// inline
// bool contains(const class QModelIndex &)
func (this *QItemSelectionRange) Contains(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange8containsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:98
// index:1
// inline
// bool contains(int, int, const class QModelIndex &)
func (this *QItemSelectionRange) Contains_1(row int, column int, parentIndex unsafe.Pointer) {
	// 1: (, row int, column int, parentIndex const QModelIndex &), (&row, &column, parentIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange8containsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, parentIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:105
// index:0
// bool intersects(const class QItemSelectionRange &)
func (this *QItemSelectionRange) Intersects(other unsafe.Pointer) {
	// 0: (, other const QItemSelectionRange &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange10intersectsERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:110
// index:0
// QItemSelectionRange intersected(const class QItemSelectionRange &)
func (this *QItemSelectionRange) Intersected(other unsafe.Pointer) {
	// 0: (, other const QItemSelectionRange &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange11intersectedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:119
// index:0
// inline
// bool isValid()
func (this *QItemSelectionRange) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:125
// index:0
// bool isEmpty()
func (this *QItemSelectionRange) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:127
// index:0
// QModelIndexList indexes()
func (this *QItemSelectionRange) Indexes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange7indexesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
