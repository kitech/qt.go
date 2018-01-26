package qtqml

// /usr/include/qt/QtQml/qqmlfile.h
// #include <qqmlfile.h>
// #include <QtQml>

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
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQQmlFileFromPointer(cthis unsafe.Pointer) *QQmlFile {
	return &QQmlFile{&qtrt.CObject{cthis}}
}
func (*QQmlFile) NewFromPointer(cthis unsafe.Pointer) *QQmlFile {
	return NewQQmlFileFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlfile.h:56
// index:0
// Public
// void QQmlFile()
func NewQQmlFile() *QQmlFile {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFileC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlfile.h:57
// index:1
// Public
// void QQmlFile(class QQmlEngine *, const class QUrl &)
func NewQQmlFile_1(arg0 *QQmlEngine /*777 QQmlEngine **/, arg1 *qtcore.QUrl) *QQmlFile {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFileC2EP10QQmlEngineRK4QUrl", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlfile.h:58
// index:2
// Public
// void QQmlFile(class QQmlEngine *, const class QString &)
func NewQQmlFile_2(arg0 *QQmlEngine /*777 QQmlEngine **/, arg1 *qtcore.QString) *QQmlFile {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFileC2EP10QQmlEngineRK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlfile.h:59
// index:0
// Public
// void ~QQmlFile()
func DeleteQQmlFile(*QQmlFile) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFileD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:63
// index:0
// Public
// bool isNull()
func (this *QQmlFile) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QQmlFile6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:64
// index:0
// Public
// bool isReady()
func (this *QQmlFile) IsReady() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QQmlFile7isReadyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:65
// index:0
// Public
// bool isError()
func (this *QQmlFile) IsError() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QQmlFile7isErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:66
// index:0
// Public
// bool isLoading()
func (this *QQmlFile) IsLoading() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QQmlFile9isLoadingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:68
// index:0
// Public
// QUrl url()
func (this *QQmlFile) Url() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QQmlFile3urlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlfile.h:70
// index:0
// Public
// QQmlFile::Status status()
func (this *QQmlFile) Status() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QQmlFile6statusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:71
// index:0
// Public
// QString error()
func (this *QQmlFile) Error() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QQmlFile5errorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlfile.h:73
// index:0
// Public
// qint64 size()
func (this *QQmlFile) Size() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QQmlFile4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtQml/qqmlfile.h:74
// index:0
// Public
// const char * data()
func (this *QQmlFile) Data() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QQmlFile4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:75
// index:0
// Public
// QByteArray dataByteArray()
func (this *QQmlFile) DataByteArray() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QQmlFile13dataByteArrayEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlfile.h:77
// index:0
// Public
// void load(class QQmlEngine *, const class QUrl &)
func (this *QQmlFile) Load(arg0 *QQmlEngine /*777 QQmlEngine **/, arg1 *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile4loadEP10QQmlEngineRK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:78
// index:1
// Public
// void load(class QQmlEngine *, const class QString &)
func (this *QQmlFile) Load_1(arg0 *QQmlEngine /*777 QQmlEngine **/, arg1 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile4loadEP10QQmlEngineRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:80
// index:0
// Public
// void clear()
func (this *QQmlFile) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:81
// index:1
// Public
// void clear(class QObject *)
func (this *QQmlFile) Clear_1(arg0 *qtcore.QObject /*777 QObject **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile5clearEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfile.h:84
// index:0
// Public
// bool connectFinished(class QObject *, const char *)
func (this *QQmlFile) ConnectFinished(arg0 *qtcore.QObject /*777 QObject **/, arg1 string) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = qtrt.CString(arg1)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile15connectFinishedEP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:85
// index:1
// Public
// bool connectFinished(class QObject *, int)
func (this *QQmlFile) ConnectFinished_1(arg0 *qtcore.QObject /*777 QObject **/, arg1 int) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile15connectFinishedEP7QObjecti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:86
// index:0
// Public
// bool connectDownloadProgress(class QObject *, const char *)
func (this *QQmlFile) ConnectDownloadProgress(arg0 *qtcore.QObject /*777 QObject **/, arg1 string) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = qtrt.CString(arg1)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile23connectDownloadProgressEP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:87
// index:1
// Public
// bool connectDownloadProgress(class QObject *, int)
func (this *QQmlFile) ConnectDownloadProgress_1(arg0 *qtcore.QObject /*777 QObject **/, arg1 int) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile23connectDownloadProgressEP7QObjecti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlfile.h:90
// index:0
// Public static
// bool isSynchronous(const class QString &)
func (this *QQmlFile) IsSynchronous(url *qtcore.QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile13isSynchronousERK7QString", ffiqt.FFI_TYPE_POINTER, url)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QQmlFile_IsSynchronous(url *qtcore.QString) bool {
	var nilthis *QQmlFile
	rv := nilthis.IsSynchronous(url)
	return rv
}

// /usr/include/qt/QtQml/qqmlfile.h:91
// index:1
// Public static
// bool isSynchronous(const class QUrl &)
func (this *QQmlFile) IsSynchronous_1(url *qtcore.QUrl) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile13isSynchronousERK4QUrl", ffiqt.FFI_TYPE_POINTER, url)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QQmlFile_IsSynchronous_1(url *qtcore.QUrl) bool {
	var nilthis *QQmlFile
	rv := nilthis.IsSynchronous_1(url)
	return rv
}

// /usr/include/qt/QtQml/qqmlfile.h:93
// index:0
// Public static
// bool isLocalFile(const class QString &)
func (this *QQmlFile) IsLocalFile(url *qtcore.QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile11isLocalFileERK7QString", ffiqt.FFI_TYPE_POINTER, url)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QQmlFile_IsLocalFile(url *qtcore.QString) bool {
	var nilthis *QQmlFile
	rv := nilthis.IsLocalFile(url)
	return rv
}

// /usr/include/qt/QtQml/qqmlfile.h:94
// index:1
// Public static
// bool isLocalFile(const class QUrl &)
func (this *QQmlFile) IsLocalFile_1(url *qtcore.QUrl) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile11isLocalFileERK4QUrl", ffiqt.FFI_TYPE_POINTER, url)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QQmlFile_IsLocalFile_1(url *qtcore.QUrl) bool {
	var nilthis *QQmlFile
	rv := nilthis.IsLocalFile_1(url)
	return rv
}

// /usr/include/qt/QtQml/qqmlfile.h:96
// index:0
// Public static
// QString urlToLocalFileOrQrc(const class QString &)
func (this *QQmlFile) UrlToLocalFileOrQrc(arg0 *qtcore.QString) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile19urlToLocalFileOrQrcERK7QString", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QQmlFile_UrlToLocalFileOrQrc(arg0 *qtcore.QString) *qtcore.QString /*123*/ {
	var nilthis *QQmlFile
	rv := nilthis.UrlToLocalFileOrQrc(arg0)
	return rv
}

// /usr/include/qt/QtQml/qqmlfile.h:97
// index:1
// Public static
// QString urlToLocalFileOrQrc(const class QUrl &)
func (this *QQmlFile) UrlToLocalFileOrQrc_1(arg0 *qtcore.QUrl) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlFile19urlToLocalFileOrQrcERK4QUrl", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QQmlFile_UrlToLocalFileOrQrc_1(arg0 *qtcore.QUrl) *qtcore.QString /*123*/ {
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
