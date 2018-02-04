package qtwidgets

// /usr/include/qt/QtWidgets/qundoview.h
// #include <qundoview.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin

type QUndoView struct {
	*QListView
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qundoview.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QWidget *)
func NewQUndoView(parent *QWidget /*777 QWidget **/) *QUndoView {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QUndoStack *, QWidget *)
func NewQUndoView_1(stack *QUndoStack /*777 QUndoStack **/, parent *QWidget /*777 QWidget **/) *QUndoView {
	var convArg0 = stack.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoStackP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QUndoView(QUndoGroup *, QWidget *)
func NewQUndoView_2(group *QUndoGroup /*777 QUndoGroup **/, parent *QWidget /*777 QWidget **/) *QUndoView {
	var convArg0 = group.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoGroupP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QUndoView()
func DeleteQUndoView(this *QUndoView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qundoview.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QUndoStack * stack()
func (this *QUndoView) Stack() *QUndoStack /*777 QUndoStack **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView5stackEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQUndoStackFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qundoview.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QUndoGroup * group()
func (this *QUndoView) Group() *QUndoGroup /*777 QUndoGroup **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView5groupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQUndoGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qundoview.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEmptyLabel(const QString &)
func (this *QUndoView) SetEmptyLabel(label *qtcore.QString) {
	var convArg0 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView13setEmptyLabelERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString emptyLabel()
func (this *QUndoView) EmptyLabel() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView10emptyLabelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qundoview.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCleanIcon(const QIcon &)
func (this *QUndoView) SetCleanIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView12setCleanIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon cleanIcon()
func (this *QUndoView) CleanIcon() *qtgui.QIcon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView9cleanIconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qundoview.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStack(QUndoStack *)
func (this *QUndoView) SetStack(stack *QUndoStack /*777 QUndoStack **/) {
	var convArg0 = stack.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView8setStackEP10QUndoStack", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGroup(QUndoGroup *)
func (this *QUndoView) SetGroup(group *QUndoGroup /*777 QUndoGroup **/) {
	var convArg0 = group.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView8setGroupEP10QUndoGroup", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
