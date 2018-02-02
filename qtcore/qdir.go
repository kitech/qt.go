package qtcore

// /usr/include/qt/QtCore/qdir.h
// #include <qdir.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
func (this *QDir) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDirFromPointer(cthis unsafe.Pointer) *QDir {
	return &QDir{&qtrt.CObject{cthis}}
}
func (*QDir) NewFromPointer(cthis unsafe.Pointer) *QDir {
	return NewQDirFromPointer(cthis)
}

// /usr/include/qt/QtCore/qdir.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDir(const QString &)
func NewQDir(path *QString) *QDir {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDirC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDir)
	return gothis
}

// /usr/include/qt/QtCore/qdir.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDir(const QString &, const QString &, QDir::SortFlags, QDir::Filters)
func NewQDir_1(path *QString, nameFilter *QString, sort int, filter int) *QDir {
	var convArg0 = path.GetCthis()
	var convArg1 = nameFilter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDirC2ERK7QStringS2_6QFlagsINS_8SortFlagEES3_INS_6FilterEE", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, sort, filter)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDir)
	return gothis
}

// /usr/include/qt/QtCore/qdir.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDir()
func DeleteQDir(this *QDir) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDirD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qdir.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDir &)
func (this *QDir) Swap(other *QDir) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &)
func (this *QDir) SetPath(path *QString) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir7setPathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:117
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path()
func (this *QDir) Path() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir4pathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString absolutePath()
func (this *QDir) AbsolutePath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir12absolutePathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QString canonicalPath()
func (this *QDir) CanonicalPath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir13canonicalPathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addResourceSearchPath(const QString &)
func (this *QDir) AddResourceSearchPath(path *QString) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir21addResourceSearchPathERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QDir_AddResourceSearchPath(path *QString) {
	var nilthis *QDir
	nilthis.AddResourceSearchPath(path)
}

// /usr/include/qt/QtCore/qdir.h:123
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setSearchPaths(const QString &, const QStringList &)
func (this *QDir) SetSearchPaths(prefix *QString, searchPaths *QStringList) {
	var convArg0 = prefix.GetCthis()
	var convArg1 = searchPaths.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14setSearchPathsERK7QStringRK11QStringList", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QDir_SetSearchPaths(prefix *QString, searchPaths *QStringList) {
	var nilthis *QDir
	nilthis.SetSearchPaths(prefix, searchPaths)
}

