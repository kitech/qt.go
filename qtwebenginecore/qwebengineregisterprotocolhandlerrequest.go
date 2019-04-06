package qtwebenginecore

// /usr/include/qt/QtWebEngineCore/qwebengineregisterprotocolhandlerrequest.h
// #include <qwebengineregisterprotocolhandlerrequest.h>
// #include <QtWebEngineCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QWebEngineRegisterProtocolHandlerRequest struct {
	*qtrt.CObject
}
type QWebEngineRegisterProtocolHandlerRequest_ITF interface {
	QWebEngineRegisterProtocolHandlerRequest_PTR() *QWebEngineRegisterProtocolHandlerRequest
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) QWebEngineRegisterProtocolHandlerRequest_PTR() *QWebEngineRegisterProtocolHandlerRequest {
	return ptr
}

func (this *QWebEngineRegisterProtocolHandlerRequest) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineRegisterProtocolHandlerRequest) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineRegisterProtocolHandlerRequestFromPointer(cthis unsafe.Pointer) *QWebEngineRegisterProtocolHandlerRequest {
	return &QWebEngineRegisterProtocolHandlerRequest{&qtrt.CObject{cthis}}
}
func (*QWebEngineRegisterProtocolHandlerRequest) NewFromPointer(cthis unsafe.Pointer) *QWebEngineRegisterProtocolHandlerRequest {
	return NewQWebEngineRegisterProtocolHandlerRequestFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineCore/qwebengineregisterprotocolhandlerrequest.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QWebEngineRegisterProtocolHandlerRequest()

/*

 */
func (*QWebEngineRegisterProtocolHandlerRequest) NewForInherit() *QWebEngineRegisterProtocolHandlerRequest {
	return NewQWebEngineRegisterProtocolHandlerRequest()
}
func NewQWebEngineRegisterProtocolHandlerRequest() *QWebEngineRegisterProtocolHandlerRequest {
	rv, err := qtrt.InvokeQtFunc6("_ZN40QWebEngineRegisterProtocolHandlerRequestC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineRegisterProtocolHandlerRequestFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWebEngineRegisterProtocolHandlerRequest)
	return gothis
}

// /usr/include/qt/QtWebEngineCore/qwebengineregisterprotocolhandlerrequest.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accept()

/*

 */
func (this *QWebEngineRegisterProtocolHandlerRequest) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN40QWebEngineRegisterProtocolHandlerRequest6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineregisterprotocolhandlerrequest.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reject()

/*

 */
func (this *QWebEngineRegisterProtocolHandlerRequest) Reject() {
	rv, err := qtrt.InvokeQtFunc6("_ZN40QWebEngineRegisterProtocolHandlerRequest6rejectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineregisterprotocolhandlerrequest.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl origin() const

/*

 */
func (this *QWebEngineRegisterProtocolHandlerRequest) Origin() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK40QWebEngineRegisterProtocolHandlerRequest6originEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebengineregisterprotocolhandlerrequest.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QString scheme() const

/*

 */
func (this *QWebEngineRegisterProtocolHandlerRequest) Scheme() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK40QWebEngineRegisterProtocolHandlerRequest6schemeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineCore/qwebengineregisterprotocolhandlerrequest.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QWebEngineRegisterProtocolHandlerRequest &) const

/*

 */
func (this *QWebEngineRegisterProtocolHandlerRequest) Operator_equal_equal(that QWebEngineRegisterProtocolHandlerRequest_ITF) bool {
	var convArg0 unsafe.Pointer
	if that != nil && that.QWebEngineRegisterProtocolHandlerRequest_PTR() != nil {
		convArg0 = that.QWebEngineRegisterProtocolHandlerRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK40QWebEngineRegisterProtocolHandlerRequesteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineCore/qwebengineregisterprotocolhandlerrequest.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QWebEngineRegisterProtocolHandlerRequest &) const

/*

 */
func (this *QWebEngineRegisterProtocolHandlerRequest) Operator_not_equal(that QWebEngineRegisterProtocolHandlerRequest_ITF) bool {
	var convArg0 unsafe.Pointer
	if that != nil && that.QWebEngineRegisterProtocolHandlerRequest_PTR() != nil {
		convArg0 = that.QWebEngineRegisterProtocolHandlerRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK40QWebEngineRegisterProtocolHandlerRequestneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQWebEngineRegisterProtocolHandlerRequest(this *QWebEngineRegisterProtocolHandlerRequest) {
	rv, err := qtrt.InvokeQtFunc6("_ZN40QWebEngineRegisterProtocolHandlerRequestD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11675() {
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
