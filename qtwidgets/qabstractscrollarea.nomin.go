//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractscrollarea.h
// #include <qabstractscrollarea.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractScrollArea(QWidget *)

/*
 */
func (*QAbstractScrollArea) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QAbstractScrollArea {
	return NewQAbstractScrollArea(parent)
}
func NewQAbstractScrollArea(parent QWidget_ITF /*777 QWidget **/) *QAbstractScrollArea {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(538689117, "_ZN19QAbstractScrollAreaC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QAbstractScrollAreaFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractScrollArea")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractScrollArea(QWidget *)

/*
 */
func (*QAbstractScrollArea) NewForInheritp() *QAbstractScrollArea {
	return NewQAbstractScrollAreap()
}
func NewQAbstractScrollAreap() *QAbstractScrollArea {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(538689117, "_ZN19QAbstractScrollAreaC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QAbstractScrollAreaFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractScrollArea")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:84
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QWidget * cornerWidget() const

/*
 */
func (this *QAbstractScrollArea) CornerWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.Qtcc3(1269609790, "_ZNK19QAbstractScrollArea12cornerWidgetEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QWidgetFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:85
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCornerWidget(QWidget *)

/*
 */
func (this *QAbstractScrollArea) SetCornerWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 Voidptr
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(322961687, "_ZN19QAbstractScrollArea15setCornerWidgetEP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:90
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QWidget * viewport() const

/*
 */
func (this *QAbstractScrollArea) Viewport() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.Qtcc3(3104358012, "_ZNK19QAbstractScrollArea8viewportEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QWidgetFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:91
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setViewport(QWidget *)

/*
 */
func (this *QAbstractScrollArea) SetViewport(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 Voidptr
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2564963984, "_ZN19QAbstractScrollArea11setViewportEP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQAbstractScrollArea(this *QAbstractScrollArea) {
	rv, err := qtrt.Qtcc1(0, "_ZN19QAbstractScrollAreaD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10182() {
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
