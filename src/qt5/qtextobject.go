package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qtextobject.h
// dst-file: /src/gui/qtextobject.go
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
  // proto:  QTextDocumentPrivate * QTextObject::docHandle();
extern void _ZNK11QTextObject9docHandleEv(void* qthis);
  // proto:  void QTextObject::~QTextObject();
extern void _ZN11QTextObjectD0Ev(void* qthis);
  // proto:  void QTextObject::QTextObject(const QTextObject & );
extern void* dector_ZN11QTextObjectC1ERKS_(void* arg0);
extern void _ZN11QTextObjectC1ERKS_(void* qthis, void* arg0);
  // proto:  QTextFormat QTextObject::format();
extern void _ZNK11QTextObject6formatEv(void* qthis);
  // proto:  int QTextObject::formatIndex();
extern void _ZNK11QTextObject11formatIndexEv(void* qthis);
  // proto:  QTextDocument * QTextObject::document();
extern void _ZNK11QTextObject8documentEv(void* qthis);
  // proto:  int QTextObject::objectIndex();
extern void _ZNK11QTextObject11objectIndexEv(void* qthis);
  // proto:  void QTextObject::QTextObject(QTextDocument * doc);
extern void* dector_ZN11QTextObjectC1EP13QTextDocument(void* arg0);
extern void _ZN11QTextObjectC1EP13QTextDocument(void* qthis, void* arg0);
  // proto:  const QMetaObject * QTextObject::metaObject();
extern void _ZNK11QTextObject10metaObjectEv(void* qthis);
  // proto:  void QTextBlockUserData::~QTextBlockUserData();
extern void _ZN18QTextBlockUserDataD0Ev(void* qthis);
  // proto:  int QTextFragment::charFormatIndex();
extern void _ZNK13QTextFragment15charFormatIndexEv(void* qthis);
  // proto:  int QTextFragment::position();
extern void _ZNK13QTextFragment8positionEv(void* qthis);
  // proto:  void QTextFragment::QTextFragment(const QTextFragment & o);
extern void* dector_ZN13QTextFragmentC1ERKS_(void* arg0);
extern void demth_ZN13QTextFragmentC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QTextFragment::contains(int position);
extern void _ZNK13QTextFragment8containsEi(void* qthis, int arg0);
  // proto:  void QTextFragment::QTextFragment();
extern void* dector_ZN13QTextFragmentC1Ev();
extern void demth_ZN13QTextFragmentC1Ev(void* qthis);
  // proto:  QString QTextFragment::text();
extern void _ZNK13QTextFragment4textEv(void* qthis);
  // proto:  QList<QGlyphRun> QTextFragment::glyphRuns(int from, int length);
extern void _ZNK13QTextFragment9glyphRunsEii(void* qthis, int arg0, int arg1);
  // proto:  bool QTextFragment::isValid();
extern void demth_ZNK13QTextFragment7isValidEv(void* qthis);
  // proto:  QTextCharFormat QTextFragment::charFormat();
extern void _ZNK13QTextFragment10charFormatEv(void* qthis);
  // proto:  int QTextFragment::length();
extern void _ZNK13QTextFragment6lengthEv(void* qthis);
  // proto:  void QTextFrameLayoutData::~QTextFrameLayoutData();
extern void _ZN20QTextFrameLayoutDataD0Ev(void* qthis);
  // proto:  const QTextDocument * QTextBlock::document();
extern void _ZNK10QTextBlock8documentEv(void* qthis);
  // proto:  QTextBlock QTextBlock::previous();
extern void _ZNK10QTextBlock8previousEv(void* qthis);
  // proto:  int QTextBlock::length();
extern void _ZNK10QTextBlock6lengthEv(void* qthis);
  // proto:  QTextBlockUserData * QTextBlock::userData();
extern void _ZNK10QTextBlock8userDataEv(void* qthis);
  // proto:  void QTextBlock::QTextBlock(const QTextBlock & o);
extern void* dector_ZN10QTextBlockC1ERKS_(void* arg0);
extern void demth_ZN10QTextBlockC1ERKS_(void* qthis, void* arg0);
  // proto:  QString QTextBlock::text();
