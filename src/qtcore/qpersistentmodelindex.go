//  header block begin
// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QPersistentModelIndex struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:107
// index:0
// void QPersistentModelIndex()
func NewQPersistentModelIndex() *QPersistentModelIndex {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QPersistentModelIndexC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QPersistentModelIndex{cthis}
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:108
// index:1
// void QPersistentModelIndex(const class QModelIndex &)
func NewQPersistentModelIndex_1(index unsafe.Pointer) *QPersistentModelIndex {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QPersistentModelIndexC2ERK11QModelIndex", ffiqt.FFI_TYPE_VOID, cthis, index)
	gopp.ErrPrint(err, rv)
	return &QPersistentModelIndex{cthis}
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:110
// index:0
// void ~QPersistentModelIndex()
func DeleteQPersistentModelIndex(*QPersistentModelIndex) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QPersistentModelIndexD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:122
// index:0
// inline
// void swap(class QPersistentModelIndex &)
func (this *QPersistentModelIndex) Swap(other unsafe.Pointer) {
	// 0: (, QPersistentModelIndex & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QPersistentModelIndex4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:127
// index:0
// int row()
func (this *QPersistentModelIndex) Row() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex3rowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:128
// index:0
// int column()
func (this *QPersistentModelIndex) Column() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex6columnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:129
// index:0
// void * internalPointer()
func (this *QPersistentModelIndex) InternalPointer() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex15internalPointerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:130
// index:0
// quintptr internalId()
func (this *QPersistentModelIndex) InternalId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex10internalIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:131
// index:0
// QModelIndex parent()
func (this *QPersistentModelIndex) Parent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex6parentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:132
// index:0
// QModelIndex sibling(int, int)
func (this *QPersistentModelIndex) Sibling(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex7siblingEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:134
// index:0
// QModelIndex child(int, int)
func (this *QPersistentModelIndex) Child(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5childEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:136
// index:0
// QVariant data(int)
func (this *QPersistentModelIndex) Data(role int) {
	// 0: (, int role), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex4dataEi", ffiqt.FFI_TYPE_VOID, this.cthis, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:137
// index:0
// Qt::ItemFlags flags()
func (this *QPersistentModelIndex) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5flagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:138
// index:0
// const QAbstractItemModel * model()
func (this *QPersistentModelIndex) Model() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5modelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:139
// index:0
// bool isValid()
func (this *QPersistentModelIndex) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
