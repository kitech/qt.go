package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZNK5QRect14marginsRemovedERK8QMargins(void* qthis, void* arg0); // 2
  // proto:  void QRect::moveTo(int x, int t);
extern void _ZN5QRect6moveToEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QRect::moveTo(const QPoint & p);
extern void _ZN5QRect6moveToERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  int QRect::right();
extern void _ZNK5QRect5rightEv(void* qthis); // 2
  // proto:  bool QRect::intersects(const QRect & r);
extern void _ZNK5QRect10intersectsERKS_(void* qthis, void* arg0); // 4
  // proto:  QPoint QRect::topLeft();
extern void _ZNK5QRect7topLeftEv(void* qthis); // 2
  // proto:  void QRect::setBottom(int pos);
extern void _ZN5QRect9setBottomEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::setTopLeft(const QPoint & p);
extern void _ZN5QRect10setTopLeftERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  QRect QRect::adjusted(int x1, int y1, int x2, int y2);
extern void _ZNK5QRect8adjustedEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  int QRect::height();
extern void _ZNK5QRect6heightEv(void* qthis); // 2
  // proto:  void QRect::setBottomLeft(const QPoint & p);
extern void _ZN5QRect13setBottomLeftERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::moveRight(int pos);
extern void _ZN5QRect9moveRightEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::moveBottomLeft(const QPoint & p);
extern void _ZN5QRect14moveBottomLeftERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::setWidth(int w);
extern void _ZN5QRect8setWidthEi(void* qthis, int32_t arg0); // 2
  // proto:  QRect QRect::marginsAdded(const QMargins & margins);
extern void _ZNK5QRect12marginsAddedERK8QMargins(void* qthis, void* arg0); // 2
  // proto:  QSize QRect::size();
extern void _ZNK5QRect4sizeEv(void* qthis); // 2
  // proto:  void QRect::moveBottom(int pos);
extern void _ZN5QRect10moveBottomEi(void* qthis, int32_t arg0); // 2
  // proto:  QRect QRect::united(const QRect & other);
extern void _ZNK5QRect6unitedERKS_(void* qthis, void* arg0); // 2
  // proto:  QPoint QRect::bottomRight();
extern void _ZNK5QRect11bottomRightEv(void* qthis); // 2
  // proto:  QRect QRect::intersected(const QRect & other);
extern void _ZNK5QRect11intersectedERKS_(void* qthis, void* arg0); // 2
  // proto:  void QRect::getRect(int * x, int * y, int * w, int * h);
extern void _ZNK5QRect7getRectEPiS0_S0_S0_(void* qthis, int32_t* arg0, int32_t* arg1, int32_t* arg2, int32_t* arg3); // 2
  // proto:  int QRect::top();
extern void _ZNK5QRect3topEv(void* qthis); // 2
  // proto:  bool QRect::contains(int x, int y, bool proper);
extern void _ZNK5QRect8containsEiib(void* qthis, int32_t arg0, int32_t arg1, bool arg2); // 2
  // proto:  bool QRect::contains(const QPoint & p, bool proper);
extern void _ZNK5QRect8containsERK6QPointb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  bool QRect::contains(const QRect & r, bool proper);
extern void _ZNK5QRect8containsERKS_b(void* qthis, void* arg0, bool arg1); // 4
  // proto:  bool QRect::contains(int x, int y);
extern void _ZNK5QRect8containsEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QRect::setTop(int pos);
extern void _ZN5QRect6setTopEi(void* qthis, int32_t arg0); // 2
  // proto:  QPoint QRect::topRight();
extern void _ZNK5QRect8topRightEv(void* qthis); // 2
  // proto:  void QRect::moveCenter(const QPoint & p);
extern void _ZN5QRect10moveCenterERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  int QRect::width();
extern void _ZNK5QRect5widthEv(void* qthis); // 2
  // proto:  bool QRect::isEmpty();
extern void _ZNK5QRect7isEmptyEv(void* qthis); // 2
  // proto:  void QRect::setX(int x);
extern void _ZN5QRect4setXEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::setRect(int x, int y, int w, int h);
extern void _ZN5QRect7setRectEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QRect::setBottomRight(const QPoint & p);
extern void _ZN5QRect14setBottomRightERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  bool QRect::isValid();
extern void _ZNK5QRect7isValidEv(void* qthis); // 2
  // proto:  void QRect::moveLeft(int pos);
extern void _ZN5QRect8moveLeftEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::setTopRight(const QPoint & p);
extern void _ZN5QRect11setTopRightERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::moveTop(int pos);
extern void _ZN5QRect7moveTopEi(void* qthis, int32_t arg0); // 2
  // proto:  QRect QRect::translated(const QPoint & p);
