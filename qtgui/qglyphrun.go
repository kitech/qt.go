package qtgui

// /usr/include/qt/QtGui/qglyphrun.h
// #include <qglyphrun.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
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
type QGlyphRun struct {
	*qtrt.CObject
}
type QGlyphRun_ITF interface {
	QGlyphRun_PTR() *QGlyphRun
}

func (ptr *QGlyphRun) QGlyphRun_PTR() *QGlyphRun { return ptr }

func (this *QGlyphRun) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGlyphRun) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGlyphRunFromPointer(cthis unsafe.Pointer) *QGlyphRun {
	return &QGlyphRun{&qtrt.CObject{cthis}}
}
func (*QGlyphRun) NewFromPointer(cthis unsafe.Pointer) *QGlyphRun {
	return NewQGlyphRunFromPointer(cthis)
}

// /usr/include/qt/QtGui/qglyphrun.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGlyphRun()

/*
Constructs an empty QGlyphRun object.
*/
func NewQGlyphRun() *QGlyphRun {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRunC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGlyphRunFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGlyphRun)
	return gothis
}

// /usr/include/qt/QtGui/qglyphrun.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QGlyphRun & operator=(QGlyphRun &&)

/*

 */
func (this *QGlyphRun) Operator_equal(other unsafe.Pointer /*333*/) *QGlyphRun {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRunaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGlyphRunFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGlyphRun)
	return rv2
}

// /usr/include/qt/QtGui/qglyphrun.h:72
// index:1
// Public Visibility=Default Availability=Available
// [8] QGlyphRun & operator=(const QGlyphRun &)

/*

 */
func (this *QGlyphRun) Operator_equal_1(other QGlyphRun_ITF) *QGlyphRun {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGlyphRun_PTR() != nil {
		convArg0 = other.QGlyphRun_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRunaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGlyphRunFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGlyphRun)
	return rv2
}

// /usr/include/qt/QtGui/qglyphrun.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGlyphRun()

/*

 */
func DeleteQGlyphRun(this *QGlyphRun) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRunD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qglyphrun.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QGlyphRun &)

/*
Swaps this glyph run instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QGlyphRun) Swap(other QGlyphRun_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGlyphRun_PTR() != nil {
		convArg0 = other.QGlyphRun_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QRawFont rawFont() const

/*
Returns the font selected for this QGlyphRun object.

See also setRawFont().
*/
func (this *QGlyphRun) RawFont() *QRawFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun7rawFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRawFont)
	return rv2
}

// /usr/include/qt/QtGui/qglyphrun.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawFont(const QRawFont &)

