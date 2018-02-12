package qtquickcontrols2

// /usr/include/qt/QtQuickControls2/qquickstyle.h
// #include <qquickstyle.h>
// #include <QtQuickControls2>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtquicktemplates2"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtquick"

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
func (this *QQuickStyle) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyle4nameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QQuickStyle_Name() string {
	var nilthis *QQuickStyle
	rv := nilthis.Name()
	return rv
}

// /usr/include/qt/QtQuickControls2/qquickstyle.h:50
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString path()
func (this *QQuickStyle) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyle4pathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QQuickStyle_Path() string {
	var nilthis *QQuickStyle
	rv := nilthis.Path()
	return rv
}

// /usr/include/qt/QtQuickControls2/qquickstyle.h:51
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setStyle(const QString &)
func (this *QQuickStyle) SetStyle(style string) {
	var tmpArg0 = qtcore.NewQString_5(style)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyle8setStyleERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QQuickStyle_SetStyle(style string) {
	var nilthis *QQuickStyle
	nilthis.SetStyle(style)
}

// /usr/include/qt/QtQuickControls2/qquickstyle.h:52
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setFallbackStyle(const QString &)
func (this *QQuickStyle) SetFallbackStyle(style string) {
	var tmpArg0 = qtcore.NewQString_5(style)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyle16setFallbackStyleERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QQuickStyle_SetFallbackStyle(style string) {
	var nilthis *QQuickStyle
	nilthis.SetFallbackStyle(style)
}

// /usr/include/qt/QtQuickControls2/qquickstyle.h:53
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList availableStyles()
func (this *QQuickStyle) AvailableStyles() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyle15availableStylesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QQuickStyle_AvailableStyles() *qtcore.QStringList /*123*/ {
	var nilthis *QQuickStyle
	rv := nilthis.AvailableStyles()
	return rv
}

func DeleteQQuickStyle(this *QQuickStyle) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
		qtrt.KeepMe()
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

//  keep block end