extern void _ZNK5QRect10translatedERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  QRect QRect::translated(int dx, int dy);
extern void _ZNK5QRect10translatedEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QRect QRect::normalized();
extern void _ZNK5QRect10normalizedEv(void* qthis); // 4
  // proto:  void QRect::setCoords(int x1, int y1, int x2, int y2);
extern void _ZN5QRect9setCoordsEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QRect::setLeft(int pos);
extern void _ZN5QRect7setLeftEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::setY(int y);
extern void _ZN5QRect4setYEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::setHeight(int h);
extern void _ZN5QRect9setHeightEi(void* qthis, int32_t arg0); // 2
  // proto:  void QRect::QRect(const QPoint & topleft, const QSize & size);
extern void _ZN5QRectC2ERK6QPointRK5QSize(void* qthis, void* arg0, void* arg1); // 1
  // proto:  void QRect::QRect();
extern void _ZN5QRectC2Ev(void* qthis); // 1
  // proto:  void QRect::QRect(int left, int top, int width, int height);
extern void _ZN5QRectC2Eiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 1
  // proto:  void QRect::QRect(const QPoint & topleft, const QPoint & bottomright);
extern void _ZN5QRectC2ERK6QPointS2_(void* qthis, void* arg0, void* arg1); // 1
  // proto:  void QRect::translate(const QPoint & p);
extern void _ZN5QRect9translateERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::translate(int dx, int dy);
extern void _ZN5QRect9translateEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QPoint QRect::bottomLeft();
extern void _ZNK5QRect10bottomLeftEv(void* qthis); // 2
  // proto:  QPoint QRect::center();
extern void _ZNK5QRect6centerEv(void* qthis); // 2
  // proto:  void QRect::moveTopRight(const QPoint & p);
extern void _ZN5QRect12moveTopRightERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::moveTopLeft(const QPoint & p);
extern void _ZN5QRect11moveTopLeftERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  int QRect::bottom();
extern void _ZNK5QRect6bottomEv(void* qthis); // 2
  // proto:  void QRect::getCoords(int * x1, int * y1, int * x2, int * y2);
extern void _ZNK5QRect9getCoordsEPiS0_S0_S0_(void* qthis, int32_t* arg0, int32_t* arg1, int32_t* arg2, int32_t* arg3); // 2
  // proto:  void QRect::setRight(int pos);
extern void _ZN5QRect8setRightEi(void* qthis, int32_t arg0); // 2
  // proto:  bool QRect::isNull();
extern void _ZNK5QRect6isNullEv(void* qthis); // 2
  // proto:  void QRect::adjust(int x1, int y1, int x2, int y2);
extern void _ZN5QRect6adjustEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  int QRect::y();
extern void _ZNK5QRect1yEv(void* qthis); // 2
  // proto:  int QRect::x();
extern void _ZNK5QRect1xEv(void* qthis); // 2
  // proto:  void QRect::moveBottomRight(const QPoint & p);
extern void _ZN5QRect15moveBottomRightERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRect::setSize(const QSize & s);
extern void _ZN5QRect7setSizeERK5QSize(void* qthis, void* arg0); // 2
  // proto:  int QRect::left();
extern void _ZNK5QRect4leftEv(void* qthis); // 2
  // proto:  QRectF QRectF::marginsRemoved(const QMarginsF & margins);
extern void _ZNK6QRectF14marginsRemovedERK9QMarginsF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::moveTo(const QPointF & p);
extern void _ZN6QRectF6moveToERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::moveTo(qreal x, qreal y);
extern void _ZN6QRectF6moveToEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QRectF::setBottomLeft(const QPointF & p);
extern void _ZN6QRectF13setBottomLeftERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  bool QRectF::intersects(const QRectF & r);
extern void _ZNK6QRectF10intersectsERKS_(void* qthis, void* arg0); // 4
  // proto:  QPointF QRectF::topLeft();
extern void _ZNK6QRectF7topLeftEv(void* qthis); // 2
  // proto:  void QRectF::setBottom(qreal pos);
extern void _ZN6QRectF9setBottomEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::setTopLeft(const QPointF & p);
extern void _ZN6QRectF10setTopLeftERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QRectF QRectF::adjusted(qreal x1, qreal y1, qreal x2, qreal y2);
extern void _ZNK6QRectF8adjustedEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  qreal QRectF::height();
extern void _ZNK6QRectF6heightEv(void* qthis); // 2
  // proto:  qreal QRectF::right();
extern void _ZNK6QRectF5rightEv(void* qthis); // 2
  // proto:  void QRectF::setLeft(qreal pos);
extern void _ZN6QRectF7setLeftEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::moveRight(qreal pos);
extern void _ZN6QRectF9moveRightEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::moveBottomLeft(const QPointF & p);
extern void _ZN6QRectF14moveBottomLeftERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::setWidth(qreal w);
extern void _ZN6QRectF8setWidthEd(void* qthis, double arg0); // 2
  // proto:  QRectF QRectF::marginsAdded(const QMarginsF & margins);
