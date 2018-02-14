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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

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
// [8] const QMetaObject * metaObject()
func (this *QUndoView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUndoView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundoview.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QWidget *)
func NewQUndoView(parent QWidget_ITF /*777 QWidget **/) *QUndoView {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QUndoStack *, QWidget *)
func NewQUndoView_1(stack QUndoStack_ITF /*777 QUndoStack **/, parent QWidget_ITF /*777 QWidget **/) *QUndoView {
	var convArg0 = stack.QUndoStack_PTR().GetCthis()
	var convArg1 = parent.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoStackP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QUndoGroup *, QWidget *)
func NewQUndoView_2(group QUndoGroup_ITF /*777 QUndoGroup **/, parent QWidget_ITF /*777 QWidget **/) *QUndoView {
	var convArg0 = group.QUndoGroup_PTR().GetCthis()
	var convArg1 = parent.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoGroupP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QUndoView()
func DeleteQUndoView(this *QUndoView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qundoview.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QUndoStack * stack()
func (this *QUndoView) Stack() *QUndoStack /*777 QUndoStack **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUndoView5stackEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQUndoStackFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundoview.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QUndoGroup * group()
func (this *QUndoView) Group() *QUndoGroup /*777 QUndoGroup **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUndoView5groupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQUndoGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundoview.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEmptyLabel(const QString &)
func (this *QUndoView) SetEmptyLabel(label string) {
	var tmpArg0 = qtcore.NewQString_5(label)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoView13setEmptyLabelERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString emptyLabel()
func (this *QUndoView) EmptyLabel() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUndoView10emptyLabelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qundoview.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCleanIcon(const QIcon &)
func (this *QUndoView) SetCleanIcon(icon qtgui.QIcon_ITF) {
	var convArg0 = icon.QIcon_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoView12setCleanIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon cleanIcon()
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
func (this *QUndoView) SetStack(stack QUndoStack_ITF /*777 QUndoStack **/) {
	var convArg0 = stack.QUndoStack_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoView8setStackEP10QUndoStack", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGroup(QUndoGroup *)
func (this *QUndoView) SetGroup(group QUndoGroup_ITF /*777 QUndoGroup **/) {
	var convArg0 = group.QUndoGroup_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUndoView8setGroupEP10QUndoGroup", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init() {
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
