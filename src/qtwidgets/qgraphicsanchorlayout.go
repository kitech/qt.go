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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:79
// index:0
// void QGraphicsAnchorLayout(class QGraphicsLayoutItem *)
func NewQGraphicsAnchorLayout(parent unsafe.Pointer) *QGraphicsAnchorLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayoutC2EP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsAnchorLayout{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:80
// index:0
// virtual
// void ~QGraphicsAnchorLayout()
func DeleteQGraphicsAnchorLayout(*QGraphicsAnchorLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:82
// index:0
// QGraphicsAnchor * addAnchor(class QGraphicsLayoutItem *, Qt::AnchorPoint, class QGraphicsLayoutItem *, Qt::AnchorPoint)
func (this *QGraphicsAnchorLayout) AddAnchor(firstItem unsafe.Pointer, firstEdge int, secondItem unsafe.Pointer, secondEdge int) {
	// 0: (, QGraphicsLayoutItem * firstItem, Qt::AnchorPoint firstEdge, QGraphicsLayoutItem * secondItem, Qt::AnchorPoint secondEdge), (firstItem, &firstEdge, secondItem, &secondEdge)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout9addAnchorEP19QGraphicsLayoutItemN2Qt11AnchorPointES1_S3_", ffiqt.FFI_TYPE_VOID, this.cthis, firstItem, &firstEdge, secondItem, &secondEdge)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:84
// index:0
// QGraphicsAnchor * anchor(class QGraphicsLayoutItem *, Qt::AnchorPoint, class QGraphicsLayoutItem *, Qt::AnchorPoint)
func (this *QGraphicsAnchorLayout) Anchor(firstItem unsafe.Pointer, firstEdge int, secondItem unsafe.Pointer, secondEdge int) {
	// 0: (, QGraphicsLayoutItem * firstItem, Qt::AnchorPoint firstEdge, QGraphicsLayoutItem * secondItem, Qt::AnchorPoint secondEdge), (firstItem, &firstEdge, secondItem, &secondEdge)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout6anchorEP19QGraphicsLayoutItemN2Qt11AnchorPointES1_S3_", ffiqt.FFI_TYPE_VOID, this.cthis, firstItem, &firstEdge, secondItem, &secondEdge)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:87
// index:0
// void addCornerAnchors(class QGraphicsLayoutItem *, Qt::Corner, class QGraphicsLayoutItem *, Qt::Corner)
func (this *QGraphicsAnchorLayout) AddCornerAnchors(firstItem unsafe.Pointer, firstCorner int, secondItem unsafe.Pointer, secondCorner int) {
	// 0: (, QGraphicsLayoutItem * firstItem, Qt::Corner firstCorner, QGraphicsLayoutItem * secondItem, Qt::Corner secondCorner), (firstItem, &firstCorner, secondItem, &secondCorner)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout16addCornerAnchorsEP19QGraphicsLayoutItemN2Qt6CornerES1_S3_", ffiqt.FFI_TYPE_VOID, this.cthis, firstItem, &firstCorner, secondItem, &secondCorner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:94
// index:0
// void setHorizontalSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetHorizontalSpacing(spacing float64) {
	// 0: (, qreal spacing), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout20setHorizontalSpacingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:95
// index:0
// void setVerticalSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetVerticalSpacing(spacing float64) {
	// 0: (, qreal spacing), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout18setVerticalSpacingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:96
// index:0
// void setSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetSpacing(spacing float64) {
	// 0: (, qreal spacing), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout10setSpacingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:97
// index:0
// qreal horizontalSpacing()
func (this *QGraphicsAnchorLayout) HorizontalSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout17horizontalSpacingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:98
// index:0
// qreal verticalSpacing()
func (this *QGraphicsAnchorLayout) VerticalSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout15verticalSpacingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:100
// index:0
// virtual
// void removeAt(int)
func (this *QGraphicsAnchorLayout) RemoveAt(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout8removeAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:101
// index:0
// virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsAnchorLayout) SetGeometry(rect unsafe.Pointer) {
	// 0: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout11setGeometryERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:102
// index:0
// virtual
// int count()
func (this *QGraphicsAnchorLayout) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:103
// index:0
// virtual
// QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsAnchorLayout) ItemAt(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout6itemAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:105
// index:0
// virtual
// void invalidate()
func (this *QGraphicsAnchorLayout) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout10invalidateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
