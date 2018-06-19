package qtwidgets

// /usr/include/qt/QtWidgets/qdial.h
// #include <qdial.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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

// bool event(QEvent *)
func (this *QDial) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QDial) InheritResizeEvent(f func(re *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QDial) InheritPaintEvent(f func(pe *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QDial) InheritMousePressEvent(f func(me *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QDial) InheritMouseReleaseEvent(f func(me *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QDial) InheritMouseMoveEvent(f func(me *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void sliderChange(QAbstractSlider::SliderChange)
func (this *QDial) InheritSliderChange(f func(change int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "sliderChange", f)
}

// void initStyleOption(QStyleOptionSlider *)
func (this *QDial) InheritInitStyleOption(f func(option *QStyleOptionSlider /*777 QStyleOptionSlider **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QDial struct {
	*QAbstractSlider
}
type QDial_ITF interface {
	QAbstractSlider_ITF
	QDial_PTR() *QDial
}

func (ptr *QDial) QDial_PTR() *QDial { return ptr }

func (this *QDial) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSlider.GetCthis()
	}
}
func (this *QDial) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSlider = NewQAbstractSliderFromPointer(cthis)
}
func NewQDialFromPointer(cthis unsafe.Pointer) *QDial {
	bcthis0 := NewQAbstractSliderFromPointer(cthis)
	return &QDial{bcthis0}
}
func (*QDial) NewFromPointer(cthis unsafe.Pointer) *QDial {
	return NewQDialFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdial.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QDial) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDial10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdial.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDial(QWidget *)

/*
Constructs a dial.

The parent argument is sent to the QAbstractSlider constructor.
*/
func NewQDial(parent QWidget_ITF /*777 QWidget **/) *QDial {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDialC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDialFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDial")
	return gothis
}

// /usr/include/qt/QtWidgets/qdial.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDial(QWidget *)

/*
Constructs a dial.

The parent argument is sent to the QAbstractSlider constructor.
*/
func NewQDial__() *QDial {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDialC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDialFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDial")
	return gothis
}

// /usr/include/qt/QtWidgets/qdial.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDial()

/*

 */
func DeleteQDial(this *QDial) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDialD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdial.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wrapping() const

/*

 */
func (this *QDial) Wrapping() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDial8wrappingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdial.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] int notchSize() const

/*

 */
func (this *QDial) NotchSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDial9notchSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdial.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNotchTarget(double)

/*

 */
func (this *QDial) SetNotchTarget(target float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDial14setNotchTargetEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), target)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal notchTarget() const

/*

 */
func (this *QDial) NotchTarget() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDial11notchTargetEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qdial.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool notchesVisible() const

/*

 */
func (this *QDial) NotchesVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDial14notchesVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdial.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QDial) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDial8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qdial.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QDial) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDial15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qdial.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNotchesVisible(bool)

/*

 */
func (this *QDial) SetNotchesVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDial17setNotchesVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWrapping(bool)

/*

 */
func (this *QDial) SetWrapping(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDial11setWrappingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:84
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QDial) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDial5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdial.h:85
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QDial) ResizeEvent(re qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if re != nil && re.QResizeEvent_PTR() != nil {
		convArg0 = re.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDial11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:86
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QDial) PaintEvent(pe qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if pe != nil && pe.QPaintEvent_PTR() != nil {
		convArg0 = pe.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDial10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:88
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QDial) MousePressEvent(me qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if me != nil && me.QMouseEvent_PTR() != nil {
		convArg0 = me.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDial15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:89
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QDial) MouseReleaseEvent(me qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if me != nil && me.QMouseEvent_PTR() != nil {
		convArg0 = me.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDial17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QDial) MouseMoveEvent(me qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if me != nil && me.QMouseEvent_PTR() != nil {
		convArg0 = me.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDial14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:92
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void sliderChange(QAbstractSlider::SliderChange)

/*
Reimplemented from QAbstractSlider::sliderChange().
*/
func (this *QDial) SliderChange(change int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDial12sliderChangeEN15QAbstractSlider12SliderChangeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), change)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:93
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionSlider *) const

/*
Initialize option with the values from this QDial. This method is useful for subclasses when they need a QStyleOptionSlider, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QDial) InitStyleOption(option QStyleOptionSlider_ITF /*777 QStyleOptionSlider **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionSlider_PTR() != nil {
		convArg0 = option.QStyleOptionSlider_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDial15initStyleOptionEP18QStyleOptionSlider", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
}

//  keep block end
