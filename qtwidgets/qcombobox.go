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
func QComboBoxFromptr(cthis unsafe.Pointer) *QComboBox {
	bcthis0 := QWidgetFromptr(cthis)
	return &QComboBox{bcthis0}
}
func (*QComboBox) Fromptr(cthis unsafe.Pointer) *QComboBox {
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
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(400444357, "_ZN9QComboBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
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
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(400444357, "_ZN9QComboBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
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
	rv, err := qtrt.Qtcc1(2452855331, "_ZNK9QComboBox15maxVisibleItemsEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:92
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaxVisibleItems(int)

/*
 */
func (this *QComboBox) SetMaxVisibleItems(maxItems int) {
	rv, err := qtrt.Qtcc1(2605451080, "_ZN9QComboBox18setMaxVisibleItemsEi", qtrt.FFITY_POINTER, this.GetCthis(), maxItems)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:94
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int count() const

/*
 */
func (this *QComboBox) Count() int {
	rv, err := qtrt.Qtcc1(484839320, "_ZNK9QComboBox5countEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:95
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaxCount(int)

/*
 */
func (this *QComboBox) SetMaxCount(max int) {
	rv, err := qtrt.Qtcc1(804866272, "_ZN9QComboBox11setMaxCountEi", qtrt.FFITY_POINTER, this.GetCthis(), max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:96
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int maxCount() const

/*
 */
func (this *QComboBox) MaxCount() int {
	rv, err := qtrt.Qtcc1(3494713860, "_ZNK9QComboBox8maxCountEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQComboBox(this *QComboBox) {
	rv, err := qtrt.Qtcc1(0, "_ZN9QComboBoxD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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

func init_unused_10119() {
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
