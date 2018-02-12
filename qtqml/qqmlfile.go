package qtqml

// /usr/include/qt/QtQml/qqmlfile.h
// #include <qqmlfile.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QQmlFile struct {
	*qtrt.CObject
}

func (this *QQmlFile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlFile) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlFileFromPointer(cthis unsafe.Pointer) *QQmlFile {
	return &QQmlFile{&qtrt.CObject{cthis}}
}
func (*QQmlFile) NewFromPointer(cthis unsafe.Pointer) *QQmlFile {
	return NewQQmlFileFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlfile.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlFile()
func NewQQmlFile() *QQmlFile {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFileC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlFile)
	return gothis
}

// /usr/include/qt/QtQml/qqmlfile.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlFile(QQmlEngine *, const QUrl &)
func NewQQmlFile_1(arg0 *QQmlEngine /*777 QQmlEngine **/, arg1 *qtcore.QUrl) *QQmlFile {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFileC2EP10QQmlEngineRK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlFile)
	return gothis
}

// /usr/include/qt/QtQml/qqmlfile.h:58
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlFile(QQmlEngine *, const QString &)
func NewQQmlFile_2(arg0 *QQmlEngine /*777 QQmlEngine **/, arg1 string) *QQmlFile {
	var convArg0 = arg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFileC2EP10QQmlEngineRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlFile)
	return gothis
}

// /usr/include/qt/QtQml/qqmlfile.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QQmlFile()
func DeleteQQmlFile(this *QQmlFile) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFileD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlfile.h:63
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QQmlFile) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QQmlFile6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReady()
func (this *QQmlFile) IsReady() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QQmlFile7isReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isError()
func (this *QQmlFile) IsError() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QQmlFile7isErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLoading()
func (this *QQmlFile) IsLoading() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QQmlFile9isLoadingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url()
func (this *QQmlFile) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QQmlFile3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQml/qqmlfile.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] QQmlFile::Status status()
func (this *QQmlFile) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QQmlFile6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString error()
func (this *QQmlFile) Error() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QQmlFile5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlfile.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 size()
func (this *QQmlFile) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QQmlFile4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtQml/qqmlfile.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * data()
func (this *QQmlFile) Data() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QQmlFile4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray dataByteArray()
func (this *QQmlFile) DataByteArray() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QQmlFile13dataByteArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtQml/qqmlfile.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void load(QQmlEngine *, const QUrl &)
func (this *QQmlFile) Load(arg0 *QQmlEngine /*777 QQmlEngine **/, arg1 *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile4loadEP10QQmlEngineRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:78
// index:1
// Public Visibility=Default Availability=Available
// [-2] void load(QQmlEngine *, const QString &)
func (this *QQmlFile) Load_1(arg0 *QQmlEngine /*777 QQmlEngine **/, arg1 string) {
	var convArg0 = arg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile4loadEP10QQmlEngineRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QQmlFile) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void clear(QObject *)
func (this *QQmlFile) Clear_1(arg0 *qtcore.QObject /*777 QObject **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile5clearEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool connectFinished(QObject *, const char *)
func (this *QQmlFile) ConnectFinished(arg0 *qtcore.QObject /*777 QObject **/, arg1 string) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = qtrt.CString(arg1)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile15connectFinishedEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:85
// index:1
// Public Visibility=Default Availability=Available
// [1] bool connectFinished(QObject *, int)
func (this *QQmlFile) ConnectFinished_1(arg0 *qtcore.QObject /*777 QObject **/, arg1 int) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile15connectFinishedEP7QObjecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:86
// index:0
// Public Visibility=Default Availability=Available
// [1] bool connectDownloadProgress(QObject *, const char *)
func (this *QQmlFile) ConnectDownloadProgress(arg0 *qtcore.QObject /*777 QObject **/, arg1 string) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = qtrt.CString(arg1)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile23connectDownloadProgressEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:87
// index:1
// Public Visibility=Default Availability=Available
// [1] bool connectDownloadProgress(QObject *, int)
func (this *QQmlFile) ConnectDownloadProgress_1(arg0 *qtcore.QObject /*777 QObject **/, arg1 int) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile23connectDownloadProgressEP7QObjecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:90
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isSynchronous(const QString &)
func (this *QQmlFile) IsSynchronous(url string) bool {
	var tmpArg0 = qtcore.NewQString_5(url)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile13isSynchronousERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QQmlFile_IsSynchronous(url string) bool {
	var nilthis *QQmlFile
	rv := nilthis.IsSynchronous(url)
	return rv
}

// /usr/include/qt/QtQml/qqmlfile.h:91
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isSynchronous(const QUrl &)
func (this *QQmlFile) IsSynchronous_1(url *qtcore.QUrl) bool {
	var convArg0 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile13isSynchronousERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QQmlFile_IsSynchronous_1(url *qtcore.QUrl) bool {
	var nilthis *QQmlFile
	rv := nilthis.IsSynchronous_1(url)
	return rv
}

// /usr/include/qt/QtQml/qqmlfile.h:93
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isLocalFile(const QString &)
func (this *QQmlFile) IsLocalFile(url string) bool {
	var tmpArg0 = qtcore.NewQString_5(url)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile11isLocalFileERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QQmlFile_IsLocalFile(url string) bool {
	var nilthis *QQmlFile
	rv := nilthis.IsLocalFile(url)
	return rv
}

// /usr/include/qt/QtQml/qqmlfile.h:94
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isLocalFile(const QUrl &)
func (this *QQmlFile) IsLocalFile_1(url *qtcore.QUrl) bool {
	var convArg0 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile11isLocalFileERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QQmlFile_IsLocalFile_1(url *qtcore.QUrl) bool {
	var nilthis *QQmlFile
	rv := nilthis.IsLocalFile_1(url)
	return rv
}

// /usr/include/qt/QtQml/qqmlfile.h:96
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString urlToLocalFileOrQrc(const QString &)
func (this *QQmlFile) UrlToLocalFileOrQrc(arg0 string) string {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile19urlToLocalFileOrQrcERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QQmlFile_UrlToLocalFileOrQrc(arg0 string) string {
	var nilthis *QQmlFile
	rv := nilthis.UrlToLocalFileOrQrc(arg0)
	return rv
}

// /usr/include/qt/QtQml/qqmlfile.h:97
// index:1
// Public static Visibility=Default Availability=Available
// [8] QString urlToLocalFileOrQrc(const QUrl &)
func (this *QQmlFile) UrlToLocalFileOrQrc_1(arg0 *qtcore.QUrl) string {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlFile19urlToLocalFileOrQrcERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QQmlFile_UrlToLocalFileOrQrc_1(arg0 *qtcore.QUrl) string {
	var nilthis *QQmlFile
	rv := nilthis.UrlToLocalFileOrQrc_1(arg0)
	return rv
}

type QQmlFile__Status = int

const QQmlFile__Null QQmlFile__Status = 0
const QQmlFile__Ready QQmlFile__Status = 1
const QQmlFile__Error QQmlFile__Status = 2
const QQmlFile__Loading QQmlFile__Status = 3

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
