package qtquick

// /usr/include/qt/QtQuick/qsggeometry.h
// #include <qsggeometry.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QSGGeometry struct {
	*qtrt.CObject
}

func (this *QSGGeometry) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSGGeometry) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSGGeometryFromPointer(cthis unsafe.Pointer) *QSGGeometry {
	return &QSGGeometry{&qtrt.CObject{cthis}}
}
func (*QSGGeometry) NewFromPointer(cthis unsafe.Pointer) *QSGGeometry {
	return NewQSGGeometryFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsggeometry.h:146
// index:0
// Public virtual
// void ~QSGGeometry()
func DeleteQSGGeometry(*QSGGeometry) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:149
// index:0
// Public
// void setDrawingMode(unsigned int)
func (this *QSGGeometry) SetDrawingMode(mode uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry14setDrawingModeEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:150
// index:0
// Public inline
// unsigned int drawingMode()
func (this *QSGGeometry) DrawingMode() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry11drawingModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("uint", rv).(uint) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:152
// index:0
// Public
// void allocate(int, int)
func (this *QSGGeometry) Allocate(vertexCount int, indexCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry8allocateEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), vertexCount, indexCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:154
// index:0
// Public inline
// int vertexCount()
func (this *QSGGeometry) VertexCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry11vertexCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:156
// index:0
// Public inline
// void * vertexData()
func (this *QSGGeometry) VertexData() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry10vertexDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:161
// index:1
// Public inline
// const void * vertexData()
func (this *QSGGeometry) VertexData_1() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry10vertexDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:157
// index:0
// Public inline
// QSGGeometry::Point2D * vertexDataAsPoint2D()
func (this *QSGGeometry) VertexDataAsPoint2D() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry19vertexDataAsPoint2DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:162
// index:1
// Public inline
// const QSGGeometry::Point2D * vertexDataAsPoint2D()
func (this *QSGGeometry) VertexDataAsPoint2D_1() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry19vertexDataAsPoint2DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:158
// index:0
// Public inline
// QSGGeometry::TexturedPoint2D * vertexDataAsTexturedPoint2D()
func (this *QSGGeometry) VertexDataAsTexturedPoint2D() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry27vertexDataAsTexturedPoint2DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:163
// index:1
// Public inline
// const QSGGeometry::TexturedPoint2D * vertexDataAsTexturedPoint2D()
func (this *QSGGeometry) VertexDataAsTexturedPoint2D_1() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry27vertexDataAsTexturedPoint2DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:159
// index:0
// Public inline
// QSGGeometry::ColoredPoint2D * vertexDataAsColoredPoint2D()
func (this *QSGGeometry) VertexDataAsColoredPoint2D() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry26vertexDataAsColoredPoint2DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:164
// index:1
// Public inline
// const QSGGeometry::ColoredPoint2D * vertexDataAsColoredPoint2D()
func (this *QSGGeometry) VertexDataAsColoredPoint2D_1() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry26vertexDataAsColoredPoint2DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:166
// index:0
// Public inline
// int indexType()
func (this *QSGGeometry) IndexType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry9indexTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:168
// index:0
// Public inline
// int indexCount()
func (this *QSGGeometry) IndexCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry10indexCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:170
// index:0
// Public
// void * indexData()
func (this *QSGGeometry) IndexData() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry9indexDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:176
// index:1
// Public
// const void * indexData()
func (this *QSGGeometry) IndexData_1() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry9indexDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:171
// index:0
// Public inline
// uint * indexDataAsUInt()
func (this *QSGGeometry) IndexDataAsUInt() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry15indexDataAsUIntEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:177
// index:1
// Public inline
// const uint * indexDataAsUInt()
func (this *QSGGeometry) IndexDataAsUInt_1() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry15indexDataAsUIntEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:172
// index:0
// Public inline
// quint16 * indexDataAsUShort()
func (this *QSGGeometry) IndexDataAsUShort() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry17indexDataAsUShortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:178
// index:1
// Public inline
// const quint16 * indexDataAsUShort()
func (this *QSGGeometry) IndexDataAsUShort_1() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry17indexDataAsUShortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:174
// index:0
// Public inline
// int sizeOfIndex()
func (this *QSGGeometry) SizeOfIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry11sizeOfIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:180
// index:0
// Public inline
// int attributeCount()
func (this *QSGGeometry) AttributeCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry14attributeCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:181
// index:0
// Public inline
// const QSGGeometry::Attribute * attributes()
func (this *QSGGeometry) Attributes() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry10attributesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:182
// index:0
// Public inline
// int sizeOfVertex()
func (this *QSGGeometry) SizeOfVertex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry12sizeOfVertexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:184
// index:0
// Public static
// void updateRectGeometry(QSGGeometry *, const QRectF &)
func (this *QSGGeometry) UpdateRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF) {
	var convArg0 = g.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry18updateRectGeometryEPS_RK6QRectF", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QSGGeometry_UpdateRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF) {
	var nilthis *QSGGeometry
	nilthis.UpdateRectGeometry(g, rect)
}

// /usr/include/qt/QtQuick/qsggeometry.h:185
// index:0
// Public static
// void updateTexturedRectGeometry(QSGGeometry *, const QRectF &, const QRectF &)
func (this *QSGGeometry) UpdateTexturedRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF, sourceRect *qtcore.QRectF) {
	var convArg0 = g.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg2 = sourceRect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry26updateTexturedRectGeometryEPS_RK6QRectFS3_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}
