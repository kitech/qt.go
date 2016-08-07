package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QTextFormat QTextObject::format();
extern void* C_ZNK11QTextObject6formatEv(void* qthis); // 4
  // proto:  QTextDocumentPrivate * QTextObject::docHandle();
extern void C_ZNK11QTextObject9docHandleEv(void* qthis); // 4
  // proto:  const QMetaObject * QTextObject::metaObject();
extern void C_ZNK11QTextObject10metaObjectEv(void* qthis); // 4
  // proto:  int QTextObject::formatIndex();
extern int32_t C_ZNK11QTextObject11formatIndexEv(void* qthis); // 4
  // proto:  int QTextObject::objectIndex();
extern int32_t C_ZNK11QTextObject11objectIndexEv(void* qthis); // 4
  // proto:  QTextDocument * QTextObject::document();
extern void* C_ZNK11QTextObject8documentEv(void* qthis); // 4
  // proto:  void QTextBlockUserData::~QTextBlockUserData();
extern void C_ZN18QTextBlockUserDataD2Ev(void* qthis); // 4
  // proto:  bool QTextFragment::isValid();
extern bool C_ZNK13QTextFragment7isValidEv(void* qthis); // 2
  // proto:  int QTextFragment::charFormatIndex();
extern int32_t C_ZNK13QTextFragment15charFormatIndexEv(void* qthis); // 4
  // proto:  QString QTextFragment::text();
extern void* C_ZNK13QTextFragment4textEv(void* qthis); // 4
  // proto:  bool QTextFragment::contains(int position);
extern bool C_ZNK13QTextFragment8containsEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextCharFormat QTextFragment::charFormat();
extern void* C_ZNK13QTextFragment10charFormatEv(void* qthis); // 4
  // proto:  void QTextFragment::QTextFragment(const QTextFragment & o);
extern void* C_ZN13QTextFragmentC2ERKS_(void* arg0); // 1
  // proto:  void QTextFragment::QTextFragment();
extern void* C_ZN13QTextFragmentC2Ev(); // 1
  // proto:  int QTextFragment::length();
extern int32_t C_ZNK13QTextFragment6lengthEv(void* qthis); // 4
  // proto:  int QTextFragment::position();
extern int32_t C_ZNK13QTextFragment8positionEv(void* qthis); // 4
  // proto:  QList<QGlyphRun> QTextFragment::glyphRuns(int from, int length);
extern void C_ZNK13QTextFragment9glyphRunsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTextFrameLayoutData::~QTextFrameLayoutData();
extern void C_ZN20QTextFrameLayoutDataD2Ev(void* qthis); // 4
  // proto:  int QTextBlock::blockFormatIndex();
extern int32_t C_ZNK10QTextBlock16blockFormatIndexEv(void* qthis); // 4
  // proto:  QString QTextBlock::text();
extern void* C_ZNK10QTextBlock4textEv(void* qthis); // 4
  // proto:  int QTextBlock::lineCount();
extern int32_t C_ZNK10QTextBlock9lineCountEv(void* qthis); // 4
  // proto:  void QTextBlock::clearLayout();
extern void C_ZN10QTextBlock11clearLayoutEv(void* qthis); // 4
  // proto:  QTextCharFormat QTextBlock::charFormat();
extern void* C_ZNK10QTextBlock10charFormatEv(void* qthis); // 4
  // proto:  void QTextBlock::setRevision(int rev);
extern void C_ZN10QTextBlock11setRevisionEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextList * QTextBlock::textList();
extern void* C_ZNK10QTextBlock8textListEv(void* qthis); // 4
  // proto:  int QTextBlock::userState();
extern int32_t C_ZNK10QTextBlock9userStateEv(void* qthis); // 4
  // proto:  void QTextBlock::setUserState(int state);
extern void C_ZN10QTextBlock12setUserStateEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextBlock::iterator QTextBlock::end();
extern void C_ZNK10QTextBlock3endEv(void* qthis); // 4
  // proto:  QTextBlockFormat QTextBlock::blockFormat();
extern void* C_ZNK10QTextBlock11blockFormatEv(void* qthis); // 4
  // proto:  QVector<QTextLayout::FormatRange> QTextBlock::textFormats();
extern void C_ZNK10QTextBlock11textFormatsEv(void* qthis); // 4
  // proto:  bool QTextBlock::contains(int position);
