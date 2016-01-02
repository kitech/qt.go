package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qgraphicsitem.h
// dst-file: /src/widgets/qgraphicsitem.go
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
  // proto:  bool QGraphicsTextItem::openExternalLinks();
extern void _ZNK17QGraphicsTextItem17openExternalLinksEv(void* qthis);
  // proto:  qreal QGraphicsTextItem::textWidth();
extern void _ZNK17QGraphicsTextItem9textWidthEv(void* qthis);
  // proto:  void QGraphicsTextItem::setTextWidth(qreal width);
extern void _ZN17QGraphicsTextItem12setTextWidthEd(void* qthis, double arg0);
  // proto:  void QGraphicsTextItem::setTextCursor(const QTextCursor & cursor);
extern void _ZN17QGraphicsTextItem13setTextCursorERK11QTextCursor(void* qthis, void* arg0);
  // proto:  int QGraphicsTextItem::type();
extern void _ZNK17QGraphicsTextItem4typeEv(void* qthis);
  // proto:  QFont QGraphicsTextItem::font();
extern void _ZNK17QGraphicsTextItem4fontEv(void* qthis);
  // proto:  void QGraphicsTextItem::QGraphicsTextItem(const QString & text, QGraphicsItem * parent);
extern void* dector_ZN17QGraphicsTextItemC1ERK7QStringP13QGraphicsItem(void* arg0, void* arg1);
extern void _ZN17QGraphicsTextItemC1ERK7QStringP13QGraphicsItem(void* qthis, void* arg0, void* arg1);
  // proto:  const QMetaObject * QGraphicsTextItem::metaObject();
extern void _ZNK17QGraphicsTextItem10metaObjectEv(void* qthis);
  // proto:  void QGraphicsTextItem::setOpenExternalLinks(bool open);
extern void _ZN17QGraphicsTextItem20setOpenExternalLinksEb(void* qthis, bool arg0);
  // proto:  void QGraphicsTextItem::setTabChangesFocus(bool b);
extern void _ZN17QGraphicsTextItem18setTabChangesFocusEb(void* qthis, bool arg0);
  // proto:  QString QGraphicsTextItem::toHtml();
extern void _ZNK17QGraphicsTextItem6toHtmlEv(void* qthis);
  // proto:  void QGraphicsTextItem::setDocument(QTextDocument * document);
extern void _ZN17QGraphicsTextItem11setDocumentEP13QTextDocument(void* qthis, void* arg0);
  // proto:  void QGraphicsTextItem::setPlainText(const QString & text);
extern void _ZN17QGraphicsTextItem12setPlainTextERK7QString(void* qthis, void* arg0);
  // proto:  void QGraphicsTextItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN17QGraphicsTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QGraphicsTextItem::setFont(const QFont & font);
extern void _ZN17QGraphicsTextItem7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  void QGraphicsTextItem::setDefaultTextColor(const QColor & c);
extern void _ZN17QGraphicsTextItem19setDefaultTextColorERK6QColor(void* qthis, void* arg0);
  // proto:  QColor QGraphicsTextItem::defaultTextColor();
extern void _ZNK17QGraphicsTextItem16defaultTextColorEv(void* qthis);
  // proto:  void QGraphicsTextItem::~QGraphicsTextItem();
extern void _ZN17QGraphicsTextItemD0Ev(void* qthis);
  // proto:  QPainterPath QGraphicsTextItem::shape();
extern void _ZNK17QGraphicsTextItem5shapeEv(void* qthis);
  // proto:  QTextCursor QGraphicsTextItem::textCursor();
extern void _ZNK17QGraphicsTextItem10textCursorEv(void* qthis);
  // proto:  QRectF QGraphicsTextItem::boundingRect();
extern void _ZNK17QGraphicsTextItem12boundingRectEv(void* qthis);
  // proto:  QString QGraphicsTextItem::toPlainText();
extern void _ZNK17QGraphicsTextItem11toPlainTextEv(void* qthis);
  // proto:  void QGraphicsTextItem::setHtml(const QString & html);
extern void _ZN17QGraphicsTextItem7setHtmlERK7QString(void* qthis, void* arg0);
  // proto:  bool QGraphicsTextItem::tabChangesFocus();
extern void _ZNK17QGraphicsTextItem15tabChangesFocusEv(void* qthis);
  // proto:  void QGraphicsTextItem::QGraphicsTextItem(const QGraphicsTextItem & );
extern void* dector_ZN17QGraphicsTextItemC1ERKS_(void* arg0);
extern void _ZN17QGraphicsTextItemC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsTextItem::QGraphicsTextItem(QGraphicsItem * parent);
extern void* dector_ZN17QGraphicsTextItemC1EP13QGraphicsItem(void* arg0);
extern void _ZN17QGraphicsTextItemC1EP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QTextDocument * QGraphicsTextItem::document();
extern void _ZNK17QGraphicsTextItem8documentEv(void* qthis);
  // proto:  bool QGraphicsTextItem::isObscuredBy(const QGraphicsItem * item);
extern void _ZNK17QGraphicsTextItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsTextItem::opaqueArea();
extern void _ZNK17QGraphicsTextItem10opaqueAreaEv(void* qthis);
  // proto:  bool QGraphicsTextItem::contains(const QPointF & point);
extern void _ZNK17QGraphicsTextItem8containsERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsTextItem::adjustSize();
extern void _ZN17QGraphicsTextItem10adjustSizeEv(void* qthis);
  // proto:  void QGraphicsPixmapItem::QGraphicsPixmapItem(QGraphicsItem * parent);
extern void* dector_ZN19QGraphicsPixmapItemC1EP13QGraphicsItem(void* arg0);
extern void _ZN19QGraphicsPixmapItemC1EP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsPixmapItem::QGraphicsPixmapItem(const QPixmap & pixmap, QGraphicsItem * parent);
extern void* dector_ZN19QGraphicsPixmapItemC1ERK7QPixmapP13QGraphicsItem(void* arg0, void* arg1);
extern void _ZN19QGraphicsPixmapItemC1ERK7QPixmapP13QGraphicsItem(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsPixmapItem::~QGraphicsPixmapItem();
extern void _ZN19QGraphicsPixmapItemD0Ev(void* qthis);
  // proto:  QPainterPath QGraphicsPixmapItem::opaqueArea();
extern void _ZNK19QGraphicsPixmapItem10opaqueAreaEv(void* qthis);
  // proto:  bool QGraphicsPixmapItem::isObscuredBy(const QGraphicsItem * item);
extern void _ZNK19QGraphicsPixmapItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0);
  // proto:  int QGraphicsPixmapItem::type();
extern void _ZNK19QGraphicsPixmapItem4typeEv(void* qthis);
  // proto:  QPainterPath QGraphicsPixmapItem::shape();
extern void _ZNK19QGraphicsPixmapItem5shapeEv(void* qthis);
  // proto:  QPixmap QGraphicsPixmapItem::pixmap();
extern void _ZNK19QGraphicsPixmapItem6pixmapEv(void* qthis);
  // proto:  void QGraphicsPixmapItem::setOffset(qreal x, qreal y);
extern void demth_ZN19QGraphicsPixmapItem9setOffsetEdd(void* qthis, double arg0, double arg1);
  // proto:  void QGraphicsPixmapItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN19QGraphicsPixmapItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QGraphicsPixmapItem::QGraphicsPixmapItem(const QGraphicsPixmapItem & );
