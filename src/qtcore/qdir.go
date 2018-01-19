//  header block begin
// /usr/include/qt/QtCore/qdir.h
// #include <qdir.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
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
type QDir struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qdir.h:102
// index:0
// void QDir(const class QString &)
func NewQDir(path unsafe.Pointer) *QDir {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDirC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, path)
	gopp.ErrPrint(err, rv)
	return &QDir{cthis}
}

// /usr/include/qt/QtCore/qdir.h:105
// index:0
// void ~QDir()
func DeleteQDir(*QDir) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDirD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:113
// index:0
// inline
// void swap(class QDir &)
func (this *QDir) Swap(other unsafe.Pointer) {
	// 0: (, QDir & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:116
// index:0
// void setPath(const class QString &)
func (this *QDir) SetPath(path unsafe.Pointer) {
	// 0: (, const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir7setPathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:117
// index:0
// QString path()
func (this *QDir) Path() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir4pathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:118
// index:0
// QString absolutePath()
func (this *QDir) AbsolutePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir12absolutePathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:119
// index:0
// QString canonicalPath()
func (this *QDir) CanonicalPath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir13canonicalPathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:121
// index:0
// static
// void addResourceSearchPath(const class QString &)
func (this *QDir) AddResourceSearchPath(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir21addResourceSearchPathERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_AddResourceSearchPath(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	var nilthis *QDir
	nilthis.AddResourceSearchPath(path)
}

// /usr/include/qt/QtCore/qdir.h:123
// index:0
// static
// void setSearchPaths(const class QString &, const class QStringList &)
func (this *QDir) SetSearchPaths(prefix unsafe.Pointer, searchPaths unsafe.Pointer) {
	// 0: (const QString & prefix, const QStringList & searchPaths), (prefix, searchPaths)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14setSearchPathsERK7QStringRK11QStringList", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_SetSearchPaths(prefix unsafe.Pointer, searchPaths unsafe.Pointer) {
	// 0: (const QString & prefix, const QStringList & searchPaths), (prefix, searchPaths)
	var nilthis *QDir
	nilthis.SetSearchPaths(prefix, searchPaths)
}

// /usr/include/qt/QtCore/qdir.h:124
// index:0
// static
// void addSearchPath(const class QString &, const class QString &)
func (this *QDir) AddSearchPath(prefix unsafe.Pointer, path unsafe.Pointer) {
	// 0: (const QString & prefix, const QString & path), (prefix, path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir13addSearchPathERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_AddSearchPath(prefix unsafe.Pointer, path unsafe.Pointer) {
	// 0: (const QString & prefix, const QString & path), (prefix, path)
	var nilthis *QDir
	nilthis.AddSearchPath(prefix, path)
}

// /usr/include/qt/QtCore/qdir.h:125
// index:0
// static
// QStringList searchPaths(const class QString &)
func (this *QDir) SearchPaths(prefix unsafe.Pointer) {
	// 0: (const QString & prefix), (prefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir11searchPathsERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_SearchPaths(prefix unsafe.Pointer) {
	// 0: (const QString & prefix), (prefix)
	var nilthis *QDir
	nilthis.SearchPaths(prefix)
}

// /usr/include/qt/QtCore/qdir.h:127
// index:0
// QString dirName()
func (this *QDir) DirName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir7dirNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:128
// index:0
// QString filePath(const class QString &)
func (this *QDir) FilePath(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir8filePathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:129
// index:0
// QString absoluteFilePath(const class QString &)
func (this *QDir) AbsoluteFilePath(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir16absoluteFilePathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:130
// index:0
// QString relativeFilePath(const class QString &)
func (this *QDir) RelativeFilePath(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir16relativeFilePathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:132
// index:0
// static
// QString toNativeSeparators(const class QString &)
func (this *QDir) ToNativeSeparators(pathName unsafe.Pointer) {
	// 0: (const QString & pathName), (pathName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir18toNativeSeparatorsERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_ToNativeSeparators(pathName unsafe.Pointer) {
	// 0: (const QString & pathName), (pathName)
	var nilthis *QDir
	nilthis.ToNativeSeparators(pathName)
}

// /usr/include/qt/QtCore/qdir.h:133
// index:0
// static
// QString fromNativeSeparators(const class QString &)
func (this *QDir) FromNativeSeparators(pathName unsafe.Pointer) {
	// 0: (const QString & pathName), (pathName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir20fromNativeSeparatorsERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_FromNativeSeparators(pathName unsafe.Pointer) {
	// 0: (const QString & pathName), (pathName)
	var nilthis *QDir
	nilthis.FromNativeSeparators(pathName)
}

// /usr/include/qt/QtCore/qdir.h:135
// index:0
// bool cd(const class QString &)
func (this *QDir) Cd(dirName unsafe.Pointer) {
	// 0: (, const QString & dirName), (dirName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir2cdERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, dirName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:136
// index:0
// bool cdUp()
func (this *QDir) CdUp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4cdUpEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:138
// index:0
// QStringList nameFilters()
func (this *QDir) NameFilters() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir11nameFiltersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:139
// index:0
// void setNameFilters(const class QStringList &)
func (this *QDir) SetNameFilters(nameFilters unsafe.Pointer) {
	// 0: (, const QStringList & nameFilters), (nameFilters)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14setNameFiltersERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, nameFilters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:141
// index:0
// QDir::Filters filter()
func (this *QDir) Filter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6filterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:143
// index:0
// QDir::SortFlags sorting()
func (this *QDir) Sorting() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir7sortingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:146
// index:0
// uint count()
func (this *QDir) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:151
// index:0
// static
// QStringList nameFiltersFromString(const class QString &)
func (this *QDir) NameFiltersFromString(nameFilter unsafe.Pointer) {
	// 0: (const QString & nameFilter), (nameFilter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir21nameFiltersFromStringERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_NameFiltersFromString(nameFilter unsafe.Pointer) {
	// 0: (const QString & nameFilter), (nameFilter)
	var nilthis *QDir
	nilthis.NameFiltersFromString(nameFilter)
}

// /usr/include/qt/QtCore/qdir.h:161
// index:0
// bool mkdir(const class QString &)
func (this *QDir) Mkdir(dirName unsafe.Pointer) {
	// 0: (, const QString & dirName), (dirName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir5mkdirERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, dirName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:162
// index:0
// bool rmdir(const class QString &)
func (this *QDir) Rmdir(dirName unsafe.Pointer) {
	// 0: (, const QString & dirName), (dirName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir5rmdirERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, dirName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:163
// index:0
// bool mkpath(const class QString &)
func (this *QDir) Mkpath(dirPath unsafe.Pointer) {
	// 0: (, const QString & dirPath), (dirPath)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6mkpathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, dirPath)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:164
// index:0
// bool rmpath(const class QString &)
func (this *QDir) Rmpath(dirPath unsafe.Pointer) {
	// 0: (, const QString & dirPath), (dirPath)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6rmpathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, dirPath)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:166
// index:0
// bool removeRecursively()
func (this *QDir) RemoveRecursively() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir17removeRecursivelyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:168
// index:0
// bool isReadable()
func (this *QDir) IsReadable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir10isReadableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:169
// index:0
// bool exists()
func (this *QDir) Exists() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6existsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:183
// index:1
// bool exists(const class QString &)
func (this *QDir) Exists_1(name unsafe.Pointer) {
	// 1: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6existsERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:170
// index:0
// bool isRoot()
func (this *QDir) IsRoot() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6isRootEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:172
// index:0
// static
// bool isRelativePath(const class QString &)
func (this *QDir) IsRelativePath(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14isRelativePathERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_IsRelativePath(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	var nilthis *QDir
	nilthis.IsRelativePath(path)
}

// /usr/include/qt/QtCore/qdir.h:173
// index:0
// static inline
// bool isAbsolutePath(const class QString &)
func (this *QDir) IsAbsolutePath(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14isAbsolutePathERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_IsAbsolutePath(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	var nilthis *QDir
	nilthis.IsAbsolutePath(path)
}

// /usr/include/qt/QtCore/qdir.h:174
// index:0
// bool isRelative()
func (this *QDir) IsRelative() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir10isRelativeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:175
// index:0
// inline
// bool isAbsolute()
func (this *QDir) IsAbsolute() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir10isAbsoluteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:176
// index:0
// bool makeAbsolute()
func (this *QDir) MakeAbsolute() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir12makeAbsoluteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:181
// index:0
// bool remove(const class QString &)
func (this *QDir) Remove(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir6removeERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:182
// index:0
// bool rename(const class QString &, const class QString &)
func (this *QDir) Rename(oldName unsafe.Pointer, newName unsafe.Pointer) {
	// 0: (, const QString & oldName, const QString & newName), (oldName, newName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir6renameERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.cthis, oldName, newName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:185
// index:0
// static
// QFileInfoList drives()
func (this *QDir) Drives() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir6drivesEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_Drives() {
	// 0: (), ()
	var nilthis *QDir
	nilthis.Drives()
}

// /usr/include/qt/QtCore/qdir.h:187
// index:0
// static inline
// QChar listSeparator()
func (this *QDir) ListSeparator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir13listSeparatorEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_ListSeparator() {
	// 0: (), ()
	var nilthis *QDir
	nilthis.ListSeparator()
}

// /usr/include/qt/QtCore/qdir.h:196
// index:0
// static
// QChar separator()
func (this *QDir) Separator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir9separatorEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_Separator() {
	// 0: (), ()
	var nilthis *QDir
	nilthis.Separator()
}

// /usr/include/qt/QtCore/qdir.h:198
// index:0
// static
// bool setCurrent(const class QString &)
func (this *QDir) SetCurrent(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir10setCurrentERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_SetCurrent(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	var nilthis *QDir
	nilthis.SetCurrent(path)
}

// /usr/include/qt/QtCore/qdir.h:199
// index:0
// static inline
// QDir current()
func (this *QDir) Current() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir7currentEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_Current() {
	// 0: (), ()
	var nilthis *QDir
	nilthis.Current()
}

// /usr/include/qt/QtCore/qdir.h:200
// index:0
// static
// QString currentPath()
func (this *QDir) CurrentPath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir11currentPathEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_CurrentPath() {
	// 0: (), ()
	var nilthis *QDir
	nilthis.CurrentPath()
}

// /usr/include/qt/QtCore/qdir.h:202
// index:0
// static inline
// QDir home()
func (this *QDir) Home() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4homeEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_Home() {
	// 0: (), ()
	var nilthis *QDir
	nilthis.Home()
}

// /usr/include/qt/QtCore/qdir.h:203
// index:0
// static
// QString homePath()
func (this *QDir) HomePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir8homePathEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_HomePath() {
	// 0: (), ()
	var nilthis *QDir
	nilthis.HomePath()
}

// /usr/include/qt/QtCore/qdir.h:204
// index:0
// static inline
// QDir root()
func (this *QDir) Root() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4rootEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_Root() {
	// 0: (), ()
	var nilthis *QDir
	nilthis.Root()
}

// /usr/include/qt/QtCore/qdir.h:205
// index:0
// static
// QString rootPath()
func (this *QDir) RootPath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir8rootPathEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_RootPath() {
	// 0: (), ()
	var nilthis *QDir
	nilthis.RootPath()
}

// /usr/include/qt/QtCore/qdir.h:206
// index:0
// static inline
// QDir temp()
func (this *QDir) Temp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4tempEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_Temp() {
	// 0: (), ()
	var nilthis *QDir
	nilthis.Temp()
}

// /usr/include/qt/QtCore/qdir.h:207
// index:0
// static
// QString tempPath()
func (this *QDir) TempPath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir8tempPathEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_TempPath() {
	// 0: (), ()
	var nilthis *QDir
	nilthis.TempPath()
}

// /usr/include/qt/QtCore/qdir.h:210
// index:0
// static
// bool match(const class QStringList &, const class QString &)
func (this *QDir) Match(filters unsafe.Pointer, fileName unsafe.Pointer) {
	// 0: (const QStringList & filters, const QString & fileName), (filters, fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir5matchERK11QStringListRK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_Match(filters unsafe.Pointer, fileName unsafe.Pointer) {
	// 0: (const QStringList & filters, const QString & fileName), (filters, fileName)
	var nilthis *QDir
	nilthis.Match(filters, fileName)
}

// /usr/include/qt/QtCore/qdir.h:211
// index:1
// static
// bool match(const class QString &, const class QString &)
func (this *QDir) Match_1(filter unsafe.Pointer, fileName unsafe.Pointer) {
	// 1: (const QString & filter, const QString & fileName), (filter, fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir5matchERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_Match_1(filter unsafe.Pointer, fileName unsafe.Pointer) {
	// 1: (const QString & filter, const QString & fileName), (filter, fileName)
	var nilthis *QDir
	nilthis.Match_1(filter, fileName)
}

// /usr/include/qt/QtCore/qdir.h:214
// index:0
// static
// QString cleanPath(const class QString &)
func (this *QDir) CleanPath(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir9cleanPathERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDir_CleanPath(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	var nilthis *QDir
	nilthis.CleanPath(path)
}

// /usr/include/qt/QtCore/qdir.h:215
// index:0
// void refresh()
func (this *QDir) Refresh() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir7refreshEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
