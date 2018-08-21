package qtcore

// /usr/include/qt/QtCore/qarraydata.h
// #include <qarraydata.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 89
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QArrayData struct {
	*qtrt.CObject
}
type QArrayData_ITF interface {
	QArrayData_PTR() *QArrayData
}

func (ptr *QArrayData) QArrayData_PTR() *QArrayData { return ptr }

func (this *QArrayData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QArrayData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQArrayDataFromPointer(cthis unsafe.Pointer) *QArrayData {
	return &QArrayData{&qtrt.CObject{cthis}}
}
func (*QArrayData) NewFromPointer(cthis unsafe.Pointer) *QArrayData {
	return NewQArrayDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qarraydata.h:57
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void * data()

/*

 */
func (this *QArrayData) Data() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QArrayData4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qarraydata.h:64
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const void * data() const

/*

 */
func (this *QArrayData) Data_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QArrayData4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qarraydata.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isMutable() const

/*

 */
func (this *QArrayData) IsMutable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QArrayData9isMutableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qarraydata.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [8] size_t detachCapacity(size_t) const

/*

 */
func (this *QArrayData) DetachCapacity(newSize uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QArrayData14detachCapacityEm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newSize)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qarraydata.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QArrayData::AllocationOptions detachFlags() const

/*

 */
func (this *QArrayData) DetachFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QArrayData11detachFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qarraydata.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QArrayData::AllocationOptions cloneFlags() const

/*

 */
func (this *QArrayData) CloneFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QArrayData10cloneFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qarraydata.h:115
// index:0
// Public static Visibility=Default Availability=Available
// [8] QArrayData * allocate(size_t, size_t, size_t, QArrayData::AllocationOptions)

/*

 */
func (this *QArrayData) Allocate(objectSize uint, alignment uint, capacity uint, options int) *QArrayData /*777 QArrayData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QArrayData8allocateEmmm6QFlagsINS_16AllocationOptionEE", qtrt.FFI_TYPE_POINTER, objectSize, alignment, capacity, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQArrayDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QArrayData_Allocate(objectSize uint, alignment uint, capacity uint, options int) *QArrayData /*777 QArrayData **/ {
	var nilthis *QArrayData
	rv := nilthis.Allocate(objectSize, alignment, capacity, options)
	return rv
}

// /usr/include/qt/QtCore/qarraydata.h:115
// index:0
// Public static Visibility=Default Availability=Available
// [8] QArrayData * allocate(size_t, size_t, size_t, QArrayData::AllocationOptions)

/*

 */
func (this *QArrayData) Allocate__(objectSize uint, alignment uint, capacity uint) *QArrayData /*777 QArrayData **/ {
	// arg: 3, QArrayData::AllocationOptions=Typedef, QArrayData::AllocationOptions=Typedef, QFlags<QArrayData::AllocationOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QArrayData8allocateEmmm6QFlagsINS_16AllocationOptionEE", qtrt.FFI_TYPE_POINTER, objectSize, alignment, capacity, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQArrayDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qarraydata.h:123
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QArrayData * sharedNull()

/*

 */
func (this *QArrayData) SharedNull() *QArrayData /*777 QArrayData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QArrayData10sharedNullEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQArrayDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QArrayData_SharedNull() *QArrayData /*777 QArrayData **/ {
	var nilthis *QArrayData
	rv := nilthis.SharedNull()
	return rv
}

func DeleteQArrayData(this *QArrayData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QArrayDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QArrayData__AllocationOption = int

//
const QArrayData__CapacityReserved QArrayData__AllocationOption = 1

//
const QArrayData__Unsharable QArrayData__AllocationOption = 2

//
const QArrayData__RawData QArrayData__AllocationOption = 4

//
const QArrayData__Grow QArrayData__AllocationOption = 8

//
const QArrayData__Default QArrayData__AllocationOption = 0

func (this *QArrayData) AllocationOptionItemName(val int) string {
	switch val {
	case QArrayData__CapacityReserved: // 1
		return "CapacityReserved"
	case QArrayData__Unsharable: // 2
		return "Unsharable"
	case QArrayData__RawData: // 4
		return "RawData"
	case QArrayData__Grow: // 8
		return "Grow"
	case QArrayData__Default: // 0
		return "Default"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QArrayData_AllocationOptionItemName(val int) string {
	var nilthis *QArrayData
	return nilthis.AllocationOptionItemName(val)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
