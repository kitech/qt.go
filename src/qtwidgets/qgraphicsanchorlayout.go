//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h
// #include <qgraphicsanchorlayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QGraphicsAnchorLayout struct {
	*QGraphicsLayout
}

func (this *QGraphicsAnchorLayout) GetCthis() unsafe.Pointer {
	return this.QGraphicsLayout.GetCthis()
}
func NewQGraphicsAnchorLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsAnchorLayout {
	bcthis0 := NewQGraphicsLayoutFromPointer(cthis)
	return &QGraphicsAnchorLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:79
// index:0
// Public
// void QGraphicsAnchorLayout(class QGraphicsLayoutItem *)
func NewQGraphicsAnchorLayout(parent unsafe.Pointer) *QGraphicsAnchorLayout {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayoutC2EP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsAnchorLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:80
// index:0
// Public virtual
// void ~QGraphicsAnchorLayout()
func DeleteQGraphicsAnchorLayout(*QGraphicsAnchorLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:82
// index:0
// Public
// QGraphicsAnchor * addAnchor(class QGraphicsLayoutItem *, Qt::AnchorPoint, class QGraphicsLayoutItem *, Qt::AnchorPoint)
func (this *QGraphicsAnchorLayout) AddAnchor(firstItem unsafe.Pointer, firstEdge int, secondItem unsafe.Pointer, secondEdge int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout9addAnchorEP19QGraphicsLayoutItemN2Qt11AnchorPointES1_S3_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), firstItem, &firstEdge, secondItem, &secondEdge)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:84
// index:0
// Public
// QGraphicsAnchor * anchor(class QGraphicsLayoutItem *, Qt::AnchorPoint, class QGraphicsLayoutItem *, Qt::AnchorPoint)
func (this *QGraphicsAnchorLayout) Anchor(firstItem unsafe.Pointer, firstEdge int, secondItem unsafe.Pointer, secondEdge int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout6anchorEP19QGraphicsLayoutItemN2Qt11AnchorPointES1_S3_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), firstItem, &firstEdge, secondItem, &secondEdge)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:87
// index:0
// Public
// void addCornerAnchors(class QGraphicsLayoutItem *, Qt::Corner, class QGraphicsLayoutItem *, Qt::Corner)
func (this *QGraphicsAnchorLayout) AddCornerAnchors(firstItem unsafe.Pointer, firstCorner int, secondItem unsafe.Pointer, secondCorner int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout16addCornerAnchorsEP19QGraphicsLayoutItemN2Qt6CornerES1_S3_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), firstItem, &firstCorner, secondItem, &secondCorner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:94
// index:0
// Public
// void setHorizontalSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetHorizontalSpacing(spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout20setHorizontalSpacingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:95
// index:0
// Public
// void setVerticalSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetVerticalSpacing(spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout18setVerticalSpacingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:96
// index:0
// Public
// void setSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetSpacing(spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout10setSpacingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:97
// index:0
// Public
// qreal horizontalSpacing()
func (this *QGraphicsAnchorLayout) HorizontalSpacing() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout17horizontalSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:98
// index:0
// Public
// qreal verticalSpacing()
func (this *QGraphicsAnchorLayout) VerticalSpacing() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout15verticalSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:100
// index:0
// Public virtual
// void removeAt(int)
func (this *QGraphicsAnchorLayout) RemoveAt(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout8removeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:101
// index:0
// Public virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsAnchorLayout) SetGeometry(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout11setGeometryERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:102
// index:0
// Public virtual
// int count()
func (this *QGraphicsAnchorLayout) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:103
// index:0
// Public virtual
// QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsAnchorLayout) ItemAt(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:105
// index:0
// Public virtual
// void invalidate()
func (this *QGraphicsAnchorLayout) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:107
// index:0
// Protected virtual
// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsAnchorLayout) SizeHint(which int, constraint *qtcore.QSizeF) interface{} {
	var convArg1 = constraint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
