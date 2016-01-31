package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QLabel::setBuddy(QWidget * );
extern void C_ZN6QLabel8setBuddyEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QString QLabel::text();
extern void C_ZNK6QLabel4textEv(void* qthis); // 4
  // proto:  void QLabel::setPixmap(const QPixmap & );
extern void C_ZN6QLabel9setPixmapERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  bool QLabel::wordWrap();
extern void C_ZNK6QLabel8wordWrapEv(void* qthis); // 4
  // proto:  bool QLabel::hasSelectedText();
extern void C_ZNK6QLabel15hasSelectedTextEv(void* qthis); // 4
  // proto:  void QLabel::setOpenExternalLinks(bool open);
extern void C_ZN6QLabel20setOpenExternalLinksEb(void* qthis, bool arg0); // 4
  // proto:  Qt::Alignment QLabel::alignment();
extern void C_ZNK6QLabel9alignmentEv(void* qthis); // 4
  // proto:  const QPixmap * QLabel::pixmap();
extern void C_ZNK6QLabel6pixmapEv(void* qthis); // 4
  // proto:  QWidget * QLabel::buddy();
extern void C_ZNK6QLabel5buddyEv(void* qthis); // 4
  // proto:  QMovie * QLabel::movie();
extern void C_ZNK6QLabel5movieEv(void* qthis); // 4
  // proto:  void QLabel::setPicture(const QPicture & );
extern void C_ZN6QLabel10setPictureERK8QPicture(void* qthis, void* arg0); // 4
  // proto:  bool QLabel::hasScaledContents();
extern void C_ZNK6QLabel17hasScaledContentsEv(void* qthis); // 4
  // proto:  void QLabel::~QLabel();
extern void C_ZN6QLabelD2Ev(void* qthis); // 4
  // proto:  void QLabel::setWordWrap(bool on);
extern void C_ZN6QLabel11setWordWrapEb(void* qthis, bool arg0); // 4
  // proto:  Qt::TextInteractionFlags QLabel::textInteractionFlags();
extern void C_ZNK6QLabel20textInteractionFlagsEv(void* qthis); // 4
  // proto:  const QPicture * QLabel::picture();
extern void C_ZNK6QLabel7pictureEv(void* qthis); // 4
  // proto:  void QLabel::setIndent(int );
extern void C_ZN6QLabel9setIndentEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLabel::setMovie(QMovie * movie);
extern void C_ZN6QLabel8setMovieEP6QMovie(void* qthis, void* arg0); // 4
  // proto:  void QLabel::setMargin(int );
extern void C_ZN6QLabel9setMarginEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QLabel::sizeHint();
extern void C_ZNK6QLabel8sizeHintEv(void* qthis); // 4
  // proto:  void QLabel::setScaledContents(bool );
extern void C_ZN6QLabel17setScaledContentsEb(void* qthis, bool arg0); // 4
  // proto:  int QLabel::indent();
extern void C_ZNK6QLabel6indentEv(void* qthis); // 4
  // proto:  const QMetaObject * QLabel::metaObject();
extern void C_ZNK6QLabel10metaObjectEv(void* qthis); // 4
  // proto:  void QLabel::setSelection(int , int );
extern void C_ZN6QLabel12setSelectionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QSize QLabel::minimumSizeHint();
extern void C_ZNK6QLabel15minimumSizeHintEv(void* qthis); // 4
  // proto:  bool QLabel::openExternalLinks();
extern void C_ZNK6QLabel17openExternalLinksEv(void* qthis); // 4
  // proto:  void QLabel::setText(const QString & );
extern void C_ZN6QLabel7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QLabel::selectedText();
extern void C_ZNK6QLabel12selectedTextEv(void* qthis); // 4
  // proto:  int QLabel::heightForWidth(int );
extern void C_ZNK6QLabel14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLabel::setNum(int );
extern void C_ZN6QLabel6setNumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLabel::setNum(double );
extern void C_ZN6QLabel6setNumEd(void* qthis, double arg0); // 4
  // proto:  int QLabel::selectionStart();
extern void C_ZNK6QLabel14selectionStartEv(void* qthis); // 4
  // proto:  Qt::TextFormat QLabel::textFormat();
extern void C_ZNK6QLabel10textFormatEv(void* qthis); // 4
  // proto:  int QLabel::margin();
extern void C_ZNK6QLabel6marginEv(void* qthis); // 4
  // proto:  void QLabel::clear();
extern void C_ZN6QLabel5clearEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _linkActivated QLabel_linkActivated_signal;
//  _linkHovered QLabel_linkHovered_signal;
}

// setBuddy(class QWidget *)
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
    // invoke: void setBuddy(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel8setBuddyEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setBuddy", args)
  }

}

// text()
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
    // invoke: QString text()
    C.C_ZNK6QLabel4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "text", args)
  }

}

// setPixmap(const class QPixmap &)
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
    // invoke: void setPixmap(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel9setPixmapERK7QPixmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setPixmap", args)
  }

}

