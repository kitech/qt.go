//  header block begin
// /usr/include/qt/QtCore/qfile.h
// #include <qfile.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QFile struct {
	*qtrt.CObject
}

func (this *QFile) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQFileFromPointer(cthis unsafe.Pointer) *QFile {
	return &QFile{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qfile.h:60
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFile) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
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
// void QFile(const class QString &)
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
// void QFile(class QObject *)
func NewQFile_2(parent unsafe.Pointer) *QFile {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:69
// index:3
// Public
// void QFile(const class QString &, class QObject *)
func NewQFile_3(name *QString, parent unsafe.Pointer) *QFile {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
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
func (this *QFile) FileName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:74
// index:0
// Public
// void setFileName(const class QString &)
func (this *QFile) SetFileName(name *QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:88
// index:0
// Public static inline
// QByteArray encodeName(const class QString &)
func (this *QFile) EncodeName(fileName *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile10encodeNameERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_EncodeName(fileName *QString) {
	var nilthis *QFile
	nilthis.EncodeName(fileName)
}

// /usr/include/qt/QtCore/qfile.h:92
// index:0
// Public static inline
// QString decodeName(const class QByteArray &)
func (this *QFile) DecodeName(localFileName *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile10decodeNameERK10QByteArray", ffiqt.FFI_TYPE_POINTER, localFileName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_DecodeName(localFileName *QByteArray) {
	var nilthis *QFile
	nilthis.DecodeName(localFileName)
}

// /usr/include/qt/QtCore/qfile.h:97
// index:1
// Public static inline
// QString decodeName(const char *)
func (this *QFile) DecodeName_1(localFileName string) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile10decodeNameEPKc", ffiqt.FFI_TYPE_POINTER, localFileName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_DecodeName_1(localFileName string) {
	var nilthis *QFile
	nilthis.DecodeName_1(localFileName)
}

// /usr/include/qt/QtCore/qfile.h:107
// index:0
// Public
// bool exists()
func (this *QFile) Exists() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile6existsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:108
// index:1
// Public static
// bool exists(const class QString &)
func (this *QFile) Exists_1(fileName *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6existsERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_Exists_1(fileName *QString) {
	var nilthis *QFile
	nilthis.Exists_1(fileName)
}

// /usr/include/qt/QtCore/qfile.h:110
// index:0
// Public
// QString readLink()
func (this *QFile) ReadLink() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile8readLinkEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:111
// index:1
// Public static
// QString readLink(const class QString &)
func (this *QFile) ReadLink_1(fileName *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile8readLinkERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_ReadLink_1(fileName *QString) {
	var nilthis *QFile
	nilthis.ReadLink_1(fileName)
}

// /usr/include/qt/QtCore/qfile.h:112
// index:0
// Public inline
// QString symLinkTarget()
func (this *QFile) SymLinkTarget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile13symLinkTargetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:113
// index:1
// Public static inline
// QString symLinkTarget(const class QString &)
func (this *QFile) SymLinkTarget_1(fileName *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile13symLinkTargetERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_SymLinkTarget_1(fileName *QString) {
	var nilthis *QFile
	nilthis.SymLinkTarget_1(fileName)
}

// /usr/include/qt/QtCore/qfile.h:115
// index:0
// Public
// bool remove()
func (this *QFile) Remove() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6removeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:116
// index:1
// Public static
// bool remove(const class QString &)
func (this *QFile) Remove_1(fileName *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6removeERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_Remove_1(fileName *QString) {
	var nilthis *QFile
	nilthis.Remove_1(fileName)
}

// /usr/include/qt/QtCore/qfile.h:118
// index:0
// Public
// bool rename(const class QString &)
func (this *QFile) Rename(newName *QString) interface{} {
	var convArg0 = newName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6renameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:119
// index:1
// Public static
// bool rename(const class QString &, const class QString &)
func (this *QFile) Rename_1(oldName *QString, newName *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6renameERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, oldName, newName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_Rename_1(oldName *QString, newName *QString) {
	var nilthis *QFile
	nilthis.Rename_1(oldName, newName)
}

// /usr/include/qt/QtCore/qfile.h:121
// index:0
// Public
// bool link(const class QString &)
func (this *QFile) Link(newName *QString) interface{} {
	var convArg0 = newName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4linkERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:122
// index:1
// Public static
// bool link(const class QString &, const class QString &)
func (this *QFile) Link_1(oldname *QString, newName *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4linkERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, oldname, newName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_Link_1(oldname *QString, newName *QString) {
	var nilthis *QFile
	nilthis.Link_1(oldname, newName)
}

// /usr/include/qt/QtCore/qfile.h:124
// index:0
// Public
// bool copy(const class QString &)
func (this *QFile) Copy(newName *QString) interface{} {
	var convArg0 = newName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4copyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:125
// index:1
// Public static
// bool copy(const class QString &, const class QString &)
func (this *QFile) Copy_1(fileName *QString, newName *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4copyERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, fileName, newName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_Copy_1(fileName *QString, newName *QString) {
	var nilthis *QFile
	nilthis.Copy_1(fileName, newName)
}

// /usr/include/qt/QtCore/qfile.h:131
// index:0
// Public virtual
// qint64 size()
func (this *QFile) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:133
// index:0
// Public virtual
// bool resize(qint64)
func (this *QFile) Resize(sz int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6resizeEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &sz)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:134
// index:1
// Public static
// bool resize(const class QString &, qint64)
func (this *QFile) Resize_1(filename *QString, sz int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6resizeERK7QStringx", ffiqt.FFI_TYPE_POINTER, filename, sz)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_Resize_1(filename *QString, sz int64) {
	var nilthis *QFile
	nilthis.Resize_1(filename, sz)
}

// /usr/include/qt/QtCore/qfile.h:136
// index:0
// Public virtual
// QFileDevice::Permissions permissions()
func (this *QFile) Permissions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile11permissionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:137
// index:1
// Public static
// QFileDevice::Permissions permissions(const class QString &)
func (this *QFile) Permissions_1(filename *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile11permissionsERK7QString", ffiqt.FFI_TYPE_POINTER, filename)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFile_Permissions_1(filename *QString) {
	var nilthis *QFile
	nilthis.Permissions_1(filename)
}

//  body block end
