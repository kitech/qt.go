//  header block begin
// /usr/include/qt/QtCore/qitemselectionmodel.h
// #include <qitemselectionmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
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
type QItemSelection struct {
	*qtrt.CObject
}

func (this *QItemSelection) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:250
// index:0
// inline
// void QItemSelection()
func NewQItemSelection() *QItemSelection {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelectionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionFromPointer(cthis)
	return gothis
}
func NewQItemSelectionFromPointer(cthis unsafe.Pointer) *QItemSelection {
	return &QItemSelection{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:251
// index:1
// void QItemSelection(const class QModelIndex &, const class QModelIndex &)
func NewQItemSelection_1(topLeft unsafe.Pointer, bottomRight unsafe.Pointer) *QItemSelection {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelectionC2ERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, cthis, topLeft, bottomRight)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:255
// index:0
// void select(const class QModelIndex &, const class QModelIndex &)
func (this *QItemSelection) Select(topLeft unsafe.Pointer, bottomRight unsafe.Pointer) {
	// 0: (, topLeft const QModelIndex &, bottomRight const QModelIndex &), (topLeft, bottomRight)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelection6selectERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), topLeft, bottomRight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:256
// index:0
// bool contains(const class QModelIndex &)
func (this *QItemSelection) Contains(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QItemSelection8containsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:257
// index:0
// QModelIndexList indexes()
func (this *QItemSelection) Indexes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QItemSelection7indexesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:258
// index:0
// void merge(const class QItemSelection &, class QItemSelectionModel::SelectionFlags)
func (this *QItemSelection) Merge(other unsafe.Pointer, command int) {
	// 0: (, other const QItemSelection &, QFlags<QItemSelectionModel::SelectionFlag> command), (other, &command)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelection5mergeERKS_6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other, &command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:259
// index:0
// static
// void split(const class QItemSelectionRange &, const class QItemSelectionRange &, class QItemSelection *)
func (this *QItemSelection) Split(range_ unsafe.Pointer, other unsafe.Pointer, result unsafe.Pointer) {
	// 0: (range const QItemSelectionRange &, other const QItemSelectionRange &, result QItemSelection *), (range_, other, result)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelection5splitERK19QItemSelectionRangeS2_PS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QItemSelection_Split(range_ unsafe.Pointer, other unsafe.Pointer, result unsafe.Pointer) {
	// 0: (range const QItemSelectionRange &, other const QItemSelectionRange &, result QItemSelection *), (range_, other, result)
	var nilthis *QItemSelection
	nilthis.Split(range_, other, result)
}

//  body block end
