package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebenginefullscreenrequest.h
// #include <qwebenginefullscreenrequest.h>
// #include <QtWebEngineWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtprintsupport"
import "github.com/kitech/qt.go/qtquick"
import "github.com/kitech/qt.go/qtwebenginecore"

//  ext block end

//  body block begin

/*

 */
type QWebEngineFullScreenRequest struct {
	*qtrt.CObject
}
type QWebEngineFullScreenRequest_ITF interface {
	QWebEngineFullScreenRequest_PTR() *QWebEngineFullScreenRequest
}

func (ptr *QWebEngineFullScreenRequest) QWebEngineFullScreenRequest_PTR() *QWebEngineFullScreenRequest {
	return ptr
}

func (this *QWebEngineFullScreenRequest) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineFullScreenRequest) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineFullScreenRequestFromPointer(cthis unsafe.Pointer) *QWebEngineFullScreenRequest {
	return &QWebEngineFullScreenRequest{&qtrt.CObject{cthis}}
}
func (*QWebEngineFullScreenRequest) NewFromPointer(cthis unsafe.Pointer) *QWebEngineFullScreenRequest {
	return NewQWebEngineFullScreenRequestFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginefullscreenrequest.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reject()

/*

 */
func (this *QWebEngineFullScreenRequest) Reject() {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QWebEngineFullScreenRequest6rejectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginefullscreenrequest.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accept()

/*

 */
func (this *QWebEngineFullScreenRequest) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QWebEngineFullScreenRequest6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginefullscreenrequest.h:57
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool toggleOn() const

/*

 */
func (this *QWebEngineFullScreenRequest) ToggleOn() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QWebEngineFullScreenRequest8toggleOnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginefullscreenrequest.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QUrl & origin() const

/*

 */
func (this *QWebEngineFullScreenRequest) Origin() *qtcore.QUrl {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QWebEngineFullScreenRequest6originEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

func DeleteQWebEngineFullScreenRequest(this *QWebEngineFullScreenRequest) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QWebEngineFullScreenRequestD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11715() {
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
		qtquickwidgets.KeepMe()
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
		qtwidgets.KeepMe()
	}
	if false {
		qtprintsupport.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
	if false {
		qtwebenginecore.KeepMe()
	}
}

//  keep block end
