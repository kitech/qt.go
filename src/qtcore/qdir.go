package qtcore

// /usr/include/qt/QtCore/qdir.h
// #include <qdir.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 55
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
	*qtrt.CObject
}

func (this *QDir) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQDirFromPointer(cthis unsafe.Pointer) *QDir {
	return &QDir{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qdir.h:102
// index:0
// Public
// void QDir(const class QString &)
func NewQDir(path *QString) *QDir {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDirC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdir.h:105
// index:0
// Public
// void ~QDir()
func DeleteQDir(*QDir) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDirD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:113
// index:0
// Public inline
// void swap(class QDir &)
func (this *QDir) Swap(other *QDir) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:116
// index:0
// Public
// void setPath(const class QString &)
func (this *QDir) SetPath(path *QString) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir7setPathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:117
// index:0
// Public
// QString path()
func (this *QDir) Path() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir4pathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:118
// index:0
// Public
// QString absolutePath()
func (this *QDir) AbsolutePath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir12absolutePathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:119
// index:0
// Public
// QString canonicalPath()
func (this *QDir) CanonicalPath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir13canonicalPathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:121
// index:0
// Public static
// void addResourceSearchPath(const class QString &)
func (this *QDir) AddResourceSearchPath(path *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir21addResourceSearchPathERK7QString", ffiqt.FFI_TYPE_POINTER, path)
	gopp.ErrPrint(err, rv)
}
func QDir_AddResourceSearchPath(path *QString) {
	var nilthis *QDir
	nilthis.AddResourceSearchPath(path)
}

// /usr/include/qt/QtCore/qdir.h:123
// index:0
// Public static
// void setSearchPaths(const class QString &, const class QStringList &)
func (this *QDir) SetSearchPaths(prefix *QString, searchPaths *QStringList) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14setSearchPathsERK7QStringRK11QStringList", ffiqt.FFI_TYPE_POINTER, prefix, searchPaths)
	gopp.ErrPrint(err, rv)
}
func QDir_SetSearchPaths(prefix *QString, searchPaths *QStringList) {
	var nilthis *QDir
	nilthis.SetSearchPaths(prefix, searchPaths)
}

// /usr/include/qt/QtCore/qdir.h:124
// index:0
// Public static
// void addSearchPath(const class QString &, const class QString &)
func (this *QDir) AddSearchPath(prefix *QString, path *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir13addSearchPathERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, prefix, path)
	gopp.ErrPrint(err, rv)
}
func QDir_AddSearchPath(prefix *QString, path *QString) {
	var nilthis *QDir
	nilthis.AddSearchPath(prefix, path)
}

// /usr/include/qt/QtCore/qdir.h:127
// index:0
// Public
// QString dirName()
func (this *QDir) DirName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir7dirNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:128
// index:0
// Public
// QString filePath(const class QString &)
func (this *QDir) FilePath(fileName *QString) *QString /*123*/ {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir8filePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:129
// index:0
// Public
// QString absoluteFilePath(const class QString &)
func (this *QDir) AbsoluteFilePath(fileName *QString) *QString /*123*/ {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir16absoluteFilePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:130
// index:0
// Public
// QString relativeFilePath(const class QString &)
func (this *QDir) RelativeFilePath(fileName *QString) *QString /*123*/ {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir16relativeFilePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:132
// index:0
// Public static
// QString toNativeSeparators(const class QString &)
func (this *QDir) ToNativeSeparators(pathName *QString) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir18toNativeSeparatorsERK7QString", ffiqt.FFI_TYPE_POINTER, pathName)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_ToNativeSeparators(pathName *QString) *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.ToNativeSeparators(pathName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:133
// index:0
// Public static
// QString fromNativeSeparators(const class QString &)
func (this *QDir) FromNativeSeparators(pathName *QString) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir20fromNativeSeparatorsERK7QString", ffiqt.FFI_TYPE_POINTER, pathName)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_FromNativeSeparators(pathName *QString) *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.FromNativeSeparators(pathName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:135
// index:0
// Public
// bool cd(const class QString &)
func (this *QDir) Cd(dirName *QString) bool {
	var convArg0 = dirName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir2cdERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:136
// index:0
// Public
// bool cdUp()
func (this *QDir) CdUp() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4cdUpEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:139
// index:0
// Public
// void setNameFilters(const class QStringList &)
func (this *QDir) SetNameFilters(nameFilters *QStringList) {
	var convArg0 = nameFilters.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14setNameFiltersERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:141
// index:0
// Public
// QDir::Filters filter()
func (this *QDir) Filter() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6filterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdir.h:143
// index:0
// Public
// QDir::SortFlags sorting()
func (this *QDir) Sorting() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir7sortingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdir.h:146
// index:0
// Public
// uint count()
func (this *QDir) Count() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdir.h:161
// index:0
// Public
// bool mkdir(const class QString &)
func (this *QDir) Mkdir(dirName *QString) bool {
	var convArg0 = dirName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir5mkdirERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:162
// index:0
// Public
// bool rmdir(const class QString &)
func (this *QDir) Rmdir(dirName *QString) bool {
	var convArg0 = dirName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir5rmdirERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:163
// index:0
// Public
// bool mkpath(const class QString &)
func (this *QDir) Mkpath(dirPath *QString) bool {
	var convArg0 = dirPath.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6mkpathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:164
// index:0
// Public
// bool rmpath(const class QString &)
func (this *QDir) Rmpath(dirPath *QString) bool {
	var convArg0 = dirPath.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6rmpathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:166
// index:0
// Public
// bool removeRecursively()
func (this *QDir) RemoveRecursively() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir17removeRecursivelyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:168
// index:0
// Public
// bool isReadable()
func (this *QDir) IsReadable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir10isReadableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:169
// index:0
// Public
// bool exists()
func (this *QDir) Exists() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6existsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:183
// index:1
// Public
// bool exists(const class QString &)
func (this *QDir) Exists_1(name *QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6existsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:170
// index:0
// Public
// bool isRoot()
func (this *QDir) IsRoot() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6isRootEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:172
// index:0
// Public static
// bool isRelativePath(const class QString &)
func (this *QDir) IsRelativePath(path *QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14isRelativePathERK7QString", ffiqt.FFI_TYPE_POINTER, path)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QDir_IsRelativePath(path *QString) bool {
	var nilthis *QDir
	rv := nilthis.IsRelativePath(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:173
// index:0
// Public static inline
// bool isAbsolutePath(const class QString &)
func (this *QDir) IsAbsolutePath(path *QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14isAbsolutePathERK7QString", ffiqt.FFI_TYPE_POINTER, path)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QDir_IsAbsolutePath(path *QString) bool {
	var nilthis *QDir
	rv := nilthis.IsAbsolutePath(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:174
// index:0
// Public
// bool isRelative()
func (this *QDir) IsRelative() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir10isRelativeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:175
// index:0
// Public inline
// bool isAbsolute()
func (this *QDir) IsAbsolute() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir10isAbsoluteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:176
// index:0
// Public
// bool makeAbsolute()
func (this *QDir) MakeAbsolute() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir12makeAbsoluteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:181
// index:0
// Public
// bool remove(const class QString &)
func (this *QDir) Remove(fileName *QString) bool {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir6removeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:182
// index:0
// Public
// bool rename(const class QString &, const class QString &)
func (this *QDir) Rename(oldName *QString, newName *QString) bool {
	var convArg0 = oldName.GetCthis()
	var convArg1 = newName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir6renameERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:187
// index:0
// Public static inline
// QChar listSeparator()
func (this *QDir) ListSeparator() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir13listSeparatorEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_ListSeparator() *QChar /*123*/ {
	var nilthis *QDir
	rv := nilthis.ListSeparator()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:196
// index:0
// Public static
// QChar separator()
func (this *QDir) Separator() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir9separatorEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_Separator() *QChar /*123*/ {
	var nilthis *QDir
	rv := nilthis.Separator()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:198
// index:0
// Public static
// bool setCurrent(const class QString &)
func (this *QDir) SetCurrent(path *QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir10setCurrentERK7QString", ffiqt.FFI_TYPE_POINTER, path)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QDir_SetCurrent(path *QString) bool {
	var nilthis *QDir
	rv := nilthis.SetCurrent(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:199
// index:0
// Public static inline
// QDir current()
func (this *QDir) Current() *QDir /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir7currentEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_Current() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Current()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:200
// index:0
// Public static
// QString currentPath()
func (this *QDir) CurrentPath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir11currentPathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_CurrentPath() *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.CurrentPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:202
// index:0
// Public static inline
// QDir home()
func (this *QDir) Home() *QDir /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4homeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_Home() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Home()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:203
// index:0
// Public static
// QString homePath()
func (this *QDir) HomePath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir8homePathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_HomePath() *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.HomePath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:204
// index:0
// Public static inline
// QDir root()
func (this *QDir) Root() *QDir /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4rootEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_Root() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Root()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:205
// index:0
// Public static
// QString rootPath()
func (this *QDir) RootPath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir8rootPathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_RootPath() *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.RootPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:206
// index:0
// Public static inline
// QDir temp()
func (this *QDir) Temp() *QDir /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4tempEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_Temp() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Temp()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:207
// index:0
// Public static
// QString tempPath()
func (this *QDir) TempPath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir8tempPathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_TempPath() *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.TempPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:210
// index:0
// Public static
// bool match(const class QStringList &, const class QString &)
func (this *QDir) Match(filters *QStringList, fileName *QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir5matchERK11QStringListRK7QString", ffiqt.FFI_TYPE_POINTER, filters, fileName)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QDir_Match(filters *QStringList, fileName *QString) bool {
	var nilthis *QDir
	rv := nilthis.Match(filters, fileName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:211
// index:1
// Public static
// bool match(const class QString &, const class QString &)
func (this *QDir) Match_1(filter *QString, fileName *QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir5matchERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, filter, fileName)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QDir_Match_1(filter *QString, fileName *QString) bool {
	var nilthis *QDir
	rv := nilthis.Match_1(filter, fileName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:214
// index:0
// Public static
// QString cleanPath(const class QString &)
func (this *QDir) CleanPath(path *QString) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir9cleanPathERK7QString", ffiqt.FFI_TYPE_POINTER, path)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDir_CleanPath(path *QString) *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.CleanPath(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:215
// index:0
// Public
// void refresh()
func (this *QDir) Refresh() {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir7refreshEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