extern void _ZNK10QTextBlock4textEv(void* qthis);
  // proto:  int QTextBlock::lineCount();
extern void _ZNK10QTextBlock9lineCountEv(void* qthis);
  // proto:  bool QTextBlock::contains(int position);
extern void _ZNK10QTextBlock8containsEi(void* qthis, int arg0);
  // proto:  int QTextBlock::blockNumber();
extern void _ZNK10QTextBlock11blockNumberEv(void* qthis);
  // proto:  void QTextBlock::setRevision(int rev);
extern void _ZN10QTextBlock11setRevisionEi(void* qthis, int arg0);
  // proto:  void QTextBlock::setVisible(bool visible);
extern void _ZN10QTextBlock10setVisibleEb(void* qthis, bool arg0);
  // proto:  void QTextBlock::clearLayout();
extern void _ZN10QTextBlock11clearLayoutEv(void* qthis);
  // proto:  QTextDocumentPrivate * QTextBlock::docHandle();
extern void demth_ZNK10QTextBlock9docHandleEv(void* qthis);
  // proto:  int QTextBlock::userState();
extern void _ZNK10QTextBlock9userStateEv(void* qthis);
  // proto:  int QTextBlock::charFormatIndex();
extern void _ZNK10QTextBlock15charFormatIndexEv(void* qthis);
  // proto:  int QTextBlock::revision();
extern void _ZNK10QTextBlock8revisionEv(void* qthis);
  // proto:  int QTextBlock::position();
extern void _ZNK10QTextBlock8positionEv(void* qthis);
  // proto:  bool QTextBlock::isValid();
extern void _ZNK10QTextBlock7isValidEv(void* qthis);
  // proto:  QTextList * QTextBlock::textList();
extern void _ZNK10QTextBlock8textListEv(void* qthis);
  // proto:  QTextLayout * QTextBlock::layout();
extern void _ZNK10QTextBlock6layoutEv(void* qthis);
  // proto:  void QTextBlock::setUserData(QTextBlockUserData * data);
extern void _ZN10QTextBlock11setUserDataEP18QTextBlockUserData(void* qthis, void* arg0);
  // proto:  int QTextBlock::blockFormatIndex();
extern void _ZNK10QTextBlock16blockFormatIndexEv(void* qthis);
  // proto:  void QTextBlock::setUserState(int state);
extern void _ZN10QTextBlock12setUserStateEi(void* qthis, int arg0);
  // proto:  int QTextBlock::fragmentIndex();
extern void demth_ZNK10QTextBlock13fragmentIndexEv(void* qthis);
  // proto:  bool QTextBlock::isVisible();
extern void _ZNK10QTextBlock9isVisibleEv(void* qthis);
  // proto:  void QTextBlock::setLineCount(int count);
extern void _ZN10QTextBlock12setLineCountEi(void* qthis, int arg0);
  // proto:  QTextBlock QTextBlock::next();
extern void _ZNK10QTextBlock4nextEv(void* qthis);
  // proto:  QTextBlockFormat QTextBlock::blockFormat();
extern void _ZNK10QTextBlock11blockFormatEv(void* qthis);
  // proto:  void QTextBlock::QTextBlock();
extern void* dector_ZN10QTextBlockC1Ev();
extern void demth_ZN10QTextBlockC1Ev(void* qthis);
  // proto:  int QTextBlock::firstLineNumber();
extern void _ZNK10QTextBlock15firstLineNumberEv(void* qthis);
  // proto:  QTextCharFormat QTextBlock::charFormat();
extern void _ZNK10QTextBlock10charFormatEv(void* qthis);
  // proto:  void QTextBlockGroup::QTextBlockGroup(const QTextBlockGroup & );
extern void* dector_ZN15QTextBlockGroupC1ERKS_(void* arg0);
extern void _ZN15QTextBlockGroupC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QTextBlockGroup::metaObject();
extern void _ZNK15QTextBlockGroup10metaObjectEv(void* qthis);
  // proto:  void QTextBlockGroup::~QTextBlockGroup();
extern void _ZN15QTextBlockGroupD0Ev(void* qthis);
  // proto:  void QTextBlockGroup::QTextBlockGroup(QTextDocument * doc);
