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
// extern C begin: 9
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
// size 8
type QPersistentModelIndex struct {
	*qtrt.CObject
}
type QPersistentModelIndex_ITF interface {
	QPersistentModelIndex_PTR() *QPersistentModelIndex
}

func (ptr *QPersistentModelIndex) QPersistentModelIndex_PTR() *QPersistentModelIndex { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QPersistentModelIndexFromptr(cthis Voidptr) *QPersistentModelIndex {
	return &QPersistentModelIndex{&qtrt.CObject{cthis}}
}
func (*QPersistentModelIndex) Fromptr(cthis Voidptr) *QPersistentModelIndex {
	return QPersistentModelIndexFromptr(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPersistentModelIndex()

/*
 */
func (*QPersistentModelIndex) NewForInherit() *QPersistentModelIndex {
	return NewQPersistentModelIndex()
}
func NewQPersistentModelIndex() *QPersistentModelIndex {
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc3(4053232108, "_ZN21QPersistentModelIndexC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint2(err, rv)
	gothis := QPersistentModelIndexFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQPersistentModelIndex)
	return gothis
}

func DeleteQPersistentModelIndex(this *QPersistentModelIndex) {
	rv, err := qtrt.Qtcc1(0, "_ZN21QPersistentModelIndexD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10017() {
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
