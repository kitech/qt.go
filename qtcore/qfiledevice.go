package qtcore

// /usr/include/qt/QtCore/qfiledevice.h
// #include <qfiledevice.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// long long readData(char *, qint64)
func (this *QFileDevice) InheritReadData(f func(data string, maxlen int64) int64) {
	qtrt.SetAllInheritCallback(this, "readData", f)
}

// long long writeData(const char *, qint64)
func (this *QFileDevice) InheritWriteData(f func(data string, len_ int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

// long long readLineData(char *, qint64)
func (this *QFileDevice) InheritReadLineData(f func(data string, maxlen int64) int64) {
	qtrt.SetAllInheritCallback(this, "readLineData", f)
}

/*

 */
type QFileDevice struct {
	*QIODevice
}
type QFileDevice_ITF interface {
	QIODevice_ITF
	QFileDevice_PTR() *QFileDevice
}

func (ptr *QFileDevice) QFileDevice_PTR() *QFileDevice { return ptr }

func (this *QFileDevice) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QIODevice.GetCthis()
	}
}
func (this *QFileDevice) SetCthis(cthis unsafe.Pointer) {
	this.QIODevice = NewQIODeviceFromPointer(cthis)
}
func NewQFileDeviceFromPointer(cthis unsafe.Pointer) *QFileDevice {
	bcthis0 := NewQIODeviceFromPointer(cthis)
	return &QFileDevice{bcthis0}
}
func (*QFileDevice) NewFromPointer(cthis unsafe.Pointer) *QFileDevice {
	return NewQFileDeviceFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfiledevice.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QFileDevice) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDevice10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qfiledevice.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFileDevice()

/*

 */
func DeleteQFileDevice(this *QFileDevice) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDeviceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfiledevice.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] QFileDevice::FileError error() const

/*
Returns the file error status.

The I/O device status returns an error code. For example, if open() returns false, or a read/write operation returns -1, this function can be called to find out the reason why the operation failed.

See also unsetError().
*/
func (this *QFileDevice) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDevice5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qfiledevice.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetError()

/*
Sets the file's error to QFileDevice::NoError.

See also error().
*/
func (this *QFileDevice) UnsetError() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice10unsetErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfiledevice.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()

/*
Reimplemented from QIODevice::close().

Calls QFileDevice::flush() and closes the file. Errors from flush are ignored.

See also QIODevice::close().
*/
func (this *QFileDevice) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfiledevice.h:105
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSequential() const

/*
Reimplemented from QIODevice::isSequential().

Returns true if the file can only be manipulated sequentially; otherwise returns false.

Most files support random-access, but some special files may not.

See also QIODevice::isSequential().
*/
func (this *QFileDevice) IsSequential() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDevice12isSequentialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfiledevice.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] int handle() const

/*
Returns the file handle of the file.

This is a small positive integer, suitable for use with C library functions such as fdopen() and fcntl(). On systems that use file descriptors for sockets (i.e. Unix systems, but not Windows) the handle can be used with QSocketNotifier as well.

If the file is not open, or there is an error, handle() returns -1.

See also QSocketNotifier.
*/
func (this *QFileDevice) Handle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDevice6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfiledevice.h:108
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString fileName() const

/*
Returns the name of the file. The default implementation in QFileDevice returns a null string.
*/
func (this *QFileDevice) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDevice8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfiledevice.h:110
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 pos() const

/*
Reimplemented from QIODevice::pos().
*/
func (this *QFileDevice) Pos() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDevice3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qfiledevice.h:111
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool seek(qint64)

/*
Reimplemented from QIODevice::seek().

For random-access devices, this function sets the current position to pos, returning true on success, or false if an error occurred. For sequential devices, the default behavior is to do nothing and return false.

Seeking beyond the end of a file: If the position is beyond the end of a file, then seek() will not immediately extend the file. If a write is performed at this position, then the file will be extended. The content of the file between the previous end of file and the newly written data is UNDEFINED and varies between platforms and file systems.
*/
func (this *QFileDevice) Seek(offset int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice4seekEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfiledevice.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd() const

/*
Reimplemented from QIODevice::atEnd().

Returns true if the end of the file has been reached; otherwise returns false.

For regular empty files on Unix (e.g. those in /proc), this function returns true, since the file system reports that the size of such a file is 0. Therefore, you should not depend on atEnd() when reading data from such a file, but rather call read() until no more data can be read.
*/
func (this *QFileDevice) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDevice5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfiledevice.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool flush()

/*
Flushes any buffered data to the file. Returns true if successful; otherwise returns false.
*/
func (this *QFileDevice) Flush() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfiledevice.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 size() const

/*
Reimplemented from QIODevice::size().

Returns the size of the file.

For regular empty files on Unix (e.g. those in /proc), this function returns 0; the contents of such a file are generated on demand in response to you calling read().
*/
func (this *QFileDevice) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDevice4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qfiledevice.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool resize(qint64)

/*
Sets the file size (in bytes) sz. Returns true if the resize succeeds; false otherwise. If sz is larger than the file currently is, the new bytes will be set to 0; if sz is smaller, the file is simply truncated.

Warning: This function can fail if the file doesn't exist.

See also size().
*/
func (this *QFileDevice) Resize(sz int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice6resizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sz)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfiledevice.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QFileDevice::Permissions permissions() const

/*
Returns the complete OR-ed together combination of QFile::Permission for the file.

See also setPermissions().
*/
func (this *QFileDevice) Permissions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDevice11permissionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qfiledevice.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setPermissions(QFileDevice::Permissions)

/*
Sets the permissions for the file to the permissions specified. Returns true if successful, or false if the permissions cannot be modified.

Warning: This function does not manipulate ACLs, which may limit its effectiveness.

See also permissions().
*/
func (this *QFileDevice) SetPermissions(permissionSpec int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice14setPermissionsE6QFlagsINS_10PermissionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), permissionSpec)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfiledevice.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] uchar * map(qint64, qint64, QFileDevice::MemoryMapFlags)