extern bool C_ZNK10QTextBlock8containsEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextBlock QTextBlock::next();
extern void* C_ZNK10QTextBlock4nextEv(void* qthis); // 4
  // proto:  QTextDocumentPrivate * QTextBlock::docHandle();
extern void C_ZNK10QTextBlock9docHandleEv(void* qthis); // 2
  // proto:  int QTextBlock::fragmentIndex();
extern int32_t C_ZNK10QTextBlock13fragmentIndexEv(void* qthis); // 2
  // proto:  int QTextBlock::firstLineNumber();
extern int32_t C_ZNK10QTextBlock15firstLineNumberEv(void* qthis); // 4
  // proto:  const QTextDocument * QTextBlock::document();
extern void* C_ZNK10QTextBlock8documentEv(void* qthis); // 4
  // proto:  int QTextBlock::revision();
extern int32_t C_ZNK10QTextBlock8revisionEv(void* qthis); // 4
  // proto:  QTextBlockUserData * QTextBlock::userData();
extern void* C_ZNK10QTextBlock8userDataEv(void* qthis); // 4
  // proto:  QTextBlock::iterator QTextBlock::begin();
extern void C_ZNK10QTextBlock5beginEv(void* qthis); // 4
  // proto:  bool QTextBlock::isValid();
extern bool C_ZNK10QTextBlock7isValidEv(void* qthis); // 4
  // proto:  void QTextBlock::setLineCount(int count);
extern void C_ZN10QTextBlock12setLineCountEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTextBlock::blockNumber();
extern int32_t C_ZNK10QTextBlock11blockNumberEv(void* qthis); // 4
  // proto:  QTextBlock QTextBlock::previous();
extern void* C_ZNK10QTextBlock8previousEv(void* qthis); // 4
  // proto:  QTextLayout * QTextBlock::layout();
extern void* C_ZNK10QTextBlock6layoutEv(void* qthis); // 4
  // proto:  void QTextBlock::setUserData(QTextBlockUserData * data);
extern void C_ZN10QTextBlock11setUserDataEP18QTextBlockUserData(void* qthis, void* arg0); // 4
  // proto:  Qt::LayoutDirection QTextBlock::textDirection();
extern void C_ZNK10QTextBlock13textDirectionEv(void* qthis); // 4
  // proto:  int QTextBlock::charFormatIndex();
extern int32_t C_ZNK10QTextBlock15charFormatIndexEv(void* qthis); // 4
  // proto:  int QTextBlock::length();
extern int32_t C_ZNK10QTextBlock6lengthEv(void* qthis); // 4
  // proto:  void QTextBlock::QTextBlock(const QTextBlock & o);
extern void* C_ZN10QTextBlockC2ERKS_(void* arg0); // 1
  // proto:  void QTextBlock::QTextBlock();
extern void* C_ZN10QTextBlockC2Ev(); // 1
  // proto:  bool QTextBlock::isVisible();
extern bool C_ZNK10QTextBlock9isVisibleEv(void* qthis); // 4
  // proto:  int QTextBlock::position();
extern int32_t C_ZNK10QTextBlock8positionEv(void* qthis); // 4
  // proto:  void QTextBlock::setVisible(bool visible);
extern void C_ZN10QTextBlock10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QTextBlockGroup::metaObject();
extern void C_ZNK15QTextBlockGroup10metaObjectEv(void* qthis); // 4
  // proto:  void QTextFrame::setFrameFormat(const QTextFrameFormat & format);
extern void C_ZN10QTextFrame14setFrameFormatERK16QTextFrameFormat(void* qthis, void* arg0); // 2
  // proto:  QList<QTextFrame *> QTextFrame::childFrames();
extern void C_ZNK10QTextFrame11childFramesEv(void* qthis); // 4
  // proto:  QTextFrame * QTextFrame::parentFrame();
extern void* C_ZNK10QTextFrame11parentFrameEv(void* qthis); // 4
  // proto:  QTextCursor QTextFrame::firstCursorPosition();
extern void* C_ZNK10QTextFrame19firstCursorPositionEv(void* qthis); // 4
  // proto:  QTextCursor QTextFrame::lastCursorPosition();
extern void* C_ZNK10QTextFrame18lastCursorPositionEv(void* qthis); // 4
  // proto:  int QTextFrame::firstPosition();
