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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qmimedatabase.h:63
// index:0
// void QMimeDatabase()
func NewQMimeDatabase() *QMimeDatabase {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMimeDatabaseC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QMimeDatabase{cthis}
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
	// 0: (, const QString & nameOrAlias), (nameOrAlias)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, nameOrAlias)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:74
// index:0
// QMimeType mimeTypeForFile(const class QString &, enum QMimeDatabase::MatchMode)
func (this *QMimeDatabase) MimeTypeForFile(fileName unsafe.Pointer, mode int) {
	// 0: (, const QString & fileName, QMimeDatabase::MatchMode mode), (fileName, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK7QStringNS_9MatchModeE", ffiqt.FFI_TYPE_VOID, this.cthis, fileName, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:75
// index:1
// QMimeType mimeTypeForFile(const class QFileInfo &, enum QMimeDatabase::MatchMode)
func (this *QMimeDatabase) MimeTypeForFile_1(fileInfo unsafe.Pointer, mode int) {
	// 1: (, const QFileInfo & fileInfo, QMimeDatabase::MatchMode mode), (fileInfo, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK9QFileInfoNS_9MatchModeE", ffiqt.FFI_TYPE_VOID, this.cthis, fileInfo, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:76
// index:0
// QList<QMimeType> mimeTypesForFileName(const class QString &)
func (this *QMimeDatabase) MimeTypesForFileName(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase20mimeTypesForFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:78
// index:0
// QMimeType mimeTypeForData(const class QByteArray &)
func (this *QMimeDatabase) MimeTypeForData(data unsafe.Pointer) {
	// 0: (, const QByteArray & data), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForDataERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:79
// index:1
// QMimeType mimeTypeForData(class QIODevice *)
func (this *QMimeDatabase) MimeTypeForData_1(device unsafe.Pointer) {
	// 1: (, QIODevice * device), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForDataEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.cthis, device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:81
// index:0
// QMimeType mimeTypeForUrl(const class QUrl &)
func (this *QMimeDatabase) MimeTypeForUrl(url unsafe.Pointer) {
	// 0: (, const QUrl & url), (url)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase14mimeTypeForUrlERK4QUrl", ffiqt.FFI_TYPE_VOID, this.cthis, url)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:82
// index:0
// QMimeType mimeTypeForFileNameAndData(const class QString &, class QIODevice *)
func (this *QMimeDatabase) MimeTypeForFileNameAndData(fileName unsafe.Pointer, device unsafe.Pointer) {
	// 0: (, const QString & fileName, QIODevice * device), (fileName, device)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringP9QIODevice", ffiqt.FFI_TYPE_VOID, this.cthis, fileName, device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:83
// index:1
// QMimeType mimeTypeForFileNameAndData(const class QString &, const class QByteArray &)
func (this *QMimeDatabase) MimeTypeForFileNameAndData_1(fileName unsafe.Pointer, data unsafe.Pointer) {
	// 1: (, const QString & fileName, const QByteArray & data), (fileName, data)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, fileName, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:85
// index:0
// QString suffixForFileName(const class QString &)
func (this *QMimeDatabase) SuffixForFileName(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase17suffixForFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedatabase.h:87
// index:0
// QList<QMimeType> allMimeTypes()
func (this *QMimeDatabase) AllMimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMimeDatabase12allMimeTypesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
