package qtgui

import "unsafe"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

func init() {
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtGui/qrgb.h:78
// index:0
// Invalid inline Visibility=Default Availability=Available
// [1] bool qIsGray(QRgb)
func QIsGray(rgb uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z7qIsGrayj", qtrt.FFI_TYPE_POINTER, rgb)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qrgb.h:57
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qGreen(QRgb)
func QGreen(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z6qGreenj", qtrt.FFI_TYPE_POINTER, rgb)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:63
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qAlpha(QRgb)
func QAlpha(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z6qAlphaj", qtrt.FFI_TYPE_POINTER, rgb)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:69
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qRgba(int, int, int, int)
func QRgba(r int, g int, b int, a int) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qRgbaiiii", qtrt.FFI_TYPE_POINTER, r, g, b, a)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qrgb.h:75
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qGray(QRgb)
func QGray(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z5qGrayj", qtrt.FFI_TYPE_POINTER, rgb)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:72
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] int qGray(int, int, int)
func QGray_1(r int, g int, b int) int {
	rv, err := qtrt.InvokeQtFunc6("_Z5qGrayiii", qtrt.FFI_TYPE_POINTER, r, g, b)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:60
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qBlue(QRgb)
func QBlue(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z5qBluej", qtrt.FFI_TYPE_POINTER, rgb)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:66
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qRgb(int, int, int)
func QRgb(r int, g int, b int) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z4qRgbiii", qtrt.FFI_TYPE_POINTER, r, g, b)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qrgb.h:54
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qRed(QRgb)
func QRed(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z4qRedj", qtrt.FFI_TYPE_POINTER, rgb)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:974
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qAccessibleLocalizedActionDescription(const QString &)
func QAccessibleLocalizedActionDescription(actionName *qtcore.QString) *qtcore.QString /*123*/ {
	var convArg0 = actionName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z37qAccessibleLocalizedActionDescriptionRK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:973
// index:0
// Invalid Visibility=Default Availability=Available
// [8] const char * qAccessibleEventString(QAccessible::Event)
func QAccessibleEventString(event int) string {
	rv, err := qtrt.InvokeQtFunc6("_Z22qAccessibleEventStringN11QAccessible5EventE", qtrt.FFI_TYPE_POINTER, event)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:972
// index:0
// Invalid Visibility=Default Availability=Available
// [8] const char * qAccessibleRoleString(QAccessible::Role)
func QAccessibleRoleString(role int) string {
	rv, err := qtrt.InvokeQtFunc6("_Z21qAccessibleRoleStringN11QAccessible4RoleE", qtrt.FFI_TYPE_POINTER, role)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qrgb.h:96
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qUnpremultiply(QRgb)
func QUnpremultiply(p uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z14qUnpremultiplyj", qtrt.FFI_TYPE_POINTER, p)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qquaternion.h:323
// index:2
// Invalid inline Visibility=Default Availability=Available
// [1] bool qFuzzyCompare(const QQuaternion &, const QQuaternion &)
func QFuzzyCompare_2(q1 *QQuaternion, q2 *QQuaternion) bool {
	var convArg0 = q1.GetCthis()
	var convArg1 = q2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z13qFuzzyCompareRK11QQuaternionS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qrgb.h:81
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qPremultiply(QRgb)
func QPremultiply(x uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z12qPremultiplyj", qtrt.FFI_TYPE_POINTER, x)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

//  body block end
