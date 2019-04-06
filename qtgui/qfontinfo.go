package qtgui

// /usr/include/qt/QtGui/qfontinfo.h
// #include <qfontinfo.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
type QFontInfo struct {
	*qtrt.CObject
}
type QFontInfo_ITF interface {
	QFontInfo_PTR() *QFontInfo
}

func (ptr *QFontInfo) QFontInfo_PTR() *QFontInfo { return ptr }

func (this *QFontInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFontInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFontInfoFromPointer(cthis unsafe.Pointer) *QFontInfo {
	return &QFontInfo{&qtrt.CObject{cthis}}
}
func (*QFontInfo) NewFromPointer(cthis unsafe.Pointer) *QFontInfo {
	return NewQFontInfoFromPointer(cthis)
}

// /usr/include/qt/QtGui/qfontinfo.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontInfo(const QFont &)

/*
Constructs a font info object for font.

The font must be screen-compatible, i.e. a font you use when drawing text in widgets or pixmaps, not QPicture or QPrinter.

The font info object holds the information for the font that is passed in the constructor at the time it is created, and is not updated if the font's attributes are changed later.

Use QPainter::fontInfo() to get the font info when painting. This will give correct results also when painting on paint device that is not screen-compatible.
*/
func (*QFontInfo) NewForInherit(arg0 QFont_ITF) *QFontInfo {
	return NewQFontInfo(arg0)
}
func NewQFontInfo(arg0 QFont_ITF) *QFontInfo {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFontInfoC2ERK5QFont", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontInfo)
	return gothis
}

// /usr/include/qt/QtGui/qfontinfo.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QFontInfo()

/*

 */
func DeleteQFontInfo(this *QFontInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFontInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qfontinfo.h:57
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontInfo & operator=(const QFontInfo &)

/*

 */
func (this *QFontInfo) Operator_equal(arg0 QFontInfo_ITF) *QFontInfo {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFontInfo_PTR() != nil {
		convArg0 = arg0.QFontInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFontInfoaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontInfo)
	return rv2
}

// /usr/include/qt/QtGui/qfontinfo.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QFontInfo &)

/*
Swaps this font info instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QFontInfo) Swap(other QFontInfo_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFontInfo_PTR() != nil {
		convArg0 = other.QFontInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFontInfo4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QString family() const

/*
Returns the family name of the matched window system font.

See also QFont::family().
*/
func (this *QFontInfo) Family() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo6familyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontinfo.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QString styleName() const

/*
Returns the style name of the matched window system font on systems that support it.

This function was introduced in  Qt 4.8.

See also QFont::styleName().
*/
func (this *QFontInfo) StyleName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9styleNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontinfo.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int pixelSize() const

/*
Returns the pixel size of the matched window system font.

See also QFont::pointSize().
*/
func (this *QFontInfo) PixelSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9pixelSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontinfo.h:64
// index:0
// Public Visibility=Default Availability=Available
// [4] int pointSize() const

/*
Returns the point size of the matched window system font.

See also pointSizeF() and QFont::pointSize().
*/
func (this *QFontInfo) PointSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9pointSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontinfo.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal pointSizeF() const

/*
Returns the point size of the matched window system font.

See also QFont::pointSizeF().
*/
func (this *QFontInfo) PointSizeF() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo10pointSizeFEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontinfo.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool italic() const

/*
Returns the italic value of the matched window system font.

See also QFont::italic().
*/
func (this *QFontInfo) Italic() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo6italicEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::Style style() const

/*
Returns the style value of the matched window system font.

See also QFont::style().
*/
func (this *QFontInfo) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] int weight() const

/*
Returns the weight of the matched window system font.

See also QFont::weight() and bold().
*/
func (this *QFontInfo) Weight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo6weightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontinfo.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool bold() const

/*
Returns true if weight() would return a value greater than QFont::Normal; otherwise returns false.

See also weight() and QFont::bold().
*/
func (this *QFontInfo) Bold() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo4boldEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool underline() const

/*

 */
func (this *QFontInfo) Underline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9underlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool overline() const

/*

 */
func (this *QFontInfo) Overline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo8overlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool strikeOut() const

/*

 */
func (this *QFontInfo) StrikeOut() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9strikeOutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool fixedPitch() const

/*
Returns the fixed pitch value of the matched window system font.

See also QFont::fixedPitch().
*/
func (this *QFontInfo) FixedPitch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo10fixedPitchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::StyleHint styleHint() const

/*
Returns the style of the matched window system font.

Currently only returns the style hint set in QFont.

See also QFont::styleHint() and QFont::StyleHint.
*/
func (this *QFontInfo) StyleHint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9styleHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rawMode() const

/*

 */
func (this *QFontInfo) RawMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo7rawModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exactMatch() const

/*
Returns true if the matched window system font is exactly the same as the one specified by the font; otherwise returns false.

See also QFont::exactMatch().
*/
func (this *QFontInfo) ExactMatch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo10exactMatchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init_unused_10863() {
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
