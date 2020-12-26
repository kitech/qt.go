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
// extern C begin: 3
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
// ignore GetCthis for 1 base
func QTreeViewFromptr(cthis Voidptr) *QTreeView {
	bcthis0 := QAbstractItemViewFromptr(cthis)
	return &QTreeView{bcthis0}
}
func (*QTreeView) Fromptr(cthis Voidptr) *QTreeView {
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
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(2330190834, "_ZN9QTreeViewC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
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
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(2330190834, "_ZN9QTreeViewC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QTreeViewFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QTreeView")
	return gothis
}

// /usr/include/qt/QtWidgets/qtreeview.h:74
// index:0
// Public virtual Ignore Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)

/*
 */
func (this *QTreeView) SetModel(model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 Voidptr
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2108085925, "_ZN9QTreeView8setModelEP18QAbstractItemModel", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:155
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void hideColumn(int)

/*
 */
func (this *QTreeView) HideColumn(column int) {
	rv, err := qtrt.Qtcc3(3246213577, "_ZN9QTreeView10hideColumnEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&column))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:156
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void showColumn(int)

/*
 */
func (this *QTreeView) ShowColumn(column int) {
	rv, err := qtrt.Qtcc3(128929531, "_ZN9QTreeView10showColumnEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&column))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:157
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void expand(const QModelIndex &)

/*
 */
func (this *QTreeView) Expand(index qtcore.QModelIndex_ITF) {
	var convArg0 Voidptr
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2194613163, "_ZN9QTreeView6expandERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:158
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void collapse(const QModelIndex &)

/*
 */
func (this *QTreeView) Collapse(index qtcore.QModelIndex_ITF) {
	var convArg0 Voidptr
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1002132025, "_ZN9QTreeView8collapseERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:159
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void resizeColumnToContents(int)

/*
 */
func (this *QTreeView) ResizeColumnToContents(column int) {
	rv, err := qtrt.Qtcc3(3203152354, "_ZN9QTreeView22resizeColumnToContentsEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&column))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQTreeView(this *QTreeView) {
	rv, err := qtrt.Qtcc1(0, "_ZN9QTreeViewD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10249() {
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
