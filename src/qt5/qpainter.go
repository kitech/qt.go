package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qpainter.h
// dst-file: /src/gui/qpainter.go
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
// class sizeof(QPainter)=1
type QPainter struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QPainter) boundingRect(args ...interface{}) () {
  // boundingRect(const class QRectF &, const class QString &, const class QTextOption &)
  // boundingRect(const class QRect &, int, const class QString &)
  // boundingRect(int, int, int, int, int, const class QString &)
  // boundingRect(const class QRectF &, int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QTextOption{}) // "const QTextOption &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[2][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption
  case 1:
    // invoke: _ZN8QPainter12boundingRectERK5QRectiRK7QString
  case 2:
    // invoke: _ZN8QPainter12boundingRectEiiiiiRK7QString
  case 3:
    // invoke: _ZN8QPainter12boundingRectERK6QRectFiRK7QString
  default:
    qtrt.ErrorResolve("QPainter", "boundingRect", args)
  }

}


func (this *QPainter) drawPicture(args ...interface{}) () {
  // drawPicture(const class QPointF &, const class QPicture &)
  // drawPicture(const class QPoint &, const class QPicture &)
  // drawPicture(int, int, const class QPicture &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPicture{}) // "const QPicture &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1][1] = reflect.TypeOf(QPicture{}) // "const QPicture &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QPicture{}) // "const QPicture &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11drawPictureERK7QPointFRK8QPicture
  case 1:
    // invoke: _ZN8QPainter11drawPictureERK6QPointRK8QPicture
  case 2:
    // invoke: _ZN8QPainter11drawPictureEiiRK8QPicture
  default:
    qtrt.ErrorResolve("QPainter", "drawPicture", args)
  }

}


func (this *QPainter) worldMatrix(args ...interface{}) () {
  // worldMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter11worldMatrixEv
  default:
    qtrt.ErrorResolve("QPainter", "worldMatrix", args)
  }

}


func (this *QPainter) drawText(args ...interface{}) () {
  // drawText(const class QPointF &, const class QString &, int, int)
  // drawText(int, int, const class QString &)
  // drawText(const class QRectF &, int, const class QString &, class QRectF *)
  // drawText(const class QRect &, int, const class QString &, class QRect *)
  // drawText(const class QRectF &, const class QString &, const class QTextOption &)
  // drawText(int, int, int, int, int, const class QString &, class QRect *)
  // drawText(const class QPointF &, const class QString &)
  // drawText(const class QPoint &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][3] = reflect.TypeOf(QRectF{}) // "QRectF *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][3] = reflect.TypeOf(QRect{}) // "QRect *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][2] = reflect.TypeOf(QTextOption{}) // "const QTextOption &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[5][3] = qtrt.Int32Ty(false) // "int"
  vtys[5][4] = qtrt.Int32Ty(false) // "int"
  vtys[5][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[5][6] = reflect.TypeOf(QRect{}) // "QRect *"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[6][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[7][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawTextERK7QPointFRK7QStringii
  case 1:
    // invoke: _ZN8QPainter8drawTextEiiRK7QString
  case 2:
    // invoke: _ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_
  case 3:
    // invoke: _ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_
  case 4:
    // invoke: _ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption
  case 5:
    // invoke: _ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect
  case 6:
    // invoke: _ZN8QPainter8drawTextERK7QPointFRK7QString
  case 7:
    // invoke: _ZN8QPainter8drawTextERK6QPointRK7QString
  default:
    qtrt.ErrorResolve("QPainter", "drawText", args)
  }

}


