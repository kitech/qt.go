//  header block begin
// /usr/include/qt/QtGui/qstandarditemmodel.h
// #include <qstandarditemmodel.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QStandardItem struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:65
// index:0
// void QStandardItem()
func NewQStandardItem() *QStandardItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QStandardItem{cthis}
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:66
// index:1
// void QStandardItem(const class QString &)
func NewQStandardItem_1(text unsafe.Pointer) *QStandardItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, text)
	gopp.ErrPrint(err, rv)
	return &QStandardItem{cthis}
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:67
// index:2
// void QStandardItem(const class QIcon &, const class QString &)
func NewQStandardItem_2(icon unsafe.Pointer, text unsafe.Pointer) *QStandardItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2ERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, cthis, icon, text)
	gopp.ErrPrint(err, rv)
	return &QStandardItem{cthis}
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:68
// index:3
// void QStandardItem(int, int)
func NewQStandardItem_3(rows int, columns int) *QStandardItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2Eii", ffiqt.FFI_TYPE_VOID, cthis, &rows, &columns)
	gopp.ErrPrint(err, rv)
	return &QStandardItem{cthis}
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:69
// index:0
// virtual
// void ~QStandardItem()
func DeleteQStandardItem(*QStandardItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:71
// index:0
// virtual
// QVariant data(int)
func (this *QStandardItem) Data(role int) {
	// 0: (, int role), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4dataEi", ffiqt.FFI_TYPE_VOID, this.cthis, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:72
// index:0
// virtual
// void setData(const class QVariant &, int)
func (this *QStandardItem) SetData(value unsafe.Pointer, role int) {
	// 0: (, const QVariant & value, int role), (value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setDataERK8QVarianti", ffiqt.FFI_TYPE_VOID, this.cthis, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:74
// index:0
// inline
// QString text()
func (this *QStandardItem) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:77
// index:0
// inline
// void setText(const class QString &)
func (this *QStandardItem) SetText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:79
// index:0
// inline
// QIcon icon()
func (this *QStandardItem) Icon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4iconEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:82
// index:0
// inline
// void setIcon(const class QIcon &)
func (this *QStandardItem) SetIcon(icon unsafe.Pointer) {
	// 0: (, const QIcon & icon), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setIconERK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:85
// index:0
// inline
// QString toolTip()
func (this *QStandardItem) ToolTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem7toolTipEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:88
// index:0
// inline
// void setToolTip(const class QString &)
func (this *QStandardItem) SetToolTip(toolTip unsafe.Pointer) {
	// 0: (, const QString & toolTip), (toolTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10setToolTipERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, toolTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:92
// index:0
// inline
// QString statusTip()
func (this *QStandardItem) StatusTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem9statusTipEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:95
// index:0
// inline
// void setStatusTip(const class QString &)
func (this *QStandardItem) SetStatusTip(statusTip unsafe.Pointer) {
	// 0: (, const QString & statusTip), (statusTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12setStatusTipERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, statusTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:99
// index:0
// inline
// QString whatsThis()
func (this *QStandardItem) WhatsThis() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem9whatsThisEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:102
// index:0
// inline
// void setWhatsThis(const class QString &)
func (this *QStandardItem) SetWhatsThis(whatsThis unsafe.Pointer) {
	// 0: (, const QString & whatsThis), (whatsThis)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12setWhatsThisERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, whatsThis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:105
// index:0
// inline
// QSize sizeHint()
func (this *QStandardItem) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:108
// index:0
// inline
// void setSizeHint(const class QSize &)
func (this *QStandardItem) SetSizeHint(sizeHint unsafe.Pointer) {
	// 0: (, const QSize & sizeHint), (sizeHint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setSizeHintERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, sizeHint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:110
// index:0
// inline
// QFont font()
func (this *QStandardItem) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4fontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:113
// index:0
// inline
// void setFont(const class QFont &)
func (this *QStandardItem) SetFont(font unsafe.Pointer) {
	// 0: (, const QFont & font), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:115
// index:0
// inline
// Qt::Alignment textAlignment()
func (this *QStandardItem) TextAlignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem13textAlignmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:120
// index:0
// inline
// QBrush background()
func (this *QStandardItem) Background() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10backgroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:123
// index:0
// inline
// void setBackground(const class QBrush &)
func (this *QStandardItem) SetBackground(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:125
// index:0
// inline
// QBrush foreground()
func (this *QStandardItem) Foreground() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10foregroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:128
// index:0
// inline
// void setForeground(const class QBrush &)
func (this *QStandardItem) SetForeground(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setForegroundERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:130
// index:0
// inline
// Qt::CheckState checkState()
func (this *QStandardItem) CheckState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10checkStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:133
// index:0
// inline
// void setCheckState(Qt::CheckState)
func (this *QStandardItem) SetCheckState(checkState int) {
	// 0: (, Qt::CheckState checkState), (&checkState)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setCheckStateEN2Qt10CheckStateE", ffiqt.FFI_TYPE_VOID, this.cthis, &checkState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:135
// index:0
// inline
// QString accessibleText()
func (this *QStandardItem) AccessibleText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem14accessibleTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:138
// index:0
// inline
// void setAccessibleText(const class QString &)
func (this *QStandardItem) SetAccessibleText(accessibleText unsafe.Pointer) {
	// 0: (, const QString & accessibleText), (accessibleText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem17setAccessibleTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, accessibleText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:140
// index:0
// inline
// QString accessibleDescription()
func (this *QStandardItem) AccessibleDescription() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem21accessibleDescriptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:143
// index:0
// inline
// void setAccessibleDescription(const class QString &)
func (this *QStandardItem) SetAccessibleDescription(accessibleDescription unsafe.Pointer) {
	// 0: (, const QString & accessibleDescription), (accessibleDescription)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem24setAccessibleDescriptionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, accessibleDescription)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:145
// index:0
// Qt::ItemFlags flags()
func (this *QStandardItem) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5flagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:148
// index:0
// inline
// bool isEnabled()
func (this *QStandardItem) IsEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem9isEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:151
// index:0
// void setEnabled(_Bool)
func (this *QStandardItem) SetEnabled(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10setEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:153
// index:0
// inline
// bool isEditable()
func (this *QStandardItem) IsEditable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10isEditableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:156
// index:0
// void setEditable(_Bool)
func (this *QStandardItem) SetEditable(editable bool) {
	// 0: (, bool editable), (&editable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setEditableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &editable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:158
// index:0
// inline
// bool isSelectable()
func (this *QStandardItem) IsSelectable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem12isSelectableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:161
// index:0
// void setSelectable(_Bool)
func (this *QStandardItem) SetSelectable(selectable bool) {
	// 0: (, bool selectable), (&selectable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setSelectableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &selectable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:163
// index:0
// inline
// bool isCheckable()
func (this *QStandardItem) IsCheckable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem11isCheckableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:166
// index:0
// void setCheckable(_Bool)
func (this *QStandardItem) SetCheckable(checkable bool) {
	// 0: (, bool checkable), (&checkable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12setCheckableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &checkable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:168
// index:0
// inline
// bool isAutoTristate()
func (this *QStandardItem) IsAutoTristate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem14isAutoTristateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:171
// index:0
// void setAutoTristate(_Bool)
func (this *QStandardItem) SetAutoTristate(tristate bool) {
	// 0: (, bool tristate), (&tristate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem15setAutoTristateEb", ffiqt.FFI_TYPE_VOID, this.cthis, &tristate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:173
// index:0
// inline
// bool isUserTristate()
func (this *QStandardItem) IsUserTristate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem14isUserTristateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:176
// index:0
// void setUserTristate(_Bool)
func (this *QStandardItem) SetUserTristate(tristate bool) {
	// 0: (, bool tristate), (&tristate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem15setUserTristateEb", ffiqt.FFI_TYPE_VOID, this.cthis, &tristate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:179
// index:0
// inline
// bool isTristate()
func (this *QStandardItem) IsTristate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10isTristateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:180
// index:0
// void setTristate(_Bool)
func (this *QStandardItem) SetTristate(tristate bool) {
	// 0: (, bool tristate), (&tristate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setTristateEb", ffiqt.FFI_TYPE_VOID, this.cthis, &tristate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:184
// index:0
// inline
// bool isDragEnabled()
func (this *QStandardItem) IsDragEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem13isDragEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:187
// index:0
// void setDragEnabled(_Bool)
func (this *QStandardItem) SetDragEnabled(dragEnabled bool) {
	// 0: (, bool dragEnabled), (&dragEnabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem14setDragEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &dragEnabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:189
// index:0
// inline
// bool isDropEnabled()
func (this *QStandardItem) IsDropEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem13isDropEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:192
// index:0
// void setDropEnabled(_Bool)
func (this *QStandardItem) SetDropEnabled(dropEnabled bool) {
	// 0: (, bool dropEnabled), (&dropEnabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem14setDropEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &dropEnabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:195
// index:0
// QStandardItem * parent()
func (this *QStandardItem) Parent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem6parentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:196
// index:0
// int row()
func (this *QStandardItem) Row() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem3rowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:197
// index:0
// int column()
func (this *QStandardItem) Column() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem6columnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:198
// index:0
// QModelIndex index()
func (this *QStandardItem) Index() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5indexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:199
// index:0
// QStandardItemModel * model()
func (this *QStandardItem) Model() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5modelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:201
// index:0
// int rowCount()
func (this *QStandardItem) RowCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem8rowCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:202
// index:0
// void setRowCount(int)
func (this *QStandardItem) SetRowCount(rows int) {
	// 0: (, int rows), (&rows)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setRowCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &rows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:203
// index:0
// int columnCount()
func (this *QStandardItem) ColumnCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem11columnCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:204
// index:0
// void setColumnCount(int)
func (this *QStandardItem) SetColumnCount(columns int) {
	// 0: (, int columns), (&columns)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem14setColumnCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:206
// index:0
// bool hasChildren()
func (this *QStandardItem) HasChildren() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem11hasChildrenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:207
// index:0
// QStandardItem * child(int, int)
func (this *QStandardItem) Child(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5childEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:208
// index:0
// void setChild(int, int, class QStandardItem *)
func (this *QStandardItem) SetChild(row int, column int, item unsafe.Pointer) {
	// 0: (, int row, int column, QStandardItem * item), (&row, &column, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem8setChildEiiPS_", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:209
// index:1
// inline
// void setChild(int, class QStandardItem *)
func (this *QStandardItem) SetChild_1(row int, item unsafe.Pointer) {
	// 1: (, int row, QStandardItem * item), (&row, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem8setChildEiPS_", ffiqt.FFI_TYPE_VOID, this.cthis, &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:214
// index:0
// void insertRows(int, int)
func (this *QStandardItem) InsertRows(row int, count int) {
	// 0: (, int row, int count), (&row, &count)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10insertRowsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:215
// index:0
// void insertColumns(int, int)
func (this *QStandardItem) InsertColumns(column int, count int) {
	// 0: (, int column, int count), (&column, &count)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13insertColumnsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:217
// index:0
// void removeRow(int)
func (this *QStandardItem) RemoveRow(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9removeRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:218
// index:0
// void removeColumn(int)
func (this *QStandardItem) RemoveColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12removeColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:219
// index:0
// void removeRows(int, int)
func (this *QStandardItem) RemoveRows(row int, count int) {
	// 0: (, int row, int count), (&row, &count)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10removeRowsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:220
// index:0
// void removeColumns(int, int)
func (this *QStandardItem) RemoveColumns(column int, count int) {
	// 0: (, int column, int count), (&column, &count)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13removeColumnsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:225
// index:0
// inline
// void insertRow(int, class QStandardItem *)
func (this *QStandardItem) InsertRow(row int, item unsafe.Pointer) {
	// 0: (, int row, QStandardItem * item), (&row, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9insertRowEiPS_", ffiqt.FFI_TYPE_VOID, this.cthis, &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:226
// index:0
// inline
// void appendRow(class QStandardItem *)
func (this *QStandardItem) AppendRow(item unsafe.Pointer) {
	// 0: (, QStandardItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9appendRowEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:228
// index:0
// QStandardItem * takeChild(int, int)
func (this *QStandardItem) TakeChild(row int, column int) {
	// 0: (, int row, int column), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9takeChildEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:229
// index:0
// QList<QStandardItem *> takeRow(int)
func (this *QStandardItem) TakeRow(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7takeRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:230
// index:0
// QList<QStandardItem *> takeColumn(int)
func (this *QStandardItem) TakeColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10takeColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:232
// index:0
// void sortChildren(int, Qt::SortOrder)
func (this *QStandardItem) SortChildren(column int, order int) {
	// 0: (, int column, Qt::SortOrder order), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12sortChildrenEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:234
// index:0
// virtual
// QStandardItem * clone()
func (this *QStandardItem) Clone() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5cloneEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:237
// index:0
// virtual
// int type()
func (this *QStandardItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:240
// index:0
// virtual
// void read(class QDataStream &)
func (this *QStandardItem) Read(in unsafe.Pointer) {
	// 0: (, QDataStream & in), (in)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem4readER11QDataStream", ffiqt.FFI_TYPE_VOID, this.cthis, in)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:241
// index:0
// virtual
// void write(class QDataStream &)
func (this *QStandardItem) Write(out unsafe.Pointer) {
	// 0: (, QDataStream & out), (out)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5writeER11QDataStream", ffiqt.FFI_TYPE_VOID, this.cthis, out)
	gopp.ErrPrint(err, rv)
}

//  body block end
