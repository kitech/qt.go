//  header block begin
// /usr/include/qt/QtWidgets/qundoview.h
// #include <qundoview.h>
// #include <QtWidgets>
package qtwidgets

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	return this.QListView.GetCthis()
}
func NewQUndoViewFromPointer(cthis unsafe.Pointer) *QUndoView {
	bcthis0 := NewQListViewFromPointer(cthis)
	return &QUndoView{bcthis0}
}

// /usr/include/qt/QtWidgets/qundoview.h:59
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QUndoView) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundoview.h:65
// index:0
// Public
// void QUndoView(class QWidget *)
func NewQUndoView(parent unsafe.Pointer) *QUndoView {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:66
// index:1
// Public
// void QUndoView(class QUndoStack *, class QWidget *)
func NewQUndoView_1(stack unsafe.Pointer, parent unsafe.Pointer) *QUndoView {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoStackP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, stack, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:68
// index:2
// Public
// void QUndoView(class QUndoGroup *, class QWidget *)
func NewQUndoView_2(group unsafe.Pointer, parent unsafe.Pointer) *QUndoView {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoGroupP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, group, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:70
// index:0
// Public virtual
// void ~QUndoView()
func DeleteQUndoView(*QUndoView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:72
// index:0
// Public
// QUndoStack * stack()
func (this *QUndoView) Stack() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView5stackEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundoview.h:74
// index:0
// Public
// QUndoGroup * group()
func (this *QUndoView) Group() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView5groupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundoview.h:77
// index:0
// Public
// void setEmptyLabel(const class QString &)
func (this *QUndoView) SetEmptyLabel(label *qtcore.QString) {
	var convArg0 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView13setEmptyLabelERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:78
// index:0
// Public
// QString emptyLabel()
func (this *QUndoView) EmptyLabel() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView10emptyLabelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundoview.h:80
// index:0
// Public
// void setCleanIcon(const class QIcon &)
func (this *QUndoView) SetCleanIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView12setCleanIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:81
// index:0
// Public
// QIcon cleanIcon()
func (this *QUndoView) CleanIcon() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView9cleanIconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundoview.h:84
// index:0
// Public
// void setStack(class QUndoStack *)
func (this *QUndoView) SetStack(stack unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView8setStackEP10QUndoStack", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stack)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:86
// index:0
// Public
// void setGroup(class QUndoGroup *)
func (this *QUndoView) SetGroup(group unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView8setGroupEP10QUndoGroup", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), group)
	gopp.ErrPrint(err, rv)
}

//  body block end
