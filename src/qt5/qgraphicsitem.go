package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qgraphicsitem.h
// dst-file: /src/widgets/qgraphicsitem.go
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
// class sizeof(QGraphicsTextItem)=1
type QGraphicsTextItem struct {
  /*qbase*/ QGraphicsObject;
  qclsinst uint64 /* *mut c_void*/;
//  _linkActivated QGraphicsTextItem_linkActivated_signal;
//  _linkHovered QGraphicsTextItem_linkHovered_signal;
}

// class sizeof(QGraphicsPixmapItem)=1
type QGraphicsPixmapItem struct {
  /*qbase*/ QGraphicsItem;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsRectItem)=1
type QGraphicsRectItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsEllipseItem)=1
type QGraphicsEllipseItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsPolygonItem)=1
type QGraphicsPolygonItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsPathItem)=1
type QGraphicsPathItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsLineItem)=1
type QGraphicsLineItem struct {
  /*qbase*/ QGraphicsItem;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsItemGroup)=1
type QGraphicsItemGroup struct {
  /*qbase*/ QGraphicsItem;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAbstractGraphicsShapeItem)=1
type QAbstractGraphicsShapeItem struct {
  /*qbase*/ QGraphicsItem;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsItem)=1
type QGraphicsItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsObject)=1
type QGraphicsObject struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _childrenChanged QGraphicsObject_childrenChanged_signal;
//  _parentChanged QGraphicsObject_parentChanged_signal;
//  _heightChanged QGraphicsObject_heightChanged_signal;
//  _zChanged QGraphicsObject_zChanged_signal;
//  _visibleChanged QGraphicsObject_visibleChanged_signal;
//  _yChanged QGraphicsObject_yChanged_signal;
//  _widthChanged QGraphicsObject_widthChanged_signal;
//  _opacityChanged QGraphicsObject_opacityChanged_signal;
//  _rotationChanged QGraphicsObject_rotationChanged_signal;
//  _enabledChanged QGraphicsObject_enabledChanged_signal;
//  _xChanged QGraphicsObject_xChanged_signal;
//  _scaleChanged QGraphicsObject_scaleChanged_signal;
}

// class sizeof(QGraphicsSimpleTextItem)=1
type QGraphicsSimpleTextItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QGraphicsTextItem) openExternalLinks(args ...interface{}) () {
  // openExternalLinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem17openExternalLinksEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "openExternalLinks", args)
 }

}


func (this *QGraphicsTextItem) textWidth(args ...interface{}) () {
  // textWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem9textWidthEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "textWidth", args)
 }

}


func (this *QGraphicsTextItem) setTextWidth(args ...interface{}) () {
  // setTextWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem12setTextWidthEd
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setTextWidth", args)
 }

}


func (this *QGraphicsTextItem) setTextCursor(args ...interface{}) () {
  // setTextCursor(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem13setTextCursorERK11QTextCursor
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setTextCursor", args)
 }

}


func (this *QGraphicsTextItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "type", args)
 }

}


func (this *QGraphicsTextItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem4fontEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "font", args)
 }

}


func NewQGraphicsTextItem(args ...interface{})() {
}


func (this *QGraphicsTextItem) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem10metaObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "metaObject", args)
 }

}


func (this *QGraphicsTextItem) setOpenExternalLinks(args ...interface{}) () {
  // setOpenExternalLinks(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem20setOpenExternalLinksEb
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setOpenExternalLinks", args)
 }

}


func (this *QGraphicsTextItem) setTabChangesFocus(args ...interface{}) () {
  // setTabChangesFocus(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem18setTabChangesFocusEb
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setTabChangesFocus", args)
 }

}


func (this *QGraphicsTextItem) toHtml(args ...interface{}) () {
  // toHtml()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem6toHtmlEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "toHtml", args)
 }

}


func (this *QGraphicsTextItem) setDocument(args ...interface{}) () {
  // setDocument(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem11setDocumentEP13QTextDocument
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setDocument", args)
 }

}


func (this *QGraphicsTextItem) setPlainText(args ...interface{}) () {
  // setPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem12setPlainTextERK7QString
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setPlainText", args)
 }

}


func (this *QGraphicsTextItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "paint", args)
 }

}


func (this *QGraphicsTextItem) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem7setFontERK5QFont
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setFont", args)
 }

}


func (this *QGraphicsTextItem) setDefaultTextColor(args ...interface{}) () {
  // setDefaultTextColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem19setDefaultTextColorERK6QColor
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setDefaultTextColor", args)
 }

}


func (this *QGraphicsTextItem) defaultTextColor(args ...interface{}) () {
  // defaultTextColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem16defaultTextColorEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "defaultTextColor", args)
 }

}


func (this *QGraphicsTextItem) FreeQGraphicsTextItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "~QGraphicsTextItem", args)
 }

}


func (this *QGraphicsTextItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem5shapeEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "shape", args)
 }

}


func (this *QGraphicsTextItem) textCursor(args ...interface{}) () {
  // textCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem10textCursorEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "textCursor", args)
 }

}


func (this *QGraphicsTextItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem12boundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "boundingRect", args)
 }

}


func (this *QGraphicsTextItem) toPlainText(args ...interface{}) () {
  // toPlainText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem11toPlainTextEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "toPlainText", args)
 }

}


func (this *QGraphicsTextItem) setHtml(args ...interface{}) () {
  // setHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem7setHtmlERK7QString
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setHtml", args)
 }

}