func (this *QPainter) fillRect(args ...interface{}) () {
  // fillRect(const class QRect &, Qt::GlobalColor)
  // fillRect(int, int, int, int, const class QColor &)
  // fillRect(int, int, int, int, Qt::BrushStyle)
  // fillRect(const class QRectF &, Qt::BrushStyle)
  // fillRect(const class QRectF &, Qt::GlobalColor)
  // fillRect(int, int, int, int, Qt::GlobalColor)
  // fillRect(const class QRect &, const class QColor &)
  // fillRect(int, int, int, int, const class QBrush &)
  // fillRect(const class QRectF &, const class QBrush &)
  // fillRect(const class QRect &, const class QBrush &)
  // fillRect(const class QRect &, Qt::BrushStyle)
  // fillRect(const class QRectF &, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::GlobalColor"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = qtrt.Int32Ty(false) // "Qt::BrushStyle"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[3][1] = qtrt.Int32Ty(false) // "Qt::BrushStyle"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[4][1] = qtrt.Int32Ty(false) // "Qt::GlobalColor"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[5][3] = qtrt.Int32Ty(false) // "int"
  vtys[5][4] = qtrt.Int32Ty(false) // "Qt::GlobalColor"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[6][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[7][3] = qtrt.Int32Ty(false) // "int"
  vtys[7][4] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[8][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[9][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[10][1] = qtrt.Int32Ty(false) // "Qt::BrushStyle"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[11][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8fillRectERK5QRectN2Qt11GlobalColorE
  case 1:
    // invoke: _ZN8QPainter8fillRectEiiiiRK6QColor
  case 2:
    // invoke: _ZN8QPainter8fillRectEiiiiN2Qt10BrushStyleE
  case 3:
    // invoke: _ZN8QPainter8fillRectERK6QRectFN2Qt10BrushStyleE
  case 4:
    // invoke: _ZN8QPainter8fillRectERK6QRectFN2Qt11GlobalColorE
  case 5:
    // invoke: _ZN8QPainter8fillRectEiiiiN2Qt11GlobalColorE
  case 6:
    // invoke: _ZN8QPainter8fillRectERK5QRectRK6QColor
  case 7:
    // invoke: _ZN8QPainter8fillRectEiiiiRK6QBrush
  case 8:
    // invoke: _ZN8QPainter8fillRectERK6QRectFRK6QBrush
  case 9:
    // invoke: _ZN8QPainter8fillRectERK5QRectRK6QBrush
  case 10:
    // invoke: _ZN8QPainter8fillRectERK5QRectN2Qt10BrushStyleE
  case 11:
    // invoke: _ZN8QPainter8fillRectERK6QRectFRK6QColor
  default:
    qtrt.ErrorResolve("QPainter", "fillRect", args)
  }

}


func (this *QPainter) matrix(args ...interface{}) () {
  // matrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter6matrixEv
  default:
    qtrt.ErrorResolve("QPainter", "matrix", args)
  }

}


func (this *QPainter) opacity(args ...interface{}) () {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter7opacityEv
  default:
    qtrt.ErrorResolve("QPainter", "opacity", args)
  }

}


func (this *QPainter) drawTiledPixmap(args ...interface{}) () {
  // drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
  // drawTiledPixmap(const class QRect &, const class QPixmap &, const class QPoint &)
  // drawTiledPixmap(int, int, int, int, const class QPixmap &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[0][2] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[1][2] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[2][5] = qtrt.Int32Ty(false) // "int"
  vtys[2][6] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF
  case 1:
    // invoke: _ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint
  case 2:
    // invoke: _ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii
  default:
    qtrt.ErrorResolve("QPainter", "drawTiledPixmap", args)
  }

}


func (this *QPainter) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter13setBackgroundERK6QBrush
  default:
    qtrt.ErrorResolve("QPainter", "setBackground", args)
  }

}


func (this *QPainter) drawChord(args ...interface{}) () {
  // drawChord(const class QRectF &, int, int)
  // drawChord(const class QRect &, int, int)
  // drawChord(int, int, int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[2][5] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawChordERK6QRectFii
  case 1:
    // invoke: _ZN8QPainter9drawChordERK5QRectii
  case 2:
    // invoke: _ZN8QPainter9drawChordEiiiiii
  default:
    qtrt.ErrorResolve("QPainter", "drawChord", args)
  }

}


func (this *QPainter) drawImage(args ...interface{}) () {
  // drawImage(const class QRectF &, const class QImage &)
  // drawImage(const class QPointF &, const class QImage &, const class QRectF &, Qt::ImageConversionFlags)
  // drawImage(const class QRect &, const class QImage &, const class QRect &, Qt::ImageConversionFlags)
  // drawImage(int, int, const class QImage &, int, int, int, int, Qt::ImageConversionFlags)
  // drawImage(const class QPoint &, const class QImage &, const class QRect &, Qt::ImageConversionFlags)
  // drawImage(const class QPoint &, const class QImage &)
  // drawImage(const class QRect &, const class QImage &)
  // drawImage(const class QPointF &, const class QImage &)
  // drawImage(const class QRectF &, const class QImage &, const class QRectF &, Qt::ImageConversionFlags)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[1][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][3] = qtrt.Int64Ty(false) // "Qt::ImageConversionFlags"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[2][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[2][2] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[2][3] = qtrt.Int64Ty(false) // "Qt::ImageConversionFlags"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[3][4] = qtrt.Int32Ty(false) // "int"
  vtys[3][5] = qtrt.Int32Ty(false) // "int"
  vtys[3][6] = qtrt.Int32Ty(false) // "int"
  vtys[3][7] = qtrt.Int64Ty(false) // "Qt::ImageConversionFlags"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[4][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[4][2] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[4][3] = qtrt.Int64Ty(false) // "Qt::ImageConversionFlags"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[5][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[6][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[7][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[8][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[8][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[8][3] = qtrt.Int64Ty(false) // "Qt::ImageConversionFlags"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawImageERK6QRectFRK6QImage
  case 1:
    // invoke: _ZN8QPainter9drawImageERK7QPointFRK6QImageRK6QRectF6QFlagsIN2Qt19ImageConversionFlagEE
  case 2:
    // invoke: _ZN8QPainter9drawImageERK5QRectRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE
  case 3:
    // invoke: _ZN8QPainter9drawImageEiiRK6QImageiiii6QFlagsIN2Qt19ImageConversionFlagEE
  case 4:
    // invoke: _ZN8QPainter9drawImageERK6QPointRK6QImageRK5QRect6QFlagsIN2Qt19ImageConversionFlagEE
  case 5:
    // invoke: _ZN8QPainter9drawImageERK6QPointRK6QImage
  case 6:
    // invoke: _ZN8QPainter9drawImageERK5QRectRK6QImage
  case 7:
    // invoke: _ZN8QPainter9drawImageERK7QPointFRK6QImage
  case 8:
    // invoke: _ZN8QPainter9drawImageERK6QRectFRK6QImageS2_6QFlagsIN2Qt19ImageConversionFlagEE
  default:
    qtrt.ErrorResolve("QPainter", "drawImage", args)
  }

}


