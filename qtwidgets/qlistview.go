// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qlistview.h
// #include <qlistview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QListView struct {
	*QAbstractItemView
}
type QListView_ITF interface {
	QAbstractItemView_ITF
	QListView_PTR() *QListView
}

func (ptr *QListView) QListView_PTR() *QListView { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QListViewFromptr(cthis Voidptr) *QListView {
	bcthis0 := QAbstractItemViewFromptr(cthis)
	return &QListView{bcthis0}
}
func (*QListView) Fromptr(cthis Voidptr) *QListView {
	return QListViewFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qlistview.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListView(QWidget *)

/*
 */
func (*QListView) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QListView {
	return NewQListView(parent)
}
func NewQListView(parent QWidget_ITF /*777 QWidget **/) *QListView {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(4267936992, "_ZN9QListViewC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QListViewFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QListView")
	return gothis
}

// /usr/include/qt/QtWidgets/qlistview.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListView(QWidget *)

/*
 */
func (*QListView) NewForInheritp() *QListView {
	return NewQListViewp()
}
func NewQListViewp() *QListView {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(4267936992, "_ZN9QListViewC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QListViewFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QListView")
	return gothis
}

func DeleteQListView(this *QListView) {
	rv, err := qtrt.Qtcc1(0, "_ZN9QListViewD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QListView__Movement = int

//
const QListView__Static QListView__Movement = 0

//
const QListView__Free QListView__Movement = 1

//
const QListView__Snap QListView__Movement = 2

func (this *QListView) MovementItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QListView_MovementItemName(val int) string {
	var nilthis *QListView
	return nilthis.MovementItemName(val)
}

/*


 */
type QListView__Flow = int

//
const QListView__LeftToRight QListView__Flow = 0

//
const QListView__TopToBottom QListView__Flow = 1

func (this *QListView) FlowItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QListView_FlowItemName(val int) string {
	var nilthis *QListView
	return nilthis.FlowItemName(val)
}

/*


 */
type QListView__ResizeMode = int

//
const QListView__Fixed QListView__ResizeMode = 0

//
const QListView__Adjust QListView__ResizeMode = 1

func (this *QListView) ResizeModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QListView_ResizeModeItemName(val int) string {
	var nilthis *QListView
	return nilthis.ResizeModeItemName(val)
}

/*


 */
type QListView__LayoutMode = int

//
const QListView__SinglePass QListView__LayoutMode = 0

//
const QListView__Batched QListView__LayoutMode = 1

func (this *QListView) LayoutModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QListView_LayoutModeItemName(val int) string {
	var nilthis *QListView
	return nilthis.LayoutModeItemName(val)
}

/*


 */
type QListView__ViewMode = int

//
const QListView__ListMode QListView__ViewMode = 0

//
const QListView__IconMode QListView__ViewMode = 1

func (this *QListView) ViewModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QListView_ViewModeItemName(val int) string {
	var nilthis *QListView
	return nilthis.ViewModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10221() {
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
