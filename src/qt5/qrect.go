package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qrect.h
// dst-file: /src/core/qrect.go
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
  // proto:  QRect QRect::marginsRemoved(const QMargins & margins);
extern void* C_ZNK5QRect14marginsRemovedERK8QMargins(void* qthis, void* arg0); // 2
  // proto:  void QRect::moveTo(int x, int t);
extern void C_ZN5QRect6moveToEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QRect::moveTo(const QPoint & p);
extern void C_ZN5QRect6moveToERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  int QRect::right();
extern int32_t C_ZNK5QRect5rightEv(void* qthis); // 2
  // proto:  bool QRect::intersects(const QRect & r);
extern bool C_ZNK5QRect10intersectsERKS_(void* qthis, void* arg0); // 4
  // proto:  QPoint QRect::topLeft();
extern void* C_ZNK5QRect7topLeftEv(void* qthis); // 2
  // proto:  void QRect::setBottom(int pos);
extern void C_ZN5QRect9setBottomEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::setTopLeft(const QPoint & p);
extern void C_ZN5QRect10setTopLeftERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  QRect QRect::adjusted(int x1, int y1, int x2, int y2);
extern void* C_ZNK5QRect8adjustedEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  int QRect::height();
extern int32_t C_ZNK5QRect6heightEv(void* qthis); // 2
  // proto:  void QRect::setBottomLeft(const QPoint & p);
extern void C_ZN5QRect13setBottomLeftERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::moveRight(int pos);
extern void C_ZN5QRect9moveRightEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::moveBottomLeft(const QPoint & p);
extern void C_ZN5QRect14moveBottomLeftERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::setWidth(int w);
extern void C_ZN5QRect8setWidthEi(void* qthis, int32_t arg0); // 2
  // proto:  QRect QRect::marginsAdded(const QMargins & margins);
extern void* C_ZNK5QRect12marginsAddedERK8QMargins(void* qthis, void* arg0); // 2
  // proto:  QSize QRect::size();
extern void* C_ZNK5QRect4sizeEv(void* qthis); // 2
  // proto:  void QRect::moveBottom(int pos);
extern void C_ZN5QRect10moveBottomEi(void* qthis, int32_t arg0); // 2
  // proto:  QRect QRect::united(const QRect & other);
extern void* C_ZNK5QRect6unitedERKS_(void* qthis, void* arg0); // 2
  // proto:  QPoint QRect::bottomRight();
extern void* C_ZNK5QRect11bottomRightEv(void* qthis); // 2
  // proto:  QRect QRect::intersected(const QRect & other);
extern void* C_ZNK5QRect11intersectedERKS_(void* qthis, void* arg0); // 2
  // proto:  void QRect::getRect(int * x, int * y, int * w, int * h);
extern void C_ZNK5QRect7getRectEPiS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 2
  // proto:  int QRect::top();
extern int32_t C_ZNK5QRect3topEv(void* qthis); // 2
  // proto:  bool QRect::contains(int x, int y, bool proper);
extern bool C_ZNK5QRect8containsEiib(void* qthis, int32_t arg0, int32_t arg1, bool arg2); // 2
  // proto:  bool QRect::contains(const QPoint & p, bool proper);
extern bool C_ZNK5QRect8containsERK6QPointb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  bool QRect::contains(const QRect & r, bool proper);
extern bool C_ZNK5QRect8containsERKS_b(void* qthis, void* arg0, bool arg1); // 4
  // proto:  bool QRect::contains(int x, int y);
extern bool C_ZNK5QRect8containsEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QRect::setTop(int pos);
extern void C_ZN5QRect6setTopEi(void* qthis, int32_t arg0); // 2
  // proto:  QPoint QRect::topRight();
extern void* C_ZNK5QRect8topRightEv(void* qthis); // 2
  // proto:  void QRect::moveCenter(const QPoint & p);
extern void C_ZN5QRect10moveCenterERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  int QRect::width();
extern int32_t C_ZNK5QRect5widthEv(void* qthis); // 2
  // proto:  bool QRect::isEmpty();
extern bool C_ZNK5QRect7isEmptyEv(void* qthis); // 2
  // proto:  void QRect::setX(int x);
extern void C_ZN5QRect4setXEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::setRect(int x, int y, int w, int h);
extern void C_ZN5QRect7setRectEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QRect::setBottomRight(const QPoint & p);
extern void C_ZN5QRect14setBottomRightERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  bool QRect::isValid();
extern bool C_ZNK5QRect7isValidEv(void* qthis); // 2
  // proto:  void QRect::moveLeft(int pos);
extern void C_ZN5QRect8moveLeftEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::setTopRight(const QPoint & p);
extern void C_ZN5QRect11setTopRightERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::moveTop(int pos);
extern void C_ZN5QRect7moveTopEi(void* qthis, int32_t arg0); // 2
  // proto:  QRect QRect::translated(const QPoint & p);
extern void* C_ZNK5QRect10translatedERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  QRect QRect::translated(int dx, int dy);
extern void* C_ZNK5QRect10translatedEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QRect QRect::normalized();
extern void* C_ZNK5QRect10normalizedEv(void* qthis); // 4
  // proto:  void QRect::setCoords(int x1, int y1, int x2, int y2);
extern void C_ZN5QRect9setCoordsEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QRect::setLeft(int pos);
extern void C_ZN5QRect7setLeftEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::setY(int y);
extern void C_ZN5QRect4setYEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::setHeight(int h);
extern void C_ZN5QRect9setHeightEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::QRect(const QPoint & topleft, const QSize & size);
extern void* C_ZN5QRectC2ERK6QPointRK5QSize(void* arg0, void* arg1); // 1
  // proto:  void QRect::QRect();
extern void* C_ZN5QRectC2Ev(); // 1
  // proto:  void QRect::QRect(int left, int top, int width, int height);
extern void* C_ZN5QRectC2Eiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 1
  // proto:  void QRect::QRect(const QPoint & topleft, const QPoint & bottomright);
extern void* C_ZN5QRectC2ERK6QPointS2_(void* arg0, void* arg1); // 1
  // proto:  void QRect::translate(const QPoint & p);
