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

// /usr/include/qt/QtCore/qfile.h:60
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFile) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:65
// index:0
// void QFile()
func NewQFile() *QFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(cthis)
	return gothis
}
func NewQFileFromPointer(cthis unsafe.Pointer) *QFile {
	return &QFile{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qfile.h:66
// index:1
// void QFile(const class QString &)
func NewQFile_1(name unsafe.Pointer) *QFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, name)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:68
// index:2
// void QFile(class QObject *)
func NewQFile_2(parent unsafe.Pointer) *QFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:69
// index:3
// void QFile(const class QString &, class QObject *)
func NewQFile_3(name unsafe.Pointer, parent unsafe.Pointer) *QFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, name, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:145
// index:4
// void QFile(class QFilePrivate &, class QObject *)
func NewQFile_4(dd unsafe.Pointer, parent unsafe.Pointer) *QFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileC2ER12QFilePrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:71
// index:0
// virtual
// void ~QFile()
func DeleteQFile(*QFile) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFileD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:73
// index:0
// virtual
// QString fileName()
func (this *QFile) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile8fileNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:74
// index:0
// void setFileName(const class QString &)
func (this *QFile) SetFileName(name unsafe.Pointer) {
	// 0: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile11setFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:88
// index:0
// static inline
// QByteArray encodeName(const class QString &)
func (this *QFile) EncodeName(fileName unsafe.Pointer) {
	// 0: (fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile10encodeNameERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_EncodeName(fileName unsafe.Pointer) {
	// 0: (fileName const QString &), (fileName)
	var nilthis *QFile
	nilthis.EncodeName(fileName)
}

// /usr/include/qt/QtCore/qfile.h:92
// index:0
// static inline
// QString decodeName(const class QByteArray &)
func (this *QFile) DecodeName(localFileName unsafe.Pointer) {
	// 0: (localFileName const QByteArray &), (localFileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile10decodeNameERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_DecodeName(localFileName unsafe.Pointer) {
	// 0: (localFileName const QByteArray &), (localFileName)
	var nilthis *QFile
	nilthis.DecodeName(localFileName)
}

// /usr/include/qt/QtCore/qfile.h:97
// index:1
// static inline
// QString decodeName(const char *)
func (this *QFile) DecodeName_1(localFileName unsafe.Pointer) {
	// 1: (localFileName const char *), (localFileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile10decodeNameEPKc", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_DecodeName_1(localFileName unsafe.Pointer) {
	// 1: (localFileName const char *), (localFileName)
	var nilthis *QFile
	nilthis.DecodeName_1(localFileName)
}

// /usr/include/qt/QtCore/qfile.h:107
// index:0
// bool exists()
func (this *QFile) Exists() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile6existsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:108
// index:1
// static
// bool exists(const class QString &)
func (this *QFile) Exists_1(fileName unsafe.Pointer) {
	// 1: (fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6existsERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_Exists_1(fileName unsafe.Pointer) {
	// 1: (fileName const QString &), (fileName)
	var nilthis *QFile
	nilthis.Exists_1(fileName)
}

// /usr/include/qt/QtCore/qfile.h:110
// index:0
// QString readLink()
func (this *QFile) ReadLink() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile8readLinkEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:111
// index:1
// static
// QString readLink(const class QString &)
func (this *QFile) ReadLink_1(fileName unsafe.Pointer) {
	// 1: (fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile8readLinkERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_ReadLink_1(fileName unsafe.Pointer) {
	// 1: (fileName const QString &), (fileName)
	var nilthis *QFile
	nilthis.ReadLink_1(fileName)
}

// /usr/include/qt/QtCore/qfile.h:112
// index:0
// inline
// QString symLinkTarget()
func (this *QFile) SymLinkTarget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile13symLinkTargetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:113
// index:1
// static inline
// QString symLinkTarget(const class QString &)
func (this *QFile) SymLinkTarget_1(fileName unsafe.Pointer) {
	// 1: (fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile13symLinkTargetERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_SymLinkTarget_1(fileName unsafe.Pointer) {
	// 1: (fileName const QString &), (fileName)
	var nilthis *QFile
	nilthis.SymLinkTarget_1(fileName)
}

// /usr/include/qt/QtCore/qfile.h:115
// index:0
// bool remove()
func (this *QFile) Remove() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6removeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:116
// index:1
// static
// bool remove(const class QString &)
func (this *QFile) Remove_1(fileName unsafe.Pointer) {
	// 1: (fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6removeERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_Remove_1(fileName unsafe.Pointer) {
	// 1: (fileName const QString &), (fileName)
	var nilthis *QFile
	nilthis.Remove_1(fileName)
}

// /usr/include/qt/QtCore/qfile.h:118
// index:0
// bool rename(const class QString &)
func (this *QFile) Rename(newName unsafe.Pointer) {
	// 0: (, newName const QString &), (newName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6renameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), newName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:119
// index:1
// static
// bool rename(const class QString &, const class QString &)
func (this *QFile) Rename_1(oldName unsafe.Pointer, newName unsafe.Pointer) {
	// 1: (oldName const QString &, newName const QString &), (oldName, newName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6renameERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_Rename_1(oldName unsafe.Pointer, newName unsafe.Pointer) {
	// 1: (oldName const QString &, newName const QString &), (oldName, newName)
	var nilthis *QFile
	nilthis.Rename_1(oldName, newName)
}

// /usr/include/qt/QtCore/qfile.h:121
// index:0
// bool link(const class QString &)
func (this *QFile) Link(newName unsafe.Pointer) {
	// 0: (, newName const QString &), (newName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4linkERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), newName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:122
// index:1
// static
// bool link(const class QString &, const class QString &)
func (this *QFile) Link_1(oldname unsafe.Pointer, newName unsafe.Pointer) {
	// 1: (oldname const QString &, newName const QString &), (oldname, newName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4linkERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_Link_1(oldname unsafe.Pointer, newName unsafe.Pointer) {
	// 1: (oldname const QString &, newName const QString &), (oldname, newName)
	var nilthis *QFile
	nilthis.Link_1(oldname, newName)
}

// /usr/include/qt/QtCore/qfile.h:124
// index:0
// bool copy(const class QString &)
func (this *QFile) Copy(newName unsafe.Pointer) {
	// 0: (, newName const QString &), (newName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4copyERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), newName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:125
// index:1
// static
// bool copy(const class QString &, const class QString &)
func (this *QFile) Copy_1(fileName unsafe.Pointer, newName unsafe.Pointer) {
	// 1: (fileName const QString &, newName const QString &), (fileName, newName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4copyERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_Copy_1(fileName unsafe.Pointer, newName unsafe.Pointer) {
	// 1: (fileName const QString &, newName const QString &), (fileName, newName)
	var nilthis *QFile
	nilthis.Copy_1(fileName, newName)
}

// /usr/include/qt/QtCore/qfile.h:127
// index:0
// virtual
// bool open(QIODevice::OpenMode)
func (this *QFile) Open(flags int) {
	// 0: (, QFlags<QIODevice::OpenModeFlag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4openE6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:128
// index:1
// bool open(FILE *, QIODevice::OpenMode, QFileDevice::FileHandleFlags)
func (this *QFile) Open_1(f unsafe.Pointer, ioFlags int, handleFlags int) {
	// 1: (, f FILE *, QFlags<QIODevice::OpenModeFlag> ioFlags, QFlags<QFileDevice::FileHandleFlag> handleFlags), (f, ioFlags, handleFlags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4openEP8_IO_FILE6QFlagsIN9QIODevice12OpenModeFlagEES2_IN11QFileDevice14FileHandleFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), f, ioFlags, handleFlags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:129
// index:2
// bool open(int, QIODevice::OpenMode, QFileDevice::FileHandleFlags)
func (this *QFile) Open_2(fd int, ioFlags int, handleFlags int) {
	// 2: (, fd int, QFlags<QIODevice::OpenModeFlag> ioFlags, QFlags<QFileDevice::FileHandleFlag> handleFlags), (&fd, ioFlags, handleFlags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile4openEi6QFlagsIN9QIODevice12OpenModeFlagEES0_IN11QFileDevice14FileHandleFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &fd, ioFlags, handleFlags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:131
// index:0
// virtual
// qint64 size()
func (this *QFile) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:133
// index:0
// virtual
// bool resize(qint64)
func (this *QFile) Resize(sz int64) {
	// 0: (, sz qint64), (&sz)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6resizeEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sz)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:134
// index:1
// static
// bool resize(const class QString &, qint64)
func (this *QFile) Resize_1(filename unsafe.Pointer, sz int64) {
	// 1: (filename const QString &, sz qint64), (filename, sz)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile6resizeERK7QStringx", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_Resize_1(filename unsafe.Pointer, sz int64) {
	// 1: (filename const QString &, sz qint64), (filename, sz)
	var nilthis *QFile
	nilthis.Resize_1(filename, sz)
}

// /usr/include/qt/QtCore/qfile.h:136
// index:0
// virtual
// QFileDevice::Permissions permissions()
func (this *QFile) Permissions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFile11permissionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:137
// index:1
// static
// QFileDevice::Permissions permissions(const class QString &)
func (this *QFile) Permissions_1(filename unsafe.Pointer) {
	// 1: (filename const QString &), (filename)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile11permissionsERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_Permissions_1(filename unsafe.Pointer) {
	// 1: (filename const QString &), (filename)
	var nilthis *QFile
	nilthis.Permissions_1(filename)
}

// /usr/include/qt/QtCore/qfile.h:138
// index:0
// virtual
// bool setPermissions(QFileDevice::Permissions)
func (this *QFile) SetPermissions(permissionSpec int) {
	// 0: (, QFlags<QFileDevice::Permission> permissionSpec), (permissionSpec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile14setPermissionsE6QFlagsIN11QFileDevice10PermissionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), permissionSpec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:139
// index:1
// static
// bool setPermissions(const class QString &, QFileDevice::Permissions)
func (this *QFile) SetPermissions_1(filename unsafe.Pointer, permissionSpec int) {
	// 1: (filename const QString &, QFlags<QFileDevice::Permission> permissionSpec), (filename, permissionSpec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFile14setPermissionsERK7QString6QFlagsIN11QFileDevice10PermissionEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFile_SetPermissions_1(filename unsafe.Pointer, permissionSpec int) {
	// 1: (filename const QString &, QFlags<QFileDevice::Permission> permissionSpec), (filename, permissionSpec)
	var nilthis *QFile
	nilthis.SetPermissions_1(filename, permissionSpec)
}

//  body block end
