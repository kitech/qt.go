package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QTextFormat QTextObject::format();
extern void C_ZNK11QTextObject6formatEv(void* qthis); // 4
  // proto:  QTextDocumentPrivate * QTextObject::docHandle();
extern void C_ZNK11QTextObject9docHandleEv(void* qthis); // 4
  // proto:  const QMetaObject * QTextObject::metaObject();
extern void C_ZNK11QTextObject10metaObjectEv(void* qthis); // 4
  // proto:  int QTextObject::formatIndex();
extern void C_ZNK11QTextObject11formatIndexEv(void* qthis); // 4
  // proto:  int QTextObject::objectIndex();
extern void C_ZNK11QTextObject11objectIndexEv(void* qthis); // 4
  // proto:  QTextDocument * QTextObject::document();
extern void C_ZNK11QTextObject8documentEv(void* qthis); // 4
  // proto:  void QTextBlockUserData::~QTextBlockUserData();
extern void C_ZN18QTextBlockUserDataD2Ev(void* qthis); // 4
  // proto:  bool QTextFragment::isValid();
extern void C_ZNK13QTextFragment7isValidEv(void* qthis); // 2
  // proto:  int QTextFragment::charFormatIndex();
extern void C_ZNK13QTextFragment15charFormatIndexEv(void* qthis); // 4
  // proto:  QString QTextFragment::text();
extern void C_ZNK13QTextFragment4textEv(void* qthis); // 4
  // proto:  bool QTextFragment::contains(int position);
extern void C_ZNK13QTextFragment8containsEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextCharFormat QTextFragment::charFormat();
extern void C_ZNK13QTextFragment10charFormatEv(void* qthis); // 4
  // proto:  void QTextFragment::QTextFragment(const QTextFragment & o);
extern void C_ZN13QTextFragmentC2ERKS_(void* qthis, void* arg0); // 1
  // proto:  void QTextFragment::QTextFragment();
extern void C_ZN13QTextFragmentC2Ev(void* qthis); // 1
  // proto:  int QTextFragment::length();
extern void C_ZNK13QTextFragment6lengthEv(void* qthis); // 4
  // proto:  int QTextFragment::position();
extern void C_ZNK13QTextFragment8positionEv(void* qthis); // 4
  // proto:  QList<QGlyphRun> QTextFragment::glyphRuns(int from, int length);
extern void C_ZNK13QTextFragment9glyphRunsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTextFrameLayoutData::~QTextFrameLayoutData();
extern void C_ZN20QTextFrameLayoutDataD2Ev(void* qthis); // 4
  // proto:  int QTextBlock::blockFormatIndex();
extern void C_ZNK10QTextBlock16blockFormatIndexEv(void* qthis); // 4
  // proto:  QString QTextBlock::text();
extern void C_ZNK10QTextBlock4textEv(void* qthis); // 4
  // proto:  int QTextBlock::lineCount();
extern void C_ZNK10QTextBlock9lineCountEv(void* qthis); // 4
  // proto:  void QTextBlock::clearLayout();
extern void C_ZN10QTextBlock11clearLayoutEv(void* qthis); // 4
  // proto:  QTextCharFormat QTextBlock::charFormat();
extern void C_ZNK10QTextBlock10charFormatEv(void* qthis); // 4
  // proto:  void QTextBlock::setRevision(int rev);
extern void C_ZN10QTextBlock11setRevisionEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextList * QTextBlock::textList();
extern void C_ZNK10QTextBlock8textListEv(void* qthis); // 4
  // proto:  int QTextBlock::userState();
extern void C_ZNK10QTextBlock9userStateEv(void* qthis); // 4
  // proto:  void QTextBlock::setUserState(int state);
extern void C_ZN10QTextBlock12setUserStateEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextBlock::iterator QTextBlock::end();
extern void C_ZNK10QTextBlock3endEv(void* qthis); // 4
  // proto:  QTextBlockFormat QTextBlock::blockFormat();
extern void C_ZNK10QTextBlock11blockFormatEv(void* qthis); // 4
  // proto:  QVector<QTextLayout::FormatRange> QTextBlock::textFormats();
