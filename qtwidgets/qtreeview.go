// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qtreeview.h
// #include <qtreeview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QTreeView struct {
	*QAbstractItemView
}
type QTreeView_ITF interface {
	QAbstractItemView_ITF
	QTreeView_PTR() *QTreeView
}

func (ptr *QTreeView) QTreeView_PTR() *QTreeView { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QTreeViewFromptr(cthis unsafe.Pointer) *QTreeView {
	bcthis0 := QAbstractItemViewFromptr(cthis)
	return &QTreeView{bcthis0}
}
func (*QTreeView) Fromptr(cthis unsafe.Pointer) *QTreeView {
	return QTreeViewFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qtreeview.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeView(QWidget *)

/*
 */
func (*QTreeView) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QTreeView {
	return NewQTreeView(parent)
}
func NewQTreeView(parent QWidget_ITF /*777 QWidget **/) *QTreeView {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2330190834, "_ZN9QTreeViewC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QTreeViewFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QTreeView")
	return gothis
}

// /usr/include/qt/QtWidgets/qtreeview.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeView(QWidget *)

/*
 */
func (*QTreeView) NewForInheritp() *QTreeView {
	return NewQTreeViewp()
}
func NewQTreeViewp() *QTreeView {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2330190834, "_ZN9QTreeViewC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QTreeViewFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QTreeView")
	return gothis
}

func DeleteQTreeView(this *QTreeView) {
	rv, err := qtrt.Qtcc1(0, "_ZN9QTreeViewD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10159() {
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
