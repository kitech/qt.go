//  header block begin
// /usr/include/qt/QtCore/qitemselectionmodel.h
// #include <qitemselectionmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:250
// index:0
// inline
// void QItemSelection()
func NewQItemSelection() *QItemSelection {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelectionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QItemSelection{cthis}
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:251
// index:1
// void QItemSelection(const class QModelIndex &, const class QModelIndex &)
func NewQItemSelection_1(topLeft unsafe.Pointer, bottomRight unsafe.Pointer) *QItemSelection {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelectionC2ERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, cthis, topLeft, bottomRight)
	gopp.ErrPrint(err, rv)
	return &QItemSelection{cthis}
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:255
// index:0
// void select(const class QModelIndex &, const class QModelIndex &)
func (this *QItemSelection) Select(topLeft unsafe.Pointer, bottomRight unsafe.Pointer) {
	// 0: (, const QModelIndex & topLeft, const QModelIndex & bottomRight), (topLeft, bottomRight)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelection6selectERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.cthis, topLeft, bottomRight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:256
// index:0
// bool contains(const class QModelIndex &)
func (this *QItemSelection) Contains(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QItemSelection8containsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:257
// index:0
// QModelIndexList indexes()
func (this *QItemSelection) Indexes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QItemSelection7indexesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:259
// index:0
// static
// void split(const class QItemSelectionRange &, const class QItemSelectionRange &, class QItemSelection *)
func (this *QItemSelection) Split(range_ unsafe.Pointer, other unsafe.Pointer, result unsafe.Pointer) {
	// 0: (const QItemSelectionRange & range, const QItemSelectionRange & other, QItemSelection * result), (range_, other, result)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelection5splitERK19QItemSelectionRangeS2_PS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QItemSelection_Split(range_ unsafe.Pointer, other unsafe.Pointer, result unsafe.Pointer) {
	// 0: (const QItemSelectionRange & range, const QItemSelectionRange & other, QItemSelection * result), (range_, other, result)
	var nilthis *QItemSelection
	nilthis.Split(range_, other, result)
}

//  body block end
