//  header block begin
// /usr/include/qt/QtCore/qitemselectionmodel.h
// #include <qitemselectionmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
func NewQItemSelectionRangeFromPointer(cthis unsafe.Pointer) *QItemSelectionRange {
	return &QItemSelectionRange{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:56
// index:0
// Public inline
// void QItemSelectionRange()
func NewQItemSelectionRange() *QItemSelectionRange {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:70
// index:1
// Public inline
// void QItemSelectionRange(const class QModelIndex &, const class QModelIndex &)
func NewQItemSelectionRange_1(topL *QModelIndex, bottomR *QModelIndex) *QItemSelectionRange {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = topL.GetCthis()
	var convArg1 = bottomR.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2ERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:71
// index:2
// Public inline
// void QItemSelectionRange(const class QModelIndex &)
func NewQItemSelectionRange_2(index *QModelIndex) *QItemSelectionRange {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionRangeC2ERK11QModelIndex", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionRangeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:73
// index:0
// Public inline
// void swap(class QItemSelectionRange &)
func (this *QItemSelectionRange) Swap(other *QItemSelectionRange) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionRange4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:79
// index:0
// Public inline
// int top()
func (this *QItemSelectionRange) Top() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange3topEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:80
// index:0
// Public inline
// int left()
func (this *QItemSelectionRange) Left() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange4leftEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:81
// index:0
// Public inline
// int bottom()
func (this *QItemSelectionRange) Bottom() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange6bottomEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:82
// index:0
// Public inline
// int right()
func (this *QItemSelectionRange) Right() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange5rightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:83
// index:0
// Public inline
// int width()
func (this *QItemSelectionRange) Width() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:84
// index:0
// Public inline
// int height()
func (this *QItemSelectionRange) Height() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:86
// index:0
// Public inline
// const QPersistentModelIndex & topLeft()
func (this *QItemSelectionRange) TopLeft() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange7topLeftEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:87
// index:0
// Public inline
// const QPersistentModelIndex & bottomRight()
func (this *QItemSelectionRange) BottomRight() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange11bottomRightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:88
// index:0
// Public inline
// QModelIndex parent()
func (this *QItemSelectionRange) Parent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:89
// index:0
// Public inline
// const QAbstractItemModel * model()
func (this *QItemSelectionRange) Model() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:91
// index:0
// Public inline
// bool contains(const class QModelIndex &)
func (this *QItemSelectionRange) Contains(index *QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange8containsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:98
// index:1
// Public inline
// bool contains(int, int, const class QModelIndex &)
func (this *QItemSelectionRange) Contains_1(row int, column int, parentIndex *QModelIndex) interface{} {
	var convArg2 = parentIndex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange8containsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:105
// index:0
// Public
// bool intersects(const class QItemSelectionRange &)
func (this *QItemSelectionRange) Intersects(other *QItemSelectionRange) interface{} {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange10intersectsERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:110
// index:0
// Public
// QItemSelectionRange intersected(const class QItemSelectionRange &)
func (this *QItemSelectionRange) Intersected(other *QItemSelectionRange) interface{} {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange11intersectedERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:119
// index:0
// Public inline
// bool isValid()
func (this *QItemSelectionRange) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:125
// index:0
// Public
// bool isEmpty()
func (this *QItemSelectionRange) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:127
// index:0
// Public
// QModelIndexList indexes()
func (this *QItemSelectionRange) Indexes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionRange7indexesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
