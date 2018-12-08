package qtcore

// /usr/include/qt/QtCore/qmimedatabase.h
// #include <qmimedatabase.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QMimeDatabase struct {
	*qtrt.CObject
}
type QMimeDatabase_ITF interface {
	QMimeDatabase_PTR() *QMimeDatabase
}

func (ptr *QMimeDatabase) QMimeDatabase_PTR() *QMimeDatabase { return ptr }

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

/*
Constructs a QMimeDatabase object.

It is perfectly OK to create an instance of QMimeDatabase every time you need to perform a lookup. The parsing of mimetypes is done on demand (when shared-mime-info is installed) or when the very first instance is constructed (when parsing XML files directly).
*/
func (*QMimeDatabase) NewForInherit() *QMimeDatabase {
	return NewQMimeDatabase()
}
func NewQMimeDatabase() *QMimeDatabase {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMimeDatabaseC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMimeDatabaseFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMimeDatabase)
	return gothis
}

// /usr/include/qt/QtCore/qmimedatabase.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMimeDatabase()

/*

 */
func DeleteQMimeDatabase(this *QMimeDatabase) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMimeDatabaseD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmimedatabase.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForName(const QString &) const

/*
Returns a MIME type for nameOrAlias or an invalid one if none found.
*/
func (this *QMimeDatabase) MimeTypeForName(nameOrAlias string) *QMimeType /*123*/ {
	var tmpArg0 = NewQString5(nameOrAlias)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForFile(const QString &, QMimeDatabase::MatchMode) const

/*
Returns a MIME type for fileInfo.

A valid MIME type is always returned.

The default matching algorithm looks at both the file name and the file contents, if necessary. The file extension has priority over the contents, but the contents will be used if the file extension is unknown, or matches multiple MIME types. If fileInfo is a Unix symbolic link, the file that it refers to will be used instead. If the file doesn't match any known pattern or data, the default MIME type (application/octet-stream) is returned.

When mode is set to MatchExtension, only the file name is used, not the file contents. The file doesn't even have to exist. If the file name doesn't match any known pattern, the default MIME type (application/octet-stream) is returned. If multiple MIME types match this file, the first one (alphabetically) is returned.

When mode is set to MatchContent, and the file is readable, only the file contents are used to determine the MIME type. This is equivalent to calling mimeTypeForData with a QFile as input device.

fileInfo may refer to an absolute or relative path.

See also QMimeType::isDefault() and mimeTypeForData().
*/
func (this *QMimeDatabase) MimeTypeForFile(fileName string, mode int) *QMimeType /*123*/ {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK7QStringNS_9MatchModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForFile(const QString &, QMimeDatabase::MatchMode) const

/*
Returns a MIME type for fileInfo.

A valid MIME type is always returned.

The default matching algorithm looks at both the file name and the file contents, if necessary. The file extension has priority over the contents, but the contents will be used if the file extension is unknown, or matches multiple MIME types. If fileInfo is a Unix symbolic link, the file that it refers to will be used instead. If the file doesn't match any known pattern or data, the default MIME type (application/octet-stream) is returned.

When mode is set to MatchExtension, only the file name is used, not the file contents. The file doesn't even have to exist. If the file name doesn't match any known pattern, the default MIME type (application/octet-stream) is returned. If multiple MIME types match this file, the first one (alphabetically) is returned.

When mode is set to MatchContent, and the file is readable, only the file contents are used to determine the MIME type. This is equivalent to calling mimeTypeForData with a QFile as input device.

fileInfo may refer to an absolute or relative path.

See also QMimeType::isDefault() and mimeTypeForData().
*/
func (this *QMimeDatabase) MimeTypeForFilep(fileName string) *QMimeType /*123*/ {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QMimeDatabase::MatchMode=Enum, QMimeDatabase::MatchMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK7QStringNS_9MatchModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:75
// index:1
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForFile(const QFileInfo &, QMimeDatabase::MatchMode) const

/*
Returns a MIME type for fileInfo.

A valid MIME type is always returned.

The default matching algorithm looks at both the file name and the file contents, if necessary. The file extension has priority over the contents, but the contents will be used if the file extension is unknown, or matches multiple MIME types. If fileInfo is a Unix symbolic link, the file that it refers to will be used instead. If the file doesn't match any known pattern or data, the default MIME type (application/octet-stream) is returned.

When mode is set to MatchExtension, only the file name is used, not the file contents. The file doesn't even have to exist. If the file name doesn't match any known pattern, the default MIME type (application/octet-stream) is returned. If multiple MIME types match this file, the first one (alphabetically) is returned.

When mode is set to MatchContent, and the file is readable, only the file contents are used to determine the MIME type. This is equivalent to calling mimeTypeForData with a QFile as input device.

fileInfo may refer to an absolute or relative path.

See also QMimeType::isDefault() and mimeTypeForData().
*/
func (this *QMimeDatabase) MimeTypeForFile1(fileInfo QFileInfo_ITF, mode int) *QMimeType /*123*/ {
	var convArg0 unsafe.Pointer
	if fileInfo != nil && fileInfo.QFileInfo_PTR() != nil {
		convArg0 = fileInfo.QFileInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK9QFileInfoNS_9MatchModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:75
// index:1
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForFile(const QFileInfo &, QMimeDatabase::MatchMode) const

/*
Returns a MIME type for fileInfo.

A valid MIME type is always returned.

The default matching algorithm looks at both the file name and the file contents, if necessary. The file extension has priority over the contents, but the contents will be used if the file extension is unknown, or matches multiple MIME types. If fileInfo is a Unix symbolic link, the file that it refers to will be used instead. If the file doesn't match any known pattern or data, the default MIME type (application/octet-stream) is returned.

When mode is set to MatchExtension, only the file name is used, not the file contents. The file doesn't even have to exist. If the file name doesn't match any known pattern, the default MIME type (application/octet-stream) is returned. If multiple MIME types match this file, the first one (alphabetically) is returned.

When mode is set to MatchContent, and the file is readable, only the file contents are used to determine the MIME type. This is equivalent to calling mimeTypeForData with a QFile as input device.

fileInfo may refer to an absolute or relative path.

See also QMimeType::isDefault() and mimeTypeForData().
*/
func (this *QMimeDatabase) MimeTypeForFile1p(fileInfo QFileInfo_ITF) *QMimeType /*123*/ {
	var convArg0 unsafe.Pointer
	if fileInfo != nil && fileInfo.QFileInfo_PTR() != nil {
		convArg0 = fileInfo.QFileInfo_PTR().GetCthis()
	}
	// arg: 1, QMimeDatabase::MatchMode=Enum, QMimeDatabase::MatchMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForFileERK9QFileInfoNS_9MatchModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForData(const QByteArray &) const

/*
Returns a MIME type for data.

A valid MIME type is always returned. If data doesn't match any known MIME type data, the default MIME type (application/octet-stream) is returned.
*/
func (this *QMimeDatabase) MimeTypeForData(data QByteArray_ITF) *QMimeType /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:79
// index:1
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForData(QIODevice *) const

/*
Returns a MIME type for data.

A valid MIME type is always returned. If data doesn't match any known MIME type data, the default MIME type (application/octet-stream) is returned.
*/
func (this *QMimeDatabase) MimeTypeForData1(device QIODevice_ITF /*777 QIODevice **/) *QMimeType /*123*/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase15mimeTypeForDataEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForUrl(const QUrl &) const

/*
Returns a MIME type for url.

If the URL is a local file, this calls mimeTypeForFile.

Otherwise the matching is done based on the file name only, except for schemes where file names don't mean much, like HTTP. This method always returns the default mimetype for HTTP URLs, use QNetworkAccessManager to handle HTTP URLs properly.

A valid MIME type is always returned. If url doesn't match any known MIME type data, the default MIME type (application/octet-stream) is returned.
*/
func (this *QMimeDatabase) MimeTypeForUrl(url QUrl_ITF) *QMimeType /*123*/ {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase14mimeTypeForUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForFileNameAndData(const QString &, QIODevice *) const

/*
Returns a MIME type for the given fileName and device data.

This overload can be useful when the file is remote, and we started to download some of its data in a device. This allows to do full MIME type matching for remote files as well.

If the device is not open, it will be opened by this function, and closed after the MIME type detection is completed.

A valid MIME type is always returned. If device data doesn't match any known MIME type data, the default MIME type (application/octet-stream) is returned.

This method looks at both the file name and the file contents, if necessary. The file extension has priority over the contents, but the contents will be used if the file extension is unknown, or matches multiple MIME types.
*/
func (this *QMimeDatabase) MimeTypeForFileNameAndData(fileName string, device QIODevice_ITF /*777 QIODevice **/) *QMimeType /*123*/ {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg1 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:83
// index:1
// Public Visibility=Default Availability=Available
// [8] QMimeType mimeTypeForFileNameAndData(const QString &, const QByteArray &) const

/*
Returns a MIME type for the given fileName and device data.

This overload can be useful when the file is remote, and we started to download some of its data in a device. This allows to do full MIME type matching for remote files as well.

If the device is not open, it will be opened by this function, and closed after the MIME type detection is completed.

A valid MIME type is always returned. If device data doesn't match any known MIME type data, the default MIME type (application/octet-stream) is returned.

This method looks at both the file name and the file contents, if necessary. The file extension has priority over the contents, but the contents will be used if the file extension is unknown, or matches multiple MIME types.
*/
func (this *QMimeDatabase) MimeTypeForFileNameAndData1(fileName string, data QByteArray_ITF) *QMimeType /*123*/ {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg1 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimedatabase.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QString suffixForFileName(const QString &) const

/*
Returns the suffix for the file fileName, as known by the MIME database.

This allows to pre-select "tar.bz2" for foo.tar.bz2, but still only "txt" for my.file.with.dots.txt.
*/
func (this *QMimeDatabase) SuffixForFileName(fileName string) string {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase17suffixForFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

/*
This enum specifies how matching a file to a MIME type is performed.


*/
type QMimeDatabase__MatchMode = int

//
const QMimeDatabase__MatchDefault QMimeDatabase__MatchMode = 0

//
const QMimeDatabase__MatchExtension QMimeDatabase__MatchMode = 1

//
const QMimeDatabase__MatchContent QMimeDatabase__MatchMode = 2

func (this *QMimeDatabase) MatchModeItemName(val int) string {
	switch val {
	case QMimeDatabase__MatchDefault: // 0
		return "MatchDefault"
	case QMimeDatabase__MatchExtension: // 1
		return "MatchExtension"
	case QMimeDatabase__MatchContent: // 2
		return "MatchContent"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QMimeDatabase_MatchModeItemName(val int) string {
	var nilthis *QMimeDatabase
	return nilthis.MatchModeItemName(val)
}

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