func (this *QGraphicsTextItem) tabChangesFocus(args ...interface{}) () {
  // tabChangesFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem15tabChangesFocusEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "tabChangesFocus", args)
 }

}


func (this *QGraphicsTextItem) document(args ...interface{}) () {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem8documentEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "document", args)
 }

}


func (this *QGraphicsTextItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem12isObscuredByEPK13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "isObscuredBy", args)
 }

}


func (this *QGraphicsTextItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem10opaqueAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "opaqueArea", args)
 }

}


func (this *QGraphicsTextItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem8containsERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "contains", args)
 }

}


func (this *QGraphicsTextItem) adjustSize(args ...interface{}) () {
  // adjustSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem10adjustSizeEv
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "adjustSize", args)
 }

}


func NewQGraphicsPixmapItem(args ...interface{})() {
}


func (this *QGraphicsPixmapItem) FreeQGraphicsPixmapItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "~QGraphicsPixmapItem", args)
 }

}


func (this *QGraphicsPixmapItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem10opaqueAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "opaqueArea", args)
 }

}


func (this *QGraphicsPixmapItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem12isObscuredByEPK13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "isObscuredBy", args)
 }

}


func (this *QGraphicsPixmapItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "type", args)
 }

}


func (this *QGraphicsPixmapItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem5shapeEv
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "shape", args)
 }

}


func (this *QGraphicsPixmapItem) pixmap(args ...interface{}) () {
  // pixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem6pixmapEv
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "pixmap", args)
 }

}


func (this *QGraphicsPixmapItem) setOffset(args ...interface{}) () {
  // setOffset(qreal, qreal)
  // setOffset(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsPixmapItem9setOffsetEdd
  case 1:
    // invoke: _ZN19QGraphicsPixmapItem9setOffsetERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "setOffset", args)
 }

}


func (this *QGraphicsPixmapItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsPixmapItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "paint", args)
 }

}


func (this *QGraphicsPixmapItem) offset(args ...interface{}) () {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem6offsetEv
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "offset", args)
 }

}


func (this *QGraphicsPixmapItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem12boundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "boundingRect", args)
 }

}


func (this *QGraphicsPixmapItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem8containsERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "contains", args)
 }

}


func (this *QGraphicsPixmapItem) setPixmap(args ...interface{}) () {
  // setPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsPixmapItem9setPixmapERK7QPixmap
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "setPixmap", args)
 }

}


func (this *QGraphicsRectItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem12isObscuredByEPK13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "isObscuredBy", args)
 }

}


func (this *QGraphicsRectItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem12boundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "boundingRect", args)
 }

}


func NewQGraphicsRectItem(args ...interface{})() {
}


func (this *QGraphicsRectItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "type", args)
 }

}


func (this *QGraphicsRectItem) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem4rectEv
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "rect", args)
 }

}


func (this *QGraphicsRectItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem5shapeEv
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "shape", args)
 }

}


func (this *QGraphicsRectItem) FreeQGraphicsRectItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "~QGraphicsRectItem", args)
 }

}


func (this *QGraphicsRectItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem10opaqueAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "opaqueArea", args)
 }

}


func (this *QGraphicsRectItem) setRect(args ...interface{}) () {
  // setRect(const class QRectF &)
  // setRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRectItem7setRectERK6QRectF
  case 1:
    // invoke: _ZN17QGraphicsRectItem7setRectEdddd
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "setRect", args)
 }

}


func (this *QGraphicsRectItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem8containsERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "contains", args)
 }

}


func (this *QGraphicsRectItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "paint", args)
 }

}


func (this *QGraphicsEllipseItem) setStartAngle(args ...interface{}) () {
  // setStartAngle(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItem13setStartAngleEi
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "setStartAngle", args)
 }

}


func NewQGraphicsEllipseItem(args ...interface{})() {
}


func (this *QGraphicsEllipseItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem8containsERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "contains", args)
 }

}


func (this *QGraphicsEllipseItem) setRect(args ...interface{}) () {
  // setRect(const class QRectF &)
  // setRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItem7setRectERK6QRectF
  case 1:
    // invoke: _ZN20QGraphicsEllipseItem7setRectEdddd
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "setRect", args)
 }

}


func (this *QGraphicsEllipseItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "paint", args)
 }

}


func (this *QGraphicsEllipseItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem12isObscuredByEPK13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "isObscuredBy", args)
 }

}


func (this *QGraphicsEllipseItem) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem4rectEv
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "rect", args)
 }

}


func (this *QGraphicsEllipseItem) spanAngle(args ...interface{}) () {
  // spanAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem9spanAngleEv
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "spanAngle", args)
 }

}


func (this *QGraphicsEllipseItem) startAngle(args ...interface{}) () {
  // startAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem10startAngleEv
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "startAngle", args)
 }

}


func (this *QGraphicsEllipseItem) setSpanAngle(args ...interface{}) () {
  // setSpanAngle(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItem12setSpanAngleEi
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "setSpanAngle", args)
 }

}


func (this *QGraphicsEllipseItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "type", args)
 }

}


func (this *QGraphicsEllipseItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem12boundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "boundingRect", args)
 }

}


func (this *QGraphicsEllipseItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem5shapeEv
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "shape", args)
 }

}


func (this *QGraphicsEllipseItem) FreeQGraphicsEllipseItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "~QGraphicsEllipseItem", args)
 }

}


func (this *QGraphicsEllipseItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem10opaqueAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "opaqueArea", args)
 }

}


