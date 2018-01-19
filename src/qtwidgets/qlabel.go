//  header block begin
// /usr/include/qt/QtWidgets/qlabel.h
// #include <qlabel.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qlabel.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QLabel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:72
// index:0
// virtual
// void ~QLabel()
func DeleteQLabel(*QLabel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:74
// index:0
// QString text()
func (this *QLabel) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:75
// index:0
// const QPixmap * pixmap()
func (this *QLabel) Pixmap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel6pixmapEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:77
// index:0
// const QPicture * picture()
func (this *QLabel) Picture() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel7pictureEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:80
// index:0
// QMovie * movie()
func (this *QLabel) Movie() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel5movieEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:83
// index:0
// Qt::TextFormat textFormat()
func (this *QLabel) TextFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel10textFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:84
// index:0
// void setTextFormat(Qt::TextFormat)
func (this *QLabel) SetTextFormat(arg0 int) {
	// 0: (, Qt::TextFormat arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel13setTextFormatEN2Qt10TextFormatE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:86
// index:0
// Qt::Alignment alignment()
func (this *QLabel) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel9alignmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:89
// index:0
// void setWordWrap(_Bool)
func (this *QLabel) SetWordWrap(on bool) {
	// 0: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel11setWordWrapEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:90
// index:0
// bool wordWrap()
func (this *QLabel) WordWrap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel8wordWrapEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:92
// index:0
// int indent()
func (this *QLabel) Indent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel6indentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:93
// index:0
// void setIndent(int)
func (this *QLabel) SetIndent(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel9setIndentEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:95
// index:0
// int margin()
func (this *QLabel) Margin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel6marginEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:96
// index:0
// void setMargin(int)
func (this *QLabel) SetMargin(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel9setMarginEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:98
// index:0
// bool hasScaledContents()
func (this *QLabel) HasScaledContents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel17hasScaledContentsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:99
// index:0
// void setScaledContents(_Bool)
func (this *QLabel) SetScaledContents(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel17setScaledContentsEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:100
// index:0
// virtual
// QSize sizeHint()
func (this *QLabel) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:101
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QLabel) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:103
// index:0
// void setBuddy(class QWidget *)
func (this *QLabel) SetBuddy(arg0 unsafe.Pointer) {
	// 0: (, QWidget * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel8setBuddyEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:104
// index:0
// QWidget * buddy()
func (this *QLabel) Buddy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel5buddyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:106
// index:0
// virtual
// int heightForWidth(int)
func (this *QLabel) HeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:108
// index:0
// bool openExternalLinks()
func (this *QLabel) OpenExternalLinks() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel17openExternalLinksEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:109
// index:0
// void setOpenExternalLinks(_Bool)
func (this *QLabel) SetOpenExternalLinks(open bool) {
	// 0: (, bool open), (&open)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel20setOpenExternalLinksEb", ffiqt.FFI_TYPE_VOID, this.cthis, &open)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:112
// index:0
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QLabel) TextInteractionFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel20textInteractionFlagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:114
// index:0
// void setSelection(int, int)
func (this *QLabel) SetSelection(arg0 int, arg1 int) {
	// 0: (, int arg0, int arg1), (&arg0, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel12setSelectionEii", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:115
// index:0
// bool hasSelectedText()
func (this *QLabel) HasSelectedText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel15hasSelectedTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:116
// index:0
// QString selectedText()
func (this *QLabel) SelectedText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel12selectedTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:117
// index:0
// int selectionStart()
func (this *QLabel) SelectionStart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLabel14selectionStartEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:120
// index:0
// void setText(const class QString &)
func (this *QLabel) SetText(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:121
// index:0
// void setPixmap(const class QPixmap &)
func (this *QLabel) SetPixmap(arg0 unsafe.Pointer) {
	// 0: (, const QPixmap & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel9setPixmapERK7QPixmap", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:123
// index:0
// void setPicture(const class QPicture &)
func (this *QLabel) SetPicture(arg0 unsafe.Pointer) {
	// 0: (, const QPicture & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel10setPictureERK8QPicture", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:126
// index:0
// void setMovie(class QMovie *)
func (this *QLabel) SetMovie(movie unsafe.Pointer) {
	// 0: (, QMovie * movie), (movie)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel8setMovieEP6QMovie", ffiqt.FFI_TYPE_VOID, this.cthis, movie)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:128
// index:0
// void setNum(int)
func (this *QLabel) SetNum(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel6setNumEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:129
// index:1
// void setNum(double)
func (this *QLabel) SetNum_1(arg0 float64) {
	// 1: (, double arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel6setNumEd", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:130
// index:0
// void clear()
func (this *QLabel) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:133
// index:0
// void linkActivated(const class QString &)
func (this *QLabel) LinkActivated(link unsafe.Pointer) {
	// 0: (, const QString & link), (link)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel13linkActivatedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, link)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:134
// index:0
// void linkHovered(const class QString &)
func (this *QLabel) LinkHovered(link unsafe.Pointer) {
	// 0: (, const QString & link), (link)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLabel11linkHoveredERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, link)
	gopp.ErrPrint(err, rv)
}

//  body block end
