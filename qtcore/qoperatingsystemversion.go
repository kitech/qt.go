package qtcore

// /usr/include/qt/QtCore/qoperatingsystemversion.h
// #include <qoperatingsystemversion.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QOperatingSystemVersion struct {
	*qtrt.CObject
}
type QOperatingSystemVersion_ITF interface {
	QOperatingSystemVersion_PTR() *QOperatingSystemVersion
}

func (ptr *QOperatingSystemVersion) QOperatingSystemVersion_PTR() *QOperatingSystemVersion { return ptr }

func (this *QOperatingSystemVersion) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QOperatingSystemVersion) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQOperatingSystemVersionFromPointer(cthis unsafe.Pointer) *QOperatingSystemVersion {
	return &QOperatingSystemVersion{&qtrt.CObject{cthis}}
}
func (*QOperatingSystemVersion) NewFromPointer(cthis unsafe.Pointer) *QOperatingSystemVersion {
	return NewQOperatingSystemVersionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QOperatingSystemVersion(enum QOperatingSystemVersion::OSType, int, int, int)

/*
Constructs a QOperatingSystemVersion consisting of the OS type osType, and major, minor, and micro version numbers vmajor, vminor and vmicro, respectively.
*/
func NewQOperatingSystemVersion(osType int, vmajor int, vminor int, vmicro int) *QOperatingSystemVersion {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QOperatingSystemVersionC2ENS_6OSTypeEiii", qtrt.FFI_TYPE_POINTER, osType, vmajor, vminor, vmicro)
	qtrt.ErrPrint(err, rv)
	gothis := NewQOperatingSystemVersionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQOperatingSystemVersion)
	return gothis
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QOperatingSystemVersion(enum QOperatingSystemVersion::OSType, int, int, int)

/*
Constructs a QOperatingSystemVersion consisting of the OS type osType, and major, minor, and micro version numbers vmajor, vminor and vmicro, respectively.
*/
func NewQOperatingSystemVersion__(osType int, vmajor int) *QOperatingSystemVersion {
	// arg: 2, int=Int, =Invalid,
	vminor := int(-1)
	// arg: 3, int=Int, =Invalid,
	vmicro := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN23QOperatingSystemVersionC2ENS_6OSTypeEiii", qtrt.FFI_TYPE_POINTER, osType, vmajor, vminor, vmicro)
	qtrt.ErrPrint(err, rv)
	gothis := NewQOperatingSystemVersionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQOperatingSystemVersion)
	return gothis
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QOperatingSystemVersion(enum QOperatingSystemVersion::OSType, int, int, int)

/*
Constructs a QOperatingSystemVersion consisting of the OS type osType, and major, minor, and micro version numbers vmajor, vminor and vmicro, respectively.
*/
func NewQOperatingSystemVersion__1(osType int, vmajor int, vminor int) *QOperatingSystemVersion {
	// arg: 3, int=Int, =Invalid,
	vmicro := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN23QOperatingSystemVersionC2ENS_6OSTypeEiii", qtrt.FFI_TYPE_POINTER, osType, vmajor, vminor, vmicro)
	qtrt.ErrPrint(err, rv)
	gothis := NewQOperatingSystemVersionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQOperatingSystemVersion)
	return gothis
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:93
// index:0
// Public static Visibility=Default Availability=Available
// [16] QOperatingSystemVersion current()

/*
Returns a QOperatingSystemVersion indicating the current OS and its version number.

See also currentType().
*/
func (this *QOperatingSystemVersion) Current() *QOperatingSystemVersion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QOperatingSystemVersion7currentEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQOperatingSystemVersionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQOperatingSystemVersion)
	return rv2
}
func QOperatingSystemVersion_Current() *QOperatingSystemVersion /*123*/ {
	var nilthis *QOperatingSystemVersion
	rv := nilthis.Current()
	return rv
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:95
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] QOperatingSystemVersion::OSType currentType()

/*
Returns the current OS type without constructing a QOperatingSystemVersion instance.

See also current().
*/
func (this *QOperatingSystemVersion) CurrentType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QOperatingSystemVersion11currentTypeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QOperatingSystemVersion_CurrentType() int {
	var nilthis *QOperatingSystemVersion
	rv := nilthis.CurrentType()
	return rv
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int majorVersion() const

/*
Returns the major version number, that is, the first segment of the operating system's version number.

See the main class documentation for what the major version number is on a given operating system.

-1 indicates an unknown or absent version number component.

See also minorVersion() and microVersion().
*/
func (this *QOperatingSystemVersion) MajorVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12majorVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int minorVersion() const

/*
Returns the minor version number, that is, the second segment of the operating system's version number.

See the main class documentation for what the minor version number is on a given operating system.

-1 indicates an unknown or absent version number component.

See also majorVersion() and microVersion().
*/
func (this *QOperatingSystemVersion) MinorVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12minorVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:116
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int microVersion() const

/*
Returns the micro version number, that is, the third segment of the operating system's version number.

See the main class documentation for what the micro version number is on a given operating system.

-1 indicates an unknown or absent version number component.

See also majorVersion() and minorVersion().
*/
func (this *QOperatingSystemVersion) MicroVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12microVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int segmentCount() const

/*
Returns the number of integers stored in the version number.
*/
func (this *QOperatingSystemVersion) SegmentCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion12segmentCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QOperatingSystemVersion::OSType type() const

/*
Returns the OS type identified by the QOperatingSystemVersion.

See also name().
*/
func (this *QOperatingSystemVersion) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qoperatingsystemversion.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*
Returns a string representation of the OS type identified by the QOperatingSystemVersion.

See also type().
*/
func (this *QOperatingSystemVersion) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QOperatingSystemVersion4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

func DeleteQOperatingSystemVersion(this *QOperatingSystemVersion) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QOperatingSystemVersionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum provides symbolic names for the various operating system families supported by QOperatingSystemVersion.


*/
type QOperatingSystemVersion__OSType = int

// An unknown or unsupported operating system.
const QOperatingSystemVersion__Unknown QOperatingSystemVersion__OSType = 0

// The Microsoft Windows operating system.
const QOperatingSystemVersion__Windows QOperatingSystemVersion__OSType = 1

// The Apple macOS operating system.
const QOperatingSystemVersion__MacOS QOperatingSystemVersion__OSType = 2

// The Apple iOS operating system.
const QOperatingSystemVersion__IOS QOperatingSystemVersion__OSType = 3

// The Apple tvOS operating system.
const QOperatingSystemVersion__TvOS QOperatingSystemVersion__OSType = 4

// The Apple watchOS operating system.
const QOperatingSystemVersion__WatchOS QOperatingSystemVersion__OSType = 5

// The Google Android operating system.
const QOperatingSystemVersion__Android QOperatingSystemVersion__OSType = 6

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