extern void C_ZN5QRect9translateERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::translate(int dx, int dy);
extern void C_ZN5QRect9translateEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QPoint QRect::bottomLeft();
extern void* C_ZNK5QRect10bottomLeftEv(void* qthis); // 2
  // proto:  QPoint QRect::center();
extern void* C_ZNK5QRect6centerEv(void* qthis); // 2
  // proto:  void QRect::moveTopRight(const QPoint & p);
extern void C_ZN5QRect12moveTopRightERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::moveTopLeft(const QPoint & p);
extern void C_ZN5QRect11moveTopLeftERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  int QRect::bottom();
extern int32_t C_ZNK5QRect6bottomEv(void* qthis); // 2
  // proto:  void QRect::getCoords(int * x1, int * y1, int * x2, int * y2);
extern void C_ZNK5QRect9getCoordsEPiS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 2
  // proto:  void QRect::setRight(int pos);
extern void C_ZN5QRect8setRightEi(void* qthis, int32_t arg0); // 2
  // proto:  bool QRect::isNull();
extern bool C_ZNK5QRect6isNullEv(void* qthis); // 2
  // proto:  void QRect::adjust(int x1, int y1, int x2, int y2);
extern void C_ZN5QRect6adjustEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  int QRect::y();
extern void C_ZNK5QRect1yEv(void* qthis); // 2
  // proto:  int QRect::x();
extern void C_ZNK5QRect1xEv(void* qthis); // 2
  // proto:  void QRect::moveBottomRight(const QPoint & p);
extern void C_ZN5QRect15moveBottomRightERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::setSize(const QSize & s);
extern void C_ZN5QRect7setSizeERK5QSize(void* qthis, void* arg0); // 2
  // proto:  int QRect::left();
extern int32_t C_ZNK5QRect4leftEv(void* qthis); // 2
  // proto:  QRectF QRectF::marginsRemoved(const QMarginsF & margins);
extern void* C_ZNK6QRectF14marginsRemovedERK9QMarginsF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::moveTo(const QPointF & p);
extern void C_ZN6QRectF6moveToERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::moveTo(qreal x, qreal y);
extern void C_ZN6QRectF6moveToEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QRectF::setBottomLeft(const QPointF & p);
extern void C_ZN6QRectF13setBottomLeftERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  bool QRectF::intersects(const QRectF & r);
extern bool C_ZNK6QRectF10intersectsERKS_(void* qthis, void* arg0); // 4
  // proto:  QPointF QRectF::topLeft();
extern void* C_ZNK6QRectF7topLeftEv(void* qthis); // 2
  // proto:  void QRectF::setBottom(qreal pos);
extern void C_ZN6QRectF9setBottomEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::setTopLeft(const QPointF & p);
extern void C_ZN6QRectF10setTopLeftERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QRectF QRectF::adjusted(qreal x1, qreal y1, qreal x2, qreal y2);
extern void* C_ZNK6QRectF8adjustedEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  qreal QRectF::height();
extern double C_ZNK6QRectF6heightEv(void* qthis); // 2
  // proto:  qreal QRectF::right();
extern double C_ZNK6QRectF5rightEv(void* qthis); // 2
  // proto:  void QRectF::setLeft(qreal pos);
extern void C_ZN6QRectF7setLeftEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::moveRight(qreal pos);
extern void C_ZN6QRectF9moveRightEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::moveBottomLeft(const QPointF & p);
extern void C_ZN6QRectF14moveBottomLeftERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::setWidth(qreal w);
extern void C_ZN6QRectF8setWidthEd(void* qthis, double arg0); // 2
  // proto:  QRectF QRectF::marginsAdded(const QMarginsF & margins);
extern void* C_ZNK6QRectF12marginsAddedERK9QMarginsF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::moveTop(qreal pos);
extern void C_ZN6QRectF7moveTopEd(void* qthis, double arg0); // 2
  // proto:  QPointF QRectF::bottomRight();
extern void* C_ZNK6QRectF11bottomRightEv(void* qthis); // 2
  // proto:  qreal QRectF::bottom();
extern double C_ZNK6QRectF6bottomEv(void* qthis); // 2
  // proto:  void QRectF::getRect(qreal * x, qreal * y, qreal * w, qreal * h);
extern void C_ZNK6QRectF7getRectEPdS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 2
  // proto:  qreal QRectF::top();
extern double C_ZNK6QRectF3topEv(void* qthis); // 2
  // proto:  bool QRectF::contains(qreal x, qreal y);
extern bool C_ZNK6QRectF8containsEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  bool QRectF::contains(const QRectF & r);
extern bool C_ZNK6QRectF8containsERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QRectF::contains(const QPointF & p);
extern bool C_ZNK6QRectF8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QRect QRectF::toRect();
extern void* C_ZNK6QRectF6toRectEv(void* qthis); // 2
  // proto:  QPointF QRectF::topRight();
extern void* C_ZNK6QRectF8topRightEv(void* qthis); // 2
  // proto:  void QRectF::moveCenter(const QPointF & p);
extern void C_ZN6QRectF10moveCenterERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  qreal QRectF::width();
extern double C_ZNK6QRectF5widthEv(void* qthis); // 2
  // proto:  bool QRectF::isEmpty();
extern bool C_ZNK6QRectF7isEmptyEv(void* qthis); // 2
  // proto:  void QRectF::QRectF(const QPointF & topleft, const QPointF & bottomRight);
extern void* C_ZN6QRectFC2ERK7QPointFS2_(void* arg0, void* arg1); // 1
  // proto:  void QRectF::QRectF(qreal left, qreal top, qreal width, qreal height);
extern void* C_ZN6QRectFC2Edddd(double arg0, double arg1, double arg2, double arg3); // 1
  // proto:  void QRectF::QRectF();
extern void* C_ZN6QRectFC2Ev(); // 1
  // proto:  void QRectF::QRectF(const QPointF & topleft, const QSizeF & size);
extern void* C_ZN6QRectFC2ERK7QPointFRK6QSizeF(void* arg0, void* arg1); // 1
  // proto:  void QRectF::QRectF(const QRect & rect);
