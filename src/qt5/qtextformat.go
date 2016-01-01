package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qtextformat.h
// dst-file: /src/gui/qtextformat.go
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
// class sizeof(QTextLength)=16
type QTextLength struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextImageFormat)=1
type QTextImageFormat struct {
  /*qbase*/ QTextCharFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextFormat)=1
type QTextFormat struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextBlockFormat)=1
type QTextBlockFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextCharFormat)=1
type QTextCharFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextTableFormat)=1
type QTextTableFormat struct {
  /*qbase*/ QTextFrameFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextTableCellFormat)=1
type QTextTableCellFormat struct {
  /*qbase*/ QTextCharFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextListFormat)=1
type QTextListFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextFrameFormat)=1
type QTextFrameFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTextLength) value(args ...interface{}) () {
  // value(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLength5valueEd
  default:
    qtrt.ErrorResolve("QTextLength", "value", args)
  }

}


func NewQTextLength(args ...interface{}) QTextLength {
  return QTextLength{}
}


func (this *QTextLength) rawValue(args ...interface{}) () {
  // rawValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLength8rawValueEv
  default:
    qtrt.ErrorResolve("QTextLength", "rawValue", args)
  }

}


func NewQTextImageFormat(args ...interface{}) QTextImageFormat {
  return QTextImageFormat{}
}


func (this *QTextImageFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextImageFormat", "isValid", args)
  }

}


func (this *QTextImageFormat) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat5widthEv
  default:
    qtrt.ErrorResolve("QTextImageFormat", "width", args)
  }

}


func (this *QTextImageFormat) setHeight(args ...interface{}) () {
  // setHeight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextImageFormat9setHeightEd
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setHeight", args)
  }

}


func (this *QTextImageFormat) setWidth(args ...interface{}) () {
  // setWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextImageFormat8setWidthEd
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setWidth", args)
  }

}


func (this *QTextImageFormat) setName(args ...interface{}) () {
  // setName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextImageFormat7setNameERK7QString
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setName", args)
  }

}


func (this *QTextImageFormat) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat4nameEv
  default:
    qtrt.ErrorResolve("QTextImageFormat", "name", args)
  }

}


func (this *QTextImageFormat) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat6heightEv
  default:
    qtrt.ErrorResolve("QTextImageFormat", "height", args)
  }

}


func (this *QTextFormat) toBlockFormat(args ...interface{}) () {
  // toBlockFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toBlockFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toBlockFormat", args)
  }

}


func (this *QTextFormat) stringProperty(args ...interface{}) () {
  // stringProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat14stringPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "stringProperty", args)
  }

}


func (this *QTextFormat) lengthVectorProperty(args ...interface{}) () {
  // lengthVectorProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat20lengthVectorPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "lengthVectorProperty", args)
  }

}


func (this *QTextFormat) objectIndex(args ...interface{}) () {
  // objectIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat11objectIndexEv
  default:
    qtrt.ErrorResolve("QTextFormat", "objectIndex", args)
  }

}


func (this *QTextFormat) setObjectIndex(args ...interface{}) () {
  // setObjectIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat14setObjectIndexEi
  default:
    qtrt.ErrorResolve("QTextFormat", "setObjectIndex", args)
  }

}


func (this *QTextFormat) clearForeground(args ...interface{}) () {
  // clearForeground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat15clearForegroundEv
  default:
    qtrt.ErrorResolve("QTextFormat", "clearForeground", args)
  }

}


func (this *QTextFormat) isTableCellFormat(args ...interface{}) () {
  // isTableCellFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat17isTableCellFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isTableCellFormat", args)
  }

}


func (this *QTextFormat) FreeQTextFormat(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextFormat", "~QTextFormat", args)
  }

}


func (this *QTextFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isValid", args)
  }

}


func NewQTextFormat(args ...interface{}) QTextFormat {
  return QTextFormat{}
}


func (this *QTextFormat) lengthProperty(args ...interface{}) () {
  // lengthProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat14lengthPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "lengthProperty", args)
  }

}


func (this *QTextFormat) merge(args ...interface{}) () {
  // merge(const class QTextFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFormat{}) // "const QTextFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat5mergeERKS_
  default:
    qtrt.ErrorResolve("QTextFormat", "merge", args)
  }

}


