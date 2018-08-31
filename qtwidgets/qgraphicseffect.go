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

// void draw(QPainter *)
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

// void drawSource(QPainter *)
func (this *QGraphicsEffect) InheritDrawSource(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawSource", f)
}

// QPixmap sourcePixmap(Qt::CoordinateSystem, QPoint *, QGraphicsEffect::PixmapPadMode)
func (this *QGraphicsEffect) InheritSourcePixmap(f func(system int, offset *qtcore.QPoint /*777 QPoint **/, mode int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sourcePixmap", f)
}

/*

 */
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

/*

 */
func (this *QGraphicsEffect) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsEffect(QObject *)

/*
Constructs a new QGraphicsEffect instance having the specified parent.
*/
func (*QGraphicsEffect) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsEffect {
	return NewQGraphicsEffect(parent)
}
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

/*
Constructs a new QGraphicsEffect instance having the specified parent.
*/
func (*QGraphicsEffect) NewForInherit__() *QGraphicsEffect {
	return NewQGraphicsEffect__()
}
func NewQGraphicsEffect__() *QGraphicsEffect {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
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

/*

 */
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

/*
Returns the effective bounding rectangle for this effect, given the provided rect in the device coordinates. When writing you own custom effect, you must call updateBoundingRect() whenever any parameters are changed that may cause this this function to return a different value.

See also sourceBoundingRect().
*/
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

/*
Returns the effective bounding rectangle for this effect, i.e., the bounding rectangle of the source in device coordinates, adjusted by any margins applied by the effect itself.

See also boundingRectFor() and updateBoundingRect().
*/
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

/*

 */
func (this *QGraphicsEffect) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(bool)

/*

 */
func (this *QGraphicsEffect) SetEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()

/*
Schedules a redraw of the effect. Call this function whenever the effect needs to be redrawn. This function does not trigger a redraw of the source.

See also updateBoundingRect().
*/
func (this *QGraphicsEffect) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void enabledChanged(bool)

/*
This signal is emitted whenever the effect is enabled or disabled. The enabled parameter holds the effects's new enabled state.

Note: Notifier signal for property enabled.

See also isEnabled().
*/
func (this *QGraphicsEffect) EnabledChanged(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect14enabledChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:99
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void draw(QPainter *)

/*
This pure virtual function draws the effect and is called whenever the source needs to be drawn.

Reimplement this function in a QGraphicsEffect subclass to provide the effect's drawing implementation, using painter.

For example:


  MyGraphicsEffect::draw(QPainter *painter)
  {
      ...
      QPoint offset;
      if (sourceIsPixmap()) {
          // No point in drawing in device coordinates (pixmap will be scaled anyways).
          const QPixmap pixmap = sourcePixmap(Qt::LogicalCoordinates, &offset);
          ...
          painter->drawPixmap(offset, pixmap);
      } else {
          // Draw pixmap in device coordinates to avoid pixmap scaling;
          const QPixmap pixmap = sourcePixmap(Qt::DeviceCoordinates, &offset);
          painter->setWorldTransform(QTransform());
          ...
          painter->drawPixmap(offset, pixmap);
      }
      ...
  }



This function should not be called explicitly by the user, since it is meant for reimplementation purposes only.
*/
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

/*
This virtual function is called by QGraphicsEffect to notify the effect that the source has changed. If the effect applies any cache, then this cache must be purged in order to reflect the new appearance of the source.

The flags describes what has changed.
*/
func (this *QGraphicsEffect) SourceChanged(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect13sourceChangedE6QFlagsINS_10ChangeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:101
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateBoundingRect()

/*
This function notifies the effect framework when the effect's bounding rectangle has changed. As a custom effect author, you must call this function whenever you change any parameters that will cause the virtual boundingRectFor() function to return a different value.

This function will call update() if this is necessary.

See also boundingRectFor(), boundingRect(), and sourceBoundingRect().
*/
func (this *QGraphicsEffect) UpdateBoundingRect() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect18updateBoundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:103
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool sourceIsPixmap() const

/*
Returns true if the source effectively is a pixmap, e.g., a QGraphicsPixmapItem.

This function is useful for optimization purposes. For instance, there's no point in drawing the source in device coordinates to avoid pixmap scaling if this function returns true - the source pixmap will be scaled anyways.
*/
func (this *QGraphicsEffect) SourceIsPixmap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect14sourceIsPixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:104
// index:0
// Protected Visibility=Default Availability=Available
// [32] QRectF sourceBoundingRect(Qt::CoordinateSystem) const

/*
Returns the bounding rectangle of the source mapped to the given system.

Calling this function with Qt::DeviceCoordinates outside of QGraphicsEffect::draw() will give undefined results, as there is no device context available.

See also draw().
*/
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

/*
Returns the bounding rectangle of the source mapped to the given system.

Calling this function with Qt::DeviceCoordinates outside of QGraphicsEffect::draw() will give undefined results, as there is no device context available.

See also draw().
*/
func (this *QGraphicsEffect) SourceBoundingRect__() *qtcore.QRectF /*123*/ {
	// arg: 0, Qt::CoordinateSystem=Elaborated, Qt::CoordinateSystem=Enum, , Invalid
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

/*
Draws the source directly using the given painter.

This function should only be called from QGraphicsEffect::draw().

For example:


  MyGraphicsOpacityEffect::draw(QPainter *painter)
  {
      // Fully opaque; draw directly without going through a pixmap.
      if (qFuzzyCompare(m_opacity, 1)) {
          drawSource(painter);
          return;
      }
      ...
  }



See also QGraphicsEffect::draw().
*/
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
// [32] QPixmap sourcePixmap(Qt::CoordinateSystem, QPoint *, QGraphicsEffect::PixmapPadMode) const

/*
Returns a pixmap with the source painted into it.

The system specifies which coordinate system to be used for the source. The optional offset parameter returns the offset where the pixmap should be painted at using the current painter. For control on how the pixmap is padded use the mode parameter.

The returned pixmap is clipped to the current painter's device rectangle when system is Qt::DeviceCoordinates.

Calling this function with Qt::DeviceCoordinates outside of QGraphicsEffect::draw() will give undefined results, as there is no device context available.

See also draw() and boundingRect().
*/
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
// [32] QPixmap sourcePixmap(Qt::CoordinateSystem, QPoint *, QGraphicsEffect::PixmapPadMode) const

/*
Returns a pixmap with the source painted into it.

The system specifies which coordinate system to be used for the source. The optional offset parameter returns the offset where the pixmap should be painted at using the current painter. For control on how the pixmap is padded use the mode parameter.

The returned pixmap is clipped to the current painter's device rectangle when system is Qt::DeviceCoordinates.

Calling this function with Qt::DeviceCoordinates outside of QGraphicsEffect::draw() will give undefined results, as there is no device context available.

See also draw() and boundingRect().
*/
func (this *QGraphicsEffect) SourcePixmap__() *qtgui.QPixmap /*123*/ {
	// arg: 0, Qt::CoordinateSystem=Elaborated, Qt::CoordinateSystem=Enum, , Invalid
	system := 0
	// arg: 1, QPoint *=Pointer, QPoint=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, QGraphicsEffect::PixmapPadMode=Enum, QGraphicsEffect::PixmapPadMode=Enum, , Invalid
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
// [32] QPixmap sourcePixmap(Qt::CoordinateSystem, QPoint *, QGraphicsEffect::PixmapPadMode) const

/*
Returns a pixmap with the source painted into it.

The system specifies which coordinate system to be used for the source. The optional offset parameter returns the offset where the pixmap should be painted at using the current painter. For control on how the pixmap is padded use the mode parameter.

The returned pixmap is clipped to the current painter's device rectangle when system is Qt::DeviceCoordinates.

Calling this function with Qt::DeviceCoordinates outside of QGraphicsEffect::draw() will give undefined results, as there is no device context available.

See also draw() and boundingRect().
*/
func (this *QGraphicsEffect) SourcePixmap__1(system int) *qtgui.QPixmap /*123*/ {
	// arg: 1, QPoint *=Pointer, QPoint=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, QGraphicsEffect::PixmapPadMode=Enum, QGraphicsEffect::PixmapPadMode=Enum, , Invalid
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
// [32] QPixmap sourcePixmap(Qt::CoordinateSystem, QPoint *, QGraphicsEffect::PixmapPadMode) const

/*
Returns a pixmap with the source painted into it.

The system specifies which coordinate system to be used for the source. The optional offset parameter returns the offset where the pixmap should be painted at using the current painter. For control on how the pixmap is padded use the mode parameter.

The returned pixmap is clipped to the current painter's device rectangle when system is Qt::DeviceCoordinates.

Calling this function with Qt::DeviceCoordinates outside of QGraphicsEffect::draw() will give undefined results, as there is no device context available.

See also draw() and boundingRect().
*/
func (this *QGraphicsEffect) SourcePixmap__2(system int, offset qtcore.QPoint_ITF /*777 QPoint **/) *qtgui.QPixmap /*123*/ {
	var convArg1 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg1 = offset.QPoint_PTR().GetCthis()
	}
	// arg: 2, QGraphicsEffect::PixmapPadMode=Enum, QGraphicsEffect::PixmapPadMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), system, convArg1, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

/*


 */
type QGraphicsEffect__ChangeFlag = int

//
const QGraphicsEffect__SourceAttached QGraphicsEffect__ChangeFlag = 1

//
const QGraphicsEffect__SourceDetached QGraphicsEffect__ChangeFlag = 2

//
const QGraphicsEffect__SourceBoundingRectChanged QGraphicsEffect__ChangeFlag = 4

//
const QGraphicsEffect__SourceInvalidated QGraphicsEffect__ChangeFlag = 8

func (this *QGraphicsEffect) ChangeFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsEffect_ChangeFlagItemName(val int) string {
	var nilthis *QGraphicsEffect
	return nilthis.ChangeFlagItemName(val)
}

/*
This enum describes how the pixmap returned from sourcePixmap should be padded.


*/
type QGraphicsEffect__PixmapPadMode = int

// The pixmap should not receive any additional padding.
const QGraphicsEffect__NoPad QGraphicsEffect__PixmapPadMode = 0

// The pixmap should be padded to ensure it has a completely transparent border.
const QGraphicsEffect__PadToTransparentBorder QGraphicsEffect__PixmapPadMode = 1

// The pixmap should be padded to match the effective bounding rectangle of the effect.
const QGraphicsEffect__PadToEffectiveBoundingRect QGraphicsEffect__PixmapPadMode = 2

func (this *QGraphicsEffect) PixmapPadModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsEffect_PixmapPadModeItemName(val int) string {
	var nilthis *QGraphicsEffect
	return nilthis.PixmapPadModeItemName(val)
}

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