extern void* dector_ZN15QTextBlockGroupC1EP13QTextDocument(void* arg0);
extern void _ZN15QTextBlockGroupC1EP13QTextDocument(void* qthis, void* arg0);
  // proto:  QTextFrameFormat QTextFrame::frameFormat();
extern void _ZNK10QTextFrame11frameFormatEv(void* qthis);
  // proto:  QTextFrameLayoutData * QTextFrame::layoutData();
extern void _ZNK10QTextFrame10layoutDataEv(void* qthis);
  // proto:  void QTextFrame::setLayoutData(QTextFrameLayoutData * data);
extern void _ZN10QTextFrame13setLayoutDataEP20QTextFrameLayoutData(void* qthis, void* arg0);
  // proto:  void QTextFrame::setFrameFormat(const QTextFrameFormat & format);
extern void demth_ZN10QTextFrame14setFrameFormatERK16QTextFrameFormat(void* qthis, void* arg0);
  // proto:  void QTextFrame::QTextFrame(const QTextFrame & );
extern void* dector_ZN10QTextFrameC1ERKS_(void* arg0);
extern void _ZN10QTextFrameC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QTextFrame::metaObject();
extern void _ZNK10QTextFrame10metaObjectEv(void* qthis);
  // proto:  QTextFrame * QTextFrame::parentFrame();
extern void _ZNK10QTextFrame11parentFrameEv(void* qthis);
  // proto:  int QTextFrame::firstPosition();
extern void _ZNK10QTextFrame13firstPositionEv(void* qthis);
  // proto:  QList<QTextFrame *> QTextFrame::childFrames();
extern void _ZNK10QTextFrame11childFramesEv(void* qthis);
  // proto:  void QTextFrame::~QTextFrame();
extern void _ZN10QTextFrameD0Ev(void* qthis);
  // proto:  QTextCursor QTextFrame::lastCursorPosition();
extern void _ZNK10QTextFrame18lastCursorPositionEv(void* qthis);
  // proto:  void QTextFrame::QTextFrame(QTextDocument * doc);
extern void* dector_ZN10QTextFrameC1EP13QTextDocument(void* arg0);
extern void _ZN10QTextFrameC1EP13QTextDocument(void* qthis, void* arg0);
  // proto:  int QTextFrame::lastPosition();
extern void _ZNK10QTextFrame12lastPositionEv(void* qthis);
  // proto:  QTextCursor QTextFrame::firstCursorPosition();
extern void _ZNK10QTextFrame19firstCursorPositionEv(void* qthis);
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

// class sizeof(QTextObject)=1
type QTextObject struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextBlockUserData)=8
type QTextBlockUserData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextFragment)=16
type QTextFragment struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextFrameLayoutData)=8
type QTextFrameLayoutData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextBlock)=16
type QTextBlock struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextBlockGroup)=1
type QTextBlockGroup struct {
  /*qbase*/ QTextObject;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextFrame)=1
type QTextFrame struct {
  /*qbase*/ QTextObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QTextDocumentPrivate * QTextObject::docHandle();
func (this *QTextObject) docHandle(args ...interface{}) () {
  // docHandle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextObject9docHandleEv
    // invoke: QTextDocumentPrivate * docHandle()
    C._ZNK11QTextObject9docHandleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "docHandle", args)
  }

}

  // proto:  void QTextObject::~QTextObject();
func (this *QTextObject) FreeQTextObject(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextObject", "~QTextObject", args)
  }

}

  // proto:  void QTextObject::QTextObject(const QTextObject & );
func NewQTextObject(args ...interface{}) QTextObject {
  return QTextObject{}
}

  // proto:  QTextFormat QTextObject::format();
func (this *QTextObject) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextObject6formatEv
    // invoke: QTextFormat format()
    C._ZNK11QTextObject6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "format", args)
  }

}

  // proto:  int QTextObject::formatIndex();
func (this *QTextObject) formatIndex(args ...interface{}) () {
  // formatIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextObject11formatIndexEv
    // invoke: int formatIndex()
    C._ZNK11QTextObject11formatIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "formatIndex", args)
  }

}

  // proto:  QTextDocument * QTextObject::document();
