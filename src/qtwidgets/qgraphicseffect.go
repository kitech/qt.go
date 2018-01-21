package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func NewQGraphicsEffectFromPointer(cthis unsafe.Pointer) *QGraphicsEffect {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsEffect{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:64
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsEffect) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:82
// index:0
// Public
// void QGraphicsEffect(class QObject *)
func NewQGraphicsEffect(parent *qtcore.QObject /*444 QObject **/) *QGraphicsEffect {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffectC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsEffectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:83
// index:0
// Public virtual
// void ~QGraphicsEffect()
func DeleteQGraphicsEffect(*QGraphicsEffect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:85
// index:0
// Public virtual
// QRectF boundingRectFor(const class QRectF &)
func (this *QGraphicsEffect) BoundingRectFor(sourceRect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = sourceRect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect15boundingRectForERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:86
// index:0
// Public
// QRectF boundingRect()
func (this *QGraphicsEffect) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:88
// index:0
// Public
// bool isEnabled()
func (this *QGraphicsEffect) IsEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:91
// index:0
// Public
// void setEnabled(_Bool)
func (this *QGraphicsEffect) SetEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:92
// index:0
// Public
// void update()
func (this *QGraphicsEffect) Update() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect6updateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:95
// index:0
// Public
// void enabledChanged(_Bool)
func (this *QGraphicsEffect) EnabledChanged(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect14enabledChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:99
// index:0
// Protected pure virtual
// void draw(class QPainter *)
func (this *QGraphicsEffect) Draw(painter *qtgui.QPainter /*444 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect4drawEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:101
// index:0
// Protected
// void updateBoundingRect()
func (this *QGraphicsEffect) UpdateBoundingRect() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect18updateBoundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:103
// index:0
// Protected
// bool sourceIsPixmap()
func (this *QGraphicsEffect) SourceIsPixmap() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect14sourceIsPixmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:104
// index:0
// Protected
// QRectF sourceBoundingRect(Qt::CoordinateSystem)
func (this *QGraphicsEffect) SourceBoundingRect(system int) *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect18sourceBoundingRectEN2Qt16CoordinateSystemE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &system)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:105
// index:0
// Protected
// void drawSource(class QPainter *)
func (this *QGraphicsEffect) DrawSource(painter *qtgui.QPainter /*444 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect10drawSourceEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:106
// index:0
// Protected
// QPixmap sourcePixmap(Qt::CoordinateSystem, class QPoint *, enum QGraphicsEffect::PixmapPadMode)
func (this *QGraphicsEffect) SourcePixmap(system int, offset *qtcore.QPoint /*444 QPoint **/, mode int) *qtgui.QPixmap /*123*/ {
	var convArg1 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &system, convArg1, &mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
