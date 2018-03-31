package qtwebenginecore

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinfo.h
// #include <qwebengineurlrequestinfo.h>
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
type QWebEngineUrlRequestInfo struct {
	*qtrt.CObject
}
type QWebEngineUrlRequestInfo_ITF interface {
	QWebEngineUrlRequestInfo_PTR() *QWebEngineUrlRequestInfo
}

func (ptr *QWebEngineUrlRequestInfo) QWebEngineUrlRequestInfo_PTR() *QWebEngineUrlRequestInfo {
	return ptr
}

func (this *QWebEngineUrlRequestInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineUrlRequestInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineUrlRequestInfoFromPointer(cthis unsafe.Pointer) *QWebEngineUrlRequestInfo {
	return &QWebEngineUrlRequestInfo{&qtrt.CObject{cthis}}
}
func (*QWebEngineUrlRequestInfo) NewFromPointer(cthis unsafe.Pointer) *QWebEngineUrlRequestInfo {
	return NewQWebEngineUrlRequestInfoFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinfo.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineUrlRequestInfo::ResourceType resourceType() const

/*

 */
func (this *QWebEngineUrlRequestInfo) ResourceType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QWebEngineUrlRequestInfo12resourceTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinfo.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineUrlRequestInfo::NavigationType navigationType() const

/*

 */
func (this *QWebEngineUrlRequestInfo) NavigationType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QWebEngineUrlRequestInfo14navigationTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinfo.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl requestUrl() const

/*

 */
func (this *QWebEngineUrlRequestInfo) RequestUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QWebEngineUrlRequestInfo10requestUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinfo.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl firstPartyUrl() const

/*

 */
func (this *QWebEngineUrlRequestInfo) FirstPartyUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QWebEngineUrlRequestInfo13firstPartyUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinfo.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray requestMethod() const

/*

 */
func (this *QWebEngineUrlRequestInfo) RequestMethod() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QWebEngineUrlRequestInfo13requestMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinfo.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool changed() const

/*

 */
func (this *QWebEngineUrlRequestInfo) Changed() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QWebEngineUrlRequestInfo7changedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinfo.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void block(bool)

/*

 */
func (this *QWebEngineUrlRequestInfo) Block(shouldBlock bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QWebEngineUrlRequestInfo5blockEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), shouldBlock)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinfo.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redirect(const QUrl &)

/*

 */
func (this *QWebEngineUrlRequestInfo) Redirect(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QWebEngineUrlRequestInfo8redirectERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlrequestinfo.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpHeader(const QByteArray &, const QByteArray &)

/*

 */
func (this *QWebEngineUrlRequestInfo) SetHttpHeader(name qtcore.QByteArray_ITF, value qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if name != nil && name.QByteArray_PTR() != nil {
		convArg0 = name.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QByteArray_PTR() != nil {
		convArg1 = value.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QWebEngineUrlRequestInfo13setHttpHeaderERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

func DeleteQWebEngineUrlRequestInfo(this *QWebEngineUrlRequestInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QWebEngineUrlRequestInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QWebEngineUrlRequestInfo__ResourceType = int

//
const QWebEngineUrlRequestInfo__ResourceTypeMainFrame QWebEngineUrlRequestInfo__ResourceType = 0

//
const QWebEngineUrlRequestInfo__ResourceTypeSubFrame QWebEngineUrlRequestInfo__ResourceType = 1

//
const QWebEngineUrlRequestInfo__ResourceTypeStylesheet QWebEngineUrlRequestInfo__ResourceType = 2

//
const QWebEngineUrlRequestInfo__ResourceTypeScript QWebEngineUrlRequestInfo__ResourceType = 3

//
const QWebEngineUrlRequestInfo__ResourceTypeImage QWebEngineUrlRequestInfo__ResourceType = 4

//
const QWebEngineUrlRequestInfo__ResourceTypeFontResource QWebEngineUrlRequestInfo__ResourceType = 5

//
const QWebEngineUrlRequestInfo__ResourceTypeSubResource QWebEngineUrlRequestInfo__ResourceType = 6

//
const QWebEngineUrlRequestInfo__ResourceTypeObject QWebEngineUrlRequestInfo__ResourceType = 7

//
const QWebEngineUrlRequestInfo__ResourceTypeMedia QWebEngineUrlRequestInfo__ResourceType = 8

//
const QWebEngineUrlRequestInfo__ResourceTypeWorker QWebEngineUrlRequestInfo__ResourceType = 9

//
const QWebEngineUrlRequestInfo__ResourceTypeSharedWorker QWebEngineUrlRequestInfo__ResourceType = 10

//
const QWebEngineUrlRequestInfo__ResourceTypePrefetch QWebEngineUrlRequestInfo__ResourceType = 11

//
const QWebEngineUrlRequestInfo__ResourceTypeFavicon QWebEngineUrlRequestInfo__ResourceType = 12

//
const QWebEngineUrlRequestInfo__ResourceTypeXhr QWebEngineUrlRequestInfo__ResourceType = 13

//
const QWebEngineUrlRequestInfo__ResourceTypePing QWebEngineUrlRequestInfo__ResourceType = 14

//
const QWebEngineUrlRequestInfo__ResourceTypeServiceWorker QWebEngineUrlRequestInfo__ResourceType = 15

//
const QWebEngineUrlRequestInfo__ResourceTypeCspReport QWebEngineUrlRequestInfo__ResourceType = 16

//
const QWebEngineUrlRequestInfo__ResourceTypePluginResource QWebEngineUrlRequestInfo__ResourceType = 17

//
const QWebEngineUrlRequestInfo__ResourceTypeLast QWebEngineUrlRequestInfo__ResourceType = 18

//
const QWebEngineUrlRequestInfo__ResourceTypeUnknown QWebEngineUrlRequestInfo__ResourceType = 255

/*


 */
type QWebEngineUrlRequestInfo__NavigationType = int

//
const QWebEngineUrlRequestInfo__NavigationTypeLink QWebEngineUrlRequestInfo__NavigationType = 0

//
const QWebEngineUrlRequestInfo__NavigationTypeTyped QWebEngineUrlRequestInfo__NavigationType = 1

//
const QWebEngineUrlRequestInfo__NavigationTypeFormSubmitted QWebEngineUrlRequestInfo__NavigationType = 2

//
const QWebEngineUrlRequestInfo__NavigationTypeBackForward QWebEngineUrlRequestInfo__NavigationType = 3

//
const QWebEngineUrlRequestInfo__NavigationTypeReload QWebEngineUrlRequestInfo__NavigationType = 4

//
const QWebEngineUrlRequestInfo__NavigationTypeOther QWebEngineUrlRequestInfo__NavigationType = 5

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