extern void* C_ZN6QRectFC2ERK5QRect(void* arg0); // 1
  // proto:  QSizeF QRectF::size();
extern void* C_ZNK6QRectF4sizeEv(void* qthis); // 2
  // proto:  void QRectF::setX(qreal pos);
extern void C_ZN6QRectF4setXEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::setRect(qreal x, qreal y, qreal w, qreal h);
extern void C_ZN6QRectF7setRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QRectF::moveLeft(qreal pos);
extern void C_ZN6QRectF8moveLeftEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::setBottomRight(const QPointF & p);
extern void C_ZN6QRectF14setBottomRightERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  bool QRectF::isValid();
extern bool C_ZNK6QRectF7isValidEv(void* qthis); // 2
  // proto:  void QRectF::getCoords(qreal * x1, qreal * y1, qreal * x2, qreal * y2);
extern void C_ZNK6QRectF9getCoordsEPdS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 2
  // proto:  void QRectF::setTopRight(const QPointF & p);
extern void C_ZN6QRectF11setTopRightERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QRectF QRectF::translated(qreal dx, qreal dy);
extern void* C_ZNK6QRectF10translatedEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QRectF QRectF::translated(const QPointF & p);
extern void* C_ZNK6QRectF10translatedERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QRectF QRectF::normalized();
extern void* C_ZNK6QRectF10normalizedEv(void* qthis); // 4
  // proto:  void QRectF::setTop(qreal pos);
extern void C_ZN6QRectF6setTopEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::moveBottom(qreal pos);
extern void C_ZN6QRectF10moveBottomEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::setY(qreal pos);
extern void C_ZN6QRectF4setYEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::setCoords(qreal x1, qreal y1, qreal x2, qreal y2);
extern void C_ZN6QRectF9setCoordsEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QRectF::setHeight(qreal h);
extern void C_ZN6QRectF9setHeightEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::translate(const QPointF & p);
extern void C_ZN6QRectF9translateERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::translate(qreal dx, qreal dy);
extern void C_ZN6QRectF9translateEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QPointF QRectF::bottomLeft();
extern void* C_ZNK6QRectF10bottomLeftEv(void* qthis); // 2
  // proto:  QPointF QRectF::center();
extern void* C_ZNK6QRectF6centerEv(void* qthis); // 2
  // proto:  void QRectF::moveTopRight(const QPointF & p);
extern void C_ZN6QRectF12moveTopRightERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::moveTopLeft(const QPointF & p);
extern void C_ZN6QRectF11moveTopLeftERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QRect QRectF::toAlignedRect();
extern void* C_ZNK6QRectF13toAlignedRectEv(void* qthis); // 4
  // proto:  QRectF QRectF::intersected(const QRectF & other);
extern void* C_ZNK6QRectF11intersectedERKS_(void* qthis, void* arg0); // 2
  // proto:  QRectF QRectF::united(const QRectF & other);
extern void* C_ZNK6QRectF6unitedERKS_(void* qthis, void* arg0); // 2
  // proto:  void QRectF::setRight(qreal pos);
extern void C_ZN6QRectF8setRightEd(void* qthis, double arg0); // 2
  // proto:  bool QRectF::isNull();
extern bool C_ZNK6QRectF6isNullEv(void* qthis); // 2
  // proto:  void QRectF::adjust(qreal x1, qreal y1, qreal x2, qreal y2);
extern void C_ZN6QRectF6adjustEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  qreal QRectF::y();
extern void C_ZNK6QRectF1yEv(void* qthis); // 2
  // proto:  qreal QRectF::x();
extern void C_ZNK6QRectF1xEv(void* qthis); // 2
  // proto:  qreal QRectF::left();
extern double C_ZNK6QRectF4leftEv(void* qthis); // 2
  // proto:  void QRectF::setSize(const QSizeF & s);
extern void C_ZN6QRectF7setSizeERK6QSizeF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::moveBottomRight(const QPointF & p);
extern void C_ZN6QRectF15moveBottomRightERK7QPointF(void* qthis, void* arg0); // 2
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

// class sizeof(QRect)=16
type QRect struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QRectF)=32
type QRectF struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// marginsRemoved(const class QMargins &)
func (this *QRect) Marginsremoved(args ...interface{}) (ret interface{}) {
  // marginsRemoved(const class QMargins &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMargins{}) // "const QMargins &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect14marginsRemovedERK8QMargins
    // invoke: QRect marginsRemoved(const class QMargins &)
    var arg0 = args[0].(QMargins).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QRect14marginsRemovedERK8QMargins(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "marginsRemoved", args)
  }

  return
}

// moveTo(int, int)
func (this *QRect) Moveto(args ...interface{}) () {
  // moveTo(int, int)
  // moveTo(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect6moveToEii
    // invoke: void moveTo(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN5QRect6moveToEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN5QRect6moveToERK6QPoint
    // invoke: void moveTo(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect6moveToERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveTo", args)
  }

  return
}

// right()
func (this *QRect) Right(args ...interface{}) (ret interface{}) {
  // right()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect5rightEv
    // invoke: int right()
    var ret0 = C.C_ZNK5QRect5rightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "right", args)
  }

  return
}

// intersects(const class QRect &)
func (this *QRect) Intersects(args ...interface{}) (ret interface{}) {
  // intersects(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect10intersectsERKS_
    // invoke: bool intersects(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QRect10intersectsERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "intersects", args)
  }

  return
}

// topLeft()
func (this *QRect) Topleft(args ...interface{}) (ret interface{}) {
  // topLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect7topLeftEv
    // invoke: QPoint topLeft()
    var ret0 = C.C_ZNK5QRect7topLeftEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "topLeft", args)
  }

  return
}

// setBottom(int)
func (this *QRect) Setbottom(args ...interface{}) () {
  // setBottom(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect9setBottomEi
    // invoke: void setBottom(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect9setBottomEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setBottom", args)
  }

  return
}

