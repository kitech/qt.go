//  header block begin
// /usr/include/qt/QtWidgets/qlayoutitem.h
// #include <qlayoutitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QWidgetItem struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:126
// index:0
// inline
// void QWidgetItem(class QWidget *)
func NewQWidgetItem(w unsafe.Pointer) *QWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItemC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, w)
	gopp.ErrPrint(err, rv)
	return &QWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:127
// index:0
// virtual
// void ~QWidgetItem()
func DeleteQWidgetItem(*QWidgetItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:129
// index:0
// virtual
// QSize sizeHint()
func (this *QWidgetItem) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:130
// index:0
// virtual
// QSize minimumSize()
func (this *QWidgetItem) MinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem11minimumSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:131
// index:0
// virtual
// QSize maximumSize()
func (this *QWidgetItem) MaximumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem11maximumSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:132
// index:0
// virtual
// Qt::Orientations expandingDirections()
func (this *QWidgetItem) ExpandingDirections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem19expandingDirectionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:133
// index:0
// virtual
// bool isEmpty()
func (this *QWidgetItem) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:134
// index:0
// virtual
// void setGeometry(const class QRect &)
func (this *QWidgetItem) SetGeometry(arg0 unsafe.Pointer) {
	// 0: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItem11setGeometryERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:135
// index:0
// virtual
// QRect geometry()
func (this *QWidgetItem) Geometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem8geometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:136
// index:0
// virtual
// QWidget * widget()
func (this *QWidgetItem) Widget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItem6widgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:138
// index:0
// virtual
// bool hasHeightForWidth()
func (this *QWidgetItem) HasHeightForWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem17hasHeightForWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:139
// index:0
// virtual
// int heightForWidth(int)
func (this *QWidgetItem) HeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:140
// index:0
// virtual
// QSizePolicy::ControlTypes controlTypes()
func (this *QWidgetItem) ControlTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem12controlTypesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
