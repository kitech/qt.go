//  header block begin
// /usr/include/qt/QtWidgets/qlistwidget.h
// #include <qlistwidget.h>
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
type QListWidgetItem struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qlistwidget.h:64
// index:0
// void QListWidgetItem(class QListWidget *, int)
func NewQListWidgetItem(view unsafe.Pointer, type_ int) *QListWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItemC2EP11QListWidgeti", ffiqt.FFI_TYPE_VOID, cthis, view, &type_)
	gopp.ErrPrint(err, rv)
	return &QListWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qlistwidget.h:65
// index:1
// void QListWidgetItem(const class QString &, class QListWidget *, int)
func NewQListWidgetItem_1(text unsafe.Pointer, view unsafe.Pointer, type_ int) *QListWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItemC2ERK7QStringP11QListWidgeti", ffiqt.FFI_TYPE_VOID, cthis, text, view, &type_)
	gopp.ErrPrint(err, rv)
	return &QListWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qlistwidget.h:66
// index:2
// void QListWidgetItem(const class QIcon &, const class QString &, class QListWidget *, int)
func NewQListWidgetItem_2(icon unsafe.Pointer, text unsafe.Pointer, view unsafe.Pointer, type_ int) *QListWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItemC2ERK5QIconRK7QStringP11QListWidgeti", ffiqt.FFI_TYPE_VOID, cthis, icon, text, view, &type_)
	gopp.ErrPrint(err, rv)
	return &QListWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qlistwidget.h:69