// setTopLeft(const class QPoint &)
func (this *QRect) Settopleft(args ...interface{}) () {
  // setTopLeft(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect10setTopLeftERK6QPoint
    // invoke: void setTopLeft(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect10setTopLeftERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setTopLeft", args)
  }

  return
}

// adjusted(int, int, int, int)
func (this *QRect) Adjusted(args ...interface{}) (ret interface{}) {
  // adjusted(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect8adjustedEiiii
    // invoke: QRect adjusted(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK5QRect8adjustedEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "adjusted", args)
  }

  return
}

// height()
func (this *QRect) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect6heightEv
    // invoke: int height()
    var ret0 = C.C_ZNK5QRect6heightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "height", args)
  }

  return
}

// setBottomLeft(const class QPoint &)
func (this *QRect) Setbottomleft(args ...interface{}) () {
  // setBottomLeft(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect13setBottomLeftERK6QPoint
    // invoke: void setBottomLeft(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect13setBottomLeftERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setBottomLeft", args)
  }

  return
}

// moveRight(int)
func (this *QRect) Moveright(args ...interface{}) () {
  // moveRight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect9moveRightEi
    // invoke: void moveRight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect9moveRightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveRight", args)
  }

  return
}

// moveBottomLeft(const class QPoint &)
func (this *QRect) Movebottomleft(args ...interface{}) () {
  // moveBottomLeft(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect14moveBottomLeftERK6QPoint
    // invoke: void moveBottomLeft(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect14moveBottomLeftERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveBottomLeft", args)
  }

  return
}

// setWidth(int)
func (this *QRect) Setwidth(args ...interface{}) () {
  // setWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect8setWidthEi
    // invoke: void setWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect8setWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setWidth", args)
  }

  return
}

// marginsAdded(const class QMargins &)
func (this *QRect) Marginsadded(args ...interface{}) (ret interface{}) {
  // marginsAdded(const class QMargins &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMargins{}) // "const QMargins &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect12marginsAddedERK8QMargins
    // invoke: QRect marginsAdded(const class QMargins &)
    var arg0 = args[0].(QMargins).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QRect12marginsAddedERK8QMargins(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "marginsAdded", args)
  }

  return
}

// size()
func (this *QRect) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect4sizeEv
    // invoke: QSize size()
    var ret0 = C.C_ZNK5QRect4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "size", args)
  }

  return
}

// moveBottom(int)
func (this *QRect) Movebottom(args ...interface{}) () {
  // moveBottom(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect10moveBottomEi
    // invoke: void moveBottom(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect10moveBottomEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveBottom", args)
  }

  return
}

// united(const class QRect &)
func (this *QRect) United(args ...interface{}) (ret interface{}) {
  // united(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect6unitedERKS_
    // invoke: QRect united(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QRect6unitedERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "united", args)
  }

  return
}

// bottomRight()
func (this *QRect) Bottomright(args ...interface{}) (ret interface{}) {
  // bottomRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect11bottomRightEv
    // invoke: QPoint bottomRight()
    var ret0 = C.C_ZNK5QRect11bottomRightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "bottomRight", args)
  }

  return
}

// intersected(const class QRect &)
func (this *QRect) Intersected(args ...interface{}) (ret interface{}) {
  // intersected(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect11intersectedERKS_
    // invoke: QRect intersected(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QRect11intersectedERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "intersected", args)
  }

  return
}

// getRect(int *, int *, int *, int *)
func (this *QRect) Getrect(args ...interface{}) () {
  // getRect(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect7getRectEPiS0_S0_S0_
    // invoke: void getRect(int *, int *, int *, int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK5QRect7getRectEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRect", "getRect", args)
  }

  return
}

// top()
func (this *QRect) Top(args ...interface{}) (ret interface{}) {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect3topEv
    // invoke: int top()
    var ret0 = C.C_ZNK5QRect3topEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "top", args)
  }

  return
}

// contains(int, int, _Bool)
func (this *QRect) Contains(args ...interface{}) (ret interface{}) {
  // contains(int, int, _Bool)
  // contains(const class QPoint &, _Bool)
  // contains(const class QRect &, _Bool)
  // contains(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1][1] = qtrt.BoolTy(false) // "bool"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[2][1] = qtrt.BoolTy(false) // "bool"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect8containsEiib
    // invoke: bool contains(int, int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK5QRect8containsEiib(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK5QRect8containsERK6QPointb
    // invoke: bool contains(const class QPoint &, _Bool)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK5QRect8containsERK6QPointb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 2:
    // invoke: _ZNK5QRect8containsERKS_b
    // invoke: bool contains(const class QRect &, _Bool)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK5QRect8containsERKS_b(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 3:
    // invoke: _ZNK5QRect8containsEii
    // invoke: bool contains(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK5QRect8containsEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "contains", args)
  }

  return
}

// setTop(int)
func (this *QRect) Settop(args ...interface{}) () {
  // setTop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect6setTopEi
    // invoke: void setTop(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect6setTopEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setTop", args)
  }

  return
}

// topRight()
func (this *QRect) Topright(args ...interface{}) (ret interface{}) {
  // topRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect8topRightEv
    // invoke: QPoint topRight()
    var ret0 = C.C_ZNK5QRect8topRightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "topRight", args)
  }

  return
}

// moveCenter(const class QPoint &)
func (this *QRect) Movecenter(args ...interface{}) () {
  // moveCenter(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect10moveCenterERK6QPoint
    // invoke: void moveCenter(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect10moveCenterERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveCenter", args)
  }

  return
}

// width()
func (this *QRect) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect5widthEv
    // invoke: int width()
    var ret0 = C.C_ZNK5QRect5widthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "width", args)
  }

  return
}

// isEmpty()
func (this *QRect) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK5QRect7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "isEmpty", args)
  }

  return
}

// setX(int)
func (this *QRect) Setx(args ...interface{}) () {
  // setX(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect4setXEi
    // invoke: void setX(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect4setXEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setX", args)
  }

  return
}

