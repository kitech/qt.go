//  header block begin
// block begin header
// /usr/include/qt/QtCore/qsysinfo.h
// #include <qsysinfo.h>
// #include <QtCore>
package qtcore

//  header block end

//  main block begin
// block begin main
//  main block end

//  use block begin
// block begin use
//  use block end

//  ext block begin
// block begin ext

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
// block begin body
type QSysInfo struct {
	*qtrt.CObject
}

func (this *QSysInfo) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQSysInfoFromPointer(cthis unsafe.Pointer) *QSysInfo {
	return &QSysInfo{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qsysinfo.h:220
// index:0
// Public static inline
// QSysInfo::WinVersion windowsVersion()
func (this *QSysInfo) WindowsVersion() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo14windowsVersionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSysInfo_WindowsVersion() {
	var nilthis *QSysInfo
	nilthis.WindowsVersion()
}

// /usr/include/qt/QtCore/qsysinfo.h:227
// index:0
// Public static inline
// QSysInfo::MacVersion macVersion()
func (this *QSysInfo) MacVersion() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo10macVersionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSysInfo_MacVersion() {
	var nilthis *QSysInfo
	nilthis.MacVersion()
}

// /usr/include/qt/QtCore/qsysinfo.h:232
// index:0
// Public static
// QString buildCpuArchitecture()
func (this *QSysInfo) BuildCpuArchitecture() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo20buildCpuArchitectureEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSysInfo_BuildCpuArchitecture() {
	var nilthis *QSysInfo
	nilthis.BuildCpuArchitecture()
}

// /usr/include/qt/QtCore/qsysinfo.h:233
// index:0
// Public static
// QString currentCpuArchitecture()
func (this *QSysInfo) CurrentCpuArchitecture() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo22currentCpuArchitectureEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSysInfo_CurrentCpuArchitecture() {
	var nilthis *QSysInfo
	nilthis.CurrentCpuArchitecture()
}

// /usr/include/qt/QtCore/qsysinfo.h:234
// index:0
// Public static
// QString buildAbi()
func (this *QSysInfo) BuildAbi() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo8buildAbiEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSysInfo_BuildAbi() {
	var nilthis *QSysInfo
	nilthis.BuildAbi()
}

// /usr/include/qt/QtCore/qsysinfo.h:236
// index:0
// Public static
// QString kernelType()
func (this *QSysInfo) KernelType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo10kernelTypeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSysInfo_KernelType() {
	var nilthis *QSysInfo
	nilthis.KernelType()
}

// /usr/include/qt/QtCore/qsysinfo.h:237
// index:0
// Public static
// QString kernelVersion()
func (this *QSysInfo) KernelVersion() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo13kernelVersionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSysInfo_KernelVersion() {
	var nilthis *QSysInfo
	nilthis.KernelVersion()
}

// /usr/include/qt/QtCore/qsysinfo.h:238
// index:0
// Public static
// QString productType()
func (this *QSysInfo) ProductType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo11productTypeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSysInfo_ProductType() {
	var nilthis *QSysInfo
	nilthis.ProductType()
}

// /usr/include/qt/QtCore/qsysinfo.h:239
// index:0
// Public static
// QString productVersion()
func (this *QSysInfo) ProductVersion() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo14productVersionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSysInfo_ProductVersion() {
	var nilthis *QSysInfo
	nilthis.ProductVersion()
}

// /usr/include/qt/QtCore/qsysinfo.h:240
// index:0
// Public static
// QString prettyProductName()
func (this *QSysInfo) PrettyProductName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo17prettyProductNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSysInfo_PrettyProductName() {
	var nilthis *QSysInfo
	nilthis.PrettyProductName()
}

// /usr/include/qt/QtCore/qsysinfo.h:242
// index:0
// Public static
// QString machineHostName()
func (this *QSysInfo) MachineHostName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSysInfo15machineHostNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSysInfo_MachineHostName() {
	var nilthis *QSysInfo
	nilthis.MachineHostName()
}

//  body block end