func (this *QPainter) setClipping(args ...interface{}) () {
  // setClipping(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11setClippingEb
  default:
    qtrt.ErrorResolve("QPainter", "setClipping", args)
  }

}


func (this *QPainter) setBrush(args ...interface{}) () {
  // setBrush(const class QBrush &)
  // setBrush(Qt::BrushStyle)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::BrushStyle"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8setBrushERK6QBrush
  case 1:
    // invoke: _ZN8QPainter8setBrushEN2Qt10BrushStyleE
  default:
    qtrt.ErrorResolve("QPainter", "setBrush", args)
  }

}


func (this *QPainter) setMatrix(args ...interface{}) () {
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
    // invoke: _ZN8QPainter9setMatrixERK7QMatrixb
  default:
    qtrt.ErrorResolve("QPainter", "setMatrix", args)
  }

}


func (this *QPainter) eraseRect(args ...interface{}) () {
  // eraseRect(const class QRectF &)
  // eraseRect(int, int, int, int)
  // eraseRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9eraseRectERK6QRectF
  case 1:
    // invoke: _ZN8QPainter9eraseRectEiiii
  case 2:
    // invoke: _ZN8QPainter9eraseRectERK5QRect
  default:
    qtrt.ErrorResolve("QPainter", "eraseRect", args)
  }

}


func (this *QPainter) translate(args ...interface{}) () {
  // translate(const class QPoint &)
  // translate(qreal, qreal)
  // translate(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9translateERK6QPoint
  case 1:
    // invoke: _ZN8QPainter9translateEdd
  case 2:
    // invoke: _ZN8QPainter9translateERK7QPointF
  default:
    qtrt.ErrorResolve("QPainter", "translate", args)
  }

}


func (this *QPainter) viewTransformEnabled(args ...interface{}) () {
  // viewTransformEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter20viewTransformEnabledEv
  default:
    qtrt.ErrorResolve("QPainter", "viewTransformEnabled", args)
  }

}


func (this *QPainter) setPen(args ...interface{}) () {
  // setPen(const class QPen &)
  // setPen(Qt::PenStyle)
  // setPen(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::PenStyle"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter6setPenERK4QPen
  case 1:
    // invoke: _ZN8QPainter6setPenEN2Qt8PenStyleE
  case 2:
    // invoke: _ZN8QPainter6setPenERK6QColor
  default:
    qtrt.ErrorResolve("QPainter", "setPen", args)
  }

}


func (this *QPainter) drawLines(args ...interface{}) () {
  // drawLines(const class QLineF *, int)
  // drawLines(const class QPointF *, int)
  // drawLines(const class QLine *, int)
  // drawLines(const QVector<class QLine> &)
  // drawLines(const class QPoint *, int)
  // drawLines(const QVector<class QLineF> &)
  // drawLines(const QVector<class QPoint> &)
  // drawLines(const QVector<class QPointF> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLine{}) // "const QLine *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  // vtys[3][0] = reflect.TypeOf(QVector<QLine>{}) // "const QVector<QLine> &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPoint{}) // "const QPoint *"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  // vtys[5][0] = reflect.TypeOf(QVector<QLineF>{}) // "const QVector<QLineF> &"
  vtys[6] = make(map[int32]reflect.Type)
  // vtys[6][0] = reflect.TypeOf(QVector<QPoint>{}) // "const QVector<QPoint> &"
  vtys[7] = make(map[int32]reflect.Type)
  // vtys[7][0] = reflect.TypeOf(QVector<QPointF>{}) // "const QVector<QPointF> &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawLinesEPK6QLineFi
  case 1:
    // invoke: _ZN8QPainter9drawLinesEPK7QPointFi
  case 2:
    // invoke: _ZN8QPainter9drawLinesEPK5QLinei
  case 3:
    // invoke: _ZN8QPainter9drawLinesERK7QVectorI5QLineE
  case 4:
    // invoke: _ZN8QPainter9drawLinesEPK6QPointi
  case 5:
    // invoke: _ZN8QPainter9drawLinesERK7QVectorI6QLineFE
  case 6:
    // invoke: _ZN8QPainter9drawLinesERK7QVectorI6QPointE
  case 7:
    // invoke: _ZN8QPainter9drawLinesERK7QVectorI7QPointFE
  default:
    qtrt.ErrorResolve("QPainter", "drawLines", args)
  }

}


