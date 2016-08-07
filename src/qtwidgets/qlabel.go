package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
import "qtgui"
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
extern void* C_ZNK6QLabel4textEv(void* qthis); // 4
  // proto:  void QLabel::setPixmap(const QPixmap & );
extern void C_ZN6QLabel9setPixmapERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  bool QLabel::wordWrap();
extern bool C_ZNK6QLabel8wordWrapEv(void* qthis); // 4
  // proto:  bool QLabel::hasSelectedText();
extern bool C_ZNK6QLabel15hasSelectedTextEv(void* qthis); // 4
  // proto:  void QLabel::setOpenExternalLinks(bool open);
extern void C_ZN6QLabel20setOpenExternalLinksEb(void* qthis, bool arg0); // 4
  // proto:  Qt::Alignment QLabel::alignment();
extern void C_ZNK6QLabel9alignmentEv(void* qthis); // 4
  // proto:  const QPixmap * QLabel::pixmap();
extern void* C_ZNK6QLabel6pixmapEv(void* qthis); // 4
  // proto:  QWidget * QLabel::buddy();
extern void* C_ZNK6QLabel5buddyEv(void* qthis); // 4
  // proto:  QMovie * QLabel::movie();
extern void* C_ZNK6QLabel5movieEv(void* qthis); // 4
  // proto:  void QLabel::setPicture(const QPicture & );
extern void C_ZN6QLabel10setPictureERK8QPicture(void* qthis, void* arg0); // 4
  // proto:  bool QLabel::hasScaledContents();
extern bool C_ZNK6QLabel17hasScaledContentsEv(void* qthis); // 4
  // proto:  void QLabel::~QLabel();
extern void C_ZN6QLabelD2Ev(void* qthis); // 4
  // proto:  void QLabel::setWordWrap(bool on);
extern void C_ZN6QLabel11setWordWrapEb(void* qthis, bool arg0); // 4
  // proto:  Qt::TextInteractionFlags QLabel::textInteractionFlags();
extern void C_ZNK6QLabel20textInteractionFlagsEv(void* qthis); // 4
  // proto:  const QPicture * QLabel::picture();
extern void* C_ZNK6QLabel7pictureEv(void* qthis); // 4
  // proto:  void QLabel::setIndent(int );
extern void C_ZN6QLabel9setIndentEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLabel::setMovie(QMovie * movie);
extern void C_ZN6QLabel8setMovieEP6QMovie(void* qthis, void* arg0); // 4
  // proto:  void QLabel::setMargin(int );
extern void C_ZN6QLabel9setMarginEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QLabel::sizeHint();
extern void* C_ZNK6QLabel8sizeHintEv(void* qthis); // 4
  // proto:  void QLabel::setScaledContents(bool );
extern void C_ZN6QLabel17setScaledContentsEb(void* qthis, bool arg0); // 4
  // proto:  int QLabel::indent();
extern int32_t C_ZNK6QLabel6indentEv(void* qthis); // 4
  // proto:  const QMetaObject * QLabel::metaObject();
extern void C_ZNK6QLabel10metaObjectEv(void* qthis); // 4
  // proto:  void QLabel::setSelection(int , int );
extern void C_ZN6QLabel12setSelectionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QSize QLabel::minimumSizeHint();
extern void* C_ZNK6QLabel15minimumSizeHintEv(void* qthis); // 4
  // proto:  bool QLabel::openExternalLinks();
extern bool C_ZNK6QLabel17openExternalLinksEv(void* qthis); // 4
  // proto:  void QLabel::setText(const QString & );
extern void C_ZN6QLabel7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QLabel::selectedText();
extern void* C_ZNK6QLabel12selectedTextEv(void* qthis); // 4
  // proto:  int QLabel::heightForWidth(int );
extern int32_t C_ZNK6QLabel14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLabel::setNum(int );
extern void C_ZN6QLabel6setNumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLabel::setNum(double );
extern void C_ZN6QLabel6setNumEd(void* qthis, double arg0); // 4
  // proto:  int QLabel::selectionStart();
extern int32_t C_ZNK6QLabel14selectionStartEv(void* qthis); // 4
  // proto:  Qt::TextFormat QLabel::textFormat();
extern void C_ZNK6QLabel10textFormatEv(void* qthis); // 4
  // proto:  int QLabel::margin();
extern int32_t C_ZNK6QLabel6marginEv(void* qthis); // 4
  // proto:  void QLabel::clear();
