package qtwidgets

// /usr/include/qt/QtWidgets/qlayoutitem.h
// #include <qlayoutitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
// size 24
type QWidgetItem struct {
	*QLayoutItem
}
type QWidgetItem_ITF interface {
	QLayoutItem_ITF
	QWidgetItem_PTR() *QWidgetItem
}

func (ptr *QWidgetItem) QWidgetItem_PTR() *QWidgetItem { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QWidgetItemFromptr(cthis Voidptr) *QWidgetItem {
	bcthis0 := QLayoutItemFromptr(cthis)
	return &QWidgetItem{bcthis0}
}
func (*QWidgetItem) Fromptr(cthis Voidptr) *QWidgetItem {
	return QWidgetItemFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QWidgetItem(QWidget *)

/*
 */
func (*QWidgetItem) NewForInherit(w QWidget_ITF /*777 QWidget **/) *QWidgetItem {
	return NewQWidgetItem(w)
}
func NewQWidgetItem(w QWidget_ITF /*777 QWidget **/) *QWidgetItem {
	var convArg0 Voidptr
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(24)
	rv, err := qtrt.Qtcc3(3510836749, "_ZN11QWidgetItemC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QWidgetItemFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:133
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
 */
func (this *QWidgetItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(1339409613, "_ZNK11QWidgetItem8sizeHintEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:134
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize minimumSize() const

/*
 */
func (this *QWidgetItem) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(8341184, "_ZNK11QWidgetItem11minimumSizeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:135
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize maximumSize() const

/*
 */
func (this *QWidgetItem) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(1076903606, "_ZNK11QWidgetItem11maximumSizeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:136
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections() const

/*
 */
func (this *QWidgetItem) ExpandingDirections() int {
	rv, err := qtrt.Qtcc3(531281421, "_ZNK11QWidgetItem19expandingDirectionsEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:137
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
 */
func (this *QWidgetItem) IsEmpty() bool {
	rv, err := qtrt.Qtcc3(3246172656, "_ZNK11QWidgetItem7isEmptyEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:141
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QWidget * widget()

/*
 */
func (this *QWidgetItem) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.Qtcc3(4104848494, "_ZN11QWidgetItem6widgetEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QWidgetFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:147
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [4] int heightForWidth(int) const

/*
 */
func (this *QWidgetItem) HeightForWidth(arg0 int) int {
	rv, err := qtrt.Qtcc3(733877506, "_ZNK11QWidgetItem14heightForWidthEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQWidgetItem(this *QWidgetItem) {
	rv, err := qtrt.Qtcc1(0, "_ZN11QWidgetItemD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10107() {
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