func (this *QPainter) setBrushOrigin(args ...interface{}) () {
  // setBrushOrigin(int, int)
  // setBrushOrigin(const class QPointF &)
  // setBrushOrigin(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter14setBrushOriginEii
  case 1:
    // invoke: _ZN8QPainter14setBrushOriginERK7QPointF
  case 2:
    // invoke: _ZN8QPainter14setBrushOriginERK6QPoint
  default:
    qtrt.ErrorResolve("QPainter", "setBrushOrigin", args)
  }

}


func (this *QPainter) worldTransform(args ...interface{}) () {
  // worldTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter14worldTransformEv
  default:
    qtrt.ErrorResolve("QPainter", "worldTransform", args)
  }

}


func (this *QPainter) drawRects(args ...interface{}) () {
  // drawRects(const QVector<class QRect> &)
  // drawRects(const class QRect *, int)
  // drawRects(const class QRectF *, int)
  // drawRects(const QVector<class QRectF> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  // vtys[0][0] = reflect.TypeOf(QVector<QRect>{}) // "const QVector<QRect> &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRectF{}) // "const QRectF *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  // vtys[3][0] = reflect.TypeOf(QVector<QRectF>{}) // "const QVector<QRectF> &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawRectsERK7QVectorI5QRectE
  case 1:
    // invoke: _ZN8QPainter9drawRectsEPK5QRecti
  case 2:
    // invoke: _ZN8QPainter9drawRectsEPK6QRectFi
  case 3:
    // invoke: _ZN8QPainter9drawRectsERK7QVectorI6QRectFE
  default:
    qtrt.ErrorResolve("QPainter", "drawRects", args)
  }

}


func (this *QPainter) drawEllipse(args ...interface{}) () {
  // drawEllipse(const class QPoint &, int, int)
  // drawEllipse(const class QRectF &)
  // drawEllipse(const class QRect &)
  // drawEllipse(const class QPointF &, qreal, qreal)
  // drawEllipse(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[4][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11drawEllipseERK6QPointii
  case 1:
    // invoke: _ZN8QPainter11drawEllipseERK6QRectF
  case 2:
    // invoke: _ZN8QPainter11drawEllipseERK5QRect
  case 3:
    // invoke: _ZN8QPainter11drawEllipseERK7QPointFdd
  case 4:
    // invoke: _ZN8QPainter11drawEllipseEiiii
  default:
    qtrt.ErrorResolve("QPainter", "drawEllipse", args)
  }

}


func (this *QPainter) drawArc(args ...interface{}) () {
  // drawArc(int, int, int, int, int, int)
  // drawArc(const class QRectF &, int, int)
  // drawArc(const class QRect &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter7drawArcEiiiiii
  case 1:
    // invoke: _ZN8QPainter7drawArcERK6QRectFii
  case 2:
    // invoke: _ZN8QPainter7drawArcERK5QRectii
  default:
    qtrt.ErrorResolve("QPainter", "drawArc", args)
  }

}


func (this *QPainter) drawPolyline(args ...interface{}) () {
  // drawPolyline(const class QPolygonF &)
  // drawPolyline(const class QPolygon &)
  // drawPolyline(const class QPoint *, int)
  // drawPolyline(const class QPointF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPoint{}) // "const QPoint *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12drawPolylineERK9QPolygonF
  case 1:
    // invoke: _ZN8QPainter12drawPolylineERK8QPolygon
  case 2:
    // invoke: _ZN8QPainter12drawPolylineEPK6QPointi
  case 3:
    // invoke: _ZN8QPainter12drawPolylineEPK7QPointFi
  default:
    qtrt.ErrorResolve("QPainter", "drawPolyline", args)
  }

}


func (this *QPainter) hasClipping(args ...interface{}) () {
  // hasClipping()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter11hasClippingEv
  default:
    qtrt.ErrorResolve("QPainter", "hasClipping", args)
  }

}