/*
Maps size bytes of the file into memory starting at offset. A file should be open for a map to succeed but the file does not need to stay open after the memory has been mapped. When the QFile is destroyed or a new file is opened with this object, any maps that have not been unmapped will automatically be unmapped.

The mapping will have the same open mode as the file (read and/or write), except when using MapPrivateOption, in which case it is always possible to write to the mapped memory.

Any mapping options can be passed through flags.

Returns a pointer to the memory or 0 if there is an error.

See also unmap().
*/
func (this *QFileDevice) Map(offset int64, size int64, flags int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice3mapExxNS_14MemoryMapFlagsE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset, size, flags)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qfiledevice.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] uchar * map(qint64, qint64, QFileDevice::MemoryMapFlags)

/*
Maps size bytes of the file into memory starting at offset. A file should be open for a map to succeed but the file does not need to stay open after the memory has been mapped. When the QFile is destroyed or a new file is opened with this object, any maps that have not been unmapped will automatically be unmapped.

The mapping will have the same open mode as the file (read and/or write), except when using MapPrivateOption, in which case it is always possible to write to the mapped memory.

Any mapping options can be passed through flags.

Returns a pointer to the memory or 0 if there is an error.

See also unmap().
*/
func (this *QFileDevice) Map__(offset int64, size int64) unsafe.Pointer /*666*/ {
	// arg: 2, QFileDevice::MemoryMapFlags=Enum, QFileDevice::MemoryMapFlags=Enum, , Invalid
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice3mapExxNS_14MemoryMapFlagsE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset, size, flags)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qfiledevice.h:128
// index:0
// Public Visibility=Default Availability=Available
// [1] bool unmap(uchar *)

/*
Unmaps the memory address.

Returns true if the unmap succeeds; false otherwise.

See also map().
*/
func (this *QFileDevice) Unmap(address unsafe.Pointer /*666*/) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice5unmapEPh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), address)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfiledevice.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime fileTime(QFileDevice::FileTime) const

/*
Returns the file time specified by time. If the time cannot be determined return QDateTime() (an invalid date time).

This function was introduced in  Qt 5.10.

See also setFileTime(), FileTime, and QDateTime::isValid().
*/
func (this *QFileDevice) FileTime(time int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDevice8fileTimeENS_8FileTimeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), time)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qfiledevice.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setFileTime(const QDateTime &, QFileDevice::FileTime)

/*
Sets the file time specified by fileTime to newDate, returning true if successful; otherwise returns false.

Note: The file must be open to use this function.

This function was introduced in  Qt 5.10.

See also fileTime() and FileTime.
*/
func (this *QFileDevice) SetFileTime(newDate QDateTime_ITF, fileTime int) bool {
	var convArg0 unsafe.Pointer
	if newDate != nil && newDate.QDateTime_PTR() != nil {
		convArg0 = newDate.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice11setFileTimeERK9QDateTimeNS_8FileTimeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fileTime)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfiledevice.h:134
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QFileDevice()

/*

 */
func NewQFileDevice() *QFileDevice {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDeviceC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileDeviceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileDevice")
	return gothis
}

// /usr/include/qt/QtCore/qfiledevice.h:138
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QFileDevice(QObject *)

/*

 */
func NewQFileDevice_1(parent QObject_ITF /*777 QObject **/) *QFileDevice {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDeviceC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileDeviceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileDevice")
	return gothis
}

// /usr/include/qt/QtCore/qfiledevice.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)

/*
Reimplemented from QIODevice::readData().
*/
func (this *QFileDevice) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice8readDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qfiledevice.h:143
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)

