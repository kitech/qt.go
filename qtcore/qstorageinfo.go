package qtcore

// /usr/include/qt/QtCore/qstorageinfo.h
// #include <qstorageinfo.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

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
func (this *QStorageInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStorageInfoFromPointer(cthis unsafe.Pointer) *QStorageInfo {
	return &QStorageInfo{&qtrt.CObject{cthis}}
}
func (*QStorageInfo) NewFromPointer(cthis unsafe.Pointer) *QStorageInfo {
	return NewQStorageInfoFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstorageinfo.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStorageInfo()
func NewQStorageInfo() *QStorageInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfoC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStorageInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStorageInfo)
	return gothis
}

// /usr/include/qt/QtCore/qstorageinfo.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStorageInfo(const QString &)
func NewQStorageInfo_1(path string) *QStorageInfo {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfoC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStorageInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStorageInfo)
	return gothis
}

// /usr/include/qt/QtCore/qstorageinfo.h:60
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QStorageInfo(const QDir &)
func NewQStorageInfo_2(dir *QDir) *QStorageInfo {
	var convArg0 = dir.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfoC2ERK4QDir", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStorageInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStorageInfo)
	return gothis
}

// /usr/include/qt/QtCore/qstorageinfo.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QStorageInfo()
func DeleteQStorageInfo(this *QStorageInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qstorageinfo.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QStorageInfo &)
func (this *QStorageInfo) Swap(other *QStorageInfo) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfo4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &)
func (this *QStorageInfo) SetPath(path string) {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfo7setPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString rootPath()
func (this *QStorageInfo) RootPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo8rootPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstorageinfo.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray device()
func (this *QStorageInfo) Device() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstorageinfo.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray subvolume()
func (this *QStorageInfo) Subvolume() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo9subvolumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstorageinfo.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray fileSystemType()
func (this *QStorageInfo) FileSystemType() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo14fileSystemTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstorageinfo.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QStorageInfo) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstorageinfo.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString displayName()
func (this *QStorageInfo) DisplayName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo11displayNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstorageinfo.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 bytesTotal()
func (this *QStorageInfo) BytesTotal() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo10bytesTotalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstorageinfo.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 bytesFree()
func (this *QStorageInfo) BytesFree() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo9bytesFreeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstorageinfo.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 bytesAvailable()
func (this *QStorageInfo) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstorageinfo.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int blockSize()
func (this *QStorageInfo) BlockSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo9blockSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstorageinfo.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isRoot()
func (this *QStorageInfo) IsRoot() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo6isRootEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly()
func (this *QStorageInfo) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReady()
func (this *QStorageInfo) IsReady() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo7isReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QStorageInfo) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh()
func (this *QStorageInfo) Refresh() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfo7refreshEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStorageInfo root()
func (this *QStorageInfo) Root() *QStorageInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfo4rootEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStorageInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStorageInfo)
	return rv2
}
func QStorageInfo_Root() *QStorageInfo /*123*/ {
	var nilthis *QStorageInfo
	rv := nilthis.Root()
	return rv
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
		qtrt.KeepMe()
	}
}

//  keep block end