extern int32_t C_ZNK10QTextFrame13firstPositionEv(void* qthis); // 4
  // proto:  QTextFrame::iterator QTextFrame::begin();
extern void C_ZNK10QTextFrame5beginEv(void* qthis); // 4
  // proto:  QTextFrameLayoutData * QTextFrame::layoutData();
extern void* C_ZNK10QTextFrame10layoutDataEv(void* qthis); // 4
  // proto:  QTextFrame::iterator QTextFrame::end();
extern void C_ZNK10QTextFrame3endEv(void* qthis); // 4
  // proto:  void QTextFrame::~QTextFrame();
extern void C_ZN10QTextFrameD2Ev(void* qthis); // 4
  // proto:  void QTextFrame::QTextFrame(QTextDocument * doc);
extern void* C_ZN10QTextFrameC2EP13QTextDocument(void* arg0); // 3
  // proto:  const QMetaObject * QTextFrame::metaObject();
extern void C_ZNK10QTextFrame10metaObjectEv(void* qthis); // 4
  // proto:  void QTextFrame::setLayoutData(QTextFrameLayoutData * data);
extern void C_ZN10QTextFrame13setLayoutDataEP20QTextFrameLayoutData(void* qthis, void* arg0); // 4
  // proto:  QTextFrameFormat QTextFrame::frameFormat();
extern void* C_ZNK10QTextFrame11frameFormatEv(void* qthis); // 2
  // proto:  int QTextFrame::lastPosition();
extern int32_t C_ZNK10QTextFrame12lastPositionEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTextObject)=1
type QTextObject struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextBlockUserData)=8
type QTextBlockUserData struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextFragment)=16
type QTextFragment struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextFrameLayoutData)=8
type QTextFrameLayoutData struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextBlock)=16
type QTextBlock struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextBlockGroup)=1
type QTextBlockGroup struct {
  /*qbase*/ QTextObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextFrame)=1
type QTextFrame struct {
  /*qbase*/ QTextObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// format()
func (this *QTextObject) Format(args ...interface{}) (ret interface{}) {
  // format()
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
    // invoke: _ZNK11QTextObject6formatEv
    // invoke: QTextFormat format()
    var ret0 = C.C_ZNK11QTextObject6formatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextFormat{}) // "QTextFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextObject", "format", args)
  }

  return
}

// docHandle()
func (this *QTextObject) Dochandle(args ...interface{}) () {
  // docHandle()
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
    // invoke: _ZNK11QTextObject9docHandleEv
    // invoke: QTextDocumentPrivate * docHandle()
    C.C_ZNK11QTextObject9docHandleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "docHandle", args)
  }

  return
}

// metaObject()
func (this *QTextObject) Metaobject(args ...interface{}) () {
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
    // invoke: _ZNK11QTextObject10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QTextObject10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextObject", "metaObject", args)
  }

  return
}

// formatIndex()
func (this *QTextObject) Formatindex(args ...interface{}) (ret interface{}) {
  // formatIndex()
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
    // invoke: _ZNK11QTextObject11formatIndexEv
    // invoke: int formatIndex()
    var ret0 = C.C_ZNK11QTextObject11formatIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextObject", "formatIndex", args)
  }

  return
}

// objectIndex()
func (this *QTextObject) Objectindex(args ...interface{}) (ret interface{}) {
  // objectIndex()
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
    // invoke: _ZNK11QTextObject11objectIndexEv
    // invoke: int objectIndex()
    var ret0 = C.C_ZNK11QTextObject11objectIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextObject", "objectIndex", args)
  }

  return
}

// document()
func (this *QTextObject) Document(args ...interface{}) (ret interface{}) {
  // document()
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
    // invoke: _ZNK11QTextObject8documentEv
    // invoke: QTextDocument * document()
    var ret0 = C.C_ZNK11QTextObject8documentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextObject", "document", args)
  }

  return
}

// ~QTextBlockUserData()
func (this *QTextBlockUserData) Freeqtextblockuserdata(args ...interface{}) () {
  // ~QTextBlockUserData()
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
    // invoke: _ZN18QTextBlockUserDataD0Ev
    // invoke: void ~QTextBlockUserData()
    C.C_ZN18QTextBlockUserDataD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockUserData", "~QTextBlockUserData", args)
  }

  return
}