func (this *QTextFormat) colorProperty(args ...interface{}) () {
  // colorProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13colorPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "colorProperty", args)
  }

}


func (this *QTextFormat) setForeground(args ...interface{}) () {
  // setForeground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat13setForegroundERK6QBrush
  default:
    qtrt.ErrorResolve("QTextFormat", "setForeground", args)
  }

}


func (this *QTextFormat) boolProperty(args ...interface{}) () {
  // boolProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12boolPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "boolProperty", args)
  }

}


func (this *QTextFormat) isListFormat(args ...interface{}) () {
  // isListFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12isListFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isListFormat", args)
  }

}


func (this *QTextFormat) isImageFormat(args ...interface{}) () {
  // isImageFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isImageFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isImageFormat", args)
  }

}


func (this *QTextFormat) clearProperty(args ...interface{}) () {
  // clearProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat13clearPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "clearProperty", args)
  }

}


func (this *QTextFormat) toFrameFormat(args ...interface{}) () {
  // toFrameFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toFrameFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toFrameFormat", args)
  }

}


func (this *QTextFormat) brushProperty(args ...interface{}) () {
  // brushProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13brushPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "brushProperty", args)
  }

}


func (this *QTextFormat) propertyCount(args ...interface{}) () {
  // propertyCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13propertyCountEv
  default:
    qtrt.ErrorResolve("QTextFormat", "propertyCount", args)
  }

}


func (this *QTextFormat) penProperty(args ...interface{}) () {
  // penProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat11penPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "penProperty", args)
  }

}


func (this *QTextFormat) property(args ...interface{}) () {
  // property(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat8propertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "property", args)
  }

}


func (this *QTextFormat) isTableFormat(args ...interface{}) () {
  // isTableFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isTableFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isTableFormat", args)
  }

}


func (this *QTextFormat) setProperty(args ...interface{}) () {
  // setProperty(int, const QVector<class QTextLength> &)
  // setProperty(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  // vtys[0][1] = reflect.TypeOf(QVector<QTextLength>{}) // "const QVector<QTextLength> &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat11setPropertyEiRK7QVectorI11QTextLengthE
  case 1:
    // invoke: _ZN11QTextFormat11setPropertyEiRK8QVariant
  default:
    qtrt.ErrorResolve("QTextFormat", "setProperty", args)
  }

}


func (this *QTextFormat) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextFormat", "type", args)
  }

}


func (this *QTextFormat) isCharFormat(args ...interface{}) () {
  // isCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12isCharFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isCharFormat", args)
  }

}


func (this *QTextFormat) clearBackground(args ...interface{}) () {
  // clearBackground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat15clearBackgroundEv
  default:
    qtrt.ErrorResolve("QTextFormat", "clearBackground", args)
  }

}


func (this *QTextFormat) isBlockFormat(args ...interface{}) () {
  // isBlockFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isBlockFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isBlockFormat", args)
  }

}


func (this *QTextFormat) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10backgroundEv
  default:
    qtrt.ErrorResolve("QTextFormat", "background", args)
  }

}


func (this *QTextFormat) doubleProperty(args ...interface{}) () {
  // doubleProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat14doublePropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "doubleProperty", args)
  }

}


func (this *QTextFormat) swap(args ...interface{}) () {
  // swap(class QTextFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFormat{}) // "QTextFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat4swapERS_
  default:
    qtrt.ErrorResolve("QTextFormat", "swap", args)
  }

}


func (this *QTextFormat) toImageFormat(args ...interface{}) () {
  // toImageFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toImageFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toImageFormat", args)
  }

}


func (this *QTextFormat) hasProperty(args ...interface{}) () {
  // hasProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat11hasPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "hasProperty", args)
  }

}


func (this *QTextFormat) foreground(args ...interface{}) () {
  // foreground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10foregroundEv
  default:
    qtrt.ErrorResolve("QTextFormat", "foreground", args)
  }

}


func (this *QTextFormat) setObjectType(args ...interface{}) () {
  // setObjectType(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat13setObjectTypeEi
  default:
    qtrt.ErrorResolve("QTextFormat", "setObjectType", args)
  }

}


func (this *QTextFormat) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat13setBackgroundERK6QBrush
  default:
    qtrt.ErrorResolve("QTextFormat", "setBackground", args)
  }

}


