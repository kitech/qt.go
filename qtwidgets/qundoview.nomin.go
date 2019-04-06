//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qundoview.h
// #include <qundoview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
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

// /usr/include/qt/QtWidgets/qundoview.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QUndoGroup *, QWidget *)

/*
Constructs a new view with parent parent.
*/
func (*QUndoView) NewForInherit2(group QUndoGroup_ITF /*777 QUndoGroup **/, parent QWidget_ITF /*777 QWidget **/) *QUndoView {
	return NewQUndoView2(group, parent)
}
func NewQUndoView2(group QUndoGroup_ITF /*777 QUndoGroup **/, parent QWidget_ITF /*777 QWidget **/) *QUndoView {
	var convArg0 unsafe.Pointer
	if group != nil && group.QUndoGroup_PTR() != nil {
		convArg0 = group.QUndoGroup_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoGroupP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoView")
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QUndoGroup *, QWidget *)

/*
Constructs a new view with parent parent.
*/
func (*QUndoView) NewForInherit2p(group QUndoGroup_ITF /*777 QUndoGroup **/) *QUndoView {
	return NewQUndoView2p(group)
}
func NewQUndoView2p(group QUndoGroup_ITF /*777 QUndoGroup **/) *QUndoView {
	var convArg0 unsafe.Pointer
	if group != nil && group.QUndoGroup_PTR() != nil {
		convArg0 = group.QUndoGroup_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoGroupP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoView")
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QUndoGroup * group() const

/*
Returns the group displayed by this view.

If the view is not looking at group, this function returns 0.

See also setGroup() and setStack().
*/
func (this *QUndoView) Group() *QUndoGroup /*777 QUndoGroup **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUndoView5groupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQUndoGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundoview.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGroup(QUndoGroup *)

/*
Sets the group displayed by this view to group. If group is 0, the view will be empty.

The view will update itself autmiatically whenever the active stack of the group changes.

See also group() and setStack().
*/
func (this *QUndoView) SetGroup(group QUndoGroup_ITF /*777 QUndoGroup **/) {
	var convArg0 unsafe.Pointer
	if group != nil && group.QUndoGroup_PTR() != nil {
		convArg0 = group.QUndoGroup_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoView8setGroupEP10QUndoGroup", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11370() {
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
