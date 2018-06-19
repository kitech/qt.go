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
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QStandardPaths struct {
	*qtrt.CObject
}
type QStandardPaths_ITF interface {
	QStandardPaths_PTR() *QStandardPaths
}

func (ptr *QStandardPaths) QStandardPaths_PTR() *QStandardPaths { return ptr }

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
// [8] QString writableLocation(QStandardPaths::StandardLocation)

/*
Returns the directory where files of type should be written to, or an empty string if the location cannot be determined.

Note: The storage location returned can be a directory that does not exist; i.e., it may need to be created by the system or the user.
*/
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
// [8] QStringList standardLocations(QStandardPaths::StandardLocation)

/*
Returns all the directories where files of type belong.

The list of directories is sorted from high to low priority, starting with writableLocation() if it can be determined. This list is empty if no locations for type are defined.

See also writableLocation().
*/
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
// [8] QString locate(QStandardPaths::StandardLocation, const QString &, QStandardPaths::LocateOptions)

/*
Tries to find a file or directory called fileName in the standard locations for type.

The full path to the first file or directory (depending on options) found is returned. If no such file or directory can be found, an empty string is returned.
*/
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

// /usr/include/qt/QtCore/qstandardpaths.h:91
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString locate(QStandardPaths::StandardLocation, const QString &, QStandardPaths::LocateOptions)