extern void C_ZNK10QTextBlock11textFormatsEv(void* qthis); // 4
  // proto:  bool QTextBlock::contains(int position);
extern void C_ZNK10QTextBlock8containsEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextBlock QTextBlock::next();
extern void C_ZNK10QTextBlock4nextEv(void* qthis); // 4
  // proto:  QTextDocumentPrivate * QTextBlock::docHandle();
extern void C_ZNK10QTextBlock9docHandleEv(void* qthis); // 2
  // proto:  int QTextBlock::fragmentIndex();
extern void C_ZNK10QTextBlock13fragmentIndexEv(void* qthis); // 2
  // proto:  int QTextBlock::firstLineNumber();
extern void C_ZNK10QTextBlock15firstLineNumberEv(void* qthis); // 4
  // proto:  const QTextDocument * QTextBlock::document();
extern void C_ZNK10QTextBlock8documentEv(void* qthis); // 4
  // proto:  int QTextBlock::revision();
extern void C_ZNK10QTextBlock8revisionEv(void* qthis); // 4
  // proto:  QTextBlockUserData * QTextBlock::userData();
extern void C_ZNK10QTextBlock8userDataEv(void* qthis); // 4
  // proto:  QTextBlock::iterator QTextBlock::begin();
extern void C_ZNK10QTextBlock5beginEv(void* qthis); // 4
  // proto:  bool QTextBlock::isValid();
extern void C_ZNK10QTextBlock7isValidEv(void* qthis); // 4
  // proto:  void QTextBlock::setLineCount(int count);
extern void C_ZN10QTextBlock12setLineCountEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTextBlock::blockNumber();
extern void C_ZNK10QTextBlock11blockNumberEv(void* qthis); // 4
  // proto:  QTextBlock QTextBlock::previous();
extern void C_ZNK10QTextBlock8previousEv(void* qthis); // 4
  // proto:  QTextLayout * QTextBlock::layout();
extern void C_ZNK10QTextBlock6layoutEv(void* qthis); // 4
  // proto:  void QTextBlock::setUserData(QTextBlockUserData * data);
extern void C_ZN10QTextBlock11setUserDataEP18QTextBlockUserData(void* qthis, void* arg0); // 4
  // proto:  Qt::LayoutDirection QTextBlock::textDirection();
extern void C_ZNK10QTextBlock13textDirectionEv(void* qthis); // 4
  // proto:  int QTextBlock::charFormatIndex();
extern void C_ZNK10QTextBlock15charFormatIndexEv(void* qthis); // 4
  // proto:  int QTextBlock::length();
extern void C_ZNK10QTextBlock6lengthEv(void* qthis); // 4
  // proto:  void QTextBlock::QTextBlock(const QTextBlock & o);
extern void C_ZN10QTextBlockC2ERKS_(void* qthis, void* arg0); // 1
  // proto:  void QTextBlock::QTextBlock();
extern void C_ZN10QTextBlockC2Ev(void* qthis); // 1
  // proto:  bool QTextBlock::isVisible();
extern void C_ZNK10QTextBlock9isVisibleEv(void* qthis); // 4
  // proto:  int QTextBlock::position();
extern void C_ZNK10QTextBlock8positionEv(void* qthis); // 4
  // proto:  void QTextBlock::setVisible(bool visible);
extern void C_ZN10QTextBlock10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QTextBlockGroup::metaObject();
extern void C_ZNK15QTextBlockGroup10metaObjectEv(void* qthis); // 4
  // proto:  void QTextFrame::setFrameFormat(const QTextFrameFormat & format);
extern void C_ZN10QTextFrame14setFrameFormatERK16QTextFrameFormat(void* qthis, void* arg0); // 2
  // proto:  QList<QTextFrame *> QTextFrame::childFrames();
extern void C_ZNK10QTextFrame11childFramesEv(void* qthis); // 4
  // proto:  QTextFrame * QTextFrame::parentFrame();
extern void C_ZNK10QTextFrame11parentFrameEv(void* qthis); // 4
  // proto:  QTextCursor QTextFrame::firstCursorPosition();