func (this *QTextFormat) toTableFormat(args ...interface{}) () {
  // toTableFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toTableFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toTableFormat", args)
  }

}


func (this *QTextFormat) isFrameFormat(args ...interface{}) () {
  // isFrameFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isFrameFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isFrameFormat", args)
  }

}


func (this *QTextFormat) intProperty(args ...interface{}) () {
  // intProperty(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat11intPropertyEi
  default:
    qtrt.ErrorResolve("QTextFormat", "intProperty", args)
  }

}


func (this *QTextFormat) toCharFormat(args ...interface{}) () {
  // toCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12toCharFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toCharFormat", args)
  }

}


func (this *QTextFormat) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat7isEmptyEv
  default:
    qtrt.ErrorResolve("QTextFormat", "isEmpty", args)
  }

}


func (this *QTextFormat) toTableCellFormat(args ...interface{}) () {
  // toTableCellFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat17toTableCellFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toTableCellFormat", args)
  }

}


func (this *QTextFormat) objectType(args ...interface{}) () {
  // objectType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10objectTypeEv
  default:
    qtrt.ErrorResolve("QTextFormat", "objectType", args)
  }

}


func (this *QTextFormat) toListFormat(args ...interface{}) () {
  // toListFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12toListFormatEv
  default:
    qtrt.ErrorResolve("QTextFormat", "toListFormat", args)
  }

}


func (this *QTextFormat) properties(args ...interface{}) () {
  // properties()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10propertiesEv
  default:
    qtrt.ErrorResolve("QTextFormat", "properties", args)
  }

}


func (this *QTextBlockFormat) indent(args ...interface{}) () {
  // indent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat6indentEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "indent", args)
  }

}


func (this *QTextBlockFormat) setTextIndent(args ...interface{}) () {
  // setTextIndent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat13setTextIndentEd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setTextIndent", args)
  }

}


func (this *QTextBlockFormat) setNonBreakableLines(args ...interface{}) () {
  // setNonBreakableLines(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat20setNonBreakableLinesEb
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setNonBreakableLines", args)
  }

}


func (this *QTextBlockFormat) setIndent(args ...interface{}) () {
  // setIndent(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat9setIndentEi
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setIndent", args)
  }

}


func (this *QTextBlockFormat) textIndent(args ...interface{}) () {
  // textIndent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat10textIndentEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "textIndent", args)
  }

}


func (this *QTextBlockFormat) lineHeight(args ...interface{}) () {
  // lineHeight()
  // lineHeight(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat10lineHeightEv
  case 1:
    // invoke: _ZNK16QTextBlockFormat10lineHeightEdd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "lineHeight", args)
  }

}


func NewQTextBlockFormat(args ...interface{}) QTextBlockFormat {
  return QTextBlockFormat{}
}


func (this *QTextBlockFormat) setRightMargin(args ...interface{}) () {
  // setRightMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat14setRightMarginEd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setRightMargin", args)
  }

}


func (this *QTextBlockFormat) topMargin(args ...interface{}) () {
  // topMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat9topMarginEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "topMargin", args)
  }

}


func (this *QTextBlockFormat) rightMargin(args ...interface{}) () {
  // rightMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat11rightMarginEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "rightMargin", args)
  }

}


func (this *QTextBlockFormat) bottomMargin(args ...interface{}) () {
  // bottomMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat12bottomMarginEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "bottomMargin", args)
  }

}


func (this *QTextBlockFormat) setTopMargin(args ...interface{}) () {
  // setTopMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat12setTopMarginEd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setTopMargin", args)
  }

}


func (this *QTextBlockFormat) leftMargin(args ...interface{}) () {
  // leftMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat10leftMarginEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "leftMargin", args)
  }

}


func (this *QTextBlockFormat) setLineHeight(args ...interface{}) () {
  // setLineHeight(qreal, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat13setLineHeightEdi
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setLineHeight", args)
  }

}


func (this *QTextBlockFormat) setBottomMargin(args ...interface{}) () {
  // setBottomMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat15setBottomMarginEd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setBottomMargin", args)
  }

}


func (this *QTextBlockFormat) lineHeightType(args ...interface{}) () {
  // lineHeightType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat14lineHeightTypeEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "lineHeightType", args)
  }

}


