package qtwidgets

// /usr/include/qt/QtWidgets/qcombobox.h
// #include <qcombobox.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
// void focusInEvent(class QFocusEvent *)
func (this *QComboBox) InheritFocusInEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QComboBox) InheritFocusOutEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QComboBox) InheritChangeEvent(f func(e *qtcore.QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "changeEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QComboBox) InheritResizeEvent(f func(e *qtgui.QResizeEvent /*777 QResizeEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QComboBox) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "paintEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QComboBox) InheritShowEvent(f func(e *qtgui.QShowEvent /*777 QShowEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(class QHideEvent *)
func (this *QComboBox) InheritHideEvent(f func(e *qtgui.QHideEvent /*777 QHideEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "hideEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QComboBox) InheritMousePressEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QComboBox) InheritMouseReleaseEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QComboBox) InheritKeyPressEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QComboBox) InheritKeyReleaseEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QComboBox) InheritWheelEvent(f func(e *qtgui.QWheelEvent /*777 QWheelEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QComboBox) InheritContextMenuEvent(f func(e *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void inputMethodEvent(class QInputMethodEvent *)
func (this *QComboBox) InheritInputMethodEvent(f func(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void initStyleOption(class QStyleOptionComboBox *)
func (this *QComboBox) InheritInitStyleOption(f func(option *QStyleOptionComboBox /*777 QStyleOptionComboBox **/)) {
	ffiqt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QComboBox struct {
	*QWidget
}

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
// [8] const QMetaObject * metaObject()
func (this *QComboBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QComboBox(QWidget *)
func NewQComboBox(parent *QWidget /*777 QWidget **/) *QComboBox {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBoxC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQComboBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qcombobox.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QComboBox()
func DeleteQComboBox(this *QComboBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBoxD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qcombobox.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxVisibleItems()
func (this *QComboBox) MaxVisibleItems() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox15maxVisibleItemsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxVisibleItems(int)
func (this *QComboBox) SetMaxVisibleItems(maxItems int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox18setMaxVisibleItemsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), maxItems)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QComboBox) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxCount(int)
func (this *QComboBox) SetMaxCount(max int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setMaxCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxCount()
func (this *QComboBox) MaxCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8maxCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoCompletion()
func (this *QComboBox) AutoCompletion() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox14autoCompletionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoCompletion(_Bool)
func (this *QComboBox) SetAutoCompletion(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox17setAutoCompletionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity autoCompletionCaseSensitivity()
func (this *QComboBox) AutoCompletionCaseSensitivity() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox29autoCompletionCaseSensitivityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoCompletionCaseSensitivity(Qt::CaseSensitivity)
func (this *QComboBox) SetAutoCompletionCaseSensitivity(sensitivity int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox32setAutoCompletionCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sensitivity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool duplicatesEnabled()
func (this *QComboBox) DuplicatesEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox17duplicatesEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDuplicatesEnabled(_Bool)
func (this *QComboBox) SetDuplicatesEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox20setDuplicatesEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrame(_Bool)
func (this *QComboBox) SetFrame(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8setFrameEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFrame()
func (this *QComboBox) HasFrame() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8hasFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int findText(const QString &, Qt::MatchFlags)
func (this *QComboBox) FindText(text *qtcore.QString, flags int) int {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8findTextERK7QString6QFlagsIN2Qt9MatchFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int findData(const QVariant &, int, Qt::MatchFlags)
func (this *QComboBox) FindData(data *qtcore.QVariant, role int, flags int) int {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8findDataERK8QVarianti6QFlagsIN2Qt9MatchFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:126
// index:0
// Public Visibility=Default Availability=Available
// [4] QComboBox::InsertPolicy insertPolicy()
func (this *QComboBox) InsertPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12insertPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInsertPolicy(enum QComboBox::InsertPolicy)
func (this *QComboBox) SetInsertPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setInsertPolicyENS_12InsertPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] QComboBox::SizeAdjustPolicy sizeAdjustPolicy()
func (this *QComboBox) SizeAdjustPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16sizeAdjustPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeAdjustPolicy(enum QComboBox::SizeAdjustPolicy)
func (this *QComboBox) SetSizeAdjustPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19setSizeAdjustPolicyENS_16SizeAdjustPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimumContentsLength()
func (this *QComboBox) MinimumContentsLength() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox21minimumContentsLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumContentsLength(int)
func (this *QComboBox) SetMinimumContentsLength(characters int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox24setMinimumContentsLengthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), characters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize()
func (this *QComboBox) IconSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8iconSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)
func (this *QComboBox) SetIconSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setIconSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEditable()
func (this *QComboBox) IsEditable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox10isEditableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEditable(_Bool)
func (this *QComboBox) SetEditable(editable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setEditableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineEdit(QLineEdit *)
func (this *QComboBox) SetLineEdit(edit *QLineEdit /*777 QLineEdit **/) {
	var convArg0 = edit.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setLineEditEP9QLineEdit", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:147
// index:0
// Public Visibility=Default Availability=Available
// [8] QLineEdit * lineEdit()
func (this *QComboBox) LineEdit() *QLineEdit /*777 QLineEdit **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8lineEditEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValidator(const QValidator *)
func (this *QComboBox) SetValidator(v *qtgui.QValidator /*777 const QValidator **/) {
	var convArg0 = v.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox12setValidatorEPK10QValidator", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:150
// index:0
// Public Visibility=Default Availability=Available
// [8] const QValidator * validator()
func (this *QComboBox) Validator() *qtgui.QValidator /*777 const QValidator **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox9validatorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQValidatorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompleter(QCompleter *)
func (this *QComboBox) SetCompleter(c *QCompleter /*777 QCompleter **/) {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox12setCompleterEP10QCompleter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QCompleter * completer()
func (this *QComboBox) Completer() *QCompleter /*777 QCompleter **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox9completerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:158
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegate()
func (this *QComboBox) ItemDelegate() *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12itemDelegateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemDelegate(QAbstractItemDelegate *)
func (this *QComboBox) SetItemDelegate(delegate *QAbstractItemDelegate /*777 QAbstractItemDelegate **/) {
	var convArg0 = delegate.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setItemDelegateEP21QAbstractItemDelegate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:161
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemModel * model()
func (this *QComboBox) Model() *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)
func (this *QComboBox) SetModel(model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/) {
	var convArg0 = model.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:164
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex rootModelIndex()
func (this *QComboBox) RootModelIndex() *qtcore.QModelIndex /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox14rootModelIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRootModelIndex(const QModelIndex &)
func (this *QComboBox) SetRootModelIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox17setRootModelIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] int modelColumn()
func (this *QComboBox) ModelColumn() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11modelColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModelColumn(int)
func (this *QComboBox) SetModelColumn(visibleColumn int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox14setModelColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visibleColumn)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:170
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex()
func (this *QComboBox) CurrentIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] QString currentText()
func (this *QComboBox) CurrentText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11currentTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:172
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant currentData(int)
func (this *QComboBox) CurrentData(role int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11currentDataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:174
// index:0
// Public Visibility=Default Availability=Available
// [8] QString itemText(int)
func (this *QComboBox) ItemText(index int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemTextEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:175
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon itemIcon(int)
func (this *QComboBox) ItemIcon(index int) *qtgui.QIcon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemIconEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:176
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant itemData(int, int)
func (this *QComboBox) ItemData(index int, role int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemDataEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:178
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(const QString &, const QVariant &)
func (this *QComboBox) AddItem(text *qtcore.QString, userData *qtcore.QVariant) {
	var convArg0 = text.GetCthis()
	var convArg1 = userData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7addItemERK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:179
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(const QIcon &, const QString &, const QVariant &)
func (this *QComboBox) AddItem_1(icon *qtgui.QIcon, text *qtcore.QString, userData *qtcore.QVariant) {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	var convArg2 = userData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:181
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addItems(const QStringList &)
func (this *QComboBox) AddItems(texts *qtcore.QStringList) {
	var convArg0 = texts.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8addItemsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:184
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void insertItem(int, const QString &, const QVariant &)
func (this *QComboBox) InsertItem(index int, text *qtcore.QString, userData *qtcore.QVariant) {
	var convArg1 = text.GetCthis()
	var convArg2 = userData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:185
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertItem(int, const QIcon &, const QString &, const QVariant &)
func (this *QComboBox) InsertItem_1(index int, icon *qtgui.QIcon, text *qtcore.QString, userData *qtcore.QVariant) {
	var convArg1 = icon.GetCthis()
	var convArg2 = text.GetCthis()
	var convArg3 = userData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertItems(int, const QStringList &)
func (this *QComboBox) InsertItems(index int, texts *qtcore.QStringList) {
	var convArg1 = texts.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11insertItemsEiRK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertSeparator(int)
func (this *QComboBox) InsertSeparator(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15insertSeparatorEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(int)
func (this *QComboBox) RemoveItem(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10removeItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemText(int, const QString &)
func (this *QComboBox) SetItemText(index int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemIcon(int, const QIcon &)
func (this *QComboBox) SetItemIcon(index int, icon *qtgui.QIcon) {
	var convArg1 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemIconEiRK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemData(int, const QVariant &, int)
func (this *QComboBox) SetItemData(index int, value *qtcore.QVariant, role int) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemDataEiRK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:196
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemView * view()
func (this *QComboBox) View() *QAbstractItemView /*777 QAbstractItemView **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox4viewEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractItemViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setView(QAbstractItemView *)
func (this *QComboBox) SetView(itemView *QAbstractItemView /*777 QAbstractItemView **/) {
	var convArg0 = itemView.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7setViewEP17QAbstractItemView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:199
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QComboBox) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:200
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QComboBox) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:202
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void showPopup()
func (this *QComboBox) ShowPopup() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9showPopupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:203
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void hidePopup()
func (this *QComboBox) HidePopup() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9hidePopupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:205
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QComboBox) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:206
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QComboBox) InputMethodQuery(arg0 int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:207
// index:1
// Public Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery, const QVariant &)
func (this *QComboBox) InputMethodQuery_1(query int, argument *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	var convArg1 = argument.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16inputMethodQueryEN2Qt16InputMethodQueryERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), query, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qcombobox.h:210
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QComboBox) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearEditText()
func (this *QComboBox) ClearEditText() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox13clearEditTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:212
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEditText(const QString &)
func (this *QComboBox) SetEditText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setEditTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)
func (this *QComboBox) SetCurrentIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setCurrentIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentText(const QString &)
func (this *QComboBox) SetCurrentText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox14setCurrentTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editTextChanged(const QString &)
func (this *QComboBox) EditTextChanged(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15editTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated(int)
func (this *QComboBox) Activated(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9activatedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:219
// index:1
// Public Visibility=Default Availability=Available
// [-2] void activated(const QString &)
func (this *QComboBox) Activated_1(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9activatedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void highlighted(int)
func (this *QComboBox) Highlighted(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11highlightedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:221
// index:1
// Public Visibility=Default Availability=Available
// [-2] void highlighted(const QString &)
func (this *QComboBox) Highlighted_1(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11highlightedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentIndexChanged(int)
func (this *QComboBox) CurrentIndexChanged(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19currentIndexChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:223
// index:1
// Public Visibility=Default Availability=Available
// [-2] void currentIndexChanged(const QString &)
func (this *QComboBox) CurrentIndexChanged_1(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19currentIndexChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentTextChanged(const QString &)
func (this *QComboBox) CurrentTextChanged(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox18currentTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:227
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QComboBox) FocusInEvent(e *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:228
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QComboBox) FocusOutEvent(e *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:229
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QComboBox) ChangeEvent(e *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:230
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QComboBox) ResizeEvent(e *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:231
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QComboBox) PaintEvent(e *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:232
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QComboBox) ShowEvent(e *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:233
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)
func (this *QComboBox) HideEvent(e *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:234
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QComboBox) MousePressEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:235
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QComboBox) MouseReleaseEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:236
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QComboBox) KeyPressEvent(e *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:237
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QComboBox) KeyReleaseEvent(e *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:239
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QComboBox) WheelEvent(e *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:242
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QComboBox) ContextMenuEvent(e *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:244
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)
func (this *QComboBox) InputMethodEvent(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:245
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionComboBox *)
func (this *QComboBox) InitStyleOption(option *QStyleOptionComboBox /*777 QStyleOptionComboBox **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox15initStyleOptionEP20QStyleOptionComboBox", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QComboBox__InsertPolicy = int

const QComboBox__NoInsert QComboBox__InsertPolicy = 0
const QComboBox__InsertAtTop QComboBox__InsertPolicy = 1
const QComboBox__InsertAtCurrent QComboBox__InsertPolicy = 2
const QComboBox__InsertAtBottom QComboBox__InsertPolicy = 3
const QComboBox__InsertAfterCurrent QComboBox__InsertPolicy = 4
const QComboBox__InsertBeforeCurrent QComboBox__InsertPolicy = 5
const QComboBox__InsertAlphabetically QComboBox__InsertPolicy = 6

type QComboBox__SizeAdjustPolicy = int

const QComboBox__AdjustToContents QComboBox__SizeAdjustPolicy = 0
const QComboBox__AdjustToContentsOnFirstShow QComboBox__SizeAdjustPolicy = 1
const QComboBox__AdjustToMinimumContentsLength QComboBox__SizeAdjustPolicy = 2
const QComboBox__AdjustToMinimumContentsLengthWithIcon QComboBox__SizeAdjustPolicy = 3

//  body block end
