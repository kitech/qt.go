package qtquick

// /usr/include/qt/QtQuick/qquickimageprovider.h
// #include <qquickimageprovider.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

/*

 */
type QQuickImageResponse struct {
	*qtcore.QObject
}
type QQuickImageResponse_ITF interface {
	qtcore.QObject_ITF
	QQuickImageResponse_PTR() *QQuickImageResponse
}

func (ptr *QQuickImageResponse) QQuickImageResponse_PTR() *QQuickImageResponse { return ptr }

func (this *QQuickImageResponse) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickImageResponse) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQuickImageResponseFromPointer(cthis unsafe.Pointer) *QQuickImageResponse {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQuickImageResponse{bcthis0}
}
func (*QQuickImageResponse) NewFromPointer(cthis unsafe.Pointer) *QQuickImageResponse {
	return NewQQuickImageResponseFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQuickImageResponse) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QQuickImageResponse10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickImageResponse()

/*

 */
func (*QQuickImageResponse) NewForInherit() *QQuickImageResponse {
	return NewQQuickImageResponse()
}
func NewQQuickImageResponse() *QQuickImageResponse {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageResponseC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickImageResponseFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickImageResponse")
	return gothis
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickImageResponse()

/*

 */
func DeleteQQuickImageResponse(this *QQuickImageResponse) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageResponseD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QQuickTextureFactory * textureFactory() const

/*

 */
func (this *QQuickImageResponse) TextureFactory() *QQuickTextureFactory /*777 QQuickTextureFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QQuickImageResponse14textureFactoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickTextureFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString errorString() const

/*

 */
func (this *QQuickImageResponse) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QQuickImageResponse11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void cancel()

/*

 */
func (this *QQuickImageResponse) Cancel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageResponse6cancelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()

/*

 */
func (this *QQuickImageResponse) Finished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageResponse8finishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11553() {
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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