func (this *QTextObject) document(args ...interface{}) () {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextObject8documentEv
    // invoke: QTextDocument * document()
    C._ZNK11QTextObject8documentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "document", args)
  }

}

  // proto:  int QTextObject::objectIndex();
func (this *QTextObject) objectIndex(args ...interface{}) () {
  // objectIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextObject11objectIndexEv
    // invoke: int objectIndex()
    C._ZNK11QTextObject11objectIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "objectIndex", args)
  }

}

  // proto:  const QMetaObject * QTextObject::metaObject();
func (this *QTextObject) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextObject10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QTextObject10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "metaObject", args)
  }

}

  // proto:  void QTextBlockUserData::~QTextBlockUserData();
func (this *QTextBlockUserData) FreeQTextBlockUserData(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextBlockUserData", "~QTextBlockUserData", args)
  }

}

  // proto:  int QTextFragment::charFormatIndex();
func (this *QTextFragment) charFormatIndex(args ...interface{}) () {
  // charFormatIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextFragment15charFormatIndexEv
    // invoke: int charFormatIndex()
    C._ZNK13QTextFragment15charFormatIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "charFormatIndex", args)
  }

}

  // proto:  int QTextFragment::position();
func (this *QTextFragment) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextFragment8positionEv
    // invoke: int position()
    C._ZNK13QTextFragment8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "position", args)
  }

}

  // proto:  void QTextFragment::QTextFragment(const QTextFragment & o);
func NewQTextFragment(args ...interface{}) QTextFragment {
  return QTextFragment{}
}

  // proto:  bool QTextFragment::contains(int position);
func (this *QTextFragment) contains(args ...interface{}) () {
  // contains(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextFragment8containsEi
    // invoke: bool contains(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK13QTextFragment8containsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFragment", "contains", args)
  }

}

  // proto:  QString QTextFragment::text();
func (this *QTextFragment) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextFragment4textEv
    // invoke: QString text()
    C._ZNK13QTextFragment4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "text", args)
  }

}

  // proto:  QList<QGlyphRun> QTextFragment::glyphRuns(int from, int length);
func (this *QTextFragment) glyphRuns(args ...interface{}) () {
  // glyphRuns(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextFragment9glyphRunsEii
    // invoke: QList<QGlyphRun> glyphRuns(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK13QTextFragment9glyphRunsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextFragment", "glyphRuns", args)
  }

}

  // proto:  bool QTextFragment::isValid();
func (this *QTextFragment) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextFragment7isValidEv
    // invoke: bool isValid()
    C.demth_ZNK13QTextFragment7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "isValid", args)
  }

}

  // proto:  QTextCharFormat QTextFragment::charFormat();
func (this *QTextFragment) charFormat(args ...interface{}) () {
  // charFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextFragment10charFormatEv
    // invoke: QTextCharFormat charFormat()
    C._ZNK13QTextFragment10charFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "charFormat", args)
  }

}

  // proto:  int QTextFragment::length();
func (this *QTextFragment) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextFragment6lengthEv
    // invoke: int length()
    C._ZNK13QTextFragment6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "length", args)
  }

}

  // proto:  void QTextFrameLayoutData::~QTextFrameLayoutData();
func (this *QTextFrameLayoutData) FreeQTextFrameLayoutData(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextFrameLayoutData", "~QTextFrameLayoutData", args)
  }

}

  // proto:  const QTextDocument * QTextBlock::document();
func (this *QTextBlock) document(args ...interface{}) () {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock8documentEv
    // invoke: const QTextDocument * document()
    C._ZNK10QTextBlock8documentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "document", args)
  }

}

  // proto:  QTextBlock QTextBlock::previous();
func (this *QTextBlock) previous(args ...interface{}) () {
  // previous()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock8previousEv
    // invoke: QTextBlock previous()
    C._ZNK10QTextBlock8previousEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "previous", args)
  }

}

  // proto:  int QTextBlock::length();
func (this *QTextBlock) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock6lengthEv
    // invoke: int length()
    C._ZNK10QTextBlock6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "length", args)
  }

}

  // proto:  QTextBlockUserData * QTextBlock::userData();
