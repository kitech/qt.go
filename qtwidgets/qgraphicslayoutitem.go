package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h
// #include <qgraphicslayoutitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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

// void setGraphicsItem(class QGraphicsItem *)
func (this *QGraphicsLayoutItem) InheritSetGraphicsItem(f func(item *QGraphicsItem /*777 QGraphicsItem **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setGraphicsItem", f)
}

// void setOwnedByLayout(_Bool)
func (this *QGraphicsLayoutItem) InheritSetOwnedByLayout(f func(ownedByLayout bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setOwnedByLayout", f)
}

// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsLayoutItem) InheritSizeHint(f func(which int, constraint *qtcore.QSizeF) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sizeHint", f)
}

type QGraphicsLayoutItem struct {
	*qtrt.CObject
}
type QGraphicsLayoutItem_ITF interface {
	QGraphicsLayoutItem_PTR() *QGraphicsLayoutItem
}

func (ptr *QGraphicsLayoutItem) QGraphicsLayoutItem_PTR() *QGraphicsLayoutItem { return ptr }

func (this *QGraphicsLayoutItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGraphicsLayoutItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGraphicsLayoutItemFromPointer(cthis unsafe.Pointer) *QGraphicsLayoutItem {
	return &QGraphicsLayoutItem{&qtrt.CObject{cthis}}
}
func (*QGraphicsLayoutItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsLayoutItem {
	return NewQGraphicsLayoutItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLayoutItem(QGraphicsLayoutItem *, _Bool)
func NewQGraphicsLayoutItem(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, isLayout bool) *QGraphicsLayoutItem {
	var convArg0 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItemC2EPS_b", qtrt.FFI_TYPE_POINTER, convArg0, isLayout)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLayoutItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsLayoutItem()
func DeleteQGraphicsLayoutItem(this *QGraphicsLayoutItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizePolicy(const QSizePolicy &)
func (this *QGraphicsLayoutItem) SetSizePolicy(policy QSizePolicy_ITF) {
	var convArg0 = policy.QSizePolicy_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSizePolicy(QSizePolicy::Policy, QSizePolicy::Policy, QSizePolicy::ControlType)
func (this *QGraphicsLayoutItem) SetSizePolicy_1(hPolicy int, vPolicy int, controlType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem13setSizePolicyEN11QSizePolicy6PolicyES1_NS0_11ControlTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hPolicy, vPolicy, controlType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:62
// index:0
// Public Visibility=Default Availability=Available
// [4] QSizePolicy sizePolicy()
func (this *QGraphicsLayoutItem) SizePolicy() *QSizePolicy /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem10sizePolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizePolicy)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumSize(const QSizeF &)
func (this *QGraphicsLayoutItem) SetMinimumSize(size qtcore.QSizeF_ITF) {
	var convArg0 = size.QSizeF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:65
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setMinimumSize(qreal, qreal)
func (this *QGraphicsLayoutItem) SetMinimumSize_1(w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMinimumSizeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:66
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF minimumSize()
func (this *QGraphicsLayoutItem) MinimumSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumWidth(qreal)
func (this *QGraphicsLayoutItem) SetMinimumWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setMinimumWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal minimumWidth()
func (this *QGraphicsLayoutItem) MinimumWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12minimumWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumHeight(qreal)
func (this *QGraphicsLayoutItem) SetMinimumHeight(height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setMinimumHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal minimumHeight()
func (this *QGraphicsLayoutItem) MinimumHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13minimumHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreferredSize(const QSizeF &)
func (this *QGraphicsLayoutItem) SetPreferredSize(size qtcore.QSizeF_ITF) {
	var convArg0 = size.QSizeF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setPreferredSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:73
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setPreferredSize(qreal, qreal)
func (this *QGraphicsLayoutItem) SetPreferredSize_1(w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setPreferredSizeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:74
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF preferredSize()
func (this *QGraphicsLayoutItem) PreferredSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13preferredSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreferredWidth(qreal)
func (this *QGraphicsLayoutItem) SetPreferredWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem17setPreferredWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal preferredWidth()
func (this *QGraphicsLayoutItem) PreferredWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem14preferredWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreferredHeight(qreal)
func (this *QGraphicsLayoutItem) SetPreferredHeight(height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem18setPreferredHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal preferredHeight()
func (this *QGraphicsLayoutItem) PreferredHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem15preferredHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumSize(const QSizeF &)
func (this *QGraphicsLayoutItem) SetMaximumSize(size qtcore.QSizeF_ITF) {
	var convArg0 = size.QSizeF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMaximumSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:81
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setMaximumSize(qreal, qreal)
func (this *QGraphicsLayoutItem) SetMaximumSize_1(w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMaximumSizeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:82
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF maximumSize()
func (this *QGraphicsLayoutItem) MaximumSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumWidth(qreal)
func (this *QGraphicsLayoutItem) SetMaximumWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setMaximumWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal maximumWidth()
func (this *QGraphicsLayoutItem) MaximumWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12maximumWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumHeight(qreal)
func (this *QGraphicsLayoutItem) SetMaximumHeight(height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setMaximumHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal maximumHeight()
func (this *QGraphicsLayoutItem) MaximumHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13maximumHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRectF &)
func (this *QGraphicsLayoutItem) SetGeometry(rect qtcore.QRectF_ITF) {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem11setGeometryERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:89
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF geometry()
func (this *QGraphicsLayoutItem) Geometry() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsLayoutItem) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem18getContentsMarginsEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:91
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF contentsRect()
func (this *QGraphicsLayoutItem) ContentsRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12contentsRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:93
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF effectiveSizeHint(Qt::SizeHint, const QSizeF &)
func (this *QGraphicsLayoutItem) EffectiveSizeHint(which int, constraint qtcore.QSizeF_ITF) *qtcore.QSizeF /*123*/ {
	var convArg1 = constraint.QSizeF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem17effectiveSizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateGeometry()
func (this *QGraphicsLayoutItem) UpdateGeometry() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14updateGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * parentLayoutItem()
func (this *QGraphicsLayoutItem) ParentLayoutItem() *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem16parentLayoutItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParentLayoutItem(QGraphicsLayoutItem *)
func (this *QGraphicsLayoutItem) SetParentLayoutItem(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) {
	var convArg0 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem19setParentLayoutItemEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLayout()
func (this *QGraphicsLayoutItem) IsLayout() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8isLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * graphicsItem()
func (this *QGraphicsLayoutItem) GraphicsItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12graphicsItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ownedByLayout()
func (this *QGraphicsLayoutItem) OwnedByLayout() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13ownedByLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:105
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setGraphicsItem(QGraphicsItem *)
func (this *QGraphicsLayoutItem) SetGraphicsItem(item QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 = item.QGraphicsItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setGraphicsItemEP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setOwnedByLayout(_Bool)
func (this *QGraphicsLayoutItem) SetOwnedByLayout(ownedByLayout bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setOwnedByLayoutEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ownedByLayout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:109
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &)
func (this *QGraphicsLayoutItem) SizeHint(which int, constraint qtcore.QSizeF_ITF) *qtcore.QSizeF /*123*/ {
	var convArg1 = constraint.QSizeF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
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
