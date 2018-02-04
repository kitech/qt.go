package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>

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
// void draw(class QPainter *)
func (this *QGraphicsEffect) InheritDraw(f func(painter *qtgui.QPainter /*777 QPainter **/)) {
	ffiqt.SetAllInheritCallback(this, "draw", f)
}

// void sourceChanged(QGraphicsEffect::ChangeFlags)
func (this *QGraphicsEffect) InheritSourceChanged(f func(flags int)) {
	ffiqt.SetAllInheritCallback(this, "sourceChanged", f)
}

// void updateBoundingRect()
func (this *QGraphicsEffect) InheritUpdateBoundingRect(f func()) {
	ffiqt.SetAllInheritCallback(this, "updateBoundingRect", f)
}

// bool sourceIsPixmap()
func (this *QGraphicsEffect) InheritSourceIsPixmap(f func() bool) {
	ffiqt.SetAllInheritCallback(this, "sourceIsPixmap", f)
}

// QRectF sourceBoundingRect(Qt::CoordinateSystem)
func (this *QGraphicsEffect) InheritSourceBoundingRect(f func(system int) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "sourceBoundingRect", f)
}

// void drawSource(class QPainter *)
func (this *QGraphicsEffect) InheritDrawSource(f func(painter *qtgui.QPainter /*777 QPainter **/)) {
	ffiqt.SetAllInheritCallback(this, "drawSource", f)
}

// QPixmap sourcePixmap(Qt::CoordinateSystem, class QPoint *, enum QGraphicsEffect::PixmapPadMode)
func (this *QGraphicsEffect) InheritSourcePixmap(f func(system int, offset *qtcore.QPoint /*777 QPoint **/, mode int) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "sourcePixmap", f)
}

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
func (this *QGraphicsEffect) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGraphicsEffectFromPointer(cthis unsafe.Pointer) *QGraphicsEffect {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsEffect{bcthis0}
}
func (*QGraphicsEffect) NewFromPointer(cthis unsafe.Pointer) *QGraphicsEffect {
	return NewQGraphicsEffectFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGraphicsEffect) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsEffect(QObject *)
func NewQGraphicsEffect(parent *qtcore.QObject /*777 QObject **/) *QGraphicsEffect {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffectC1EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsEffect()
func DeleteQGraphicsEffect(this *QGraphicsEffect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffectD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRectFor(const QRectF &)
func (this *QGraphicsEffect) BoundingRectFor(sourceRect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = sourceRect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect15boundingRectForERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:86
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QGraphicsEffect) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled()
func (this *QGraphicsEffect) IsEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QGraphicsEffect) SetEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()
func (this *QGraphicsEffect) Update() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect6updateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void enabledChanged(_Bool)
func (this *QGraphicsEffect) EnabledChanged(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect14enabledChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:99
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void draw(QPainter *)
func (this *QGraphicsEffect) Draw(painter *qtgui.QPainter /*777 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect4drawEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:100
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void sourceChanged(QGraphicsEffect::ChangeFlags)
func (this *QGraphicsEffect) SourceChanged(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect13sourceChangedE6QFlagsINS_10ChangeFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:101
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateBoundingRect()
func (this *QGraphicsEffect) UpdateBoundingRect() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect18updateBoundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:103
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool sourceIsPixmap()
func (this *QGraphicsEffect) SourceIsPixmap() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect14sourceIsPixmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:104
// index:0
// Protected Visibility=Default Availability=Available
// [32] QRectF sourceBoundingRect(Qt::CoordinateSystem)
func (this *QGraphicsEffect) SourceBoundingRect(system int) *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect18sourceBoundingRectEN2Qt16CoordinateSystemE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), system)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:105
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void drawSource(QPainter *)
func (this *QGraphicsEffect) DrawSource(painter *qtgui.QPainter /*777 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsEffect10drawSourceEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [32] QPixmap sourcePixmap(Qt::CoordinateSystem, QPoint *, enum QGraphicsEffect::PixmapPadMode)
func (this *QGraphicsEffect) SourcePixmap(system int, offset *qtcore.QPoint /*777 QPoint **/, mode int) *qtgui.QPixmap /*123*/ {
	var convArg1 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), system, convArg1, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

type QGraphicsEffect__ChangeFlag = int

const QGraphicsEffect__SourceAttached QGraphicsEffect__ChangeFlag = 1
const QGraphicsEffect__SourceDetached QGraphicsEffect__ChangeFlag = 2
const QGraphicsEffect__SourceBoundingRectChanged QGraphicsEffect__ChangeFlag = 4
const QGraphicsEffect__SourceInvalidated QGraphicsEffect__ChangeFlag = 8

type QGraphicsEffect__PixmapPadMode = int

const QGraphicsEffect__NoPad QGraphicsEffect__PixmapPadMode = 0
const QGraphicsEffect__PadToTransparentBorder QGraphicsEffect__PixmapPadMode = 1
const QGraphicsEffect__PadToEffectiveBoundingRect QGraphicsEffect__PixmapPadMode = 2

//  body block end
