package qtgui

// /usr/include/qt/QtGui/qstandarditemmodel.h
// #include <qstandarditemmodel.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// QStandardItem & operator=(const QStandardItem &)
func (this *QStandardItem) InheritOperator_equal(f func(other *QStandardItem) unsafe.Pointer /*555*/) {
	qtrt.SetAllInheritCallback(this, "operator=", f)
}

// void emitDataChanged()
func (this *QStandardItem) InheritEmitDataChanged(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "emitDataChanged", f)
}

/*

 */
type QStandardItem struct {
	*qtrt.CObject
}
type QStandardItem_ITF interface {
	QStandardItem_PTR() *QStandardItem
}

func (ptr *QStandardItem) QStandardItem_PTR() *QStandardItem { return ptr }

func (this *QStandardItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStandardItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
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

/*

 */
func NewQStandardItem() *QStandardItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItemC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStandardItem)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStandardItem(const QString &)

/*

 */
func NewQStandardItem_1(text string) *QStandardItem {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItemC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStandardItem)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:67
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QStandardItem(const QIcon &, const QString &)

/*

 */
func NewQStandardItem_2(icon QIcon_ITF, text string) *QStandardItem {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItemC2ERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStandardItem)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:68
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QStandardItem(int, int)

/*

 */
func NewQStandardItem_3(rows int, columns int) *QStandardItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItemC2Eii", qtrt.FFI_TYPE_POINTER, rows, columns)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStandardItem)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:68
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QStandardItem(int, int)

/*

 */
