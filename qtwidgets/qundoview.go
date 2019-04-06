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

/*

 */
type QUndoView struct {
	*QListView
}
type QUndoView_ITF interface {
	QListView_ITF
	QUndoView_PTR() *QUndoView
}

func (ptr *QUndoView) QUndoView_PTR() *QUndoView { return ptr }

func (this *QUndoView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QListView.GetCthis()
	}
}
func (this *QUndoView) SetCthis(cthis unsafe.Pointer) {
	this.QListView = NewQListViewFromPointer(cthis)
}
func NewQUndoViewFromPointer(cthis unsafe.Pointer) *QUndoView {
	bcthis0 := NewQListViewFromPointer(cthis)
	return &QUndoView{bcthis0}
}
func (*QUndoView) NewFromPointer(cthis unsafe.Pointer) *QUndoView {
	return NewQUndoViewFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qundoview.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QUndoView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUndoView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundoview.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QWidget *)

/*
Constructs a new view with parent parent.
*/
func (*QUndoView) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QUndoView {
	return NewQUndoView(parent)
}
func NewQUndoView(parent QWidget_ITF /*777 QWidget **/) *QUndoView {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoView")
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QWidget *)

/*
Constructs a new view with parent parent.
*/
func (*QUndoView) NewForInheritp() *QUndoView {
	return NewQUndoViewp()
}
func NewQUndoViewp() *QUndoView {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoView")
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QUndoStack *, QWidget *)

/*
Constructs a new view with parent parent.
*/
func (*QUndoView) NewForInherit1(stack QUndoStack_ITF /*777 QUndoStack **/, parent QWidget_ITF /*777 QWidget **/) *QUndoView {
	return NewQUndoView1(stack, parent)
}
func NewQUndoView1(stack QUndoStack_ITF /*777 QUndoStack **/, parent QWidget_ITF /*777 QWidget **/) *QUndoView {
	var convArg0 unsafe.Pointer
	if stack != nil && stack.QUndoStack_PTR() != nil {
		convArg0 = stack.QUndoStack_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoStackP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoView")
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QUndoStack *, QWidget *)

/*
Constructs a new view with parent parent.
*/
func (*QUndoView) NewForInherit1p(stack QUndoStack_ITF /*777 QUndoStack **/) *QUndoView {
	return NewQUndoView1p(stack)
}
func NewQUndoView1p(stack QUndoStack_ITF /*777 QUndoStack **/) *QUndoView {
	var convArg0 unsafe.Pointer
	if stack != nil && stack.QUndoStack_PTR() != nil {
		convArg0 = stack.QUndoStack_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoStackP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoView")
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QUndoView()

/*

 */
func DeleteQUndoView(this *QUndoView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qundoview.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QUndoStack * stack() const

/*
Returns the stack currently displayed by this view. If the view is looking at a QUndoGroup, this the group's active stack.

See also setStack() and setGroup().
*/
func (this *QUndoView) Stack() *QUndoStack /*777 QUndoStack **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUndoView5stackEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQUndoStackFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundoview.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEmptyLabel(const QString &)

/*

 */
func (this *QUndoView) SetEmptyLabel(label string) {
	var tmpArg0 = qtcore.NewQString5(label)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoView13setEmptyLabelERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString emptyLabel() const

/*

 */
func (this *QUndoView) EmptyLabel() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUndoView10emptyLabelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qundoview.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCleanIcon(const QIcon &)

/*

 */
func (this *QUndoView) SetCleanIcon(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoView12setCleanIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon cleanIcon() const

/*

 */
func (this *QUndoView) CleanIcon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUndoView9cleanIconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qundoview.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStack(QUndoStack *)

/*
Sets the stack displayed by this view to stack. If stack is 0, the view will be empty.

If the view was previously looking at a QUndoGroup, the group is set to 0.

See also stack() and setGroup().
*/
func (this *QUndoView) SetStack(stack QUndoStack_ITF /*777 QUndoStack **/) {
	var convArg0 unsafe.Pointer
	if stack != nil && stack.QUndoStack_PTR() != nil {
		convArg0 = stack.QUndoStack_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoView8setStackEP10QUndoStack", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11369() {
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
