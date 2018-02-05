package qtwidgets

// /usr/include/qt/QtWidgets/qlabel.h
// #include <qlabel.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool event(class QEvent *)
func (this *QLabel) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QLabel) InheritKeyPressEvent(f func(ev *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QLabel) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QLabel) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QLabel) InheritMousePressEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QLabel) InheritMouseMoveEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QLabel) InheritMouseReleaseEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QLabel) InheritContextMenuEvent(f func(ev *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QLabel) InheritFocusInEvent(f func(ev *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QLabel) InheritFocusOutEvent(f func(ev *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// bool focusNextPrevChild(_Bool)
func (this *QLabel) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

type QLabel struct {
	*QFrame
}

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
// [8] const QMetaObject * metaObject()
func (this *QLabel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLabel(QWidget *, Qt::WindowFlags)
func NewQLabel(parent *QWidget /*777 QWidget **/, f int) *QLabel {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabelC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQLabelFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLabel(const QString &, QWidget *, Qt::WindowFlags)
func NewQLabel_1(text *qtcore.QString, parent *QWidget /*777 QWidget **/, f int) *QLabel {
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabelC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQLabelFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLabel()
func DeleteQLabel(this *QLabel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlabel.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QLabel) Text() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] const QPixmap * pixmap()
func (this *QLabel) Pixmap() *qtgui.QPixmap /*777 const QPixmap **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel6pixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] const QPicture * picture()
func (this *QLabel) Picture() *qtgui.QPicture /*777 const QPicture **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel7pictureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPictureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QMovie * movie()
func (this *QLabel) Movie() *qtgui.QMovie /*777 QMovie **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel5movieEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQMovieFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextFormat textFormat()
func (this *QLabel) TextFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel10textFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextFormat(Qt::TextFormat)
func (this *QLabel) SetTextFormat(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel13setTextFormatEN2Qt10TextFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment()
func (this *QLabel) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)
func (this *QLabel) SetAlignment(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordWrap(_Bool)
func (this *QLabel) SetWordWrap(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel11setWordWrapEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wordWrap()
func (this *QLabel) WordWrap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel8wordWrapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int indent()
func (this *QLabel) Indent() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel6indentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlabel.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndent(int)
func (this *QLabel) SetIndent(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel9setIndentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] int margin()
func (this *QLabel) Margin() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel6marginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlabel.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMargin(int)
func (this *QLabel) SetMargin(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel9setMarginEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasScaledContents()
func (this *QLabel) HasScaledContents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel17hasScaledContentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScaledContents(_Bool)
func (this *QLabel) SetScaledContents(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel17setScaledContentsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:100
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QLabel) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QLabel) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBuddy(QWidget *)
func (this *QLabel) SetBuddy(arg0 *QWidget /*777 QWidget **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel8setBuddyEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * buddy()
func (this *QLabel) Buddy() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel5buddyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:106
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int)
func (this *QLabel) HeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlabel.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool openExternalLinks()
func (this *QLabel) OpenExternalLinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel17openExternalLinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpenExternalLinks(_Bool)
func (this *QLabel) SetOpenExternalLinks(open bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel20setOpenExternalLinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), open)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextInteractionFlags(Qt::TextInteractionFlags)
func (this *QLabel) SetTextInteractionFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextInteractionFlags textInteractionFlags()
func (this *QLabel) TextInteractionFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel20textInteractionFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelection(int, int)
func (this *QLabel) SetSelection(arg0 int, arg1 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel12setSelectionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasSelectedText()
func (this *QLabel) HasSelectedText() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel15hasSelectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedText()
func (this *QLabel) SelectedText() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel12selectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionStart()
func (this *QLabel) SelectionStart() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLabel14selectionStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlabel.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QLabel) SetText(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixmap(const QPixmap &)
func (this *QLabel) SetPixmap(arg0 *qtgui.QPixmap) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel9setPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPicture(const QPicture &)
func (this *QLabel) SetPicture(arg0 *qtgui.QPicture) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel10setPictureERK8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovie(QMovie *)
func (this *QLabel) SetMovie(movie *qtgui.QMovie /*777 QMovie **/) {
	var convArg0 = movie.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel8setMovieEP6QMovie", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNum(int)
func (this *QLabel) SetNum(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel6setNumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:129
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setNum(double)
func (this *QLabel) SetNum_1(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel6setNumEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QLabel) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void linkActivated(const QString &)
func (this *QLabel) LinkActivated(link *qtcore.QString) {
	var convArg0 = link.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel13linkActivatedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void linkHovered(const QString &)
func (this *QLabel) LinkHovered(link *qtcore.QString) {
	var convArg0 = link.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel11linkHoveredERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:137
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QLabel) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:138
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QLabel) KeyPressEvent(ev *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:139
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QLabel) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:140
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QLabel) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:141
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QLabel) MousePressEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QLabel) MouseMoveEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:143
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QLabel) MouseReleaseEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:145
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QLabel) ContextMenuEvent(ev *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:147
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QLabel) FocusInEvent(ev *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:148
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QLabel) FocusOutEvent(ev *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:149
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(_Bool)
func (this *QLabel) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLabel18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