func (this *QGraphicsPolygonItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem5shapeEv
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "shape", args)
 }

}


func (this *QGraphicsPolygonItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem12isObscuredByEPK13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "isObscuredBy", args)
 }

}


func (this *QGraphicsPolygonItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsPolygonItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "paint", args)
 }

}


func NewQGraphicsPolygonItem(args ...interface{})() {
}


func (this *QGraphicsPolygonItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem12boundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "boundingRect", args)
 }

}


func (this *QGraphicsPolygonItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "type", args)
 }

}


func (this *QGraphicsPolygonItem) FreeQGraphicsPolygonItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "~QGraphicsPolygonItem", args)
 }

}


func (this *QGraphicsPolygonItem) polygon(args ...interface{}) () {
  // polygon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem7polygonEv
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "polygon", args)
 }

}


func (this *QGraphicsPolygonItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem10opaqueAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "opaqueArea", args)
 }

}


func (this *QGraphicsPolygonItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem8containsERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "contains", args)
 }

}


func (this *QGraphicsPolygonItem) setPolygon(args ...interface{}) () {
  // setPolygon(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsPolygonItem10setPolygonERK9QPolygonF
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "setPolygon", args)
 }

}


func (this *QGraphicsPathItem) setPath(args ...interface{}) () {
  // setPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsPathItem7setPathERK12QPainterPath
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "setPath", args)
 }

}


func NewQGraphicsPathItem(args ...interface{})() {
}


func (this *QGraphicsPathItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem8containsERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "contains", args)
 }

}


func (this *QGraphicsPathItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem12boundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "boundingRect", args)
 }

}


func (this *QGraphicsPathItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "type", args)
 }

}


func (this *QGraphicsPathItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem10opaqueAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "opaqueArea", args)
 }

}


func (this *QGraphicsPathItem) path(args ...interface{}) () {
  // path()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem4pathEv
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "path", args)
 }

}


func (this *QGraphicsPathItem) FreeQGraphicsPathItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "~QGraphicsPathItem", args)
 }

}


func (this *QGraphicsPathItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem5shapeEv
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "shape", args)
 }

}


func (this *QGraphicsPathItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem12isObscuredByEPK13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "isObscuredBy", args)
 }

}


func (this *QGraphicsPathItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsPathItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "paint", args)
 }

}


func (this *QGraphicsLineItem) setPen(args ...interface{}) () {
  // setPen(const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "const QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsLineItem6setPenERK4QPen
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "setPen", args)
 }

}


func NewQGraphicsLineItem(args ...interface{})() {
}


func (this *QGraphicsLineItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem12isObscuredByEPK13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "isObscuredBy", args)
 }

}


func (this *QGraphicsLineItem) line(args ...interface{}) () {
  // line()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem4lineEv
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "line", args)
 }

}


func (this *QGraphicsLineItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem10opaqueAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "opaqueArea", args)
 }

}


func (this *QGraphicsLineItem) setLine(args ...interface{}) () {
  // setLine(qreal, qreal, qreal, qreal)
  // setLine(const class QLineF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsLineItem7setLineEdddd
  case 1:
    // invoke: _ZN17QGraphicsLineItem7setLineERK6QLineF
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "setLine", args)
 }

}


func (this *QGraphicsLineItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem12boundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "boundingRect", args)
 }

}


func (this *QGraphicsLineItem) pen(args ...interface{}) () {
  // pen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem3penEv
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "pen", args)
 }

}


func (this *QGraphicsLineItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem5shapeEv
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "shape", args)
 }

}


func (this *QGraphicsLineItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsLineItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "paint", args)
 }

}


func (this *QGraphicsLineItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "type", args)
 }

}


func (this *QGraphicsLineItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem8containsERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "contains", args)
 }

}


func (this *QGraphicsLineItem) FreeQGraphicsLineItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "~QGraphicsLineItem", args)
 }

}


func (this *QGraphicsItemGroup) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsItemGroup12isObscuredByEPK13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "isObscuredBy", args)
 }

}


func (this *QGraphicsItemGroup) FreeQGraphicsItemGroup(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "~QGraphicsItemGroup", args)
 }

}


func NewQGraphicsItemGroup(args ...interface{})() {
}


func (this *QGraphicsItemGroup) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "type", args)
 }

}


func (this *QGraphicsItemGroup) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsItemGroup12boundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "boundingRect", args)
 }

}


func (this *QGraphicsItemGroup) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "paint", args)
 }

}


func (this *QGraphicsItemGroup) removeFromGroup(args ...interface{}) () {
  // removeFromGroup(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsItemGroup15removeFromGroupEP13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "removeFromGroup", args)
 }

}


func (this *QGraphicsItemGroup) addToGroup(args ...interface{}) () {
  // addToGroup(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsItemGroup10addToGroupEP13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "addToGroup", args)
 }

}


func (this *QGraphicsItemGroup) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsItemGroup10opaqueAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "opaqueArea", args)
 }

}


func (this *QAbstractGraphicsShapeItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAbstractGraphicsShapeItem12isObscuredByEPK13QGraphicsItem
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "isObscuredBy", args)
 }

}


func (this *QAbstractGraphicsShapeItem) brush(args ...interface{}) () {
  // brush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAbstractGraphicsShapeItem5brushEv
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "brush", args)
 }

}


func NewQAbstractGraphicsShapeItem(args ...interface{})() {
}


func (this *QAbstractGraphicsShapeItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAbstractGraphicsShapeItem10opaqueAreaEv
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "opaqueArea", args)
 }

}


