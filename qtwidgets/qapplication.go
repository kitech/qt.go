package qtwidgets

// /usr/include/qt/QtWidgets/qapplication.h
// #include <qapplication.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
func (this *QApplication) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QApplication struct {
	*qtgui.QGuiApplication
}
type QApplication_ITF interface {
	qtgui.QGuiApplication_ITF
	QApplication_PTR() *QApplication
}

func (ptr *QApplication) QApplication_PTR() *QApplication { return ptr }

func (this *QApplication) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGuiApplication.GetCthis()
	}
}
func (this *QApplication) SetCthis(cthis unsafe.Pointer) {
	this.QGuiApplication = qtgui.NewQGuiApplicationFromPointer(cthis)
}
func NewQApplicationFromPointer(cthis unsafe.Pointer) *QApplication {
	bcthis0 := qtgui.NewQGuiApplicationFromPointer(cthis)
	return &QApplication{bcthis0}
}
func (*QApplication) NewFromPointer(cthis unsafe.Pointer) *QApplication {
	return NewQApplicationFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qapplication.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QApplication) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QApplication10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qapplication.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QApplication(int &, char **, int)

/*
Initializes the window system and constructs an application object with argc command line arguments in argv.

Warning: The data referred to by argc and argv must stay valid for the entire lifetime of the QApplication object. In addition, argc must be greater than zero and argv must contain at least one valid character string.

The global qApp pointer refers to this application object. Only one application object should be created.

This application object must be constructed before any paint devices (including widgets, pixmaps, bitmaps etc.).

Note: argc and argv might be changed as Qt removes command line arguments that it recognizes.

All Qt programs automatically support the following command line options:


-style= style, sets the application GUI style. Possible values depend on your system configuration. If you compiled Qt with additional styles or have additional styles as plugins these will be available to the -style command line option. You can also set the style for all Qt applications by setting the QT_STYLE_OVERRIDE environment variable.
-style style, is the same as listed above.
-stylesheet= stylesheet, sets the application styleSheet. The value must be a path to a file that contains the Style Sheet.Note: Relative URLs in the Style Sheet file are relative to the Style Sheet file's path.
-stylesheet stylesheet, is the same as listed above.
-widgetcount, prints debug message at the end about number of widgets left undestroyed and maximum number of widgets existed at the same time
-reverse, sets the application's layout direction to Qt::RightToLeft
-qmljsdebugger=, activates the QML/JS debugger with a specified port. The value must be of format port:1234[,block], where block is optional and will make the application wait until a debugger connects to it.


See also QCoreApplication::arguments().
*/
func (*QApplication) NewForInherit(argc int, argv []string, arg2 int) *QApplication {
	return NewQApplication(argc, argv, arg2)
}
func NewQApplication(argc int, argv []string, arg2 int) *QApplication {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplicationC2ERiPPci", qtrt.FFI_TYPE_POINTER, &argc, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQApplicationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QApplication")
	return gothis
}

// /usr/include/qt/QtWidgets/qapplication.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QApplication(int &, char **, int)

/*
Initializes the window system and constructs an application object with argc command line arguments in argv.

Warning: The data referred to by argc and argv must stay valid for the entire lifetime of the QApplication object. In addition, argc must be greater than zero and argv must contain at least one valid character string.

The global qApp pointer refers to this application object. Only one application object should be created.

This application object must be constructed before any paint devices (including widgets, pixmaps, bitmaps etc.).

Note: argc and argv might be changed as Qt removes command line arguments that it recognizes.

All Qt programs automatically support the following command line options:


-style= style, sets the application GUI style. Possible values depend on your system configuration. If you compiled Qt with additional styles or have additional styles as plugins these will be available to the -style command line option. You can also set the style for all Qt applications by setting the QT_STYLE_OVERRIDE environment variable.
-style style, is the same as listed above.
-stylesheet= stylesheet, sets the application styleSheet. The value must be a path to a file that contains the Style Sheet.Note: Relative URLs in the Style Sheet file are relative to the Style Sheet file's path.
-stylesheet stylesheet, is the same as listed above.
-widgetcount, prints debug message at the end about number of widgets left undestroyed and maximum number of widgets existed at the same time
-reverse, sets the application's layout direction to Qt::RightToLeft
-qmljsdebugger=, activates the QML/JS debugger with a specified port. The value must be of format port:1234[,block], where block is optional and will make the application wait until a debugger connects to it.


See also QCoreApplication::arguments().
*/
func (*QApplication) NewForInheritp(argc int, argv []string) *QApplication {
	return NewQApplicationp(argc, argv)
}
func NewQApplicationp(argc int, argv []string) *QApplication {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	// arg: 2, int=Int, =Invalid, , Invalid
	arg2 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplicationC2ERiPPci", qtrt.FFI_TYPE_POINTER, &argc, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQApplicationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QApplication")
	return gothis
}

