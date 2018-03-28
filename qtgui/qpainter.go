package qtgui

// /usr/include/qt/QtGui/qpainter.h
// #include <qpainter.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QPainter struct {
	*qtrt.CObject
}
type QPainter_ITF interface {
	QPainter_PTR() *QPainter
}

func (ptr *QPainter) QPainter_PTR() *QPainter { return ptr }

func (this *QPainter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPainter) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPainterFromPointer(cthis unsafe.Pointer) *QPainter {
	return &QPainter{&qtrt.CObject{cthis}}
}
func (*QPainter) NewFromPointer(cthis unsafe.Pointer) *QPainter {
	return NewQPainterFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpainter.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPainter()

/*
Constructs a painter.

See also begin() and end().
*/
func NewQPainter() *QPainter {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainterC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPainterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPainter)
	return gothis
}

// /usr/include/qt/QtGui/qpainter.h:125
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPainter(QPaintDevice *)

/*
Constructs a painter.

See also begin() and end().
*/
func NewQPainter_1(arg0 QPaintDevice_ITF /*777 QPaintDevice **/) *QPainter {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintDevice_PTR() != nil {
		convArg0 = arg0.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainterC2EP12QPaintDevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPainterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPainter)
	return gothis
}

// /usr/include/qt/QtGui/qpainter.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPainter()

/*

 */
