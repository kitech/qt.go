//  header block begin
// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QModelIndex struct {
	*qtrt.CObject
}

func (this *QModelIndex) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQModelIndexFromPointer(cthis unsafe.Pointer) *QModelIndex {
	return &QModelIndex{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:58
// index:0
// Public inline
// void QModelIndex()
func NewQModelIndex() *QModelIndex {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QModelIndexC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQModelIndexFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:60
// index:0
// Public inline
// int row()
func (this *QModelIndex) Row() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:61
// index:0
// Public inline
// int column()
func (this *QModelIndex) Column() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex6columnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:62
// index:0
// Public inline
// quintptr internalId()
func (this *QModelIndex) InternalId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex10internalIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:63
// index:0
// Public inline
// void * internalPointer()
func (this *QModelIndex) InternalPointer() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex15internalPointerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:64
// index:0
// Public inline
// QModelIndex parent()
func (this *QModelIndex) Parent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:65
// index:0
// Public inline
// QModelIndex sibling(int, int)
func (this *QModelIndex) Sibling(row int, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex7siblingEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:67
// index:0
// Public inline
// QModelIndex child(int, int)
func (this *QModelIndex) Child(row int, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex5childEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:69
// index:0
// Public inline
// QVariant data(int)
func (this *QModelIndex) Data(role int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex4dataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:70
// index:0
// Public inline
// Qt::ItemFlags flags()
func (this *QModelIndex) Flags() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:71
// index:0
// Public inline
// const QAbstractItemModel * model()
func (this *QModelIndex) Model() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:72
// index:0
// Public inline
// bool isValid()
func (this *QModelIndex) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
