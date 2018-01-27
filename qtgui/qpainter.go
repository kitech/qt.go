package qtgui

// /usr/include/qt/QtGui/qpainter.h
// #include <qpainter.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
type QPainter struct {
	*qtrt.CObject
}

func (this *QPainter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPainter) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPainterFromPointer(cthis unsafe.Pointer) *QPainter {
	return &QPainter{&qtrt.CObject{cthis}}
}
func (*QPainter) NewFromPointer(cthis unsafe.Pointer) *QPainter {
	return NewQPainterFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpainter.h:124
// index:0
// Public
// void QPainter()
func NewQPainter() *QPainter {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainterC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPainterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpainter.h:125
// index:1
// Public
// void QPainter(QPaintDevice *)
func NewQPainter_1(arg0 *QPaintDevice /*777 QPaintDevice **/) *QPainter {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainterC2EP12QPaintDevice", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPainterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpainter.h:126
// index:0
// Public
// void ~QPainter()
func DeleteQPainter(*QPainter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:128
// index:0
// Public
// QPaintDevice * device()
func (this *QPainter) Device() *QPaintDevice /*777 QPaintDevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:130
// index:0
// Public
// bool begin(QPaintDevice *)
func (this *QPainter) Begin(arg0 *QPaintDevice /*777 QPaintDevice **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter5beginEP12QPaintDevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:131
// index:0
// Public
// bool end()
func (this *QPainter) End() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:132
// index:0
// Public
// bool isActive()
func (this *QPainter) IsActive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter8isActiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:134
// index:0
// Public
// void initFrom(const QPaintDevice *)
func (this *QPainter) InitFrom(device *QPaintDevice /*777 const QPaintDevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8initFromEPK12QPaintDevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:180
// index:0
// Public
// void setCompositionMode(QPainter::CompositionMode)
func (this *QPainter) SetCompositionMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter18setCompositionModeENS_15CompositionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:181
// index:0
// Public
// QPainter::CompositionMode compositionMode()
func (this *QPainter) CompositionMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter15compositionModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:183
// index:0
// Public
// const QFont & font()
func (this *QPainter) Font() *QFont {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:184
// index:0
// Public
// void setFont(const QFont &)
func (this *QPainter) SetFont(f *QFont) {
	var convArg0 = f.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:186
// index:0
// Public
// QFontMetrics fontMetrics()
func (this *QPainter) FontMetrics() *QFontMetrics /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11fontMetricsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:187
// index:0
// Public
// QFontInfo fontInfo()
func (this *QPainter) FontInfo() *QFontInfo /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter8fontInfoEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQFontInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:189
// index:0
// Public
// void setPen(const QColor &)
func (this *QPainter) SetPen(color *QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter6setPenERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:190
// index:1
// Public
// void setPen(const QPen &)
func (this *QPainter) SetPen_1(pen *QPen) {
	var convArg0 = pen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter6setPenERK4QPen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:191
// index:2
// Public
// void setPen(Qt::PenStyle)
func (this *QPainter) SetPen_2(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter6setPenEN2Qt8PenStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:192
// index:0
// Public
// const QPen & pen()
func (this *QPainter) Pen() *QPen {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter3penEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:194
// index:0
// Public
// void setBrush(const QBrush &)
func (this *QPainter) SetBrush(brush *QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8setBrushERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:195
// index:1
// Public
// void setBrush(Qt::BrushStyle)
func (this *QPainter) SetBrush_1(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8setBrushEN2Qt10BrushStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:196
// index:0
// Public
// const QBrush & brush()
func (this *QPainter) Brush() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter5brushEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:199
// index:0
// Public
// void setBackgroundMode(Qt::BGMode)
func (this *QPainter) SetBackgroundMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17setBackgroundModeEN2Qt6BGModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:200
// index:0
// Public
// Qt::BGMode backgroundMode()
func (this *QPainter) BackgroundMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter14backgroundModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:202
// index:0
// Public
// QPoint brushOrigin()
func (this *QPainter) BrushOrigin() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11brushOriginEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:203
// index:0
// Public inline
// void setBrushOrigin(int, int)
func (this *QPainter) SetBrushOrigin(x int, y int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:204
// index:1
// Public inline
// void setBrushOrigin(const QPoint &)
func (this *QPainter) SetBrushOrigin_1(arg0 *qtcore.QPoint) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:205
// index:2
// Public
// void setBrushOrigin(const QPointF &)
func (this *QPainter) SetBrushOrigin_2(arg0 *qtcore.QPointF) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:207
// index:0
// Public
// void setBackground(const QBrush &)
func (this *QPainter) SetBackground(bg *QBrush) {
	var convArg0 = bg.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:208
// index:0
// Public
// const QBrush & background()
func (this *QPainter) Background() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter10backgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:210
// index:0
// Public
// qreal opacity()
func (this *QPainter) Opacity() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter7opacityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qpainter.h:211
// index:0
// Public
// void setOpacity(qreal)
func (this *QPainter) SetOpacity(opacity float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10setOpacityEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:214
// index:0
// Public
// QRegion clipRegion()
func (this *QPainter) ClipRegion() *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter10clipRegionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:215
// index:0
// Public
// QPainterPath clipPath()
func (this *QPainter) ClipPath() *QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter8clipPathEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:217
// index:0
// Public
// void setClipRect(const QRectF &, Qt::ClipOperation)
func (this *QPainter) SetClipRect(arg0 *qtcore.QRectF, op int) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setClipRectERK6QRectFN2Qt13ClipOperationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:218
// index:1
// Public
// void setClipRect(const QRect &, Qt::ClipOperation)
func (this *QPainter) SetClipRect_1(arg0 *qtcore.QRect, op int) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setClipRectERK5QRectN2Qt13ClipOperationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:219
// index:2
// Public inline
// void setClipRect(int, int, int, int, Qt::ClipOperation)
func (this *QPainter) SetClipRect_2(x int, y int, w int, h int, op int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setClipRectEiiiiN2Qt13ClipOperationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, op)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:221
// index:0
// Public
// void setClipRegion(const QRegion &, Qt::ClipOperation)
func (this *QPainter) SetClipRegion(arg0 *QRegion, op int) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13setClipRegionERK7QRegionN2Qt13ClipOperationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:223
// index:0
// Public
// void setClipPath(const QPainterPath &, Qt::ClipOperation)
func (this *QPainter) SetClipPath(path *QPainterPath, op int) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setClipPathERK12QPainterPathN2Qt13ClipOperationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, op)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:225
// index:0
// Public
// void setClipping(bool)
func (this *QPainter) SetClipping(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setClippingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:226
// index:0
// Public
// bool hasClipping()
func (this *QPainter) HasClipping() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11hasClippingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:228
// index:0
// Public
// QRectF clipBoundingRect()
func (this *QPainter) ClipBoundingRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter16clipBoundingRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:230
// index:0
// Public
// void save()
func (this *QPainter) Save() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter4saveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:231
// index:0
// Public
// void restore()
func (this *QPainter) Restore() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7restoreEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:234
// index:0
// Public
// void setMatrix(const QMatrix &, bool)
func (this *QPainter) SetMatrix(matrix *QMatrix, combine bool) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9setMatrixERK7QMatrixb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:235
// index:0
// Public
// const QMatrix & matrix()
func (this *QPainter) Matrix() *QMatrix {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter6matrixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:236
// index:0
// Public
// const QMatrix & deviceMatrix()
func (this *QPainter) DeviceMatrix() *QMatrix {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter12deviceMatrixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:237
// index:0
// Public
// void resetMatrix()
func (this *QPainter) ResetMatrix() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11resetMatrixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:239
// index:0
// Public
// void setTransform(const QTransform &, bool)
func (this *QPainter) SetTransform(transform *QTransform, combine bool) {
	var convArg0 = transform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12setTransformERK10QTransformb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:240
// index:0
// Public
// const QTransform & transform()
func (this *QPainter) Transform() *QTransform {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter9transformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:241
// index:0
// Public
// const QTransform & deviceTransform()
func (this *QPainter) DeviceTransform() *QTransform {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter15deviceTransformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:242
// index:0
// Public
// void resetTransform()
func (this *QPainter) ResetTransform() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14resetTransformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:244
// index:0
// Public
// void setWorldMatrix(const QMatrix &, bool)
func (this *QPainter) SetWorldMatrix(matrix *QMatrix, combine bool) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14setWorldMatrixERK7QMatrixb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:245
// index:0
// Public
// const QMatrix & worldMatrix()
func (this *QPainter) WorldMatrix() *QMatrix {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11worldMatrixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:247
// index:0
// Public
// void setWorldTransform(const QTransform &, bool)
func (this *QPainter) SetWorldTransform(matrix *QTransform, combine bool) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17setWorldTransformERK10QTransformb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:248
// index:0
// Public
// const QTransform & worldTransform()
func (this *QPainter) WorldTransform() *QTransform {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter14worldTransformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:250
// index:0
// Public
// QMatrix combinedMatrix()
func (this *QPainter) CombinedMatrix() *QMatrix /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter14combinedMatrixEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:251
// index:0
// Public
// QTransform combinedTransform()
func (this *QPainter) CombinedTransform() *QTransform /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter17combinedTransformEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:253
// index:0
// Public
// void setMatrixEnabled(bool)
func (this *QPainter) SetMatrixEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter16setMatrixEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:254
// index:0
// Public
// bool matrixEnabled()
func (this *QPainter) MatrixEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter13matrixEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:256
// index:0
// Public
// void setWorldMatrixEnabled(bool)
func (this *QPainter) SetWorldMatrixEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter21setWorldMatrixEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:257
// index:0
// Public
// bool worldMatrixEnabled()
func (this *QPainter) WorldMatrixEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter18worldMatrixEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:259
// index:0
// Public
// void scale(qreal, qreal)
func (this *QPainter) Scale(sx float64, sy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter5scaleEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sx, sy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:260
// index:0
// Public
// void shear(qreal, qreal)
func (this *QPainter) Shear(sh float64, sv float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter5shearEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sh, sv)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:261
// index:0
// Public
// void rotate(qreal)
func (this *QPainter) Rotate(a float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter6rotateEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:263
// index:0
// Public
// void translate(const QPointF &)
func (this *QPainter) Translate(offset *qtcore.QPointF) {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9translateERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:264
// index:1
// Public inline
// void translate(const QPoint &)
func (this *QPainter) Translate_1(offset *qtcore.QPoint) {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9translateERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:265
// index:2
// Public inline
// void translate(qreal, qreal)
func (this *QPainter) Translate_2(dx float64, dy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9translateEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:267
// index:0
// Public
// QRect window()
func (this *QPainter) Window() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter6windowEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:268
// index:0
// Public
// void setWindow(const QRect &)
func (this *QPainter) SetWindow(window *qtcore.QRect) {
	var convArg0 = window.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9setWindowERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:269
// index:1
// Public inline
// void setWindow(int, int, int, int)
func (this *QPainter) SetWindow_1(x int, y int, w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9setWindowEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:271
// index:0
// Public
// QRect viewport()
func (this *QPainter) Viewport() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter8viewportEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:272
// index:0
// Public
// void setViewport(const QRect &)
func (this *QPainter) SetViewport(viewport *qtcore.QRect) {
	var convArg0 = viewport.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setViewportERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:273
// index:1
// Public inline
// void setViewport(int, int, int, int)
func (this *QPainter) SetViewport_1(x int, y int, w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setViewportEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:275
// index:0
// Public
// void setViewTransformEnabled(bool)
func (this *QPainter) SetViewTransformEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter23setViewTransformEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:276
// index:0
// Public
// bool viewTransformEnabled()
func (this *QPainter) ViewTransformEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter20viewTransformEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:279
// index:0
// Public
// void strokePath(const QPainterPath &, const QPen &)
func (this *QPainter) StrokePath(path *QPainterPath, pen *QPen) {
	var convArg0 = path.GetCthis()
	var convArg1 = pen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10strokePathERK12QPainterPathRK4QPen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:280
// index:0
// Public
// void fillPath(const QPainterPath &, const QBrush &)
func (this *QPainter) FillPath(path *QPainterPath, brush *QBrush) {
	var convArg0 = path.GetCthis()
	var convArg1 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillPathERK12QPainterPathRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:281
// index:0
// Public
// void drawPath(const QPainterPath &)
func (this *QPainter) DrawPath(path *QPainterPath) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawPathERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:283
// index:0
// Public inline
// void drawPoint(const QPointF &)
func (this *QPainter) DrawPoint(pt *qtcore.QPointF) {
	var convArg0 = pt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawPointERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:284
// index:1
// Public inline
// void drawPoint(const QPoint &)
func (this *QPainter) DrawPoint_1(p *qtcore.QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawPointERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:285
// index:2
// Public inline
// void drawPoint(int, int)
func (this *QPainter) DrawPoint_2(x int, y int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawPointEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:287
// index:0
// Public
// void drawPoints(const QPointF *, int)
func (this *QPainter) DrawPoints(points *qtcore.QPointF /*777 const QPointF **/, pointCount int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPointsEPK7QPointFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:288
// index:1
// Public inline
// void drawPoints(const QPolygonF &)
func (this *QPainter) DrawPoints_1(points *QPolygonF) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPointsERK9QPolygonF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:289
// index:2
// Public
// void drawPoints(const QPoint *, int)
func (this *QPainter) DrawPoints_2(points *qtcore.QPoint /*777 const QPoint **/, pointCount int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPointsEPK6QPointi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:290
// index:3
// Public inline
// void drawPoints(const QPolygon &)
func (this *QPainter) DrawPoints_3(points *QPolygon) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPointsERK8QPolygon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:292
// index:0
// Public inline
// void drawLine(const QLineF &)
func (this *QPainter) DrawLine(line *qtcore.QLineF) {
	var convArg0 = line.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawLineERK6QLineF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:293
// index:1
// Public inline
// void drawLine(const QLine &)
func (this *QPainter) DrawLine_1(line *qtcore.QLine) {
	var convArg0 = line.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawLineERK5QLine", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:294
// index:2
// Public inline
// void drawLine(int, int, int, int)
func (this *QPainter) DrawLine_2(x1 int, y1 int, x2 int, y2 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawLineEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:295
// index:3
// Public inline
// void drawLine(const QPoint &, const QPoint &)
func (this *QPainter) DrawLine_3(p1 *qtcore.QPoint, p2 *qtcore.QPoint) {
	var convArg0 = p1.GetCthis()
	var convArg1 = p2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawLineERK6QPointS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:296
// index:4
// Public inline
// void drawLine(const QPointF &, const QPointF &)
func (this *QPainter) DrawLine_4(p1 *qtcore.QPointF, p2 *qtcore.QPointF) {
	var convArg0 = p1.GetCthis()
	var convArg1 = p2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawLineERK7QPointFS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:298
// index:0
// Public
// void drawLines(const QLineF *, int)
func (this *QPainter) DrawLines(lines *qtcore.QLineF /*777 const QLineF **/, lineCount int) {
	var convArg0 = lines.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK6QLineFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:300
// index:1
// Public
// void drawLines(const QPointF *, int)
func (this *QPainter) DrawLines_1(pointPairs *qtcore.QPointF /*777 const QPointF **/, lineCount int) {
	var convArg0 = pointPairs.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK7QPointFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:302
// index:2
// Public
// void drawLines(const QLine *, int)
func (this *QPainter) DrawLines_2(lines *qtcore.QLine /*777 const QLine **/, lineCount int) {
	var convArg0 = lines.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK5QLinei", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:304
// index:3
// Public
// void drawLines(const QPoint *, int)
func (this *QPainter) DrawLines_3(pointPairs *qtcore.QPoint /*777 const QPoint **/, lineCount int) {
	var convArg0 = pointPairs.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK6QPointi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:307
// index:0
// Public inline
// void drawRect(const QRectF &)
func (this *QPainter) DrawRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:308
// index:1
// Public inline
// void drawRect(int, int, int, int)
func (this *QPainter) DrawRect_1(x1 int, y1 int, w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawRectEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:309
// index:2
// Public inline
// void drawRect(const QRect &)
func (this *QPainter) DrawRect_2(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:311
// index:0
// Public
// void drawRects(const QRectF *, int)
func (this *QPainter) DrawRects(rects *qtcore.QRectF /*777 const QRectF **/, rectCount int) {
	var convArg0 = rects.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawRectsEPK6QRectFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rectCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:313
// index:1
// Public
// void drawRects(const QRect *, int)
func (this *QPainter) DrawRects_1(rects *qtcore.QRect /*777 const QRect **/, rectCount int) {
	var convArg0 = rects.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawRectsEPK5QRecti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rectCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:316
// index:0
// Public
// void drawEllipse(const QRectF &)
func (this *QPainter) DrawEllipse(r *qtcore.QRectF) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:317
// index:1
// Public
// void drawEllipse(const QRect &)
func (this *QPainter) DrawEllipse_1(r *qtcore.QRect) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:318
// index:2
// Public inline
// void drawEllipse(int, int, int, int)
func (this *QPainter) DrawEllipse_2(x int, y int, w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawEllipseEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:320
// index:3
// Public inline
// void drawEllipse(const QPointF &, qreal, qreal)
func (this *QPainter) DrawEllipse_3(center *qtcore.QPointF, rx float64, ry float64) {
	var convArg0 = center.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK7QPointFdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rx, ry)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:321
// index:4
// Public inline
// void drawEllipse(const QPoint &, int, int)
func (this *QPainter) DrawEllipse_4(center *qtcore.QPoint, rx int, ry int) {
	var convArg0 = center.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK6QPointii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rx, ry)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:323
// index:0
// Public
// void drawPolyline(const QPointF *, int)
func (this *QPainter) DrawPolyline(points *qtcore.QPointF /*777 const QPointF **/, pointCount int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawPolylineEPK7QPointFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:324
// index:1
// Public inline
// void drawPolyline(const QPolygonF &)
func (this *QPainter) DrawPolyline_1(polyline *QPolygonF) {
	var convArg0 = polyline.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawPolylineERK9QPolygonF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:325
// index:2
// Public
// void drawPolyline(const QPoint *, int)
func (this *QPainter) DrawPolyline_2(points *qtcore.QPoint /*777 const QPoint **/, pointCount int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawPolylineEPK6QPointi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:326
// index:3
// Public inline
// void drawPolyline(const QPolygon &)
func (this *QPainter) DrawPolyline_3(polygon *QPolygon) {
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawPolylineERK8QPolygon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:328
// index:0
// Public
// void drawPolygon(const QPointF *, int, Qt::FillRule)
func (this *QPainter) DrawPolygon(points *qtcore.QPointF /*777 const QPointF **/, pointCount int, fillRule int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPolygonEPK7QPointFiN2Qt8FillRuleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, fillRule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:329
// index:1
// Public inline
// void drawPolygon(const QPolygonF &, Qt::FillRule)
func (this *QPainter) DrawPolygon_1(polygon *QPolygonF, fillRule int) {
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPolygonERK9QPolygonFN2Qt8FillRuleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:330
// index:2
// Public
// void drawPolygon(const QPoint *, int, Qt::FillRule)
func (this *QPainter) DrawPolygon_2(points *qtcore.QPoint /*777 const QPoint **/, pointCount int, fillRule int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPolygonEPK6QPointiN2Qt8FillRuleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount, fillRule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:331
// index:3
// Public inline
// void drawPolygon(const QPolygon &, Qt::FillRule)
func (this *QPainter) DrawPolygon_3(polygon *QPolygon, fillRule int) {
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPolygonERK8QPolygonN2Qt8FillRuleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:333
// index:0
// Public
// void drawConvexPolygon(const QPointF *, int)
func (this *QPainter) DrawConvexPolygon(points *qtcore.QPointF /*777 const QPointF **/, pointCount int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonEPK7QPointFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:334
// index:1
// Public inline
// void drawConvexPolygon(const QPolygonF &)
func (this *QPainter) DrawConvexPolygon_1(polygon *QPolygonF) {
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonERK9QPolygonF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:335
// index:2
// Public
// void drawConvexPolygon(const QPoint *, int)
func (this *QPainter) DrawConvexPolygon_2(points *qtcore.QPoint /*777 const QPoint **/, pointCount int) {
	var convArg0 = points.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonEPK6QPointi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:336
// index:3
// Public inline
// void drawConvexPolygon(const QPolygon &)
func (this *QPainter) DrawConvexPolygon_3(polygon *QPolygon) {
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonERK8QPolygon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:338
// index:0
// Public
// void drawArc(const QRectF &, int, int)
func (this *QPainter) DrawArc(rect *qtcore.QRectF, a int, alen int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawArcERK6QRectFii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:339
// index:1
// Public inline
// void drawArc(const QRect &, int, int)
func (this *QPainter) DrawArc_1(arg0 *qtcore.QRect, a int, alen int) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawArcERK5QRectii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:340
// index:2
// Public inline
// void drawArc(int, int, int, int, int, int)
func (this *QPainter) DrawArc_2(x int, y int, w int, h int, a int, alen int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawArcEiiiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, a, alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:342
// index:0
// Public
// void drawPie(const QRectF &, int, int)
func (this *QPainter) DrawPie(rect *qtcore.QRectF, a int, alen int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawPieERK6QRectFii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:343
// index:1
// Public inline
// void drawPie(int, int, int, int, int, int)
func (this *QPainter) DrawPie_1(x int, y int, w int, h int, a int, alen int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawPieEiiiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, a, alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:344
// index:2
// Public inline
// void drawPie(const QRect &, int, int)
func (this *QPainter) DrawPie_2(arg0 *qtcore.QRect, a int, alen int) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawPieERK5QRectii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:346
// index:0
// Public
// void drawChord(const QRectF &, int, int)
func (this *QPainter) DrawChord(rect *qtcore.QRectF, a int, alen int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawChordERK6QRectFii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:347
// index:1
// Public inline
// void drawChord(int, int, int, int, int, int)
func (this *QPainter) DrawChord_1(x int, y int, w int, h int, a int, alen int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawChordEiiiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, a, alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:348
// index:2
// Public inline
// void drawChord(const QRect &, int, int)
func (this *QPainter) DrawChord_2(arg0 *qtcore.QRect, a int, alen int) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawChordERK5QRectii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, a, alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:350
// index:0
// Public
// void drawRoundedRect(const QRectF &, qreal, qreal, Qt::SizeMode)
func (this *QPainter) DrawRoundedRect(rect *qtcore.QRectF, xRadius float64, yRadius float64, mode int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectERK6QRectFddN2Qt8SizeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRadius, yRadius, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:352
// index:1
// Public inline
// void drawRoundedRect(int, int, int, int, qreal, qreal, Qt::SizeMode)
func (this *QPainter) DrawRoundedRect_1(x int, y int, w int, h int, xRadius float64, yRadius float64, mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectEiiiiddN2Qt8SizeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xRadius, yRadius, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:354
// index:2
// Public inline
// void drawRoundedRect(const QRect &, qreal, qreal, Qt::SizeMode)
func (this *QPainter) DrawRoundedRect_2(rect *qtcore.QRect, xRadius float64, yRadius float64, mode int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectERK5QRectddN2Qt8SizeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRadius, yRadius, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:357
// index:0
// Public
// void drawRoundRect(const QRectF &, int, int)
func (this *QPainter) DrawRoundRect(r *qtcore.QRectF, xround int, yround int) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK6QRectFii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xround, yround)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:358
// index:1
// Public inline
// void drawRoundRect(int, int, int, int, int, int)
func (this *QPainter) DrawRoundRect_1(x int, y int, w int, h int, arg4 int, arg5 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectEiiiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, arg4, arg5)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:359
// index:2
// Public inline
// void drawRoundRect(const QRect &, int, int)
func (this *QPainter) DrawRoundRect_2(r *qtcore.QRect, xround int, yround int) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK5QRectii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xround, yround)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:361
// index:0
// Public
// void drawTiledPixmap(const QRectF &, const QPixmap &, const QPointF &)
func (this *QPainter) DrawTiledPixmap(rect *qtcore.QRectF, pm *QPixmap, offset *qtcore.QPointF) {
	var convArg0 = rect.GetCthis()
	var convArg1 = pm.GetCthis()
	var convArg2 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:362
// index:1
// Public inline
// void drawTiledPixmap(int, int, int, int, const QPixmap &, int, int)
func (this *QPainter) DrawTiledPixmap_1(x int, y int, w int, h int, arg4 *QPixmap, sx int, sy int) {
	var convArg4 = arg4.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, sx, sy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:363
// index:2
// Public inline
// void drawTiledPixmap(const QRect &, const QPixmap &, const QPoint &)
func (this *QPainter) DrawTiledPixmap_2(arg0 *qtcore.QRect, arg1 *QPixmap, arg2 *qtcore.QPoint) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	var convArg2 = arg2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:365
// index:0
// Public
// void drawPicture(const QPointF &, const QPicture &)
func (this *QPainter) DrawPicture(p *qtcore.QPointF, picture *QPicture) {
	var convArg0 = p.GetCthis()
	var convArg1 = picture.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPictureERK7QPointFRK8QPicture", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:366
// index:1
// Public inline
// void drawPicture(int, int, const QPicture &)
func (this *QPainter) DrawPicture_1(x int, y int, picture *QPicture) {
	var convArg2 = picture.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPictureEiiRK8QPicture", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:367
// index:2
// Public inline
// void drawPicture(const QPoint &, const QPicture &)
func (this *QPainter) DrawPicture_2(p *qtcore.QPoint, picture *QPicture) {
	var convArg0 = p.GetCthis()
	var convArg1 = picture.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPictureERK6QPointRK8QPicture", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:370
// index:0
// Public
// void drawPixmap(const QRectF &, const QPixmap &, const QRectF &)
func (this *QPainter) DrawPixmap(targetRect *qtcore.QRectF, pixmap *QPixmap, sourceRect *qtcore.QRectF) {
	var convArg0 = targetRect.GetCthis()
	var convArg1 = pixmap.GetCthis()
	var convArg2 = sourceRect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QRectFRK7QPixmapS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:371
// index:1
// Public inline
// void drawPixmap(const QRect &, const QPixmap &, const QRect &)
func (this *QPainter) DrawPixmap_1(targetRect *qtcore.QRect, pixmap *QPixmap, sourceRect *qtcore.QRect) {
	var convArg0 = targetRect.GetCthis()
	var convArg1 = pixmap.GetCthis()
	var convArg2 = sourceRect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK5QRectRK7QPixmapS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:372
// index:2
// Public inline
// void drawPixmap(int, int, int, int, const QPixmap &, int, int, int, int)
func (this *QPainter) DrawPixmap_2(x int, y int, w int, h int, pm *QPixmap, sx int, sy int, sw int, sh int) {
	var convArg4 = pm.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiiiRK7QPixmapiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, sx, sy, sw, sh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:374
// index:3
// Public inline
// void drawPixmap(int, int, const QPixmap &, int, int, int, int)
func (this *QPainter) DrawPixmap_3(x int, y int, pm *QPixmap, sx int, sy int, sw int, sh int) {
	var convArg2 = pm.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiRK7QPixmapiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, sx, sy, sw, sh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:376
// index:4
// Public inline
// void drawPixmap(const QPointF &, const QPixmap &, const QRectF &)
func (this *QPainter) DrawPixmap_4(p *qtcore.QPointF, pm *QPixmap, sr *qtcore.QRectF) {
	var convArg0 = p.GetCthis()
	var convArg1 = pm.GetCthis()
	var convArg2 = sr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmapRK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:377
// index:5
// Public inline
// void drawPixmap(const QPoint &, const QPixmap &, const QRect &)
func (this *QPainter) DrawPixmap_5(p *qtcore.QPoint, pm *QPixmap, sr *qtcore.QRect) {
	var convArg0 = p.GetCthis()
	var convArg1 = pm.GetCthis()
	var convArg2 = sr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QPointRK7QPixmapRK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:378
// index:6
// Public
// void drawPixmap(const QPointF &, const QPixmap &)
func (this *QPainter) DrawPixmap_6(p *qtcore.QPointF, pm *QPixmap) {
	var convArg0 = p.GetCthis()
	var convArg1 = pm.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:379
// index:7
// Public inline
// void drawPixmap(const QPoint &, const QPixmap &)
func (this *QPainter) DrawPixmap_7(p *qtcore.QPoint, pm *QPixmap) {
	var convArg0 = p.GetCthis()
	var convArg1 = pm.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QPointRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:380
// index:8
// Public inline
// void drawPixmap(int, int, const QPixmap &)
func (this *QPainter) DrawPixmap_8(x int, y int, pm *QPixmap) {
	var convArg2 = pm.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:381
// index:9
// Public inline
// void drawPixmap(const QRect &, const QPixmap &)
func (this *QPainter) DrawPixmap_9(r *qtcore.QRect, pm *QPixmap) {
	var convArg0 = r.GetCthis()
	var convArg1 = pm.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK5QRectRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:382
// index:10
// Public inline
// void drawPixmap(int, int, int, int, const QPixmap &)
func (this *QPainter) DrawPixmap_10(x int, y int, w int, h int, pm *QPixmap) {
	var convArg4 = pm.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiiiRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:387
// index:0
// Public
// void drawImage(const QRectF &, const QImage &, const QRectF &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage(targetRect *qtcore.QRectF, image *QImage, sourceRect *qtcore.QRectF, flags int) {
	var convArg0 = targetRect.GetCthis()
	var convArg1 = image.GetCthis()
	var convArg2 = sourceRect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QRectFRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:389
// index:1
// Public inline
// void drawImage(const QRect &, const QImage &, const QRect &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_1(targetRect *qtcore.QRect, image *QImage, sourceRect *qtcore.QRect, flags int) {
	var convArg0 = targetRect.GetCthis()
	var convArg1 = image.GetCthis()
	var convArg2 = sourceRect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK5QRectRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:391
// index:2
// Public inline
// void drawImage(const QPointF &, const QImage &, const QRectF &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_2(p *qtcore.QPointF, image *QImage, sr *qtcore.QRectF, flags int) {
	var convArg0 = p.GetCthis()
	var convArg1 = image.GetCthis()
	var convArg2 = sr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK7QPointFRK6QImageRK6QRectF6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:393
// index:3
// Public inline
// void drawImage(const QPoint &, const QImage &, const QRect &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_3(p *qtcore.QPoint, image *QImage, sr *qtcore.QRect, flags int) {
	var convArg0 = p.GetCthis()
	var convArg1 = image.GetCthis()
	var convArg2 = sr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QPointRK6QImageRK5QRect6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:395
// index:4
// Public inline
// void drawImage(const QRectF &, const QImage &)
func (this *QPainter) DrawImage_4(r *qtcore.QRectF, image *QImage) {
	var convArg0 = r.GetCthis()
	var convArg1 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QRectFRK6QImage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:396
// index:5
// Public inline
// void drawImage(const QRect &, const QImage &)
func (this *QPainter) DrawImage_5(r *qtcore.QRect, image *QImage) {
	var convArg0 = r.GetCthis()
	var convArg1 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK5QRectRK6QImage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:397
// index:6
// Public
// void drawImage(const QPointF &, const QImage &)
func (this *QPainter) DrawImage_6(p *qtcore.QPointF, image *QImage) {
	var convArg0 = p.GetCthis()
	var convArg1 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK7QPointFRK6QImage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:398
// index:7
// Public inline
// void drawImage(const QPoint &, const QImage &)
func (this *QPainter) DrawImage_7(p *qtcore.QPoint, image *QImage) {
	var convArg0 = p.GetCthis()
	var convArg1 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QPointRK6QImage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:399
// index:8
// Public inline
// void drawImage(int, int, const QImage &, int, int, int, int, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_8(x int, y int, image *QImage, sx int, sy int, sw int, sh int, flags int) {
	var convArg2 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageEiiRK6QImageiiii6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, sx, sy, sw, sh, flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:402
// index:0
// Public
// void setLayoutDirection(Qt::LayoutDirection)
func (this *QPainter) SetLayoutDirection(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter18setLayoutDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:403
// index:0
// Public
// Qt::LayoutDirection layoutDirection()
func (this *QPainter) LayoutDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter15layoutDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:406
// index:0
// Public
// void drawGlyphRun(const QPointF &, const QGlyphRun &)
func (this *QPainter) DrawGlyphRun(position *qtcore.QPointF, glyphRun *QGlyphRun) {
	var convArg0 = position.GetCthis()
	var convArg1 = glyphRun.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawGlyphRunERK7QPointFRK9QGlyphRun", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:409
// index:0
// Public
// void drawStaticText(const QPointF &, const QStaticText &)
func (this *QPainter) DrawStaticText(topLeftPosition *qtcore.QPointF, staticText *QStaticText) {
	var convArg0 = topLeftPosition.GetCthis()
	var convArg1 = staticText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextERK7QPointFRK11QStaticText", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:410
// index:1
// Public inline
// void drawStaticText(const QPoint &, const QStaticText &)
func (this *QPainter) DrawStaticText_1(topLeftPosition *qtcore.QPoint, staticText *QStaticText) {
	var convArg0 = topLeftPosition.GetCthis()
	var convArg1 = staticText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextERK6QPointRK11QStaticText", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:411
// index:2
// Public inline
// void drawStaticText(int, int, const QStaticText &)
func (this *QPainter) DrawStaticText_2(left int, top int, staticText *QStaticText) {
	var convArg2 = staticText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextEiiRK11QStaticText", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, top, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:413
// index:0
// Public
// void drawText(const QPointF &, const QString &)
func (this *QPainter) DrawText(p *qtcore.QPointF, s *qtcore.QString) {
	var convArg0 = p.GetCthis()
	var convArg1 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK7QPointFRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:414
// index:1
// Public inline
// void drawText(const QPoint &, const QString &)
func (this *QPainter) DrawText_1(p *qtcore.QPoint, s *qtcore.QString) {
	var convArg0 = p.GetCthis()
	var convArg1 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QPointRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:415
// index:2
// Public inline
// void drawText(int, int, const QString &)
func (this *QPainter) DrawText_2(x int, y int, s *qtcore.QString) {
	var convArg2 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextEiiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:417
// index:3
// Public
// void drawText(const QPointF &, const QString &, int, int)
func (this *QPainter) DrawText_3(p *qtcore.QPointF, str *qtcore.QString, tf int, justificationPadding int) {
	var convArg0 = p.GetCthis()
	var convArg1 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK7QPointFRK7QStringii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, tf, justificationPadding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:419
// index:4
// Public
// void drawText(const QRectF &, int, const QString &, QRectF *)
func (this *QPainter) DrawText_4(r *qtcore.QRectF, flags int, text *qtcore.QString, br *qtcore.QRectF /*777 QRectF **/) {
	var convArg0 = r.GetCthis()
	var convArg2 = text.GetCthis()
	var convArg3 = br.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:420
// index:5
// Public
// void drawText(const QRect &, int, const QString &, QRect *)
func (this *QPainter) DrawText_5(r *qtcore.QRect, flags int, text *qtcore.QString, br *qtcore.QRect /*777 QRect **/) {
	var convArg0 = r.GetCthis()
	var convArg2 = text.GetCthis()
	var convArg3 = br.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:421
// index:6
// Public inline
// void drawText(int, int, int, int, int, const QString &, QRect *)
func (this *QPainter) DrawText_6(x int, y int, w int, h int, flags int, text *qtcore.QString, br *qtcore.QRect /*777 QRect **/) {
	var convArg5 = text.GetCthis()
	var convArg6 = br.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, flags, convArg5, convArg6)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:423
// index:7
// Public
// void drawText(const QRectF &, const QString &, const QTextOption &)
func (this *QPainter) DrawText_7(r *qtcore.QRectF, text *qtcore.QString, o *QTextOption) {
	var convArg0 = r.GetCthis()
	var convArg1 = text.GetCthis()
	var convArg2 = o.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:425
// index:0
// Public
// QRectF boundingRect(const QRectF &, int, const QString &)
func (this *QPainter) BoundingRect(rect *qtcore.QRectF, flags int, text *qtcore.QString) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK6QRectFiRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, flags, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:426
// index:1
// Public
// QRect boundingRect(const QRect &, int, const QString &)
func (this *QPainter) BoundingRect_1(rect *qtcore.QRect, flags int, text *qtcore.QString) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK5QRectiRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, flags, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:427
// index:2
// Public inline
// QRect boundingRect(int, int, int, int, int, const QString &)
func (this *QPainter) BoundingRect_2(x int, y int, w int, h int, flags int, text *qtcore.QString) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg5 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12boundingRectEiiiiiRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, w, h, flags, convArg5)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:429
// index:3
// Public
// QRectF boundingRect(const QRectF &, const QString &, const QTextOption &)
func (this *QPainter) BoundingRect_3(rect *qtcore.QRectF, text *qtcore.QString, o *QTextOption) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	var convArg1 = text.GetCthis()
	var convArg2 = o.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:431
// index:0
// Public
// void drawTextItem(const QPointF &, const QTextItem &)
func (this *QPainter) DrawTextItem(p *qtcore.QPointF, ti *QTextItem) {
	var convArg0 = p.GetCthis()
	var convArg1 = ti.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawTextItemERK7QPointFRK9QTextItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:432
// index:1
// Public inline
// void drawTextItem(int, int, const QTextItem &)
func (this *QPainter) DrawTextItem_1(x int, y int, ti *QTextItem) {
	var convArg2 = ti.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawTextItemEiiRK9QTextItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:433
// index:2
// Public inline
// void drawTextItem(const QPoint &, const QTextItem &)
func (this *QPainter) DrawTextItem_2(p *qtcore.QPoint, ti *QTextItem) {
	var convArg0 = p.GetCthis()
	var convArg1 = ti.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawTextItemERK6QPointRK9QTextItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:435
// index:0
// Public
// void fillRect(const QRectF &, const QBrush &)
func (this *QPainter) FillRect(arg0 *qtcore.QRectF, arg1 *QBrush) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:436
// index:1
// Public inline
// void fillRect(int, int, int, int, const QBrush &)
func (this *QPainter) FillRect_1(x int, y int, w int, h int, arg4 *QBrush) {
	var convArg4 = arg4.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:437
// index:2
// Public
// void fillRect(const QRect &, const QBrush &)
func (this *QPainter) FillRect_2(arg0 *qtcore.QRect, arg1 *QBrush) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:439
// index:3
// Public
// void fillRect(const QRectF &, const QColor &)
func (this *QPainter) FillRect_3(arg0 *qtcore.QRectF, color *QColor) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFRK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:440
// index:4
// Public inline
// void fillRect(int, int, int, int, const QColor &)
func (this *QPainter) FillRect_4(x int, y int, w int, h int, color *QColor) {
	var convArg4 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiRK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:441
// index:5
// Public
// void fillRect(const QRect &, const QColor &)
func (this *QPainter) FillRect_5(arg0 *qtcore.QRect, color *QColor) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectRK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:443
// index:6
// Public inline
// void fillRect(int, int, int, int, Qt::GlobalColor)
func (this *QPainter) FillRect_6(x int, y int, w int, h int, c int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiN2Qt11GlobalColorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:444
// index:7
// Public inline
// void fillRect(const QRect &, Qt::GlobalColor)
func (this *QPainter) FillRect_7(r *qtcore.QRect, c int) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectN2Qt11GlobalColorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:445
// index:8
// Public inline
// void fillRect(const QRectF &, Qt::GlobalColor)
func (this *QPainter) FillRect_8(r *qtcore.QRectF, c int) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFN2Qt11GlobalColorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:447
// index:9
// Public inline
// void fillRect(int, int, int, int, Qt::BrushStyle)
func (this *QPainter) FillRect_9(x int, y int, w int, h int, style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiN2Qt10BrushStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:448
// index:10
// Public inline
// void fillRect(const QRect &, Qt::BrushStyle)
func (this *QPainter) FillRect_10(r *qtcore.QRect, style int) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectN2Qt10BrushStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:449
// index:11
// Public inline
// void fillRect(const QRectF &, Qt::BrushStyle)
func (this *QPainter) FillRect_11(r *qtcore.QRectF, style int) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFN2Qt10BrushStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:451
// index:0
// Public
// void eraseRect(const QRectF &)
func (this *QPainter) EraseRect(arg0 *qtcore.QRectF) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9eraseRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:452
// index:1
// Public inline
// void eraseRect(int, int, int, int)
func (this *QPainter) EraseRect_1(x int, y int, w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9eraseRectEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:453
// index:2
// Public inline
// void eraseRect(const QRect &)
func (this *QPainter) EraseRect_2(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9eraseRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:455
// index:0
// Public
// void setRenderHint(QPainter::RenderHint, bool)
func (this *QPainter) SetRenderHint(hint int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13setRenderHintENS_10RenderHintEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hint, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:457
// index:0
// Public
// QPainter::RenderHints renderHints()
func (this *QPainter) RenderHints() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11renderHintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpainter.h:458
// index:0
// Public inline
// bool testRenderHint(QPainter::RenderHint)
func (this *QPainter) TestRenderHint(hint int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter14testRenderHintENS_10RenderHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hint)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpainter.h:460
// index:0
// Public
// QPaintEngine * paintEngine()
func (this *QPainter) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11paintEngineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:462
// index:0
// Public static
// void setRedirected(const QPaintDevice *, QPaintDevice *, const QPoint &)
func (this *QPainter) SetRedirected(device *QPaintDevice /*777 const QPaintDevice **/, replacement *QPaintDevice /*777 QPaintDevice **/, offset *qtcore.QPoint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13setRedirectedEPK12QPaintDevicePS0_RK6QPoint", ffiqt.FFI_TYPE_POINTER, device, replacement, offset)
	gopp.ErrPrint(err, rv)
}
func QPainter_SetRedirected(device *QPaintDevice /*777 const QPaintDevice **/, replacement *QPaintDevice /*777 QPaintDevice **/, offset *qtcore.QPoint) {
	var nilthis *QPainter
	nilthis.SetRedirected(device, replacement, offset)
}

// /usr/include/qt/QtGui/qpainter.h:464
// index:0
// Public static
// QPaintDevice * redirected(const QPaintDevice *, QPoint *)
func (this *QPainter) Redirected(device *QPaintDevice /*777 const QPaintDevice **/, offset *qtcore.QPoint /*777 QPoint **/) *QPaintDevice /*777 QPaintDevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10redirectedEPK12QPaintDeviceP6QPoint", ffiqt.FFI_TYPE_POINTER, device, offset)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QPainter_Redirected(device *QPaintDevice /*777 const QPaintDevice **/, offset *qtcore.QPoint /*777 QPoint **/) *QPaintDevice /*777 QPaintDevice **/ {
	var nilthis *QPainter
	rv := nilthis.Redirected(device, offset)
	return rv
}

// /usr/include/qt/QtGui/qpainter.h:465
// index:0
// Public static
// void restoreRedirected(const QPaintDevice *)
func (this *QPainter) RestoreRedirected(device *QPaintDevice /*777 const QPaintDevice **/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17restoreRedirectedEPK12QPaintDevice", ffiqt.FFI_TYPE_POINTER, device)
	gopp.ErrPrint(err, rv)
}
func QPainter_RestoreRedirected(device *QPaintDevice /*777 const QPaintDevice **/) {
	var nilthis *QPainter
	nilthis.RestoreRedirected(device)
}

// /usr/include/qt/QtGui/qpainter.h:467
// index:0
// Public
// void beginNativePainting()
func (this *QPainter) BeginNativePainting() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter19beginNativePaintingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:468
// index:0
// Public
// void endNativePainting()
func (this *QPainter) EndNativePainting() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17endNativePaintingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QPainter__RenderHint = int

const QPainter__Antialiasing QPainter__RenderHint = 1
const QPainter__TextAntialiasing QPainter__RenderHint = 2
const QPainter__SmoothPixmapTransform QPainter__RenderHint = 4
const QPainter__HighQualityAntialiasing QPainter__RenderHint = 8
const QPainter__NonCosmeticDefaultPen QPainter__RenderHint = 16
const QPainter__Qt4CompatiblePainting QPainter__RenderHint = 32

type QPainter__PixmapFragmentHint = int

const QPainter__OpaqueHint QPainter__PixmapFragmentHint = 1

type QPainter__CompositionMode = int

const QPainter__CompositionMode_SourceOver QPainter__CompositionMode = 0
const QPainter__CompositionMode_DestinationOver QPainter__CompositionMode = 1
const QPainter__CompositionMode_Clear QPainter__CompositionMode = 2
const QPainter__CompositionMode_Source QPainter__CompositionMode = 3
const QPainter__CompositionMode_Destination QPainter__CompositionMode = 4
const QPainter__CompositionMode_SourceIn QPainter__CompositionMode = 5
const QPainter__CompositionMode_DestinationIn QPainter__CompositionMode = 6
const QPainter__CompositionMode_SourceOut QPainter__CompositionMode = 7
const QPainter__CompositionMode_DestinationOut QPainter__CompositionMode = 8
const QPainter__CompositionMode_SourceAtop QPainter__CompositionMode = 9
const QPainter__CompositionMode_DestinationAtop QPainter__CompositionMode = 10
const QPainter__CompositionMode_Xor QPainter__CompositionMode = 11
const QPainter__CompositionMode_Plus QPainter__CompositionMode = 12
const QPainter__CompositionMode_Multiply QPainter__CompositionMode = 13
const QPainter__CompositionMode_Screen QPainter__CompositionMode = 14
const QPainter__CompositionMode_Overlay QPainter__CompositionMode = 15
const QPainter__CompositionMode_Darken QPainter__CompositionMode = 16
const QPainter__CompositionMode_Lighten QPainter__CompositionMode = 17
const QPainter__CompositionMode_ColorDodge QPainter__CompositionMode = 18
const QPainter__CompositionMode_ColorBurn QPainter__CompositionMode = 19
const QPainter__CompositionMode_HardLight QPainter__CompositionMode = 20
const QPainter__CompositionMode_SoftLight QPainter__CompositionMode = 21
const QPainter__CompositionMode_Difference QPainter__CompositionMode = 22
const QPainter__CompositionMode_Exclusion QPainter__CompositionMode = 23
const QPainter__RasterOp_SourceOrDestination QPainter__CompositionMode = 24
const QPainter__RasterOp_SourceAndDestination QPainter__CompositionMode = 25
const QPainter__RasterOp_SourceXorDestination QPainter__CompositionMode = 26
const QPainter__RasterOp_NotSourceAndNotDestination QPainter__CompositionMode = 27
const QPainter__RasterOp_NotSourceOrNotDestination QPainter__CompositionMode = 28
const QPainter__RasterOp_NotSourceXorDestination QPainter__CompositionMode = 29
const QPainter__RasterOp_NotSource QPainter__CompositionMode = 30
const QPainter__RasterOp_NotSourceAndDestination QPainter__CompositionMode = 31
const QPainter__RasterOp_SourceAndNotDestination QPainter__CompositionMode = 32
const QPainter__RasterOp_NotSourceOrDestination QPainter__CompositionMode = 33
const QPainter__RasterOp_SourceOrNotDestination QPainter__CompositionMode = 34
const QPainter__RasterOp_ClearDestination QPainter__CompositionMode = 35
const QPainter__RasterOp_SetDestination QPainter__CompositionMode = 36
const QPainter__RasterOp_NotDestination QPainter__CompositionMode = 37

//  body block end
