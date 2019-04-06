package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QTextImageFormat struct {
	*QTextCharFormat
}
type QTextImageFormat_ITF interface {
	QTextCharFormat_ITF
	QTextImageFormat_PTR() *QTextImageFormat
}

func (ptr *QTextImageFormat) QTextImageFormat_PTR() *QTextImageFormat { return ptr }

func (this *QTextImageFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextCharFormat.GetCthis()
	}
}
func (this *QTextImageFormat) SetCthis(cthis unsafe.Pointer) {
	this.QTextCharFormat = NewQTextCharFormatFromPointer(cthis)
}
func NewQTextImageFormatFromPointer(cthis unsafe.Pointer) *QTextImageFormat {
	bcthis0 := NewQTextCharFormatFromPointer(cthis)
	return &QTextImageFormat{bcthis0}
}
func (*QTextImageFormat) NewFromPointer(cthis unsafe.Pointer) *QTextImageFormat {
	return NewQTextImageFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:742
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextImageFormat()

/*

 */
func (*QTextImageFormat) NewForInherit() *QTextImageFormat {
	return NewQTextImageFormat()
}
func NewQTextImageFormat() *QTextImageFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextImageFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextImageFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextImageFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:763
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QTextImageFormat(const QTextFormat &)

/*

 */
func (*QTextImageFormat) NewForInherit1(format QTextFormat_ITF) *QTextImageFormat {
	return NewQTextImageFormat1(format)
}
func NewQTextImageFormat1(format QTextFormat_ITF) *QTextImageFormat {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextFormat_PTR() != nil {
		convArg0 = format.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextImageFormatC2ERK11QTextFormat", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextImageFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextImageFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:744
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the format is valid (i.e. is not InvalidFormat); otherwise returns false.
*/
func (this *QTextImageFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextImageFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:746
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setName(const QString &)

/*

 */
func (this *QTextImageFormat) SetName(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextImageFormat7setNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:747
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString name() const

/*

 */
func (this *QTextImageFormat) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextImageFormat4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextformat.h:750
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidth(qreal)

/*

 */
func (this *QTextImageFormat) SetWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextImageFormat8setWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:751
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal width() const

/*

 */
func (this *QTextImageFormat) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextImageFormat5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:754
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeight(qreal)

/*

 */
func (this *QTextImageFormat) SetHeight(height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextImageFormat9setHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:755
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal height() const

/*

 */
func (this *QTextImageFormat) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextImageFormat6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:758
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setQuality(int)

/*

 */
func (this *QTextImageFormat) SetQuality(quality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextImageFormat10setQualityEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), quality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:758
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setQuality(int)

/*

 */
func (this *QTextImageFormat) SetQualityp() {
	// arg: 0, int=Int, =Invalid, , Invalid
	quality := int(100)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextImageFormat10setQualityEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), quality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:759
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int quality() const

/*

 */
func (this *QTextImageFormat) Quality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextImageFormat7qualityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQTextImageFormat(this *QTextImageFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextImageFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10769() {
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
