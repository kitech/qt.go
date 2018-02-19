package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void draw(class QPainter *)
func (this *QGraphicsEffect) InheritDraw(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "draw", f)
}

// void sourceChanged(QGraphicsEffect::ChangeFlags)
func (this *QGraphicsEffect) InheritSourceChanged(f func(flags int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "sourceChanged", f)
}

// void updateBoundingRect()
func (this *QGraphicsEffect) InheritUpdateBoundingRect(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateBoundingRect", f)
}

// bool sourceIsPixmap()
func (this *QGraphicsEffect) InheritSourceIsPixmap(f func() bool) {
	qtrt.SetAllInheritCallback(this, "sourceIsPixmap", f)
}

// QRectF sourceBoundingRect(Qt::CoordinateSystem)
func (this *QGraphicsEffect) InheritSourceBoundingRect(f func(system int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sourceBoundingRect", f)
}

// void drawSource(class QPainter *)
func (this *QGraphicsEffect) InheritDrawSource(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawSource", f)
}

// QPixmap sourcePixmap(Qt::CoordinateSystem, class QPoint *, enum QGraphicsEffect::PixmapPadMode)
func (this *QGraphicsEffect) InheritSourcePixmap(f func(system int, offset *qtcore.QPoint /*777 QPoint **/, mode int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sourcePixmap", f)
}

type QGraphicsEffect struct {
	*qtcore.QObject
}
type QGraphicsEffect_ITF interface {
	qtcore.QObject_ITF
	QGraphicsEffect_PTR() *QGraphicsEffect
}

func (ptr *QGraphicsEffect) QGraphicsEffect_PTR() *QGraphicsEffect { return ptr }

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
// [8] const QMetaObject * metaObject() const
func (this *QGraphicsEffect) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsEffect(QObject *)
func NewQGraphicsEffect(parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsEffect {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsEffect")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsEffect(QObject *)
func NewQGraphicsEffect__() *QGraphicsEffect {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsEffect")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsEffect()
func DeleteQGraphicsEffect(this *QGraphicsEffect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRectFor(const QRectF &) const
func (this *QGraphicsEffect) BoundingRectFor(sourceRect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if sourceRect != nil && sourceRect.QRectF_PTR() != nil {
		convArg0 = sourceRect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect15boundingRectForERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:86
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect() const
func (this *QGraphicsEffect) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled() const
func (this *QGraphicsEffect) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QGraphicsEffect) SetEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()
func (this *QGraphicsEffect) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void enabledChanged(_Bool)
func (this *QGraphicsEffect) EnabledChanged(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect14enabledChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:99
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void draw(QPainter *)
func (this *QGraphicsEffect) Draw(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect4drawEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:100
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void sourceChanged(QGraphicsEffect::ChangeFlags)
func (this *QGraphicsEffect) SourceChanged(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect13sourceChangedE6QFlagsINS_10ChangeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:101
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateBoundingRect()
func (this *QGraphicsEffect) UpdateBoundingRect() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect18updateBoundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:103
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool sourceIsPixmap() const
func (this *QGraphicsEffect) SourceIsPixmap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect14sourceIsPixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:104
// index:0
// Protected Visibility=Default Availability=Available
// [32] QRectF sourceBoundingRect(Qt::CoordinateSystem) const
func (this *QGraphicsEffect) SourceBoundingRect(system int) *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect18sourceBoundingRectEN2Qt16CoordinateSystemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), system)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:104
// index:0
// Protected Visibility=Default Availability=Available
// [32] QRectF sourceBoundingRect(Qt::CoordinateSystem) const
func (this *QGraphicsEffect) SourceBoundingRect__() *qtcore.QRectF /*123*/ {
	// arg: 0, Qt::CoordinateSystem=Elaborated, Qt::CoordinateSystem=Enum,
	system := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect18sourceBoundingRectEN2Qt16CoordinateSystemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), system)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:105
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void drawSource(QPainter *)
func (this *QGraphicsEffect) DrawSource(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect10drawSourceEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [32] QPixmap sourcePixmap(Qt::CoordinateSystem, QPoint *, enum QGraphicsEffect::PixmapPadMode) const
func (this *QGraphicsEffect) SourcePixmap(system int, offset qtcore.QPoint_ITF /*777 QPoint **/, mode int) *qtgui.QPixmap /*123*/ {
	var convArg1 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg1 = offset.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), system, convArg1, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [32] QPixmap sourcePixmap(Qt::CoordinateSystem, QPoint *, enum QGraphicsEffect::PixmapPadMode) const
func (this *QGraphicsEffect) SourcePixmap__() *qtgui.QPixmap /*123*/ {
	// arg: 0, Qt::CoordinateSystem=Elaborated, Qt::CoordinateSystem=Enum,
	system := 0
	// arg: 1, QPoint *=Pointer, QPoint=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, QGraphicsEffect::PixmapPadMode=Enum, QGraphicsEffect::PixmapPadMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), system, convArg1, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [32] QPixmap sourcePixmap(Qt::CoordinateSystem, QPoint *, enum QGraphicsEffect::PixmapPadMode) const
func (this *QGraphicsEffect) SourcePixmap__1(system int) *qtgui.QPixmap /*123*/ {
	// arg: 1, QPoint *=Pointer, QPoint=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, QGraphicsEffect::PixmapPadMode=Enum, QGraphicsEffect::PixmapPadMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), system, convArg1, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [32] QPixmap sourcePixmap(Qt::CoordinateSystem, QPoint *, enum QGraphicsEffect::PixmapPadMode) const
func (this *QGraphicsEffect) SourcePixmap__2(system int, offset qtcore.QPoint_ITF /*777 QPoint **/) *qtgui.QPixmap /*123*/ {
	var convArg1 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg1 = offset.QPoint_PTR().GetCthis()
	}
	// arg: 2, QGraphicsEffect::PixmapPadMode=Enum, QGraphicsEffect::PixmapPadMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), system, convArg1, mode)
	qtrt.ErrPrint(err, rv)
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

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