extern void _ZNK6QRectF12marginsAddedERK9QMarginsF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::moveTop(qreal pos);
extern void _ZN6QRectF7moveTopEd(void* qthis, double arg0); // 2
  // proto:  QPointF QRectF::bottomRight();
extern void _ZNK6QRectF11bottomRightEv(void* qthis); // 2
  // proto:  qreal QRectF::bottom();
extern void _ZNK6QRectF6bottomEv(void* qthis); // 2
  // proto:  void QRectF::getRect(qreal * x, qreal * y, qreal * w, qreal * h);
extern void _ZNK6QRectF7getRectEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3); // 2
  // proto:  qreal QRectF::top();
extern void _ZNK6QRectF3topEv(void* qthis); // 2
  // proto:  bool QRectF::contains(qreal x, qreal y);
extern void _ZNK6QRectF8containsEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  bool QRectF::contains(const QRectF & r);
extern void _ZNK6QRectF8containsERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QRectF::contains(const QPointF & p);
extern void _ZNK6QRectF8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QRect QRectF::toRect();
extern void _ZNK6QRectF6toRectEv(void* qthis); // 2
  // proto:  QPointF QRectF::topRight();
extern void _ZNK6QRectF8topRightEv(void* qthis); // 2
  // proto:  void QRectF::moveCenter(const QPointF & p);
extern void _ZN6QRectF10moveCenterERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  qreal QRectF::width();
extern void _ZNK6QRectF5widthEv(void* qthis); // 2
  // proto:  bool QRectF::isEmpty();
extern void _ZNK6QRectF7isEmptyEv(void* qthis); // 2
  // proto:  void QRectF::QRectF(const QPointF & topleft, const QPointF & bottomRight);
extern void _ZN6QRectFC2ERK7QPointFS2_(void* qthis, void* arg0, void* arg1); // 1
  // proto:  void QRectF::QRectF(qreal left, qreal top, qreal width, qreal height);
extern void _ZN6QRectFC2Edddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 1
  // proto:  void QRectF::QRectF();
extern void _ZN6QRectFC2Ev(void* qthis); // 1
  // proto:  void QRectF::QRectF(const QPointF & topleft, const QSizeF & size);
extern void _ZN6QRectFC2ERK7QPointFRK6QSizeF(void* qthis, void* arg0, void* arg1); // 1
  // proto:  void QRectF::QRectF(const QRect & rect);
extern void _ZN6QRectFC2ERK5QRect(void* qthis, void* arg0); // 1
  // proto:  QSizeF QRectF::size();
extern void _ZNK6QRectF4sizeEv(void* qthis); // 2
  // proto:  void QRectF::setX(qreal pos);
extern void _ZN6QRectF4setXEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::setRect(qreal x, qreal y, qreal w, qreal h);
extern void _ZN6QRectF7setRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QRectF::moveLeft(qreal pos);
extern void _ZN6QRectF8moveLeftEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::setBottomRight(const QPointF & p);
extern void _ZN6QRectF14setBottomRightERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  bool QRectF::isValid();
extern void _ZNK6QRectF7isValidEv(void* qthis); // 2
  // proto:  void QRectF::getCoords(qreal * x1, qreal * y1, qreal * x2, qreal * y2);
extern void _ZNK6QRectF9getCoordsEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3); // 2
  // proto:  void QRectF::setTopRight(const QPointF & p);
extern void _ZN6QRectF11setTopRightERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QRectF QRectF::translated(qreal dx, qreal dy);
extern void _ZNK6QRectF10translatedEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QRectF QRectF::translated(const QPointF & p);
extern void _ZNK6QRectF10translatedERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QRectF QRectF::normalized();
extern void _ZNK6QRectF10normalizedEv(void* qthis); // 4
  // proto:  void QRectF::setTop(qreal pos);
extern void _ZN6QRectF6setTopEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::moveBottom(qreal pos);
extern void _ZN6QRectF10moveBottomEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::setY(qreal pos);
extern void _ZN6QRectF4setYEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::setCoords(qreal x1, qreal y1, qreal x2, qreal y2);
extern void _ZN6QRectF9setCoordsEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QRectF::setHeight(qreal h);
extern void _ZN6QRectF9setHeightEd(void* qthis, double arg0); // 2
  // proto:  void QRectF::translate(const QPointF & p);
extern void _ZN6QRectF9translateERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::translate(qreal dx, qreal dy);
extern void _ZN6QRectF9translateEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QPointF QRectF::bottomLeft();
extern void _ZNK6QRectF10bottomLeftEv(void* qthis); // 2
  // proto:  QPointF QRectF::center();
