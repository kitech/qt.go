package qtcore

// /usr/include/qt/QtCore/qarraydata.h
// #include <qarraydata.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 89
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QArrayData struct {
	*qtrt.CObject
}

func (this *QArrayData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QArrayData) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQArrayDataFromPointer(cthis unsafe.Pointer) *QArrayData {
	return &QArrayData{&qtrt.CObject{cthis}}
}
func (*QArrayData) NewFromPointer(cthis unsafe.Pointer) *QArrayData {
	return NewQArrayDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qarraydata.h:57
// index:0
// Public inline
// void * data()
func (this *QArrayData) Data() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QArrayData4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qarraydata.h:64
// index:1
// Public inline
// const void * data()
func (this *QArrayData) Data_1() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QArrayData4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qarraydata.h:74
// index:0
// Public inline
// bool isMutable()
func (this *QArrayData) IsMutable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QArrayData9isMutableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qarraydata.h:92
// index:0
// Public inline
// size_t detachCapacity(size_t)
func (this *QArrayData) DetachCapacity(newSize uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QArrayData14detachCapacityEm", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newSize)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qarraydata.h:99
// index:0
// Public inline
// QArrayData::AllocationOptions detachFlags()
func (this *QArrayData) DetachFlags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QArrayData11detachFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qarraydata.h:107
// index:0
// Public inline
// QArrayData::AllocationOptions cloneFlags()
func (this *QArrayData) CloneFlags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QArrayData10cloneFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qarraydata.h:123
// index:0
// Public static inline
// QArrayData * sharedNull()
func (this *QArrayData) SharedNull() *QArrayData /*777 QArrayData **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QArrayData10sharedNullEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQArrayDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QArrayData_SharedNull() *QArrayData /*777 QArrayData **/ {
	var nilthis *QArrayData
	rv := nilthis.SharedNull()
	return rv
}

type QArrayData__AllocationOption = int

const QArrayData__CapacityReserved QArrayData__AllocationOption = 1
const QArrayData__Unsharable QArrayData__AllocationOption = 2
const QArrayData__RawData QArrayData__AllocationOption = 4
const QArrayData__Grow QArrayData__AllocationOption = 8
const QArrayData__Default QArrayData__AllocationOption = 0

//  body block end
