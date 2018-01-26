package qtquick

// /usr/include/qt/QtQuick/qquickimageprovider.h
// #include <qquickimageprovider.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQuickImageResponse struct {
	*qtcore.QObject
}

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

// /usr/include/qt/QtQuick/qquickimageprovider.h:73
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QQuickImageResponse) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QQuickImageResponse10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:75
// index:0
// Public
// void QQuickImageResponse()
func NewQQuickImageResponse() *QQuickImageResponse {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickImageResponseC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickImageResponseFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:76
// index:0
// Public virtual
// void ~QQuickImageResponse()
func DeleteQQuickImageResponse(*QQuickImageResponse) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickImageResponseD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:78
// index:0
// Public pure virtual
// QQuickTextureFactory * textureFactory()
func (this *QQuickImageResponse) TextureFactory() *QQuickTextureFactory /*777 QQuickTextureFactory **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QQuickImageResponse14textureFactoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickTextureFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:79
// index:0
// Public virtual
// QString errorString()
func (this *QQuickImageResponse) ErrorString() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QQuickImageResponse11errorStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:82
// index:0
// Public virtual
// void cancel()
func (this *QQuickImageResponse) Cancel() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickImageResponse6cancelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:85
// index:0
// Public
// void finished()
func (this *QQuickImageResponse) Finished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickImageResponse8finishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
