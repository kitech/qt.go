package qtcore

// /usr/include/qt/QtCore/qsharedmemory.h
// #include <qsharedmemory.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QSharedMemory struct {
	*QObject
}

func (this *QSharedMemory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSharedMemory) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQSharedMemoryFromPointer(cthis unsafe.Pointer) *QSharedMemory {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QSharedMemory{bcthis0}
}
func (*QSharedMemory) NewFromPointer(cthis unsafe.Pointer) *QSharedMemory {
	return NewQSharedMemoryFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsharedmemory.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSharedMemory) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsharedmemory.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSharedMemory(QObject *)
func NewQSharedMemory(parent *QObject /*777 QObject **/) *QSharedMemory {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemoryC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSharedMemoryFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsharedmemory.h:78
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSharedMemory(const QString &, QObject *)
func NewQSharedMemory_1(key string, parent *QObject /*777 QObject **/) *QSharedMemory {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemoryC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSharedMemoryFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsharedmemory.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSharedMemory()
func DeleteQSharedMemory(this *QSharedMemory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsharedmemory.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QString &)
func (this *QSharedMemory) SetKey(key string) {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6setKeyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString key()
func (this *QSharedMemory) Key() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory3keyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsharedmemory.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNativeKey(const QString &)
func (this *QSharedMemory) SetNativeKey(key string) {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory12setNativeKeyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QString nativeKey()
func (this *QSharedMemory) NativeKey() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory9nativeKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsharedmemory.h:86
// index:0
// Public Visibility=Default Availability=Available
// [1] bool create(int, enum QSharedMemory::AccessMode)
func (this *QSharedMemory) Create(size int, mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6createEiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int size()
func (this *QSharedMemory) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsharedmemory.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool attach(enum QSharedMemory::AccessMode)
func (this *QSharedMemory) Attach(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6attachENS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAttached()
func (this *QSharedMemory) IsAttached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory10isAttachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool detach()
func (this *QSharedMemory) Detach() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] void * data()
func (this *QSharedMemory) Data() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qsharedmemory.h:95
// index:1
// Public Visibility=Default Availability=Available
// [8] const void * data()
func (this *QSharedMemory) Data_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qsharedmemory.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] const void * constData()
func (this *QSharedMemory) ConstData() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory9constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qsharedmemory.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool lock()
func (this *QSharedMemory) Lock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory4lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool unlock()
func (this *QSharedMemory) Unlock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] QSharedMemory::SharedMemoryError error()
func (this *QSharedMemory) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QSharedMemory) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

type QSharedMemory__AccessMode = int

const QSharedMemory__ReadOnly QSharedMemory__AccessMode = 0
const QSharedMemory__ReadWrite QSharedMemory__AccessMode = 1

type QSharedMemory__SharedMemoryError = int

const QSharedMemory__NoError QSharedMemory__SharedMemoryError = 0
const QSharedMemory__PermissionDenied QSharedMemory__SharedMemoryError = 1
const QSharedMemory__InvalidSize QSharedMemory__SharedMemoryError = 2
const QSharedMemory__KeyError QSharedMemory__SharedMemoryError = 3
const QSharedMemory__AlreadyExists QSharedMemory__SharedMemoryError = 4
const QSharedMemory__NotFound QSharedMemory__SharedMemoryError = 5
const QSharedMemory__LockError QSharedMemory__SharedMemoryError = 6
const QSharedMemory__OutOfResources QSharedMemory__SharedMemoryError = 7
const QSharedMemory__UnknownError QSharedMemory__SharedMemoryError = 8

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
		qtrt.KeepMe()
	}
}

//  keep block end
