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
// extern C begin: 31
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
type QItemSelection struct {
	*qtrt.CObject
}
type QItemSelection_ITF interface {
	QItemSelection_PTR() *QItemSelection
}

func (ptr *QItemSelection) QItemSelection_PTR() *QItemSelection { return ptr }

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

/*

 */
func (*QItemSelection) NewForInherit() *QItemSelection {
	return NewQItemSelection()
}
func NewQItemSelection() *QItemSelection {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelectionC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemSelection)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:251
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QItemSelection(const QModelIndex &, const QModelIndex &)

/*

 */
func (*QItemSelection) NewForInherit1(topLeft QModelIndex_ITF, bottomRight QModelIndex_ITF) *QItemSelection {
	return NewQItemSelection1(topLeft, bottomRight)
}
func NewQItemSelection1(topLeft QModelIndex_ITF, bottomRight QModelIndex_ITF) *QItemSelection {
	var convArg0 unsafe.Pointer
	if topLeft != nil && topLeft.QModelIndex_PTR() != nil {
		convArg0 = topLeft.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if bottomRight != nil && bottomRight.QModelIndex_PTR() != nil {
		convArg1 = bottomRight.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelectionC2ERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemSelection)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void select(const QModelIndex &, const QModelIndex &)

/*
Selects the model item index using the specified command, and emits selectionChanged().

See also QItemSelectionModel::SelectionFlags.
*/
func (this *QItemSelection) Select(topLeft QModelIndex_ITF, bottomRight QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if topLeft != nil && topLeft.QModelIndex_PTR() != nil {
		convArg0 = topLeft.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if bottomRight != nil && bottomRight.QModelIndex_PTR() != nil {
		convArg1 = bottomRight.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelection6selectERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:256
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QModelIndex &) const

/*

 */
func (this *QItemSelection) Contains(index QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QItemSelection8containsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:257
// index:0
// Public Visibility=Default Availability=Available
// [8] QModelIndexList indexes() const

/*

 */
func (this *QItemSelection) Indexes() *QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QItemSelection7indexesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:258
// index:0
// Public Visibility=Default Availability=Available
// [-2] void merge(const QItemSelection &, QItemSelectionModel::SelectionFlags)

/*

 */
func (this *QItemSelection) Merge(other QItemSelection_ITF, command int) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QItemSelection_PTR() != nil {
		convArg0 = other.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelection5mergeERKS_6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:259
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void split(const QItemSelectionRange &, const QItemSelectionRange &, QItemSelection *)

/*

 */
func (this *QItemSelection) Split(range_ QItemSelectionRange_ITF, other QItemSelectionRange_ITF, result QItemSelection_ITF /*777 QItemSelection **/) {
	var convArg0 unsafe.Pointer
	if range_ != nil && range_.QItemSelectionRange_PTR() != nil {
		convArg0 = range_.QItemSelectionRange_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if other != nil && other.QItemSelectionRange_PTR() != nil {
		convArg1 = other.QItemSelectionRange_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if result != nil && result.QItemSelection_PTR() != nil {
		convArg2 = result.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelection5splitERK19QItemSelectionRangeS2_PS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}
func QItemSelection_Split(range_ QItemSelectionRange_ITF, other QItemSelectionRange_ITF, result QItemSelection_ITF /*777 QItemSelection **/) {
	var nilthis *QItemSelection
	nilthis.Split(range_, other, result)
}

func DeleteQItemSelection(this *QItemSelection) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QItemSelectionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