func DeleteQPainter(this *QPainter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpainter.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QPaintDevice * device() const

/*
Returns the paint device on which this painter is currently painting, or 0 if the painter is not active.

See also isActive().
*/
func (this *QPainter) Device() *QPaintDevice /*777 QPaintDevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpainter.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool begin(QPaintDevice *)

/*
Begins painting the paint device and returns true if successful; otherwise returns false.

Notice that all painter settings (setPen(), setBrush() etc.) are reset to default values when begin() is called.

The errors that can occur are serious problems, such as these:


  painter->begin(0); // impossible - paint device cannot be 0

  QPixmap image(0, 0);
  painter->begin(&image); // impossible - image.isNull() == true;

  painter->begin(myWidget);
  painter2->begin(myWidget); // impossible - only one painter at a time



Note that most of the time, you can use one of the constructors instead of begin(), and that end() is automatically done at destruction.

Warning: A paint device can only be painted by one painter at a time.

Warning: Painting on a QImage with the format QImage::Format_Indexed8 is not supported.

See also end() and QPainter().
*/
func (this *QPainter) Begin(arg0 QPaintDevice_ITF /*777 QPaintDevice **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintDevice_PTR() != nil {
		convArg0 = arg0.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter5beginEP12QPaintDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool end()

/*
Ends painting. Any resources used while painting are released. You don't normally need to call this since it is called by the destructor.

Returns true if the painter is no longer active; otherwise returns false.

See also begin() and isActive().
*/
func (this *QPainter) End() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:132
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive() const

/*
Returns true if begin() has been called and end() has not yet been called; otherwise returns false.

See also begin() and QPaintDevice::paintingActive().
*/
func (this *QPainter) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void initFrom(const QPaintDevice *)

/*

 */
func (this *QPainter) InitFrom(device QPaintDevice_ITF /*777 const QPaintDevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QPaintDevice_PTR() != nil {
		convArg0 = device.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8initFromEPK12QPaintDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompositionMode(QPainter::CompositionMode)

/*
Sets the composition mode to the given mode.

Warning: Only a QPainter operating on a QImage fully supports all composition modes. The RasterOp modes are supported for X11 as described in compositionMode().

See also compositionMode().
*/
func (this *QPainter) SetCompositionMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter18setCompositionModeENS_15CompositionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:181
// index:0
// Public Visibility=Default Availability=Available
// [4] QPainter::CompositionMode compositionMode() const

/*
Returns the current composition mode.

See also CompositionMode and setCompositionMode().
*/
func (this *QPainter) CompositionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter15compositionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:183
// index:0
// Public Visibility=Default Availability=Available
// [16] const QFont & font() const

/*
Returns the currently set font used for drawing text.

See also setFont(), drawText(), and Settings.
*/
func (this *QPainter) Font() *QFont {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)

/*
Sets the painter's font to the given font.

This font is used by subsequent drawText() functions. The text color is the same as the pen color.

If you set a font that isn't available, Qt finds a close match. font() will return what you set using setFont() and fontInfo() returns the font actually being used (which may be the same).

See also font(), drawText(), and Settings.
*/
func (this *QPainter) SetFont(f QFont_ITF) {
	var convArg0 unsafe.Pointer
	if f != nil && f.QFont_PTR() != nil {
		convArg0 = f.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:186
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontMetrics fontMetrics() const

/*
Returns the font metrics for the painter if the painter is active. Otherwise, the return value is undefined.

See also font(), isActive(), and Settings.
*/
func (this *QPainter) FontMetrics() *QFontMetrics /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11fontMetricsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontMetrics)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:187
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontInfo fontInfo() const

/*
Returns the font info for the painter if the painter is active. Otherwise, the return value is undefined.

See also font(), isActive(), and Settings.
*/
func (this *QPainter) FontInfo() *QFontInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter8fontInfoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontInfo)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPen(const QColor &)

/*
Sets the painter's pen to be the given pen.

The pen defines how to draw lines and outlines, and it also defines the text color.

See also pen() and Settings.
*/
func (this *QPainter) SetPen(color QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter6setPenERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:190
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPen(const QPen &)

/*
Sets the painter's pen to be the given pen.

The pen defines how to draw lines and outlines, and it also defines the text color.

See also pen() and Settings.
*/
func (this *QPainter) SetPen_1(pen QPen_ITF) {
	var convArg0 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg0 = pen.QPen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter6setPenERK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:191
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setPen(Qt::PenStyle)

/*
Sets the painter's pen to be the given pen.

The pen defines how to draw lines and outlines, and it also defines the text color.

See also pen() and Settings.
*/
func (this *QPainter) SetPen_2(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter6setPenEN2Qt8PenStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:192
// index:0
// Public Visibility=Default Availability=Available
// [8] const QPen & pen() const

/*
Returns the painter's current pen.

See also setPen() and Settings.
*/
func (this *QPainter) Pen() *QPen {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter3penEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPen)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBrush(const QBrush &)

/*
Sets the painter's brush to the given brush.

The painter's brush defines how shapes are filled.

See also brush() and Settings.
*/
func (this *QPainter) SetBrush(brush QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8setBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:195
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setBrush(Qt::BrushStyle)

/*
Sets the painter's brush to the given brush.

The painter's brush defines how shapes are filled.

See also brush() and Settings.
*/
func (this *QPainter) SetBrush_1(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8setBrushEN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:196
// index:0
// Public Visibility=Default Availability=Available
// [8] const QBrush & brush() const

/*
Returns the painter's current brush.

See also QPainter::setBrush() and Settings.
*/
func (this *QPainter) Brush() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter5brushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundMode(Qt::BGMode)

/*
Sets the background mode of the painter to the given mode

Qt::TransparentMode (the default) draws stippled lines and text without setting the background pixels. Qt::OpaqueMode fills these space with the current background color.

Note that in order to draw a bitmap or pixmap transparently, you must use QPixmap::setMask().

See also backgroundMode(), setBackground(), and Settings.
*/
func (this *QPainter) SetBackgroundMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17setBackgroundModeEN2Qt6BGModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:200
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::BGMode backgroundMode() const

/*
Returns the current background mode.

See also setBackgroundMode() and Settings.
*/
func (this *QPainter) BackgroundMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter14backgroundModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:202
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint brushOrigin() const

/*
Returns the currently set brush origin.

See also setBrushOrigin() and Settings.
*/
func (this *QPainter) BrushOrigin() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11brushOriginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:203
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBrushOrigin(int, int)

/*
Sets the brush origin to position.

The brush origin specifies the (0, 0) coordinate of the painter's brush.

Note that while the brushOrigin() was necessary to adopt the parent's background for a widget in Qt 3, this is no longer the case since the Qt 4 painter doesn't paint the background unless you explicitly tell it to do so by setting the widget's autoFillBackground property to true.

See also brushOrigin() and Settings.
*/
func (this *QPainter) SetBrushOrigin(x int, y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:204
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setBrushOrigin(const QPoint &)

/*
Sets the brush origin to position.

The brush origin specifies the (0, 0) coordinate of the painter's brush.

Note that while the brushOrigin() was necessary to adopt the parent's background for a widget in Qt 3, this is no longer the case since the Qt 4 painter doesn't paint the background unless you explicitly tell it to do so by setting the widget's autoFillBackground property to true.

See also brushOrigin() and Settings.
*/
func (this *QPainter) SetBrushOrigin_1(arg0 qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:205
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setBrushOrigin(const QPointF &)

/*
Sets the brush origin to position.

The brush origin specifies the (0, 0) coordinate of the painter's brush.

Note that while the brushOrigin() was necessary to adopt the parent's background for a widget in Qt 3, this is no longer the case since the Qt 4 painter doesn't paint the background unless you explicitly tell it to do so by setting the widget's autoFillBackground property to true.

See also brushOrigin() and Settings.
*/
func (this *QPainter) SetBrushOrigin_2(arg0 qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPointF_PTR() != nil {
		convArg0 = arg0.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)

/*
Sets the background brush of the painter to the given brush.

The background brush is the brush that is filled in when drawing opaque text, stippled lines and bitmaps. The background brush has no effect in transparent background mode (which is the default).

See also background(), setBackgroundMode(), and Settings.
*/
func (this *QPainter) SetBackground(bg QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if bg != nil && bg.QBrush_PTR() != nil {
		convArg0 = bg.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13setBackgroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:208
// index:0
// Public Visibility=Default Availability=Available
// [8] const QBrush & background() const

/*
Returns the current background brush.

See also setBackground() and Settings.
*/
func (this *QPainter) Background() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter10backgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:210
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opacity() const

/*
Returns the opacity of the painter. The default value is 1.

This function was introduced in  Qt 4.2.

See also setOpacity().
*/
func (this *QPainter) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainter.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacity(qreal)

/*
Sets the opacity of the painter to opacity. The value should be in the range 0.0 to 1.0, where 0.0 is fully transparent and 1.0 is fully opaque.

Opacity set on the painter will apply to all drawing operations individually.

This function was introduced in  Qt 4.2.

See also opacity().
*/
func (this *QPainter) SetOpacity(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10setOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:214
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion clipRegion() const

/*
Returns the currently set clip region. Note that the clip region is given in logical coordinates.

Warning: QPainter does not store the combined clip explicitly as this is handled by the underlying QPaintEngine, so the path is recreated on demand and transformed to the current logical coordinate system. This is potentially an expensive operation.

See also setClipRegion(), clipPath(), and setClipping().
*/
func (this *QPainter) ClipRegion() *QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter10clipRegionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:215
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath clipPath() const

/*
Returns the current clip path in logical coordinates.

Warning: QPainter does not store the combined clip explicitly as this is handled by the underlying QPaintEngine, so the path is recreated on demand and transformed to the current logical coordinate system. This is potentially an expensive operation.

See also setClipPath(), clipRegion(), and setClipping().
*/
func (this *QPainter) ClipPath() *QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter8clipPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipRect(const QRectF &, Qt::ClipOperation)

/*
Enables clipping, and sets the clip region to the given rectangle using the given clip operation. The default operation is to replace the current clip rectangle.

Note that the clip rectangle is specified in logical (painter) coordinates.

See also clipRegion(), setClipping(), and Clipping.
*/
func (this *QPainter) SetClipRect(arg0 qtcore.QRectF_ITF, op int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRectF_PTR() != nil {
		convArg0 = arg0.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipRectERK6QRectFN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipRect(const QRectF &, Qt::ClipOperation)

/*
Enables clipping, and sets the clip region to the given rectangle using the given clip operation. The default operation is to replace the current clip rectangle.

Note that the clip rectangle is specified in logical (painter) coordinates.

See also clipRegion(), setClipping(), and Clipping.
*/
func (this *QPainter) SetClipRect__(arg0 qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRectF_PTR() != nil {
		convArg0 = arg0.QRectF_PTR().GetCthis()
	}
	// arg: 1, Qt::ClipOperation=Elaborated, Qt::ClipOperation=Enum,
	op := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipRectERK6QRectFN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:218
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setClipRect(const QRect &, Qt::ClipOperation)

/*
Enables clipping, and sets the clip region to the given rectangle using the given clip operation. The default operation is to replace the current clip rectangle.

Note that the clip rectangle is specified in logical (painter) coordinates.

See also clipRegion(), setClipping(), and Clipping.
*/
func (this *QPainter) SetClipRect_1(arg0 qtcore.QRect_ITF, op int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipRectERK5QRectN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:218
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setClipRect(const QRect &, Qt::ClipOperation)

/*
Enables clipping, and sets the clip region to the given rectangle using the given clip operation. The default operation is to replace the current clip rectangle.

Note that the clip rectangle is specified in logical (painter) coordinates.

See also clipRegion(), setClipping(), and Clipping.
*/
func (this *QPainter) SetClipRect_1_(arg0 qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	// arg: 1, Qt::ClipOperation=Elaborated, Qt::ClipOperation=Enum,
	op := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipRectERK5QRectN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:219
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void setClipRect(int, int, int, int, Qt::ClipOperation)

/*
Enables clipping, and sets the clip region to the given rectangle using the given clip operation. The default operation is to replace the current clip rectangle.

Note that the clip rectangle is specified in logical (painter) coordinates.

See also clipRegion(), setClipping(), and Clipping.
*/
func (this *QPainter) SetClipRect_2(x int, y int, w int, h int, op int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipRectEiiiiN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:219
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void setClipRect(int, int, int, int, Qt::ClipOperation)

/*
Enables clipping, and sets the clip region to the given rectangle using the given clip operation. The default operation is to replace the current clip rectangle.

Note that the clip rectangle is specified in logical (painter) coordinates.

See also clipRegion(), setClipping(), and Clipping.
*/
func (this *QPainter) SetClipRect_2_(x int, y int, w int, h int) {
	// arg: 4, Qt::ClipOperation=Elaborated, Qt::ClipOperation=Enum,
	op := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipRectEiiiiN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipRegion(const QRegion &, Qt::ClipOperation)

/*
Sets the clip region to the given region using the specified clip operation. The default clip operation is to replace the current clip region.

Note that the clip region is given in logical coordinates.

See also clipRegion(), setClipRect(), and Clipping.
*/
func (this *QPainter) SetClipRegion(arg0 QRegion_ITF, op int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRegion_PTR() != nil {
		convArg0 = arg0.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13setClipRegionERK7QRegionN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipRegion(const QRegion &, Qt::ClipOperation)

/*
Sets the clip region to the given region using the specified clip operation. The default clip operation is to replace the current clip region.

Note that the clip region is given in logical coordinates.

See also clipRegion(), setClipRect(), and Clipping.
*/
func (this *QPainter) SetClipRegion__(arg0 QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRegion_PTR() != nil {
		convArg0 = arg0.QRegion_PTR().GetCthis()
	}
	// arg: 1, Qt::ClipOperation=Elaborated, Qt::ClipOperation=Enum,
	op := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13setClipRegionERK7QRegionN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipPath(const QPainterPath &, Qt::ClipOperation)

/*
Enables clipping, and sets the clip path for the painter to the given path, with the clip operation.

Note that the clip path is specified in logical (painter) coordinates.

See also clipPath(), clipRegion(), and Clipping.
*/
func (this *QPainter) SetClipPath(path QPainterPath_ITF, op int) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipPathERK12QPainterPathN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipPath(const QPainterPath &, Qt::ClipOperation)

/*
Enables clipping, and sets the clip path for the painter to the given path, with the clip operation.

Note that the clip path is specified in logical (painter) coordinates.

See also clipPath(), clipRegion(), and Clipping.
*/
func (this *QPainter) SetClipPath__(path QPainterPath_ITF) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	// arg: 1, Qt::ClipOperation=Elaborated, Qt::ClipOperation=Enum,
	op := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClipPathERK12QPainterPathN2Qt13ClipOperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipping(bool)

/*
Enables clipping if enable is true, or disables clipping if enable is false.

See also hasClipping() and Clipping.
*/
func (this *QPainter) SetClipping(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setClippingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:226
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasClipping() const

/*
Returns true if clipping has been set; otherwise returns false.

See also setClipping() and Clipping.
*/
func (this *QPainter) HasClipping() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11hasClippingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:228
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF clipBoundingRect() const

/*
Returns the bounding rectangle of the current clip if there is a clip; otherwise returns an empty rectangle. Note that the clip region is given in logical coordinates.

The bounding rectangle is not guaranteed to be tight.

This function was introduced in  Qt 4.8.

See also setClipRect(), setClipPath(), and setClipRegion().
*/
func (this *QPainter) ClipBoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter16clipBoundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void save()

/*
Saves the current painter state (pushes the state onto a stack). A save() must be followed by a corresponding restore(); the end() function unwinds the stack.

See also restore().
*/
func (this *QPainter) Save() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter4saveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:231
// index:0
// Public Visibility=Default Availability=Available
// [-2] void restore()

/*
Restores the current painter state (pops a saved state off the stack).

See also save().
*/
func (this *QPainter) Restore() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7restoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:234
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(const QMatrix &, bool)

/*

 */
func (this *QPainter) SetMatrix(matrix QMatrix_ITF, combine bool) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9setMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:234
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(const QMatrix &, bool)

/*

 */
func (this *QPainter) SetMatrix__(matrix QMatrix_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid,
	combine := false
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9setMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:235
// index:0
// Public Visibility=Default Availability=Available
// [48] const QMatrix & matrix() const

/*

 */
func (this *QPainter) Matrix() *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter6matrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:236
// index:0
// Public Visibility=Default Availability=Available
// [48] const QMatrix & deviceMatrix() const

/*

 */
func (this *QPainter) DeviceMatrix() *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter12deviceMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetMatrix()

/*

 */
func (this *QPainter) ResetMatrix() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11resetMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransform(const QTransform &, bool)

/*
Sets the world transformation matrix. If combine is true, the specified transform is combined with the current matrix; otherwise it replaces the current matrix.

This function was introduced in  Qt 4.3.

See also transform() and setWorldTransform().
*/
func (this *QPainter) SetTransform(transform QTransform_ITF, combine bool) {
	var convArg0 unsafe.Pointer
	if transform != nil && transform.QTransform_PTR() != nil {
		convArg0 = transform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12setTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransform(const QTransform &, bool)

/*
Sets the world transformation matrix. If combine is true, the specified transform is combined with the current matrix; otherwise it replaces the current matrix.

This function was introduced in  Qt 4.3.

See also transform() and setWorldTransform().
*/
func (this *QPainter) SetTransform__(transform QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if transform != nil && transform.QTransform_PTR() != nil {
		convArg0 = transform.QTransform_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid,
	combine := false
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12setTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:240
// index:0
// Public Visibility=Default Availability=Available
// [88] const QTransform & transform() const

/*
Returns the world transformation matrix.

See also setTransform() and worldTransform().
*/
func (this *QPainter) Transform() *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter9transformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:241
// index:0
// Public Visibility=Default Availability=Available
// [88] const QTransform & deviceTransform() const

/*
Returns the matrix that transforms from logical coordinates to device coordinates of the platform dependent paint device.

This function is only needed when using platform painting commands on the platform dependent handle (Qt::HANDLE), and the platform does not do transformations nativly.

The QPaintEngine::PaintEngineFeature enum can be queried to determine whether the platform performs the transformations or not.

See also worldTransform() and QPaintEngine::hasFeature().
*/
func (this *QPainter) DeviceTransform() *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter15deviceTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetTransform()

/*
Resets any transformations that were made using translate(), scale(), shear(), rotate(), setWorldTransform(), setViewport() and setWindow().

See also Coordinate Transformations.
*/
func (this *QPainter) ResetTransform() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14resetTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:244
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorldMatrix(const QMatrix &, bool)

/*

 */
func (this *QPainter) SetWorldMatrix(matrix QMatrix_ITF, combine bool) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setWorldMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:244
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorldMatrix(const QMatrix &, bool)

/*

 */
func (this *QPainter) SetWorldMatrix__(matrix QMatrix_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid,
	combine := false
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setWorldMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:245
// index:0
// Public Visibility=Default Availability=Available
// [48] const QMatrix & worldMatrix() const

/*

 */
func (this *QPainter) WorldMatrix() *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11worldMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorldTransform(const QTransform &, bool)

/*
Sets the world transformation matrix. If combine is true, the specified matrix is combined with the current matrix; otherwise it replaces the current matrix.

See also worldTransform(), transform(), and setTransform().
*/
func (this *QPainter) SetWorldTransform(matrix QTransform_ITF, combine bool) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17setWorldTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorldTransform(const QTransform &, bool)

/*
Sets the world transformation matrix. If combine is true, the specified matrix is combined with the current matrix; otherwise it replaces the current matrix.

See also worldTransform(), transform(), and setTransform().
*/
func (this *QPainter) SetWorldTransform__(matrix QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid,
	combine := false
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17setWorldTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:248
// index:0
// Public Visibility=Default Availability=Available
// [88] const QTransform & worldTransform() const

/*
Returns the world transformation matrix.

See also setWorldTransform().
*/
func (this *QPainter) WorldTransform() *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter14worldTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:250
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix combinedMatrix() const

/*

 */
func (this *QPainter) CombinedMatrix() *QMatrix /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter14combinedMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:251
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform combinedTransform() const

/*
Returns the transformation matrix combining the current window/viewport and world transformation.

See also setWorldTransform(), setWindow(), and setViewport().
*/
func (this *QPainter) CombinedTransform() *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter17combinedTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:253
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrixEnabled(bool)

/*

 */
func (this *QPainter) SetMatrixEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter16setMatrixEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:254
// index:0
// Public Visibility=Default Availability=Available
// [1] bool matrixEnabled() const

/*

 */
func (this *QPainter) MatrixEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter13matrixEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorldMatrixEnabled(bool)

/*
Enables transformations if enable is true, or disables transformations if enable is false. The world transformation matrix is not changed.

This function was introduced in  Qt 4.2.

See also worldMatrixEnabled(), worldTransform(), and Coordinate Transformations.
*/
func (this *QPainter) SetWorldMatrixEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter21setWorldMatrixEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:257
// index:0
// Public Visibility=Default Availability=Available
// [1] bool worldMatrixEnabled() const

/*
Returns true if world transformation is enabled; otherwise returns false.

This function was introduced in  Qt 4.2.

See also setWorldMatrixEnabled(), worldTransform(), and Coordinate System.
*/
func (this *QPainter) WorldMatrixEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter18worldMatrixEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scale(qreal, qreal)

/*
Scales the coordinate system by (sx, sy).

See also setWorldTransform() and Coordinate Transformations.
*/
func (this *QPainter) Scale(sx float64, sy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter5scaleEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sx, sy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:260
// index:0
// Public Visibility=Default Availability=Available
// [-2] void shear(qreal, qreal)

/*
Shears the coordinate system by (sh, sv).

See also setWorldTransform() and Coordinate Transformations.
*/
func (this *QPainter) Shear(sh float64, sv float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter5shearEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sh, sv)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rotate(qreal)

/*
Rotates the coordinate system clockwise. The given angle parameter is in degrees.

See also setWorldTransform() and Coordinate Transformations.
*/
func (this *QPainter) Rotate(a float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter6rotateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:263
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(const QPointF &)

/*
Translates the coordinate system by the given offset; i.e. the given offset is added to points.

See also setWorldTransform() and Coordinate Transformations.
*/
func (this *QPainter) Translate(offset qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPointF_PTR() != nil {
		convArg0 = offset.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9translateERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:264
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPoint &)

/*
Translates the coordinate system by the given offset; i.e. the given offset is added to points.

See also setWorldTransform() and Coordinate Transformations.
*/
func (this *QPainter) Translate_1(offset qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg0 = offset.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9translateERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:265
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void translate(qreal, qreal)

/*
Translates the coordinate system by the given offset; i.e. the given offset is added to points.

See also setWorldTransform() and Coordinate Transformations.
*/
func (this *QPainter) Translate_2(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:267
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect window() const

/*
Returns the window rectangle.

See also setWindow() and setViewTransformEnabled().
*/
func (this *QPainter) Window() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindow(const QRect &)

/*
Sets the painter's window to the given rectangle, and enables view transformations.

The window rectangle is part of the view transformation. The window specifies the logical coordinate system. Its sister, the viewport(), specifies the device coordinate system.

The default window rectangle is the same as the device's rectangle.

See also window(), viewTransformEnabled(), and Window-Viewport Conversion.
*/
func (this *QPainter) SetWindow(window qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QRect_PTR() != nil {
		convArg0 = window.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9setWindowERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:269
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setWindow(int, int, int, int)

/*
Sets the painter's window to the given rectangle, and enables view transformations.

The window rectangle is part of the view transformation. The window specifies the logical coordinate system. Its sister, the viewport(), specifies the device coordinate system.

The default window rectangle is the same as the device's rectangle.

See also window(), viewTransformEnabled(), and Window-Viewport Conversion.
*/
func (this *QPainter) SetWindow_1(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9setWindowEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:271
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect viewport() const

/*
Returns the viewport rectangle.

See also setViewport() and setViewTransformEnabled().
*/
func (this *QPainter) Viewport() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter8viewportEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:272
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewport(const QRect &)

/*
Sets the painter's viewport rectangle to the given rectangle, and enables view transformations.

The viewport rectangle is part of the view transformation. The viewport specifies the device coordinate system. Its sister, the window(), specifies the logical coordinate system.

The default viewport rectangle is the same as the device's rectangle.

See also viewport(), viewTransformEnabled(), and Window-Viewport Conversion.
*/
func (this *QPainter) SetViewport(viewport qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if viewport != nil && viewport.QRect_PTR() != nil {
		convArg0 = viewport.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setViewportERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:273
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setViewport(int, int, int, int)

/*
Sets the painter's viewport rectangle to the given rectangle, and enables view transformations.

The viewport rectangle is part of the view transformation. The viewport specifies the device coordinate system. Its sister, the window(), specifies the logical coordinate system.

The default viewport rectangle is the same as the device's rectangle.

See also viewport(), viewTransformEnabled(), and Window-Viewport Conversion.
*/
func (this *QPainter) SetViewport_1(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11setViewportEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewTransformEnabled(bool)

/*
Enables view transformations if enable is true, or disables view transformations if enable is false.

See also viewTransformEnabled() and Window-Viewport Conversion.
*/
func (this *QPainter) SetViewTransformEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter23setViewTransformEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:276
// index:0
// Public Visibility=Default Availability=Available
// [1] bool viewTransformEnabled() const

/*
Returns true if view transformation is enabled; otherwise returns false.

See also setViewTransformEnabled() and worldTransform().
*/
func (this *QPainter) ViewTransformEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter20viewTransformEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:279
// index:0
// Public Visibility=Default Availability=Available
// [-2] void strokePath(const QPainterPath &, const QPen &)

/*
Draws the outline (strokes) the path path with the pen specified by pen

See also fillPath() and Drawing.
*/
func (this *QPainter) StrokePath(path QPainterPath_ITF, pen QPen_ITF) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg1 = pen.QPen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10strokePathERK12QPainterPathRK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:280
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fillPath(const QPainterPath &, const QBrush &)

/*
Fills the given path using the given brush. The outline is not drawn.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawPath().
*/
func (this *QPainter) FillPath(path QPainterPath_ITF, brush QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg1 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillPathERK12QPainterPathRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:281
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPath(const QPainterPath &)

/*
Draws the given painter path using the current pen for outline and the current brush for filling.



  QPainterPath path;
  path.moveTo(20, 80);
  path.lineTo(20, 30);
  path.cubicTo(80, 0, 50, 50, 80, 80);

  QPainter painter(this);
  painter.drawPath(path);





See also the Painter Paths example and the Vector Deformation example.
*/
func (this *QPainter) DrawPath(path QPainterPath_ITF) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawPathERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:283
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawPoint(const QPointF &)

/*
Draws a single point at the given position using the current pen's color.

See also Coordinate System.
*/
func (this *QPainter) DrawPoint(pt qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPointF_PTR() != nil {
		convArg0 = pt.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:284
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPoint(const QPoint &)

/*
Draws a single point at the given position using the current pen's color.

See also Coordinate System.
*/
func (this *QPainter) DrawPoint_1(p qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawPointERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:285
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawPoint(int, int)

/*
Draws a single point at the given position using the current pen's color.

See also Coordinate System.
*/
func (this *QPainter) DrawPoint_2(x int, y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawPointEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:287
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPoints(const QPointF *, int)

/*
Draws the first pointCount points in the array points using the current pen's color.

See also Coordinate System.
*/
func (this *QPainter) DrawPoints(points qtcore.QPointF_ITF /*777 const QPointF **/, pointCount int) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPointF_PTR() != nil {
		convArg0 = points.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPointsEPK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:288
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPoints(const QPolygonF &)

/*
Draws the first pointCount points in the array points using the current pen's color.

See also Coordinate System.
*/
func (this *QPainter) DrawPoints_1(points QPolygonF_ITF) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPolygonF_PTR() != nil {
		convArg0 = points.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPointsERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:289
// index:2
// Public Visibility=Default Availability=Available
// [-2] void drawPoints(const QPoint *, int)

/*
Draws the first pointCount points in the array points using the current pen's color.

See also Coordinate System.
*/
func (this *QPainter) DrawPoints_2(points qtcore.QPoint_ITF /*777 const QPoint **/, pointCount int) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPoint_PTR() != nil {
		convArg0 = points.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPointsEPK6QPointi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:290
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawPoints(const QPolygon &)

/*
Draws the first pointCount points in the array points using the current pen's color.

See also Coordinate System.
*/
func (this *QPainter) DrawPoints_3(points QPolygon_ITF) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPolygon_PTR() != nil {
		convArg0 = points.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPointsERK8QPolygon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:292
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawLine(const QLineF &)

/*
Draws a line defined by line.



  QLineF line(10.0, 80.0, 90.0, 20.0);

  QPainter(this);
  painter.drawLine(line);





See also drawLines(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawLine(line qtcore.QLineF_ITF) {
	var convArg0 unsafe.Pointer
	if line != nil && line.QLineF_PTR() != nil {
		convArg0 = line.QLineF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawLineERK6QLineF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:293
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawLine(const QLine &)

/*
Draws a line defined by line.



  QLineF line(10.0, 80.0, 90.0, 20.0);

  QPainter(this);
  painter.drawLine(line);





See also drawLines(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawLine_1(line qtcore.QLine_ITF) {
	var convArg0 unsafe.Pointer
	if line != nil && line.QLine_PTR() != nil {
		convArg0 = line.QLine_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawLineERK5QLine", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:294
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawLine(int, int, int, int)

/*
Draws a line defined by line.



  QLineF line(10.0, 80.0, 90.0, 20.0);

  QPainter(this);
  painter.drawLine(line);





See also drawLines(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawLine_2(x1 int, y1 int, x2 int, y2 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawLineEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:295
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawLine(const QPoint &, const QPoint &)

/*
Draws a line defined by line.



  QLineF line(10.0, 80.0, 90.0, 20.0);

  QPainter(this);
  painter.drawLine(line);





See also drawLines(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawLine_3(p1 qtcore.QPoint_ITF, p2 qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p1 != nil && p1.QPoint_PTR() != nil {
		convArg0 = p1.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if p2 != nil && p2.QPoint_PTR() != nil {
		convArg1 = p2.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawLineERK6QPointS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:296
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void drawLine(const QPointF &, const QPointF &)

/*
Draws a line defined by line.



  QLineF line(10.0, 80.0, 90.0, 20.0);

  QPainter(this);
  painter.drawLine(line);





See also drawLines(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawLine_4(p1 qtcore.QPointF_ITF, p2 qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p1 != nil && p1.QPointF_PTR() != nil {
		convArg0 = p1.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if p2 != nil && p2.QPointF_PTR() != nil {
		convArg1 = p2.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawLineERK7QPointFS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:298
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawLines(const QLineF *, int)

/*
Draws the first lineCount lines in the array lines using the current pen.

See also drawLine() and drawPolyline().
*/
func (this *QPainter) DrawLines(lines qtcore.QLineF_ITF /*777 const QLineF **/, lineCount int) {
	var convArg0 unsafe.Pointer
	if lines != nil && lines.QLineF_PTR() != nil {
		convArg0 = lines.QLineF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK6QLineFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:300
// index:1
// Public Visibility=Default Availability=Available
// [-2] void drawLines(const QPointF *, int)

/*
Draws the first lineCount lines in the array lines using the current pen.

See also drawLine() and drawPolyline().
*/
func (this *QPainter) DrawLines_1(pointPairs qtcore.QPointF_ITF /*777 const QPointF **/, lineCount int) {
	var convArg0 unsafe.Pointer
	if pointPairs != nil && pointPairs.QPointF_PTR() != nil {
		convArg0 = pointPairs.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:302
// index:2
// Public Visibility=Default Availability=Available
// [-2] void drawLines(const QLine *, int)

/*
Draws the first lineCount lines in the array lines using the current pen.

See also drawLine() and drawPolyline().
*/
func (this *QPainter) DrawLines_2(lines qtcore.QLine_ITF /*777 const QLine **/, lineCount int) {
	var convArg0 unsafe.Pointer
	if lines != nil && lines.QLine_PTR() != nil {
		convArg0 = lines.QLine_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK5QLinei", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:304
// index:3
// Public Visibility=Default Availability=Available
// [-2] void drawLines(const QPoint *, int)

/*
Draws the first lineCount lines in the array lines using the current pen.

See also drawLine() and drawPolyline().
*/
func (this *QPainter) DrawLines_3(pointPairs qtcore.QPoint_ITF /*777 const QPoint **/, lineCount int) {
	var convArg0 unsafe.Pointer
	if pointPairs != nil && pointPairs.QPoint_PTR() != nil {
		convArg0 = pointPairs.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK6QPointi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:307
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawRect(const QRectF &)

/*
Draws the current rectangle with the current pen and brush.

A filled rectangle has a size of rectangle.size(). A stroked rectangle has a size of rectangle.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawRect(rectangle);





See also drawRects(), drawPolygon(), and Coordinate System.
*/
func (this *QPainter) DrawRect(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:308
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawRect(int, int, int, int)

/*
Draws the current rectangle with the current pen and brush.

A filled rectangle has a size of rectangle.size(). A stroked rectangle has a size of rectangle.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawRect(rectangle);





See also drawRects(), drawPolygon(), and Coordinate System.
*/
func (this *QPainter) DrawRect_1(x1 int, y1 int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawRectEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:309
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawRect(const QRect &)

/*
Draws the current rectangle with the current pen and brush.

A filled rectangle has a size of rectangle.size(). A stroked rectangle has a size of rectangle.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawRect(rectangle);





See also drawRects(), drawPolygon(), and Coordinate System.
*/
func (this *QPainter) DrawRect_2(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawRects(const QRectF *, int)

/*
Draws the first rectCount of the given rectangles using the current pen and brush.

See also drawRect().
*/
func (this *QPainter) DrawRects(rects qtcore.QRectF_ITF /*777 const QRectF **/, rectCount int) {
	var convArg0 unsafe.Pointer
	if rects != nil && rects.QRectF_PTR() != nil {
		convArg0 = rects.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawRectsEPK6QRectFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rectCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:313
// index:1
// Public Visibility=Default Availability=Available
// [-2] void drawRects(const QRect *, int)

/*
Draws the first rectCount of the given rectangles using the current pen and brush.

See also drawRect().
*/
func (this *QPainter) DrawRects_1(rects qtcore.QRect_ITF /*777 const QRect **/, rectCount int) {
	var convArg0 unsafe.Pointer
	if rects != nil && rects.QRect_PTR() != nil {
		convArg0 = rects.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawRectsEPK5QRecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rectCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:316
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawEllipse(const QRectF &)

/*
Draws the ellipse defined by the given rectangle.

A filled ellipse has a size of rectangle.size(). A stroked ellipse has a size of rectangle.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawEllipse(rectangle);





See also drawPie() and Coordinate System.
*/
func (this *QPainter) DrawEllipse(r qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:317
// index:1
// Public Visibility=Default Availability=Available
// [-2] void drawEllipse(const QRect &)

/*
Draws the ellipse defined by the given rectangle.

A filled ellipse has a size of rectangle.size(). A stroked ellipse has a size of rectangle.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawEllipse(rectangle);





See also drawPie() and Coordinate System.
*/
func (this *QPainter) DrawEllipse_1(r qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:318
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawEllipse(int, int, int, int)

/*
Draws the ellipse defined by the given rectangle.

A filled ellipse has a size of rectangle.size(). A stroked ellipse has a size of rectangle.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawEllipse(rectangle);





See also drawPie() and Coordinate System.
*/
func (this *QPainter) DrawEllipse_2(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawEllipseEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:320
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawEllipse(const QPointF &, qreal, qreal)

/*
Draws the ellipse defined by the given rectangle.

A filled ellipse has a size of rectangle.size(). A stroked ellipse has a size of rectangle.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawEllipse(rectangle);





See also drawPie() and Coordinate System.
*/
func (this *QPainter) DrawEllipse_3(center qtcore.QPointF_ITF, rx float64, ry float64) {
	var convArg0 unsafe.Pointer
	if center != nil && center.QPointF_PTR() != nil {
		convArg0 = center.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK7QPointFdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rx, ry)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:321
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void drawEllipse(const QPoint &, int, int)

/*
Draws the ellipse defined by the given rectangle.

A filled ellipse has a size of rectangle.size(). A stroked ellipse has a size of rectangle.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawEllipse(rectangle);





See also drawPie() and Coordinate System.
*/
func (this *QPainter) DrawEllipse_4(center qtcore.QPoint_ITF, rx int, ry int) {
	var convArg0 unsafe.Pointer
	if center != nil && center.QPoint_PTR() != nil {
		convArg0 = center.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK6QPointii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rx, ry)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:323
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPolyline(const QPointF *, int)

/*
Draws the polyline defined by the first pointCount points in points using the current pen.

Note that unlike the drawPolygon() function the last point is not connected to the first, neither is the polyline filled.



  static const QPointF points[3] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
  };

  QPainter painter(this);
  painter.drawPolyline(points, 3);





See also drawLines(), drawPolygon(), and Coordinate System.
*/
func (this *QPainter) DrawPolyline(points qtcore.QPointF_ITF /*777 const QPointF **/, pointCount int) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPointF_PTR() != nil {
		convArg0 = points.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawPolylineEPK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:324
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPolyline(const QPolygonF &)

/*
Draws the polyline defined by the first pointCount points in points using the current pen.

Note that unlike the drawPolygon() function the last point is not connected to the first, neither is the polyline filled.



  static const QPointF points[3] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
  };

  QPainter painter(this);
  painter.drawPolyline(points, 3);





See also drawLines(), drawPolygon(), and Coordinate System.
*/
func (this *QPainter) DrawPolyline_1(polyline QPolygonF_ITF) {
	var convArg0 unsafe.Pointer
	if polyline != nil && polyline.QPolygonF_PTR() != nil {
		convArg0 = polyline.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawPolylineERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:325
// index:2
// Public Visibility=Default Availability=Available
// [-2] void drawPolyline(const QPoint *, int)

/*
Draws the polyline defined by the first pointCount points in points using the current pen.

Note that unlike the drawPolygon() function the last point is not connected to the first, neither is the polyline filled.



  static const QPointF points[3] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
  };

  QPainter painter(this);
  painter.drawPolyline(points, 3);





See also drawLines(), drawPolygon(), and Coordinate System.
*/
func (this *QPainter) DrawPolyline_2(points qtcore.QPoint_ITF /*777 const QPoint **/, pointCount int) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPoint_PTR() != nil {
		convArg0 = points.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawPolylineEPK6QPointi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:326
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawPolyline(const QPolygon &)

/*
Draws the polyline defined by the first pointCount points in points using the current pen.

Note that unlike the drawPolygon() function the last point is not connected to the first, neither is the polyline filled.



  static const QPointF points[3] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
  };

  QPainter painter(this);
  painter.drawPolyline(points, 3);





See also drawLines(), drawPolygon(), and Coordinate System.
*/
func (this *QPainter) DrawPolyline_3(polygon QPolygon_ITF) {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygon_PTR() != nil {
		convArg0 = polygon.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawPolylineERK8QPolygon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:328
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPointF *, int, Qt::FillRule)

/*
Draws the polygon defined by the first pointCount points in the array points using the current pen and brush.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush().

If fillRule is Qt::WindingFill, the polygon is filled using the winding fill algorithm. If fillRule is Qt::OddEvenFill, the polygon is filled using the odd-even fill algorithm. See Qt::FillRule for a more detailed description of these fill rules.

See also drawConvexPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawPolygon(points qtcore.QPointF_ITF /*777 const QPointF **/, pointCount int, fillRule int) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPointF_PTR() != nil {
		convArg0 = points.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonEPK7QPointFiN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:328
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPointF *, int, Qt::FillRule)

/*
Draws the polygon defined by the first pointCount points in the array points using the current pen and brush.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush().

If fillRule is Qt::WindingFill, the polygon is filled using the winding fill algorithm. If fillRule is Qt::OddEvenFill, the polygon is filled using the odd-even fill algorithm. See Qt::FillRule for a more detailed description of these fill rules.

See also drawConvexPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawPolygon__(points qtcore.QPointF_ITF /*777 const QPointF **/, pointCount int) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPointF_PTR() != nil {
		convArg0 = points.QPointF_PTR().GetCthis()
	}
	// arg: 2, Qt::FillRule=Elaborated, Qt::FillRule=Enum,
	fillRule := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonEPK7QPointFiN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:329
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPolygonF &, Qt::FillRule)

/*
Draws the polygon defined by the first pointCount points in the array points using the current pen and brush.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush().

If fillRule is Qt::WindingFill, the polygon is filled using the winding fill algorithm. If fillRule is Qt::OddEvenFill, the polygon is filled using the odd-even fill algorithm. See Qt::FillRule for a more detailed description of these fill rules.

See also drawConvexPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawPolygon_1(polygon QPolygonF_ITF, fillRule int) {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonERK9QPolygonFN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:329
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPolygonF &, Qt::FillRule)

/*
Draws the polygon defined by the first pointCount points in the array points using the current pen and brush.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush().

If fillRule is Qt::WindingFill, the polygon is filled using the winding fill algorithm. If fillRule is Qt::OddEvenFill, the polygon is filled using the odd-even fill algorithm. See Qt::FillRule for a more detailed description of these fill rules.

See also drawConvexPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawPolygon_1_(polygon QPolygonF_ITF) {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	// arg: 1, Qt::FillRule=Elaborated, Qt::FillRule=Enum,
	fillRule := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonERK9QPolygonFN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:330
// index:2
// Public Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPoint *, int, Qt::FillRule)

/*
Draws the polygon defined by the first pointCount points in the array points using the current pen and brush.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush().

If fillRule is Qt::WindingFill, the polygon is filled using the winding fill algorithm. If fillRule is Qt::OddEvenFill, the polygon is filled using the odd-even fill algorithm. See Qt::FillRule for a more detailed description of these fill rules.

See also drawConvexPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawPolygon_2(points qtcore.QPoint_ITF /*777 const QPoint **/, pointCount int, fillRule int) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPoint_PTR() != nil {
		convArg0 = points.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonEPK6QPointiN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:330
// index:2
// Public Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPoint *, int, Qt::FillRule)

/*
Draws the polygon defined by the first pointCount points in the array points using the current pen and brush.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush().

If fillRule is Qt::WindingFill, the polygon is filled using the winding fill algorithm. If fillRule is Qt::OddEvenFill, the polygon is filled using the odd-even fill algorithm. See Qt::FillRule for a more detailed description of these fill rules.

See also drawConvexPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawPolygon_2_(points qtcore.QPoint_ITF /*777 const QPoint **/, pointCount int) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPoint_PTR() != nil {
		convArg0 = points.QPoint_PTR().GetCthis()
	}
	// arg: 2, Qt::FillRule=Elaborated, Qt::FillRule=Enum,
	fillRule := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonEPK6QPointiN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:331
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPolygon &, Qt::FillRule)

/*
Draws the polygon defined by the first pointCount points in the array points using the current pen and brush.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush().

If fillRule is Qt::WindingFill, the polygon is filled using the winding fill algorithm. If fillRule is Qt::OddEvenFill, the polygon is filled using the odd-even fill algorithm. See Qt::FillRule for a more detailed description of these fill rules.

See also drawConvexPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawPolygon_3(polygon QPolygon_ITF, fillRule int) {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygon_PTR() != nil {
		convArg0 = polygon.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonERK8QPolygonN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:331
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawPolygon(const QPolygon &, Qt::FillRule)

/*
Draws the polygon defined by the first pointCount points in the array points using the current pen and brush.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush().

If fillRule is Qt::WindingFill, the polygon is filled using the winding fill algorithm. If fillRule is Qt::OddEvenFill, the polygon is filled using the odd-even fill algorithm. See Qt::FillRule for a more detailed description of these fill rules.

See also drawConvexPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawPolygon_3_(polygon QPolygon_ITF) {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygon_PTR() != nil {
		convArg0 = polygon.QPolygon_PTR().GetCthis()
	}
	// arg: 1, Qt::FillRule=Elaborated, Qt::FillRule=Enum,
	fillRule := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPolygonERK8QPolygonN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:333
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawConvexPolygon(const QPointF *, int)

/*
Draws the convex polygon defined by the first pointCount points in the array points using the current pen.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawConvexPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush(). If the supplied polygon is not convex, i.e. it contains at least one angle larger than 180 degrees, the results are undefined.

On some platforms (e.g. X11), the drawConvexPolygon() function can be faster than the drawPolygon() function.

See also drawPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawConvexPolygon(points qtcore.QPointF_ITF /*777 const QPointF **/, pointCount int) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPointF_PTR() != nil {
		convArg0 = points.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonEPK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:334
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawConvexPolygon(const QPolygonF &)

/*
Draws the convex polygon defined by the first pointCount points in the array points using the current pen.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawConvexPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush(). If the supplied polygon is not convex, i.e. it contains at least one angle larger than 180 degrees, the results are undefined.

On some platforms (e.g. X11), the drawConvexPolygon() function can be faster than the drawPolygon() function.

See also drawPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawConvexPolygon_1(polygon QPolygonF_ITF) {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:335
// index:2
// Public Visibility=Default Availability=Available
// [-2] void drawConvexPolygon(const QPoint *, int)

/*
Draws the convex polygon defined by the first pointCount points in the array points using the current pen.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawConvexPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush(). If the supplied polygon is not convex, i.e. it contains at least one angle larger than 180 degrees, the results are undefined.

On some platforms (e.g. X11), the drawConvexPolygon() function can be faster than the drawPolygon() function.

See also drawPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawConvexPolygon_2(points qtcore.QPoint_ITF /*777 const QPoint **/, pointCount int) {
	var convArg0 unsafe.Pointer
	if points != nil && points.QPoint_PTR() != nil {
		convArg0 = points.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonEPK6QPointi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:336
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawConvexPolygon(const QPolygon &)

/*
Draws the convex polygon defined by the first pointCount points in the array points using the current pen.



  static const QPointF points[4] = {
      QPointF(10.0, 80.0),
      QPointF(20.0, 10.0),
      QPointF(80.0, 30.0),
      QPointF(90.0, 70.0)
  };

  QPainter painter(this);
  painter.drawConvexPolygon(points, 4);





The first point is implicitly connected to the last point, and the polygon is filled with the current brush(). If the supplied polygon is not convex, i.e. it contains at least one angle larger than 180 degrees, the results are undefined.

On some platforms (e.g. X11), the drawConvexPolygon() function can be faster than the drawPolygon() function.

See also drawPolygon(), drawPolyline(), and Coordinate System.
*/
func (this *QPainter) DrawConvexPolygon_3(polygon QPolygon_ITF) {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygon_PTR() != nil {
		convArg0 = polygon.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonERK8QPolygon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawArc(const QRectF &, int, int)

/*
Draws the arc defined by the given rectangle, startAngle and spanAngle.

The startAngle and spanAngle must be specified in 1/16th of a degree, i.e. a full circle equals 5760 (16 * 360). Positive values for the angles mean counter-clockwise while negative values mean the clockwise direction. Zero degrees is at the 3 o'clock position.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);
  int startAngle = 30 * 16;
  int spanAngle = 120 * 16;

  QPainter painter(this);
  painter.drawArc(rectangle, startAngle, spanAngle);





See also drawPie(), drawChord(), and Coordinate System.
*/
func (this *QPainter) DrawArc(rect qtcore.QRectF_ITF, a int, alen int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawArcERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:339
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawArc(const QRect &, int, int)

/*
Draws the arc defined by the given rectangle, startAngle and spanAngle.

The startAngle and spanAngle must be specified in 1/16th of a degree, i.e. a full circle equals 5760 (16 * 360). Positive values for the angles mean counter-clockwise while negative values mean the clockwise direction. Zero degrees is at the 3 o'clock position.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);
  int startAngle = 30 * 16;
  int spanAngle = 120 * 16;

  QPainter painter(this);
  painter.drawArc(rectangle, startAngle, spanAngle);





See also drawPie(), drawChord(), and Coordinate System.
*/
func (this *QPainter) DrawArc_1(arg0 qtcore.QRect_ITF, a int, alen int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawArcERK5QRectii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:340
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawArc(int, int, int, int, int, int)

/*
Draws the arc defined by the given rectangle, startAngle and spanAngle.

The startAngle and spanAngle must be specified in 1/16th of a degree, i.e. a full circle equals 5760 (16 * 360). Positive values for the angles mean counter-clockwise while negative values mean the clockwise direction. Zero degrees is at the 3 o'clock position.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);
  int startAngle = 30 * 16;
  int spanAngle = 120 * 16;

  QPainter painter(this);
  painter.drawArc(rectangle, startAngle, spanAngle);





See also drawPie(), drawChord(), and Coordinate System.
*/
func (this *QPainter) DrawArc_2(x int, y int, w int, h int, a int, alen int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawArcEiiiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:342
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPie(const QRectF &, int, int)

/*
Draws a pie defined by the given rectangle, startAngle and spanAngle.

The pie is filled with the current brush().

The startAngle and spanAngle must be specified in 1/16th of a degree, i.e. a full circle equals 5760 (16 * 360). Positive values for the angles mean counter-clockwise while negative values mean the clockwise direction. Zero degrees is at the 3 o'clock position.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);
  int startAngle = 30 * 16;
  int spanAngle = 120 * 16;

  QPainter painter(this);
  painter.drawPie(rectangle, startAngle, spanAngle);





See also drawEllipse(), drawChord(), and Coordinate System.
*/
func (this *QPainter) DrawPie(rect qtcore.QRectF_ITF, a int, alen int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawPieERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:343
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPie(int, int, int, int, int, int)

/*
Draws a pie defined by the given rectangle, startAngle and spanAngle.

The pie is filled with the current brush().

The startAngle and spanAngle must be specified in 1/16th of a degree, i.e. a full circle equals 5760 (16 * 360). Positive values for the angles mean counter-clockwise while negative values mean the clockwise direction. Zero degrees is at the 3 o'clock position.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);
  int startAngle = 30 * 16;
  int spanAngle = 120 * 16;

  QPainter painter(this);
  painter.drawPie(rectangle, startAngle, spanAngle);





See also drawEllipse(), drawChord(), and Coordinate System.
*/
func (this *QPainter) DrawPie_1(x int, y int, w int, h int, a int, alen int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawPieEiiiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:344
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawPie(const QRect &, int, int)

/*
Draws a pie defined by the given rectangle, startAngle and spanAngle.

The pie is filled with the current brush().

The startAngle and spanAngle must be specified in 1/16th of a degree, i.e. a full circle equals 5760 (16 * 360). Positive values for the angles mean counter-clockwise while negative values mean the clockwise direction. Zero degrees is at the 3 o'clock position.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);
  int startAngle = 30 * 16;
  int spanAngle = 120 * 16;

  QPainter painter(this);
  painter.drawPie(rectangle, startAngle, spanAngle);





See also drawEllipse(), drawChord(), and Coordinate System.
*/
func (this *QPainter) DrawPie_2(arg0 qtcore.QRect_ITF, a int, alen int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter7drawPieERK5QRectii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:346
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawChord(const QRectF &, int, int)

/*
Draws the chord defined by the given rectangle, startAngle and spanAngle. The chord is filled with the current brush().

The startAngle and spanAngle must be specified in 1/16th of a degree, i.e. a full circle equals 5760 (16 * 360). Positive values for the angles mean counter-clockwise while negative values mean the clockwise direction. Zero degrees is at the 3 o'clock position.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);
  int startAngle = 30 * 16;
  int spanAngle = 120 * 16;

  QPainter painter(this);
  painter.drawChord(rect, startAngle, spanAngle);





See also drawArc(), drawPie(), and Coordinate System.
*/
func (this *QPainter) DrawChord(rect qtcore.QRectF_ITF, a int, alen int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawChordERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:347
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawChord(int, int, int, int, int, int)

/*
Draws the chord defined by the given rectangle, startAngle and spanAngle. The chord is filled with the current brush().

The startAngle and spanAngle must be specified in 1/16th of a degree, i.e. a full circle equals 5760 (16 * 360). Positive values for the angles mean counter-clockwise while negative values mean the clockwise direction. Zero degrees is at the 3 o'clock position.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);
  int startAngle = 30 * 16;
  int spanAngle = 120 * 16;

  QPainter painter(this);
  painter.drawChord(rect, startAngle, spanAngle);





See also drawArc(), drawPie(), and Coordinate System.
*/
func (this *QPainter) DrawChord_1(x int, y int, w int, h int, a int, alen int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawChordEiiiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:348
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawChord(const QRect &, int, int)

/*
Draws the chord defined by the given rectangle, startAngle and spanAngle. The chord is filled with the current brush().

The startAngle and spanAngle must be specified in 1/16th of a degree, i.e. a full circle equals 5760 (16 * 360). Positive values for the angles mean counter-clockwise while negative values mean the clockwise direction. Zero degrees is at the 3 o'clock position.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);
  int startAngle = 30 * 16;
  int spanAngle = 120 * 16;

  QPainter painter(this);
  painter.drawChord(rect, startAngle, spanAngle);





See also drawArc(), drawPie(), and Coordinate System.
*/
func (this *QPainter) DrawChord_2(arg0 qtcore.QRect_ITF, a int, alen int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawChordERK5QRectii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:350
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawRoundedRect(const QRectF &, qreal, qreal, Qt::SizeMode)

/*
Draws the given rectangle rect with rounded corners.

The xRadius and yRadius arguments specify the radii of the ellipses defining the corners of the rounded rectangle. When mode is Qt::RelativeSize, xRadius and yRadius are specified in percentage of half the rectangle's width and height respectively, and should be in the range 0.0 to 100.0.

A filled rectangle has a size of rect.size(). A stroked rectangle has a size of rect.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawRoundedRect(rectangle, 20.0, 15.0);





This function was introduced in  Qt 4.4.

See also drawRect() and QPen.
*/
func (this *QPainter) DrawRoundedRect(rect qtcore.QRectF_ITF, xRadius float64, yRadius float64, mode int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectERK6QRectFddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:350
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawRoundedRect(const QRectF &, qreal, qreal, Qt::SizeMode)

/*
Draws the given rectangle rect with rounded corners.

The xRadius and yRadius arguments specify the radii of the ellipses defining the corners of the rounded rectangle. When mode is Qt::RelativeSize, xRadius and yRadius are specified in percentage of half the rectangle's width and height respectively, and should be in the range 0.0 to 100.0.

A filled rectangle has a size of rect.size(). A stroked rectangle has a size of rect.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawRoundedRect(rectangle, 20.0, 15.0);





This function was introduced in  Qt 4.4.

See also drawRect() and QPen.
*/
func (this *QPainter) DrawRoundedRect__(rect qtcore.QRectF_ITF, xRadius float64, yRadius float64) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 3, Qt::SizeMode=Elaborated, Qt::SizeMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectERK6QRectFddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:352
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundedRect(int, int, int, int, qreal, qreal, Qt::SizeMode)

/*
Draws the given rectangle rect with rounded corners.

The xRadius and yRadius arguments specify the radii of the ellipses defining the corners of the rounded rectangle. When mode is Qt::RelativeSize, xRadius and yRadius are specified in percentage of half the rectangle's width and height respectively, and should be in the range 0.0 to 100.0.

A filled rectangle has a size of rect.size(). A stroked rectangle has a size of rect.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawRoundedRect(rectangle, 20.0, 15.0);





This function was introduced in  Qt 4.4.

See also drawRect() and QPen.
*/
func (this *QPainter) DrawRoundedRect_1(x int, y int, w int, h int, xRadius float64, yRadius float64, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectEiiiiddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:352
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundedRect(int, int, int, int, qreal, qreal, Qt::SizeMode)

/*
Draws the given rectangle rect with rounded corners.

The xRadius and yRadius arguments specify the radii of the ellipses defining the corners of the rounded rectangle. When mode is Qt::RelativeSize, xRadius and yRadius are specified in percentage of half the rectangle's width and height respectively, and should be in the range 0.0 to 100.0.

A filled rectangle has a size of rect.size(). A stroked rectangle has a size of rect.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawRoundedRect(rectangle, 20.0, 15.0);





This function was introduced in  Qt 4.4.

See also drawRect() and QPen.
*/
func (this *QPainter) DrawRoundedRect_1_(x int, y int, w int, h int, xRadius float64, yRadius float64) {
	// arg: 6, Qt::SizeMode=Elaborated, Qt::SizeMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectEiiiiddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:354
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundedRect(const QRect &, qreal, qreal, Qt::SizeMode)

/*
Draws the given rectangle rect with rounded corners.

The xRadius and yRadius arguments specify the radii of the ellipses defining the corners of the rounded rectangle. When mode is Qt::RelativeSize, xRadius and yRadius are specified in percentage of half the rectangle's width and height respectively, and should be in the range 0.0 to 100.0.

A filled rectangle has a size of rect.size(). A stroked rectangle has a size of rect.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawRoundedRect(rectangle, 20.0, 15.0);





This function was introduced in  Qt 4.4.

See also drawRect() and QPen.
*/
func (this *QPainter) DrawRoundedRect_2(rect qtcore.QRect_ITF, xRadius float64, yRadius float64, mode int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectERK5QRectddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:354
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundedRect(const QRect &, qreal, qreal, Qt::SizeMode)

/*
Draws the given rectangle rect with rounded corners.

The xRadius and yRadius arguments specify the radii of the ellipses defining the corners of the rounded rectangle. When mode is Qt::RelativeSize, xRadius and yRadius are specified in percentage of half the rectangle's width and height respectively, and should be in the range 0.0 to 100.0.

A filled rectangle has a size of rect.size(). A stroked rectangle has a size of rect.size() plus the pen width.



  QRectF rectangle(10.0, 20.0, 80.0, 60.0);

  QPainter painter(this);
  painter.drawRoundedRect(rectangle, 20.0, 15.0);





This function was introduced in  Qt 4.4.

See also drawRect() and QPen.
*/
func (this *QPainter) DrawRoundedRect_2_(rect qtcore.QRect_ITF, xRadius float64, yRadius float64) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	// arg: 3, Qt::SizeMode=Elaborated, Qt::SizeMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectERK5QRectddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:357
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawRoundRect(const QRectF &, int, int)

/*

 */
func (this *QPainter) DrawRoundRect(r qtcore.QRectF_ITF, xround int, yround int) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xround, yround)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:357
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawRoundRect(const QRectF &, int, int)

/*

 */
func (this *QPainter) DrawRoundRect__(r qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	xround := int(25)
	// arg: 2, int=Int, =Invalid,
	yround := int(25)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xround, yround)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:357
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawRoundRect(const QRectF &, int, int)

/*

 */
func (this *QPainter) DrawRoundRect__1(r qtcore.QRectF_ITF, xround int) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	yround := int(25)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xround, yround)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:358
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundRect(int, int, int, int, int, int)

/*

 */
func (this *QPainter) DrawRoundRect_1(x int, y int, w int, h int, arg4 int, arg5 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectEiiiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, arg4, arg5)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:358
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundRect(int, int, int, int, int, int)

/*

 */
func (this *QPainter) DrawRoundRect_1_(x int, y int, w int, h int) {
	// arg: 4, int=Int, =Invalid,
	arg4 := int(25)
	// arg: 5, int=Int, =Invalid,
	arg5 := int(25)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectEiiiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, arg4, arg5)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:358
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundRect(int, int, int, int, int, int)

/*

 */
func (this *QPainter) DrawRoundRect_1_1(x int, y int, w int, h int, arg4 int) {
	// arg: 5, int=Int, =Invalid,
	arg5 := int(25)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectEiiiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, arg4, arg5)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:359
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundRect(const QRect &, int, int)

/*

 */
func (this *QPainter) DrawRoundRect_2(r qtcore.QRect_ITF, xround int, yround int) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK5QRectii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xround, yround)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:359
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundRect(const QRect &, int, int)

/*

 */
func (this *QPainter) DrawRoundRect_2_(r qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	xround := int(25)
	// arg: 2, int=Int, =Invalid,
	yround := int(25)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK5QRectii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xround, yround)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:359
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawRoundRect(const QRect &, int, int)

/*

 */
func (this *QPainter) DrawRoundRect_2_1(r qtcore.QRect_ITF, xround int) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	yround := int(25)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK5QRectii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xround, yround)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:361
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawTiledPixmap(const QRectF &, const QPixmap &, const QPointF &)

/*
Draws a tiled pixmap, inside the given rectangle with its origin at the given position.

Calling drawTiledPixmap() is similar to calling drawPixmap() several times to fill (tile) an area with a pixmap, but is potentially much more efficient depending on the underlying window system.

See also drawPixmap().
*/
func (this *QPainter) DrawTiledPixmap(rect qtcore.QRectF_ITF, pm QPixmap_ITF, offset qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pm != nil && pm.QPixmap_PTR() != nil {
		convArg1 = pm.QPixmap_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if offset != nil && offset.QPointF_PTR() != nil {
		convArg2 = offset.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:361
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawTiledPixmap(const QRectF &, const QPixmap &, const QPointF &)

/*
Draws a tiled pixmap, inside the given rectangle with its origin at the given position.

Calling drawTiledPixmap() is similar to calling drawPixmap() several times to fill (tile) an area with a pixmap, but is potentially much more efficient depending on the underlying window system.

See also drawPixmap().
*/
func (this *QPainter) DrawTiledPixmap__(rect qtcore.QRectF_ITF, pm QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pm != nil && pm.QPixmap_PTR() != nil {
		convArg1 = pm.QPixmap_PTR().GetCthis()
	}
	// arg: 2, const QPointF &=LValueReference, QPointF=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:362
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawTiledPixmap(int, int, int, int, const QPixmap &, int, int)

/*
Draws a tiled pixmap, inside the given rectangle with its origin at the given position.

Calling drawTiledPixmap() is similar to calling drawPixmap() several times to fill (tile) an area with a pixmap, but is potentially much more efficient depending on the underlying window system.

See also drawPixmap().
*/
func (this *QPainter) DrawTiledPixmap_1(x int, y int, w int, h int, arg4 QPixmap_ITF, sx int, sy int) {
	var convArg4 unsafe.Pointer
	if arg4 != nil && arg4.QPixmap_PTR() != nil {
		convArg4 = arg4.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, sx, sy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:362
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawTiledPixmap(int, int, int, int, const QPixmap &, int, int)

/*
Draws a tiled pixmap, inside the given rectangle with its origin at the given position.

Calling drawTiledPixmap() is similar to calling drawPixmap() several times to fill (tile) an area with a pixmap, but is potentially much more efficient depending on the underlying window system.

See also drawPixmap().
*/
func (this *QPainter) DrawTiledPixmap_1_(x int, y int, w int, h int, arg4 QPixmap_ITF) {
	var convArg4 unsafe.Pointer
	if arg4 != nil && arg4.QPixmap_PTR() != nil {
		convArg4 = arg4.QPixmap_PTR().GetCthis()
	}
	// arg: 5, int=Int, =Invalid,
	sx := int(0)
	// arg: 6, int=Int, =Invalid,
	sy := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, sx, sy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:362
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawTiledPixmap(int, int, int, int, const QPixmap &, int, int)

/*
Draws a tiled pixmap, inside the given rectangle with its origin at the given position.

Calling drawTiledPixmap() is similar to calling drawPixmap() several times to fill (tile) an area with a pixmap, but is potentially much more efficient depending on the underlying window system.

See also drawPixmap().
*/
func (this *QPainter) DrawTiledPixmap_1_1(x int, y int, w int, h int, arg4 QPixmap_ITF, sx int) {
	var convArg4 unsafe.Pointer
	if arg4 != nil && arg4.QPixmap_PTR() != nil {
		convArg4 = arg4.QPixmap_PTR().GetCthis()
	}
	// arg: 6, int=Int, =Invalid,
	sy := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, sx, sy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:363
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawTiledPixmap(const QRect &, const QPixmap &, const QPoint &)

/*
Draws a tiled pixmap, inside the given rectangle with its origin at the given position.

Calling drawTiledPixmap() is similar to calling drawPixmap() several times to fill (tile) an area with a pixmap, but is potentially much more efficient depending on the underlying window system.

See also drawPixmap().
*/
func (this *QPainter) DrawTiledPixmap_2(arg0 qtcore.QRect_ITF, arg1 QPixmap_ITF, arg2 qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QPixmap_PTR() != nil {
		convArg1 = arg1.QPixmap_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if arg2 != nil && arg2.QPoint_PTR() != nil {
		convArg2 = arg2.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:363
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawTiledPixmap(const QRect &, const QPixmap &, const QPoint &)

/*
Draws a tiled pixmap, inside the given rectangle with its origin at the given position.

Calling drawTiledPixmap() is similar to calling drawPixmap() several times to fill (tile) an area with a pixmap, but is potentially much more efficient depending on the underlying window system.

See also drawPixmap().
*/
func (this *QPainter) DrawTiledPixmap_2_(arg0 qtcore.QRect_ITF, arg1 QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QPixmap_PTR() != nil {
		convArg1 = arg1.QPixmap_PTR().GetCthis()
	}
	// arg: 2, const QPoint &=LValueReference, QPoint=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:365
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPicture(const QPointF &, const QPicture &)

/*
Replays the given picture at the given point.

The QPicture class is a paint device that records and replays QPainter commands. A picture serializes the painter commands to an IO device in a platform-independent format. Everything that can be painted on a widget or pixmap can also be stored in a picture.

This function does exactly the same as QPicture::play() when called with point = QPoint(0, 0).



  QPicture picture;
  QPointF point(10.0, 20.0)
  picture.load("drawing.pic");

  QPainter painter(this);
  painter.drawPicture(0, 0, picture);





See also QPicture::play().
*/
func (this *QPainter) DrawPicture(p qtcore.QPointF_ITF, picture QPicture_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if picture != nil && picture.QPicture_PTR() != nil {
		convArg1 = picture.QPicture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPictureERK7QPointFRK8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:366
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPicture(int, int, const QPicture &)

/*
Replays the given picture at the given point.

The QPicture class is a paint device that records and replays QPainter commands. A picture serializes the painter commands to an IO device in a platform-independent format. Everything that can be painted on a widget or pixmap can also be stored in a picture.

This function does exactly the same as QPicture::play() when called with point = QPoint(0, 0).



  QPicture picture;
  QPointF point(10.0, 20.0)
  picture.load("drawing.pic");

  QPainter painter(this);
  painter.drawPicture(0, 0, picture);





See also QPicture::play().
*/
func (this *QPainter) DrawPicture_1(x int, y int, picture QPicture_ITF) {
	var convArg2 unsafe.Pointer
	if picture != nil && picture.QPicture_PTR() != nil {
		convArg2 = picture.QPicture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPictureEiiRK8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:367
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawPicture(const QPoint &, const QPicture &)

/*
Replays the given picture at the given point.

The QPicture class is a paint device that records and replays QPainter commands. A picture serializes the painter commands to an IO device in a platform-independent format. Everything that can be painted on a widget or pixmap can also be stored in a picture.

This function does exactly the same as QPicture::play() when called with point = QPoint(0, 0).



  QPicture picture;
  QPointF point(10.0, 20.0)
  picture.load("drawing.pic");

  QPainter painter(this);
  painter.drawPicture(0, 0, picture);





See also QPicture::play().
*/
func (this *QPainter) DrawPicture_2(p qtcore.QPoint_ITF, picture QPicture_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if picture != nil && picture.QPicture_PTR() != nil {
		convArg1 = picture.QPicture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter11drawPictureERK6QPointRK8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:370
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawPixmap(const QRectF &, const QPixmap &, const QRectF &)

/*
Draws the rectangular portion source of the given pixmap into the given target in the paint device.

Note: The pixmap is scaled to fit the rectangle, if both the pixmap and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QPixmap::devicePixelRatio().



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QPixmap pixmap(":myPixmap.png");

  QPainter(this);
  painter.drawPixmap(target, pixmap, source);





If pixmap is a QBitmap it is drawn with the bits that are "set" using the pens color. If backgroundMode is Qt::OpaqueMode, the "unset" bits are drawn using the color of the background brush; if backgroundMode is Qt::TransparentMode, the "unset" bits are transparent. Drawing bitmaps with gradient or texture colors is not supported.

See also drawImage() and QPixmap::devicePixelRatio().
*/
func (this *QPainter) DrawPixmap(targetRect qtcore.QRectF_ITF, pixmap QPixmap_ITF, sourceRect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if targetRect != nil && targetRect.QRectF_PTR() != nil {
		convArg0 = targetRect.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceRect != nil && sourceRect.QRectF_PTR() != nil {
		convArg2 = sourceRect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QRectFRK7QPixmapS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:371
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(const QRect &, const QPixmap &, const QRect &)

/*
Draws the rectangular portion source of the given pixmap into the given target in the paint device.

Note: The pixmap is scaled to fit the rectangle, if both the pixmap and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QPixmap::devicePixelRatio().



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QPixmap pixmap(":myPixmap.png");

  QPainter(this);
  painter.drawPixmap(target, pixmap, source);





If pixmap is a QBitmap it is drawn with the bits that are "set" using the pens color. If backgroundMode is Qt::OpaqueMode, the "unset" bits are drawn using the color of the background brush; if backgroundMode is Qt::TransparentMode, the "unset" bits are transparent. Drawing bitmaps with gradient or texture colors is not supported.

See also drawImage() and QPixmap::devicePixelRatio().
*/
func (this *QPainter) DrawPixmap_1(targetRect qtcore.QRect_ITF, pixmap QPixmap_ITF, sourceRect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if targetRect != nil && targetRect.QRect_PTR() != nil {
		convArg0 = targetRect.QRect_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceRect != nil && sourceRect.QRect_PTR() != nil {
		convArg2 = sourceRect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK5QRectRK7QPixmapS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:372
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(int, int, int, int, const QPixmap &, int, int, int, int)

/*
Draws the rectangular portion source of the given pixmap into the given target in the paint device.

Note: The pixmap is scaled to fit the rectangle, if both the pixmap and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QPixmap::devicePixelRatio().



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QPixmap pixmap(":myPixmap.png");

  QPainter(this);
  painter.drawPixmap(target, pixmap, source);





If pixmap is a QBitmap it is drawn with the bits that are "set" using the pens color. If backgroundMode is Qt::OpaqueMode, the "unset" bits are drawn using the color of the background brush; if backgroundMode is Qt::TransparentMode, the "unset" bits are transparent. Drawing bitmaps with gradient or texture colors is not supported.

See also drawImage() and QPixmap::devicePixelRatio().
*/
func (this *QPainter) DrawPixmap_2(x int, y int, w int, h int, pm QPixmap_ITF, sx int, sy int, sw int, sh int) {
	var convArg4 unsafe.Pointer
	if pm != nil && pm.QPixmap_PTR() != nil {
		convArg4 = pm.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiiiRK7QPixmapiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, sx, sy, sw, sh)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:374
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(int, int, const QPixmap &, int, int, int, int)

/*
Draws the rectangular portion source of the given pixmap into the given target in the paint device.

Note: The pixmap is scaled to fit the rectangle, if both the pixmap and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QPixmap::devicePixelRatio().



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QPixmap pixmap(":myPixmap.png");

  QPainter(this);
  painter.drawPixmap(target, pixmap, source);





If pixmap is a QBitmap it is drawn with the bits that are "set" using the pens color. If backgroundMode is Qt::OpaqueMode, the "unset" bits are drawn using the color of the background brush; if backgroundMode is Qt::TransparentMode, the "unset" bits are transparent. Drawing bitmaps with gradient or texture colors is not supported.

See also drawImage() and QPixmap::devicePixelRatio().
*/
func (this *QPainter) DrawPixmap_3(x int, y int, pm QPixmap_ITF, sx int, sy int, sw int, sh int) {
	var convArg2 unsafe.Pointer
	if pm != nil && pm.QPixmap_PTR() != nil {
		convArg2 = pm.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiRK7QPixmapiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, sx, sy, sw, sh)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:376
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(const QPointF &, const QPixmap &, const QRectF &)

/*
Draws the rectangular portion source of the given pixmap into the given target in the paint device.

Note: The pixmap is scaled to fit the rectangle, if both the pixmap and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QPixmap::devicePixelRatio().



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QPixmap pixmap(":myPixmap.png");

  QPainter(this);
  painter.drawPixmap(target, pixmap, source);





If pixmap is a QBitmap it is drawn with the bits that are "set" using the pens color. If backgroundMode is Qt::OpaqueMode, the "unset" bits are drawn using the color of the background brush; if backgroundMode is Qt::TransparentMode, the "unset" bits are transparent. Drawing bitmaps with gradient or texture colors is not supported.

See also drawImage() and QPixmap::devicePixelRatio().
*/
func (this *QPainter) DrawPixmap_4(p qtcore.QPointF_ITF, pm QPixmap_ITF, sr qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pm != nil && pm.QPixmap_PTR() != nil {
		convArg1 = pm.QPixmap_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sr != nil && sr.QRectF_PTR() != nil {
		convArg2 = sr.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmapRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:377
// index:5
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(const QPoint &, const QPixmap &, const QRect &)

/*
Draws the rectangular portion source of the given pixmap into the given target in the paint device.

Note: The pixmap is scaled to fit the rectangle, if both the pixmap and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QPixmap::devicePixelRatio().



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QPixmap pixmap(":myPixmap.png");

  QPainter(this);
  painter.drawPixmap(target, pixmap, source);





If pixmap is a QBitmap it is drawn with the bits that are "set" using the pens color. If backgroundMode is Qt::OpaqueMode, the "unset" bits are drawn using the color of the background brush; if backgroundMode is Qt::TransparentMode, the "unset" bits are transparent. Drawing bitmaps with gradient or texture colors is not supported.

See also drawImage() and QPixmap::devicePixelRatio().
*/
func (this *QPainter) DrawPixmap_5(p qtcore.QPoint_ITF, pm QPixmap_ITF, sr qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pm != nil && pm.QPixmap_PTR() != nil {
		convArg1 = pm.QPixmap_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sr != nil && sr.QRect_PTR() != nil {
		convArg2 = sr.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QPointRK7QPixmapRK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:378
// index:6
// Public Visibility=Default Availability=Available
// [-2] void drawPixmap(const QPointF &, const QPixmap &)

/*
Draws the rectangular portion source of the given pixmap into the given target in the paint device.

Note: The pixmap is scaled to fit the rectangle, if both the pixmap and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QPixmap::devicePixelRatio().



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QPixmap pixmap(":myPixmap.png");

  QPainter(this);
  painter.drawPixmap(target, pixmap, source);





If pixmap is a QBitmap it is drawn with the bits that are "set" using the pens color. If backgroundMode is Qt::OpaqueMode, the "unset" bits are drawn using the color of the background brush; if backgroundMode is Qt::TransparentMode, the "unset" bits are transparent. Drawing bitmaps with gradient or texture colors is not supported.

See also drawImage() and QPixmap::devicePixelRatio().
*/
func (this *QPainter) DrawPixmap_6(p qtcore.QPointF_ITF, pm QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pm != nil && pm.QPixmap_PTR() != nil {
		convArg1 = pm.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:379
// index:7
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(const QPoint &, const QPixmap &)

/*
Draws the rectangular portion source of the given pixmap into the given target in the paint device.

Note: The pixmap is scaled to fit the rectangle, if both the pixmap and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QPixmap::devicePixelRatio().



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QPixmap pixmap(":myPixmap.png");

  QPainter(this);
  painter.drawPixmap(target, pixmap, source);





If pixmap is a QBitmap it is drawn with the bits that are "set" using the pens color. If backgroundMode is Qt::OpaqueMode, the "unset" bits are drawn using the color of the background brush; if backgroundMode is Qt::TransparentMode, the "unset" bits are transparent. Drawing bitmaps with gradient or texture colors is not supported.

See also drawImage() and QPixmap::devicePixelRatio().
*/
func (this *QPainter) DrawPixmap_7(p qtcore.QPoint_ITF, pm QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pm != nil && pm.QPixmap_PTR() != nil {
		convArg1 = pm.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QPointRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:380
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(int, int, const QPixmap &)

/*
Draws the rectangular portion source of the given pixmap into the given target in the paint device.

Note: The pixmap is scaled to fit the rectangle, if both the pixmap and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QPixmap::devicePixelRatio().



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QPixmap pixmap(":myPixmap.png");

  QPainter(this);
  painter.drawPixmap(target, pixmap, source);





If pixmap is a QBitmap it is drawn with the bits that are "set" using the pens color. If backgroundMode is Qt::OpaqueMode, the "unset" bits are drawn using the color of the background brush; if backgroundMode is Qt::TransparentMode, the "unset" bits are transparent. Drawing bitmaps with gradient or texture colors is not supported.

See also drawImage() and QPixmap::devicePixelRatio().
*/
func (this *QPainter) DrawPixmap_8(x int, y int, pm QPixmap_ITF) {
	var convArg2 unsafe.Pointer
	if pm != nil && pm.QPixmap_PTR() != nil {
		convArg2 = pm.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:381
// index:9
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(const QRect &, const QPixmap &)

/*
Draws the rectangular portion source of the given pixmap into the given target in the paint device.

Note: The pixmap is scaled to fit the rectangle, if both the pixmap and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QPixmap::devicePixelRatio().



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QPixmap pixmap(":myPixmap.png");

  QPainter(this);
  painter.drawPixmap(target, pixmap, source);





If pixmap is a QBitmap it is drawn with the bits that are "set" using the pens color. If backgroundMode is Qt::OpaqueMode, the "unset" bits are drawn using the color of the background brush; if backgroundMode is Qt::TransparentMode, the "unset" bits are transparent. Drawing bitmaps with gradient or texture colors is not supported.

See also drawImage() and QPixmap::devicePixelRatio().
*/
func (this *QPainter) DrawPixmap_9(r qtcore.QRect_ITF, pm QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pm != nil && pm.QPixmap_PTR() != nil {
		convArg1 = pm.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK5QRectRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:382
// index:10
// Public inline Visibility=Default Availability=Available
// [-2] void drawPixmap(int, int, int, int, const QPixmap &)

/*
Draws the rectangular portion source of the given pixmap into the given target in the paint device.

Note: The pixmap is scaled to fit the rectangle, if both the pixmap and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QPixmap::devicePixelRatio().



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QPixmap pixmap(":myPixmap.png");

  QPainter(this);
  painter.drawPixmap(target, pixmap, source);





If pixmap is a QBitmap it is drawn with the bits that are "set" using the pens color. If backgroundMode is Qt::OpaqueMode, the "unset" bits are drawn using the color of the background brush; if backgroundMode is Qt::TransparentMode, the "unset" bits are transparent. Drawing bitmaps with gradient or texture colors is not supported.

See also drawImage() and QPixmap::devicePixelRatio().
*/
func (this *QPainter) DrawPixmap_10(x int, y int, w int, h int, pm QPixmap_ITF) {
	var convArg4 unsafe.Pointer
	if pm != nil && pm.QPixmap_PTR() != nil {
		convArg4 = pm.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiiiRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:387
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawImage(const QRectF &, const QImage &, const QRectF &, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage(targetRect qtcore.QRectF_ITF, image QImage_ITF, sourceRect qtcore.QRectF_ITF, flags int) {
	var convArg0 unsafe.Pointer
	if targetRect != nil && targetRect.QRectF_PTR() != nil {
		convArg0 = targetRect.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceRect != nil && sourceRect.QRectF_PTR() != nil {
		convArg2 = sourceRect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QRectFRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:387
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawImage(const QRectF &, const QImage &, const QRectF &, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage__(targetRect qtcore.QRectF_ITF, image QImage_ITF, sourceRect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if targetRect != nil && targetRect.QRectF_PTR() != nil {
		convArg0 = targetRect.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceRect != nil && sourceRect.QRectF_PTR() != nil {
		convArg2 = sourceRect.QRectF_PTR().GetCthis()
	}
	// arg: 3, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QRectFRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:389
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QRect &, const QImage &, const QRect &, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_1(targetRect qtcore.QRect_ITF, image QImage_ITF, sourceRect qtcore.QRect_ITF, flags int) {
	var convArg0 unsafe.Pointer
	if targetRect != nil && targetRect.QRect_PTR() != nil {
		convArg0 = targetRect.QRect_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceRect != nil && sourceRect.QRect_PTR() != nil {
		convArg2 = sourceRect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK5QRectRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:389
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QRect &, const QImage &, const QRect &, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_1_(targetRect qtcore.QRect_ITF, image QImage_ITF, sourceRect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if targetRect != nil && targetRect.QRect_PTR() != nil {
		convArg0 = targetRect.QRect_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceRect != nil && sourceRect.QRect_PTR() != nil {
		convArg2 = sourceRect.QRect_PTR().GetCthis()
	}
	// arg: 3, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK5QRectRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:391
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QPointF &, const QImage &, const QRectF &, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_2(p qtcore.QPointF_ITF, image QImage_ITF, sr qtcore.QRectF_ITF, flags int) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sr != nil && sr.QRectF_PTR() != nil {
		convArg2 = sr.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK7QPointFRK6QImageRK6QRectF6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:391
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QPointF &, const QImage &, const QRectF &, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_2_(p qtcore.QPointF_ITF, image QImage_ITF, sr qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sr != nil && sr.QRectF_PTR() != nil {
		convArg2 = sr.QRectF_PTR().GetCthis()
	}
	// arg: 3, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK7QPointFRK6QImageRK6QRectF6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:393
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QPoint &, const QImage &, const QRect &, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_3(p qtcore.QPoint_ITF, image QImage_ITF, sr qtcore.QRect_ITF, flags int) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sr != nil && sr.QRect_PTR() != nil {
		convArg2 = sr.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QPointRK6QImageRK5QRect6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:393
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QPoint &, const QImage &, const QRect &, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_3_(p qtcore.QPoint_ITF, image QImage_ITF, sr qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sr != nil && sr.QRect_PTR() != nil {
		convArg2 = sr.QRect_PTR().GetCthis()
	}
	// arg: 3, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QPointRK6QImageRK5QRect6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:395
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QRectF &, const QImage &)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_4(r qtcore.QRectF_ITF, image QImage_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QRectFRK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:396
// index:5
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QRect &, const QImage &)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_5(r qtcore.QRect_ITF, image QImage_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK5QRectRK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:397
// index:6
// Public Visibility=Default Availability=Available
// [-2] void drawImage(const QPointF &, const QImage &)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_6(p qtcore.QPointF_ITF, image QImage_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK7QPointFRK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:398
// index:7
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(const QPoint &, const QImage &)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_7(p qtcore.QPoint_ITF, image QImage_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg1 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QPointRK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:399
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(int, int, const QImage &, int, int, int, int, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_8(x int, y int, image QImage_ITF, sx int, sy int, sw int, sh int, flags int) {
	var convArg2 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg2 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageEiiRK6QImageiiii6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, sx, sy, sw, sh, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:399
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(int, int, const QImage &, int, int, int, int, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_8_(x int, y int, image QImage_ITF) {
	var convArg2 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg2 = image.QImage_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid,
	sx := int(0)
	// arg: 4, int=Int, =Invalid,
	sy := int(0)
	// arg: 5, int=Int, =Invalid,
	sw := int(-1)
	// arg: 6, int=Int, =Invalid,
	sh := int(-1)
	// arg: 7, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageEiiRK6QImageiiii6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, sx, sy, sw, sh, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:399
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(int, int, const QImage &, int, int, int, int, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_8_1(x int, y int, image QImage_ITF, sx int) {
	var convArg2 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg2 = image.QImage_PTR().GetCthis()
	}
	// arg: 4, int=Int, =Invalid,
	sy := int(0)
	// arg: 5, int=Int, =Invalid,
	sw := int(-1)
	// arg: 6, int=Int, =Invalid,
	sh := int(-1)
	// arg: 7, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageEiiRK6QImageiiii6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, sx, sy, sw, sh, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:399
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(int, int, const QImage &, int, int, int, int, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_8_2(x int, y int, image QImage_ITF, sx int, sy int) {
	var convArg2 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg2 = image.QImage_PTR().GetCthis()
	}
	// arg: 5, int=Int, =Invalid,
	sw := int(-1)
	// arg: 6, int=Int, =Invalid,
	sh := int(-1)
	// arg: 7, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageEiiRK6QImageiiii6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, sx, sy, sw, sh, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:399
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(int, int, const QImage &, int, int, int, int, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_8_3(x int, y int, image QImage_ITF, sx int, sy int, sw int) {
	var convArg2 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg2 = image.QImage_PTR().GetCthis()
	}
	// arg: 6, int=Int, =Invalid,
	sh := int(-1)
	// arg: 7, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageEiiRK6QImageiiii6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, sx, sy, sw, sh, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:399
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void drawImage(int, int, const QImage &, int, int, int, int, Qt::ImageConversionFlags)

/*
Draws the rectangular portion source of the given image into the target rectangle in the paint device.

Note: The image is scaled to fit the rectangle, if both the image and rectangle size disagree.

Note: See Drawing High Resolution Versions of Pixmaps and Images on how this is affected by QImage::devicePixelRatio().

If the image needs to be modified to fit in a lower-resolution result (e.g. converting from 32-bit to 8-bit), use the flags to specify how you would prefer this to happen.



  QRectF target(10.0, 20.0, 80.0, 60.0);
  QRectF source(0.0, 0.0, 70.0, 40.0);
  QImage image(":/images/myImage.png");

  QPainter painter(this);
  painter.drawImage(target, image, source);





See also drawPixmap() and QImage::devicePixelRatio().
*/
func (this *QPainter) DrawImage_8_4(x int, y int, image QImage_ITF, sx int, sy int, sw int, sh int) {
	var convArg2 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg2 = image.QImage_PTR().GetCthis()
	}
	// arg: 7, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9drawImageEiiRK6QImageiiii6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, sx, sy, sw, sh, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:402
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayoutDirection(Qt::LayoutDirection)

/*
Sets the layout direction used by the painter when drawing text, to the specified direction.

The default is Qt::LayoutDirectionAuto, which will implicitly determine the direction from the text drawn.

See also QTextOption::setTextDirection(), layoutDirection(), drawText(), and Settings.
*/
func (this *QPainter) SetLayoutDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter18setLayoutDirectionEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:403
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection layoutDirection() const

/*
Returns the layout direction used by the painter when drawing text.

See also QTextOption::textDirection(), setLayoutDirection(), drawText(), and Settings.
*/
func (this *QPainter) LayoutDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter15layoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:406
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawGlyphRun(const QPointF &, const QGlyphRun &)

/*
Draws the glyphs represented by glyphs at position. The position gives the edge of the baseline for the string of glyphs. The glyphs will be retrieved from the font selected on glyphs and at offsets given by the positions in glyphs.

This function was introduced in  Qt 4.8.

See also QGlyphRun::setRawFont(), QGlyphRun::setPositions(), and QGlyphRun::setGlyphIndexes().
*/
func (this *QPainter) DrawGlyphRun(position qtcore.QPointF_ITF, glyphRun QGlyphRun_ITF) {
	var convArg0 unsafe.Pointer
	if position != nil && position.QPointF_PTR() != nil {
		convArg0 = position.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if glyphRun != nil && glyphRun.QGlyphRun_PTR() != nil {
		convArg1 = glyphRun.QGlyphRun_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawGlyphRunERK7QPointFRK9QGlyphRun", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:409
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawStaticText(const QPointF &, const QStaticText &)

/*
Draws the given staticText at the given topLeftPosition.

The text will be drawn using the font and the transformation set on the painter. If the font and/or transformation set on the painter are different from the ones used to initialize the layout of the QStaticText, then the layout will have to be recalculated. Use QStaticText::prepare() to initialize staticText with the font and transformation with which it will later be drawn.

If topLeftPosition is not the same as when staticText was initialized, or when it was last drawn, then there will be a slight overhead when translating the text to its new position.

Note: If the painter's transformation is not affine, then staticText will be drawn using regular calls to drawText(), losing any potential for performance improvement.

Note: The y-position is used as the top of the font.

This function was introduced in  Qt 4.7.

See also QStaticText.
*/
func (this *QPainter) DrawStaticText(topLeftPosition qtcore.QPointF_ITF, staticText QStaticText_ITF) {
	var convArg0 unsafe.Pointer
	if topLeftPosition != nil && topLeftPosition.QPointF_PTR() != nil {
		convArg0 = topLeftPosition.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if staticText != nil && staticText.QStaticText_PTR() != nil {
		convArg1 = staticText.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextERK7QPointFRK11QStaticText", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:410
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawStaticText(const QPoint &, const QStaticText &)

/*
Draws the given staticText at the given topLeftPosition.

The text will be drawn using the font and the transformation set on the painter. If the font and/or transformation set on the painter are different from the ones used to initialize the layout of the QStaticText, then the layout will have to be recalculated. Use QStaticText::prepare() to initialize staticText with the font and transformation with which it will later be drawn.

If topLeftPosition is not the same as when staticText was initialized, or when it was last drawn, then there will be a slight overhead when translating the text to its new position.

Note: If the painter's transformation is not affine, then staticText will be drawn using regular calls to drawText(), losing any potential for performance improvement.

Note: The y-position is used as the top of the font.

This function was introduced in  Qt 4.7.

See also QStaticText.
*/
func (this *QPainter) DrawStaticText_1(topLeftPosition qtcore.QPoint_ITF, staticText QStaticText_ITF) {
	var convArg0 unsafe.Pointer
	if topLeftPosition != nil && topLeftPosition.QPoint_PTR() != nil {
		convArg0 = topLeftPosition.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if staticText != nil && staticText.QStaticText_PTR() != nil {
		convArg1 = staticText.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextERK6QPointRK11QStaticText", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:411
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawStaticText(int, int, const QStaticText &)

/*
Draws the given staticText at the given topLeftPosition.

The text will be drawn using the font and the transformation set on the painter. If the font and/or transformation set on the painter are different from the ones used to initialize the layout of the QStaticText, then the layout will have to be recalculated. Use QStaticText::prepare() to initialize staticText with the font and transformation with which it will later be drawn.

If topLeftPosition is not the same as when staticText was initialized, or when it was last drawn, then there will be a slight overhead when translating the text to its new position.

Note: If the painter's transformation is not affine, then staticText will be drawn using regular calls to drawText(), losing any potential for performance improvement.

Note: The y-position is used as the top of the font.

This function was introduced in  Qt 4.7.

See also QStaticText.
*/
func (this *QPainter) DrawStaticText_2(left int, top int, staticText QStaticText_ITF) {
	var convArg2 unsafe.Pointer
	if staticText != nil && staticText.QStaticText_PTR() != nil {
		convArg2 = staticText.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextEiiRK11QStaticText", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:413
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QPointF &, const QString &)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText(p qtcore.QPointF_ITF, s string) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(s)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK7QPointFRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:414
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawText(const QPoint &, const QString &)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText_1(p qtcore.QPoint_ITF, s string) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(s)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QPointRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:415
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawText(int, int, const QString &)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText_2(x int, y int, s string) {
	var tmpArg2 = qtcore.NewQString_5(s)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextEiiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:417
// index:3
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QPointF &, const QString &, int, int)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText_3(p qtcore.QPointF_ITF, str string, tf int, justificationPadding int) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(str)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK7QPointFRK7QStringii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, tf, justificationPadding)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:419
// index:4
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QRectF &, int, const QString &, QRectF *)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText_4(r qtcore.QRectF_ITF, flags int, text string, br qtcore.QRectF_ITF /*777 QRectF **/) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 unsafe.Pointer
	if br != nil && br.QRectF_PTR() != nil {
		convArg3 = br.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:419
// index:4
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QRectF &, int, const QString &, QRectF *)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText_4_(r qtcore.QRectF_ITF, flags int, text string) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QRectF *=Pointer, QRectF=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:420
// index:5
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QRect &, int, const QString &, QRect *)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText_5(r qtcore.QRect_ITF, flags int, text string, br qtcore.QRect_ITF /*777 QRect **/) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 unsafe.Pointer
	if br != nil && br.QRect_PTR() != nil {
		convArg3 = br.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:420
// index:5
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QRect &, int, const QString &, QRect *)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText_5_(r qtcore.QRect_ITF, flags int, text string) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QRect *=Pointer, QRect=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:421
// index:6
// Public inline Visibility=Default Availability=Available
// [-2] void drawText(int, int, int, int, int, const QString &, QRect *)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText_6(x int, y int, w int, h int, flags int, text string, br qtcore.QRect_ITF /*777 QRect **/) {
	var tmpArg5 = qtcore.NewQString_5(text)
	var convArg5 = tmpArg5.GetCthis()
	var convArg6 unsafe.Pointer
	if br != nil && br.QRect_PTR() != nil {
		convArg6 = br.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, flags, convArg5, convArg6)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:421
// index:6
// Public inline Visibility=Default Availability=Available
// [-2] void drawText(int, int, int, int, int, const QString &, QRect *)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText_6_(x int, y int, w int, h int, flags int, text string) {
	var tmpArg5 = qtcore.NewQString_5(text)
	var convArg5 = tmpArg5.GetCthis()
	// arg: 6, QRect *=Pointer, QRect=Record,
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, flags, convArg5, convArg6)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:423
// index:7
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QRectF &, const QString &, const QTextOption &)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText_7(r qtcore.QRectF_ITF, text string, o QTextOption_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if o != nil && o.QTextOption_PTR() != nil {
		convArg2 = o.QTextOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:423
// index:7
// Public Visibility=Default Availability=Available
// [-2] void drawText(const QRectF &, const QString &, const QTextOption &)

/*
Draws the given text with the currently defined text direction, beginning at the given position.

This function does not handle the newline character (\n), as it cannot break text into multiple lines, and it cannot display the newline character. Use the QPainter::drawText() overload that takes a rectangle instead if you want to draw multiple lines of text with the newline character, or if you want the text to be wrapped.

By default, QPainter draws text anti-aliased.

Note: The y-position is used as the baseline of the font.

See also setFont() and setPen().
*/
func (this *QPainter) DrawText_7_(r qtcore.QRectF_ITF, text string) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QTextOption &=LValueReference, QTextOption=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:425
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QRectF &, int, const QString &)

/*
Returns the bounding rectangle of the text as it will appear when drawn inside the given rectangle with the specified flags using the currently set font(); i.e the function tells you where the drawText() function will draw when given the same arguments.

If the text does not fit within the given rectangle using the specified flags, the function returns the required rectangle.

The flags argument is a bitwise OR of the following flags:


Qt::AlignLeft
Qt::AlignRight
Qt::AlignHCenter
Qt::AlignTop
Qt::AlignBottom
Qt::AlignVCenter
Qt::AlignCenter
Qt::TextSingleLine
Qt::TextExpandTabs
Qt::TextShowMnemonic
Qt::TextWordWrap
Qt::TextIncludeTrailingSpaces


If several of the horizontal or several of the vertical alignment flags are set, the resulting alignment is undefined.

See also drawText(), Qt::Alignment, and Qt::TextFlag.
*/
func (this *QPainter) BoundingRect(rect qtcore.QRectF_ITF, flags int, text string) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK6QRectFiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:426
// index:1
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect(const QRect &, int, const QString &)

/*
Returns the bounding rectangle of the text as it will appear when drawn inside the given rectangle with the specified flags using the currently set font(); i.e the function tells you where the drawText() function will draw when given the same arguments.

If the text does not fit within the given rectangle using the specified flags, the function returns the required rectangle.

The flags argument is a bitwise OR of the following flags:


Qt::AlignLeft
Qt::AlignRight
Qt::AlignHCenter
Qt::AlignTop
Qt::AlignBottom
Qt::AlignVCenter
Qt::AlignCenter
Qt::TextSingleLine
Qt::TextExpandTabs
Qt::TextShowMnemonic
Qt::TextWordWrap
Qt::TextIncludeTrailingSpaces


If several of the horizontal or several of the vertical alignment flags are set, the resulting alignment is undefined.

See also drawText(), Qt::Alignment, and Qt::TextFlag.
*/
func (this *QPainter) BoundingRect_1(rect qtcore.QRect_ITF, flags int, text string) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK5QRectiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:427
// index:2
// Public inline Visibility=Default Availability=Available
// [16] QRect boundingRect(int, int, int, int, int, const QString &)

/*
Returns the bounding rectangle of the text as it will appear when drawn inside the given rectangle with the specified flags using the currently set font(); i.e the function tells you where the drawText() function will draw when given the same arguments.

If the text does not fit within the given rectangle using the specified flags, the function returns the required rectangle.

The flags argument is a bitwise OR of the following flags:


Qt::AlignLeft
Qt::AlignRight
Qt::AlignHCenter
Qt::AlignTop
Qt::AlignBottom
Qt::AlignVCenter
Qt::AlignCenter
Qt::TextSingleLine
Qt::TextExpandTabs
Qt::TextShowMnemonic
Qt::TextWordWrap
Qt::TextIncludeTrailingSpaces


If several of the horizontal or several of the vertical alignment flags are set, the resulting alignment is undefined.

See also drawText(), Qt::Alignment, and Qt::TextFlag.
*/
func (this *QPainter) BoundingRect_2(x int, y int, w int, h int, flags int, text string) *qtcore.QRect /*123*/ {
	var tmpArg5 = qtcore.NewQString_5(text)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12boundingRectEiiiiiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, flags, convArg5)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:429
// index:3
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QRectF &, const QString &, const QTextOption &)

/*
Returns the bounding rectangle of the text as it will appear when drawn inside the given rectangle with the specified flags using the currently set font(); i.e the function tells you where the drawText() function will draw when given the same arguments.

If the text does not fit within the given rectangle using the specified flags, the function returns the required rectangle.

The flags argument is a bitwise OR of the following flags:


Qt::AlignLeft
Qt::AlignRight
Qt::AlignHCenter
Qt::AlignTop
Qt::AlignBottom
Qt::AlignVCenter
Qt::AlignCenter
Qt::TextSingleLine
Qt::TextExpandTabs
Qt::TextShowMnemonic
Qt::TextWordWrap
Qt::TextIncludeTrailingSpaces


If several of the horizontal or several of the vertical alignment flags are set, the resulting alignment is undefined.

See also drawText(), Qt::Alignment, and Qt::TextFlag.
*/
func (this *QPainter) BoundingRect_3(rect qtcore.QRectF_ITF, text string, o QTextOption_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if o != nil && o.QTextOption_PTR() != nil {
		convArg2 = o.QTextOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:429
// index:3
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QRectF &, const QString &, const QTextOption &)

/*
Returns the bounding rectangle of the text as it will appear when drawn inside the given rectangle with the specified flags using the currently set font(); i.e the function tells you where the drawText() function will draw when given the same arguments.

If the text does not fit within the given rectangle using the specified flags, the function returns the required rectangle.

The flags argument is a bitwise OR of the following flags:


Qt::AlignLeft
Qt::AlignRight
Qt::AlignHCenter
Qt::AlignTop
Qt::AlignBottom
Qt::AlignVCenter
Qt::AlignCenter
Qt::TextSingleLine
Qt::TextExpandTabs
Qt::TextShowMnemonic
Qt::TextWordWrap
Qt::TextIncludeTrailingSpaces


If several of the horizontal or several of the vertical alignment flags are set, the resulting alignment is undefined.

See also drawText(), Qt::Alignment, and Qt::TextFlag.
*/
func (this *QPainter) BoundingRect_3_(rect qtcore.QRectF_ITF, text string) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QTextOption &=LValueReference, QTextOption=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:431
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawTextItem(const QPointF &, const QTextItem &)

/*

 */
func (this *QPainter) DrawTextItem(p qtcore.QPointF_ITF, ti QTextItem_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if ti != nil && ti.QTextItem_PTR() != nil {
		convArg1 = ti.QTextItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawTextItemERK7QPointFRK9QTextItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:432
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void drawTextItem(int, int, const QTextItem &)

/*

 */
func (this *QPainter) DrawTextItem_1(x int, y int, ti QTextItem_ITF) {
	var convArg2 unsafe.Pointer
	if ti != nil && ti.QTextItem_PTR() != nil {
		convArg2 = ti.QTextItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawTextItemEiiRK9QTextItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:433
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void drawTextItem(const QPoint &, const QTextItem &)

/*

 */
func (this *QPainter) DrawTextItem_2(p qtcore.QPoint_ITF, ti QTextItem_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if ti != nil && ti.QTextItem_PTR() != nil {
		convArg1 = ti.QTextItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter12drawTextItemERK6QPointRK9QTextItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:435
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fillRect(const QRectF &, const QBrush &)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect(arg0 qtcore.QRectF_ITF, arg1 QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRectF_PTR() != nil {
		convArg0 = arg0.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QBrush_PTR() != nil {
		convArg1 = arg1.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:436
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(int, int, int, int, const QBrush &)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect_1(x int, y int, w int, h int, arg4 QBrush_ITF) {
	var convArg4 unsafe.Pointer
	if arg4 != nil && arg4.QBrush_PTR() != nil {
		convArg4 = arg4.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:437
// index:2
// Public Visibility=Default Availability=Available
// [-2] void fillRect(const QRect &, const QBrush &)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect_2(arg0 qtcore.QRect_ITF, arg1 QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QBrush_PTR() != nil {
		convArg1 = arg1.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:439
// index:3
// Public Visibility=Default Availability=Available
// [-2] void fillRect(const QRectF &, const QColor &)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect_3(arg0 qtcore.QRectF_ITF, color QColor_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRectF_PTR() != nil {
		convArg0 = arg0.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg1 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:440
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(int, int, int, int, const QColor &)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect_4(x int, y int, w int, h int, color QColor_ITF) {
	var convArg4 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg4 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:441
// index:5
// Public Visibility=Default Availability=Available
// [-2] void fillRect(const QRect &, const QColor &)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect_5(arg0 qtcore.QRect_ITF, color QColor_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg1 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:443
// index:6
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(int, int, int, int, Qt::GlobalColor)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect_6(x int, y int, w int, h int, c int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiN2Qt11GlobalColorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:444
// index:7
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(const QRect &, Qt::GlobalColor)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect_7(r qtcore.QRect_ITF, c int) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectN2Qt11GlobalColorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:445
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(const QRectF &, Qt::GlobalColor)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect_8(r qtcore.QRectF_ITF, c int) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFN2Qt11GlobalColorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:447
// index:9
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(int, int, int, int, Qt::BrushStyle)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect_9(x int, y int, w int, h int, style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:448
// index:10
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(const QRect &, Qt::BrushStyle)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect_10(r qtcore.QRect_ITF, style int) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:449
// index:11
// Public inline Visibility=Default Availability=Available
// [-2] void fillRect(const QRectF &, Qt::BrushStyle)

/*
Fills the given rectangle with the brush specified.

Alternatively, you can specify a QColor instead of a QBrush; the QBrush constructor (taking a QColor argument) will automatically create a solid pattern brush.

See also drawRect().
*/
func (this *QPainter) FillRect_11(r qtcore.QRectF_ITF, style int) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFN2Qt10BrushStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:451
// index:0
// Public Visibility=Default Availability=Available
// [-2] void eraseRect(const QRectF &)

/*
Erases the area inside the given rectangle. Equivalent to calling


  fillRect(rectangle, background()).



See also fillRect().
*/
func (this *QPainter) EraseRect(arg0 qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRectF_PTR() != nil {
		convArg0 = arg0.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9eraseRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:452
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void eraseRect(int, int, int, int)

/*
Erases the area inside the given rectangle. Equivalent to calling


  fillRect(rectangle, background()).



See also fillRect().
*/
func (this *QPainter) EraseRect_1(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9eraseRectEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:453
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void eraseRect(const QRect &)

/*
Erases the area inside the given rectangle. Equivalent to calling


  fillRect(rectangle, background()).



See also fillRect().
*/
func (this *QPainter) EraseRect_2(arg0 qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter9eraseRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:455
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHint(QPainter::RenderHint, bool)

/*
Sets the given render hint on the painter if on is true; otherwise clears the render hint.

See also setRenderHints(), renderHints(), and Rendering Quality.
*/
func (this *QPainter) SetRenderHint(hint int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13setRenderHintENS_10RenderHintEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:455
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHint(QPainter::RenderHint, bool)

/*
Sets the given render hint on the painter if on is true; otherwise clears the render hint.

See also setRenderHints(), renderHints(), and Rendering Quality.
*/
func (this *QPainter) SetRenderHint__(hint int) {
	// arg: 1, bool=Bool, =Invalid,
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13setRenderHintENS_10RenderHintEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:456
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHints(QPainter::RenderHints, bool)

/*
Sets the given render hints on the painter if on is true; otherwise clears the render hints.

This function was introduced in  Qt 4.2.

See also setRenderHint(), renderHints(), and Rendering Quality.
*/
func (this *QPainter) SetRenderHints(hints int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setRenderHintsE6QFlagsINS_10RenderHintEEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:456
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHints(QPainter::RenderHints, bool)

/*
Sets the given render hints on the painter if on is true; otherwise clears the render hints.

This function was introduced in  Qt 4.2.

See also setRenderHint(), renderHints(), and Rendering Quality.
*/
func (this *QPainter) SetRenderHints__(hints int) {
	// arg: 1, bool=Bool, =Invalid,
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter14setRenderHintsE6QFlagsINS_10RenderHintEEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:457
// index:0
// Public Visibility=Default Availability=Available
// [4] QPainter::RenderHints renderHints() const

/*
Returns a flag that specifies the rendering hints that are set for this painter.

See also setRenderHints(), testRenderHint(), and Rendering Quality.
*/
func (this *QPainter) RenderHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11renderHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:458
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool testRenderHint(QPainter::RenderHint) const

/*
Returns true if hint is set; otherwise returns false.

This function was introduced in  Qt 4.3.

See also renderHints() and setRenderHint().
*/
func (this *QPainter) TestRenderHint(hint int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter14testRenderHintENS_10RenderHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:460
// index:0
// Public Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine() const

/*
Returns the paint engine that the painter is currently operating on if the painter is active; otherwise 0.

See also isActive().
*/
func (this *QPainter) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPainter11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpainter.h:462
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setRedirected(const QPaintDevice *, QPaintDevice *, const QPoint &)

/*

 */
func (this *QPainter) SetRedirected(device QPaintDevice_ITF /*777 const QPaintDevice **/, replacement QPaintDevice_ITF /*777 QPaintDevice **/, offset qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QPaintDevice_PTR() != nil {
		convArg0 = device.QPaintDevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if replacement != nil && replacement.QPaintDevice_PTR() != nil {
		convArg1 = replacement.QPaintDevice_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg2 = offset.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13setRedirectedEPK12QPaintDevicePS0_RK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}
func QPainter_SetRedirected(device QPaintDevice_ITF /*777 const QPaintDevice **/, replacement QPaintDevice_ITF /*777 QPaintDevice **/, offset qtcore.QPoint_ITF) {
	var nilthis *QPainter
	nilthis.SetRedirected(device, replacement, offset)
}

// /usr/include/qt/QtGui/qpainter.h:462
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setRedirected(const QPaintDevice *, QPaintDevice *, const QPoint &)

/*

 */
func (this *QPainter) SetRedirected__(device QPaintDevice_ITF /*777 const QPaintDevice **/, replacement QPaintDevice_ITF /*777 QPaintDevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QPaintDevice_PTR() != nil {
		convArg0 = device.QPaintDevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if replacement != nil && replacement.QPaintDevice_PTR() != nil {
		convArg1 = replacement.QPaintDevice_PTR().GetCthis()
	}
	// arg: 2, const QPoint &=LValueReference, QPoint=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter13setRedirectedEPK12QPaintDevicePS0_RK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:464
// index:0
// Public static Visibility=Default Availability=Available
// [8] QPaintDevice * redirected(const QPaintDevice *, QPoint *)

/*

 */
func (this *QPainter) Redirected(device QPaintDevice_ITF /*777 const QPaintDevice **/, offset qtcore.QPoint_ITF /*777 QPoint **/) *QPaintDevice /*777 QPaintDevice **/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QPaintDevice_PTR() != nil {
		convArg0 = device.QPaintDevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg1 = offset.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10redirectedEPK12QPaintDeviceP6QPoint", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QPainter_Redirected(device QPaintDevice_ITF /*777 const QPaintDevice **/, offset qtcore.QPoint_ITF /*777 QPoint **/) *QPaintDevice /*777 QPaintDevice **/ {
	var nilthis *QPainter
	rv := nilthis.Redirected(device, offset)
	return rv
}

// /usr/include/qt/QtGui/qpainter.h:464
// index:0
// Public static Visibility=Default Availability=Available
// [8] QPaintDevice * redirected(const QPaintDevice *, QPoint *)

/*

 */
func (this *QPainter) Redirected__(device QPaintDevice_ITF /*777 const QPaintDevice **/) *QPaintDevice /*777 QPaintDevice **/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QPaintDevice_PTR() != nil {
		convArg0 = device.QPaintDevice_PTR().GetCthis()
	}
	// arg: 1, QPoint *=Pointer, QPoint=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter10redirectedEPK12QPaintDeviceP6QPoint", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpainter.h:465
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void restoreRedirected(const QPaintDevice *)

/*

 */
func (this *QPainter) RestoreRedirected(device QPaintDevice_ITF /*777 const QPaintDevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QPaintDevice_PTR() != nil {
		convArg0 = device.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17restoreRedirectedEPK12QPaintDevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QPainter_RestoreRedirected(device QPaintDevice_ITF /*777 const QPaintDevice **/) {
	var nilthis *QPainter
	nilthis.RestoreRedirected(device)
}

// /usr/include/qt/QtGui/qpainter.h:467
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginNativePainting()

/*
Flushes the painting pipeline and prepares for the user issuing commands directly to the underlying graphics context. Must be followed by a call to endNativePainting().

Note that only the states the underlying paint engine changes will be reset to their respective default states. The states we reset may change from release to release. The following states are currently reset in the OpenGL 2 engine:


blending is disabled
the depth, stencil and scissor tests are disabled
the active texture unit is reset to 0
the depth mask, depth function and the clear depth are reset to their default values
the stencil mask, stencil operation and stencil function are reset to their default values
the current color is reset to solid white


If, for example, the OpenGL polygon mode is changed by the user inside a beginNativePaint()/endNativePainting() block, it will not be reset to the default state by endNativePainting(). Here is an example that shows intermixing of painter commands and raw OpenGL commands:


  QPainter painter(this);
  painter.fillRect(0, 0, 128, 128, Qt::green);
  painter.beginNativePainting();

  glEnable(GL_SCISSOR_TEST);
  glScissor(0, 0, 64, 64);

  glClearColor(1, 0, 0, 1);
  glClear(GL_COLOR_BUFFER_BIT);

  glDisable(GL_SCISSOR_TEST);

  painter.endNativePainting();



This function was introduced in  Qt 4.6.

See also endNativePainting().
*/
func (this *QPainter) BeginNativePainting() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter19beginNativePaintingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:468
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endNativePainting()

/*
Restores the painter after manually issuing native painting commands. Lets the painter restore any native state that it relies on before calling any other painter commands.

This function was introduced in  Qt 4.6.

See also beginNativePainting().
*/
func (this *QPainter) EndNativePainting() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPainter17endNativePaintingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QPainter__RenderHint = int

//
const QPainter__Antialiasing QPainter__RenderHint = 1

//
const QPainter__TextAntialiasing QPainter__RenderHint = 2

//
const QPainter__SmoothPixmapTransform QPainter__RenderHint = 4

//
const QPainter__HighQualityAntialiasing QPainter__RenderHint = 8

//
const QPainter__NonCosmeticDefaultPen QPainter__RenderHint = 16

//
const QPainter__Qt4CompatiblePainting QPainter__RenderHint = 32

/*


 */
type QPainter__PixmapFragmentHint = int

//
const QPainter__OpaqueHint QPainter__PixmapFragmentHint = 1

/*
Defines the modes supported for digital image compositing. Composition modes are used to specify how the pixels in one image, the source, are merged with the pixel in another image, the destination.

Please note that the bitwise raster operation modes, denoted with a RasterOp prefix, are only natively supported in the X11 and raster paint engines. This means that the only way to utilize these modes on the Mac is via a QImage. The RasterOp denoted blend modes are not supported for pens and brushes with alpha components. Also, turning on the QPainter::Antialiasing render hint will effectively disable the RasterOp modes.





The most common type is SourceOver (often referred to as just alpha blending) where the source pixel is blended on top of the destination pixel in such a way that the alpha component of the source defines the translucency of the pixel.

Several composition modes require an alpha channel in the source or target images to have an effect. For optimal performance the image format Format_ARGB32_Premultiplied is preferred.

When a composition mode is set it applies to all painting operator, pens, brushes, gradients and pixmap/image drawing.



See also compositionMode(), setCompositionMode(), Composition Modes, and Image Composition Example.

*/
type QPainter__CompositionMode = int

// This is the default mode. The alpha of the source is used to blend the pixel on top of the destination.
const QPainter__CompositionMode_SourceOver QPainter__CompositionMode = 0

// The alpha of the destination is used to blend it on top of the source pixels. This mode is the inverse of CompositionMode_SourceOver.
const QPainter__CompositionMode_DestinationOver QPainter__CompositionMode = 1

// The pixels in the destination are cleared (set to fully transparent) independent of the source.
const QPainter__CompositionMode_Clear QPainter__CompositionMode = 2

// The output is the source pixel. (This means a basic copy operation and is identical to SourceOver when the source pixel is opaque).
const QPainter__CompositionMode_Source QPainter__CompositionMode = 3

// The output is the destination pixel. This means that the blending has no effect. This mode is the inverse of CompositionMode_Source.
const QPainter__CompositionMode_Destination QPainter__CompositionMode = 4

// The output is the source, where the alpha is reduced by that of the destination.
const QPainter__CompositionMode_SourceIn QPainter__CompositionMode = 5

// The output is the destination, where the alpha is reduced by that of the source. This mode is the inverse of CompositionMode_SourceIn.
const QPainter__CompositionMode_DestinationIn QPainter__CompositionMode = 6

// The output is the source, where the alpha is reduced by the inverse of destination.
const QPainter__CompositionMode_SourceOut QPainter__CompositionMode = 7

// The output is the destination, where the alpha is reduced by the inverse of the source. This mode is the inverse of CompositionMode_SourceOut.
const QPainter__CompositionMode_DestinationOut QPainter__CompositionMode = 8

// The source pixel is blended on top of the destination, with the alpha of the source pixel reduced by the alpha of the destination pixel.
const QPainter__CompositionMode_SourceAtop QPainter__CompositionMode = 9

//
const QPainter__CompositionMode_DestinationAtop QPainter__CompositionMode = 10

//
const QPainter__CompositionMode_Xor QPainter__CompositionMode = 11

//
const QPainter__CompositionMode_Plus QPainter__CompositionMode = 12

//
const QPainter__CompositionMode_Multiply QPainter__CompositionMode = 13

//
const QPainter__CompositionMode_Screen QPainter__CompositionMode = 14

//
const QPainter__CompositionMode_Overlay QPainter__CompositionMode = 15

//
const QPainter__CompositionMode_Darken QPainter__CompositionMode = 16

//
const QPainter__CompositionMode_Lighten QPainter__CompositionMode = 17

//
const QPainter__CompositionMode_ColorDodge QPainter__CompositionMode = 18

//
const QPainter__CompositionMode_ColorBurn QPainter__CompositionMode = 19

//
const QPainter__CompositionMode_HardLight QPainter__CompositionMode = 20

//
const QPainter__CompositionMode_SoftLight QPainter__CompositionMode = 21

//
const QPainter__CompositionMode_Difference QPainter__CompositionMode = 22

//
const QPainter__CompositionMode_Exclusion QPainter__CompositionMode = 23

//
const QPainter__RasterOp_SourceOrDestination QPainter__CompositionMode = 24

//
const QPainter__RasterOp_SourceAndDestination QPainter__CompositionMode = 25

//
const QPainter__RasterOp_SourceXorDestination QPainter__CompositionMode = 26

//
const QPainter__RasterOp_NotSourceAndNotDestination QPainter__CompositionMode = 27

//
const QPainter__RasterOp_NotSourceOrNotDestination QPainter__CompositionMode = 28

//
const QPainter__RasterOp_NotSourceXorDestination QPainter__CompositionMode = 29

//
const QPainter__RasterOp_NotSource QPainter__CompositionMode = 30

//
const QPainter__RasterOp_NotSourceAndDestination QPainter__CompositionMode = 31

//
const QPainter__RasterOp_SourceAndNotDestination QPainter__CompositionMode = 32

//
const QPainter__RasterOp_NotSourceOrDestination QPainter__CompositionMode = 33

//
const QPainter__RasterOp_SourceOrNotDestination QPainter__CompositionMode = 34

//
const QPainter__RasterOp_ClearDestination QPainter__CompositionMode = 35

//
const QPainter__RasterOp_SetDestination QPainter__CompositionMode = 36

//
const QPainter__RasterOp_NotDestination QPainter__CompositionMode = 37

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
}

//  keep block end
