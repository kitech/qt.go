//  header block begin
// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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

// /usr/include/qt/QtCore/qabstractitemmodel.h:393
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractListModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractListModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:396
// index:0
// void QAbstractListModel(class QObject *)
func NewQAbstractListModel(parent unsafe.Pointer) *QAbstractListModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractListModelC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractListModelFromPointer(cthis)
	return gothis
}
func NewQAbstractListModelFromPointer(cthis unsafe.Pointer) *QAbstractListModel {
	bcthis0 := NewQAbstractItemModelFromPointer(cthis)
	return &QAbstractListModel{bcthis0}
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:409
// index:1
// void QAbstractListModel(class QAbstractItemModelPrivate &, class QObject *)
func NewQAbstractListModel_1(dd unsafe.Pointer, parent unsafe.Pointer) *QAbstractListModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractListModelC1ER25QAbstractItemModelPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractListModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:397
// index:0
// virtual
// void ~QAbstractListModel()
func DeleteQAbstractListModel(*QAbstractListModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractListModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:399
// index:0
// virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QAbstractListModel) Index(row int, column int, parent unsafe.Pointer) {
	// 0: (, row int, column int, parent const QModelIndex &), (&row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractListModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:400
// index:0
// virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QAbstractListModel) Sibling(row int, column int, idx unsafe.Pointer) {
	// 0: (, row int, column int, idx const QModelIndex &), (&row, &column, idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractListModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:401
// index:0
// virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QAbstractListModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, data const QMimeData *, action Qt::DropAction, row int, column int, parent const QModelIndex &), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractListModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:404
// index:0
// virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QAbstractListModel) Flags(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractListModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

//  body block end