// isValid()
func (this *QTextFragment) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
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
    // invoke: _ZNK13QTextFragment7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK13QTextFragment7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFragment", "isValid", args)
  }

  return
}

// charFormatIndex()
func (this *QTextFragment) Charformatindex(args ...interface{}) (ret interface{}) {
  // charFormatIndex()
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
    // invoke: _ZNK13QTextFragment15charFormatIndexEv
    // invoke: int charFormatIndex()
    var ret0 = C.C_ZNK13QTextFragment15charFormatIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFragment", "charFormatIndex", args)
  }

  return
}

// text()
func (this *QTextFragment) Text(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK13QTextFragment4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK13QTextFragment4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFragment", "text", args)
  }

  return
}

// contains(int)
func (this *QTextFragment) Contains(args ...interface{}) (ret interface{}) {
  // contains(int)
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
    // invoke: _ZNK13QTextFragment8containsEi
    // invoke: bool contains(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QTextFragment8containsEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFragment", "contains", args)
  }

  return
}

// charFormat()
func (this *QTextFragment) Charformat(args ...interface{}) (ret interface{}) {
  // charFormat()
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
    // invoke: _ZNK13QTextFragment10charFormatEv
    // invoke: QTextCharFormat charFormat()
    var ret0 = C.C_ZNK13QTextFragment10charFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCharFormat{}) // "QTextCharFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFragment", "charFormat", args)
  }

  return
}

// QTextFragment(const class QTextFragment &)
func NewQTextFragment(args ...interface{}) *QTextFragment {
  // QTextFragment(const class QTextFragment &)
  // QTextFragment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFragment{}) // "const QTextFragment &"
  vtys[1] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextFragmentC1ERKS_
    // invoke: void QTextFragment(const class QTextFragment &)
    var arg0 = args[0].(*QTextFragment).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QTextFragmentC2ERKS_(arg0)
    return &QTextFragment{Qclsinst:qthis}
  case 1:
    // invoke: _ZN13QTextFragmentC1Ev
    // invoke: void QTextFragment()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QTextFragmentC2Ev()
    return &QTextFragment{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextFragment", "QTextFragment", args)
  }

  return nil // QTextFragment{Qclsinst:qthis}
}

// length()
func (this *QTextFragment) Length(args ...interface{}) (ret interface{}) {
  // length()
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
    // invoke: _ZNK13QTextFragment6lengthEv
    // invoke: int length()
    var ret0 = C.C_ZNK13QTextFragment6lengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFragment", "length", args)
  }

  return
}

// position()
func (this *QTextFragment) Position(args ...interface{}) (ret interface{}) {
  // position()
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
    // invoke: _ZNK13QTextFragment8positionEv
    // invoke: int position()
    var ret0 = C.C_ZNK13QTextFragment8positionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFragment", "position", args)
  }

  return
}

// glyphRuns(int, int)
func (this *QTextFragment) Glyphruns(args ...interface{}) () {
  // glyphRuns(int, int)
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
    // invoke: _ZNK13QTextFragment9glyphRunsEii
    // invoke: QList<QGlyphRun> glyphRuns(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK13QTextFragment9glyphRunsEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextFragment", "glyphRuns", args)
  }

  return
}

// ~QTextFrameLayoutData()
func (this *QTextFrameLayoutData) Freeqtextframelayoutdata(args ...interface{}) () {
  // ~QTextFrameLayoutData()
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
    // invoke: _ZN20QTextFrameLayoutDataD0Ev
    // invoke: void ~QTextFrameLayoutData()
    C.C_ZN20QTextFrameLayoutDataD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameLayoutData", "~QTextFrameLayoutData", args)
  }

  return
}

// blockFormatIndex()
func (this *QTextBlock) Blockformatindex(args ...interface{}) (ret interface{}) {
  // blockFormatIndex()
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
    // invoke: _ZNK10QTextBlock16blockFormatIndexEv
    // invoke: int blockFormatIndex()
    var ret0 = C.C_ZNK10QTextBlock16blockFormatIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "blockFormatIndex", args)
  }

  return
}

// text()
func (this *QTextBlock) Text(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK10QTextBlock4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK10QTextBlock4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "text", args)
  }

  return
}

