package qtcore

// /usr/include/qt/QtCore/qitemselectionmodel.h
// #include <qitemselectionmodel.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQItemSelectionFromPointer(cthis unsafe.Pointer) *QItemSelection {
	return &QItemSelection{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:250
// index:0
// Public inline
// void QItemSelection()
func NewQItemSelection() *QItemSelection {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelectionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:251
// index:1
// Public
// void QItemSelection(const class QModelIndex &, const class QModelIndex &)
func NewQItemSelection_1(topLeft *QModelIndex, bottomRight *QModelIndex) *QItemSelection {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = topLeft.GetCthis()
	var convArg1 = bottomRight.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelectionC2ERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:255
// index:0
// Public
// void select(const class QModelIndex &, const class QModelIndex &)
func (this *QItemSelection) Select(topLeft *QModelIndex, bottomRight *QModelIndex) {
	var convArg0 = topLeft.GetCthis()
	var convArg1 = bottomRight.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelection6selectERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:256
// index:0
// Public
// bool contains(const class QModelIndex &)
func (this *QItemSelection) Contains(index *QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QItemSelection8containsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:259
// index:0
// Public static
// void split(const class QItemSelectionRange &, const class QItemSelectionRange &, class QItemSelection *)
func (this *QItemSelection) Split(range_ *QItemSelectionRange, other *QItemSelectionRange, result *QItemSelection /*444 QItemSelection **/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QItemSelection5splitERK19QItemSelectionRangeS2_PS_", ffiqt.FFI_TYPE_POINTER, range_, other, result)
	gopp.ErrPrint(err, rv)
}
func QItemSelection_Split(range_ *QItemSelectionRange, other *QItemSelectionRange, result *QItemSelection /*444 QItemSelection **/) {
	var nilthis *QItemSelection
	nilthis.Split(range_, other, result)
}

//  body block end
