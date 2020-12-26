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
// size 16
type QAbstractItemModel struct {
	*QObject
}
type QAbstractItemModel_ITF interface {
	QObject_ITF
	QAbstractItemModel_PTR() *QAbstractItemModel
}

func (ptr *QAbstractItemModel) QAbstractItemModel_PTR() *QAbstractItemModel { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QAbstractItemModelFromptr(cthis Voidptr) *QAbstractItemModel {
	bcthis0 := QObjectFromptr(cthis)
	return &QAbstractItemModel{bcthis0}
}
func (*QAbstractItemModel) Fromptr(cthis Voidptr) *QAbstractItemModel {
	return QAbstractItemModelFromptr(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemModel(QObject *)

/*
 */
func (*QAbstractItemModel) NewForInherit(parent QObject_ITF /*777 QObject **/) *QAbstractItemModel {
	return NewQAbstractItemModel(parent)
}
func NewQAbstractItemModel(parent QObject_ITF /*777 QObject **/) *QAbstractItemModel {
	var convArg0 Voidptr
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(1538334395, "_ZN18QAbstractItemModelC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QAbstractItemModelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractItemModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemModel(QObject *)

/*
 */
func (*QAbstractItemModel) NewForInheritp() *QAbstractItemModel {
	return NewQAbstractItemModelp()
}
func NewQAbstractItemModelp() *QAbstractItemModel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(1538334395, "_ZN18QAbstractItemModelC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QAbstractItemModelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractItemModel")
	return gothis
}

func DeleteQAbstractItemModel(this *QAbstractItemModel) {
	rv, err := qtrt.Qtcc3(23633593, "_ZN18QAbstractItemModelD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QAbstractItemModel__LayoutChangeHint = int

//
const QAbstractItemModel__NoLayoutChangeHint QAbstractItemModel__LayoutChangeHint = 0

//
const QAbstractItemModel__VerticalSortHint QAbstractItemModel__LayoutChangeHint = 1

//
const QAbstractItemModel__HorizontalSortHint QAbstractItemModel__LayoutChangeHint = 2

func (this *QAbstractItemModel) LayoutChangeHintItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemModel_LayoutChangeHintItemName(val int) string {
	var nilthis *QAbstractItemModel
	return nilthis.LayoutChangeHintItemName(val)
}

/*


 */
type QAbstractItemModel__CheckIndexOption = int

//
const QAbstractItemModel__NoOption QAbstractItemModel__CheckIndexOption = 0

//
const QAbstractItemModel__IndexIsValid QAbstractItemModel__CheckIndexOption = 1

//
const QAbstractItemModel__DoNotUseParent QAbstractItemModel__CheckIndexOption = 2

//
const QAbstractItemModel__ParentIsInvalid QAbstractItemModel__CheckIndexOption = 4

func (this *QAbstractItemModel) CheckIndexOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemModel_CheckIndexOptionItemName(val int) string {
	var nilthis *QAbstractItemModel
	return nilthis.CheckIndexOptionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10019() {
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