func (this *QTextBlock) userData(args ...interface{}) () {
  // userData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock8userDataEv
    // invoke: QTextBlockUserData * userData()
    C._ZNK10QTextBlock8userDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "userData", args)
  }

}

  // proto:  void QTextBlock::QTextBlock(const QTextBlock & o);
func NewQTextBlock(args ...interface{}) QTextBlock {
  return QTextBlock{}
}

  // proto:  QString QTextBlock::text();
func (this *QTextBlock) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock4textEv
    // invoke: QString text()
    C._ZNK10QTextBlock4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "text", args)
  }

}

  // proto:  int QTextBlock::lineCount();
func (this *QTextBlock) lineCount(args ...interface{}) () {
  // lineCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock9lineCountEv
    // invoke: int lineCount()
    C._ZNK10QTextBlock9lineCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "lineCount", args)
  }

}

  // proto:  bool QTextBlock::contains(int position);
func (this *QTextBlock) contains(args ...interface{}) () {
  // contains(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock8containsEi
    // invoke: bool contains(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTextBlock8containsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "contains", args)
  }

}

  // proto:  int QTextBlock::blockNumber();
func (this *QTextBlock) blockNumber(args ...interface{}) () {
  // blockNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock11blockNumberEv
    // invoke: int blockNumber()
    C._ZNK10QTextBlock11blockNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "blockNumber", args)
  }

}

  // proto:  void QTextBlock::setRevision(int rev);
func (this *QTextBlock) setRevision(args ...interface{}) () {
  // setRevision(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextBlock11setRevisionEi
    // invoke: void setRevision(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTextBlock11setRevisionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setRevision", args)
  }

}

  // proto:  void QTextBlock::setVisible(bool visible);
func (this *QTextBlock) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextBlock10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN10QTextBlock10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setVisible", args)
  }

}

  // proto:  void QTextBlock::clearLayout();
func (this *QTextBlock) clearLayout(args ...interface{}) () {
  // clearLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextBlock11clearLayoutEv
    // invoke: void clearLayout()
    C._ZN10QTextBlock11clearLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "clearLayout", args)
  }

}

  // proto:  QTextDocumentPrivate * QTextBlock::docHandle();
func (this *QTextBlock) docHandle(args ...interface{}) () {
  // docHandle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock9docHandleEv
    // invoke: QTextDocumentPrivate * docHandle()
    C.demth_ZNK10QTextBlock9docHandleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "docHandle", args)
  }

}

  // proto:  int QTextBlock::userState();
func (this *QTextBlock) userState(args ...interface{}) () {
  // userState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock9userStateEv
    // invoke: int userState()
    C._ZNK10QTextBlock9userStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "userState", args)
  }

}

  // proto:  int QTextBlock::charFormatIndex();
func (this *QTextBlock) charFormatIndex(args ...interface{}) () {
  // charFormatIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock15charFormatIndexEv
    // invoke: int charFormatIndex()
    C._ZNK10QTextBlock15charFormatIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "charFormatIndex", args)
  }

}

  // proto:  int QTextBlock::revision();
func (this *QTextBlock) revision(args ...interface{}) () {
  // revision()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock8revisionEv
    // invoke: int revision()
    C._ZNK10QTextBlock8revisionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "revision", args)
  }

}

  // proto:  int QTextBlock::position();
func (this *QTextBlock) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock8positionEv
    // invoke: int position()
    C._ZNK10QTextBlock8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "position", args)
  }

}

  // proto:  bool QTextBlock::isValid();
func (this *QTextBlock) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock7isValidEv
    // invoke: bool isValid()
    C._ZNK10QTextBlock7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "isValid", args)
  }

}

  // proto:  QTextList * QTextBlock::textList();
func (this *QTextBlock) textList(args ...interface{}) () {
  // textList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock8textListEv
    // invoke: QTextList * textList()
    C._ZNK10QTextBlock8textListEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "textList", args)
  }

}

  // proto:  QTextLayout * QTextBlock::layout();
