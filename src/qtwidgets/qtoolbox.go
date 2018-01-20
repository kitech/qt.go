//  header block begin
// /usr/include/qt/QtWidgets/qtoolbox.h
// #include <qtoolbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
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
	*QFrame
}

func (this *QToolBox) GetCthis() unsafe.Pointer {
	return this.QFrame.GetCthis()
}

// /usr/include/qt/QtWidgets/qtoolbox.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QToolBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:60
// index:0
// void QToolBox(class QWidget *, Qt::WindowFlags)
func NewQToolBox(parent unsafe.Pointer, f int) *QToolBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBoxC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, parent, &f)
	gopp.ErrPrint(err, rv)
	gothis := NewQToolBoxFromPointer(cthis)
	return gothis
}
func NewQToolBoxFromPointer(cthis unsafe.Pointer) *QToolBox {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QToolBox{bcthis0}
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
	// 0: (, widget QWidget *, text const QString &), (widget, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox7addItemEP7QWidgetRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:64
// index:1
// int addItem(class QWidget *, const class QIcon &, const class QString &)
func (this *QToolBox) AddItem_1(widget unsafe.Pointer, icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, widget QWidget *, icon const QIcon &, text const QString &), (widget, icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox7addItemEP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget, icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:65
// index:0
// int insertItem(int, class QWidget *, const class QString &)
func (this *QToolBox) InsertItem(index int, widget unsafe.Pointer, text unsafe.Pointer) {
	// 0: (, index int, widget QWidget *, text const QString &), (&index, widget, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10insertItemEiP7QWidgetRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, widget, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:66
// index:1
// int insertItem(int, class QWidget *, const class QIcon &, const class QString &)
func (this *QToolBox) InsertItem_1(index int, widget unsafe.Pointer, icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, index int, widget QWidget *, icon const QIcon &, text const QString &), (&index, widget, icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10insertItemEiP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, widget, icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:68
// index:0
// void removeItem(int)
func (this *QToolBox) RemoveItem(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10removeItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:70
// index:0
// void setItemEnabled(int, _Bool)
func (this *QToolBox) SetItemEnabled(index int, enabled bool) {
	// 0: (, index int, enabled bool), (&index, &enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14setItemEnabledEib", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:71
// index:0
// bool isItemEnabled(int)
func (this *QToolBox) IsItemEnabled(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox13isItemEnabledEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:73
// index:0
// void setItemText(int, const class QString &)
func (this *QToolBox) SetItemText(index int, text unsafe.Pointer) {
	// 0: (, index int, text const QString &), (&index, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11setItemTextEiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:74
// index:0
// QString itemText(int)
func (this *QToolBox) ItemText(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox8itemTextEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:76
// index:0
// void setItemIcon(int, const class QIcon &)
func (this *QToolBox) SetItemIcon(index int, icon unsafe.Pointer) {
	// 0: (, index int, icon const QIcon &), (&index, icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11setItemIconEiRK5QIcon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:77
// index:0
// QIcon itemIcon(int)
func (this *QToolBox) ItemIcon(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox8itemIconEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:80
// index:0
// void setItemToolTip(int, const class QString &)
func (this *QToolBox) SetItemToolTip(index int, toolTip unsafe.Pointer) {
	// 0: (, index int, toolTip const QString &), (&index, toolTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14setItemToolTipEiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, toolTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:81
// index:0
// QString itemToolTip(int)
func (this *QToolBox) ItemToolTip(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox11itemToolTipEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:84
// index:0
// int currentIndex()
func (this *QToolBox) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:85
// index:0
// QWidget * currentWidget()
func (this *QToolBox) CurrentWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox13currentWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:86
// index:0
// QWidget * widget(int)
func (this *QToolBox) Widget(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox6widgetEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:87
// index:0
// int indexOf(class QWidget *)
func (this *QToolBox) IndexOf(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox7indexOfEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:88
// index:0
// int count()
func (this *QToolBox) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:91
// index:0
// void setCurrentIndex(int)
func (this *QToolBox) SetCurrentIndex(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox15setCurrentIndexEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:92
// index:0
// void setCurrentWidget(class QWidget *)
func (this *QToolBox) SetCurrentWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:95
// index:0
// void currentChanged(int)
func (this *QToolBox) CurrentChanged(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14currentChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:98
// index:0
// virtual
// bool event(class QEvent *)
func (this *QToolBox) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:99
// index:0
// virtual
// void itemInserted(int)
func (this *QToolBox) ItemInserted(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox12itemInsertedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:100
// index:0
// virtual
// void itemRemoved(int)
func (this *QToolBox) ItemRemoved(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11itemRemovedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:101
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QToolBox) ShowEvent(e unsafe.Pointer) {
	// 0: (, e QShowEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:102
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QToolBox) ChangeEvent(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
