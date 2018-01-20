//  header block begin
// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QGraphicsEffect struct {
	*qtcore.QObject
}

func (this *QGraphicsEffect) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:64
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsEffect) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:82
// index:0
// void QGraphicsEffect(class QObject *)
func NewQGraphicsEffect(parent unsafe.Pointer) *QGraphicsEffect {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffectC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsEffectFromPointer(cthis)
	return gothis
}
func NewQGraphicsEffectFromPointer(cthis unsafe.Pointer) *QGraphicsEffect {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsEffect{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:98
// index:1
// void QGraphicsEffect(class QGraphicsEffectPrivate &, class QObject *)
func NewQGraphicsEffect_1(d unsafe.Pointer, parent unsafe.Pointer) *QGraphicsEffect {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffectC1ER22QGraphicsEffectPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, d, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsEffectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:83
// index:0
// virtual
// void ~QGraphicsEffect()
func DeleteQGraphicsEffect(*QGraphicsEffect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:85
// index:0
// virtual
// QRectF boundingRectFor(const class QRectF &)
func (this *QGraphicsEffect) BoundingRectFor(sourceRect unsafe.Pointer) {
	// 0: (, sourceRect const QRectF &), (sourceRect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect15boundingRectForERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceRect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:86
// index:0
// QRectF boundingRect()
func (this *QGraphicsEffect) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:88
// index:0
// bool isEnabled()
func (this *QGraphicsEffect) IsEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect9isEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:91
// index:0
// void setEnabled(_Bool)
func (this *QGraphicsEffect) SetEnabled(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect10setEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:92
// index:0
// void update()
func (this *QGraphicsEffect) Update() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect6updateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:95
// index:0
// void enabledChanged(_Bool)
func (this *QGraphicsEffect) EnabledChanged(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect14enabledChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:99
// index:0
// pure virtual
// void draw(class QPainter *)
func (this *QGraphicsEffect) Draw(painter unsafe.Pointer) {
	// 0: (, painter QPainter *), (painter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect4drawEP8QPainter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:100
// index:0
// virtual
// void sourceChanged(QGraphicsEffect::ChangeFlags)
func (this *QGraphicsEffect) SourceChanged(flags int) {
	// 0: (, QFlags<QGraphicsEffect::ChangeFlag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect13sourceChangedE6QFlagsINS_10ChangeFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:101
// index:0
// void updateBoundingRect()
func (this *QGraphicsEffect) UpdateBoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect18updateBoundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:103
// index:0
// bool sourceIsPixmap()
func (this *QGraphicsEffect) SourceIsPixmap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect14sourceIsPixmapEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:104
// index:0
// QRectF sourceBoundingRect(Qt::CoordinateSystem)
func (this *QGraphicsEffect) SourceBoundingRect(system int) {
	// 0: (, system Qt::CoordinateSystem), (&system)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect18sourceBoundingRectEN2Qt16CoordinateSystemE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &system)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:105
// index:0
// void drawSource(class QPainter *)
func (this *QGraphicsEffect) DrawSource(painter unsafe.Pointer) {
	// 0: (, painter QPainter *), (painter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect10drawSourceEP8QPainter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:106
// index:0
// QPixmap sourcePixmap(Qt::CoordinateSystem, class QPoint *, enum QGraphicsEffect::PixmapPadMode)
func (this *QGraphicsEffect) SourcePixmap(system int, offset unsafe.Pointer, mode int) {
	// 0: (, system Qt::CoordinateSystem, offset QPoint *, mode QGraphicsEffect::PixmapPadMode), (&system, offset, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &system, offset, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:120
// index:0
// QGraphicsEffectSource * source()
func (this *QGraphicsEffect) Source() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect6sourceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
