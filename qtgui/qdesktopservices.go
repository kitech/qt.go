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
// extern C begin: 20
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
type QDesktopServices struct {
	*qtrt.CObject
}
type QDesktopServices_ITF interface {
	QDesktopServices_PTR() *QDesktopServices
}

func (ptr *QDesktopServices) QDesktopServices_PTR() *QDesktopServices { return ptr }

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

/*
Opens the given url in the appropriate Web browser for the user's desktop environment, and returns true if successful; otherwise returns false.

If the URL is a reference to a local file (i.e., the URL scheme is "file") then it will be opened with a suitable application instead of a Web browser.

The following example opens a file on the Windows file system residing on a path that contains spaces:


  QDesktopServices::openUrl(QUrl("file:///C:/Documents and Settings/All Users/Desktop", QUrl::TolerantMode));



If a mailto URL is specified, the user's e-mail client will be used to open a composer window containing the options specified in the URL, similar to the way mailto links are handled by a Web browser.

For example, the following URL contains a recipient (user@foo.com), a subject (Test), and a message body (Just a test):


  mailto:user@foo.com?subject=Test&body=Just a test



Warning: Although many e-mail clients can send attachments and are Unicode-aware, the user may have configured their client without these features. Also, certain e-mail clients (e.g., Lotus Notes) have problems with long URLs.

Warning: A return value of true indicates that the application has successfully requested the operating system to open the URL in an external application. The external application may still fail to launch or fail to open the requested URL. This result will not be reported back to the application.

Warning: URLs passed to this function on iOS will not load unless their schemes are listed in the LSApplicationQueriesSchemes key of the application's Info.plist file. For more information, see the Apple Developer Documentation for canOpenURL(_:). For example, the following lines enable URLs with the HTTPS scheme:


  <key>LSApplicationQueriesSchemes</key>
  <array>
      <string>https</string>
  </array>



See also setUrlHandler().
*/
func (this *QDesktopServices) OpenUrl(url qtcore.QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDesktopServices7openUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDesktopServices_OpenUrl(url qtcore.QUrl_ITF) bool {
	var nilthis *QDesktopServices
	rv := nilthis.OpenUrl(url)
	return rv
}

// /usr/include/qt/QtGui/qdesktopservices.h:60
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setUrlHandler(const QString &, QObject *, const char *)

/*
Sets the handler for the given scheme to be the handler method provided by the receiver object.

This function provides a way to customize the behavior of openUrl(). If openUrl() is called with a URL with the specified scheme then the given method on the receiver object is called instead of QDesktopServices launching an external application.

The provided method must be implemented as a slot that only accepts a single QUrl argument.

If setUrlHandler() is used to set a new handler for a scheme which already has a handler, the existing handler is simply replaced with the new one. Since QDesktopServices does not take ownership of handlers, no objects are deleted when a handler is replaced.

Note that the handler will always be called from within the same thread that calls QDesktopServices::openUrl().

See also openUrl() and unsetUrlHandler().
*/
func (this *QDesktopServices) SetUrlHandler(scheme string, receiver qtcore.QObject_ITF /*777 QObject **/, method string) {
	var tmpArg0 = qtcore.NewQString_5(scheme)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg1 = receiver.QObject_PTR().GetCthis()
	}
	var convArg2 = qtrt.CString(method)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDesktopServices13setUrlHandlerERK7QStringP7QObjectPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}
func QDesktopServices_SetUrlHandler(scheme string, receiver qtcore.QObject_ITF /*777 QObject **/, method string) {
	var nilthis *QDesktopServices
	nilthis.SetUrlHandler(scheme, receiver, method)
}

// /usr/include/qt/QtGui/qdesktopservices.h:61
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void unsetUrlHandler(const QString &)

/*
Removes a previously set URL handler for the specified scheme.

See also setUrlHandler().
*/
func (this *QDesktopServices) UnsetUrlHandler(scheme string) {
	var tmpArg0 = qtcore.NewQString_5(scheme)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDesktopServices15unsetUrlHandlerERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QDesktopServices_UnsetUrlHandler(scheme string) {
	var nilthis *QDesktopServices
	nilthis.UnsetUrlHandler(scheme)
}

func DeleteQDesktopServices(this *QDesktopServices) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDesktopServicesD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
		qtcore.KeepMe()
	}
}

//  keep block end