func (this *QAbstractGraphicsShapeItem) setBrush(args ...interface{}) () {
  // setBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractGraphicsShapeItem8setBrushERK6QBrush
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "setBrush", args)
 }

}


func (this *QAbstractGraphicsShapeItem) setPen(args ...interface{}) () {
  // setPen(const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "const QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractGraphicsShapeItem6setPenERK4QPen
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "setPen", args)
 }

}


func (this *QAbstractGraphicsShapeItem) pen(args ...interface{}) () {
  // pen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAbstractGraphicsShapeItem3penEv
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "pen", args)
 }

}


func (this *QAbstractGraphicsShapeItem) FreeQAbstractGraphicsShapeItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "~QAbstractGraphicsShapeItem", args)
 }

}


func NewQGraphicsItem(args ...interface{})() {
}


func (this *QGraphicsItem) mapFromParent(args ...interface{}) () {
  // mapFromParent(const class QPainterPath &)
  // mapFromParent(const class QRectF &)
  // mapFromParent(qreal, qreal, qreal, qreal)
  // mapFromParent(qreal, qreal)
  // mapFromParent(const class QPointF &)
  // mapFromParent(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem13mapFromParentERK12QPainterPath
  case 1:
    // invoke: _ZNK13QGraphicsItem13mapFromParentERK6QRectF
  case 2:
    // invoke: _ZNK13QGraphicsItem13mapFromParentEdddd
  case 3:
    // invoke: _ZNK13QGraphicsItem13mapFromParentEdd
  case 4:
    // invoke: _ZNK13QGraphicsItem13mapFromParentERK7QPointF
  case 5:
    // invoke: _ZNK13QGraphicsItem13mapFromParentERK9QPolygonF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapFromParent", args)
 }

}


func (this *QGraphicsItem) mapFromItem(args ...interface{}) () {
  // mapFromItem(const class QGraphicsItem *, const class QPointF &)
  // mapFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
  // mapFromItem(const class QGraphicsItem *, const class QRectF &)
  // mapFromItem(const class QGraphicsItem *, qreal, qreal)
  // mapFromItem(const class QGraphicsItem *, const class QPainterPath &)
  // mapFromItem(const class QGraphicsItem *, const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[2][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[4][1] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[5][1] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK7QPointF
  case 1:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_dddd
  case 2:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK6QRectF
  case 3:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_dd
  case 4:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK12QPainterPath
  case 5:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK9QPolygonF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapFromItem", args)
 }

}


func (this *QGraphicsItem) focusItem(args ...interface{}) () {
  // focusItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9focusItemEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "focusItem", args)
 }

}


func (this *QGraphicsItem) parentObject(args ...interface{}) () {
  // parentObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12parentObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "parentObject", args)
 }

}


func (this *QGraphicsItem) setTransformOriginPoint(args ...interface{}) () {
  // setTransformOriginPoint(const class QPointF &)
  // setTransformOriginPoint(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem23setTransformOriginPointERK7QPointF
  case 1:
    // invoke: _ZN13QGraphicsItem23setTransformOriginPointEdd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setTransformOriginPoint", args)
 }

}


func (this *QGraphicsItem) ungrabMouse(args ...interface{}) () {
  // ungrabMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11ungrabMouseEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "ungrabMouse", args)
 }

}


func (this *QGraphicsItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsItem", "type", args)
 }

}


func (this *QGraphicsItem) isSelected(args ...interface{}) () {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10isSelectedEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isSelected", args)
 }

}


func (this *QGraphicsItem) parentWidget(args ...interface{}) () {
  // parentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12parentWidgetEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "parentWidget", args)
 }

}


func (this *QGraphicsItem) resetTransform(args ...interface{}) () {
  // resetTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem14resetTransformEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "resetTransform", args)
 }

}


func (this *QGraphicsItem) boundingRegion(args ...interface{}) () {
  // boundingRegion(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14boundingRegionERK10QTransform
  default:
    qtrt.ErrorResolve("QGraphicsItem", "boundingRegion", args)
 }

}


func (this *QGraphicsItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsItem", "paint", args)
 }

}


func (this *QGraphicsItem) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8isActiveEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isActive", args)
 }

}


func (this *QGraphicsItem) mapToParent(args ...interface{}) () {
  // mapToParent(const class QPolygonF &)
  // mapToParent(const class QPainterPath &)
  // mapToParent(const class QPointF &)
  // mapToParent(qreal, qreal, qreal, qreal)
  // mapToParent(qreal, qreal)
  // mapToParent(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK9QPolygonF
  case 1:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK12QPainterPath
  case 2:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK7QPointF
  case 3:
    // invoke: _ZNK13QGraphicsItem11mapToParentEdddd
  case 4:
    // invoke: _ZNK13QGraphicsItem11mapToParentEdd
  case 5:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapToParent", args)
 }

}


func (this *QGraphicsItem) isWidget(args ...interface{}) () {
  // isWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8isWidgetEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isWidget", args)
 }

}


func (this *QGraphicsItem) setParentItem(args ...interface{}) () {
  // setParentItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem13setParentItemEPS_
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setParentItem", args)
 }

}


func (this *QGraphicsItem) mapToItem(args ...interface{}) () {
  // mapToItem(const class QGraphicsItem *, const class QRectF &)
  // mapToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
  // mapToItem(const class QGraphicsItem *, qreal, qreal)
  // mapToItem(const class QGraphicsItem *, const class QPainterPath &)
  // mapToItem(const class QGraphicsItem *, const class QPointF &)
  // mapToItem(const class QGraphicsItem *, const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[3][1] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[4][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[5][1] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK6QRectF
  case 1:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_dddd
  case 2:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_dd
  case 3:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK12QPainterPath
  case 4:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK7QPointF
  case 5:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK9QPolygonF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapToItem", args)
 }

}


