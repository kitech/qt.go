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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qcollator.h:58
// index:0
// void ~QCollatorSortKey()
func DeleteQCollatorSortKey(*QCollatorSortKey) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCollatorSortKeyD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:64
// index:0
// inline
// void swap(class QCollatorSortKey &)
func (this *QCollatorSortKey) Swap(other unsafe.Pointer) {
	// 0: (, QCollatorSortKey & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCollatorSortKey4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:67
// index:0
// int compare(const class QCollatorSortKey &)
func (this *QCollatorSortKey) Compare(key unsafe.Pointer) {
	// 0: (, const QCollatorSortKey & key), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QCollatorSortKey7compareERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, key)
	gopp.ErrPrint(err, rv)
}

//  body block end
