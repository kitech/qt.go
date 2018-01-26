package qtcore

// /usr/include/qt/QtCore/qsysinfo.h
// #include <qsysinfo.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
type QSysInfo struct {
	*qtrt.CObject
}

func (this *QSysInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSysInfo) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSysInfoFromPointer(cthis unsafe.Pointer) *QSysInfo {
	return &QSysInfo{&qtrt.CObject{cthis}}
}
func (*QSysInfo) NewFromPointer(cthis unsafe.Pointer) *QSysInfo {
	return NewQSysInfoFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsysinfo.h:220
// index:0
// Public static inline
// QSysInfo::WinVersion windowsVersion()
func (this *QSysInfo) WindowsVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo14windowsVersionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QSysInfo_WindowsVersion() int {
	var nilthis *QSysInfo
	rv := nilthis.WindowsVersion()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:227
// index:0
// Public static inline
// QSysInfo::MacVersion macVersion()
func (this *QSysInfo) MacVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo10macVersionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QSysInfo_MacVersion() int {
	var nilthis *QSysInfo
	rv := nilthis.MacVersion()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:232
// index:0
// Public static
// QString buildCpuArchitecture()
func (this *QSysInfo) BuildCpuArchitecture() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo20buildCpuArchitectureEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSysInfo_BuildCpuArchitecture() *QString /*123*/ {
	var nilthis *QSysInfo
	rv := nilthis.BuildCpuArchitecture()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:233
// index:0
// Public static
// QString currentCpuArchitecture()
func (this *QSysInfo) CurrentCpuArchitecture() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo22currentCpuArchitectureEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSysInfo_CurrentCpuArchitecture() *QString /*123*/ {
	var nilthis *QSysInfo
	rv := nilthis.CurrentCpuArchitecture()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:234
// index:0
// Public static
// QString buildAbi()
func (this *QSysInfo) BuildAbi() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo8buildAbiEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSysInfo_BuildAbi() *QString /*123*/ {
	var nilthis *QSysInfo
	rv := nilthis.BuildAbi()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:236
// index:0
// Public static
// QString kernelType()
func (this *QSysInfo) KernelType() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo10kernelTypeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSysInfo_KernelType() *QString /*123*/ {
	var nilthis *QSysInfo
	rv := nilthis.KernelType()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:237
// index:0
// Public static
// QString kernelVersion()
func (this *QSysInfo) KernelVersion() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo13kernelVersionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSysInfo_KernelVersion() *QString /*123*/ {
	var nilthis *QSysInfo
	rv := nilthis.KernelVersion()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:238
// index:0
// Public static
// QString productType()
func (this *QSysInfo) ProductType() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo11productTypeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSysInfo_ProductType() *QString /*123*/ {
	var nilthis *QSysInfo
	rv := nilthis.ProductType()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:239
// index:0
// Public static
// QString productVersion()
func (this *QSysInfo) ProductVersion() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo14productVersionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSysInfo_ProductVersion() *QString /*123*/ {
	var nilthis *QSysInfo
	rv := nilthis.ProductVersion()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:240
// index:0
// Public static
// QString prettyProductName()
func (this *QSysInfo) PrettyProductName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo17prettyProductNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSysInfo_PrettyProductName() *QString /*123*/ {
	var nilthis *QSysInfo
	rv := nilthis.PrettyProductName()
	return rv
}

// /usr/include/qt/QtCore/qsysinfo.h:242
// index:0
// Public static
// QString machineHostName()
func (this *QSysInfo) MachineHostName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo15machineHostNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSysInfo_MachineHostName() *QString /*123*/ {
	var nilthis *QSysInfo
	rv := nilthis.MachineHostName()
	return rv
}

type QSysInfo__Sizes = int

const QSysInfo__WordSize QSysInfo__Sizes = 64

type QSysInfo__Endian = int

const QSysInfo__BigEndian QSysInfo__Endian = 0
const QSysInfo__LittleEndian QSysInfo__Endian = 1
const QSysInfo__ByteOrder QSysInfo__Endian = 1

type QSysInfo__WinVersion = int

const QSysInfo__WV_None QSysInfo__WinVersion = 0
const QSysInfo__WV_32s QSysInfo__WinVersion = 1
const QSysInfo__WV_95 QSysInfo__WinVersion = 2
const QSysInfo__WV_98 QSysInfo__WinVersion = 3
const QSysInfo__WV_Me QSysInfo__WinVersion = 4
const QSysInfo__WV_DOS_based QSysInfo__WinVersion = 15
const QSysInfo__WV_NT QSysInfo__WinVersion = 16
const QSysInfo__WV_2000 QSysInfo__WinVersion = 32
const QSysInfo__WV_XP QSysInfo__WinVersion = 48
const QSysInfo__WV_2003 QSysInfo__WinVersion = 64
const QSysInfo__WV_VISTA QSysInfo__WinVersion = 128
const QSysInfo__WV_WINDOWS7 QSysInfo__WinVersion = 144
const QSysInfo__WV_WINDOWS8 QSysInfo__WinVersion = 160
const QSysInfo__WV_WINDOWS8_1 QSysInfo__WinVersion = 176
const QSysInfo__WV_WINDOWS10 QSysInfo__WinVersion = 192
const QSysInfo__WV_NT_based QSysInfo__WinVersion = 240
const QSysInfo__WV_4_0 QSysInfo__WinVersion = 16
const QSysInfo__WV_5_0 QSysInfo__WinVersion = 32
const QSysInfo__WV_5_1 QSysInfo__WinVersion = 48
const QSysInfo__WV_5_2 QSysInfo__WinVersion = 64
const QSysInfo__WV_6_0 QSysInfo__WinVersion = 128
const QSysInfo__WV_6_1 QSysInfo__WinVersion = 144
const QSysInfo__WV_6_2 QSysInfo__WinVersion = 160
const QSysInfo__WV_6_3 QSysInfo__WinVersion = 176
const QSysInfo__WV_10_0 QSysInfo__WinVersion = 192
const QSysInfo__WV_CE QSysInfo__WinVersion = 256
const QSysInfo__WV_CENET QSysInfo__WinVersion = 512
const QSysInfo__WV_CE_5 QSysInfo__WinVersion = 768
const QSysInfo__WV_CE_6 QSysInfo__WinVersion = 1024
const QSysInfo__WV_CE_based QSysInfo__WinVersion = 3840

type QSysInfo__MacVersion = int

const QSysInfo__MV_None QSysInfo__MacVersion = 65535
const QSysInfo__MV_Unknown QSysInfo__MacVersion = 0
const QSysInfo__MV_9 QSysInfo__MacVersion = 1
const QSysInfo__MV_10_0 QSysInfo__MacVersion = 2
const QSysInfo__MV_10_1 QSysInfo__MacVersion = 3
const QSysInfo__MV_10_2 QSysInfo__MacVersion = 4
const QSysInfo__MV_10_3 QSysInfo__MacVersion = 5
const QSysInfo__MV_10_4 QSysInfo__MacVersion = 6
const QSysInfo__MV_10_5 QSysInfo__MacVersion = 7
const QSysInfo__MV_10_6 QSysInfo__MacVersion = 8
const QSysInfo__MV_10_7 QSysInfo__MacVersion = 9
const QSysInfo__MV_10_8 QSysInfo__MacVersion = 10
const QSysInfo__MV_10_9 QSysInfo__MacVersion = 11
const QSysInfo__MV_10_10 QSysInfo__MacVersion = 12
const QSysInfo__MV_10_11 QSysInfo__MacVersion = 13
const QSysInfo__MV_10_12 QSysInfo__MacVersion = 14
const QSysInfo__MV_CHEETAH QSysInfo__MacVersion = 2
const QSysInfo__MV_PUMA QSysInfo__MacVersion = 3
const QSysInfo__MV_JAGUAR QSysInfo__MacVersion = 4
const QSysInfo__MV_PANTHER QSysInfo__MacVersion = 5
const QSysInfo__MV_TIGER QSysInfo__MacVersion = 6
const QSysInfo__MV_LEOPARD QSysInfo__MacVersion = 7
const QSysInfo__MV_SNOWLEOPARD QSysInfo__MacVersion = 8
const QSysInfo__MV_LION QSysInfo__MacVersion = 9
const QSysInfo__MV_MOUNTAINLION QSysInfo__MacVersion = 10
const QSysInfo__MV_MAVERICKS QSysInfo__MacVersion = 11
const QSysInfo__MV_YOSEMITE QSysInfo__MacVersion = 12
const QSysInfo__MV_ELCAPITAN QSysInfo__MacVersion = 13
const QSysInfo__MV_SIERRA QSysInfo__MacVersion = 14
const QSysInfo__MV_IOS QSysInfo__MacVersion = 256
const QSysInfo__MV_IOS_4_3 QSysInfo__MacVersion = 323
const QSysInfo__MV_IOS_5_0 QSysInfo__MacVersion = 336
const QSysInfo__MV_IOS_5_1 QSysInfo__MacVersion = 337
const QSysInfo__MV_IOS_6_0 QSysInfo__MacVersion = 352
const QSysInfo__MV_IOS_6_1 QSysInfo__MacVersion = 353
const QSysInfo__MV_IOS_7_0 QSysInfo__MacVersion = 368
const QSysInfo__MV_IOS_7_1 QSysInfo__MacVersion = 369
const QSysInfo__MV_IOS_8_0 QSysInfo__MacVersion = 384
const QSysInfo__MV_IOS_8_1 QSysInfo__MacVersion = 385
const QSysInfo__MV_IOS_8_2 QSysInfo__MacVersion = 386
const QSysInfo__MV_IOS_8_3 QSysInfo__MacVersion = 387
const QSysInfo__MV_IOS_8_4 QSysInfo__MacVersion = 388
const QSysInfo__MV_IOS_9_0 QSysInfo__MacVersion = 400
const QSysInfo__MV_IOS_9_1 QSysInfo__MacVersion = 401
const QSysInfo__MV_IOS_9_2 QSysInfo__MacVersion = 402
const QSysInfo__MV_IOS_9_3 QSysInfo__MacVersion = 403
const QSysInfo__MV_IOS_10_0 QSysInfo__MacVersion = 416
const QSysInfo__MV_TVOS QSysInfo__MacVersion = 512
const QSysInfo__MV_TVOS_9_0 QSysInfo__MacVersion = 656
const QSysInfo__MV_TVOS_9_1 QSysInfo__MacVersion = 657
const QSysInfo__MV_TVOS_9_2 QSysInfo__MacVersion = 658
const QSysInfo__MV_TVOS_10_0 QSysInfo__MacVersion = 672
const QSysInfo__MV_WATCHOS QSysInfo__MacVersion = 1024
const QSysInfo__MV_WATCHOS_2_0 QSysInfo__MacVersion = 1056
const QSysInfo__MV_WATCHOS_2_1 QSysInfo__MacVersion = 1057
const QSysInfo__MV_WATCHOS_2_2 QSysInfo__MacVersion = 1058
const QSysInfo__MV_WATCHOS_3_0 QSysInfo__MacVersion = 1072

//  body block end
