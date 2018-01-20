//  header block begin
// /usr/include/qt/QtWidgets/qlayoutitem.h
// #include <qlayoutitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 59
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
type QLayoutItem struct {
	*qtrt.CObject
}

func (this *QLayoutItem) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:63
// index:0
// inline
// void QLayoutItem(Qt::Alignment)
func NewQLayoutItem(alignment int) *QLayoutItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItemC1E6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, cthis, &alignment)
	gopp.ErrPrint(err, rv)
	gothis := NewQLayoutItemFromPointer(cthis)
	return gothis
}
func NewQLayoutItemFromPointer(cthis unsafe.Pointer) *QLayoutItem {
	return &QLayoutItem{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:64
// index:0
// virtual
// void ~QLayoutItem()
func DeleteQLayoutItem(*QLayoutItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:65
// index:0
// pure virtual
// QSize sizeHint()
func (this *QLayoutItem) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:66
// index:0
// pure virtual
// QSize minimumSize()
func (this *QLayoutItem) MinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem11minimumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:67
// index:0
// pure virtual
// QSize maximumSize()
func (this *QLayoutItem) MaximumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem11maximumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:68
// index:0
// pure virtual
// Qt::Orientations expandingDirections()
func (this *QLayoutItem) ExpandingDirections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem19expandingDirectionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:69
// index:0
// pure virtual
// void setGeometry(const class QRect &)
func (this *QLayoutItem) SetGeometry(arg0 unsafe.Pointer) {
	// 0: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItem11setGeometryERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:70
// index:0
// pure virtual
// QRect geometry()
func (this *QLayoutItem) Geometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem8geometryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:71
// index:0
// pure virtual
// bool isEmpty()
func (this *QLayoutItem) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:72
// index:0
// virtual
// bool hasHeightForWidth()
func (this *QLayoutItem) HasHeightForWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem17hasHeightForWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:73
// index:0
// virtual
// int heightForWidth(int)
func (this *QLayoutItem) HeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:74
// index:0
// virtual
// int minimumHeightForWidth(int)
func (this *QLayoutItem) MinimumHeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem21minimumHeightForWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:75
// index:0
// virtual
// void invalidate()
func (this *QLayoutItem) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItem10invalidateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:77
// index:0
// virtual
// QWidget * widget()
func (this *QLayoutItem) Widget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItem6widgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:78
// index:0
// virtual
// QLayout * layout()
func (this *QLayoutItem) Layout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItem6layoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:79
// index:0
// virtual
// QSpacerItem * spacerItem()
func (this *QLayoutItem) SpacerItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItem10spacerItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:81
// index:0
// inline
// Qt::Alignment alignment()
func (this *QLayoutItem) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem9alignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:82
// index:0
// void setAlignment(Qt::Alignment)
func (this *QLayoutItem) SetAlignment(a int) {
	// 0: (, QFlags<Qt::AlignmentFlag> a), (&a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLayoutItem12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:83
// index:0
// virtual
// QSizePolicy::ControlTypes controlTypes()
func (this *QLayoutItem) ControlTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLayoutItem12controlTypesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
