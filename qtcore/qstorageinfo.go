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
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QStorageInfo struct {
	*qtrt.CObject
}
type QStorageInfo_ITF interface {
	QStorageInfo_PTR() *QStorageInfo
}

func (ptr *QStorageInfo) QStorageInfo_PTR() *QStorageInfo { return ptr }

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

/*
Constructs an empty QStorageInfo object.

Objects created with the default constructor will be invalid and therefore not ready for use.

See also setPath(), isReady(), and isValid().
*/
func (*QStorageInfo) NewForInherit() *QStorageInfo {
	return NewQStorageInfo()
}
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

/*
Constructs an empty QStorageInfo object.

Objects created with the default constructor will be invalid and therefore not ready for use.

See also setPath(), isReady(), and isValid().
*/
func (*QStorageInfo) NewForInherit1(path string) *QStorageInfo {
	return NewQStorageInfo1(path)
}
func NewQStorageInfo1(path string) *QStorageInfo {
	var tmpArg0 = NewQString5(path)
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

/*
Constructs an empty QStorageInfo object.

Objects created with the default constructor will be invalid and therefore not ready for use.

See also setPath(), isReady(), and isValid().
*/
func (*QStorageInfo) NewForInherit2(dir QDir_ITF) *QStorageInfo {
	return NewQStorageInfo2(dir)
}
func NewQStorageInfo2(dir QDir_ITF) *QStorageInfo {
	var convArg0 unsafe.Pointer
	if dir != nil && dir.QDir_PTR() != nil {
		convArg0 = dir.QDir_PTR().GetCthis()
	}
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

/*

 */
func DeleteQStorageInfo(this *QStorageInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qstorageinfo.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QStorageInfo & operator=(const QStorageInfo &)

/*

 */
func (this *QStorageInfo) Operator_equal(other QStorageInfo_ITF) *QStorageInfo {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStorageInfo_PTR() != nil {
		convArg0 = other.QStorageInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfoaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStorageInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStorageInfo)
	return rv2
}

// /usr/include/qt/QtCore/qstorageinfo.h:66
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QStorageInfo & operator=(QStorageInfo &&)

/*

 */
func (this *QStorageInfo) Operator_equal1(other unsafe.Pointer /*333*/) *QStorageInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfoaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStorageInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStorageInfo)
	return rv2
}

// /usr/include/qt/QtCore/qstorageinfo.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QStorageInfo &)

/*
Swaps this volume info with other. This function is very fast and never fails.
*/
func (this *QStorageInfo) Swap(other QStorageInfo_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStorageInfo_PTR() != nil {
		convArg0 = other.QStorageInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfo4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &)

/*
Sets this QStorageInfo object to the filesystem mounted where path is located.

path can either be a root path of the filesystem, a directory, or a file within that filesystem.

See also rootPath().
*/
func (this *QStorageInfo) SetPath(path string) {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfo7setPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString rootPath() const

/*
Returns the mount point of the filesystem this QStorageInfo object represents.

On Windows, it returns the volume letter in case the volume is not mounted to a directory.

Note that the value returned by rootPath() is the real mount point of a volume, and may not be equal to the value passed to the constructor or setPath() method. For example, if you have only the root volume in the system, and pass '/directory' to setPath(), then this method will return '/'.

See also setPath() and device().
*/
func (this *QStorageInfo) RootPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo8rootPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstorageinfo.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray device() const

/*
Returns the device for this volume.

For example, on Unix filesystems (including macOS), this returns the devpath like /dev/sda0 for local storages. On Windows, it returns the UNC path starting with \\\\?\\ for local storages (in other words, the volume GUID).

See also rootPath() and subvolume().
*/
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
// [8] QByteArray subvolume() const

/*
Returns the subvolume name for this volume.

Some filesystem types allow multiple subvolumes inside one device, which may be mounted in different paths. If the subvolume could be detected, it is returned here. The format of the subvolume name is specific to each filesystem type.

If this volume was not mounted from a subvolume of a larger filesystem or if the subvolume could not be detected, this function returns an empty byte array.

This function was introduced in  Qt 5.9.

See also device().
*/
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
// [8] QByteArray fileSystemType() const

/*
Returns the type name of the filesystem.

This is a platform-dependent function, and filesystem names can vary between different operating systems. For example, on Windows filesystems they can be named NTFS, and on Linux they can be named ntfs-3g or fuseblk.

See also name().
*/
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
// [8] QString name() const

/*
Returns the human-readable name of a filesystem, usually called label.

Not all filesystems support this feature. In this case, the value returned by this method could be empty. An empty string is returned if the file system does not support labels, or if no label is set.

On Linux, retrieving the volume's label requires udev to be present in the system.

See also fileSystemType().
*/
func (this *QStorageInfo) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstorageinfo.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString displayName() const

/*
Returns the volume's name, if available, or the root path if not.
*/
func (this *QStorageInfo) DisplayName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo11displayNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstorageinfo.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 bytesTotal() const

/*
Returns the total volume size in bytes.

Returns -1 if QStorageInfo object is not valid.

See also bytesFree() and bytesAvailable().
*/
func (this *QStorageInfo) BytesTotal() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo10bytesTotalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstorageinfo.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 bytesFree() const

/*
Returns the number of free bytes in a volume. Note that if there are quotas on the filesystem, this value can be larger than the value returned by bytesAvailable().

Returns -1 if QStorageInfo object is not valid.

See also bytesTotal() and bytesAvailable().
*/
func (this *QStorageInfo) BytesFree() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo9bytesFreeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstorageinfo.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 bytesAvailable() const

/*
Returns the size (in bytes) available for the current user. It returns the total size available if the user is the root user or a system administrator.

This size can be less than or equal to the free size returned by bytesFree() function.

Returns -1 if QStorageInfo object is not valid.

See also bytesTotal() and bytesFree().
*/
func (this *QStorageInfo) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstorageinfo.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int blockSize() const

/*
Returns the optimal transfer block size for this filesystem.

Returns -1 if QStorageInfo could not determine the size or if the QStorageInfo object is not valid.

This function was introduced in  Qt 5.6.
*/
func (this *QStorageInfo) BlockSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo9blockSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstorageinfo.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isRoot() const

/*
Returns true if this QStorageInfo represents the system root volume; false otherwise.

On Unix filesystems, the root volume is a volume mounted on /. On Windows, the root volume is the volume where the OS is installed.

See also root().
*/
func (this *QStorageInfo) IsRoot() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo6isRootEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly() const

/*
Returns true if the current filesystem is protected from writing; false otherwise.
*/
func (this *QStorageInfo) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReady() const

/*
Returns true if the current filesystem is ready to work; false otherwise. For example, false is returned if the CD volume is not inserted.

Note that fileSystemType(), name(), bytesTotal(), bytesFree(), and bytesAvailable() will return invalid data until the volume is ready.

See also isValid().
*/
func (this *QStorageInfo) IsReady() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo7isReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the QStorageInfo specified by rootPath exists and is mounted correctly.

See also isReady().
*/
func (this *QStorageInfo) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStorageInfo7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstorageinfo.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh()

/*
Resets QStorageInfo's internal cache.

QStorageInfo caches information about storage to speed up performance. QStorageInfo retrieves information during object construction and/or when calling the setPath() method. You have to manually reset the cache by calling this function to update storage information.
*/
func (this *QStorageInfo) Refresh() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStorageInfo7refreshEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstorageinfo.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStorageInfo root()

/*
Returns a QStorageInfo object that represents the system root volume.

On Unix systems this call returns the root ('/') volume; in Windows the volume where the operating system is installed.

See also isRoot().
*/
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

func init_unused_10555() {
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
