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
// extern C begin: 29
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
type QAbstractTableModel struct {
	*QAbstractItemModel
}
type QAbstractTableModel_ITF interface {
	QAbstractItemModel_ITF
	QAbstractTableModel_PTR() *QAbstractTableModel
}

func (ptr *QAbstractTableModel) QAbstractTableModel_PTR() *QAbstractTableModel { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QAbstractTableModelFromptr(cthis Voidptr) *QAbstractTableModel {
	bcthis0 := QAbstractItemModelFromptr(cthis)
	return &QAbstractTableModel{bcthis0}
}
func (*QAbstractTableModel) Fromptr(cthis Voidptr) *QAbstractTableModel {
	return QAbstractTableModelFromptr(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:392
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractTableModel(QObject *)

/*
 */
func (*QAbstractTableModel) NewForInherit(parent QObject_ITF /*777 QObject **/) *QAbstractTableModel {
	return NewQAbstractTableModel(parent)
}
func NewQAbstractTableModel(parent QObject_ITF /*777 QObject **/) *QAbstractTableModel {
	var convArg0 Voidptr
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(2295160465, "_ZN19QAbstractTableModelC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QAbstractTableModelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractTableModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:392
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractTableModel(QObject *)

/*
 */
func (*QAbstractTableModel) NewForInheritp() *QAbstractTableModel {
	return NewQAbstractTableModelp()
}
func NewQAbstractTableModelp() *QAbstractTableModel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(2295160465, "_ZN19QAbstractTableModelC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QAbstractTableModelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractTableModel")
	return gothis
}

func DeleteQAbstractTableModel(this *QAbstractTableModel) {
	rv, err := qtrt.Qtcc3(2737077743, "_ZN19QAbstractTableModelD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10021() {
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