/*
Sets the font in which to look up the glyph indexes to the rawFont specified.

See also rawFont() and setGlyphIndexes().
*/
func (this *QGlyphRun) SetRawFont(rawFont QRawFont_ITF) {
	var convArg0 unsafe.Pointer
	if rawFont != nil && rawFont.QRawFont_PTR() != nil {
		convArg0 = rawFont.QRawFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun10setRawFontERK8QRawFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawData(const quint32 *, const QPointF *, int)

/*
Sets the glyph indexes and positions of this QGlyphRun to use the first size elements in the arrays glyphIndexArray and glyphPositionArray. The data is not copied. The caller must guarantee that the arrays are not deleted as long as this QGlyphRun and any copies of it exists.

See also setGlyphIndexes() and setPositions().
*/
func (this *QGlyphRun) SetRawData(glyphIndexArray unsafe.Pointer /*666*/, glyphPositionArray qtcore.QPointF_ITF /*777 const QPointF **/, size int) {
	var convArg1 unsafe.Pointer
	if glyphPositionArray != nil && glyphPositionArray.QPointF_PTR() != nil {
		convArg1 = glyphPositionArray.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun10setRawDataEPKjPK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndexArray, convArg1, size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears all data in the QGlyphRun object.
*/
func (this *QGlyphRun) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGlyphRun &) const

/*

 */
func (this *QGlyphRun) Operator_equal_equal(other QGlyphRun_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGlyphRun_PTR() != nil {
		convArg0 = other.QGlyphRun_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRuneqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qglyphrun.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QGlyphRun &) const

/*

 */
func (this *QGlyphRun) Operator_not_equal(other QGlyphRun_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGlyphRun_PTR() != nil {
		convArg0 = other.QGlyphRun_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRunneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qglyphrun.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOverline(_Bool)

/*
Indicates that this QGlyphRun should be painted with an overline decoration if overline is true. Otherwise the QGlyphRun should be painted with no overline decoration.

See also overline(), setFlag(), and setFlags().
*/
func (this *QGlyphRun) SetOverline(overline bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun11setOverlineEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), overline)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool overline() const

/*
Returns true if this QGlyphRun should be painted with an overline decoration.

See also setOverline() and flags().
*/
func (this *QGlyphRun) Overline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun8overlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qglyphrun.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUnderline(_Bool)

/*
Indicates that this QGlyphRun should be painted with an underline decoration if underline is true. Otherwise the QGlyphRun should be painted with no underline decoration.

See also underline(), setFlag(), and setFlags().
*/
func (this *QGlyphRun) SetUnderline(underline bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun12setUnderlineEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), underline)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool underline() const

/*
Returns true if this QGlyphRun should be painted with an underline decoration.

See also setUnderline() and flags().
*/
func (this *QGlyphRun) Underline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun9underlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qglyphrun.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStrikeOut(_Bool)

/*
Indicates that this QGlyphRun should be painted with an strike out decoration if strikeOut is true. Otherwise the QGlyphRun should be painted with no strike out decoration.

See also strikeOut(), setFlag(), and setFlags().
*/
func (this *QGlyphRun) SetStrikeOut(strikeOut bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun12setStrikeOutEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), strikeOut)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool strikeOut() const

/*
Returns true if this QGlyphRun should be painted with a strike out decoration.

See also setStrikeOut() and flags().
*/
func (this *QGlyphRun) StrikeOut() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun9strikeOutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qglyphrun.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRightToLeft(_Bool)

/*
Indicates that this QGlyphRun contains glyphs that should be ordered from the right to left if rightToLeft is true. Otherwise the order of the glyphs is assumed to be left to right.

This function was introduced in  Qt 5.0.

See also isRightToLeft(), setFlag(), and setFlags().
*/
func (this *QGlyphRun) SetRightToLeft(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun14setRightToLeftEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRightToLeft() const

/*
Returns true if this QGlyphRun contains glyphs that are painted from the right to the left.

This function was introduced in  Qt 5.0.

See also setRightToLeft() and flags().
*/
func (this *QGlyphRun) IsRightToLeft() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun13isRightToLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qglyphrun.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(enum QGlyphRun::GlyphRunFlag, _Bool)

/*
If enabled is true, then flag is enabled; otherwise, it is disabled.

This function was introduced in  Qt 5.0.

See also flags() and setFlags().
*/
func (this *QGlyphRun) SetFlag(flag int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun7setFlagENS_12GlyphRunFlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(enum QGlyphRun::GlyphRunFlag, _Bool)

/*
If enabled is true, then flag is enabled; otherwise, it is disabled.

This function was introduced in  Qt 5.0.

See also flags() and setFlags().
*/
func (this *QGlyphRun) SetFlag__(flag int) {
	// arg: 1, bool=Bool, =Invalid,
	enabled := true
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun7setFlagENS_12GlyphRunFlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QGlyphRun::GlyphRunFlags)

/*
Sets the flags of this QGlyphRun to flags.

This function was introduced in  Qt 5.0.

See also setFlag() and flags().
*/
func (this *QGlyphRun) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun8setFlagsE6QFlagsINS_12GlyphRunFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] QGlyphRun::GlyphRunFlags flags() const

/*
Returns the flags set for this QGlyphRun.

This function was introduced in  Qt 5.0.

See also setFlags(), setFlag(), and setFlag().
*/
func (this *QGlyphRun) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBoundingRect(const QRectF &)

/*
Sets the bounding rect of the glyphs in this QGlyphRun to be boundingRect. This rectangle will be returned by boundingRect() unless it is empty, in which case the bounding rectangle of the glyphs in the glyph run will be returned instead.

Note: Unless you are implementing text shaping, you should not have to use this function. It is used specifically when the QGlyphRun should represent an area which is smaller than the area of the glyphs it contains. This could happen e.g. if the glyph run is retrieved by calling QTextLayout::glyphRuns() and the specified range only includes part of a ligature (where two or more characters are combined to a single glyph.) When this is the case, the bounding rect should only include the appropriate part of the ligature glyph, based on a calculation of the average width of the characters in the ligature.

In order to support such a case (an example is selections which should be drawn with a different color than the main text color), it is necessary to clip the painting mechanism to the rectangle returned from boundingRect() to avoid drawing the entire ligature glyph.

This function was introduced in  Qt 5.0.

See also boundingRect().
*/
func (this *QGlyphRun) SetBoundingRect(boundingRect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if boundingRect != nil && boundingRect.QRectF_PTR() != nil {
		convArg0 = boundingRect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun15setBoundingRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:113
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect() const

/*
Returns the smallest rectangle that contains all glyphs in this QGlyphRun. If a bounding rect has been set using setBoundingRect(), then this will be returned. Otherwise the bounding rect will be calculated based on the font metrics of the glyphs in the glyph run.

This function was introduced in  Qt 5.0.

See also setBoundingRect().
*/
func (this *QGlyphRun) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qglyphrun.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the QGlyphRun does not contain any glyphs.

This function was introduced in  Qt 5.0.
*/
func (this *QGlyphRun) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QGlyphRun__GlyphRunFlag = int

//
const QGlyphRun__Overline QGlyphRun__GlyphRunFlag = 1

//
const QGlyphRun__Underline QGlyphRun__GlyphRunFlag = 2

//
const QGlyphRun__StrikeOut QGlyphRun__GlyphRunFlag = 4

//
const QGlyphRun__RightToLeft QGlyphRun__GlyphRunFlag = 8

//
const QGlyphRun__SplitLigature QGlyphRun__GlyphRunFlag = 16

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
