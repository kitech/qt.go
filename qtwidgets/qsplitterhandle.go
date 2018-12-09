// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qsplitter.h
// #include <qsplitter.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 37
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

// void paintEvent(QPaintEvent *)
func (this *QSplitterHandle) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QSplitterHandle) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QSplitterHandle) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QSplitterHandle) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QSplitterHandle) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// bool event(QEvent *)
func (this *QSplitterHandle) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void moveSplitter(int)
func (this *QSplitterHandle) InheritMoveSplitter(f func(p int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "moveSplitter", f)
}

// int closestLegalPosition(int)
func (this *QSplitterHandle) InheritClosestLegalPosition(f func(p int) int) {
	qtrt.SetAllInheritCallback(this, "closestLegalPosition", f)
}

/*

 */
type QSplitterHandle struct {
	*QWidget
}
type QSplitterHandle_ITF interface {
	QWidget_ITF
	QSplitterHandle_PTR() *QSplitterHandle
}

func (ptr *QSplitterHandle) QSplitterHandle_PTR() *QSplitterHandle { return ptr }

func (this *QSplitterHandle) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QSplitterHandle) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQSplitterHandleFromPointer(cthis unsafe.Pointer) *QSplitterHandle {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QSplitterHandle{bcthis0}
}
func (*QSplitterHandle) NewFromPointer(cthis unsafe.Pointer) *QSplitterHandle {
	return NewQSplitterHandleFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qsplitter.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSplitterHandle) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSplitterHandle10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplitter.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSplitterHandle(Qt::Orientation, QSplitter *)

/*

 */
func (*QSplitterHandle) NewForInherit(o int, parent QSplitter_ITF /*777 QSplitter **/) *QSplitterHandle {
	return NewQSplitterHandle(o, parent)
}
func NewQSplitterHandle(o int, parent QSplitter_ITF /*777 QSplitter **/) *QSplitterHandle {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QSplitter_PTR() != nil {
		convArg1 = parent.QSplitter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSplitterHandleC2EN2Qt11OrientationEP9QSplitter", qtrt.FFI_TYPE_POINTER, o, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplitterHandleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplitterHandle")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplitter.h:141
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSplitterHandle()

/*

 */
func DeleteQSplitterHandle(this *QSplitterHandle) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSplitterHandleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qsplitter.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)

/*

 */
func (this *QSplitterHandle) SetOrientation(o int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSplitterHandle14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), o)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:144
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation() const

/*

 */
func (this *QSplitterHandle) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSplitterHandle11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:145
// index:0
// Public Visibility=Default Availability=Available
// [1] bool opaqueResize() const

/*

 */
func (this *QSplitterHandle) OpaqueResize() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSplitterHandle12opaqueResizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:146
// index:0
// Public Visibility=Default Availability=Available
// [8] QSplitter * splitter() const

/*

 */
func (this *QSplitterHandle) Splitter() *QSplitter /*777 QSplitter **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSplitterHandle8splitterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSplitterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplitter.h:148
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QSplitterHandle) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSplitterHandle8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:151
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*

 */
func (this *QSplitterHandle) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSplitterHandle10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:152
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*

 */
func (this *QSplitterHandle) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSplitterHandle14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:153
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*

 */
func (this *QSplitterHandle) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSplitterHandle15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:154
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*

 */
func (this *QSplitterHandle) MouseReleaseEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSplitterHandle17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:155
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QSplitterHandle) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSplitterHandle11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:156
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QSplitterHandle) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSplitterHandle5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:158
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void moveSplitter(int)

/*
Moves the left or top edge of the splitter handle at index as close as possible to position pos, which is the distance from the left or top edge of the widget.

For right-to-left languages such as Arabic and Hebrew, the layout of horizontal splitters is reversed. pos is then the distance from the right edge of the widget.

See also splitterMoved(), closestLegalPosition(), and getRange().
*/
func (this *QSplitterHandle) MoveSplitter(p int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSplitterHandle12moveSplitterEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), p)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:159
// index:0
// Protected Visibility=Default Availability=Available
// [4] int closestLegalPosition(int)

/*
Returns the closest legal position to pos of the widget at index.

For right-to-left languages such as Arabic and Hebrew, the layout of horizontal splitters is reversed. Positions are then measured from the right edge of the widget.

See also getRange().
*/
func (this *QSplitterHandle) ClosestLegalPosition(p int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSplitterHandle20closestLegalPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), p)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