// index:0
// virtual
// void ~QListWidgetItem()
func DeleteQListWidgetItem(*QListWidgetItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:71
// index:0
// virtual
// QListWidgetItem * clone()
func (this *QListWidgetItem) Clone() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem5cloneEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:73
// index:0
// inline
// QListWidget * listWidget()
func (this *QListWidgetItem) ListWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem10listWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:75
// index:0
// inline
// void setSelected(_Bool)
func (this *QListWidgetItem) SetSelected(select_ bool) {
	// 0: (, bool select), (&select_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem11setSelectedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:76
// index:0
// inline
// bool isSelected()
func (this *QListWidgetItem) IsSelected() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem10isSelectedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:78
// index:0
// inline
// void setHidden(_Bool)
func (this *QListWidgetItem) SetHidden(hide bool) {
	// 0: (, bool hide), (&hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem9setHiddenEb", ffiqt.FFI_TYPE_VOID, this.cthis, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:79
// index:0
// inline
// bool isHidden()
func (this *QListWidgetItem) IsHidden() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem8isHiddenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:81
// index:0
// inline
// Qt::ItemFlags flags()
func (this *QListWidgetItem) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem5flagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:84
// index:0
// inline
// QString text()
func (this *QListWidgetItem) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:86
// index:0
// inline
// void setText(const class QString &)
func (this *QListWidgetItem) SetText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:88
// index:0
// inline
// QIcon icon()
func (this *QListWidgetItem) Icon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem4iconEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:90
// index:0
// inline
// void setIcon(const class QIcon &)
func (this *QListWidgetItem) SetIcon(icon unsafe.Pointer) {
	// 0: (, const QIcon & icon), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem7setIconERK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:92
// index:0
// inline
// QString statusTip()
func (this *QListWidgetItem) StatusTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem9statusTipEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:94
// index:0
// inline
// void setStatusTip(const class QString &)
func (this *QListWidgetItem) SetStatusTip(statusTip unsafe.Pointer) {
	// 0: (, const QString & statusTip), (statusTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem12setStatusTipERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, statusTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:97
// index:0
// inline
// QString toolTip()
func (this *QListWidgetItem) ToolTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem7toolTipEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:99
// index:0
// inline
// void setToolTip(const class QString &)
func (this *QListWidgetItem) SetToolTip(toolTip unsafe.Pointer) {
	// 0: (, const QString & toolTip), (toolTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem10setToolTipERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, toolTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:103
// index:0
// inline
// QString whatsThis()
func (this *QListWidgetItem) WhatsThis() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem9whatsThisEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:105
// index:0
// inline
// void setWhatsThis(const class QString &)
func (this *QListWidgetItem) SetWhatsThis(whatsThis unsafe.Pointer) {
	// 0: (, const QString & whatsThis), (whatsThis)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem12setWhatsThisERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, whatsThis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:108
// index:0
// inline
// QFont font()
func (this *QListWidgetItem) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem4fontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:110
// index:0
// inline
// void setFont(const class QFont &)
func (this *QListWidgetItem) SetFont(font unsafe.Pointer) {
	// 0: (, const QFont & font), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:112
// index:0
// inline
// int textAlignment()
func (this *QListWidgetItem) TextAlignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem13textAlignmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:114
// index:0
// inline
// void setTextAlignment(int)
func (this *QListWidgetItem) SetTextAlignment(alignment int) {
	// 0: (, int alignment), (&alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem16setTextAlignmentEi", ffiqt.FFI_TYPE_VOID, this.cthis, &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:117
// index:0
// inline
// QColor backgroundColor()
func (this *QListWidgetItem) BackgroundColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem15backgroundColorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:119
// index:0
// inline virtual
// void setBackgroundColor(const class QColor &)
func (this *QListWidgetItem) SetBackgroundColor(color unsafe.Pointer) {
	// 0: (, const QColor & color), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem18setBackgroundColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:122
// index:0
// inline
// QBrush background()
func (this *QListWidgetItem) Background() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem10backgroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:124
// index:0
// inline
// void setBackground(const class QBrush &)
func (this *QListWidgetItem) SetBackground(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:127
// index:0
// inline
// QColor textColor()
func (this *QListWidgetItem) TextColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem9textColorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:129
// index:0
// inline
// void setTextColor(const class QColor &)
func (this *QListWidgetItem) SetTextColor(color unsafe.Pointer) {
	// 0: (, const QColor & color), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem12setTextColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:132
// index:0
// inline
// QBrush foreground()
func (this *QListWidgetItem) Foreground() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem10foregroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:134
// index:0
// inline
// void setForeground(const class QBrush &)
func (this *QListWidgetItem) SetForeground(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem13setForegroundERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:137
// index:0
// inline
// Qt::CheckState checkState()
func (this *QListWidgetItem) CheckState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem10checkStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:139
// index:0
// inline
// void setCheckState(Qt::CheckState)
func (this *QListWidgetItem) SetCheckState(state int) {
	// 0: (, Qt::CheckState state), (&state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem13setCheckStateEN2Qt10CheckStateE", ffiqt.FFI_TYPE_VOID, this.cthis, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:142
// index:0
// inline
// QSize sizeHint()
func (this *QListWidgetItem) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:144
// index:0
// inline
// void setSizeHint(const class QSize &)
func (this *QListWidgetItem) SetSizeHint(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem11setSizeHintERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:147
// index:0
// virtual
// QVariant data(int)
func (this *QListWidgetItem) Data(role int) {
	// 0: (, int role), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem4dataEi", ffiqt.FFI_TYPE_VOID, this.cthis, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:148
// index:0
// virtual
// void setData(int, const class QVariant &)
func (this *QListWidgetItem) SetData(role int, value unsafe.Pointer) {
	// 0: (, int role, const QVariant & value), (&role, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem7setDataEiRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &role, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:153
// index:0
// virtual
// void read(class QDataStream &)
func (this *QListWidgetItem) Read(in unsafe.Pointer) {
	// 0: (, QDataStream & in), (in)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem4readER11QDataStream", ffiqt.FFI_TYPE_VOID, this.cthis, in)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:154
// index:0
// virtual
// void write(class QDataStream &)
func (this *QListWidgetItem) Write(out unsafe.Pointer) {
	// 0: (, QDataStream & out), (out)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem5writeER11QDataStream", ffiqt.FFI_TYPE_VOID, this.cthis, out)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:158
// index:0
// inline
// int type()
func (this *QListWidgetItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
