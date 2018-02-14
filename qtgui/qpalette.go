package qtgui

// /usr/include/qt/QtGui/qpalette.h
// #include <qpalette.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 95
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

type QPalette struct {
	*qtrt.CObject
}
type QPalette_ITF interface {
	QPalette_PTR() *QPalette
}

func (ptr *QPalette) QPalette_PTR() *QPalette { return ptr }

func (this *QPalette) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPalette) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPaletteFromPointer(cthis unsafe.Pointer) *QPalette {
	return &QPalette{&qtrt.CObject{cthis}}
}
func (*QPalette) NewFromPointer(cthis unsafe.Pointer) *QPalette {
	return NewQPaletteFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpalette.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPalette()
func NewQPalette() *QPalette {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPaletteC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPalette(const QColor &)
func NewQPalette_1(button QColor_ITF) *QPalette {
	var convArg0 unsafe.Pointer
	if button != nil && button.QColor_PTR() != nil {
		convArg0 = button.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QColor", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:60
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPalette(Qt::GlobalColor)
func NewQPalette_2(button int) *QPalette {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPaletteC2EN2Qt11GlobalColorE", qtrt.FFI_TYPE_POINTER, button)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:61
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPalette(const QColor &, const QColor &)
func NewQPalette_3(button QColor_ITF, window QColor_ITF) *QPalette {
	var convArg0 unsafe.Pointer
	if button != nil && button.QColor_PTR() != nil {
		convArg0 = button.QColor_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if window != nil && window.QColor_PTR() != nil {
		convArg1 = window.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QColorS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:62
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QPalette(const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &)
func NewQPalette_4(windowText QBrush_ITF, button QBrush_ITF, light QBrush_ITF, dark QBrush_ITF, mid QBrush_ITF, text QBrush_ITF, bright_text QBrush_ITF, base QBrush_ITF, window QBrush_ITF) *QPalette {
	var convArg0 unsafe.Pointer
	if windowText != nil && windowText.QBrush_PTR() != nil {
		convArg0 = windowText.QBrush_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if button != nil && button.QBrush_PTR() != nil {
		convArg1 = button.QBrush_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if light != nil && light.QBrush_PTR() != nil {
		convArg2 = light.QBrush_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if dark != nil && dark.QBrush_PTR() != nil {
		convArg3 = dark.QBrush_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if mid != nil && mid.QBrush_PTR() != nil {
		convArg4 = mid.QBrush_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if text != nil && text.QBrush_PTR() != nil {
		convArg5 = text.QBrush_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if bright_text != nil && bright_text.QBrush_PTR() != nil {
		convArg6 = bright_text.QBrush_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if base != nil && base.QBrush_PTR() != nil {
		convArg7 = base.QBrush_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if window != nil && window.QBrush_PTR() != nil {
		convArg8 = window.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QBrushS2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:65
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QPalette(const QColor &, const QColor &, const QColor &, const QColor &, const QColor &, const QColor &, const QColor &)
func NewQPalette_5(windowText QColor_ITF, window QColor_ITF, light QColor_ITF, dark QColor_ITF, mid QColor_ITF, text QColor_ITF, base QColor_ITF) *QPalette {
	var convArg0 unsafe.Pointer
	if windowText != nil && windowText.QColor_PTR() != nil {
		convArg0 = windowText.QColor_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if window != nil && window.QColor_PTR() != nil {
		convArg1 = window.QColor_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if light != nil && light.QColor_PTR() != nil {
		convArg2 = light.QColor_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if dark != nil && dark.QColor_PTR() != nil {
		convArg3 = dark.QColor_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if mid != nil && mid.QColor_PTR() != nil {
		convArg4 = mid.QColor_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if text != nil && text.QColor_PTR() != nil {
		convArg5 = text.QColor_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if base != nil && base.QColor_PTR() != nil {
		convArg6 = base.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QColorS2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPalette()
func DeleteQPalette(this *QPalette) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPaletteD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpalette.h:69
// index:0
// Public Visibility=Default Availability=Available
// [16] QPalette & operator=(const QPalette &)
func (this *QPalette) Operator_equal(palette QPalette_ITF) *QPalette {
	var convArg0 unsafe.Pointer
	if palette != nil && palette.QPalette_PTR() != nil {
		convArg0 = palette.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPaletteaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPalette)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:74
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QPalette & operator=(QPalette &&)
func (this *QPalette) Operator_equal_1(other unsafe.Pointer /*333*/) *QPalette {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPaletteaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPalette)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPalette &)
func (this *QPalette) Swap(other QPalette_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPalette_PTR() != nil {
		convArg0 = other.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPalette4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:104
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPalette::ColorGroup currentColorGroup()
func (this *QPalette) CurrentColorGroup() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette17currentColorGroupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpalette.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCurrentColorGroup(enum QPalette::ColorGroup)
func (this *QPalette) SetCurrentColorGroup(cg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPalette20setCurrentColorGroupENS_10ColorGroupE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QColor & color(enum QPalette::ColorGroup, enum QPalette::ColorRole)
func (this *QPalette) Color(cg int, cr int) *QColor {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette5colorENS_10ColorGroupENS_9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cg, cr)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:121
// index:1
// Public inline Visibility=Default Availability=Available
// [16] const QColor & color(enum QPalette::ColorRole)
func (this *QPalette) Color_1(cr int) *QColor {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette5colorENS_9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cr)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] const QBrush & brush(enum QPalette::ColorGroup, enum QPalette::ColorRole)
func (this *QPalette) Brush(cg int, cr int) *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette5brushENS_10ColorGroupENS_9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cg, cr)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:122
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & brush(enum QPalette::ColorRole)
func (this *QPalette) Brush_1(cr int) *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette5brushENS_9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cr)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setColor(enum QPalette::ColorGroup, enum QPalette::ColorRole, const QColor &)
func (this *QPalette) SetColor(cg int, cr int, color QColor_ITF) {
	var convArg2 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg2 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPalette8setColorENS_10ColorGroupENS_9ColorRoleERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cg, cr, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:111
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setColor(enum QPalette::ColorRole, const QColor &)
func (this *QPalette) SetColor_1(cr int, color QColor_ITF) {
	var convArg1 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg1 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPalette8setColorENS_9ColorRoleERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cr, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBrush(enum QPalette::ColorRole, const QBrush &)
func (this *QPalette) SetBrush(cr int, brush QBrush_ITF) {
	var convArg1 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg1 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPalette8setBrushENS_9ColorRoleERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cr, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:114
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setBrush(enum QPalette::ColorGroup, enum QPalette::ColorRole, const QBrush &)
func (this *QPalette) SetBrush_1(cg int, cr int, brush QBrush_ITF) {
	var convArg2 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg2 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPalette8setBrushENS_10ColorGroupENS_9ColorRoleERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cg, cr, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBrushSet(enum QPalette::ColorGroup, enum QPalette::ColorRole)
func (this *QPalette) IsBrushSet(cg int, cr int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette10isBrushSetENS_10ColorGroupENS_9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cg, cr)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpalette.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColorGroup(enum QPalette::ColorGroup, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &)
func (this *QPalette) SetColorGroup(cr int, windowText QBrush_ITF, button QBrush_ITF, light QBrush_ITF, dark QBrush_ITF, mid QBrush_ITF, text QBrush_ITF, bright_text QBrush_ITF, base QBrush_ITF, window QBrush_ITF) {
	var convArg1 unsafe.Pointer
	if windowText != nil && windowText.QBrush_PTR() != nil {
		convArg1 = windowText.QBrush_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if button != nil && button.QBrush_PTR() != nil {
		convArg2 = button.QBrush_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if light != nil && light.QBrush_PTR() != nil {
		convArg3 = light.QBrush_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if dark != nil && dark.QBrush_PTR() != nil {
		convArg4 = dark.QBrush_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if mid != nil && mid.QBrush_PTR() != nil {
		convArg5 = mid.QBrush_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if text != nil && text.QBrush_PTR() != nil {
		convArg6 = text.QBrush_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if bright_text != nil && bright_text.QBrush_PTR() != nil {
		convArg7 = bright_text.QBrush_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if base != nil && base.QBrush_PTR() != nil {
		convArg8 = base.QBrush_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if window != nil && window.QBrush_PTR() != nil {
		convArg9 = window.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPalette13setColorGroupENS_10ColorGroupERK6QBrushS3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cr, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:119
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEqual(enum QPalette::ColorGroup, enum QPalette::ColorGroup)
func (this *QPalette) IsEqual(cr1 int, cr2 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette7isEqualENS_10ColorGroupES0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cr1, cr2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpalette.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & foreground()
func (this *QPalette) Foreground() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette10foregroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & windowText()
func (this *QPalette) WindowText() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette10windowTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & button()
func (this *QPalette) Button() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette6buttonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & light()
func (this *QPalette) Light() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette5lightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & dark()
func (this *QPalette) Dark() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette4darkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & mid()
func (this *QPalette) Mid() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette3midEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & text()
func (this *QPalette) Text() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & base()
func (this *QPalette) Base() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette4baseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:131
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & alternateBase()
func (this *QPalette) AlternateBase() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette13alternateBaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & toolTipBase()
func (this *QPalette) ToolTipBase() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette11toolTipBaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & toolTipText()
func (this *QPalette) ToolTipText() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette11toolTipTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:134
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & background()
func (this *QPalette) Background() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette10backgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & window()
func (this *QPalette) Window() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & midlight()
func (this *QPalette) Midlight() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette8midlightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:137
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & brightText()
func (this *QPalette) BrightText() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette10brightTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:138
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & buttonText()
func (this *QPalette) ButtonText() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette10buttonTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:139
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & shadow()
func (this *QPalette) Shadow() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette6shadowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & highlight()
func (this *QPalette) Highlight() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette9highlightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:141
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & highlightedText()
func (this *QPalette) HighlightedText() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette15highlightedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:142
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & link()
func (this *QPalette) Link() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette4linkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:143
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & linkVisited()
func (this *QPalette) LinkVisited() *QBrush {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette11linkVisitedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:145
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QPalette &)
func (this *QPalette) Operator_equal_equal(p QPalette_ITF) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPalette_PTR() != nil {
		convArg0 = p.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPaletteeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpalette.h:146
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QPalette &)
func (this *QPalette) Operator_not_equal(p QPalette_ITF) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPalette_PTR() != nil {
		convArg0 = p.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPaletteneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpalette.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCopyOf(const QPalette &)
func (this *QPalette) IsCopyOf(p QPalette_ITF) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPalette_PTR() != nil {
		convArg0 = p.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette8isCopyOfERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpalette.h:152
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 cacheKey()
func (this *QPalette) CacheKey() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette8cacheKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qpalette.h:154
// index:0
// Public Visibility=Default Availability=Available
// [16] QPalette resolve(const QPalette &)
func (this *QPalette) Resolve(arg0 QPalette_ITF) *QPalette /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPalette_PTR() != nil {
		convArg0 = arg0.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette7resolveERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPalette)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:155
// index:1
// Public inline Visibility=Default Availability=Available
// [4] uint resolve()
func (this *QPalette) Resolve_1() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPalette7resolveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qpalette.h:156
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void resolve(uint)
func (this *QPalette) Resolve_2(mask uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPalette7resolveEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mask)
	qtrt.ErrPrint(err, rv)
}

type QPalette__ColorGroup = int

const QPalette__Active QPalette__ColorGroup = 0
const QPalette__Disabled QPalette__ColorGroup = 1
const QPalette__Inactive QPalette__ColorGroup = 2
const QPalette__NColorGroups QPalette__ColorGroup = 3
const QPalette__Current QPalette__ColorGroup = 4
const QPalette__All QPalette__ColorGroup = 5
const QPalette__Normal QPalette__ColorGroup = 0

type QPalette__ColorRole = int

const QPalette__WindowText QPalette__ColorRole = 0
const QPalette__Button QPalette__ColorRole = 1
const QPalette__Light QPalette__ColorRole = 2
const QPalette__Midlight QPalette__ColorRole = 3
const QPalette__Dark QPalette__ColorRole = 4
const QPalette__Mid QPalette__ColorRole = 5
const QPalette__Text QPalette__ColorRole = 6
const QPalette__BrightText QPalette__ColorRole = 7
const QPalette__ButtonText QPalette__ColorRole = 8
const QPalette__Base QPalette__ColorRole = 9
const QPalette__Window QPalette__ColorRole = 10
const QPalette__Shadow QPalette__ColorRole = 11
const QPalette__Highlight QPalette__ColorRole = 12
const QPalette__HighlightedText QPalette__ColorRole = 13
const QPalette__Link QPalette__ColorRole = 14
const QPalette__LinkVisited QPalette__ColorRole = 15
const QPalette__AlternateBase QPalette__ColorRole = 16
const QPalette__NoRole QPalette__ColorRole = 17
const QPalette__ToolTipBase QPalette__ColorRole = 18
const QPalette__ToolTipText QPalette__ColorRole = 19
const QPalette__NColorRoles QPalette__ColorRole = 20
const QPalette__Foreground QPalette__ColorRole = 0
const QPalette__Background QPalette__ColorRole = 10

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
