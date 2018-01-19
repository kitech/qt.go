//  header block begin
// /usr/include/qt/QtWidgets/qtablewidget.h
// #include <qtablewidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QTableWidgetItem struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qtablewidget.h:82
// index:0
// void QTableWidgetItem(int)
func NewQTableWidgetItem(type_ int) *QTableWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	return &QTableWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qtablewidget.h:83
// index:1
// void QTableWidgetItem(const class QString &, int)
func NewQTableWidgetItem_1(text unsafe.Pointer, type_ int) *QTableWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemC2ERK7QStringi", ffiqt.FFI_TYPE_VOID, cthis, text, &type_)
	gopp.ErrPrint(err, rv)
	return &QTableWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qtablewidget.h:84
// index:2
// void QTableWidgetItem(const class QIcon &, const class QString &, int)
func NewQTableWidgetItem_2(icon unsafe.Pointer, text unsafe.Pointer, type_ int) *QTableWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemC2ERK5QIconRK7QStringi", ffiqt.FFI_TYPE_VOID, cthis, icon, text, &type_)
	gopp.ErrPrint(err, rv)
	return &QTableWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qtablewidget.h:86
// index:0
// virtual
// void ~QTableWidgetItem()
func DeleteQTableWidgetItem(*QTableWidgetItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:88
// index:0
// virtual
// QTableWidgetItem * clone()
func (this *QTableWidgetItem) Clone() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem5cloneEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:90
// index:0
// inline
// QTableWidget * tableWidget()
func (this *QTableWidgetItem) TableWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem11tableWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:92
// index:0
// inline
// int row()
func (this *QTableWidgetItem) Row() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem3rowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:93
// index:0
// inline
// int column()
func (this *QTableWidgetItem) Column() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem6columnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:95
// index:0
// inline
// void setSelected(_Bool)
func (this *QTableWidgetItem) SetSelected(select_ bool) {
	// 0: (, bool select), (&select_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem11setSelectedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:96
// index:0
// inline
// bool isSelected()
func (this *QTableWidgetItem) IsSelected() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10isSelectedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:98
// index:0
// inline
// Qt::ItemFlags flags()
func (this *QTableWidgetItem) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem5flagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:101
// index:0
// inline
// QString text()
func (this *QTableWidgetItem) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:103
// index:0
// inline
// void setText(const class QString &)
func (this *QTableWidgetItem) SetText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:105
// index:0
// inline
// QIcon icon()
func (this *QTableWidgetItem) Icon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4iconEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:107
// index:0
// inline
// void setIcon(const class QIcon &)
func (this *QTableWidgetItem) SetIcon(icon unsafe.Pointer) {
	// 0: (, const QIcon & icon), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setIconERK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:109
// index:0
// inline
// QString statusTip()
func (this *QTableWidgetItem) StatusTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem9statusTipEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:111
// index:0
// inline
// void setStatusTip(const class QString &)
func (this *QTableWidgetItem) SetStatusTip(statusTip unsafe.Pointer) {
	// 0: (, const QString & statusTip), (statusTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem12setStatusTipERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, statusTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:114
// index:0
// inline
// QString toolTip()
func (this *QTableWidgetItem) ToolTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem7toolTipEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:116
// index:0
// inline
// void setToolTip(const class QString &)
func (this *QTableWidgetItem) SetToolTip(toolTip unsafe.Pointer) {
	// 0: (, const QString & toolTip), (toolTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem10setToolTipERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, toolTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:120
// index:0
// inline
// QString whatsThis()
func (this *QTableWidgetItem) WhatsThis() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem9whatsThisEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:122
// index:0
// inline
// void setWhatsThis(const class QString &)
func (this *QTableWidgetItem) SetWhatsThis(whatsThis unsafe.Pointer) {
	// 0: (, const QString & whatsThis), (whatsThis)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem12setWhatsThisERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, whatsThis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:125
// index:0
// inline
// QFont font()
func (this *QTableWidgetItem) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4fontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:127
// index:0
// inline
// void setFont(const class QFont &)
func (this *QTableWidgetItem) SetFont(font unsafe.Pointer) {
	// 0: (, const QFont & font), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:129
// index:0
// inline
// int textAlignment()
func (this *QTableWidgetItem) TextAlignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem13textAlignmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:131
// index:0
// inline
// void setTextAlignment(int)
func (this *QTableWidgetItem) SetTextAlignment(alignment int) {
	// 0: (, int alignment), (&alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem16setTextAlignmentEi", ffiqt.FFI_TYPE_VOID, this.cthis, &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:134
// index:0
// inline
// QColor backgroundColor()
func (this *QTableWidgetItem) BackgroundColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem15backgroundColorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:136
// index:0
// inline
// void setBackgroundColor(const class QColor &)
func (this *QTableWidgetItem) SetBackgroundColor(color unsafe.Pointer) {
	// 0: (, const QColor & color), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem18setBackgroundColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:139
// index:0
// inline
// QBrush background()
func (this *QTableWidgetItem) Background() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10backgroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:141
// index:0
// inline
// void setBackground(const class QBrush &)
func (this *QTableWidgetItem) SetBackground(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:144
// index:0
// inline
// QColor textColor()
func (this *QTableWidgetItem) TextColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem9textColorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:146
// index:0
// inline
// void setTextColor(const class QColor &)
func (this *QTableWidgetItem) SetTextColor(color unsafe.Pointer) {
	// 0: (, const QColor & color), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem12setTextColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:149
// index:0
// inline
// QBrush foreground()
func (this *QTableWidgetItem) Foreground() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10foregroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:151
// index:0
// inline
// void setForeground(const class QBrush &)
func (this *QTableWidgetItem) SetForeground(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem13setForegroundERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:154
// index:0
// inline
// Qt::CheckState checkState()
func (this *QTableWidgetItem) CheckState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10checkStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:156
// index:0
// inline
// void setCheckState(Qt::CheckState)
func (this *QTableWidgetItem) SetCheckState(state int) {
	// 0: (, Qt::CheckState state), (&state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem13setCheckStateEN2Qt10CheckStateE", ffiqt.FFI_TYPE_VOID, this.cthis, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:159
// index:0
// inline
// QSize sizeHint()
func (this *QTableWidgetItem) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:161
// index:0
// inline
// void setSizeHint(const class QSize &)
func (this *QTableWidgetItem) SetSizeHint(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem11setSizeHintERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:164
// index:0
// virtual
// QVariant data(int)
func (this *QTableWidgetItem) Data(role int) {
	// 0: (, int role), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4dataEi", ffiqt.FFI_TYPE_VOID, this.cthis, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:165
// index:0
// virtual
// void setData(int, const class QVariant &)
func (this *QTableWidgetItem) SetData(role int, value unsafe.Pointer) {
	// 0: (, int role, const QVariant & value), (&role, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setDataEiRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &role, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:170
// index:0
// virtual
// void read(class QDataStream &)
func (this *QTableWidgetItem) Read(in unsafe.Pointer) {
	// 0: (, QDataStream & in), (in)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem4readER11QDataStream", ffiqt.FFI_TYPE_VOID, this.cthis, in)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:171
// index:0
// virtual
// void write(class QDataStream &)
func (this *QTableWidgetItem) Write(out unsafe.Pointer) {
	// 0: (, QDataStream & out), (out)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem5writeER11QDataStream", ffiqt.FFI_TYPE_VOID, this.cthis, out)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:175
// index:0
// inline
// int type()
func (this *QTableWidgetItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