// wordWrap()
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
    // invoke: bool wordWrap()
    C.C_ZNK6QLabel8wordWrapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "wordWrap", args)
  }

}

// hasSelectedText()
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
    // invoke: bool hasSelectedText()
    C.C_ZNK6QLabel15hasSelectedTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "hasSelectedText", args)
  }

}

// setOpenExternalLinks(_Bool)
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
    // invoke: void setOpenExternalLinks(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel20setOpenExternalLinksEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setOpenExternalLinks", args)
  }

}

// alignment()
func (this *QLabel) alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK6QLabel9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "alignment", args)
  }

}

// pixmap()
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
    // invoke: const QPixmap * pixmap()
    C.C_ZNK6QLabel6pixmapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "pixmap", args)
  }

}

// buddy()
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
    // invoke: QWidget * buddy()
    C.C_ZNK6QLabel5buddyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "buddy", args)
  }

}

// movie()
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
    // invoke: QMovie * movie()
    C.C_ZNK6QLabel5movieEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "movie", args)
  }

}

// setPicture(const class QPicture &)
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
    // invoke: void setPicture(const class QPicture &)
    var arg0 = args[0].(QPicture).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel10setPictureERK8QPicture(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setPicture", args)
  }

}

// hasScaledContents()
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
    // invoke: bool hasScaledContents()
    C.C_ZNK6QLabel17hasScaledContentsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "hasScaledContents", args)
  }

}

// ~QLabel()
func (this *QLabel) FreeQLabel(args ...interface{}) () {
  // ~QLabel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabelD0Ev
    // invoke: void ~QLabel()
    C.C_ZN6QLabelD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "~QLabel", args)
  }

}

// setWordWrap(_Bool)
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
    // invoke: void setWordWrap(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel11setWordWrapEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setWordWrap", args)
  }

}

// textInteractionFlags()
func (this *QLabel) textInteractionFlags(args ...interface{}) () {
  // textInteractionFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel20textInteractionFlagsEv
    // invoke: Qt::TextInteractionFlags textInteractionFlags()
    C.C_ZNK6QLabel20textInteractionFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "textInteractionFlags", args)
  }

}

// picture()
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
    // invoke: const QPicture * picture()
    C.C_ZNK6QLabel7pictureEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "picture", args)
  }

}

// setIndent(int)
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
    // invoke: void setIndent(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel9setIndentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setIndent", args)
  }

}

// setMovie(class QMovie *)
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
    // invoke: void setMovie(class QMovie *)
    var arg0 = args[0].(QMovie).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel8setMovieEP6QMovie(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setMovie", args)
  }

}

// setMargin(int)
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
    // invoke: void setMargin(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel9setMarginEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setMargin", args)
  }

}

// sizeHint()
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
    // invoke: QSize sizeHint()
    C.C_ZNK6QLabel8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "sizeHint", args)
  }

}

// setScaledContents(_Bool)
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
    // invoke: void setScaledContents(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel17setScaledContentsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setScaledContents", args)
  }

}

// indent()
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
    // invoke: int indent()
    C.C_ZNK6QLabel6indentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "indent", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK6QLabel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "metaObject", args)
  }

}

// setSelection(int, int)
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
    // invoke: void setSelection(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN6QLabel12setSelectionEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLabel", "setSelection", args)
  }

}

// minimumSizeHint()
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
    // invoke: QSize minimumSizeHint()
    C.C_ZNK6QLabel15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "minimumSizeHint", args)
  }

}

// openExternalLinks()
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
    // invoke: bool openExternalLinks()
    C.C_ZNK6QLabel17openExternalLinksEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "openExternalLinks", args)
  }

}

// setText(const class QString &)
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
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setText", args)
  }

}

// selectedText()
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
    // invoke: QString selectedText()
    C.C_ZNK6QLabel12selectedTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "selectedText", args)
  }

}

// heightForWidth(int)
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
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK6QLabel14heightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "heightForWidth", args)
  }

}

// setNum(int)
func (this *QLabel) setNum(args ...interface{}) () {
  // setNum(int)
  // setNum(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel6setNumEi
    // invoke: void setNum(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel6setNumEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN6QLabel6setNumEd
    // invoke: void setNum(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel6setNumEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setNum", args)
  }

}

// selectionStart()
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
    // invoke: int selectionStart()
    C.C_ZNK6QLabel14selectionStartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "selectionStart", args)
  }

}

// textFormat()
func (this *QLabel) textFormat(args ...interface{}) () {
  // textFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel10textFormatEv
    // invoke: Qt::TextFormat textFormat()
    C.C_ZNK6QLabel10textFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "textFormat", args)
  }

}

// margin()
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
    // invoke: int margin()
    C.C_ZNK6QLabel6marginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "margin", args)
  }

}

// clear()
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
    // invoke: void clear()
    C.C_ZN6QLabel5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "clear", args)
  }

}

// <= body block end

