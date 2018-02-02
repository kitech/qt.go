package qtgui

// /usr/include/qt/QtGui/qstatictext.h
// #include <qstatictext.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 49
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QStaticText struct {
	*qtrt.CObject
}

func (this *QStaticText) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStaticText) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStaticTextFromPointer(cthis unsafe.Pointer) *QStaticText {
	return &QStaticText{&qtrt.CObject{cthis}}
}
func (*QStaticText) NewFromPointer(cthis unsafe.Pointer) *QStaticText {
	return NewQStaticTextFromPointer(cthis)
}

// /usr/include/qt/QtGui/qstatictext.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStaticText()
func NewQStaticText() *QStaticText {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticTextC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStaticTextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStaticText)
	return gothis
}

// /usr/include/qt/QtGui/qstatictext.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStaticText(const QString &)
func NewQStaticText_1(text *qtcore.QString) *QStaticText {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticTextC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStaticTextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStaticText)
	return gothis
}

// /usr/include/qt/QtGui/qstatictext.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QStaticText()
func DeleteQStaticText(this *QStaticText) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticTextD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qstatictext.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QStaticText &)
func (this *QStaticText) Swap(other *QStaticText) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QStaticText) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QStaticText) Text() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qstatictext.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextFormat(Qt::TextFormat)
func (this *QStaticText) SetTextFormat(textFormat int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText13setTextFormatEN2Qt10TextFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), textFormat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextFormat textFormat()
func (this *QStaticText) TextFormat() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText10textFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qstatictext.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextWidth(qreal)
func (this *QStaticText) SetTextWidth(textWidth float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText12setTextWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), textWidth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal textWidth()
func (this *QStaticText) TextWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText9textWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qstatictext.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextOption(const QTextOption &)
func (this *QStaticText) SetTextOption(textOption *QTextOption) {
	var convArg0 = textOption.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText13setTextOptionERK11QTextOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:85
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextOption textOption()
func (this *QStaticText) TextOption() *QTextOption /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText10textOptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextOption)
	return rv2
}

// /usr/include/qt/QtGui/qstatictext.h:87
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size()
func (this *QStaticText) Size() *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qstatictext.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prepare(const QTransform &, const QFont &)
func (this *QStaticText) Prepare(matrix *QTransform, font *QFont) {
	var convArg0 = matrix.GetCthis()
	var convArg1 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText7prepareERK10QTransformRK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPerformanceHint(enum QStaticText::PerformanceHint)
func (this *QStaticText) SetPerformanceHint(performanceHint int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText18setPerformanceHintENS_15PerformanceHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), performanceHint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] QStaticText::PerformanceHint performanceHint()
func (this *QStaticText) PerformanceHint() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText15performanceHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

type QStaticText__PerformanceHint = int

const QStaticText__ModerateCaching QStaticText__PerformanceHint = 0
const QStaticText__AggressiveCaching QStaticText__PerformanceHint = 1

//  body block end
