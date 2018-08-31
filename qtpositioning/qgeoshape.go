package qtpositioning

// /usr/include/qt/QtPositioning/qgeoshape.h
// #include <qgeoshape.h>
// #include <QtPositioning>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QGeoShape struct {
	*qtrt.CObject
}
type QGeoShape_ITF interface {
	QGeoShape_PTR() *QGeoShape
}

func (ptr *QGeoShape) QGeoShape_PTR() *QGeoShape { return ptr }

func (this *QGeoShape) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGeoShape) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGeoShapeFromPointer(cthis unsafe.Pointer) *QGeoShape {
	return &QGeoShape{&qtrt.CObject{cthis}}
}
func (*QGeoShape) NewFromPointer(cthis unsafe.Pointer) *QGeoShape {
	return NewQGeoShapeFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeoshape.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoShape()

/*
Constructs a new invalid geo shape of UnknownType.
*/
func (*QGeoShape) NewForInherit() *QGeoShape {
	return NewQGeoShape()
}
func NewQGeoShape() *QGeoShape {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGeoShapeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoShapeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoShape)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeoshape.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGeoShape()

/*

 */
func DeleteQGeoShape(this *QGeoShape) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGeoShapeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeoshape.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] QGeoShape::ShapeType type() const

/*
Returns the type of this geo shape.

Note: Getter function for property type.
*/
func (this *QGeoShape) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGeoShape4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qgeoshape.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns whether this geo shape is valid.

Note: Getter function for property isValid.
*/
func (this *QGeoShape) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGeoShape7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoshape.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns whether this geo shape is empty.

An empty geo shape is a region which has a geometrical area of 0.

Note: Getter function for property isEmpty.
*/
func (this *QGeoShape) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGeoShape7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoshape.h:77
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QGeoCoordinate &) const

/*
Returns whether the coordinate coordinate is contained within this geo shape.
*/
func (this *QGeoShape) Contains(coordinate QGeoCoordinate_ITF) bool {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGeoShape8containsERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoshape.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoRectangle boundingGeoRectangle() const

/*
Returns a QGeoRectangle representing the geographical bounding rectangle of the geo shape, that defines the latitudinal/longitudinal bounds of the geo shape.

This function was introduced in  Qt 5.9.
*/
func (this *QGeoShape) BoundingGeoRectangle() *QGeoRectangle /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGeoShape20boundingGeoRectangleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoRectangleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoRectangle)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeoshape.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate center() const

/*
Returns the coordinate located at the geometric center of the geo shape.

This function was introduced in  Qt 5.5.
*/
func (this *QGeoShape) Center() *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGeoShape6centerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeoshape.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void extendShape(const QGeoCoordinate &)

/*

 */
func (this *QGeoShape) ExtendShape(coordinate QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGeoShape11extendShapeERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoshape.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGeoShape &) const

/*

 */
func (this *QGeoShape) Operator_equal_equal(other QGeoShape_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoShape_PTR() != nil {
		convArg0 = other.QGeoShape_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGeoShapeeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoshape.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QGeoShape &) const

/*

 */
func (this *QGeoShape) Operator_not_equal(other QGeoShape_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoShape_PTR() != nil {
		convArg0 = other.QGeoShape_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGeoShapeneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoshape.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoShape & operator=(const QGeoShape &)

/*

 */
func (this *QGeoShape) Operator_equal(other QGeoShape_ITF) *QGeoShape {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoShape_PTR() != nil {
		convArg0 = other.QGeoShape_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGeoShapeaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoShapeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoShape)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeoshape.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns a string representation of this geo shape.

This function was introduced in  Qt 5.5.
*/
func (this *QGeoShape) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGeoShape8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

/*
Describes the type of the shape.


*/
type QGeoShape__ShapeType = int

// A shape of unknown type.
const QGeoShape__UnknownType QGeoShape__ShapeType = 0

// A rectangular shape.
const QGeoShape__RectangleType QGeoShape__ShapeType = 1

// A circular shape.
const QGeoShape__CircleType QGeoShape__ShapeType = 2

// A path type.
const QGeoShape__PathType QGeoShape__ShapeType = 3

//
const QGeoShape__PolygonType QGeoShape__ShapeType = 4

func (this *QGeoShape) ShapeTypeItemName(val int) string {
	switch val {
	case QGeoShape__UnknownType: // 0
		return "UnknownType"
	case QGeoShape__RectangleType: // 1
		return "RectangleType"
	case QGeoShape__CircleType: // 2
		return "CircleType"
	case QGeoShape__PathType: // 3
		return "PathType"
	case QGeoShape__PolygonType: // 4
		return "PolygonType"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGeoShape_ShapeTypeItemName(val int) string {
	var nilthis *QGeoShape
	return nilthis.ShapeTypeItemName(val)
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