func QSGGeometry_UpdateTexturedRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF, sourceRect *qtcore.QRectF) {
	var nilthis *QSGGeometry
	nilthis.UpdateTexturedRectGeometry(g, rect, sourceRect)
}

// /usr/include/qt/QtQuick/qsggeometry.h:186
// index:0
// Public static
// void updateColoredRectGeometry(QSGGeometry *, const QRectF &)
func (this *QSGGeometry) UpdateColoredRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF) {
	var convArg0 = g.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry25updateColoredRectGeometryEPS_RK6QRectF", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QSGGeometry_UpdateColoredRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF) {
	var nilthis *QSGGeometry
	nilthis.UpdateColoredRectGeometry(g, rect)
}

// /usr/include/qt/QtQuick/qsggeometry.h:188
// index:0
// Public
// void setIndexDataPattern(enum QSGGeometry::DataPattern)
func (this *QSGGeometry) SetIndexDataPattern(p int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry19setIndexDataPatternENS_11DataPatternE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:189
// index:0
// Public inline
// QSGGeometry::DataPattern indexDataPattern()
func (this *QSGGeometry) IndexDataPattern() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry16indexDataPatternEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:191
// index:0
// Public
// void setVertexDataPattern(enum QSGGeometry::DataPattern)
func (this *QSGGeometry) SetVertexDataPattern(p int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry20setVertexDataPatternENS_11DataPatternE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:192
// index:0
// Public inline
// QSGGeometry::DataPattern vertexDataPattern()
func (this *QSGGeometry) VertexDataPattern() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry17vertexDataPatternEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:194
// index:0
// Public
// void markIndexDataDirty()
func (this *QSGGeometry) MarkIndexDataDirty() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry18markIndexDataDirtyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:195
// index:0
// Public
// void markVertexDataDirty()
func (this *QSGGeometry) MarkVertexDataDirty() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry19markVertexDataDirtyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:197
// index:0
// Public
// float lineWidth()
func (this *QSGGeometry) LineWidth() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGGeometry9lineWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:198
// index:0
// Public
// void setLineWidth(float)
func (this *QSGGeometry) SetLineWidth(w float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGGeometry12setLineWidthEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

type QSGGeometry__AttributeType = int

const QSGGeometry__UnknownAttribute QSGGeometry__AttributeType = 0
const QSGGeometry__PositionAttribute QSGGeometry__AttributeType = 1
const QSGGeometry__ColorAttribute QSGGeometry__AttributeType = 2
const QSGGeometry__TexCoordAttribute QSGGeometry__AttributeType = 3
const QSGGeometry__TexCoord1Attribute QSGGeometry__AttributeType = 4
const QSGGeometry__TexCoord2Attribute QSGGeometry__AttributeType = 5

type QSGGeometry__DataPattern = int

const QSGGeometry__AlwaysUploadPattern QSGGeometry__DataPattern = 0
const QSGGeometry__StreamPattern QSGGeometry__DataPattern = 1
const QSGGeometry__DynamicPattern QSGGeometry__DataPattern = 2
const QSGGeometry__StaticPattern QSGGeometry__DataPattern = 3

type QSGGeometry__DrawingMode = int

const QSGGeometry__DrawPoints QSGGeometry__DrawingMode = 0
const QSGGeometry__DrawLines QSGGeometry__DrawingMode = 1
const QSGGeometry__DrawLineLoop QSGGeometry__DrawingMode = 2
const QSGGeometry__DrawLineStrip QSGGeometry__DrawingMode = 3
const QSGGeometry__DrawTriangles QSGGeometry__DrawingMode = 4
const QSGGeometry__DrawTriangleStrip QSGGeometry__DrawingMode = 5
const QSGGeometry__DrawTriangleFan QSGGeometry__DrawingMode = 6

type QSGGeometry__Type = int

const QSGGeometry__ByteType QSGGeometry__Type = 5120
const QSGGeometry__UnsignedByteType QSGGeometry__Type = 5121
const QSGGeometry__ShortType QSGGeometry__Type = 5122
const QSGGeometry__UnsignedShortType QSGGeometry__Type = 5123
const QSGGeometry__IntType QSGGeometry__Type = 5124
const QSGGeometry__UnsignedIntType QSGGeometry__Type = 5125
const QSGGeometry__FloatType QSGGeometry__Type = 5126

//  body block end