extern void* dector_ZN19QGraphicsPixmapItemC1ERKS_(void* arg0);
extern void _ZN19QGraphicsPixmapItemC1ERKS_(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsPixmapItem::offset();
extern void _ZNK19QGraphicsPixmapItem6offsetEv(void* qthis);
  // proto:  QRectF QGraphicsPixmapItem::boundingRect();
extern void _ZNK19QGraphicsPixmapItem12boundingRectEv(void* qthis);
  // proto:  bool QGraphicsPixmapItem::contains(const QPointF & point);
extern void _ZNK19QGraphicsPixmapItem8containsERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsPixmapItem::setPixmap(const QPixmap & pixmap);
extern void _ZN19QGraphicsPixmapItem9setPixmapERK7QPixmap(void* qthis, void* arg0);
  // proto:  void QGraphicsPixmapItem::setOffset(const QPointF & offset);
extern void _ZN19QGraphicsPixmapItem9setOffsetERK7QPointF(void* qthis, void* arg0);
  // proto:  bool QGraphicsRectItem::isObscuredBy(const QGraphicsItem * item);
extern void _ZNK17QGraphicsRectItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QRectF QGraphicsRectItem::boundingRect();
extern void _ZNK17QGraphicsRectItem12boundingRectEv(void* qthis);
  // proto:  void QGraphicsRectItem::QGraphicsRectItem(const QGraphicsRectItem & );
extern void* dector_ZN17QGraphicsRectItemC1ERKS_(void* arg0);
extern void _ZN17QGraphicsRectItemC1ERKS_(void* qthis, void* arg0);
  // proto:  int QGraphicsRectItem::type();
extern void _ZNK17QGraphicsRectItem4typeEv(void* qthis);
  // proto:  QRectF QGraphicsRectItem::rect();
extern void _ZNK17QGraphicsRectItem4rectEv(void* qthis);
  // proto:  QPainterPath QGraphicsRectItem::shape();
extern void _ZNK17QGraphicsRectItem5shapeEv(void* qthis);
  // proto:  void QGraphicsRectItem::~QGraphicsRectItem();
extern void _ZN17QGraphicsRectItemD0Ev(void* qthis);
  // proto:  void QGraphicsRectItem::QGraphicsRectItem(const QRectF & rect, QGraphicsItem * parent);
extern void* dector_ZN17QGraphicsRectItemC1ERK6QRectFP13QGraphicsItem(void* arg0, void* arg1);
extern void _ZN17QGraphicsRectItemC1ERK6QRectFP13QGraphicsItem(void* qthis, void* arg0, void* arg1);
  // proto:  QPainterPath QGraphicsRectItem::opaqueArea();
extern void _ZNK17QGraphicsRectItem10opaqueAreaEv(void* qthis);
  // proto:  void QGraphicsRectItem::setRect(const QRectF & rect);
extern void _ZN17QGraphicsRectItem7setRectERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsRectItem::setRect(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZN17QGraphicsRectItem7setRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsRectItem::QGraphicsRectItem(QGraphicsItem * parent);
extern void* dector_ZN17QGraphicsRectItemC1EP13QGraphicsItem(void* arg0);
extern void _ZN17QGraphicsRectItemC1EP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  bool QGraphicsRectItem::contains(const QPointF & point);
extern void _ZNK17QGraphicsRectItem8containsERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsRectItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QGraphicsRectItem::QGraphicsRectItem(qreal x, qreal y, qreal w, qreal h, QGraphicsItem * parent);
extern void* dector_ZN17QGraphicsRectItemC1EddddP13QGraphicsItem(double arg0, double arg1, double arg2, double arg3, void* arg4);
extern void _ZN17QGraphicsRectItemC1EddddP13QGraphicsItem(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4);
  // proto:  void QGraphicsEllipseItem::setStartAngle(int angle);
extern void _ZN20QGraphicsEllipseItem13setStartAngleEi(void* qthis, int arg0);
  // proto:  void QGraphicsEllipseItem::QGraphicsEllipseItem(const QGraphicsEllipseItem & );
extern void* dector_ZN20QGraphicsEllipseItemC1ERKS_(void* arg0);
extern void _ZN20QGraphicsEllipseItemC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QGraphicsEllipseItem::contains(const QPointF & point);
extern void _ZNK20QGraphicsEllipseItem8containsERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsEllipseItem::QGraphicsEllipseItem(const QRectF & rect, QGraphicsItem * parent);
extern void* dector_ZN20QGraphicsEllipseItemC1ERK6QRectFP13QGraphicsItem(void* arg0, void* arg1);
extern void _ZN20QGraphicsEllipseItemC1ERK6QRectFP13QGraphicsItem(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsEllipseItem::setRect(const QRectF & rect);
extern void _ZN20QGraphicsEllipseItem7setRectERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsEllipseItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN20QGraphicsEllipseItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  bool QGraphicsEllipseItem::isObscuredBy(const QGraphicsItem * item);
extern void _ZNK20QGraphicsEllipseItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QRectF QGraphicsEllipseItem::rect();
extern void _ZNK20QGraphicsEllipseItem4rectEv(void* qthis);
  // proto:  int QGraphicsEllipseItem::spanAngle();
extern void _ZNK20QGraphicsEllipseItem9spanAngleEv(void* qthis);
  // proto:  int QGraphicsEllipseItem::startAngle();
extern void _ZNK20QGraphicsEllipseItem10startAngleEv(void* qthis);
  // proto:  void QGraphicsEllipseItem::QGraphicsEllipseItem(qreal x, qreal y, qreal w, qreal h, QGraphicsItem * parent);
extern void* dector_ZN20QGraphicsEllipseItemC1EddddP13QGraphicsItem(double arg0, double arg1, double arg2, double arg3, void* arg4);
extern void _ZN20QGraphicsEllipseItemC1EddddP13QGraphicsItem(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4);
  // proto:  void QGraphicsEllipseItem::setRect(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZN20QGraphicsEllipseItem7setRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsEllipseItem::setSpanAngle(int angle);
extern void _ZN20QGraphicsEllipseItem12setSpanAngleEi(void* qthis, int arg0);
  // proto:  int QGraphicsEllipseItem::type();
extern void _ZNK20QGraphicsEllipseItem4typeEv(void* qthis);
  // proto:  QRectF QGraphicsEllipseItem::boundingRect();
extern void _ZNK20QGraphicsEllipseItem12boundingRectEv(void* qthis);
  // proto:  QPainterPath QGraphicsEllipseItem::shape();
extern void _ZNK20QGraphicsEllipseItem5shapeEv(void* qthis);
  // proto:  void QGraphicsEllipseItem::~QGraphicsEllipseItem();
extern void _ZN20QGraphicsEllipseItemD0Ev(void* qthis);
  // proto:  void QGraphicsEllipseItem::QGraphicsEllipseItem(QGraphicsItem * parent);
extern void* dector_ZN20QGraphicsEllipseItemC1EP13QGraphicsItem(void* arg0);
extern void _ZN20QGraphicsEllipseItemC1EP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsEllipseItem::opaqueArea();
extern void _ZNK20QGraphicsEllipseItem10opaqueAreaEv(void* qthis);
  // proto:  QPainterPath QGraphicsPolygonItem::shape();
extern void _ZNK20QGraphicsPolygonItem5shapeEv(void* qthis);
  // proto:  bool QGraphicsPolygonItem::isObscuredBy(const QGraphicsItem * item);
extern void _ZNK20QGraphicsPolygonItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsPolygonItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN20QGraphicsPolygonItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QGraphicsPolygonItem::QGraphicsPolygonItem(QGraphicsItem * parent);
extern void* dector_ZN20QGraphicsPolygonItemC1EP13QGraphicsItem(void* arg0);
extern void _ZN20QGraphicsPolygonItemC1EP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QRectF QGraphicsPolygonItem::boundingRect();
extern void _ZNK20QGraphicsPolygonItem12boundingRectEv(void* qthis);
  // proto:  int QGraphicsPolygonItem::type();
extern void _ZNK20QGraphicsPolygonItem4typeEv(void* qthis);
  // proto:  void QGraphicsPolygonItem::~QGraphicsPolygonItem();
extern void _ZN20QGraphicsPolygonItemD0Ev(void* qthis);
  // proto:  QPolygonF QGraphicsPolygonItem::polygon();
extern void _ZNK20QGraphicsPolygonItem7polygonEv(void* qthis);
  // proto:  void QGraphicsPolygonItem::QGraphicsPolygonItem(const QGraphicsPolygonItem & );
extern void* dector_ZN20QGraphicsPolygonItemC1ERKS_(void* arg0);
extern void _ZN20QGraphicsPolygonItemC1ERKS_(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsPolygonItem::opaqueArea();
extern void _ZNK20QGraphicsPolygonItem10opaqueAreaEv(void* qthis);
  // proto:  void QGraphicsPolygonItem::QGraphicsPolygonItem(const QPolygonF & polygon, QGraphicsItem * parent);
extern void* dector_ZN20QGraphicsPolygonItemC1ERK9QPolygonFP13QGraphicsItem(void* arg0, void* arg1);
extern void _ZN20QGraphicsPolygonItemC1ERK9QPolygonFP13QGraphicsItem(void* qthis, void* arg0, void* arg1);
  // proto:  bool QGraphicsPolygonItem::contains(const QPointF & point);
extern void _ZNK20QGraphicsPolygonItem8containsERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsPolygonItem::setPolygon(const QPolygonF & polygon);
extern void _ZN20QGraphicsPolygonItem10setPolygonERK9QPolygonF(void* qthis, void* arg0);
  // proto:  void QGraphicsPathItem::setPath(const QPainterPath & path);
extern void _ZN17QGraphicsPathItem7setPathERK12QPainterPath(void* qthis, void* arg0);
  // proto:  void QGraphicsPathItem::QGraphicsPathItem(const QPainterPath & path, QGraphicsItem * parent);
extern void* dector_ZN17QGraphicsPathItemC1ERK12QPainterPathP13QGraphicsItem(void* arg0, void* arg1);
extern void _ZN17QGraphicsPathItemC1ERK12QPainterPathP13QGraphicsItem(void* qthis, void* arg0, void* arg1);
  // proto:  bool QGraphicsPathItem::contains(const QPointF & point);
extern void _ZNK17QGraphicsPathItem8containsERK7QPointF(void* qthis, void* arg0);
  // proto:  QRectF QGraphicsPathItem::boundingRect();
extern void _ZNK17QGraphicsPathItem12boundingRectEv(void* qthis);
  // proto:  void QGraphicsPathItem::QGraphicsPathItem(const QGraphicsPathItem & );
extern void* dector_ZN17QGraphicsPathItemC1ERKS_(void* arg0);
extern void _ZN17QGraphicsPathItemC1ERKS_(void* qthis, void* arg0);
  // proto:  int QGraphicsPathItem::type();
extern void _ZNK17QGraphicsPathItem4typeEv(void* qthis);
  // proto:  QPainterPath QGraphicsPathItem::opaqueArea();
extern void _ZNK17QGraphicsPathItem10opaqueAreaEv(void* qthis);
  // proto:  QPainterPath QGraphicsPathItem::path();
extern void _ZNK17QGraphicsPathItem4pathEv(void* qthis);
  // proto:  void QGraphicsPathItem::~QGraphicsPathItem();
extern void _ZN17QGraphicsPathItemD0Ev(void* qthis);
  // proto:  QPainterPath QGraphicsPathItem::shape();
extern void _ZNK17QGraphicsPathItem5shapeEv(void* qthis);
  // proto:  bool QGraphicsPathItem::isObscuredBy(const QGraphicsItem * item);
extern void _ZNK17QGraphicsPathItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsPathItem::QGraphicsPathItem(QGraphicsItem * parent);
extern void* dector_ZN17QGraphicsPathItemC1EP13QGraphicsItem(void* arg0);
extern void _ZN17QGraphicsPathItemC1EP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsPathItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN17QGraphicsPathItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QGraphicsLineItem::setPen(const QPen & pen);
extern void _ZN17QGraphicsLineItem6setPenERK4QPen(void* qthis, void* arg0);
  // proto:  void QGraphicsLineItem::QGraphicsLineItem(QGraphicsItem * parent);
extern void* dector_ZN17QGraphicsLineItemC1EP13QGraphicsItem(void* arg0);
extern void _ZN17QGraphicsLineItemC1EP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  bool QGraphicsLineItem::isObscuredBy(const QGraphicsItem * item);
extern void _ZNK17QGraphicsLineItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsLineItem::QGraphicsLineItem(const QLineF & line, QGraphicsItem * parent);
extern void* dector_ZN17QGraphicsLineItemC1ERK6QLineFP13QGraphicsItem(void* arg0, void* arg1);
extern void _ZN17QGraphicsLineItemC1ERK6QLineFP13QGraphicsItem(void* qthis, void* arg0, void* arg1);
  // proto:  QLineF QGraphicsLineItem::line();
extern void _ZNK17QGraphicsLineItem4lineEv(void* qthis);
  // proto:  QPainterPath QGraphicsLineItem::opaqueArea();
extern void _ZNK17QGraphicsLineItem10opaqueAreaEv(void* qthis);
  // proto:  void QGraphicsLineItem::setLine(qreal x1, qreal y1, qreal x2, qreal y2);
extern void demth_ZN17QGraphicsLineItem7setLineEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  QRectF QGraphicsLineItem::boundingRect();
extern void _ZNK17QGraphicsLineItem12boundingRectEv(void* qthis);
  // proto:  QPen QGraphicsLineItem::pen();
extern void _ZNK17QGraphicsLineItem3penEv(void* qthis);
  // proto:  void QGraphicsLineItem::setLine(const QLineF & line);
extern void _ZN17QGraphicsLineItem7setLineERK6QLineF(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsLineItem::shape();
extern void _ZNK17QGraphicsLineItem5shapeEv(void* qthis);
  // proto:  void QGraphicsLineItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN17QGraphicsLineItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  int QGraphicsLineItem::type();
extern void _ZNK17QGraphicsLineItem4typeEv(void* qthis);
  // proto:  void QGraphicsLineItem::QGraphicsLineItem(const QGraphicsLineItem & );
extern void* dector_ZN17QGraphicsLineItemC1ERKS_(void* arg0);
extern void _ZN17QGraphicsLineItemC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsLineItem::QGraphicsLineItem(qreal x1, qreal y1, qreal x2, qreal y2, QGraphicsItem * parent);
extern void* dector_ZN17QGraphicsLineItemC1EddddP13QGraphicsItem(double arg0, double arg1, double arg2, double arg3, void* arg4);
extern void _ZN17QGraphicsLineItemC1EddddP13QGraphicsItem(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4);
  // proto:  bool QGraphicsLineItem::contains(const QPointF & point);
extern void _ZNK17QGraphicsLineItem8containsERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsLineItem::~QGraphicsLineItem();
extern void _ZN17QGraphicsLineItemD0Ev(void* qthis);
  // proto:  bool QGraphicsItemGroup::isObscuredBy(const QGraphicsItem * item);
extern void _ZNK18QGraphicsItemGroup12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsItemGroup::~QGraphicsItemGroup();
extern void _ZN18QGraphicsItemGroupD0Ev(void* qthis);
  // proto:  void QGraphicsItemGroup::QGraphicsItemGroup(QGraphicsItem * parent);
extern void* dector_ZN18QGraphicsItemGroupC1EP13QGraphicsItem(void* arg0);
extern void _ZN18QGraphicsItemGroupC1EP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  int QGraphicsItemGroup::type();
extern void _ZNK18QGraphicsItemGroup4typeEv(void* qthis);
  // proto:  QRectF QGraphicsItemGroup::boundingRect();
extern void _ZNK18QGraphicsItemGroup12boundingRectEv(void* qthis);
  // proto:  void QGraphicsItemGroup::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QGraphicsItemGroup::removeFromGroup(QGraphicsItem * item);
extern void _ZN18QGraphicsItemGroup15removeFromGroupEP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsItemGroup::addToGroup(QGraphicsItem * item);
extern void _ZN18QGraphicsItemGroup10addToGroupEP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsItemGroup::opaqueArea();
extern void _ZNK18QGraphicsItemGroup10opaqueAreaEv(void* qthis);
  // proto:  void QGraphicsItemGroup::QGraphicsItemGroup(const QGraphicsItemGroup & );
extern void* dector_ZN18QGraphicsItemGroupC1ERKS_(void* arg0);
extern void _ZN18QGraphicsItemGroupC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QAbstractGraphicsShapeItem::isObscuredBy(const QGraphicsItem * item);
extern void _ZNK26QAbstractGraphicsShapeItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QBrush QAbstractGraphicsShapeItem::brush();
extern void _ZNK26QAbstractGraphicsShapeItem5brushEv(void* qthis);
  // proto:  void QAbstractGraphicsShapeItem::QAbstractGraphicsShapeItem(QGraphicsItem * parent);
extern void* dector_ZN26QAbstractGraphicsShapeItemC1EP13QGraphicsItem(void* arg0);
extern void _ZN26QAbstractGraphicsShapeItemC1EP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QPainterPath QAbstractGraphicsShapeItem::opaqueArea();
extern void _ZNK26QAbstractGraphicsShapeItem10opaqueAreaEv(void* qthis);
  // proto:  void QAbstractGraphicsShapeItem::QAbstractGraphicsShapeItem(const QAbstractGraphicsShapeItem & );
extern void* dector_ZN26QAbstractGraphicsShapeItemC1ERKS_(void* arg0);
extern void _ZN26QAbstractGraphicsShapeItemC1ERKS_(void* qthis, void* arg0);
  // proto:  void QAbstractGraphicsShapeItem::setBrush(const QBrush & brush);
extern void _ZN26QAbstractGraphicsShapeItem8setBrushERK6QBrush(void* qthis, void* arg0);
  // proto:  void QAbstractGraphicsShapeItem::setPen(const QPen & pen);
extern void _ZN26QAbstractGraphicsShapeItem6setPenERK4QPen(void* qthis, void* arg0);
  // proto:  QPen QAbstractGraphicsShapeItem::pen();
extern void _ZNK26QAbstractGraphicsShapeItem3penEv(void* qthis);
  // proto:  void QAbstractGraphicsShapeItem::~QAbstractGraphicsShapeItem();
extern void _ZN26QAbstractGraphicsShapeItemD0Ev(void* qthis);
  // proto:  void QGraphicsItem::QGraphicsItem(const QGraphicsItem & );
extern void* dector_ZN13QGraphicsItemC1ERKS_(void* arg0);
extern void _ZN13QGraphicsItemC1ERKS_(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsItem::mapFromParent(const QPainterPath & path);
extern void _ZNK13QGraphicsItem13mapFromParentERK12QPainterPath(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QPointF & point);
extern void _ZNK13QGraphicsItem11mapFromItemEPKS_RK7QPointF(void* qthis, void* arg0, void* arg1);
  // proto:  QGraphicsItem * QGraphicsItem::focusItem();
extern void _ZNK13QGraphicsItem9focusItemEv(void* qthis);
  // proto:  QGraphicsObject * QGraphicsItem::parentObject();
extern void _ZNK13QGraphicsItem12parentObjectEv(void* qthis);
  // proto:  void QGraphicsItem::setTransformOriginPoint(const QPointF & origin);
extern void _ZN13QGraphicsItem23setTransformOriginPointERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsItem::ungrabMouse();
extern void _ZN13QGraphicsItem11ungrabMouseEv(void* qthis);
  // proto:  int QGraphicsItem::type();
extern void _ZNK13QGraphicsItem4typeEv(void* qthis);
  // proto:  bool QGraphicsItem::isSelected();
extern void _ZNK13QGraphicsItem10isSelectedEv(void* qthis);
  // proto:  QPolygonF QGraphicsItem::mapFromItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem11mapFromItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4);
  // proto:  QGraphicsWidget * QGraphicsItem::parentWidget();
extern void _ZNK13QGraphicsItem12parentWidgetEv(void* qthis);
  // proto:  void QGraphicsItem::resetTransform();
extern void _ZN13QGraphicsItem14resetTransformEv(void* qthis);
  // proto:  QRegion QGraphicsItem::boundingRegion(const QTransform & itemToDeviceTransform);
extern void _ZNK13QGraphicsItem14boundingRegionERK10QTransform(void* qthis, void* arg0);
  // proto:  void QGraphicsItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN13QGraphicsItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  bool QGraphicsItem::isActive();
extern void _ZNK13QGraphicsItem8isActiveEv(void* qthis);
  // proto:  void QGraphicsItem::QGraphicsItem(QGraphicsItem * parent);
extern void* dector_ZN13QGraphicsItemC1EPS_(void* arg0);
extern void _ZN13QGraphicsItemC1EPS_(void* qthis, void* arg0);
  // proto:  QPolygonF QGraphicsItem::mapToParent(const QPolygonF & polygon);
extern void _ZNK13QGraphicsItem11mapToParentERK9QPolygonF(void* qthis, void* arg0);
  // proto:  bool QGraphicsItem::isWidget();
extern void _ZNK13QGraphicsItem8isWidgetEv(void* qthis);
  // proto:  void QGraphicsItem::setParentItem(QGraphicsItem * parent);
extern void _ZN13QGraphicsItem13setParentItemEPS_(void* qthis, void* arg0);
  // proto:  QPolygonF QGraphicsItem::mapToItem(const QGraphicsItem * item, const QRectF & rect);
extern void _ZNK13QGraphicsItem9mapToItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1);
  // proto:  QGraphicsWidget * QGraphicsItem::window();
extern void _ZNK13QGraphicsItem6windowEv(void* qthis);
  // proto:  QPointF QGraphicsItem::scenePos();
extern void _ZNK13QGraphicsItem8scenePosEv(void* qthis);
  // proto:  bool QGraphicsItem::handlesChildEvents();
extern void _ZNK13QGraphicsItem18handlesChildEventsEv(void* qthis);
  // proto:  void QGraphicsItem::setOpacity(qreal opacity);
extern void _ZN13QGraphicsItem10setOpacityEd(void* qthis, double arg0);
  // proto:  QTransform QGraphicsItem::sceneTransform();
extern void _ZNK13QGraphicsItem14sceneTransformEv(void* qthis);
  // proto:  void QGraphicsItem::setZValue(qreal z);
extern void _ZN13QGraphicsItem9setZValueEd(void* qthis, double arg0);
  // proto:  QPolygonF QGraphicsItem::mapFromParent(const QRectF & rect);
extern void _ZNK13QGraphicsItem13mapFromParentERK6QRectF(void* qthis, void* arg0);
  // proto:  QPolygonF QGraphicsItem::mapFromParent(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem13mapFromParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  bool QGraphicsItem::isObscured(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem10isObscuredEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsItem::installSceneEventFilter(QGraphicsItem * filterItem);
extern void _ZN13QGraphicsItem23installSceneEventFilterEPS_(void* qthis, void* arg0);
  // proto:  void QGraphicsItem::setY(qreal y);
extern void _ZN13QGraphicsItem4setYEd(void* qthis, double arg0);
  // proto:  QRectF QGraphicsItem::mapRectToItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem13mapRectToItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4);
  // proto:  QGraphicsItem * QGraphicsItem::parentItem();
extern void _ZNK13QGraphicsItem10parentItemEv(void* qthis);
  // proto:  void QGraphicsItem::clearFocus();
extern void _ZN13QGraphicsItem10clearFocusEv(void* qthis);
  // proto:  bool QGraphicsItem::isWindow();
extern void _ZNK13QGraphicsItem8isWindowEv(void* qthis);
  // proto:  QPointF QGraphicsItem::transformOriginPoint();
extern void _ZNK13QGraphicsItem20transformOriginPointEv(void* qthis);
  // proto:  QRectF QGraphicsItem::boundingRect();
extern void _ZNK13QGraphicsItem12boundingRectEv(void* qthis);
  // proto:  QRectF QGraphicsItem::childrenBoundingRect();
extern void _ZNK13QGraphicsItem20childrenBoundingRectEv(void* qthis);
  // proto:  bool QGraphicsItem::isObscured(const QRectF & rect);
extern void _ZNK13QGraphicsItem10isObscuredERK6QRectF(void* qthis, void* arg0);
  // proto:  QPolygonF QGraphicsItem::mapFromScene(const QRectF & rect);
extern void _ZNK13QGraphicsItem12mapFromSceneERK6QRectF(void* qthis, void* arg0);
  // proto:  bool QGraphicsItem::hasCursor();
extern void _ZNK13QGraphicsItem9hasCursorEv(void* qthis);
  // proto:  void QGraphicsItem::setGraphicsEffect(QGraphicsEffect * effect);
extern void _ZN13QGraphicsItem17setGraphicsEffectEP15QGraphicsEffect(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsItem::mapToParent(const QPainterPath & path);
extern void _ZNK13QGraphicsItem11mapToParentERK12QPainterPath(void* qthis, void* arg0);
  // proto:  void QGraphicsItem::ensureVisible(qreal x, qreal y, qreal w, qreal h, int xmargin, int ymargin);
extern void demth_ZN13QGraphicsItem13ensureVisibleEddddii(void* qthis, double arg0, double arg1, double arg2, double arg3, int arg4, int arg5);
  // proto:  QPolygonF QGraphicsItem::mapToItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem9mapToItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4);
  // proto:  QPointF QGraphicsItem::mapToItem(const QGraphicsItem * item, qreal x, qreal y);
extern void demth_ZNK13QGraphicsItem9mapToItemEPKS_dd(void* qthis, void* arg0, double arg1, double arg2);
  // proto:  QRectF QGraphicsItem::mapRectToParent(const QRectF & rect);
extern void _ZNK13QGraphicsItem15mapRectToParentERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsItem::setToolTip(const QString & toolTip);
extern void _ZN13QGraphicsItem10setToolTipERK7QString(void* qthis, void* arg0);
  // proto:  qreal QGraphicsItem::rotation();
extern void _ZNK13QGraphicsItem8rotationEv(void* qthis);
  // proto:  QGraphicsScene * QGraphicsItem::scene();
extern void _ZNK13QGraphicsItem5sceneEv(void* qthis);
  // proto:  QPainterPath QGraphicsItem::mapToItem(const QGraphicsItem * item, const QPainterPath & path);
extern void _ZNK13QGraphicsItem9mapToItemEPKS_RK12QPainterPath(void* qthis, void* arg0, void* arg1);
  // proto:  QRectF QGraphicsItem::mapRectToParent(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem15mapRectToParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  QPolygonF QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QRectF & rect);
extern void _ZNK13QGraphicsItem11mapFromItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1);
  // proto:  QRectF QGraphicsItem::mapRectFromParent(const QRectF & rect);
extern void _ZNK13QGraphicsItem17mapRectFromParentERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsItem::setFocusProxy(QGraphicsItem * item);
extern void _ZN13QGraphicsItem13setFocusProxyEPS_(void* qthis, void* arg0);
  // proto:  bool QGraphicsItem::acceptDrops();
extern void _ZNK13QGraphicsItem11acceptDropsEv(void* qthis);
  // proto:  QPointF QGraphicsItem::mapToParent(const QPointF & point);
extern void _ZNK13QGraphicsItem11mapToParentERK7QPointF(void* qthis, void* arg0);
  // proto:  QRectF QGraphicsItem::mapRectFromScene(const QRectF & rect);
extern void _ZNK13QGraphicsItem16mapRectFromSceneERK6QRectF(void* qthis, void* arg0);
  // proto:  QGraphicsItem * QGraphicsItem::focusScopeItem();
extern void _ZNK13QGraphicsItem14focusScopeItemEv(void* qthis);
  // proto:  void QGraphicsItem::removeSceneEventFilter(QGraphicsItem * filterItem);
extern void _ZN13QGraphicsItem22removeSceneEventFilterEPS_(void* qthis, void* arg0);
  // proto:  QGraphicsItem * QGraphicsItem::focusProxy();
extern void _ZNK13QGraphicsItem10focusProxyEv(void* qthis);
  // proto:  QPointF QGraphicsItem::mapToItem(const QGraphicsItem * item, const QPointF & point);
extern void _ZNK13QGraphicsItem9mapToItemEPKS_RK7QPointF(void* qthis, void* arg0, void* arg1);
  // proto:  QRectF QGraphicsItem::sceneBoundingRect();
extern void _ZNK13QGraphicsItem17sceneBoundingRectEv(void* qthis);
  // proto:  void QGraphicsItem::~QGraphicsItem();
extern void _ZN13QGraphicsItemD0Ev(void* qthis);
  // proto:  void QGraphicsItem::setX(qreal x);
extern void _ZN13QGraphicsItem4setXEd(void* qthis, double arg0);
  // proto:  void QGraphicsItem::update(qreal x, qreal y, qreal width, qreal height);
extern void demth_ZN13QGraphicsItem6updateEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsItem::setSelected(bool selected);
extern void _ZN13QGraphicsItem11setSelectedEb(void* qthis, bool arg0);
  // proto:  QRectF QGraphicsItem::mapRectToItem(const QGraphicsItem * item, const QRectF & rect);
extern void _ZNK13QGraphicsItem13mapRectToItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsItem::stackBefore(const QGraphicsItem * sibling);
extern void _ZN13QGraphicsItem11stackBeforeEPKS_(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsItem::mapFromItem(const QGraphicsItem * item, qreal x, qreal y);
extern void demth_ZNK13QGraphicsItem11mapFromItemEPKS_dd(void* qthis, void* arg0, double arg1, double arg2);
  // proto:  void QGraphicsItem::resetMatrix();
extern void _ZN13QGraphicsItem11resetMatrixEv(void* qthis);
  // proto:  QPainterPath QGraphicsItem::opaqueArea();
extern void _ZNK13QGraphicsItem10opaqueAreaEv(void* qthis);
  // proto:  void QGraphicsItem::unsetCursor();
extern void _ZN13QGraphicsItem11unsetCursorEv(void* qthis);
  // proto:  QPointF QGraphicsItem::mapFromParent(qreal x, qreal y);
extern void demth_ZNK13QGraphicsItem13mapFromParentEdd(void* qthis, double arg0, double arg1);
  // proto:  QRectF QGraphicsItem::mapRectToScene(const QRectF & rect);
extern void _ZNK13QGraphicsItem14mapRectToSceneERK6QRectF(void* qthis, void* arg0);
  // proto:  QRectF QGraphicsItem::mapRectFromItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem15mapRectFromItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4);
  // proto:  qreal QGraphicsItem::scale();
extern void _ZNK13QGraphicsItem5scaleEv(void* qthis);
  // proto:  void QGraphicsItem::setBoundingRegionGranularity(qreal granularity);
extern void _ZN13QGraphicsItem28setBoundingRegionGranularityEd(void* qthis, double arg0);
  // proto:  void QGraphicsItem::setAcceptDrops(bool on);
extern void _ZN13QGraphicsItem14setAcceptDropsEb(void* qthis, bool arg0);
  // proto:  QPolygonF QGraphicsItem::mapFromScene(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem12mapFromSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsItem::ungrabKeyboard();
extern void _ZN13QGraphicsItem14ungrabKeyboardEv(void* qthis);
  // proto:  void QGraphicsItem::setEnabled(bool enabled);
extern void _ZN13QGraphicsItem10setEnabledEb(void* qthis, bool arg0);
  // proto:  QGraphicsEffect * QGraphicsItem::graphicsEffect();
extern void _ZNK13QGraphicsItem14graphicsEffectEv(void* qthis);
  // proto:  bool QGraphicsItem::acceptHoverEvents();
extern void _ZNK13QGraphicsItem17acceptHoverEventsEv(void* qthis);
  // proto:  QGraphicsWidget * QGraphicsItem::topLevelWidget();
extern void _ZNK13QGraphicsItem14topLevelWidgetEv(void* qthis);
  // proto:  QList<QGraphicsTransform *> QGraphicsItem::transformations();
extern void _ZNK13QGraphicsItem15transformationsEv(void* qthis);
  // proto:  QPolygonF QGraphicsItem::mapToScene(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem10mapToSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  QPointF QGraphicsItem::mapToScene(qreal x, qreal y);
extern void demth_ZNK13QGraphicsItem10mapToSceneEdd(void* qthis, double arg0, double arg1);
  // proto:  QRectF QGraphicsItem::mapRectFromScene(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem16mapRectFromSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsItem::advance(int phase);
extern void _ZN13QGraphicsItem7advanceEi(void* qthis, int arg0);
  // proto:  QMatrix QGraphicsItem::sceneMatrix();
extern void _ZNK13QGraphicsItem11sceneMatrixEv(void* qthis);
  // proto:  void QGraphicsItem::setFiltersChildEvents(bool enabled);
extern void _ZN13QGraphicsItem21setFiltersChildEventsEb(void* qthis, bool arg0);
  // proto:  QPolygonF QGraphicsItem::mapToScene(const QPolygonF & polygon);
extern void _ZNK13QGraphicsItem10mapToSceneERK9QPolygonF(void* qthis, void* arg0);
  // proto:  QTransform QGraphicsItem::itemTransform(const QGraphicsItem * other, bool * ok);
extern void _ZNK13QGraphicsItem13itemTransformEPKS_Pb(void* qthis, void* arg0, bool* arg1);
  // proto:  void QGraphicsItem::setTransformOriginPoint(qreal ax, qreal ay);
extern void demth_ZN13QGraphicsItem23setTransformOriginPointEdd(void* qthis, double arg0, double arg1);
  // proto:  void QGraphicsItem::moveBy(qreal dx, qreal dy);
extern void demth_ZN13QGraphicsItem6moveByEdd(void* qthis, double arg0, double arg1);
  // proto:  QPolygonF QGraphicsItem::mapFromScene(const QPolygonF & polygon);
extern void _ZNK13QGraphicsItem12mapFromSceneERK9QPolygonF(void* qthis, void* arg0);
  // proto:  QGraphicsItemGroup * QGraphicsItem::group();
extern void _ZNK13QGraphicsItem5groupEv(void* qthis);
  // proto:  QPainterPath QGraphicsItem::shape();
extern void _ZNK13QGraphicsItem5shapeEv(void* qthis);
  // proto:  QPointF QGraphicsItem::mapFromScene(qreal x, qreal y);
extern void demth_ZNK13QGraphicsItem12mapFromSceneEdd(void* qthis, double arg0, double arg1);
  // proto:  void QGraphicsItem::scroll(qreal dx, qreal dy, const QRectF & rect);
extern void _ZN13QGraphicsItem6scrollEddRK6QRectF(void* qthis, double arg0, double arg1, void* arg2);
  // proto:  bool QGraphicsItem::isObscuredBy(const QGraphicsItem * item);
extern void _ZNK13QGraphicsItem12isObscuredByEPKS_(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsItem::mapFromParent(const QPointF & point);
extern void _ZNK13QGraphicsItem13mapFromParentERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsItem::setData(int key, const QVariant & value);
extern void _ZN13QGraphicsItem7setDataEiRK8QVariant(void* qthis, int arg0, void* arg1);
  // proto:  QGraphicsItem * QGraphicsItem::commonAncestorItem(const QGraphicsItem * other);
extern void _ZNK13QGraphicsItem18commonAncestorItemEPKS_(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsItem::mapFromScene(const QPainterPath & path);
extern void _ZNK13QGraphicsItem12mapFromSceneERK12QPainterPath(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsItem::mapToScene(const QPainterPath & path);
extern void _ZNK13QGraphicsItem10mapToSceneERK12QPainterPath(void* qthis, void* arg0);
  // proto:  QPolygonF QGraphicsItem::mapToParent(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem11mapToParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsItem::setGroup(QGraphicsItemGroup * group);
extern void _ZN13QGraphicsItem8setGroupEP18QGraphicsItemGroup(void* qthis, void* arg0);
  // proto:  QRectF QGraphicsItem::mapRectFromParent(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem17mapRectFromParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsItem::show();
extern void demth_ZN13QGraphicsItem4showEv(void* qthis);
  // proto:  QRectF QGraphicsItem::mapRectFromItem(const QGraphicsItem * item, const QRectF & rect);
extern void _ZNK13QGraphicsItem15mapRectFromItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1);
  // proto:  qreal QGraphicsItem::y();
extern void demth_ZNK13QGraphicsItem1yEv(void* qthis);
  // proto:  QPointF QGraphicsItem::mapFromScene(const QPointF & point);
extern void _ZNK13QGraphicsItem12mapFromSceneERK7QPointF(void* qthis, void* arg0);
  // proto:  bool QGraphicsItem::hasFocus();
extern void _ZNK13QGraphicsItem8hasFocusEv(void* qthis);
  // proto:  QPainterPath QGraphicsItem::clipPath();
extern void _ZNK13QGraphicsItem8clipPathEv(void* qthis);
  // proto:  void QGraphicsItem::setPos(qreal x, qreal y);
extern void demth_ZN13QGraphicsItem6setPosEdd(void* qthis, double arg0, double arg1);
  // proto:  bool QGraphicsItem::isEnabled();
extern void _ZNK13QGraphicsItem9isEnabledEv(void* qthis);
  // proto:  bool QGraphicsItem::contains(const QPointF & point);
extern void _ZNK13QGraphicsItem8containsERK7QPointF(void* qthis, void* arg0);
  // proto:  bool QGraphicsItem::isPanel();
extern void _ZNK13QGraphicsItem7isPanelEv(void* qthis);
  // proto:  bool QGraphicsItem::filtersChildEvents();
extern void _ZNK13QGraphicsItem18filtersChildEventsEv(void* qthis);
  // proto:  void QGraphicsItem::grabKeyboard();
extern void _ZN13QGraphicsItem12grabKeyboardEv(void* qthis);
  // proto:  QPainterPath QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QPainterPath & path);
extern void _ZNK13QGraphicsItem11mapFromItemEPKS_RK12QPainterPath(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsItem::setActive(bool active);
extern void _ZN13QGraphicsItem9setActiveEb(void* qthis, bool arg0);
  // proto:  QGraphicsObject * QGraphicsItem::toGraphicsObject();
extern void _ZN13QGraphicsItem16toGraphicsObjectEv(void* qthis);
  // proto:  QPolygonF QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QPolygonF & polygon);
extern void _ZNK13QGraphicsItem11mapFromItemEPKS_RK9QPolygonF(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsItem::setHandlesChildEvents(bool enabled);
extern void _ZN13QGraphicsItem21setHandlesChildEventsEb(void* qthis, bool arg0);
  // proto:  QPolygonF QGraphicsItem::mapFromParent(const QPolygonF & polygon);
extern void _ZNK13QGraphicsItem13mapFromParentERK9QPolygonF(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsItem::mapToParent(qreal x, qreal y);
extern void demth_ZNK13QGraphicsItem11mapToParentEdd(void* qthis, double arg0, double arg1);
  // proto:  void QGraphicsItem::setMatrix(const QMatrix & matrix, bool combine);
extern void _ZN13QGraphicsItem9setMatrixERK7QMatrixb(void* qthis, void* arg0, bool arg1);
  // proto:  void QGraphicsItem::update(const QRectF & rect);
extern void _ZN13QGraphicsItem6updateERK6QRectF(void* qthis, void* arg0);
  // proto:  QPolygonF QGraphicsItem::mapToItem(const QGraphicsItem * item, const QPolygonF & polygon);
extern void _ZNK13QGraphicsItem9mapToItemEPKS_RK9QPolygonF(void* qthis, void* arg0, void* arg1);
  // proto:  QTransform QGraphicsItem::transform();
extern void _ZNK13QGraphicsItem9transformEv(void* qthis);
  // proto:  QVariant QGraphicsItem::data(int key);
extern void _ZNK13QGraphicsItem4dataEi(void* qthis, int arg0);
  // proto:  void QGraphicsItem::hide();
extern void demth_ZN13QGraphicsItem4hideEv(void* qthis);
  // proto:  bool QGraphicsItem::isUnderMouse();
extern void _ZNK13QGraphicsItem12isUnderMouseEv(void* qthis);
  // proto:  void QGraphicsItem::setAcceptTouchEvents(bool enabled);
extern void _ZN13QGraphicsItem20setAcceptTouchEventsEb(void* qthis, bool arg0);
  // proto:  void QGraphicsItem::setAcceptHoverEvents(bool enabled);
extern void _ZN13QGraphicsItem20setAcceptHoverEventsEb(void* qthis, bool arg0);
  // proto:  QList<QGraphicsItem *> QGraphicsItem::childItems();
extern void _ZNK13QGraphicsItem10childItemsEv(void* qthis);
  // proto:  bool QGraphicsItem::isAncestorOf(const QGraphicsItem * child);
extern void _ZNK13QGraphicsItem12isAncestorOfEPKS_(void* qthis, void* arg0);
  // proto:  qreal QGraphicsItem::opacity();
extern void _ZNK13QGraphicsItem7opacityEv(void* qthis);
  // proto:  bool QGraphicsItem::isVisibleTo(const QGraphicsItem * parent);
extern void _ZNK13QGraphicsItem11isVisibleToEPKS_(void* qthis, void* arg0);
  // proto:  QString QGraphicsItem::toolTip();
extern void _ZNK13QGraphicsItem7toolTipEv(void* qthis);
  // proto:  QCursor QGraphicsItem::cursor();
extern void _ZNK13QGraphicsItem6cursorEv(void* qthis);
  // proto:  QPointF QGraphicsItem::mapToScene(const QPointF & point);
extern void _ZNK13QGraphicsItem10mapToSceneERK7QPointF(void* qthis, void* arg0);
  // proto:  qreal QGraphicsItem::zValue();
extern void _ZNK13QGraphicsItem6zValueEv(void* qthis);
  // proto:  QMatrix QGraphicsItem::matrix();
extern void _ZNK13QGraphicsItem6matrixEv(void* qthis);
  // proto:  QRectF QGraphicsItem::mapRectToScene(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsItem14mapRectToSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsItem::setPos(const QPointF & pos);
extern void _ZN13QGraphicsItem6setPosERK7QPointF(void* qthis, void* arg0);
  // proto:  QGraphicsItem * QGraphicsItem::panel();
extern void _ZNK13QGraphicsItem5panelEv(void* qthis);
  // proto:  bool QGraphicsItem::isClipped();
extern void _ZNK13QGraphicsItem9isClippedEv(void* qthis);
  // proto:  QGraphicsItem * QGraphicsItem::topLevelItem();
extern void _ZNK13QGraphicsItem12topLevelItemEv(void* qthis);
  // proto:  QPolygonF QGraphicsItem::mapToScene(const QRectF & rect);
extern void _ZNK13QGraphicsItem10mapToSceneERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsItem::setScale(qreal scale);
extern void _ZN13QGraphicsItem8setScaleEd(void* qthis, double arg0);
  // proto:  void QGraphicsItem::setCursor(const QCursor & cursor);
extern void _ZN13QGraphicsItem9setCursorERK7QCursor(void* qthis, void* arg0);
  // proto:  bool QGraphicsItem::isVisible();
extern void _ZNK13QGraphicsItem9isVisibleEv(void* qthis);
  // proto:  QPointF QGraphicsItem::pos();
extern void _ZNK13QGraphicsItem3posEv(void* qthis);
  // proto:  bool QGraphicsItem::isBlockedByModalPanel(QGraphicsItem ** blockingPanel);
extern void _ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_(void* qthis, void* arg0);
  // proto:  qreal QGraphicsItem::effectiveOpacity();
extern void _ZNK13QGraphicsItem16effectiveOpacityEv(void* qthis);
  // proto:  void QGraphicsItem::ensureVisible(const QRectF & rect, int xmargin, int ymargin);
extern void _ZN13QGraphicsItem13ensureVisibleERK6QRectFii(void* qthis, void* arg0, int arg1, int arg2);
  // proto:  qreal QGraphicsItem::boundingRegionGranularity();
extern void _ZNK13QGraphicsItem25boundingRegionGranularityEv(void* qthis);
  // proto:  qreal QGraphicsItem::x();
extern void demth_ZNK13QGraphicsItem1xEv(void* qthis);
  // proto:  void QGraphicsItem::grabMouse();
extern void _ZN13QGraphicsItem9grabMouseEv(void* qthis);
  // proto:  void QGraphicsItem::setVisible(bool visible);
extern void _ZN13QGraphicsItem10setVisibleEb(void* qthis, bool arg0);
  // proto:  void QGraphicsItem::setRotation(qreal angle);
extern void _ZN13QGraphicsItem11setRotationEd(void* qthis, double arg0);
  // proto:  QTransform QGraphicsItem::deviceTransform(const QTransform & viewportTransform);
extern void _ZNK13QGraphicsItem15deviceTransformERK10QTransform(void* qthis, void* arg0);
  // proto:  bool QGraphicsItem::acceptTouchEvents();
extern void _ZNK13QGraphicsItem17acceptTouchEventsEv(void* qthis);
  // proto:  void QGraphicsItem::setTransform(const QTransform & matrix, bool combine);
extern void _ZN13QGraphicsItem12setTransformERK10QTransformb(void* qthis, void* arg0, bool arg1);
  // proto:  QPolygonF QGraphicsItem::mapToParent(const QRectF & rect);
extern void _ZNK13QGraphicsItem11mapToParentERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsObject::QGraphicsObject(QGraphicsItem * parent);
extern void* dector_ZN15QGraphicsObjectC1EP13QGraphicsItem(void* arg0);
extern void _ZN15QGraphicsObjectC1EP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsObject::~QGraphicsObject();
extern void _ZN15QGraphicsObjectD0Ev(void* qthis);
  // proto:  const QMetaObject * QGraphicsObject::metaObject();
extern void _ZNK15QGraphicsObject10metaObjectEv(void* qthis);
  // proto:  int QGraphicsSimpleTextItem::type();
extern void _ZNK23QGraphicsSimpleTextItem4typeEv(void* qthis);
  // proto:  QFont QGraphicsSimpleTextItem::font();
extern void _ZNK23QGraphicsSimpleTextItem4fontEv(void* qthis);
  // proto:  void QGraphicsSimpleTextItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN23QGraphicsSimpleTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QGraphicsSimpleTextItem::~QGraphicsSimpleTextItem();
extern void _ZN23QGraphicsSimpleTextItemD0Ev(void* qthis);
  // proto:  void QGraphicsSimpleTextItem::setText(const QString & text);
extern void _ZN23QGraphicsSimpleTextItem7setTextERK7QString(void* qthis, void* arg0);
  // proto:  QString QGraphicsSimpleTextItem::text();
extern void _ZNK23QGraphicsSimpleTextItem4textEv(void* qthis);
  // proto:  void QGraphicsSimpleTextItem::QGraphicsSimpleTextItem(const QString & text, QGraphicsItem * parent);
extern void* dector_ZN23QGraphicsSimpleTextItemC1ERK7QStringP13QGraphicsItem(void* arg0, void* arg1);
extern void _ZN23QGraphicsSimpleTextItemC1ERK7QStringP13QGraphicsItem(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsSimpleTextItem::QGraphicsSimpleTextItem(const QGraphicsSimpleTextItem & );
extern void* dector_ZN23QGraphicsSimpleTextItemC1ERKS_(void* arg0);
extern void _ZN23QGraphicsSimpleTextItemC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QGraphicsSimpleTextItem::isObscuredBy(const QGraphicsItem * item);
extern void _ZNK23QGraphicsSimpleTextItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsSimpleTextItem::shape();
extern void _ZNK23QGraphicsSimpleTextItem5shapeEv(void* qthis);
  // proto:  void QGraphicsSimpleTextItem::QGraphicsSimpleTextItem(QGraphicsItem * parent);
extern void* dector_ZN23QGraphicsSimpleTextItemC1EP13QGraphicsItem(void* arg0);
extern void _ZN23QGraphicsSimpleTextItemC1EP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsSimpleTextItem::setFont(const QFont & font);
extern void _ZN23QGraphicsSimpleTextItem7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsSimpleTextItem::opaqueArea();
extern void _ZNK23QGraphicsSimpleTextItem10opaqueAreaEv(void* qthis);
  // proto:  QRectF QGraphicsSimpleTextItem::boundingRect();
extern void _ZNK23QGraphicsSimpleTextItem12boundingRectEv(void* qthis);
  // proto:  bool QGraphicsSimpleTextItem::contains(const QPointF & point);
extern void _ZNK23QGraphicsSimpleTextItem8containsERK7QPointF(void* qthis, void* arg0);
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

  // proto:  bool QGraphicsTextItem::openExternalLinks();
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

  // proto:  qreal QGraphicsTextItem::textWidth();
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

  // proto:  void QGraphicsTextItem::setTextWidth(qreal width);
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

  // proto:  void QGraphicsTextItem::setTextCursor(const QTextCursor & cursor);
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

  // proto:  int QGraphicsTextItem::type();
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

  // proto:  QFont QGraphicsTextItem::font();
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

  // proto:  void QGraphicsTextItem::QGraphicsTextItem(const QString & text, QGraphicsItem * parent);
func NewQGraphicsTextItem(args ...interface{}) QGraphicsTextItem {
  return QGraphicsTextItem{}
}

  // proto:  const QMetaObject * QGraphicsTextItem::metaObject();
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

  // proto:  void QGraphicsTextItem::setOpenExternalLinks(bool open);
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

  // proto:  void QGraphicsTextItem::setTabChangesFocus(bool b);
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

  // proto:  QString QGraphicsTextItem::toHtml();
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

  // proto:  void QGraphicsTextItem::setDocument(QTextDocument * document);
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

  // proto:  void QGraphicsTextItem::setPlainText(const QString & text);
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

  // proto:  void QGraphicsTextItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
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

  // proto:  void QGraphicsTextItem::setFont(const QFont & font);
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

  // proto:  void QGraphicsTextItem::setDefaultTextColor(const QColor & c);
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

  // proto:  QColor QGraphicsTextItem::defaultTextColor();
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

  // proto:  void QGraphicsTextItem::~QGraphicsTextItem();
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

  // proto:  QPainterPath QGraphicsTextItem::shape();
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

  // proto:  QTextCursor QGraphicsTextItem::textCursor();
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

  // proto:  QRectF QGraphicsTextItem::boundingRect();
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

  // proto:  QString QGraphicsTextItem::toPlainText();
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

  // proto:  void QGraphicsTextItem::setHtml(const QString & html);
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

  // proto:  bool QGraphicsTextItem::tabChangesFocus();
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

  // proto:  QTextDocument * QGraphicsTextItem::document();
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

  // proto:  bool QGraphicsTextItem::isObscuredBy(const QGraphicsItem * item);
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

  // proto:  QPainterPath QGraphicsTextItem::opaqueArea();
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

  // proto:  bool QGraphicsTextItem::contains(const QPointF & point);
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

  // proto:  void QGraphicsTextItem::adjustSize();
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

  // proto:  void QGraphicsPixmapItem::QGraphicsPixmapItem(QGraphicsItem * parent);
func NewQGraphicsPixmapItem(args ...interface{}) QGraphicsPixmapItem {
  return QGraphicsPixmapItem{}
}

  // proto:  void QGraphicsPixmapItem::~QGraphicsPixmapItem();
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

  // proto:  QPainterPath QGraphicsPixmapItem::opaqueArea();
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

  // proto:  bool QGraphicsPixmapItem::isObscuredBy(const QGraphicsItem * item);
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

  // proto:  int QGraphicsPixmapItem::type();
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

  // proto:  QPainterPath QGraphicsPixmapItem::shape();
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

  // proto:  QPixmap QGraphicsPixmapItem::pixmap();
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

  // proto:  void QGraphicsPixmapItem::setOffset(qreal x, qreal y);
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

  // proto:  void QGraphicsPixmapItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
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

  // proto:  QPointF QGraphicsPixmapItem::offset();
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

  // proto:  QRectF QGraphicsPixmapItem::boundingRect();
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

  // proto:  bool QGraphicsPixmapItem::contains(const QPointF & point);
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

  // proto:  void QGraphicsPixmapItem::setPixmap(const QPixmap & pixmap);
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

  // proto:  bool QGraphicsRectItem::isObscuredBy(const QGraphicsItem * item);
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

  // proto:  QRectF QGraphicsRectItem::boundingRect();
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

  // proto:  void QGraphicsRectItem::QGraphicsRectItem(const QGraphicsRectItem & );
func NewQGraphicsRectItem(args ...interface{}) QGraphicsRectItem {
  return QGraphicsRectItem{}
}

  // proto:  int QGraphicsRectItem::type();
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

  // proto:  QRectF QGraphicsRectItem::rect();
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

  // proto:  QPainterPath QGraphicsRectItem::shape();
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

  // proto:  void QGraphicsRectItem::~QGraphicsRectItem();
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

  // proto:  QPainterPath QGraphicsRectItem::opaqueArea();
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

  // proto:  void QGraphicsRectItem::setRect(const QRectF & rect);
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

  // proto:  bool QGraphicsRectItem::contains(const QPointF & point);
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

  // proto:  void QGraphicsRectItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
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

  // proto:  void QGraphicsEllipseItem::setStartAngle(int angle);
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

  // proto:  void QGraphicsEllipseItem::QGraphicsEllipseItem(const QGraphicsEllipseItem & );
func NewQGraphicsEllipseItem(args ...interface{}) QGraphicsEllipseItem {
  return QGraphicsEllipseItem{}
}

  // proto:  bool QGraphicsEllipseItem::contains(const QPointF & point);
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

  // proto:  void QGraphicsEllipseItem::setRect(const QRectF & rect);
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

  // proto:  void QGraphicsEllipseItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
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

  // proto:  bool QGraphicsEllipseItem::isObscuredBy(const QGraphicsItem * item);
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

  // proto:  QRectF QGraphicsEllipseItem::rect();
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

  // proto:  int QGraphicsEllipseItem::spanAngle();
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

  // proto:  int QGraphicsEllipseItem::startAngle();
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

  // proto:  void QGraphicsEllipseItem::setSpanAngle(int angle);
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

  // proto:  int QGraphicsEllipseItem::type();
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

  // proto:  QRectF QGraphicsEllipseItem::boundingRect();
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

  // proto:  QPainterPath QGraphicsEllipseItem::shape();
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

  // proto:  void QGraphicsEllipseItem::~QGraphicsEllipseItem();
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

  // proto:  QPainterPath QGraphicsEllipseItem::opaqueArea();
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

  // proto:  QPainterPath QGraphicsPolygonItem::shape();
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

  // proto:  bool QGraphicsPolygonItem::isObscuredBy(const QGraphicsItem * item);
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

  // proto:  void QGraphicsPolygonItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
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

  // proto:  void QGraphicsPolygonItem::QGraphicsPolygonItem(QGraphicsItem * parent);
func NewQGraphicsPolygonItem(args ...interface{}) QGraphicsPolygonItem {
  return QGraphicsPolygonItem{}
}

  // proto:  QRectF QGraphicsPolygonItem::boundingRect();
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

  // proto:  int QGraphicsPolygonItem::type();
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

  // proto:  void QGraphicsPolygonItem::~QGraphicsPolygonItem();
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

  // proto:  QPolygonF QGraphicsPolygonItem::polygon();
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

  // proto:  QPainterPath QGraphicsPolygonItem::opaqueArea();
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

  // proto:  bool QGraphicsPolygonItem::contains(const QPointF & point);
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

  // proto:  void QGraphicsPolygonItem::setPolygon(const QPolygonF & polygon);
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

  // proto:  void QGraphicsPathItem::setPath(const QPainterPath & path);
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

  // proto:  void QGraphicsPathItem::QGraphicsPathItem(const QPainterPath & path, QGraphicsItem * parent);
func NewQGraphicsPathItem(args ...interface{}) QGraphicsPathItem {
  return QGraphicsPathItem{}
}

  // proto:  bool QGraphicsPathItem::contains(const QPointF & point);
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

  // proto:  QRectF QGraphicsPathItem::boundingRect();
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

  // proto:  int QGraphicsPathItem::type();
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

  // proto:  QPainterPath QGraphicsPathItem::opaqueArea();
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

  // proto:  QPainterPath QGraphicsPathItem::path();
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

  // proto:  void QGraphicsPathItem::~QGraphicsPathItem();
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

  // proto:  QPainterPath QGraphicsPathItem::shape();
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

  // proto:  bool QGraphicsPathItem::isObscuredBy(const QGraphicsItem * item);
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

  // proto:  void QGraphicsPathItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
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

  // proto:  void QGraphicsLineItem::setPen(const QPen & pen);
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

  // proto:  void QGraphicsLineItem::QGraphicsLineItem(QGraphicsItem * parent);
func NewQGraphicsLineItem(args ...interface{}) QGraphicsLineItem {
  return QGraphicsLineItem{}
}

  // proto:  bool QGraphicsLineItem::isObscuredBy(const QGraphicsItem * item);
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

  // proto:  QLineF QGraphicsLineItem::line();
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

  // proto:  QPainterPath QGraphicsLineItem::opaqueArea();
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

  // proto:  void QGraphicsLineItem::setLine(qreal x1, qreal y1, qreal x2, qreal y2);
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

  // proto:  QRectF QGraphicsLineItem::boundingRect();
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

  // proto:  QPen QGraphicsLineItem::pen();
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

  // proto:  QPainterPath QGraphicsLineItem::shape();
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

  // proto:  void QGraphicsLineItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
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

  // proto:  int QGraphicsLineItem::type();
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

  // proto:  bool QGraphicsLineItem::contains(const QPointF & point);
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

  // proto:  void QGraphicsLineItem::~QGraphicsLineItem();
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

  // proto:  bool QGraphicsItemGroup::isObscuredBy(const QGraphicsItem * item);
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

  // proto:  void QGraphicsItemGroup::~QGraphicsItemGroup();
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

  // proto:  void QGraphicsItemGroup::QGraphicsItemGroup(QGraphicsItem * parent);
func NewQGraphicsItemGroup(args ...interface{}) QGraphicsItemGroup {
  return QGraphicsItemGroup{}
}

  // proto:  int QGraphicsItemGroup::type();
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

  // proto:  QRectF QGraphicsItemGroup::boundingRect();
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

  // proto:  void QGraphicsItemGroup::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
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

  // proto:  void QGraphicsItemGroup::removeFromGroup(QGraphicsItem * item);
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

  // proto:  void QGraphicsItemGroup::addToGroup(QGraphicsItem * item);
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

  // proto:  QPainterPath QGraphicsItemGroup::opaqueArea();
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

  // proto:  bool QAbstractGraphicsShapeItem::isObscuredBy(const QGraphicsItem * item);
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

  // proto:  QBrush QAbstractGraphicsShapeItem::brush();
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

  // proto:  void QAbstractGraphicsShapeItem::QAbstractGraphicsShapeItem(QGraphicsItem * parent);
func NewQAbstractGraphicsShapeItem(args ...interface{}) QAbstractGraphicsShapeItem {
  return QAbstractGraphicsShapeItem{}
}

  // proto:  QPainterPath QAbstractGraphicsShapeItem::opaqueArea();
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

  // proto:  void QAbstractGraphicsShapeItem::setBrush(const QBrush & brush);
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

  // proto:  void QAbstractGraphicsShapeItem::setPen(const QPen & pen);
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

  // proto:  QPen QAbstractGraphicsShapeItem::pen();
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

  // proto:  void QAbstractGraphicsShapeItem::~QAbstractGraphicsShapeItem();
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

  // proto:  void QGraphicsItem::QGraphicsItem(const QGraphicsItem & );
func NewQGraphicsItem(args ...interface{}) QGraphicsItem {
  return QGraphicsItem{}
}

  // proto:  QPainterPath QGraphicsItem::mapFromParent(const QPainterPath & path);
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

  // proto:  QPointF QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QPointF & point);
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

  // proto:  QGraphicsItem * QGraphicsItem::focusItem();
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

  // proto:  QGraphicsObject * QGraphicsItem::parentObject();
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

  // proto:  void QGraphicsItem::setTransformOriginPoint(const QPointF & origin);
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

  // proto:  void QGraphicsItem::ungrabMouse();
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

  // proto:  int QGraphicsItem::type();
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

  // proto:  bool QGraphicsItem::isSelected();
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

  // proto:  QGraphicsWidget * QGraphicsItem::parentWidget();
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

  // proto:  void QGraphicsItem::resetTransform();
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

  // proto:  QRegion QGraphicsItem::boundingRegion(const QTransform & itemToDeviceTransform);
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

  // proto:  void QGraphicsItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
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

  // proto:  bool QGraphicsItem::isActive();
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

  // proto:  QPolygonF QGraphicsItem::mapToParent(const QPolygonF & polygon);
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

  // proto:  bool QGraphicsItem::isWidget();
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

  // proto:  void QGraphicsItem::setParentItem(QGraphicsItem * parent);
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

  // proto:  QPolygonF QGraphicsItem::mapToItem(const QGraphicsItem * item, const QRectF & rect);
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

  // proto:  QGraphicsWidget * QGraphicsItem::window();
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

  // proto:  QPointF QGraphicsItem::scenePos();
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

  // proto:  bool QGraphicsItem::handlesChildEvents();
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

  // proto:  void QGraphicsItem::setOpacity(qreal opacity);
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

  // proto:  QTransform QGraphicsItem::sceneTransform();
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

  // proto:  void QGraphicsItem::setZValue(qreal z);
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

  // proto:  bool QGraphicsItem::isObscured(qreal x, qreal y, qreal w, qreal h);
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

  // proto:  void QGraphicsItem::installSceneEventFilter(QGraphicsItem * filterItem);
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

  // proto:  void QGraphicsItem::setY(qreal y);
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

  // proto:  QRectF QGraphicsItem::mapRectToItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
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

  // proto:  QGraphicsItem * QGraphicsItem::parentItem();
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

  // proto:  void QGraphicsItem::clearFocus();
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

  // proto:  bool QGraphicsItem::isWindow();
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

  // proto:  QPointF QGraphicsItem::transformOriginPoint();
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

  // proto:  QRectF QGraphicsItem::boundingRect();
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

  // proto:  QRectF QGraphicsItem::childrenBoundingRect();
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

  // proto:  QPolygonF QGraphicsItem::mapFromScene(const QRectF & rect);
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

  // proto:  bool QGraphicsItem::hasCursor();
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

  // proto:  void QGraphicsItem::setGraphicsEffect(QGraphicsEffect * effect);
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

  // proto:  void QGraphicsItem::ensureVisible(qreal x, qreal y, qreal w, qreal h, int xmargin, int ymargin);
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

  // proto:  QRectF QGraphicsItem::mapRectToParent(const QRectF & rect);
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

  // proto:  void QGraphicsItem::setToolTip(const QString & toolTip);
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

  // proto:  qreal QGraphicsItem::rotation();
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

  // proto:  QGraphicsScene * QGraphicsItem::scene();
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

  // proto:  QRectF QGraphicsItem::mapRectFromParent(const QRectF & rect);
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

  // proto:  void QGraphicsItem::setFocusProxy(QGraphicsItem * item);
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

  // proto:  bool QGraphicsItem::acceptDrops();
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

  // proto:  QRectF QGraphicsItem::mapRectFromScene(const QRectF & rect);
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

  // proto:  QGraphicsItem * QGraphicsItem::focusScopeItem();
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

  // proto:  void QGraphicsItem::removeSceneEventFilter(QGraphicsItem * filterItem);
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

  // proto:  QGraphicsItem * QGraphicsItem::focusProxy();
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

  // proto:  QRectF QGraphicsItem::sceneBoundingRect();
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

  // proto:  void QGraphicsItem::~QGraphicsItem();
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

  // proto:  void QGraphicsItem::setX(qreal x);
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

  // proto:  void QGraphicsItem::update(qreal x, qreal y, qreal width, qreal height);
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

  // proto:  void QGraphicsItem::setSelected(bool selected);
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

  // proto:  void QGraphicsItem::stackBefore(const QGraphicsItem * sibling);
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

  // proto:  void QGraphicsItem::resetMatrix();
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

  // proto:  QPainterPath QGraphicsItem::opaqueArea();
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

  // proto:  void QGraphicsItem::unsetCursor();
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

  // proto:  QRectF QGraphicsItem::mapRectToScene(const QRectF & rect);
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

  // proto:  QRectF QGraphicsItem::mapRectFromItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
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

  // proto:  qreal QGraphicsItem::scale();
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

  // proto:  void QGraphicsItem::setBoundingRegionGranularity(qreal granularity);
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

  // proto:  void QGraphicsItem::setAcceptDrops(bool on);
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

  // proto:  void QGraphicsItem::ungrabKeyboard();
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

  // proto:  void QGraphicsItem::setEnabled(bool enabled);
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

  // proto:  QGraphicsEffect * QGraphicsItem::graphicsEffect();
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

  // proto:  bool QGraphicsItem::acceptHoverEvents();
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

  // proto:  QGraphicsWidget * QGraphicsItem::topLevelWidget();
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

  // proto:  QList<QGraphicsTransform *> QGraphicsItem::transformations();
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

  // proto:  QPolygonF QGraphicsItem::mapToScene(qreal x, qreal y, qreal w, qreal h);
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

  // proto:  void QGraphicsItem::advance(int phase);
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

  // proto:  QMatrix QGraphicsItem::sceneMatrix();
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

  // proto:  void QGraphicsItem::setFiltersChildEvents(bool enabled);
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

  // proto:  QTransform QGraphicsItem::itemTransform(const QGraphicsItem * other, bool * ok);
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

  // proto:  void QGraphicsItem::moveBy(qreal dx, qreal dy);
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

  // proto:  QGraphicsItemGroup * QGraphicsItem::group();
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

  // proto:  QPainterPath QGraphicsItem::shape();
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

  // proto:  void QGraphicsItem::scroll(qreal dx, qreal dy, const QRectF & rect);
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

  // proto:  bool QGraphicsItem::isObscuredBy(const QGraphicsItem * item);
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

  // proto:  void QGraphicsItem::setData(int key, const QVariant & value);
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

  // proto:  QGraphicsItem * QGraphicsItem::commonAncestorItem(const QGraphicsItem * other);
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

  // proto:  void QGraphicsItem::setGroup(QGraphicsItemGroup * group);
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

  // proto:  void QGraphicsItem::show();
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

  // proto:  qreal QGraphicsItem::y();
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

  // proto:  bool QGraphicsItem::hasFocus();
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

  // proto:  QPainterPath QGraphicsItem::clipPath();
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

  // proto:  void QGraphicsItem::setPos(qreal x, qreal y);
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

  // proto:  bool QGraphicsItem::isEnabled();
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

  // proto:  bool QGraphicsItem::contains(const QPointF & point);
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

  // proto:  bool QGraphicsItem::isPanel();
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

  // proto:  bool QGraphicsItem::filtersChildEvents();
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

  // proto:  void QGraphicsItem::grabKeyboard();
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

  // proto:  void QGraphicsItem::setActive(bool active);
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

  // proto:  QGraphicsObject * QGraphicsItem::toGraphicsObject();
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

  // proto:  void QGraphicsItem::setHandlesChildEvents(bool enabled);
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

  // proto:  void QGraphicsItem::setMatrix(const QMatrix & matrix, bool combine);
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

  // proto:  QTransform QGraphicsItem::transform();
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

  // proto:  QVariant QGraphicsItem::data(int key);
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

  // proto:  void QGraphicsItem::hide();
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

  // proto:  bool QGraphicsItem::isUnderMouse();
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

  // proto:  void QGraphicsItem::setAcceptTouchEvents(bool enabled);
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

  // proto:  void QGraphicsItem::setAcceptHoverEvents(bool enabled);
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

  // proto:  QList<QGraphicsItem *> QGraphicsItem::childItems();
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

  // proto:  bool QGraphicsItem::isAncestorOf(const QGraphicsItem * child);
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

  // proto:  qreal QGraphicsItem::opacity();
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

  // proto:  bool QGraphicsItem::isVisibleTo(const QGraphicsItem * parent);
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

  // proto:  QString QGraphicsItem::toolTip();
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

  // proto:  QCursor QGraphicsItem::cursor();
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

  // proto:  qreal QGraphicsItem::zValue();
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

  // proto:  QMatrix QGraphicsItem::matrix();
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

  // proto:  QGraphicsItem * QGraphicsItem::panel();
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

  // proto:  bool QGraphicsItem::isClipped();
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

  // proto:  QGraphicsItem * QGraphicsItem::topLevelItem();
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

  // proto:  void QGraphicsItem::setScale(qreal scale);
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

  // proto:  void QGraphicsItem::setCursor(const QCursor & cursor);
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

  // proto:  bool QGraphicsItem::isVisible();
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

  // proto:  QPointF QGraphicsItem::pos();
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

  // proto:  bool QGraphicsItem::isBlockedByModalPanel(QGraphicsItem ** blockingPanel);
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

  // proto:  qreal QGraphicsItem::effectiveOpacity();
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

  // proto:  qreal QGraphicsItem::boundingRegionGranularity();
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

  // proto:  qreal QGraphicsItem::x();
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

  // proto:  void QGraphicsItem::grabMouse();
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

  // proto:  void QGraphicsItem::setVisible(bool visible);
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

  // proto:  void QGraphicsItem::setRotation(qreal angle);
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

  // proto:  QTransform QGraphicsItem::deviceTransform(const QTransform & viewportTransform);
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

  // proto:  bool QGraphicsItem::acceptTouchEvents();
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

  // proto:  void QGraphicsItem::setTransform(const QTransform & matrix, bool combine);
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

  // proto:  void QGraphicsObject::QGraphicsObject(QGraphicsItem * parent);
func NewQGraphicsObject(args ...interface{}) QGraphicsObject {
  return QGraphicsObject{}
}

  // proto:  void QGraphicsObject::~QGraphicsObject();
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

  // proto:  const QMetaObject * QGraphicsObject::metaObject();
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

  // proto:  int QGraphicsSimpleTextItem::type();
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

  // proto:  QFont QGraphicsSimpleTextItem::font();
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

  // proto:  void QGraphicsSimpleTextItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
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

  // proto:  void QGraphicsSimpleTextItem::~QGraphicsSimpleTextItem();
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

  // proto:  void QGraphicsSimpleTextItem::setText(const QString & text);
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

  // proto:  QString QGraphicsSimpleTextItem::text();
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

  // proto:  void QGraphicsSimpleTextItem::QGraphicsSimpleTextItem(const QString & text, QGraphicsItem * parent);
func NewQGraphicsSimpleTextItem(args ...interface{}) QGraphicsSimpleTextItem {
  return QGraphicsSimpleTextItem{}
}

  // proto:  bool QGraphicsSimpleTextItem::isObscuredBy(const QGraphicsItem * item);
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

  // proto:  QPainterPath QGraphicsSimpleTextItem::shape();
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

  // proto:  void QGraphicsSimpleTextItem::setFont(const QFont & font);
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

  // proto:  QPainterPath QGraphicsSimpleTextItem::opaqueArea();
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

  // proto:  QRectF QGraphicsSimpleTextItem::boundingRect();
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

  // proto:  bool QGraphicsSimpleTextItem::contains(const QPointF & point);
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

