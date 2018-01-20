//  header block begin
// /usr/include/qt/QtGui/qdesktopservices.h
// #include <qdesktopservices.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QDesktopServices struct {
	*qtrt.CObject
}

func (this *QDesktopServices) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQDesktopServicesFromPointer(cthis unsafe.Pointer) *QDesktopServices {
	return &QDesktopServices{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qdesktopservices.h:59
// index:0
// Public static
// bool openUrl(const class QUrl &)
func (this *QDesktopServices) OpenUrl(url *qtcore.QUrl) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDesktopServices7openUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, url)
	gopp.ErrPrint(err, rv)
	return rv
}
func QDesktopServices_OpenUrl(url *qtcore.QUrl) {
	var nilthis *QDesktopServices
	nilthis.OpenUrl(url)
}

// /usr/include/qt/QtGui/qdesktopservices.h:60
// index:0
// Public static
// void setUrlHandler(const class QString &, class QObject *, const char *)
func (this *QDesktopServices) SetUrlHandler(scheme *qtcore.QString, receiver unsafe.Pointer, method string) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDesktopServices13setUrlHandlerERK7QStringP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, scheme, receiver, method)
	gopp.ErrPrint(err, rv)
}
func QDesktopServices_SetUrlHandler(scheme *qtcore.QString, receiver unsafe.Pointer, method string) {
	var nilthis *QDesktopServices
	nilthis.SetUrlHandler(scheme, receiver, method)
}

// /usr/include/qt/QtGui/qdesktopservices.h:61
// index:0
// Public static
// void unsetUrlHandler(const class QString &)
func (this *QDesktopServices) UnsetUrlHandler(scheme *qtcore.QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDesktopServices15unsetUrlHandlerERK7QString", ffiqt.FFI_TYPE_POINTER, scheme)
	gopp.ErrPrint(err, rv)
}
func QDesktopServices_UnsetUrlHandler(scheme *qtcore.QString) {
	var nilthis *QDesktopServices
	nilthis.UnsetUrlHandler(scheme)
}

//  body block end