func NewQStandardItem_3_(rows int) *QStandardItem {
	// arg: 1, int=Int, =Invalid, , Invalid
	columns := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItemC2Eii", qtrt.FFI_TYPE_POINTER, rows, columns)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStandardItem)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStandardItem()

/*

 */
func DeleteQStandardItem(this *QStandardItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(int) const

/*
Reimplemented from QAbstractItemModel::data().

See also setData().
*/
func (this *QStandardItem) Data(role int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem4dataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(int) const

/*
Reimplemented from QAbstractItemModel::data().

See also setData().
*/
func (this *QStandardItem) Data__() *qtcore.QVariant /*123*/ {
	// arg: 0, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::UserRole + 1*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem4dataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setData(const QVariant &, int)

/*
Reimplemented from QAbstractItemModel::setData().

See also data().
*/
func (this *QStandardItem) SetData(value qtcore.QVariant_ITF, role int) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg0 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem7setDataERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setData(const QVariant &, int)

/*
Reimplemented from QAbstractItemModel::setData().

See also data().
*/
func (this *QStandardItem) SetData__(value qtcore.QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg0 = value.QVariant_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::UserRole + 1*/
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem7setDataERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString text() const

/*

 */
func (this *QStandardItem) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*

 */
func (this *QStandardItem) SetText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QIcon icon() const

/*

 */
func (this *QStandardItem) Icon() *QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)

/*

 */
func (this *QStandardItem) SetIcon(icon QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toolTip() const

/*

 */
func (this *QStandardItem) ToolTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)

/*

 */
func (this *QStandardItem) SetToolTip(toolTip string) {
	var tmpArg0 = qtcore.NewQString_5(toolTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString statusTip() const

/*

 */
func (this *QStandardItem) StatusTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem9statusTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setStatusTip(const QString &)

/*

 */
func (this *QStandardItem) SetStatusTip(statusTip string) {
	var tmpArg0 = qtcore.NewQString_5(statusTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem12setStatusTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString whatsThis() const

/*

 */
func (this *QStandardItem) WhatsThis() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem9whatsThisEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)

/*

 */
func (this *QStandardItem) SetWhatsThis(whatsThis string) {
	var tmpArg0 = qtcore.NewQString_5(whatsThis)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem12setWhatsThisERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*

 */
func (this *QStandardItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSizeHint(const QSize &)

/*

 */
func (this *QStandardItem) SetSizeHint(sizeHint qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if sizeHint != nil && sizeHint.QSize_PTR() != nil {
		convArg0 = sizeHint.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem11setSizeHintERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QFont font() const

/*

 */
func (this *QStandardItem) Font() *QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)

/*

 */
func (this *QStandardItem) SetFont(font QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Alignment textAlignment() const

/*

 */
func (this *QStandardItem) TextAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem13textAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextAlignment(Qt::Alignment)

/*

 */
func (this *QStandardItem) SetTextAlignment(textAlignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem16setTextAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), textAlignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush background() const

/*

 */
func (this *QStandardItem) Background() *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem10backgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)

/*

 */
func (this *QStandardItem) SetBackground(brush QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem13setBackgroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush foreground() const

/*

 */
func (this *QStandardItem) Foreground() *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem10foregroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setForeground(const QBrush &)

/*

 */
func (this *QStandardItem) SetForeground(brush QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem13setForegroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::CheckState checkState() const

/*

 */
func (this *QStandardItem) CheckState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem10checkStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCheckState(Qt::CheckState)

/*

 */
func (this *QStandardItem) SetCheckState(checkState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem13setCheckStateEN2Qt10CheckStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checkState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString accessibleText() const

/*

 */
func (this *QStandardItem) AccessibleText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem14accessibleTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:138
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAccessibleText(const QString &)

/*

 */
func (this *QStandardItem) SetAccessibleText(accessibleText string) {
	var tmpArg0 = qtcore.NewQString_5(accessibleText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem17setAccessibleTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString accessibleDescription() const

/*

 */
func (this *QStandardItem) AccessibleDescription() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem21accessibleDescriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:143
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAccessibleDescription(const QString &)

/*

 */
func (this *QStandardItem) SetAccessibleDescription(accessibleDescription string) {
	var tmpArg0 = qtcore.NewQString_5(accessibleDescription)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem24setAccessibleDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:145
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags() const

/*
Reimplemented from QAbstractItemModel::flags().
*/
func (this *QStandardItem) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(Qt::ItemFlags)

/*

 */
func (this *QStandardItem) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem8setFlagsE6QFlagsIN2Qt8ItemFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:148
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEnabled() const

/*

 */
func (this *QStandardItem) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(bool)

/*

 */
func (this *QStandardItem) SetEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:153
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEditable() const

/*

 */
func (this *QStandardItem) IsEditable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem10isEditableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEditable(bool)

/*

 */
func (this *QStandardItem) SetEditable(editable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem11setEditableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), editable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSelectable() const

/*

 */
func (this *QStandardItem) IsSelectable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem12isSelectableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectable(bool)

/*

 */
func (this *QStandardItem) SetSelectable(selectable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem13setSelectableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), selectable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:163
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isCheckable() const

/*

 */
func (this *QStandardItem) IsCheckable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem11isCheckableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCheckable(bool)

/*

 */
func (this *QStandardItem) SetCheckable(checkable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem12setCheckableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checkable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:168
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAutoTristate() const

/*

 */
func (this *QStandardItem) IsAutoTristate() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem14isAutoTristateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoTristate(bool)

/*

 */
func (this *QStandardItem) SetAutoTristate(tristate bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem15setAutoTristateEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tristate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:173
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUserTristate() const

/*

 */
func (this *QStandardItem) IsUserTristate() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem14isUserTristateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserTristate(bool)

/*

 */
func (this *QStandardItem) SetUserTristate(tristate bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem15setUserTristateEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tristate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:179
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTristate() const

/*

 */
func (this *QStandardItem) IsTristate() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem10isTristateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTristate(bool)

/*

 */
func (this *QStandardItem) SetTristate(tristate bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem11setTristateEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tristate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:184
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDragEnabled() const

/*

 */
func (this *QStandardItem) IsDragEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem13isDragEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragEnabled(bool)

/*

 */
func (this *QStandardItem) SetDragEnabled(dragEnabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem14setDragEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dragEnabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:189
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDropEnabled() const

/*

 */
func (this *QStandardItem) IsDropEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem13isDropEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDropEnabled(bool)

/*

 */
func (this *QStandardItem) SetDropEnabled(dropEnabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem14setDropEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dropEnabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:195
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * parent() const

/*
Reimplemented from QAbstractItemModel::parent().
*/
func (this *QStandardItem) Parent() *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:196
// index:0
// Public Visibility=Default Availability=Available
// [4] int row() const

/*

 */
func (this *QStandardItem) Row() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem3rowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:197
// index:0
// Public Visibility=Default Availability=Available
// [4] int column() const

/*

 */
func (this *QStandardItem) Column() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem6columnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:198
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex index() const

/*
Reimplemented from QAbstractItemModel::index().
*/
func (this *QStandardItem) Index() *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem5indexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:199
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItemModel * model() const

/*

 */
func (this *QStandardItem) Model() *QStandardItemModel /*777 QStandardItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:201
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowCount() const

/*
Reimplemented from QAbstractItemModel::rowCount().

See also setRowCount().
*/
func (this *QStandardItem) RowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem8rowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:202
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowCount(int)

/*
Sets the number of rows in this model to rows. If this is less than rowCount(), the data in the unwanted rows is discarded.

This function was introduced in  Qt 4.2.

See also rowCount() and setColumnCount().
*/
func (this *QStandardItem) SetRowCount(rows int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem11setRowCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rows)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:203
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnCount() const

/*
Reimplemented from QAbstractItemModel::columnCount().

See also setColumnCount().
*/
func (this *QStandardItem) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnCount(int)

/*
Sets the number of columns in this model to columns. If this is less than columnCount(), the data in the unwanted columns is discarded.

This function was introduced in  Qt 4.2.

See also columnCount() and setRowCount().
*/
func (this *QStandardItem) SetColumnCount(columns int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem14setColumnCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:206
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasChildren() const

/*
Reimplemented from QAbstractItemModel::hasChildren().
*/
func (this *QStandardItem) HasChildren() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem11hasChildrenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:207
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * child(int, int) const

/*

 */
func (this *QStandardItem) Child(row int, column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem5childEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:207
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * child(int, int) const

/*

 */
func (this *QStandardItem) Child__(row int) *QStandardItem /*777 QStandardItem **/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem5childEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChild(int, int, QStandardItem *)

/*

 */
func (this *QStandardItem) SetChild(row int, column int, item QStandardItem_ITF /*777 QStandardItem **/) {
	var convArg2 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg2 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem8setChildEiiPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:209
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setChild(int, QStandardItem *)

/*

 */
func (this *QStandardItem) SetChild_1(row int, item QStandardItem_ITF /*777 QStandardItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg1 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem8setChildEiPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertRows(int, int)

/*
Reimplemented from QAbstractItemModel::insertRows().
*/
func (this *QStandardItem) InsertRows(row int, count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem10insertRowsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertColumns(int, int)

/*
Reimplemented from QAbstractItemModel::insertColumns().
*/
func (this *QStandardItem) InsertColumns(column int, count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem13insertColumnsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeRow(int)

/*

 */
func (this *QStandardItem) RemoveRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem9removeRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeColumn(int)

/*

 */
func (this *QStandardItem) RemoveColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem12removeColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeRows(int, int)

/*
Reimplemented from QAbstractItemModel::removeRows().
*/
func (this *QStandardItem) RemoveRows(row int, count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem10removeRowsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeColumns(int, int)

/*
Reimplemented from QAbstractItemModel::removeColumns().
*/
func (this *QStandardItem) RemoveColumns(column int, count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem13removeColumnsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:225
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void insertRow(int, QStandardItem *)

/*
Inserts a row at row containing items. If necessary, the column count is increased to the size of items.

This function was introduced in  Qt 4.2.

See also takeRow(), appendRow(), and insertColumn().
*/
func (this *QStandardItem) InsertRow(row int, item QStandardItem_ITF /*777 QStandardItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg1 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem9insertRowEiPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void appendRow(QStandardItem *)

/*
Appends a row containing items. If necessary, the column count is increased to the size of items.

This function was introduced in  Qt 4.2.

See also insertRow() and appendColumn().
*/
func (this *QStandardItem) AppendRow(item QStandardItem_ITF /*777 QStandardItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg0 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem9appendRowEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:228
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * takeChild(int, int)

/*

 */
func (this *QStandardItem) TakeChild(row int, column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem9takeChildEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:228
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * takeChild(int, int)

/*

 */
func (this *QStandardItem) TakeChild__(row int) *QStandardItem /*777 QStandardItem **/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem9takeChildEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:232
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortChildren(int, Qt::SortOrder)

/*

 */
func (this *QStandardItem) SortChildren(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem12sortChildrenEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:232
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortChildren(int, Qt::SortOrder)

/*

 */
func (this *QStandardItem) SortChildren__(column int) {
	// arg: 1, Qt::SortOrder=Elaborated, Qt::SortOrder=Enum, , Invalid
	order := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem12sortChildrenEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:234
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStandardItem * clone() const

/*

 */
func (this *QStandardItem) Clone() *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem5cloneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:237
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type() const

/*

 */
func (this *QStandardItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:240
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void read(QDataStream &)

/*

 */
func (this *QStandardItem) Read(in qtcore.QDataStream_ITF) {
	var convArg0 unsafe.Pointer
	if in != nil && in.QDataStream_PTR() != nil {
		convArg0 = in.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem4readER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:241
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void write(QDataStream &) const

/*

 */
func (this *QStandardItem) Write(out qtcore.QDataStream_ITF) {
	var convArg0 unsafe.Pointer
	if out != nil && out.QDataStream_PTR() != nil {
		convArg0 = out.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem5writeER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:243
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool operator<(const QStandardItem &) const

/*

 */
func (this *QStandardItem) Operator_less_than(other QStandardItem_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStandardItem_PTR() != nil {
		convArg0 = other.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItemltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:248
// index:0
// Protected Visibility=Default Availability=Available
// [16] QStandardItem & operator=(const QStandardItem &)

/*

 */
func (this *QStandardItem) Operator_equal(other QStandardItem_ITF) *QStandardItem {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStandardItem_PTR() != nil {
		convArg0 = other.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItemaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStandardItem)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:251
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void emitDataChanged()

/*

 */
func (this *QStandardItem) EmitDataChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem15emitDataChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QStandardItem__ItemType = int

//
const QStandardItem__Type QStandardItem__ItemType = 0

//
const QStandardItem__UserType QStandardItem__ItemType = 1000

func (this *QStandardItem) ItemTypeItemName(val int) string {
	switch val {
	case QStandardItem__Type: // 0
		return "Type"
	case QStandardItem__UserType: // 1000
		return "UserType"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStandardItem_ItemTypeItemName(val int) string {
	var nilthis *QStandardItem
	return nilthis.ItemTypeItemName(val)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
