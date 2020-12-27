package qtgui

// /usr/include/qt/QtGui/qtextoption.h
// #include <qtextoption.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*
 */
// size 32
type QTextOption struct {
	*qtrt.CObject
}
type QTextOption_ITF interface {
	QTextOption_PTR() *QTextOption
}

func (ptr *QTextOption) QTextOption_PTR() *QTextOption { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QTextOptionFromptr(cthis Voidptr) *QTextOption {
	return &QTextOption{&qtrt.CObject{cthis}}
}
func (*QTextOption) Fromptr(cthis Voidptr) *QTextOption {
	return QTextOptionFromptr(cthis)
}

// /usr/include/qt/QtGui/qtextoption.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextOption()

/*
 */
func (*QTextOption) NewForInherit() *QTextOption {
	return NewQTextOption()
}
func NewQTextOption() *QTextOption {
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc3(1571547807, "_ZN11QTextOptionC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint3(err, rv)
	gothis := QTextOptionFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQTextOption)
	return gothis
}

func DeleteQTextOption(this *QTextOption) {
	rv, err := qtrt.Qtcc3(3229405734, "_ZN11QTextOptionD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QTextOption__TabType = int

//
const QTextOption__LeftTab QTextOption__TabType = 0

//
const QTextOption__RightTab QTextOption__TabType = 1

//
const QTextOption__CenterTab QTextOption__TabType = 2

//
const QTextOption__DelimiterTab QTextOption__TabType = 3

func (this *QTextOption) TabTypeItemName(val int) string {
	switch val {
	case QTextOption__LeftTab: // 0
		return "LeftTab"
	case QTextOption__RightTab: // 1
		return "RightTab"
	case QTextOption__CenterTab: // 2
		return "CenterTab"
	case QTextOption__DelimiterTab: // 3
		return "DelimiterTab"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextOption_TabTypeItemName(val int) string {
	var nilthis *QTextOption
	return nilthis.TabTypeItemName(val)
}

/*


 */
type QTextOption__WrapMode = int

//
const QTextOption__NoWrap QTextOption__WrapMode = 0

//
const QTextOption__WordWrap QTextOption__WrapMode = 1

//
const QTextOption__ManualWrap QTextOption__WrapMode = 2

//
const QTextOption__WrapAnywhere QTextOption__WrapMode = 3

//
const QTextOption__WrapAtWordBoundaryOrAnywhere QTextOption__WrapMode = 4

func (this *QTextOption) WrapModeItemName(val int) string {
	switch val {
	case QTextOption__NoWrap: // 0
		return "NoWrap"
	case QTextOption__WordWrap: // 1
		return "WordWrap"
	case QTextOption__ManualWrap: // 2
		return "ManualWrap"
	case QTextOption__WrapAnywhere: // 3
		return "WrapAnywhere"
	case QTextOption__WrapAtWordBoundaryOrAnywhere: // 4
		return "WrapAtWordBoundaryOrAnywhere"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextOption_WrapModeItemName(val int) string {
	var nilthis *QTextOption
	return nilthis.WrapModeItemName(val)
}

/*


 */
type QTextOption__Flag = int

//
const QTextOption__ShowTabsAndSpaces QTextOption__Flag = 1

//
const QTextOption__ShowLineAndParagraphSeparators QTextOption__Flag = 2

//
const QTextOption__AddSpaceForLineAndParagraphSeparators QTextOption__Flag = 4

//
const QTextOption__SuppressColors QTextOption__Flag = 8

//
const QTextOption__ShowDocumentTerminator QTextOption__Flag = 16

//
const QTextOption__IncludeTrailingSpaces QTextOption__Flag = -2147483648

func (this *QTextOption) FlagItemName(val int) string {
	switch val {
	case QTextOption__ShowTabsAndSpaces: // 1
		return "ShowTabsAndSpaces"
	case QTextOption__ShowLineAndParagraphSeparators: // 2
		return "ShowLineAndParagraphSeparators"
	case QTextOption__AddSpaceForLineAndParagraphSeparators: // 4
		return "AddSpaceForLineAndParagraphSeparators"
	case QTextOption__SuppressColors: // 8
		return "SuppressColors"
	case QTextOption__ShowDocumentTerminator: // 16
		return "ShowDocumentTerminator"
	case QTextOption__IncludeTrailingSpaces: // -2147483648
		return "IncludeTrailingSpaces"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextOption_FlagItemName(val int) string {
	var nilthis *QTextOption
	return nilthis.FlagItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10155() {
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
}

//  keep block end
