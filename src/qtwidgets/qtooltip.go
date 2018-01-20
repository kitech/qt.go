//  header block begin
// /usr/include/qt/QtWidgets/qtooltip.h
// #include <qtooltip.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
}

//  ext block end

//  body block begin
type QToolTip struct {
	*qtrt.CObject
}

func (this *QToolTip) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQToolTipFromPointer(cthis unsafe.Pointer) *QToolTip {
	return &QToolTip{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qtooltip.h:56
// index:0
// Public static
// void showText(const class QPoint &, const class QString &, class QWidget *)
func (this *QToolTip) ShowText(pos *qtcore.QPoint, text *qtcore.QString, w unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidget", ffiqt.FFI_TYPE_POINTER, pos, text, w)
	gopp.ErrPrint(err, rv)
}
func QToolTip_ShowText(pos *qtcore.QPoint, text *qtcore.QString, w unsafe.Pointer) {
	var nilthis *QToolTip
	nilthis.ShowText(pos, text, w)
}

// /usr/include/qt/QtWidgets/qtooltip.h:57
// index:1
// Public static
// void showText(const class QPoint &, const class QString &, class QWidget *, const class QRect &)
func (this *QToolTip) ShowText_1(pos *qtcore.QPoint, text *qtcore.QString, w unsafe.Pointer, rect *qtcore.QRect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRect", ffiqt.FFI_TYPE_POINTER, pos, text, w, rect)
	gopp.ErrPrint(err, rv)
}
func QToolTip_ShowText_1(pos *qtcore.QPoint, text *qtcore.QString, w unsafe.Pointer, rect *qtcore.QRect) {
	var nilthis *QToolTip
	nilthis.ShowText_1(pos, text, w, rect)
}

// /usr/include/qt/QtWidgets/qtooltip.h:58
// index:2
// Public static
// void showText(const class QPoint &, const class QString &, class QWidget *, const class QRect &, int)
func (this *QToolTip) ShowText_2(pos *qtcore.QPoint, text *qtcore.QString, w unsafe.Pointer, rect *qtcore.QRect, msecShowTime int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRecti", ffiqt.FFI_TYPE_POINTER, pos, text, w, rect, msecShowTime)
	gopp.ErrPrint(err, rv)
}
func QToolTip_ShowText_2(pos *qtcore.QPoint, text *qtcore.QString, w unsafe.Pointer, rect *qtcore.QRect, msecShowTime int) {
	var nilthis *QToolTip
	nilthis.ShowText_2(pos, text, w, rect, msecShowTime)
}

// /usr/include/qt/QtWidgets/qtooltip.h:59
// index:0
// Public static inline
// void hideText()
func (this *QToolTip) HideText() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip8hideTextEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QToolTip_HideText() {
	var nilthis *QToolTip
	nilthis.HideText()
}

// /usr/include/qt/QtWidgets/qtooltip.h:61
// index:0
// Public static
// bool isVisible()
func (this *QToolTip) IsVisible() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip9isVisibleEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QToolTip_IsVisible() {
	var nilthis *QToolTip
	nilthis.IsVisible()
}

// /usr/include/qt/QtWidgets/qtooltip.h:62
// index:0
// Public static
// QString text()
func (this *QToolTip) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip4textEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QToolTip_Text() {
	var nilthis *QToolTip
	nilthis.Text()
}

// /usr/include/qt/QtWidgets/qtooltip.h:64
// index:0
// Public static
// QPalette palette()
func (this *QToolTip) Palette() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip7paletteEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QToolTip_Palette() {
	var nilthis *QToolTip
	nilthis.Palette()
}

// /usr/include/qt/QtWidgets/qtooltip.h:65
// index:0
// Public static
// void setPalette(const class QPalette &)
func (this *QToolTip) SetPalette(arg0 *qtgui.QPalette) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip10setPaletteERK8QPalette", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QToolTip_SetPalette(arg0 *qtgui.QPalette) {
	var nilthis *QToolTip
	nilthis.SetPalette(arg0)
}

// /usr/include/qt/QtWidgets/qtooltip.h:66
// index:0
// Public static
// QFont font()
func (this *QToolTip) Font() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip4fontEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QToolTip_Font() {
	var nilthis *QToolTip
	nilthis.Font()
}

// /usr/include/qt/QtWidgets/qtooltip.h:67
// index:0
// Public static
// void setFont(const class QFont &)
func (this *QToolTip) SetFont(arg0 *qtgui.QFont) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QToolTip_SetFont(arg0 *qtgui.QFont) {
	var nilthis *QToolTip
	nilthis.SetFont(arg0)
}

//  body block end
