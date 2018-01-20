//  header block begin
// /usr/include/qt/QtCore/qcollator.h
// #include <qcollator.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 106
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
type QCollatorSortKey struct {
	*qtrt.CObject
}

func (this *QCollatorSortKey) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQCollatorSortKeyFromPointer(cthis unsafe.Pointer) *QCollatorSortKey {
	return &QCollatorSortKey{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qcollator.h:58
// index:0
// Public
// void ~QCollatorSortKey()
func DeleteQCollatorSortKey(*QCollatorSortKey) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCollatorSortKeyD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:64
// index:0
// Public inline
// void swap(class QCollatorSortKey &)
func (this *QCollatorSortKey) Swap(other *QCollatorSortKey) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCollatorSortKey4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:67
// index:0
// Public
// int compare(const class QCollatorSortKey &)
func (this *QCollatorSortKey) Compare(key *QCollatorSortKey) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QCollatorSortKey7compareERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcollator.h:70
// index:0
// Protected
// void QCollatorSortKey(class QCollatorSortKeyPrivate *)
func NewQCollatorSortKey(arg0 unsafe.Pointer) *QCollatorSortKey {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCollatorSortKeyC2EP23QCollatorSortKeyPrivate", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCollatorSortKeyFromPointer(cthis)
	return gothis
}

//  body block end
