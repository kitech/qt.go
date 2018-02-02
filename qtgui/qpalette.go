package qtgui

// /usr/include/qt/QtGui/qpalette.h
// #include <qpalette.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 95
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

type QPalette struct {
	*qtrt.CObject
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPalette(const QColor &)
func NewQPalette_1(button *QColor) *QPalette {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QColor", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:60
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPalette(Qt::GlobalColor)
func NewQPalette_2(button int) *QPalette {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2EN2Qt11GlobalColorE", ffiqt.FFI_TYPE_POINTER, button)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:61
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPalette(const QColor &, const QColor &)
func NewQPalette_3(button *QColor, window *QColor) *QPalette {
	var convArg0 = button.GetCthis()
	var convArg1 = window.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QColorS2_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:62
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QPalette(const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &)
func NewQPalette_4(windowText *QBrush, button *QBrush, light *QBrush, dark *QBrush, mid *QBrush, text *QBrush, bright_text *QBrush, base *QBrush, window *QBrush) *QPalette {
	var convArg0 = windowText.GetCthis()
	var convArg1 = button.GetCthis()
	var convArg2 = light.GetCthis()
	var convArg3 = dark.GetCthis()
	var convArg4 = mid.GetCthis()
	var convArg5 = text.GetCthis()
	var convArg6 = bright_text.GetCthis()
	var convArg7 = base.GetCthis()
	var convArg8 = window.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QBrushS2_S2_S2_S2_S2_S2_S2_S2_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:65
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QPalette(const QColor &, const QColor &, const QColor &, const QColor &, const QColor &, const QColor &, const QColor &)
func NewQPalette_5(windowText *QColor, window *QColor, light *QColor, dark *QColor, mid *QColor, text *QColor, base *QColor) *QPalette {
	var convArg0 = windowText.GetCthis()
	var convArg1 = window.GetCthis()
	var convArg2 = light.GetCthis()
	var convArg3 = dark.GetCthis()
	var convArg4 = mid.GetCthis()
	var convArg5 = text.GetCthis()
	var convArg6 = base.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QColorS2_S2_S2_S2_S2_S2_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPalette)
	return gothis
}

// /usr/include/qt/QtGui/qpalette.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPalette()
func DeleteQPalette(this *QPalette) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpalette.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPalette &)
func (this *QPalette) Swap(other *QPalette) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:104
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPalette::ColorGroup currentColorGroup()
func (this *QPalette) CurrentColorGroup() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette17currentColorGroupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpalette.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCurrentColorGroup(enum QPalette::ColorGroup)
func (this *QPalette) SetCurrentColorGroup(cg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette20setCurrentColorGroupENS_10ColorGroupE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QColor & color(enum QPalette::ColorGroup, enum QPalette::ColorRole)
func (this *QPalette) Color(cg int, cr int) *QColor {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette5colorENS_10ColorGroupENS_9ColorRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cg, cr)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:121
// index:1
// Public inline Visibility=Default Availability=Available
// [16] const QColor & color(enum QPalette::ColorRole)
func (this *QPalette) Color_1(cr int) *QColor {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette5colorENS_9ColorRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cr)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] const QBrush & brush(enum QPalette::ColorGroup, enum QPalette::ColorRole)
func (this *QPalette) Brush(cg int, cr int) *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette5brushENS_10ColorGroupENS_9ColorRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cg, cr)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:122
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & brush(enum QPalette::ColorRole)
func (this *QPalette) Brush_1(cr int) *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette5brushENS_9ColorRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cr)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setColor(enum QPalette::ColorGroup, enum QPalette::ColorRole, const QColor &)
func (this *QPalette) SetColor(cg int, cr int, color *QColor) {
	var convArg2 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette8setColorENS_10ColorGroupENS_9ColorRoleERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cg, cr, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:111
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setColor(enum QPalette::ColorRole, const QColor &)
func (this *QPalette) SetColor_1(cr int, color *QColor) {
	var convArg1 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette8setColorENS_9ColorRoleERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cr, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBrush(enum QPalette::ColorRole, const QBrush &)
func (this *QPalette) SetBrush(cr int, brush *QBrush) {
	var convArg1 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette8setBrushENS_9ColorRoleERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cr, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:114
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setBrush(enum QPalette::ColorGroup, enum QPalette::ColorRole, const QBrush &)
func (this *QPalette) SetBrush_1(cg int, cr int, brush *QBrush) {
	var convArg2 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette8setBrushENS_10ColorGroupENS_9ColorRoleERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cg, cr, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBrushSet(enum QPalette::ColorGroup, enum QPalette::ColorRole)
func (this *QPalette) IsBrushSet(cg int, cr int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10isBrushSetENS_10ColorGroupENS_9ColorRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cg, cr)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpalette.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColorGroup(enum QPalette::ColorGroup, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &, const QBrush &)
func (this *QPalette) SetColorGroup(cr int, windowText *QBrush, button *QBrush, light *QBrush, dark *QBrush, mid *QBrush, text *QBrush, bright_text *QBrush, base *QBrush, window *QBrush) {
	var convArg1 = windowText.GetCthis()
	var convArg2 = button.GetCthis()
	var convArg3 = light.GetCthis()
	var convArg4 = dark.GetCthis()
	var convArg5 = mid.GetCthis()
	var convArg6 = text.GetCthis()
	var convArg7 = bright_text.GetCthis()
	var convArg8 = base.GetCthis()
	var convArg9 = window.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette13setColorGroupENS_10ColorGroupERK6QBrushS3_S3_S3_S3_S3_S3_S3_S3_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cr, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:119
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEqual(enum QPalette::ColorGroup, enum QPalette::ColorGroup)
func (this *QPalette) IsEqual(cr1 int, cr2 int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette7isEqualENS_10ColorGroupES0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cr1, cr2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpalette.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & foreground()
func (this *QPalette) Foreground() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10foregroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & windowText()
func (this *QPalette) WindowText() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10windowTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & button()
func (this *QPalette) Button() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette6buttonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & light()
func (this *QPalette) Light() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette5lightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & dark()
func (this *QPalette) Dark() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette4darkEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & mid()
func (this *QPalette) Mid() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette3midEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & text()
func (this *QPalette) Text() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & base()
func (this *QPalette) Base() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette4baseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:131
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & alternateBase()
func (this *QPalette) AlternateBase() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette13alternateBaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & toolTipBase()
func (this *QPalette) ToolTipBase() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette11toolTipBaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & toolTipText()
func (this *QPalette) ToolTipText() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette11toolTipTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:134
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & background()
func (this *QPalette) Background() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10backgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & window()
func (this *QPalette) Window() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & midlight()
func (this *QPalette) Midlight() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette8midlightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:137
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & brightText()
func (this *QPalette) BrightText() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10brightTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:138
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & buttonText()
func (this *QPalette) ButtonText() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10buttonTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:139
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & shadow()
func (this *QPalette) Shadow() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette6shadowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & highlight()
func (this *QPalette) Highlight() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette9highlightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:141
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & highlightedText()
func (this *QPalette) HighlightedText() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette15highlightedTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:142
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & link()
func (this *QPalette) Link() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette4linkEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:143
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QBrush & linkVisited()
func (this *QPalette) LinkVisited() *QBrush {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette11linkVisitedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCopyOf(const QPalette &)
func (this *QPalette) IsCopyOf(p *QPalette) bool {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette8isCopyOfERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpalette.h:152
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 cacheKey()
func (this *QPalette) CacheKey() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette8cacheKeyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qpalette.h:154
// index:0
// Public Visibility=Default Availability=Available
// [16] QPalette resolve(const QPalette &)
func (this *QPalette) Resolve(arg0 *QPalette) *QPalette /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette7resolveERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPalette)
	return rv2
}

// /usr/include/qt/QtGui/qpalette.h:155
// index:1
// Public inline Visibility=Default Availability=Available
// [4] uint resolve()
func (this *QPalette) Resolve_1() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette7resolveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qpalette.h:156
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void resolve(uint)
func (this *QPalette) Resolve_2(mask uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette7resolveEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mask)
	gopp.ErrPrint(err, rv)
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