extern void C_ZNK10QTextFrame19firstCursorPositionEv(void* qthis); // 4
  // proto:  QTextCursor QTextFrame::lastCursorPosition();
extern void C_ZNK10QTextFrame18lastCursorPositionEv(void* qthis); // 4
  // proto:  int QTextFrame::firstPosition();
extern void C_ZNK10QTextFrame13firstPositionEv(void* qthis); // 4
  // proto:  QTextFrame::iterator QTextFrame::begin();
extern void C_ZNK10QTextFrame5beginEv(void* qthis); // 4
  // proto:  QTextFrameLayoutData * QTextFrame::layoutData();
extern void C_ZNK10QTextFrame10layoutDataEv(void* qthis); // 4
  // proto:  QTextFrame::iterator QTextFrame::end();
extern void C_ZNK10QTextFrame3endEv(void* qthis); // 4
  // proto:  void QTextFrame::~QTextFrame();
extern void C_ZN10QTextFrameD2Ev(void* qthis); // 4
  // proto:  void QTextFrame::QTextFrame(QTextDocument * doc);
extern void C_ZN10QTextFrameC2EP13QTextDocument(void* qthis, void* arg0); // 3
  // proto:  const QMetaObject * QTextFrame::metaObject();
extern void C_ZNK10QTextFrame10metaObjectEv(void* qthis); // 4
  // proto:  void QTextFrame::setLayoutData(QTextFrameLayoutData * data);
extern void C_ZN10QTextFrame13setLayoutDataEP20QTextFrameLayoutData(void* qthis, void* arg0); // 4
  // proto:  QTextFrameFormat QTextFrame::frameFormat();
extern void C_ZNK10QTextFrame11frameFormatEv(void* qthis); // 2
  // proto:  int QTextFrame::lastPosition();
extern void C_ZNK10QTextFrame12lastPositionEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextBlockUserData)=8
type QTextBlockUserData struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextFragment)=16
type QTextFragment struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextFrameLayoutData)=8
type QTextFrameLayoutData struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextBlock)=16
type QTextBlock struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextBlockGroup)=1
type QTextBlockGroup struct {
  /*qbase*/ QTextObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextFrame)=1
type QTextFrame struct {
  /*qbase*/ QTextObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// format()
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
    C.C_ZNK11QTextObject6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "format", args)
  }

}

// docHandle()
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
    C.C_ZNK11QTextObject9docHandleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "docHandle", args)
  }

}

// metaObject()
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
    C.C_ZNK11QTextObject10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "metaObject", args)
  }

}

// formatIndex()
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
    C.C_ZNK11QTextObject11formatIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "formatIndex", args)
  }

}

// objectIndex()
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
    C.C_ZNK11QTextObject11objectIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "objectIndex", args)
  }

}

// document()
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
    C.C_ZNK11QTextObject8documentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "document", args)
  }

}

// ~QTextBlockUserData()
func (this *QTextBlockUserData) FreeQTextBlockUserData(args ...interface{}) () {
  // ~QTextBlockUserData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QTextBlockUserDataD0Ev
    // invoke: void ~QTextBlockUserData()
    C.C_ZN18QTextBlockUserDataD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockUserData", "~QTextBlockUserData", args)
  }

}

// isValid()
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
    C.C_ZNK13QTextFragment7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "isValid", args)
  }

}

// charFormatIndex()
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
    C.C_ZNK13QTextFragment15charFormatIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "charFormatIndex", args)
  }

}

// text()
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
    C.C_ZNK13QTextFragment4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "text", args)
  }

}

// contains(int)
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
    C.C_ZNK13QTextFragment8containsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFragment", "contains", args)
  }

}

// charFormat()
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
    C.C_ZNK13QTextFragment10charFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "charFormat", args)
  }

}

// QTextFragment(const class QTextFragment &)
func NewQTextFragment(args ...interface{}) QTextFragment {
  // QTextFragment(const class QTextFragment &)
  // QTextFragment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFragment{}) // "const QTextFragment &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextFragmentC1ERKS_
    // invoke: void QTextFragment(const class QTextFragment &)
    var arg0 = args[0].(QTextFragment).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN13QTextFragmentC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN13QTextFragmentC1Ev
    // invoke: void QTextFragment()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN13QTextFragmentC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextFragment", "QTextFragment", args)
  }

  return QTextFragment{}
}

