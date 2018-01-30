package qtgui

// /usr/include/qt/QtGui/qstandarditemmodel.h
// #include <qstandarditemmodel.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStandardItem) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQStandardItemFromPointer(cthis unsafe.Pointer) *QStandardItem {
	return &QStandardItem{&qtrt.CObject{cthis}}
}
func (*QStandardItem) NewFromPointer(cthis unsafe.Pointer) *QStandardItem {
	return NewQStandardItemFromPointer(cthis)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStandardItem()
func NewQStandardItem() *QStandardItem {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStandardItem(const QString &)
func NewQStandardItem_1(text *qtcore.QString) *QStandardItem {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:67
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QStandardItem(const QIcon &, const QString &)
func NewQStandardItem_2(icon *QIcon, text *qtcore.QString) *QStandardItem {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2ERK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:68
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QStandardItem(int, int)
func NewQStandardItem_3(rows int, columns int) *QStandardItem {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemC2Eii", ffiqt.FFI_TYPE_POINTER, rows, columns)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStandardItem()
func DeleteQStandardItem(*QStandardItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(int)
func (this *QStandardItem) Data(role int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4dataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setData(const QVariant &, int)
func (this *QStandardItem) SetData(value *qtcore.QVariant, role int) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setDataERK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString text()
func (this *QStandardItem) Text() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QStandardItem) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QStandardItem) Icon() *QIcon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4iconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QStandardItem) SetIcon(icon *QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toolTip()
func (this *QStandardItem) ToolTip() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem7toolTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)
func (this *QStandardItem) SetToolTip(toolTip *qtcore.QString) {
	var convArg0 = toolTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10setToolTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString statusTip()
func (this *QStandardItem) StatusTip() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem9statusTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setStatusTip(const QString &)
func (this *QStandardItem) SetStatusTip(statusTip *qtcore.QString) {
	var convArg0 = statusTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12setStatusTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString whatsThis()
func (this *QStandardItem) WhatsThis() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem9whatsThisEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)
func (this *QStandardItem) SetWhatsThis(whatsThis *qtcore.QString) {
	var convArg0 = whatsThis.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12setWhatsThisERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QStandardItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSizeHint(const QSize &)
func (this *QStandardItem) SetSizeHint(sizeHint *qtcore.QSize) {
	var convArg0 = sizeHint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setSizeHintERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QFont font()
func (this *QStandardItem) Font() *QFont /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QStandardItem) SetFont(font *QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Alignment textAlignment()
func (this *QStandardItem) TextAlignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem13textAlignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush background()
func (this *QStandardItem) Background() *QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10backgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)
func (this *QStandardItem) SetBackground(brush *QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush foreground()
func (this *QStandardItem) Foreground() *QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10foregroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setForeground(const QBrush &)
func (this *QStandardItem) SetForeground(brush *QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setForegroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::CheckState checkState()
func (this *QStandardItem) CheckState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10checkStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCheckState(Qt::CheckState)
func (this *QStandardItem) SetCheckState(checkState int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setCheckStateEN2Qt10CheckStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), checkState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString accessibleText()
func (this *QStandardItem) AccessibleText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem14accessibleTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:138
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAccessibleText(const QString &)
func (this *QStandardItem) SetAccessibleText(accessibleText *qtcore.QString) {
	var convArg0 = accessibleText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem17setAccessibleTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString accessibleDescription()
func (this *QStandardItem) AccessibleDescription() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem21accessibleDescriptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:143
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAccessibleDescription(const QString &)
func (this *QStandardItem) SetAccessibleDescription(accessibleDescription *qtcore.QString) {
	var convArg0 = accessibleDescription.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem24setAccessibleDescriptionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:145
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags()
func (this *QStandardItem) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(Qt::ItemFlags)
func (this *QStandardItem) SetFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem8setFlagsE6QFlagsIN2Qt8ItemFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:148
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEnabled()
func (this *QStandardItem) IsEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QStandardItem) SetEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:153
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEditable()
func (this *QStandardItem) IsEditable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10isEditableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEditable(_Bool)
func (this *QStandardItem) SetEditable(editable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setEditableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSelectable()
func (this *QStandardItem) IsSelectable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem12isSelectableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectable(_Bool)
func (this *QStandardItem) SetSelectable(selectable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13setSelectableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), selectable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:163
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isCheckable()
func (this *QStandardItem) IsCheckable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem11isCheckableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCheckable(_Bool)
func (this *QStandardItem) SetCheckable(checkable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12setCheckableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), checkable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:168
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAutoTristate()
func (this *QStandardItem) IsAutoTristate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem14isAutoTristateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoTristate(_Bool)
func (this *QStandardItem) SetAutoTristate(tristate bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem15setAutoTristateEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), tristate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:173
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUserTristate()
func (this *QStandardItem) IsUserTristate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem14isUserTristateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserTristate(_Bool)
func (this *QStandardItem) SetUserTristate(tristate bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem15setUserTristateEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), tristate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:179
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTristate()
func (this *QStandardItem) IsTristate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem10isTristateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTristate(_Bool)
func (this *QStandardItem) SetTristate(tristate bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setTristateEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), tristate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:184
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDragEnabled()
func (this *QStandardItem) IsDragEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem13isDragEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragEnabled(_Bool)
func (this *QStandardItem) SetDragEnabled(dragEnabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem14setDragEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dragEnabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:189
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDropEnabled()
func (this *QStandardItem) IsDropEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem13isDropEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDropEnabled(_Bool)
func (this *QStandardItem) SetDropEnabled(dropEnabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem14setDropEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dropEnabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:195
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * parent()
func (this *QStandardItem) Parent() *QStandardItem /*777 QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:196
// index:0
// Public Visibility=Default Availability=Available
// [4] int row()
func (this *QStandardItem) Row() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:197
// index:0
// Public Visibility=Default Availability=Available
// [4] int column()
func (this *QStandardItem) Column() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem6columnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:198
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex index()
func (this *QStandardItem) Index() *qtcore.QModelIndex /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5indexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:199
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItemModel * model()
func (this *QStandardItem) Model() *QStandardItemModel /*777 QStandardItemModel **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:201
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowCount()
func (this *QStandardItem) RowCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem8rowCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:202
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowCount(int)
func (this *QStandardItem) SetRowCount(rows int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem11setRowCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), rows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:203
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnCount()
func (this *QStandardItem) ColumnCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnCount(int)
func (this *QStandardItem) SetColumnCount(columns int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem14setColumnCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:206
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasChildren()
func (this *QStandardItem) HasChildren() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem11hasChildrenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:207
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * child(int, int)
func (this *QStandardItem) Child(row int, column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5childEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChild(int, int, QStandardItem *)
func (this *QStandardItem) SetChild(row int, column int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg2 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem8setChildEiiPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:209
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setChild(int, QStandardItem *)
func (this *QStandardItem) SetChild_1(row int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem8setChildEiPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertRows(int, int)
func (this *QStandardItem) InsertRows(row int, count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10insertRowsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertColumns(int, int)
func (this *QStandardItem) InsertColumns(column int, count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13insertColumnsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeRow(int)
func (this *QStandardItem) RemoveRow(row int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9removeRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeColumn(int)
func (this *QStandardItem) RemoveColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12removeColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeRows(int, int)
func (this *QStandardItem) RemoveRows(row int, count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem10removeRowsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeColumns(int, int)
func (this *QStandardItem) RemoveColumns(column int, count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem13removeColumnsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:225
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void insertRow(int, QStandardItem *)
func (this *QStandardItem) InsertRow(row int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9insertRowEiPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void appendRow(QStandardItem *)
func (this *QStandardItem) AppendRow(item *QStandardItem /*777 QStandardItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9appendRowEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:228
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * takeChild(int, int)
func (this *QStandardItem) TakeChild(row int, column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem9takeChildEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:232
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortChildren(int, Qt::SortOrder)
func (this *QStandardItem) SortChildren(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem12sortChildrenEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:234
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStandardItem * clone()
func (this *QStandardItem) Clone() *QStandardItem /*777 QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5cloneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:237
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type()
func (this *QStandardItem) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:240
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void read(QDataStream &)
func (this *QStandardItem) Read(in *qtcore.QDataStream) {
	var convArg0 = in.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem4readER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:241
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void write(QDataStream &)
func (this *QStandardItem) Write(out *qtcore.QDataStream) {
	var convArg0 = out.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStandardItem5writeER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:251
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void emitDataChanged()
func (this *QStandardItem) EmitDataChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStandardItem15emitDataChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QStandardItem__ItemType = int

const QStandardItem__Type QStandardItem__ItemType = 0
const QStandardItem__UserType QStandardItem__ItemType = 1000

//  body block end