extern void _ZNK6QRectF6centerEv(void* qthis); // 2
  // proto:  void QRectF::moveTopRight(const QPointF & p);
extern void _ZN6QRectF12moveTopRightERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::moveTopLeft(const QPointF & p);
extern void _ZN6QRectF11moveTopLeftERK7QPointF(void* qthis, void* arg0); // 2
  // proto:  QRect QRectF::toAlignedRect();
extern void _ZNK6QRectF13toAlignedRectEv(void* qthis); // 4
  // proto:  QRectF QRectF::intersected(const QRectF & other);
extern void _ZNK6QRectF11intersectedERKS_(void* qthis, void* arg0); // 2
  // proto:  QRectF QRectF::united(const QRectF & other);
extern void _ZNK6QRectF6unitedERKS_(void* qthis, void* arg0); // 2
  // proto:  void QRectF::setRight(qreal pos);
extern void _ZN6QRectF8setRightEd(void* qthis, double arg0); // 2
  // proto:  bool QRectF::isNull();
extern void _ZNK6QRectF6isNullEv(void* qthis); // 2
  // proto:  void QRectF::adjust(qreal x1, qreal y1, qreal x2, qreal y2);
extern void _ZN6QRectF6adjustEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  qreal QRectF::y();
extern void _ZNK6QRectF1yEv(void* qthis); // 2
  // proto:  qreal QRectF::x();
extern void _ZNK6QRectF1xEv(void* qthis); // 2
  // proto:  qreal QRectF::left();
extern void _ZNK6QRectF4leftEv(void* qthis); // 2
  // proto:  void QRectF::setSize(const QSizeF & s);
extern void _ZN6QRectF7setSizeERK6QSizeF(void* qthis, void* arg0); // 2
  // proto:  void QRectF::moveBottomRight(const QPointF & p);
extern void _ZN6QRectF15moveBottomRightERK7QPointF(void* qthis, void* arg0); // 2
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
func (this *QRect) marginsRemoved(args ...interface{}) () {
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
    C._ZNK5QRect14marginsRemovedERK8QMargins(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "marginsRemoved", args)
  }

}

// moveTo(int, int)
func (this *QRect) moveTo(args ...interface{}) () {
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
    C._ZN5QRect6moveToEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN5QRect6moveToERK6QPoint
    // invoke: void moveTo(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QRect6moveToERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveTo", args)
  }

}

// right()
func (this *QRect) right(args ...interface{}) () {
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
    C._ZNK5QRect5rightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "right", args)
  }

}

// intersects(const class QRect &)
func (this *QRect) intersects(args ...interface{}) () {
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
    C._ZNK5QRect10intersectsERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "intersects", args)
  }

}

// topLeft()
func (this *QRect) topLeft(args ...interface{}) () {
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
    C._ZNK5QRect7topLeftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "topLeft", args)
  }

}

// setBottom(int)
func (this *QRect) setBottom(args ...interface{}) () {
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
    C._ZN5QRect9setBottomEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setBottom", args)
  }

}

// setTopLeft(const class QPoint &)
func (this *QRect) setTopLeft(args ...interface{}) () {
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
    C._ZN5QRect10setTopLeftERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setTopLeft", args)
  }

}

// adjusted(int, int, int, int)
func (this *QRect) adjusted(args ...interface{}) () {
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
    C._ZNK5QRect8adjustedEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRect", "adjusted", args)
  }

}

// height()
func (this *QRect) height(args ...interface{}) () {
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
    C._ZNK5QRect6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "height", args)
  }

}

// setBottomLeft(const class QPoint &)
func (this *QRect) setBottomLeft(args ...interface{}) () {
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
    C._ZN5QRect13setBottomLeftERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setBottomLeft", args)
  }

}

// moveRight(int)
func (this *QRect) moveRight(args ...interface{}) () {
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
    C._ZN5QRect9moveRightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveRight", args)
  }

}

// moveBottomLeft(const class QPoint &)
func (this *QRect) moveBottomLeft(args ...interface{}) () {
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
    C._ZN5QRect14moveBottomLeftERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveBottomLeft", args)
  }

}

// setWidth(int)
func (this *QRect) setWidth(args ...interface{}) () {
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
    C._ZN5QRect8setWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setWidth", args)
  }

}

// marginsAdded(const class QMargins &)
func (this *QRect) marginsAdded(args ...interface{}) () {
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
    C._ZNK5QRect12marginsAddedERK8QMargins(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "marginsAdded", args)
  }

}

// size()
func (this *QRect) size(args ...interface{}) () {
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
    C._ZNK5QRect4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "size", args)
  }

}

// moveBottom(int)
func (this *QRect) moveBottom(args ...interface{}) () {
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
    C._ZN5QRect10moveBottomEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveBottom", args)
  }

}

