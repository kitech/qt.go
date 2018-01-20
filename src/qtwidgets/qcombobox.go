//  header block begin
// /usr/include/qt/QtWidgets/qcombobox.h
// #include <qcombobox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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
type QComboBox struct {
	*QWidget
}

func (this *QComboBox) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qcombobox.h:61
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QComboBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:85
// index:0
// void QComboBox(class QWidget *)
func NewQComboBox(parent unsafe.Pointer) *QComboBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQComboBoxFromPointer(cthis)
	return gothis
}
func NewQComboBoxFromPointer(cthis unsafe.Pointer) *QComboBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QComboBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qcombobox.h:249
// index:1
// void QComboBox(class QComboBoxPrivate &, class QWidget *)
func NewQComboBox_1(arg0 unsafe.Pointer, arg1 unsafe.Pointer) *QComboBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBoxC2ER16QComboBoxPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, arg0, arg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQComboBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcombobox.h:86
// index:0
// virtual
// void ~QComboBox()
func DeleteQComboBox(*QComboBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:88
// index:0
// int maxVisibleItems()
func (this *QComboBox) MaxVisibleItems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox15maxVisibleItemsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:89
// index:0
// void setMaxVisibleItems(int)
func (this *QComboBox) SetMaxVisibleItems(maxItems int) {
	// 0: (, maxItems int), (&maxItems)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox18setMaxVisibleItemsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &maxItems)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:91
// index:0
// int count()
func (this *QComboBox) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:92
// index:0
// void setMaxCount(int)
func (this *QComboBox) SetMaxCount(max int) {
	// 0: (, max int), (&max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setMaxCountEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:93
// index:0
// int maxCount()
func (this *QComboBox) MaxCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8maxCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:96
// index:0
// bool autoCompletion()
func (this *QComboBox) AutoCompletion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox14autoCompletionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:97
// index:0
// void setAutoCompletion(_Bool)
func (this *QComboBox) SetAutoCompletion(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox17setAutoCompletionEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:99
// index:0
// Qt::CaseSensitivity autoCompletionCaseSensitivity()
func (this *QComboBox) AutoCompletionCaseSensitivity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox29autoCompletionCaseSensitivityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:100
// index:0
// void setAutoCompletionCaseSensitivity(Qt::CaseSensitivity)
func (this *QComboBox) SetAutoCompletionCaseSensitivity(sensitivity int) {
	// 0: (, sensitivity Qt::CaseSensitivity), (&sensitivity)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox32setAutoCompletionCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sensitivity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:103
// index:0
// bool duplicatesEnabled()
func (this *QComboBox) DuplicatesEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox17duplicatesEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:104
// index:0
// void setDuplicatesEnabled(_Bool)
func (this *QComboBox) SetDuplicatesEnabled(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox20setDuplicatesEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:106
// index:0
// void setFrame(_Bool)
func (this *QComboBox) SetFrame(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8setFrameEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:107
// index:0
// bool hasFrame()
func (this *QComboBox) HasFrame() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8hasFrameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:109
// index:0
// inline
// int findText(const class QString &, Qt::MatchFlags)
func (this *QComboBox) FindText(text unsafe.Pointer, flags int) {
	// 0: (, text const QString &, QFlags<Qt::MatchFlag> flags), (text, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8findTextERK7QString6QFlagsIN2Qt9MatchFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:112
// index:0
// int findData(const class QVariant &, int, Qt::MatchFlags)
func (this *QComboBox) FindData(data unsafe.Pointer, role int, flags int) {
	// 0: (, data const QVariant &, role int, QFlags<Qt::MatchFlag> flags), (data, &role, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8findDataERK8QVarianti6QFlagsIN2Qt9MatchFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &role, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:126
// index:0
// QComboBox::InsertPolicy insertPolicy()
func (this *QComboBox) InsertPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12insertPolicyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:127
// index:0
// void setInsertPolicy(enum QComboBox::InsertPolicy)
func (this *QComboBox) SetInsertPolicy(policy int) {
	// 0: (, policy QComboBox::InsertPolicy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setInsertPolicyENS_12InsertPolicyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:137
// index:0
// QComboBox::SizeAdjustPolicy sizeAdjustPolicy()
func (this *QComboBox) SizeAdjustPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16sizeAdjustPolicyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:138
// index:0
// void setSizeAdjustPolicy(enum QComboBox::SizeAdjustPolicy)
func (this *QComboBox) SetSizeAdjustPolicy(policy int) {
	// 0: (, policy QComboBox::SizeAdjustPolicy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19setSizeAdjustPolicyENS_16SizeAdjustPolicyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:139
// index:0
// int minimumContentsLength()
func (this *QComboBox) MinimumContentsLength() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox21minimumContentsLengthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:140
// index:0
// void setMinimumContentsLength(int)
func (this *QComboBox) SetMinimumContentsLength(characters int) {
	// 0: (, characters int), (&characters)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox24setMinimumContentsLengthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &characters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:141
// index:0
// QSize iconSize()
func (this *QComboBox) IconSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8iconSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:142
// index:0
// void setIconSize(const class QSize &)
func (this *QComboBox) SetIconSize(size unsafe.Pointer) {
	// 0: (, size const QSize &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setIconSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:144
// index:0
// bool isEditable()
func (this *QComboBox) IsEditable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox10isEditableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:145
// index:0
// void setEditable(_Bool)
func (this *QComboBox) SetEditable(editable bool) {
	// 0: (, editable bool), (&editable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setEditableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &editable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:146
// index:0
// void setLineEdit(class QLineEdit *)
func (this *QComboBox) SetLineEdit(edit unsafe.Pointer) {
	// 0: (, edit QLineEdit *), (edit)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setLineEditEP9QLineEdit", ffiqt.FFI_TYPE_VOID, this.GetCthis(), edit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:147
// index:0
// QLineEdit * lineEdit()
func (this *QComboBox) LineEdit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8lineEditEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:149
// index:0
// void setValidator(const class QValidator *)
func (this *QComboBox) SetValidator(v unsafe.Pointer) {
	// 0: (, v const QValidator *), (v)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox12setValidatorEPK10QValidator", ffiqt.FFI_TYPE_VOID, this.GetCthis(), v)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:150
// index:0
// const QValidator * validator()
func (this *QComboBox) Validator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox9validatorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:154
// index:0
// void setCompleter(class QCompleter *)
func (this *QComboBox) SetCompleter(c unsafe.Pointer) {
	// 0: (, c QCompleter *), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox12setCompleterEP10QCompleter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:155
// index:0
// QCompleter * completer()
func (this *QComboBox) Completer() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox9completerEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:158
// index:0
// QAbstractItemDelegate * itemDelegate()
func (this *QComboBox) ItemDelegate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12itemDelegateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:159
// index:0
// void setItemDelegate(class QAbstractItemDelegate *)
func (this *QComboBox) SetItemDelegate(delegate unsafe.Pointer) {
	// 0: (, delegate QAbstractItemDelegate *), (delegate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setItemDelegateEP21QAbstractItemDelegate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:161
// index:0
// QAbstractItemModel * model()
func (this *QComboBox) Model() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox5modelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:162
// index:0
// void setModel(class QAbstractItemModel *)
func (this *QComboBox) SetModel(model unsafe.Pointer) {
	// 0: (, model QAbstractItemModel *), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:164
// index:0
// QModelIndex rootModelIndex()
func (this *QComboBox) RootModelIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox14rootModelIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:165
// index:0
// void setRootModelIndex(const class QModelIndex &)
func (this *QComboBox) SetRootModelIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox17setRootModelIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:167
// index:0
// int modelColumn()
func (this *QComboBox) ModelColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11modelColumnEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:168
// index:0
// void setModelColumn(int)
func (this *QComboBox) SetModelColumn(visibleColumn int) {
	// 0: (, visibleColumn int), (&visibleColumn)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox14setModelColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visibleColumn)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:170
// index:0
// int currentIndex()
func (this *QComboBox) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:171
// index:0
// QString currentText()
func (this *QComboBox) CurrentText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11currentTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:172
// index:0
// QVariant currentData(int)
func (this *QComboBox) CurrentData(role int) {
	// 0: (, role int), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11currentDataEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:174
// index:0
// QString itemText(int)
func (this *QComboBox) ItemText(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemTextEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:175
// index:0
// QIcon itemIcon(int)
func (this *QComboBox) ItemIcon(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemIconEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:176
// index:0
// QVariant itemData(int, int)
func (this *QComboBox) ItemData(index int, role int) {
	// 0: (, index int, role int), (&index, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemDataEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:178
// index:0
// inline
// void addItem(const class QString &, const class QVariant &)
func (this *QComboBox) AddItem(text unsafe.Pointer, userData unsafe.Pointer) {
	// 0: (, text const QString &, userData const QVariant &), (text, userData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7addItemERK7QStringRK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, userData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:179
// index:1
// inline
// void addItem(const class QIcon &, const class QString &, const class QVariant &)
func (this *QComboBox) AddItem_1(icon unsafe.Pointer, text unsafe.Pointer, userData unsafe.Pointer) {
	// 1: (, icon const QIcon &, text const QString &, userData const QVariant &), (icon, text, userData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon, text, userData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:181
// index:0
// inline
// void addItems(const class QStringList &)
func (this *QComboBox) AddItems(texts unsafe.Pointer) {
	// 0: (, texts const QStringList &), (texts)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8addItemsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), texts)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:184
// index:0
// inline
// void insertItem(int, const class QString &, const class QVariant &)
func (this *QComboBox) InsertItem(index int, text unsafe.Pointer, userData unsafe.Pointer) {
	// 0: (, index int, text const QString &, userData const QVariant &), (&index, text, userData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, text, userData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:185
// index:1
// void insertItem(int, const class QIcon &, const class QString &, const class QVariant &)
func (this *QComboBox) InsertItem_1(index int, icon unsafe.Pointer, text unsafe.Pointer, userData unsafe.Pointer) {
	// 1: (, index int, icon const QIcon &, text const QString &, userData const QVariant &), (&index, icon, text, userData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, icon, text, userData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:187
// index:0
// void insertItems(int, const class QStringList &)
func (this *QComboBox) InsertItems(index int, texts unsafe.Pointer) {
	// 0: (, index int, texts const QStringList &), (&index, texts)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11insertItemsEiRK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, texts)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:188
// index:0
// void insertSeparator(int)
func (this *QComboBox) InsertSeparator(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15insertSeparatorEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:190
// index:0
// void removeItem(int)
func (this *QComboBox) RemoveItem(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10removeItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:192
// index:0
// void setItemText(int, const class QString &)
func (this *QComboBox) SetItemText(index int, text unsafe.Pointer) {
	// 0: (, index int, text const QString &), (&index, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemTextEiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:193
// index:0
// void setItemIcon(int, const class QIcon &)
func (this *QComboBox) SetItemIcon(index int, icon unsafe.Pointer) {
	// 0: (, index int, icon const QIcon &), (&index, icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemIconEiRK5QIcon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:194
// index:0
// void setItemData(int, const class QVariant &, int)
func (this *QComboBox) SetItemData(index int, value unsafe.Pointer, role int) {
	// 0: (, index int, value const QVariant &, role int), (&index, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemDataEiRK8QVarianti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:196
// index:0
// QAbstractItemView * view()
func (this *QComboBox) View() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox4viewEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:197
// index:0
// void setView(class QAbstractItemView *)
func (this *QComboBox) SetView(itemView unsafe.Pointer) {
	// 0: (, itemView QAbstractItemView *), (itemView)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7setViewEP17QAbstractItemView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), itemView)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:199
// index:0
// virtual
// QSize sizeHint()
func (this *QComboBox) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:200
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QComboBox) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:202
// index:0
// virtual
// void showPopup()
func (this *QComboBox) ShowPopup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9showPopupEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:203
// index:0
// virtual
// void hidePopup()
func (this *QComboBox) HidePopup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9hidePopupEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:205
// index:0
// virtual
// bool event(class QEvent *)
func (this *QComboBox) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:206
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QComboBox) InputMethodQuery(arg0 int) {
	// 0: (, Qt::InputMethodQuery arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:207
// index:1
// QVariant inputMethodQuery(Qt::InputMethodQuery, const class QVariant &)
func (this *QComboBox) InputMethodQuery_1(query int, argument unsafe.Pointer) {
	// 1: (, query Qt::InputMethodQuery, argument const QVariant &), (&query, argument)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16inputMethodQueryEN2Qt16InputMethodQueryERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &query, argument)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:210
// index:0
// void clear()
func (this *QComboBox) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:211
// index:0
// void clearEditText()
func (this *QComboBox) ClearEditText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox13clearEditTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:212
// index:0
// void setEditText(const class QString &)
func (this *QComboBox) SetEditText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setEditTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:213
// index:0
// void setCurrentIndex(int)
func (this *QComboBox) SetCurrentIndex(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setCurrentIndexEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:214
// index:0
// void setCurrentText(const class QString &)
func (this *QComboBox) SetCurrentText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox14setCurrentTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:217
// index:0
// void editTextChanged(const class QString &)
func (this *QComboBox) EditTextChanged(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15editTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:218
// index:0
// void activated(int)
func (this *QComboBox) Activated(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9activatedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:219
// index:1
// void activated(const class QString &)
func (this *QComboBox) Activated_1(arg0 unsafe.Pointer) {
	// 1: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9activatedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:220
// index:0
// void highlighted(int)
func (this *QComboBox) Highlighted(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11highlightedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:221
// index:1
// void highlighted(const class QString &)
func (this *QComboBox) Highlighted_1(arg0 unsafe.Pointer) {
	// 1: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11highlightedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:222
// index:0
// void currentIndexChanged(int)
func (this *QComboBox) CurrentIndexChanged(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19currentIndexChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:223
// index:1
// void currentIndexChanged(const class QString &)
func (this *QComboBox) CurrentIndexChanged_1(arg0 unsafe.Pointer) {
	// 1: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19currentIndexChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:224
// index:0
// void currentTextChanged(const class QString &)
func (this *QComboBox) CurrentTextChanged(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox18currentTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:227
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QComboBox) FocusInEvent(e unsafe.Pointer) {
	// 0: (, e QFocusEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:228
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QComboBox) FocusOutEvent(e unsafe.Pointer) {
	// 0: (, e QFocusEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:229
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QComboBox) ChangeEvent(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:230
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QComboBox) ResizeEvent(e unsafe.Pointer) {
	// 0: (, e QResizeEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:231
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QComboBox) PaintEvent(e unsafe.Pointer) {
	// 0: (, e QPaintEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:232
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QComboBox) ShowEvent(e unsafe.Pointer) {
	// 0: (, e QShowEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:233
// index:0
// virtual
// void hideEvent(class QHideEvent *)
func (this *QComboBox) HideEvent(e unsafe.Pointer) {
	// 0: (, e QHideEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:234
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QComboBox) MousePressEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:235
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QComboBox) MouseReleaseEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:236
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QComboBox) KeyPressEvent(e unsafe.Pointer) {
	// 0: (, e QKeyEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:237
// index:0
// virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QComboBox) KeyReleaseEvent(e unsafe.Pointer) {
	// 0: (, e QKeyEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:239
// index:0
// virtual
// void wheelEvent(class QWheelEvent *)
func (this *QComboBox) WheelEvent(e unsafe.Pointer) {
	// 0: (, e QWheelEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:242
// index:0
// virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QComboBox) ContextMenuEvent(e unsafe.Pointer) {
	// 0: (, e QContextMenuEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:244
// index:0
// virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QComboBox) InputMethodEvent(arg0 unsafe.Pointer) {
	// 0: (, QInputMethodEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:245
// index:0
// void initStyleOption(class QStyleOptionComboBox *)
func (this *QComboBox) InitStyleOption(option unsafe.Pointer) {
	// 0: (, option QStyleOptionComboBox *), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox15initStyleOptionEP20QStyleOptionComboBox", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
