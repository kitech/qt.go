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
func NewQSysInfoFromPointer(cthis unsafe.Pointer) *QSysInfo {
	return &QSysInfo{&qtrt.CObject{cthis}}
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

//  body block end
