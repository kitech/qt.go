//  header block begin
// /usr/include/qt/QtGui/qpainter.h
// #include <qpainter.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	return this.Cthis
}

// /usr/include/qt/QtGui/qpainter.h:124
// index:0
// void QPainter()
func NewQPainter() *QPainter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainterC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPainterFromPointer(cthis)
	return gothis
}
func NewQPainterFromPointer(cthis unsafe.Pointer) *QPainter {
	return &QPainter{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpainter.h:125
// index:1
// void QPainter(class QPaintDevice *)
func NewQPainter_1(arg0 unsafe.Pointer) *QPainter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainterC2EP12QPaintDevice", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPainterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpainter.h:126
// index:0
// void ~QPainter()
func DeleteQPainter(*QPainter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:128
// index:0
// QPaintDevice * device()
func (this *QPainter) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter6deviceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:130
// index:0
// bool begin(class QPaintDevice *)
func (this *QPainter) Begin(arg0 unsafe.Pointer) {
	// 0: (, QPaintDevice * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter5beginEP12QPaintDevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:131
// index:0
// bool end()
func (this *QPainter) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter3endEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:132
// index:0
// bool isActive()
func (this *QPainter) IsActive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter8isActiveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:134
// index:0
// void initFrom(const class QPaintDevice *)
func (this *QPainter) InitFrom(device unsafe.Pointer) {
	// 0: (, device const QPaintDevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8initFromEPK12QPaintDevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:180
// index:0
// void setCompositionMode(enum QPainter::CompositionMode)
func (this *QPainter) SetCompositionMode(mode int) {
	// 0: (, mode QPainter::CompositionMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter18setCompositionModeENS_15CompositionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:181
// index:0
// QPainter::CompositionMode compositionMode()
func (this *QPainter) CompositionMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter15compositionModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:183
// index:0
// const QFont & font()
func (this *QPainter) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter4fontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:184
// index:0
// void setFont(const class QFont &)
func (this *QPainter) SetFont(f unsafe.Pointer) {
	// 0: (, f const QFont &), (f)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:186
// index:0
// QFontMetrics fontMetrics()
func (this *QPainter) FontMetrics() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11fontMetricsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:187
// index:0
// QFontInfo fontInfo()
func (this *QPainter) FontInfo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter8fontInfoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:189
// index:0
// void setPen(const class QColor &)
func (this *QPainter) SetPen(color unsafe.Pointer) {
	// 0: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter6setPenERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:190
// index:1
// void setPen(const class QPen &)
func (this *QPainter) SetPen_1(pen unsafe.Pointer) {
	// 1: (, pen const QPen &), (pen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter6setPenERK4QPen", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:191
// index:2
// void setPen(Qt::PenStyle)
func (this *QPainter) SetPen_2(style int) {
	// 2: (, style Qt::PenStyle), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter6setPenEN2Qt8PenStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:192
// index:0
// const QPen & pen()
func (this *QPainter) Pen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter3penEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:194
// index:0
// void setBrush(const class QBrush &)
func (this *QPainter) SetBrush(brush unsafe.Pointer) {
	// 0: (, brush const QBrush &), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8setBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:195
// index:1
// void setBrush(Qt::BrushStyle)
func (this *QPainter) SetBrush_1(style int) {
	// 1: (, style Qt::BrushStyle), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8setBrushEN2Qt10BrushStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:196
// index:0
// const QBrush & brush()
func (this *QPainter) Brush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter5brushEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:199
// index:0
// void setBackgroundMode(Qt::BGMode)
func (this *QPainter) SetBackgroundMode(mode int) {
	// 0: (, mode Qt::BGMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17setBackgroundModeEN2Qt6BGModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:200
// index:0
// Qt::BGMode backgroundMode()
func (this *QPainter) BackgroundMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter14backgroundModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:202
// index:0
// QPoint brushOrigin()
func (this *QPainter) BrushOrigin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11brushOriginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:203
// index:0
// inline
// void setBrushOrigin(int, int)
func (this *QPainter) SetBrushOrigin(x int, y int) {
	// 0: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:204
// index:1
// inline
// void setBrushOrigin(const class QPoint &)
func (this *QPainter) SetBrushOrigin_1(arg0 unsafe.Pointer) {
	// 1: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:205
// index:2
// void setBrushOrigin(const class QPointF &)
func (this *QPainter) SetBrushOrigin_2(arg0 unsafe.Pointer) {
	// 2: (, const QPointF & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14setBrushOriginERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:207
// index:0
// void setBackground(const class QBrush &)
func (this *QPainter) SetBackground(bg unsafe.Pointer) {
	// 0: (, bg const QBrush &), (bg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), bg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:208
// index:0
// const QBrush & background()
func (this *QPainter) Background() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter10backgroundEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:210
// index:0
// qreal opacity()
func (this *QPainter) Opacity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter7opacityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:211
// index:0
// void setOpacity(qreal)
func (this *QPainter) SetOpacity(opacity float64) {
	// 0: (, opacity qreal), (&opacity)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10setOpacityEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:214
// index:0
// QRegion clipRegion()
func (this *QPainter) ClipRegion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter10clipRegionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:215
// index:0
// QPainterPath clipPath()
func (this *QPainter) ClipPath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter8clipPathEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:217
// index:0
// void setClipRect(const class QRectF &, Qt::ClipOperation)
func (this *QPainter) SetClipRect(arg0 unsafe.Pointer, op int) {
	// 0: (, const QRectF & arg0, op Qt::ClipOperation), (arg0, &op)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setClipRectERK6QRectFN2Qt13ClipOperationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &op)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:218
// index:1
// void setClipRect(const class QRect &, Qt::ClipOperation)
func (this *QPainter) SetClipRect_1(arg0 unsafe.Pointer, op int) {
	// 1: (, const QRect & arg0, op Qt::ClipOperation), (arg0, &op)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setClipRectERK5QRectN2Qt13ClipOperationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &op)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:219
// index:2
// inline
// void setClipRect(int, int, int, int, Qt::ClipOperation)
func (this *QPainter) SetClipRect_2(x int, y int, w int, h int, op int) {
	// 2: (, x int, y int, w int, h int, op Qt::ClipOperation), (&x, &y, &w, &h, &op)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setClipRectEiiiiN2Qt13ClipOperationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &op)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:221
// index:0
// void setClipRegion(const class QRegion &, Qt::ClipOperation)
func (this *QPainter) SetClipRegion(arg0 unsafe.Pointer, op int) {
	// 0: (, const QRegion & arg0, op Qt::ClipOperation), (arg0, &op)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13setClipRegionERK7QRegionN2Qt13ClipOperationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &op)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:223
// index:0
// void setClipPath(const class QPainterPath &, Qt::ClipOperation)
func (this *QPainter) SetClipPath(path unsafe.Pointer, op int) {
	// 0: (, path const QPainterPath &, op Qt::ClipOperation), (path, &op)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setClipPathERK12QPainterPathN2Qt13ClipOperationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, &op)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:225
// index:0
// void setClipping(_Bool)
func (this *QPainter) SetClipping(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setClippingEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:226
// index:0
// bool hasClipping()
func (this *QPainter) HasClipping() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11hasClippingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:228
// index:0
// QRectF clipBoundingRect()
func (this *QPainter) ClipBoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter16clipBoundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:230
// index:0
// void save()
func (this *QPainter) Save() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter4saveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:231
// index:0
// void restore()
func (this *QPainter) Restore() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7restoreEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:234
// index:0
// void setMatrix(const class QMatrix &, _Bool)
func (this *QPainter) SetMatrix(matrix unsafe.Pointer, combine bool) {
	// 0: (, matrix const QMatrix &, combine bool), (matrix, &combine)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9setMatrixERK7QMatrixb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix, &combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:235
// index:0
// const QMatrix & matrix()
func (this *QPainter) Matrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter6matrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:236
// index:0
// const QMatrix & deviceMatrix()
func (this *QPainter) DeviceMatrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter12deviceMatrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:237
// index:0
// void resetMatrix()
func (this *QPainter) ResetMatrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11resetMatrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:239
// index:0
// void setTransform(const class QTransform &, _Bool)
func (this *QPainter) SetTransform(transform unsafe.Pointer, combine bool) {
	// 0: (, transform const QTransform &, combine bool), (transform, &combine)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12setTransformERK10QTransformb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), transform, &combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:240
// index:0
// const QTransform & transform()
func (this *QPainter) Transform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter9transformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:241
// index:0
// const QTransform & deviceTransform()
func (this *QPainter) DeviceTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter15deviceTransformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:242
// index:0
// void resetTransform()
func (this *QPainter) ResetTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14resetTransformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:244
// index:0
// void setWorldMatrix(const class QMatrix &, _Bool)
func (this *QPainter) SetWorldMatrix(matrix unsafe.Pointer, combine bool) {
	// 0: (, matrix const QMatrix &, combine bool), (matrix, &combine)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14setWorldMatrixERK7QMatrixb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix, &combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:245
// index:0
// const QMatrix & worldMatrix()
func (this *QPainter) WorldMatrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11worldMatrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:247
// index:0
// void setWorldTransform(const class QTransform &, _Bool)
func (this *QPainter) SetWorldTransform(matrix unsafe.Pointer, combine bool) {
	// 0: (, matrix const QTransform &, combine bool), (matrix, &combine)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17setWorldTransformERK10QTransformb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix, &combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:248
// index:0
// const QTransform & worldTransform()
func (this *QPainter) WorldTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter14worldTransformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:250
// index:0
// QMatrix combinedMatrix()
func (this *QPainter) CombinedMatrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter14combinedMatrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:251
// index:0
// QTransform combinedTransform()
func (this *QPainter) CombinedTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter17combinedTransformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:253
// index:0
// void setMatrixEnabled(_Bool)
func (this *QPainter) SetMatrixEnabled(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter16setMatrixEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:254
// index:0
// bool matrixEnabled()
func (this *QPainter) MatrixEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter13matrixEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:256
// index:0
// void setWorldMatrixEnabled(_Bool)
func (this *QPainter) SetWorldMatrixEnabled(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter21setWorldMatrixEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:257
// index:0
// bool worldMatrixEnabled()
func (this *QPainter) WorldMatrixEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter18worldMatrixEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:259
// index:0
// void scale(qreal, qreal)
func (this *QPainter) Scale(sx float64, sy float64) {
	// 0: (, sx qreal, sy qreal), (&sx, &sy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter5scaleEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sx, &sy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:260
// index:0
// void shear(qreal, qreal)
func (this *QPainter) Shear(sh float64, sv float64) {
	// 0: (, sh qreal, sv qreal), (&sh, &sv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter5shearEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sh, &sv)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:261
// index:0
// void rotate(qreal)
func (this *QPainter) Rotate(a float64) {
	// 0: (, a qreal), (&a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter6rotateEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:263
// index:0
// void translate(const class QPointF &)
func (this *QPainter) Translate(offset unsafe.Pointer) {
	// 0: (, offset const QPointF &), (offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9translateERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:264
// index:1
// inline
// void translate(const class QPoint &)
func (this *QPainter) Translate_1(offset unsafe.Pointer) {
	// 1: (, offset const QPoint &), (offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9translateERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:265
// index:2
// inline
// void translate(qreal, qreal)
func (this *QPainter) Translate_2(dx float64, dy float64) {
	// 2: (, dx qreal, dy qreal), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9translateEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:267
// index:0
// QRect window()
func (this *QPainter) Window() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter6windowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:268
// index:0
// void setWindow(const class QRect &)
func (this *QPainter) SetWindow(window unsafe.Pointer) {
	// 0: (, window const QRect &), (window)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9setWindowERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), window)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:269
// index:1
// inline
// void setWindow(int, int, int, int)
func (this *QPainter) SetWindow_1(x int, y int, w int, h int) {
	// 1: (, x int, y int, w int, h int), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9setWindowEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:271
// index:0
// QRect viewport()
func (this *QPainter) Viewport() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter8viewportEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:272
// index:0
// void setViewport(const class QRect &)
func (this *QPainter) SetViewport(viewport unsafe.Pointer) {
	// 0: (, viewport const QRect &), (viewport)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setViewportERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), viewport)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:273
// index:1
// inline
// void setViewport(int, int, int, int)
func (this *QPainter) SetViewport_1(x int, y int, w int, h int) {
	// 1: (, x int, y int, w int, h int), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11setViewportEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:275
// index:0
// void setViewTransformEnabled(_Bool)
func (this *QPainter) SetViewTransformEnabled(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter23setViewTransformEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:276
// index:0
// bool viewTransformEnabled()
func (this *QPainter) ViewTransformEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter20viewTransformEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:279
// index:0
// void strokePath(const class QPainterPath &, const class QPen &)
func (this *QPainter) StrokePath(path unsafe.Pointer, pen unsafe.Pointer) {
	// 0: (, path const QPainterPath &, pen const QPen &), (path, pen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10strokePathERK12QPainterPathRK4QPen", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, pen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:280
// index:0
// void fillPath(const class QPainterPath &, const class QBrush &)
func (this *QPainter) FillPath(path unsafe.Pointer, brush unsafe.Pointer) {
	// 0: (, path const QPainterPath &, brush const QBrush &), (path, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillPathERK12QPainterPathRK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:281
// index:0
// void drawPath(const class QPainterPath &)
func (this *QPainter) DrawPath(path unsafe.Pointer) {
	// 0: (, path const QPainterPath &), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawPathERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:283
// index:0
// inline
// void drawPoint(const class QPointF &)
func (this *QPainter) DrawPoint(pt unsafe.Pointer) {
	// 0: (, pt const QPointF &), (pt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawPointERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:284
// index:1
// inline
// void drawPoint(const class QPoint &)
func (this *QPainter) DrawPoint_1(p unsafe.Pointer) {
	// 1: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawPointERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:285
// index:2
// inline
// void drawPoint(int, int)
func (this *QPainter) DrawPoint_2(x int, y int) {
	// 2: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawPointEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:287
// index:0
// void drawPoints(const class QPointF *, int)
func (this *QPainter) DrawPoints(points unsafe.Pointer, pointCount int) {
	// 0: (, points const QPointF *, pointCount int), (points, &pointCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPointsEPK7QPointFi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:288
// index:1
// inline
// void drawPoints(const class QPolygonF &)
func (this *QPainter) DrawPoints_1(points unsafe.Pointer) {
	// 1: (, points const QPolygonF &), (points)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPointsERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:289
// index:2
// void drawPoints(const class QPoint *, int)
func (this *QPainter) DrawPoints_2(points unsafe.Pointer, pointCount int) {
	// 2: (, points const QPoint *, pointCount int), (points, &pointCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPointsEPK6QPointi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:290
// index:3
// inline
// void drawPoints(const class QPolygon &)
func (this *QPainter) DrawPoints_3(points unsafe.Pointer) {
	// 3: (, points const QPolygon &), (points)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPointsERK8QPolygon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:292
// index:0
// inline
// void drawLine(const class QLineF &)
func (this *QPainter) DrawLine(line unsafe.Pointer) {
	// 0: (, line const QLineF &), (line)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawLineERK6QLineF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), line)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:293
// index:1
// inline
// void drawLine(const class QLine &)
func (this *QPainter) DrawLine_1(line unsafe.Pointer) {
	// 1: (, line const QLine &), (line)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawLineERK5QLine", ffiqt.FFI_TYPE_VOID, this.GetCthis(), line)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:294
// index:2
// inline
// void drawLine(int, int, int, int)
func (this *QPainter) DrawLine_2(x1 int, y1 int, x2 int, y2 int) {
	// 2: (, x1 int, y1 int, x2 int, y2 int), (&x1, &y1, &x2, &y2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawLineEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:295
// index:3
// inline
// void drawLine(const class QPoint &, const class QPoint &)
func (this *QPainter) DrawLine_3(p1 unsafe.Pointer, p2 unsafe.Pointer) {
	// 3: (, p1 const QPoint &, p2 const QPoint &), (p1, p2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawLineERK6QPointS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p1, p2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:296
// index:4
// inline
// void drawLine(const class QPointF &, const class QPointF &)
func (this *QPainter) DrawLine_4(p1 unsafe.Pointer, p2 unsafe.Pointer) {
	// 4: (, p1 const QPointF &, p2 const QPointF &), (p1, p2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawLineERK7QPointFS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p1, p2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:298
// index:0
// void drawLines(const class QLineF *, int)
func (this *QPainter) DrawLines(lines unsafe.Pointer, lineCount int) {
	// 0: (, lines const QLineF *, lineCount int), (lines, &lineCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK6QLineFi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), lines, &lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:300
// index:1
// void drawLines(const class QPointF *, int)
func (this *QPainter) DrawLines_1(pointPairs unsafe.Pointer, lineCount int) {
	// 1: (, pointPairs const QPointF *, lineCount int), (pointPairs, &lineCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK7QPointFi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pointPairs, &lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:302
// index:2
// void drawLines(const class QLine *, int)
func (this *QPainter) DrawLines_2(lines unsafe.Pointer, lineCount int) {
	// 2: (, lines const QLine *, lineCount int), (lines, &lineCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK5QLinei", ffiqt.FFI_TYPE_VOID, this.GetCthis(), lines, &lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:304
// index:3
// void drawLines(const class QPoint *, int)
func (this *QPainter) DrawLines_3(pointPairs unsafe.Pointer, lineCount int) {
	// 3: (, pointPairs const QPoint *, lineCount int), (pointPairs, &lineCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawLinesEPK6QPointi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pointPairs, &lineCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:307
// index:0
// inline
// void drawRect(const class QRectF &)
func (this *QPainter) DrawRect(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:308
// index:1
// inline
// void drawRect(int, int, int, int)
func (this *QPainter) DrawRect_1(x1 int, y1 int, w int, h int) {
	// 1: (, x1 int, y1 int, w int, h int), (&x1, &y1, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawRectEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x1, &y1, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:309
// index:2
// inline
// void drawRect(const class QRect &)
func (this *QPainter) DrawRect_2(rect unsafe.Pointer) {
	// 2: (, rect const QRect &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawRectERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:311
// index:0
// void drawRects(const class QRectF *, int)
func (this *QPainter) DrawRects(rects unsafe.Pointer, rectCount int) {
	// 0: (, rects const QRectF *, rectCount int), (rects, &rectCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawRectsEPK6QRectFi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rects, &rectCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:313
// index:1
// void drawRects(const class QRect *, int)
func (this *QPainter) DrawRects_1(rects unsafe.Pointer, rectCount int) {
	// 1: (, rects const QRect *, rectCount int), (rects, &rectCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawRectsEPK5QRecti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rects, &rectCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:316
// index:0
// void drawEllipse(const class QRectF &)
func (this *QPainter) DrawEllipse(r unsafe.Pointer) {
	// 0: (, r const QRectF &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:317
// index:1
// void drawEllipse(const class QRect &)
func (this *QPainter) DrawEllipse_1(r unsafe.Pointer) {
	// 1: (, r const QRect &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:318
// index:2
// inline
// void drawEllipse(int, int, int, int)
func (this *QPainter) DrawEllipse_2(x int, y int, w int, h int) {
	// 2: (, x int, y int, w int, h int), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawEllipseEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:320
// index:3
// inline
// void drawEllipse(const class QPointF &, qreal, qreal)
func (this *QPainter) DrawEllipse_3(center unsafe.Pointer, rx float64, ry float64) {
	// 3: (, center const QPointF &, rx qreal, ry qreal), (center, &rx, &ry)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK7QPointFdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), center, &rx, &ry)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:321
// index:4
// inline
// void drawEllipse(const class QPoint &, int, int)
func (this *QPainter) DrawEllipse_4(center unsafe.Pointer, rx int, ry int) {
	// 4: (, center const QPoint &, rx int, ry int), (center, &rx, &ry)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawEllipseERK6QPointii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), center, &rx, &ry)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:323
// index:0
// void drawPolyline(const class QPointF *, int)
func (this *QPainter) DrawPolyline(points unsafe.Pointer, pointCount int) {
	// 0: (, points const QPointF *, pointCount int), (points, &pointCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawPolylineEPK7QPointFi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:324
// index:1
// inline
// void drawPolyline(const class QPolygonF &)
func (this *QPainter) DrawPolyline_1(polyline unsafe.Pointer) {
	// 1: (, polyline const QPolygonF &), (polyline)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawPolylineERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polyline)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:325
// index:2
// void drawPolyline(const class QPoint *, int)
func (this *QPainter) DrawPolyline_2(points unsafe.Pointer, pointCount int) {
	// 2: (, points const QPoint *, pointCount int), (points, &pointCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawPolylineEPK6QPointi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:326
// index:3
// inline
// void drawPolyline(const class QPolygon &)
func (this *QPainter) DrawPolyline_3(polygon unsafe.Pointer) {
	// 3: (, polygon const QPolygon &), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawPolylineERK8QPolygon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:328
// index:0
// void drawPolygon(const class QPointF *, int, Qt::FillRule)
func (this *QPainter) DrawPolygon(points unsafe.Pointer, pointCount int, fillRule int) {
	// 0: (, points const QPointF *, pointCount int, fillRule Qt::FillRule), (points, &pointCount, &fillRule)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPolygonEPK7QPointFiN2Qt8FillRuleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount, &fillRule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:329
// index:1
// inline
// void drawPolygon(const class QPolygonF &, Qt::FillRule)
func (this *QPainter) DrawPolygon_1(polygon unsafe.Pointer, fillRule int) {
	// 1: (, polygon const QPolygonF &, fillRule Qt::FillRule), (polygon, &fillRule)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPolygonERK9QPolygonFN2Qt8FillRuleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon, &fillRule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:330
// index:2
// void drawPolygon(const class QPoint *, int, Qt::FillRule)
func (this *QPainter) DrawPolygon_2(points unsafe.Pointer, pointCount int, fillRule int) {
	// 2: (, points const QPoint *, pointCount int, fillRule Qt::FillRule), (points, &pointCount, &fillRule)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPolygonEPK6QPointiN2Qt8FillRuleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount, &fillRule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:331
// index:3
// inline
// void drawPolygon(const class QPolygon &, Qt::FillRule)
func (this *QPainter) DrawPolygon_3(polygon unsafe.Pointer, fillRule int) {
	// 3: (, polygon const QPolygon &, fillRule Qt::FillRule), (polygon, &fillRule)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPolygonERK8QPolygonN2Qt8FillRuleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon, &fillRule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:333
// index:0
// void drawConvexPolygon(const class QPointF *, int)
func (this *QPainter) DrawConvexPolygon(points unsafe.Pointer, pointCount int) {
	// 0: (, points const QPointF *, pointCount int), (points, &pointCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonEPK7QPointFi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:334
// index:1
// inline
// void drawConvexPolygon(const class QPolygonF &)
func (this *QPainter) DrawConvexPolygon_1(polygon unsafe.Pointer) {
	// 1: (, polygon const QPolygonF &), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:335
// index:2
// void drawConvexPolygon(const class QPoint *, int)
func (this *QPainter) DrawConvexPolygon_2(points unsafe.Pointer, pointCount int) {
	// 2: (, points const QPoint *, pointCount int), (points, &pointCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonEPK6QPointi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), points, &pointCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:336
// index:3
// inline
// void drawConvexPolygon(const class QPolygon &)
func (this *QPainter) DrawConvexPolygon_3(polygon unsafe.Pointer) {
	// 3: (, polygon const QPolygon &), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17drawConvexPolygonERK8QPolygon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:338
// index:0
// void drawArc(const class QRectF &, int, int)
func (this *QPainter) DrawArc(rect unsafe.Pointer, a int, alen int) {
	// 0: (, rect const QRectF &, a int, alen int), (rect, &a, &alen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawArcERK6QRectFii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &a, &alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:339
// index:1
// inline
// void drawArc(const class QRect &, int, int)
func (this *QPainter) DrawArc_1(arg0 unsafe.Pointer, a int, alen int) {
	// 1: (, const QRect & arg0, a int, alen int), (arg0, &a, &alen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawArcERK5QRectii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &a, &alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:340
// index:2
// inline
// void drawArc(int, int, int, int, int, int)
func (this *QPainter) DrawArc_2(x int, y int, w int, h int, a int, alen int) {
	// 2: (, x int, y int, w int, h int, a int, alen int), (&x, &y, &w, &h, &a, &alen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawArcEiiiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &a, &alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:342
// index:0
// void drawPie(const class QRectF &, int, int)
func (this *QPainter) DrawPie(rect unsafe.Pointer, a int, alen int) {
	// 0: (, rect const QRectF &, a int, alen int), (rect, &a, &alen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawPieERK6QRectFii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &a, &alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:343
// index:1
// inline
// void drawPie(int, int, int, int, int, int)
func (this *QPainter) DrawPie_1(x int, y int, w int, h int, a int, alen int) {
	// 1: (, x int, y int, w int, h int, a int, alen int), (&x, &y, &w, &h, &a, &alen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawPieEiiiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &a, &alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:344
// index:2
// inline
// void drawPie(const class QRect &, int, int)
func (this *QPainter) DrawPie_2(arg0 unsafe.Pointer, a int, alen int) {
	// 2: (, const QRect & arg0, a int, alen int), (arg0, &a, &alen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter7drawPieERK5QRectii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &a, &alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:346
// index:0
// void drawChord(const class QRectF &, int, int)
func (this *QPainter) DrawChord(rect unsafe.Pointer, a int, alen int) {
	// 0: (, rect const QRectF &, a int, alen int), (rect, &a, &alen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawChordERK6QRectFii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &a, &alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:347
// index:1
// inline
// void drawChord(int, int, int, int, int, int)
func (this *QPainter) DrawChord_1(x int, y int, w int, h int, a int, alen int) {
	// 1: (, x int, y int, w int, h int, a int, alen int), (&x, &y, &w, &h, &a, &alen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawChordEiiiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &a, &alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:348
// index:2
// inline
// void drawChord(const class QRect &, int, int)
func (this *QPainter) DrawChord_2(arg0 unsafe.Pointer, a int, alen int) {
	// 2: (, const QRect & arg0, a int, alen int), (arg0, &a, &alen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawChordERK5QRectii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &a, &alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:350
// index:0
// void drawRoundedRect(const class QRectF &, qreal, qreal, Qt::SizeMode)
func (this *QPainter) DrawRoundedRect(rect unsafe.Pointer, xRadius float64, yRadius float64, mode int) {
	// 0: (, rect const QRectF &, xRadius qreal, yRadius qreal, mode Qt::SizeMode), (rect, &xRadius, &yRadius, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectERK6QRectFddN2Qt8SizeModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &xRadius, &yRadius, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:352
// index:1
// inline
// void drawRoundedRect(int, int, int, int, qreal, qreal, Qt::SizeMode)
func (this *QPainter) DrawRoundedRect_1(x int, y int, w int, h int, xRadius float64, yRadius float64, mode int) {
	// 1: (, x int, y int, w int, h int, xRadius qreal, yRadius qreal, mode Qt::SizeMode), (&x, &y, &w, &h, &xRadius, &yRadius, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectEiiiiddN2Qt8SizeModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &xRadius, &yRadius, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:354
// index:2
// inline
// void drawRoundedRect(const class QRect &, qreal, qreal, Qt::SizeMode)
func (this *QPainter) DrawRoundedRect_2(rect unsafe.Pointer, xRadius float64, yRadius float64, mode int) {
	// 2: (, rect const QRect &, xRadius qreal, yRadius qreal, mode Qt::SizeMode), (rect, &xRadius, &yRadius, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawRoundedRectERK5QRectddN2Qt8SizeModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &xRadius, &yRadius, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:357
// index:0
// void drawRoundRect(const class QRectF &, int, int)
func (this *QPainter) DrawRoundRect(r unsafe.Pointer, xround int, yround int) {
	// 0: (, r const QRectF &, xround int, yround int), (r, &xround, &yround)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK6QRectFii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, &xround, &yround)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:358
// index:1
// inline
// void drawRoundRect(int, int, int, int, int, int)
func (this *QPainter) DrawRoundRect_1(x int, y int, w int, h int, arg4 int, arg5 int) {
	// 1: (, x int, y int, w int, h int, int arg4, int arg5), (&x, &y, &w, &h, &arg4, &arg5)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectEiiiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &arg4, &arg5)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:359
// index:2
// inline
// void drawRoundRect(const class QRect &, int, int)
func (this *QPainter) DrawRoundRect_2(r unsafe.Pointer, xround int, yround int) {
	// 2: (, r const QRect &, xround int, yround int), (r, &xround, &yround)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13drawRoundRectERK5QRectii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, &xround, &yround)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:361
// index:0
// void drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
func (this *QPainter) DrawTiledPixmap(rect unsafe.Pointer, pm unsafe.Pointer, offset unsafe.Pointer) {
	// 0: (, rect const QRectF &, pm const QPixmap &, offset const QPointF &), (rect, pm, offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, pm, offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:362
// index:1
// inline
// void drawTiledPixmap(int, int, int, int, const class QPixmap &, int, int)
func (this *QPainter) DrawTiledPixmap_1(x int, y int, w int, h int, arg4 unsafe.Pointer, sx int, sy int) {
	// 1: (, x int, y int, w int, h int, const QPixmap & arg4, sx int, sy int), (&x, &y, &w, &h, arg4, &sx, &sy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, arg4, &sx, &sy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:363
// index:2
// inline
// void drawTiledPixmap(const class QRect &, const class QPixmap &, const class QPoint &)
func (this *QPainter) DrawTiledPixmap_2(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 unsafe.Pointer) {
	// 2: (, const QRect & arg0, const QPixmap & arg1, const QPoint & arg2), (arg0, arg1, arg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, arg1, arg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:365
// index:0
// void drawPicture(const class QPointF &, const class QPicture &)
func (this *QPainter) DrawPicture(p unsafe.Pointer, picture unsafe.Pointer) {
	// 0: (, p const QPointF &, picture const QPicture &), (p, picture)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPictureERK7QPointFRK8QPicture", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, picture)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:366
// index:1
// inline
// void drawPicture(int, int, const class QPicture &)
func (this *QPainter) DrawPicture_1(x int, y int, picture unsafe.Pointer) {
	// 1: (, x int, y int, picture const QPicture &), (&x, &y, picture)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPictureEiiRK8QPicture", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, picture)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:367
// index:2
// inline
// void drawPicture(const class QPoint &, const class QPicture &)
func (this *QPainter) DrawPicture_2(p unsafe.Pointer, picture unsafe.Pointer) {
	// 2: (, p const QPoint &, picture const QPicture &), (p, picture)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter11drawPictureERK6QPointRK8QPicture", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, picture)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:370
// index:0
// void drawPixmap(const class QRectF &, const class QPixmap &, const class QRectF &)
func (this *QPainter) DrawPixmap(targetRect unsafe.Pointer, pixmap unsafe.Pointer, sourceRect unsafe.Pointer) {
	// 0: (, targetRect const QRectF &, pixmap const QPixmap &, sourceRect const QRectF &), (targetRect, pixmap, sourceRect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QRectFRK7QPixmapS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), targetRect, pixmap, sourceRect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:371
// index:1
// inline
// void drawPixmap(const class QRect &, const class QPixmap &, const class QRect &)
func (this *QPainter) DrawPixmap_1(targetRect unsafe.Pointer, pixmap unsafe.Pointer, sourceRect unsafe.Pointer) {
	// 1: (, targetRect const QRect &, pixmap const QPixmap &, sourceRect const QRect &), (targetRect, pixmap, sourceRect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK5QRectRK7QPixmapS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), targetRect, pixmap, sourceRect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:372
// index:2
// inline
// void drawPixmap(int, int, int, int, const class QPixmap &, int, int, int, int)
func (this *QPainter) DrawPixmap_2(x int, y int, w int, h int, pm unsafe.Pointer, sx int, sy int, sw int, sh int) {
	// 2: (, x int, y int, w int, h int, pm const QPixmap &, sx int, sy int, sw int, sh int), (&x, &y, &w, &h, pm, &sx, &sy, &sw, &sh)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiiiRK7QPixmapiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, pm, &sx, &sy, &sw, &sh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:374
// index:3
// inline
// void drawPixmap(int, int, const class QPixmap &, int, int, int, int)
func (this *QPainter) DrawPixmap_3(x int, y int, pm unsafe.Pointer, sx int, sy int, sw int, sh int) {
	// 3: (, x int, y int, pm const QPixmap &, sx int, sy int, sw int, sh int), (&x, &y, pm, &sx, &sy, &sw, &sh)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiRK7QPixmapiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, pm, &sx, &sy, &sw, &sh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:376
// index:4
// inline
// void drawPixmap(const class QPointF &, const class QPixmap &, const class QRectF &)
func (this *QPainter) DrawPixmap_4(p unsafe.Pointer, pm unsafe.Pointer, sr unsafe.Pointer) {
	// 4: (, p const QPointF &, pm const QPixmap &, sr const QRectF &), (p, pm, sr)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmapRK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, pm, sr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:377
// index:5
// inline
// void drawPixmap(const class QPoint &, const class QPixmap &, const class QRect &)
func (this *QPainter) DrawPixmap_5(p unsafe.Pointer, pm unsafe.Pointer, sr unsafe.Pointer) {
	// 5: (, p const QPoint &, pm const QPixmap &, sr const QRect &), (p, pm, sr)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QPointRK7QPixmapRK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, pm, sr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:378
// index:6
// void drawPixmap(const class QPointF &, const class QPixmap &)
func (this *QPainter) DrawPixmap_6(p unsafe.Pointer, pm unsafe.Pointer) {
	// 6: (, p const QPointF &, pm const QPixmap &), (p, pm)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, pm)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:379
// index:7
// inline
// void drawPixmap(const class QPoint &, const class QPixmap &)
func (this *QPainter) DrawPixmap_7(p unsafe.Pointer, pm unsafe.Pointer) {
	// 7: (, p const QPoint &, pm const QPixmap &), (p, pm)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK6QPointRK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, pm)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:380
// index:8
// inline
// void drawPixmap(int, int, const class QPixmap &)
func (this *QPainter) DrawPixmap_8(x int, y int, pm unsafe.Pointer) {
	// 8: (, x int, y int, pm const QPixmap &), (&x, &y, pm)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiRK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, pm)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:381
// index:9
// inline
// void drawPixmap(const class QRect &, const class QPixmap &)
func (this *QPainter) DrawPixmap_9(r unsafe.Pointer, pm unsafe.Pointer) {
	// 9: (, r const QRect &, pm const QPixmap &), (r, pm)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapERK5QRectRK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, pm)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:382
// index:10
// inline
// void drawPixmap(int, int, int, int, const class QPixmap &)
func (this *QPainter) DrawPixmap_10(x int, y int, w int, h int, pm unsafe.Pointer) {
	// 10: (, x int, y int, w int, h int, pm const QPixmap &), (&x, &y, &w, &h, pm)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10drawPixmapEiiiiRK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, pm)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:384
// index:0
// void drawPixmapFragments(const class QPainter::PixmapFragment *, int, const class QPixmap &, QPainter::PixmapFragmentHints)
func (this *QPainter) DrawPixmapFragments(fragments unsafe.Pointer, fragmentCount int, pixmap unsafe.Pointer, hints int) {
	// 0: (, fragments const QPainter::PixmapFragment *, fragmentCount int, pixmap const QPixmap &, QFlags<QPainter::PixmapFragmentHint> hints), (fragments, &fragmentCount, pixmap, hints)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter19drawPixmapFragmentsEPKNS_14PixmapFragmentEiRK7QPixmap6QFlagsINS_18PixmapFragmentHintEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fragments, &fragmentCount, pixmap, hints)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:387
// index:0
// void drawImage(const class QRectF &, const class QImage &, const class QRectF &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage(targetRect unsafe.Pointer, image unsafe.Pointer, sourceRect unsafe.Pointer, flags int) {
	// 0: (, targetRect const QRectF &, image const QImage &, sourceRect const QRectF &, QFlags<Qt::ImageConversionFlag> flags), (targetRect, image, sourceRect, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QRectFRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), targetRect, image, sourceRect, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:389
// index:1
// inline
// void drawImage(const class QRect &, const class QImage &, const class QRect &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_1(targetRect unsafe.Pointer, image unsafe.Pointer, sourceRect unsafe.Pointer, flags int) {
	// 1: (, targetRect const QRect &, image const QImage &, sourceRect const QRect &, QFlags<Qt::ImageConversionFlag> flags), (targetRect, image, sourceRect, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK5QRectRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), targetRect, image, sourceRect, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:391
// index:2
// inline
// void drawImage(const class QPointF &, const class QImage &, const class QRectF &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_2(p unsafe.Pointer, image unsafe.Pointer, sr unsafe.Pointer, flags int) {
	// 2: (, p const QPointF &, image const QImage &, sr const QRectF &, QFlags<Qt::ImageConversionFlag> flags), (p, image, sr, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK7QPointFRK6QImageRK6QRectF6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, image, sr, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:393
// index:3
// inline
// void drawImage(const class QPoint &, const class QImage &, const class QRect &, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_3(p unsafe.Pointer, image unsafe.Pointer, sr unsafe.Pointer, flags int) {
	// 3: (, p const QPoint &, image const QImage &, sr const QRect &, QFlags<Qt::ImageConversionFlag> flags), (p, image, sr, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QPointRK6QImageRK5QRect6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, image, sr, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:395
// index:4
// inline
// void drawImage(const class QRectF &, const class QImage &)
func (this *QPainter) DrawImage_4(r unsafe.Pointer, image unsafe.Pointer) {
	// 4: (, r const QRectF &, image const QImage &), (r, image)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QRectFRK6QImage", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, image)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:396
// index:5
// inline
// void drawImage(const class QRect &, const class QImage &)
func (this *QPainter) DrawImage_5(r unsafe.Pointer, image unsafe.Pointer) {
	// 5: (, r const QRect &, image const QImage &), (r, image)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK5QRectRK6QImage", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, image)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:397
// index:6
// void drawImage(const class QPointF &, const class QImage &)
func (this *QPainter) DrawImage_6(p unsafe.Pointer, image unsafe.Pointer) {
	// 6: (, p const QPointF &, image const QImage &), (p, image)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK7QPointFRK6QImage", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, image)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:398
// index:7
// inline
// void drawImage(const class QPoint &, const class QImage &)
func (this *QPainter) DrawImage_7(p unsafe.Pointer, image unsafe.Pointer) {
	// 7: (, p const QPoint &, image const QImage &), (p, image)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageERK6QPointRK6QImage", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, image)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:399
// index:8
// inline
// void drawImage(int, int, const class QImage &, int, int, int, int, Qt::ImageConversionFlags)
func (this *QPainter) DrawImage_8(x int, y int, image unsafe.Pointer, sx int, sy int, sw int, sh int, flags int) {
	// 8: (, x int, y int, image const QImage &, sx int, sy int, sw int, sh int, QFlags<Qt::ImageConversionFlag> flags), (&x, &y, image, &sx, &sy, &sw, &sh, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9drawImageEiiRK6QImageiiii6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, image, &sx, &sy, &sw, &sh, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:402
// index:0
// void setLayoutDirection(Qt::LayoutDirection)
func (this *QPainter) SetLayoutDirection(direction int) {
	// 0: (, direction Qt::LayoutDirection), (&direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter18setLayoutDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:403
// index:0
// Qt::LayoutDirection layoutDirection()
func (this *QPainter) LayoutDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter15layoutDirectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:406
// index:0
// void drawGlyphRun(const class QPointF &, const class QGlyphRun &)
func (this *QPainter) DrawGlyphRun(position unsafe.Pointer, glyphRun unsafe.Pointer) {
	// 0: (, position const QPointF &, glyphRun const QGlyphRun &), (position, glyphRun)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawGlyphRunERK7QPointFRK9QGlyphRun", ffiqt.FFI_TYPE_VOID, this.GetCthis(), position, glyphRun)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:409
// index:0
// void drawStaticText(const class QPointF &, const class QStaticText &)
func (this *QPainter) DrawStaticText(topLeftPosition unsafe.Pointer, staticText unsafe.Pointer) {
	// 0: (, topLeftPosition const QPointF &, staticText const QStaticText &), (topLeftPosition, staticText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextERK7QPointFRK11QStaticText", ffiqt.FFI_TYPE_VOID, this.GetCthis(), topLeftPosition, staticText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:410
// index:1
// inline
// void drawStaticText(const class QPoint &, const class QStaticText &)
func (this *QPainter) DrawStaticText_1(topLeftPosition unsafe.Pointer, staticText unsafe.Pointer) {
	// 1: (, topLeftPosition const QPoint &, staticText const QStaticText &), (topLeftPosition, staticText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextERK6QPointRK11QStaticText", ffiqt.FFI_TYPE_VOID, this.GetCthis(), topLeftPosition, staticText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:411
// index:2
// inline
// void drawStaticText(int, int, const class QStaticText &)
func (this *QPainter) DrawStaticText_2(left int, top int, staticText unsafe.Pointer) {
	// 2: (, left int, top int, staticText const QStaticText &), (&left, &top, staticText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14drawStaticTextEiiRK11QStaticText", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &left, &top, staticText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:413
// index:0
// void drawText(const class QPointF &, const class QString &)
func (this *QPainter) DrawText(p unsafe.Pointer, s unsafe.Pointer) {
	// 0: (, p const QPointF &, s const QString &), (p, s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK7QPointFRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:414
// index:1
// inline
// void drawText(const class QPoint &, const class QString &)
func (this *QPainter) DrawText_1(p unsafe.Pointer, s unsafe.Pointer) {
	// 1: (, p const QPoint &, s const QString &), (p, s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QPointRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:415
// index:2
// inline
// void drawText(int, int, const class QString &)
func (this *QPainter) DrawText_2(x int, y int, s unsafe.Pointer) {
	// 2: (, x int, y int, s const QString &), (&x, &y, s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextEiiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:417
// index:3
// void drawText(const class QPointF &, const class QString &, int, int)
func (this *QPainter) DrawText_3(p unsafe.Pointer, str unsafe.Pointer, tf int, justificationPadding int) {
	// 3: (, p const QPointF &, str const QString &, tf int, justificationPadding int), (p, str, &tf, &justificationPadding)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK7QPointFRK7QStringii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, str, &tf, &justificationPadding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:419
// index:4
// void drawText(const class QRectF &, int, const class QString &, class QRectF *)
func (this *QPainter) DrawText_4(r unsafe.Pointer, flags int, text unsafe.Pointer, br unsafe.Pointer) {
	// 4: (, r const QRectF &, flags int, text const QString &, br QRectF *), (r, &flags, text, br)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, &flags, text, br)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:420
// index:5
// void drawText(const class QRect &, int, const class QString &, class QRect *)
func (this *QPainter) DrawText_5(r unsafe.Pointer, flags int, text unsafe.Pointer, br unsafe.Pointer) {
	// 5: (, r const QRect &, flags int, text const QString &, br QRect *), (r, &flags, text, br)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, &flags, text, br)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:421
// index:6
// inline
// void drawText(int, int, int, int, int, const class QString &, class QRect *)
func (this *QPainter) DrawText_6(x int, y int, w int, h int, flags int, text unsafe.Pointer, br unsafe.Pointer) {
	// 6: (, x int, y int, w int, h int, flags int, text const QString &, br QRect *), (&x, &y, &w, &h, &flags, text, br)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &flags, text, br)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:423
// index:7
// void drawText(const class QRectF &, const class QString &, const class QTextOption &)
func (this *QPainter) DrawText_7(r unsafe.Pointer, text unsafe.Pointer, o unsafe.Pointer) {
	// 7: (, r const QRectF &, text const QString &, o const QTextOption &), (r, text, o)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, text, o)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:425
// index:0
// QRectF boundingRect(const class QRectF &, int, const class QString &)
func (this *QPainter) BoundingRect(rect unsafe.Pointer, flags int, text unsafe.Pointer) {
	// 0: (, rect const QRectF &, flags int, text const QString &), (rect, &flags, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK6QRectFiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &flags, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:426
// index:1
// QRect boundingRect(const class QRect &, int, const class QString &)
func (this *QPainter) BoundingRect_1(rect unsafe.Pointer, flags int, text unsafe.Pointer) {
	// 1: (, rect const QRect &, flags int, text const QString &), (rect, &flags, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK5QRectiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &flags, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:427
// index:2
// inline
// QRect boundingRect(int, int, int, int, int, const class QString &)
func (this *QPainter) BoundingRect_2(x int, y int, w int, h int, flags int, text unsafe.Pointer) {
	// 2: (, x int, y int, w int, h int, flags int, text const QString &), (&x, &y, &w, &h, &flags, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12boundingRectEiiiiiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &flags, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:429
// index:3
// QRectF boundingRect(const class QRectF &, const class QString &, const class QTextOption &)
func (this *QPainter) BoundingRect_3(rect unsafe.Pointer, text unsafe.Pointer, o unsafe.Pointer) {
	// 3: (, rect const QRectF &, text const QString &, o const QTextOption &), (rect, text, o)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, text, o)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:431
// index:0
// void drawTextItem(const class QPointF &, const class QTextItem &)
func (this *QPainter) DrawTextItem(p unsafe.Pointer, ti unsafe.Pointer) {
	// 0: (, p const QPointF &, ti const QTextItem &), (p, ti)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawTextItemERK7QPointFRK9QTextItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, ti)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:432
// index:1
// inline
// void drawTextItem(int, int, const class QTextItem &)
func (this *QPainter) DrawTextItem_1(x int, y int, ti unsafe.Pointer) {
	// 1: (, x int, y int, ti const QTextItem &), (&x, &y, ti)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawTextItemEiiRK9QTextItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, ti)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:433
// index:2
// inline
// void drawTextItem(const class QPoint &, const class QTextItem &)
func (this *QPainter) DrawTextItem_2(p unsafe.Pointer, ti unsafe.Pointer) {
	// 2: (, p const QPoint &, ti const QTextItem &), (p, ti)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter12drawTextItemERK6QPointRK9QTextItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p, ti)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:435
// index:0
// void fillRect(const class QRectF &, const class QBrush &)
func (this *QPainter) FillRect(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, const QRectF & arg0, const QBrush & arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFRK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:436
// index:1
// inline
// void fillRect(int, int, int, int, const class QBrush &)
func (this *QPainter) FillRect_1(x int, y int, w int, h int, arg4 unsafe.Pointer) {
	// 1: (, x int, y int, w int, h int, const QBrush & arg4), (&x, &y, &w, &h, arg4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiRK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, arg4)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:437
// index:2
// void fillRect(const class QRect &, const class QBrush &)
func (this *QPainter) FillRect_2(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 2: (, const QRect & arg0, const QBrush & arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectRK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:439
// index:3
// void fillRect(const class QRectF &, const class QColor &)
func (this *QPainter) FillRect_3(arg0 unsafe.Pointer, color unsafe.Pointer) {
	// 3: (, const QRectF & arg0, color const QColor &), (arg0, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFRK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:440
// index:4
// inline
// void fillRect(int, int, int, int, const class QColor &)
func (this *QPainter) FillRect_4(x int, y int, w int, h int, color unsafe.Pointer) {
	// 4: (, x int, y int, w int, h int, color const QColor &), (&x, &y, &w, &h, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiRK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:441
// index:5
// void fillRect(const class QRect &, const class QColor &)
func (this *QPainter) FillRect_5(arg0 unsafe.Pointer, color unsafe.Pointer) {
	// 5: (, const QRect & arg0, color const QColor &), (arg0, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectRK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:443
// index:6
// inline
// void fillRect(int, int, int, int, Qt::GlobalColor)
func (this *QPainter) FillRect_6(x int, y int, w int, h int, c int) {
	// 6: (, x int, y int, w int, h int, c Qt::GlobalColor), (&x, &y, &w, &h, &c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiN2Qt11GlobalColorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:444
// index:7
// inline
// void fillRect(const class QRect &, Qt::GlobalColor)
func (this *QPainter) FillRect_7(r unsafe.Pointer, c int) {
	// 7: (, r const QRect &, c Qt::GlobalColor), (r, &c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectN2Qt11GlobalColorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:445
// index:8
// inline
// void fillRect(const class QRectF &, Qt::GlobalColor)
func (this *QPainter) FillRect_8(r unsafe.Pointer, c int) {
	// 8: (, r const QRectF &, c Qt::GlobalColor), (r, &c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFN2Qt11GlobalColorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:447
// index:9
// inline
// void fillRect(int, int, int, int, Qt::BrushStyle)
func (this *QPainter) FillRect_9(x int, y int, w int, h int, style int) {
	// 9: (, x int, y int, w int, h int, style Qt::BrushStyle), (&x, &y, &w, &h, &style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectEiiiiN2Qt10BrushStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:448
// index:10
// inline
// void fillRect(const class QRect &, Qt::BrushStyle)
func (this *QPainter) FillRect_10(r unsafe.Pointer, style int) {
	// 10: (, r const QRect &, style Qt::BrushStyle), (r, &style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK5QRectN2Qt10BrushStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:449
// index:11
// inline
// void fillRect(const class QRectF &, Qt::BrushStyle)
func (this *QPainter) FillRect_11(r unsafe.Pointer, style int) {
	// 11: (, r const QRectF &, style Qt::BrushStyle), (r, &style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter8fillRectERK6QRectFN2Qt10BrushStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:451
// index:0
// void eraseRect(const class QRectF &)
func (this *QPainter) EraseRect(arg0 unsafe.Pointer) {
	// 0: (, const QRectF & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9eraseRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:452
// index:1
// inline
// void eraseRect(int, int, int, int)
func (this *QPainter) EraseRect_1(x int, y int, w int, h int) {
	// 1: (, x int, y int, w int, h int), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9eraseRectEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:453
// index:2
// inline
// void eraseRect(const class QRect &)
func (this *QPainter) EraseRect_2(arg0 unsafe.Pointer) {
	// 2: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter9eraseRectERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:455
// index:0
// void setRenderHint(enum QPainter::RenderHint, _Bool)
func (this *QPainter) SetRenderHint(hint int, on bool) {
	// 0: (, hint QPainter::RenderHint, on bool), (&hint, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13setRenderHintENS_10RenderHintEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &hint, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:456
// index:0
// void setRenderHints(QPainter::RenderHints, _Bool)
func (this *QPainter) SetRenderHints(hints int, on bool) {
	// 0: (, QFlags<QPainter::RenderHint> hints, on bool), (hints, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter14setRenderHintsE6QFlagsINS_10RenderHintEEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), hints, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:457
// index:0
// QPainter::RenderHints renderHints()
func (this *QPainter) RenderHints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11renderHintsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:458
// index:0
// inline
// bool testRenderHint(enum QPainter::RenderHint)
func (this *QPainter) TestRenderHint(hint int) {
	// 0: (, hint QPainter::RenderHint), (&hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter14testRenderHintENS_10RenderHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:460
// index:0
// QPaintEngine * paintEngine()
func (this *QPainter) PaintEngine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPainter11paintEngineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:462
// index:0
// static
// void setRedirected(const class QPaintDevice *, class QPaintDevice *, const class QPoint &)
func (this *QPainter) SetRedirected(device unsafe.Pointer, replacement unsafe.Pointer, offset unsafe.Pointer) {
	// 0: (device const QPaintDevice *, replacement QPaintDevice *, offset const QPoint &), (device, replacement, offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter13setRedirectedEPK12QPaintDevicePS0_RK6QPoint", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPainter_SetRedirected(device unsafe.Pointer, replacement unsafe.Pointer, offset unsafe.Pointer) {
	// 0: (device const QPaintDevice *, replacement QPaintDevice *, offset const QPoint &), (device, replacement, offset)
	var nilthis *QPainter
	nilthis.SetRedirected(device, replacement, offset)
}

// /usr/include/qt/QtGui/qpainter.h:464
// index:0
// static
// QPaintDevice * redirected(const class QPaintDevice *, class QPoint *)
func (this *QPainter) Redirected(device unsafe.Pointer, offset unsafe.Pointer) {
	// 0: (device const QPaintDevice *, offset QPoint *), (device, offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter10redirectedEPK12QPaintDeviceP6QPoint", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPainter_Redirected(device unsafe.Pointer, offset unsafe.Pointer) {
	// 0: (device const QPaintDevice *, offset QPoint *), (device, offset)
	var nilthis *QPainter
	nilthis.Redirected(device, offset)
}

// /usr/include/qt/QtGui/qpainter.h:465
// index:0
// static
// void restoreRedirected(const class QPaintDevice *)
func (this *QPainter) RestoreRedirected(device unsafe.Pointer) {
	// 0: (device const QPaintDevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17restoreRedirectedEPK12QPaintDevice", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPainter_RestoreRedirected(device unsafe.Pointer) {
	// 0: (device const QPaintDevice *), (device)
	var nilthis *QPainter
	nilthis.RestoreRedirected(device)
}

// /usr/include/qt/QtGui/qpainter.h:467
// index:0
// void beginNativePainting()
func (this *QPainter) BeginNativePainting() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter19beginNativePaintingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainter.h:468
// index:0
// void endNativePainting()
func (this *QPainter) EndNativePainting() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPainter17endNativePaintingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
