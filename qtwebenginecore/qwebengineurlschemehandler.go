package qtwebenginecore

// /usr/include/qt/QtWebEngineCore/qwebengineurlschemehandler.h
// #include <qwebengineurlschemehandler.h>
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
type QWebEngineUrlSchemeHandler struct {
	*qtcore.QObject
}
type QWebEngineUrlSchemeHandler_ITF interface {
	qtcore.QObject_ITF
	QWebEngineUrlSchemeHandler_PTR() *QWebEngineUrlSchemeHandler
}

func (ptr *QWebEngineUrlSchemeHandler) QWebEngineUrlSchemeHandler_PTR() *QWebEngineUrlSchemeHandler {
	return ptr
}

func (this *QWebEngineUrlSchemeHandler) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWebEngineUrlSchemeHandler) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWebEngineUrlSchemeHandlerFromPointer(cthis unsafe.Pointer) *QWebEngineUrlSchemeHandler {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWebEngineUrlSchemeHandler{bcthis0}
}
func (*QWebEngineUrlSchemeHandler) NewFromPointer(cthis unsafe.Pointer) *QWebEngineUrlSchemeHandler {
	return NewQWebEngineUrlSchemeHandlerFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlschemehandler.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWebEngineUrlSchemeHandler) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QWebEngineUrlSchemeHandler10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlschemehandler.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineUrlSchemeHandler(QObject *)

/*

 */
func (*QWebEngineUrlSchemeHandler) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QWebEngineUrlSchemeHandler {
	return NewQWebEngineUrlSchemeHandler(parent)
}
func NewQWebEngineUrlSchemeHandler(parent qtcore.QObject_ITF /*777 QObject **/) *QWebEngineUrlSchemeHandler {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWebEngineUrlSchemeHandlerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineUrlSchemeHandlerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEngineUrlSchemeHandler")
	return gothis
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlschemehandler.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineUrlSchemeHandler(QObject *)

/*

 */
func (*QWebEngineUrlSchemeHandler) NewForInheritp() *QWebEngineUrlSchemeHandler {
	return NewQWebEngineUrlSchemeHandlerp()
}
func NewQWebEngineUrlSchemeHandlerp() *QWebEngineUrlSchemeHandler {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWebEngineUrlSchemeHandlerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineUrlSchemeHandlerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEngineUrlSchemeHandler")
	return gothis
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlschemehandler.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWebEngineUrlSchemeHandler()

/*

 */
func DeleteQWebEngineUrlSchemeHandler(this *QWebEngineUrlSchemeHandler) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWebEngineUrlSchemeHandlerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlschemehandler.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void requestStarted(QWebEngineUrlRequestJob *)

/*

 */
func (this *QWebEngineUrlSchemeHandler) RequestStarted(arg0 QWebEngineUrlRequestJob_ITF /*777 QWebEngineUrlRequestJob **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWebEngineUrlRequestJob_PTR() != nil {
		convArg0 = arg0.QWebEngineUrlRequestJob_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWebEngineUrlSchemeHandler14requestStartedEP23QWebEngineUrlRequestJob", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlschemehandler.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void _q_destroyedUrlSchemeHandler(QWebEngineUrlSchemeHandler *)

/*

 */
func (this *QWebEngineUrlSchemeHandler) _q_destroyedUrlSchemeHandler(arg0 QWebEngineUrlSchemeHandler_ITF /*777 QWebEngineUrlSchemeHandler **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWebEngineUrlSchemeHandler_PTR() != nil {
		convArg0 = arg0.QWebEngineUrlSchemeHandler_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWebEngineUrlSchemeHandler28_q_destroyedUrlSchemeHandlerEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
