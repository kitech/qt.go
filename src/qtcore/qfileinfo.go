//  header block begin
// /usr/include/qt/QtCore/qfileinfo.h
// #include <qfileinfo.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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
type QFileInfo struct {
	*qtrt.CObject
}

func (this *QFileInfo) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQFileInfoFromPointer(cthis unsafe.Pointer) *QFileInfo {
	return &QFileInfo{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qfileinfo.h:60
// index:0
// Public
// void QFileInfo(class QFileInfoPrivate *)
func NewQFileInfo(d unsafe.Pointer) *QFileInfo {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoC2EP16QFileInfoPrivate", ffiqt.FFI_TYPE_VOID, cthis, d)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:62
// index:1
// Public
// void QFileInfo()
func NewQFileInfo_1() *QFileInfo {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:63
// index:2
// Public
// void QFileInfo(const class QString &)
func NewQFileInfo_2(file *QString) *QFileInfo {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:64
// index:3
// Public
// void QFileInfo(const class QFile &)
func NewQFileInfo_3(file *QFile) *QFileInfo {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoC2ERK5QFile", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:65
// index:4
// Public
// void QFileInfo(const class QDir &, const class QString &)
func NewQFileInfo_4(dir *QDir, file *QString) *QFileInfo {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = dir.GetCthis()
	var convArg1 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoC2ERK4QDirRK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:67
// index:0
// Public
// void ~QFileInfo()
func DeleteQFileInfo(*QFileInfo) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:74
// index:0
// Public inline
// void swap(class QFileInfo &)
func (this *QFileInfo) Swap(other *QFileInfo) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:80
// index:0
// Public
// void setFile(const class QString &)
func (this *QFileInfo) SetFile(file *QString) {
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:81
// index:1
// Public
// void setFile(const class QFile &)
func (this *QFileInfo) SetFile_1(file *QFile) {
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK5QFile", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:82
// index:2
// Public
// void setFile(const class QDir &, const class QString &)
func (this *QFileInfo) SetFile_2(dir *QDir, file *QString) {
	var convArg0 = dir.GetCthis()
	var convArg1 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK4QDirRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:83
// index:0
// Public
// bool exists()
func (this *QFileInfo) Exists() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo6existsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:84
// index:1
// Public static
// bool exists(const class QString &)
func (this *QFileInfo) Exists_1(file *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo6existsERK7QString", ffiqt.FFI_TYPE_POINTER, file)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFileInfo_Exists_1(file *QString) {
	var nilthis *QFileInfo
	nilthis.Exists_1(file)
}

// /usr/include/qt/QtCore/qfileinfo.h:85
// index:0
// Public
// void refresh()
func (this *QFileInfo) Refresh() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo7refreshEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:87
// index:0
// Public
// QString filePath()
func (this *QFileInfo) FilePath() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8filePathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:88
// index:0
// Public
// QString absoluteFilePath()
func (this *QFileInfo) AbsoluteFilePath() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo16absoluteFilePathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:89
// index:0
// Public
// QString canonicalFilePath()
func (this *QFileInfo) CanonicalFilePath() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo17canonicalFilePathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:90
// index:0
// Public
// QString fileName()
func (this *QFileInfo) FileName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:91
// index:0
// Public
// QString baseName()
func (this *QFileInfo) BaseName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8baseNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:92
// index:0
// Public
// QString completeBaseName()
func (this *QFileInfo) CompleteBaseName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo16completeBaseNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:93
// index:0
// Public
// QString suffix()
func (this *QFileInfo) Suffix() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo6suffixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:94
// index:0
// Public
// QString bundleName()
func (this *QFileInfo) BundleName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo10bundleNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:95
// index:0
// Public
// QString completeSuffix()
func (this *QFileInfo) CompleteSuffix() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo14completeSuffixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:97
// index:0
// Public
// QString path()
func (this *QFileInfo) Path() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo4pathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:98
// index:0
// Public
// QString absolutePath()
func (this *QFileInfo) AbsolutePath() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo12absolutePathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:99
// index:0
// Public
// QString canonicalPath()
func (this *QFileInfo) CanonicalPath() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo13canonicalPathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:100
// index:0
// Public
// QDir dir()
func (this *QFileInfo) Dir() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo3dirEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:101
// index:0
// Public
// QDir absoluteDir()
func (this *QFileInfo) AbsoluteDir() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo11absoluteDirEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:103
// index:0
// Public
// bool isReadable()
func (this *QFileInfo) IsReadable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo10isReadableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:104
// index:0
// Public
// bool isWritable()
func (this *QFileInfo) IsWritable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo10isWritableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:105
// index:0
// Public
// bool isExecutable()
func (this *QFileInfo) IsExecutable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo12isExecutableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:106
// index:0
// Public
// bool isHidden()
func (this *QFileInfo) IsHidden() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8isHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:107
// index:0
// Public
// bool isNativePath()
func (this *QFileInfo) IsNativePath() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo12isNativePathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:109
// index:0
// Public
// bool isRelative()
func (this *QFileInfo) IsRelative() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo10isRelativeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:110
// index:0
// Public inline
// bool isAbsolute()
func (this *QFileInfo) IsAbsolute() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo10isAbsoluteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:111
// index:0
// Public
// bool makeAbsolute()
func (this *QFileInfo) MakeAbsolute() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo12makeAbsoluteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:113
// index:0
// Public
// bool isFile()
func (this *QFileInfo) IsFile() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo6isFileEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:114
// index:0
// Public
// bool isDir()
func (this *QFileInfo) IsDir() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo5isDirEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:115
// index:0
// Public
// bool isSymLink()
func (this *QFileInfo) IsSymLink() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo9isSymLinkEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:116
// index:0
// Public
// bool isRoot()
func (this *QFileInfo) IsRoot() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo6isRootEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:117
// index:0
// Public
// bool isBundle()
func (this *QFileInfo) IsBundle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8isBundleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:119
// index:0
// Public
// QString readLink()
func (this *QFileInfo) ReadLink() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8readLinkEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:120
// index:0
// Public inline
// QString symLinkTarget()
func (this *QFileInfo) SymLinkTarget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo13symLinkTargetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:122
// index:0
// Public
// QString owner()
func (this *QFileInfo) Owner() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo5ownerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:123
// index:0
// Public
// uint ownerId()
func (this *QFileInfo) OwnerId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo7ownerIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:124
// index:0
// Public
// QString group()
func (this *QFileInfo) Group() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo5groupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:125
// index:0
// Public
// uint groupId()
func (this *QFileInfo) GroupId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo7groupIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:128
// index:0
// Public
// QFile::Permissions permissions()
func (this *QFileInfo) Permissions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo11permissionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:130
// index:0
// Public
// qint64 size()
func (this *QFileInfo) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:135
// index:0
// Public
// QDateTime created()
func (this *QFileInfo) Created() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo7createdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:137
// index:0
// Public
// QDateTime birthTime()
func (this *QFileInfo) BirthTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo9birthTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:138
// index:0
// Public
// QDateTime metadataChangeTime()
func (this *QFileInfo) MetadataChangeTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo18metadataChangeTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:139
// index:0
// Public
// QDateTime lastModified()
func (this *QFileInfo) LastModified() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo12lastModifiedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:140
// index:0
// Public
// QDateTime lastRead()
func (this *QFileInfo) LastRead() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8lastReadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:141
// index:0
// Public
// QDateTime fileTime(class QFile::FileTime)
func (this *QFileInfo) FileTime(time int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8fileTimeEN11QFileDevice8FileTimeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &time)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:143
// index:0
// Public
// bool caching()
func (this *QFileInfo) Caching() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo7cachingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:144
// index:0
// Public
// void setCaching(_Bool)
func (this *QFileInfo) SetCaching(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo10setCachingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

//  body block end