extern void C_ZN6QLabel5clearEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QLabel)=1
type QLabel struct {
  /*qbase*/ QFrame;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _linkActivated QLabel_linkActivated_signal;
//  _linkHovered QLabel_linkHovered_signal;
}

// setBuddy(class QWidget *)
func (this *QLabel) Setbuddy(args ...interface{}) () {
  // setBuddy(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel8setBuddyEP7QWidget
    // invoke: void setBuddy(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel8setBuddyEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setBuddy", args)
  }

  return
}

// text()
func (this *QLabel) Text(args ...interface{}) (ret interface{}) {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK6QLabel4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "text", args)
  }

  return
}

// setPixmap(const class QPixmap &)
func (this *QLabel) Setpixmap(args ...interface{}) () {
  // setPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPixmap{}) // "const QPixmap &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel9setPixmapERK7QPixmap
    // invoke: void setPixmap(const class QPixmap &)
    var arg0 = args[0].(*qtgui.QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel9setPixmapERK7QPixmap(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setPixmap", args)
  }

  return
}

// wordWrap()
func (this *QLabel) Wordwrap(args ...interface{}) (ret interface{}) {
  // wordWrap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel8wordWrapEv
    // invoke: bool wordWrap()
    var ret0 = C.C_ZNK6QLabel8wordWrapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "wordWrap", args)
  }

  return
}

// hasSelectedText()
func (this *QLabel) Hasselectedtext(args ...interface{}) (ret interface{}) {
  // hasSelectedText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel15hasSelectedTextEv
    // invoke: bool hasSelectedText()
    var ret0 = C.C_ZNK6QLabel15hasSelectedTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "hasSelectedText", args)
  }

  return
}

// setOpenExternalLinks(_Bool)
func (this *QLabel) Setopenexternallinks(args ...interface{}) () {
  // setOpenExternalLinks(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel20setOpenExternalLinksEb
    // invoke: void setOpenExternalLinks(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel20setOpenExternalLinksEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setOpenExternalLinks", args)
  }

  return
}

// alignment()
func (this *QLabel) Alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK6QLabel9alignmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "alignment", args)
  }

  return
}

// pixmap()
func (this *QLabel) Pixmap(args ...interface{}) (ret interface{}) {
  // pixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel6pixmapEv
    // invoke: const QPixmap * pixmap()
    var ret0 = C.C_ZNK6QLabel6pixmapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QPixmap{}) // "const QPixmap *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "pixmap", args)
  }

  return
}

// buddy()
func (this *QLabel) Buddy(args ...interface{}) (ret interface{}) {
  // buddy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel5buddyEv
    // invoke: QWidget * buddy()
    var ret0 = C.C_ZNK6QLabel5buddyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "buddy", args)
  }

  return
}

// movie()
func (this *QLabel) Movie(args ...interface{}) (ret interface{}) {
  // movie()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel5movieEv
    // invoke: QMovie * movie()
    var ret0 = C.C_ZNK6QLabel5movieEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QMovie{}) // "QMovie *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "movie", args)
  }

  return
}

// setPicture(const class QPicture &)
func (this *QLabel) Setpicture(args ...interface{}) () {
  // setPicture(const class QPicture &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPicture{}) // "const QPicture &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel10setPictureERK8QPicture
    // invoke: void setPicture(const class QPicture &)
    var arg0 = args[0].(*qtgui.QPicture).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel10setPictureERK8QPicture(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setPicture", args)
  }

  return
}

// hasScaledContents()
func (this *QLabel) Hasscaledcontents(args ...interface{}) (ret interface{}) {
  // hasScaledContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel17hasScaledContentsEv
    // invoke: bool hasScaledContents()
    var ret0 = C.C_ZNK6QLabel17hasScaledContentsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "hasScaledContents", args)
  }

  return
}

// ~QLabel()
func (this *QLabel) Freeqlabel(args ...interface{}) () {
  // ~QLabel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabelD0Ev
    // invoke: void ~QLabel()
    C.C_ZN6QLabelD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "~QLabel", args)
  }

  return
}

// setWordWrap(_Bool)
func (this *QLabel) Setwordwrap(args ...interface{}) () {
  // setWordWrap(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel11setWordWrapEb
    // invoke: void setWordWrap(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel11setWordWrapEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setWordWrap", args)
  }

  return
}

// textInteractionFlags()
func (this *QLabel) Textinteractionflags(args ...interface{}) () {
  // textInteractionFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel20textInteractionFlagsEv
    // invoke: Qt::TextInteractionFlags textInteractionFlags()
    C.C_ZNK6QLabel20textInteractionFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "textInteractionFlags", args)
  }

  return
}

