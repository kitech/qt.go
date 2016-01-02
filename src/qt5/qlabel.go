package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qlabel.h
// dst-file: /src/widgets/qlabel.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QPicture * QLabel::picture();
extern void _ZNK6QLabel7pictureEv(void* qthis);
  // proto:  void QLabel::setNum(double );
extern void _ZN6QLabel6setNumEd(void* qthis, double arg0);
  // proto:  void QLabel::setPicture(const QPicture & );
extern void _ZN6QLabel10setPictureERK8QPicture(void* qthis, void* arg0);
  // proto:  void QLabel::setText(const QString & );
extern void _ZN6QLabel7setTextERK7QString(void* qthis, void* arg0);
  // proto:  const QPixmap * QLabel::pixmap();
extern void _ZNK6QLabel6pixmapEv(void* qthis);
  // proto:  void QLabel::setIndent(int );
extern void _ZN6QLabel9setIndentEi(void* qthis, int arg0);
  // proto:  const QMetaObject * QLabel::metaObject();
extern void _ZNK6QLabel10metaObjectEv(void* qthis);
  // proto:  void QLabel::~QLabel();
extern void _ZN6QLabelD0Ev(void* qthis);
  // proto:  void QLabel::setSelection(int , int );
extern void _ZN6QLabel12setSelectionEii(void* qthis, int arg0, int arg1);
  // proto:  bool QLabel::hasScaledContents();
extern void _ZNK6QLabel17hasScaledContentsEv(void* qthis);
  // proto:  QString QLabel::text();
extern void _ZNK6QLabel4textEv(void* qthis);
  // proto:  int QLabel::heightForWidth(int );
extern void _ZNK6QLabel14heightForWidthEi(void* qthis, int arg0);
  // proto:  bool QLabel::openExternalLinks();
extern void _ZNK6QLabel17openExternalLinksEv(void* qthis);
  // proto:  void QLabel::setNum(int );
extern void _ZN6QLabel6setNumEi(void* qthis, int arg0);
  // proto:  void QLabel::setPixmap(const QPixmap & );
extern void _ZN6QLabel9setPixmapERK7QPixmap(void* qthis, void* arg0);
  // proto:  void QLabel::setOpenExternalLinks(bool open);
extern void _ZN6QLabel20setOpenExternalLinksEb(void* qthis, bool arg0);
  // proto:  QWidget * QLabel::buddy();
extern void _ZNK6QLabel5buddyEv(void* qthis);
  // proto:  bool QLabel::wordWrap();
extern void _ZNK6QLabel8wordWrapEv(void* qthis);
  // proto:  void QLabel::setWordWrap(bool on);
extern void _ZN6QLabel11setWordWrapEb(void* qthis, bool arg0);
  // proto:  void QLabel::clear();
extern void _ZN6QLabel5clearEv(void* qthis);
  // proto:  void QLabel::setMargin(int );
extern void _ZN6QLabel9setMarginEi(void* qthis, int arg0);
  // proto:  QSize QLabel::minimumSizeHint();
extern void _ZNK6QLabel15minimumSizeHintEv(void* qthis);
  // proto:  int QLabel::selectionStart();
extern void _ZNK6QLabel14selectionStartEv(void* qthis);
  // proto:  bool QLabel::hasSelectedText();
extern void _ZNK6QLabel15hasSelectedTextEv(void* qthis);
  // proto:  void QLabel::setBuddy(QWidget * );
extern void _ZN6QLabel8setBuddyEP7QWidget(void* qthis, void* arg0);
  // proto:  void QLabel::QLabel(const QLabel & );
extern void* dector_ZN6QLabelC1ERKS_(void* arg0);
extern void _ZN6QLabelC1ERKS_(void* qthis, void* arg0);
  // proto:  int QLabel::indent();
extern void _ZNK6QLabel6indentEv(void* qthis);
  // proto:  QSize QLabel::sizeHint();
extern void _ZNK6QLabel8sizeHintEv(void* qthis);
  // proto:  int QLabel::margin();
