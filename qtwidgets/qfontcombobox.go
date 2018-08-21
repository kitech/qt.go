package qtwidgets

// /usr/include/qt/QtWidgets/qfontcombobox.h
// #include <qfontcombobox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QFontComboBox) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QFontComboBox struct {
	*QComboBox
}
type QFontComboBox_ITF interface {
	QComboBox_ITF
	QFontComboBox_PTR() *QFontComboBox
}

func (ptr *QFontComboBox) QFontComboBox_PTR() *QFontComboBox { return ptr }

func (this *QFontComboBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QComboBox.GetCthis()
	}
}
func (this *QFontComboBox) SetCthis(cthis unsafe.Pointer) {
	this.QComboBox = NewQComboBoxFromPointer(cthis)
}
func NewQFontComboBoxFromPointer(cthis unsafe.Pointer) *QFontComboBox {
	bcthis0 := NewQComboBoxFromPointer(cthis)
	return &QFontComboBox{bcthis0}
}
func (*QFontComboBox) NewFromPointer(cthis unsafe.Pointer) *QFontComboBox {
	return NewQFontComboBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QFontComboBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontComboBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontComboBox(QWidget *)

/*
Constructs a font combobox with the given parent.
*/
func NewQFontComboBox(parent QWidget_ITF /*777 QWidget **/) *QFontComboBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontComboBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFontComboBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontComboBox(QWidget *)

/*
Constructs a font combobox with the given parent.
*/
func NewQFontComboBox__() *QFontComboBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontComboBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFontComboBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFontComboBox()

/*

 */
func DeleteQFontComboBox(this *QFontComboBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWritingSystem(QFontDatabase::WritingSystem)

/*

 */
func (this *QFontComboBox) SetWritingSystem(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBox16setWritingSystemEN13QFontDatabase13WritingSystemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] QFontDatabase::WritingSystem writingSystem() const

/*

 */
func (this *QFontComboBox) WritingSystem() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontComboBox13writingSystemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFontFilters(QFontComboBox::FontFilters)

/*

 */
func (this *QFontComboBox) SetFontFilters(filters int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBox14setFontFiltersE6QFlagsINS_10FontFilterEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] QFontComboBox::FontFilters fontFilters() const

/*

 */
func (this *QFontComboBox) FontFilters() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontComboBox11fontFiltersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:80
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont currentFont() const

/*

 */
func (this *QFontComboBox) CurrentFont() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontComboBox11currentFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QFontComboBox) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontComboBox8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentFont(const QFont &)

/*

 */
func (this *QFontComboBox) SetCurrentFont(f qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if f != nil && f.QFont_PTR() != nil {
		convArg0 = f.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBox14setCurrentFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentFontChanged(const QFont &)

/*
This signal is emitted whenever the current font changes, with the new font.

Note: Notifier signal for property currentFont.

See also currentFont.
*/
func (this *QFontComboBox) CurrentFontChanged(f qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if f != nil && f.QFont_PTR() != nil {
		convArg0 = f.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBox18currentFontChangedERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QFontComboBox) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontComboBox5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QFontComboBox__FontFilter = int

//
const QFontComboBox__AllFonts QFontComboBox__FontFilter = 0

//
const QFontComboBox__ScalableFonts QFontComboBox__FontFilter = 1

//
const QFontComboBox__NonScalableFonts QFontComboBox__FontFilter = 2

//
const QFontComboBox__MonospacedFonts QFontComboBox__FontFilter = 4

//
const QFontComboBox__ProportionalFonts QFontComboBox__FontFilter = 8

func (this *QFontComboBox) FontFilterItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFontComboBox_FontFilterItemName(val int) string {
	var nilthis *QFontComboBox
	return nilthis.FontFilterItemName(val)
}

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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
