package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qtextobject.h
// dst-file: /src/gui/qtextobject.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
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
  default:
    qtrt.ErrorResolve("QTextObject", "docHandle", args)
 }

}


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


func NewQTextObject(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QTextObject", "format", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextObject", "formatIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextObject", "document", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextObject", "objectIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextObject", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFragment", "charFormatIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFragment", "position", args)
 }

}


func NewQTextFragment(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QTextFragment", "contains", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFragment", "text", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFragment", "glyphRuns", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFragment", "isValid", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFragment", "charFormat", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFragment", "length", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "document", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "previous", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "length", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "userData", args)
 }

}


func NewQTextBlock(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "text", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "lineCount", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "contains", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "blockNumber", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "setRevision", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "setVisible", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "clearLayout", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "docHandle", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "userState", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "charFormatIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "revision", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "position", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "isValid", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "textList", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "layout", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "setUserData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "blockFormatIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "setUserState", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "fragmentIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "isVisible", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "setLineCount", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "next", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "blockFormat", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "firstLineNumber", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBlock", "charFormat", args)
 }

}


func NewQTextBlockGroup(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QTextBlockGroup", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFrame", "frameFormat", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFrame", "layoutData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFrame", "setLayoutData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFrame", "setFrameFormat", args)
 }

}


func NewQTextFrame(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QTextFrame", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFrame", "parentFrame", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFrame", "firstPosition", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFrame", "childFrames", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFrame", "lastCursorPosition", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFrame", "lastPosition", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextFrame", "firstCursorPosition", args)
 }

}

// <= body block end

