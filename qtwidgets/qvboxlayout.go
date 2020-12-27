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
// extern C begin: 2
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
type QVBoxLayout struct {
	*QBoxLayout
}
type QVBoxLayout_ITF interface {
	QBoxLayout_ITF
	QVBoxLayout_PTR() *QVBoxLayout
}

func (ptr *QVBoxLayout) QVBoxLayout_PTR() *QVBoxLayout { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QVBoxLayoutFromptr(cthis Voidptr) *QVBoxLayout {
	bcthis0 := QBoxLayoutFromptr(cthis)
	return &QVBoxLayout{bcthis0}
}
func (*QVBoxLayout) Fromptr(cthis Voidptr) *QVBoxLayout {
	return QVBoxLayoutFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QVBoxLayout()

/*
 */
func (*QVBoxLayout) NewForInherit() *QVBoxLayout {
	return NewQVBoxLayout()
}
func NewQVBoxLayout() *QVBoxLayout {
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc3(538351151, "_ZN11QVBoxLayoutC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint3(err, rv)
	gothis := QVBoxLayoutFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QVBoxLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:131
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QVBoxLayout(QWidget *)

/*
 */
func (*QVBoxLayout) NewForInherit1(parent QWidget_ITF /*777 QWidget **/) *QVBoxLayout {
	return NewQVBoxLayout1(parent)
}
func NewQVBoxLayout1(parent QWidget_ITF /*777 QWidget **/) *QVBoxLayout {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc3(368368412, "_ZN11QVBoxLayoutC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QVBoxLayoutFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QVBoxLayout")
	return gothis
}

func DeleteQVBoxLayout(this *QVBoxLayout) {
	rv, err := qtrt.Qtcc3(3183586966, "_ZN11QVBoxLayoutD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10211() {
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
