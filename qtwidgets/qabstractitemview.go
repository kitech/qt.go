// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractitemview.h
// #include <qabstractitemview.h>
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

/*
 */
// size 48
type QAbstractItemView struct {
	*QAbstractScrollArea
}
type QAbstractItemView_ITF interface {
	QAbstractScrollArea_ITF
	QAbstractItemView_PTR() *QAbstractItemView
}

func (ptr *QAbstractItemView) QAbstractItemView_PTR() *QAbstractItemView { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QAbstractItemViewFromptr(cthis unsafe.Pointer) *QAbstractItemView {
	bcthis0 := QAbstractScrollAreaFromptr(cthis)
	return &QAbstractItemView{bcthis0}
}
func (*QAbstractItemView) Fromptr(cthis unsafe.Pointer) *QAbstractItemView {
	return QAbstractItemViewFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemView(QWidget *)

/*
 */
func (*QAbstractItemView) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QAbstractItemView {
	return NewQAbstractItemView(parent)
}
func NewQAbstractItemView(parent QWidget_ITF /*777 QWidget **/) *QAbstractItemView {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2582094347, "_ZN17QAbstractItemViewC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QAbstractItemViewFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractItemView")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemView(QWidget *)

/*
 */
func (*QAbstractItemView) NewForInheritp() *QAbstractItemView {
	return NewQAbstractItemViewp()
}
func NewQAbstractItemViewp() *QAbstractItemView {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2582094347, "_ZN17QAbstractItemViewC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QAbstractItemViewFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractItemView")
	return gothis
}

func DeleteQAbstractItemView(this *QAbstractItemView) {
	rv, err := qtrt.Qtcc1(0, "_ZN17QAbstractItemViewD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QAbstractItemView__SelectionMode = int

//
const QAbstractItemView__NoSelection QAbstractItemView__SelectionMode = 0

//
const QAbstractItemView__SingleSelection QAbstractItemView__SelectionMode = 1

//
const QAbstractItemView__MultiSelection QAbstractItemView__SelectionMode = 2

//
const QAbstractItemView__ExtendedSelection QAbstractItemView__SelectionMode = 3

//
const QAbstractItemView__ContiguousSelection QAbstractItemView__SelectionMode = 4

func (this *QAbstractItemView) SelectionModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_SelectionModeItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.SelectionModeItemName(val)
}

/*


 */
type QAbstractItemView__SelectionBehavior = int

//
const QAbstractItemView__SelectItems QAbstractItemView__SelectionBehavior = 0

//
const QAbstractItemView__SelectRows QAbstractItemView__SelectionBehavior = 1

//
const QAbstractItemView__SelectColumns QAbstractItemView__SelectionBehavior = 2

func (this *QAbstractItemView) SelectionBehaviorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_SelectionBehaviorItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.SelectionBehaviorItemName(val)
}

/*


 */
type QAbstractItemView__ScrollHint = int

//
const QAbstractItemView__EnsureVisible QAbstractItemView__ScrollHint = 0

//
const QAbstractItemView__PositionAtTop QAbstractItemView__ScrollHint = 1

//
const QAbstractItemView__PositionAtBottom QAbstractItemView__ScrollHint = 2

//
const QAbstractItemView__PositionAtCenter QAbstractItemView__ScrollHint = 3

func (this *QAbstractItemView) ScrollHintItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_ScrollHintItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.ScrollHintItemName(val)
}

/*


 */
type QAbstractItemView__EditTrigger = int

//
const QAbstractItemView__NoEditTriggers QAbstractItemView__EditTrigger = 0

//
const QAbstractItemView__CurrentChanged QAbstractItemView__EditTrigger = 1

//
const QAbstractItemView__DoubleClicked QAbstractItemView__EditTrigger = 2

//
const QAbstractItemView__SelectedClicked QAbstractItemView__EditTrigger = 4

//
const QAbstractItemView__EditKeyPressed QAbstractItemView__EditTrigger = 8

//
const QAbstractItemView__AnyKeyPressed QAbstractItemView__EditTrigger = 16

//
const QAbstractItemView__AllEditTriggers QAbstractItemView__EditTrigger = 31

func (this *QAbstractItemView) EditTriggerItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_EditTriggerItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.EditTriggerItemName(val)
}

/*


 */
type QAbstractItemView__ScrollMode = int

//
const QAbstractItemView__ScrollPerItem QAbstractItemView__ScrollMode = 0

//
const QAbstractItemView__ScrollPerPixel QAbstractItemView__ScrollMode = 1

func (this *QAbstractItemView) ScrollModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_ScrollModeItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.ScrollModeItemName(val)
}

/*


 */
type QAbstractItemView__DragDropMode = int

//
const QAbstractItemView__NoDragDrop QAbstractItemView__DragDropMode = 0

//
const QAbstractItemView__DragOnly QAbstractItemView__DragDropMode = 1

//
const QAbstractItemView__DropOnly QAbstractItemView__DragDropMode = 2

//
const QAbstractItemView__DragDrop QAbstractItemView__DragDropMode = 3

//
const QAbstractItemView__InternalMove QAbstractItemView__DragDropMode = 4

func (this *QAbstractItemView) DragDropModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_DragDropModeItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.DragDropModeItemName(val)
}

/*


 */
type QAbstractItemView__CursorAction = int

//
const QAbstractItemView__MoveUp QAbstractItemView__CursorAction = 0

//
const QAbstractItemView__MoveDown QAbstractItemView__CursorAction = 1

//
const QAbstractItemView__MoveLeft QAbstractItemView__CursorAction = 2

//
const QAbstractItemView__MoveRight QAbstractItemView__CursorAction = 3

//
const QAbstractItemView__MoveHome QAbstractItemView__CursorAction = 4

//
const QAbstractItemView__MoveEnd QAbstractItemView__CursorAction = 5

//
const QAbstractItemView__MovePageUp QAbstractItemView__CursorAction = 6

//
const QAbstractItemView__MovePageDown QAbstractItemView__CursorAction = 7

//
const QAbstractItemView__MoveNext QAbstractItemView__CursorAction = 8

//
const QAbstractItemView__MovePrevious QAbstractItemView__CursorAction = 9

func (this *QAbstractItemView) CursorActionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_CursorActionItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.CursorActionItemName(val)
}

/*


 */
type QAbstractItemView__State = int

//
const QAbstractItemView__NoState QAbstractItemView__State = 0

//
const QAbstractItemView__DraggingState QAbstractItemView__State = 1

//
const QAbstractItemView__DragSelectingState QAbstractItemView__State = 2

//
const QAbstractItemView__EditingState QAbstractItemView__State = 3

//
const QAbstractItemView__ExpandingState QAbstractItemView__State = 4

//
const QAbstractItemView__CollapsingState QAbstractItemView__State = 5

//
const QAbstractItemView__AnimatingState QAbstractItemView__State = 6

func (this *QAbstractItemView) StateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_StateItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.StateItemName(val)
}

/*


 */
type QAbstractItemView__DropIndicatorPosition = int

//
const QAbstractItemView__OnItem QAbstractItemView__DropIndicatorPosition = 0

//
const QAbstractItemView__AboveItem QAbstractItemView__DropIndicatorPosition = 1

//
const QAbstractItemView__BelowItem QAbstractItemView__DropIndicatorPosition = 2

//
const QAbstractItemView__OnViewport QAbstractItemView__DropIndicatorPosition = 3

func (this *QAbstractItemView) DropIndicatorPositionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_DropIndicatorPositionItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.DropIndicatorPositionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10097() {
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
