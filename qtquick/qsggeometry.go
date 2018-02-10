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
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"
import "qt.go/qtgui"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSGGeometryFromPointer(cthis unsafe.Pointer) *QSGGeometry {
	return &QSGGeometry{&qtrt.CObject{cthis}}
}
func (*QSGGeometry) NewFromPointer(cthis unsafe.Pointer) *QSGGeometry {
	return NewQSGGeometryFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsggeometry.h:138
// index:0
// Public static Visibility=Default Availability=Available
// [16] const QSGGeometry::AttributeSet & defaultAttributes_Point2D()
func (this *QSGGeometry) DefaultAttributes_Point2D() unsafe.Pointer /*555*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry25defaultAttributes_Point2DEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QSGGeometry_DefaultAttributes_Point2D() unsafe.Pointer /*555*/ {
	var nilthis *QSGGeometry
	rv := nilthis.DefaultAttributes_Point2D()
	return rv
}

// /usr/include/qt/QtQuick/qsggeometry.h:139
// index:0
// Public static Visibility=Default Availability=Available
// [16] const QSGGeometry::AttributeSet & defaultAttributes_TexturedPoint2D()
func (this *QSGGeometry) DefaultAttributes_TexturedPoint2D() unsafe.Pointer /*555*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry33defaultAttributes_TexturedPoint2DEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QSGGeometry_DefaultAttributes_TexturedPoint2D() unsafe.Pointer /*555*/ {
	var nilthis *QSGGeometry
	rv := nilthis.DefaultAttributes_TexturedPoint2D()
	return rv
}

// /usr/include/qt/QtQuick/qsggeometry.h:140
// index:0
// Public static Visibility=Default Availability=Available
// [16] const QSGGeometry::AttributeSet & defaultAttributes_ColoredPoint2D()
func (this *QSGGeometry) DefaultAttributes_ColoredPoint2D() unsafe.Pointer /*555*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry32defaultAttributes_ColoredPoint2DEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QSGGeometry_DefaultAttributes_ColoredPoint2D() unsafe.Pointer /*555*/ {
	var nilthis *QSGGeometry
	rv := nilthis.DefaultAttributes_ColoredPoint2D()
	return rv
}

// /usr/include/qt/QtQuick/qsggeometry.h:146
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGGeometry()
func DeleteQSGGeometry(this *QSGGeometry) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 128)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsggeometry.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDrawingMode(unsigned int)
func (this *QSGGeometry) SetDrawingMode(mode uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry14setDrawingModeEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:150
// index:0
// Public inline Visibility=Default Availability=Available
// [4] unsigned int drawingMode()
func (this *QSGGeometry) DrawingMode() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry11drawingModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("uint", rv).(uint) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void allocate(int, int)
func (this *QSGGeometry) Allocate(vertexCount int, indexCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry8allocateEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), vertexCount, indexCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int vertexCount()
func (this *QSGGeometry) VertexCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry11vertexCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:156
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void * vertexData()
func (this *QSGGeometry) VertexData() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry10vertexDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:161
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const void * vertexData()
func (this *QSGGeometry) VertexData_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry10vertexDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:157
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGGeometry::Point2D * vertexDataAsPoint2D()
func (this *QSGGeometry) VertexDataAsPoint2D() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry19vertexDataAsPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:162
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const QSGGeometry::Point2D * vertexDataAsPoint2D()
func (this *QSGGeometry) VertexDataAsPoint2D_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry19vertexDataAsPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGGeometry::TexturedPoint2D * vertexDataAsTexturedPoint2D()
func (this *QSGGeometry) VertexDataAsTexturedPoint2D() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry27vertexDataAsTexturedPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:163
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const QSGGeometry::TexturedPoint2D * vertexDataAsTexturedPoint2D()
func (this *QSGGeometry) VertexDataAsTexturedPoint2D_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry27vertexDataAsTexturedPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:159
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGGeometry::ColoredPoint2D * vertexDataAsColoredPoint2D()
func (this *QSGGeometry) VertexDataAsColoredPoint2D() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry26vertexDataAsColoredPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:164
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const QSGGeometry::ColoredPoint2D * vertexDataAsColoredPoint2D()
func (this *QSGGeometry) VertexDataAsColoredPoint2D_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry26vertexDataAsColoredPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:166
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indexType()
func (this *QSGGeometry) IndexType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry9indexTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:168
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indexCount()
func (this *QSGGeometry) IndexCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry10indexCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:170
// index:0
// Public Visibility=Default Availability=Available
// [8] void * indexData()
func (this *QSGGeometry) IndexData() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry9indexDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:176
// index:1
// Public Visibility=Default Availability=Available
// [8] const void * indexData()
func (this *QSGGeometry) IndexData_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry9indexDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:171
// index:0
// Public inline Visibility=Default Availability=Available
// [8] uint * indexDataAsUInt()
func (this *QSGGeometry) IndexDataAsUInt() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry15indexDataAsUIntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:177
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const uint * indexDataAsUInt()
func (this *QSGGeometry) IndexDataAsUInt_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry15indexDataAsUIntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:172
// index:0
// Public inline Visibility=Default Availability=Available
// [8] quint16 * indexDataAsUShort()
func (this *QSGGeometry) IndexDataAsUShort() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry17indexDataAsUShortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:178
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const quint16 * indexDataAsUShort()
func (this *QSGGeometry) IndexDataAsUShort_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry17indexDataAsUShortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:174
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int sizeOfIndex()
func (this *QSGGeometry) SizeOfIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry11sizeOfIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:180
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int attributeCount()
func (this *QSGGeometry) AttributeCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry14attributeCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:181
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QSGGeometry::Attribute * attributes()
func (this *QSGGeometry) Attributes() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry10attributesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:182
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int sizeOfVertex()
func (this *QSGGeometry) SizeOfVertex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry12sizeOfVertexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:184
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void updateRectGeometry(QSGGeometry *, const QRectF &)
func (this *QSGGeometry) UpdateRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF) {
	var convArg0 = g.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry18updateRectGeometryEPS_RK6QRectF", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QSGGeometry_UpdateRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF) {
	var nilthis *QSGGeometry
	nilthis.UpdateRectGeometry(g, rect)
}