func (this *QTextBlock) layout(args ...interface{}) () {
  // layout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock6layoutEv
    // invoke: QTextLayout * layout()
    C._ZNK10QTextBlock6layoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "layout", args)
  }

}

  // proto:  void QTextBlock::setUserData(QTextBlockUserData * data);
func (this *QTextBlock) setUserData(args ...interface{}) () {
  // setUserData(class QTextBlockUserData *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlockUserData{}) // "QTextBlockUserData *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextBlock11setUserDataEP18QTextBlockUserData
    // invoke: void setUserData(class QTextBlockUserData *)
    var arg0 = args[0].(QTextBlockUserData).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTextBlock11setUserDataEP18QTextBlockUserData(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setUserData", args)
  }

}

  // proto:  int QTextBlock::blockFormatIndex();
func (this *QTextBlock) blockFormatIndex(args ...interface{}) () {
  // blockFormatIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock16blockFormatIndexEv
    // invoke: int blockFormatIndex()
    C._ZNK10QTextBlock16blockFormatIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "blockFormatIndex", args)
  }

}

  // proto:  void QTextBlock::setUserState(int state);
func (this *QTextBlock) setUserState(args ...interface{}) () {
  // setUserState(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextBlock12setUserStateEi
    // invoke: void setUserState(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTextBlock12setUserStateEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setUserState", args)
  }

}

  // proto:  int QTextBlock::fragmentIndex();
func (this *QTextBlock) fragmentIndex(args ...interface{}) () {
  // fragmentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock13fragmentIndexEv
    // invoke: int fragmentIndex()
    C.demth_ZNK10QTextBlock13fragmentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "fragmentIndex", args)
  }

}

  // proto:  bool QTextBlock::isVisible();
func (this *QTextBlock) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock9isVisibleEv
    // invoke: bool isVisible()
    C._ZNK10QTextBlock9isVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "isVisible", args)
  }

}

  // proto:  void QTextBlock::setLineCount(int count);
func (this *QTextBlock) setLineCount(args ...interface{}) () {
  // setLineCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextBlock12setLineCountEi
    // invoke: void setLineCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTextBlock12setLineCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setLineCount", args)
  }

}

  // proto:  QTextBlock QTextBlock::next();
func (this *QTextBlock) next(args ...interface{}) () {
  // next()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock4nextEv
    // invoke: QTextBlock next()
    C._ZNK10QTextBlock4nextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "next", args)
  }

}

  // proto:  QTextBlockFormat QTextBlock::blockFormat();
func (this *QTextBlock) blockFormat(args ...interface{}) () {
  // blockFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock11blockFormatEv
    // invoke: QTextBlockFormat blockFormat()
    C._ZNK10QTextBlock11blockFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "blockFormat", args)
  }

}

  // proto:  int QTextBlock::firstLineNumber();
func (this *QTextBlock) firstLineNumber(args ...interface{}) () {
  // firstLineNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock15firstLineNumberEv
    // invoke: int firstLineNumber()
    C._ZNK10QTextBlock15firstLineNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "firstLineNumber", args)
  }

}

  // proto:  QTextCharFormat QTextBlock::charFormat();
func (this *QTextBlock) charFormat(args ...interface{}) () {
  // charFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock10charFormatEv
    // invoke: QTextCharFormat charFormat()
    C._ZNK10QTextBlock10charFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "charFormat", args)
  }

}

  // proto:  void QTextBlockGroup::QTextBlockGroup(const QTextBlockGroup & );
func NewQTextBlockGroup(args ...interface{}) QTextBlockGroup {
  return QTextBlockGroup{}
}

  // proto:  const QMetaObject * QTextBlockGroup::metaObject();
func (this *QTextBlockGroup) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextBlockGroup10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK15QTextBlockGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockGroup", "metaObject", args)
  }

}

  // proto:  void QTextBlockGroup::~QTextBlockGroup();
func (this *QTextBlockGroup) FreeQTextBlockGroup(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextBlockGroup", "~QTextBlockGroup", args)
  }

}

  // proto:  QTextFrameFormat QTextFrame::frameFormat();
