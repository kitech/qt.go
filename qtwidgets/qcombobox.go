// +build !minimal

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
// extern C begin: 9
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

/*
 */
// size 48
type QComboBox struct {
	*QWidget
}
type QComboBox_ITF interface {
	QWidget_ITF
	QComboBox_PTR() *QComboBox
}

func (ptr *QComboBox) QComboBox_PTR() *QComboBox { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QComboBoxFromptr(cthis Voidptr) *QComboBox {
	bcthis0 := QWidgetFromptr(cthis)
	return &QComboBox{bcthis0}
}
func (*QComboBox) Fromptr(cthis Voidptr) *QComboBox {
	return QComboBoxFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qcombobox.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QComboBox(QWidget *)

/*
 */
func (*QComboBox) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QComboBox {
	return NewQComboBox(parent)
}
func NewQComboBox(parent QWidget_ITF /*777 QWidget **/) *QComboBox {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(400444357, "_ZN9QComboBoxC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QComboBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QComboBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qcombobox.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QComboBox(QWidget *)

/*
 */
func (*QComboBox) NewForInheritp() *QComboBox {
	return NewQComboBoxp()
}
func NewQComboBoxp() *QComboBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(400444357, "_ZN9QComboBoxC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QComboBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QComboBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qcombobox.h:91
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int maxVisibleItems() const

/*
 */
func (this *QComboBox) MaxVisibleItems() int {
	rv, err := qtrt.Qtcc3(2452855331, "_ZNK9QComboBox15maxVisibleItemsEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:92
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaxVisibleItems(int)

/*
 */
func (this *QComboBox) SetMaxVisibleItems(maxItems int) {
	rv, err := qtrt.Qtcc3(2605451080, "_ZN9QComboBox18setMaxVisibleItemsEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&maxItems))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:94
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int count() const

/*
 */
func (this *QComboBox) Count() int {
	rv, err := qtrt.Qtcc3(484839320, "_ZNK9QComboBox5countEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:95
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaxCount(int)

/*
 */
func (this *QComboBox) SetMaxCount(max int) {
	rv, err := qtrt.Qtcc3(804866272, "_ZN9QComboBox11setMaxCountEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&max))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:96
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int maxCount() const

/*
 */
func (this *QComboBox) MaxCount() int {
	rv, err := qtrt.Qtcc3(3494713860, "_ZNK9QComboBox8maxCountEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:111
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool duplicatesEnabled() const

/*
 */
func (this *QComboBox) DuplicatesEnabled() bool {
	rv, err := qtrt.Qtcc3(3571501781, "_ZNK9QComboBox17duplicatesEnabledEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:112
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setDuplicatesEnabled(bool)

/*
 */
func (this *QComboBox) SetDuplicatesEnabled(enable bool) {
	rv, err := qtrt.Qtcc3(1134441101, "_ZN9QComboBox20setDuplicatesEnabledEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&enable))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:159
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setEditable(bool)

/*
 */
func (this *QComboBox) SetEditable(editable bool) {
	rv, err := qtrt.Qtcc3(1670976916, "_ZN9QComboBox11setEditableEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&editable))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:184
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int currentIndex() const

/*
 */
func (this *QComboBox) CurrentIndex() int {
	rv, err := qtrt.Qtcc3(3220285656, "_ZNK9QComboBox12currentIndexEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:185
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString currentText() const

/*
 */
func (this *QComboBox) CurrentText() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(1212789540, "_ZNK9QComboBox11currentTextEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qcombobox.h:188
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString itemText(int) const

/*
 */
func (this *QComboBox) ItemText(index int) string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(2984951329, "_ZNK9QComboBox8itemTextEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qcombobox.h:192
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void addItem(const QString &, const QVariant &)

/*
 */
func (this *QComboBox) AddItem(text string, userData qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 Voidptr
	if userData != nil && userData.QVariant_PTR() != nil {
		convArg1 = userData.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(931319215, "_ZN9QComboBox7addItemERK7QStringRK8QVariant", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:192
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void addItem(const QString &, const QVariant &)

/*
 */
func (this *QComboBox) AddItemp(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QVariant &=LValueReference, QVariant=Record, , Invalid
	var convArg1 = qtcore.NewQVariant()
	rv, err := qtrt.Qtcc3(931319215, "_ZN9QComboBox7addItemERK7QStringRK8QVariant", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:198
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void insertItem(int, const QString &, const QVariant &)

/*
 */
func (this *QComboBox) InsertItem(index int, text string, userData qtcore.QVariant_ITF) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 Voidptr
	if userData != nil && userData.QVariant_PTR() != nil {
		convArg2 = userData.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(3465040126, "_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&index), Voidptr(&convArg1), Voidptr(&convArg2))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:198
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void insertItem(int, const QString &, const QVariant &)

/*
 */
func (this *QComboBox) InsertItemp(index int, text string) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QVariant &=LValueReference, QVariant=Record, , Invalid
	var convArg2 = qtcore.NewQVariant()
	rv, err := qtrt.Qtcc3(3465040126, "_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&index), Voidptr(&convArg1), Voidptr(&convArg2))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:204
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void removeItem(int)

/*
 */
func (this *QComboBox) RemoveItem(index int) {
	rv, err := qtrt.Qtcc3(3369550016, "_ZN9QComboBox10removeItemEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:206
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setItemText(int, const QString &)

/*
 */
func (this *QComboBox) SetItemText(index int, text string) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.Qtcc3(198468933, "_ZN9QComboBox11setItemTextEiRK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&index), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:210
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QAbstractItemView * view() const

/*
 */
func (this *QComboBox) View() *QAbstractItemView /*777 QAbstractItemView **/ {
	rv, err := qtrt.Qtcc3(3310046969, "_ZNK9QComboBox4viewEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QAbstractItemViewFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcombobox.h:224
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clear()

/*
 */
func (this *QComboBox) Clear() {
	rv, err := qtrt.Qtcc3(3991980318, "_ZN9QComboBox5clearEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:225
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clearEditText()

/*
 */
func (this *QComboBox) ClearEditText() {
	rv, err := qtrt.Qtcc3(2293827830, "_ZN9QComboBox13clearEditTextEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:226
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setEditText(const QString &)

/*
 */
func (this *QComboBox) SetEditText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(3736805172, "_ZN9QComboBox11setEditTextERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:227
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)

/*
 */
func (this *QComboBox) SetCurrentIndex(index int) {
	rv, err := qtrt.Qtcc3(934993843, "_ZN9QComboBox15setCurrentIndexEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:228
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCurrentText(const QString &)

/*
 */
func (this *QComboBox) SetCurrentText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(3375923067, "_ZN9QComboBox14setCurrentTextERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:231
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void editTextChanged(const QString &)

/*
 */
func (this *QComboBox) EditTextChanged(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(681558797, "_ZN9QComboBox15editTextChangedERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:232
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void activated(int)

/*
 */
func (this *QComboBox) Activated(index int) {
	rv, err := qtrt.Qtcc3(3913087343, "_ZN9QComboBox9activatedEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:233
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void textActivated(const QString &)

/*
 */
func (this *QComboBox) TextActivated(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(1528827286, "_ZN9QComboBox13textActivatedERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:234
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void highlighted(int)

/*
 */
func (this *QComboBox) Highlighted(index int) {
	rv, err := qtrt.Qtcc3(1991475282, "_ZN9QComboBox11highlightedEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:235
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void textHighlighted(const QString &)

/*
 */
func (this *QComboBox) TextHighlighted(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(1845413289, "_ZN9QComboBox15textHighlightedERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:236
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void currentIndexChanged(int)

/*
 */
func (this *QComboBox) CurrentIndexChanged(index int) {
	rv, err := qtrt.Qtcc3(3753237607, "_ZN9QComboBox19currentIndexChangedEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQComboBox(this *QComboBox) {
	rv, err := qtrt.Qtcc3(2960591526, "_ZN9QComboBoxD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QComboBox__InsertPolicy = int

//
const QComboBox__NoInsert QComboBox__InsertPolicy = 0

//
const QComboBox__InsertAtTop QComboBox__InsertPolicy = 1

//
const QComboBox__InsertAtCurrent QComboBox__InsertPolicy = 2

//
const QComboBox__InsertAtBottom QComboBox__InsertPolicy = 3

//
const QComboBox__InsertAfterCurrent QComboBox__InsertPolicy = 4

//
const QComboBox__InsertBeforeCurrent QComboBox__InsertPolicy = 5

//
const QComboBox__InsertAlphabetically QComboBox__InsertPolicy = 6

func (this *QComboBox) InsertPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QComboBox_InsertPolicyItemName(val int) string {
	var nilthis *QComboBox
	return nilthis.InsertPolicyItemName(val)
}

/*


 */
type QComboBox__SizeAdjustPolicy = int

//
const QComboBox__AdjustToContents QComboBox__SizeAdjustPolicy = 0

//
const QComboBox__AdjustToContentsOnFirstShow QComboBox__SizeAdjustPolicy = 1

//
const QComboBox__AdjustToMinimumContentsLength QComboBox__SizeAdjustPolicy = 2

//
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

func init_unused_10209() {
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
