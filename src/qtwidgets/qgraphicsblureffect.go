//  header block begin
// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QGraphicsBlurEffect struct {
	*QGraphicsEffect
}

func (this *QGraphicsBlurEffect) GetCthis() unsafe.Pointer {
	return this.QGraphicsEffect.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:157
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsBlurEffect) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:170
// index:0
// void QGraphicsBlurEffect(class QObject *)
func NewQGraphicsBlurEffect(parent unsafe.Pointer) *QGraphicsBlurEffect {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffectC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsBlurEffectFromPointer(cthis)
	return gothis
}
func NewQGraphicsBlurEffectFromPointer(cthis unsafe.Pointer) *QGraphicsBlurEffect {
	bcthis0 := NewQGraphicsEffectFromPointer(cthis)
	return &QGraphicsBlurEffect{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:171
// index:0
// virtual
// void ~QGraphicsBlurEffect()
func DeleteQGraphicsBlurEffect(*QGraphicsBlurEffect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:173
// index:0
// virtual
// QRectF boundingRectFor(const class QRectF &)
func (this *QGraphicsBlurEffect) BoundingRectFor(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect15boundingRectForERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:174
// index:0
// qreal blurRadius()
func (this *QGraphicsBlurEffect) BlurRadius() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect10blurRadiusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:175
// index:0
// QGraphicsBlurEffect::BlurHints blurHints()
func (this *QGraphicsBlurEffect) BlurHints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect9blurHintsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:178
// index:0
// void setBlurRadius(qreal)
func (this *QGraphicsBlurEffect) SetBlurRadius(blurRadius float64) {
	// 0: (, blurRadius qreal), (&blurRadius)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect13setBlurRadiusEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &blurRadius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:179
// index:0
// void setBlurHints(QGraphicsBlurEffect::BlurHints)
func (this *QGraphicsBlurEffect) SetBlurHints(hints int) {
	// 0: (, QFlags<QGraphicsBlurEffect::BlurHint> hints), (hints)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect12setBlurHintsE6QFlagsINS_8BlurHintEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), hints)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:182
// index:0
// void blurRadiusChanged(qreal)
func (this *QGraphicsBlurEffect) BlurRadiusChanged(blurRadius float64) {
	// 0: (, blurRadius qreal), (&blurRadius)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect17blurRadiusChangedEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &blurRadius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:183
// index:0
// void blurHintsChanged(QGraphicsBlurEffect::BlurHints)
func (this *QGraphicsBlurEffect) BlurHintsChanged(hints int) {
	// 0: (, QFlags<QGraphicsBlurEffect::BlurHint> hints), (hints)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect16blurHintsChangedE6QFlagsINS_8BlurHintEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), hints)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:186
// index:0
// virtual
// void draw(class QPainter *)
func (this *QGraphicsBlurEffect) Draw(painter unsafe.Pointer) {
	// 0: (, painter QPainter *), (painter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect4drawEP8QPainter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter)
	gopp.ErrPrint(err, rv)
}

//  body block end