func (this *QPainter) drawPixmap(args ...interface{}) () {
  // drawPixmap(const class QRectF &, const class QPixmap &, const class QRectF &)
  // drawPixmap(int, int, const class QPixmap &, int, int, int, int)
  // drawPixmap(int, int, const class QPixmap &)
  // drawPixmap(int, int, int, int, const class QPixmap &)
  // drawPixmap(const class QPointF &, const class QPixmap &)
  // drawPixmap(const class QPoint &, const class QPixmap &, const class QRect &)
  // drawPixmap(const class QRect &, const class QPixmap &)
  // drawPixmap(int, int, int, int, const class QPixmap &, int, int, int, int)
  // drawPixmap(const class QPointF &, const class QPixmap &, const class QRectF &)
  // drawPixmap(const class QPoint &, const class QPixmap &)
  // drawPixmap(const class QRect &, const class QPixmap &, const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[0][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[1][5] = qtrt.Int32Ty(false) // "int"
  vtys[1][6] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[3][4] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[4][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[5][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[5][2] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[6][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[7][3] = qtrt.Int32Ty(false) // "int"
  vtys[7][4] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[7][5] = qtrt.Int32Ty(false) // "int"
  vtys[7][6] = qtrt.Int32Ty(false) // "int"
  vtys[7][7] = qtrt.Int32Ty(false) // "int"
  vtys[7][8] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[8][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[8][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[9][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[10][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[10][2] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10drawPixmapERK6QRectFRK7QPixmapS2_
  case 1:
    // invoke: _ZN8QPainter10drawPixmapEiiRK7QPixmapiiii
  case 2:
    // invoke: _ZN8QPainter10drawPixmapEiiRK7QPixmap
  case 3:
    // invoke: _ZN8QPainter10drawPixmapEiiiiRK7QPixmap
  case 4:
    // invoke: _ZN8QPainter10drawPixmapERK7QPointFRK7QPixmap
  case 5:
    // invoke: _ZN8QPainter10drawPixmapERK6QPointRK7QPixmapRK5QRect
  case 6:
    // invoke: _ZN8QPainter10drawPixmapERK5QRectRK7QPixmap
  case 7:
    // invoke: _ZN8QPainter10drawPixmapEiiiiRK7QPixmapiiii
  case 8:
    // invoke: _ZN8QPainter10drawPixmapERK7QPointFRK7QPixmapRK6QRectF
  case 9:
    // invoke: _ZN8QPainter10drawPixmapERK6QPointRK7QPixmap
  case 10:
    // invoke: _ZN8QPainter10drawPixmapERK5QRectRK7QPixmapS2_
  default:
    qtrt.ErrorResolve("QPainter", "drawPixmap", args)
  }

}


func (this *QPainter) drawStaticText(args ...interface{}) () {
  // drawStaticText(int, int, const class QStaticText &)
  // drawStaticText(const class QPointF &, const class QStaticText &)
  // drawStaticText(const class QPoint &, const class QStaticText &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QStaticText{}) // "const QStaticText &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QStaticText{}) // "const QStaticText &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2][1] = reflect.TypeOf(QStaticText{}) // "const QStaticText &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter14drawStaticTextEiiRK11QStaticText
  case 1:
    // invoke: _ZN8QPainter14drawStaticTextERK7QPointFRK11QStaticText
  case 2:
    // invoke: _ZN8QPainter14drawStaticTextERK6QPointRK11QStaticText
  default:
    qtrt.ErrorResolve("QPainter", "drawStaticText", args)
  }

}


func (this *QPainter) strokePath(args ...interface{}) () {
  // strokePath(const class QPainterPath &, const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10strokePathERK12QPainterPathRK4QPen
  default:
    qtrt.ErrorResolve("QPainter", "strokePath", args)
  }

}


func (this *QPainter) drawConvexPolygon(args ...interface{}) () {
  // drawConvexPolygon(const class QPoint *, int)
  // drawConvexPolygon(const class QPointF *, int)
  // drawConvexPolygon(const class QPolygonF &)
  // drawConvexPolygon(const class QPolygon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter17drawConvexPolygonEPK6QPointi
  case 1:
    // invoke: _ZN8QPainter17drawConvexPolygonEPK7QPointFi
  case 2:
    // invoke: _ZN8QPainter17drawConvexPolygonERK9QPolygonF
  case 3:
    // invoke: _ZN8QPainter17drawConvexPolygonERK8QPolygon
  default:
    qtrt.ErrorResolve("QPainter", "drawConvexPolygon", args)
  }

}


func (this *QPainter) drawPath(args ...interface{}) () {
  // drawPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawPathERK12QPainterPath
  default:
    qtrt.ErrorResolve("QPainter", "drawPath", args)
  }

}


func (this *QPainter) combinedMatrix(args ...interface{}) () {
  // combinedMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter14combinedMatrixEv
  default:
    qtrt.ErrorResolve("QPainter", "combinedMatrix", args)
  }

}


func (this *QPainter) setMatrixEnabled(args ...interface{}) () {
  // setMatrixEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter16setMatrixEnabledEb
  default:
    qtrt.ErrorResolve("QPainter", "setMatrixEnabled", args)
  }

}


