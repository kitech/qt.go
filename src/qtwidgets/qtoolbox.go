package qtwidgets

// /usr/include/qt/QtWidgets/qtoolbox.h
// #include <qtoolbox.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
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
	if this == nil {
		return nil
	} else {
		return this.QFrame.GetCthis()
	}
}
func (this *QToolBox) SetCthis(cthis unsafe.Pointer) {
	this.QFrame = NewQFrameFromPointer(cthis)
}
func NewQToolBoxFromPointer(cthis unsafe.Pointer) *QToolBox {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QToolBox{bcthis0}
}
func (*QToolBox) NewFromPointer(cthis unsafe.Pointer) *QToolBox {
	return NewQToolBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QToolBox) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbox.h:60
// index:0
// Public
// void QToolBox(class QWidget *, Qt::WindowFlags)
func NewQToolBox(parent *QWidget /*444 QWidget **/, f int) *QToolBox {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBoxC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQToolBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbox.h:61
// index:0
// Public virtual
// void ~QToolBox()
func DeleteQToolBox(*QToolBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:63
// index:0
// Public
// int addItem(class QWidget *, const class QString &)
func (this *QToolBox) AddItem(widget *QWidget /*444 QWidget **/, text *qtcore.QString) int {
	var convArg0 = widget.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox7addItemEP7QWidgetRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:64
// index:1
// Public
// int addItem(class QWidget *, const class QIcon &, const class QString &)
func (this *QToolBox) AddItem_1(widget *QWidget /*444 QWidget **/, icon *qtgui.QIcon, text *qtcore.QString) int {
	var convArg0 = widget.GetCthis()
	var convArg1 = icon.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox7addItemEP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:65
// index:0
// Public
// int insertItem(int, class QWidget *, const class QString &)
func (this *QToolBox) InsertItem(index int, widget *QWidget /*444 QWidget **/, text *qtcore.QString) int {
	var convArg1 = widget.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10insertItemEiP7QWidgetRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:66
// index:1
// Public
// int insertItem(int, class QWidget *, const class QIcon &, const class QString &)
func (this *QToolBox) InsertItem_1(index int, widget *QWidget /*444 QWidget **/, icon *qtgui.QIcon, text *qtcore.QString) int {
	var convArg1 = widget.GetCthis()
	var convArg2 = icon.GetCthis()
	var convArg3 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10insertItemEiP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:68
// index:0
// Public
// void removeItem(int)
func (this *QToolBox) RemoveItem(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10removeItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:70
// index:0
// Public
// void setItemEnabled(int, _Bool)
func (this *QToolBox) SetItemEnabled(index int, enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14setItemEnabledEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:71
// index:0
// Public
// bool isItemEnabled(int)
func (this *QToolBox) IsItemEnabled(index int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox13isItemEnabledEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbox.h:73
// index:0
// Public
// void setItemText(int, const class QString &)
func (this *QToolBox) SetItemText(index int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11setItemTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:74
// index:0
// Public
// QString itemText(int)
func (this *QToolBox) ItemText(index int) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox8itemTextEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbox.h:76
// index:0
// Public
// void setItemIcon(int, const class QIcon &)
func (this *QToolBox) SetItemIcon(index int, icon *qtgui.QIcon) {
	var convArg1 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11setItemIconEiRK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:77
// index:0
// Public
// QIcon itemIcon(int)
func (this *QToolBox) ItemIcon(index int) *qtgui.QIcon /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox8itemIconEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbox.h:80
// index:0
// Public
// void setItemToolTip(int, const class QString &)
func (this *QToolBox) SetItemToolTip(index int, toolTip *qtcore.QString) {
	var convArg1 = toolTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14setItemToolTipEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:81
// index:0
// Public
// QString itemToolTip(int)
func (this *QToolBox) ItemToolTip(index int) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox11itemToolTipEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbox.h:84
// index:0
// Public
// int currentIndex()
func (this *QToolBox) CurrentIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:85
// index:0
// Public
// QWidget * currentWidget()
func (this *QToolBox) CurrentWidget() *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox13currentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbox.h:86
// index:0
// Public
// QWidget * widget(int)
func (this *QToolBox) Widget(index int) *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox6widgetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbox.h:87
// index:0
// Public
// int indexOf(class QWidget *)
func (this *QToolBox) IndexOf(widget *QWidget /*444 QWidget **/) int {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox7indexOfEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:88
// index:0
// Public
// int count()
func (this *QToolBox) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:91
// index:0
// Public
// void setCurrentIndex(int)
func (this *QToolBox) SetCurrentIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox15setCurrentIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:92
// index:0
// Public
// void setCurrentWidget(class QWidget *)
func (this *QToolBox) SetCurrentWidget(widget *QWidget /*444 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:95
// index:0
// Public
// void currentChanged(int)
func (this *QToolBox) CurrentChanged(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14currentChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:98
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QToolBox) Event(e *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbox.h:99
// index:0
// Protected virtual
// void itemInserted(int)
func (this *QToolBox) ItemInserted(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox12itemInsertedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:100
// index:0
// Protected virtual
// void itemRemoved(int)
func (this *QToolBox) ItemRemoved(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11itemRemovedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:101
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QToolBox) ShowEvent(e *qtgui.QShowEvent /*444 QShowEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:102
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QToolBox) ChangeEvent(arg0 *qtcore.QEvent /*444 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
