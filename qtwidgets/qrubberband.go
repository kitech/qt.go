package qtwidgets

// /usr/include/qt/QtWidgets/qrubberband.h
// #include <qrubberband.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 65
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QRubberBand) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QRubberBand) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QRubberBand) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QRubberBand) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QRubberBand) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void moveEvent(class QMoveEvent *)
func (this *QRubberBand) InheritMoveEvent(f func(arg0 *qtgui.QMoveEvent /*777 QMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "moveEvent", f)
}

// void initStyleOption(class QStyleOptionRubberBand *)
func (this *QRubberBand) InheritInitStyleOption(f func(option *QStyleOptionRubberBand /*777 QStyleOptionRubberBand **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QRubberBand struct {
	*QWidget
}
type QRubberBand_ITF interface {
	QWidget_ITF
	QRubberBand_PTR() *QRubberBand
}

func (ptr *QRubberBand) QRubberBand_PTR() *QRubberBand { return ptr }

func (this *QRubberBand) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QRubberBand) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQRubberBandFromPointer(cthis unsafe.Pointer) *QRubberBand {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QRubberBand{bcthis0}
}
func (*QRubberBand) NewFromPointer(cthis unsafe.Pointer) *QRubberBand {
	return NewQRubberBandFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qrubberband.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QRubberBand) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRubberBand10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qrubberband.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRubberBand(enum QRubberBand::Shape, QWidget *)

/*
Constructs a rubber band of shape s, with parent p.

By default a rectangular rubber band (s is Rectangle) will use a mask, so that a small border of the rectangle is all that is visible. Some styles (e.g., native macOS) will change this and call QWidget::setWindowOpacity() to make a semi-transparent filled selection rectangle.
*/
func NewQRubberBand(arg0 int, arg1 QWidget_ITF /*777 QWidget **/) *QRubberBand {
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QWidget_PTR() != nil {
		convArg1 = arg1.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBandC2ENS_5ShapeEP7QWidget", qtrt.FFI_TYPE_POINTER, arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRubberBandFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRubberBand")
	return gothis
}

// /usr/include/qt/QtWidgets/qrubberband.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRubberBand(enum QRubberBand::Shape, QWidget *)

/*
Constructs a rubber band of shape s, with parent p.

By default a rectangular rubber band (s is Rectangle) will use a mask, so that a small border of the rectangle is all that is visible. Some styles (e.g., native macOS) will change this and call QWidget::setWindowOpacity() to make a semi-transparent filled selection rectangle.
*/
func NewQRubberBand__(arg0 int) *QRubberBand {
	// arg: 1, QWidget *=Pointer, QWidget=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBandC2ENS_5ShapeEP7QWidget", qtrt.FFI_TYPE_POINTER, arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRubberBandFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRubberBand")
	return gothis
}

// /usr/include/qt/QtWidgets/qrubberband.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRubberBand()

/*

 */
func DeleteQRubberBand(this *QRubberBand) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBandD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qrubberband.h:62
// index:0
// Public Visibility=Default Availability=Available
// [4] QRubberBand::Shape shape() const

/*
Returns the shape of this rubber band. The shape can only be set upon construction.
*/
func (this *QRubberBand) Shape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRubberBand5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)

/*
Sets the geometry of the rubber band to rect, specified in the coordinate system of its parent widget.

See also QWidget::geometry.
*/
func (this *QRubberBand) SetGeometry(r qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:66
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setGeometry(int, int, int, int)

/*
Sets the geometry of the rubber band to rect, specified in the coordinate system of its parent widget.

See also QWidget::geometry.
*/
func (this *QRubberBand) SetGeometry_1(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand11setGeometryEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void move(int, int)

/*
Moves the rubberband to point (x, y).

See also resize().
*/
func (this *QRubberBand) Move(x int, y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand4moveEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:68
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void move(const QPoint &)

/*
Moves the rubberband to point (x, y).

See also resize().
*/
func (this *QRubberBand) Move_1(p qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand4moveERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void resize(int, int)

/*
Resizes the rubberband so that its width is width, and its height is height.

See also move().
*/
func (this *QRubberBand) Resize(w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand6resizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:72
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void resize(const QSize &)

/*
Resizes the rubberband so that its width is width, and its height is height.

See also move().
*/
func (this *QRubberBand) Resize_1(s qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand6resizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:76
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QRubberBand) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qrubberband.h:77
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QRubberBand) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:78
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QRubberBand) ChangeEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:79
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QRubberBand) ShowEvent(arg0 qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:80
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QRubberBand) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:81
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void moveEvent(QMoveEvent *)

/*
Reimplemented from QWidget::moveEvent().
*/
func (this *QRubberBand) MoveEvent(arg0 qtgui.QMoveEvent_ITF /*777 QMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMoveEvent_PTR() != nil {
		convArg0 = arg0.QMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRubberBand9moveEventEP10QMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:82
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionRubberBand *) const

/*
Initialize option with the values from this QRubberBand. This method is useful for subclasses when they need a QStyleOptionRubberBand, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QRubberBand) InitStyleOption(option QStyleOptionRubberBand_ITF /*777 QStyleOptionRubberBand **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionRubberBand_PTR() != nil {
		convArg0 = option.QStyleOptionRubberBand_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRubberBand15initStyleOptionEP22QStyleOptionRubberBand", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum specifies what shape a QRubberBand should have. This is a drawing hint that is passed down to the style system, and can be interpreted by each QStyle.


*/
type QRubberBand__Shape = int

// A QRubberBand can represent a vertical or horizontal line. Geometry is still given in rect() and the line will fill the given geometry on most styles.
const QRubberBand__Line QRubberBand__Shape = 0

// A QRubberBand can represent a rectangle. Some styles will interpret this as a filled (often semi-transparent) rectangle, or a rectangular outline.
const QRubberBand__Rectangle QRubberBand__Shape = 1

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
}

//  keep block end
