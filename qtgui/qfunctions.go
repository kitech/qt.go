package qtgui

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

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
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrgb.h:57
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qGreen(QRgb)
func QGreen(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z6qGreenj", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:63
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qAlpha(QRgb)
func QAlpha(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z6qAlphaj", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:69
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qRgba(int, int, int, int)
func QRgba(r int, g int, b int, a int) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qRgbaiiii", qtrt.FFI_TYPE_POINTER, r, g, b, a)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qrgb.h:75
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qGray(QRgb)
func QGray(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z5qGrayj", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:72
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] int qGray(int, int, int)
func QGray_1(r int, g int, b int) int {
	rv, err := qtrt.InvokeQtFunc6("_Z5qGrayiii", qtrt.FFI_TYPE_POINTER, r, g, b)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:60
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qBlue(QRgb)
func QBlue(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z5qBluej", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:66
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qRgb(int, int, int)
func QRgb(r int, g int, b int) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z4qRgbiii", qtrt.FFI_TYPE_POINTER, r, g, b)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qrgb.h:54
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qRed(QRgb)
func QRed(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z4qRedj", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:974
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qAccessibleLocalizedActionDescription(const QString &)
func QAccessibleLocalizedActionDescription(actionName string) string {
	var tmpArg0 = qtcore.NewQString_5(actionName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z37qAccessibleLocalizedActionDescriptionRK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:973
// index:0
// Invalid Visibility=Default Availability=Available
// [8] const char * qAccessibleEventString(QAccessible::Event)
func QAccessibleEventString(event int) string {
	rv, err := qtrt.InvokeQtFunc6("_Z22qAccessibleEventStringN11QAccessible5EventE", qtrt.FFI_TYPE_POINTER, event)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:972
// index:0
// Invalid Visibility=Default Availability=Available
// [8] const char * qAccessibleRoleString(QAccessible::Role)
func QAccessibleRoleString(role int) string {
	rv, err := qtrt.InvokeQtFunc6("_Z21qAccessibleRoleStringN11QAccessible4RoleE", qtrt.FFI_TYPE_POINTER, role)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qicon.h:149
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qt_findAtNxFile(const QString &, qreal, qreal *)
func Qt_findAtNxFile(baseFileName string, targetDevicePixelRatio float64, sourceDevicePixelRatio unsafe.Pointer /*666*/) string {
	var tmpArg0 = qtcore.NewQString_5(baseFileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z15qt_findAtNxFileRK7QStringdPd", qtrt.FFI_TYPE_POINTER, convArg0, targetDevicePixelRatio, sourceDevicePixelRatio)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qrgb.h:96
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qUnpremultiply(QRgb)
func QUnpremultiply(p uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z14qUnpremultiplyj", qtrt.FFI_TYPE_POINTER, p)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qquaternion.h:323
// index:2
// Invalid inline Visibility=Default Availability=Available
// [1] bool qFuzzyCompare(const QQuaternion &, const QQuaternion &)
func QFuzzyCompare_2(q1 QQuaternion_ITF, q2 QQuaternion_ITF) bool {
	var convArg0 unsafe.Pointer
	if q1 != nil && q1.QQuaternion_PTR() != nil {
		convArg0 = q1.QQuaternion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if q2 != nil && q2.QQuaternion_PTR() != nil {
		convArg1 = q2.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z13qFuzzyCompareRK11QQuaternionS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrgb.h:81
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qPremultiply(QRgb)
func QPremultiply(x uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z12qPremultiplyj", qtrt.FFI_TYPE_POINTER, x)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

//  body block end