// united(const class QRect &)
func (this *QRect) united(args ...interface{}) () {
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
    C._ZNK5QRect6unitedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "united", args)
  }

}

// bottomRight()
func (this *QRect) bottomRight(args ...interface{}) () {
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
    C._ZNK5QRect11bottomRightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "bottomRight", args)
  }

}

// intersected(const class QRect &)
func (this *QRect) intersected(args ...interface{}) () {
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
    C._ZNK5QRect11intersectedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "intersected", args)
  }

}

// getRect(int *, int *, int *, int *)
func (this *QRect) getRect(args ...interface{}) () {
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
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C._ZNK5QRect7getRectEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRect", "getRect", args)
  }

}

// top()
func (this *QRect) top(args ...interface{}) () {
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
    C._ZNK5QRect3topEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "top", args)
  }

}

// contains(int, int, _Bool)
func (this *QRect) contains(args ...interface{}) () {
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
    C._ZNK5QRect8containsEiib(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK5QRect8containsERK6QPointb
    // invoke: bool contains(const class QPoint &, _Bool)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZNK5QRect8containsERK6QPointb(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK5QRect8containsERKS_b
    // invoke: bool contains(const class QRect &, _Bool)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZNK5QRect8containsERKS_b(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZNK5QRect8containsEii
    // invoke: bool contains(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK5QRect8containsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRect", "contains", args)
  }

}

// setTop(int)
func (this *QRect) setTop(args ...interface{}) () {
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
    C._ZN5QRect6setTopEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setTop", args)
  }

}

// topRight()
func (this *QRect) topRight(args ...interface{}) () {
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
    C._ZNK5QRect8topRightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "topRight", args)
  }

}

// moveCenter(const class QPoint &)
func (this *QRect) moveCenter(args ...interface{}) () {
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
    C._ZN5QRect10moveCenterERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveCenter", args)
  }

}

// width()
func (this *QRect) width(args ...interface{}) () {
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
    C._ZNK5QRect5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "width", args)
  }

}

// isEmpty()
func (this *QRect) isEmpty(args ...interface{}) () {
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
    C._ZNK5QRect7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "isEmpty", args)
  }

}

// setX(int)
func (this *QRect) setX(args ...interface{}) () {
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
    C._ZN5QRect4setXEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setX", args)
  }

}

// setRect(int, int, int, int)
func (this *QRect) setRect(args ...interface{}) () {
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
    C._ZN5QRect7setRectEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRect", "setRect", args)
  }

}

// setBottomRight(const class QPoint &)
func (this *QRect) setBottomRight(args ...interface{}) () {
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
    C._ZN5QRect14setBottomRightERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setBottomRight", args)
  }

}

// isValid()
func (this *QRect) isValid(args ...interface{}) () {
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
    C._ZNK5QRect7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "isValid", args)
  }

}

// moveLeft(int)
func (this *QRect) moveLeft(args ...interface{}) () {
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
    C._ZN5QRect8moveLeftEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveLeft", args)
  }

}

// setTopRight(const class QPoint &)
func (this *QRect) setTopRight(args ...interface{}) () {
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
    C._ZN5QRect11setTopRightERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setTopRight", args)
  }

}

// moveTop(int)
func (this *QRect) moveTop(args ...interface{}) () {
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
    C._ZN5QRect7moveTopEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveTop", args)
  }

}

// translated(const class QPoint &)
func (this *QRect) translated(args ...interface{}) () {
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
    C._ZNK5QRect10translatedERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK5QRect10translatedEii
    // invoke: QRect translated(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK5QRect10translatedEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRect", "translated", args)
  }

}

// normalized()
func (this *QRect) normalized(args ...interface{}) () {
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
    C._ZNK5QRect10normalizedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "normalized", args)
  }

}

// setCoords(int, int, int, int)
func (this *QRect) setCoords(args ...interface{}) () {
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
    C._ZN5QRect9setCoordsEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRect", "setCoords", args)
  }

}

// setLeft(int)
func (this *QRect) setLeft(args ...interface{}) () {
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
    C._ZN5QRect7setLeftEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setLeft", args)
  }

}

// setY(int)
func (this *QRect) setY(args ...interface{}) () {
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
    C._ZN5QRect4setYEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setY", args)
  }

}

// setHeight(int)
func (this *QRect) setHeight(args ...interface{}) () {
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
    C._ZN5QRect9setHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setHeight", args)
  }

}

// QRect(const class QPoint &, const class QSize &)
func NewQRect(args ...interface{}) QRect {
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
    C._ZN5QRectC2ERK6QPointRK5QSize(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN5QRectC1Ev
    // invoke: void QRect()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QRectC2Ev(qthis)
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
    C._ZN5QRectC2Eiiii(qthis, arg0, arg1, arg2, arg3)
  case 3:
    // invoke: _ZN5QRectC1ERK6QPointS2_
    // invoke: void QRect(const class QPoint &, const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QRectC2ERK6QPointS2_(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRect", "QRect", args)
  }

  return QRect{}
}

