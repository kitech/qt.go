package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicslayout.h
// #include <qgraphicslayout.h>
// #include <QtWidgets>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
// void addChildLayoutItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLayout) InheritAddChildLayoutItem(f func(layoutItem *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/)) {
	ffiqt.SetAllInheritCallback(this, "addChildLayoutItem", f)
}

type QGraphicsLayout struct {
	*QGraphicsLayoutItem
}

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
func NewQGraphicsLayout(parent *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) *QGraphicsLayout {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayoutC1EP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsLayout()
func DeleteQGraphicsLayout(this *QGraphicsLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayoutD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsLayout) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout18setContentsMarginsEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsLayout) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activate()
func (this *QGraphicsLayout) Activate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout8activateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActivated()
func (this *QGraphicsLayout) IsActivated() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout11isActivatedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QGraphicsLayout) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateGeometry()
func (this *QGraphicsLayout) UpdateGeometry() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout14updateGeometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void widgetEvent(QEvent *)
func (this *QGraphicsLayout) WidgetEvent(e *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout11widgetEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:70
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int count()
func (this *QGraphicsLayout) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsLayout) ItemAt(i int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:72
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void removeAt(int)
func (this *QGraphicsLayout) RemoveAt(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout8removeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:74
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setInstantInvalidatePropagation(_Bool)
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
// Public static Visibility=Default Availability=Available
// [1] bool instantInvalidatePropagation()
func (this *QGraphicsLayout) InstantInvalidatePropagation() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout28instantInvalidatePropagationEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
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
func (this *QGraphicsLayout) AddChildLayoutItem(layoutItem *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) {
	var convArg0 = layoutItem.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsLayout18addChildLayoutItemEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
