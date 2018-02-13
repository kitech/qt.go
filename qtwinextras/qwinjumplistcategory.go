package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h
// #include <qwinjumplistcategory.h>
// #include <Qtwinextras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QWinJumpListCategory struct {
	*qtrt.CObject
}
type QWinJumpListCategory_ITF interface {
	QWinJumpListCategory_PTR() *QWinJumpListCategory
}

func (ptr *QWinJumpListCategory) QWinJumpListCategory_PTR() *QWinJumpListCategory { return ptr }

func (this *QWinJumpListCategory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWinJumpListCategory) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWinJumpListCategoryFromPointer(cthis unsafe.Pointer) *QWinJumpListCategory {
	return &QWinJumpListCategory{&qtrt.CObject{cthis}}
}
func (*QWinJumpListCategory) NewFromPointer(cthis unsafe.Pointer) *QWinJumpListCategory {
	return NewQWinJumpListCategoryFromPointer(cthis)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinJumpListCategory(const QString &)
func NewQWinJumpListCategory(title string) *QWinJumpListCategory {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinJumpListCategoryC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWinJumpListCategoryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWinJumpListCategory)
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QWinJumpListCategory()
func DeleteQWinJumpListCategory(this *QWinJumpListCategory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinJumpListCategoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] QWinJumpListCategory::Type type()
func (this *QWinJumpListCategory) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinJumpListCategory4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QWinJumpListCategory) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinJumpListCategory9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QWinJumpListCategory) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinJumpListCategory10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title()
func (this *QWinJumpListCategory) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinJumpListCategory5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)
func (this *QWinJumpListCategory) SetTitle(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinJumpListCategory8setTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QWinJumpListCategory) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinJumpListCategory5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QWinJumpListCategory) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinJumpListCategory7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addItem(QWinJumpListItem *)
func (this *QWinJumpListCategory) AddItem(item *QWinJumpListItem /*777 QWinJumpListItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinJumpListCategory7addItemEP16QWinJumpListItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QWinJumpListItem * addDestination(const QString &)
func (this *QWinJumpListCategory) AddDestination(filePath string) *QWinJumpListItem /*777 QWinJumpListItem **/ {
	var tmpArg0 = qtcore.NewQString_5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinJumpListCategory14addDestinationERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWinJumpListItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QWinJumpListItem * addLink(const QString &, const QString &, const QStringList &)
func (this *QWinJumpListCategory) AddLink(title string, executablePath string, arguments *qtcore.QStringList) *QWinJumpListItem /*777 QWinJumpListItem **/ {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(executablePath)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = arguments.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinJumpListCategory7addLinkERK7QStringS2_RK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWinJumpListItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:82
// index:1
// Public Visibility=Default Availability=Available
// [8] QWinJumpListItem * addLink(const QIcon &, const QString &, const QString &, const QStringList &)
func (this *QWinJumpListCategory) AddLink_1(icon *qtgui.QIcon, title string, executablePath string, arguments *qtcore.QStringList) *QWinJumpListItem /*777 QWinJumpListItem **/ {
	var convArg0 = icon.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(executablePath)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 = arguments.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinJumpListCategory7addLinkERK5QIconRK7QStringS5_RK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWinJumpListItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] QWinJumpListItem * addSeparator()
func (this *QWinJumpListCategory) AddSeparator() *QWinJumpListItem /*777 QWinJumpListItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinJumpListCategory12addSeparatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWinJumpListItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistcategory.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QWinJumpListCategory) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinJumpListCategory5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

type QWinJumpListCategory__Type = int

const QWinJumpListCategory__Custom QWinJumpListCategory__Type = 0
const QWinJumpListCategory__Recent QWinJumpListCategory__Type = 1
const QWinJumpListCategory__Frequent QWinJumpListCategory__Type = 2
const QWinJumpListCategory__Tasks QWinJumpListCategory__Type = 3

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