func (this *QTextFrame) frameFormat(args ...interface{}) () {
  // frameFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextFrame11frameFormatEv
    // invoke: QTextFrameFormat frameFormat()
    C._ZNK10QTextFrame11frameFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "frameFormat", args)
  }

}

  // proto:  QTextFrameLayoutData * QTextFrame::layoutData();
func (this *QTextFrame) layoutData(args ...interface{}) () {
  // layoutData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextFrame10layoutDataEv
    // invoke: QTextFrameLayoutData * layoutData()
    C._ZNK10QTextFrame10layoutDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "layoutData", args)
  }

}

  // proto:  void QTextFrame::setLayoutData(QTextFrameLayoutData * data);
func (this *QTextFrame) setLayoutData(args ...interface{}) () {
  // setLayoutData(class QTextFrameLayoutData *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFrameLayoutData{}) // "QTextFrameLayoutData *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextFrame13setLayoutDataEP20QTextFrameLayoutData
    // invoke: void setLayoutData(class QTextFrameLayoutData *)
    var arg0 = args[0].(QTextFrameLayoutData).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTextFrame13setLayoutDataEP20QTextFrameLayoutData(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrame", "setLayoutData", args)
  }

}

  // proto:  void QTextFrame::setFrameFormat(const QTextFrameFormat & format);
func (this *QTextFrame) setFrameFormat(args ...interface{}) () {
  // setFrameFormat(const class QTextFrameFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFrameFormat{}) // "const QTextFrameFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextFrame14setFrameFormatERK16QTextFrameFormat
    // invoke: void setFrameFormat(const class QTextFrameFormat &)
    var arg0 = args[0].(QTextFrameFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN10QTextFrame14setFrameFormatERK16QTextFrameFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrame", "setFrameFormat", args)
  }

}

  // proto:  void QTextFrame::QTextFrame(const QTextFrame & );
func NewQTextFrame(args ...interface{}) QTextFrame {
  return QTextFrame{}
}

  // proto:  const QMetaObject * QTextFrame::metaObject();
func (this *QTextFrame) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextFrame10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK10QTextFrame10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "metaObject", args)
  }

}

  // proto:  QTextFrame * QTextFrame::parentFrame();
func (this *QTextFrame) parentFrame(args ...interface{}) () {
  // parentFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextFrame11parentFrameEv
    // invoke: QTextFrame * parentFrame()
    C._ZNK10QTextFrame11parentFrameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "parentFrame", args)
  }

}

  // proto:  int QTextFrame::firstPosition();
func (this *QTextFrame) firstPosition(args ...interface{}) () {
  // firstPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextFrame13firstPositionEv
    // invoke: int firstPosition()
    C._ZNK10QTextFrame13firstPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "firstPosition", args)
  }

}

  // proto:  QList<QTextFrame *> QTextFrame::childFrames();
func (this *QTextFrame) childFrames(args ...interface{}) () {
  // childFrames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextFrame11childFramesEv
    // invoke: QList<QTextFrame *> childFrames()
    C._ZNK10QTextFrame11childFramesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "childFrames", args)
  }

}

  // proto:  void QTextFrame::~QTextFrame();
func (this *QTextFrame) FreeQTextFrame(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextFrame", "~QTextFrame", args)
  }

}

  // proto:  QTextCursor QTextFrame::lastCursorPosition();
func (this *QTextFrame) lastCursorPosition(args ...interface{}) () {
  // lastCursorPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextFrame18lastCursorPositionEv
    // invoke: QTextCursor lastCursorPosition()
    C._ZNK10QTextFrame18lastCursorPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "lastCursorPosition", args)
  }

}

  // proto:  int QTextFrame::lastPosition();
func (this *QTextFrame) lastPosition(args ...interface{}) () {
  // lastPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextFrame12lastPositionEv
    // invoke: int lastPosition()
    C._ZNK10QTextFrame12lastPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "lastPosition", args)
  }

}

  // proto:  QTextCursor QTextFrame::firstCursorPosition();
func (this *QTextFrame) firstCursorPosition(args ...interface{}) () {
  // firstCursorPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextFrame19firstCursorPositionEv
    // invoke: QTextCursor firstCursorPosition()
    C._ZNK10QTextFrame19firstCursorPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "firstCursorPosition", args)
  }

}

// <= body block end