func (this *QTextBlockFormat) setLeftMargin(args ...interface{}) () {
  // setLeftMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormat13setLeftMarginEd
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setLeftMargin", args)
  }

}


func (this *QTextBlockFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "isValid", args)
  }

}


func (this *QTextBlockFormat) nonBreakableLines(args ...interface{}) () {
  // nonBreakableLines()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat17nonBreakableLinesEv
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "nonBreakableLines", args)
  }

}


func (this *QTextCharFormat) setFontLetterSpacing(args ...interface{}) () {
  // setFontLetterSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat20setFontLetterSpacingEd
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontLetterSpacing", args)
  }

}


func (this *QTextCharFormat) isAnchor(args ...interface{}) () {
  // isAnchor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat8isAnchorEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "isAnchor", args)
  }

}


func (this *QTextCharFormat) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  // setFont(const class QFont &, enum QTextCharFormat::FontPropertiesInheritanceBehavior)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1][1] = qtrt.Int32Ty(false) // "QTextCharFormat::FontPropertiesInheritanceBehavior"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat7setFontERK5QFont
  case 1:
    // invoke: _ZN15QTextCharFormat7setFontERK5QFontNS_33FontPropertiesInheritanceBehaviorE
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFont", args)
  }

}


func (this *QTextCharFormat) fontOverline(args ...interface{}) () {
  // fontOverline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat12fontOverlineEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontOverline", args)
  }

}


func (this *QTextCharFormat) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat4fontEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "font", args)
  }

}


func (this *QTextCharFormat) fontFamily(args ...interface{}) () {
  // fontFamily()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10fontFamilyEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontFamily", args)
  }

}


func (this *QTextCharFormat) fontStrikeOut(args ...interface{}) () {
  // fontStrikeOut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat13fontStrikeOutEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStrikeOut", args)
  }

}


func (this *QTextCharFormat) setFontPointSize(args ...interface{}) () {
  // setFontPointSize(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat16setFontPointSizeEd
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontPointSize", args)
  }

}


func (this *QTextCharFormat) setUnderlineColor(args ...interface{}) () {
  // setUnderlineColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat17setUnderlineColorERK6QColor
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setUnderlineColor", args)
  }

}


func (this *QTextCharFormat) tableCellRowSpan(args ...interface{}) () {
  // tableCellRowSpan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat16tableCellRowSpanEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "tableCellRowSpan", args)
  }

}


func (this *QTextCharFormat) setFontUnderline(args ...interface{}) () {
  // setFontUnderline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat16setFontUnderlineEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontUnderline", args)
  }

}


func (this *QTextCharFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "isValid", args)
  }

}


func (this *QTextCharFormat) fontItalic(args ...interface{}) () {
  // fontItalic()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10fontItalicEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontItalic", args)
  }

}


func (this *QTextCharFormat) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat10setToolTipERK7QString
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setToolTip", args)
  }

}


func (this *QTextCharFormat) setTextOutline(args ...interface{}) () {
  // setTextOutline(const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "const QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat14setTextOutlineERK4QPen
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTextOutline", args)
  }

}


func (this *QTextCharFormat) setTableCellRowSpan(args ...interface{}) () {
  // setTableCellRowSpan(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat19setTableCellRowSpanEi
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTableCellRowSpan", args)
  }

}


func (this *QTextCharFormat) setAnchor(args ...interface{}) () {
  // setAnchor(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat9setAnchorEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchor", args)
  }

}


func (this *QTextCharFormat) fontPointSize(args ...interface{}) () {
  // fontPointSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat13fontPointSizeEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontPointSize", args)
  }

}


func NewQTextCharFormat(args ...interface{}) QTextCharFormat {
  return QTextCharFormat{}
}


func (this *QTextCharFormat) setFontStrikeOut(args ...interface{}) () {
  // setFontStrikeOut(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat16setFontStrikeOutEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontStrikeOut", args)
  }

}


func (this *QTextCharFormat) fontWordSpacing(args ...interface{}) () {
  // fontWordSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat15fontWordSpacingEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontWordSpacing", args)
  }

}


func (this *QTextCharFormat) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat7toolTipEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "toolTip", args)
  }

}