// lineCount()
func (this *QTextBlock) Linecount(args ...interface{}) (ret interface{}) {
  // lineCount()
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
    // invoke: _ZNK10QTextBlock9lineCountEv
    // invoke: int lineCount()
    var ret0 = C.C_ZNK10QTextBlock9lineCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "lineCount", args)
  }

  return
}

// clearLayout()
func (this *QTextBlock) Clearlayout(args ...interface{}) () {
  // clearLayout()
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
    // invoke: _ZN10QTextBlock11clearLayoutEv
    // invoke: void clearLayout()
    C.C_ZN10QTextBlock11clearLayoutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "clearLayout", args)
  }

  return
}

// charFormat()
func (this *QTextBlock) Charformat(args ...interface{}) (ret interface{}) {
  // charFormat()
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
    // invoke: _ZNK10QTextBlock10charFormatEv
    // invoke: QTextCharFormat charFormat()
    var ret0 = C.C_ZNK10QTextBlock10charFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCharFormat{}) // "QTextCharFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "charFormat", args)
  }

  return
}

// setRevision(int)
func (this *QTextBlock) Setrevision(args ...interface{}) () {
  // setRevision(int)
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
    // invoke: _ZN10QTextBlock11setRevisionEi
    // invoke: void setRevision(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextBlock11setRevisionEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setRevision", args)
  }

  return
}

// textList()
func (this *QTextBlock) Textlist(args ...interface{}) (ret interface{}) {
  // textList()
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
    // invoke: _ZNK10QTextBlock8textListEv
    // invoke: QTextList * textList()
    var ret0 = C.C_ZNK10QTextBlock8textListEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextList{}) // "QTextList *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "textList", args)
  }

  return
}

// userState()
func (this *QTextBlock) Userstate(args ...interface{}) (ret interface{}) {
  // userState()
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
    // invoke: _ZNK10QTextBlock9userStateEv
    // invoke: int userState()
    var ret0 = C.C_ZNK10QTextBlock9userStateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "userState", args)
  }

  return
}

// setUserState(int)
func (this *QTextBlock) Setuserstate(args ...interface{}) () {
  // setUserState(int)
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
    // invoke: _ZN10QTextBlock12setUserStateEi
    // invoke: void setUserState(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextBlock12setUserStateEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setUserState", args)
  }

  return
}

// end()
func (this *QTextBlock) End(args ...interface{}) () {
  // end()
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
    // invoke: _ZNK10QTextBlock3endEv
    // invoke: QTextBlock::iterator end()
    C.C_ZNK10QTextBlock3endEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "end", args)
  }

  return
}

// blockFormat()
func (this *QTextBlock) Blockformat(args ...interface{}) (ret interface{}) {
  // blockFormat()
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
    // invoke: _ZNK10QTextBlock11blockFormatEv
    // invoke: QTextBlockFormat blockFormat()
    var ret0 = C.C_ZNK10QTextBlock11blockFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlockFormat{}) // "QTextBlockFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "blockFormat", args)
  }

  return
}

// textFormats()
func (this *QTextBlock) Textformats(args ...interface{}) () {
  // textFormats()
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
    // invoke: _ZNK10QTextBlock11textFormatsEv
    // invoke: QVector<QTextLayout::FormatRange> textFormats()
    C.C_ZNK10QTextBlock11textFormatsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "textFormats", args)
  }

  return
}

// contains(int)
func (this *QTextBlock) Contains(args ...interface{}) (ret interface{}) {
  // contains(int)
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
    // invoke: _ZNK10QTextBlock8containsEi
    // invoke: bool contains(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTextBlock8containsEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "contains", args)
  }

  return
}

// next()
func (this *QTextBlock) Next(args ...interface{}) (ret interface{}) {
  // next()
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
    // invoke: _ZNK10QTextBlock4nextEv
    // invoke: QTextBlock next()
    var ret0 = C.C_ZNK10QTextBlock4nextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlock{}) // "QTextBlock"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "next", args)
  }

  return
}

// docHandle()
func (this *QTextBlock) Dochandle(args ...interface{}) () {
  // docHandle()
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
    // invoke: _ZNK10QTextBlock9docHandleEv
    // invoke: QTextDocumentPrivate * docHandle()
    C.C_ZNK10QTextBlock9docHandleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "docHandle", args)
  }

  return
}

