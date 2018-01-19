//  header block begin
// /usr/include/qt/QtWidgets/qcombobox.h
// #include <qcombobox.h>
// #include <QtWidgets>
package qtwidgets

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qcombobox.h:61
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QComboBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:85
// index:0
// void QComboBox(class QWidget *)
func NewQComboBox(parent unsafe.Pointer) *QComboBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QComboBox{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox15maxVisibleItemsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:89
// index:0
// void setMaxVisibleItems(int)
func (this *QComboBox) SetMaxVisibleItems(maxItems int) {
	// 0: (, int maxItems), (&maxItems)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox18setMaxVisibleItemsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &maxItems)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:91
// index:0
// int count()
func (this *QComboBox) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:92
// index:0
// void setMaxCount(int)
func (this *QComboBox) SetMaxCount(max int) {
	// 0: (, int max), (&max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setMaxCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:93
// index:0
// int maxCount()
func (this *QComboBox) MaxCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8maxCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:96
// index:0
// bool autoCompletion()
func (this *QComboBox) AutoCompletion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox14autoCompletionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:97
// index:0
// void setAutoCompletion(_Bool)
func (this *QComboBox) SetAutoCompletion(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox17setAutoCompletionEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:99
// index:0
// Qt::CaseSensitivity autoCompletionCaseSensitivity()
func (this *QComboBox) AutoCompletionCaseSensitivity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox29autoCompletionCaseSensitivityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:100
// index:0
// void setAutoCompletionCaseSensitivity(Qt::CaseSensitivity)
func (this *QComboBox) SetAutoCompletionCaseSensitivity(sensitivity int) {
	// 0: (, Qt::CaseSensitivity sensitivity), (&sensitivity)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox32setAutoCompletionCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, &sensitivity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:103
// index:0
// bool duplicatesEnabled()
func (this *QComboBox) DuplicatesEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox17duplicatesEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:104
// index:0
// void setDuplicatesEnabled(_Bool)
func (this *QComboBox) SetDuplicatesEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox20setDuplicatesEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:106
// index:0
// void setFrame(_Bool)
func (this *QComboBox) SetFrame(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8setFrameEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:107
// index:0
// bool hasFrame()
func (this *QComboBox) HasFrame() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8hasFrameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:126
// index:0
// QComboBox::InsertPolicy insertPolicy()
func (this *QComboBox) InsertPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12insertPolicyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:127
// index:0
// void setInsertPolicy(enum QComboBox::InsertPolicy)
func (this *QComboBox) SetInsertPolicy(policy int) {
	// 0: (, QComboBox::InsertPolicy policy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setInsertPolicyENS_12InsertPolicyE", ffiqt.FFI_TYPE_VOID, this.cthis, &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:137
// index:0
// QComboBox::SizeAdjustPolicy sizeAdjustPolicy()
func (this *QComboBox) SizeAdjustPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16sizeAdjustPolicyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:138
// index:0
// void setSizeAdjustPolicy(enum QComboBox::SizeAdjustPolicy)
func (this *QComboBox) SetSizeAdjustPolicy(policy int) {
	// 0: (, QComboBox::SizeAdjustPolicy policy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19setSizeAdjustPolicyENS_16SizeAdjustPolicyE", ffiqt.FFI_TYPE_VOID, this.cthis, &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:139
// index:0
// int minimumContentsLength()
func (this *QComboBox) MinimumContentsLength() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox21minimumContentsLengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:140
// index:0
// void setMinimumContentsLength(int)
func (this *QComboBox) SetMinimumContentsLength(characters int) {
	// 0: (, int characters), (&characters)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox24setMinimumContentsLengthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &characters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:141
// index:0
// QSize iconSize()
func (this *QComboBox) IconSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8iconSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:142
// index:0
// void setIconSize(const class QSize &)
func (this *QComboBox) SetIconSize(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setIconSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:144
// index:0
// bool isEditable()
func (this *QComboBox) IsEditable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox10isEditableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:145
// index:0
// void setEditable(_Bool)
func (this *QComboBox) SetEditable(editable bool) {
	// 0: (, bool editable), (&editable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setEditableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &editable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:146
// index:0
// void setLineEdit(class QLineEdit *)
func (this *QComboBox) SetLineEdit(edit unsafe.Pointer) {
	// 0: (, QLineEdit * edit), (edit)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setLineEditEP9QLineEdit", ffiqt.FFI_TYPE_VOID, this.cthis, edit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:147
// index:0
// QLineEdit * lineEdit()
func (this *QComboBox) LineEdit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8lineEditEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:149
// index:0
// void setValidator(const class QValidator *)
func (this *QComboBox) SetValidator(v unsafe.Pointer) {
	// 0: (, const QValidator * v), (v)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox12setValidatorEPK10QValidator", ffiqt.FFI_TYPE_VOID, this.cthis, v)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:150
// index:0
// const QValidator * validator()
func (this *QComboBox) Validator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox9validatorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:154
// index:0
// void setCompleter(class QCompleter *)
func (this *QComboBox) SetCompleter(c unsafe.Pointer) {
	// 0: (, QCompleter * c), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox12setCompleterEP10QCompleter", ffiqt.FFI_TYPE_VOID, this.cthis, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:155
// index:0
// QCompleter * completer()
func (this *QComboBox) Completer() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox9completerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:158
// index:0
// QAbstractItemDelegate * itemDelegate()
func (this *QComboBox) ItemDelegate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12itemDelegateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:159
// index:0
// void setItemDelegate(class QAbstractItemDelegate *)
func (this *QComboBox) SetItemDelegate(delegate unsafe.Pointer) {
	// 0: (, QAbstractItemDelegate * delegate), (delegate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setItemDelegateEP21QAbstractItemDelegate", ffiqt.FFI_TYPE_VOID, this.cthis, delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:161
// index:0
// QAbstractItemModel * model()
func (this *QComboBox) Model() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox5modelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:162
// index:0
// void setModel(class QAbstractItemModel *)
func (this *QComboBox) SetModel(model unsafe.Pointer) {
	// 0: (, QAbstractItemModel * model), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.cthis, model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:164
// index:0
// QModelIndex rootModelIndex()
func (this *QComboBox) RootModelIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox14rootModelIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:165
// index:0
// void setRootModelIndex(const class QModelIndex &)
func (this *QComboBox) SetRootModelIndex(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox17setRootModelIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:167
// index:0
// int modelColumn()
func (this *QComboBox) ModelColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11modelColumnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:168
// index:0
// void setModelColumn(int)
func (this *QComboBox) SetModelColumn(visibleColumn int) {
	// 0: (, int visibleColumn), (&visibleColumn)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox14setModelColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &visibleColumn)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:170
// index:0
// int currentIndex()
func (this *QComboBox) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:171
// index:0
// QString currentText()
func (this *QComboBox) CurrentText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11currentTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:172
// index:0
// QVariant currentData(int)
func (this *QComboBox) CurrentData(role int) {
	// 0: (, int role), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11currentDataEi", ffiqt.FFI_TYPE_VOID, this.cthis, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:174
// index:0
// QString itemText(int)
func (this *QComboBox) ItemText(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemTextEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:175
// index:0
// QIcon itemIcon(int)
func (this *QComboBox) ItemIcon(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemIconEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:176
// index:0
// QVariant itemData(int, int)
func (this *QComboBox) ItemData(index int, role int) {
	// 0: (, int index, int role), (&index, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemDataEii", ffiqt.FFI_TYPE_VOID, this.cthis, &index, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:178
// index:0
// inline
// void addItem(const class QString &, const class QVariant &)
func (this *QComboBox) AddItem(text unsafe.Pointer, userData unsafe.Pointer) {
	// 0: (, const QString & text, const QVariant & userData), (text, userData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7addItemERK7QStringRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, text, userData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:179
// index:1
// inline
// void addItem(const class QIcon &, const class QString &, const class QVariant &)
func (this *QComboBox) AddItem_1(icon unsafe.Pointer, text unsafe.Pointer, userData unsafe.Pointer) {
	// 1: (, const QIcon & icon, const QString & text, const QVariant & userData), (icon, text, userData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, icon, text, userData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:181
// index:0
// inline
// void addItems(const class QStringList &)
func (this *QComboBox) AddItems(texts unsafe.Pointer) {
	// 0: (, const QStringList & texts), (texts)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8addItemsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, texts)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:184
// index:0
// inline
// void insertItem(int, const class QString &, const class QVariant &)
func (this *QComboBox) InsertItem(index int, text unsafe.Pointer, userData unsafe.Pointer) {
	// 0: (, int index, const QString & text, const QVariant & userData), (&index, text, userData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &index, text, userData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:185
// index:1
// void insertItem(int, const class QIcon &, const class QString &, const class QVariant &)
func (this *QComboBox) InsertItem_1(index int, icon unsafe.Pointer, text unsafe.Pointer, userData unsafe.Pointer) {
	// 1: (, int index, const QIcon & icon, const QString & text, const QVariant & userData), (&index, icon, text, userData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &index, icon, text, userData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:187
// index:0
// void insertItems(int, const class QStringList &)
func (this *QComboBox) InsertItems(index int, texts unsafe.Pointer) {
	// 0: (, int index, const QStringList & texts), (&index, texts)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11insertItemsEiRK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, &index, texts)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:188
// index:0
// void insertSeparator(int)
func (this *QComboBox) InsertSeparator(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15insertSeparatorEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:190
// index:0
// void removeItem(int)
func (this *QComboBox) RemoveItem(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10removeItemEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:192
// index:0
// void setItemText(int, const class QString &)
func (this *QComboBox) SetItemText(index int, text unsafe.Pointer) {
	// 0: (, int index, const QString & text), (&index, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemTextEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &index, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:193
// index:0
// void setItemIcon(int, const class QIcon &)
func (this *QComboBox) SetItemIcon(index int, icon unsafe.Pointer) {
	// 0: (, int index, const QIcon & icon), (&index, icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemIconEiRK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, &index, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:194
// index:0
// void setItemData(int, const class QVariant &, int)
func (this *QComboBox) SetItemData(index int, value unsafe.Pointer, role int) {
	// 0: (, int index, const QVariant & value, int role), (&index, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemDataEiRK8QVarianti", ffiqt.FFI_TYPE_VOID, this.cthis, &index, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:196
// index:0
// QAbstractItemView * view()
func (this *QComboBox) View() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox4viewEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:197
// index:0
// void setView(class QAbstractItemView *)
func (this *QComboBox) SetView(itemView unsafe.Pointer) {
	// 0: (, QAbstractItemView * itemView), (itemView)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7setViewEP17QAbstractItemView", ffiqt.FFI_TYPE_VOID, this.cthis, itemView)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:199
// index:0
// virtual
// QSize sizeHint()
func (this *QComboBox) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:200
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QComboBox) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:202
// index:0
// virtual
// void showPopup()
func (this *QComboBox) ShowPopup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9showPopupEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:203
// index:0
// virtual
// void hidePopup()
func (this *QComboBox) HidePopup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9hidePopupEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:205
// index:0
// virtual
// bool event(class QEvent *)
func (this *QComboBox) Event(event unsafe.Pointer) {
	// 0: (, QEvent * event), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.cthis, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:206
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QComboBox) InputMethodQuery(arg0 int) {
	// 0: (, Qt::InputMethodQuery arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:207
// index:1
// QVariant inputMethodQuery(Qt::InputMethodQuery, const class QVariant &)
func (this *QComboBox) InputMethodQuery_1(query int, argument unsafe.Pointer) {
	// 1: (, Qt::InputMethodQuery query, const QVariant & argument), (&query, argument)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16inputMethodQueryEN2Qt16InputMethodQueryERK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &query, argument)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:210
// index:0
// void clear()
func (this *QComboBox) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:211
// index:0
// void clearEditText()
func (this *QComboBox) ClearEditText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox13clearEditTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:212
// index:0
// void setEditText(const class QString &)
func (this *QComboBox) SetEditText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setEditTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:213
// index:0
// void setCurrentIndex(int)
func (this *QComboBox) SetCurrentIndex(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setCurrentIndexEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:214
// index:0
// void setCurrentText(const class QString &)
func (this *QComboBox) SetCurrentText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox14setCurrentTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:217
// index:0
// void editTextChanged(const class QString &)
func (this *QComboBox) EditTextChanged(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15editTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:218
// index:0
// void activated(int)
func (this *QComboBox) Activated(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9activatedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:219
// index:1
// void activated(const class QString &)
func (this *QComboBox) Activated_1(arg0 unsafe.Pointer) {
	// 1: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9activatedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:220
// index:0
// void highlighted(int)
func (this *QComboBox) Highlighted(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11highlightedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:221
// index:1
// void highlighted(const class QString &)
func (this *QComboBox) Highlighted_1(arg0 unsafe.Pointer) {
	// 1: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11highlightedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:222
// index:0
// void currentIndexChanged(int)
func (this *QComboBox) CurrentIndexChanged(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19currentIndexChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:223
// index:1
// void currentIndexChanged(const class QString &)
func (this *QComboBox) CurrentIndexChanged_1(arg0 unsafe.Pointer) {
	// 1: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19currentIndexChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:224
// index:0
// void currentTextChanged(const class QString &)
func (this *QComboBox) CurrentTextChanged(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox18currentTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
