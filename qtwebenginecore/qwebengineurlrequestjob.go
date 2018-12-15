package qtwebenginecore

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestjob.h
// #include <qwebengineurlrequestjob.h>
// #include <QtWebEngineCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtquick"

//  ext block end

//  body block begin

/*

 */
type QWebEngineUrlRequestJob struct {
	*qtcore.QObject
}
type QWebEngineUrlRequestJob_ITF interface {
	qtcore.QObject_ITF
	QWebEngineUrlRequestJob_PTR() *QWebEngineUrlRequestJob
}

func (ptr *QWebEngineUrlRequestJob) QWebEngineUrlRequestJob_PTR() *QWebEngineUrlRequestJob { return ptr }

func (this *QWebEngineUrlRequestJob) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWebEngineUrlRequestJob) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWebEngineUrlRequestJobFromPointer(cthis unsafe.Pointer) *QWebEngineUrlRequestJob {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWebEngineUrlRequestJob{bcthis0}
}
func (*QWebEngineUrlRequestJob) NewFromPointer(cthis unsafe.Pointer) *QWebEngineUrlRequestJob {
	return NewQWebEngineUrlRequestJobFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestjob.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWebEngineUrlRequestJob) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWebEngineUrlRequestJob10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestjob.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWebEngineUrlRequestJob()

/*

 */
func DeleteQWebEngineUrlRequestJob(this *QWebEngineUrlRequestJob) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWebEngineUrlRequestJobD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestjob.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl requestUrl() const

/*

 */
func (this *QWebEngineUrlRequestJob) RequestUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWebEngineUrlRequestJob10requestUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestjob.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray requestMethod() const

/*

 */
func (this *QWebEngineUrlRequestJob) RequestMethod() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWebEngineUrlRequestJob13requestMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestjob.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl initiator() const

/*

 */
func (this *QWebEngineUrlRequestJob) Initiator() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWebEngineUrlRequestJob9initiatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestjob.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reply(const QByteArray &, QIODevice *)

/*

 */
func (this *QWebEngineUrlRequestJob) Reply(contentType qtcore.QByteArray_ITF, device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if contentType != nil && contentType.QByteArray_PTR() != nil {
		convArg0 = contentType.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg1 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWebEngineUrlRequestJob5replyERK10QByteArrayP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestjob.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fail(QWebEngineUrlRequestJob::Error)

/*

 */
func (this *QWebEngineUrlRequestJob) Fail(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWebEngineUrlRequestJob4failENS_5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestjob.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redirect(const QUrl &)

/*

 */
func (this *QWebEngineUrlRequestJob) Redirect(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWebEngineUrlRequestJob8redirectERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QWebEngineUrlRequestJob__Error = int

//
const QWebEngineUrlRequestJob__NoError QWebEngineUrlRequestJob__Error = 0

//
const QWebEngineUrlRequestJob__UrlNotFound QWebEngineUrlRequestJob__Error = 1

//
const QWebEngineUrlRequestJob__UrlInvalid QWebEngineUrlRequestJob__Error = 2

//
const QWebEngineUrlRequestJob__RequestAborted QWebEngineUrlRequestJob__Error = 3

//
const QWebEngineUrlRequestJob__RequestDenied QWebEngineUrlRequestJob__Error = 4

//
const QWebEngineUrlRequestJob__RequestFailed QWebEngineUrlRequestJob__Error = 5

func (this *QWebEngineUrlRequestJob) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWebEngineUrlRequestJob_ErrorItemName(val int) string {
	var nilthis *QWebEngineUrlRequestJob
	return nilthis.ErrorItemName(val)
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
	if false {
		qtqml.KeepMe()
	}
	if false {
		qtpositioning.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
}

//  keep block end