// fragmentIndex()
func (this *QTextBlock) Fragmentindex(args ...interface{}) (ret interface{}) {
  // fragmentIndex()
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
    // invoke: _ZNK10QTextBlock13fragmentIndexEv
    // invoke: int fragmentIndex()
    var ret0 = C.C_ZNK10QTextBlock13fragmentIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "fragmentIndex", args)
  }

  return
}

// firstLineNumber()
func (this *QTextBlock) Firstlinenumber(args ...interface{}) (ret interface{}) {
  // firstLineNumber()
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
    // invoke: _ZNK10QTextBlock15firstLineNumberEv
    // invoke: int firstLineNumber()
    var ret0 = C.C_ZNK10QTextBlock15firstLineNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "firstLineNumber", args)
  }

  return
}

// document()
func (this *QTextBlock) Document(args ...interface{}) (ret interface{}) {
  // document()
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
    // invoke: _ZNK10QTextBlock8documentEv
    // invoke: const QTextDocument * document()
    var ret0 = C.C_ZNK10QTextBlock8documentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextDocument{}) // "const QTextDocument *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "document", args)
  }

  return
}

// revision()
func (this *QTextBlock) Revision(args ...interface{}) (ret interface{}) {
  // revision()
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
    // invoke: _ZNK10QTextBlock8revisionEv
    // invoke: int revision()
    var ret0 = C.C_ZNK10QTextBlock8revisionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "revision", args)
  }

  return
}

// userData()
func (this *QTextBlock) Userdata(args ...interface{}) (ret interface{}) {
  // userData()
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
    // invoke: _ZNK10QTextBlock8userDataEv
    // invoke: QTextBlockUserData * userData()
    var ret0 = C.C_ZNK10QTextBlock8userDataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlockUserData{}) // "QTextBlockUserData *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "userData", args)
  }

  return
}

// begin()
func (this *QTextBlock) Begin(args ...interface{}) () {
  // begin()
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
    // invoke: _ZNK10QTextBlock5beginEv
    // invoke: QTextBlock::iterator begin()
    C.C_ZNK10QTextBlock5beginEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "begin", args)
  }

  return
}

// isValid()
func (this *QTextBlock) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
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
    // invoke: _ZNK10QTextBlock7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK10QTextBlock7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "isValid", args)
  }

  return
}

// setLineCount(int)
func (this *QTextBlock) Setlinecount(args ...interface{}) () {
  // setLineCount(int)
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
    // invoke: _ZN10QTextBlock12setLineCountEi
    // invoke: void setLineCount(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextBlock12setLineCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setLineCount", args)
  }

  return
}

// blockNumber()
func (this *QTextBlock) Blocknumber(args ...interface{}) (ret interface{}) {
  // blockNumber()
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
    // invoke: _ZNK10QTextBlock11blockNumberEv
    // invoke: int blockNumber()
    var ret0 = C.C_ZNK10QTextBlock11blockNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "blockNumber", args)
  }

  return
}

// previous()
func (this *QTextBlock) Previous(args ...interface{}) (ret interface{}) {
  // previous()
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
    // invoke: _ZNK10QTextBlock8previousEv
    // invoke: QTextBlock previous()
    var ret0 = C.C_ZNK10QTextBlock8previousEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlock{}) // "QTextBlock"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "previous", args)
  }

  return
}

// layout()
func (this *QTextBlock) Layout(args ...interface{}) (ret interface{}) {
  // layout()
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
    // invoke: _ZNK10QTextBlock6layoutEv
    // invoke: QTextLayout * layout()
    var ret0 = C.C_ZNK10QTextBlock6layoutEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextLayout{}) // "QTextLayout *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "layout", args)
  }

  return
}

// setUserData(class QTextBlockUserData *)
func (this *QTextBlock) Setuserdata(args ...interface{}) () {
  // setUserData(class QTextBlockUserData *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlockUserData{}) // "QTextBlockUserData *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextBlock11setUserDataEP18QTextBlockUserData
    // invoke: void setUserData(class QTextBlockUserData *)
    var arg0 = args[0].(*QTextBlockUserData).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextBlock11setUserDataEP18QTextBlockUserData(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setUserData", args)
  }

  return
}

// textDirection()
func (this *QTextBlock) Textdirection(args ...interface{}) () {
  // textDirection()
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
    // invoke: _ZNK10QTextBlock13textDirectionEv
    // invoke: Qt::LayoutDirection textDirection()
    C.C_ZNK10QTextBlock13textDirectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlock", "textDirection", args)
  }

  return
}

