package qtgui

// /usr/include/qt/QtGui/qrasterwindow.h
// #include <qrasterwindow.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

// int metric(QPaintDevice::PaintDeviceMetric)
func (this *QRasterWindow) InheritMetric(f func(metric int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

// QPaintDevice * redirected(QPoint *)
func (this *QRasterWindow) InheritRedirected(f func(arg0 *qtcore.QPoint /*777 QPoint **/) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "redirected", f)
}

/*

 */
type QRasterWindow struct {
	*QPaintDeviceWindow
}
type QRasterWindow_ITF interface {
	QPaintDeviceWindow_ITF
	QRasterWindow_PTR() *QRasterWindow
}

func (ptr *QRasterWindow) QRasterWindow_PTR() *QRasterWindow { return ptr }

func (this *QRasterWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPaintDeviceWindow.GetCthis()
	}
}
func (this *QRasterWindow) SetCthis(cthis unsafe.Pointer) {
	this.QPaintDeviceWindow = NewQPaintDeviceWindowFromPointer(cthis)
}
func NewQRasterWindowFromPointer(cthis unsafe.Pointer) *QRasterWindow {
	bcthis0 := NewQPaintDeviceWindowFromPointer(cthis)
	return &QRasterWindow{bcthis0}
}
func (*QRasterWindow) NewFromPointer(cthis unsafe.Pointer) *QRasterWindow {
	return NewQRasterWindowFromPointer(cthis)
}

// /usr/include/qt/QtGui/qrasterwindow.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QRasterWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QRasterWindow10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qrasterwindow.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRasterWindow(QWindow *)

/*
Constructs a new QRasterWindow with parent.
*/
func (*QRasterWindow) NewForInherit(parent QWindow_ITF /*777 QWindow **/) *QRasterWindow {
	return NewQRasterWindow(parent)
}
func NewQRasterWindow(parent QWindow_ITF /*777 QWindow **/) *QRasterWindow {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWindow_PTR() != nil {
		convArg0 = parent.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QRasterWindowC2EP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRasterWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRasterWindow")
	return gothis
}

// /usr/include/qt/QtGui/qrasterwindow.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRasterWindow(QWindow *)

/*
Constructs a new QRasterWindow with parent.
*/
func (*QRasterWindow) NewForInheritp() *QRasterWindow {
	return NewQRasterWindowp()
}
func NewQRasterWindowp() *QRasterWindow {
	// arg: 0, QWindow *=Pointer, QWindow=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QRasterWindowC2EP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRasterWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRasterWindow")
	return gothis
}

// /usr/include/qt/QtGui/qrasterwindow.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRasterWindow()

/*

 */
func DeleteQRasterWindow(this *QRasterWindow) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QRasterWindowD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 64)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qrasterwindow.h:60
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(QPaintDevice::PaintDeviceMetric) const

/*

 */
func (this *QRasterWindow) Metric(metric int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QRasterWindow6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrasterwindow.h:61
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPaintDevice * redirected(QPoint *) const

/*

 */
func (this *QRasterWindow) Redirected(arg0 qtcore.QPoint_ITF /*777 QPoint **/) *QPaintDevice /*777 QPaintDevice **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QRasterWindow10redirectedEP6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_10929() {
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
