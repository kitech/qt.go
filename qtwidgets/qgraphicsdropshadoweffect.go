package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>

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
type QGraphicsDropShadowEffect struct {
	*QGraphicsEffect
}

func (this *QGraphicsDropShadowEffect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsEffect.GetCthis()
	}
}
func (this *QGraphicsDropShadowEffect) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsEffect = NewQGraphicsEffectFromPointer(cthis)
}
func NewQGraphicsDropShadowEffectFromPointer(cthis unsafe.Pointer) *QGraphicsDropShadowEffect {
	bcthis0 := NewQGraphicsEffectFromPointer(cthis)
	return &QGraphicsDropShadowEffect{bcthis0}
}
func (*QGraphicsDropShadowEffect) NewFromPointer(cthis unsafe.Pointer) *QGraphicsDropShadowEffect {
	return NewQGraphicsDropShadowEffectFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:198
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsDropShadowEffect) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:205
// index:0
// Public
// void QGraphicsDropShadowEffect(QObject *)
func NewQGraphicsDropShadowEffect(parent *qtcore.QObject /*777 QObject **/) *QGraphicsDropShadowEffect {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffectC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsDropShadowEffectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:206
// index:0
// Public virtual
// void ~QGraphicsDropShadowEffect()
func DeleteQGraphicsDropShadowEffect(*QGraphicsDropShadowEffect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:208
// index:0
// Public virtual
// QRectF boundingRectFor(const QRectF &)
func (this *QGraphicsDropShadowEffect) BoundingRectFor(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect15boundingRectForERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:209
// index:0
// Public
// QPointF offset()
func (this *QGraphicsDropShadowEffect) Offset() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect6offsetEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:211
// index:0
// Public inline
// qreal xOffset()
func (this *QGraphicsDropShadowEffect) XOffset() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect7xOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:214
// index:0
// Public inline
// qreal yOffset()
func (this *QGraphicsDropShadowEffect) YOffset() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect7yOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:217
// index:0
// Public
// qreal blurRadius()
func (this *QGraphicsDropShadowEffect) BlurRadius() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect10blurRadiusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:218
// index:0
// Public
// QColor color()
func (this *QGraphicsDropShadowEffect) Color() *qtgui.QColor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect5colorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:221
// index:0
// Public
// void setOffset(const QPointF &)
func (this *QGraphicsDropShadowEffect) SetOffset(ofs *qtcore.QPointF) {
	var convArg0 = ofs.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect9setOffsetERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:223
// index:1
// Public inline
// void setOffset(qreal, qreal)
func (this *QGraphicsDropShadowEffect) SetOffset_1(dx float64, dy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect9setOffsetEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:226
// index:2
// Public inline
// void setOffset(qreal)
func (this *QGraphicsDropShadowEffect) SetOffset_2(d float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect9setOffsetEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), d)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:229
// index:0
// Public inline
// void setXOffset(qreal)
func (this *QGraphicsDropShadowEffect) SetXOffset(dx float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect10setXOffsetEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:232
// index:0
// Public inline
// void setYOffset(qreal)
func (this *QGraphicsDropShadowEffect) SetYOffset(dy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect10setYOffsetEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:235
// index:0
// Public
// void setBlurRadius(qreal)
func (this *QGraphicsDropShadowEffect) SetBlurRadius(blurRadius float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect13setBlurRadiusEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), blurRadius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:236
// index:0
// Public
// void setColor(const QColor &)
func (this *QGraphicsDropShadowEffect) SetColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect8setColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:239
// index:0
// Public
// void offsetChanged(const QPointF &)
func (this *QGraphicsDropShadowEffect) OffsetChanged(offset *qtcore.QPointF) {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect13offsetChangedERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:240
// index:0
// Public
// void blurRadiusChanged(qreal)
func (this *QGraphicsDropShadowEffect) BlurRadiusChanged(blurRadius float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect17blurRadiusChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), blurRadius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:241
// index:0
// Public
// void colorChanged(const QColor &)
func (this *QGraphicsDropShadowEffect) ColorChanged(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect12colorChangedERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:244
// index:0
// Protected virtual
// void draw(QPainter *)
func (this *QGraphicsDropShadowEffect) Draw(painter *qtgui.QPainter /*777 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect4drawEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