// length()
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
    C.C_ZNK13QTextFragment6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "length", args)
  }

}

// position()
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
    C.C_ZNK13QTextFragment8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFragment", "position", args)
  }

}

// glyphRuns(int, int)
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
    C.C_ZNK13QTextFragment9glyphRunsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextFragment", "glyphRuns", args)
  }

}

// ~QTextFrameLayoutData()
func (this *QTextFrameLayoutData) FreeQTextFrameLayoutData(args ...interface{}) () {
  // ~QTextFrameLayoutData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextFrameLayoutDataD0Ev
    // invoke: void ~QTextFrameLayoutData()
    C.C_ZN20QTextFrameLayoutDataD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameLayoutData", "~QTextFrameLayoutData", args)
  }

}

// blockFormatIndex()
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
    C.C_ZNK10QTextBlock16blockFormatIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "blockFormatIndex", args)
  }

}

// text()
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
    C.C_ZNK10QTextBlock4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "text", args)
  }

}

// lineCount()
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
    C.C_ZNK10QTextBlock9lineCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "lineCount", args)
  }

}

// clearLayout()
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
    C.C_ZN10QTextBlock11clearLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "clearLayout", args)
  }

}

// charFormat()
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
    C.C_ZNK10QTextBlock10charFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "charFormat", args)
  }

}

// setRevision(int)
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
    C.C_ZN10QTextBlock11setRevisionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setRevision", args)
  }

}

// textList()
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
    C.C_ZNK10QTextBlock8textListEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "textList", args)
  }

}

// userState()
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
    C.C_ZNK10QTextBlock9userStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "userState", args)
  }

}

// setUserState(int)
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
    C.C_ZN10QTextBlock12setUserStateEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setUserState", args)
  }

}

// end()
func (this *QTextBlock) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock3endEv
    // invoke: QTextBlock::iterator end()
    C.C_ZNK10QTextBlock3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "end", args)
  }

}

// blockFormat()
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
    C.C_ZNK10QTextBlock11blockFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "blockFormat", args)
  }

}

// textFormats()
func (this *QTextBlock) textFormats(args ...interface{}) () {
  // textFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock11textFormatsEv
    // invoke: QVector<QTextLayout::FormatRange> textFormats()
    C.C_ZNK10QTextBlock11textFormatsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "textFormats", args)
  }

}

// contains(int)
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
    C.C_ZNK10QTextBlock8containsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "contains", args)
  }

}

// next()
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
    C.C_ZNK10QTextBlock4nextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "next", args)
  }

}

// docHandle()
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
    C.C_ZNK10QTextBlock9docHandleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "docHandle", args)
  }

}

// fragmentIndex()
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
    C.C_ZNK10QTextBlock13fragmentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "fragmentIndex", args)
  }

}

// firstLineNumber()
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
    C.C_ZNK10QTextBlock15firstLineNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "firstLineNumber", args)
  }

}

// document()
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
    C.C_ZNK10QTextBlock8documentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "document", args)
  }

}

// revision()
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
    C.C_ZNK10QTextBlock8revisionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "revision", args)
  }

}

// userData()
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
    C.C_ZNK10QTextBlock8userDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "userData", args)
  }

}

// begin()
func (this *QTextBlock) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock5beginEv
    // invoke: QTextBlock::iterator begin()
    C.C_ZNK10QTextBlock5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "begin", args)
  }

}

// isValid()
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
    C.C_ZNK10QTextBlock7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "isValid", args)
  }

}

// setLineCount(int)
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
    C.C_ZN10QTextBlock12setLineCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setLineCount", args)
  }

}

// blockNumber()
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
    C.C_ZNK10QTextBlock11blockNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "blockNumber", args)
  }

}

// previous()
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
    C.C_ZNK10QTextBlock8previousEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "previous", args)
  }

}

