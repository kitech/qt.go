package qtmultimedia

// /usr/include/qt/QtMultimedia/qcamerafocus.h
// #include <qcamerafocus.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 35
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QCameraFocusZone struct {
	*qtrt.CObject
}
type QCameraFocusZone_ITF interface {
	QCameraFocusZone_PTR() *QCameraFocusZone
}

func (ptr *QCameraFocusZone) QCameraFocusZone_PTR() *QCameraFocusZone { return ptr }

func (this *QCameraFocusZone) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCameraFocusZone) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCameraFocusZoneFromPointer(cthis unsafe.Pointer) *QCameraFocusZone {
	return &QCameraFocusZone{&qtrt.CObject{cthis}}
}
func (*QCameraFocusZone) NewFromPointer(cthis unsafe.Pointer) *QCameraFocusZone {
	return NewQCameraFocusZoneFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCameraFocusZone()

/*

 */
func NewQCameraFocusZone() *QCameraFocusZone {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCameraFocusZoneC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFocusZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCameraFocusZone)
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCameraFocusZone(const QRectF &, QCameraFocusZone::FocusZoneStatus)

/*

 */
func NewQCameraFocusZone_1(area qtcore.QRectF_ITF, status int) *QCameraFocusZone {
	var convArg0 unsafe.Pointer
	if area != nil && area.QRectF_PTR() != nil {
		convArg0 = area.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCameraFocusZoneC2ERK6QRectFNS_15FocusZoneStatusE", qtrt.FFI_TYPE_POINTER, convArg0, status)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFocusZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCameraFocusZone)
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCameraFocusZone(const QRectF &, QCameraFocusZone::FocusZoneStatus)

/*

 */
func NewQCameraFocusZone_1_(area qtcore.QRectF_ITF) *QCameraFocusZone {
	var convArg0 unsafe.Pointer
	if area != nil && area.QRectF_PTR() != nil {
		convArg0 = area.QRectF_PTR().GetCthis()
	}
	// arg: 1, QCameraFocusZone::FocusZoneStatus=Enum, QCameraFocusZone::FocusZoneStatus=Enum, , Invalid
	status := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCameraFocusZoneC2ERK6QRectFNS_15FocusZoneStatusE", qtrt.FFI_TYPE_POINTER, convArg0, status)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFocusZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCameraFocusZone)
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QCameraFocusZone & operator=(const QCameraFocusZone &)

/*

 */
func (this *QCameraFocusZone) Operator_equal(other QCameraFocusZone_ITF) *QCameraFocusZone {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCameraFocusZone_PTR() != nil {
		convArg0 = other.QCameraFocusZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCameraFocusZoneaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCameraFocusZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCameraFocusZone)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QCameraFocusZone &) const

/*

 */
func (this *QCameraFocusZone) Operator_equal_equal(other QCameraFocusZone_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCameraFocusZone_PTR() != nil {
		convArg0 = other.QCameraFocusZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QCameraFocusZoneeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QCameraFocusZone &) const

/*

 */
func (this *QCameraFocusZone) Operator_not_equal(other QCameraFocusZone_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCameraFocusZone_PTR() != nil {
		convArg0 = other.QCameraFocusZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QCameraFocusZoneneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCameraFocusZone()

/*

 */
func DeleteQCameraFocusZone(this *QCameraFocusZone) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCameraFocusZoneD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QCameraFocusZone) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QCameraFocusZone7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:81
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF area() const

/*

 */
func (this *QCameraFocusZone) Area() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QCameraFocusZone4areaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QCameraFocusZone::FocusZoneStatus status() const

/*

 */
func (this *QCameraFocusZone) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QCameraFocusZone6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatus(QCameraFocusZone::FocusZoneStatus)

/*

 */
func (this *QCameraFocusZone) SetStatus(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCameraFocusZone9setStatusENS_15FocusZoneStatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QCameraFocusZone__FocusZoneStatus = int

//
const QCameraFocusZone__Invalid QCameraFocusZone__FocusZoneStatus = 0

//
const QCameraFocusZone__Unused QCameraFocusZone__FocusZoneStatus = 1

//
const QCameraFocusZone__Selected QCameraFocusZone__FocusZoneStatus = 2

//
const QCameraFocusZone__Focused QCameraFocusZone__FocusZoneStatus = 3

func (this *QCameraFocusZone) FocusZoneStatusItemName(val int) string {
	switch val {
	case QCameraFocusZone__Invalid: // 0
		return "Invalid"
	case QCameraFocusZone__Unused: // 1
		return "Unused"
	case QCameraFocusZone__Selected: // 2
		return "Selected"
	case QCameraFocusZone__Focused: // 3
		return "Focused"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QCameraFocusZone_FocusZoneStatusItemName(val int) string {
	var nilthis *QCameraFocusZone
	return nilthis.FocusZoneStatusItemName(val)
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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
