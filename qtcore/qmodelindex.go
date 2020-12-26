// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
// size 24
type QModelIndex struct {
	*qtrt.CObject
}
type QModelIndex_ITF interface {
	QModelIndex_PTR() *QModelIndex
}

func (ptr *QModelIndex) QModelIndex_PTR() *QModelIndex { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QModelIndexFromptr(cthis Voidptr) *QModelIndex {
	return &QModelIndex{&qtrt.CObject{cthis}}
}
func (*QModelIndex) Fromptr(cthis Voidptr) *QModelIndex {
	return QModelIndexFromptr(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QModelIndex()

/*
 */
func (*QModelIndex) NewForInherit() *QModelIndex {
	return NewQModelIndex()
}
func NewQModelIndex() *QModelIndex {
	cthis := qtrt.Malloc(24)
	rv, err := qtrt.Qtcc3(756365278, "_ZN11QModelIndexC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint2(err, rv)
	gothis := QModelIndexFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQModelIndex)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:62
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int row() const

/*
 */
func (this *QModelIndex) Row() int {
	rv, err := qtrt.Qtcc3(3724840192, "_ZNK11QModelIndex3rowEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:63
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int column() const

/*
 */
func (this *QModelIndex) Column() int {
	rv, err := qtrt.Qtcc3(3046022154, "_ZNK11QModelIndex6columnEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:64
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] quintptr internalId() const

/*
 */
func (this *QModelIndex) InternalId() uint64 {
	rv, err := qtrt.Qtcc3(47055808, "_ZNK11QModelIndex10internalIdEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:65
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] void * internalPointer() const

/*
 */
func (this *QModelIndex) InternalPointer() unsafe.Pointer /*666*/ {
	rv, err := qtrt.Qtcc3(2841223468, "_ZNK11QModelIndex15internalPointerEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return Voidptr(uintptr(rv))
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:66
// index:0
// Public inline Indirect Visibility=Default Availability=Available
// [24] QModelIndex parent() const

/*
 */
func (this *QModelIndex) Parent() *QModelIndex /*123*/ {
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(1304276158, "_ZNK11QModelIndex6parentEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:67
// index:0
// Public inline Indirect Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int) const

/*
 */
func (this *QModelIndex) Sibling(row int, column int) *QModelIndex /*123*/ {
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(128912225, "_ZNK11QModelIndex7siblingEii", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, Voidptr(&sretobj), this.Addr(), Voidptr(&row), Voidptr(&column))
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:68
// index:0
// Public inline Indirect Visibility=Default Availability=Available
// [24] QModelIndex siblingAtColumn(int) const

/*
 */
func (this *QModelIndex) SiblingAtColumn(column int) *QModelIndex /*123*/ {
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(2281150179, "_ZNK11QModelIndex15siblingAtColumnEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), this.Addr(), Voidptr(&column))
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:69
// index:0
// Public inline Indirect Visibility=Default Availability=Available
// [24] QModelIndex siblingAtRow(int) const

/*
 */
func (this *QModelIndex) SiblingAtRow(row int) *QModelIndex /*123*/ {
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(3167847458, "_ZNK11QModelIndex12siblingAtRowEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), this.Addr(), Voidptr(&row))
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

func DeleteQModelIndex(this *QModelIndex) {
	rv, err := qtrt.Qtcc3(2965507943, "_ZN11QModelIndexD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10015() {
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