// layout()
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
    C.C_ZNK10QTextBlock6layoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "layout", args)
  }

}

// setUserData(class QTextBlockUserData *)
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
    C.C_ZN10QTextBlock11setUserDataEP18QTextBlockUserData(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setUserData", args)
  }

}

// textDirection()
func (this *QTextBlock) textDirection(args ...interface{}) () {
  // textDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextBlock13textDirectionEv
    // invoke: Qt::LayoutDirection textDirection()
    C.C_ZNK10QTextBlock13textDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "textDirection", args)
  }

}

// charFormatIndex()
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
    C.C_ZNK10QTextBlock15charFormatIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "charFormatIndex", args)
  }

}

// length()
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
    C.C_ZNK10QTextBlock6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "length", args)
  }

}

// QTextBlock(const class QTextBlock &)
func NewQTextBlock(args ...interface{}) QTextBlock {
  // QTextBlock(const class QTextBlock &)
  // QTextBlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextBlockC1ERKS_
    // invoke: void QTextBlock(const class QTextBlock &)
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN10QTextBlockC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN10QTextBlockC1Ev
    // invoke: void QTextBlock()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN10QTextBlockC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextBlock", "QTextBlock", args)
  }

  return QTextBlock{}
}

// isVisible()
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
    C.C_ZNK10QTextBlock9isVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "isVisible", args)
  }

}

// position()
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
    C.C_ZNK10QTextBlock8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "position", args)
  }

}

// setVisible(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextBlock10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setVisible", args)
  }

}

// metaObject()
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
    C.C_ZNK15QTextBlockGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockGroup", "metaObject", args)
  }

}

// setFrameFormat(const class QTextFrameFormat &)
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
    C.C_ZN10QTextFrame14setFrameFormatERK16QTextFrameFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrame", "setFrameFormat", args)
  }

}

// childFrames()
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
    C.C_ZNK10QTextFrame11childFramesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "childFrames", args)
  }

}

// parentFrame()
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
    C.C_ZNK10QTextFrame11parentFrameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "parentFrame", args)
  }

}

// firstCursorPosition()
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
    C.C_ZNK10QTextFrame19firstCursorPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "firstCursorPosition", args)
  }

}

// lastCursorPosition()
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
    C.C_ZNK10QTextFrame18lastCursorPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "lastCursorPosition", args)
  }

}

// firstPosition()
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
    C.C_ZNK10QTextFrame13firstPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "firstPosition", args)
  }

}

// begin()
func (this *QTextFrame) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextFrame5beginEv
    // invoke: QTextFrame::iterator begin()
    C.C_ZNK10QTextFrame5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "begin", args)
  }

}

// layoutData()
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
    C.C_ZNK10QTextFrame10layoutDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "layoutData", args)
  }

}

// end()
func (this *QTextFrame) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextFrame3endEv
    // invoke: QTextFrame::iterator end()
    C.C_ZNK10QTextFrame3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "end", args)
  }

}

// ~QTextFrame()
func (this *QTextFrame) FreeQTextFrame(args ...interface{}) () {
  // ~QTextFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextFrameD0Ev
    // invoke: void ~QTextFrame()
    C.C_ZN10QTextFrameD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "~QTextFrame", args)
  }

}

// QTextFrame(class QTextDocument *)
func NewQTextFrame(args ...interface{}) QTextFrame {
  // QTextFrame(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextFrameC1EP13QTextDocument
    // invoke: void QTextFrame(class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN10QTextFrameC2EP13QTextDocument(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTextFrame", "QTextFrame", args)
  }

  return QTextFrame{}
}

// metaObject()
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
    C.C_ZNK10QTextFrame10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "metaObject", args)
  }

}

// setLayoutData(class QTextFrameLayoutData *)
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
    C.C_ZN10QTextFrame13setLayoutDataEP20QTextFrameLayoutData(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrame", "setLayoutData", args)
  }

}

// frameFormat()
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
    C.C_ZNK10QTextFrame11frameFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "frameFormat", args)
  }

}

// lastPosition()
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
    C.C_ZNK10QTextFrame12lastPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "lastPosition", args)
  }

}

// <= body block end

