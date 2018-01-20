//  header block begin
// /usr/include/qt/QtWidgets/qgraphicslayout.h
// #include <qgraphicslayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 41
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
type QGraphicsLayout struct {
	*QGraphicsLayoutItem
}

func (this *QGraphicsLayout) GetCthis() unsafe.Pointer {
	return this.QGraphicsLayoutItem.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:57
// index:0
// void QGraphicsLayout(class QGraphicsLayoutItem *)
func NewQGraphicsLayout(parent unsafe.Pointer) *QGraphicsLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayoutC1EP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutFromPointer(cthis)
	return gothis
}
func NewQGraphicsLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsLayout {
	bcthis0 := NewQGraphicsLayoutItemFromPointer(cthis)
	return &QGraphicsLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:77
// index:1
// void QGraphicsLayout(class QGraphicsLayoutPrivate &, class QGraphicsLayoutItem *)
func NewQGraphicsLayout_1(arg0 unsafe.Pointer, arg1 unsafe.Pointer) *QGraphicsLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayoutC1ER22QGraphicsLayoutPrivateP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, arg0, arg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:58
// index:0
// virtual
// void ~QGraphicsLayout()
func DeleteQGraphicsLayout(*QGraphicsLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:60
// index:0
// void setContentsMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsLayout) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	// 0: (, left qreal, top qreal, right qreal, bottom qreal), (&left, &top, &right, &bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout18setContentsMarginsEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:61
// index:0
// virtual
// void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsLayout) GetContentsMargins(left unsafe.Pointer, top unsafe.Pointer, right unsafe.Pointer, bottom unsafe.Pointer) {
	// 0: (, left qreal *, top qreal *, right qreal *, bottom qreal *), (left, top, right, bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:63
// index:0
// void activate()
func (this *QGraphicsLayout) Activate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout8activateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:64
// index:0
// bool isActivated()
func (this *QGraphicsLayout) IsActivated() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout11isActivatedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:65
// index:0
// virtual
// void invalidate()
func (this *QGraphicsLayout) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout10invalidateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:66
// index:0
// virtual
// void updateGeometry()
func (this *QGraphicsLayout) UpdateGeometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout14updateGeometryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:68
// index:0
// virtual
// void widgetEvent(class QEvent *)
func (this *QGraphicsLayout) WidgetEvent(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout11widgetEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:70
// index:0
// pure virtual
// int count()
func (this *QGraphicsLayout) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:71
// index:0
// pure virtual
// QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsLayout) ItemAt(i int) {
	// 0: (, i int), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout6itemAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:72
// index:0
// pure virtual
// void removeAt(int)
func (this *QGraphicsLayout) RemoveAt(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout8removeAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:74
// index:0
// static
// void setInstantInvalidatePropagation(_Bool)
func (this *QGraphicsLayout) SetInstantInvalidatePropagation(enable bool) {
	// 0: (enable bool), (enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout31setInstantInvalidatePropagationEb", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGraphicsLayout_SetInstantInvalidatePropagation(enable bool) {
	// 0: (enable bool), (enable)
	var nilthis *QGraphicsLayout
	nilthis.SetInstantInvalidatePropagation(enable)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:75
// index:0
// static
// bool instantInvalidatePropagation()
func (this *QGraphicsLayout) InstantInvalidatePropagation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout28instantInvalidatePropagationEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGraphicsLayout_InstantInvalidatePropagation() {
	// 0: (), ()
	var nilthis *QGraphicsLayout
	nilthis.InstantInvalidatePropagation()
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:78
// index:0
// void addChildLayoutItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLayout) AddChildLayoutItem(layoutItem unsafe.Pointer) {
	// 0: (, layoutItem QGraphicsLayoutItem *), (layoutItem)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout18addChildLayoutItemEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), layoutItem)
	gopp.ErrPrint(err, rv)
}

//  body block end
