package qtquickcontrols2

// /usr/include/qt/QtQuickControls2/qquickstyle.h
// #include <qquickstyle.h>
// #include <QtQuickControls2>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtnetwork"
import "qt.go/qtquicktemplates2"
import "qt.go/qtcore"
import "qt.go/qtqml"
import "qt.go/qtgui"
import "qt.go/qtquick"

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
		qtnetwork.KeepMe()
	}
	if false {
		qtquicktemplates2.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
}

//  ext block end

//  body block begin

type QQuickStyle struct {
	*qtrt.CObject
}

func (this *QQuickStyle) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQuickStyle) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQuickStyleFromPointer(cthis unsafe.Pointer) *QQuickStyle {
	return &QQuickStyle{&qtrt.CObject{cthis}}
}
func (*QQuickStyle) NewFromPointer(cthis unsafe.Pointer) *QQuickStyle {
	return NewQQuickStyleFromPointer(cthis)
}

// /usr/include/qt/QtQuickControls2/qquickstyle.h:49
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString name()
func (this *QQuickStyle) Name() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyle4nameEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}
func QQuickStyle_Name() *qtcore.QString /*123*/ {
	var nilthis *QQuickStyle
	rv := nilthis.Name()
	return rv
}

// /usr/include/qt/QtQuickControls2/qquickstyle.h:50
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString path()
func (this *QQuickStyle) Path() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyle4pathEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}
func QQuickStyle_Path() *qtcore.QString /*123*/ {
	var nilthis *QQuickStyle
	rv := nilthis.Path()
	return rv
}

// /usr/include/qt/QtQuickControls2/qquickstyle.h:51
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setStyle(const QString &)
func (this *QQuickStyle) SetStyle(style *qtcore.QString) {
	var convArg0 = style.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyle8setStyleERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QQuickStyle_SetStyle(style *qtcore.QString) {
	var nilthis *QQuickStyle
	nilthis.SetStyle(style)
}

// /usr/include/qt/QtQuickControls2/qquickstyle.h:52
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setFallbackStyle(const QString &)
func (this *QQuickStyle) SetFallbackStyle(style *qtcore.QString) {
	var convArg0 = style.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyle16setFallbackStyleERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QQuickStyle_SetFallbackStyle(style *qtcore.QString) {
	var nilthis *QQuickStyle
	nilthis.SetFallbackStyle(style)
}

func DeleteQQuickStyle(this *QQuickStyle) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
