package qtcore

// /usr/include/qt/QtCore/qfileinfo.h
// #include <qfileinfo.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QFileInfo struct {
	*qtrt.CObject
}

func (this *QFileInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFileInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFileInfoFromPointer(cthis unsafe.Pointer) *QFileInfo {
	return &QFileInfo{&qtrt.CObject{cthis}}
}
func (*QFileInfo) NewFromPointer(cthis unsafe.Pointer) *QFileInfo {
	return NewQFileInfoFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfileinfo.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileInfo()
func NewQFileInfo() *QFileInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFileInfo)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFileInfo(const QString &)
func NewQFileInfo_1(file string) *QFileInfo {
	var tmpArg0 = NewQString_5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFileInfo)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:64
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QFileInfo(const QFile &)
func NewQFileInfo_2(file *QFile) *QFileInfo {
	var convArg0 = file.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoC2ERK5QFile", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFileInfo)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:65
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QFileInfo(const QDir &, const QString &)
func NewQFileInfo_3(dir *QDir, file string) *QFileInfo {
	var convArg0 = dir.GetCthis()
	var tmpArg1 = NewQString_5(file)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoC2ERK4QDirRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFileInfo)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QFileInfo()
func DeleteQFileInfo(this *QFileInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfileinfo.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QFileInfo &)
func (this *QFileInfo) Swap(other *QFileInfo) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFile(const QString &)
func (this *QFileInfo) SetFile(file string) {
	var tmpArg0 = NewQString_5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFile(const QFile &)
func (this *QFileInfo) SetFile_1(file *QFile) {
	var convArg0 = file.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK5QFile", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:82
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setFile(const QDir &, const QString &)
func (this *QFileInfo) SetFile_2(dir *QDir, file string) {
	var convArg0 = dir.GetCthis()
	var tmpArg1 = NewQString_5(file)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK4QDirRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exists()
func (this *QFileInfo) Exists() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo6existsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:84
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool exists(const QString &)
func (this *QFileInfo) Exists_1(file string) bool {
	var tmpArg0 = NewQString_5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo6existsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFileInfo_Exists_1(file string) bool {
	var nilthis *QFileInfo
	rv := nilthis.Exists_1(file)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh()
func (this *QFileInfo) Refresh() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo7refreshEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath()
func (this *QFileInfo) FilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8filePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QString absoluteFilePath()
func (this *QFileInfo) AbsoluteFilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo16absoluteFilePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QString canonicalFilePath()
func (this *QFileInfo) CanonicalFilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo17canonicalFilePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QFileInfo) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QString baseName()
func (this *QFileInfo) BaseName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8baseNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QString completeBaseName()
func (this *QFileInfo) CompleteBaseName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo16completeBaseNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QString suffix()
func (this *QFileInfo) Suffix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo6suffixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QString bundleName()
func (this *QFileInfo) BundleName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10bundleNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QString completeSuffix()
func (this *QFileInfo) CompleteSuffix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo14completeSuffixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path()
func (this *QFileInfo) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QString absolutePath()
func (this *QFileInfo) AbsolutePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo12absolutePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QString canonicalPath()
func (this *QFileInfo) CanonicalPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo13canonicalPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QDir dir()
func (this *QFileInfo) Dir() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo3dirEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QDir absoluteDir()
func (this *QFileInfo) AbsoluteDir() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo11absoluteDirEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadable()
func (this *QFileInfo) IsReadable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10isReadableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWritable()
func (this *QFileInfo) IsWritable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10isWritableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExecutable()
func (this *QFileInfo) IsExecutable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo12isExecutableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isHidden()
func (this *QFileInfo) IsHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8isHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNativePath()
func (this *QFileInfo) IsNativePath() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo12isNativePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRelative()
func (this *QFileInfo) IsRelative() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10isRelativeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAbsolute()
func (this *QFileInfo) IsAbsolute() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10isAbsoluteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool makeAbsolute()
func (this *QFileInfo) MakeAbsolute() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo12makeAbsoluteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFile()
func (this *QFileInfo) IsFile() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo6isFileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDir()
func (this *QFileInfo) IsDir() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo5isDirEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSymLink()
func (this *QFileInfo) IsSymLink() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo9isSymLinkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRoot()
func (this *QFileInfo) IsRoot() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo6isRootEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBundle()
func (this *QFileInfo) IsBundle() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8isBundleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readLink()
func (this *QFileInfo) ReadLink() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8readLinkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString symLinkTarget()
func (this *QFileInfo) SymLinkTarget() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo13symLinkTargetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QString owner()
func (this *QFileInfo) Owner() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo5ownerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:123
// index:0
// Public Visibility=Default Availability=Available
// [4] uint ownerId()
func (this *QFileInfo) OwnerId() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo7ownerIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qfileinfo.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QString group()
func (this *QFileInfo) Group() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo5groupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:125
// index:0
// Public Visibility=Default Availability=Available
// [4] uint groupId()
func (this *QFileInfo) GroupId() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo7groupIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qfileinfo.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool permission(QFile::Permissions)
func (this *QFileInfo) Permission(permissions int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10permissionE6QFlagsIN11QFileDevice10PermissionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), permissions)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] QFile::Permissions permissions()
func (this *QFileInfo) Permissions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo11permissionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 size()
func (this *QFileInfo) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qfileinfo.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime created()
func (this *QFileInfo) Created() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo7createdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime birthTime()
func (this *QFileInfo) BirthTime() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo9birthTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime metadataChangeTime()
func (this *QFileInfo) MetadataChangeTime() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo18metadataChangeTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime lastModified()
func (this *QFileInfo) LastModified() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo12lastModifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:140
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime lastRead()
func (this *QFileInfo) LastRead() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8lastReadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime fileTime(QFile::FileTime)
func (this *QFileInfo) FileTime(time int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8fileTimeEN11QFileDevice8FileTimeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), time)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool caching()
func (this *QFileInfo) Caching() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo7cachingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCaching(_Bool)
func (this *QFileInfo) SetCaching(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo10setCachingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
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
		qtrt.KeepMe()
	}
}

//  keep block end
