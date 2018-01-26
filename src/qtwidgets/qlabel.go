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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QLabel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:70
// index:0
// Public
// void QLabel(class QWidget *, Qt::WindowFlags)
func NewQLabel(parent *QWidget /*777 QWidget **/, f int) *QLabel {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabelC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQLabelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:71
// index:1
// Public
// void QLabel(const class QString &, class QWidget *, Qt::WindowFlags)
func NewQLabel_1(text *qtcore.QString, parent *QWidget /*777 QWidget **/, f int) *QLabel {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabelC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQLabelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:72
// index:0
// Public virtual
// void ~QLabel()
func DeleteQLabel(*QLabel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:74
// index:0
// Public
// QString text()
func (this *QLabel) Text() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel4textEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:75
// index:0
// Public
// const QPixmap * pixmap()
func (this *QLabel) Pixmap() *qtgui.QPixmap /*777 const QPixmap **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel6pixmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:77
// index:0
// Public
// const QPicture * picture()
func (this *QLabel) Picture() *qtgui.QPicture /*777 const QPicture **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel7pictureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPictureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:80
// index:0
// Public
// QMovie * movie()
func (this *QLabel) Movie() *qtgui.QMovie /*777 QMovie **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel5movieEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQMovieFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:83
// index:0
// Public
// Qt::TextFormat textFormat()
func (this *QLabel) TextFormat() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel10textFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:84
// index:0
// Public
// void setTextFormat(Qt::TextFormat)
func (this *QLabel) SetTextFormat(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel13setTextFormatEN2Qt10TextFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:86
// index:0
// Public
// Qt::Alignment alignment()
func (this *QLabel) Alignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:89
// index:0
// Public
// void setWordWrap(_Bool)
func (this *QLabel) SetWordWrap(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel11setWordWrapEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:90
// index:0
// Public
// bool wordWrap()
func (this *QLabel) WordWrap() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel8wordWrapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:92
// index:0
// Public
// int indent()
func (this *QLabel) Indent() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel6indentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qlabel.h:93
// index:0
// Public
// void setIndent(int)
func (this *QLabel) SetIndent(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel9setIndentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:95
// index:0
// Public
// int margin()
func (this *QLabel) Margin() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel6marginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qlabel.h:96
// index:0
// Public
// void setMargin(int)
func (this *QLabel) SetMargin(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel9setMarginEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:98
// index:0
// Public
// bool hasScaledContents()
func (this *QLabel) HasScaledContents() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel17hasScaledContentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:99
// index:0
// Public
// void setScaledContents(_Bool)
func (this *QLabel) SetScaledContents(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel17setScaledContentsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:100
// index:0
// Public virtual
// QSize sizeHint()
func (this *QLabel) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:101
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QLabel) MinimumSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:103
// index:0
// Public
// void setBuddy(class QWidget *)
func (this *QLabel) SetBuddy(arg0 *QWidget /*777 QWidget **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel8setBuddyEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:104
// index:0
// Public
// QWidget * buddy()
func (this *QLabel) Buddy() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel5buddyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:106
// index:0
// Public virtual
// int heightForWidth(int)
func (this *QLabel) HeightForWidth(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qlabel.h:108
// index:0
// Public
// bool openExternalLinks()
func (this *QLabel) OpenExternalLinks() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel17openExternalLinksEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:109
// index:0
// Public
// void setOpenExternalLinks(_Bool)
func (this *QLabel) SetOpenExternalLinks(open bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel20setOpenExternalLinksEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), open)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:111
// index:0
// Public
// void setTextInteractionFlags(Qt::TextInteractionFlags)
func (this *QLabel) SetTextInteractionFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:112
// index:0
// Public
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QLabel) TextInteractionFlags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel20textInteractionFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:114
// index:0
// Public
// void setSelection(int, int)
func (this *QLabel) SetSelection(arg0 int, arg1 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel12setSelectionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:115
// index:0
// Public
// bool hasSelectedText()
func (this *QLabel) HasSelectedText() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel15hasSelectedTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:116
// index:0
// Public
// QString selectedText()
func (this *QLabel) SelectedText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel12selectedTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlabel.h:117
// index:0
// Public
// int selectionStart()
func (this *QLabel) SelectionStart() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel14selectionStartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qlabel.h:120
// index:0
// Public
// void setText(const class QString &)
func (this *QLabel) SetText(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:121
// index:0
// Public
// void setPixmap(const class QPixmap &)
func (this *QLabel) SetPixmap(arg0 *qtgui.QPixmap) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel9setPixmapERK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:123
// index:0
// Public
// void setPicture(const class QPicture &)
func (this *QLabel) SetPicture(arg0 *qtgui.QPicture) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel10setPictureERK8QPicture", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:126
// index:0
// Public
// void setMovie(class QMovie *)
func (this *QLabel) SetMovie(movie *qtgui.QMovie /*777 QMovie **/) {
	var convArg0 = movie.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel8setMovieEP6QMovie", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:128
// index:0
// Public
// void setNum(int)
func (this *QLabel) SetNum(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel6setNumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:129
// index:1
// Public
// void setNum(double)
func (this *QLabel) SetNum_1(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel6setNumEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:130
// index:0
// Public
// void clear()
func (this *QLabel) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:133
// index:0
// Public
// void linkActivated(const class QString &)
func (this *QLabel) LinkActivated(link *qtcore.QString) {
	var convArg0 = link.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel13linkActivatedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:134
// index:0
// Public
// void linkHovered(const class QString &)
func (this *QLabel) LinkHovered(link *qtcore.QString) {
	var convArg0 = link.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel11linkHoveredERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:137
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QLabel) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlabel.h:138
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QLabel) KeyPressEvent(ev *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:139
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QLabel) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:140
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QLabel) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:141
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QLabel) MousePressEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:142
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QLabel) MouseMoveEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:143
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QLabel) MouseReleaseEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:145
// index:0
// Protected virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QLabel) ContextMenuEvent(ev *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:147
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QLabel) FocusInEvent(ev *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:148
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QLabel) FocusOutEvent(ev *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:149
// index:0
// Protected virtual
// bool focusNextPrevChild(_Bool)
func (this *QLabel) FocusNextPrevChild(next bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
