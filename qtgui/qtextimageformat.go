package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QTextImageFormat struct {
	*QTextCharFormat
}

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

// /usr/include/qt/QtGui/qtextformat.h:735
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextImageFormat()
func NewQTextImageFormat() *QTextImageFormat {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextImageFormatC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextImageFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:752
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QTextImageFormat(const QTextFormat &)
func NewQTextImageFormat_1(format *QTextFormat) *QTextImageFormat {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextImageFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextImageFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:737
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextImageFormat) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextImageFormat7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:739
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setName(const QString &)
func (this *QTextImageFormat) SetName(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextImageFormat7setNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:740
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString name()
func (this *QTextImageFormat) Name() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextImageFormat4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:743
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidth(qreal)
func (this *QTextImageFormat) SetWidth(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextImageFormat8setWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:744
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal width()
func (this *QTextImageFormat) Width() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextImageFormat5widthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:747
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeight(qreal)
func (this *QTextImageFormat) SetHeight(height float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextImageFormat9setHeightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:748
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal height()
func (this *QTextImageFormat) Height() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextImageFormat6heightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

//  body block end
