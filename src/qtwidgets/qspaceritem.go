//  header block begin
// /usr/include/qt/QtWidgets/qlayoutitem.h
// #include <qlayoutitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QSpacerItem struct {
	*QLayoutItem
}

func (this *QSpacerItem) GetCthis() unsafe.Pointer {
	return this.QLayoutItem.GetCthis()
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:95
// index:0
// inline
// void QSpacerItem(int, int, class QSizePolicy::Policy, class QSizePolicy::Policy)
func NewQSpacerItem(w int, h int, hData int, vData int) *QSpacerItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSpacerItemC2EiiN11QSizePolicy6PolicyES1_", ffiqt.FFI_TYPE_VOID, cthis, &w, &h, &hData, &vData)
	gopp.ErrPrint(err, rv)
	gothis := NewQSpacerItemFromPointer(cthis)
	return gothis
}
func NewQSpacerItemFromPointer(cthis unsafe.Pointer) *QSpacerItem {
	bcthis0 := NewQLayoutItemFromPointer(cthis)
	return &QSpacerItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:99
// index:0
// virtual
// void ~QSpacerItem()
func DeleteQSpacerItem(*QSpacerItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSpacerItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:101
// index:0
// void changeSize(int, int, class QSizePolicy::Policy, class QSizePolicy::Policy)
func (this *QSpacerItem) ChangeSize(w int, h int, hData int, vData int) {
	// 0: (, w int, h int, hData QSizePolicy::Policy, vData QSizePolicy::Policy), (&w, &h, &hData, &vData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSpacerItem10changeSizeEiiN11QSizePolicy6PolicyES1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &h, &hData, &vData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:104
// index:0
// virtual
// QSize sizeHint()
func (this *QSpacerItem) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:105
// index:0
// virtual
// QSize minimumSize()
func (this *QSpacerItem) MinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem11minimumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:106
// index:0
// virtual
// QSize maximumSize()
func (this *QSpacerItem) MaximumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem11maximumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:107
// index:0
// virtual
// Qt::Orientations expandingDirections()
func (this *QSpacerItem) ExpandingDirections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem19expandingDirectionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:108
// index:0
// virtual
// bool isEmpty()
func (this *QSpacerItem) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:109
// index:0
// virtual
// void setGeometry(const class QRect &)
func (this *QSpacerItem) SetGeometry(arg0 unsafe.Pointer) {
	// 0: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSpacerItem11setGeometryERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:110
// index:0
// virtual
// QRect geometry()
func (this *QSpacerItem) Geometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem8geometryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:111
// index:0
// virtual
// QSpacerItem * spacerItem()
func (this *QSpacerItem) SpacerItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSpacerItem10spacerItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:112
// index:0
// inline
// QSizePolicy sizePolicy()
func (this *QSpacerItem) SizePolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem10sizePolicyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