func (this *QPainter) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter7setFontERK5QFont
  default:
    qtrt.ErrorResolve("QPainter", "setFont", args)
  }

}


func (this *QPainter) setWindow(args ...interface{}) () {
  // setWindow(const class QRect &)
  // setWindow(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9setWindowERK5QRect
  case 1:
    // invoke: _ZN8QPainter9setWindowEiiii
  default:
    qtrt.ErrorResolve("QPainter", "setWindow", args)
  }

}


func (this *QPainter) deviceMatrix(args ...interface{}) () {
  // deviceMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter12deviceMatrixEv
  default:
    qtrt.ErrorResolve("QPainter", "deviceMatrix", args)
  }

}


func (this *QPainter) drawPie(args ...interface{}) () {
  // drawPie(int, int, int, int, int, int)
  // drawPie(const class QRect &, int, int)
  // drawPie(const class QRectF &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter7drawPieEiiiiii
  case 1:
    // invoke: _ZN8QPainter7drawPieERK5QRectii
  case 2:
    // invoke: _ZN8QPainter7drawPieERK6QRectFii
  default:
    qtrt.ErrorResolve("QPainter", "drawPie", args)
  }

}


func (this *QPainter) setWorldMatrixEnabled(args ...interface{}) () {
  // setWorldMatrixEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter21setWorldMatrixEnabledEb
  default:
    qtrt.ErrorResolve("QPainter", "setWorldMatrixEnabled", args)
  }

}


func NewQPainter(args ...interface{}) QPainter {
  return QPainter{}
}


func (this *QPainter) drawPoints(args ...interface{}) () {
  // drawPoints(const class QPolygon &)
  // drawPoints(const class QPointF *, int)
  // drawPoints(const class QPolygonF &)
  // drawPoints(const class QPoint *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPoint{}) // "const QPoint *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10drawPointsERK8QPolygon
  case 1:
    // invoke: _ZN8QPainter10drawPointsEPK7QPointFi
  case 2:
    // invoke: _ZN8QPainter10drawPointsERK9QPolygonF
  case 3:
    // invoke: _ZN8QPainter10drawPointsEPK6QPointi
  default:
    qtrt.ErrorResolve("QPainter", "drawPoints", args)
  }

}


func (this *QPainter) drawRect(args ...interface{}) () {
  // drawRect(int, int, int, int)
  // drawRect(const class QRectF &)
  // drawRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawRectEiiii
  case 1:
    // invoke: _ZN8QPainter8drawRectERK6QRectF
  case 2:
    // invoke: _ZN8QPainter8drawRectERK5QRect
  default:
    qtrt.ErrorResolve("QPainter", "drawRect", args)
  }

}


func (this *QPainter) clipRegion(args ...interface{}) () {
  // clipRegion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter10clipRegionEv
  default:
    qtrt.ErrorResolve("QPainter", "clipRegion", args)
  }

}


func (this *QPainter) drawLine(args ...interface{}) () {
  // drawLine(const class QLineF &)
  // drawLine(const class QPointF &, const class QPointF &)
  // drawLine(const class QLine &)
  // drawLine(int, int, int, int)
  // drawLine(const class QPoint &, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLine{}) // "const QLine &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[4][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawLineERK6QLineF
  case 1:
    // invoke: _ZN8QPainter8drawLineERK7QPointFS2_
  case 2:
    // invoke: _ZN8QPainter8drawLineERK5QLine
  case 3:
    // invoke: _ZN8QPainter8drawLineEiiii
  case 4:
    // invoke: _ZN8QPainter8drawLineERK6QPointS2_
  default:
    qtrt.ErrorResolve("QPainter", "drawLine", args)
  }

}


func (this *QPainter) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter6deviceEv
  default:
    qtrt.ErrorResolve("QPainter", "device", args)
  }

}


func (this *QPainter) setViewport(args ...interface{}) () {
  // setViewport(const class QRect &)
  // setViewport(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11setViewportERK5QRect
  case 1:
    // invoke: _ZN8QPainter11setViewportEiiii
  default:
    qtrt.ErrorResolve("QPainter", "setViewport", args)
  }

}


