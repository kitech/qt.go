package qtcore

// /usr/include/qt/QtCore/qitemselectionmodel.h
// #include <qitemselectionmodel.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"

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
func (this *QItemSelection) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQItemSelectionFromPointer(cthis unsafe.Pointer) *QItemSelection {
	return &QItemSelection{&qtrt.CObject{cthis}}
}
func (*QItemSelection) NewFromPointer(cthis unsafe.Pointer) *QItemSelection {
	return NewQItemSelectionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:250
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QItemSelection()
func NewQItemSelection() *QItemSelection {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelectionC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemSelection)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:251
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QItemSelection(const QModelIndex &, const QModelIndex &)
func NewQItemSelection_1(topLeft *QModelIndex, bottomRight *QModelIndex) *QItemSelection {
	var convArg0 = topLeft.GetCthis()
	var convArg1 = bottomRight.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelectionC2ERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemSelection)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void select(const QModelIndex &, const QModelIndex &)
func (this *QItemSelection) Select(topLeft *QModelIndex, bottomRight *QModelIndex) {
	var convArg0 = topLeft.GetCthis()
	var convArg1 = bottomRight.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelection6selectERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:256
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QModelIndex &)
func (this *QItemSelection) Contains(index *QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QItemSelection8containsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:258
// index:0
// Public Visibility=Default Availability=Available
// [-2] void merge(const QItemSelection &, QItemSelectionModel::SelectionFlags)
func (this *QItemSelection) Merge(other *QItemSelection, command int) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelection5mergeERKS_6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:259
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void split(const QItemSelectionRange &, const QItemSelectionRange &, QItemSelection *)
func (this *QItemSelection) Split(range_ *QItemSelectionRange, other *QItemSelectionRange, result *QItemSelection /*777 QItemSelection **/) {
	var convArg0 = range_.GetCthis()
	var convArg1 = other.GetCthis()
	var convArg2 = result.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelection5splitERK19QItemSelectionRangeS2_PS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}
func QItemSelection_Split(range_ *QItemSelectionRange, other *QItemSelectionRange, result *QItemSelection /*777 QItemSelection **/) {
	var nilthis *QItemSelection
	nilthis.Split(range_, other, result)
}

func DeleteQItemSelection(this *QItemSelection) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelectionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
