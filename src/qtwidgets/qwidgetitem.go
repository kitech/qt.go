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
	*QLayoutItem
}

func (this *QWidgetItem) GetCthis() unsafe.Pointer {
	return this.QLayoutItem.GetCthis()
}
func NewQWidgetItemFromPointer(cthis unsafe.Pointer) *QWidgetItem {
	bcthis0 := NewQLayoutItemFromPointer(cthis)
	return &QWidgetItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:126
// index:0
// Public inline
// void QWidgetItem(class QWidget *)
func NewQWidgetItem(w unsafe.Pointer) *QWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItemC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, w)
	gopp.ErrPrint(err, rv)
	gothis := NewQWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:127
// index:0
// Public virtual
// void ~QWidgetItem()
func DeleteQWidgetItem(*QWidgetItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:129
// index:0
// Public virtual
// QSize sizeHint()
func (this *QWidgetItem) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:130
// index:0
// Public virtual
// QSize minimumSize()
func (this *QWidgetItem) MinimumSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:131
// index:0
// Public virtual
// QSize maximumSize()
func (this *QWidgetItem) MaximumSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem11maximumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:132
// index:0
// Public virtual
// Qt::Orientations expandingDirections()
func (this *QWidgetItem) ExpandingDirections() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem19expandingDirectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:133
// index:0
// Public virtual
// bool isEmpty()
func (this *QWidgetItem) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:134
// index:0
// Public virtual
// void setGeometry(const class QRect &)
func (this *QWidgetItem) SetGeometry(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItem11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:135
// index:0
// Public virtual
// QRect geometry()
func (this *QWidgetItem) Geometry() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem8geometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:136
// index:0
// Public virtual
// QWidget * widget()
func (this *QWidgetItem) Widget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItem6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:138
// index:0
// Public virtual
// bool hasHeightForWidth()
func (this *QWidgetItem) HasHeightForWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:139
// index:0
// Public virtual
// int heightForWidth(int)
func (this *QWidgetItem) HeightForWidth(arg0 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:140
// index:0
// Public virtual
// QSizePolicy::ControlTypes controlTypes()
func (this *QWidgetItem) ControlTypes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem12controlTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
