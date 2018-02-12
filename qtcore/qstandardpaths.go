package qtcore

// /usr/include/qt/QtCore/qstandardpaths.h
// #include <qstandardpaths.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 62
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStandardPathsFromPointer(cthis unsafe.Pointer) *QStandardPaths {
	return &QStandardPaths{&qtrt.CObject{cthis}}
}
func (*QStandardPaths) NewFromPointer(cthis unsafe.Pointer) *QStandardPaths {
	return NewQStandardPathsFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstandardpaths.h:81
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString writableLocation(enum QStandardPaths::StandardLocation)
func (this *QStandardPaths) WritableLocation(type_ int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths16writableLocationENS_16StandardLocationE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QStandardPaths_WritableLocation(type_ int) string {
	var nilthis *QStandardPaths
	rv := nilthis.WritableLocation(type_)
	return rv
}

// /usr/include/qt/QtCore/qstandardpaths.h:82
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList standardLocations(enum QStandardPaths::StandardLocation)
func (this *QStandardPaths) StandardLocations(type_ int) *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths17standardLocationsENS_16StandardLocationE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}
func QStandardPaths_StandardLocations(type_ int) *QStringList /*123*/ {
	var nilthis *QStandardPaths
	rv := nilthis.StandardLocations(type_)
	return rv
}

// /usr/include/qt/QtCore/qstandardpaths.h:91
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString locate(enum QStandardPaths::StandardLocation, const QString &, QStandardPaths::LocateOptions)
func (this *QStandardPaths) Locate(type_ int, fileName string, options int) string {
	var tmpArg1 = NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths6locateENS_16StandardLocationERK7QString6QFlagsINS_12LocateOptionEE", qtrt.FFI_TYPE_POINTER, type_, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QStandardPaths_Locate(type_ int, fileName string, options int) string {
	var nilthis *QStandardPaths
	rv := nilthis.Locate(type_, fileName, options)
	return rv
}

// /usr/include/qt/QtCore/qstandardpaths.h:92
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList locateAll(enum QStandardPaths::StandardLocation, const QString &, QStandardPaths::LocateOptions)
func (this *QStandardPaths) LocateAll(type_ int, fileName string, options int) *QStringList /*123*/ {
	var tmpArg1 = NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths9locateAllENS_16StandardLocationERK7QString6QFlagsINS_12LocateOptionEE", qtrt.FFI_TYPE_POINTER, type_, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}
func QStandardPaths_LocateAll(type_ int, fileName string, options int) *QStringList /*123*/ {
	var nilthis *QStandardPaths
	rv := nilthis.LocateAll(type_, fileName, options)
	return rv
}

// /usr/include/qt/QtCore/qstandardpaths.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString displayName(enum QStandardPaths::StandardLocation)
func (this *QStandardPaths) DisplayName(type_ int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths11displayNameENS_16StandardLocationE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QStandardPaths_DisplayName(type_ int) string {
	var nilthis *QStandardPaths
	rv := nilthis.DisplayName(type_)
	return rv
}

// /usr/include/qt/QtCore/qstandardpaths.h:97
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString findExecutable(const QString &, const QStringList &)
func (this *QStandardPaths) FindExecutable(executableName string, paths *QStringList) string {
	var tmpArg0 = NewQString_5(executableName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = paths.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths14findExecutableERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QStandardPaths_FindExecutable(executableName string, paths *QStringList) string {
	var nilthis *QStandardPaths
	rv := nilthis.FindExecutable(executableName, paths)
	return rv
}

// /usr/include/qt/QtCore/qstandardpaths.h:100
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void enableTestMode(_Bool)
func (this *QStandardPaths) EnableTestMode(testMode bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths14enableTestModeEb", qtrt.FFI_TYPE_POINTER, testMode)
	qtrt.ErrPrint(err, rv)
}
func QStandardPaths_EnableTestMode(testMode bool) {
	var nilthis *QStandardPaths
	nilthis.EnableTestMode(testMode)
}

// /usr/include/qt/QtCore/qstandardpaths.h:102
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setTestModeEnabled(_Bool)
func (this *QStandardPaths) SetTestModeEnabled(testMode bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths18setTestModeEnabledEb", qtrt.FFI_TYPE_POINTER, testMode)
	qtrt.ErrPrint(err, rv)
}
func QStandardPaths_SetTestModeEnabled(testMode bool) {
	var nilthis *QStandardPaths
	nilthis.SetTestModeEnabled(testMode)
}

// /usr/include/qt/QtCore/qstandardpaths.h:103
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isTestModeEnabled()
func (this *QStandardPaths) IsTestModeEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths17isTestModeEnabledEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QStandardPaths_IsTestModeEnabled() bool {
	var nilthis *QStandardPaths
	rv := nilthis.IsTestModeEnabled()
	return rv
}

func DeleteQStandardPaths(this *QStandardPaths) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPathsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
