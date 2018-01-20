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

// /usr/include/qt/QtWidgets/qtooltip.h:56
// index:0
// static
// void showText(const class QPoint &, const class QString &, class QWidget *)
func (this *QToolTip) ShowText(pos unsafe.Pointer, text unsafe.Pointer, w unsafe.Pointer) {
	// 0: (pos const QPoint &, text const QString &, w QWidget *), (pos, text, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QToolTip_ShowText(pos unsafe.Pointer, text unsafe.Pointer, w unsafe.Pointer) {
	// 0: (pos const QPoint &, text const QString &, w QWidget *), (pos, text, w)
	var nilthis *QToolTip
	nilthis.ShowText(pos, text, w)
}

// /usr/include/qt/QtWidgets/qtooltip.h:57
// index:1
// static
// void showText(const class QPoint &, const class QString &, class QWidget *, const class QRect &)
func (this *QToolTip) ShowText_1(pos unsafe.Pointer, text unsafe.Pointer, w unsafe.Pointer, rect unsafe.Pointer) {
	// 1: (pos const QPoint &, text const QString &, w QWidget *, rect const QRect &), (pos, text, w, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRect", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QToolTip_ShowText_1(pos unsafe.Pointer, text unsafe.Pointer, w unsafe.Pointer, rect unsafe.Pointer) {
	// 1: (pos const QPoint &, text const QString &, w QWidget *, rect const QRect &), (pos, text, w, rect)
	var nilthis *QToolTip
	nilthis.ShowText_1(pos, text, w, rect)
}

// /usr/include/qt/QtWidgets/qtooltip.h:58
// index:2
// static
// void showText(const class QPoint &, const class QString &, class QWidget *, const class QRect &, int)
func (this *QToolTip) ShowText_2(pos unsafe.Pointer, text unsafe.Pointer, w unsafe.Pointer, rect unsafe.Pointer, msecShowTime int) {
	// 2: (pos const QPoint &, text const QString &, w QWidget *, rect const QRect &, msecShowTime int), (pos, text, w, rect, msecShowTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRecti", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QToolTip_ShowText_2(pos unsafe.Pointer, text unsafe.Pointer, w unsafe.Pointer, rect unsafe.Pointer, msecShowTime int) {
	// 2: (pos const QPoint &, text const QString &, w QWidget *, rect const QRect &, msecShowTime int), (pos, text, w, rect, msecShowTime)
	var nilthis *QToolTip
	nilthis.ShowText_2(pos, text, w, rect, msecShowTime)
}

// /usr/include/qt/QtWidgets/qtooltip.h:59
// index:0
// static inline
// void hideText()
func (this *QToolTip) HideText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip8hideTextEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QToolTip_HideText() {
	// 0: (), ()
	var nilthis *QToolTip
	nilthis.HideText()
}

// /usr/include/qt/QtWidgets/qtooltip.h:61
// index:0
// static
// bool isVisible()
func (this *QToolTip) IsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip9isVisibleEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QToolTip_IsVisible() {
	// 0: (), ()
	var nilthis *QToolTip
	nilthis.IsVisible()
}

// /usr/include/qt/QtWidgets/qtooltip.h:62
// index:0
// static
// QString text()
func (this *QToolTip) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip4textEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QToolTip_Text() {
	// 0: (), ()
	var nilthis *QToolTip
	nilthis.Text()
}

// /usr/include/qt/QtWidgets/qtooltip.h:64
// index:0
// static
// QPalette palette()
func (this *QToolTip) Palette() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip7paletteEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QToolTip_Palette() {
	// 0: (), ()
	var nilthis *QToolTip
	nilthis.Palette()
}

// /usr/include/qt/QtWidgets/qtooltip.h:65
// index:0
// static
// void setPalette(const class QPalette &)
func (this *QToolTip) SetPalette(arg0 unsafe.Pointer) {
	// 0: (const QPalette & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip10setPaletteERK8QPalette", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QToolTip_SetPalette(arg0 unsafe.Pointer) {
	// 0: (const QPalette & arg0), (arg0)
	var nilthis *QToolTip
	nilthis.SetPalette(arg0)
}

// /usr/include/qt/QtWidgets/qtooltip.h:66
// index:0
// static
// QFont font()
func (this *QToolTip) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip4fontEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QToolTip_Font() {
	// 0: (), ()
	var nilthis *QToolTip
	nilthis.Font()
}

// /usr/include/qt/QtWidgets/qtooltip.h:67
// index:0
// static
// void setFont(const class QFont &)
func (this *QToolTip) SetFont(arg0 unsafe.Pointer) {
	// 0: (const QFont & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolTip7setFontERK5QFont", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QToolTip_SetFont(arg0 unsafe.Pointer) {
	// 0: (const QFont & arg0), (arg0)
	var nilthis *QToolTip
	nilthis.SetFont(arg0)
}

//  body block end
