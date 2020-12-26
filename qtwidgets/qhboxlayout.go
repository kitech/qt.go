package qtwidgets

// /usr/include/qt/QtWidgets/qboxlayout.h
// #include <qboxlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
// size 32
type QHBoxLayout struct {
	*QBoxLayout
}
type QHBoxLayout_ITF interface {
	QBoxLayout_ITF
	QHBoxLayout_PTR() *QHBoxLayout
}

func (ptr *QHBoxLayout) QHBoxLayout_PTR() *QHBoxLayout { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QHBoxLayoutFromptr(cthis Voidptr) *QHBoxLayout {
	bcthis0 := QBoxLayoutFromptr(cthis)
	return &QHBoxLayout{bcthis0}
}
func (*QHBoxLayout) Fromptr(cthis Voidptr) *QHBoxLayout {
	return QHBoxLayoutFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHBoxLayout()

/*
 */
func (*QHBoxLayout) NewForInherit() *QHBoxLayout {
	return NewQHBoxLayout()
}
func NewQHBoxLayout() *QHBoxLayout {
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc3(2369781699, "_ZN11QHBoxLayoutC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint2(err, rv)
	gothis := QHBoxLayoutFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QHBoxLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:118
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHBoxLayout(QWidget *)

/*
 */
func (*QHBoxLayout) NewForInherit1(parent QWidget_ITF /*777 QWidget **/) *QHBoxLayout {
	return NewQHBoxLayout1(parent)
}
func NewQHBoxLayout1(parent QWidget_ITF /*777 QWidget **/) *QHBoxLayout {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc3(2008149469, "_ZN11QHBoxLayoutC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QHBoxLayoutFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QHBoxLayout")
	return gothis
}

func DeleteQHBoxLayout(this *QHBoxLayout) {
	rv, err := qtrt.Qtcc3(283690874, "_ZN11QHBoxLayoutD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10201() {
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
