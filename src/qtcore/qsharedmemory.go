//  header block begin
// /usr/include/qt/QtCore/qsharedmemory.h
// #include <qsharedmemory.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
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
type QSharedMemory struct {
	*QObject
}

func (this *QSharedMemory) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQSharedMemoryFromPointer(cthis unsafe.Pointer) *QSharedMemory {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QSharedMemory{bcthis0}
}

// /usr/include/qt/QtCore/qsharedmemory.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSharedMemory) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:77
// index:0
// Public
// void QSharedMemory(class QObject *)
func NewQSharedMemory(parent unsafe.Pointer) *QSharedMemory {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemoryC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSharedMemoryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsharedmemory.h:78
// index:1
// Public
// void QSharedMemory(const class QString &, class QObject *)
func NewQSharedMemory_1(key *QString, parent unsafe.Pointer) *QSharedMemory {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemoryC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSharedMemoryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsharedmemory.h:79
// index:0
// Public virtual
// void ~QSharedMemory()
func DeleteQSharedMemory(*QSharedMemory) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemoryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:81
// index:0
// Public
// void setKey(const class QString &)
func (this *QSharedMemory) SetKey(key *QString) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory6setKeyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:82
// index:0
// Public
// QString key()
func (this *QSharedMemory) Key() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory3keyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:83
// index:0
// Public
// void setNativeKey(const class QString &)
func (this *QSharedMemory) SetNativeKey(key *QString) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory12setNativeKeyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:84
// index:0
// Public
// QString nativeKey()
func (this *QSharedMemory) NativeKey() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory9nativeKeyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:86
// index:0
// Public
// bool create(int, enum QSharedMemory::AccessMode)
func (this *QSharedMemory) Create(size int, mode int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory6createEiNS_10AccessModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size, &mode)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:87
// index:0
// Public
// int size()
func (this *QSharedMemory) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:89
// index:0
// Public
// bool attach(enum QSharedMemory::AccessMode)
func (this *QSharedMemory) Attach(mode int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory6attachENS_10AccessModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:90
// index:0
// Public
// bool isAttached()
func (this *QSharedMemory) IsAttached() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory10isAttachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:91
// index:0
// Public
// bool detach()
func (this *QSharedMemory) Detach() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory6detachEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:93
// index:0
// Public
// void * data()
func (this *QSharedMemory) Data() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:95
// index:1
// Public
// const void * data()
func (this *QSharedMemory) Data_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:94
// index:0
// Public
// const void * constData()
func (this *QSharedMemory) ConstData() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory9constDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:98
// index:0
// Public
// bool lock()
func (this *QSharedMemory) Lock() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory4lockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:99
// index:0
// Public
// bool unlock()
func (this *QSharedMemory) Unlock() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory6unlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:102
// index:0
// Public
// QSharedMemory::SharedMemoryError error()
func (this *QSharedMemory) Error() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsharedmemory.h:103
// index:0
// Public
// QString errorString()
func (this *QSharedMemory) ErrorString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
