//  header block begin
// /usr/include/qt/QtWidgets/qtreewidget.h
// #include <qtreewidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QTreeWidgetItem struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qtreewidget.h:67
// index:0
// void QTreeWidgetItem(int)
func NewQTreeWidgetItem(type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	return &QTreeWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qtreewidget.h:68
// index:1
// void QTreeWidgetItem(const class QStringList &, int)
func NewQTreeWidgetItem_1(strings unsafe.Pointer, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2ERK11QStringListi", ffiqt.FFI_TYPE_VOID, cthis, strings, &type_)
	gopp.ErrPrint(err, rv)
	return &QTreeWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qtreewidget.h:69
// index:2
// void QTreeWidgetItem(class QTreeWidget *, int)
func NewQTreeWidgetItem_2(view unsafe.Pointer, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgeti", ffiqt.FFI_TYPE_VOID, cthis, view, &type_)
	gopp.ErrPrint(err, rv)
	return &QTreeWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qtreewidget.h:70
// index:3
// void QTreeWidgetItem(class QTreeWidget *, const class QStringList &, int)
func NewQTreeWidgetItem_3(view unsafe.Pointer, strings unsafe.Pointer, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgetRK11QStringListi", ffiqt.FFI_TYPE_VOID, cthis, view, strings, &type_)
	gopp.ErrPrint(err, rv)
	return &QTreeWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qtreewidget.h:71
// index:4
// void QTreeWidgetItem(class QTreeWidget *, class QTreeWidgetItem *, int)
func NewQTreeWidgetItem_4(view unsafe.Pointer, after unsafe.Pointer, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgetPS_i", ffiqt.FFI_TYPE_VOID, cthis, view, after, &type_)
	gopp.ErrPrint(err, rv)
	return &QTreeWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qtreewidget.h:72
// index:5
// void QTreeWidgetItem(class QTreeWidgetItem *, int)
func NewQTreeWidgetItem_5(parent unsafe.Pointer, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_i", ffiqt.FFI_TYPE_VOID, cthis, parent, &type_)
	gopp.ErrPrint(err, rv)
	return &QTreeWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qtreewidget.h:73
// index:6
// void QTreeWidgetItem(class QTreeWidgetItem *, const class QStringList &, int)
func NewQTreeWidgetItem_6(parent unsafe.Pointer, strings unsafe.Pointer, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_RK11QStringListi", ffiqt.FFI_TYPE_VOID, cthis, parent, strings, &type_)
	gopp.ErrPrint(err, rv)
	return &QTreeWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qtreewidget.h:74
// index:7
// void QTreeWidgetItem(class QTreeWidgetItem *, class QTreeWidgetItem *, int)
func NewQTreeWidgetItem_7(parent unsafe.Pointer, after unsafe.Pointer, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_S0_i", ffiqt.FFI_TYPE_VOID, cthis, parent, after, &type_)
	gopp.ErrPrint(err, rv)
	return &QTreeWidgetItem{cthis}
}

// /usr/include/qt/QtWidgets/qtreewidget.h:76
// index:0
// virtual
// void ~QTreeWidgetItem()
func DeleteQTreeWidgetItem(*QTreeWidgetItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:78
// index:0
// virtual
// QTreeWidgetItem * clone()
func (this *QTreeWidgetItem) Clone() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5cloneEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:80
// index:0
// inline
// QTreeWidget * treeWidget()
func (this *QTreeWidgetItem) TreeWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10treeWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:82
// index:0
// inline
// void setSelected(_Bool)
func (this *QTreeWidgetItem) SetSelected(select_ bool) {
	// 0: (, bool select), (&select_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setSelectedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:83
// index:0
// inline
// bool isSelected()
func (this *QTreeWidgetItem) IsSelected() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10isSelectedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:85
// index:0
// inline
// void setHidden(_Bool)
func (this *QTreeWidgetItem) SetHidden(hide bool) {
	// 0: (, bool hide), (&hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem9setHiddenEb", ffiqt.FFI_TYPE_VOID, this.cthis, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:86
// index:0
// inline
// bool isHidden()
func (this *QTreeWidgetItem) IsHidden() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem8isHiddenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:88
// index:0
// inline
// void setExpanded(_Bool)
func (this *QTreeWidgetItem) SetExpanded(expand bool) {
	// 0: (, bool expand), (&expand)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setExpandedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &expand)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:89
// index:0
// inline
// bool isExpanded()
func (this *QTreeWidgetItem) IsExpanded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10isExpandedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:91
// index:0
// inline
// void setFirstColumnSpanned(_Bool)
func (this *QTreeWidgetItem) SetFirstColumnSpanned(span bool) {
	// 0: (, bool span), (&span)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem21setFirstColumnSpannedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &span)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:92
// index:0
// inline
// bool isFirstColumnSpanned()
func (this *QTreeWidgetItem) IsFirstColumnSpanned() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem20isFirstColumnSpannedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:94
// index:0
// inline
// void setDisabled(_Bool)
func (this *QTreeWidgetItem) SetDisabled(disabled bool) {
	// 0: (, bool disabled), (&disabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setDisabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &disabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:95
// index:0
// inline
// bool isDisabled()
func (this *QTreeWidgetItem) IsDisabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10isDisabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:98
// index:0
// void setChildIndicatorPolicy(class QTreeWidgetItem::ChildIndicatorPolicy)
func (this *QTreeWidgetItem) SetChildIndicatorPolicy(policy int) {
	// 0: (, QTreeWidgetItem::ChildIndicatorPolicy policy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem23setChildIndicatorPolicyENS_20ChildIndicatorPolicyE", ffiqt.FFI_TYPE_VOID, this.cthis, &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:99
// index:0
// QTreeWidgetItem::ChildIndicatorPolicy childIndicatorPolicy()
func (this *QTreeWidgetItem) ChildIndicatorPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem20childIndicatorPolicyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:101
// index:0
// Qt::ItemFlags flags()
func (this *QTreeWidgetItem) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5flagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:104
// index:0
// inline
// QString text(int)
func (this *QTreeWidgetItem) Text(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4textEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:106
// index:0
// inline
// void setText(int, const class QString &)
func (this *QTreeWidgetItem) SetText(column int, text unsafe.Pointer) {
	// 0: (, int column, const QString & text), (&column, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setTextEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &column, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:108
// index:0
// inline
// QIcon icon(int)
func (this *QTreeWidgetItem) Icon(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4iconEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:110
// index:0
// inline
// void setIcon(int, const class QIcon &)
func (this *QTreeWidgetItem) SetIcon(column int, icon unsafe.Pointer) {
	// 0: (, int column, const QIcon & icon), (&column, icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setIconEiRK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, &column, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:112
// index:0
// inline
// QString statusTip(int)
func (this *QTreeWidgetItem) StatusTip(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem9statusTipEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:114
// index:0
// inline
// void setStatusTip(int, const class QString &)
func (this *QTreeWidgetItem) SetStatusTip(column int, statusTip unsafe.Pointer) {
	// 0: (, int column, const QString & statusTip), (&column, statusTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem12setStatusTipEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &column, statusTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:117
// index:0
// inline
// QString toolTip(int)
func (this *QTreeWidgetItem) ToolTip(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem7toolTipEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:119
// index:0
// inline
// void setToolTip(int, const class QString &)
func (this *QTreeWidgetItem) SetToolTip(column int, toolTip unsafe.Pointer) {
	// 0: (, int column, const QString & toolTip), (&column, toolTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem10setToolTipEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &column, toolTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:123
// index:0
// inline
// QString whatsThis(int)
func (this *QTreeWidgetItem) WhatsThis(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem9whatsThisEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:125
// index:0
// inline
// void setWhatsThis(int, const class QString &)
func (this *QTreeWidgetItem) SetWhatsThis(column int, whatsThis unsafe.Pointer) {
	// 0: (, int column, const QString & whatsThis), (&column, whatsThis)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem12setWhatsThisEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &column, whatsThis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:128
// index:0
// inline
// QFont font(int)
func (this *QTreeWidgetItem) Font(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4fontEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:130
// index:0
// inline
// void setFont(int, const class QFont &)
func (this *QTreeWidgetItem) SetFont(column int, font unsafe.Pointer) {
	// 0: (, int column, const QFont & font), (&column, font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setFontEiRK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, &column, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:132
// index:0
// inline
// int textAlignment(int)
func (this *QTreeWidgetItem) TextAlignment(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem13textAlignmentEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:134
// index:0
// inline
// void setTextAlignment(int, int)
func (this *QTreeWidgetItem) SetTextAlignment(column int, alignment int) {
	// 0: (, int column, int alignment), (&column, &alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem16setTextAlignmentEii", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:137
// index:0
// inline
// QColor backgroundColor(int)
func (this *QTreeWidgetItem) BackgroundColor(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem15backgroundColorEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:139
// index:0
// inline
// void setBackgroundColor(int, const class QColor &)
func (this *QTreeWidgetItem) SetBackgroundColor(column int, color unsafe.Pointer) {
	// 0: (, int column, const QColor & color), (&column, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem18setBackgroundColorEiRK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, &column, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:142
// index:0
// inline
// QBrush background(int)
func (this *QTreeWidgetItem) Background(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10backgroundEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:144
// index:0
// inline
// void setBackground(int, const class QBrush &)
func (this *QTreeWidgetItem) SetBackground(column int, brush unsafe.Pointer) {
	// 0: (, int column, const QBrush & brush), (&column, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem13setBackgroundEiRK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, &column, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:147
// index:0
// inline
// QColor textColor(int)
func (this *QTreeWidgetItem) TextColor(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem9textColorEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:149
// index:0
// inline
// void setTextColor(int, const class QColor &)
func (this *QTreeWidgetItem) SetTextColor(column int, color unsafe.Pointer) {
	// 0: (, int column, const QColor & color), (&column, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem12setTextColorEiRK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, &column, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:152
// index:0
// inline
// QBrush foreground(int)
func (this *QTreeWidgetItem) Foreground(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10foregroundEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:154
// index:0
// inline
// void setForeground(int, const class QBrush &)
func (this *QTreeWidgetItem) SetForeground(column int, brush unsafe.Pointer) {
	// 0: (, int column, const QBrush & brush), (&column, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem13setForegroundEiRK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, &column, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:157
// index:0
// inline
// Qt::CheckState checkState(int)
func (this *QTreeWidgetItem) CheckState(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10checkStateEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:159
// index:0
// inline
// void setCheckState(int, Qt::CheckState)
func (this *QTreeWidgetItem) SetCheckState(column int, state int) {
	// 0: (, int column, Qt::CheckState state), (&column, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem13setCheckStateEiN2Qt10CheckStateE", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:162
// index:0
// inline
// QSize sizeHint(int)
func (this *QTreeWidgetItem) SizeHint(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem8sizeHintEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:164
// index:0
// inline
// void setSizeHint(int, const class QSize &)
func (this *QTreeWidgetItem) SetSizeHint(column int, size unsafe.Pointer) {
	// 0: (, int column, const QSize & size), (&column, size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setSizeHintEiRK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, &column, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:167
// index:0
// virtual
// QVariant data(int, int)
func (this *QTreeWidgetItem) Data(column int, role int) {
	// 0: (, int column, int role), (&column, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4dataEii", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:168
// index:0
// virtual
// void setData(int, int, const class QVariant &)
func (this *QTreeWidgetItem) SetData(column int, role int, value unsafe.Pointer) {
	// 0: (, int column, int role, const QVariant & value), (&column, &role, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setDataEiiRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &role, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:173
// index:0
// virtual
// void read(class QDataStream &)
func (this *QTreeWidgetItem) Read(in unsafe.Pointer) {
	// 0: (, QDataStream & in), (in)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem4readER11QDataStream", ffiqt.FFI_TYPE_VOID, this.cthis, in)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:174
// index:0
// virtual
// void write(class QDataStream &)
func (this *QTreeWidgetItem) Write(out unsafe.Pointer) {
	// 0: (, QDataStream & out), (out)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5writeER11QDataStream", ffiqt.FFI_TYPE_VOID, this.cthis, out)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:178
// index:0
// inline
// QTreeWidgetItem * parent()
func (this *QTreeWidgetItem) Parent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem6parentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:179
// index:0
// inline
// QTreeWidgetItem * child(int)
func (this *QTreeWidgetItem) Child(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5childEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:185
// index:0
// inline
// int childCount()
func (this *QTreeWidgetItem) ChildCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10childCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:186
// index:0
// inline
// int columnCount()
func (this *QTreeWidgetItem) ColumnCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem11columnCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:187
// index:0
// inline
// int indexOfChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) IndexOfChild(child unsafe.Pointer) {
	// 0: (, QTreeWidgetItem * child), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem12indexOfChildEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:189
// index:0
// void addChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) AddChild(child unsafe.Pointer) {
	// 0: (, QTreeWidgetItem * child), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem8addChildEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:190
// index:0
// void insertChild(int, class QTreeWidgetItem *)
func (this *QTreeWidgetItem) InsertChild(index int, child unsafe.Pointer) {
	// 0: (, int index, QTreeWidgetItem * child), (&index, child)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11insertChildEiPS_", ffiqt.FFI_TYPE_VOID, this.cthis, &index, child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:191
// index:0
// void removeChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) RemoveChild(child unsafe.Pointer) {
	// 0: (, QTreeWidgetItem * child), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11removeChildEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:192
// index:0
// QTreeWidgetItem * takeChild(int)
func (this *QTreeWidgetItem) TakeChild(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem9takeChildEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:196
// index:0
// QList<QTreeWidgetItem *> takeChildren()
func (this *QTreeWidgetItem) TakeChildren() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem12takeChildrenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:198
// index:0
// inline
// int type()
func (this *QTreeWidgetItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:199
// index:0
// inline
// void sortChildren(int, Qt::SortOrder)
func (this *QTreeWidgetItem) SortChildren(column int, order int) {
	// 0: (, int column, Qt::SortOrder order), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem12sortChildrenEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &order)
	gopp.ErrPrint(err, rv)
}

//  body block end