extern void _ZNK6QLabel6marginEv(void* qthis);
  // proto:  QMovie * QLabel::movie();
extern void _ZNK6QLabel5movieEv(void* qthis);
  // proto:  void QLabel::setScaledContents(bool );
extern void _ZN6QLabel17setScaledContentsEb(void* qthis, bool arg0);
  // proto:  void QLabel::setMovie(QMovie * movie);
extern void _ZN6QLabel8setMovieEP6QMovie(void* qthis, void* arg0);
  // proto:  QString QLabel::selectedText();
extern void _ZNK6QLabel12selectedTextEv(void* qthis);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QLabel)=1
type QLabel struct {
  /*qbase*/ QFrame;
  qclsinst uint64 /* *mut c_void*/;
//  _linkActivated QLabel_linkActivated_signal;
//  _linkHovered QLabel_linkHovered_signal;
}

  // proto:  const QPicture * QLabel::picture();
func (this *QLabel) picture(args ...interface{}) () {
  // picture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel7pictureEv
  default:
    qtrt.ErrorResolve("QLabel", "picture", args)
  }

}

  // proto:  void QLabel::setNum(double );
func (this *QLabel) setNum(args ...interface{}) () {
  // setNum(double)
  // setNum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel6setNumEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN6QLabel6setNumEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "setNum", args)
  }

}

  // proto:  void QLabel::setPicture(const QPicture & );
func (this *QLabel) setPicture(args ...interface{}) () {
  // setPicture(const class QPicture &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPicture{}) // "const QPicture &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel10setPictureERK8QPicture
    var arg0 = args[0].(QPicture).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "setPicture", args)
  }

}

  // proto:  void QLabel::setText(const QString & );
func (this *QLabel) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel7setTextERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "setText", args)
  }

}

  // proto:  const QPixmap * QLabel::pixmap();
func (this *QLabel) pixmap(args ...interface{}) () {
  // pixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel6pixmapEv
  default:
    qtrt.ErrorResolve("QLabel", "pixmap", args)
  }

}

  // proto:  void QLabel::setIndent(int );
func (this *QLabel) setIndent(args ...interface{}) () {
  // setIndent(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel9setIndentEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "setIndent", args)
  }

}

  // proto:  const QMetaObject * QLabel::metaObject();
func (this *QLabel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel10metaObjectEv
  default:
    qtrt.ErrorResolve("QLabel", "metaObject", args)
  }

}

  // proto:  void QLabel::~QLabel();
func (this *QLabel) FreeQLabel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLabel", "~QLabel", args)
  }

}

  // proto:  void QLabel::setSelection(int , int );
func (this *QLabel) setSelection(args ...interface{}) () {
  // setSelection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel12setSelectionEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QLabel", "setSelection", args)
  }

}

  // proto:  bool QLabel::hasScaledContents();
func (this *QLabel) hasScaledContents(args ...interface{}) () {
  // hasScaledContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel17hasScaledContentsEv
  default:
    qtrt.ErrorResolve("QLabel", "hasScaledContents", args)
  }

}

  // proto:  QString QLabel::text();
func (this *QLabel) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel4textEv
  default:
    qtrt.ErrorResolve("QLabel", "text", args)
  }

}

  // proto:  int QLabel::heightForWidth(int );
func (this *QLabel) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel14heightForWidthEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "heightForWidth", args)
  }

}

  // proto:  bool QLabel::openExternalLinks();
func (this *QLabel) openExternalLinks(args ...interface{}) () {
  // openExternalLinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel17openExternalLinksEv
  default:
    qtrt.ErrorResolve("QLabel", "openExternalLinks", args)
  }

}

  // proto:  void QLabel::setPixmap(const QPixmap & );
func (this *QLabel) setPixmap(args ...interface{}) () {
  // setPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel9setPixmapERK7QPixmap
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "setPixmap", args)
  }

}

  // proto:  void QLabel::setOpenExternalLinks(bool open);
