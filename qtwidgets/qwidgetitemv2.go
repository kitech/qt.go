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
// extern C begin: 8
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
// size 88
type QWidgetItemV2 struct {
	*QWidgetItem
}
type QWidgetItemV2_ITF interface {
	QWidgetItem_ITF
	QWidgetItemV2_PTR() *QWidgetItemV2
}

func (ptr *QWidgetItemV2) QWidgetItemV2_PTR() *QWidgetItemV2 { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QWidgetItemV2Fromptr(cthis Voidptr) *QWidgetItemV2 {
	bcthis0 := QWidgetItemFromptr(cthis)
	return &QWidgetItemV2{bcthis0}
}
func (*QWidgetItemV2) Fromptr(cthis Voidptr) *QWidgetItemV2 {
	return QWidgetItemV2Fromptr(cthis)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWidgetItemV2(QWidget *)

/*
 */
func (*QWidgetItemV2) NewForInherit(widget QWidget_ITF /*777 QWidget **/) *QWidgetItemV2 {
	return NewQWidgetItemV2(widget)
}
func NewQWidgetItemV2(widget QWidget_ITF /*777 QWidget **/) *QWidgetItemV2 {
	var convArg0 Voidptr
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(88)
	rv, err := qtrt.Qtcc3(1692828260, "_ZN13QWidgetItemV2C2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QWidgetItemV2Fromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQWidgetItemV2)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:159
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
 */
func (this *QWidgetItemV2) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(3325180898, "_ZNK13QWidgetItemV28sizeHintEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:160
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize minimumSize() const

/*
 */
func (this *QWidgetItemV2) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(3287127018, "_ZNK13QWidgetItemV211minimumSizeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:161
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize maximumSize() const

/*
 */
func (this *QWidgetItemV2) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(2208492444, "_ZNK13QWidgetItemV211maximumSizeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:162
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [4] int heightForWidth(int) const

/*
 */
func (this *QWidgetItemV2) HeightForWidth(width int) int {
	rv, err := qtrt.Qtcc3(4230616517, "_ZNK13QWidgetItemV214heightForWidthEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&width))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQWidgetItemV2(this *QWidgetItemV2) {
	rv, err := qtrt.Qtcc1(0, "_ZN13QWidgetItemV2D2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QWidgetItemV2__ = int

//
const QWidgetItemV2__Dirty QWidgetItemV2__ = -123

//
const QWidgetItemV2__HfwCacheMaxSize QWidgetItemV2__ = 3

func (this *QWidgetItemV2) ItemName(val int) string {
	switch val {
	case QWidgetItemV2__Dirty: // -123
		return "Dirty"
	case QWidgetItemV2__HfwCacheMaxSize: // 3
		return "HfwCacheMaxSize"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWidgetItemV2_ItemName(val int) string {
	var nilthis *QWidgetItemV2
	return nilthis.ItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10109() {
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
