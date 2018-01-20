//  header block begin
// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h
// #include <qgraphicslayoutitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QGraphicsLayoutItem struct {
	*qtrt.CObject
}

func (this *QGraphicsLayoutItem) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:57
// index:0
// void QGraphicsLayoutItem(class QGraphicsLayoutItem *, _Bool)
func NewQGraphicsLayoutItem(parent unsafe.Pointer, isLayout bool) *QGraphicsLayoutItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItemC1EPS_b", ffiqt.FFI_TYPE_VOID, cthis, parent, &isLayout)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutItemFromPointer(cthis)
	return gothis
}
func NewQGraphicsLayoutItemFromPointer(cthis unsafe.Pointer) *QGraphicsLayoutItem {
	return &QGraphicsLayoutItem{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:107
// index:1
// void QGraphicsLayoutItem(class QGraphicsLayoutItemPrivate &)
func NewQGraphicsLayoutItem_1(dd unsafe.Pointer) *QGraphicsLayoutItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItemC1ER26QGraphicsLayoutItemPrivate", ffiqt.FFI_TYPE_VOID, cthis, dd)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:58
// index:0
// virtual
// void ~QGraphicsLayoutItem()
func DeleteQGraphicsLayoutItem(*QGraphicsLayoutItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:60
// index:0
// void setSizePolicy(const class QSizePolicy &)
func (this *QGraphicsLayoutItem) SetSizePolicy(policy unsafe.Pointer) {
	// 0: (, policy const QSizePolicy &), (policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy", ffiqt.FFI_TYPE_VOID, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:61
// index:1
// void setSizePolicy(class QSizePolicy::Policy, class QSizePolicy::Policy, class QSizePolicy::ControlType)
func (this *QGraphicsLayoutItem) SetSizePolicy_1(hPolicy int, vPolicy int, controlType int) {
	// 1: (, hPolicy QSizePolicy::Policy, vPolicy QSizePolicy::Policy, controlType QSizePolicy::ControlType), (&hPolicy, &vPolicy, &controlType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem13setSizePolicyEN11QSizePolicy6PolicyES1_NS0_11ControlTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &hPolicy, &vPolicy, &controlType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:62
// index:0
// QSizePolicy sizePolicy()
func (this *QGraphicsLayoutItem) SizePolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem10sizePolicyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:64
// index:0
// void setMinimumSize(const class QSizeF &)
func (this *QGraphicsLayoutItem) SetMinimumSize(size unsafe.Pointer) {
	// 0: (, size const QSizeF &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:65
// index:1
// inline
// void setMinimumSize(qreal, qreal)
func (this *QGraphicsLayoutItem) SetMinimumSize_1(w float64, h float64) {
	// 1: (, w qreal, h qreal), (&w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMinimumSizeEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:66
// index:0
// QSizeF minimumSize()
func (this *QGraphicsLayoutItem) MinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem11minimumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:67
// index:0
// void setMinimumWidth(qreal)
func (this *QGraphicsLayoutItem) SetMinimumWidth(width float64) {
	// 0: (, width qreal), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setMinimumWidthEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:68
// index:0
// inline
// qreal minimumWidth()
func (this *QGraphicsLayoutItem) MinimumWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12minimumWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:69
// index:0
// void setMinimumHeight(qreal)
func (this *QGraphicsLayoutItem) SetMinimumHeight(height float64) {
	// 0: (, height qreal), (&height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setMinimumHeightEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:70
// index:0
// inline
// qreal minimumHeight()
func (this *QGraphicsLayoutItem) MinimumHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13minimumHeightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:72
// index:0
// void setPreferredSize(const class QSizeF &)
func (this *QGraphicsLayoutItem) SetPreferredSize(size unsafe.Pointer) {
	// 0: (, size const QSizeF &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setPreferredSizeERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:73
// index:1
// inline
// void setPreferredSize(qreal, qreal)
func (this *QGraphicsLayoutItem) SetPreferredSize_1(w float64, h float64) {
	// 1: (, w qreal, h qreal), (&w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setPreferredSizeEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:74
// index:0
// QSizeF preferredSize()
func (this *QGraphicsLayoutItem) PreferredSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13preferredSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:75
// index:0
// void setPreferredWidth(qreal)
func (this *QGraphicsLayoutItem) SetPreferredWidth(width float64) {
	// 0: (, width qreal), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem17setPreferredWidthEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:76
// index:0
// inline
// qreal preferredWidth()
func (this *QGraphicsLayoutItem) PreferredWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem14preferredWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:77
// index:0
// void setPreferredHeight(qreal)
func (this *QGraphicsLayoutItem) SetPreferredHeight(height float64) {
	// 0: (, height qreal), (&height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem18setPreferredHeightEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:78
// index:0
// inline
// qreal preferredHeight()
func (this *QGraphicsLayoutItem) PreferredHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem15preferredHeightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:80
// index:0
// void setMaximumSize(const class QSizeF &)
func (this *QGraphicsLayoutItem) SetMaximumSize(size unsafe.Pointer) {
	// 0: (, size const QSizeF &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMaximumSizeERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:81
// index:1
// inline
// void setMaximumSize(qreal, qreal)
func (this *QGraphicsLayoutItem) SetMaximumSize_1(w float64, h float64) {
	// 1: (, w qreal, h qreal), (&w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMaximumSizeEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:82
// index:0
// QSizeF maximumSize()
func (this *QGraphicsLayoutItem) MaximumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem11maximumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:83
// index:0
// void setMaximumWidth(qreal)
func (this *QGraphicsLayoutItem) SetMaximumWidth(width float64) {
	// 0: (, width qreal), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setMaximumWidthEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:84
// index:0
// inline
// qreal maximumWidth()
func (this *QGraphicsLayoutItem) MaximumWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12maximumWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:85
// index:0
// void setMaximumHeight(qreal)
func (this *QGraphicsLayoutItem) SetMaximumHeight(height float64) {
	// 0: (, height qreal), (&height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setMaximumHeightEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:86
// index:0
// inline
// qreal maximumHeight()
func (this *QGraphicsLayoutItem) MaximumHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13maximumHeightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:88
// index:0
// virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsLayoutItem) SetGeometry(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem11setGeometryERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:89
// index:0
// QRectF geometry()
func (this *QGraphicsLayoutItem) Geometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8geometryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:90
// index:0
// virtual
// void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsLayoutItem) GetContentsMargins(left unsafe.Pointer, top unsafe.Pointer, right unsafe.Pointer, bottom unsafe.Pointer) {
	// 0: (, left qreal *, top qreal *, right qreal *, bottom qreal *), (left, top, right, bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem18getContentsMarginsEPdS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:91
// index:0
// QRectF contentsRect()
func (this *QGraphicsLayoutItem) ContentsRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12contentsRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:93
// index:0
// QSizeF effectiveSizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsLayoutItem) EffectiveSizeHint(which int, constraint unsafe.Pointer) {
	// 0: (, which Qt::SizeHint, constraint const QSizeF &), (&which, constraint)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem17effectiveSizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &which, constraint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:95
// index:0
// virtual
// void updateGeometry()
func (this *QGraphicsLayoutItem) UpdateGeometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14updateGeometryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:97
// index:0
// QGraphicsLayoutItem * parentLayoutItem()
func (this *QGraphicsLayoutItem) ParentLayoutItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem16parentLayoutItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:98
// index:0
// void setParentLayoutItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLayoutItem) SetParentLayoutItem(parent unsafe.Pointer) {
	// 0: (, parent QGraphicsLayoutItem *), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem19setParentLayoutItemEPS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:100
// index:0
// bool isLayout()
func (this *QGraphicsLayoutItem) IsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8isLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:101
// index:0
// QGraphicsItem * graphicsItem()
func (this *QGraphicsLayoutItem) GraphicsItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12graphicsItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:102
// index:0
// bool ownedByLayout()
func (this *QGraphicsLayoutItem) OwnedByLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13ownedByLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:105
// index:0
// void setGraphicsItem(class QGraphicsItem *)
func (this *QGraphicsLayoutItem) SetGraphicsItem(item unsafe.Pointer) {
	// 0: (, item QGraphicsItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setGraphicsItemEP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:106
// index:0
// void setOwnedByLayout(_Bool)
func (this *QGraphicsLayoutItem) SetOwnedByLayout(ownedByLayout bool) {
	// 0: (, ownedByLayout bool), (&ownedByLayout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setOwnedByLayoutEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ownedByLayout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:109
// index:0
// pure virtual
// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsLayoutItem) SizeHint(which int, constraint unsafe.Pointer) {
	// 0: (, which Qt::SizeHint, constraint const QSizeF &), (&which, constraint)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8sizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &which, constraint)
	gopp.ErrPrint(err, rv)
}

//  body block end
