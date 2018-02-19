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
func (this *QMimeDatabase) MimeTypeForName(nameOrAlias string) *QMimeType /*123*/ {
	var tmpArg0 = NewQString_5(nameOrAlias)
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
// [8] QMimeType mimeTypeForFile(const QString &, enum QMimeDatabase::MatchMode) const
func (this *QMimeDatabase) MimeTypeForFile(fileName string, mode int) *QMimeType /*123*/ {
	var tmpArg0 = NewQString_5(fileName)
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
// [8] QMimeType mimeTypeForFile(const QString &, enum QMimeDatabase::MatchMode) const
func (this *QMimeDatabase) MimeTypeForFile__(fileName string) *QMimeType /*123*/ {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QMimeDatabase::MatchMode=Enum, QMimeDatabase::MatchMode=Enum,
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
// [8] QMimeType mimeTypeForFile(const QFileInfo &, enum QMimeDatabase::MatchMode) const
func (this *QMimeDatabase) MimeTypeForFile_1(fileInfo QFileInfo_ITF, mode int) *QMimeType /*123*/ {
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
// [8] QMimeType mimeTypeForFile(const QFileInfo &, enum QMimeDatabase::MatchMode) const
func (this *QMimeDatabase) MimeTypeForFile_1_(fileInfo QFileInfo_ITF) *QMimeType /*123*/ {
	var convArg0 unsafe.Pointer
	if fileInfo != nil && fileInfo.QFileInfo_PTR() != nil {
		convArg0 = fileInfo.QFileInfo_PTR().GetCthis()
	}
	// arg: 1, QMimeDatabase::MatchMode=Enum, QMimeDatabase::MatchMode=Enum,
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
func (this *QMimeDatabase) MimeTypeForData_1(device QIODevice_ITF /*777 QIODevice **/) *QMimeType /*123*/ {
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
func (this *QMimeDatabase) MimeTypeForFileNameAndData(fileName string, device QIODevice_ITF /*777 QIODevice **/) *QMimeType /*123*/ {
	var tmpArg0 = NewQString_5(fileName)
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
func (this *QMimeDatabase) MimeTypeForFileNameAndData_1(fileName string, data QByteArray_ITF) *QMimeType /*123*/ {
	var tmpArg0 = NewQString_5(fileName)
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
func (this *QMimeDatabase) SuffixForFileName(fileName string) string {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMimeDatabase17suffixForFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

type QMimeDatabase__MatchMode = int

const QMimeDatabase__MatchDefault QMimeDatabase__MatchMode = 0
const QMimeDatabase__MatchExtension QMimeDatabase__MatchMode = 1
const QMimeDatabase__MatchContent QMimeDatabase__MatchMode = 2

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