/*
Tries to find a file or directory called fileName in the standard locations for type.

The full path to the first file or directory (depending on options) found is returned. If no such file or directory can be found, an empty string is returned.
*/
func (this *QStandardPaths) Locate__(type_ int, fileName string) string {
	var tmpArg1 = NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QStandardPaths::LocateOptions=Typedef, QStandardPaths::LocateOptions=Typedef, QFlags<QStandardPaths::LocateOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths6locateENS_16StandardLocationERK7QString6QFlagsINS_12LocateOptionEE", qtrt.FFI_TYPE_POINTER, type_, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstandardpaths.h:92
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList locateAll(QStandardPaths::StandardLocation, const QString &, QStandardPaths::LocateOptions)

/*
Tries to find all files or directories called fileName in the standard locations for type.

The options flag allows to specify whether to look for files or directories.

Returns the list of all the files that were found.
*/
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

// /usr/include/qt/QtCore/qstandardpaths.h:92
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList locateAll(QStandardPaths::StandardLocation, const QString &, QStandardPaths::LocateOptions)

/*
Tries to find all files or directories called fileName in the standard locations for type.

The options flag allows to specify whether to look for files or directories.

Returns the list of all the files that were found.
*/
func (this *QStandardPaths) LocateAll__(type_ int, fileName string) *QStringList /*123*/ {
	var tmpArg1 = NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QStandardPaths::LocateOptions=Typedef, QStandardPaths::LocateOptions=Typedef, QFlags<QStandardPaths::LocateOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths9locateAllENS_16StandardLocationERK7QString6QFlagsINS_12LocateOptionEE", qtrt.FFI_TYPE_POINTER, type_, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qstandardpaths.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString displayName(QStandardPaths::StandardLocation)

/*
Returns a localized display name for the given location type or an empty QString if no relevant location can be found.
*/
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

/*
Finds the executable named executableName in the paths specified by paths, or the system paths if paths is empty.

On most operating systems the system path is determined by the PATH environment variable.

The directories where to search for the executable can be set in the paths argument. To search in both your own paths and the system paths, call findExecutable twice, once with paths set and once with paths empty.

Symlinks are not resolved, in order to preserve behavior for the case of executables whose behavior depends on the name they are invoked with.

Note: On Windows, the usual executable extensions (from the PATHEXT environment variable) are automatically appended, so that for instance findExecutable("foo") will find foo.exe or foo.bat if present.

Returns the absolute file path to the executable, or an empty string if not found.
*/
func (this *QStandardPaths) FindExecutable(executableName string, paths QStringList_ITF) string {
	var tmpArg0 = NewQString_5(executableName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if paths != nil && paths.QStringList_PTR() != nil {
		convArg1 = paths.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths14findExecutableERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QStandardPaths_FindExecutable(executableName string, paths QStringList_ITF) string {
	var nilthis *QStandardPaths
	rv := nilthis.FindExecutable(executableName, paths)
	return rv
}

// /usr/include/qt/QtCore/qstandardpaths.h:97
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString findExecutable(const QString &, const QStringList &)

/*
Finds the executable named executableName in the paths specified by paths, or the system paths if paths is empty.

On most operating systems the system path is determined by the PATH environment variable.

The directories where to search for the executable can be set in the paths argument. To search in both your own paths and the system paths, call findExecutable twice, once with paths set and once with paths empty.

Symlinks are not resolved, in order to preserve behavior for the case of executables whose behavior depends on the name they are invoked with.

Note: On Windows, the usual executable extensions (from the PATHEXT environment variable) are automatically appended, so that for instance findExecutable("foo") will find foo.exe or foo.bat if present.

Returns the absolute file path to the executable, or an empty string if not found.
*/
func (this *QStandardPaths) FindExecutable__(executableName string) string {
	var tmpArg0 = NewQString_5(executableName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStandardPaths14findExecutableERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstandardpaths.h:100
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void enableTestMode(bool)

/*

 */
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
// [-2] void setTestModeEnabled(bool)

/*
If testMode is true, this enables a special "test mode" in QStandardPaths, which changes writable locations to point to test directories, in order to prevent auto tests from reading from or writing to the current user's configuration.

This affects the locations into which test programs might write files: GenericDataLocation, DataLocation, ConfigLocation, GenericConfigLocation, AppConfigLocation, GenericCacheLocation, CacheLocation. Other locations are not affected.

On Unix, XDG_DATA_HOME is set to ~/.qttest/share, XDG_CONFIG_HOME is set to ~/.qttest/config, and XDG_CACHE_HOME is set to ~/.qttest/cache.

On macOS, data goes to ~/.qttest/Application Support, cache goes to ~/.qttest/Cache, and config goes to ~/.qttest/Preferences.

On Windows, everything goes to a "qttest" directory under Application Data.
*/
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

/*

 */
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

/*
This enum describes the different locations that can be queried using methods such as QStandardPaths::writableLocation, QStandardPaths::standardLocations, and QStandardPaths::displayName.

Some of the values in this enum represent a user configuration. Such enum values will return the same paths in different applications, so they could be used to share data with other applications. Other values are specific to this application. Each enum value in the table below describes whether it's application-specific or generic.

Application-specific directories should be assumed to be unreachable by other applications. Therefore, files placed there might not be readable by other applications, even if run by the same user. On the other hand, generic directories should be assumed to be accessible by all applications run by this user, but should still be assumed to be unreachable by applications by other users.

Data interchange with other users is out of the scope of QStandardPaths.



The following table gives examples of paths on different operating systems. The first path is the writable path (unless noted). Other, additional paths, if any, represent non-writable locations.


 Path typemacOSWindows
DesktopLocation"~/Desktop""C:/Users/<USER>/Desktop"
DocumentsLocation"~/Documents""C:/Users/<USER>/Documents"
FontsLocation"/System/Library/Fonts" (not writable)"C:/Windows/Fonts" (not writable)
ApplicationsLocation"/Applications" (not writable)"C:/Users/<USER>/AppData/Roaming/Microsoft/Windows/Start Menu/Programs"
MusicLocation"~/Music""C:/Users/<USER>/Music"
MoviesLocation"~/Movies""C:/Users/<USER>/Videos"
PicturesLocation"~/Pictures""C:/Users/<USER>/Pictures"
TempLocationrandomly generated by the OS"C:/Users/<USER>/AppData/Local/Temp"
HomeLocation"~""C:/Users/<USER>"
DataLocation"~/Library/Application Support/<APPNAME>", "/Library/Application Support/<APPNAME>". "<APPDIR>/../Resources""C:/Users/<USER>/AppData/Local/<APPNAME>", "C:/ProgramData/<APPNAME>", "<APPDIR>", "<APPDIR>/data", "<APPDIR>/data/<APPNAME>"
CacheLocation"~/Library/Caches/<APPNAME>", "/Library/Caches/<APPNAME>""C:/Users/<USER>/AppData/Local/<APPNAME>/cache"
GenericDataLocation"~/Library/Application Support", "/Library/Application Support""C:/Users/<USER>/AppData/Local", "C:/ProgramData", "<APPDIR>", "<APPDIR>/data"
RuntimeLocation"~/Library/Application Support""C:/Users/<USER>"
ConfigLocation"~/Library/Preferences""C:/Users/<USER>/AppData/Local/<APPNAME>", "C:/ProgramData/<APPNAME>"
GenericConfigLocation"~/Library/Preferences""C:/Users/<USER>/AppData/Local", "C:/ProgramData"
DownloadLocation"~/Downloads""C:/Users/<USER>/Documents"
GenericCacheLocation"~/Library/Caches", "/Library/Caches""C:/Users/<USER>/AppData/Local/cache"
AppDataLocation"~/Library/Application Support/<APPNAME>", "/Library/Application Support/<APPNAME>". "<APPDIR>/../Resources""C:/Users/<USER>/AppData/Roaming/<APPNAME>", "C:/ProgramData/<APPNAME>", "<APPDIR>", "<APPDIR>/data", "<APPDIR>/data/<APPNAME>"
AppLocalDataLocation"~/Library/Application Support/<APPNAME>", "/Library/Application Support/<APPNAME>". "<APPDIR>/../Resources""C:/Users/<USER>/AppData/Local/<APPNAME>", "C:/ProgramData/<APPNAME>", "<APPDIR>", "<APPDIR>/data", "<APPDIR>/data/<APPNAME>"
AppConfigLocation"~/Library/Preferences/<APPNAME>""C:/Users/<USER>/AppData/Local/<APPNAME>", "C:/ProgramData/<APPNAME>"



 Path typeLinux
DesktopLocation"~/Desktop"
DocumentsLocation"~/Documents"
FontsLocation"~/.fonts"
ApplicationsLocation"~/.local/share/applications", "/usr/local/share/applications", "/usr/share/applications"
MusicLocation"~/Music"
MoviesLocation"~/Videos"
PicturesLocation"~/Pictures"
TempLocation"/tmp"
HomeLocation"~"
DataLocation"~/.local/share/<APPNAME>", "/usr/local/share/<APPNAME>", "/usr/share/<APPNAME>"
CacheLocation"~/.cache/<APPNAME>"
GenericDataLocation"~/.local/share", "/usr/local/share", "/usr/share"
RuntimeLocation"/run/user/<USER>"
ConfigLocation"~/.config", "/etc/xdg"
GenericConfigLocation"~/.config", "/etc/xdg"
DownloadLocation"~/Downloads"
GenericCacheLocation"~/.cache"
AppDataLocation"~/.local/share/<APPNAME>", "/usr/local/share/<APPNAME>", "/usr/share/<APPNAME>"
AppLocalDataLocation"~/.local/share/<APPNAME>", "/usr/local/share/<APPNAME>", "/usr/share/<APPNAME>"
AppConfigLocation"~/.config/<APPNAME>", "/etc/xdg/<APPNAME>"



 Path typeAndroidiOS
DesktopLocation"<APPROOT>/files""<APPROOT>/Documents/Desktop"
DocumentsLocation"<USER>/Documents", "<USER>/<APPNAME>/Documents""<APPROOT>/Documents"
FontsLocation"/system/fonts" (not writable)"<APPROOT>/Library/Fonts"
ApplicationsLocationnot supported (directory not readable)not supported
MusicLocation"<USER>/Music", "<USER>/<APPNAME>/Music""<APPROOT>/Documents/Music"
MoviesLocation"<USER>/Movies", "<USER>/<APPNAME>/Movies""<APPROOT>/Documents/Movies"
PicturesLocation"<USER>/Pictures", "<USER>/<APPNAME>/Pictures""<APPROOT>/Documents/Pictures", "assets-library://"
TempLocation"<APPROOT>/cache""<APPROOT>/tmp"
HomeLocation"<APPROOT>/files""<APPROOT>" (not writable)
DataLocation"<APPROOT>/files", "<USER>/<APPNAME>/files""<APPROOT>/Library/Application Support"
CacheLocation"<APPROOT>/cache", "<USER>/<APPNAME>/cache""<APPROOT>/Library/Caches"
GenericDataLocation"<USER>""<APPROOT>/Documents"
RuntimeLocation"<APPROOT>/cache"not supported
ConfigLocation"<APPROOT>/files/settings""<APPROOT>/Library/Preferences"
GenericConfigLocation"<APPROOT>/files/settings" (there is no shared settings)"<APPROOT>/Library/Preferences"
DownloadLocation"<USER>/Downloads", "<USER>/<APPNAME>/Downloads""<APPROOT>/Documents/Downloads"
GenericCacheLocation"<APPROOT>/cache" (there is no shared cache)"<APPROOT>/Library/Caches"
AppDataLocation"<APPROOT>/files", "<USER>/<APPNAME>/files""<APPROOT>/Library/Application Support"
AppConfigLocation"<APPROOT>/files/settings""<APPROOT>/Library/Preferences/<APPNAME>"
AppLocalDataLocation"<APPROOT>/files", "<USER>/<APPNAME>/files""<APPROOT>/Library/Application Support"


In the table above, <APPNAME> is usually the organization name, the application name, or both, or a unique name generated at packaging. Similarly, <APPROOT> is the location where this application is installed (often a sandbox). <APPDIR> is the directory containing the application executable.

The paths above should not be relied upon, as they may change according to OS configuration, locale, or they may change in future Qt versions.

Note: On Android, applications with open files on the external storage (<USER> locations), will be killed if the external storage is unmounted.


See also writableLocation(), standardLocations(), displayName(), locate(), and locateAll().

*/
type QStandardPaths__StandardLocation = int

// Returns the user's desktop directory. This is a generic value. On systems with no concept of a desktop.
const QStandardPaths__DesktopLocation QStandardPaths__StandardLocation = 0

// Returns the directory containing user document files. This is a generic value. The returned path is never empty.
const QStandardPaths__DocumentsLocation QStandardPaths__StandardLocation = 1

// Returns the directory containing user's fonts. This is a generic value. Note that installing fonts may require additional, platform-specific operations.
const QStandardPaths__FontsLocation QStandardPaths__StandardLocation = 2

// Returns the directory containing the user applications (either executables, application bundles, or shortcuts to them). This is a generic value. Note that installing applications may require additional, platform-specific operations. Files, folders or shortcuts in this directory are platform-specific.
const QStandardPaths__ApplicationsLocation QStandardPaths__StandardLocation = 3

// Returns the directory containing the user's music or other audio files. This is a generic value. If no directory specific for music files exists, a sensible fallback for storing user documents is returned.
const QStandardPaths__MusicLocation QStandardPaths__StandardLocation = 4

// Returns the directory containing the user's movies and videos. This is a generic value. If no directory specific for movie files exists, a sensible fallback for storing user documents is returned.
const QStandardPaths__MoviesLocation QStandardPaths__StandardLocation = 5

// Returns the directory containing the user's pictures or photos. This is a generic value. If no directory specific for picture files exists, a sensible fallback for storing user documents is returned.
const QStandardPaths__PicturesLocation QStandardPaths__StandardLocation = 6

// Returns a directory where temporary files can be stored. The returned value might be application-specific, shared among other applications for this user, or even system-wide. The returned path is never empty.
const QStandardPaths__TempLocation QStandardPaths__StandardLocation = 7

// Returns the user's home directory (the same as QDir::homePath()). On Unix systems, this is equal to the HOME environment variable. This value might be generic or application-specific, but the returned path is never empty.
const QStandardPaths__HomeLocation QStandardPaths__StandardLocation = 8

// Returns the same value as AppLocalDataLocation. This enumeration value is deprecated. Using AppDataLocation is preferable since on Windows, the roaming path is recommended.
const QStandardPaths__DataLocation QStandardPaths__StandardLocation = 9

//
const QStandardPaths__CacheLocation QStandardPaths__StandardLocation = 10

//
const QStandardPaths__GenericDataLocation QStandardPaths__StandardLocation = 11

//
const QStandardPaths__RuntimeLocation QStandardPaths__StandardLocation = 12

//
const QStandardPaths__ConfigLocation QStandardPaths__StandardLocation = 13

//
const QStandardPaths__DownloadLocation QStandardPaths__StandardLocation = 14

//
const QStandardPaths__GenericCacheLocation QStandardPaths__StandardLocation = 15

//
const QStandardPaths__GenericConfigLocation QStandardPaths__StandardLocation = 16

//
const QStandardPaths__AppDataLocation QStandardPaths__StandardLocation = 17

//
const QStandardPaths__AppConfigLocation QStandardPaths__StandardLocation = 18

//
const QStandardPaths__AppLocalDataLocation QStandardPaths__StandardLocation = 9

/*


 */
type QStandardPaths__LocateOption = int

//
const QStandardPaths__LocateFile QStandardPaths__LocateOption = 0

//
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
