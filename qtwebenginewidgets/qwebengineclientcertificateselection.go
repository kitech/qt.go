package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebengineclientcertificateselection.h
// #include <qwebengineclientcertificateselection.h>
// #include <QtWebEngineWidgets>

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
type QWebEngineClientCertificateSelection struct {
	*qtrt.CObject
}
type QWebEngineClientCertificateSelection_ITF interface {
	QWebEngineClientCertificateSelection_PTR() *QWebEngineClientCertificateSelection
}

func (ptr *QWebEngineClientCertificateSelection) QWebEngineClientCertificateSelection_PTR() *QWebEngineClientCertificateSelection {
	return ptr
}

func (this *QWebEngineClientCertificateSelection) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineClientCertificateSelection) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineClientCertificateSelectionFromPointer(cthis unsafe.Pointer) *QWebEngineClientCertificateSelection {
	return &QWebEngineClientCertificateSelection{&qtrt.CObject{cthis}}
}
func (*QWebEngineClientCertificateSelection) NewFromPointer(cthis unsafe.Pointer) *QWebEngineClientCertificateSelection {
	return NewQWebEngineClientCertificateSelectionFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineclientcertificateselection.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QWebEngineClientCertificateSelection()

/*

 */
func DeleteQWebEngineClientCertificateSelection(this *QWebEngineClientCertificateSelection) {
	rv, err := qtrt.InvokeQtFunc6("_ZN36QWebEngineClientCertificateSelectionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineclientcertificateselection.h:60
// index:0
// Public Visibility=Default Availability=Available
// [16] QWebEngineClientCertificateSelection & operator=(const QWebEngineClientCertificateSelection &)

/*

 */
func (this *QWebEngineClientCertificateSelection) Operator_equal(arg0 QWebEngineClientCertificateSelection_ITF) *QWebEngineClientCertificateSelection {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWebEngineClientCertificateSelection_PTR() != nil {
		convArg0 = arg0.QWebEngineClientCertificateSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN36QWebEngineClientCertificateSelectionaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineClientCertificateSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineClientCertificateSelection)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineclientcertificateselection.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl host() const

/*

 */
func (this *QWebEngineClientCertificateSelection) Host() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK36QWebEngineClientCertificateSelection4hostEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineclientcertificateselection.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void select(const QSslCertificate &)

/*

 */
func (this *QWebEngineClientCertificateSelection) Select(certificate qtnetwork.QSslCertificate_ITF) {
	var convArg0 unsafe.Pointer
	if certificate != nil && certificate.QSslCertificate_PTR() != nil {
		convArg0 = certificate.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN36QWebEngineClientCertificateSelection6selectERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineclientcertificateselection.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectNone()

/*

 */
func (this *QWebEngineClientCertificateSelection) SelectNone() {
	rv, err := qtrt.InvokeQtFunc6("_ZN36QWebEngineClientCertificateSelection10selectNoneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