func (this *QLabel) setOpenExternalLinks(args ...interface{}) () {
  // setOpenExternalLinks(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel20setOpenExternalLinksEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "setOpenExternalLinks", args)
  }

}

  // proto:  QWidget * QLabel::buddy();
func (this *QLabel) buddy(args ...interface{}) () {
  // buddy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel5buddyEv
  default:
    qtrt.ErrorResolve("QLabel", "buddy", args)
  }

}

  // proto:  bool QLabel::wordWrap();
func (this *QLabel) wordWrap(args ...interface{}) () {
  // wordWrap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel8wordWrapEv
  default:
    qtrt.ErrorResolve("QLabel", "wordWrap", args)
  }

}

  // proto:  void QLabel::setWordWrap(bool on);
func (this *QLabel) setWordWrap(args ...interface{}) () {
  // setWordWrap(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel11setWordWrapEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "setWordWrap", args)
  }

}

  // proto:  void QLabel::clear();
func (this *QLabel) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel5clearEv
  default:
    qtrt.ErrorResolve("QLabel", "clear", args)
  }

}

  // proto:  void QLabel::setMargin(int );
func (this *QLabel) setMargin(args ...interface{}) () {
  // setMargin(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel9setMarginEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "setMargin", args)
  }

}

  // proto:  QSize QLabel::minimumSizeHint();
func (this *QLabel) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QLabel", "minimumSizeHint", args)
  }

}

  // proto:  int QLabel::selectionStart();
func (this *QLabel) selectionStart(args ...interface{}) () {
  // selectionStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel14selectionStartEv
  default:
    qtrt.ErrorResolve("QLabel", "selectionStart", args)
  }

}

  // proto:  bool QLabel::hasSelectedText();
func (this *QLabel) hasSelectedText(args ...interface{}) () {
  // hasSelectedText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel15hasSelectedTextEv
  default:
    qtrt.ErrorResolve("QLabel", "hasSelectedText", args)
  }

}

  // proto:  void QLabel::setBuddy(QWidget * );
func (this *QLabel) setBuddy(args ...interface{}) () {
  // setBuddy(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel8setBuddyEP7QWidget
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "setBuddy", args)
  }

}

  // proto:  void QLabel::QLabel(const QLabel & );
func NewQLabel(args ...interface{}) QLabel {
  return QLabel{}
}

  // proto:  int QLabel::indent();
func (this *QLabel) indent(args ...interface{}) () {
  // indent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel6indentEv
  default:
    qtrt.ErrorResolve("QLabel", "indent", args)
  }

}

  // proto:  QSize QLabel::sizeHint();
func (this *QLabel) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel8sizeHintEv
  default:
    qtrt.ErrorResolve("QLabel", "sizeHint", args)
  }

}

  // proto:  int QLabel::margin();
func (this *QLabel) margin(args ...interface{}) () {
  // margin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel6marginEv
  default:
    qtrt.ErrorResolve("QLabel", "margin", args)
  }

}

  // proto:  QMovie * QLabel::movie();
func (this *QLabel) movie(args ...interface{}) () {
  // movie()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel5movieEv
  default:
    qtrt.ErrorResolve("QLabel", "movie", args)
  }

}

  // proto:  void QLabel::setScaledContents(bool );
func (this *QLabel) setScaledContents(args ...interface{}) () {
  // setScaledContents(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel17setScaledContentsEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "setScaledContents", args)
  }

}

  // proto:  void QLabel::setMovie(QMovie * movie);
func (this *QLabel) setMovie(args ...interface{}) () {
  // setMovie(class QMovie *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMovie{}) // "QMovie *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel8setMovieEP6QMovie
    var arg0 = args[0].(QMovie).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QLabel", "setMovie", args)
  }

}

  // proto:  QString QLabel::selectedText();
func (this *QLabel) selectedText(args ...interface{}) () {
  // selectedText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel12selectedTextEv
  default:
    qtrt.ErrorResolve("QLabel", "selectedText", args)
  }

}

// <= body block end

