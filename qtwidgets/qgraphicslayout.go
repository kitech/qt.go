package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicslayout.h
// #include <qgraphicslayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void addChildLayoutItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLayout) InheritAddChildLayoutItem(f func(layoutItem *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "addChildLayoutItem", f)
}

type QGraphicsLayout struct {
	*QGraphicsLayoutItem
}
type QGraphicsLayout_ITF interface {
	QGraphicsLayoutItem_ITF
	QGraphicsLayout_PTR() *QGraphicsLayout
}

func (ptr *QGraphicsLayout) QGraphicsLayout_PTR() *QGraphicsLayout { return ptr }

func (this *QGraphicsLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsLayoutItem.GetCthis()
	}
}
func (this *QGraphicsLayout) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsLayoutItem = NewQGraphicsLayoutItemFromPointer(cthis)
}
func NewQGraphicsLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsLayout {
	bcthis0 := NewQGraphicsLayoutItemFromPointer(cthis)
	return &QGraphicsLayout{bcthis0}
}
func (*QGraphicsLayout) NewFromPointer(cthis unsafe.Pointer) *QGraphicsLayout {
	return NewQGraphicsLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLayout(QGraphicsLayoutItem *)
func NewQGraphicsLayout(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) *QGraphicsLayout {
	var convArg0 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsLayout()
func DeleteQGraphicsLayout(this *QGraphicsLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsLayout) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout18setContentsMarginsEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsLayout) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activate()
func (this *QGraphicsLayout) Activate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout8activateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActivated()
func (this *QGraphicsLayout) IsActivated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsLayout11isActivatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QGraphicsLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateGeometry()
func (this *QGraphicsLayout) UpdateGeometry() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout14updateGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void widgetEvent(QEvent *)
func (this *QGraphicsLayout) WidgetEvent(e qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 = e.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout11widgetEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:70
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int count()
func (this *QGraphicsLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsLayout) ItemAt(i int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:72
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void removeAt(int)
func (this *QGraphicsLayout) RemoveAt(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout8removeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:74
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setInstantInvalidatePropagation(_Bool)
func (this *QGraphicsLayout) SetInstantInvalidatePropagation(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout31setInstantInvalidatePropagationEb", qtrt.FFI_TYPE_POINTER, enable)
	qtrt.ErrPrint(err, rv)
}
func QGraphicsLayout_SetInstantInvalidatePropagation(enable bool) {
	var nilthis *QGraphicsLayout
	nilthis.SetInstantInvalidatePropagation(enable)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:75
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool instantInvalidatePropagation()
func (this *QGraphicsLayout) InstantInvalidatePropagation() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout28instantInvalidatePropagationEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QGraphicsLayout_InstantInvalidatePropagation() bool {
	var nilthis *QGraphicsLayout
	rv := nilthis.InstantInvalidatePropagation()
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:78
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addChildLayoutItem(QGraphicsLayoutItem *)
func (this *QGraphicsLayout) AddChildLayoutItem(layoutItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) {
	var convArg0 = layoutItem.QGraphicsLayoutItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout18addChildLayoutItemEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
