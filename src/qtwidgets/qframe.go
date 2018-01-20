//  header block begin
// /usr/include/qt/QtWidgets/qframe.h
// #include <qframe.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QFrame struct {
	*QWidget
}

func (this *QFrame) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qframe.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFrame) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:64
// index:0
// void QFrame(class QWidget *, Qt::WindowFlags)
func NewQFrame(parent unsafe.Pointer, f int) *QFrame {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrameC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, parent, &f)
	gopp.ErrPrint(err, rv)
	gothis := NewQFrameFromPointer(cthis)
	return gothis
}
func NewQFrameFromPointer(cthis unsafe.Pointer) *QFrame {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QFrame{bcthis0}
}

// /usr/include/qt/QtWidgets/qframe.h:118
// index:1
// void QFrame(class QFramePrivate &, class QWidget *, Qt::WindowFlags)
func NewQFrame_1(dd unsafe.Pointer, parent unsafe.Pointer, f int) *QFrame {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrameC2ER13QFramePrivateP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, dd, parent, &f)
	gopp.ErrPrint(err, rv)
	gothis := NewQFrameFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qframe.h:65
// index:0
// virtual
// void ~QFrame()
func DeleteQFrame(*QFrame) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrameD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:67
// index:0
// int frameStyle()
func (this *QFrame) FrameStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame10frameStyleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:68
// index:0
// void setFrameStyle(int)
func (this *QFrame) SetFrameStyle(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame13setFrameStyleEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:70
// index:0
// int frameWidth()
func (this *QFrame) FrameWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame10frameWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:72
// index:0
// virtual
// QSize sizeHint()
func (this *QFrame) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:96
// index:0
// QFrame::Shape frameShape()
func (this *QFrame) FrameShape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame10frameShapeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:97
// index:0
// void setFrameShape(enum QFrame::Shape)
func (this *QFrame) SetFrameShape(arg0 int) {
	// 0: (, QFrame::Shape arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame13setFrameShapeENS_5ShapeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:98
// index:0
// QFrame::Shadow frameShadow()
func (this *QFrame) FrameShadow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame11frameShadowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:99
// index:0
// void setFrameShadow(enum QFrame::Shadow)
func (this *QFrame) SetFrameShadow(arg0 int) {
	// 0: (, QFrame::Shadow arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame14setFrameShadowENS_6ShadowE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:101
// index:0
// int lineWidth()
func (this *QFrame) LineWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame9lineWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:102
// index:0
// void setLineWidth(int)
func (this *QFrame) SetLineWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame12setLineWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:104
// index:0
// int midLineWidth()
func (this *QFrame) MidLineWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame12midLineWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:105
// index:0
// void setMidLineWidth(int)
func (this *QFrame) SetMidLineWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame15setMidLineWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:107
// index:0
// QRect frameRect()
func (this *QFrame) FrameRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame9frameRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:108
// index:0
// void setFrameRect(const class QRect &)
func (this *QFrame) SetFrameRect(arg0 unsafe.Pointer) {
	// 0: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame12setFrameRectERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:111
// index:0
// virtual
// bool event(class QEvent *)
func (this *QFrame) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:112
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QFrame) PaintEvent(arg0 unsafe.Pointer) {
	// 0: (, QPaintEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:113
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QFrame) ChangeEvent(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:114
// index:0
// void drawFrame(class QPainter *)
func (this *QFrame) DrawFrame(arg0 unsafe.Pointer) {
	// 0: (, QPainter * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame9drawFrameEP8QPainter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:119
// index:0
// void initStyleOption(class QStyleOptionFrame *)
func (this *QFrame) InitStyleOption(option unsafe.Pointer) {
	// 0: (, option QStyleOptionFrame *), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame15initStyleOptionEP17QStyleOptionFrame", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
