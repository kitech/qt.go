package qtcore

// /usr/include/qt/QtCore/qstandardpaths.h
// #include <qstandardpaths.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 60
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStandardPaths) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQStandardPathsFromPointer(cthis unsafe.Pointer) *QStandardPaths {
	return &QStandardPaths{&qtrt.CObject{cthis}}
}
func (*QStandardPaths) NewFromPointer(cthis unsafe.Pointer) *QStandardPaths {
	return NewQStandardPathsFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstandardpaths.h:81
// index:0
// Public static
// QString writableLocation(enum QStandardPaths::StandardLocation)
func (this *QStandardPaths) WritableLocation(type_ int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths16writableLocationENS_16StandardLocationE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QStandardPaths_WritableLocation(type_ int) *QString /*123*/ {
	var nilthis *QStandardPaths
	rv := nilthis.WritableLocation(type_)
	return rv
}

// /usr/include/qt/QtCore/qstandardpaths.h:94
// index:0
// Public static
// QString displayName(enum QStandardPaths::StandardLocation)
func (this *QStandardPaths) DisplayName(type_ int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths11displayNameENS_16StandardLocationE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QStandardPaths_DisplayName(type_ int) *QString /*123*/ {
	var nilthis *QStandardPaths
	rv := nilthis.DisplayName(type_)
	return rv
}

// /usr/include/qt/QtCore/qstandardpaths.h:97
// index:0
// Public static
// QString findExecutable(const class QString &, const class QStringList &)
func (this *QStandardPaths) FindExecutable(executableName *QString, paths *QStringList) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths14findExecutableERK7QStringRK11QStringList", ffiqt.FFI_TYPE_POINTER, executableName, paths)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QStandardPaths_FindExecutable(executableName *QString, paths *QStringList) *QString /*123*/ {
	var nilthis *QStandardPaths
	rv := nilthis.FindExecutable(executableName, paths)
	return rv
}

// /usr/include/qt/QtCore/qstandardpaths.h:100
// index:0
// Public static
// void enableTestMode(_Bool)
func (this *QStandardPaths) EnableTestMode(testMode bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths14enableTestModeEb", ffiqt.FFI_TYPE_POINTER, testMode)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_EnableTestMode(testMode bool) {
	var nilthis *QStandardPaths
	nilthis.EnableTestMode(testMode)
}

// /usr/include/qt/QtCore/qstandardpaths.h:102
// index:0
// Public static
// void setTestModeEnabled(_Bool)
func (this *QStandardPaths) SetTestModeEnabled(testMode bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths18setTestModeEnabledEb", ffiqt.FFI_TYPE_POINTER, testMode)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_SetTestModeEnabled(testMode bool) {
	var nilthis *QStandardPaths
	nilthis.SetTestModeEnabled(testMode)
}

// /usr/include/qt/QtCore/qstandardpaths.h:103
// index:0
// Public static
// bool isTestModeEnabled()
func (this *QStandardPaths) IsTestModeEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths17isTestModeEnabledEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QStandardPaths_IsTestModeEnabled() bool {
	var nilthis *QStandardPaths
	rv := nilthis.IsTestModeEnabled()
	return rv
}

type QStandardPaths__StandardLocation = int

const QStandardPaths__DesktopLocation QStandardPaths__StandardLocation = 0
const QStandardPaths__DocumentsLocation QStandardPaths__StandardLocation = 1
const QStandardPaths__FontsLocation QStandardPaths__StandardLocation = 2
const QStandardPaths__ApplicationsLocation QStandardPaths__StandardLocation = 3
const QStandardPaths__MusicLocation QStandardPaths__StandardLocation = 4
const QStandardPaths__MoviesLocation QStandardPaths__StandardLocation = 5
const QStandardPaths__PicturesLocation QStandardPaths__StandardLocation = 6
const QStandardPaths__TempLocation QStandardPaths__StandardLocation = 7
const QStandardPaths__HomeLocation QStandardPaths__StandardLocation = 8
const QStandardPaths__DataLocation QStandardPaths__StandardLocation = 9
const QStandardPaths__CacheLocation QStandardPaths__StandardLocation = 10
const QStandardPaths__GenericDataLocation QStandardPaths__StandardLocation = 11
const QStandardPaths__RuntimeLocation QStandardPaths__StandardLocation = 12
const QStandardPaths__ConfigLocation QStandardPaths__StandardLocation = 13
const QStandardPaths__DownloadLocation QStandardPaths__StandardLocation = 14
const QStandardPaths__GenericCacheLocation QStandardPaths__StandardLocation = 15
const QStandardPaths__GenericConfigLocation QStandardPaths__StandardLocation = 16
const QStandardPaths__AppDataLocation QStandardPaths__StandardLocation = 17
const QStandardPaths__AppConfigLocation QStandardPaths__StandardLocation = 18
const QStandardPaths__AppLocalDataLocation QStandardPaths__StandardLocation = 9

type QStandardPaths__LocateOption = int

const QStandardPaths__LocateFile QStandardPaths__LocateOption = 0
const QStandardPaths__LocateDirectory QStandardPaths__LocateOption = 1

//  body block end
