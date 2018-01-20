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

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:79
// index:0
// void QGraphicsAnchorLayout(class QGraphicsLayoutItem *)
func NewQGraphicsAnchorLayout(parent unsafe.Pointer) *QGraphicsAnchorLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayoutC2EP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsAnchorLayoutFromPointer(cthis)
	return gothis
}
func NewQGraphicsAnchorLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsAnchorLayout {
	bcthis0 := NewQGraphicsLayoutFromPointer(cthis)
	return &QGraphicsAnchorLayout{bcthis0}
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
	// 0: (, firstItem QGraphicsLayoutItem *, firstEdge Qt::AnchorPoint, secondItem QGraphicsLayoutItem *, secondEdge Qt::AnchorPoint), (firstItem, &firstEdge, secondItem, &secondEdge)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout9addAnchorEP19QGraphicsLayoutItemN2Qt11AnchorPointES1_S3_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), firstItem, &firstEdge, secondItem, &secondEdge)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:84
// index:0
// QGraphicsAnchor * anchor(class QGraphicsLayoutItem *, Qt::AnchorPoint, class QGraphicsLayoutItem *, Qt::AnchorPoint)
func (this *QGraphicsAnchorLayout) Anchor(firstItem unsafe.Pointer, firstEdge int, secondItem unsafe.Pointer, secondEdge int) {
	// 0: (, firstItem QGraphicsLayoutItem *, firstEdge Qt::AnchorPoint, secondItem QGraphicsLayoutItem *, secondEdge Qt::AnchorPoint), (firstItem, &firstEdge, secondItem, &secondEdge)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout6anchorEP19QGraphicsLayoutItemN2Qt11AnchorPointES1_S3_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), firstItem, &firstEdge, secondItem, &secondEdge)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:87
// index:0
// void addCornerAnchors(class QGraphicsLayoutItem *, Qt::Corner, class QGraphicsLayoutItem *, Qt::Corner)
func (this *QGraphicsAnchorLayout) AddCornerAnchors(firstItem unsafe.Pointer, firstCorner int, secondItem unsafe.Pointer, secondCorner int) {
	// 0: (, firstItem QGraphicsLayoutItem *, firstCorner Qt::Corner, secondItem QGraphicsLayoutItem *, secondCorner Qt::Corner), (firstItem, &firstCorner, secondItem, &secondCorner)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout16addCornerAnchorsEP19QGraphicsLayoutItemN2Qt6CornerES1_S3_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), firstItem, &firstCorner, secondItem, &secondCorner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:90
// index:0
// void addAnchors(class QGraphicsLayoutItem *, class QGraphicsLayoutItem *, Qt::Orientations)
func (this *QGraphicsAnchorLayout) AddAnchors(firstItem unsafe.Pointer, secondItem unsafe.Pointer, orientations int) {
	// 0: (, firstItem QGraphicsLayoutItem *, secondItem QGraphicsLayoutItem *, QFlags<Qt::Orientation> orientations), (firstItem, secondItem, &orientations)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout10addAnchorsEP19QGraphicsLayoutItemS1_6QFlagsIN2Qt11OrientationEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), firstItem, secondItem, &orientations)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:94
// index:0
// void setHorizontalSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetHorizontalSpacing(spacing float64) {
	// 0: (, spacing qreal), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout20setHorizontalSpacingEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:95
// index:0
// void setVerticalSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetVerticalSpacing(spacing float64) {
	// 0: (, spacing qreal), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout18setVerticalSpacingEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:96
// index:0
// void setSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetSpacing(spacing float64) {
	// 0: (, spacing qreal), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout10setSpacingEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:97
// index:0
// qreal horizontalSpacing()
func (this *QGraphicsAnchorLayout) HorizontalSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout17horizontalSpacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:98
// index:0
// qreal verticalSpacing()
func (this *QGraphicsAnchorLayout) VerticalSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout15verticalSpacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:100
// index:0
// virtual
// void removeAt(int)
func (this *QGraphicsAnchorLayout) RemoveAt(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout8removeAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:101
// index:0
// virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsAnchorLayout) SetGeometry(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout11setGeometryERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:102
// index:0
// virtual
// int count()
func (this *QGraphicsAnchorLayout) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:103
// index:0
// virtual
// QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsAnchorLayout) ItemAt(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout6itemAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:105
// index:0
// virtual
// void invalidate()
func (this *QGraphicsAnchorLayout) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout10invalidateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:107
// index:0
// virtual
// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsAnchorLayout) SizeHint(which int, constraint unsafe.Pointer) {
	// 0: (, which Qt::SizeHint, constraint const QSizeF &), (&which, constraint)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &which, constraint)
	gopp.ErrPrint(err, rv)
}

//  body block end