// /usr/include/qt/QtQuick/qsggeometry.h:185
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void updateTexturedRectGeometry(QSGGeometry *, const QRectF &, const QRectF &)
func (this *QSGGeometry) UpdateTexturedRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF, sourceRect *qtcore.QRectF) {
	var convArg0 = g.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg2 = sourceRect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry26updateTexturedRectGeometryEPS_RK6QRectFS3_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}
func QSGGeometry_UpdateTexturedRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF, sourceRect *qtcore.QRectF) {
	var nilthis *QSGGeometry
	nilthis.UpdateTexturedRectGeometry(g, rect, sourceRect)
}

// /usr/include/qt/QtQuick/qsggeometry.h:186
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void updateColoredRectGeometry(QSGGeometry *, const QRectF &)
func (this *QSGGeometry) UpdateColoredRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF) {
	var convArg0 = g.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry25updateColoredRectGeometryEPS_RK6QRectF", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QSGGeometry_UpdateColoredRectGeometry(g *QSGGeometry /*777 QSGGeometry **/, rect *qtcore.QRectF) {
	var nilthis *QSGGeometry
	nilthis.UpdateColoredRectGeometry(g, rect)
}

// /usr/include/qt/QtQuick/qsggeometry.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndexDataPattern(enum QSGGeometry::DataPattern)
func (this *QSGGeometry) SetIndexDataPattern(p int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry19setIndexDataPatternENS_11DataPatternE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:189
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGGeometry::DataPattern indexDataPattern()
func (this *QSGGeometry) IndexDataPattern() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry16indexDataPatternEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVertexDataPattern(enum QSGGeometry::DataPattern)
func (this *QSGGeometry) SetVertexDataPattern(p int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry20setVertexDataPatternENS_11DataPatternE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:192
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGGeometry::DataPattern vertexDataPattern()
func (this *QSGGeometry) VertexDataPattern() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry17vertexDataPatternEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void markIndexDataDirty()
func (this *QSGGeometry) MarkIndexDataDirty() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry18markIndexDataDirtyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void markVertexDataDirty()
func (this *QSGGeometry) MarkVertexDataDirty() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry19markVertexDataDirtyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:197
// index:0
// Public Visibility=Default Availability=Available
// [4] float lineWidth()
func (this *QSGGeometry) LineWidth() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry9lineWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineWidth(float)
func (this *QSGGeometry) SetLineWidth(w float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry12setLineWidthEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
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