func (this *QGraphicsItem) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6windowEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "window", args)
 }

}


func (this *QGraphicsItem) scenePos(args ...interface{}) () {
  // scenePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8scenePosEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scenePos", args)
 }

}


func (this *QGraphicsItem) handlesChildEvents(args ...interface{}) () {
  // handlesChildEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem18handlesChildEventsEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "handlesChildEvents", args)
 }

}


func (this *QGraphicsItem) setOpacity(args ...interface{}) () {
  // setOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem10setOpacityEd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setOpacity", args)
 }

}


func (this *QGraphicsItem) sceneTransform(args ...interface{}) () {
  // sceneTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14sceneTransformEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "sceneTransform", args)
 }

}


func (this *QGraphicsItem) setZValue(args ...interface{}) () {
  // setZValue(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem9setZValueEd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setZValue", args)
 }

}


func (this *QGraphicsItem) isObscured(args ...interface{}) () {
  // isObscured(qreal, qreal, qreal, qreal)
  // isObscured(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10isObscuredEdddd
  case 1:
    // invoke: _ZNK13QGraphicsItem10isObscuredERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isObscured", args)
 }

}


func (this *QGraphicsItem) installSceneEventFilter(args ...interface{}) () {
  // installSceneEventFilter(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem23installSceneEventFilterEPS_
  default:
    qtrt.ErrorResolve("QGraphicsItem", "installSceneEventFilter", args)
 }

}


func (this *QGraphicsItem) setY(args ...interface{}) () {
  // setY(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem4setYEd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setY", args)
 }

}


func (this *QGraphicsItem) mapRectToItem(args ...interface{}) () {
  // mapRectToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
  // mapRectToItem(const class QGraphicsItem *, const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[1][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem13mapRectToItemEPKS_dddd
  case 1:
    // invoke: _ZNK13QGraphicsItem13mapRectToItemEPKS_RK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectToItem", args)
 }

}


func (this *QGraphicsItem) parentItem(args ...interface{}) () {
  // parentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10parentItemEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "parentItem", args)
 }

}


func (this *QGraphicsItem) clearFocus(args ...interface{}) () {
  // clearFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem10clearFocusEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "clearFocus", args)
 }

}


func (this *QGraphicsItem) isWindow(args ...interface{}) () {
  // isWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8isWindowEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isWindow", args)
 }

}


func (this *QGraphicsItem) transformOriginPoint(args ...interface{}) () {
  // transformOriginPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem20transformOriginPointEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "transformOriginPoint", args)
 }

}


func (this *QGraphicsItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12boundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "boundingRect", args)
 }

}


func (this *QGraphicsItem) childrenBoundingRect(args ...interface{}) () {
  // childrenBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem20childrenBoundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "childrenBoundingRect", args)
 }

}


func (this *QGraphicsItem) mapFromScene(args ...interface{}) () {
  // mapFromScene(const class QRectF &)
  // mapFromScene(qreal, qreal, qreal, qreal)
  // mapFromScene(const class QPolygonF &)
  // mapFromScene(qreal, qreal)
  // mapFromScene(const class QPainterPath &)
  // mapFromScene(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK6QRectF
  case 1:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneEdddd
  case 2:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK9QPolygonF
  case 3:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneEdd
  case 4:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK12QPainterPath
  case 5:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapFromScene", args)
 }

}


func (this *QGraphicsItem) hasCursor(args ...interface{}) () {
  // hasCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9hasCursorEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "hasCursor", args)
 }

}


func (this *QGraphicsItem) setGraphicsEffect(args ...interface{}) () {
  // setGraphicsEffect(class QGraphicsEffect *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsEffect{}) // "QGraphicsEffect *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem17setGraphicsEffectEP15QGraphicsEffect
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setGraphicsEffect", args)
 }

}


func (this *QGraphicsItem) ensureVisible(args ...interface{}) () {
  // ensureVisible(qreal, qreal, qreal, qreal, int, int)
  // ensureVisible(const class QRectF &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem13ensureVisibleEddddii
  case 1:
    // invoke: _ZN13QGraphicsItem13ensureVisibleERK6QRectFii
  default:
    qtrt.ErrorResolve("QGraphicsItem", "ensureVisible", args)
 }

}


func (this *QGraphicsItem) mapRectToParent(args ...interface{}) () {
  // mapRectToParent(const class QRectF &)
  // mapRectToParent(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem15mapRectToParentERK6QRectF
  case 1:
    // invoke: _ZNK13QGraphicsItem15mapRectToParentEdddd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectToParent", args)
 }

}


func (this *QGraphicsItem) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem10setToolTipERK7QString
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setToolTip", args)
 }

}


func (this *QGraphicsItem) rotation(args ...interface{}) () {
  // rotation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8rotationEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "rotation", args)
 }

}


func (this *QGraphicsItem) scene(args ...interface{}) () {
  // scene()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5sceneEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scene", args)
 }

}


func (this *QGraphicsItem) mapRectFromParent(args ...interface{}) () {
  // mapRectFromParent(const class QRectF &)
  // mapRectFromParent(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem17mapRectFromParentERK6QRectF
  case 1:
    // invoke: _ZNK13QGraphicsItem17mapRectFromParentEdddd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectFromParent", args)
 }

}


