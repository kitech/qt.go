//  header block begin
// /usr/include/qt/QtCore/qmimedatabase.h
// #include <qmimedatabase.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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

// /usr/include/qt/QtCore/qmimedatabase.h:63
// index:0
// void QMimeDatabase()
func NewQMimeDatabase() *QMimeDatabase {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMimeDatabaseC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMimeDatabaseFromPointer(cthis)
	return gothis
}
func NewQMimeDatabaseFromPointer(cthis unsafe.Pointer) *QMimeDatabase {
	return &QMimeDatabase{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qmimedatabase.h:64
// index:0
// void ~QMimeDatabase()
func DeleteQMimeDatabase(*QMimeDatabase) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMimeDatabaseD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:66
// index:0
// QMimeType mimeTypeForName(const class QString &)
func (this *QMimeDatabase) MimeTypeForName(nameOrAlias unsafe.Pointer) {
	// 0: (, nameOrAlias const QString &), (nameOrAlias)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForNameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), nameOrAlias)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:74
// index:0
// QMimeType mimeTypeForFile(const class QString &, enum QMimeDatabase::MatchMode)
func (this *QMimeDatabase) MimeTypeForFile(fileName unsafe.Pointer, mode int) {
	// 0: (, fileName const QString &, mode QMimeDatabase::MatchMode), (fileName, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK7QStringNS_9MatchModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:75
// index:1
// QMimeType mimeTypeForFile(const class QFileInfo &, enum QMimeDatabase::MatchMode)
func (this *QMimeDatabase) MimeTypeForFile_1(fileInfo unsafe.Pointer, mode int) {
	// 1: (, fileInfo const QFileInfo &, mode QMimeDatabase::MatchMode), (fileInfo, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK9QFileInfoNS_9MatchModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileInfo, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:76
// index:0
// QList<QMimeType> mimeTypesForFileName(const class QString &)
func (this *QMimeDatabase) MimeTypesForFileName(fileName unsafe.Pointer) {
	// 0: (, fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase20mimeTypesForFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:78
// index:0
// QMimeType mimeTypeForData(const class QByteArray &)
func (this *QMimeDatabase) MimeTypeForData(data unsafe.Pointer) {
	// 0: (, data const QByteArray &), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForDataERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:79
// index:1
// QMimeType mimeTypeForData(class QIODevice *)
func (this *QMimeDatabase) MimeTypeForData_1(device unsafe.Pointer) {
	// 1: (, device QIODevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForDataEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:81
// index:0
// QMimeType mimeTypeForUrl(const class QUrl &)
func (this *QMimeDatabase) MimeTypeForUrl(url unsafe.Pointer) {
	// 0: (, url const QUrl &), (url)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase14mimeTypeForUrlERK4QUrl", ffiqt.FFI_TYPE_VOID, this.GetCthis(), url)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:82
// index:0
// QMimeType mimeTypeForFileNameAndData(const class QString &, class QIODevice *)
func (this *QMimeDatabase) MimeTypeForFileNameAndData(fileName unsafe.Pointer, device unsafe.Pointer) {
	// 0: (, fileName const QString &, device QIODevice *), (fileName, device)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringP9QIODevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName, device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:83
// index:1
// QMimeType mimeTypeForFileNameAndData(const class QString &, const class QByteArray &)
func (this *QMimeDatabase) MimeTypeForFileNameAndData_1(fileName unsafe.Pointer, data unsafe.Pointer) {
	// 1: (, fileName const QString &, data const QByteArray &), (fileName, data)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:85
// index:0
// QString suffixForFileName(const class QString &)
func (this *QMimeDatabase) SuffixForFileName(fileName unsafe.Pointer) {
	// 0: (, fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase17suffixForFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:87
// index:0
// QList<QMimeType> allMimeTypes()
func (this *QMimeDatabase) AllMimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase12allMimeTypesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
