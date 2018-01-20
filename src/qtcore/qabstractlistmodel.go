//  header block begin
// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QAbstractListModel struct {
	*QAbstractItemModel
}

func (this *QAbstractListModel) GetCthis() unsafe.Pointer {
	return this.QAbstractItemModel.GetCthis()
}
func NewQAbstractListModelFromPointer(cthis unsafe.Pointer) *QAbstractListModel {
	bcthis0 := NewQAbstractItemModelFromPointer(cthis)
	return &QAbstractListModel{bcthis0}
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:393
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractListModel) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractListModel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:396
// index:0
// Public
// void QAbstractListModel(class QObject *)
func NewQAbstractListModel(parent unsafe.Pointer) *QAbstractListModel {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractListModelC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractListModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:397
// index:0
// Public virtual
// void ~QAbstractListModel()
func DeleteQAbstractListModel(*QAbstractListModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractListModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:399
// index:0
// Public virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QAbstractListModel) Index(row int, column int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractListModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:400
// index:0
// Public virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QAbstractListModel) Sibling(row int, column int, idx *QModelIndex) interface{} {
	var convArg2 = idx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractListModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:401
// index:0
// Public virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QAbstractListModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent *QModelIndex) interface{} {
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractListModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data, &action, &row, &column, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:404
// index:0
// Public virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QAbstractListModel) Flags(index *QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractListModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
