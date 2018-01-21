package qtcore

// /usr/include/qt/QtCore/qstorageinfo.h
// #include <qstorageinfo.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
type QStorageInfo struct {
	*qtrt.CObject
}

func (this *QStorageInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQStorageInfoFromPointer(cthis unsafe.Pointer) *QStorageInfo {
	return &QStorageInfo{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qstorageinfo.h:58
// index:0
// Public
// void QStorageInfo()
func NewQStorageInfo() *QStorageInfo {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfoC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStorageInfoFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstorageinfo.h:59
// index:1
// Public
// void QStorageInfo(const class QString &)
func NewQStorageInfo_1(path *QString) *QStorageInfo {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfoC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStorageInfoFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstorageinfo.h:60
// index:2
// Public
// void QStorageInfo(const class QDir &)
func NewQStorageInfo_2(dir *QDir) *QStorageInfo {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = dir.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfoC2ERK4QDir", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStorageInfoFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstorageinfo.h:62
// index:0
// Public
// void ~QStorageInfo()
func DeleteQStorageInfo(*QStorageInfo) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfoD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:69
// index:0
// Public inline
// void swap(class QStorageInfo &)
func (this *QStorageInfo) Swap(other *QStorageInfo) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfo4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:72
// index:0
// Public
// void setPath(const class QString &)
func (this *QStorageInfo) SetPath(path *QString) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfo7setPathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:74
// index:0
// Public
// QString rootPath()
func (this *QStorageInfo) RootPath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo8rootPathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstorageinfo.h:75
// index:0
// Public
// QByteArray device()
func (this *QStorageInfo) Device() *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstorageinfo.h:76
// index:0
// Public
// QByteArray subvolume()
func (this *QStorageInfo) Subvolume() *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo9subvolumeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstorageinfo.h:77
// index:0
// Public
// QByteArray fileSystemType()
func (this *QStorageInfo) FileSystemType() *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo14fileSystemTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstorageinfo.h:78
// index:0
// Public
// QString name()
func (this *QStorageInfo) Name() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstorageinfo.h:79
// index:0
// Public
// QString displayName()
func (this *QStorageInfo) DisplayName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo11displayNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstorageinfo.h:81
// index:0
// Public
// qint64 bytesTotal()
func (this *QStorageInfo) BytesTotal() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo10bytesTotalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstorageinfo.h:82
// index:0
// Public
// qint64 bytesFree()
func (this *QStorageInfo) BytesFree() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo9bytesFreeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstorageinfo.h:83
// index:0
// Public
// qint64 bytesAvailable()
func (this *QStorageInfo) BytesAvailable() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo14bytesAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstorageinfo.h:84
// index:0
// Public
// int blockSize()
func (this *QStorageInfo) BlockSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo9blockSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstorageinfo.h:86
// index:0
// Public inline
// bool isRoot()
func (this *QStorageInfo) IsRoot() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo6isRootEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:87
// index:0
// Public
// bool isReadOnly()
func (this *QStorageInfo) IsReadOnly() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo10isReadOnlyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:88
// index:0
// Public
// bool isReady()
func (this *QStorageInfo) IsReady() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo7isReadyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:89
// index:0
// Public
// bool isValid()
func (this *QStorageInfo) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStorageInfo7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:91
// index:0
// Public
// void refresh()
func (this *QStorageInfo) Refresh() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfo7refreshEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:94
// index:0
// Public static
// QStorageInfo root()
func (this *QStorageInfo) Root() *QStorageInfo /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStorageInfo4rootEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStorageInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QStorageInfo_Root() *QStorageInfo /*123*/ {
	var nilthis *QStorageInfo
	rv := nilthis.Root()
	return rv
}

//  body block end
