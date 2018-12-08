package qtpositioning

// /usr/include/qt/QtPositioning/qgeorectangle.h
// #include <qgeorectangle.h>
// #include <QtPositioning>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QGeoRectangle struct {
	*QGeoShape
}
type QGeoRectangle_ITF interface {
	QGeoShape_ITF
	QGeoRectangle_PTR() *QGeoRectangle
}

func (ptr *QGeoRectangle) QGeoRectangle_PTR() *QGeoRectangle { return ptr }

func (this *QGeoRectangle) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGeoShape.GetCthis()
	}
}
func (this *QGeoRectangle) SetCthis(cthis unsafe.Pointer) {
	this.QGeoShape = NewQGeoShapeFromPointer(cthis)
}
func NewQGeoRectangleFromPointer(cthis unsafe.Pointer) *QGeoRectangle {
	bcthis0 := NewQGeoShapeFromPointer(cthis)
	return &QGeoRectangle{bcthis0}
}
func (*QGeoRectangle) NewFromPointer(cthis unsafe.Pointer) *QGeoRectangle {
	return NewQGeoRectangleFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoRectangle()

/*
Constructs a new, invalid geo rectangle.
*/
func (*QGeoRectangle) NewForInherit() *QGeoRectangle {
	return NewQGeoRectangle()
}
func NewQGeoRectangle() *QGeoRectangle {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangleC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoRectangleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoRectangle)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGeoRectangle(const QGeoCoordinate &, double, double)

/*
Constructs a new, invalid geo rectangle.
*/
func (*QGeoRectangle) NewForInherit1(center QGeoCoordinate_ITF, degreesWidth float64, degreesHeight float64) *QGeoRectangle {
	return NewQGeoRectangle1(center, degreesWidth, degreesHeight)
}
func NewQGeoRectangle1(center QGeoCoordinate_ITF, degreesWidth float64, degreesHeight float64) *QGeoRectangle {
	var convArg0 unsafe.Pointer
	if center != nil && center.QGeoCoordinate_PTR() != nil {
		convArg0 = center.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangleC2ERK14QGeoCoordinatedd", qtrt.FFI_TYPE_POINTER, convArg0, degreesWidth, degreesHeight)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoRectangleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoRectangle)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:63
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QGeoRectangle(const QGeoCoordinate &, const QGeoCoordinate &)

/*
Constructs a new, invalid geo rectangle.
*/
func (*QGeoRectangle) NewForInherit2(topLeft QGeoCoordinate_ITF, bottomRight QGeoCoordinate_ITF) *QGeoRectangle {
	return NewQGeoRectangle2(topLeft, bottomRight)
}
func NewQGeoRectangle2(topLeft QGeoCoordinate_ITF, bottomRight QGeoCoordinate_ITF) *QGeoRectangle {
	var convArg0 unsafe.Pointer
	if topLeft != nil && topLeft.QGeoCoordinate_PTR() != nil {
		convArg0 = topLeft.QGeoCoordinate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if bottomRight != nil && bottomRight.QGeoCoordinate_PTR() != nil {
		convArg1 = bottomRight.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangleC2ERK14QGeoCoordinateS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoRectangleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoRectangle)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:66
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QGeoRectangle(const QGeoShape &)

/*
Constructs a new, invalid geo rectangle.
*/
func (*QGeoRectangle) NewForInherit3(other QGeoShape_ITF) *QGeoRectangle {
	return NewQGeoRectangle3(other)
}
func NewQGeoRectangle3(other QGeoShape_ITF) *QGeoRectangle {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoShape_PTR() != nil {
		convArg0 = other.QGeoShape_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangleC2ERK9QGeoShape", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoRectangleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoRectangle)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGeoRectangle()

/*

 */
func DeleteQGeoRectangle(this *QGeoRectangle) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoRectangle & operator=(const QGeoRectangle &)

/*

 */
func (this *QGeoRectangle) Operator_equal(other QGeoRectangle_ITF) *QGeoRectangle {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoRectangle_PTR() != nil {
		convArg0 = other.QGeoRectangle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangleaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoRectangleFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoRectangle)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGeoRectangle &) const

/*

 */
func (this *QGeoRectangle) Operator_equal_equal(other QGeoRectangle_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoRectangle_PTR() != nil {
		convArg0 = other.QGeoRectangle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangleeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QGeoRectangle &) const

/*

 */
func (this *QGeoRectangle) Operator_not_equal(other QGeoRectangle_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoRectangle_PTR() != nil {
		convArg0 = other.QGeoRectangle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangleneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTopLeft(const QGeoCoordinate &)

/*
Sets the top left coordinate of this geo rectangle to topLeft.

Note: Setter function for property topLeft.

See also topLeft().
*/
func (this *QGeoRectangle) SetTopLeft(topLeft QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if topLeft != nil && topLeft.QGeoCoordinate_PTR() != nil {
		convArg0 = topLeft.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangle10setTopLeftERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate topLeft() const

/*
Returns the top left coordinate of this geo rectangle.

Note: Getter function for property topLeft.

See also setTopLeft().
*/
func (this *QGeoRectangle) TopLeft() *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle7topLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTopRight(const QGeoCoordinate &)

/*
Sets the top right coordinate of this geo rectangle to topRight.

Note: Setter function for property topRight.

See also topRight().
*/
func (this *QGeoRectangle) SetTopRight(topRight QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if topRight != nil && topRight.QGeoCoordinate_PTR() != nil {
		convArg0 = topRight.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangle11setTopRightERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate topRight() const

/*
Returns the top right coordinate of this geo rectangle.

Note: Getter function for property topRight.

See also setTopRight().
*/
func (this *QGeoRectangle) TopRight() *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle8topRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBottomLeft(const QGeoCoordinate &)

/*
Sets the bottom left coordinate of this geo rectangle to bottomLeft.

Note: Setter function for property bottomLeft.

See also bottomLeft().
*/
func (this *QGeoRectangle) SetBottomLeft(bottomLeft QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if bottomLeft != nil && bottomLeft.QGeoCoordinate_PTR() != nil {
		convArg0 = bottomLeft.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangle13setBottomLeftERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate bottomLeft() const

/*
Returns the bottom left coordinate of this geo rectangle.

Note: Getter function for property bottomLeft.

See also setBottomLeft().
*/
func (this *QGeoRectangle) BottomLeft() *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle10bottomLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBottomRight(const QGeoCoordinate &)

/*
Sets the bottom right coordinate of this geo rectangle to bottomRight.

Note: Setter function for property bottomRight.

See also bottomRight().
*/
func (this *QGeoRectangle) SetBottomRight(bottomRight QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if bottomRight != nil && bottomRight.QGeoCoordinate_PTR() != nil {
		convArg0 = bottomRight.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangle14setBottomRightERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate bottomRight() const

/*
Returns the bottom right coordinate of this geo rectangle.

Note: Getter function for property bottomRight.

See also setBottomRight().
*/
func (this *QGeoRectangle) BottomRight() *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle11bottomRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCenter(const QGeoCoordinate &)

/*
Sets the center of this geo rectangle to center.

If this causes the geo rectangle to cross on of the poles the height of the geo rectangle will be truncated such that the geo rectangle only extends up to the pole. The center of the geo rectangle will be unchanged, and the height will be adjusted such that the center point is at the center of the truncated geo rectangle.

Note: Setter function for property center.

See also center().
*/
func (this *QGeoRectangle) SetCenter(center QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if center != nil && center.QGeoCoordinate_PTR() != nil {
		convArg0 = center.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangle9setCenterERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate center() const

/*
Returns the center of this geo rectangle. Equivalent to QGeoShape::center().

Note: Getter function for property center.

See also setCenter().
*/
func (this *QGeoRectangle) Center() *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle6centerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidth(double)

/*
Sets the width of this geo rectangle in degrees to degreesWidth.

Note: Setter function for property width.

See also width().
*/
func (this *QGeoRectangle) SetWidth(degreesWidth float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangle8setWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), degreesWidth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] double width() const

/*
Returns the width of this geo rectangle in degrees.

The return value is undefined if this geo rectangle is invalid.

Note: Getter function for property width.

See also setWidth().
*/
func (this *QGeoRectangle) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeight(double)

/*
Sets the height of this geo rectangle in degrees to degreesHeight.

Note: Setter function for property height.

See also height().
*/
func (this *QGeoRectangle) SetHeight(degreesHeight float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangle9setHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), degreesHeight)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] double height() const

/*
Returns the height of this geo rectangle in degrees.

The return value is undefined if this geo rectangle is invalid.

Note: Getter function for property height.

See also setHeight().
*/
func (this *QGeoRectangle) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QGeoRectangle &) const

/*
Returns whether the geo rectangle rectangle is contained within this geo rectangle.
*/
func (this *QGeoRectangle) Contains(rectangle QGeoRectangle_ITF) bool {
	var convArg0 unsafe.Pointer
	if rectangle != nil && rectangle.QGeoRectangle_PTR() != nil {
		convArg0 = rectangle.QGeoRectangle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle8containsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:101
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QGeoRectangle &) const

/*
Returns whether the geo rectangle rectangle intersects this geo rectangle.

If the top or bottom edges of both geo rectangles are at one of the poles the geo rectangles are considered to be intersecting, since the longitude is irrelevant when the edges are at the pole.
*/
func (this *QGeoRectangle) Intersects(rectangle QGeoRectangle_ITF) bool {
	var convArg0 unsafe.Pointer
	if rectangle != nil && rectangle.QGeoRectangle_PTR() != nil {
		convArg0 = rectangle.QGeoRectangle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle10intersectsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(double, double)

/*
Translates this geo rectangle by degreesLatitude northwards and degreesLongitude eastwards.

Negative values of degreesLatitude and degreesLongitude correspond to southward and westward translation respectively.

If the translation would have caused the geo rectangle to cross a pole the geo rectangle will be translated until the top or bottom edge of the geo rectangle touches the pole but not further.
*/
func (this *QGeoRectangle) Translate(degreesLatitude float64, degreesLongitude float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangle9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), degreesLatitude, degreesLongitude)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoRectangle translated(double, double) const

/*
Returns a copy of this geo rectangle translated by degreesLatitude northwards and degreesLongitude eastwards.

Negative values of degreesLatitude and degreesLongitude correspond to southward and westward translation respectively.

See also translate().
*/
func (this *QGeoRectangle) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoRectangle /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle10translatedEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), degreesLatitude, degreesLongitude)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoRectangleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoRectangle)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void extendRectangle(const QGeoCoordinate &)

/*
Extends the geo rectangle to also cover the coordinate coordinate

This function was introduced in  Qt 5.9.
*/
func (this *QGeoRectangle) ExtendRectangle(coordinate QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangle15extendRectangleERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoRectangle united(const QGeoRectangle &) const

/*
Returns the smallest geo rectangle which contains both this geo rectangle and rectangle.

If the centers of the two geo rectangles are separated by exactly 180.0 degrees then the width is set to 360.0 degrees with the leftmost longitude set to -180.0 degrees and the rightmost longitude set to 180.0 degrees. This is done to ensure that the result is independent of the order of the operands.
*/
func (this *QGeoRectangle) United(rectangle QGeoRectangle_ITF) *QGeoRectangle /*123*/ {
	var convArg0 unsafe.Pointer
	if rectangle != nil && rectangle.QGeoRectangle_PTR() != nil {
		convArg0 = rectangle.QGeoRectangle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle6unitedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoRectangleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoRectangle)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoRectangle operator|(const QGeoRectangle &) const

/*

 */
func (this *QGeoRectangle) Operator_or(rectangle QGeoRectangle_ITF) *QGeoRectangle /*123*/ {
	var convArg0 unsafe.Pointer
	if rectangle != nil && rectangle.QGeoRectangle_PTR() != nil {
		convArg0 = rectangle.QGeoRectangle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangleorERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoRectangleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoRectangle)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoRectangle & operator|=(const QGeoRectangle &)

/*

 */
func (this *QGeoRectangle) Operator_or_equal(rectangle QGeoRectangle_ITF) *QGeoRectangle {
	var convArg0 unsafe.Pointer
	if rectangle != nil && rectangle.QGeoRectangle_PTR() != nil {
		convArg0 = rectangle.QGeoRectangle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGeoRectangleoRERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoRectangleFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoRectangle)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeorectangle.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns the geo rectangle properties as a string.

This function was introduced in  Qt 5.5.
*/
func (this *QGeoRectangle) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGeoRectangle8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
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
}

//  keep block end