func (this *QPainter) drawTextItem(args ...interface{}) () {
  // drawTextItem(int, int, const class QTextItem &)
  // drawTextItem(const class QPoint &, const class QTextItem &)
  // drawTextItem(const class QPointF &, const class QTextItem &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1][1] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2][1] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12drawTextItemEiiRK9QTextItem
  case 1:
    // invoke: _ZN8QPainter12drawTextItemERK6QPointRK9QTextItem
  case 2:
    // invoke: _ZN8QPainter12drawTextItemERK7QPointFRK9QTextItem
  default:
    qtrt.ErrorResolve("QPainter", "drawTextItem", args)
  }

}


func (this *QPainter) save(args ...interface{}) () {
  // save()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter4saveEv
  default:
    qtrt.ErrorResolve("QPainter", "save", args)
  }

}


func (this *QPainter) combinedTransform(args ...interface{}) () {
  // combinedTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter17combinedTransformEv
  default:
    qtrt.ErrorResolve("QPainter", "combinedTransform", args)
  }

}


func (this *QPainter) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter3endEv
  default:
    qtrt.ErrorResolve("QPainter", "end", args)
  }

}


func (this *QPainter) drawRoundRect(args ...interface{}) () {
  // drawRoundRect(const class QRect &, int, int)
  // drawRoundRect(const class QRectF &, int, int)
  // drawRoundRect(int, int, int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[2][5] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter13drawRoundRectERK5QRectii
  case 1:
    // invoke: _ZN8QPainter13drawRoundRectERK6QRectFii
  case 2:
    // invoke: _ZN8QPainter13drawRoundRectEiiiiii
  default:
    qtrt.ErrorResolve("QPainter", "drawRoundRect", args)
  }

}


func (this *QPainter) setWorldTransform(args ...interface{}) () {
  // setWorldTransform(const class QTransform &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter17setWorldTransformERK10QTransformb
  default:
    qtrt.ErrorResolve("QPainter", "setWorldTransform", args)
  }

}


func (this *QPainter) restore(args ...interface{}) () {
  // restore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter7restoreEv
  default:
    qtrt.ErrorResolve("QPainter", "restore", args)
  }

}


func (this *QPainter) drawPoint(args ...interface{}) () {
  // drawPoint(const class QPoint &)
  // drawPoint(int, int)
  // drawPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawPointERK6QPoint
  case 1:
    // invoke: _ZN8QPainter9drawPointEii
  case 2:
    // invoke: _ZN8QPainter9drawPointERK7QPointF
  default:
    qtrt.ErrorResolve("QPainter", "drawPoint", args)
  }

}


func (this *QPainter) redirected_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainter", "redirected", args)
  }

}


func (this *QPainter) shear(args ...interface{}) () {
  // shear(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter5shearEdd
  default:
    qtrt.ErrorResolve("QPainter", "shear", args)
  }

}


func (this *QPainter) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter4fontEv
  default:
    qtrt.ErrorResolve("QPainter", "font", args)
  }

}


func (this *QPainter) deviceTransform(args ...interface{}) () {
  // deviceTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter15deviceTransformEv
  default:
    qtrt.ErrorResolve("QPainter", "deviceTransform", args)
  }

}


func (this *QPainter) resetMatrix(args ...interface{}) () {
  // resetMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11resetMatrixEv
  default:
    qtrt.ErrorResolve("QPainter", "resetMatrix", args)
  }

}


func (this *QPainter) paintEngine(args ...interface{}) () {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter11paintEngineEv
  default:
    qtrt.ErrorResolve("QPainter", "paintEngine", args)
  }

}


func (this *QPainter) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter8isActiveEv
  default:
    qtrt.ErrorResolve("QPainter", "isActive", args)
  }

}


func (this *QPainter) restoreRedirected_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainter", "restoreRedirected", args)
  }

}


func (this *QPainter) worldMatrixEnabled(args ...interface{}) () {
  // worldMatrixEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter18worldMatrixEnabledEv
  default:
    qtrt.ErrorResolve("QPainter", "worldMatrixEnabled", args)
  }

}


func (this *QPainter) transform(args ...interface{}) () {
  // transform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter9transformEv
  default:
    qtrt.ErrorResolve("QPainter", "transform", args)
  }

}


func (this *QPainter) setRedirected_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainter", "setRedirected", args)
  }

}


func (this *QPainter) fontMetrics(args ...interface{}) () {
  // fontMetrics()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter11fontMetricsEv
  default:
    qtrt.ErrorResolve("QPainter", "fontMetrics", args)
  }

}


func (this *QPainter) drawGlyphRun(args ...interface{}) () {
  // drawGlyphRun(const class QPointF &, const class QGlyphRun &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QGlyphRun{}) // "const QGlyphRun &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12drawGlyphRunERK7QPointFRK9QGlyphRun
  default:
    qtrt.ErrorResolve("QPainter", "drawGlyphRun", args)
  }

}


