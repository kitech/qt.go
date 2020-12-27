package qtgui

// /usr/include/qt/QtGui/qdesktopservices.h
// #include <qdesktopservices.h>
// #include <QtGui>

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
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*
 */
// size 1
type QDesktopServices struct {
	*qtrt.CObject
}
type QDesktopServices_ITF interface {
	QDesktopServices_PTR() *QDesktopServices
}

func (ptr *QDesktopServices) QDesktopServices_PTR() *QDesktopServices { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QDesktopServicesFromptr(cthis Voidptr) *QDesktopServices {
	return &QDesktopServices{&qtrt.CObject{cthis}}
}
func (*QDesktopServices) Fromptr(cthis Voidptr) *QDesktopServices {
	return QDesktopServicesFromptr(cthis)
}

// /usr/include/qt/QtGui/qdesktopservices.h:59
// index:0
// Public static Extend Visibility=Default Availability=Available
// [1] bool openUrl(const QUrl &)

/*
 */
func (this *QDesktopServices) OpenUrl(url qtcore.QUrl_ITF) bool {
	var convArg0 Voidptr
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2434727476, "_ZN16QDesktopServices7openUrlERK4QUrl", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}
func QDesktopServices_OpenUrl(url qtcore.QUrl_ITF) bool {
	var nilthis *QDesktopServices
	rv := nilthis.OpenUrl(url)
	return rv
}

// /usr/include/qt/QtGui/qdesktopservices.h:60
// index:0
// Public static Ignore Visibility=Default Availability=Available
// [-2] void setUrlHandler(const QString &, QObject *, const char *)

/*
 */
func (this *QDesktopServices) SetUrlHandler(scheme string, receiver qtcore.QObject_ITF /*777 QObject **/, method string) {
	var tmpArg0 = qtcore.NewQString5(scheme)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 Voidptr
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg1 = receiver.QObject_PTR().GetCthis()
	}
	var convArg2 = qtrt.CStringRef(&method)
	rv, err := qtrt.Qtcc3(254877298, "_ZN16QDesktopServices13setUrlHandlerERK7QStringP7QObjectPKc", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
}
func QDesktopServices_SetUrlHandler(scheme string, receiver qtcore.QObject_ITF /*777 QObject **/, method string) {
	var nilthis *QDesktopServices
	nilthis.SetUrlHandler(scheme, receiver, method)
}

// /usr/include/qt/QtGui/qdesktopservices.h:61
// index:0
// Public static Ignore Visibility=Default Availability=Available
// [-2] void unsetUrlHandler(const QString &)

/*
 */
func (this *QDesktopServices) UnsetUrlHandler(scheme string) {
	var tmpArg0 = qtcore.NewQString5(scheme)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(1369796868, "_ZN16QDesktopServices15unsetUrlHandlerERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
}
func QDesktopServices_UnsetUrlHandler(scheme string) {
	var nilthis *QDesktopServices
	nilthis.UnsetUrlHandler(scheme)
}

func DeleteQDesktopServices(this *QDesktopServices) {
	rv, err := qtrt.Qtcc3(1858443276, "_ZN16QDesktopServicesD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10163() {
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
		qtcore.KeepMe()
	}
}

//  keep block end
