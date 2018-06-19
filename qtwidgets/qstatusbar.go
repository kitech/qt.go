package qtwidgets

// /usr/include/qt/QtWidgets/qstatusbar.h
// #include <qstatusbar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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

// void showEvent(QShowEvent *)
func (this *QStatusBar) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QStatusBar) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QStatusBar) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void reformat()
func (this *QStatusBar) InheritReformat(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "reformat", f)
}

// void hideOrShow()
func (this *QStatusBar) InheritHideOrShow(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideOrShow", f)
}

// bool event(QEvent *)
func (this *QStatusBar) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QStatusBar struct {
	*QWidget
}
type QStatusBar_ITF interface {
	QWidget_ITF
	QStatusBar_PTR() *QStatusBar
}

func (ptr *QStatusBar) QStatusBar_PTR() *QStatusBar { return ptr }

func (this *QStatusBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QStatusBar) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQStatusBarFromPointer(cthis unsafe.Pointer) *QStatusBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QStatusBar{bcthis0}
}
func (*QStatusBar) NewFromPointer(cthis unsafe.Pointer) *QStatusBar {
	return NewQStatusBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QStatusBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStatusBar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstatusbar.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStatusBar(QWidget *)

/*
Constructs a status bar with a size grip and the given parent.

See also setSizeGripEnabled().
*/
func NewQStatusBar(parent QWidget_ITF /*777 QWidget **/) *QStatusBar {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStatusBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStatusBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qstatusbar.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStatusBar(QWidget *)

/*
Constructs a status bar with a size grip and the given parent.

See also setSizeGripEnabled().
*/
func NewQStatusBar__() *QStatusBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStatusBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStatusBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qstatusbar.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStatusBar()

/*

 */
func DeleteQStatusBar(this *QStatusBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int)

/*
Adds the given widget to this status bar, reparenting the widget if it isn't already a child of this QStatusBar object. The stretch parameter is used to compute a suitable size for the given widget as the status bar grows and shrinks. The default stretch factor is 0, i.e giving the widget a minimum of space.

The widget is located to the far left of the first permanent widget (see addPermanentWidget()) and may be obscured by temporary messages.

See also insertWidget(), removeWidget(), and addPermanentWidget().
*/
func (this *QStatusBar) AddWidget(widget QWidget_ITF /*777 QWidget **/, stretch int) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar9addWidgetEP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int)

/*
Adds the given widget to this status bar, reparenting the widget if it isn't already a child of this QStatusBar object. The stretch parameter is used to compute a suitable size for the given widget as the status bar grows and shrinks. The default stretch factor is 0, i.e giving the widget a minimum of space.

The widget is located to the far left of the first permanent widget (see addPermanentWidget()) and may be obscured by temporary messages.

See also insertWidget(), removeWidget(), and addPermanentWidget().
*/
func (this *QStatusBar) AddWidget__(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar9addWidgetEP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertWidget(int, QWidget *, int)

/*
Inserts the given widget at the given index to this status bar, reparenting the widget if it isn't already a child of this QStatusBar object. If index is out of range, the widget is appended (in which case it is the actual index of the widget that is returned).

The stretch parameter is used to compute a suitable size for the given widget as the status bar grows and shrinks. The default stretch factor is 0, i.e giving the widget a minimum of space.

The widget is located to the far left of the first permanent widget (see addPermanentWidget()) and may be obscured by temporary messages.

This function was introduced in  Qt 4.2.

See also addWidget(), removeWidget(), and addPermanentWidget().
*/
func (this *QStatusBar) InsertWidget(index int, widget QWidget_ITF /*777 QWidget **/, stretch int) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar12insertWidgetEiP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertWidget(int, QWidget *, int)

/*
Inserts the given widget at the given index to this status bar, reparenting the widget if it isn't already a child of this QStatusBar object. If index is out of range, the widget is appended (in which case it is the actual index of the widget that is returned).

The stretch parameter is used to compute a suitable size for the given widget as the status bar grows and shrinks. The default stretch factor is 0, i.e giving the widget a minimum of space.

The widget is located to the far left of the first permanent widget (see addPermanentWidget()) and may be obscured by temporary messages.

This function was introduced in  Qt 4.2.

See also addWidget(), removeWidget(), and addPermanentWidget().
*/
func (this *QStatusBar) InsertWidget__(index int, widget QWidget_ITF /*777 QWidget **/) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar12insertWidgetEiP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPermanentWidget(QWidget *, int)

/*
Adds the given widget permanently to this status bar, reparenting the widget if it isn't already a child of this QStatusBar object. The stretch parameter is used to compute a suitable size for the given widget as the status bar grows and shrinks. The default stretch factor is 0, i.e giving the widget a minimum of space.

Permanently means that the widget may not be obscured by temporary messages. It is is located at the far right of the status bar.

See also insertPermanentWidget(), removeWidget(), and addWidget().
*/
func (this *QStatusBar) AddPermanentWidget(widget QWidget_ITF /*777 QWidget **/, stretch int) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar18addPermanentWidgetEP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPermanentWidget(QWidget *, int)

/*
Adds the given widget permanently to this status bar, reparenting the widget if it isn't already a child of this QStatusBar object. The stretch parameter is used to compute a suitable size for the given widget as the status bar grows and shrinks. The default stretch factor is 0, i.e giving the widget a minimum of space.

Permanently means that the widget may not be obscured by temporary messages. It is is located at the far right of the status bar.

See also insertPermanentWidget(), removeWidget(), and addWidget().
*/
func (this *QStatusBar) AddPermanentWidget__(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar18addPermanentWidgetEP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertPermanentWidget(int, QWidget *, int)

/*
Inserts the given widget at the given index permanently to this status bar, reparenting the widget if it isn't already a child of this QStatusBar object. If index is out of range, the widget is appended (in which case it is the actual index of the widget that is returned).

The stretch parameter is used to compute a suitable size for the given widget as the status bar grows and shrinks. The default stretch factor is 0, i.e giving the widget a minimum of space.

Permanently means that the widget may not be obscured by temporary messages. It is is located at the far right of the status bar.

This function was introduced in  Qt 4.2.

See also addPermanentWidget(), removeWidget(), and addWidget().
*/
func (this *QStatusBar) InsertPermanentWidget(index int, widget QWidget_ITF /*777 QWidget **/, stretch int) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertPermanentWidget(int, QWidget *, int)

/*
Inserts the given widget at the given index permanently to this status bar, reparenting the widget if it isn't already a child of this QStatusBar object. If index is out of range, the widget is appended (in which case it is the actual index of the widget that is returned).

The stretch parameter is used to compute a suitable size for the given widget as the status bar grows and shrinks. The default stretch factor is 0, i.e giving the widget a minimum of space.

Permanently means that the widget may not be obscured by temporary messages. It is is located at the far right of the status bar.

This function was introduced in  Qt 4.2.

See also addPermanentWidget(), removeWidget(), and addWidget().
*/
func (this *QStatusBar) InsertPermanentWidget__(index int, widget QWidget_ITF /*777 QWidget **/) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeWidget(QWidget *)

/*
Removes the specified widget from the status bar.

Note: This function does not delete the widget but hides it. To add the widget again, you must call both the addWidget() and show() functions.

See also addWidget(), addPermanentWidget(), and clearMessage().
*/
func (this *QStatusBar) RemoveWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar12removeWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeGripEnabled(bool)

/*

 */
func (this *QStatusBar) SetSizeGripEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar18setSizeGripEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSizeGripEnabled() const

/*

 */
func (this *QStatusBar) IsSizeGripEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStatusBar17isSizeGripEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstatusbar.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString currentMessage() const

/*
Returns the temporary message currently shown, or an empty string if there is no such message.

See also showMessage().
*/
func (this *QStatusBar) CurrentMessage() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStatusBar14currentMessageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qstatusbar.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, int)

/*
Hides the normal status indications and displays the given message for the specified number of milli-seconds (timeout). If timeout is 0 (default), the message remains displayed until the clearMessage() slot is called or until the showMessage() slot is called again to change the message.

Note that showMessage() is called to show temporary explanations of tool tip texts, so passing a timeout of 0 is not sufficient to display a permanent message.

See also messageChanged(), currentMessage(), and clearMessage().
*/
func (this *QStatusBar) ShowMessage(text string, timeout int) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar11showMessageERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, timeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, int)

/*
Hides the normal status indications and displays the given message for the specified number of milli-seconds (timeout). If timeout is 0 (default), the message remains displayed until the clearMessage() slot is called or until the showMessage() slot is called again to change the message.

Note that showMessage() is called to show temporary explanations of tool tip texts, so passing a timeout of 0 is not sufficient to display a permanent message.

See also messageChanged(), currentMessage(), and clearMessage().
*/
func (this *QStatusBar) ShowMessage__(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	timeout := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar11showMessageERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, timeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMessage()

/*
Removes any temporary message being shown.

See also currentMessage(), showMessage(), and removeWidget().
*/
func (this *QStatusBar) ClearMessage() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar12clearMessageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void messageChanged(const QString &)

/*
This signal is emitted whenever the temporary status message changes. The new temporary message is passed in the message parameter which is a null-string when the message has been removed.

See also showMessage() and clearMessage().
*/
func (this *QStatusBar) MessageChanged(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar14messageChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:82
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QStatusBar) ShowEvent(arg0 qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:83
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().

Shows the temporary message, if appropriate, in response to the paint event.
*/
func (this *QStatusBar) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:84
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QStatusBar) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:87
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void reformat()

/*
Changes the status bar's appearance to account for item changes.

Special subclasses may need this function, but geometry management will usually take care of any necessary rearrangements.
*/
func (this *QStatusBar) Reformat() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar8reformatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:88
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void hideOrShow()

/*
Ensures that the right widgets are visible.

Used by the showMessage() and clearMessage() functions.
*/
func (this *QStatusBar) HideOrShow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar10hideOrShowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:89
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QStatusBar) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