func (this *QGraphicsItem) setFocusProxy(args ...interface{}) () {
  // setFocusProxy(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem13setFocusProxyEPS_
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setFocusProxy", args)
 }

}


func (this *QGraphicsItem) acceptDrops(args ...interface{}) () {
  // acceptDrops()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11acceptDropsEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "acceptDrops", args)
 }

}


func (this *QGraphicsItem) mapRectFromScene(args ...interface{}) () {
  // mapRectFromScene(const class QRectF &)
  // mapRectFromScene(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem16mapRectFromSceneERK6QRectF
  case 1:
    // invoke: _ZNK13QGraphicsItem16mapRectFromSceneEdddd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectFromScene", args)
 }

}


func (this *QGraphicsItem) focusScopeItem(args ...interface{}) () {
  // focusScopeItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14focusScopeItemEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "focusScopeItem", args)
 }

}


func (this *QGraphicsItem) removeSceneEventFilter(args ...interface{}) () {
  // removeSceneEventFilter(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem22removeSceneEventFilterEPS_
  default:
    qtrt.ErrorResolve("QGraphicsItem", "removeSceneEventFilter", args)
 }

}


func (this *QGraphicsItem) focusProxy(args ...interface{}) () {
  // focusProxy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10focusProxyEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "focusProxy", args)
 }

}


func (this *QGraphicsItem) sceneBoundingRect(args ...interface{}) () {
  // sceneBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem17sceneBoundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "sceneBoundingRect", args)
 }

}


func (this *QGraphicsItem) FreeQGraphicsItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsItem", "~QGraphicsItem", args)
 }

}


func (this *QGraphicsItem) setX(args ...interface{}) () {
  // setX(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem4setXEd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setX", args)
 }

}


func (this *QGraphicsItem) update(args ...interface{}) () {
  // update(qreal, qreal, qreal, qreal)
  // update(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem6updateEdddd
  case 1:
    // invoke: _ZN13QGraphicsItem6updateERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "update", args)
 }

}


func (this *QGraphicsItem) setSelected(args ...interface{}) () {
  // setSelected(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11setSelectedEb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setSelected", args)
 }

}


func (this *QGraphicsItem) stackBefore(args ...interface{}) () {
  // stackBefore(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11stackBeforeEPKS_
  default:
    qtrt.ErrorResolve("QGraphicsItem", "stackBefore", args)
 }

}


func (this *QGraphicsItem) resetMatrix(args ...interface{}) () {
  // resetMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11resetMatrixEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "resetMatrix", args)
 }

}


func (this *QGraphicsItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10opaqueAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "opaqueArea", args)
 }

}


func (this *QGraphicsItem) unsetCursor(args ...interface{}) () {
  // unsetCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11unsetCursorEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "unsetCursor", args)
 }

}


func (this *QGraphicsItem) mapRectToScene(args ...interface{}) () {
  // mapRectToScene(const class QRectF &)
  // mapRectToScene(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14mapRectToSceneERK6QRectF
  case 1:
    // invoke: _ZNK13QGraphicsItem14mapRectToSceneEdddd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectToScene", args)
 }

}


func (this *QGraphicsItem) mapRectFromItem(args ...interface{}) () {
  // mapRectFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
  // mapRectFromItem(const class QGraphicsItem *, const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[1][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem15mapRectFromItemEPKS_dddd
  case 1:
    // invoke: _ZNK13QGraphicsItem15mapRectFromItemEPKS_RK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectFromItem", args)
 }

}


func (this *QGraphicsItem) scale(args ...interface{}) () {
  // scale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5scaleEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scale", args)
 }

}


func (this *QGraphicsItem) setBoundingRegionGranularity(args ...interface{}) () {
  // setBoundingRegionGranularity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem28setBoundingRegionGranularityEd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setBoundingRegionGranularity", args)
 }

}


func (this *QGraphicsItem) setAcceptDrops(args ...interface{}) () {
  // setAcceptDrops(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem14setAcceptDropsEb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setAcceptDrops", args)
 }

}


func (this *QGraphicsItem) ungrabKeyboard(args ...interface{}) () {
  // ungrabKeyboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem14ungrabKeyboardEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "ungrabKeyboard", args)
 }

}


func (this *QGraphicsItem) setEnabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem10setEnabledEb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setEnabled", args)
 }

}


func (this *QGraphicsItem) graphicsEffect(args ...interface{}) () {
  // graphicsEffect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14graphicsEffectEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "graphicsEffect", args)
 }

}


func (this *QGraphicsItem) acceptHoverEvents(args ...interface{}) () {
  // acceptHoverEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem17acceptHoverEventsEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "acceptHoverEvents", args)
 }

}


func (this *QGraphicsItem) topLevelWidget(args ...interface{}) () {
  // topLevelWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14topLevelWidgetEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "topLevelWidget", args)
 }

}


func (this *QGraphicsItem) transformations(args ...interface{}) () {
  // transformations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem15transformationsEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "transformations", args)
 }

}


func (this *QGraphicsItem) mapToScene(args ...interface{}) () {
  // mapToScene(qreal, qreal, qreal, qreal)
  // mapToScene(qreal, qreal)
  // mapToScene(const class QPolygonF &)
  // mapToScene(const class QPainterPath &)
  // mapToScene(const class QPointF &)
  // mapToScene(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10mapToSceneEdddd
  case 1:
    // invoke: _ZNK13QGraphicsItem10mapToSceneEdd
  case 2:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK9QPolygonF
  case 3:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK12QPainterPath
  case 4:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK7QPointF
  case 5:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapToScene", args)
 }

}


