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
	*qtrt.CObject
}

func (this *QStandardItem) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQStandardItemFromPointer(cthis unsafe.Pointer) *QStandardItem {
	return &QStandardItem{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:65
// index:0
// Public
// void QStandardItem()
func NewQStandardItem() *QStandardItem {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:66
// index:1
// Public
// void QStandardItem(const class QString &)
func NewQStandardItem_1(text *qtcore.QString) *QStandardItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:67
// index:2
// Public
// void QStandardItem(const class QIcon &, const class QString &)
func NewQStandardItem_2(icon *QIcon, text *qtcore.QString) *QStandardItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2ERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:68
// index:3
// Public
// void QStandardItem(int, int)
func NewQStandardItem_3(rows int, columns int) *QStandardItem {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2Eii", ffiqt.FFI_TYPE_VOID, cthis, &rows, &columns)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:69
// index:0
// Public virtual
// void ~QStandardItem()
func DeleteQStandardItem(*QStandardItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:71
// index:0
// Public virtual
// QVariant data(int)
func (this *QStandardItem) Data(role int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4dataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:72
// index:0
// Public virtual
// void setData(const class QVariant &, int)
func (this *QStandardItem) SetData(value *qtcore.QVariant, role int) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setDataERK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:74
// index:0
// Public inline
// QString text()
func (this *QStandardItem) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:77
// index:0
// Public inline
// void setText(const class QString &)
func (this *QStandardItem) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:79
// index:0
// Public inline
// QIcon icon()
func (this *QStandardItem) Icon() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4iconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:82
// index:0
// Public inline
// void setIcon(const class QIcon &)
func (this *QStandardItem) SetIcon(icon *QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:85
// index:0
// Public inline
// QString toolTip()
func (this *QStandardItem) ToolTip() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem7toolTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:88
// index:0
// Public inline
// void setToolTip(const class QString &)
func (this *QStandardItem) SetToolTip(toolTip *qtcore.QString) {
	var convArg0 = toolTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10setToolTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:92
// index:0
// Public inline
// QString statusTip()
func (this *QStandardItem) StatusTip() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem9statusTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:95
// index:0
// Public inline
// void setStatusTip(const class QString &)
func (this *QStandardItem) SetStatusTip(statusTip *qtcore.QString) {
	var convArg0 = statusTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12setStatusTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:99
// index:0
// Public inline
// QString whatsThis()
func (this *QStandardItem) WhatsThis() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem9whatsThisEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:102
// index:0
// Public inline
// void setWhatsThis(const class QString &)
func (this *QStandardItem) SetWhatsThis(whatsThis *qtcore.QString) {
	var convArg0 = whatsThis.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12setWhatsThisERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:105
// index:0
// Public inline
// QSize sizeHint()
func (this *QStandardItem) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:108
// index:0
// Public inline
// void setSizeHint(const class QSize &)
func (this *QStandardItem) SetSizeHint(sizeHint *qtcore.QSize) {
	var convArg0 = sizeHint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setSizeHintERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:110
// index:0
// Public inline
// QFont font()
func (this *QStandardItem) Font() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:113
// index:0
// Public inline
// void setFont(const class QFont &)
func (this *QStandardItem) SetFont(font *QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:115
// index:0
// Public inline
// Qt::Alignment textAlignment()
func (this *QStandardItem) TextAlignment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem13textAlignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:120
// index:0
// Public inline
// QBrush background()
func (this *QStandardItem) Background() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10backgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:123
// index:0
// Public inline
// void setBackground(const class QBrush &)
func (this *QStandardItem) SetBackground(brush *QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:125
// index:0
// Public inline
// QBrush foreground()
func (this *QStandardItem) Foreground() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10foregroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:128
// index:0
// Public inline
// void setForeground(const class QBrush &)
func (this *QStandardItem) SetForeground(brush *QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setForegroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:130
// index:0
// Public inline
// Qt::CheckState checkState()
func (this *QStandardItem) CheckState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10checkStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:133
// index:0
// Public inline
// void setCheckState(Qt::CheckState)
func (this *QStandardItem) SetCheckState(checkState int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setCheckStateEN2Qt10CheckStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &checkState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:135
// index:0
// Public inline
// QString accessibleText()
func (this *QStandardItem) AccessibleText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem14accessibleTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:138
// index:0
// Public inline
// void setAccessibleText(const class QString &)
func (this *QStandardItem) SetAccessibleText(accessibleText *qtcore.QString) {
	var convArg0 = accessibleText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem17setAccessibleTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:140
// index:0
// Public inline
// QString accessibleDescription()
func (this *QStandardItem) AccessibleDescription() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem21accessibleDescriptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:143
// index:0
// Public inline
// void setAccessibleDescription(const class QString &)
func (this *QStandardItem) SetAccessibleDescription(accessibleDescription *qtcore.QString) {
	var convArg0 = accessibleDescription.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem24setAccessibleDescriptionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:145
// index:0
// Public
// Qt::ItemFlags flags()
func (this *QStandardItem) Flags() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:148
// index:0
// Public inline
// bool isEnabled()
func (this *QStandardItem) IsEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:151
// index:0
// Public
// void setEnabled(_Bool)
func (this *QStandardItem) SetEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:153
// index:0
// Public inline
// bool isEditable()
func (this *QStandardItem) IsEditable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10isEditableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:156
// index:0
// Public
// void setEditable(_Bool)
func (this *QStandardItem) SetEditable(editable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setEditableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &editable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:158
// index:0
// Public inline
// bool isSelectable()
func (this *QStandardItem) IsSelectable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem12isSelectableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:161
// index:0
// Public
// void setSelectable(_Bool)
func (this *QStandardItem) SetSelectable(selectable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setSelectableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &selectable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:163
// index:0
// Public inline
// bool isCheckable()
func (this *QStandardItem) IsCheckable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem11isCheckableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:166
// index:0
// Public
// void setCheckable(_Bool)
func (this *QStandardItem) SetCheckable(checkable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12setCheckableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &checkable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:168
// index:0
// Public inline
// bool isAutoTristate()
func (this *QStandardItem) IsAutoTristate() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem14isAutoTristateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:171
// index:0
// Public
// void setAutoTristate(_Bool)
func (this *QStandardItem) SetAutoTristate(tristate bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem15setAutoTristateEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &tristate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:173
// index:0
// Public inline
// bool isUserTristate()
func (this *QStandardItem) IsUserTristate() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem14isUserTristateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:176
// index:0
// Public
// void setUserTristate(_Bool)
func (this *QStandardItem) SetUserTristate(tristate bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem15setUserTristateEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &tristate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:179
// index:0
// Public inline
// bool isTristate()
func (this *QStandardItem) IsTristate() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10isTristateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:180
// index:0
// Public
// void setTristate(_Bool)
func (this *QStandardItem) SetTristate(tristate bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setTristateEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &tristate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:184
// index:0
// Public inline
// bool isDragEnabled()
func (this *QStandardItem) IsDragEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem13isDragEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:187
// index:0
// Public
// void setDragEnabled(_Bool)
func (this *QStandardItem) SetDragEnabled(dragEnabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem14setDragEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dragEnabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:189
// index:0
// Public inline
// bool isDropEnabled()
func (this *QStandardItem) IsDropEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem13isDropEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:192
// index:0
// Public
// void setDropEnabled(_Bool)
func (this *QStandardItem) SetDropEnabled(dropEnabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem14setDropEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dropEnabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:195
// index:0
// Public
// QStandardItem * parent()
func (this *QStandardItem) Parent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:196
// index:0
// Public
// int row()
func (this *QStandardItem) Row() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:197
// index:0
// Public
// int column()
func (this *QStandardItem) Column() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem6columnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:198
// index:0
// Public
// QModelIndex index()
func (this *QStandardItem) Index() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5indexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:199
// index:0
// Public
// QStandardItemModel * model()
func (this *QStandardItem) Model() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:201
// index:0
// Public
// int rowCount()
func (this *QStandardItem) RowCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem8rowCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:202
// index:0
// Public
// void setRowCount(int)
func (this *QStandardItem) SetRowCount(rows int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setRowCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &rows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:203
// index:0
// Public
// int columnCount()
func (this *QStandardItem) ColumnCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:204
// index:0
// Public
// void setColumnCount(int)
func (this *QStandardItem) SetColumnCount(columns int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem14setColumnCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:206
// index:0
// Public
// bool hasChildren()
func (this *QStandardItem) HasChildren() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem11hasChildrenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:207
// index:0
// Public
// QStandardItem * child(int, int)
func (this *QStandardItem) Child(row int, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5childEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:208
// index:0
// Public
// void setChild(int, int, class QStandardItem *)
func (this *QStandardItem) SetChild(row int, column int, item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem8setChildEiiPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:209
// index:1
// Public inline
// void setChild(int, class QStandardItem *)
func (this *QStandardItem) SetChild_1(row int, item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem8setChildEiPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:214
// index:0
// Public
// void insertRows(int, int)
func (this *QStandardItem) InsertRows(row int, count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10insertRowsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:215
// index:0
// Public
// void insertColumns(int, int)
func (this *QStandardItem) InsertColumns(column int, count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13insertColumnsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:217
// index:0
// Public
// void removeRow(int)
func (this *QStandardItem) RemoveRow(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9removeRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:218
// index:0
// Public
// void removeColumn(int)
func (this *QStandardItem) RemoveColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12removeColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:219
// index:0
// Public
// void removeRows(int, int)
func (this *QStandardItem) RemoveRows(row int, count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10removeRowsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:220
// index:0
// Public
// void removeColumns(int, int)
func (this *QStandardItem) RemoveColumns(column int, count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13removeColumnsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:225
// index:0
// Public inline
// void insertRow(int, class QStandardItem *)
func (this *QStandardItem) InsertRow(row int, item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9insertRowEiPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:226
// index:0
// Public inline
// void appendRow(class QStandardItem *)
func (this *QStandardItem) AppendRow(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9appendRowEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:228
// index:0
// Public
// QStandardItem * takeChild(int, int)
func (this *QStandardItem) TakeChild(row int, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9takeChildEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:229
// index:0
// Public
// QList<QStandardItem *> takeRow(int)
func (this *QStandardItem) TakeRow(row int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7takeRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:230
// index:0
// Public
// QList<QStandardItem *> takeColumn(int)
func (this *QStandardItem) TakeColumn(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10takeColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:232
// index:0
// Public
// void sortChildren(int, Qt::SortOrder)
func (this *QStandardItem) SortChildren(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12sortChildrenEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:234
// index:0
// Public virtual
// QStandardItem * clone()
func (this *QStandardItem) Clone() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5cloneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:237
// index:0
// Public virtual
// int type()
func (this *QStandardItem) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:240
// index:0
// Public virtual
// void read(class QDataStream &)
func (this *QStandardItem) Read(in *qtcore.QDataStream) {
	var convArg0 = in.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem4readER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:241
// index:0
// Public virtual
// void write(class QDataStream &)
func (this *QStandardItem) Write(out *qtcore.QDataStream) {
	var convArg0 = out.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5writeER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:251
// index:0
// Protected
// void emitDataChanged()
func (this *QStandardItem) EmitDataChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem15emitDataChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
