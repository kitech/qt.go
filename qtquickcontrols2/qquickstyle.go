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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtquicktemplates2"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtquick"

//  ext block end

//  body block begin

/*

 */
type QQuickStyle struct {
	*qtrt.CObject
}
type QQuickStyle_ITF interface {
	QQuickStyle_PTR() *QQuickStyle
}

func (ptr *QQuickStyle) QQuickStyle_PTR() *QQuickStyle { return ptr }

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

/*
Returns the name of the application style.

Note: The application style can be specified by passing a -style command line argument. Therefore name() may not return a fully resolved value if called before constructing a QGuiApplication.
*/
func (this *QQuickStyle) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyle4nameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
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

/*
Returns the path of an overridden application style, or an empty string if the style is one of the built-in Qt Quick Controls 2 styles.

Note: The application style can be specified by passing a -style command line argument. Therefore path() may not return a fully resolved value if called before constructing a QGuiApplication.
*/
func (this *QQuickStyle) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQuickStyle4pathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
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

/*
Sets the application style to style.

Note: The style must be configured before loading QML that imports Qt Quick Controls 2. It is not possible to change the style after the QML types have been registered.
*/
func (this *QQuickStyle) SetStyle(style string) {
	var tmpArg0 = qtcore.NewQString5(style)
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

/*
Sets the application fallback style to style.

Note: The fallback style must be the name of one of the built-in Qt Quick Controls 2 styles, e.g. "Material".

Note: The style must be configured before loading QML that imports Qt Quick Controls 2. It is not possible to change the style after the QML types have been registered.

The fallback style can be also specified by setting the QT_QUICK_CONTROLS_FALLBACK_STYLE environment variable.

This function was introduced in  Qt 5.8.
*/
func (this *QQuickStyle) SetFallbackStyle(style string) {
	var tmpArg0 = qtcore.NewQString5(style)
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

/*
Returns the names of the available built-in styles.

Note: The method must be called after creating an instance of QGuiApplication.

This function was introduced in  Qt 5.9.
*/
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
		log.Println(123)
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
