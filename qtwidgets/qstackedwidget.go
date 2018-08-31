package qtwidgets

// /usr/include/qt/QtWidgets/qstackedwidget.h
// #include <qstackedwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
func (this *QStackedWidget) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QStackedWidget struct {
	*QFrame
}
type QStackedWidget_ITF interface {
	QFrame_ITF
	QStackedWidget_PTR() *QStackedWidget
}

func (ptr *QStackedWidget) QStackedWidget_PTR() *QStackedWidget { return ptr }

func (this *QStackedWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QFrame.GetCthis()
	}
}
func (this *QStackedWidget) SetCthis(cthis unsafe.Pointer) {
	this.QFrame = NewQFrameFromPointer(cthis)
}
func NewQStackedWidgetFromPointer(cthis unsafe.Pointer) *QStackedWidget {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QStackedWidget{bcthis0}
}
func (*QStackedWidget) NewFromPointer(cthis unsafe.Pointer) *QStackedWidget {
	return NewQStackedWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QStackedWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStackedWidget(QWidget *)

/*
Constructs a QStackedWidget with the given parent.

See also addWidget() and insertWidget().
*/
func (*QStackedWidget) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QStackedWidget {
	return NewQStackedWidget(parent)
}
func NewQStackedWidget(parent QWidget_ITF /*777 QWidget **/) *QStackedWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStackedWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStackedWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStackedWidget(QWidget *)

/*
Constructs a QStackedWidget with the given parent.

See also addWidget() and insertWidget().
*/
func (*QStackedWidget) NewForInherit__() *QStackedWidget {
	return NewQStackedWidget__()
}
func NewQStackedWidget__() *QStackedWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStackedWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStackedWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStackedWidget()

/*

 */
func DeleteQStackedWidget(this *QStackedWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:62
// index:0
// Public Visibility=Default Availability=Available
// [4] int addWidget(QWidget *)

/*
Appends the given widget to the QStackedWidget and returns the index position. Ownership of widget is passed on to the QStackedWidget.

If the QStackedWidget is empty before this function is called, widget becomes the current widget.

See also insertWidget(), removeWidget(), and setCurrentWidget().
*/
func (this *QStackedWidget) AddWidget(w QWidget_ITF /*777 QWidget **/) int {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget9addWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertWidget(int, QWidget *)

/*
Inserts the given widget at the given index in the QStackedWidget. Ownership of widget is passed on to the QStackedWidget. If index is out of range, the widget is appended (in which case it is the actual index of the widget that is returned).

If the QStackedWidget was empty before this function is called, the given widget becomes the current widget.

Inserting a new widget at an index less than or equal to the current index will increment the current index, but keep the current widget.

See also addWidget(), removeWidget(), and setCurrentWidget().
*/
func (this *QStackedWidget) InsertWidget(index int, w QWidget_ITF /*777 QWidget **/) int {
	var convArg1 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg1 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget12insertWidgetEiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeWidget(QWidget *)

/*
Removes widget from the QStackedWidget. i.e., widget is not deleted but simply removed from the stacked layout, causing it to be hidden.

Note: Parent object and parent widget of widget will remain the QStackedWidget. If the application wants to reuse the removed widget, then it is recommended to re-parent it.

See also addWidget(), insertWidget(), and currentWidget().
*/
func (this *QStackedWidget) RemoveWidget(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget12removeWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * currentWidget() const

/*
Returns the current widget, or 0 if there are no child widgets.

See also currentIndex() and setCurrentWidget().
*/
func (this *QStackedWidget) CurrentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget13currentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex() const

/*

 */
func (this *QStackedWidget) CurrentIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(QWidget *) const

/*
Returns the index of the given widget, or -1 if the given widget is not a child of the QStackedWidget.

See also currentIndex() and widget().
*/
func (this *QStackedWidget) IndexOf(arg0 QWidget_ITF /*777 QWidget **/) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget7indexOfEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget(int) const

/*
Returns the widget at the given index, or 0 if there is no such widget.

See also currentWidget() and indexOf().
*/
func (this *QStackedWidget) Widget(arg0 int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget6widgetEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:71
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*

 */
func (this *QStackedWidget) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)

/*

 */
func (this *QStackedWidget) SetCurrentIndex(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget15setCurrentIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentWidget(QWidget *)

/*
Sets the current widget to be the specified widget. The new current widget must already be contained in this stacked widget.

See also currentWidget() and setCurrentIndex().
*/
func (this *QStackedWidget) SetCurrentWidget(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget16setCurrentWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentChanged(int)

/*
This signal is emitted whenever the current widget changes.

The parameter holds the index of the new current widget, or -1 if there isn't a new one (for example, if there are no widgets in the QStackedWidget).

Note: Notifier signal for property currentIndex.

See also currentWidget() and setCurrentWidget().
*/
func (this *QStackedWidget) CurrentChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget14currentChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void widgetRemoved(int)

/*
This signal is emitted whenever a widget is removed. The widget's index is passed as parameter.

See also removeWidget().
*/
func (this *QStackedWidget) WidgetRemoved(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget13widgetRemovedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:82
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QStackedWidget) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
