package qtcore

// /usr/include/qt/QtCore/qmimedatabase.h
// #include <qmimedatabase.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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

type QMimeDatabase struct {
	*qtrt.CObject
}

func (this *QMimeDatabase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMimeDatabase) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMimeDatabaseFromPointer(cthis unsafe.Pointer) *QMimeDatabase {
	return &QMimeDatabase{&qtrt.CObject{cthis}}
}
func (*QMimeDatabase) NewFromPointer(cthis unsafe.Pointer) *QMimeDatabase {
	return NewQMimeDatabaseFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmimedatabase.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMimeDatabase()
func NewQMimeDatabase() *QMimeDatabase {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMimeDatabaseC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQMimeDatabaseFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMimeDatabase)
	return gothis
}

// /usr/include/qt/QtCore/qmimedatabase.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMimeDatabase()
func DeleteQMimeDatabase(this *QMimeDatabase) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMimeDatabaseD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmimedatabase.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForName(const QString &)
func (this *QMimeDatabase) MimeTypeForName(nameOrAlias *QString) *QMimeType /*123*/ {
	var convArg0 = nameOrAlias.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForFile(const QString &, enum QMimeDatabase::MatchMode)
func (this *QMimeDatabase) MimeTypeForFile(fileName *QString, mode int) *QMimeType /*123*/ {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK7QStringNS_9MatchModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:75
// index:1
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForFile(const QFileInfo &, enum QMimeDatabase::MatchMode)
func (this *QMimeDatabase) MimeTypeForFile_1(fileInfo *QFileInfo, mode int) *QMimeType /*123*/ {
	var convArg0 = fileInfo.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK9QFileInfoNS_9MatchModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForData(const QByteArray &)
func (this *QMimeDatabase) MimeTypeForData(data *QByteArray) *QMimeType /*123*/ {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForDataERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:79
// index:1
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForData(QIODevice *)
func (this *QMimeDatabase) MimeTypeForData_1(device *QIODevice /*777 QIODevice **/) *QMimeType /*123*/ {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForDataEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForUrl(const QUrl &)
func (this *QMimeDatabase) MimeTypeForUrl(url *QUrl) *QMimeType /*123*/ {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase14mimeTypeForUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForFileNameAndData(const QString &, QIODevice *)
func (this *QMimeDatabase) MimeTypeForFileNameAndData(fileName *QString, device *QIODevice /*777 QIODevice **/) *QMimeType /*123*/ {
	var convArg0 = fileName.GetCthis()
	var convArg1 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:83
// index:1
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForFileNameAndData(const QString &, const QByteArray &)
func (this *QMimeDatabase) MimeTypeForFileNameAndData_1(fileName *QString, data *QByteArray) *QMimeType /*123*/ {
	var convArg0 = fileName.GetCthis()
	var convArg1 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QString suffixForFileName(const QString &)
func (this *QMimeDatabase) SuffixForFileName(fileName *QString) *QString /*123*/ {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase17suffixForFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

type QMimeDatabase__MatchMode = int

const QMimeDatabase__MatchDefault QMimeDatabase__MatchMode = 0
const QMimeDatabase__MatchExtension QMimeDatabase__MatchMode = 1
const QMimeDatabase__MatchContent QMimeDatabase__MatchMode = 2

//  body block end
