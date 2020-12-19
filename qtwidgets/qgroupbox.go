// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qgroupbox.h
// #include <qgroupbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
type QGroupBox struct {
	*QWidget
}
type QGroupBox_ITF interface {
	QWidget_ITF
	QGroupBox_PTR() *QGroupBox
}

func (ptr *QGroupBox) QGroupBox_PTR() *QGroupBox { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QGroupBoxFromptr(cthis unsafe.Pointer) *QGroupBox {
	bcthis0 := QWidgetFromptr(cthis)
	return &QGroupBox{bcthis0}
}
func (*QGroupBox) Fromptr(cthis unsafe.Pointer) *QGroupBox {
	return QGroupBoxFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(QWidget *)

/*
 */
func (*QGroupBox) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QGroupBox {
	return NewQGroupBox(parent)
}
func NewQGroupBox(parent QWidget_ITF /*777 QWidget **/) *QGroupBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(662756536, "_ZN9QGroupBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QGroupBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(QWidget *)

/*
 */
func (*QGroupBox) NewForInheritp() *QGroupBox {
	return NewQGroupBoxp()
}
func NewQGroupBoxp() *QGroupBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(662756536, "_ZN9QGroupBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QGroupBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(const QString &, QWidget *)

/*
 */
func (*QGroupBox) NewForInherit1(title string, parent QWidget_ITF /*777 QWidget **/) *QGroupBox {
	return NewQGroupBox1(title, parent)
}
func NewQGroupBox1(title string, parent QWidget_ITF /*777 QWidget **/) *QGroupBox {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2800323070, "_ZN9QGroupBoxC2ERK7QStringP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QGroupBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(const QString &, QWidget *)

/*
 */
func (*QGroupBox) NewForInherit1p(title string) *QGroupBox {
	return NewQGroupBox1p(title)
}
func NewQGroupBox1p(title string) *QGroupBox {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2800323070, "_ZN9QGroupBoxC2ERK7QStringP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QGroupBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

func DeleteQGroupBox(this *QGroupBox) {
	rv, err := qtrt.Qtcc1(0, "_ZN9QGroupBoxD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10127() {
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
