package qtcore

// /usr/include/qt/QtCore/qsysinfo.h
// #include <qsysinfo.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
type QSysInfo struct {
	*qtrt.CObject
}
type QSysInfo_ITF interface {
	QSysInfo_PTR() *QSysInfo
}

func (ptr *QSysInfo) QSysInfo_PTR() *QSysInfo { return ptr }

func (this *QSysInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSysInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSysInfoFromPointer(cthis unsafe.Pointer) *QSysInfo {
	return &QSysInfo{&qtrt.CObject{cthis}}
}
func (*QSysInfo) NewFromPointer(cthis unsafe.Pointer) *QSysInfo {
	return NewQSysInfoFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsysinfo.h:220
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] QSysInfo::WinVersion windowsVersion()

/*

 */
func (this *QSysInfo) WindowsVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfo14windowsVersionEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QSysInfo_WindowsVersion() int {
	var nilthis *QSysInfo
	rv := nilthis.WindowsVersion()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:227
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] QSysInfo::MacVersion macVersion()

/*

 */
func (this *QSysInfo) MacVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfo10macVersionEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QSysInfo_MacVersion() int {
	var nilthis *QSysInfo
	rv := nilthis.MacVersion()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:232
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString buildCpuArchitecture()

/*
Returns the architecture of the CPU that Qt was compiled for, in text format. Note that this may not match the actual CPU that the application is running on if there's an emulation layer or if the CPU supports multiple architectures (like x86-64 processors supporting i386 applications). To detect that, use currentCpuArchitecture().

Values returned by this function are stable and will not change over time, so applications can rely on the returned value as an identifier, except that new CPU types may be added over time.

Typical returned values are (note: list not exhaustive):


"arm"
"arm64"
"i386"
"ia64"
"mips"
"mips64"
"power"
"power64"
"sparc"
"sparcv9"
"x86_64"


This function was introduced in  Qt 5.4.

See also QSysInfo::buildAbi() and QSysInfo::currentCpuArchitecture().
*/
func (this *QSysInfo) BuildCpuArchitecture() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfo20buildCpuArchitectureEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QSysInfo_BuildCpuArchitecture() string {
	var nilthis *QSysInfo
	rv := nilthis.BuildCpuArchitecture()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:233
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString currentCpuArchitecture()

/*
Returns the architecture of the CPU that the application is running on, in text format. Note that this function depends on what the OS will report and may not detect the actual CPU architecture if the OS hides that information or is unable to provide it. For example, a 32-bit OS running on a 64-bit CPU is usually unable to determine the CPU is actually capable of running 64-bit programs.

Values returned by this function are mostly stable: an attempt will be made to ensure that they stay constant over time and match the values returned by QSysInfo::builldCpuArchitecture(). However, due to the nature of the operating system functions being used, there may be discrepancies.

Typical returned values are (note: list not exhaustive):


"arm"
"arm64"
"i386"
"ia64"
"mips"
"mips64"
"power"
"power64"
"sparc"
"sparcv9"
"x86_64"


This function was introduced in  Qt 5.4.

See also QSysInfo::buildAbi() and QSysInfo::buildCpuArchitecture().
*/
func (this *QSysInfo) CurrentCpuArchitecture() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfo22currentCpuArchitectureEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QSysInfo_CurrentCpuArchitecture() string {
	var nilthis *QSysInfo
	rv := nilthis.CurrentCpuArchitecture()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:234
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString buildAbi()

/*
Returns the full architecture string that Qt was compiled for. This string is useful for identifying different, incompatible builds. For example, it can be used as an identifier to request an upgrade package from a server.

The values returned from this function are kept stable as follows: the mandatory components of the result will not change in future versions of Qt, but optional suffixes may be added.

The returned value is composed of three or more parts, separated by dashes ("-"). They are:


 ComponentValue
CPU ArchitectureThe same as QSysInfo::buildCpuArchitecture(), such as "arm", "i386", "mips" or "x86_64"
Endianness"little_endian" or "big_endian"
Word sizeWhether it's a 32- or 64-bit application. Possible values are: "llp64" (Windows 64-bit), "lp64" (Unix 64-bit), "ilp32" (32-bit)
(Optional) ABIZero or more components identifying different ABIs possible in this architecture. Currently, Qt has optional ABI components for ARM and MIPS processors: one component is the main ABI (such as "eabi", "o32", "n32", "o64"); another is whether the calling convention is using hardware floating point registers ("hardfloat" is present).Additionally, if Qt was configured with -qreal float, the ABI option tag "qreal_float" will be present. If Qt was configured with another type as qreal, that type is present after "qreal_", with all characters other than letters and digits escaped by an underscore, followed by two hex digits. For example, -qreal long double becomes "qreal_long_20double".



This function was introduced in  Qt 5.4.

See also QSysInfo::buildCpuArchitecture().
*/
func (this *QSysInfo) BuildAbi() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfo8buildAbiEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QSysInfo_BuildAbi() string {
	var nilthis *QSysInfo
	rv := nilthis.BuildAbi()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:236
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString kernelType()

/*
Returns the type of the operating system kernel Qt was compiled for. It's also the kernel the application is running on, unless the host operating system is running a form of compatibility or virtualization layer.

Values returned by this function are stable and will not change over time, so applications can rely on the returned value as an identifier, except that new OS kernel types may be added over time.

On Windows, this function returns the type of Windows kernel, like "winnt". On Unix systems, it returns the same as the output of uname -s (lowercased).

Note: This function may return surprising values: it returns "linux" for all operating systems running Linux (including Android), "qnx" for all operating systems running QNX, "freebsd" for Debian/kFreeBSD, and "darwin" for macOS and iOS. For information on the type of product the application is running on, see productType().

This function was introduced in  Qt 5.4.

See also QFileSelector, kernelVersion(), productType(), productVersion(), and prettyProductName().
*/
func (this *QSysInfo) KernelType() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfo10kernelTypeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QSysInfo_KernelType() string {
	var nilthis *QSysInfo
	rv := nilthis.KernelType()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:237
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString kernelVersion()

/*
Returns the release version of the operating system kernel. On Windows, it returns the version of the NT kernel. On Unix systems, including Android and macOS, it returns the same as the uname -r command would return.

If the version could not be determined, this function may return an empty string.

This function was introduced in  Qt 5.4.

See also kernelType(), productType(), productVersion(), and prettyProductName().
*/
func (this *QSysInfo) KernelVersion() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfo13kernelVersionEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QSysInfo_KernelVersion() string {
	var nilthis *QSysInfo
	rv := nilthis.KernelVersion()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:238
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString productType()

/*
Returns the product name of the operating system this application is running in. If the application is running on some sort of emulation or virtualization layer (such as WINE on a Unix system), this function will inspect the emulation / virtualization layer.

Values returned by this function are stable and will not change over time, so applications can rely on the returned value as an identifier, except that new OS types may be added over time.

Linux and Android note: this function returns "android" for Linux systems running Android userspace, notably when using the Bionic library. For all other Linux systems, regardless of C library being used, it tries to determine the distribution name and returns that. If determining the distribution name failed, it returns "unknown".

macOS note: this function returns "osx" for all macOS systems, regardless of Apple naming convention. The returned string will be updated for Qt 6. Note that this function erroneously returned "macos" for macOS 10.12 in Qt versions 5.6.2, 5.7.1, and 5.8.0.

Darwin, iOS, tvOS, and watchOS note: this function returns "ios" for iOS systems, "tvos" for tvOS systems, "watchos" for watchOS systems, and "darwin" in case the system could not be determined.

FreeBSD note: this function returns "debian" for Debian/kFreeBSD and "unknown" otherwise.

Windows note: this function "winrt" for WinRT builds, and "windows" for normal desktop builds.

For other Unix-type systems, this function usually returns "unknown".

This function was introduced in  Qt 5.4.

See also QFileSelector, kernelType(), kernelVersion(), productVersion(), and prettyProductName().
*/
func (this *QSysInfo) ProductType() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfo11productTypeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QSysInfo_ProductType() string {
	var nilthis *QSysInfo
	rv := nilthis.ProductType()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:239
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString productVersion()

/*
Returns the product version of the operating system in string form. If the version could not be determined, this function returns "unknown".

It will return the Android, iOS, macOS, Windows full-product versions on those systems.

Typical returned values are (note: list not exhaustive):


"2016.09" (Amazon Linux AMI 2016.09)
"7.1" (Android Nougat)
"25" (Fedora 25)
"10.1" (iOS 10.1)
"10.12" (macOS Sierra)
"10.0" (tvOS 10)
"16.10" (Ubuntu 16.10)
"3.1" (watchOS 3.1)
"7 SP 1" (Windows 7 Service Pack 1)
"8.1" (Windows 8.1)
"10" (Windows 10)
"Server 2016" (Windows Server 2016)


On Linux systems, it will try to determine the distribution version and will return that. This is also done on Debian/kFreeBSD, so this function will return Debian version in that case.

In all other Unix-type systems, this function always returns "unknown".

Note: The version string returned from this function is not guaranteed to be orderable. On Linux, the version of the distribution may jump unexpectedly, please refer to the distribution's documentation for versioning practices.

This function was introduced in  Qt 5.4.

See also kernelType(), kernelVersion(), productType(), and prettyProductName().
*/
func (this *QSysInfo) ProductVersion() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfo14productVersionEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QSysInfo_ProductVersion() string {
	var nilthis *QSysInfo
	rv := nilthis.ProductVersion()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:240
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString prettyProductName()

/*
Returns a prettier form of productType() and productVersion(), containing other tokens like the operating system type, codenames and other information. The result of this function is suitable for displaying to the user, but not for long-term storage, as the string may change with updates to Qt.

If productType() is "unknown", this function will instead use the kernelType() and kernelVersion() functions.

This function was introduced in  Qt 5.4.

See also kernelType(), kernelVersion(), productType(), and productVersion().
*/
func (this *QSysInfo) PrettyProductName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfo17prettyProductNameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QSysInfo_PrettyProductName() string {
	var nilthis *QSysInfo
	rv := nilthis.PrettyProductName()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:242
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString machineHostName()

/*
Returns this machine's host name, if one is configured. Note that hostnames are not guaranteed to be globally unique, especially if they were configured automatically.

This function does not guarantee the returned host name is a Fully Qualified Domain Name (FQDN). For that, use QHostInfo to resolve the returned name to an FQDN.

This function returns the same as QHostInfo::localHostName().

This function was introduced in  Qt 5.6.

See also QHostInfo::localDomainName.
*/
func (this *QSysInfo) MachineHostName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfo15machineHostNameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QSysInfo_MachineHostName() string {
	var nilthis *QSysInfo
	rv := nilthis.MachineHostName()
	return rv
}

func DeleteQSysInfo(this *QSysInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSysInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum provides platform-specific information about the sizes of data structures used by the underlying architecture.


*/
type QSysInfo__Sizes = int

//
const QSysInfo__WordSize QSysInfo__Sizes = 64

/*
QSysInfo::ByteOrderBigEndian or LittleEndianEquals BigEndian or LittleEndian, depending on the platform's byte order.

*/
type QSysInfo__Endian = int

// Big-endian byte order (also called Network byte order)
const QSysInfo__BigEndian QSysInfo__Endian = 0

// Little-endian byte order
const QSysInfo__LittleEndian QSysInfo__Endian = 1

//
const QSysInfo__ByteOrder QSysInfo__Endian = 1

/*


 */
type QSysInfo__WinVersion = int

//
const QSysInfo__WV_None QSysInfo__WinVersion = 0

//
const QSysInfo__WV_32s QSysInfo__WinVersion = 1

//
const QSysInfo__WV_95 QSysInfo__WinVersion = 2

//
const QSysInfo__WV_98 QSysInfo__WinVersion = 3

//
const QSysInfo__WV_Me QSysInfo__WinVersion = 4

//
const QSysInfo__WV_DOS_based QSysInfo__WinVersion = 15

//
const QSysInfo__WV_NT QSysInfo__WinVersion = 16

//
const QSysInfo__WV_2000 QSysInfo__WinVersion = 32

//
const QSysInfo__WV_XP QSysInfo__WinVersion = 48

//
const QSysInfo__WV_2003 QSysInfo__WinVersion = 64

//
const QSysInfo__WV_VISTA QSysInfo__WinVersion = 128

//
const QSysInfo__WV_WINDOWS7 QSysInfo__WinVersion = 144

//
const QSysInfo__WV_WINDOWS8 QSysInfo__WinVersion = 160

//
const QSysInfo__WV_WINDOWS8_1 QSysInfo__WinVersion = 176

//
const QSysInfo__WV_WINDOWS10 QSysInfo__WinVersion = 192

//
const QSysInfo__WV_NT_based QSysInfo__WinVersion = 240

//
const QSysInfo__WV_4_0 QSysInfo__WinVersion = 16

//
const QSysInfo__WV_5_0 QSysInfo__WinVersion = 32

//
const QSysInfo__WV_5_1 QSysInfo__WinVersion = 48

//
const QSysInfo__WV_5_2 QSysInfo__WinVersion = 64

//
const QSysInfo__WV_6_0 QSysInfo__WinVersion = 128

//
const QSysInfo__WV_6_1 QSysInfo__WinVersion = 144

//
const QSysInfo__WV_6_2 QSysInfo__WinVersion = 160

//
const QSysInfo__WV_6_3 QSysInfo__WinVersion = 176

//
const QSysInfo__WV_10_0 QSysInfo__WinVersion = 192

//
const QSysInfo__WV_CE QSysInfo__WinVersion = 256

//
const QSysInfo__WV_CENET QSysInfo__WinVersion = 512

//
const QSysInfo__WV_CE_5 QSysInfo__WinVersion = 768

//
const QSysInfo__WV_CE_6 QSysInfo__WinVersion = 1024

//
const QSysInfo__WV_CE_based QSysInfo__WinVersion = 3840

/*


 */
type QSysInfo__MacVersion = int

//
const QSysInfo__MV_None QSysInfo__MacVersion = 65535

//
const QSysInfo__MV_Unknown QSysInfo__MacVersion = 0

//
const QSysInfo__MV_9 QSysInfo__MacVersion = 1

//
const QSysInfo__MV_10_0 QSysInfo__MacVersion = 2

//
const QSysInfo__MV_10_1 QSysInfo__MacVersion = 3

//
const QSysInfo__MV_10_2 QSysInfo__MacVersion = 4

//
const QSysInfo__MV_10_3 QSysInfo__MacVersion = 5

//
const QSysInfo__MV_10_4 QSysInfo__MacVersion = 6

//
const QSysInfo__MV_10_5 QSysInfo__MacVersion = 7

//
const QSysInfo__MV_10_6 QSysInfo__MacVersion = 8

//
const QSysInfo__MV_10_7 QSysInfo__MacVersion = 9

//
const QSysInfo__MV_10_8 QSysInfo__MacVersion = 10

//
const QSysInfo__MV_10_9 QSysInfo__MacVersion = 11

//
const QSysInfo__MV_10_10 QSysInfo__MacVersion = 12

//
const QSysInfo__MV_10_11 QSysInfo__MacVersion = 13

//
const QSysInfo__MV_10_12 QSysInfo__MacVersion = 14

//
const QSysInfo__MV_CHEETAH QSysInfo__MacVersion = 2

//
const QSysInfo__MV_PUMA QSysInfo__MacVersion = 3

//
const QSysInfo__MV_JAGUAR QSysInfo__MacVersion = 4

//
const QSysInfo__MV_PANTHER QSysInfo__MacVersion = 5

//
const QSysInfo__MV_TIGER QSysInfo__MacVersion = 6

//
const QSysInfo__MV_LEOPARD QSysInfo__MacVersion = 7

//
const QSysInfo__MV_SNOWLEOPARD QSysInfo__MacVersion = 8

//
const QSysInfo__MV_LION QSysInfo__MacVersion = 9

//
const QSysInfo__MV_MOUNTAINLION QSysInfo__MacVersion = 10

//
const QSysInfo__MV_MAVERICKS QSysInfo__MacVersion = 11

//
const QSysInfo__MV_YOSEMITE QSysInfo__MacVersion = 12

//
const QSysInfo__MV_ELCAPITAN QSysInfo__MacVersion = 13

//
const QSysInfo__MV_SIERRA QSysInfo__MacVersion = 14

//
const QSysInfo__MV_IOS QSysInfo__MacVersion = 256

//
const QSysInfo__MV_IOS_4_3 QSysInfo__MacVersion = 323

//
const QSysInfo__MV_IOS_5_0 QSysInfo__MacVersion = 336

//
const QSysInfo__MV_IOS_5_1 QSysInfo__MacVersion = 337

//
const QSysInfo__MV_IOS_6_0 QSysInfo__MacVersion = 352

//
const QSysInfo__MV_IOS_6_1 QSysInfo__MacVersion = 353

//
const QSysInfo__MV_IOS_7_0 QSysInfo__MacVersion = 368

//
const QSysInfo__MV_IOS_7_1 QSysInfo__MacVersion = 369

//
const QSysInfo__MV_IOS_8_0 QSysInfo__MacVersion = 384

//
const QSysInfo__MV_IOS_8_1 QSysInfo__MacVersion = 385

//
const QSysInfo__MV_IOS_8_2 QSysInfo__MacVersion = 386

//
const QSysInfo__MV_IOS_8_3 QSysInfo__MacVersion = 387

//
const QSysInfo__MV_IOS_8_4 QSysInfo__MacVersion = 388

//
const QSysInfo__MV_IOS_9_0 QSysInfo__MacVersion = 400

//
const QSysInfo__MV_IOS_9_1 QSysInfo__MacVersion = 401

//
const QSysInfo__MV_IOS_9_2 QSysInfo__MacVersion = 402

//
const QSysInfo__MV_IOS_9_3 QSysInfo__MacVersion = 403

//
const QSysInfo__MV_IOS_10_0 QSysInfo__MacVersion = 416

//
const QSysInfo__MV_TVOS QSysInfo__MacVersion = 512

//
const QSysInfo__MV_TVOS_9_0 QSysInfo__MacVersion = 656

//
const QSysInfo__MV_TVOS_9_1 QSysInfo__MacVersion = 657

//
const QSysInfo__MV_TVOS_9_2 QSysInfo__MacVersion = 658

//
const QSysInfo__MV_TVOS_10_0 QSysInfo__MacVersion = 672

//
const QSysInfo__MV_WATCHOS QSysInfo__MacVersion = 1024

//
const QSysInfo__MV_WATCHOS_2_0 QSysInfo__MacVersion = 1056

//
const QSysInfo__MV_WATCHOS_2_1 QSysInfo__MacVersion = 1057

//
const QSysInfo__MV_WATCHOS_2_2 QSysInfo__MacVersion = 1058

//
const QSysInfo__MV_WATCHOS_3_0 QSysInfo__MacVersion = 1072

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
