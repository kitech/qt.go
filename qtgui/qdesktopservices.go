package qtgui

// /usr/include/qt/QtGui/qdesktopservices.h
// #include <qdesktopservices.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDesktopServices) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDesktopServicesFromPointer(cthis unsafe.Pointer) *QDesktopServices {
	return &QDesktopServices{&qtrt.CObject{cthis}}
}
func (*QDesktopServices) NewFromPointer(cthis unsafe.Pointer) *QDesktopServices {
	return NewQDesktopServicesFromPointer(cthis)
}

// /usr/include/qt/QtGui/qdesktopservices.h:59
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool openUrl(const QUrl &)
func (this *QDesktopServices) OpenUrl(url *qtcore.QUrl) bool {
	var convArg0 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDesktopServices7openUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QDesktopServices_OpenUrl(url *qtcore.QUrl) bool {
	var nilthis *QDesktopServices
	rv := nilthis.OpenUrl(url)
	return rv
}

// /usr/include/qt/QtGui/qdesktopservices.h:60
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setUrlHandler(const QString &, QObject *, const char *)
func (this *QDesktopServices) SetUrlHandler(scheme *qtcore.QString, receiver *qtcore.QObject /*777 QObject **/, method string) {
	var convArg0 = scheme.GetCthis()
	var convArg1 = receiver.GetCthis()
	var convArg2 = qtrt.CString(method)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDesktopServices13setUrlHandlerERK7QStringP7QObjectPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}
func QDesktopServices_SetUrlHandler(scheme *qtcore.QString, receiver *qtcore.QObject /*777 QObject **/, method string) {
	var nilthis *QDesktopServices
	nilthis.SetUrlHandler(scheme, receiver, method)
}

// /usr/include/qt/QtGui/qdesktopservices.h:61
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void unsetUrlHandler(const QString &)
func (this *QDesktopServices) UnsetUrlHandler(scheme *qtcore.QString) {
	var convArg0 = scheme.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDesktopServices15unsetUrlHandlerERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QDesktopServices_UnsetUrlHandler(scheme *qtcore.QString) {
	var nilthis *QDesktopServices
	nilthis.UnsetUrlHandler(scheme)
}

func DeleteQDesktopServices(this *QDesktopServices) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDesktopServicesD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
