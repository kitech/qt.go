package qtwebenginecore

// /usr/include/qt/QtWebEngineCore/qwebenginequotarequest.h
// #include <qwebenginequotarequest.h>
// #include <QtWebEngineCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QWebEngineQuotaRequest struct {
	*qtrt.CObject
}
type QWebEngineQuotaRequest_ITF interface {
	QWebEngineQuotaRequest_PTR() *QWebEngineQuotaRequest
}

func (ptr *QWebEngineQuotaRequest) QWebEngineQuotaRequest_PTR() *QWebEngineQuotaRequest { return ptr }

func (this *QWebEngineQuotaRequest) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineQuotaRequest) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineQuotaRequestFromPointer(cthis unsafe.Pointer) *QWebEngineQuotaRequest {
	return &QWebEngineQuotaRequest{&qtrt.CObject{cthis}}
}
func (*QWebEngineQuotaRequest) NewFromPointer(cthis unsafe.Pointer) *QWebEngineQuotaRequest {
	return NewQWebEngineQuotaRequestFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineCore/qwebenginequotarequest.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QWebEngineQuotaRequest()

/*

 */
func (*QWebEngineQuotaRequest) NewForInherit() *QWebEngineQuotaRequest {
	return NewQWebEngineQuotaRequest()
}
func NewQWebEngineQuotaRequest() *QWebEngineQuotaRequest {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineQuotaRequestC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineQuotaRequestFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWebEngineQuotaRequest)
	return gothis
}

// /usr/include/qt/QtWebEngineCore/qwebenginequotarequest.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accept()

/*

 */
func (this *QWebEngineQuotaRequest) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineQuotaRequest6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginequotarequest.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reject()

/*

 */
func (this *QWebEngineQuotaRequest) Reject() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineQuotaRequest6rejectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginequotarequest.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl origin() const

/*

 */
func (this *QWebEngineQuotaRequest) Origin() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineQuotaRequest6originEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebenginequotarequest.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 requestedSize() const

/*

 */
func (this *QWebEngineQuotaRequest) RequestedSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineQuotaRequest13requestedSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtWebEngineCore/qwebenginequotarequest.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QWebEngineQuotaRequest &) const

/*

 */
func (this *QWebEngineQuotaRequest) Operator_equal_equal(that QWebEngineQuotaRequest_ITF) bool {
	var convArg0 unsafe.Pointer
	if that != nil && that.QWebEngineQuotaRequest_PTR() != nil {
		convArg0 = that.QWebEngineQuotaRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineQuotaRequesteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineCore/qwebenginequotarequest.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QWebEngineQuotaRequest &) const

/*

 */
func (this *QWebEngineQuotaRequest) Operator_not_equal(that QWebEngineQuotaRequest_ITF) bool {
	var convArg0 unsafe.Pointer
	if that != nil && that.QWebEngineQuotaRequest_PTR() != nil {
		convArg0 = that.QWebEngineQuotaRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWebEngineQuotaRequestneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQWebEngineQuotaRequest(this *QWebEngineQuotaRequest) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWebEngineQuotaRequestD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11673() {
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