func (this *QTextCharFormat) setAnchorNames(args ...interface{}) () {
  // setAnchorNames(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat14setAnchorNamesERK11QStringList
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorNames", args)
  }

}


func (this *QTextCharFormat) anchorNames(args ...interface{}) () {
  // anchorNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11anchorNamesEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorNames", args)
  }

}


func (this *QTextCharFormat) setFontFixedPitch(args ...interface{}) () {
  // setFontFixedPitch(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat17setFontFixedPitchEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontFixedPitch", args)
  }

}


func (this *QTextCharFormat) setFontItalic(args ...interface{}) () {
  // setFontItalic(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat13setFontItalicEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontItalic", args)
  }

}


func (this *QTextCharFormat) setFontFamily(args ...interface{}) () {
  // setFontFamily(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat13setFontFamilyERK7QString
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontFamily", args)
  }

}


func (this *QTextCharFormat) fontFixedPitch(args ...interface{}) () {
  // fontFixedPitch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat14fontFixedPitchEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontFixedPitch", args)
  }

}


func (this *QTextCharFormat) setAnchorHref(args ...interface{}) () {
  // setAnchorHref(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat13setAnchorHrefERK7QString
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorHref", args)
  }

}


func (this *QTextCharFormat) fontStretch(args ...interface{}) () {
  // fontStretch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11fontStretchEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStretch", args)
  }

}


func (this *QTextCharFormat) setFontKerning(args ...interface{}) () {
  // setFontKerning(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat14setFontKerningEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontKerning", args)
  }

}


func (this *QTextCharFormat) tableCellColumnSpan(args ...interface{}) () {
  // tableCellColumnSpan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat19tableCellColumnSpanEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "tableCellColumnSpan", args)
  }

}


func (this *QTextCharFormat) fontLetterSpacing(args ...interface{}) () {
  // fontLetterSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat17fontLetterSpacingEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontLetterSpacing", args)
  }

}


func (this *QTextCharFormat) anchorHref(args ...interface{}) () {
  // anchorHref()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10anchorHrefEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorHref", args)
  }

}


func (this *QTextCharFormat) anchorName(args ...interface{}) () {
  // anchorName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10anchorNameEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorName", args)
  }

}


func (this *QTextCharFormat) setFontStretch(args ...interface{}) () {
  // setFontStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat14setFontStretchEi
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontStretch", args)
  }

}


func (this *QTextCharFormat) setAnchorName(args ...interface{}) () {
  // setAnchorName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat13setAnchorNameERK7QString
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorName", args)
  }

}


func (this *QTextCharFormat) fontKerning(args ...interface{}) () {
  // fontKerning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11fontKerningEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontKerning", args)
  }

}


func (this *QTextCharFormat) setFontWeight(args ...interface{}) () {
  // setFontWeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat13setFontWeightEi
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontWeight", args)
  }

}


func (this *QTextCharFormat) fontUnderline(args ...interface{}) () {
  // fontUnderline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat13fontUnderlineEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontUnderline", args)
  }

}


func (this *QTextCharFormat) setFontWordSpacing(args ...interface{}) () {
  // setFontWordSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat18setFontWordSpacingEd
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontWordSpacing", args)
  }

}


func (this *QTextCharFormat) underlineColor(args ...interface{}) () {
  // underlineColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat14underlineColorEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "underlineColor", args)
  }

}


func (this *QTextCharFormat) fontWeight(args ...interface{}) () {
  // fontWeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10fontWeightEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontWeight", args)
  }

}


func (this *QTextCharFormat) setFontOverline(args ...interface{}) () {
  // setFontOverline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat15setFontOverlineEb
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontOverline", args)
  }

}


func (this *QTextCharFormat) setTableCellColumnSpan(args ...interface{}) () {
  // setTableCellColumnSpan(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat22setTableCellColumnSpanEi
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTableCellColumnSpan", args)
  }

}


func (this *QTextCharFormat) textOutline(args ...interface{}) () {
  // textOutline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11textOutlineEv
  default:
    qtrt.ErrorResolve("QTextCharFormat", "textOutline", args)
  }

}


func NewQTextTableFormat(args ...interface{}) QTextTableFormat {
  return QTextTableFormat{}
}


func (this *QTextTableFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "isValid", args)
  }

}


