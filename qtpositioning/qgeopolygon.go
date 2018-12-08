package qtpositioning

// /usr/include/qt/QtPositioning/qgeopolygon.h
// #include <qgeopolygon.h>
// #include <QtPositioning>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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
type QGeoPolygon struct {
	*QGeoShape
}
type QGeoPolygon_ITF interface {
	QGeoShape_ITF
	QGeoPolygon_PTR() *QGeoPolygon
}

func (ptr *QGeoPolygon) QGeoPolygon_PTR() *QGeoPolygon { return ptr }

func (this *QGeoPolygon) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGeoShape.GetCthis()
	}
}
func (this *QGeoPolygon) SetCthis(cthis unsafe.Pointer) {
	this.QGeoShape = NewQGeoShapeFromPointer(cthis)
}
func NewQGeoPolygonFromPointer(cthis unsafe.Pointer) *QGeoPolygon {
	bcthis0 := NewQGeoShapeFromPointer(cthis)
	return &QGeoPolygon{bcthis0}
}
func (*QGeoPolygon) NewFromPointer(cthis unsafe.Pointer) *QGeoPolygon {
	return NewQGeoPolygonFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoPolygon()

/*
Constructs a new, empty geo path.
*/
func (*QGeoPolygon) NewForInherit() *QGeoPolygon {
	return NewQGeoPolygon()
}
func NewQGeoPolygon() *QGeoPolygon {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoPolygonC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoPolygonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoPolygon)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGeoPolygon(const QGeoShape &)

/*
Constructs a new, empty geo path.
*/
func (*QGeoPolygon) NewForInherit1(other QGeoShape_ITF) *QGeoPolygon {
	return NewQGeoPolygon1(other)
}
func NewQGeoPolygon1(other QGeoShape_ITF) *QGeoPolygon {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoShape_PTR() != nil {
		convArg0 = other.QGeoShape_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoPolygonC2ERK9QGeoShape", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoPolygonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoPolygon)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGeoPolygon()

/*

 */
func DeleteQGeoPolygon(this *QGeoPolygon) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoPolygonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoPolygon & operator=(const QGeoPolygon &)

/*

 */
func (this *QGeoPolygon) Operator_equal(other QGeoPolygon_ITF) *QGeoPolygon {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoPolygon_PTR() != nil {
		convArg0 = other.QGeoPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoPolygonaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoPolygon)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGeoPolygon &) const

/*

 */
func (this *QGeoPolygon) Operator_equal_equal(other QGeoPolygon_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoPolygon_PTR() != nil {
		convArg0 = other.QGeoPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoPolygoneqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QGeoPolygon &) const

/*

 */
func (this *QGeoPolygon) Operator_not_equal(other QGeoPolygon_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoPolygon_PTR() != nil {
		convArg0 = other.QGeoPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoPolygonneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(double, double)

/*
Translates this geo path by degreesLatitude northwards and degreesLongitude eastwards.

Negative values of degreesLatitude and degreesLongitude correspond to southward and westward translation respectively.
*/
func (this *QGeoPolygon) Translate(degreesLatitude float64, degreesLongitude float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoPolygon9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), degreesLatitude, degreesLongitude)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoPolygon translated(double, double) const

/*
Returns a copy of this geo polygon translated by degreesLatitude northwards and degreesLongitude eastwards.

Negative values of degreesLatitude and degreesLongitude correspond to southward and westward translation respectively.

See also translate().
*/
func (this *QGeoPolygon) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoPolygon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoPolygon10translatedEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), degreesLatitude, degreesLongitude)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoPolygon)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] double length(int, int) const

/*
Returns the length of the polygon's perimeter, in meters, from the element indexFrom to the element indexTo. The length is intended to be the sum of the shortest distances for each pair of adjacent points.
*/
func (this *QGeoPolygon) Length(indexFrom int, indexTo int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoPolygon6lengthEii", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), indexFrom, indexTo)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] double length(int, int) const

/*
Returns the length of the polygon's perimeter, in meters, from the element indexFrom to the element indexTo. The length is intended to be the sum of the shortest distances for each pair of adjacent points.
*/
func (this *QGeoPolygon) Lengthp() float64 {
	// arg: 0, int=Int, =Invalid, , Invalid
	indexFrom := int(0)
	// arg: 1, int=Int, =Invalid, , Invalid
	indexTo := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoPolygon6lengthEii", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), indexFrom, indexTo)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] double length(int, int) const

/*
Returns the length of the polygon's perimeter, in meters, from the element indexFrom to the element indexTo. The length is intended to be the sum of the shortest distances for each pair of adjacent points.
*/
func (this *QGeoPolygon) Lengthp1(indexFrom int) float64 {
	// arg: 1, int=Int, =Invalid, , Invalid
	indexTo := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoPolygon6lengthEii", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), indexFrom, indexTo)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int size() const

/*
Returns the number of elements in the polygon.

This function was introduced in  Qt 5.10.
*/
func (this *QGeoPolygon) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoPolygon4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addCoordinate(const QGeoCoordinate &)

/*
Appends coordinate to the polygon.
*/
func (this *QGeoPolygon) AddCoordinate(coordinate QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoPolygon13addCoordinateERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertCoordinate(int, const QGeoCoordinate &)

/*
Inserts coordinate at the specified index.
*/
func (this *QGeoPolygon) InsertCoordinate(index int, coordinate QGeoCoordinate_ITF) {
	var convArg1 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg1 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoPolygon16insertCoordinateEiRK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void replaceCoordinate(int, const QGeoCoordinate &)

/*
Replaces the path element at the specified index with coordinate.
*/
func (this *QGeoPolygon) ReplaceCoordinate(index int, coordinate QGeoCoordinate_ITF) {
	var convArg1 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg1 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoPolygon17replaceCoordinateEiRK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate coordinateAt(int) const

/*
Returns the coordinate at index .
*/
func (this *QGeoPolygon) CoordinateAt(index int) *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoPolygon12coordinateAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool containsCoordinate(const QGeoCoordinate &) const

/*
Returns true if the polygon's perimeter contains coordinate as one of the elements.
*/
func (this *QGeoPolygon) ContainsCoordinate(coordinate QGeoCoordinate_ITF) bool {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoPolygon18containsCoordinateERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeCoordinate(const QGeoCoordinate &)

/*
Removes the last occurrence of coordinate from the polygon.
*/
func (this *QGeoPolygon) RemoveCoordinate(coordinate QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoPolygon16removeCoordinateERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:84
// index:1
// Public Visibility=Default Availability=Available
// [-2] void removeCoordinate(int)

/*
Removes the last occurrence of coordinate from the polygon.
*/
func (this *QGeoPolygon) RemoveCoordinate1(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QGeoPolygon16removeCoordinateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopolygon.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns the geo path properties as a string.
*/
func (this *QGeoPolygon) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QGeoPolygon8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