func (this *QPainter) resetTransform(args ...interface{}) () {
  // resetTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter14resetTransformEv
  default:
    qtrt.ErrorResolve("QPainter", "resetTransform", args)
  }

}


func (this *QPainter) brush(args ...interface{}) () {
  // brush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter5brushEv
  default:
    qtrt.ErrorResolve("QPainter", "brush", args)
  }

}


func (this *QPainter) FreeQPainter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainter", "~QPainter", args)
  }

}


func (this *QPainter) begin(args ...interface{}) () {
  // begin(class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter5beginEP12QPaintDevice
  default:
    qtrt.ErrorResolve("QPainter", "begin", args)
  }

}


func (this *QPainter) scale(args ...interface{}) () {
  // scale(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter5scaleEdd
  default:
    qtrt.ErrorResolve("QPainter", "scale", args)
  }

}


func (this *QPainter) setWorldMatrix(args ...interface{}) () {
  // setWorldMatrix(const class QMatrix &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter14setWorldMatrixERK7QMatrixb
  default:
    qtrt.ErrorResolve("QPainter", "setWorldMatrix", args)
  }

}


func (this *QPainter) clipPath(args ...interface{}) () {
  // clipPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter8clipPathEv
  default:
    qtrt.ErrorResolve("QPainter", "clipPath", args)
  }

}


func (this *QPainter) brushOrigin(args ...interface{}) () {
  // brushOrigin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter11brushOriginEv
  default:
    qtrt.ErrorResolve("QPainter", "brushOrigin", args)
  }

}


func (this *QPainter) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter10backgroundEv
  default:
    qtrt.ErrorResolve("QPainter", "background", args)
  }

}


func (this *QPainter) viewport(args ...interface{}) () {
  // viewport()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter8viewportEv
  default:
    qtrt.ErrorResolve("QPainter", "viewport", args)
  }

}


func (this *QPainter) fillPath(args ...interface{}) () {
  // fillPath(const class QPainterPath &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8fillPathERK12QPainterPathRK6QBrush
  default:
    qtrt.ErrorResolve("QPainter", "fillPath", args)
  }

}


func (this *QPainter) matrixEnabled(args ...interface{}) () {
  // matrixEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter13matrixEnabledEv
  default:
    qtrt.ErrorResolve("QPainter", "matrixEnabled", args)
  }

}


func (this *QPainter) setTransform(args ...interface{}) () {
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
    // invoke: _ZN8QPainter12setTransformERK10QTransformb
  default:
    qtrt.ErrorResolve("QPainter", "setTransform", args)
  }

}


func (this *QPainter) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter6windowEv
  default:
    qtrt.ErrorResolve("QPainter", "window", args)
  }

}


func (this *QPainter) initFrom(args ...interface{}) () {
  // initFrom(const class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8initFromEPK12QPaintDevice
  default:
    qtrt.ErrorResolve("QPainter", "initFrom", args)
  }

}


func (this *QPainter) fontInfo(args ...interface{}) () {
  // fontInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter8fontInfoEv
  default:
    qtrt.ErrorResolve("QPainter", "fontInfo", args)
  }

}


func (this *QPainter) endNativePainting(args ...interface{}) () {
  // endNativePainting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter17endNativePaintingEv
  default:
    qtrt.ErrorResolve("QPainter", "endNativePainting", args)
  }

}


func (this *QPainter) setViewTransformEnabled(args ...interface{}) () {
  // setViewTransformEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter23setViewTransformEnabledEb
  default:
    qtrt.ErrorResolve("QPainter", "setViewTransformEnabled", args)
  }

}


func (this *QPainter) setOpacity(args ...interface{}) () {
  // setOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10setOpacityEd
  default:
    qtrt.ErrorResolve("QPainter", "setOpacity", args)
  }

}


func (this *QPainter) pen(args ...interface{}) () {
  // pen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter3penEv
  default:
    qtrt.ErrorResolve("QPainter", "pen", args)
  }

}


func (this *QPainter) rotate(args ...interface{}) () {
  // rotate(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter6rotateEd
  default:
    qtrt.ErrorResolve("QPainter", "rotate", args)
  }

}


func (this *QPainter) clipBoundingRect(args ...interface{}) () {
  // clipBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter16clipBoundingRectEv
  default:
    qtrt.ErrorResolve("QPainter", "clipBoundingRect", args)
  }

}


func (this *QPainter) beginNativePainting(args ...interface{}) () {
  // beginNativePainting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter19beginNativePaintingEv
  default:
    qtrt.ErrorResolve("QPainter", "beginNativePainting", args)
  }

}

// <= body block end

