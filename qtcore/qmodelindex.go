package qtcore

// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QModelIndex) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQModelIndexFromPointer(cthis unsafe.Pointer) *QModelIndex {
	return &QModelIndex{&qtrt.CObject{cthis}}
}
func (*QModelIndex) NewFromPointer(cthis unsafe.Pointer) *QModelIndex {
	return NewQModelIndexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QModelIndex()
func NewQModelIndex() *QModelIndex {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QModelIndexC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int row()
func (this *QModelIndex) Row() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int column()
func (this *QModelIndex) Column() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex6columnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [8] quintptr internalId()
func (this *QModelIndex) InternalId() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex10internalIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void * internalPointer()
func (this *QModelIndex) InternalPointer() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex15internalPointerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex parent()
func (this *QModelIndex) Parent() *QModelIndex /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int)
func (this *QModelIndex) Sibling(row int, column int) *QModelIndex /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex7siblingEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex child(int, int)
func (this *QModelIndex) Child(row int, column int) *QModelIndex /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex5childEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QVariant data(int)
func (this *QModelIndex) Data(role int) *QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex4dataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags()
func (this *QModelIndex) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QAbstractItemModel * model()
func (this *QModelIndex) Model() *QAbstractItemModel /*777 const QAbstractItemModel **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QModelIndex) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QModelIndex7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
