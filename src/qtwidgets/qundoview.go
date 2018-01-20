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

// /usr/include/qt/QtWidgets/qundoview.h:59
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QUndoView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:65
// index:0
// void QUndoView(class QWidget *)
func NewQUndoView(parent unsafe.Pointer) *QUndoView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(cthis)
	return gothis
}
func NewQUndoViewFromPointer(cthis unsafe.Pointer) *QUndoView {
	bcthis0 := NewQListViewFromPointer(cthis)
	return &QUndoView{bcthis0}
}

// /usr/include/qt/QtWidgets/qundoview.h:66
// index:1
// void QUndoView(class QUndoStack *, class QWidget *)
func NewQUndoView_1(stack unsafe.Pointer, parent unsafe.Pointer) *QUndoView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoStackP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, stack, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:68
// index:2
// void QUndoView(class QUndoGroup *, class QWidget *)
func NewQUndoView_2(group unsafe.Pointer, parent unsafe.Pointer) *QUndoView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewC2EP10QUndoGroupP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, group, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qundoview.h:70
// index:0
// virtual
// void ~QUndoView()
func DeleteQUndoView(*QUndoView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:72
// index:0
// QUndoStack * stack()
func (this *QUndoView) Stack() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView5stackEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:74
// index:0
// QUndoGroup * group()
func (this *QUndoView) Group() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView5groupEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:77
// index:0
// void setEmptyLabel(const class QString &)
func (this *QUndoView) SetEmptyLabel(label unsafe.Pointer) {
	// 0: (, label const QString &), (label)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView13setEmptyLabelERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:78
// index:0
// QString emptyLabel()
func (this *QUndoView) EmptyLabel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView10emptyLabelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:80
// index:0
// void setCleanIcon(const class QIcon &)
func (this *QUndoView) SetCleanIcon(icon unsafe.Pointer) {
	// 0: (, icon const QIcon &), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView12setCleanIconERK5QIcon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:81
// index:0
// QIcon cleanIcon()
func (this *QUndoView) CleanIcon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUndoView9cleanIconEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:84
// index:0
// void setStack(class QUndoStack *)
func (this *QUndoView) SetStack(stack unsafe.Pointer) {
	// 0: (, stack QUndoStack *), (stack)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView8setStackEP10QUndoStack", ffiqt.FFI_TYPE_VOID, this.GetCthis(), stack)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundoview.h:86
// index:0
// void setGroup(class QUndoGroup *)
func (this *QUndoView) SetGroup(group unsafe.Pointer) {
	// 0: (, group QUndoGroup *), (group)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUndoView8setGroupEP10QUndoGroup", ffiqt.FFI_TYPE_VOID, this.GetCthis(), group)
	gopp.ErrPrint(err, rv)
}

//  body block end
