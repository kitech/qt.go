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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qfileinfo.h:60
// index:0
// void QFileInfo(class QFileInfoPrivate *)
func NewQFileInfo(d unsafe.Pointer) *QFileInfo {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoC2EP16QFileInfoPrivate", ffiqt.FFI_TYPE_VOID, cthis, d)
	gopp.ErrPrint(err, rv)
	return &QFileInfo{cthis}
}

// /usr/include/qt/QtCore/qfileinfo.h:62
// index:1
// void QFileInfo()
func NewQFileInfo_1() *QFileInfo {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QFileInfo{cthis}
}

// /usr/include/qt/QtCore/qfileinfo.h:63
// index:2
// void QFileInfo(const class QString &)
func NewQFileInfo_2(file unsafe.Pointer) *QFileInfo {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, file)
	gopp.ErrPrint(err, rv)
	return &QFileInfo{cthis}
}

// /usr/include/qt/QtCore/qfileinfo.h:64
// index:3
// void QFileInfo(const class QFile &)
func NewQFileInfo_3(file unsafe.Pointer) *QFileInfo {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoC2ERK5QFile", ffiqt.FFI_TYPE_VOID, cthis, file)
	gopp.ErrPrint(err, rv)
	return &QFileInfo{cthis}
}

// /usr/include/qt/QtCore/qfileinfo.h:65
// index:4
// void QFileInfo(const class QDir &, const class QString &)
func NewQFileInfo_4(dir unsafe.Pointer, file unsafe.Pointer) *QFileInfo {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoC2ERK4QDirRK7QString", ffiqt.FFI_TYPE_VOID, cthis, dir, file)
	gopp.ErrPrint(err, rv)
	return &QFileInfo{cthis}
}

