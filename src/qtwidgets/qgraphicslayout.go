//  header block begin
// /usr/include/qt/QtWidgets/qgraphicslayout.h
// #include <qgraphicslayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
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
func NewQGraphicsLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsLayout {
	bcthis0 := NewQGraphicsLayoutItemFromPointer(cthis)
	return &QGraphicsLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:57
// index:0
// Public
// void QGraphicsLayout(class QGraphicsLayoutItem *)
func NewQGraphicsLayout(parent unsafe.Pointer) *QGraphicsLayout {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayoutC1EP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:58
// index:0
// Public virtual
// void ~QGraphicsLayout()
func DeleteQGraphicsLayout(*QGraphicsLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:60
// index:0
// Public
// void setContentsMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsLayout) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout18setContentsMarginsEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:61
// index:0
// Public virtual
// void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsLayout) GetContentsMargins(left unsafe.Pointer, top unsafe.Pointer, right unsafe.Pointer, bottom unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:63
// index:0
// Public
// void activate()
func (this *QGraphicsLayout) Activate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout8activateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:64
// index:0
// Public
// bool isActivated()
func (this *QGraphicsLayout) IsActivated() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout11isActivatedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:65
// index:0
// Public virtual
// void invalidate()
func (this *QGraphicsLayout) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:66
// index:0
// Public virtual
// void updateGeometry()
func (this *QGraphicsLayout) UpdateGeometry() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout14updateGeometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:68
// index:0
// Public virtual
// void widgetEvent(class QEvent *)
func (this *QGraphicsLayout) WidgetEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout11widgetEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:70
// index:0
// Public pure virtual
// int count()
func (this *QGraphicsLayout) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:71
// index:0
// Public pure virtual
// QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsLayout) ItemAt(i int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:72
// index:0
// Public pure virtual
// void removeAt(int)
func (this *QGraphicsLayout) RemoveAt(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout8removeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:74
// index:0
// Public static
// void setInstantInvalidatePropagation(_Bool)
func (this *QGraphicsLayout) SetInstantInvalidatePropagation(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout31setInstantInvalidatePropagationEb", ffiqt.FFI_TYPE_POINTER, enable)
	gopp.ErrPrint(err, rv)
}
func QGraphicsLayout_SetInstantInvalidatePropagation(enable bool) {
	var nilthis *QGraphicsLayout
	nilthis.SetInstantInvalidatePropagation(enable)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:75
// index:0
// Public static
// bool instantInvalidatePropagation()
func (this *QGraphicsLayout) InstantInvalidatePropagation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout28instantInvalidatePropagationEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGraphicsLayout_InstantInvalidatePropagation() {
	var nilthis *QGraphicsLayout
	nilthis.InstantInvalidatePropagation()
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:78
// index:0
// Protected
// void addChildLayoutItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLayout) AddChildLayoutItem(layoutItem unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout18addChildLayoutItemEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), layoutItem)
	gopp.ErrPrint(err, rv)
}

//  body block end