// picture()
func (this *QLabel) Picture(args ...interface{}) (ret interface{}) {
  // picture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel7pictureEv
    // invoke: const QPicture * picture()
    var ret0 = C.C_ZNK6QLabel7pictureEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QPicture{}) // "const QPicture *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "picture", args)
  }

  return
}

// setIndent(int)
func (this *QLabel) Setindent(args ...interface{}) () {
  // setIndent(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel9setIndentEi
    // invoke: void setIndent(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel9setIndentEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setIndent", args)
  }

  return
}

// setMovie(class QMovie *)
func (this *QLabel) Setmovie(args ...interface{}) () {
  // setMovie(class QMovie *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QMovie{}) // "QMovie *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel8setMovieEP6QMovie
    // invoke: void setMovie(class QMovie *)
    var arg0 = args[0].(*qtgui.QMovie).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel8setMovieEP6QMovie(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setMovie", args)
  }

  return
}

// setMargin(int)
func (this *QLabel) Setmargin(args ...interface{}) () {
  // setMargin(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel9setMarginEi
    // invoke: void setMargin(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel9setMarginEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setMargin", args)
  }

  return
}

// sizeHint()
func (this *QLabel) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK6QLabel8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "sizeHint", args)
  }

  return
}

// setScaledContents(_Bool)
func (this *QLabel) Setscaledcontents(args ...interface{}) () {
  // setScaledContents(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel17setScaledContentsEb
    // invoke: void setScaledContents(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel17setScaledContentsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setScaledContents", args)
  }

  return
}

// indent()
func (this *QLabel) Indent(args ...interface{}) (ret interface{}) {
  // indent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel6indentEv
    // invoke: int indent()
    var ret0 = C.C_ZNK6QLabel6indentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "indent", args)
  }

  return
}

// metaObject()
func (this *QLabel) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK6QLabel10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "metaObject", args)
  }

  return
}

// setSelection(int, int)
func (this *QLabel) Setselection(args ...interface{}) () {
  // setSelection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel12setSelectionEii
    // invoke: void setSelection(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN6QLabel12setSelectionEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLabel", "setSelection", args)
  }

  return
}

// minimumSizeHint()
func (this *QLabel) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK6QLabel15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "minimumSizeHint", args)
  }

  return
}

// openExternalLinks()
func (this *QLabel) Openexternallinks(args ...interface{}) (ret interface{}) {
  // openExternalLinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel17openExternalLinksEv
    // invoke: bool openExternalLinks()
    var ret0 = C.C_ZNK6QLabel17openExternalLinksEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "openExternalLinks", args)
  }

  return
}

// setText(const class QString &)
func (this *QLabel) Settext(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setText", args)
  }

  return
}

// selectedText()
func (this *QLabel) Selectedtext(args ...interface{}) (ret interface{}) {
  // selectedText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel12selectedTextEv
    // invoke: QString selectedText()
    var ret0 = C.C_ZNK6QLabel12selectedTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "selectedText", args)
  }

  return
}

// heightForWidth(int)
func (this *QLabel) Heightforwidth(args ...interface{}) (ret interface{}) {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QLabel14heightForWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "heightForWidth", args)
  }

  return
}

// setNum(int)
func (this *QLabel) Setnum(args ...interface{}) () {
  // setNum(int)
  // setNum(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "double"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel6setNumEi
    // invoke: void setNum(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel6setNumEi(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN6QLabel6setNumEd
    // invoke: void setNum(double)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QLabel6setNumEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLabel", "setNum", args)
  }

  return
}

// selectionStart()
func (this *QLabel) Selectionstart(args ...interface{}) (ret interface{}) {
  // selectionStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel14selectionStartEv
    // invoke: int selectionStart()
    var ret0 = C.C_ZNK6QLabel14selectionStartEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "selectionStart", args)
  }

  return
}

// textFormat()
func (this *QLabel) Textformat(args ...interface{}) () {
  // textFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel10textFormatEv
    // invoke: Qt::TextFormat textFormat()
    C.C_ZNK6QLabel10textFormatEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "textFormat", args)
  }

  return
}

// margin()
func (this *QLabel) Margin(args ...interface{}) (ret interface{}) {
  // margin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QLabel6marginEv
    // invoke: int margin()
    var ret0 = C.C_ZNK6QLabel6marginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLabel", "margin", args)
  }

  return
}

// clear()
func (this *QLabel) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QLabel5clearEv
    // invoke: void clear()
    C.C_ZN6QLabel5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLabel", "clear", args)
  }

  return
}

// <= body block end