// translate(const class QPoint &)
func (this *QRect) translate(args ...interface{}) () {
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
    C._ZN5QRect9translateERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN5QRect9translateEii
    // invoke: void translate(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN5QRect9translateEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRect", "translate", args)
  }

}

// bottomLeft()
func (this *QRect) bottomLeft(args ...interface{}) () {
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
    C._ZNK5QRect10bottomLeftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "bottomLeft", args)
  }

}

// center()
func (this *QRect) center(args ...interface{}) () {
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
    C._ZNK5QRect6centerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "center", args)
  }

}

// moveTopRight(const class QPoint &)
func (this *QRect) moveTopRight(args ...interface{}) () {
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
    C._ZN5QRect12moveTopRightERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveTopRight", args)
  }

}

// moveTopLeft(const class QPoint &)
func (this *QRect) moveTopLeft(args ...interface{}) () {
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
    C._ZN5QRect11moveTopLeftERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveTopLeft", args)
  }

}

// bottom()
func (this *QRect) bottom(args ...interface{}) () {
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
    C._ZNK5QRect6bottomEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "bottom", args)
  }

}

// getCoords(int *, int *, int *, int *)
func (this *QRect) getCoords(args ...interface{}) () {
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
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C._ZNK5QRect9getCoordsEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRect", "getCoords", args)
  }

}

// setRight(int)
func (this *QRect) setRight(args ...interface{}) () {
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
    C._ZN5QRect8setRightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setRight", args)
  }

}

// isNull()
func (this *QRect) isNull(args ...interface{}) () {
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
    C._ZNK5QRect6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "isNull", args)
  }

}

// adjust(int, int, int, int)
func (this *QRect) adjust(args ...interface{}) () {
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
    C._ZN5QRect6adjustEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRect", "adjust", args)
  }

}

// y()
func (this *QRect) y(args ...interface{}) () {
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
    C._ZNK5QRect1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "y", args)
  }

}

// x()
func (this *QRect) x(args ...interface{}) () {
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
    C._ZNK5QRect1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "x", args)
  }

}

// moveBottomRight(const class QPoint &)
func (this *QRect) moveBottomRight(args ...interface{}) () {
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
    C._ZN5QRect15moveBottomRightERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "moveBottomRight", args)
  }

}

// setSize(const class QSize &)
func (this *QRect) setSize(args ...interface{}) () {
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
    C._ZN5QRect7setSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRect", "setSize", args)
  }

}

// left()
func (this *QRect) left(args ...interface{}) () {
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
    C._ZNK5QRect4leftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRect", "left", args)
  }

}

// marginsRemoved(const class QMarginsF &)
func (this *QRectF) marginsRemoved(args ...interface{}) () {
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
    C._ZNK6QRectF14marginsRemovedERK9QMarginsF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "marginsRemoved", args)
  }

}

// moveTo(const class QPointF &)
func (this *QRectF) moveTo(args ...interface{}) () {
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
    C._ZN6QRectF6moveToERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN6QRectF6moveToEdd
    // invoke: void moveTo(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN6QRectF6moveToEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRectF", "moveTo", args)
  }

}

// setBottomLeft(const class QPointF &)
func (this *QRectF) setBottomLeft(args ...interface{}) () {
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
    C._ZN6QRectF13setBottomLeftERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setBottomLeft", args)
  }

}

// intersects(const class QRectF &)
func (this *QRectF) intersects(args ...interface{}) () {
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
    C._ZNK6QRectF10intersectsERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "intersects", args)
  }

}

// topLeft()
func (this *QRectF) topLeft(args ...interface{}) () {
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
    C._ZNK6QRectF7topLeftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "topLeft", args)
  }

}

// setBottom(qreal)
func (this *QRectF) setBottom(args ...interface{}) () {
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
    C._ZN6QRectF9setBottomEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setBottom", args)
  }

}

// setTopLeft(const class QPointF &)
func (this *QRectF) setTopLeft(args ...interface{}) () {
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
    C._ZN6QRectF10setTopLeftERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setTopLeft", args)
  }

}

// adjusted(qreal, qreal, qreal, qreal)
func (this *QRectF) adjusted(args ...interface{}) () {
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
    C._ZNK6QRectF8adjustedEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRectF", "adjusted", args)
  }

}

// height()
func (this *QRectF) height(args ...interface{}) () {
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
    C._ZNK6QRectF6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "height", args)
  }

}

// right()
func (this *QRectF) right(args ...interface{}) () {
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
    C._ZNK6QRectF5rightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "right", args)
  }

}

