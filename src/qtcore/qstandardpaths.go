//  header block begin
// /usr/include/qt/QtCore/qstandardpaths.h
// #include <qstandardpaths.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 62
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
type QStandardPaths struct {
	*qtrt.CObject
}

func (this *QStandardPaths) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qstandardpaths.h:81
// index:0
// static
// QString writableLocation(enum QStandardPaths::StandardLocation)
func (this *QStandardPaths) WritableLocation(type_ int) {
	// 0: (type QStandardPaths::StandardLocation), (type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths16writableLocationENS_16StandardLocationE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_WritableLocation(type_ int) {
	// 0: (type QStandardPaths::StandardLocation), (type_)
	var nilthis *QStandardPaths
	nilthis.WritableLocation(type_)
}

// /usr/include/qt/QtCore/qstandardpaths.h:82
// index:0
// static
// QStringList standardLocations(enum QStandardPaths::StandardLocation)
func (this *QStandardPaths) StandardLocations(type_ int) {
	// 0: (type QStandardPaths::StandardLocation), (type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths17standardLocationsENS_16StandardLocationE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_StandardLocations(type_ int) {
	// 0: (type QStandardPaths::StandardLocation), (type_)
	var nilthis *QStandardPaths
	nilthis.StandardLocations(type_)
}

// /usr/include/qt/QtCore/qstandardpaths.h:91
// index:0
// static
// QString locate(enum QStandardPaths::StandardLocation, const class QString &, QStandardPaths::LocateOptions)
func (this *QStandardPaths) Locate(type_ int, fileName unsafe.Pointer, options int) {
	// 0: (type QStandardPaths::StandardLocation, fileName const QString &, QFlags<QStandardPaths::LocateOption> options), (type_, fileName, options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths6locateENS_16StandardLocationERK7QString6QFlagsINS_12LocateOptionEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_Locate(type_ int, fileName unsafe.Pointer, options int) {
	// 0: (type QStandardPaths::StandardLocation, fileName const QString &, QFlags<QStandardPaths::LocateOption> options), (type_, fileName, options)
	var nilthis *QStandardPaths
	nilthis.Locate(type_, fileName, options)
}

// /usr/include/qt/QtCore/qstandardpaths.h:92
// index:0
// static
// QStringList locateAll(enum QStandardPaths::StandardLocation, const class QString &, QStandardPaths::LocateOptions)
func (this *QStandardPaths) LocateAll(type_ int, fileName unsafe.Pointer, options int) {
	// 0: (type QStandardPaths::StandardLocation, fileName const QString &, QFlags<QStandardPaths::LocateOption> options), (type_, fileName, options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths9locateAllENS_16StandardLocationERK7QString6QFlagsINS_12LocateOptionEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_LocateAll(type_ int, fileName unsafe.Pointer, options int) {
	// 0: (type QStandardPaths::StandardLocation, fileName const QString &, QFlags<QStandardPaths::LocateOption> options), (type_, fileName, options)
	var nilthis *QStandardPaths
	nilthis.LocateAll(type_, fileName, options)
}

// /usr/include/qt/QtCore/qstandardpaths.h:94
// index:0
// static
// QString displayName(enum QStandardPaths::StandardLocation)
func (this *QStandardPaths) DisplayName(type_ int) {
	// 0: (type QStandardPaths::StandardLocation), (type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths11displayNameENS_16StandardLocationE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_DisplayName(type_ int) {
	// 0: (type QStandardPaths::StandardLocation), (type_)
	var nilthis *QStandardPaths
	nilthis.DisplayName(type_)
}

// /usr/include/qt/QtCore/qstandardpaths.h:97
// index:0
// static
// QString findExecutable(const class QString &, const class QStringList &)
func (this *QStandardPaths) FindExecutable(executableName unsafe.Pointer, paths unsafe.Pointer) {
	// 0: (executableName const QString &, paths const QStringList &), (executableName, paths)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths14findExecutableERK7QStringRK11QStringList", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_FindExecutable(executableName unsafe.Pointer, paths unsafe.Pointer) {
	// 0: (executableName const QString &, paths const QStringList &), (executableName, paths)
	var nilthis *QStandardPaths
	nilthis.FindExecutable(executableName, paths)
}

// /usr/include/qt/QtCore/qstandardpaths.h:100
// index:0
// static
// void enableTestMode(_Bool)
func (this *QStandardPaths) EnableTestMode(testMode bool) {
	// 0: (testMode bool), (testMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths14enableTestModeEb", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_EnableTestMode(testMode bool) {
	// 0: (testMode bool), (testMode)
	var nilthis *QStandardPaths
	nilthis.EnableTestMode(testMode)
}

// /usr/include/qt/QtCore/qstandardpaths.h:102
// index:0
// static
// void setTestModeEnabled(_Bool)
func (this *QStandardPaths) SetTestModeEnabled(testMode bool) {
	// 0: (testMode bool), (testMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths18setTestModeEnabledEb", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_SetTestModeEnabled(testMode bool) {
	// 0: (testMode bool), (testMode)
	var nilthis *QStandardPaths
	nilthis.SetTestModeEnabled(testMode)
}

// /usr/include/qt/QtCore/qstandardpaths.h:103
// index:0
// static
// bool isTestModeEnabled()
func (this *QStandardPaths) IsTestModeEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths17isTestModeEnabledEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_IsTestModeEnabled() {
	// 0: (), ()
	var nilthis *QStandardPaths
	nilthis.IsTestModeEnabled()
}

//  body block end
