package qtgui

// /usr/include/qt/QtGui/qpaintengine.h
// #include <qpaintengine.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 206
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
type QTextItem struct {
	*qtrt.CObject
}
type QTextItem_ITF interface {
	QTextItem_PTR() *QTextItem
}

func (ptr *QTextItem) QTextItem_PTR() *QTextItem { return ptr }

func (this *QTextItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextItemFromPointer(cthis unsafe.Pointer) *QTextItem {
	return &QTextItem{&qtrt.CObject{cthis}}
}
func (*QTextItem) NewFromPointer(cthis unsafe.Pointer) *QTextItem {
	return NewQTextItemFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpaintengine.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal descent() const

/*

 */
func (this *QTextItem) Descent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem7descentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpaintengine.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal ascent() const

/*

 */
func (this *QTextItem) Ascent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem6ascentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpaintengine.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal width() const

/*

 */
func (this *QTextItem) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpaintengine.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextItem::RenderFlags renderFlags() const

/*

 */
func (this *QTextItem) RenderFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem11renderFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpaintengine.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*

 */
func (this *QTextItem) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qpaintengine.h:81
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font() const

/*

 */
func (this *QTextItem) Font() *QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextItem4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

func DeleteQTextItem(this *QTextItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QTextItem__RenderFlag = int

//
const QTextItem__RightToLeft QTextItem__RenderFlag = 1

//
const QTextItem__Overline QTextItem__RenderFlag = 16

//
const QTextItem__Underline QTextItem__RenderFlag = 32

//
const QTextItem__StrikeOut QTextItem__RenderFlag = 64

//
const QTextItem__Dummy QTextItem__RenderFlag = -1

func (this *QTextItem) RenderFlagItemName(val int) string {
	switch val {
	case QTextItem__RightToLeft: // 1
		return "RightToLeft"
	case QTextItem__Overline: // 16
		return "Overline"
	case QTextItem__Underline: // 32
		return "Underline"
	case QTextItem__StrikeOut: // 64
		return "StrikeOut"
	case QTextItem__Dummy: // -1
		return "Dummy"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextItem_RenderFlagItemName(val int) string {
	var nilthis *QTextItem
	return nilthis.RenderFlagItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10913() {
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
