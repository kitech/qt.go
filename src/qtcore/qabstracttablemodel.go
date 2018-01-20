//  header block begin
// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 62
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
type QAbstractTableModel struct {
	*QAbstractItemModel
}

func (this *QAbstractTableModel) GetCthis() unsafe.Pointer {
	return this.QAbstractItemModel.GetCthis()
}
func NewQAbstractTableModelFromPointer(cthis unsafe.Pointer) *QAbstractTableModel {
	bcthis0 := NewQAbstractItemModelFromPointer(cthis)
	return &QAbstractTableModel{bcthis0}
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:367
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractTableModel) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTableModel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:370
// index:0
// Public
// void QAbstractTableModel(class QObject *)
func NewQAbstractTableModel(parent unsafe.Pointer) *QAbstractTableModel {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTableModelC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractTableModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:371
// index:0
// Public virtual
// void ~QAbstractTableModel()
func DeleteQAbstractTableModel(*QAbstractTableModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTableModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:373
// index:0
// Public virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QAbstractTableModel) Index(row int, column int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTableModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:374
// index:0
// Public virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QAbstractTableModel) Sibling(row int, column int, idx *QModelIndex) interface{} {
	var convArg2 = idx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTableModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:375
// index:0
// Public virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QAbstractTableModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent *QModelIndex) interface{} {
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTableModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data, &action, &row, &column, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:378
// index:0
// Public virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QAbstractTableModel) Flags(index *QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTableModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
