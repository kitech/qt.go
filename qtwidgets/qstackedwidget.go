// +build !minimal

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
// extern C begin: 10
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

/*
 */
// size 48
type QStackedWidget struct {
	*QFrame
}
type QStackedWidget_ITF interface {
	QFrame_ITF
	QStackedWidget_PTR() *QStackedWidget
}

func (ptr *QStackedWidget) QStackedWidget_PTR() *QStackedWidget { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QStackedWidgetFromptr(cthis Voidptr) *QStackedWidget {
	bcthis0 := QFrameFromptr(cthis)
	return &QStackedWidget{bcthis0}
}
func (*QStackedWidget) Fromptr(cthis Voidptr) *QStackedWidget {
	return QStackedWidgetFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStackedWidget(QWidget *)

/*
 */
func (*QStackedWidget) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QStackedWidget {
	return NewQStackedWidget(parent)
}
func NewQStackedWidget(parent QWidget_ITF /*777 QWidget **/) *QStackedWidget {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(4054824524, "_ZN14QStackedWidgetC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QStackedWidgetFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QStackedWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStackedWidget(QWidget *)

/*
 */
func (*QStackedWidget) NewForInheritp() *QStackedWidget {
	return NewQStackedWidgetp()
}
func NewQStackedWidgetp() *QStackedWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(4054824524, "_ZN14QStackedWidgetC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QStackedWidgetFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QStackedWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:62
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int addWidget(QWidget *)

/*
 */
func (this *QStackedWidget) AddWidget(w QWidget_ITF /*777 QWidget **/) int {
	var convArg0 Voidptr
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1249653093, "_ZN14QStackedWidget9addWidgetEP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:63
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int insertWidget(int, QWidget *)

/*
 */
func (this *QStackedWidget) InsertWidget(index int, w QWidget_ITF /*777 QWidget **/) int {
	var convArg1 Voidptr
	if w != nil && w.QWidget_PTR() != nil {
		convArg1 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2385140436, "_ZN14QStackedWidget12insertWidgetEiP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&index), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:64
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void removeWidget(QWidget *)

/*
 */
func (this *QStackedWidget) RemoveWidget(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 Voidptr
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(598937263, "_ZN14QStackedWidget12removeWidgetEP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:66
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QWidget * currentWidget() const

/*
 */
func (this *QStackedWidget) CurrentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.Qtcc3(4210146573, "_ZNK14QStackedWidget13currentWidgetEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return /*==*/ QWidgetFromptr(rv.Ptr()) // 444
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:67
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int currentIndex() const

/*
 */
func (this *QStackedWidget) CurrentIndex() int {
	rv, err := qtrt.Qtcc3(2777648257, "_ZNK14QStackedWidget12currentIndexEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:69
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int indexOf(QWidget *) const

/*
 */
func (this *QStackedWidget) IndexOf(arg0 QWidget_ITF /*777 QWidget **/) int {
	var convArg0 Voidptr
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2588077492, "_ZNK14QStackedWidget7indexOfEP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:70
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QWidget * widget(int) const

/*
 */
func (this *QStackedWidget) Widget(arg0 int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.Qtcc3(137583665, "_ZNK14QStackedWidget6widgetEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint3(err, rv)
	return /*==*/ QWidgetFromptr(rv.Ptr()) // 444
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:71
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int count() const

/*
 */
func (this *QStackedWidget) Count() int {
	rv, err := qtrt.Qtcc3(3735907962, "_ZNK14QStackedWidget5countEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:74
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)

/*
 */
func (this *QStackedWidget) SetCurrentIndex(index int) {
	rv, err := qtrt.Qtcc3(3030298062, "_ZN14QStackedWidget15setCurrentIndexEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:75
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCurrentWidget(QWidget *)

/*
 */
func (this *QStackedWidget) SetCurrentWidget(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 Voidptr
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1926882632, "_ZN14QStackedWidget16setCurrentWidgetEP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:78
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void currentChanged(int)

/*
 */
func (this *QStackedWidget) CurrentChanged(arg0 int) {
	rv, err := qtrt.Qtcc3(3435686077, "_ZN14QStackedWidget14currentChangedEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:79
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void widgetRemoved(int)

/*
 */
func (this *QStackedWidget) WidgetRemoved(index int) {
	rv, err := qtrt.Qtcc3(3890011385, "_ZN14QStackedWidget13widgetRemovedEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint3(err, rv)
}

func DeleteQStackedWidget(this *QStackedWidget) {
	rv, err := qtrt.Qtcc3(2840884496, "_ZN14QStackedWidgetD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10249() {
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