// setLeft(qreal)
func (this *QRectF) setLeft(args ...interface{}) () {
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
    C._ZN6QRectF7setLeftEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setLeft", args)
  }

}

// moveRight(qreal)
func (this *QRectF) moveRight(args ...interface{}) () {
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
    C._ZN6QRectF9moveRightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveRight", args)
  }

}

// moveBottomLeft(const class QPointF &)
func (this *QRectF) moveBottomLeft(args ...interface{}) () {
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
    C._ZN6QRectF14moveBottomLeftERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveBottomLeft", args)
  }

}

// setWidth(qreal)
func (this *QRectF) setWidth(args ...interface{}) () {
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
    C._ZN6QRectF8setWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setWidth", args)
  }

}

// marginsAdded(const class QMarginsF &)
func (this *QRectF) marginsAdded(args ...interface{}) () {
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
    C._ZNK6QRectF12marginsAddedERK9QMarginsF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "marginsAdded", args)
  }

}

// moveTop(qreal)
func (this *QRectF) moveTop(args ...interface{}) () {
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
    C._ZN6QRectF7moveTopEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveTop", args)
  }

}

// bottomRight()
func (this *QRectF) bottomRight(args ...interface{}) () {
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
    C._ZNK6QRectF11bottomRightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "bottomRight", args)
  }

}

// bottom()
func (this *QRectF) bottom(args ...interface{}) () {
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
    C._ZNK6QRectF6bottomEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "bottom", args)
  }

}

// getRect(qreal *, qreal *, qreal *, qreal *)
func (this *QRectF) getRect(args ...interface{}) () {
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
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C._ZNK6QRectF7getRectEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRectF", "getRect", args)
  }

}

// top()
func (this *QRectF) top(args ...interface{}) () {
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
    C._ZNK6QRectF3topEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "top", args)
  }

}

// contains(qreal, qreal)
func (this *QRectF) contains(args ...interface{}) () {
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
    C._ZNK6QRectF8containsEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK6QRectF8containsERKS_
    // invoke: bool contains(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK6QRectF8containsERKS_(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK6QRectF8containsERK7QPointF
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK6QRectF8containsERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "contains", args)
  }

}

// toRect()
func (this *QRectF) toRect(args ...interface{}) () {
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
    C._ZNK6QRectF6toRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "toRect", args)
  }

}

// topRight()
func (this *QRectF) topRight(args ...interface{}) () {
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
    C._ZNK6QRectF8topRightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "topRight", args)
  }

}

// moveCenter(const class QPointF &)
func (this *QRectF) moveCenter(args ...interface{}) () {
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
    C._ZN6QRectF10moveCenterERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveCenter", args)
  }

}

// width()
func (this *QRectF) width(args ...interface{}) () {
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
    C._ZNK6QRectF5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "width", args)
  }

}

// isEmpty()
func (this *QRectF) isEmpty(args ...interface{}) () {
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
    C._ZNK6QRectF7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "isEmpty", args)
  }

}

// QRectF(const class QPointF &, const class QPointF &)
func NewQRectF(args ...interface{}) QRectF {
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
    C._ZN6QRectFC2ERK7QPointFS2_(qthis, arg0, arg1)
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
    C._ZN6QRectFC2Edddd(qthis, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN6QRectFC1Ev
    // invoke: void QRectF()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QRectFC2Ev(qthis)
  case 3:
    // invoke: _ZN6QRectFC1ERK7QPointFRK6QSizeF
    // invoke: void QRectF(const class QPointF &, const class QSizeF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QSizeF).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QRectFC2ERK7QPointFRK6QSizeF(qthis, arg0, arg1)
  case 4:
    // invoke: _ZN6QRectFC1ERK5QRect
    // invoke: void QRectF(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QRectFC2ERK5QRect(qthis, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "QRectF", args)
  }

  return QRectF{}
}

// size()
func (this *QRectF) size(args ...interface{}) () {
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
    C._ZNK6QRectF4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "size", args)
  }

}

// setX(qreal)
func (this *QRectF) setX(args ...interface{}) () {
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
    C._ZN6QRectF4setXEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setX", args)
  }

}

// setRect(qreal, qreal, qreal, qreal)
func (this *QRectF) setRect(args ...interface{}) () {
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
    C._ZN6QRectF7setRectEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRectF", "setRect", args)
  }

}

// moveLeft(qreal)
func (this *QRectF) moveLeft(args ...interface{}) () {
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
    C._ZN6QRectF8moveLeftEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveLeft", args)
  }

}

// setBottomRight(const class QPointF &)
func (this *QRectF) setBottomRight(args ...interface{}) () {
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
    C._ZN6QRectF14setBottomRightERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setBottomRight", args)
  }

}

