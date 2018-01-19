//  header block begin
// /usr/include/qt/QtWidgets/qstackedlayout.h
// #include <qstackedlayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QStackedLayout struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QStackedLayout) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:66
// index:0
// void QStackedLayout()
func NewQStackedLayout() *QStackedLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QStackedLayout{cthis}
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:67
// index:1
// void QStackedLayout(class QWidget *)
func NewQStackedLayout_1(parent unsafe.Pointer) *QStackedLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QStackedLayout{cthis}
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:68
// index:2
// void QStackedLayout(class QLayout *)
func NewQStackedLayout_2(parentLayout unsafe.Pointer) *QStackedLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutC2EP7QLayout", ffiqt.FFI_TYPE_VOID, cthis, parentLayout)
	gopp.ErrPrint(err, rv)
	return &QStackedLayout{cthis}
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:69
// index:0
// virtual
// void ~QStackedLayout()
func DeleteQStackedLayout(*QStackedLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:71
// index:0
// int addWidget(class QWidget *)
func (this *QStackedLayout) AddWidget(w unsafe.Pointer) {
	// 0: (, QWidget * w), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout9addWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:72
// index:0
// int insertWidget(int, class QWidget *)
func (this *QStackedLayout) InsertWidget(index int, w unsafe.Pointer) {
	// 0: (, int index, QWidget * w), (&index, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout12insertWidgetEiP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &index, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:74
// index:0
// QWidget * currentWidget()
func (this *QStackedLayout) CurrentWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout13currentWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:75
// index:0
// int currentIndex()
func (this *QStackedLayout) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:77
// index:0
// QWidget * widget(int)
func (this *QStackedLayout) Widget(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout6widgetEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:78
// index:0
// virtual
// int count()
func (this *QStackedLayout) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:80
// index:0
// QStackedLayout::StackingMode stackingMode()
func (this *QStackedLayout) StackingMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout12stackingModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:81
// index:0
// void setStackingMode(enum QStackedLayout::StackingMode)
func (this *QStackedLayout) SetStackingMode(stackingMode int) {
	// 0: (, QStackedLayout::StackingMode stackingMode), (&stackingMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout15setStackingModeENS_12StackingModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &stackingMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:84
// index:0
// virtual
// void addItem(class QLayoutItem *)
func (this *QStackedLayout) AddItem(item unsafe.Pointer) {
	// 0: (, QLayoutItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:85
// index:0
// virtual
// QSize sizeHint()
func (this *QStackedLayout) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:86
// index:0
// virtual
// QSize minimumSize()
func (this *QStackedLayout) MinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout11minimumSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:87
// index:0
// virtual
// QLayoutItem * itemAt(int)
func (this *QStackedLayout) ItemAt(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout6itemAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:88
// index:0
// virtual
// QLayoutItem * takeAt(int)
func (this *QStackedLayout) TakeAt(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout6takeAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:89
// index:0
// virtual
// void setGeometry(const class QRect &)
func (this *QStackedLayout) SetGeometry(rect unsafe.Pointer) {
	// 0: (, const QRect & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout11setGeometryERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:90
// index:0
// virtual
// bool hasHeightForWidth()
func (this *QStackedLayout) HasHeightForWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout17hasHeightForWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:91
// index:0
// virtual
// int heightForWidth(int)
func (this *QStackedLayout) HeightForWidth(width int) {
	// 0: (, int width), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:94
// index:0
// void widgetRemoved(int)
func (this *QStackedLayout) WidgetRemoved(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout13widgetRemovedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:95
// index:0
// void currentChanged(int)
func (this *QStackedLayout) CurrentChanged(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout14currentChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:98
// index:0
// void setCurrentIndex(int)
func (this *QStackedLayout) SetCurrentIndex(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout15setCurrentIndexEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:99
// index:0
// void setCurrentWidget(class QWidget *)
func (this *QStackedLayout) SetCurrentWidget(w unsafe.Pointer) {
	// 0: (, QWidget * w), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, w)
	gopp.ErrPrint(err, rv)
}

//  body block end
