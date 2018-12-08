package qtwidgets

// /usr/include/qt/QtWidgets/qlabel.h
// #include <qlabel.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
func (this *QLabel) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QLabel) InheritKeyPressEvent(f func(ev *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QLabel) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void changeEvent(QEvent *)
func (this *QLabel) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QLabel) InheritMousePressEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QLabel) InheritMouseMoveEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QLabel) InheritMouseReleaseEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void contextMenuEvent(QContextMenuEvent *)
func (this *QLabel) InheritContextMenuEvent(f func(ev *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QLabel) InheritFocusInEvent(f func(ev *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QLabel) InheritFocusOutEvent(f func(ev *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// bool focusNextPrevChild(bool)
func (this *QLabel) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

/*

 */
type QLabel struct {
	*QFrame
}
type QLabel_ITF interface {
	QFrame_ITF
	QLabel_PTR() *QLabel
}

func (ptr *QLabel) QLabel_PTR() *QLabel { return ptr }

func (this *QLabel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QFrame.GetCthis()
	}
}
func (this *QLabel) SetCthis(cthis unsafe.Pointer) {
	this.QFrame = NewQFrameFromPointer(cthis)
}
func NewQLabelFromPointer(cthis unsafe.Pointer) *QLabel {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QLabel{bcthis0}
}
func (*QLabel) NewFromPointer(cthis unsafe.Pointer) *QLabel {
	return NewQLabelFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlabel.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QLabel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlabel.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLabel(QWidget *, Qt::WindowFlags)

/*
Constructs an empty label.

The parent and widget flag f, arguments are passed to the QFrame constructor.

See also setAlignment(), setFrameStyle(), and setIndent().
*/
func (*QLabel) NewForInherit(parent QWidget_ITF /*777 QWidget **/, f int) *QLabel {
	return NewQLabel(parent, f)
}
func NewQLabel(parent QWidget_ITF /*777 QWidget **/, f int) *QLabel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabelC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLabelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLabel(QWidget *, Qt::WindowFlags)

/*
Constructs an empty label.

The parent and widget flag f, arguments are passed to the QFrame constructor.

See also setAlignment(), setFrameStyle(), and setIndent().
*/
func (*QLabel) NewForInheritp() *QLabel {
	return NewQLabelp()
}
func NewQLabelp() *QLabel {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabelC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLabelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLabel(QWidget *, Qt::WindowFlags)

/*
Constructs an empty label.

The parent and widget flag f, arguments are passed to the QFrame constructor.

See also setAlignment(), setFrameStyle(), and setIndent().
*/
func (*QLabel) NewForInheritp1(parent QWidget_ITF /*777 QWidget **/) *QLabel {
	return NewQLabelp1(parent)
}
func NewQLabelp1(parent QWidget_ITF /*777 QWidget **/) *QLabel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabelC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLabelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLabel(const QString &, QWidget *, Qt::WindowFlags)

/*
Constructs an empty label.

The parent and widget flag f, arguments are passed to the QFrame constructor.

See also setAlignment(), setFrameStyle(), and setIndent().
*/
func (*QLabel) NewForInherit1(text string, parent QWidget_ITF /*777 QWidget **/, f int) *QLabel {
	return NewQLabel1(text, parent, f)
}
func NewQLabel1(text string, parent QWidget_ITF /*777 QWidget **/, f int) *QLabel {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabelC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLabelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLabel(const QString &, QWidget *, Qt::WindowFlags)

/*
Constructs an empty label.

The parent and widget flag f, arguments are passed to the QFrame constructor.

See also setAlignment(), setFrameStyle(), and setIndent().
*/
func (*QLabel) NewForInherit1p(text string) *QLabel {
	return NewQLabel1p(text)
}
func NewQLabel1p(text string) *QLabel {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabelC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLabelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLabel(const QString &, QWidget *, Qt::WindowFlags)

/*
Constructs an empty label.

The parent and widget flag f, arguments are passed to the QFrame constructor.

See also setAlignment(), setFrameStyle(), and setIndent().
*/
func (*QLabel) NewForInherit1p1(text string, parent QWidget_ITF /*777 QWidget **/) *QLabel {
	return NewQLabel1p1(text, parent)
}
func NewQLabel1p1(text string, parent QWidget_ITF /*777 QWidget **/) *QLabel {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabelC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLabelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLabel()

/*

 */
func DeleteQLabel(this *QLabel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlabel.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*

 */
func (this *QLabel) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlabel.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] const QPixmap * pixmap() const

/*

 */
func (this *QLabel) Pixmap() *qtgui.QPixmap /*777 const QPixmap **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel6pixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlabel.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] const QPicture * picture() const

/*
Returns the label's picture or 0 if the label doesn't have a picture.

See also setPicture().
*/
func (this *QLabel) Picture() *qtgui.QPicture /*777 const QPicture **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel7pictureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQPictureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlabel.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QMovie * movie() const

/*
Returns a pointer to the label's movie, or 0 if no movie has been set.

See also setMovie().
*/
func (this *QLabel) Movie() *qtgui.QMovie /*777 QMovie **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel5movieEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQMovieFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlabel.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextFormat textFormat() const

/*

 */
func (this *QLabel) TextFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel10textFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextFormat(Qt::TextFormat)

/*

 */
func (this *QLabel) SetTextFormat(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel13setTextFormatEN2Qt10TextFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const

/*

 */
func (this *QLabel) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)

/*

 */
func (this *QLabel) SetAlignment(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordWrap(bool)

/*

 */
func (this *QLabel) SetWordWrap(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel11setWordWrapEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wordWrap() const

/*

 */
func (this *QLabel) WordWrap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel8wordWrapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int indent() const

/*

 */
func (this *QLabel) Indent() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel6indentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlabel.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndent(int)

/*

 */
func (this *QLabel) SetIndent(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel9setIndentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] int margin() const

/*

 */
func (this *QLabel) Margin() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel6marginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlabel.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMargin(int)

/*

 */
func (this *QLabel) SetMargin(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel9setMarginEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasScaledContents() const

/*

 */
func (this *QLabel) HasScaledContents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel17hasScaledContentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScaledContents(bool)

/*

 */
func (this *QLabel) SetScaledContents(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel17setScaledContentsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:100
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QLabel) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QLabel) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBuddy(QWidget *)

/*
Sets this label's buddy to buddy.

When the user presses the shortcut key indicated by this label, the keyboard focus is transferred to the label's buddy widget.

The buddy mechanism is only available for QLabels that contain text in which one character is prefixed with an ampersand, '&'. This character is set as the shortcut key. See the QKeySequence::mnemonic() documentation for details (to display an actual ampersand, use '&&').

In a dialog, you might create two data entry widgets and a label for each, and set up the geometry layout so each label is just to the left of its data entry widget (its "buddy"), for example:


  QLineEdit *nameEdit  = new QLineEdit(this);
  QLabel    *nameLabel = new QLabel("&Name:", this);
  nameLabel->setBuddy(nameEdit);
  QLineEdit *phoneEdit  = new QLineEdit(this);
  QLabel    *phoneLabel = new QLabel("&Phone:", this);
  phoneLabel->setBuddy(phoneEdit);
  // (layout setup not shown)



With the code above, the focus jumps to the Name field when the user presses Alt+N, and to the Phone field when the user presses Alt+P.

To unset a previously set buddy, call this function with buddy set to 0.

See also buddy(), setText(), QShortcut, and setAlignment().
*/
func (this *QLabel) SetBuddy(arg0 QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel8setBuddyEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * buddy() const

/*
Returns this label's buddy, or 0 if no buddy is currently set.

See also setBuddy().
*/
func (this *QLabel) Buddy() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel5buddyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlabel.h:106
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int) const

/*
Reimplemented from QWidget::heightForWidth().
*/
func (this *QLabel) HeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlabel.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool openExternalLinks() const

/*

 */
func (this *QLabel) OpenExternalLinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel17openExternalLinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpenExternalLinks(bool)

/*

 */
func (this *QLabel) SetOpenExternalLinks(open bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel20setOpenExternalLinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), open)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextInteractionFlags(Qt::TextInteractionFlags)

/*

 */
func (this *QLabel) SetTextInteractionFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextInteractionFlags textInteractionFlags() const

/*

 */
func (this *QLabel) TextInteractionFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel20textInteractionFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelection(int, int)

/*
Selects text from position start and for length characters.

Note: The textInteractionFlags set on the label need to include either TextSelectableByMouse or TextSelectableByKeyboard.

This function was introduced in  Qt 4.7.

See also selectedText().
*/
func (this *QLabel) SetSelection(arg0 int, arg1 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel12setSelectionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasSelectedText() const

/*

 */
func (this *QLabel) HasSelectedText() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel15hasSelectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedText() const

/*

 */
func (this *QLabel) SelectedText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel12selectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlabel.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionStart() const

/*
selectionStart() returns the index of the first selected character in the label or -1 if no text is selected.

Note: The textInteractionFlags set on the label need to include either TextSelectableByMouse or TextSelectableByKeyboard.

This function was introduced in  Qt 4.7.

See also selectedText().
*/
func (this *QLabel) SelectionStart() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel14selectionStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlabel.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*

 */
func (this *QLabel) SetText(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixmap(const QPixmap &)

/*

 */
func (this *QLabel) SetPixmap(arg0 qtgui.QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPixmap_PTR() != nil {
		convArg0 = arg0.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel9setPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPicture(const QPicture &)

/*
Sets the label contents to picture. Any previous content is cleared.

The buddy shortcut, if any, is disabled.

See also picture() and setBuddy().
*/
func (this *QLabel) SetPicture(arg0 qtgui.QPicture_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPicture_PTR() != nil {
		convArg0 = arg0.QPicture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel10setPictureERK8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovie(QMovie *)

/*
Sets the label contents to movie. Any previous content is cleared. The label does NOT take ownership of the movie.

The buddy shortcut, if any, is disabled.

See also movie() and setBuddy().
*/
func (this *QLabel) SetMovie(movie qtgui.QMovie_ITF /*777 QMovie **/) {
	var convArg0 unsafe.Pointer
	if movie != nil && movie.QMovie_PTR() != nil {
		convArg0 = movie.QMovie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel8setMovieEP6QMovie", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNum(int)

/*
Sets the label contents to plain text containing the textual representation of integer num. Any previous content is cleared. Does nothing if the integer's string representation is the same as the current contents of the label.

The buddy shortcut, if any, is disabled.

See also setText(), QString::setNum(), and setBuddy().
*/
func (this *QLabel) SetNum(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel6setNumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:129
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setNum(double)

/*
Sets the label contents to plain text containing the textual representation of integer num. Any previous content is cleared. Does nothing if the integer's string representation is the same as the current contents of the label.

The buddy shortcut, if any, is disabled.

See also setText(), QString::setNum(), and setBuddy().
*/
func (this *QLabel) SetNum1(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel6setNumEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears any label contents.
*/
func (this *QLabel) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void linkActivated(const QString &)

/*
This signal is emitted when the user clicks a link. The URL referred to by the anchor is passed in link.

This function was introduced in  Qt 4.2.

See also linkHovered().
*/
func (this *QLabel) LinkActivated(link string) {
	var tmpArg0 = qtcore.NewQString5(link)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel13linkActivatedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void linkHovered(const QString &)

/*
This signal is emitted when the user hovers over a link. The URL referred to by the anchor is passed in link.

This function was introduced in  Qt 4.2.

See also linkActivated().
*/
func (this *QLabel) LinkHovered(link string) {
	var tmpArg0 = qtcore.NewQString5(link)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel11linkHoveredERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:137
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QLabel) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:138
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QLabel) KeyPressEvent(ev qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QKeyEvent_PTR() != nil {
		convArg0 = ev.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:139
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QLabel) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:140
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QLabel) ChangeEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:141
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QLabel) MousePressEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QLabel) MouseMoveEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:143
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QLabel) MouseReleaseEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:145
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*
Reimplemented from QWidget::contextMenuEvent().
*/
func (this *QLabel) ContextMenuEvent(ev qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QContextMenuEvent_PTR() != nil {
		convArg0 = ev.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:147
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusInEvent().
*/
func (this *QLabel) FocusInEvent(ev qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QFocusEvent_PTR() != nil {
		convArg0 = ev.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:148
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusOutEvent().
*/
func (this *QLabel) FocusOutEvent(ev qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QFocusEvent_PTR() != nil {
		convArg0 = ev.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:149
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(bool)

/*
Reimplemented from QWidget::focusNextPrevChild().
*/
func (this *QLabel) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
