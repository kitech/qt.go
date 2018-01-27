package qtcore

// /usr/include/qt/QtCore/qfile.h
// #include <qfile.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QFile struct {
	*qtrt.CObject
}

func (this *QFile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFile) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQFileFromPointer(cthis unsafe.Pointer) *QFile {
	return &QFile{&qtrt.CObject{cthis}}
}
func (*QFile) NewFromPointer(cthis unsafe.Pointer) *QFile {
	return NewQFileFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfile.h:60
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFile) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qfile.h:65
// index:0
// Public
// void QFile()
func NewQFile() *QFile {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:66
// index:1
// Public
// void QFile(const QString &)
func NewQFile_1(name *QString) *QFile {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:68
// index:2
// Public
// void QFile(QObject *)
func NewQFile_2(parent *QObject /*777 QObject **/) *QFile {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:69
// index:3
// Public
// void QFile(const QString &, QObject *)
func NewQFile_3(name *QString, parent *QObject /*777 QObject **/) *QFile {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = name.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:71
// index:0
// Public virtual
// void ~QFile()
func DeleteQFile(*QFile) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:73
// index:0
// Public virtual
// QString fileName()
func (this *QFile) FileName() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile8fileNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qfile.h:74
// index:0
// Public
// void setFileName(const QString &)
func (this *QFile) SetFileName(name *QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:88
// index:0
// Public static inline
// QByteArray encodeName(const QString &)
func (this *QFile) EncodeName(fileName *QString) *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile10encodeNameERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QFile_EncodeName(fileName *QString) *QByteArray /*123*/ {
	var nilthis *QFile
	rv := nilthis.EncodeName(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:92
// index:0
// Public static inline
// QString decodeName(const QByteArray &)
func (this *QFile) DecodeName(localFileName *QByteArray) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile10decodeNameERK10QByteArray", ffiqt.FFI_TYPE_POINTER, localFileName)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QFile_DecodeName(localFileName *QByteArray) *QString /*123*/ {
	var nilthis *QFile
	rv := nilthis.DecodeName(localFileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:97
// index:1
// Public static inline
// QString decodeName(const char *)
func (this *QFile) DecodeName_1(localFileName string) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile10decodeNameEPKc", ffiqt.FFI_TYPE_POINTER, localFileName)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QFile_DecodeName_1(localFileName string) *QString /*123*/ {
	var nilthis *QFile
	rv := nilthis.DecodeName_1(localFileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:107
// index:0
// Public
// bool exists()
func (this *QFile) Exists() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile6existsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:108
// index:1
// Public static
// bool exists(const QString &)
func (this *QFile) Exists_1(fileName *QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6existsERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QFile_Exists_1(fileName *QString) bool {
	var nilthis *QFile
	rv := nilthis.Exists_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:110
// index:0
// Public
// QString readLink()
func (this *QFile) ReadLink() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile8readLinkEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qfile.h:111
// index:1
// Public static
// QString readLink(const QString &)
func (this *QFile) ReadLink_1(fileName *QString) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile8readLinkERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QFile_ReadLink_1(fileName *QString) *QString /*123*/ {
	var nilthis *QFile
	rv := nilthis.ReadLink_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:112
// index:0
// Public inline
// QString symLinkTarget()
func (this *QFile) SymLinkTarget() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile13symLinkTargetEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qfile.h:113
// index:1
// Public static inline
// QString symLinkTarget(const QString &)
func (this *QFile) SymLinkTarget_1(fileName *QString) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile13symLinkTargetERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QFile_SymLinkTarget_1(fileName *QString) *QString /*123*/ {
	var nilthis *QFile
	rv := nilthis.SymLinkTarget_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:115
// index:0
// Public
// bool remove()
func (this *QFile) Remove() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6removeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:116
// index:1
// Public static
// bool remove(const QString &)
func (this *QFile) Remove_1(fileName *QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6removeERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QFile_Remove_1(fileName *QString) bool {
	var nilthis *QFile
	rv := nilthis.Remove_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:118
// index:0
// Public
// bool rename(const QString &)
func (this *QFile) Rename(newName *QString) bool {
	var convArg0 = newName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6renameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:119
// index:1
// Public static
// bool rename(const QString &, const QString &)
func (this *QFile) Rename_1(oldName *QString, newName *QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6renameERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, oldName, newName)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QFile_Rename_1(oldName *QString, newName *QString) bool {
	var nilthis *QFile
	rv := nilthis.Rename_1(oldName, newName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:121
// index:0
// Public
// bool link(const QString &)
func (this *QFile) Link(newName *QString) bool {
	var convArg0 = newName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4linkERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:122
// index:1
// Public static
// bool link(const QString &, const QString &)
func (this *QFile) Link_1(oldname *QString, newName *QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4linkERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, oldname, newName)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QFile_Link_1(oldname *QString, newName *QString) bool {
	var nilthis *QFile
	rv := nilthis.Link_1(oldname, newName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:124
// index:0
// Public
// bool copy(const QString &)
func (this *QFile) Copy(newName *QString) bool {
	var convArg0 = newName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4copyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:125
// index:1
// Public static
// bool copy(const QString &, const QString &)
func (this *QFile) Copy_1(fileName *QString, newName *QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4copyERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, fileName, newName)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QFile_Copy_1(fileName *QString, newName *QString) bool {
	var nilthis *QFile
	rv := nilthis.Copy_1(fileName, newName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:131
// index:0
// Public virtual
// qint64 size()
func (this *QFile) Size() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qfile.h:133
// index:0
// Public virtual
// bool resize(qint64)
func (this *QFile) Resize(sz int64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6resizeEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sz)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:134
// index:1
// Public static
// bool resize(const QString &, qint64)
func (this *QFile) Resize_1(filename *QString, sz int64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6resizeERK7QStringx", ffiqt.FFI_TYPE_POINTER, filename, sz)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QFile_Resize_1(filename *QString, sz int64) bool {
	var nilthis *QFile
	rv := nilthis.Resize_1(filename, sz)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:136
// index:0
// Public virtual
// QFileDevice::Permissions permissions()
func (this *QFile) Permissions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile11permissionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qfile.h:137
// index:1
// Public static
// QFileDevice::Permissions permissions(const QString &)
func (this *QFile) Permissions_1(filename *QString) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile11permissionsERK7QString", ffiqt.FFI_TYPE_POINTER, filename)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QFile_Permissions_1(filename *QString) int {
	var nilthis *QFile
	rv := nilthis.Permissions_1(filename)
	return rv
}

//  body block end
