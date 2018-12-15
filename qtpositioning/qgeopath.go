package qtpositioning

// /usr/include/qt/QtPositioning/qgeopath.h
// #include <qgeopath.h>
// #include <QtPositioning>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QGeoPath struct {
	*QGeoShape
}
type QGeoPath_ITF interface {
	QGeoShape_ITF
	QGeoPath_PTR() *QGeoPath
}

func (ptr *QGeoPath) QGeoPath_PTR() *QGeoPath { return ptr }

func (this *QGeoPath) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGeoShape.GetCthis()
	}
}
func (this *QGeoPath) SetCthis(cthis unsafe.Pointer) {
	this.QGeoShape = NewQGeoShapeFromPointer(cthis)
}
func NewQGeoPathFromPointer(cthis unsafe.Pointer) *QGeoPath {
	bcthis0 := NewQGeoShapeFromPointer(cthis)
	return &QGeoPath{bcthis0}
}
func (*QGeoPath) NewFromPointer(cthis unsafe.Pointer) *QGeoPath {
	return NewQGeoPathFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeopath.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoPath()

/*
Constructs a new, empty geo path.
*/
func (*QGeoPath) NewForInherit() *QGeoPath {
	return NewQGeoPath()
}
func NewQGeoPath() *QGeoPath {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGeoPathC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoPathFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoPath)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeopath.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGeoPath(const QGeoShape &)

/*
Constructs a new, empty geo path.
*/
func (*QGeoPath) NewForInherit1(other QGeoShape_ITF) *QGeoPath {
	return NewQGeoPath1(other)
}
func NewQGeoPath1(other QGeoShape_ITF) *QGeoPath {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoShape_PTR() != nil {
		convArg0 = other.QGeoShape_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGeoPathC2ERK9QGeoShape", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoPathFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoPath)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeopath.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGeoPath()

/*

 */
func DeleteQGeoPath(this *QGeoPath) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGeoPathD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeopath.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoPath & operator=(const QGeoPath &)

/*

 */
func (this *QGeoPath) Operator_equal(other QGeoPath_ITF) *QGeoPath {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoPath_PTR() != nil {
		convArg0 = other.QGeoPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGeoPathaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoPathFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoPath)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeopath.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGeoPath &) const

/*

 */
func (this *QGeoPath) Operator_equal_equal(other QGeoPath_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoPath_PTR() != nil {
		convArg0 = other.QGeoPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGeoPatheqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeopath.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QGeoPath &) const

/*

 */
func (this *QGeoPath) Operator_not_equal(other QGeoPath_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoPath_PTR() != nil {
		convArg0 = other.QGeoPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGeoPathneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeopath.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidth(const qreal &)

/*

 */
func (this *QGeoPath) SetWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGeoPath8setWidthERKd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopath.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal width() const

/*
Returns the width of the path, in meters. This information is used in the contains method The default value is 0.

Note: Getter function for property width.

See also setWidth().
*/
func (this *QGeoPath) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGeoPath5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeopath.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(double, double)

/*
Translates this geo path by degreesLatitude northwards and degreesLongitude eastwards.

Negative values of degreesLatitude and degreesLongitude correspond to southward and westward translation respectively.
*/
func (this *QGeoPath) Translate(degreesLatitude float64, degreesLongitude float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGeoPath9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), degreesLatitude, degreesLongitude)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopath.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoPath translated(double, double) const

/*
Returns a copy of this geo path translated by degreesLatitude northwards and degreesLongitude eastwards.

Negative values of degreesLatitude and degreesLongitude correspond to southward and westward translation respectively.

See also translate().
*/
func (this *QGeoPath) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGeoPath10translatedEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), degreesLatitude, degreesLongitude)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoPath)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeopath.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] double length(int, int) const

/*
Returns the length of the path, in meters, from the element indexFrom to the element indexTo. The length is intended to be the sum of the shortest distances for each pair of adjacent points.
*/
func (this *QGeoPath) Length(indexFrom int, indexTo int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGeoPath6lengthEii", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), indexFrom, indexTo)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeopath.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] double length(int, int) const

/*
Returns the length of the path, in meters, from the element indexFrom to the element indexTo. The length is intended to be the sum of the shortest distances for each pair of adjacent points.
*/
func (this *QGeoPath) Lengthp() float64 {
	// arg: 0, int=Int, =Invalid, , Invalid
	indexFrom := int(0)
	// arg: 1, int=Int, =Invalid, , Invalid
	indexTo := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGeoPath6lengthEii", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), indexFrom, indexTo)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeopath.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] double length(int, int) const

/*
Returns the length of the path, in meters, from the element indexFrom to the element indexTo. The length is intended to be the sum of the shortest distances for each pair of adjacent points.
*/
func (this *QGeoPath) Lengthp1(indexFrom int) float64 {
	// arg: 1, int=Int, =Invalid, , Invalid
	indexTo := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGeoPath6lengthEii", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), indexFrom, indexTo)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeopath.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addCoordinate(const QGeoCoordinate &)

/*
Appends coordinate to the path.
*/
func (this *QGeoPath) AddCoordinate(coordinate QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGeoPath13addCoordinateERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopath.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertCoordinate(int, const QGeoCoordinate &)

/*
Inserts coordinate at the specified index.
*/
func (this *QGeoPath) InsertCoordinate(index int, coordinate QGeoCoordinate_ITF) {
	var convArg1 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg1 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGeoPath16insertCoordinateEiRK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopath.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void replaceCoordinate(int, const QGeoCoordinate &)

/*
Replaces the path element at the specified index with coordinate.
*/
func (this *QGeoPath) ReplaceCoordinate(index int, coordinate QGeoCoordinate_ITF) {
	var convArg1 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg1 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGeoPath17replaceCoordinateEiRK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopath.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate coordinateAt(int) const

/*
Returns the coordinate at index .
*/
func (this *QGeoPath) CoordinateAt(index int) *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGeoPath12coordinateAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeopath.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool containsCoordinate(const QGeoCoordinate &) const

/*
Returns true if the path contains coordinate as one of the elements.
*/
func (this *QGeoPath) ContainsCoordinate(coordinate QGeoCoordinate_ITF) bool {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGeoPath18containsCoordinateERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeopath.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeCoordinate(const QGeoCoordinate &)

/*
Removes the last occurrence of coordinate from the path.
*/
func (this *QGeoPath) RemoveCoordinate(coordinate QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGeoPath16removeCoordinateERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopath.h:87
// index:1
// Public Visibility=Default Availability=Available
// [-2] void removeCoordinate(int)

/*
Removes the last occurrence of coordinate from the path.
*/
func (this *QGeoPath) RemoveCoordinate1(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGeoPath16removeCoordinateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopath.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns the geo path properties as a string.
*/
func (this *QGeoPath) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGeoPath8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
