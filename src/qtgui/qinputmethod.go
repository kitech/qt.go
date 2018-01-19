//  header block begin
// /usr/include/qt/QtGui/qinputmethod.h
// #include <qinputmethod.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QInputMethod struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qinputmethod.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QInputMethod) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:68
// index:0
// QTransform inputItemTransform()
func (this *QInputMethod) InputItemTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod18inputItemTransformEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:69
// index:0
// void setInputItemTransform(const class QTransform &)
func (this *QInputMethod) SetInputItemTransform(transform unsafe.Pointer) {
	// 0: (, const QTransform & transform), (transform)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod21setInputItemTransformERK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, transform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:71
// index:0
// QRectF inputItemRectangle()
func (this *QInputMethod) InputItemRectangle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod18inputItemRectangleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:72
// index:0
// void setInputItemRectangle(const class QRectF &)
func (this *QInputMethod) SetInputItemRectangle(rect unsafe.Pointer) {
	// 0: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod21setInputItemRectangleERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:75
// index:0
// QRectF cursorRectangle()
func (this *QInputMethod) CursorRectangle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod15cursorRectangleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:76
// index:0
// QRectF anchorRectangle()
func (this *QInputMethod) AnchorRectangle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod15anchorRectangleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:79
// index:0
// QRectF keyboardRectangle()
func (this *QInputMethod) KeyboardRectangle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod17keyboardRectangleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:81
// index:0
// QRectF inputItemClipRectangle()
func (this *QInputMethod) InputItemClipRectangle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod22inputItemClipRectangleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:89
// index:0
// bool isVisible()
func (this *QInputMethod) IsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod9isVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:90
// index:0
// void setVisible(_Bool)
func (this *QInputMethod) SetVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:92
// index:0
// bool isAnimating()
func (this *QInputMethod) IsAnimating() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod11isAnimatingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:94
// index:0
// QLocale locale()
func (this *QInputMethod) Locale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod6localeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:95
// index:0
// Qt::LayoutDirection inputDirection()
func (this *QInputMethod) InputDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod14inputDirectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:97
// index:0
// static
// QVariant queryFocusObject(Qt::InputMethodQuery, class QVariant)
func (this *QInputMethod) QueryFocusObject(query int, argument unsafe.Pointer) {
	// 0: (Qt::InputMethodQuery query, QVariant argument), (query, argument)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod16queryFocusObjectEN2Qt16InputMethodQueryE8QVariant", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QInputMethod_QueryFocusObject(query int, argument unsafe.Pointer) {
	// 0: (Qt::InputMethodQuery query, QVariant argument), (query, argument)
	var nilthis *QInputMethod
	nilthis.QueryFocusObject(query, argument)
}

// /usr/include/qt/QtGui/qinputmethod.h:100
// index:0
// void show()
func (this *QInputMethod) Show() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod4showEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:101
// index:0
// void hide()
func (this *QInputMethod) Hide() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod4hideEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:104
// index:0
// void reset()
func (this *QInputMethod) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod5resetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:105
// index:0
// void commit()
func (this *QInputMethod) Commit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod6commitEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:107
// index:0
// void invokeAction(enum QInputMethod::Action, int)
func (this *QInputMethod) InvokeAction(a int, cursorPosition int) {
	// 0: (, QInputMethod::Action a, int cursorPosition), (&a, &cursorPosition)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod12invokeActionENS_6ActionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &a, &cursorPosition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:110
// index:0
// void cursorRectangleChanged()
func (this *QInputMethod) CursorRectangleChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod22cursorRectangleChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:111
// index:0
// void anchorRectangleChanged()
func (this *QInputMethod) AnchorRectangleChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod22anchorRectangleChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:112
// index:0
// void keyboardRectangleChanged()
func (this *QInputMethod) KeyboardRectangleChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod24keyboardRectangleChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:113
// index:0
// void inputItemClipRectangleChanged()
func (this *QInputMethod) InputItemClipRectangleChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod29inputItemClipRectangleChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:114
// index:0
// void visibleChanged()
func (this *QInputMethod) VisibleChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod14visibleChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:115
// index:0
// void animatingChanged()
func (this *QInputMethod) AnimatingChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod16animatingChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:116
// index:0
// void localeChanged()
func (this *QInputMethod) LocaleChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod13localeChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:117
// index:0
// void inputDirectionChanged(Qt::LayoutDirection)
func (this *QInputMethod) InputDirectionChanged(newDirection int) {
	// 0: (, Qt::LayoutDirection newDirection), (&newDirection)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod21inputDirectionChangedEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_VOID, this.cthis, &newDirection)
	gopp.ErrPrint(err, rv)
}

//  body block end