// /usr/include/qt/QtWidgets/qapplication.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QApplication()

/*

 */
func DeleteQApplication(this *QApplication) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplicationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qapplication.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStyle * style()

/*
Returns the application's style object.

See also setStyle() and QStyle.
*/
func (this *QApplication) Style() *QStyle /*777 QStyle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication5styleEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QApplication_Style() *QStyle /*777 QStyle **/ {
	var nilthis *QApplication
	rv := nilthis.Style()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:99
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setStyle(QStyle *)

/*
Sets the application's GUI style to style. Ownership of the style object is transferred to QApplication, so QApplication will delete the style object on application exit or when a new style is set and the old style is still the parent of the application object.

Example usage:


  QApplication::setStyle(QStyleFactory::create("Fusion"));



When switching application styles, the color palette is set back to the initial colors or the system defaults. This is necessary since certain styles have to adapt the color palette to be fully style-guide compliant.

Setting the style before a palette has been set, i.e., before creating QApplication, will cause the application to use QStyle::standardPalette() for the palette.

Warning: Qt style sheets are currently not supported for custom QStyle subclasses. We plan to address this in some future release.

See also style(), QStyle, setPalette(), and desktopSettingsAware().
*/
func (this *QApplication) SetStyle(arg0 QStyle_ITF /*777 QStyle **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStyle_PTR() != nil {
		convArg0 = arg0.QStyle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication8setStyleEP6QStyle", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetStyle(arg0 QStyle_ITF /*777 QStyle **/) {
	var nilthis *QApplication
	nilthis.SetStyle(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:100
// index:1
// Public static Visibility=Default Availability=Available
// [8] QStyle * setStyle(const QString &)

/*
Sets the application's GUI style to style. Ownership of the style object is transferred to QApplication, so QApplication will delete the style object on application exit or when a new style is set and the old style is still the parent of the application object.

Example usage:


  QApplication::setStyle(QStyleFactory::create("Fusion"));



When switching application styles, the color palette is set back to the initial colors or the system defaults. This is necessary since certain styles have to adapt the color palette to be fully style-guide compliant.

Setting the style before a palette has been set, i.e., before creating QApplication, will cause the application to use QStyle::standardPalette() for the palette.

Warning: Qt style sheets are currently not supported for custom QStyle subclasses. We plan to address this in some future release.

See also style(), QStyle, setPalette(), and desktopSettingsAware().
*/
func (this *QApplication) SetStyle1(arg0 string) *QStyle /*777 QStyle **/ {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication8setStyleERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QApplication_SetStyle1(arg0 string) *QStyle /*777 QStyle **/ {
	var nilthis *QApplication
	rv := nilthis.SetStyle1(arg0)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:103
// index:0
// Public static Visibility=Default Availability=Available
// [4] int colorSpec()

/*

 */
func (this *QApplication) ColorSpec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication9colorSpecEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_ColorSpec() int {
	var nilthis *QApplication
	rv := nilthis.ColorSpec()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:104
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setColorSpec(int)

/*

 */
func (this *QApplication) SetColorSpec(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication12setColorSpecEi", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetColorSpec(arg0 int) {
	var nilthis *QApplication
	nilthis.SetColorSpec(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:111
// index:0
// Public static Visibility=Default Availability=Available
// [16] QPalette palette(const QWidget *)

/*
If a widget is passed, the default palette for the widget's class is returned. This may or may not be the application palette. In most cases there is no special palette for certain types of widgets, but one notable exception is the popup menu under Windows, if the user has defined a special background color for menus in the display settings.

See also setPalette() and QWidget::palette().
*/
func (this *QApplication) Palette(arg0 QWidget_ITF /*777 const QWidget **/) *qtgui.QPalette /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication7paletteEPK7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPalette)
	return rv2
}
func QApplication_Palette(arg0 QWidget_ITF /*777 const QWidget **/) *qtgui.QPalette /*123*/ {
	var nilthis *QApplication
	rv := nilthis.Palette(arg0)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:112
// index:1
// Public static Visibility=Default Availability=Available
// [16] QPalette palette(const char *)

/*
If a widget is passed, the default palette for the widget's class is returned. This may or may not be the application palette. In most cases there is no special palette for certain types of widgets, but one notable exception is the popup menu under Windows, if the user has defined a special background color for menus in the display settings.

See also setPalette() and QWidget::palette().
*/
func (this *QApplication) Palette1(className string) *qtgui.QPalette /*123*/ {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication7paletteEPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPalette)
	return rv2
}
func QApplication_Palette1(className string) *qtgui.QPalette /*123*/ {
	var nilthis *QApplication
	rv := nilthis.Palette1(className)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:113
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setPalette(const QPalette &, const char *)

/*
Changes the default application palette to palette.

If className is passed, the change applies only to widgets that inherit className (as reported by QObject::inherits()). If className is left 0, the change affects all widgets, thus overriding any previously set class specific palettes.

The palette may be changed according to the current GUI style in QStyle::polish().

Warning: Do not use this function in conjunction with Qt Style Sheets. When using style sheets, the palette of a widget can be customized using the "color", "background-color", "selection-color", "selection-background-color" and "alternate-background-color".

Note: Some styles do not use the palette for all drawing, for instance, if they make use of native theme engines. This is the case for the Windows Vista and macOS styles.

See also QWidget::setPalette(), palette(), and QStyle::polish().
*/
func (this *QApplication) SetPalette(arg0 qtgui.QPalette_ITF, className string) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPalette_PTR() != nil {
		convArg0 = arg0.QPalette_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication10setPaletteERK8QPalettePKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetPalette(arg0 qtgui.QPalette_ITF, className string) {
	var nilthis *QApplication
	nilthis.SetPalette(arg0, className)
}

// /usr/include/qt/QtWidgets/qapplication.h:113
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setPalette(const QPalette &, const char *)

/*
Changes the default application palette to palette.

If className is passed, the change applies only to widgets that inherit className (as reported by QObject::inherits()). If className is left 0, the change affects all widgets, thus overriding any previously set class specific palettes.

The palette may be changed according to the current GUI style in QStyle::polish().

Warning: Do not use this function in conjunction with Qt Style Sheets. When using style sheets, the palette of a widget can be customized using the "color", "background-color", "selection-color", "selection-background-color" and "alternate-background-color".

Note: Some styles do not use the palette for all drawing, for instance, if they make use of native theme engines. This is the case for the Windows Vista and macOS styles.

See also QWidget::setPalette(), palette(), and QStyle::polish().
*/
func (this *QApplication) SetPalettep(arg0 qtgui.QPalette_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPalette_PTR() != nil {
		convArg0 = arg0.QPalette_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication10setPaletteERK8QPalettePKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [16] QFont font()

/*
Returns the default application font.

See also setFont(), fontMetrics(), and QWidget::font().
*/
func (this *QApplication) Font() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication4fontEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}
func QApplication_Font() *qtgui.QFont /*123*/ {
	var nilthis *QApplication
	rv := nilthis.Font()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:115
// index:1
// Public static Visibility=Default Availability=Available
// [16] QFont font(const QWidget *)

/*
Returns the default application font.

See also setFont(), fontMetrics(), and QWidget::font().
*/
func (this *QApplication) Font1(arg0 QWidget_ITF /*777 const QWidget **/) *qtgui.QFont /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication4fontEPK7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}
func QApplication_Font1(arg0 QWidget_ITF /*777 const QWidget **/) *qtgui.QFont /*123*/ {
	var nilthis *QApplication
	rv := nilthis.Font1(arg0)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:116
// index:2
// Public static Visibility=Default Availability=Available
// [16] QFont font(const char *)

/*
Returns the default application font.

See also setFont(), fontMetrics(), and QWidget::font().
*/
func (this *QApplication) Font2(className string) *qtgui.QFont /*123*/ {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication4fontEPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}
func QApplication_Font2(className string) *qtgui.QFont /*123*/ {
	var nilthis *QApplication
	rv := nilthis.Font2(className)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:117
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setFont(const QFont &, const char *)

/*
Changes the default application font to font. If className is passed, the change applies only to classes that inherit className (as reported by QObject::inherits()).

On application start-up, the default font depends on the window system. It can vary depending on both the window system version and the locale. This function lets you override the default font; but overriding may be a bad idea because, for example, some locales need extra large fonts to support their special characters.

Warning: Do not use this function in conjunction with Qt Style Sheets. The font of an application can be customized using the "font" style sheet property. To set a bold font for all QPushButtons, set the application styleSheet() as "QPushButton { font: bold }"

See also font(), fontMetrics(), and QWidget::setFont().
*/
func (this *QApplication) SetFont(arg0 qtgui.QFont_ITF, className string) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication7setFontERK5QFontPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetFont(arg0 qtgui.QFont_ITF, className string) {
	var nilthis *QApplication
	nilthis.SetFont(arg0, className)
}

// /usr/include/qt/QtWidgets/qapplication.h:117
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setFont(const QFont &, const char *)

/*
Changes the default application font to font. If className is passed, the change applies only to classes that inherit className (as reported by QObject::inherits()).

On application start-up, the default font depends on the window system. It can vary depending on both the window system version and the locale. This function lets you override the default font; but overriding may be a bad idea because, for example, some locales need extra large fonts to support their special characters.

Warning: Do not use this function in conjunction with Qt Style Sheets. The font of an application can be customized using the "font" style sheet property. To set a bold font for all QPushButtons, set the application styleSheet() as "QPushButton { font: bold }"

See also font(), fontMetrics(), and QWidget::setFont().
*/
func (this *QApplication) SetFontp(arg0 qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication7setFontERK5QFontPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QFontMetrics fontMetrics()

/*
Returns display (screen) font metrics for the application font.

See also font(), setFont(), QWidget::fontMetrics(), and QPainter::fontMetrics().
*/
func (this *QApplication) FontMetrics() *qtgui.QFontMetrics /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication11fontMetricsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFontMetrics)
	return rv2
}
func QApplication_FontMetrics() *qtgui.QFontMetrics /*123*/ {
	var nilthis *QApplication
	rv := nilthis.FontMetrics()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setWindowIcon(const QIcon &)

/*

 */
func (this *QApplication) SetWindowIcon(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication13setWindowIconERK5QIcon", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetWindowIcon(icon qtgui.QIcon_ITF) {
	var nilthis *QApplication
	nilthis.SetWindowIcon(icon)
}

// /usr/include/qt/QtWidgets/qapplication.h:122
// index:0
// Public static Visibility=Default Availability=Available
// [8] QIcon windowIcon()

/*

 */
func (this *QApplication) WindowIcon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication10windowIconEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}
func QApplication_WindowIcon() *qtgui.QIcon /*123*/ {
	var nilthis *QApplication
	rv := nilthis.WindowIcon()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QWidgetList allWidgets()

/*
Returns a list of all the widgets in the application.

The list is empty (QList::isEmpty()) if there are no widgets.

Note: Some of the widgets may be hidden.

Example:


  void updateAllWidgets()
  {
      foreach (QWidget *widget, QApplication::allWidgets())
          widget->update();
  }



See also topLevelWidgets() and QWidget::isVisible().
*/
func (this *QApplication) AllWidgets() *QWidgetList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication10allWidgetsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*222*/ NewQWidgetListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}
func QApplication_AllWidgets() *QWidgetList /*667*/ {
	var nilthis *QApplication
	rv := nilthis.AllWidgets()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:126
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QWidgetList topLevelWidgets()

/*
Returns a list of the top-level widgets (windows) in the application.

Note: Some of the top-level widgets may be hidden, for example a tooltip if no tooltip is currently shown.

Example:


  void showAllHiddenTopLevelWidgets()
  {
      foreach (QWidget *widget, QApplication::topLevelWidgets()) {
          if (widget->isHidden())
              widget->show();
      }
  }



See also allWidgets(), QWidget::isWindow(), and QWidget::isHidden().
*/
func (this *QApplication) TopLevelWidgets() *QWidgetList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication15topLevelWidgetsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*222*/ NewQWidgetListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}
func QApplication_TopLevelWidgets() *QWidgetList /*667*/ {
	var nilthis *QApplication
	rv := nilthis.TopLevelWidgets()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:128
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDesktopWidget * desktop()

/*
Returns the desktop widget (also called the root window).

The desktop may be composed of multiple screens, so it would be incorrect, for example, to attempt to center some widget in the desktop's geometry. QDesktopWidget has various functions for obtaining useful geometries upon the desktop, such as QDesktopWidget::screenGeometry() and QDesktopWidget::availableGeometry().

On X11, it is also possible to draw on the desktop.
*/
func (this *QApplication) Desktop() *QDesktopWidget /*777 QDesktopWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication7desktopEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQDesktopWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QApplication_Desktop() *QDesktopWidget /*777 QDesktopWidget **/ {
	var nilthis *QApplication
	rv := nilthis.Desktop()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:130
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * activePopupWidget()

/*
Returns the active popup widget.

A popup widget is a special top-level widget that sets the Qt::WType_Popup widget flag, e.g. the QMenu widget. When the application opens a popup widget, all events are sent to the popup. Normal widgets and modal widgets cannot be accessed before the popup widget is closed.

Only other popup widgets may be opened when a popup widget is shown. The popup widgets are organized in a stack. This function returns the active popup widget at the top of the stack.

See also activeModalWidget() and topLevelWidgets().
*/
func (this *QApplication) ActivePopupWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication17activePopupWidgetEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QApplication_ActivePopupWidget() *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.ActivePopupWidget()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:131
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * activeModalWidget()

/*
Returns the active modal widget.

A modal widget is a special top-level widget which is a subclass of QDialog that specifies the modal parameter of the constructor as true. A modal widget must be closed before the user can continue with other parts of the program.

Modal widgets are organized in a stack. This function returns the active modal widget at the top of the stack.

See also activePopupWidget() and topLevelWidgets().
*/
func (this *QApplication) ActiveModalWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication17activeModalWidgetEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QApplication_ActiveModalWidget() *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.ActiveModalWidget()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:132
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * focusWidget()

/*
Returns the application widget that has the keyboard input focus, or 0 if no widget in this application has the focus.

See also QWidget::setFocus(), QWidget::hasFocus(), activeWindow(), and focusChanged().
*/
func (this *QApplication) FocusWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication11focusWidgetEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QApplication_FocusWidget() *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.FocusWidget()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:134
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * activeWindow()

/*
Returns the application top-level window that has the keyboard input focus, or 0 if no application window has the focus. There might be an activeWindow() even if there is no focusWidget(), for example if no widget in that window accepts key events.

See also setActiveWindow(), QWidget::setFocus(), QWidget::hasFocus(), and focusWidget().
*/
func (this *QApplication) ActiveWindow() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication12activeWindowEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QApplication_ActiveWindow() *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.ActiveWindow()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:135
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setActiveWindow(QWidget *)

/*
Sets the active window to the active widget in response to a system event. The function is called from the platform specific event handlers.

Warning: This function does not set the keyboard focus to the active widget. Call QWidget::activateWindow() instead.

It sets the activeWindow() and focusWidget() attributes and sends proper WindowActivate/WindowDeactivate and FocusIn/FocusOut events to all appropriate widgets. The window will then be painted in active state (e.g. cursors in line edits will blink), and it will have tool tips enabled.

See also activeWindow() and QWidget::activateWindow().
*/
func (this *QApplication) SetActiveWindow(act QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if act != nil && act.QWidget_PTR() != nil {
		convArg0 = act.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication15setActiveWindowEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetActiveWindow(act QWidget_ITF /*777 QWidget **/) {
	var nilthis *QApplication
	nilthis.SetActiveWindow(act)
}

// /usr/include/qt/QtWidgets/qapplication.h:137
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * widgetAt(const QPoint &)

/*
Returns the widget at global screen position point, or 0 if there is no Qt widget there.

This function can be slow.

See also QCursor::pos(), QWidget::grabMouse(), and QWidget::grabKeyboard().
*/
func (this *QApplication) WidgetAt(p qtcore.QPoint_ITF) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication8widgetAtERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QApplication_WidgetAt(p qtcore.QPoint_ITF) *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.WidgetAt(p)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:138
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QWidget * widgetAt(int, int)

/*
Returns the widget at global screen position point, or 0 if there is no Qt widget there.

This function can be slow.

See also QCursor::pos(), QWidget::grabMouse(), and QWidget::grabKeyboard().
*/
func (this *QApplication) WidgetAt1(x int, y int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication8widgetAtEii", qtrt.FFI_TYPE_POINTER, x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QApplication_WidgetAt1(x int, y int) *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.WidgetAt1(x, y)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:139
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * topLevelAt(const QPoint &)

/*
Returns the top-level widget at the given point; returns 0 if there is no such widget.
*/
func (this *QApplication) TopLevelAt(p qtcore.QPoint_ITF) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication10topLevelAtERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QApplication_TopLevelAt(p qtcore.QPoint_ITF) *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.TopLevelAt(p)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:140
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QWidget * topLevelAt(int, int)

/*
Returns the top-level widget at the given point; returns 0 if there is no such widget.
*/
func (this *QApplication) TopLevelAt1(x int, y int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication10topLevelAtEii", qtrt.FFI_TYPE_POINTER, x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QApplication_TopLevelAt1(x int, y int) *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.TopLevelAt1(x, y)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:145
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void beep()

/*
Sounds the bell, using the default volume and sound. The function is not available in Qt for Embedded Linux.
*/
func (this *QApplication) Beep() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication4beepEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QApplication_Beep() {
	var nilthis *QApplication
	nilthis.Beep()
}

// /usr/include/qt/QtWidgets/qapplication.h:146
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void alert(QWidget *, int)

/*
Causes an alert to be shown for widget if the window is not the active window. The alert is shown for msec miliseconds. If msec is zero (the default), then the alert is shown indefinitely until the window becomes active again.

Currently this function does nothing on Qt for Embedded Linux.

On macOS, this works more at the application level and will cause the application icon to bounce in the dock.

On Windows, this causes the window's taskbar entry to flash for a time. If msec is zero, the flashing will stop and the taskbar entry will turn a different color (currently orange).

On X11, this will cause the window to be marked as "demands attention", the window must not be hidden (i.e. not have hide() called on it, but be visible in some sort of way) in order for this to work.

This function was introduced in  Qt 4.3.
*/
func (this *QApplication) Alert(widget QWidget_ITF /*777 QWidget **/, duration int) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication5alertEP7QWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, duration)
	qtrt.ErrPrint(err, rv)
}
func QApplication_Alert(widget QWidget_ITF /*777 QWidget **/, duration int) {
	var nilthis *QApplication
	nilthis.Alert(widget, duration)
}

// /usr/include/qt/QtWidgets/qapplication.h:146
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void alert(QWidget *, int)

/*
Causes an alert to be shown for widget if the window is not the active window. The alert is shown for msec miliseconds. If msec is zero (the default), then the alert is shown indefinitely until the window becomes active again.

Currently this function does nothing on Qt for Embedded Linux.

On macOS, this works more at the application level and will cause the application icon to bounce in the dock.

On Windows, this causes the window's taskbar entry to flash for a time. If msec is zero, the flashing will stop and the taskbar entry will turn a different color (currently orange).

On X11, this will cause the window to be marked as "demands attention", the window must not be hidden (i.e. not have hide() called on it, but be visible in some sort of way) in order for this to work.

This function was introduced in  Qt 4.3.
*/
func (this *QApplication) Alertp(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	duration := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication5alertEP7QWidgeti", qtrt.FFI_TYPE_POINTER, convArg0, duration)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:148
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setCursorFlashTime(int)

/*

 */
func (this *QApplication) SetCursorFlashTime(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication18setCursorFlashTimeEi", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetCursorFlashTime(arg0 int) {
	var nilthis *QApplication
	nilthis.SetCursorFlashTime(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:149
// index:0
// Public static Visibility=Default Availability=Available
// [4] int cursorFlashTime()

/*

 */
func (this *QApplication) CursorFlashTime() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication15cursorFlashTimeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_CursorFlashTime() int {
	var nilthis *QApplication
	rv := nilthis.CursorFlashTime()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:151
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDoubleClickInterval(int)

/*

 */
func (this *QApplication) SetDoubleClickInterval(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication22setDoubleClickIntervalEi", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetDoubleClickInterval(arg0 int) {
	var nilthis *QApplication
	nilthis.SetDoubleClickInterval(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:152
// index:0
// Public static Visibility=Default Availability=Available
// [4] int doubleClickInterval()

/*

 */
func (this *QApplication) DoubleClickInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication19doubleClickIntervalEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_DoubleClickInterval() int {
	var nilthis *QApplication
	rv := nilthis.DoubleClickInterval()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:154
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setKeyboardInputInterval(int)

/*

 */
func (this *QApplication) SetKeyboardInputInterval(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication24setKeyboardInputIntervalEi", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetKeyboardInputInterval(arg0 int) {
	var nilthis *QApplication
	nilthis.SetKeyboardInputInterval(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:155
// index:0
// Public static Visibility=Default Availability=Available
// [4] int keyboardInputInterval()

/*

 */
func (this *QApplication) KeyboardInputInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication21keyboardInputIntervalEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_KeyboardInputInterval() int {
	var nilthis *QApplication
	rv := nilthis.KeyboardInputInterval()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:161
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setGlobalStrut(const QSize &)

/*

 */
func (this *QApplication) SetGlobalStrut(arg0 qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication14setGlobalStrutERK5QSize", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetGlobalStrut(arg0 qtcore.QSize_ITF) {
	var nilthis *QApplication
	nilthis.SetGlobalStrut(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:162
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSize globalStrut()

/*

 */
func (this *QApplication) GlobalStrut() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication11globalStrutEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}
func QApplication_GlobalStrut() *qtcore.QSize /*123*/ {
	var nilthis *QApplication
	rv := nilthis.GlobalStrut()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:164
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setStartDragTime(int)

/*

 */
func (this *QApplication) SetStartDragTime(ms int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication16setStartDragTimeEi", qtrt.FFI_TYPE_POINTER, ms)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetStartDragTime(ms int) {
	var nilthis *QApplication
	nilthis.SetStartDragTime(ms)
}

// /usr/include/qt/QtWidgets/qapplication.h:165
// index:0
// Public static Visibility=Default Availability=Available
// [4] int startDragTime()

/*

 */
func (this *QApplication) StartDragTime() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication13startDragTimeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_StartDragTime() int {
	var nilthis *QApplication
	rv := nilthis.StartDragTime()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:166
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setStartDragDistance(int)

/*

 */
func (this *QApplication) SetStartDragDistance(l int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication20setStartDragDistanceEi", qtrt.FFI_TYPE_POINTER, l)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetStartDragDistance(l int) {
	var nilthis *QApplication
	nilthis.SetStartDragDistance(l)
}

// /usr/include/qt/QtWidgets/qapplication.h:167
// index:0
// Public static Visibility=Default Availability=Available
// [4] int startDragDistance()

/*

 */
func (this *QApplication) StartDragDistance() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication17startDragDistanceEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_StartDragDistance() int {
	var nilthis *QApplication
	rv := nilthis.StartDragDistance()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:169
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isEffectEnabled(Qt::UIEffect)

/*
Returns true if effect is enabled; otherwise returns false.

By default, Qt will try to use the desktop settings. To prevent this, call setDesktopSettingsAware(false).

Note: All effects are disabled on screens running at less than 16-bit color depth.

See also setEffectEnabled() and Qt::UIEffect.
*/
func (this *QApplication) IsEffectEnabled(arg0 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication15isEffectEnabledEN2Qt8UIEffectE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QApplication_IsEffectEnabled(arg0 int) bool {
	var nilthis *QApplication
	rv := nilthis.IsEffectEnabled(arg0)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:170
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setEffectEnabled(Qt::UIEffect, bool)

/*
Enables the UI effect effect if enable is true, otherwise the effect will not be used.

Note: All effects are disabled on screens running at less than 16-bit color depth.

See also isEffectEnabled(), Qt::UIEffect, and setDesktopSettingsAware().
*/
func (this *QApplication) SetEffectEnabled(arg0 int, enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication16setEffectEnabledEN2Qt8UIEffectEb", qtrt.FFI_TYPE_POINTER, arg0, enable)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetEffectEnabled(arg0 int, enable bool) {
	var nilthis *QApplication
	nilthis.SetEffectEnabled(arg0, enable)
}

// /usr/include/qt/QtWidgets/qapplication.h:170
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setEffectEnabled(Qt::UIEffect, bool)

/*
Enables the UI effect effect if enable is true, otherwise the effect will not be used.

Note: All effects are disabled on screens running at less than 16-bit color depth.

See also isEffectEnabled(), Qt::UIEffect, and setDesktopSettingsAware().
*/
func (this *QApplication) SetEffectEnabledp(arg0 int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	enable := true
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication16setEffectEnabledEN2Qt8UIEffectEb", qtrt.FFI_TYPE_POINTER, arg0, enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [4] int exec()

/*
Enters the main event loop and waits until exit() is called, then returns the value that was set to exit() (which is 0 if exit() is called via quit()).

It is necessary to call this function to start event handling. The main event loop receives events from the window system and dispatches these to the application widgets.

Generally, no user interaction can take place before calling exec(). As a special case, modal widgets like QMessageBox can be used before calling exec(), because modal widgets call exec() to start a local event loop.

To make your application perform idle processing, i.e., executing a special function whenever there are no pending events, use a QTimer with 0 timeout. More advanced idle processing schemes can be achieved using processEvents().

We recommend that you connect clean-up code to the aboutToQuit() signal, instead of putting it in your application's main() function. This is because, on some platforms the QApplication::exec() call may not return. For example, on the Windows platform, when the user logs off, the system terminates the process after Qt closes all top-level windows. Hence, there is no guarantee that the application will have time to exit its event loop and execute code at the end of the main() function, after the QApplication::exec() call.

See also quitOnLastWindowClosed, QCoreApplication::quit(), QCoreApplication::exit(), QCoreApplication::processEvents(), and QCoreApplication::exec().
*/
func (this *QApplication) Exec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication4execEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_Exec() int {
	var nilthis *QApplication
	rv := nilthis.Exec()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:180
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool notify(QObject *, QEvent *)

/*
Reimplemented from QGuiApplication::notify().
*/
func (this *QApplication) Notify(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication6notifyEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qapplication.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusChanged(QWidget *, QWidget *)

/*
This signal is emitted when the widget that has keyboard focus changed from old to now, i.e., because the user pressed the tab-key, clicked into a widget or changed the active window. Both old and now can be the null-pointer.

The signal is emitted after both widget have been notified about the change through QFocusEvent.

This function was introduced in  Qt 4.1.

See also QWidget::setFocus(), QWidget::clearFocus(), and Qt::FocusReason.
*/
func (this *QApplication) FocusChanged(old QWidget_ITF /*777 QWidget **/, now QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if old != nil && old.QWidget_PTR() != nil {
		convArg0 = old.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if now != nil && now.QWidget_PTR() != nil {
		convArg1 = now.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication12focusChangedEP7QWidgetS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:193
// index:0
// Public Visibility=Default Availability=Available
// [8] QString styleSheet() const

/*

 */
func (this *QApplication) StyleSheet() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QApplication10styleSheetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qapplication.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyleSheet(const QString &)

/*

 */
func (this *QApplication) SetStyleSheet(sheet string) {
	var tmpArg0 = qtcore.NewQString5(sheet)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication13setStyleSheetERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoSipEnabled(const bool)

/*

 */
func (this *QApplication) SetAutoSipEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication17setAutoSipEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:199
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoSipEnabled() const

/*

 */
func (this *QApplication) AutoSipEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QApplication14autoSipEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qapplication.h:200
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void closeAllWindows()

/*
Closes all top-level windows.

This function is particularly useful for applications with many top-level windows. It could, for example, be connected to a Exit entry in the File menu:


      const QIcon exitIcon = QIcon::fromTheme("application-exit");
      QAction *exitAct = fileMenu->addAction(exitIcon, tr("E&xit"), qApp, &QApplication::closeAllWindows);
      exitAct->setShortcuts(QKeySequence::Quit);
      exitAct->setStatusTip(tr("Exit the application"));
      fileMenu->addAction(exitAct);



The windows are closed in random order, until one window does not accept the close event. The application quits when the last window was successfully closed; this can be turned off by setting quitOnLastWindowClosed to false.

See also quitOnLastWindowClosed, lastWindowClosed(), QWidget::close(), QWidget::closeEvent(), lastWindowClosed(), QCoreApplication::quit(), topLevelWidgets(), and QWidget::isWindow().
*/
func (this *QApplication) CloseAllWindows() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication15closeAllWindowsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QApplication_CloseAllWindows() {
	var nilthis *QApplication
	nilthis.CloseAllWindows()
}

// /usr/include/qt/QtWidgets/qapplication.h:201
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void aboutQt()

/*
Displays a simple message box about Qt. The message includes the version number of Qt being used by the application.

This is useful for inclusion in the Help menu of an application, as shown in the Menus example.

This function is a convenience slot for QMessageBox::aboutQt().
*/
func (this *QApplication) AboutQt() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication7aboutQtEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QApplication_AboutQt() {
	var nilthis *QApplication
	nilthis.AboutQt()
}

// /usr/include/qt/QtWidgets/qapplication.h:204
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QGuiApplication::event().
*/
func (this *QApplication) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QApplication__ColorSpec = int

//
const QApplication__NormalColor QApplication__ColorSpec = 0

//
const QApplication__CustomColor QApplication__ColorSpec = 1

//
const QApplication__ManyColor QApplication__ColorSpec = 2

func (this *QApplication) ColorSpecItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QApplication_ColorSpecItemName(val int) string {
	var nilthis *QApplication
	return nilthis.ColorSpecItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11069() {
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
