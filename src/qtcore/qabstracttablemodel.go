//  header block begin
// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 64
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

// /usr/include/qt/QtCore/qabstractitemmodel.h:367
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractTableModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTableModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:370
// index:0
// void QAbstractTableModel(class QObject *)
func NewQAbstractTableModel(parent unsafe.Pointer) *QAbstractTableModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTableModelC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractTableModelFromPointer(cthis)
	return gothis
}
func NewQAbstractTableModelFromPointer(cthis unsafe.Pointer) *QAbstractTableModel {
	bcthis0 := NewQAbstractItemModelFromPointer(cthis)
	return &QAbstractTableModel{bcthis0}
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:383
// index:1
// void QAbstractTableModel(class QAbstractItemModelPrivate &, class QObject *)
func NewQAbstractTableModel_1(dd unsafe.Pointer, parent unsafe.Pointer) *QAbstractTableModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTableModelC1ER25QAbstractItemModelPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractTableModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:371
// index:0
// virtual
// void ~QAbstractTableModel()
func DeleteQAbstractTableModel(*QAbstractTableModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTableModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:373
// index:0
// virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QAbstractTableModel) Index(row int, column int, parent unsafe.Pointer) {
	// 0: (, row int, column int, parent const QModelIndex &), (&row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTableModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:374
// index:0
// virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QAbstractTableModel) Sibling(row int, column int, idx unsafe.Pointer) {
	// 0: (, row int, column int, idx const QModelIndex &), (&row, &column, idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTableModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:375
// index:0
// virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QAbstractTableModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, data const QMimeData *, action Qt::DropAction, row int, column int, parent const QModelIndex &), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTableModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:378
// index:0
// virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QAbstractTableModel) Flags(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTableModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

//  body block end
