//  header block begin
// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QGraphicsDropShadowEffect struct {
	*QGraphicsEffect
}

func (this *QGraphicsDropShadowEffect) GetCthis() unsafe.Pointer {
	return this.QGraphicsEffect.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:198
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsDropShadowEffect) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:205
// index:0
// void QGraphicsDropShadowEffect(class QObject *)
func NewQGraphicsDropShadowEffect(parent unsafe.Pointer) *QGraphicsDropShadowEffect {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffectC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsDropShadowEffectFromPointer(cthis)
	return gothis
}
func NewQGraphicsDropShadowEffectFromPointer(cthis unsafe.Pointer) *QGraphicsDropShadowEffect {
	bcthis0 := NewQGraphicsEffectFromPointer(cthis)
	return &QGraphicsDropShadowEffect{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:206
// index:0
// virtual
// void ~QGraphicsDropShadowEffect()
func DeleteQGraphicsDropShadowEffect(*QGraphicsDropShadowEffect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:208
// index:0
// virtual
// QRectF boundingRectFor(const class QRectF &)
func (this *QGraphicsDropShadowEffect) BoundingRectFor(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect15boundingRectForERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:209
// index:0
// QPointF offset()
func (this *QGraphicsDropShadowEffect) Offset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect6offsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:211
// index:0
// inline
// qreal xOffset()
func (this *QGraphicsDropShadowEffect) XOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect7xOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:214
// index:0
// inline
// qreal yOffset()
func (this *QGraphicsDropShadowEffect) YOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect7yOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:217
// index:0
// qreal blurRadius()
func (this *QGraphicsDropShadowEffect) BlurRadius() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect10blurRadiusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:218
// index:0
// QColor color()
func (this *QGraphicsDropShadowEffect) Color() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect5colorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:221
// index:0
// void setOffset(const class QPointF &)
func (this *QGraphicsDropShadowEffect) SetOffset(ofs unsafe.Pointer) {
	// 0: (, ofs const QPointF &), (ofs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect9setOffsetERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ofs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:223
// index:1
// inline
// void setOffset(qreal, qreal)
func (this *QGraphicsDropShadowEffect) SetOffset_1(dx float64, dy float64) {
	// 1: (, dx qreal, dy qreal), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect9setOffsetEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:226
// index:2
// inline
// void setOffset(qreal)
func (this *QGraphicsDropShadowEffect) SetOffset_2(d float64) {
	// 2: (, d qreal), (&d)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect9setOffsetEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &d)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:229
// index:0
// inline
// void setXOffset(qreal)
func (this *QGraphicsDropShadowEffect) SetXOffset(dx float64) {
	// 0: (, dx qreal), (&dx)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect10setXOffsetEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:232
// index:0
// inline
// void setYOffset(qreal)
func (this *QGraphicsDropShadowEffect) SetYOffset(dy float64) {
	// 0: (, dy qreal), (&dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect10setYOffsetEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:235
// index:0
// void setBlurRadius(qreal)
func (this *QGraphicsDropShadowEffect) SetBlurRadius(blurRadius float64) {
	// 0: (, blurRadius qreal), (&blurRadius)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect13setBlurRadiusEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &blurRadius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:236
// index:0
// void setColor(const class QColor &)
func (this *QGraphicsDropShadowEffect) SetColor(color unsafe.Pointer) {
	// 0: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect8setColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:239
// index:0
// void offsetChanged(const class QPointF &)
func (this *QGraphicsDropShadowEffect) OffsetChanged(offset unsafe.Pointer) {
	// 0: (, offset const QPointF &), (offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect13offsetChangedERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:240
// index:0
// void blurRadiusChanged(qreal)
func (this *QGraphicsDropShadowEffect) BlurRadiusChanged(blurRadius float64) {
	// 0: (, blurRadius qreal), (&blurRadius)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect17blurRadiusChangedEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &blurRadius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:241
// index:0
// void colorChanged(const class QColor &)
func (this *QGraphicsDropShadowEffect) ColorChanged(color unsafe.Pointer) {
	// 0: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect12colorChangedERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:244
// index:0
// virtual
// void draw(class QPainter *)
func (this *QGraphicsDropShadowEffect) Draw(painter unsafe.Pointer) {
	// 0: (, painter QPainter *), (painter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect4drawEP8QPainter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter)
	gopp.ErrPrint(err, rv)
}

//  body block end