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
func QAbstractTableModelFromptr(cthis unsafe.Pointer) *QAbstractTableModel {
	bcthis0 := QAbstractItemModelFromptr(cthis)
	return &QAbstractTableModel{bcthis0}
}
func (*QAbstractTableModel) Fromptr(cthis unsafe.Pointer) *QAbstractTableModel {
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
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc1(2295160465, "_ZN19QAbstractTableModelC2EP7QObject", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
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
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc1(2295160465, "_ZN19QAbstractTableModelC2EP7QObject", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QAbstractTableModelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractTableModel")
	return gothis
}

func DeleteQAbstractTableModel(this *QAbstractTableModel) {
	rv, err := qtrt.Qtcc1(0, "_ZN19QAbstractTableModelD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
