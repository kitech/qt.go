package qtpositioning

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

func init_unused_10013() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:111
// index:87
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QGeoPositionInfoSource::PositioningMethods::enum_type, int)

/*

 */
func Operator_or87(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN22QGeoPositionInfoSource17PositioningMethodEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeocoordinate.h:125
// index:51
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QGeoCoordinate &, uint)

/*

 */
func QHash51(coordinate QGeoCoordinate_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK14QGeoCoordinatej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

//  body block end