// /usr/include/qt/QtCore/qfileinfo.h:67
// index:0
// void ~QFileInfo()
func DeleteQFileInfo(*QFileInfo) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfoD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:74
// index:0
// inline
// void swap(class QFileInfo &)
func (this *QFileInfo) Swap(other unsafe.Pointer) {
	// 0: (, QFileInfo & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:80
// index:0
// void setFile(const class QString &)
func (this *QFileInfo) SetFile(file unsafe.Pointer) {
	// 0: (, const QString & file), (file)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, file)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:81
// index:1
// void setFile(const class QFile &)
func (this *QFileInfo) SetFile_1(file unsafe.Pointer) {
	// 1: (, const QFile & file), (file)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK5QFile", ffiqt.FFI_TYPE_VOID, this.cthis, file)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:82
// index:2
// void setFile(const class QDir &, const class QString &)
func (this *QFileInfo) SetFile_2(dir unsafe.Pointer, file unsafe.Pointer) {
	// 2: (, const QDir & dir, const QString & file), (dir, file)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK4QDirRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, dir, file)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:83
// index:0
// bool exists()
func (this *QFileInfo) Exists() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo6existsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:84
// index:1
// static
// bool exists(const class QString &)
func (this *QFileInfo) Exists_1(file unsafe.Pointer) {
	// 1: (const QString & file), (file)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo6existsERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFileInfo_Exists_1(file unsafe.Pointer) {
	// 1: (const QString & file), (file)
	var nilthis *QFileInfo
	nilthis.Exists_1(file)
}

// /usr/include/qt/QtCore/qfileinfo.h:85
// index:0
// void refresh()
func (this *QFileInfo) Refresh() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo7refreshEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:87
// index:0
// QString filePath()
func (this *QFileInfo) FilePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8filePathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:88
// index:0
// QString absoluteFilePath()
func (this *QFileInfo) AbsoluteFilePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo16absoluteFilePathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:89
// index:0
// QString canonicalFilePath()
func (this *QFileInfo) CanonicalFilePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo17canonicalFilePathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:90
// index:0
// QString fileName()
func (this *QFileInfo) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8fileNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:91
// index:0
// QString baseName()
func (this *QFileInfo) BaseName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8baseNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:92
// index:0
// QString completeBaseName()
func (this *QFileInfo) CompleteBaseName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo16completeBaseNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:93
// index:0
// QString suffix()
func (this *QFileInfo) Suffix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo6suffixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:94
// index:0
// QString bundleName()
func (this *QFileInfo) BundleName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo10bundleNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:95
// index:0
// QString completeSuffix()
func (this *QFileInfo) CompleteSuffix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo14completeSuffixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:97
// index:0
// QString path()
func (this *QFileInfo) Path() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo4pathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:98
// index:0
// QString absolutePath()
func (this *QFileInfo) AbsolutePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo12absolutePathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:99
// index:0
// QString canonicalPath()
func (this *QFileInfo) CanonicalPath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo13canonicalPathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:100
// index:0
// QDir dir()
func (this *QFileInfo) Dir() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo3dirEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:101
// index:0
// QDir absoluteDir()
func (this *QFileInfo) AbsoluteDir() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo11absoluteDirEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:103
// index:0
// bool isReadable()
func (this *QFileInfo) IsReadable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo10isReadableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:104
// index:0
// bool isWritable()
func (this *QFileInfo) IsWritable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo10isWritableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:105
// index:0
// bool isExecutable()
func (this *QFileInfo) IsExecutable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo12isExecutableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:106
// index:0
// bool isHidden()
func (this *QFileInfo) IsHidden() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8isHiddenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:107
// index:0
// bool isNativePath()
func (this *QFileInfo) IsNativePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo12isNativePathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:109
// index:0
// bool isRelative()
func (this *QFileInfo) IsRelative() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo10isRelativeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:110
// index:0
// inline
// bool isAbsolute()
func (this *QFileInfo) IsAbsolute() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo10isAbsoluteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:111
// index:0
// bool makeAbsolute()
func (this *QFileInfo) MakeAbsolute() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo12makeAbsoluteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:113
// index:0
// bool isFile()
func (this *QFileInfo) IsFile() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo6isFileEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:114
// index:0
// bool isDir()
func (this *QFileInfo) IsDir() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo5isDirEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:115
// index:0
// bool isSymLink()
func (this *QFileInfo) IsSymLink() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo9isSymLinkEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:116
// index:0
// bool isRoot()
func (this *QFileInfo) IsRoot() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo6isRootEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:117
// index:0
// bool isBundle()
func (this *QFileInfo) IsBundle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8isBundleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:119
// index:0
// QString readLink()
func (this *QFileInfo) ReadLink() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8readLinkEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:120
// index:0
// inline
// QString symLinkTarget()
func (this *QFileInfo) SymLinkTarget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo13symLinkTargetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:122
// index:0
// QString owner()
func (this *QFileInfo) Owner() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo5ownerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:123
// index:0
// uint ownerId()
func (this *QFileInfo) OwnerId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo7ownerIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:124
// index:0
// QString group()
func (this *QFileInfo) Group() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo5groupEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:125
// index:0
// uint groupId()
func (this *QFileInfo) GroupId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo7groupIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:128
// index:0
// QFile::Permissions permissions()
func (this *QFileInfo) Permissions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo11permissionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:130
// index:0
// qint64 size()
func (this *QFileInfo) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:135
// index:0
// QDateTime created()
func (this *QFileInfo) Created() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo7createdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:137
// index:0
// QDateTime birthTime()
func (this *QFileInfo) BirthTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo9birthTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:138
// index:0
// QDateTime metadataChangeTime()
func (this *QFileInfo) MetadataChangeTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo18metadataChangeTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:139
// index:0
// QDateTime lastModified()
func (this *QFileInfo) LastModified() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo12lastModifiedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:140
// index:0
// QDateTime lastRead()
func (this *QFileInfo) LastRead() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8lastReadEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:141
// index:0
// QDateTime fileTime(class QFile::FileTime)
func (this *QFileInfo) FileTime(time int) {
	// 0: (, QFile::FileTime time), (&time)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo8fileTimeEN11QFileDevice8FileTimeE", ffiqt.FFI_TYPE_VOID, this.cthis, &time)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:143
// index:0
// bool caching()
func (this *QFileInfo) Caching() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFileInfo7cachingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:144
// index:0
// void setCaching(_Bool)
func (this *QFileInfo) SetCaching(on bool) {
	// 0: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFileInfo10setCachingEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

//  body block end
