package qtquick

// /usr/include/qt/QtQuick/qsggeometry.h
// #include <qsggeometry.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

/*

 */
type QSGGeometry struct {
	*qtrt.CObject
}
type QSGGeometry_ITF interface {
	QSGGeometry_PTR() *QSGGeometry
}

func (ptr *QSGGeometry) QSGGeometry_PTR() *QSGGeometry { return ptr }

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

/*
Convenience function which returns attributes to be used for 2D solid color drawing.
*/
func (this *QSGGeometry) DefaultAttributes_Point2D() unsafe.Pointer /*555*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry25defaultAttributes_Point2DEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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

/*
Convenience function which returns attributes to be used for textured 2D drawing.
*/
func (this *QSGGeometry) DefaultAttributes_TexturedPoint2D() unsafe.Pointer /*555*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry33defaultAttributes_TexturedPoint2DEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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

/*
Convenience function which returns attributes to be used for per vertex colored 2D drawing.
*/
func (this *QSGGeometry) DefaultAttributes_ColoredPoint2D() unsafe.Pointer /*555*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry32defaultAttributes_ColoredPoint2DEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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

/*

 */
func DeleteQSGGeometry(this *QSGGeometry) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 128)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsggeometry.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDrawingMode(unsigned int)

/*
Sets the mode to be used for drawing this geometry.

The default value is QSGGeometry::DrawTriangleStrip.

See also drawingMode() and DrawingMode.
*/
func (this *QSGGeometry) SetDrawingMode(mode uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry14setDrawingModeEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:150
// index:0
// Public inline Visibility=Default Availability=Available
// [4] unsigned int drawingMode() const

/*
Returns the drawing mode of this geometry.

The default value is GL_TRIANGLE_STRIP.

See also setDrawingMode().
*/
func (this *QSGGeometry) DrawingMode() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry11drawingModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("uint", rv).(uint) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void allocate(int, int)

/*
Resizes the vertex and index data of this geometry object to fit vertexCount vertices and indexCount indices.

Vertex and index data will be invalidated after this call and the caller must mark the associated geometry node as dirty, by calling node->markDirty(QSGNode::DirtyGeometry) to ensure that the renderer has a chance to update internal buffers.
*/
func (this *QSGGeometry) Allocate(vertexCount int, indexCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry8allocateEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), vertexCount, indexCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void allocate(int, int)

/*
Resizes the vertex and index data of this geometry object to fit vertexCount vertices and indexCount indices.

Vertex and index data will be invalidated after this call and the caller must mark the associated geometry node as dirty, by calling node->markDirty(QSGNode::DirtyGeometry) to ensure that the renderer has a chance to update internal buffers.
*/
func (this *QSGGeometry) Allocate__(vertexCount int) {
	// arg: 1, int=Int, =Invalid, , Invalid
	indexCount := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry8allocateEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), vertexCount, indexCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int vertexCount() const

/*
Returns the number of vertices in this geometry object.
*/
func (this *QSGGeometry) VertexCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry11vertexCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:156
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void * vertexData()

/*
Returns a pointer to the raw vertex data of this geometry object.

See also vertexDataAsPoint2D() and vertexDataAsTexturedPoint2D().
*/
func (this *QSGGeometry) VertexData() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry10vertexDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:161
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const void * vertexData() const

/*
Returns a pointer to the raw vertex data of this geometry object.

See also vertexDataAsPoint2D() and vertexDataAsTexturedPoint2D().
*/
func (this *QSGGeometry) VertexData_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry10vertexDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:157
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGGeometry::Point2D * vertexDataAsPoint2D()

/*
Convenience function to access the vertex data as a mutable array of QSGGeometry::Point2D.
*/
func (this *QSGGeometry) VertexDataAsPoint2D() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry19vertexDataAsPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:162
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const QSGGeometry::Point2D * vertexDataAsPoint2D() const

/*
Convenience function to access the vertex data as a mutable array of QSGGeometry::Point2D.
*/
func (this *QSGGeometry) VertexDataAsPoint2D_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry19vertexDataAsPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGGeometry::TexturedPoint2D * vertexDataAsTexturedPoint2D()

/*
Convenience function to access the vertex data as a mutable array of QSGGeometry::TexturedPoint2D.
*/
func (this *QSGGeometry) VertexDataAsTexturedPoint2D() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry27vertexDataAsTexturedPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:163
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const QSGGeometry::TexturedPoint2D * vertexDataAsTexturedPoint2D() const

/*
Convenience function to access the vertex data as a mutable array of QSGGeometry::TexturedPoint2D.
*/
func (this *QSGGeometry) VertexDataAsTexturedPoint2D_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry27vertexDataAsTexturedPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:159
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGGeometry::ColoredPoint2D * vertexDataAsColoredPoint2D()

/*
Convenience function to access the vertex data as a mutable array of QSGGeometry::ColoredPoint2D.
*/
func (this *QSGGeometry) VertexDataAsColoredPoint2D() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry26vertexDataAsColoredPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:164
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const QSGGeometry::ColoredPoint2D * vertexDataAsColoredPoint2D() const

/*
Convenience function to access the vertex data as a mutable array of QSGGeometry::ColoredPoint2D.
*/
func (this *QSGGeometry) VertexDataAsColoredPoint2D_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry26vertexDataAsColoredPoint2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:166
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indexType() const

/*
Returns the primitive type used for indices in this geometry object.
*/
func (this *QSGGeometry) IndexType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry9indexTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:168
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indexCount() const

/*
Returns the number of indices in this geometry object.
*/
func (this *QSGGeometry) IndexCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry10indexCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:170
// index:0
// Public Visibility=Default Availability=Available
// [8] void * indexData()

/*
Returns a pointer to the raw index data of this geometry object.

See also indexDataAsUShort() and indexDataAsUInt().
*/
func (this *QSGGeometry) IndexData() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry9indexDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:176
// index:1
// Public Visibility=Default Availability=Available
// [8] const void * indexData() const

/*
Returns a pointer to the raw index data of this geometry object.

See also indexDataAsUShort() and indexDataAsUInt().
*/
func (this *QSGGeometry) IndexData_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry9indexDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:171
// index:0
// Public inline Visibility=Default Availability=Available
// [8] uint * indexDataAsUInt()

/*
Convenience function to access the index data as a mutable array of 32-bit unsigned integers.
*/
func (this *QSGGeometry) IndexDataAsUInt() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry15indexDataAsUIntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:177
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const uint * indexDataAsUInt() const

/*
Convenience function to access the index data as a mutable array of 32-bit unsigned integers.
*/
func (this *QSGGeometry) IndexDataAsUInt_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry15indexDataAsUIntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:172
// index:0
// Public inline Visibility=Default Availability=Available
// [8] quint16 * indexDataAsUShort()

/*
Convenience function to access the index data as a mutable array of 16-bit unsigned integers.
*/
func (this *QSGGeometry) IndexDataAsUShort() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry17indexDataAsUShortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:178
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const quint16 * indexDataAsUShort() const

/*
Convenience function to access the index data as a mutable array of 16-bit unsigned integers.
*/
func (this *QSGGeometry) IndexDataAsUShort_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry17indexDataAsUShortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:174
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int sizeOfIndex() const

/*
Returns the byte size of the index type.

This value is either 1 when index type is GL_UNSIGNED_BYTE or 2 when index type is GL_UNSIGNED_SHORT. For Desktop OpenGL, GL_UNSIGNED_INT with the value 4 is also supported.
*/
func (this *QSGGeometry) SizeOfIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry11sizeOfIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:180
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int attributeCount() const

/*
Returns the number of attributes in the attrbute set used by this geometry.
*/
func (this *QSGGeometry) AttributeCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry14attributeCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:181
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QSGGeometry::Attribute * attributes() const

/*
Returns an array with the attributes of this geometry. The size of the array is given with attributeCount().
*/
func (this *QSGGeometry) Attributes() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry10attributesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsggeometry.h:182
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int sizeOfVertex() const

/*
Returns the size in bytes of one vertex.

This value comes from the attributes.
*/
func (this *QSGGeometry) SizeOfVertex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry12sizeOfVertexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:184
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void updateRectGeometry(QSGGeometry *, const QRectF &)

/*
Updates the geometry g with the coordinates in rect.

The function assumes the geometry object contains a single triangle strip of QSGGeometry::Point2D vertices
*/
func (this *QSGGeometry) UpdateRectGeometry(g QSGGeometry_ITF /*777 QSGGeometry **/, rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if g != nil && g.QSGGeometry_PTR() != nil {
		convArg0 = g.QSGGeometry_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry18updateRectGeometryEPS_RK6QRectF", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QSGGeometry_UpdateRectGeometry(g QSGGeometry_ITF /*777 QSGGeometry **/, rect qtcore.QRectF_ITF) {
	var nilthis *QSGGeometry
	nilthis.UpdateRectGeometry(g, rect)
}

// /usr/include/qt/QtQuick/qsggeometry.h:185
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void updateTexturedRectGeometry(QSGGeometry *, const QRectF &, const QRectF &)

/*
Updates the geometry g with the coordinates in rect and texture coordinates from textureRect.

textureRect should be in normalized coordinates.

g is assumed to be a triangle strip of four vertices of type QSGGeometry::TexturedPoint2D.
*/
func (this *QSGGeometry) UpdateTexturedRectGeometry(g QSGGeometry_ITF /*777 QSGGeometry **/, rect qtcore.QRectF_ITF, sourceRect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if g != nil && g.QSGGeometry_PTR() != nil {
		convArg0 = g.QSGGeometry_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceRect != nil && sourceRect.QRectF_PTR() != nil {
		convArg2 = sourceRect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry26updateTexturedRectGeometryEPS_RK6QRectFS3_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}
func QSGGeometry_UpdateTexturedRectGeometry(g QSGGeometry_ITF /*777 QSGGeometry **/, rect qtcore.QRectF_ITF, sourceRect qtcore.QRectF_ITF) {
	var nilthis *QSGGeometry
	nilthis.UpdateTexturedRectGeometry(g, rect, sourceRect)
}

// /usr/include/qt/QtQuick/qsggeometry.h:186
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void updateColoredRectGeometry(QSGGeometry *, const QRectF &)

/*
Updates the geometry g with the coordinates in rect.

The function assumes the geometry object contains a single triangle strip of QSGGeometry::ColoredPoint2D vertices
*/
func (this *QSGGeometry) UpdateColoredRectGeometry(g QSGGeometry_ITF /*777 QSGGeometry **/, rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if g != nil && g.QSGGeometry_PTR() != nil {
		convArg0 = g.QSGGeometry_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry25updateColoredRectGeometryEPS_RK6QRectF", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QSGGeometry_UpdateColoredRectGeometry(g QSGGeometry_ITF /*777 QSGGeometry **/, rect qtcore.QRectF_ITF) {
	var nilthis *QSGGeometry
	nilthis.UpdateColoredRectGeometry(g, rect)
}

// /usr/include/qt/QtQuick/qsggeometry.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndexDataPattern(QSGGeometry::DataPattern)

/*
Sets the usage pattern for indices to p.

The default is AlwaysUploadPattern. When set to anything other than the default, the user must call markIndexDataDirty() after changing the index data, in addition to calling QSGNode::markDirty() with QSGNode::DirtyGeometry.

See also indexDataPattern().
*/
func (this *QSGGeometry) SetIndexDataPattern(p int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry19setIndexDataPatternENS_11DataPatternE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), p)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:189
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGGeometry::DataPattern indexDataPattern() const

/*
Returns the usage pattern for indices in this geometry. The default pattern is AlwaysUploadPattern.

See also setIndexDataPattern().
*/
func (this *QSGGeometry) IndexDataPattern() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry16indexDataPatternEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVertexDataPattern(QSGGeometry::DataPattern)

/*
Sets the usage pattern for vertices to p.

The default is AlwaysUploadPattern. When set to anything other than the default, the user must call markVertexDataDirty() after changing the vertex data, in addition to calling QSGNode::markDirty() with QSGNode::DirtyGeometry.

See also vertexDataPattern().
*/
func (this *QSGGeometry) SetVertexDataPattern(p int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry20setVertexDataPatternENS_11DataPatternE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), p)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:192
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGGeometry::DataPattern vertexDataPattern() const

/*
Returns the usage pattern for vertices in this geometry. The default pattern is AlwaysUploadPattern.

See also setVertexDataPattern().
*/
func (this *QSGGeometry) VertexDataPattern() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry17vertexDataPatternEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void markIndexDataDirty()

/*
Mark that the vertices in this geometry has changed and must be uploaded again.

This function only has an effect when the usage pattern for vertices is StaticData and the renderer that renders this geometry uploads the geometry into Vertex Buffer Objects (VBOs).
*/
func (this *QSGGeometry) MarkIndexDataDirty() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry18markIndexDataDirtyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void markVertexDataDirty()

/*
Mark that the vertices in this geometry has changed and must be uploaded again.

This function only has an effect when the usage pattern for vertices is StaticData and the renderer that renders this geometry uploads the geometry into Vertex Buffer Objects (VBOs).
*/
func (this *QSGGeometry) MarkVertexDataDirty() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry19markVertexDataDirtyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsggeometry.h:197
// index:0
// Public Visibility=Default Availability=Available
// [4] float lineWidth() const

/*
Gets the current line or point width or to be used for this geometry. This property only applies to line width when the drawingMode is DrawLines, DarwLineStrip, or DrawLineLoop. For desktop OpenGL, it also applies to point size when the drawingMode is DrawPoints.

The default value is 1.0

Note: When not using OpenGL, support for point and line drawing may be limited. For example, some APIs do not support point sprites and so setting a size other than 1 is not possible. Some backends may be able implement support via geometry shaders, but this is not guaranteed to be always available.

See also setLineWidth() and drawingMode().
*/
func (this *QSGGeometry) LineWidth() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGGeometry9lineWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtQuick/qsggeometry.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineWidth(float)

/*
Sets the line or point width to be used for this geometry to width. This property only applies to line width when the drawingMode is DrawLines, DrawLineStrip, or DrawLineLoop. For Desktop OpenGL, it also applies to point size when the drawingMode is DrawPoints.

Note: How line width and point size are treated is implementation dependent: The application should not rely on these, but rather create triangles or similar to draw areas. On OpenGL ES, line width support is limited and point size is unsupported.

See also lineWidth() and drawingMode().
*/
func (this *QSGGeometry) SetLineWidth(w float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGGeometry12setLineWidthEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QSGGeometry__AttributeType = int

//
const QSGGeometry__UnknownAttribute QSGGeometry__AttributeType = 0

//
const QSGGeometry__PositionAttribute QSGGeometry__AttributeType = 1

//
const QSGGeometry__ColorAttribute QSGGeometry__AttributeType = 2

//
const QSGGeometry__TexCoordAttribute QSGGeometry__AttributeType = 3

//
const QSGGeometry__TexCoord1Attribute QSGGeometry__AttributeType = 4

//
const QSGGeometry__TexCoord2Attribute QSGGeometry__AttributeType = 5

/*
The DataPattern enum is used to specify the use pattern for the vertex and index data in a geometry object.


*/
type QSGGeometry__DataPattern = int

// The data is always uploaded. This means that the user does not need to explicitly mark index and vertex data as dirty after changing it. This is the default.
const QSGGeometry__AlwaysUploadPattern QSGGeometry__DataPattern = 0

// The data is modified for almost every time it is drawn. This is a hint that may provide better performance. When set, the user must make sure to mark the data as dirty after changing it.
const QSGGeometry__StreamPattern QSGGeometry__DataPattern = 1

// The data is modified repeatedly and drawn many times. This is a hint that may provide better performance. When set the user must make sure to mark the data as dirty after changing it.
const QSGGeometry__DynamicPattern QSGGeometry__DataPattern = 2

// The data is modified once and drawn many times. This is a hint that may provide better performance. When set the user must make sure to mark the data as dirty after changing it.
const QSGGeometry__StaticPattern QSGGeometry__DataPattern = 3

/*
The values correspond to OpenGL enum values like GL_POINTS, GL_LINES, etc. QSGGeometry provies its own type in order to be able to provide the same API with non-OpenGL backends as well.

ConstantValue

*/
type QSGGeometry__DrawingMode = int

//
const QSGGeometry__DrawPoints QSGGeometry__DrawingMode = 0

//
const QSGGeometry__DrawLines QSGGeometry__DrawingMode = 1

//
const QSGGeometry__DrawLineLoop QSGGeometry__DrawingMode = 2

//
const QSGGeometry__DrawLineStrip QSGGeometry__DrawingMode = 3

//
const QSGGeometry__DrawTriangles QSGGeometry__DrawingMode = 4

//
const QSGGeometry__DrawTriangleStrip QSGGeometry__DrawingMode = 5

//
const QSGGeometry__DrawTriangleFan QSGGeometry__DrawingMode = 6

/*
The values correspond to OpenGL type constants like GL_BYTE, GL_UNSIGNED_BYTE, etc. QSGGeometry provies its own type in order to be able to provide the same API with non-OpenGL backends as well.

ConstantValue

*/
type QSGGeometry__Type = int

//
const QSGGeometry__ByteType QSGGeometry__Type = 5120

//
const QSGGeometry__UnsignedByteType QSGGeometry__Type = 5121

//
const QSGGeometry__ShortType QSGGeometry__Type = 5122

//
const QSGGeometry__UnsignedShortType QSGGeometry__Type = 5123

//
const QSGGeometry__IntType QSGGeometry__Type = 5124

//
const QSGGeometry__UnsignedIntType QSGGeometry__Type = 5125

//
const QSGGeometry__FloatType QSGGeometry__Type = 5126

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
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