func (this *QGraphicsItem) advance(args ...interface{}) () {
  // advance(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem7advanceEi
  default:
    qtrt.ErrorResolve("QGraphicsItem", "advance", args)
 }

}


func (this *QGraphicsItem) sceneMatrix(args ...interface{}) () {
  // sceneMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11sceneMatrixEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "sceneMatrix", args)
 }

}


func (this *QGraphicsItem) setFiltersChildEvents(args ...interface{}) () {
  // setFiltersChildEvents(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem21setFiltersChildEventsEb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setFiltersChildEvents", args)
 }

}


func (this *QGraphicsItem) itemTransform(args ...interface{}) () {
  // itemTransform(const class QGraphicsItem *, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem13itemTransformEPKS_Pb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "itemTransform", args)
 }

}


func (this *QGraphicsItem) moveBy(args ...interface{}) () {
  // moveBy(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem6moveByEdd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "moveBy", args)
 }

}


func (this *QGraphicsItem) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5groupEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "group", args)
 }

}


func (this *QGraphicsItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5shapeEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "shape", args)
 }

}


func (this *QGraphicsItem) scroll(args ...interface{}) () {
  // scroll(qreal, qreal, const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem6scrollEddRK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scroll", args)
 }

}


func (this *QGraphicsItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12isObscuredByEPKS_
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isObscuredBy", args)
 }

}


func (this *QGraphicsItem) setData(args ...interface{}) () {
  // setData(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem7setDataEiRK8QVariant
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setData", args)
 }

}


func (this *QGraphicsItem) commonAncestorItem(args ...interface{}) () {
  // commonAncestorItem(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem18commonAncestorItemEPKS_
  default:
    qtrt.ErrorResolve("QGraphicsItem", "commonAncestorItem", args)
 }

}


func (this *QGraphicsItem) setGroup(args ...interface{}) () {
  // setGroup(class QGraphicsItemGroup *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItemGroup{}) // "QGraphicsItemGroup *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem8setGroupEP18QGraphicsItemGroup
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setGroup", args)
 }

}


func (this *QGraphicsItem) show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem4showEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "show", args)
 }

}


func (this *QGraphicsItem) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem1yEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "y", args)
 }

}


func (this *QGraphicsItem) hasFocus(args ...interface{}) () {
  // hasFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8hasFocusEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "hasFocus", args)
 }

}


func (this *QGraphicsItem) clipPath(args ...interface{}) () {
  // clipPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8clipPathEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "clipPath", args)
 }

}


func (this *QGraphicsItem) setPos(args ...interface{}) () {
  // setPos(qreal, qreal)
  // setPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem6setPosEdd
  case 1:
    // invoke: _ZN13QGraphicsItem6setPosERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setPos", args)
 }

}


func (this *QGraphicsItem) isEnabled(args ...interface{}) () {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9isEnabledEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isEnabled", args)
 }

}


func (this *QGraphicsItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8containsERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsItem", "contains", args)
 }

}


func (this *QGraphicsItem) isPanel(args ...interface{}) () {
  // isPanel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem7isPanelEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isPanel", args)
 }

}


func (this *QGraphicsItem) filtersChildEvents(args ...interface{}) () {
  // filtersChildEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem18filtersChildEventsEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "filtersChildEvents", args)
 }

}


func (this *QGraphicsItem) grabKeyboard(args ...interface{}) () {
  // grabKeyboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem12grabKeyboardEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "grabKeyboard", args)
 }

}


func (this *QGraphicsItem) setActive(args ...interface{}) () {
  // setActive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem9setActiveEb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setActive", args)
 }

}


func (this *QGraphicsItem) toGraphicsObject(args ...interface{}) () {
  // toGraphicsObject()
  // toGraphicsObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem16toGraphicsObjectEv
  case 1:
    // invoke: _ZNK13QGraphicsItem16toGraphicsObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "toGraphicsObject", args)
 }

}


func (this *QGraphicsItem) setHandlesChildEvents(args ...interface{}) () {
  // setHandlesChildEvents(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem21setHandlesChildEventsEb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setHandlesChildEvents", args)
 }

}


func (this *QGraphicsItem) setMatrix(args ...interface{}) () {
  // setMatrix(const class QMatrix &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem9setMatrixERK7QMatrixb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setMatrix", args)
 }

}


func (this *QGraphicsItem) transform(args ...interface{}) () {
  // transform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9transformEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "transform", args)
 }

}


func (this *QGraphicsItem) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem4dataEi
  default:
    qtrt.ErrorResolve("QGraphicsItem", "data", args)
 }

}


func (this *QGraphicsItem) hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem4hideEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "hide", args)
 }

}


func (this *QGraphicsItem) isUnderMouse(args ...interface{}) () {
  // isUnderMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12isUnderMouseEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isUnderMouse", args)
 }

}


func (this *QGraphicsItem) setAcceptTouchEvents(args ...interface{}) () {
  // setAcceptTouchEvents(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem20setAcceptTouchEventsEb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setAcceptTouchEvents", args)
 }

}


func (this *QGraphicsItem) setAcceptHoverEvents(args ...interface{}) () {
  // setAcceptHoverEvents(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem20setAcceptHoverEventsEb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setAcceptHoverEvents", args)
 }

}


func (this *QGraphicsItem) childItems(args ...interface{}) () {
  // childItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10childItemsEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "childItems", args)
 }

}


