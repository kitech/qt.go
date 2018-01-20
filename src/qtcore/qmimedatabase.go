//  header block begin
// /usr/include/qt/QtCore/qmimedatabase.h
// #include <qmimedatabase.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QMimeDatabase struct {
	*qtrt.CObject
}

func (this *QMimeDatabase) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQMimeDatabaseFromPointer(cthis unsafe.Pointer) *QMimeDatabase {
	return &QMimeDatabase{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qmimedatabase.h:63
// index:0
// Public
// void QMimeDatabase()
func NewQMimeDatabase() *QMimeDatabase {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMimeDatabaseC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMimeDatabaseFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmimedatabase.h:64
// index:0
// Public
// void ~QMimeDatabase()
func DeleteQMimeDatabase(*QMimeDatabase) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMimeDatabaseD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:66
// index:0
// Public
// QMimeType mimeTypeForName(const class QString &)
func (this *QMimeDatabase) MimeTypeForName(nameOrAlias *QString) interface{} {
	var convArg0 = nameOrAlias.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedatabase.h:74
// index:0
// Public
// QMimeType mimeTypeForFile(const class QString &, enum QMimeDatabase::MatchMode)
func (this *QMimeDatabase) MimeTypeForFile(fileName *QString, mode int) interface{} {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK7QStringNS_9MatchModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedatabase.h:75
// index:1
// Public
// QMimeType mimeTypeForFile(const class QFileInfo &, enum QMimeDatabase::MatchMode)
func (this *QMimeDatabase) MimeTypeForFile_1(fileInfo *QFileInfo, mode int) interface{} {
	var convArg0 = fileInfo.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK9QFileInfoNS_9MatchModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedatabase.h:76
// index:0
// Public
// QList<QMimeType> mimeTypesForFileName(const class QString &)
func (this *QMimeDatabase) MimeTypesForFileName(fileName *QString) interface{} {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase20mimeTypesForFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedatabase.h:78
// index:0
// Public
// QMimeType mimeTypeForData(const class QByteArray &)
func (this *QMimeDatabase) MimeTypeForData(data *QByteArray) interface{} {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForDataERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedatabase.h:79
// index:1
// Public
// QMimeType mimeTypeForData(class QIODevice *)
func (this *QMimeDatabase) MimeTypeForData_1(device unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForDataEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedatabase.h:81
// index:0
// Public
// QMimeType mimeTypeForUrl(const class QUrl &)
func (this *QMimeDatabase) MimeTypeForUrl(url *QUrl) interface{} {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase14mimeTypeForUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedatabase.h:82
// index:0
// Public
// QMimeType mimeTypeForFileNameAndData(const class QString &, class QIODevice *)
func (this *QMimeDatabase) MimeTypeForFileNameAndData(fileName *QString, device unsafe.Pointer) interface{} {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, device)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedatabase.h:83
// index:1
// Public
// QMimeType mimeTypeForFileNameAndData(const class QString &, const class QByteArray &)
func (this *QMimeDatabase) MimeTypeForFileNameAndData_1(fileName *QString, data *QByteArray) interface{} {
	var convArg0 = fileName.GetCthis()
	var convArg1 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedatabase.h:85
// index:0
// Public
// QString suffixForFileName(const class QString &)
func (this *QMimeDatabase) SuffixForFileName(fileName *QString) interface{} {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase17suffixForFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedatabase.h:87
// index:0
// Public
// QList<QMimeType> allMimeTypes()
func (this *QMimeDatabase) AllMimeTypes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase12allMimeTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