// setRect(int, int, int, int)
func (this *QRect) Setrect(args ...interface{}) () {
  // setRect(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect7setRectEiiii
    // invoke: void setRect(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN5QRect7setRectEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRect", "setRect", args)
  }

  return
}

// setBottomRight(const class QPoint &)
func (this *QRect) Setbottomright(args ...interface{}) () {
  // setBottomRight(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect14setBottomRightERK6QPoint
    // invoke: void setBottomRight(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect14setBottomRightERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setBottomRight", args)
  }

  return
}

// isValid()
func (this *QRect) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK5QRect7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "isValid", args)
  }

  return
}

// moveLeft(int)
func (this *QRect) Moveleft(args ...interface{}) () {
  // moveLeft(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect8moveLeftEi
    // invoke: void moveLeft(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect8moveLeftEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveLeft", args)
  }

  return
}

// setTopRight(const class QPoint &)
func (this *QRect) Settopright(args ...interface{}) () {
  // setTopRight(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect11setTopRightERK6QPoint
    // invoke: void setTopRight(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect11setTopRightERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setTopRight", args)
  }

  return
}

// moveTop(int)
func (this *QRect) Movetop(args ...interface{}) () {
  // moveTop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect7moveTopEi
    // invoke: void moveTop(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect7moveTopEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveTop", args)
  }

  return
}

// translated(const class QPoint &)
func (this *QRect) Translated(args ...interface{}) (ret interface{}) {
  // translated(const class QPoint &)
  // translated(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect10translatedERK6QPoint
    // invoke: QRect translated(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QRect10translatedERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK5QRect10translatedEii
    // invoke: QRect translated(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK5QRect10translatedEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "translated", args)
  }

  return
}

// normalized()
func (this *QRect) Normalized(args ...interface{}) (ret interface{}) {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect10normalizedEv
    // invoke: QRect normalized()
    var ret0 = C.C_ZNK5QRect10normalizedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "normalized", args)
  }

  return
}

// setCoords(int, int, int, int)
func (this *QRect) Setcoords(args ...interface{}) () {
  // setCoords(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect9setCoordsEiiii
    // invoke: void setCoords(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN5QRect9setCoordsEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRect", "setCoords", args)
  }

  return
}

// setLeft(int)
func (this *QRect) Setleft(args ...interface{}) () {
  // setLeft(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect7setLeftEi
    // invoke: void setLeft(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect7setLeftEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setLeft", args)
  }

  return
}

// setY(int)
func (this *QRect) Sety(args ...interface{}) () {
  // setY(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect4setYEi
    // invoke: void setY(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect4setYEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setY", args)
  }

  return
}

// setHeight(int)
func (this *QRect) Setheight(args ...interface{}) () {
  // setHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect9setHeightEi
    // invoke: void setHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect9setHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setHeight", args)
  }

  return
}

// QRect(const class QPoint &, const class QSize &)
func NewQRect(args ...interface{}) *QRect {
  // QRect(const class QPoint &, const class QSize &)
  // QRect()
  // QRect(int, int, int, int)
  // QRect(const class QPoint &, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[3][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRectC1ERK6QPointRK5QSize
    // invoke: void QRect(const class QPoint &, const class QSize &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QSize).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QRectC2ERK6QPointRK5QSize(arg0, arg1)
    return &QRect{qclsinst:qthis}
  case 1:
    // invoke: _ZN5QRectC1Ev
    // invoke: void QRect()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QRectC2Ev()
    return &QRect{qclsinst:qthis}
  case 2:
    // invoke: _ZN5QRectC1Eiiii
    // invoke: void QRect(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QRectC2Eiiii(arg0, arg1, arg2, arg3)
    return &QRect{qclsinst:qthis}
  case 3:
    // invoke: _ZN5QRectC1ERK6QPointS2_
    // invoke: void QRect(const class QPoint &, const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QRectC2ERK6QPointS2_(arg0, arg1)
    return &QRect{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRect", "QRect", args)
  }

  return nil // QRect{qclsinst:qthis}
}

// translate(const class QPoint &)
func (this *QRect) Translate(args ...interface{}) () {
  // translate(const class QPoint &)
  // translate(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect9translateERK6QPoint
    // invoke: void translate(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect9translateERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN5QRect9translateEii
    // invoke: void translate(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN5QRect9translateEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRect", "translate", args)
  }

  return
}

// bottomLeft()
func (this *QRect) Bottomleft(args ...interface{}) (ret interface{}) {
  // bottomLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect10bottomLeftEv
    // invoke: QPoint bottomLeft()
    var ret0 = C.C_ZNK5QRect10bottomLeftEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "bottomLeft", args)
  }

  return
}

// center()
func (this *QRect) Center(args ...interface{}) (ret interface{}) {
  // center()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect6centerEv
    // invoke: QPoint center()
    var ret0 = C.C_ZNK5QRect6centerEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "center", args)
  }

  return
}

// moveTopRight(const class QPoint &)
func (this *QRect) Movetopright(args ...interface{}) () {
  // moveTopRight(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect12moveTopRightERK6QPoint
    // invoke: void moveTopRight(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect12moveTopRightERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveTopRight", args)
  }

  return
}

// moveTopLeft(const class QPoint &)
func (this *QRect) Movetopleft(args ...interface{}) () {
  // moveTopLeft(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect11moveTopLeftERK6QPoint
    // invoke: void moveTopLeft(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect11moveTopLeftERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveTopLeft", args)
  }

  return
}

// bottom()
func (this *QRect) Bottom(args ...interface{}) (ret interface{}) {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect6bottomEv
    // invoke: int bottom()
    var ret0 = C.C_ZNK5QRect6bottomEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "bottom", args)
  }

  return
}

// getCoords(int *, int *, int *, int *)
func (this *QRect) Getcoords(args ...interface{}) () {
  // getCoords(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect9getCoordsEPiS0_S0_S0_
    // invoke: void getCoords(int *, int *, int *, int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK5QRect9getCoordsEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRect", "getCoords", args)
  }

  return
}