/*
Reimplemented from QIODevice::writeData().
*/
func (this *QFileDevice) WriteData(data string, len_ int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qfiledevice.h:144
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readLineData(char *, qint64)

/*
Reimplemented from QIODevice::readLineData().
*/
func (this *QFileDevice) ReadLineData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDevice12readLineDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

/*
This enum describes the errors that may be returned by the error() function.


*/
type QFileDevice__FileError = int

// No error occurred.
const QFileDevice__NoError QFileDevice__FileError = 0

// An error occurred when reading from the file.
const QFileDevice__ReadError QFileDevice__FileError = 1

// An error occurred when writing to the file.
const QFileDevice__WriteError QFileDevice__FileError = 2

// A fatal error occurred.
const QFileDevice__FatalError QFileDevice__FileError = 3

// Out of resources (e.g., too many open files, out of memory, etc.)
const QFileDevice__ResourceError QFileDevice__FileError = 4

// The file could not be opened.
const QFileDevice__OpenError QFileDevice__FileError = 5

// The operation was aborted.
const QFileDevice__AbortError QFileDevice__FileError = 6

// A timeout occurred.
const QFileDevice__TimeOutError QFileDevice__FileError = 7

// An unspecified error occurred.
const QFileDevice__UnspecifiedError QFileDevice__FileError = 8

// The file could not be removed.
const QFileDevice__RemoveError QFileDevice__FileError = 9

//
const QFileDevice__RenameError QFileDevice__FileError = 10

//
const QFileDevice__PositionError QFileDevice__FileError = 11

//
const QFileDevice__ResizeError QFileDevice__FileError = 12

//
const QFileDevice__PermissionsError QFileDevice__FileError = 13

//
const QFileDevice__CopyError QFileDevice__FileError = 14

func (this *QFileDevice) FileErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFileDevice_FileErrorItemName(val int) string {
	var nilthis *QFileDevice
	return nilthis.FileErrorItemName(val)
}

/*
This enum is used by the fileTime() and setFileTime() functions.



This enum was introduced or modified in  Qt 5.10.

See also setFileTime(), fileTime(), and QFileInfo::fileTime().

*/
type QFileDevice__FileTime = int

// When the file was most recently accessed (e.g. read or written to).
const QFileDevice__FileAccessTime QFileDevice__FileTime = 0

// When the file was created (may not be not supported on UNIX).
const QFileDevice__FileBirthTime QFileDevice__FileTime = 1

// When the file's metadata was last changed.
const QFileDevice__FileMetadataChangeTime QFileDevice__FileTime = 2

// When the file was most recently modified.
const QFileDevice__FileModificationTime QFileDevice__FileTime = 3

func (this *QFileDevice) FileTimeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFileDevice_FileTimeItemName(val int) string {
	var nilthis *QFileDevice
	return nilthis.FileTimeItemName(val)
}

/*


 */
type QFileDevice__Permission = int

//
const QFileDevice__ReadOwner QFileDevice__Permission = 16384

//
const QFileDevice__WriteOwner QFileDevice__Permission = 8192

//
const QFileDevice__ExeOwner QFileDevice__Permission = 4096

//
const QFileDevice__ReadUser QFileDevice__Permission = 1024

//
const QFileDevice__WriteUser QFileDevice__Permission = 512

//
const QFileDevice__ExeUser QFileDevice__Permission = 256

//
const QFileDevice__ReadGroup QFileDevice__Permission = 64

//
const QFileDevice__WriteGroup QFileDevice__Permission = 32

//
const QFileDevice__ExeGroup QFileDevice__Permission = 16

//
const QFileDevice__ReadOther QFileDevice__Permission = 4

//
const QFileDevice__WriteOther QFileDevice__Permission = 2

//
const QFileDevice__ExeOther QFileDevice__Permission = 1

func (this *QFileDevice) PermissionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFileDevice_PermissionItemName(val int) string {
	var nilthis *QFileDevice
	return nilthis.PermissionItemName(val)
}

/*


 */
type QFileDevice__FileHandleFlag = int

//
const QFileDevice__AutoCloseHandle QFileDevice__FileHandleFlag = 1

//
const QFileDevice__DontCloseHandle QFileDevice__FileHandleFlag = 0

func (this *QFileDevice) FileHandleFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFileDevice_FileHandleFlagItemName(val int) string {
	var nilthis *QFileDevice
	return nilthis.FileHandleFlagItemName(val)
}

/*
This enum describes special options that may be used by the map() function.



This enum was introduced or modified in  Qt 4.4.

*/
type QFileDevice__MemoryMapFlags = int

// No options.
const QFileDevice__NoOptions QFileDevice__MemoryMapFlags = 0

//
const QFileDevice__MapPrivateOption QFileDevice__MemoryMapFlags = 1

func (this *QFileDevice) MemoryMapFlagsItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFileDevice_MemoryMapFlagsItemName(val int) string {
	var nilthis *QFileDevice
	return nilthis.MemoryMapFlagsItemName(val)
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
