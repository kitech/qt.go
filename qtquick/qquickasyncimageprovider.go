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
import "gopp"
import "qt.go/cffiqt"
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
type QQuickAsyncImageProvider struct {
	*QQuickImageProvider
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
// Public
// void QQuickAsyncImageProvider()
func NewQQuickAsyncImageProvider() *QQuickAsyncImageProvider {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QQuickAsyncImageProviderC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickAsyncImageProviderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:116
// index:0
// Public virtual
// void ~QQuickAsyncImageProvider()
func DeleteQQuickAsyncImageProvider(*QQuickAsyncImageProvider) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QQuickAsyncImageProviderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:121
// index:0
// Public pure virtual
// QQuickImageResponse * requestImageResponse(const QString &, const QSize &)
func (this *QQuickAsyncImageProvider) RequestImageResponse(id *qtcore.QString, requestedSize *qtcore.QSize) *QQuickImageResponse /*777 QQuickImageResponse **/ {
	var convArg0 = id.GetCthis()
	var convArg1 = requestedSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QQuickAsyncImageProvider20requestImageResponseERK7QStringRK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickImageResponseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
