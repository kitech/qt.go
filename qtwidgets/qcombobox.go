package qtwidgets

// /usr/include/qt/QtWidgets/qcombobox.h
// #include <qcombobox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void focusInEvent(QFocusEvent *)
func (this *QComboBox) InheritFocusInEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QComboBox) InheritFocusOutEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void changeEvent(QEvent *)
func (this *QComboBox) InheritChangeEvent(f func(e *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QComboBox) InheritResizeEvent(f func(e *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QComboBox) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QComboBox) InheritShowEvent(f func(e *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QComboBox) InheritHideEvent(f func(e *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QComboBox) InheritMousePressEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QComboBox) InheritMouseReleaseEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QComboBox) InheritKeyPressEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QComboBox) InheritKeyReleaseEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QComboBox) InheritWheelEvent(f func(e *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void contextMenuEvent(QContextMenuEvent *)
func (this *QComboBox) InheritContextMenuEvent(f func(e *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void inputMethodEvent(QInputMethodEvent *)
func (this *QComboBox) InheritInputMethodEvent(f func(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void initStyleOption(QStyleOptionComboBox *)
func (this *QComboBox) InheritInitStyleOption(f func(option *QStyleOptionComboBox /*777 QStyleOptionComboBox **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QComboBox struct {
	*QWidget
}
type QComboBox_ITF interface {
	QWidget_ITF
	QComboBox_PTR() *QComboBox
}

func (ptr *QComboBox) QComboBox_PTR() *QComboBox { return ptr }

func (this *QComboBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QComboBox) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQComboBoxFromPointer(cthis unsafe.Pointer) *QComboBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QComboBox{bcthis0}
}
func (*QComboBox) NewFromPointer(cthis unsafe.Pointer) *QComboBox {
	return NewQComboBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qcombobox.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QComboBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcombobox.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QComboBox(QWidget *)

/*
Constructs a combobox with the given parent, using the default model QStandardItemModel.
*/
func (*QComboBox) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QComboBox {
	return NewQComboBox(parent)
}
func NewQComboBox(parent QWidget_ITF /*777 QWidget **/) *QComboBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQComboBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QComboBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qcombobox.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QComboBox(QWidget *)

/*
Constructs a combobox with the given parent, using the default model QStandardItemModel.
*/
func (*QComboBox) NewForInherit__() *QComboBox {
	return NewQComboBox__()
}
func NewQComboBox__() *QComboBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQComboBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QComboBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qcombobox.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QComboBox()

/*

 */
func DeleteQComboBox(this *QComboBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qcombobox.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxVisibleItems() const

/*

 */
func (this *QComboBox) MaxVisibleItems() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox15maxVisibleItemsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxVisibleItems(int)

/*

 */
func (this *QComboBox) SetMaxVisibleItems(maxItems int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox18setMaxVisibleItemsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxItems)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*

 */
func (this *QComboBox) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxCount(int)

/*

 */
func (this *QComboBox) SetMaxCount(max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11setMaxCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxCount() const

/*

 */
func (this *QComboBox) MaxCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8maxCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoCompletion() const

/*

 */
func (this *QComboBox) AutoCompletion() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox14autoCompletionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoCompletion(bool)

/*

 */
func (this *QComboBox) SetAutoCompletion(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox17setAutoCompletionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity autoCompletionCaseSensitivity() const

/*

 */
func (this *QComboBox) AutoCompletionCaseSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox29autoCompletionCaseSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoCompletionCaseSensitivity(Qt::CaseSensitivity)

/*

 */
func (this *QComboBox) SetAutoCompletionCaseSensitivity(sensitivity int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox32setAutoCompletionCaseSensitivityEN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sensitivity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool duplicatesEnabled() const

/*

 */
func (this *QComboBox) DuplicatesEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox17duplicatesEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDuplicatesEnabled(bool)

/*

 */
func (this *QComboBox) SetDuplicatesEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox20setDuplicatesEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrame(bool)

/*

 */
func (this *QComboBox) SetFrame(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox8setFrameEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFrame() const

/*

 */
func (this *QComboBox) HasFrame() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8hasFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int findText(const QString &, Qt::MatchFlags) const

/*
Returns the index of the item containing the given text; otherwise returns -1.

The flags specify how the items in the combobox are searched.
*/
func (this *QComboBox) FindText(text string, flags int) int {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8findTextERK7QString6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int findText(const QString &, Qt::MatchFlags) const

/*
Returns the index of the item containing the given text; otherwise returns -1.

The flags specify how the items in the combobox are searched.
*/
func (this *QComboBox) FindText__(text string) int {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::MatchFlags=Elaborated, Qt::MatchFlags=Typedef, QFlags<Qt::MatchFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8findTextERK7QString6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int findData(const QVariant &, int, Qt::MatchFlags) const

/*
Returns the index of the item containing the given data for the given role; otherwise returns -1.

The flags specify how the items in the combobox are searched.
*/
func (this *QComboBox) FindData(data qtcore.QVariant_ITF, role int, flags int) int {
	var convArg0 unsafe.Pointer
	if data != nil && data.QVariant_PTR() != nil {
		convArg0 = data.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8findDataERK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int findData(const QVariant &, int, Qt::MatchFlags) const

/*
Returns the index of the item containing the given data for the given role; otherwise returns -1.

The flags specify how the items in the combobox are searched.
*/
func (this *QComboBox) FindData__(data qtcore.QVariant_ITF) int {
	var convArg0 unsafe.Pointer
	if data != nil && data.QVariant_PTR() != nil {
		convArg0 = data.QVariant_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::UserRole*/
	// arg: 2, Qt::MatchFlags=Elaborated, Qt::MatchFlags=Typedef, QFlags<Qt::MatchFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8findDataERK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int findData(const QVariant &, int, Qt::MatchFlags) const

/*
Returns the index of the item containing the given data for the given role; otherwise returns -1.

The flags specify how the items in the combobox are searched.
*/
func (this *QComboBox) FindData__1(data qtcore.QVariant_ITF, role int) int {
	var convArg0 unsafe.Pointer
	if data != nil && data.QVariant_PTR() != nil {
		convArg0 = data.QVariant_PTR().GetCthis()
	}
	// arg: 2, Qt::MatchFlags=Elaborated, Qt::MatchFlags=Typedef, QFlags<Qt::MatchFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8findDataERK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:126
// index:0
// Public Visibility=Default Availability=Available
// [4] QComboBox::InsertPolicy insertPolicy() const

/*

 */
func (this *QComboBox) InsertPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox12insertPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInsertPolicy(QComboBox::InsertPolicy)

/*

 */
func (this *QComboBox) SetInsertPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox15setInsertPolicyENS_12InsertPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] QComboBox::SizeAdjustPolicy sizeAdjustPolicy() const

/*

 */
func (this *QComboBox) SizeAdjustPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox16sizeAdjustPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeAdjustPolicy(QComboBox::SizeAdjustPolicy)

/*

 */
func (this *QComboBox) SetSizeAdjustPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox19setSizeAdjustPolicyENS_16SizeAdjustPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimumContentsLength() const

/*

 */
func (this *QComboBox) MinimumContentsLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox21minimumContentsLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumContentsLength(int)

/*

 */
func (this *QComboBox) SetMinimumContentsLength(characters int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox24setMinimumContentsLengthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), characters)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize() const

/*

 */
func (this *QComboBox) IconSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8iconSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)

/*

 */
func (this *QComboBox) SetIconSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEditable() const

/*

 */
func (this *QComboBox) IsEditable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox10isEditableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEditable(bool)

/*

 */
func (this *QComboBox) SetEditable(editable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11setEditableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), editable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineEdit(QLineEdit *)

/*
Sets the line edit to use instead of the current line edit widget.

The combo box takes ownership of the line edit.

See also lineEdit().
*/
func (this *QComboBox) SetLineEdit(edit QLineEdit_ITF /*777 QLineEdit **/) {
	var convArg0 unsafe.Pointer
	if edit != nil && edit.QLineEdit_PTR() != nil {
		convArg0 = edit.QLineEdit_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11setLineEditEP9QLineEdit", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:147
// index:0
// Public Visibility=Default Availability=Available
// [8] QLineEdit * lineEdit() const

/*
Returns the line edit used to edit items in the combobox, or 0 if there is no line edit.

Only editable combo boxes have a line edit.

See also setLineEdit().
*/
func (this *QComboBox) LineEdit() *QLineEdit /*777 QLineEdit **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8lineEditEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcombobox.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValidator(const QValidator *)

/*
Sets the validator to use instead of the current validator.

Note: The validator is removed when the editable property becomes false.

See also validator().
*/
func (this *QComboBox) SetValidator(v qtgui.QValidator_ITF /*777 const QValidator **/) {
	var convArg0 unsafe.Pointer
	if v != nil && v.QValidator_PTR() != nil {
		convArg0 = v.QValidator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox12setValidatorEPK10QValidator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:150
// index:0
// Public Visibility=Default Availability=Available
// [8] const QValidator * validator() const

/*
Returns the validator that is used to constrain text input for the combobox.

See also setValidator() and editable.
*/
func (this *QComboBox) Validator() *qtgui.QValidator /*777 const QValidator **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox9validatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQValidatorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcombobox.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompleter(QCompleter *)

/*
Sets the completer to use instead of the current completer. If completer is 0, auto completion is disabled.

By default, for an editable combo box, a QCompleter that performs case insensitive inline completion is automatically created.

Note: The completer is removed when the editable property becomes false.

This function was introduced in  Qt 4.2.

See also completer().
*/
func (this *QComboBox) SetCompleter(c QCompleter_ITF /*777 QCompleter **/) {
	var convArg0 unsafe.Pointer
	if c != nil && c.QCompleter_PTR() != nil {
		convArg0 = c.QCompleter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox12setCompleterEP10QCompleter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QCompleter * completer() const

/*
Returns the completer that is used to auto complete text input for the combobox.

This function was introduced in  Qt 4.2.

See also setCompleter() and editable.
*/
func (this *QComboBox) Completer() *QCompleter /*777 QCompleter **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox9completerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcombobox.h:158
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegate() const

/*
Returns the item delegate used by the popup list view.

See also setItemDelegate().
*/
func (this *QComboBox) ItemDelegate() *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox12itemDelegateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcombobox.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemDelegate(QAbstractItemDelegate *)

/*
Sets the item delegate for the popup list view. The combobox takes ownership of the delegate.

Warning: You should not share the same instance of a delegate between comboboxes, widget mappers or views. Doing so can cause incorrect or unintuitive editing behavior since each view connected to a given delegate may receive the closeEditor() signal, and attempt to access, modify or close an editor that has already been closed.

See also itemDelegate().
*/
func (this *QComboBox) SetItemDelegate(delegate QAbstractItemDelegate_ITF /*777 QAbstractItemDelegate **/) {
	var convArg0 unsafe.Pointer
	if delegate != nil && delegate.QAbstractItemDelegate_PTR() != nil {
		convArg0 = delegate.QAbstractItemDelegate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox15setItemDelegateEP21QAbstractItemDelegate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:161
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemModel * model() const

/*
Returns the model used by the combobox.

See also setModel().
*/
func (this *QComboBox) Model() *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcombobox.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)

/*
Sets the model to be model. model must not be 0. If you want to clear the contents of a model, call clear().

See also model() and clear().
*/
func (this *QComboBox) SetModel(model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox8setModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:164
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex rootModelIndex() const

/*
Returns the root model item index for the items in the combobox.

See also setRootModelIndex().
*/
func (this *QComboBox) RootModelIndex() *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox14rootModelIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRootModelIndex(const QModelIndex &)

/*
Sets the root model item index for the items in the combobox.

See also rootModelIndex().
*/
func (this *QComboBox) SetRootModelIndex(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox17setRootModelIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] int modelColumn() const

/*

 */
func (this *QComboBox) ModelColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox11modelColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModelColumn(int)

/*

 */
func (this *QComboBox) SetModelColumn(visibleColumn int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox14setModelColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visibleColumn)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:170
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex() const

/*

 */
func (this *QComboBox) CurrentIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] QString currentText() const

/*

 */
func (this *QComboBox) CurrentText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox11currentTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qcombobox.h:172
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant currentData(int) const

/*

 */
func (this *QComboBox) CurrentData(role int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox11currentDataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:172
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant currentData(int) const

/*

 */
func (this *QComboBox) CurrentData__() *qtcore.QVariant /*123*/ {
	// arg: 0, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::UserRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox11currentDataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:174
// index:0
// Public Visibility=Default Availability=Available
// [8] QString itemText(int) const

/*
Returns the text for the given index in the combobox.

See also setItemText().
*/
func (this *QComboBox) ItemText(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8itemTextEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qcombobox.h:175
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon itemIcon(int) const

/*
Returns the icon for the given index in the combobox.

See also setItemIcon().
*/
func (this *QComboBox) ItemIcon(index int) *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8itemIconEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:176
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant itemData(int, int) const

/*
Returns the data for the given role in the given index in the combobox, or QVariant::Invalid if there is no data for this role.

See also setItemData().
*/
func (this *QComboBox) ItemData(index int, role int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8itemDataEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:176
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant itemData(int, int) const

/*
Returns the data for the given role in the given index in the combobox, or QVariant::Invalid if there is no data for this role.

See also setItemData().
*/
func (this *QComboBox) ItemData__(index int) *qtcore.QVariant /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::UserRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8itemDataEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:178
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(const QString &, const QVariant &)

/*
Adds an item to the combobox with the given text, and containing the specified userData (stored in the Qt::UserRole). The item is appended to the list of existing items.
*/
func (this *QComboBox) AddItem(text string, userData qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if userData != nil && userData.QVariant_PTR() != nil {
		convArg1 = userData.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox7addItemERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:178
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(const QString &, const QVariant &)

/*
Adds an item to the combobox with the given text, and containing the specified userData (stored in the Qt::UserRole). The item is appended to the list of existing items.
*/
func (this *QComboBox) AddItem__(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QVariant &=LValueReference, QVariant=Record, , Invalid
	var convArg1 = qtcore.NewQVariant()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox7addItemERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:179
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(const QIcon &, const QString &, const QVariant &)

/*
Adds an item to the combobox with the given text, and containing the specified userData (stored in the Qt::UserRole). The item is appended to the list of existing items.
*/
func (this *QComboBox) AddItem_1(icon qtgui.QIcon_ITF, text string, userData qtcore.QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if userData != nil && userData.QVariant_PTR() != nil {
		convArg2 = userData.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:179
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(const QIcon &, const QString &, const QVariant &)

/*
Adds an item to the combobox with the given text, and containing the specified userData (stored in the Qt::UserRole). The item is appended to the list of existing items.
*/
func (this *QComboBox) AddItem_1_(icon qtgui.QIcon_ITF, text string) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QVariant &=LValueReference, QVariant=Record, , Invalid
	var convArg2 = qtcore.NewQVariant()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:181
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addItems(const QStringList &)

/*
Adds each of the strings in the given texts to the combobox. Each item is appended to the list of existing items in turn.
*/
func (this *QComboBox) AddItems(texts qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if texts != nil && texts.QStringList_PTR() != nil {
		convArg0 = texts.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox8addItemsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:184
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void insertItem(int, const QString &, const QVariant &)

/*
Inserts the text and userData (stored in the Qt::UserRole) into the combobox at the given index.

If the index is equal to or higher than the total number of items, the new item is appended to the list of existing items. If the index is zero or negative, the new item is prepended to the list of existing items.

See also insertItems().
*/
func (this *QComboBox) InsertItem(index int, text string, userData qtcore.QVariant_ITF) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if userData != nil && userData.QVariant_PTR() != nil {
		convArg2 = userData.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:184
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void insertItem(int, const QString &, const QVariant &)

/*
Inserts the text and userData (stored in the Qt::UserRole) into the combobox at the given index.

If the index is equal to or higher than the total number of items, the new item is appended to the list of existing items. If the index is zero or negative, the new item is prepended to the list of existing items.

See also insertItems().
*/
func (this *QComboBox) InsertItem__(index int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QVariant &=LValueReference, QVariant=Record, , Invalid
	var convArg2 = qtcore.NewQVariant()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:185
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertItem(int, const QIcon &, const QString &, const QVariant &)

/*
Inserts the text and userData (stored in the Qt::UserRole) into the combobox at the given index.

If the index is equal to or higher than the total number of items, the new item is appended to the list of existing items. If the index is zero or negative, the new item is prepended to the list of existing items.

See also insertItems().
*/
func (this *QComboBox) InsertItem_1(index int, icon qtgui.QIcon_ITF, text string, userData qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg1 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 unsafe.Pointer
	if userData != nil && userData.QVariant_PTR() != nil {
		convArg3 = userData.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:185
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertItem(int, const QIcon &, const QString &, const QVariant &)

/*
Inserts the text and userData (stored in the Qt::UserRole) into the combobox at the given index.

If the index is equal to or higher than the total number of items, the new item is appended to the list of existing items. If the index is zero or negative, the new item is prepended to the list of existing items.

See also insertItems().
*/
func (this *QComboBox) InsertItem_1_(index int, icon qtgui.QIcon_ITF, text string) {
	var convArg1 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg1 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, const QVariant &=LValueReference, QVariant=Record, , Invalid
	var convArg3 = qtcore.NewQVariant()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertItems(int, const QStringList &)

/*
Inserts the strings from the list into the combobox as separate items, starting at the index specified.

If the index is equal to or higher than the total number of items, the new items are appended to the list of existing items. If the index is zero or negative, the new items are prepended to the list of existing items.

See also insertItem().
*/
func (this *QComboBox) InsertItems(index int, texts qtcore.QStringList_ITF) {
	var convArg1 unsafe.Pointer
	if texts != nil && texts.QStringList_PTR() != nil {
		convArg1 = texts.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11insertItemsEiRK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertSeparator(int)

/*
Inserts a separator item into the combobox at the given index.

If the index is equal to or higher than the total number of items, the new item is appended to the list of existing items. If the index is zero or negative, the new item is prepended to the list of existing items.

This function was introduced in  Qt 4.4.

See also insertItem().
*/
func (this *QComboBox) InsertSeparator(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox15insertSeparatorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(int)

/*
Removes the item at the given index from the combobox. This will update the current index if the index is removed.

This function does nothing if index is out of range.
*/
func (this *QComboBox) RemoveItem(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox10removeItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemText(int, const QString &)

/*
Sets the text for the item on the given index in the combobox.

See also itemText().
*/
func (this *QComboBox) SetItemText(index int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11setItemTextEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemIcon(int, const QIcon &)

/*
Sets the icon for the item on the given index in the combobox.

See also itemIcon().
*/
func (this *QComboBox) SetItemIcon(index int, icon qtgui.QIcon_ITF) {
	var convArg1 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg1 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11setItemIconEiRK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemData(int, const QVariant &, int)

/*
Sets the data role for the item on the given index in the combobox to the specified value.

See also itemData().
*/
func (this *QComboBox) SetItemData(index int, value qtcore.QVariant_ITF, role int) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11setItemDataEiRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemData(int, const QVariant &, int)

/*
Sets the data role for the item on the given index in the combobox to the specified value.

See also itemData().
*/
func (this *QComboBox) SetItemData__(index int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::UserRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11setItemDataEiRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:196
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemView * view() const

/*
Returns the list view used for the combobox popup.

See also setView().
*/
func (this *QComboBox) View() *QAbstractItemView /*777 QAbstractItemView **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox4viewEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcombobox.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setView(QAbstractItemView *)

/*
Sets the view to be used in the combobox popup to the given itemView. The combobox takes ownership of the view.

Note: If you want to use the convenience views (like QListWidget, QTableWidget or QTreeWidget), make sure to call setModel() on the combobox with the convenience widgets model before calling this function.

See also view().
*/
func (this *QComboBox) SetView(itemView QAbstractItemView_ITF /*777 QAbstractItemView **/) {
	var convArg0 unsafe.Pointer
	if itemView != nil && itemView.QAbstractItemView_PTR() != nil {
		convArg0 = itemView.QAbstractItemView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox7setViewEP17QAbstractItemView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:199
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().

This implementation caches the size hint to avoid resizing when the contents change dynamically. To invalidate the cached value change the sizeAdjustPolicy.
*/
func (this *QComboBox) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:200
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QComboBox) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:202
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void showPopup()

/*
Displays the list of items in the combobox. If the list is empty then the no items will be shown.

If you reimplement this function to show a custom pop-up, make sure you call hidePopup() to reset the internal state.

See also hidePopup().
*/
func (this *QComboBox) ShowPopup() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox9showPopupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:203
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void hidePopup()

/*
Hides the list of items in the combobox if it is currently visible and resets the internal state, so that if the custom pop-up was shown inside the reimplemented showPopup(), then you also need to reimplement the hidePopup() function to hide your custom pop-up and call the base class implementation to reset the internal state whenever your custom pop-up widget is hidden.

See also showPopup().
*/
func (this *QComboBox) HidePopup() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox9hidePopupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:205
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QComboBox) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:206
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
Reimplemented from QWidget::inputMethodQuery().
*/
func (this *QComboBox) InputMethodQuery(arg0 int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:207
// index:1
// Public Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery, const QVariant &) const

/*
Reimplemented from QWidget::inputMethodQuery().
*/
func (this *QComboBox) InputMethodQuery_1(query int, argument qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if argument != nil && argument.QVariant_PTR() != nil {
		convArg1 = argument.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox16inputMethodQueryEN2Qt16InputMethodQueryERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:210
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the combobox, removing all items.

Note: If you have set an external model on the combobox this model will still be cleared when calling this function.
*/
func (this *QComboBox) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearEditText()

/*
Clears the contents of the line edit used for editing in the combobox.
*/
func (this *QComboBox) ClearEditText() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox13clearEditTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:212
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEditText(const QString &)

/*
Sets the text in the combobox's text edit.
*/
func (this *QComboBox) SetEditText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11setEditTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)

/*

 */
func (this *QComboBox) SetCurrentIndex(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox15setCurrentIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentText(const QString &)

/*

 */
func (this *QComboBox) SetCurrentText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox14setCurrentTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editTextChanged(const QString &)

/*
This signal is emitted when the text in the combobox's line edit widget is changed. The new text is specified by text.
*/
func (this *QComboBox) EditTextChanged(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox15editTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated(int)

/*
This signal is sent when the user chooses an item in the combobox. The item's index is passed. Note that this signal is sent even when the choice is not changed. If you need to know when the choice actually changes, use signal currentIndexChanged().

Note: Signal activated is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(comboBox, QOverload<int>::of(&QComboBox::activated),
      [=](int index){ /-* ... *-/ });
*/
func (this *QComboBox) Activated(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox9activatedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:219
// index:1
// Public Visibility=Default Availability=Available
// [-2] void activated(const QString &)

/*
This signal is sent when the user chooses an item in the combobox. The item's index is passed. Note that this signal is sent even when the choice is not changed. If you need to know when the choice actually changes, use signal currentIndexChanged().

Note: Signal activated is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(comboBox, QOverload<int>::of(&QComboBox::activated),
      [=](int index){ /-* ... *-/ });
*/
func (this *QComboBox) Activated_1(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox9activatedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void highlighted(int)

/*
This signal is sent when an item in the combobox popup list is highlighted by the user. The item's index is passed.

Note: Signal highlighted is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(comboBox, QOverload<int>::of(&QComboBox::highlighted),
      [=](int index){ /-* ... *-/ });
*/
func (this *QComboBox) Highlighted(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11highlightedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:221
// index:1
// Public Visibility=Default Availability=Available
// [-2] void highlighted(const QString &)

/*
This signal is sent when an item in the combobox popup list is highlighted by the user. The item's index is passed.

Note: Signal highlighted is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(comboBox, QOverload<int>::of(&QComboBox::highlighted),
      [=](int index){ /-* ... *-/ });
*/
func (this *QComboBox) Highlighted_1(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11highlightedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentIndexChanged(int)

/*
This signal is sent whenever the currentIndex in the combobox changes either through user interaction or programmatically. The item's index is passed or -1 if the combobox becomes empty or the currentIndex was reset.

Note: Signal currentIndexChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(comboBox, QOverload<int>::of(&QComboBox::currentIndexChanged),
      [=](int index){ /-* ... *-/ });



This function was introduced in  Qt 4.1.

Note: Notifier signal for property currentIndex.
*/
func (this *QComboBox) CurrentIndexChanged(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox19currentIndexChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:223
// index:1
// Public Visibility=Default Availability=Available
// [-2] void currentIndexChanged(const QString &)

/*
This signal is sent whenever the currentIndex in the combobox changes either through user interaction or programmatically. The item's index is passed or -1 if the combobox becomes empty or the currentIndex was reset.

Note: Signal currentIndexChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(comboBox, QOverload<int>::of(&QComboBox::currentIndexChanged),
      [=](int index){ /-* ... *-/ });



This function was introduced in  Qt 4.1.

Note: Notifier signal for property currentIndex.
*/
func (this *QComboBox) CurrentIndexChanged_1(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox19currentIndexChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentTextChanged(const QString &)

/*
This signal is sent whenever currentText changes. The new value is passed as text.

This function was introduced in  Qt 5.0.

Note: Notifier signal for property currentText.
*/
func (this *QComboBox) CurrentTextChanged(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox18currentTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:227
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusInEvent().
*/
func (this *QComboBox) FocusInEvent(e qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QFocusEvent_PTR() != nil {
		convArg0 = e.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:228
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusOutEvent().
*/
func (this *QComboBox) FocusOutEvent(e qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QFocusEvent_PTR() != nil {
		convArg0 = e.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:229
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QComboBox) ChangeEvent(e qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:230
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QComboBox) ResizeEvent(e qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QResizeEvent_PTR() != nil {
		convArg0 = e.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:231
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QComboBox) PaintEvent(e qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QPaintEvent_PTR() != nil {
		convArg0 = e.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:232
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QComboBox) ShowEvent(e qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QShowEvent_PTR() != nil {
		convArg0 = e.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:233
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
Reimplemented from QWidget::hideEvent().
*/
func (this *QComboBox) HideEvent(e qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QHideEvent_PTR() != nil {
		convArg0 = e.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:234
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QComboBox) MousePressEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:235
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QComboBox) MouseReleaseEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:236
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QComboBox) KeyPressEvent(e qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QKeyEvent_PTR() != nil {
		convArg0 = e.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:237
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyReleaseEvent().
*/
func (this *QComboBox) KeyReleaseEvent(e qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QKeyEvent_PTR() != nil {
		convArg0 = e.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:239
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().
*/
func (this *QComboBox) WheelEvent(e qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QWheelEvent_PTR() != nil {
		convArg0 = e.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:242
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*
Reimplemented from QWidget::contextMenuEvent().
*/
func (this *QComboBox) ContextMenuEvent(e qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QContextMenuEvent_PTR() != nil {
		convArg0 = e.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:244
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
Reimplemented from QWidget::inputMethodEvent().
*/
func (this *QComboBox) InputMethodEvent(arg0 qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QInputMethodEvent_PTR() != nil {
		convArg0 = arg0.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:245
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionComboBox *) const

/*
Initialize option with the values from this QComboBox. This method is useful for subclasses when they need a QStyleOptionComboBox, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QComboBox) InitStyleOption(option QStyleOptionComboBox_ITF /*777 QStyleOptionComboBox **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionComboBox_PTR() != nil {
		convArg0 = option.QStyleOptionComboBox_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox15initStyleOptionEP20QStyleOptionComboBox", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum specifies what the QComboBox should do when a new string is entered by the user.


*/
type QComboBox__InsertPolicy = int

// The string will not be inserted into the combobox.
const QComboBox__NoInsert QComboBox__InsertPolicy = 0

// The string will be inserted as the first item in the combobox.
const QComboBox__InsertAtTop QComboBox__InsertPolicy = 1

// The current item will be replaced by the string.
const QComboBox__InsertAtCurrent QComboBox__InsertPolicy = 2

// The string will be inserted after the last item in the combobox.
const QComboBox__InsertAtBottom QComboBox__InsertPolicy = 3

// The string is inserted after the current item in the combobox.
const QComboBox__InsertAfterCurrent QComboBox__InsertPolicy = 4

// The string is inserted before the current item in the combobox.
const QComboBox__InsertBeforeCurrent QComboBox__InsertPolicy = 5

// The string is inserted in the alphabetic order in the combobox.
const QComboBox__InsertAlphabetically QComboBox__InsertPolicy = 6

func (this *QComboBox) InsertPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QComboBox_InsertPolicyItemName(val int) string {
	var nilthis *QComboBox
	return nilthis.InsertPolicyItemName(val)
}

/*
This enum specifies how the size hint of the QComboBox should adjust when new content is added or content changes.


*/
type QComboBox__SizeAdjustPolicy = int

// The combobox will always adjust to the contents
const QComboBox__AdjustToContents QComboBox__SizeAdjustPolicy = 0

// The combobox will adjust to its contents the first time it is shown.
const QComboBox__AdjustToContentsOnFirstShow QComboBox__SizeAdjustPolicy = 1

// Use AdjustToContents or AdjustToContentsOnFirstShow instead.
const QComboBox__AdjustToMinimumContentsLength QComboBox__SizeAdjustPolicy = 2

// The combobox will adjust to minimumContentsLength plus space for an icon. For performance reasons use this policy on large models.
const QComboBox__AdjustToMinimumContentsLengthWithIcon QComboBox__SizeAdjustPolicy = 3

func (this *QComboBox) SizeAdjustPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QComboBox_SizeAdjustPolicyItemName(val int) string {
	var nilthis *QComboBox
	return nilthis.SizeAdjustPolicyItemName(val)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
