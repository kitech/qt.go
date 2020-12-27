package qtgui

// /usr/include/qt/QtGui/qclipboard.h
// #include <qclipboard.h>
// #include <QtGui>

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

//  ext block end

//  body block begin

/*
 */
// size 16
type QClipboard struct {
	*qtcore.QObject
}
type QClipboard_ITF interface {
	qtcore.QObject_ITF
	QClipboard_PTR() *QClipboard
}

func (ptr *QClipboard) QClipboard_PTR() *QClipboard { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QClipboardFromptr(cthis Voidptr) *QClipboard {
	bcthis0 := qtcore.QObjectFromptr(cthis)
	return &QClipboard{bcthis0}
}
func (*QClipboard) Fromptr(cthis Voidptr) *QClipboard {
	return QClipboardFromptr(cthis)
}

// /usr/include/qt/QtGui/qclipboard.h:65
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clear(QClipboard::Mode)

/*
 */
func (this *QClipboard) Clear(mode int) {
	rv, err := qtrt.Qtcc3(3673203142, "_ZN10QClipboard5clearENS_4ModeE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&mode))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:65
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clear(QClipboard::Mode)

/*
 */
func (this *QClipboard) Clearp() {
	// arg: 0, QClipboard::Mode=Enum, QClipboard::Mode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.Qtcc3(3673203142, "_ZN10QClipboard5clearENS_4ModeE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&mode))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:67
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool supportsSelection() const

/*
 */
func (this *QClipboard) SupportsSelection() bool {
	rv, err := qtrt.Qtcc3(1880697911, "_ZNK10QClipboard17supportsSelectionEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtGui/qclipboard.h:68
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool supportsFindBuffer() const

/*
 */
func (this *QClipboard) SupportsFindBuffer() bool {
	rv, err := qtrt.Qtcc3(2133879840, "_ZNK10QClipboard18supportsFindBufferEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtGui/qclipboard.h:70
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool ownsSelection() const

/*
 */
func (this *QClipboard) OwnsSelection() bool {
	rv, err := qtrt.Qtcc3(317563336, "_ZNK10QClipboard13ownsSelectionEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtGui/qclipboard.h:71
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool ownsClipboard() const

/*
 */
func (this *QClipboard) OwnsClipboard() bool {
	rv, err := qtrt.Qtcc3(1848675440, "_ZNK10QClipboard13ownsClipboardEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtGui/qclipboard.h:72
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool ownsFindBuffer() const

/*
 */
func (this *QClipboard) OwnsFindBuffer() bool {
	rv, err := qtrt.Qtcc3(1054311493, "_ZNK10QClipboard14ownsFindBufferEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

func DeleteQClipboard(this *QClipboard) {
	rv, err := qtrt.Qtcc3(4282311273, "_ZN10QClipboardD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QClipboard__Mode = int

//
const QClipboard__Clipboard QClipboard__Mode = 0

//
const QClipboard__Selection QClipboard__Mode = 1

//
const QClipboard__FindBuffer QClipboard__Mode = 2

//
const QClipboard__LastMode QClipboard__Mode = 2

func (this *QClipboard) ModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QClipboard_ModeItemName(val int) string {
	var nilthis *QClipboard
	return nilthis.ModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10161() {
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
