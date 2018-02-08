package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h
// #include <qwinjumplistitem.h>
// #include <Qtwinextras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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

type QWinJumpListItem struct {
	*qtrt.CObject
}

func (this *QWinJumpListItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWinJumpListItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWinJumpListItemFromPointer(cthis unsafe.Pointer) *QWinJumpListItem {
	return &QWinJumpListItem{&qtrt.CObject{cthis}}
}
func (*QWinJumpListItem) NewFromPointer(cthis unsafe.Pointer) *QWinJumpListItem {
	return NewQWinJumpListItemFromPointer(cthis)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinJumpListItem(enum QWinJumpListItem::Type)
func NewQWinJumpListItem(type_ int) *QWinJumpListItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWinJumpListItemC2ENS_4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQWinJumpListItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWinJumpListItem)
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QWinJumpListItem()
func DeleteQWinJumpListItem(this *QWinJumpListItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWinJumpListItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setType(enum QWinJumpListItem::Type)
func (this *QWinJumpListItem) SetType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWinJumpListItem7setTypeENS_4TypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] QWinJumpListItem::Type type()
func (this *QWinJumpListItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWinJumpListItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilePath(const QString &)
func (this *QWinJumpListItem) SetFilePath(filePath string) {
	var tmpArg0 = qtcore.NewQString_5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWinJumpListItem11setFilePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath()
func (this *QWinJumpListItem) FilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWinJumpListItem8filePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorkingDirectory(const QString &)
func (this *QWinJumpListItem) SetWorkingDirectory(workingDirectory string) {
	var tmpArg0 = qtcore.NewQString_5(workingDirectory)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWinJumpListItem19setWorkingDirectoryERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString workingDirectory()
func (this *QWinJumpListItem) WorkingDirectory() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWinJumpListItem16workingDirectoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QWinJumpListItem) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWinJumpListItem7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QWinJumpListItem) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWinJumpListItem4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)
func (this *QWinJumpListItem) SetTitle(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWinJumpListItem8setTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title()
func (this *QWinJumpListItem) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWinJumpListItem5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDescription(const QString &)
func (this *QWinJumpListItem) SetDescription(description string) {
	var tmpArg0 = qtcore.NewQString_5(description)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWinJumpListItem14setDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QString description()
func (this *QWinJumpListItem) Description() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWinJumpListItem11descriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplistitem.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setArguments(const QStringList &)
func (this *QWinJumpListItem) SetArguments(arguments *qtcore.QStringList) {
	var convArg0 = arguments.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWinJumpListItem12setArgumentsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QWinJumpListItem__Type = int

const QWinJumpListItem__Destination QWinJumpListItem__Type = 0
const QWinJumpListItem__Link QWinJumpListItem__Type = 1
const QWinJumpListItem__Separator QWinJumpListItem__Type = 2

//  body block end
