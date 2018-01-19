//  header block begin
// /usr/include/qt/QtWidgets/qtoolbox.h
// #include <qtoolbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 38
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
type QToolBox struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qtoolbox.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QToolBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:61
// index:0
// virtual
// void ~QToolBox()
func DeleteQToolBox(*QToolBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:63
// index:0
// int addItem(class QWidget *, const class QString &)
func (this *QToolBox) AddItem(widget unsafe.Pointer, text unsafe.Pointer) {
	// 0: (, QWidget * widget, const QString & text), (widget, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox7addItemEP7QWidgetRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, widget, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:64
// index:1
// int addItem(class QWidget *, const class QIcon &, const class QString &)
func (this *QToolBox) AddItem_1(widget unsafe.Pointer, icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, QWidget * widget, const QIcon & icon, const QString & text), (widget, icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox7addItemEP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, widget, icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:65
// index:0
// int insertItem(int, class QWidget *, const class QString &)
func (this *QToolBox) InsertItem(index int, widget unsafe.Pointer, text unsafe.Pointer) {
	// 0: (, int index, QWidget * widget, const QString & text), (&index, widget, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10insertItemEiP7QWidgetRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &index, widget, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:66
// index:1
// int insertItem(int, class QWidget *, const class QIcon &, const class QString &)
func (this *QToolBox) InsertItem_1(index int, widget unsafe.Pointer, icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, int index, QWidget * widget, const QIcon & icon, const QString & text), (&index, widget, icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10insertItemEiP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &index, widget, icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:68
// index:0
// void removeItem(int)
func (this *QToolBox) RemoveItem(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10removeItemEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:70
// index:0
// void setItemEnabled(int, _Bool)
func (this *QToolBox) SetItemEnabled(index int, enabled bool) {
	// 0: (, int index, bool enabled), (&index, &enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14setItemEnabledEib", ffiqt.FFI_TYPE_VOID, this.cthis, &index, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:71
// index:0
// bool isItemEnabled(int)
func (this *QToolBox) IsItemEnabled(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox13isItemEnabledEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:73
// index:0
// void setItemText(int, const class QString &)
func (this *QToolBox) SetItemText(index int, text unsafe.Pointer) {
	// 0: (, int index, const QString & text), (&index, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11setItemTextEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &index, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:74
// index:0
// QString itemText(int)
func (this *QToolBox) ItemText(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox8itemTextEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:76
// index:0
// void setItemIcon(int, const class QIcon &)
func (this *QToolBox) SetItemIcon(index int, icon unsafe.Pointer) {
	// 0: (, int index, const QIcon & icon), (&index, icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11setItemIconEiRK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, &index, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:77
// index:0
// QIcon itemIcon(int)
func (this *QToolBox) ItemIcon(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox8itemIconEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:80
// index:0
// void setItemToolTip(int, const class QString &)
func (this *QToolBox) SetItemToolTip(index int, toolTip unsafe.Pointer) {
	// 0: (, int index, const QString & toolTip), (&index, toolTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14setItemToolTipEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &index, toolTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:81
// index:0
// QString itemToolTip(int)
func (this *QToolBox) ItemToolTip(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox11itemToolTipEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:84
// index:0
// int currentIndex()
func (this *QToolBox) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:85
// index:0
// QWidget * currentWidget()
func (this *QToolBox) CurrentWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox13currentWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:86
// index:0
// QWidget * widget(int)
func (this *QToolBox) Widget(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox6widgetEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:87
// index:0
// int indexOf(class QWidget *)
func (this *QToolBox) IndexOf(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox7indexOfEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:88
// index:0
// int count()
func (this *QToolBox) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:91
// index:0
// void setCurrentIndex(int)
func (this *QToolBox) SetCurrentIndex(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox15setCurrentIndexEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:92
// index:0
// void setCurrentWidget(class QWidget *)
func (this *QToolBox) SetCurrentWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:95
// index:0
// void currentChanged(int)
func (this *QToolBox) CurrentChanged(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14currentChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

//  body block end