func (this *QTextTableFormat) headerRowCount(args ...interface{}) () {
  // headerRowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat14headerRowCountEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "headerRowCount", args)
  }

}


func (this *QTextTableFormat) columns(args ...interface{}) () {
  // columns()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat7columnsEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "columns", args)
  }

}


func (this *QTextTableFormat) columnWidthConstraints(args ...interface{}) () {
  // columnWidthConstraints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat22columnWidthConstraintsEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "columnWidthConstraints", args)
  }

}


func (this *QTextTableFormat) setCellPadding(args ...interface{}) () {
  // setCellPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormat14setCellPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setCellPadding", args)
  }

}


func (this *QTextTableFormat) cellPadding(args ...interface{}) () {
  // cellPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat11cellPaddingEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "cellPadding", args)
  }

}


func (this *QTextTableFormat) setCellSpacing(args ...interface{}) () {
  // setCellSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormat14setCellSpacingEd
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setCellSpacing", args)
  }

}


func (this *QTextTableFormat) setColumns(args ...interface{}) () {
  // setColumns(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormat10setColumnsEi
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setColumns", args)
  }

}


func (this *QTextTableFormat) clearColumnWidthConstraints(args ...interface{}) () {
  // clearColumnWidthConstraints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormat27clearColumnWidthConstraintsEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "clearColumnWidthConstraints", args)
  }

}


func (this *QTextTableFormat) setHeaderRowCount(args ...interface{}) () {
  // setHeaderRowCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormat17setHeaderRowCountEi
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setHeaderRowCount", args)
  }

}


func (this *QTextTableFormat) cellSpacing(args ...interface{}) () {
  // cellSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat11cellSpacingEv
  default:
    qtrt.ErrorResolve("QTextTableFormat", "cellSpacing", args)
  }

}


func NewQTextTableCellFormat(args ...interface{}) QTextTableCellFormat {
  return QTextTableCellFormat{}
}


func (this *QTextTableCellFormat) setLeftPadding(args ...interface{}) () {
  // setLeftPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormat14setLeftPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setLeftPadding", args)
  }

}


func (this *QTextTableCellFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "isValid", args)
  }

}


func (this *QTextTableCellFormat) setTopPadding(args ...interface{}) () {
  // setTopPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormat13setTopPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setTopPadding", args)
  }

}


func (this *QTextTableCellFormat) leftPadding(args ...interface{}) () {
  // leftPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat11leftPaddingEv
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "leftPadding", args)
  }

}


func (this *QTextTableCellFormat) setPadding(args ...interface{}) () {
  // setPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormat10setPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setPadding", args)
  }

}


func (this *QTextTableCellFormat) topPadding(args ...interface{}) () {
  // topPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat10topPaddingEv
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "topPadding", args)
  }

}


func (this *QTextTableCellFormat) rightPadding(args ...interface{}) () {
  // rightPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat12rightPaddingEv
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "rightPadding", args)
  }

}


func (this *QTextTableCellFormat) bottomPadding(args ...interface{}) () {
  // bottomPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat13bottomPaddingEv
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "bottomPadding", args)
  }

}


func (this *QTextTableCellFormat) setRightPadding(args ...interface{}) () {
  // setRightPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormat15setRightPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setRightPadding", args)
  }

}


func (this *QTextTableCellFormat) setBottomPadding(args ...interface{}) () {
  // setBottomPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormat16setBottomPaddingEd
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setBottomPadding", args)
  }

}


func (this *QTextListFormat) indent(args ...interface{}) () {
  // indent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat6indentEv
  default:
    qtrt.ErrorResolve("QTextListFormat", "indent", args)
  }

}


func NewQTextListFormat(args ...interface{}) QTextListFormat {
  return QTextListFormat{}
}


func (this *QTextListFormat) setIndent(args ...interface{}) () {
  // setIndent(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextListFormat9setIndentEi
  default:
    qtrt.ErrorResolve("QTextListFormat", "setIndent", args)
  }

}


func (this *QTextListFormat) numberSuffix(args ...interface{}) () {
  // numberSuffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat12numberSuffixEv
  default:
    qtrt.ErrorResolve("QTextListFormat", "numberSuffix", args)
  }

}