// charFormatIndex()
func (this *QTextBlock) Charformatindex(args ...interface{}) (ret interface{}) {
  // charFormatIndex()
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
    // invoke: _ZNK10QTextBlock15charFormatIndexEv
    // invoke: int charFormatIndex()
    var ret0 = C.C_ZNK10QTextBlock15charFormatIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "charFormatIndex", args)
  }

  return
}

// length()
func (this *QTextBlock) Length(args ...interface{}) (ret interface{}) {
  // length()
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
    // invoke: _ZNK10QTextBlock6lengthEv
    // invoke: int length()
    var ret0 = C.C_ZNK10QTextBlock6lengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "length", args)
  }

  return
}

// QTextBlock(const class QTextBlock &)
func NewQTextBlock(args ...interface{}) *QTextBlock {
  // QTextBlock(const class QTextBlock &)
  // QTextBlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"
  vtys[1] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextBlockC1ERKS_
    // invoke: void QTextBlock(const class QTextBlock &)
    var arg0 = args[0].(*QTextBlock).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTextBlockC2ERKS_(arg0)
    return &QTextBlock{Qclsinst:qthis}
  case 1:
    // invoke: _ZN10QTextBlockC1Ev
    // invoke: void QTextBlock()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTextBlockC2Ev()
    return &QTextBlock{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextBlock", "QTextBlock", args)
  }

  return nil // QTextBlock{Qclsinst:qthis}
}

// isVisible()
func (this *QTextBlock) Isvisible(args ...interface{}) (ret interface{}) {
  // isVisible()
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
    // invoke: _ZNK10QTextBlock9isVisibleEv
    // invoke: bool isVisible()
    var ret0 = C.C_ZNK10QTextBlock9isVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "isVisible", args)
  }

  return
}

// position()
func (this *QTextBlock) Position(args ...interface{}) (ret interface{}) {
  // position()
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
    // invoke: _ZNK10QTextBlock8positionEv
    // invoke: int position()
    var ret0 = C.C_ZNK10QTextBlock8positionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBlock", "position", args)
  }

  return
}

// setVisible(_Bool)
func (this *QTextBlock) Setvisible(args ...interface{}) () {
  // setVisible(_Bool)
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
    // invoke: _ZN10QTextBlock10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextBlock10setVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlock", "setVisible", args)
  }

  return
}

// metaObject()
func (this *QTextBlockGroup) Metaobject(args ...interface{}) () {
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
    // invoke: _ZNK15QTextBlockGroup10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QTextBlockGroup10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockGroup", "metaObject", args)
  }

  return
}

// setFrameFormat(const class QTextFrameFormat &)
func (this *QTextFrame) Setframeformat(args ...interface{}) () {
  // setFrameFormat(const class QTextFrameFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFrameFormat{}) // "const QTextFrameFormat &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextFrame14setFrameFormatERK16QTextFrameFormat
    // invoke: void setFrameFormat(const class QTextFrameFormat &)
    var arg0 = args[0].(*QTextFrameFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextFrame14setFrameFormatERK16QTextFrameFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrame", "setFrameFormat", args)
  }

  return
}

// childFrames()
func (this *QTextFrame) Childframes(args ...interface{}) () {
  // childFrames()
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
    // invoke: _ZNK10QTextFrame11childFramesEv
    // invoke: QList<QTextFrame *> childFrames()
    C.C_ZNK10QTextFrame11childFramesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "childFrames", args)
  }

  return
}

// parentFrame()
func (this *QTextFrame) Parentframe(args ...interface{}) (ret interface{}) {
  // parentFrame()
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
    // invoke: _ZNK10QTextFrame11parentFrameEv
    // invoke: QTextFrame * parentFrame()
    var ret0 = C.C_ZNK10QTextFrame11parentFrameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextFrame{}) // "QTextFrame *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFrame", "parentFrame", args)
  }

  return
}

// firstCursorPosition()
func (this *QTextFrame) Firstcursorposition(args ...interface{}) (ret interface{}) {
  // firstCursorPosition()
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
    // invoke: _ZNK10QTextFrame19firstCursorPositionEv
    // invoke: QTextCursor firstCursorPosition()
    var ret0 = C.C_ZNK10QTextFrame19firstCursorPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCursor{}) // "QTextCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFrame", "firstCursorPosition", args)
  }

  return
}

