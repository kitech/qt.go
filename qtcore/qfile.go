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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFileFromPointer(cthis unsafe.Pointer) *QFile {
	return &QFile{&qtrt.CObject{cthis}}
}
func (*QFile) NewFromPointer(cthis unsafe.Pointer) *QFile {
	return NewQFileFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfile.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QFile) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qfile.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFile()
func NewQFile() *QFile {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFileC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFile(const QString &)
func NewQFile_1(name string) *QFile {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFileC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QFile(QObject *)
func NewQFile_2(parent *QObject /*777 QObject **/) *QFile {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFileC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:69
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QFile(const QString &, QObject *)
func NewQFile_3(name string, parent *QObject /*777 QObject **/) *QFile {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFileC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFile()
func DeleteQFile(this *QFile) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFileD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfile.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QFile) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfile.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QFile) SetFileName(name string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:88
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QByteArray encodeName(const QString &)
func (this *QFile) EncodeName(fileName string) *QByteArray /*123*/ {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile10encodeNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QFile_EncodeName(fileName string) *QByteArray /*123*/ {
	var nilthis *QFile
	rv := nilthis.EncodeName(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:92
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QString decodeName(const QByteArray &)
func (this *QFile) DecodeName(localFileName *QByteArray) string {
	var convArg0 = localFileName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile10decodeNameERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QFile_DecodeName(localFileName *QByteArray) string {
	var nilthis *QFile
	rv := nilthis.DecodeName(localFileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:97
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QString decodeName(const char *)
func (this *QFile) DecodeName_1(localFileName string) string {
	var convArg0 = qtrt.CString(localFileName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile10decodeNameEPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QFile_DecodeName_1(localFileName string) string {
	var nilthis *QFile
	rv := nilthis.DecodeName_1(localFileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exists()
func (this *QFile) Exists() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile6existsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:108
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool exists(const QString &)
func (this *QFile) Exists_1(fileName string) bool {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6existsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Exists_1(fileName string) bool {
	var nilthis *QFile
	rv := nilthis.Exists_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readLink()
func (this *QFile) ReadLink() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile8readLinkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfile.h:111
// index:1
// Public static Visibility=Default Availability=Available
// [8] QString readLink(const QString &)
func (this *QFile) ReadLink_1(fileName string) string {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile8readLinkERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QFile_ReadLink_1(fileName string) string {
	var nilthis *QFile
	rv := nilthis.ReadLink_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString symLinkTarget()
func (this *QFile) SymLinkTarget() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile13symLinkTargetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfile.h:113
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QString symLinkTarget(const QString &)
func (this *QFile) SymLinkTarget_1(fileName string) string {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile13symLinkTargetERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QFile_SymLinkTarget_1(fileName string) string {
	var nilthis *QFile
	rv := nilthis.SymLinkTarget_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool remove()
func (this *QFile) Remove() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6removeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:116
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool remove(const QString &)
func (this *QFile) Remove_1(fileName string) bool {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6removeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Remove_1(fileName string) bool {
	var nilthis *QFile
	rv := nilthis.Remove_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rename(const QString &)
func (this *QFile) Rename(newName string) bool {
	var tmpArg0 = NewQString_5(newName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6renameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:119
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool rename(const QString &, const QString &)
func (this *QFile) Rename_1(oldName string, newName string) bool {
	var tmpArg0 = NewQString_5(oldName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(newName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6renameERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Rename_1(oldName string, newName string) bool {
	var nilthis *QFile
	rv := nilthis.Rename_1(oldName, newName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool link(const QString &)
func (this *QFile) Link(newName string) bool {
	var tmpArg0 = NewQString_5(newName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4linkERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:122
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool link(const QString &, const QString &)
func (this *QFile) Link_1(oldname string, newName string) bool {
	var tmpArg0 = NewQString_5(oldname)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(newName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4linkERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Link_1(oldname string, newName string) bool {
	var nilthis *QFile
	rv := nilthis.Link_1(oldname, newName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:124
// index:0
// Public Visibility=Default Availability=Available
// [1] bool copy(const QString &)
func (this *QFile) Copy(newName string) bool {
	var tmpArg0 = NewQString_5(newName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4copyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:125
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool copy(const QString &, const QString &)
func (this *QFile) Copy_1(fileName string, newName string) bool {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(newName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4copyERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Copy_1(fileName string, newName string) bool {
	var nilthis *QFile
	rv := nilthis.Copy_1(fileName, newName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)
func (this *QFile) Open(flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:129
// index:1
// Public Visibility=Default Availability=Available
// [1] bool open(int, QIODevice::OpenMode, QFileDevice::FileHandleFlags)
func (this *QFile) Open_1(fd int, ioFlags int, handleFlags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4openEi6QFlagsIN9QIODevice12OpenModeFlagEES0_IN11QFileDevice14FileHandleFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fd, ioFlags, handleFlags)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:131
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 size()
func (this *QFile) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qfile.h:133
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool resize(qint64)
func (this *QFile) Resize(sz int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6resizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sz)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:134
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool resize(const QString &, qint64)
func (this *QFile) Resize_1(filename string, sz int64) bool {
	var tmpArg0 = NewQString_5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6resizeERK7QStringx", qtrt.FFI_TYPE_POINTER, convArg0, sz)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Resize_1(filename string, sz int64) bool {
	var nilthis *QFile
	rv := nilthis.Resize_1(filename, sz)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:136
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QFileDevice::Permissions permissions()
func (this *QFile) Permissions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile11permissionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qfile.h:137
// index:1
// Public static Visibility=Default Availability=Available
// [4] QFileDevice::Permissions permissions(const QString &)
func (this *QFile) Permissions_1(filename string) int {
	var tmpArg0 = NewQString_5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile11permissionsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QFile_Permissions_1(filename string) int {
	var nilthis *QFile
	rv := nilthis.Permissions_1(filename)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setPermissions(QFileDevice::Permissions)
func (this *QFile) SetPermissions(permissionSpec int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile14setPermissionsE6QFlagsIN11QFileDevice10PermissionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), permissionSpec)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:139
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool setPermissions(const QString &, QFileDevice::Permissions)
func (this *QFile) SetPermissions_1(filename string, permissionSpec int) bool {
	var tmpArg0 = NewQString_5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile14setPermissionsERK7QString6QFlagsIN11QFileDevice10PermissionEE", qtrt.FFI_TYPE_POINTER, convArg0, permissionSpec)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QFile_SetPermissions_1(filename string, permissionSpec int) bool {
	var nilthis *QFile
	rv := nilthis.SetPermissions_1(filename, permissionSpec)
	return rv
}

//  body block end
