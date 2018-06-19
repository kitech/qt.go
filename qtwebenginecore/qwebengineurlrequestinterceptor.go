package qtwebenginecore

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinterceptor.h
// #include <qwebengineurlrequestinterceptor.h>
// #include <QtWebEngineCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QWebEngineUrlRequestInterceptor struct {
	*qtcore.QObject
}
type QWebEngineUrlRequestInterceptor_ITF interface {
	qtcore.QObject_ITF
	QWebEngineUrlRequestInterceptor_PTR() *QWebEngineUrlRequestInterceptor
}

func (ptr *QWebEngineUrlRequestInterceptor) QWebEngineUrlRequestInterceptor_PTR() *QWebEngineUrlRequestInterceptor {
	return ptr
}

func (this *QWebEngineUrlRequestInterceptor) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWebEngineUrlRequestInterceptor) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWebEngineUrlRequestInterceptorFromPointer(cthis unsafe.Pointer) *QWebEngineUrlRequestInterceptor {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWebEngineUrlRequestInterceptor{bcthis0}
}
func (*QWebEngineUrlRequestInterceptor) NewFromPointer(cthis unsafe.Pointer) *QWebEngineUrlRequestInterceptor {
	return NewQWebEngineUrlRequestInterceptorFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinterceptor.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWebEngineUrlRequestInterceptor) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QWebEngineUrlRequestInterceptor10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinterceptor.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QWebEngineUrlRequestInterceptor(QObject *)

/*

 */
func NewQWebEngineUrlRequestInterceptor(p qtcore.QObject_ITF /*777 QObject **/) *QWebEngineUrlRequestInterceptor {
	var convArg0 unsafe.Pointer
	if p != nil && p.QObject_PTR() != nil {
		convArg0 = p.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN31QWebEngineUrlRequestInterceptorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineUrlRequestInterceptorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEngineUrlRequestInterceptor")
	return gothis
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinterceptor.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QWebEngineUrlRequestInterceptor(QObject *)

/*

 */
func NewQWebEngineUrlRequestInterceptor__() *QWebEngineUrlRequestInterceptor {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN31QWebEngineUrlRequestInterceptorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineUrlRequestInterceptorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEngineUrlRequestInterceptor")
	return gothis
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinterceptor.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void interceptRequest(QWebEngineUrlRequestInfo &)

/*

 */
func (this *QWebEngineUrlRequestInterceptor) InterceptRequest(info QWebEngineUrlRequestInfo_ITF) {
	var convArg0 unsafe.Pointer
	if info != nil && info.QWebEngineUrlRequestInfo_PTR() != nil {
		convArg0 = info.QWebEngineUrlRequestInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN31QWebEngineUrlRequestInterceptor16interceptRequestER24QWebEngineUrlRequestInfo", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQWebEngineUrlRequestInterceptor(this *QWebEngineUrlRequestInterceptor) {
	rv, err := qtrt.InvokeQtFunc6("_ZN31QWebEngineUrlRequestInterceptorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