// setRight(int)
func (this *QRect) Setright(args ...interface{}) () {
  // setRight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect8setRightEi
    // invoke: void setRight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect8setRightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setRight", args)
  }

  return
}

// isNull()
func (this *QRect) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK5QRect6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "isNull", args)
  }

  return
}

// adjust(int, int, int, int)
func (this *QRect) Adjust(args ...interface{}) () {
  // adjust(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect6adjustEiiii
    // invoke: void adjust(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN5QRect6adjustEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRect", "adjust", args)
  }

  return
}

// y()
func (this *QRect) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect1yEv
    // invoke: int y()
    C.C_ZNK5QRect1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "y", args)
  }

  return
}

// x()
func (this *QRect) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect1xEv
    // invoke: int x()
    C.C_ZNK5QRect1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "x", args)
  }

  return
}

// moveBottomRight(const class QPoint &)
func (this *QRect) Movebottomright(args ...interface{}) () {
  // moveBottomRight(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect15moveBottomRightERK6QPoint
    // invoke: void moveBottomRight(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect15moveBottomRightERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveBottomRight", args)
  }

  return
}

// setSize(const class QSize &)
func (this *QRect) Setsize(args ...interface{}) () {
  // setSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect7setSizeERK5QSize
    // invoke: void setSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QRect7setSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setSize", args)
  }

  return
}

// left()
func (this *QRect) Left(args ...interface{}) (ret interface{}) {
  // left()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect4leftEv
    // invoke: int left()
    var ret0 = C.C_ZNK5QRect4leftEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRect", "left", args)
  }

  return
}

// marginsRemoved(const class QMarginsF &)
func (this *QRectF) Marginsremoved(args ...interface{}) (ret interface{}) {
  // marginsRemoved(const class QMarginsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMarginsF{}) // "const QMarginsF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF14marginsRemovedERK9QMarginsF
    // invoke: QRectF marginsRemoved(const class QMarginsF &)
    var arg0 = args[0].(QMarginsF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QRectF14marginsRemovedERK9QMarginsF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "marginsRemoved", args)
  }

  return
}

// moveTo(const class QPointF &)
func (this *QRectF) Moveto(args ...interface{}) () {
  // moveTo(const class QPointF &)
  // moveTo(qreal, qreal)
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
    // invoke: _ZN6QRectF6moveToERK7QPointF
    // invoke: void moveTo(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF6moveToERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN6QRectF6moveToEdd
    // invoke: void moveTo(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN6QRectF6moveToEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRectF", "moveTo", args)
  }

  return
}

// setBottomLeft(const class QPointF &)
func (this *QRectF) Setbottomleft(args ...interface{}) () {
  // setBottomLeft(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF13setBottomLeftERK7QPointF
    // invoke: void setBottomLeft(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF13setBottomLeftERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setBottomLeft", args)
  }

  return
}

// intersects(const class QRectF &)
func (this *QRectF) Intersects(args ...interface{}) (ret interface{}) {
  // intersects(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF10intersectsERKS_
    // invoke: bool intersects(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QRectF10intersectsERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "intersects", args)
  }

  return
}

// topLeft()
func (this *QRectF) Topleft(args ...interface{}) (ret interface{}) {
  // topLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF7topLeftEv
    // invoke: QPointF topLeft()
    var ret0 = C.C_ZNK6QRectF7topLeftEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "topLeft", args)
  }

  return
}

// setBottom(qreal)
func (this *QRectF) Setbottom(args ...interface{}) () {
  // setBottom(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF9setBottomEd
    // invoke: void setBottom(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF9setBottomEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setBottom", args)
  }

  return
}

// setTopLeft(const class QPointF &)
func (this *QRectF) Settopleft(args ...interface{}) () {
  // setTopLeft(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF10setTopLeftERK7QPointF
    // invoke: void setTopLeft(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF10setTopLeftERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setTopLeft", args)
  }

  return
}

// adjusted(qreal, qreal, qreal, qreal)
func (this *QRectF) Adjusted(args ...interface{}) (ret interface{}) {
  // adjusted(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF8adjustedEdddd
    // invoke: QRectF adjusted(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK6QRectF8adjustedEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "adjusted", args)
  }

  return
}

// height()
func (this *QRectF) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6heightEv
    // invoke: qreal height()
    var ret0 = C.C_ZNK6QRectF6heightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "height", args)
  }

  return
}

// right()
func (this *QRectF) Right(args ...interface{}) (ret interface{}) {
  // right()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF5rightEv
    // invoke: qreal right()
    var ret0 = C.C_ZNK6QRectF5rightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "right", args)
  }

  return
}

// setLeft(qreal)
func (this *QRectF) Setleft(args ...interface{}) () {
  // setLeft(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF7setLeftEd
    // invoke: void setLeft(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF7setLeftEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setLeft", args)
  }

  return
}

// moveRight(qreal)
func (this *QRectF) Moveright(args ...interface{}) () {
  // moveRight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF9moveRightEd
    // invoke: void moveRight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF9moveRightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveRight", args)
  }

  return
}

// moveBottomLeft(const class QPointF &)
func (this *QRectF) Movebottomleft(args ...interface{}) () {
  // moveBottomLeft(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF14moveBottomLeftERK7QPointF
    // invoke: void moveBottomLeft(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF14moveBottomLeftERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveBottomLeft", args)
  }

  return
}

// setWidth(qreal)
func (this *QRectF) Setwidth(args ...interface{}) () {
  // setWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF8setWidthEd
    // invoke: void setWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF8setWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setWidth", args)
  }

  return
}

// marginsAdded(const class QMarginsF &)
func (this *QRectF) Marginsadded(args ...interface{}) (ret interface{}) {
  // marginsAdded(const class QMarginsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMarginsF{}) // "const QMarginsF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF12marginsAddedERK9QMarginsF
    // invoke: QRectF marginsAdded(const class QMarginsF &)
    var arg0 = args[0].(QMarginsF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QRectF12marginsAddedERK9QMarginsF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "marginsAdded", args)
  }

  return
}

