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
type QAbstractListModel struct {
	*QAbstractItemModel
}
type QAbstractListModel_ITF interface {
	QAbstractItemModel_ITF
	QAbstractListModel_PTR() *QAbstractListModel
}

func (ptr *QAbstractListModel) QAbstractListModel_PTR() *QAbstractListModel { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QAbstractListModelFromptr(cthis Voidptr) *QAbstractListModel {
	bcthis0 := QAbstractItemModelFromptr(cthis)
	return &QAbstractListModel{bcthis0}
}
func (*QAbstractListModel) Fromptr(cthis Voidptr) *QAbstractListModel {
	return QAbstractListModelFromptr(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:418
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractListModel(QObject *)

/*
 */
func (*QAbstractListModel) NewForInherit(parent QObject_ITF /*777 QObject **/) *QAbstractListModel {
	return NewQAbstractListModel(parent)
}
func NewQAbstractListModel(parent QObject_ITF /*777 QObject **/) *QAbstractListModel {
	var convArg0 Voidptr
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(3627761914, "_ZN18QAbstractListModelC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QAbstractListModelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractListModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:418
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractListModel(QObject *)

/*
 */
func (*QAbstractListModel) NewForInheritp() *QAbstractListModel {
	return NewQAbstractListModelp()
}
func NewQAbstractListModelp() *QAbstractListModel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(3627761914, "_ZN18QAbstractListModelC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QAbstractListModelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractListModel")
	return gothis
}

func DeleteQAbstractListModel(this *QAbstractListModel) {
	rv, err := qtrt.Qtcc3(319317266, "_ZN18QAbstractListModelD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10023() {
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