func (this *QGraphicsItem) isAncestorOf(args ...interface{}) () {
  // isAncestorOf(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12isAncestorOfEPKS_
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isAncestorOf", args)
 }

}


func (this *QGraphicsItem) opacity(args ...interface{}) () {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem7opacityEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "opacity", args)
 }

}


func (this *QGraphicsItem) isVisibleTo(args ...interface{}) () {
  // isVisibleTo(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11isVisibleToEPKS_
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isVisibleTo", args)
 }

}


func (this *QGraphicsItem) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem7toolTipEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "toolTip", args)
 }

}


func (this *QGraphicsItem) cursor(args ...interface{}) () {
  // cursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6cursorEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "cursor", args)
 }

}


func (this *QGraphicsItem) zValue(args ...interface{}) () {
  // zValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6zValueEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "zValue", args)
 }

}


func (this *QGraphicsItem) matrix(args ...interface{}) () {
  // matrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6matrixEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "matrix", args)
 }

}


func (this *QGraphicsItem) panel(args ...interface{}) () {
  // panel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5panelEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "panel", args)
 }

}


func (this *QGraphicsItem) isClipped(args ...interface{}) () {
  // isClipped()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9isClippedEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isClipped", args)
 }

}


func (this *QGraphicsItem) topLevelItem(args ...interface{}) () {
  // topLevelItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12topLevelItemEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "topLevelItem", args)
 }

}


func (this *QGraphicsItem) setScale(args ...interface{}) () {
  // setScale(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem8setScaleEd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setScale", args)
 }

}


func (this *QGraphicsItem) setCursor(args ...interface{}) () {
  // setCursor(const class QCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCursor{}) // "const QCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem9setCursorERK7QCursor
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setCursor", args)
 }

}


func (this *QGraphicsItem) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9isVisibleEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isVisible", args)
 }

}


func (this *QGraphicsItem) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem3posEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "pos", args)
 }

}


func (this *QGraphicsItem) isBlockedByModalPanel(args ...interface{}) () {
  // isBlockedByModalPanel(class QGraphicsItem **)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem **"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isBlockedByModalPanel", args)
 }

}


func (this *QGraphicsItem) effectiveOpacity(args ...interface{}) () {
  // effectiveOpacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem16effectiveOpacityEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "effectiveOpacity", args)
 }

}


func (this *QGraphicsItem) boundingRegionGranularity(args ...interface{}) () {
  // boundingRegionGranularity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem25boundingRegionGranularityEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "boundingRegionGranularity", args)
 }

}


func (this *QGraphicsItem) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem1xEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "x", args)
 }

}


func (this *QGraphicsItem) grabMouse(args ...interface{}) () {
  // grabMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem9grabMouseEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "grabMouse", args)
 }

}


func (this *QGraphicsItem) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem10setVisibleEb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setVisible", args)
 }

}


func (this *QGraphicsItem) setRotation(args ...interface{}) () {
  // setRotation(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11setRotationEd
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setRotation", args)
 }

}


func (this *QGraphicsItem) deviceTransform(args ...interface{}) () {
  // deviceTransform(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem15deviceTransformERK10QTransform
  default:
    qtrt.ErrorResolve("QGraphicsItem", "deviceTransform", args)
 }

}


func (this *QGraphicsItem) acceptTouchEvents(args ...interface{}) () {
  // acceptTouchEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem17acceptTouchEventsEv
  default:
    qtrt.ErrorResolve("QGraphicsItem", "acceptTouchEvents", args)
 }

}


func (this *QGraphicsItem) setTransform(args ...interface{}) () {
  // setTransform(const class QTransform &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem12setTransformERK10QTransformb
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setTransform", args)
 }

}


func NewQGraphicsObject(args ...interface{})() {
}


func (this *QGraphicsObject) FreeQGraphicsObject(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsObject", "~QGraphicsObject", args)
 }

}


func (this *QGraphicsObject) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsObject10metaObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsObject", "metaObject", args)
 }

}


func (this *QGraphicsSimpleTextItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "type", args)
 }

}


func (this *QGraphicsSimpleTextItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem4fontEv
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "font", args)
 }

}


func (this *QGraphicsSimpleTextItem) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSimpleTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "paint", args)
 }

}


func (this *QGraphicsSimpleTextItem) FreeQGraphicsSimpleTextItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "~QGraphicsSimpleTextItem", args)
 }

}


func (this *QGraphicsSimpleTextItem) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSimpleTextItem7setTextERK7QString
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "setText", args)
 }

}


func (this *QGraphicsSimpleTextItem) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem4textEv
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "text", args)
 }

}


func NewQGraphicsSimpleTextItem(args ...interface{})() {
}


func (this *QGraphicsSimpleTextItem) isObscuredBy(args ...interface{}) () {
  // isObscuredBy(const class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem12isObscuredByEPK13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "isObscuredBy", args)
 }

}


func (this *QGraphicsSimpleTextItem) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem5shapeEv
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "shape", args)
 }

}


func (this *QGraphicsSimpleTextItem) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSimpleTextItem7setFontERK5QFont
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "setFont", args)
 }

}


func (this *QGraphicsSimpleTextItem) opaqueArea(args ...interface{}) () {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem10opaqueAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "opaqueArea", args)
 }

}


func (this *QGraphicsSimpleTextItem) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem12boundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "boundingRect", args)
 }

}


func (this *QGraphicsSimpleTextItem) contains(args ...interface{}) () {
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem8containsERK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "contains", args)
 }

}

// <= body block end

