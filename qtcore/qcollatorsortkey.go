package qtcore

// /usr/include/qt/QtCore/qcollator.h
// #include <qcollator.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 102
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
type QCollatorSortKey struct {
	*qtrt.CObject
}

func (this *QCollatorSortKey) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCollatorSortKey) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQCollatorSortKeyFromPointer(cthis unsafe.Pointer) *QCollatorSortKey {
	return &QCollatorSortKey{&qtrt.CObject{cthis}}
}
func (*QCollatorSortKey) NewFromPointer(cthis unsafe.Pointer) *QCollatorSortKey {
	return NewQCollatorSortKeyFromPointer(cthis)
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
// void swap(QCollatorSortKey &)
func (this *QCollatorSortKey) Swap(other *QCollatorSortKey) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCollatorSortKey4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:67
// index:0
// Public
// int compare(const QCollatorSortKey &)
func (this *QCollatorSortKey) Compare(key *QCollatorSortKey) int {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QCollatorSortKey7compareERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