// lastCursorPosition()
func (this *QTextFrame) Lastcursorposition(args ...interface{}) (ret interface{}) {
  // lastCursorPosition()
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
    // invoke: _ZNK10QTextFrame18lastCursorPositionEv
    // invoke: QTextCursor lastCursorPosition()
    var ret0 = C.C_ZNK10QTextFrame18lastCursorPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCursor{}) // "QTextCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFrame", "lastCursorPosition", args)
  }

  return
}

// firstPosition()
func (this *QTextFrame) Firstposition(args ...interface{}) (ret interface{}) {
  // firstPosition()
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
    // invoke: _ZNK10QTextFrame13firstPositionEv
    // invoke: int firstPosition()
    var ret0 = C.C_ZNK10QTextFrame13firstPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFrame", "firstPosition", args)
  }

  return
}

// begin()
func (this *QTextFrame) Begin(args ...interface{}) () {
  // begin()
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
    // invoke: _ZNK10QTextFrame5beginEv
    // invoke: QTextFrame::iterator begin()
    C.C_ZNK10QTextFrame5beginEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "begin", args)
  }

  return
}

// layoutData()
func (this *QTextFrame) Layoutdata(args ...interface{}) (ret interface{}) {
  // layoutData()
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
    // invoke: _ZNK10QTextFrame10layoutDataEv
    // invoke: QTextFrameLayoutData * layoutData()
    var ret0 = C.C_ZNK10QTextFrame10layoutDataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextFrameLayoutData{}) // "QTextFrameLayoutData *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFrame", "layoutData", args)
  }

  return
}

// end()
func (this *QTextFrame) End(args ...interface{}) () {
  // end()
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
    // invoke: _ZNK10QTextFrame3endEv
    // invoke: QTextFrame::iterator end()
    C.C_ZNK10QTextFrame3endEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "end", args)
  }

  return
}

// ~QTextFrame()
func (this *QTextFrame) Freeqtextframe(args ...interface{}) () {
  // ~QTextFrame()
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
    // invoke: _ZN10QTextFrameD0Ev
    // invoke: void ~QTextFrame()
    C.C_ZN10QTextFrameD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "~QTextFrame", args)
  }

  return
}

// QTextFrame(class QTextDocument *)
func NewQTextFrame(args ...interface{}) *QTextFrame {
  // QTextFrame(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextFrameC1EP13QTextDocument
    // invoke: void QTextFrame(class QTextDocument *)
    var arg0 = args[0].(*QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTextFrameC2EP13QTextDocument(arg0)
    return &QTextFrame{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextFrame", "QTextFrame", args)
  }

  return nil // QTextFrame{Qclsinst:qthis}
}

// metaObject()
func (this *QTextFrame) Metaobject(args ...interface{}) () {
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
    // invoke: _ZNK10QTextFrame10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QTextFrame10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrame", "metaObject", args)
  }

  return
}

// setLayoutData(class QTextFrameLayoutData *)
func (this *QTextFrame) Setlayoutdata(args ...interface{}) () {
  // setLayoutData(class QTextFrameLayoutData *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFrameLayoutData{}) // "QTextFrameLayoutData *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextFrame13setLayoutDataEP20QTextFrameLayoutData
    // invoke: void setLayoutData(class QTextFrameLayoutData *)
    var arg0 = args[0].(*QTextFrameLayoutData).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextFrame13setLayoutDataEP20QTextFrameLayoutData(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrame", "setLayoutData", args)
  }

  return
}

// frameFormat()
func (this *QTextFrame) Frameformat(args ...interface{}) (ret interface{}) {
  // frameFormat()
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
    // invoke: _ZNK10QTextFrame11frameFormatEv
    // invoke: QTextFrameFormat frameFormat()
    var ret0 = C.C_ZNK10QTextFrame11frameFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextFrameFormat{}) // "QTextFrameFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFrame", "frameFormat", args)
  }

  return
}

// lastPosition()
func (this *QTextFrame) Lastposition(args ...interface{}) (ret interface{}) {
  // lastPosition()
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
    // invoke: _ZNK10QTextFrame12lastPositionEv
    // invoke: int lastPosition()
    var ret0 = C.C_ZNK10QTextFrame12lastPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextFrame", "lastPosition", args)
  }

  return
}

// <= body block end