// moveTop(qreal)
func (this *QRectF) Movetop(args ...interface{}) () {
  // moveTop(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF7moveTopEd
    // invoke: void moveTop(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF7moveTopEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveTop", args)
  }

  return
}

// bottomRight()
func (this *QRectF) Bottomright(args ...interface{}) (ret interface{}) {
  // bottomRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF11bottomRightEv
    // invoke: QPointF bottomRight()
    var ret0 = C.C_ZNK6QRectF11bottomRightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "bottomRight", args)
  }

  return
}

// bottom()
func (this *QRectF) Bottom(args ...interface{}) (ret interface{}) {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6bottomEv
    // invoke: qreal bottom()
    var ret0 = C.C_ZNK6QRectF6bottomEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "bottom", args)
  }

  return
}

// getRect(qreal *, qreal *, qreal *, qreal *)
func (this *QRectF) Getrect(args ...interface{}) () {
  // getRect(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF7getRectEPdS0_S0_S0_
    // invoke: void getRect(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (unsafe.Pointer)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QRectF7getRectEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRectF", "getRect", args)
  }

  return
}

// top()
func (this *QRectF) Top(args ...interface{}) (ret interface{}) {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF3topEv
    // invoke: qreal top()
    var ret0 = C.C_ZNK6QRectF3topEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "top", args)
  }

  return
}

// contains(qreal, qreal)
func (this *QRectF) Contains(args ...interface{}) (ret interface{}) {
  // contains(qreal, qreal)
  // contains(const class QRectF &)
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF8containsEdd
    // invoke: bool contains(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK6QRectF8containsEdd(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK6QRectF8containsERKS_
    // invoke: bool contains(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QRectF8containsERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 2:
    // invoke: _ZNK6QRectF8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QRectF8containsERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "contains", args)
  }

  return
}

// toRect()
func (this *QRectF) Torect(args ...interface{}) (ret interface{}) {
  // toRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6toRectEv
    // invoke: QRect toRect()
    var ret0 = C.C_ZNK6QRectF6toRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "toRect", args)
  }

  return
}

// topRight()
func (this *QRectF) Topright(args ...interface{}) (ret interface{}) {
  // topRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF8topRightEv
    // invoke: QPointF topRight()
    var ret0 = C.C_ZNK6QRectF8topRightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "topRight", args)
  }

  return
}

// moveCenter(const class QPointF &)
func (this *QRectF) Movecenter(args ...interface{}) () {
  // moveCenter(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF10moveCenterERK7QPointF
    // invoke: void moveCenter(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF10moveCenterERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveCenter", args)
  }

  return
}

// width()
func (this *QRectF) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF5widthEv
    // invoke: qreal width()
    var ret0 = C.C_ZNK6QRectF5widthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "width", args)
  }

  return
}

// isEmpty()
func (this *QRectF) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK6QRectF7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "isEmpty", args)
  }

  return
}

// QRectF(const class QPointF &, const class QPointF &)
func NewQRectF(args ...interface{}) *QRectF {
  // QRectF(const class QPointF &, const class QPointF &)
  // QRectF(qreal, qreal, qreal, qreal)
  // QRectF()
  // QRectF(const class QPointF &, const class QSizeF &)
  // QRectF(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[3][1] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectFC1ERK7QPointFS2_
    // invoke: void QRectF(const class QPointF &, const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QRectFC2ERK7QPointFS2_(arg0, arg1)
    return &QRectF{qclsinst:qthis}
  case 1:
    // invoke: _ZN6QRectFC1Edddd
    // invoke: void QRectF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QRectFC2Edddd(arg0, arg1, arg2, arg3)
    return &QRectF{qclsinst:qthis}
  case 2:
    // invoke: _ZN6QRectFC1Ev
    // invoke: void QRectF()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QRectFC2Ev()
    return &QRectF{qclsinst:qthis}
  case 3:
    // invoke: _ZN6QRectFC1ERK7QPointFRK6QSizeF
    // invoke: void QRectF(const class QPointF &, const class QSizeF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QSizeF).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QRectFC2ERK7QPointFRK6QSizeF(arg0, arg1)
    return &QRectF{qclsinst:qthis}
  case 4:
    // invoke: _ZN6QRectFC1ERK5QRect
    // invoke: void QRectF(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QRectFC2ERK5QRect(arg0)
    return &QRectF{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRectF", "QRectF", args)
  }

  return nil // QRectF{qclsinst:qthis}
}

// size()
func (this *QRectF) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF4sizeEv
    // invoke: QSizeF size()
    var ret0 = C.C_ZNK6QRectF4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizeF{}) // "QSizeF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "size", args)
  }

  return
}

// setX(qreal)
func (this *QRectF) Setx(args ...interface{}) () {
  // setX(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF4setXEd
    // invoke: void setX(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF4setXEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setX", args)
  }

  return
}

// setRect(qreal, qreal, qreal, qreal)
func (this *QRectF) Setrect(args ...interface{}) () {
  // setRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF7setRectEdddd
    // invoke: void setRect(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN6QRectF7setRectEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRectF", "setRect", args)
  }

  return
}

// moveLeft(qreal)
func (this *QRectF) Moveleft(args ...interface{}) () {
  // moveLeft(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF8moveLeftEd
    // invoke: void moveLeft(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF8moveLeftEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveLeft", args)
  }

  return
}

// setBottomRight(const class QPointF &)
func (this *QRectF) Setbottomright(args ...interface{}) () {
  // setBottomRight(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF14setBottomRightERK7QPointF
    // invoke: void setBottomRight(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF14setBottomRightERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setBottomRight", args)
  }

  return
}

// isValid()
func (this *QRectF) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK6QRectF7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "isValid", args)
  }

  return
}

// getCoords(qreal *, qreal *, qreal *, qreal *)
func (this *QRectF) Getcoords(args ...interface{}) () {
  // getCoords(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF9getCoordsEPdS0_S0_S0_
    // invoke: void getCoords(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (unsafe.Pointer)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QRectF9getCoordsEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRectF", "getCoords", args)
  }

  return
}