func (this *QTextListFormat) numberPrefix(args ...interface{}) () {
  // numberPrefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat12numberPrefixEv
  default:
    qtrt.ErrorResolve("QTextListFormat", "numberPrefix", args)
  }

}


func (this *QTextListFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextListFormat", "isValid", args)
  }

}


func (this *QTextListFormat) setNumberSuffix(args ...interface{}) () {
  // setNumberSuffix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextListFormat15setNumberSuffixERK7QString
  default:
    qtrt.ErrorResolve("QTextListFormat", "setNumberSuffix", args)
  }

}


func (this *QTextListFormat) setNumberPrefix(args ...interface{}) () {
  // setNumberPrefix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextListFormat15setNumberPrefixERK7QString
  default:
    qtrt.ErrorResolve("QTextListFormat", "setNumberPrefix", args)
  }

}


func (this *QTextFrameFormat) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat7isValidEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "isValid", args)
  }

}


func (this *QTextFrameFormat) setHeight(args ...interface{}) () {
  // setHeight(qreal)
  // setHeight(const class QTextLength &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextLength{}) // "const QTextLength &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat9setHeightEd
  case 1:
    // invoke: _ZN16QTextFrameFormat9setHeightERK11QTextLength
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setHeight", args)
  }

}


func (this *QTextFrameFormat) setBorderBrush(args ...interface{}) () {
  // setBorderBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat14setBorderBrushERK6QBrush
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBorderBrush", args)
  }

}


func (this *QTextFrameFormat) margin(args ...interface{}) () {
  // margin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat6marginEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "margin", args)
  }

}


func (this *QTextFrameFormat) borderBrush(args ...interface{}) () {
  // borderBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat11borderBrushEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "borderBrush", args)
  }

}


func (this *QTextFrameFormat) setRightMargin(args ...interface{}) () {
  // setRightMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat14setRightMarginEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setRightMargin", args)
  }

}


func (this *QTextFrameFormat) setMargin(args ...interface{}) () {
  // setMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat9setMarginEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setMargin", args)
  }

}


func (this *QTextFrameFormat) setBorder(args ...interface{}) () {
  // setBorder(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat9setBorderEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBorder", args)
  }

}


func (this *QTextFrameFormat) setWidth(args ...interface{}) () {
  // setWidth(const class QTextLength &)
  // setWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextLength{}) // "const QTextLength &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat8setWidthERK11QTextLength
  case 1:
    // invoke: _ZN16QTextFrameFormat8setWidthEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setWidth", args)
  }

}


func (this *QTextFrameFormat) bottomMargin(args ...interface{}) () {
  // bottomMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat12bottomMarginEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "bottomMargin", args)
  }

}


func (this *QTextFrameFormat) setBottomMargin(args ...interface{}) () {
  // setBottomMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat15setBottomMarginEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBottomMargin", args)
  }

}


func (this *QTextFrameFormat) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat6heightEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "height", args)
  }

}


func (this *QTextFrameFormat) rightMargin(args ...interface{}) () {
  // rightMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat11rightMarginEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "rightMargin", args)
  }

}


func (this *QTextFrameFormat) setPadding(args ...interface{}) () {
  // setPadding(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat10setPaddingEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setPadding", args)
  }

}


func (this *QTextFrameFormat) setTopMargin(args ...interface{}) () {
  // setTopMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat12setTopMarginEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setTopMargin", args)
  }

}


func (this *QTextFrameFormat) topMargin(args ...interface{}) () {
  // topMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat9topMarginEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "topMargin", args)
  }

}


func (this *QTextFrameFormat) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat5widthEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "width", args)
  }

}


func NewQTextFrameFormat(args ...interface{}) QTextFrameFormat {
  return QTextFrameFormat{}
}


func (this *QTextFrameFormat) padding(args ...interface{}) () {
  // padding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat7paddingEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "padding", args)
  }

}


func (this *QTextFrameFormat) setLeftMargin(args ...interface{}) () {
  // setLeftMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormat13setLeftMarginEd
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setLeftMargin", args)
  }

}


func (this *QTextFrameFormat) border(args ...interface{}) () {
  // border()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat6borderEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "border", args)
  }

}


func (this *QTextFrameFormat) leftMargin(args ...interface{}) () {
  // leftMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat10leftMarginEv
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "leftMargin", args)
  }

}

// <= body block end

