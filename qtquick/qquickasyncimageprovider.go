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
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

type QQuickAsyncImageProvider struct {
	*QQuickImageProvider
}
type QQuickAsyncImageProvider_ITF interface {
	QQuickImageProvider_ITF
	QQuickAsyncImageProvider_PTR() *QQuickAsyncImageProvider
}

func (ptr *QQuickAsyncImageProvider) QQuickAsyncImageProvider_PTR() *QQuickAsyncImageProvider {
	return ptr
}

func (this *QQuickAsyncImageProvider) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QQuickImageProvider.GetCthis()
	}
}
func (this *QQuickAsyncImageProvider) SetCthis(cthis unsafe.Pointer) {
	this.QQuickImageProvider = NewQQuickImageProviderFromPointer(cthis)
}
func NewQQuickAsyncImageProviderFromPointer(cthis unsafe.Pointer) *QQuickAsyncImageProvider {
	bcthis0 := NewQQuickImageProviderFromPointer(cthis)
	return &QQuickAsyncImageProvider{bcthis0}
}
func (*QQuickAsyncImageProvider) NewFromPointer(cthis unsafe.Pointer) *QQuickAsyncImageProvider {
	return NewQQuickAsyncImageProviderFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickAsyncImageProvider()
func NewQQuickAsyncImageProvider() *QQuickAsyncImageProvider {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QQuickAsyncImageProviderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickAsyncImageProviderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuickAsyncImageProvider)
	return gothis
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:116
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickAsyncImageProvider()
func DeleteQQuickAsyncImageProvider(this *QQuickAsyncImageProvider) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QQuickAsyncImageProviderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:121
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QQuickImageResponse * requestImageResponse(const QString &, const QSize &)
func (this *QQuickAsyncImageProvider) RequestImageResponse(id string, requestedSize *qtcore.QSize) *QQuickImageResponse /*777 QQuickImageResponse **/ {
	var tmpArg0 = qtcore.NewQString_5(id)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = requestedSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QQuickAsyncImageProvider20requestImageResponseERK7QStringRK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickImageResponseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