// /usr/include/qt/QtCore/qdir.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addSearchPath(const QString &, const QString &)
func (this *QDir) AddSearchPath(prefix *QString, path *QString) {
	var convArg0 = prefix.GetCthis()
	var convArg1 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir13addSearchPathERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QDir_AddSearchPath(prefix *QString, path *QString) {
	var nilthis *QDir
	nilthis.AddSearchPath(prefix, path)
}

// /usr/include/qt/QtCore/qdir.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QString dirName()
func (this *QDir) DirName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir7dirNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath(const QString &)
func (this *QDir) FilePath(fileName *QString) *QString /*123*/ {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir8filePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QString absoluteFilePath(const QString &)
func (this *QDir) AbsoluteFilePath(fileName *QString) *QString /*123*/ {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir16absoluteFilePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] QString relativeFilePath(const QString &)
func (this *QDir) RelativeFilePath(fileName *QString) *QString /*123*/ {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir16relativeFilePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:132
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString toNativeSeparators(const QString &)
func (this *QDir) ToNativeSeparators(pathName *QString) *QString /*123*/ {
	var convArg0 = pathName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir18toNativeSeparatorsERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}
func QDir_ToNativeSeparators(pathName *QString) *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.ToNativeSeparators(pathName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:133
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString fromNativeSeparators(const QString &)
func (this *QDir) FromNativeSeparators(pathName *QString) *QString /*123*/ {
	var convArg0 = pathName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir20fromNativeSeparatorsERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}
func QDir_FromNativeSeparators(pathName *QString) *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.FromNativeSeparators(pathName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cd(const QString &)
func (this *QDir) Cd(dirName *QString) bool {
	var convArg0 = dirName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir2cdERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:136
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cdUp()
func (this *QDir) CdUp() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4cdUpEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameFilters(const QStringList &)
func (this *QDir) SetNameFilters(nameFilters *QStringList) {
	var convArg0 = nameFilters.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14setNameFiltersERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:141
// index:0
// Public Visibility=Default Availability=Available
// [4] QDir::Filters filter()
func (this *QDir) Filter() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6filterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdir.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilter(QDir::Filters)
func (this *QDir) SetFilter(filter int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir9setFilterE6QFlagsINS_6FilterEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:143
// index:0
// Public Visibility=Default Availability=Available
// [4] QDir::SortFlags sorting()
func (this *QDir) Sorting() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir7sortingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdir.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSorting(QDir::SortFlags)
func (this *QDir) SetSorting(sort int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir10setSortingE6QFlagsINS_8SortFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sort)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:146
// index:0
// Public Visibility=Default Availability=Available
// [4] uint count()
func (this *QDir) Count() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdir.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty(QDir::Filters)
func (this *QDir) IsEmpty(filters int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir7isEmptyE6QFlagsINS_6FilterEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), filters)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:161
// index:0
// Public Visibility=Default Availability=Available
// [1] bool mkdir(const QString &)
func (this *QDir) Mkdir(dirName *QString) bool {
	var convArg0 = dirName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir5mkdirERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:162
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rmdir(const QString &)
func (this *QDir) Rmdir(dirName *QString) bool {
	var convArg0 = dirName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir5rmdirERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool mkpath(const QString &)
func (this *QDir) Mkpath(dirPath *QString) bool {
	var convArg0 = dirPath.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6mkpathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rmpath(const QString &)
func (this *QDir) Rmpath(dirPath *QString) bool {
	var convArg0 = dirPath.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6rmpathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool removeRecursively()
func (this *QDir) RemoveRecursively() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir17removeRecursivelyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:168
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadable()
func (this *QDir) IsReadable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir10isReadableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:169
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exists()
func (this *QDir) Exists() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6existsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:183
// index:1
// Public Visibility=Default Availability=Available
// [1] bool exists(const QString &)
func (this *QDir) Exists_1(name *QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6existsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRoot()
func (this *QDir) IsRoot() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir6isRootEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:172
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isRelativePath(const QString &)
func (this *QDir) IsRelativePath(path *QString) bool {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14isRelativePathERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
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
// Public static inline Visibility=Default Availability=Available
// [1] bool isAbsolutePath(const QString &)
func (this *QDir) IsAbsolutePath(path *QString) bool {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir14isAbsolutePathERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
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
// Public Visibility=Default Availability=Available
// [1] bool isRelative()
func (this *QDir) IsRelative() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir10isRelativeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:175
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAbsolute()
func (this *QDir) IsAbsolute() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir10isAbsoluteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:176
// index:0
// Public Visibility=Default Availability=Available
// [1] bool makeAbsolute()
func (this *QDir) MakeAbsolute() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir12makeAbsoluteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool remove(const QString &)
func (this *QDir) Remove(fileName *QString) bool {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir6removeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rename(const QString &, const QString &)
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
// Public static inline Visibility=Default Availability=Available
// [2] QChar listSeparator()
func (this *QDir) ListSeparator() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir13listSeparatorEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}
func QDir_ListSeparator() *QChar /*123*/ {
	var nilthis *QDir
	rv := nilthis.ListSeparator()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:196
// index:0
// Public static Visibility=Default Availability=Available
// [2] QChar separator()
func (this *QDir) Separator() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir9separatorEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}
func QDir_Separator() *QChar /*123*/ {
	var nilthis *QDir
	rv := nilthis.Separator()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:198
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool setCurrent(const QString &)
func (this *QDir) SetCurrent(path *QString) bool {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir10setCurrentERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
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
// Public static inline Visibility=Default Availability=Available
// [8] QDir current()
func (this *QDir) Current() *QDir /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir7currentEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}
func QDir_Current() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Current()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:200
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString currentPath()
func (this *QDir) CurrentPath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir11currentPathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}
func QDir_CurrentPath() *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.CurrentPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:202
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDir home()
func (this *QDir) Home() *QDir /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4homeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}
func QDir_Home() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Home()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:203
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString homePath()
func (this *QDir) HomePath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir8homePathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}
func QDir_HomePath() *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.HomePath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:204
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDir root()
func (this *QDir) Root() *QDir /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4rootEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}
func QDir_Root() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Root()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:205
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString rootPath()
func (this *QDir) RootPath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir8rootPathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}
func QDir_RootPath() *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.RootPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:206
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDir temp()
func (this *QDir) Temp() *QDir /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir4tempEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}
func QDir_Temp() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Temp()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:207
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString tempPath()
func (this *QDir) TempPath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir8tempPathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}
func QDir_TempPath() *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.TempPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool match(const QStringList &, const QString &)
func (this *QDir) Match(filters *QStringList, fileName *QString) bool {
	var convArg0 = filters.GetCthis()
	var convArg1 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir5matchERK11QStringListRK7QString", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
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
// Public static Visibility=Default Availability=Available
// [1] bool match(const QString &, const QString &)
func (this *QDir) Match_1(filter *QString, fileName *QString) bool {
	var convArg0 = filter.GetCthis()
	var convArg1 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir5matchERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
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
// Public static Visibility=Default Availability=Available
// [8] QString cleanPath(const QString &)
func (this *QDir) CleanPath(path *QString) *QString /*123*/ {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QDir9cleanPathERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}
func QDir_CleanPath(path *QString) *QString /*123*/ {
	var nilthis *QDir
	rv := nilthis.CleanPath(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh()
func (this *QDir) Refresh() {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QDir7refreshEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QDir__Filter = int

const QDir__Dirs QDir__Filter = 1
const QDir__Files QDir__Filter = 2
const QDir__Drives QDir__Filter = 4
const QDir__NoSymLinks QDir__Filter = 8
const QDir__AllEntries QDir__Filter = 7
const QDir__TypeMask QDir__Filter = 15
const QDir__Readable QDir__Filter = 16
const QDir__Writable QDir__Filter = 32
const QDir__Executable QDir__Filter = 64
const QDir__PermissionMask QDir__Filter = 112
const QDir__Modified QDir__Filter = 128
const QDir__Hidden QDir__Filter = 256
const QDir__System QDir__Filter = 512
const QDir__AccessMask QDir__Filter = 1008
const QDir__AllDirs QDir__Filter = 1024
const QDir__CaseSensitive QDir__Filter = 2048
const QDir__NoDot QDir__Filter = 8192
const QDir__NoDotDot QDir__Filter = 16384
const QDir__NoDotAndDotDot QDir__Filter = 24576
const QDir__NoFilter QDir__Filter = 4294967295

type QDir__SortFlag = int

const QDir__Name QDir__SortFlag = 0
const QDir__Time QDir__SortFlag = 1
const QDir__Size QDir__SortFlag = 2
const QDir__Unsorted QDir__SortFlag = 3
const QDir__SortByMask QDir__SortFlag = 3
const QDir__DirsFirst QDir__SortFlag = 4
const QDir__Reversed QDir__SortFlag = 8
const QDir__IgnoreCase QDir__SortFlag = 16
const QDir__DirsLast QDir__SortFlag = 32
const QDir__LocaleAware QDir__SortFlag = 64
const QDir__Type QDir__SortFlag = 128
const QDir__NoSort QDir__SortFlag = 4294967295

//  body block end