// isValid()
func (this *QRectF) isValid(args ...interface{}) () {
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
    C._ZNK6QRectF7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "isValid", args)
  }

}

// getCoords(qreal *, qreal *, qreal *, qreal *)
func (this *QRectF) getCoords(args ...interface{}) () {
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
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C._ZNK6QRectF9getCoordsEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRectF", "getCoords", args)
  }

}

// setTopRight(const class QPointF &)
func (this *QRectF) setTopRight(args ...interface{}) () {
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
    C._ZN6QRectF11setTopRightERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setTopRight", args)
  }

}

// translated(qreal, qreal)
func (this *QRectF) translated(args ...interface{}) () {
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
    C._ZNK6QRectF10translatedEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK6QRectF10translatedERK7QPointF
    // invoke: QRectF translated(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK6QRectF10translatedERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "translated", args)
  }

}

// normalized()
func (this *QRectF) normalized(args ...interface{}) () {
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
    C._ZNK6QRectF10normalizedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "normalized", args)
  }

}

// setTop(qreal)
func (this *QRectF) setTop(args ...interface{}) () {
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
    C._ZN6QRectF6setTopEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setTop", args)
  }

}

// moveBottom(qreal)
func (this *QRectF) moveBottom(args ...interface{}) () {
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
    C._ZN6QRectF10moveBottomEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveBottom", args)
  }

}

// setY(qreal)
func (this *QRectF) setY(args ...interface{}) () {
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
    C._ZN6QRectF4setYEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setY", args)
  }

}

// setCoords(qreal, qreal, qreal, qreal)
func (this *QRectF) setCoords(args ...interface{}) () {
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
    C._ZN6QRectF9setCoordsEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRectF", "setCoords", args)
  }

}

// setHeight(qreal)
func (this *QRectF) setHeight(args ...interface{}) () {
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
    C._ZN6QRectF9setHeightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setHeight", args)
  }

}

// translate(const class QPointF &)
func (this *QRectF) translate(args ...interface{}) () {
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
    C._ZN6QRectF9translateERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN6QRectF9translateEdd
    // invoke: void translate(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN6QRectF9translateEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRectF", "translate", args)
  }

}

// bottomLeft()
func (this *QRectF) bottomLeft(args ...interface{}) () {
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
    C._ZNK6QRectF10bottomLeftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "bottomLeft", args)
  }

}

// center()
func (this *QRectF) center(args ...interface{}) () {
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
    C._ZNK6QRectF6centerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "center", args)
  }

}

// moveTopRight(const class QPointF &)
func (this *QRectF) moveTopRight(args ...interface{}) () {
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
    C._ZN6QRectF12moveTopRightERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveTopRight", args)
  }

}

// moveTopLeft(const class QPointF &)
func (this *QRectF) moveTopLeft(args ...interface{}) () {
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
    C._ZN6QRectF11moveTopLeftERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveTopLeft", args)
  }

}

// toAlignedRect()
func (this *QRectF) toAlignedRect(args ...interface{}) () {
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
    C._ZNK6QRectF13toAlignedRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "toAlignedRect", args)
  }

}

// intersected(const class QRectF &)
func (this *QRectF) intersected(args ...interface{}) () {
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
    C._ZNK6QRectF11intersectedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "intersected", args)
  }

}

// united(const class QRectF &)
func (this *QRectF) united(args ...interface{}) () {
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
    C._ZNK6QRectF6unitedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "united", args)
  }

}

// setRight(qreal)
func (this *QRectF) setRight(args ...interface{}) () {
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
    C._ZN6QRectF8setRightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setRight", args)
  }

}

// isNull()
func (this *QRectF) isNull(args ...interface{}) () {
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
    C._ZNK6QRectF6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "isNull", args)
  }

}

// adjust(qreal, qreal, qreal, qreal)
func (this *QRectF) adjust(args ...interface{}) () {
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
    C._ZN6QRectF6adjustEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRectF", "adjust", args)
  }

}

// y()
func (this *QRectF) y(args ...interface{}) () {
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
    C._ZNK6QRectF1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "y", args)
  }

}

// x()
func (this *QRectF) x(args ...interface{}) () {
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
    C._ZNK6QRectF1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "x", args)
  }

}

// left()
func (this *QRectF) left(args ...interface{}) () {
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
    C._ZNK6QRectF4leftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRectF", "left", args)
  }

}

// setSize(const class QSizeF &)
func (this *QRectF) setSize(args ...interface{}) () {
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
    C._ZN6QRectF7setSizeERK6QSizeF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "setSize", args)
  }

}

// moveBottomRight(const class QPointF &)
func (this *QRectF) moveBottomRight(args ...interface{}) () {
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
    C._ZN6QRectF15moveBottomRightERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRectF", "moveBottomRight", args)
  }

}

// <= body block end

