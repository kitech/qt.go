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
// extern C begin: 0
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
// size 40
type QSpacerItem struct {
	*QLayoutItem
}
type QSpacerItem_ITF interface {
	QLayoutItem_ITF
	QSpacerItem_PTR() *QSpacerItem
}

func (ptr *QSpacerItem) QSpacerItem_PTR() *QSpacerItem { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QSpacerItemFromptr(cthis unsafe.Pointer) *QSpacerItem {
	bcthis0 := QLayoutItemFromptr(cthis)
	return &QSpacerItem{bcthis0}
}
func (*QSpacerItem) Fromptr(cthis unsafe.Pointer) *QSpacerItem {
	return QSpacerItemFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSpacerItem(int, int, QSizePolicy::Policy, QSizePolicy::Policy)

/*
 */
func (*QSpacerItem) NewForInherit(w int, h int, hData int, vData int) *QSpacerItem {
	return NewQSpacerItem(w, h, hData, vData)
}
func NewQSpacerItem(w int, h int, hData int, vData int) *QSpacerItem {
	cthis := qtrt.Malloc(40)
	rv, err := qtrt.Qtcc1(886124047, "_ZN11QSpacerItemC2EiiN11QSizePolicy6PolicyES1_", qtrt.FFITY_POINTER, cthis, w, h, hData, vData)
	qtrt.ErrPrint(err, rv)
	gothis := QSpacerItemFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQSpacerItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSpacerItem(int, int, QSizePolicy::Policy, QSizePolicy::Policy)

/*
 */
func (*QSpacerItem) NewForInheritp(w int, h int) *QSpacerItem {
	return NewQSpacerItemp(w, h)
}
func NewQSpacerItemp(w int, h int) *QSpacerItem {
	// arg: 2, QSizePolicy::Policy=Elaborated, QSizePolicy::Policy=Enum, , Invalid
	hData := 0
	// arg: 3, QSizePolicy::Policy=Elaborated, QSizePolicy::Policy=Enum, , Invalid
	vData := 0
	cthis := qtrt.Malloc(40)
	rv, err := qtrt.Qtcc1(886124047, "_ZN11QSpacerItemC2EiiN11QSizePolicy6PolicyES1_", qtrt.FFITY_POINTER, cthis, w, h, hData, vData)
	qtrt.ErrPrint(err, rv)
	gothis := QSpacerItemFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQSpacerItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSpacerItem(int, int, QSizePolicy::Policy, QSizePolicy::Policy)

/*
 */
func (*QSpacerItem) NewForInheritp1(w int, h int, hData int) *QSpacerItem {
	return NewQSpacerItemp1(w, h, hData)
}
func NewQSpacerItemp1(w int, h int, hData int) *QSpacerItem {
	// arg: 3, QSizePolicy::Policy=Elaborated, QSizePolicy::Policy=Enum, , Invalid
	vData := 0
	cthis := qtrt.Malloc(40)
	rv, err := qtrt.Qtcc1(886124047, "_ZN11QSpacerItemC2EiiN11QSizePolicy6PolicyES1_", qtrt.FFITY_POINTER, cthis, w, h, hData, vData)
	qtrt.ErrPrint(err, rv)
	gothis := QSpacerItemFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQSpacerItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:105
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void changeSize(int, int, QSizePolicy::Policy, QSizePolicy::Policy)

/*
 */
func (this *QSpacerItem) ChangeSize(w int, h int, hData int, vData int) {
	rv, err := qtrt.Qtcc1(2168887114, "_ZN11QSpacerItem10changeSizeEiiN11QSizePolicy6PolicyES1_", qtrt.FFITY_POINTER, this.GetCthis(), w, h, hData, vData)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:105
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void changeSize(int, int, QSizePolicy::Policy, QSizePolicy::Policy)

/*
 */
func (this *QSpacerItem) ChangeSizep(w int, h int) {
	// arg: 2, QSizePolicy::Policy=Elaborated, QSizePolicy::Policy=Enum, , Invalid
	hData := 0
	// arg: 3, QSizePolicy::Policy=Elaborated, QSizePolicy::Policy=Enum, , Invalid
	vData := 0
	rv, err := qtrt.Qtcc1(2168887114, "_ZN11QSpacerItem10changeSizeEiiN11QSizePolicy6PolicyES1_", qtrt.FFITY_POINTER, this.GetCthis(), w, h, hData, vData)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:105
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void changeSize(int, int, QSizePolicy::Policy, QSizePolicy::Policy)

/*
 */
func (this *QSpacerItem) ChangeSizep1(w int, h int, hData int) {
	// arg: 3, QSizePolicy::Policy=Elaborated, QSizePolicy::Policy=Enum, , Invalid
	vData := 0
	rv, err := qtrt.Qtcc1(2168887114, "_ZN11QSpacerItem10changeSizeEiiN11QSizePolicy6PolicyES1_", qtrt.FFITY_POINTER, this.GetCthis(), w, h, hData, vData)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:108
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
 */
func (this *QSpacerItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc1(602256895, "_ZNK11QSpacerItem8sizeHintEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.QSizeFromptr(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:109
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize minimumSize() const

/*
 */
func (this *QSpacerItem) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc1(2485939981, "_ZNK11QSpacerItem11minimumSizeEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.QSizeFromptr(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:110
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize maximumSize() const

/*
 */
func (this *QSpacerItem) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc1(3563263867, "_ZNK11QSpacerItem11maximumSizeEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.QSizeFromptr(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:111
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections() const

/*
 */
func (this *QSpacerItem) ExpandingDirections() int {
	rv, err := qtrt.Qtcc1(179331885, "_ZNK11QSpacerItem19expandingDirectionsEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:112
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
 */
func (this *QSpacerItem) IsEmpty() bool {
	rv, err := qtrt.Qtcc1(4153964196, "_ZNK11QSpacerItem7isEmptyEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQSpacerItem(this *QSpacerItem) {
	rv, err := qtrt.Qtcc1(0, "_ZN11QSpacerItemD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10101() {
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
