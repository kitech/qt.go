package qtpositioning

// /usr/include/qt/QtPositioning/qgeocircle.h
// #include <qgeocircle.h>
// #include <QtPositioning>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
type QGeoCircle struct {
	*QGeoShape
}
type QGeoCircle_ITF interface {
	QGeoShape_ITF
	QGeoCircle_PTR() *QGeoCircle
}

func (ptr *QGeoCircle) QGeoCircle_PTR() *QGeoCircle { return ptr }

func (this *QGeoCircle) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGeoShape.GetCthis()
	}
}
func (this *QGeoCircle) SetCthis(cthis unsafe.Pointer) {
	this.QGeoShape = NewQGeoShapeFromPointer(cthis)
}
func NewQGeoCircleFromPointer(cthis unsafe.Pointer) *QGeoCircle {
	bcthis0 := NewQGeoShapeFromPointer(cthis)
	return &QGeoCircle{bcthis0}
}
func (*QGeoCircle) NewFromPointer(cthis unsafe.Pointer) *QGeoCircle {
	return NewQGeoCircleFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeocircle.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoCircle()

/*
Constructs a new, invalid geo circle.
*/
func (*QGeoCircle) NewForInherit() *QGeoCircle {
	return NewQGeoCircle()
}
func NewQGeoCircle() *QGeoCircle {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QGeoCircleC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoCircleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoCircle)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeocircle.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGeoCircle(const QGeoCoordinate &, qreal)

/*
Constructs a new, invalid geo circle.
*/
func (*QGeoCircle) NewForInherit1(center QGeoCoordinate_ITF, radius float64) *QGeoCircle {
	return NewQGeoCircle1(center, radius)
}
func NewQGeoCircle1(center QGeoCoordinate_ITF, radius float64) *QGeoCircle {
	var convArg0 unsafe.Pointer
	if center != nil && center.QGeoCoordinate_PTR() != nil {
		convArg0 = center.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QGeoCircleC2ERK14QGeoCoordinated", qtrt.FFI_TYPE_POINTER, convArg0, radius)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoCircleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoCircle)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeocircle.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGeoCircle(const QGeoCoordinate &, qreal)

/*
Constructs a new, invalid geo circle.
*/
func (*QGeoCircle) NewForInherit1p(center QGeoCoordinate_ITF) *QGeoCircle {
	return NewQGeoCircle1p(center)
}
func NewQGeoCircle1p(center QGeoCoordinate_ITF) *QGeoCircle {
	var convArg0 unsafe.Pointer
	if center != nil && center.QGeoCoordinate_PTR() != nil {
		convArg0 = center.QGeoCoordinate_PTR().GetCthis()
	}
	// arg: 1, qreal=Typedef, qreal=Typedef, double, Double
	radius := float64(-1.0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QGeoCircleC2ERK14QGeoCoordinated", qtrt.FFI_TYPE_POINTER, convArg0, radius)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoCircleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoCircle)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeocircle.h:60
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QGeoCircle(const QGeoShape &)

/*
Constructs a new, invalid geo circle.
*/
func (*QGeoCircle) NewForInherit2(other QGeoShape_ITF) *QGeoCircle {
	return NewQGeoCircle2(other)
}
func NewQGeoCircle2(other QGeoShape_ITF) *QGeoCircle {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoShape_PTR() != nil {
		convArg0 = other.QGeoShape_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QGeoCircleC2ERK9QGeoShape", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoCircleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGeoCircle)
	return gothis
}

// /usr/include/qt/QtPositioning/qgeocircle.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGeoCircle()

/*

 */
func DeleteQGeoCircle(this *QGeoCircle) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QGeoCircleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeocircle.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCircle & operator=(const QGeoCircle &)

/*

 */
func (this *QGeoCircle) Operator_equal(other QGeoCircle_ITF) *QGeoCircle {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoCircle_PTR() != nil {
		convArg0 = other.QGeoCircle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QGeoCircleaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCircleFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCircle)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeocircle.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGeoCircle &) const

/*

 */
func (this *QGeoCircle) Operator_equal_equal(other QGeoCircle_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoCircle_PTR() != nil {
		convArg0 = other.QGeoCircle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QGeoCircleeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeocircle.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QGeoCircle &) const

/*

 */
func (this *QGeoCircle) Operator_not_equal(other QGeoCircle_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGeoCircle_PTR() != nil {
		convArg0 = other.QGeoCircle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QGeoCircleneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeocircle.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCenter(const QGeoCoordinate &)

/*
Sets the center coordinate of this geo circle to center.

Note: Setter function for property center.

See also center().
*/
func (this *QGeoCircle) SetCenter(center QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if center != nil && center.QGeoCoordinate_PTR() != nil {
		convArg0 = center.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QGeoCircle9setCenterERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeocircle.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCoordinate center() const

/*
Returns the center coordinate of this geo circle. Equivalent to QGeoShape::center().

Note: Getter function for property center.

See also setCenter().
*/
func (this *QGeoCircle) Center() *QGeoCoordinate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QGeoCircle6centerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCoordinateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCoordinate)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeocircle.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRadius(qreal)

/*
Sets the radius in meters of this geo circle to radius.

Note: Setter function for property radius.

See also radius().
*/
func (this *QGeoCircle) SetRadius(radius float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QGeoCircle9setRadiusEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), radius)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeocircle.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal radius() const

/*
Returns the radius in meters of this geo circle.

Note: Getter function for property radius.

See also setRadius().
*/
func (this *QGeoCircle) Radius() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QGeoCircle6radiusEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qgeocircle.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(double, double)

/*
Translates this geo circle by degreesLatitude northwards and degreesLongitude eastwards.

Negative values of degreesLatitude and degreesLongitude correspond to southward and westward translation respectively.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.
*/
func (this *QGeoCircle) Translate(degreesLatitude float64, degreesLongitude float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QGeoCircle9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), degreesLatitude, degreesLongitude)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeocircle.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QGeoCircle translated(double, double) const

/*
Returns a copy of this geo circle translated by degreesLatitude northwards and degreesLongitude eastwards.

Negative values of degreesLatitude and degreesLongitude correspond to southward and westward translation respectively.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also translate().
*/
func (this *QGeoCircle) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoCircle /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QGeoCircle10translatedEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), degreesLatitude, degreesLongitude)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoCircleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoCircle)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeocircle.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void extendCircle(const QGeoCoordinate &)

/*
Extends the geo circle to also cover the coordinate coordinate

This function was introduced in  Qt 5.9.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.
*/
func (this *QGeoCircle) ExtendCircle(coordinate QGeoCoordinate_ITF) {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QGeoCircle12extendCircleERK14QGeoCoordinate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeocircle.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns the geo circle properties as a string.

This function was introduced in  Qt 5.5.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.
*/
func (this *QGeoCircle) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QGeoCircle8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
