package qtcore

// /usr/include/qt/QtCore/qdir.h
// #include <qdir.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QDir struct {
	*qtrt.CObject
}
type QDir_ITF interface {
	QDir_PTR() *QDir
}

func (ptr *QDir) QDir_PTR() *QDir { return ptr }

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
func NewQDir(path string) *QDir {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDirC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDir)
	return gothis
}

// /usr/include/qt/QtCore/qdir.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDir(const QString &, const QString &, QDir::SortFlags, QDir::Filters)
func NewQDir_1(path string, nameFilter string, sort int, filter int) *QDir {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(nameFilter)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDirC2ERK7QStringS2_6QFlagsINS_8SortFlagEES3_INS_6FilterEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, sort, filter)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDir)
	return gothis
}

// /usr/include/qt/QtCore/qdir.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDir()
func DeleteQDir(this *QDir) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDirD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qdir.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDir &)
func (this *QDir) Swap(other QDir_ITF) {
	var convArg0 = other.QDir_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &)
func (this *QDir) SetPath(path string) {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir7setPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:117
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path()
func (this *QDir) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString absolutePath()
func (this *QDir) AbsolutePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir12absolutePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QString canonicalPath()
func (this *QDir) CanonicalPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir13canonicalPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addResourceSearchPath(const QString &)
func (this *QDir) AddResourceSearchPath(path string) {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir21addResourceSearchPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QDir_AddResourceSearchPath(path string) {
	var nilthis *QDir
	nilthis.AddResourceSearchPath(path)
}

// /usr/include/qt/QtCore/qdir.h:123
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setSearchPaths(const QString &, const QStringList &)
func (this *QDir) SetSearchPaths(prefix string, searchPaths QStringList_ITF) {
	var tmpArg0 = NewQString_5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = searchPaths.QStringList_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir14setSearchPathsERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QDir_SetSearchPaths(prefix string, searchPaths QStringList_ITF) {
	var nilthis *QDir
	nilthis.SetSearchPaths(prefix, searchPaths)
}

// /usr/include/qt/QtCore/qdir.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addSearchPath(const QString &, const QString &)
func (this *QDir) AddSearchPath(prefix string, path string) {
	var tmpArg0 = NewQString_5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(path)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir13addSearchPathERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QDir_AddSearchPath(prefix string, path string) {
	var nilthis *QDir
	nilthis.AddSearchPath(prefix, path)
}

// /usr/include/qt/QtCore/qdir.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList searchPaths(const QString &)
func (this *QDir) SearchPaths(prefix string) *QStringList /*123*/ {
	var tmpArg0 = NewQString_5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir11searchPathsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}
func QDir_SearchPaths(prefix string) *QStringList /*123*/ {
	var nilthis *QDir
	rv := nilthis.SearchPaths(prefix)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QString dirName()
func (this *QDir) DirName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir7dirNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath(const QString &)
func (this *QDir) FilePath(fileName string) string {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir8filePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QString absoluteFilePath(const QString &)
func (this *QDir) AbsoluteFilePath(fileName string) string {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir16absoluteFilePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] QString relativeFilePath(const QString &)
func (this *QDir) RelativeFilePath(fileName string) string {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir16relativeFilePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:132
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString toNativeSeparators(const QString &)
func (this *QDir) ToNativeSeparators(pathName string) string {
	var tmpArg0 = NewQString_5(pathName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir18toNativeSeparatorsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_ToNativeSeparators(pathName string) string {
	var nilthis *QDir
	rv := nilthis.ToNativeSeparators(pathName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:133
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString fromNativeSeparators(const QString &)
func (this *QDir) FromNativeSeparators(pathName string) string {
	var tmpArg0 = NewQString_5(pathName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir20fromNativeSeparatorsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_FromNativeSeparators(pathName string) string {
	var nilthis *QDir
	rv := nilthis.FromNativeSeparators(pathName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cd(const QString &)
func (this *QDir) Cd(dirName string) bool {
	var tmpArg0 = NewQString_5(dirName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir2cdERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:136
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cdUp()
func (this *QDir) CdUp() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir4cdUpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList nameFilters()
func (this *QDir) NameFilters() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir11nameFiltersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameFilters(const QStringList &)
func (this *QDir) SetNameFilters(nameFilters QStringList_ITF) {
	var convArg0 = nameFilters.QStringList_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir14setNameFiltersERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:141
// index:0
// Public Visibility=Default Availability=Available
// [4] QDir::Filters filter()
func (this *QDir) Filter() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6filterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdir.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilter(QDir::Filters)
func (this *QDir) SetFilter(filter int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir9setFilterE6QFlagsINS_6FilterEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:143
// index:0
// Public Visibility=Default Availability=Available
// [4] QDir::SortFlags sorting()
func (this *QDir) Sorting() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir7sortingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdir.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSorting(QDir::SortFlags)
func (this *QDir) SetSorting(sort int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir10setSortingE6QFlagsINS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sort)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:146
// index:0
// Public Visibility=Default Availability=Available
// [4] uint count()
func (this *QDir) Count() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdir.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty(QDir::Filters)
func (this *QDir) IsEmpty(filters int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir7isEmptyE6QFlagsINS_6FilterEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:151
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList nameFiltersFromString(const QString &)
func (this *QDir) NameFiltersFromString(nameFilter string) *QStringList /*123*/ {
	var tmpArg0 = NewQString_5(nameFilter)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir21nameFiltersFromStringERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}
func QDir_NameFiltersFromString(nameFilter string) *QStringList /*123*/ {
	var nilthis *QDir
	rv := nilthis.NameFiltersFromString(nameFilter)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:153
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList entryList(QDir::Filters, QDir::SortFlags)
func (this *QDir) EntryList(filters int, sort int) *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir9entryListE6QFlagsINS_6FilterEES0_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:154
// index:1
// Public Visibility=Default Availability=Available
// [8] QStringList entryList(const QStringList &, QDir::Filters, QDir::SortFlags)
func (this *QDir) EntryList_1(nameFilters QStringList_ITF, filters int, sort int) *QStringList /*123*/ {
	var convArg0 = nameFilters.QStringList_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir9entryListERK11QStringList6QFlagsINS_6FilterEES3_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] QFileInfoList entryInfoList(QDir::Filters, QDir::SortFlags)
func (this *QDir) EntryInfoList(filters int, sort int) *QFileInfoList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir13entryInfoListE6QFlagsINS_6FilterEES0_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:158
// index:1
// Public Visibility=Default Availability=Available
// [-2] QFileInfoList entryInfoList(const QStringList &, QDir::Filters, QDir::SortFlags)
func (this *QDir) EntryInfoList_1(nameFilters QStringList_ITF, filters int, sort int) *QFileInfoList /*667*/ {
	var convArg0 = nameFilters.QStringList_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir13entryInfoListERK11QStringList6QFlagsINS_6FilterEES3_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:161
// index:0
// Public Visibility=Default Availability=Available
// [1] bool mkdir(const QString &)
func (this *QDir) Mkdir(dirName string) bool {
	var tmpArg0 = NewQString_5(dirName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir5mkdirERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:162
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rmdir(const QString &)
func (this *QDir) Rmdir(dirName string) bool {
	var tmpArg0 = NewQString_5(dirName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir5rmdirERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool mkpath(const QString &)
func (this *QDir) Mkpath(dirPath string) bool {
	var tmpArg0 = NewQString_5(dirPath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6mkpathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rmpath(const QString &)
func (this *QDir) Rmpath(dirPath string) bool {
	var tmpArg0 = NewQString_5(dirPath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6rmpathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool removeRecursively()
func (this *QDir) RemoveRecursively() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir17removeRecursivelyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:168
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadable()
func (this *QDir) IsReadable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir10isReadableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:169
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exists()
func (this *QDir) Exists() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6existsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:183
// index:1
// Public Visibility=Default Availability=Available
// [1] bool exists(const QString &)
func (this *QDir) Exists_1(name string) bool {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6existsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRoot()
func (this *QDir) IsRoot() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6isRootEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:172
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isRelativePath(const QString &)
func (this *QDir) IsRelativePath(path string) bool {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir14isRelativePathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_IsRelativePath(path string) bool {
	var nilthis *QDir
	rv := nilthis.IsRelativePath(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:173
// index:0
// Public static inline Visibility=Default Availability=Available
// [1] bool isAbsolutePath(const QString &)
func (this *QDir) IsAbsolutePath(path string) bool {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir14isAbsolutePathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_IsAbsolutePath(path string) bool {
	var nilthis *QDir
	rv := nilthis.IsAbsolutePath(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:174
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRelative()
func (this *QDir) IsRelative() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir10isRelativeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:175
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAbsolute()
func (this *QDir) IsAbsolute() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir10isAbsoluteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:176
// index:0
// Public Visibility=Default Availability=Available
// [1] bool makeAbsolute()
func (this *QDir) MakeAbsolute() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir12makeAbsoluteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool remove(const QString &)
func (this *QDir) Remove(fileName string) bool {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir6removeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rename(const QString &, const QString &)
func (this *QDir) Rename(oldName string, newName string) bool {
	var tmpArg0 = NewQString_5(oldName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(newName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir6renameERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:185
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QFileInfoList drives()
func (this *QDir) Drives() *QFileInfoList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir6drivesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}
func QDir_Drives() *QFileInfoList /*667*/ {
	var nilthis *QDir
	rv := nilthis.Drives()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:187
// index:0
// Public static inline Visibility=Default Availability=Available
// [2] QChar listSeparator()
func (this *QDir) ListSeparator() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir13listSeparatorEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir9separatorEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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
func (this *QDir) SetCurrent(path string) bool {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir10setCurrentERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_SetCurrent(path string) bool {
	var nilthis *QDir
	rv := nilthis.SetCurrent(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:199
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDir current()
func (this *QDir) Current() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir7currentEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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
func (this *QDir) CurrentPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir11currentPathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_CurrentPath() string {
	var nilthis *QDir
	rv := nilthis.CurrentPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:202
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDir home()
func (this *QDir) Home() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir4homeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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
func (this *QDir) HomePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir8homePathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_HomePath() string {
	var nilthis *QDir
	rv := nilthis.HomePath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:204
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDir root()
func (this *QDir) Root() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir4rootEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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
func (this *QDir) RootPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir8rootPathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_RootPath() string {
	var nilthis *QDir
	rv := nilthis.RootPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:206
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDir temp()
func (this *QDir) Temp() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir4tempEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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
func (this *QDir) TempPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir8tempPathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_TempPath() string {
	var nilthis *QDir
	rv := nilthis.TempPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool match(const QStringList &, const QString &)
func (this *QDir) Match(filters QStringList_ITF, fileName string) bool {
	var convArg0 = filters.QStringList_PTR().GetCthis()
	var tmpArg1 = NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir5matchERK11QStringListRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_Match(filters QStringList_ITF, fileName string) bool {
	var nilthis *QDir
	rv := nilthis.Match(filters, fileName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:211
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool match(const QString &, const QString &)
func (this *QDir) Match_1(filter string, fileName string) bool {
	var tmpArg0 = NewQString_5(filter)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir5matchERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_Match_1(filter string, fileName string) bool {
	var nilthis *QDir
	rv := nilthis.Match_1(filter, fileName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:214
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString cleanPath(const QString &)
func (this *QDir) CleanPath(path string) string {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir9cleanPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_CleanPath(path string) string {
	var nilthis *QDir
	rv := nilthis.CleanPath(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh()
func (this *QDir) Refresh() {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir7refreshEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
const QDir__NoFilter QDir__Filter = -1

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
const QDir__NoSort QDir__SortFlag = -1

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
