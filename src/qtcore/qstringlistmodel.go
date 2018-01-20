//  header block begin
// /usr/include/qt/QtCore/qstringlistmodel.h
// #include <qstringlistmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QStringListModel struct {
	*QAbstractListModel
}

func (this *QStringListModel) GetCthis() unsafe.Pointer {
	return this.QAbstractListModel.GetCthis()
}

// /usr/include/qt/QtCore/qstringlistmodel.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QStringListModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:55
// index:0
// void QStringListModel(class QObject *)
func NewQStringListModel(parent unsafe.Pointer) *QStringListModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringListModelFromPointer(cthis)
	return gothis
}
func NewQStringListModelFromPointer(cthis unsafe.Pointer) *QStringListModel {
	bcthis0 := NewQAbstractListModelFromPointer(cthis)
	return &QStringListModel{bcthis0}
}

// /usr/include/qt/QtCore/qstringlistmodel.h:56
// index:1
// void QStringListModel(const class QStringList &, class QObject *)
func NewQStringListModel_1(strings unsafe.Pointer, parent unsafe.Pointer) *QStringListModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModelC2ERK11QStringListP7QObject", ffiqt.FFI_TYPE_VOID, cthis, strings, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringListModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstringlistmodel.h:58
// index:0
// virtual
// int rowCount(const class QModelIndex &)
func (this *QStringListModel) RowCount(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:59
// index:0
// virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QStringListModel) Sibling(row int, column int, idx unsafe.Pointer) {
	// 0: (, row int, column int, idx const QModelIndex &), (&row, &column, idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:61
// index:0
// virtual
// QVariant data(const class QModelIndex &, int)
func (this *QStringListModel) Data(index unsafe.Pointer, role int) {
	// 0: (, index const QModelIndex &, role int), (index, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:62
// index:0
// virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QStringListModel) SetData(index unsafe.Pointer, value unsafe.Pointer, role int) {
	// 0: (, index const QModelIndex &, value const QVariant &, role int), (index, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:64
// index:0
// virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QStringListModel) Flags(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:66
// index:0
// virtual
// bool insertRows(int, int, const class QModelIndex &)
func (this *QStringListModel) InsertRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, row int, count int, parent const QModelIndex &), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModel10insertRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:67
// index:0
// virtual
// bool removeRows(int, int, const class QModelIndex &)
func (this *QStringListModel) RemoveRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, row int, count int, parent const QModelIndex &), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModel10removeRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:69
// index:0
// virtual
// void sort(int, Qt::SortOrder)
func (this *QStringListModel) Sort(column int, order int) {
	// 0: (, column int, order Qt::SortOrder), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:71
// index:0
// QStringList stringList()
func (this *QStringListModel) StringList() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel10stringListEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:72
// index:0
// void setStringList(const class QStringList &)
func (this *QStringListModel) SetStringList(strings unsafe.Pointer) {
	// 0: (, strings const QStringList &), (strings)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModel13setStringListERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), strings)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:74
// index:0
// virtual
// Qt::DropActions supportedDropActions()
func (this *QStringListModel) SupportedDropActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel20supportedDropActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