// setTopRight(const class QPointF &)
func (this *QRectF) Settopright(args ...interface{}) () {
  // setTopRight(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF11setTopRightERK7QPointF
    // invoke: void setTopRight(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF11setTopRightERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setTopRight", args)
  }

  return
}

// translated(qreal, qreal)
func (this *QRectF) Translated(args ...interface{}) (ret interface{}) {
  // translated(qreal, qreal)
  // translated(const class QPointF &)
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
    // invoke: _ZNK6QRectF10translatedEdd
    // invoke: QRectF translated(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK6QRectF10translatedEdd(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK6QRectF10translatedERK7QPointF
    // invoke: QRectF translated(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QRectF10translatedERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "translated", args)
  }

  return
}

// normalized()
func (this *QRectF) Normalized(args ...interface{}) (ret interface{}) {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF10normalizedEv
    // invoke: QRectF normalized()
    var ret0 = C.C_ZNK6QRectF10normalizedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "normalized", args)
  }

  return
}

// setTop(qreal)
func (this *QRectF) Settop(args ...interface{}) () {
  // setTop(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF6setTopEd
    // invoke: void setTop(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF6setTopEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setTop", args)
  }

  return
}

// moveBottom(qreal)
func (this *QRectF) Movebottom(args ...interface{}) () {
  // moveBottom(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF10moveBottomEd
    // invoke: void moveBottom(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF10moveBottomEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveBottom", args)
  }

  return
}

// setY(qreal)
func (this *QRectF) Sety(args ...interface{}) () {
  // setY(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF4setYEd
    // invoke: void setY(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF4setYEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setY", args)
  }

  return
}

// setCoords(qreal, qreal, qreal, qreal)
func (this *QRectF) Setcoords(args ...interface{}) () {
  // setCoords(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF9setCoordsEdddd
    // invoke: void setCoords(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN6QRectF9setCoordsEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRectF", "setCoords", args)
  }

  return
}

// setHeight(qreal)
func (this *QRectF) Setheight(args ...interface{}) () {
  // setHeight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF9setHeightEd
    // invoke: void setHeight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF9setHeightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setHeight", args)
  }

  return
}

// translate(const class QPointF &)
func (this *QRectF) Translate(args ...interface{}) () {
  // translate(const class QPointF &)
  // translate(qreal, qreal)
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
    // invoke: _ZN6QRectF9translateERK7QPointF
    // invoke: void translate(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF9translateERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN6QRectF9translateEdd
    // invoke: void translate(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN6QRectF9translateEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRectF", "translate", args)
  }

  return
}

// bottomLeft()
func (this *QRectF) Bottomleft(args ...interface{}) (ret interface{}) {
  // bottomLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF10bottomLeftEv
    // invoke: QPointF bottomLeft()
    var ret0 = C.C_ZNK6QRectF10bottomLeftEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "bottomLeft", args)
  }

  return
}

// center()
func (this *QRectF) Center(args ...interface{}) (ret interface{}) {
  // center()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6centerEv
    // invoke: QPointF center()
    var ret0 = C.C_ZNK6QRectF6centerEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "center", args)
  }

  return
}

// moveTopRight(const class QPointF &)
func (this *QRectF) Movetopright(args ...interface{}) () {
  // moveTopRight(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF12moveTopRightERK7QPointF
    // invoke: void moveTopRight(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF12moveTopRightERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveTopRight", args)
  }

  return
}

// moveTopLeft(const class QPointF &)
func (this *QRectF) Movetopleft(args ...interface{}) () {
  // moveTopLeft(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF11moveTopLeftERK7QPointF
    // invoke: void moveTopLeft(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF11moveTopLeftERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveTopLeft", args)
  }

  return
}

// toAlignedRect()
func (this *QRectF) Toalignedrect(args ...interface{}) (ret interface{}) {
  // toAlignedRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF13toAlignedRectEv
    // invoke: QRect toAlignedRect()
    var ret0 = C.C_ZNK6QRectF13toAlignedRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "toAlignedRect", args)
  }

  return
}

// intersected(const class QRectF &)
func (this *QRectF) Intersected(args ...interface{}) (ret interface{}) {
  // intersected(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF11intersectedERKS_
    // invoke: QRectF intersected(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QRectF11intersectedERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "intersected", args)
  }

  return
}

// united(const class QRectF &)
func (this *QRectF) United(args ...interface{}) (ret interface{}) {
  // united(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6unitedERKS_
    // invoke: QRectF united(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QRectF6unitedERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "united", args)
  }

  return
}

// setRight(qreal)
func (this *QRectF) Setright(args ...interface{}) () {
  // setRight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF8setRightEd
    // invoke: void setRight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF8setRightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setRight", args)
  }

  return
}

// isNull()
func (this *QRectF) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK6QRectF6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "isNull", args)
  }

  return
}

// adjust(qreal, qreal, qreal, qreal)
func (this *QRectF) Adjust(args ...interface{}) () {
  // adjust(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF6adjustEdddd
    // invoke: void adjust(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN6QRectF6adjustEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRectF", "adjust", args)
  }

  return
}

// y()
func (this *QRectF) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF1yEv
    // invoke: qreal y()
    C.C_ZNK6QRectF1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "y", args)
  }

  return
}

// x()
func (this *QRectF) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF1xEv
    // invoke: qreal x()
    C.C_ZNK6QRectF1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "x", args)
  }

  return
}

// left()
func (this *QRectF) Left(args ...interface{}) (ret interface{}) {
  // left()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF4leftEv
    // invoke: qreal left()
    var ret0 = C.C_ZNK6QRectF4leftEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRectF", "left", args)
  }

  return
}

// setSize(const class QSizeF &)
func (this *QRectF) Setsize(args ...interface{}) () {
  // setSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF7setSizeERK6QSizeF
    // invoke: void setSize(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF7setSizeERK6QSizeF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setSize", args)
  }

  return
}

// moveBottomRight(const class QPointF &)
func (this *QRectF) Movebottomright(args ...interface{}) () {
  // moveBottomRight(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF15moveBottomRightERK7QPointF
    // invoke: void moveBottomRight(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QRectF15moveBottomRightERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveBottomRight", args)
  }

  return
}

// <= body block end

