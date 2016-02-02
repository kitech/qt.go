package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QGraphicsTextItem::contains(const QPointF & point);
extern bool C_ZNK17QGraphicsTextItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsTextItem::boundingRect();
extern void* C_ZNK17QGraphicsTextItem12boundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::setPlainText(const QString & text);
extern void C_ZN17QGraphicsTextItem12setPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsTextItem::setHtml(const QString & html);
extern void C_ZN17QGraphicsTextItem7setHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsTextItem::setTabChangesFocus(bool b);
extern void C_ZN17QGraphicsTextItem18setTabChangesFocusEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsTextItem::setDocument(QTextDocument * document);
extern void C_ZN17QGraphicsTextItem11setDocumentEP13QTextDocument(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsTextItem::setOpenExternalLinks(bool open);
extern void C_ZN17QGraphicsTextItem20setOpenExternalLinksEb(void* qthis, bool arg0); // 4
  // proto:  QFont QGraphicsTextItem::font();
extern void* C_ZNK17QGraphicsTextItem4fontEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::setDefaultTextColor(const QColor & c);
extern void C_ZN17QGraphicsTextItem19setDefaultTextColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsTextItem::openExternalLinks();
extern bool C_ZNK17QGraphicsTextItem17openExternalLinksEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::setTextCursor(const QTextCursor & cursor);
extern void C_ZN17QGraphicsTextItem13setTextCursorERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  QString QGraphicsTextItem::toPlainText();
extern void* C_ZNK17QGraphicsTextItem11toPlainTextEv(void* qthis); // 4
  // proto:  bool QGraphicsTextItem::tabChangesFocus();
extern bool C_ZNK17QGraphicsTextItem15tabChangesFocusEv(void* qthis); // 4
  // proto:  QTextCursor QGraphicsTextItem::textCursor();
extern void* C_ZNK17QGraphicsTextItem10textCursorEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN17QGraphicsTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QGraphicsTextItem::QGraphicsTextItem(QGraphicsItem * parent);
extern void* C_ZN17QGraphicsTextItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsTextItem::QGraphicsTextItem(const QString & text, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsTextItemC2ERK7QStringP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  int QGraphicsTextItem::type();
extern int32_t C_ZNK17QGraphicsTextItem4typeEv(void* qthis); // 4
  // proto:  Qt::TextInteractionFlags QGraphicsTextItem::textInteractionFlags();
extern void C_ZNK17QGraphicsTextItem20textInteractionFlagsEv(void* qthis); // 4
  // proto:  QString QGraphicsTextItem::toHtml();
extern void* C_ZNK17QGraphicsTextItem6toHtmlEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::adjustSize();
extern void C_ZN17QGraphicsTextItem10adjustSizeEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::~QGraphicsTextItem();
extern void C_ZN17QGraphicsTextItemD2Ev(void* qthis); // 4
  // proto:  void QGraphicsTextItem::setTextWidth(qreal width);
extern void C_ZN17QGraphicsTextItem12setTextWidthEd(void* qthis, double arg0); // 4
  // proto:  bool QGraphicsTextItem::isObscuredBy(const QGraphicsItem * item);
extern bool C_ZNK17QGraphicsTextItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  QTextDocument * QGraphicsTextItem::document();
extern void* C_ZNK17QGraphicsTextItem8documentEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsTextItem::opaqueArea();
extern void* C_ZNK17QGraphicsTextItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QColor QGraphicsTextItem::defaultTextColor();
extern void* C_ZNK17QGraphicsTextItem16defaultTextColorEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsTextItem::metaObject();
extern void C_ZNK17QGraphicsTextItem10metaObjectEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsTextItem::shape();
extern void* C_ZNK17QGraphicsTextItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsTextItem::setFont(const QFont & font);
extern void C_ZN17QGraphicsTextItem7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsTextItem::textWidth();
extern double C_ZNK17QGraphicsTextItem9textWidthEv(void* qthis); // 4
  // proto:  QRectF QGraphicsPixmapItem::boundingRect();
extern void* C_ZNK19QGraphicsPixmapItem12boundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsPixmapItem::setPixmap(const QPixmap & pixmap);
extern void C_ZN19QGraphicsPixmapItem9setPixmapERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsPixmapItem::shape();
extern void* C_ZNK19QGraphicsPixmapItem5shapeEv(void* qthis); // 4
  // proto:  Qt::TransformationMode QGraphicsPixmapItem::transformationMode();
extern void C_ZNK19QGraphicsPixmapItem18transformationModeEv(void* qthis); // 4
  // proto:  QPixmap QGraphicsPixmapItem::pixmap();
extern void* C_ZNK19QGraphicsPixmapItem6pixmapEv(void* qthis); // 4
  // proto:  bool QGraphicsPixmapItem::contains(const QPointF & point);
extern bool C_ZNK19QGraphicsPixmapItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QGraphicsPixmapItem::ShapeMode QGraphicsPixmapItem::shapeMode();
extern void C_ZNK19QGraphicsPixmapItem9shapeModeEv(void* qthis); // 4
  // proto:  void QGraphicsPixmapItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN19QGraphicsPixmapItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QGraphicsPixmapItem::setOffset(const QPointF & offset);
extern void C_ZN19QGraphicsPixmapItem9setOffsetERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsPixmapItem::setOffset(qreal x, qreal y);
extern void C_ZN19QGraphicsPixmapItem9setOffsetEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  int QGraphicsPixmapItem::type();
extern int32_t C_ZNK19QGraphicsPixmapItem4typeEv(void* qthis); // 4
  // proto:  void QGraphicsPixmapItem::~QGraphicsPixmapItem();
extern void C_ZN19QGraphicsPixmapItemD2Ev(void* qthis); // 4
  // proto:  QPointF QGraphicsPixmapItem::offset();
extern void* C_ZNK19QGraphicsPixmapItem6offsetEv(void* qthis); // 4
  // proto:  void QGraphicsPixmapItem::QGraphicsPixmapItem(QGraphicsItem * parent);
extern void* C_ZN19QGraphicsPixmapItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsPixmapItem::QGraphicsPixmapItem(const QPixmap & pixmap, QGraphicsItem * parent);
extern void* C_ZN19QGraphicsPixmapItemC2ERK7QPixmapP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  bool QGraphicsPixmapItem::isObscuredBy(const QGraphicsItem * item);
extern bool C_ZNK19QGraphicsPixmapItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsPixmapItem::opaqueArea();
extern void* C_ZNK19QGraphicsPixmapItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsRectItem::opaqueArea();
extern void* C_ZNK17QGraphicsRectItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QRectF QGraphicsRectItem::boundingRect();
extern void* C_ZNK17QGraphicsRectItem12boundingRectEv(void* qthis); // 4
  // proto:  int QGraphicsRectItem::type();
extern int32_t C_ZNK17QGraphicsRectItem4typeEv(void* qthis); // 4
  // proto:  bool QGraphicsRectItem::contains(const QPointF & point);
extern bool C_ZNK17QGraphicsRectItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsRectItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPainterPath QGraphicsRectItem::shape();
extern void* C_ZNK17QGraphicsRectItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsRectItem::QGraphicsRectItem(QGraphicsItem * parent);
extern void* C_ZN17QGraphicsRectItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsRectItem::QGraphicsRectItem(const QRectF & rect, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsRectItemC2ERK6QRectFP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  void QGraphicsRectItem::QGraphicsRectItem(qreal x, qreal y, qreal w, qreal h, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsRectItemC2EddddP13QGraphicsItem(double arg0, double arg1, double arg2, double arg3, void* arg4); // 3
  // proto:  void QGraphicsRectItem::~QGraphicsRectItem();
extern void C_ZN17QGraphicsRectItemD2Ev(void* qthis); // 4
  // proto:  bool QGraphicsRectItem::isObscuredBy(const QGraphicsItem * item);
extern bool C_ZNK17QGraphicsRectItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsRectItem::setRect(const QRectF & rect);
extern void C_ZN17QGraphicsRectItem7setRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsRectItem::setRect(qreal x, qreal y, qreal w, qreal h);
extern void C_ZN17QGraphicsRectItem7setRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QRectF QGraphicsRectItem::rect();
extern void* C_ZNK17QGraphicsRectItem4rectEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsEllipseItem::opaqueArea();
extern void* C_ZNK20QGraphicsEllipseItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QRectF QGraphicsEllipseItem::boundingRect();
extern void* C_ZNK20QGraphicsEllipseItem12boundingRectEv(void* qthis); // 4
  // proto:  int QGraphicsEllipseItem::spanAngle();
extern int32_t C_ZNK20QGraphicsEllipseItem9spanAngleEv(void* qthis); // 4
  // proto:  int QGraphicsEllipseItem::type();
extern int32_t C_ZNK20QGraphicsEllipseItem4typeEv(void* qthis); // 4
  // proto:  bool QGraphicsEllipseItem::contains(const QPointF & point);
extern bool C_ZNK20QGraphicsEllipseItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsEllipseItem::setSpanAngle(int angle);
extern void C_ZN20QGraphicsEllipseItem12setSpanAngleEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsEllipseItem::setRect(const QRectF & rect);
extern void C_ZN20QGraphicsEllipseItem7setRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsEllipseItem::setRect(qreal x, qreal y, qreal w, qreal h);
extern void C_ZN20QGraphicsEllipseItem7setRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QGraphicsEllipseItem::setStartAngle(int angle);
extern void C_ZN20QGraphicsEllipseItem13setStartAngleEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsEllipseItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN20QGraphicsEllipseItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPainterPath QGraphicsEllipseItem::shape();
extern void* C_ZNK20QGraphicsEllipseItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsEllipseItem::~QGraphicsEllipseItem();
extern void C_ZN20QGraphicsEllipseItemD2Ev(void* qthis); // 4
  // proto:  int QGraphicsEllipseItem::startAngle();
extern int32_t C_ZNK20QGraphicsEllipseItem10startAngleEv(void* qthis); // 4
  // proto:  bool QGraphicsEllipseItem::isObscuredBy(const QGraphicsItem * item);
extern bool C_ZNK20QGraphicsEllipseItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsEllipseItem::QGraphicsEllipseItem(QGraphicsItem * parent);
extern void* C_ZN20QGraphicsEllipseItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsEllipseItem::QGraphicsEllipseItem(qreal x, qreal y, qreal w, qreal h, QGraphicsItem * parent);
extern void* C_ZN20QGraphicsEllipseItemC2EddddP13QGraphicsItem(double arg0, double arg1, double arg2, double arg3, void* arg4); // 3
  // proto:  void QGraphicsEllipseItem::QGraphicsEllipseItem(const QRectF & rect, QGraphicsItem * parent);
extern void* C_ZN20QGraphicsEllipseItemC2ERK6QRectFP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  QRectF QGraphicsEllipseItem::rect();
extern void* C_ZNK20QGraphicsEllipseItem4rectEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsPolygonItem::opaqueArea();
extern void* C_ZNK20QGraphicsPolygonItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QRectF QGraphicsPolygonItem::boundingRect();
extern void* C_ZNK20QGraphicsPolygonItem12boundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsPolygonItem::setPolygon(const QPolygonF & polygon);
extern void C_ZN20QGraphicsPolygonItem10setPolygonERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  int QGraphicsPolygonItem::type();
extern int32_t C_ZNK20QGraphicsPolygonItem4typeEv(void* qthis); // 4
  // proto:  bool QGraphicsPolygonItem::contains(const QPointF & point);
extern bool C_ZNK20QGraphicsPolygonItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsPolygonItem::polygon();
extern void* C_ZNK20QGraphicsPolygonItem7polygonEv(void* qthis); // 4
  // proto:  void QGraphicsPolygonItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN20QGraphicsPolygonItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPainterPath QGraphicsPolygonItem::shape();
extern void* C_ZNK20QGraphicsPolygonItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsPolygonItem::QGraphicsPolygonItem(const QPolygonF & polygon, QGraphicsItem * parent);
extern void* C_ZN20QGraphicsPolygonItemC2ERK9QPolygonFP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  void QGraphicsPolygonItem::QGraphicsPolygonItem(QGraphicsItem * parent);
extern void* C_ZN20QGraphicsPolygonItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsPolygonItem::~QGraphicsPolygonItem();
extern void C_ZN20QGraphicsPolygonItemD2Ev(void* qthis); // 4
  // proto:  bool QGraphicsPolygonItem::isObscuredBy(const QGraphicsItem * item);
extern bool C_ZNK20QGraphicsPolygonItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  Qt::FillRule QGraphicsPolygonItem::fillRule();
extern void C_ZNK20QGraphicsPolygonItem8fillRuleEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsPathItem::opaqueArea();
extern void* C_ZNK17QGraphicsPathItem10opaqueAreaEv(void* qthis); // 4
  // proto:  void QGraphicsPathItem::~QGraphicsPathItem();
extern void C_ZN17QGraphicsPathItemD2Ev(void* qthis); // 4
  // proto:  QRectF QGraphicsPathItem::boundingRect();
extern void* C_ZNK17QGraphicsPathItem12boundingRectEv(void* qthis); // 4
  // proto:  bool QGraphicsPathItem::isObscuredBy(const QGraphicsItem * item);
extern bool C_ZNK17QGraphicsPathItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsPathItem::contains(const QPointF & point);
extern bool C_ZNK17QGraphicsPathItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsPathItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN17QGraphicsPathItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPainterPath QGraphicsPathItem::shape();
extern void* C_ZNK17QGraphicsPathItem5shapeEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsPathItem::path();
extern void* C_ZNK17QGraphicsPathItem4pathEv(void* qthis); // 4
  // proto:  void QGraphicsPathItem::setPath(const QPainterPath & path);
extern void C_ZN17QGraphicsPathItem7setPathERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsPathItem::QGraphicsPathItem(const QPainterPath & path, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsPathItemC2ERK12QPainterPathP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  void QGraphicsPathItem::QGraphicsPathItem(QGraphicsItem * parent);
extern void* C_ZN17QGraphicsPathItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  int QGraphicsPathItem::type();
extern int32_t C_ZNK17QGraphicsPathItem4typeEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsLineItem::opaqueArea();
extern void* C_ZNK17QGraphicsLineItem10opaqueAreaEv(void* qthis); // 4
  // proto:  QRectF QGraphicsLineItem::boundingRect();
extern void* C_ZNK17QGraphicsLineItem12boundingRectEv(void* qthis); // 4
  // proto:  int QGraphicsLineItem::type();
extern int32_t C_ZNK17QGraphicsLineItem4typeEv(void* qthis); // 4
  // proto:  bool QGraphicsLineItem::contains(const QPointF & point);
extern bool C_ZNK17QGraphicsLineItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsLineItem::shape();
extern void* C_ZNK17QGraphicsLineItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsLineItem::~QGraphicsLineItem();
extern void C_ZN17QGraphicsLineItemD2Ev(void* qthis); // 4
  // proto:  void QGraphicsLineItem::QGraphicsLineItem(QGraphicsItem * parent);
extern void* C_ZN17QGraphicsLineItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsLineItem::QGraphicsLineItem(const QLineF & line, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsLineItemC2ERK6QLineFP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  void QGraphicsLineItem::QGraphicsLineItem(qreal x1, qreal y1, qreal x2, qreal y2, QGraphicsItem * parent);
extern void* C_ZN17QGraphicsLineItemC2EddddP13QGraphicsItem(double arg0, double arg1, double arg2, double arg3, void* arg4); // 3
  // proto:  void QGraphicsLineItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN17QGraphicsLineItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPen QGraphicsLineItem::pen();
extern void* C_ZNK17QGraphicsLineItem3penEv(void* qthis); // 4
  // proto:  void QGraphicsLineItem::setPen(const QPen & pen);
extern void C_ZN17QGraphicsLineItem6setPenERK4QPen(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLineItem::setLine(const QLineF & line);
extern void C_ZN17QGraphicsLineItem7setLineERK6QLineF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsLineItem::setLine(qreal x1, qreal y1, qreal x2, qreal y2);
extern void C_ZN17QGraphicsLineItem7setLineEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QLineF QGraphicsLineItem::line();
extern void* C_ZNK17QGraphicsLineItem4lineEv(void* qthis); // 4
  // proto:  bool QGraphicsLineItem::isObscuredBy(const QGraphicsItem * item);
extern bool C_ZNK17QGraphicsLineItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsItemGroup::opaqueArea();
extern void* C_ZNK18QGraphicsItemGroup10opaqueAreaEv(void* qthis); // 4
  // proto:  void QGraphicsItemGroup::addToGroup(QGraphicsItem * item);
extern void C_ZN18QGraphicsItemGroup10addToGroupEP13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItemGroup::QGraphicsItemGroup(QGraphicsItem * parent);
extern void* C_ZN18QGraphicsItemGroupC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  QRectF QGraphicsItemGroup::boundingRect();
extern void* C_ZNK18QGraphicsItemGroup12boundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsItemGroup::removeFromGroup(QGraphicsItem * item);
extern void C_ZN18QGraphicsItemGroup15removeFromGroupEP13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsItemGroup::isObscuredBy(const QGraphicsItem * item);
extern bool C_ZNK18QGraphicsItemGroup12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItemGroup::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  int QGraphicsItemGroup::type();
extern int32_t C_ZNK18QGraphicsItemGroup4typeEv(void* qthis); // 4
  // proto:  void QGraphicsItemGroup::~QGraphicsItemGroup();
extern void C_ZN18QGraphicsItemGroupD2Ev(void* qthis); // 4
  // proto:  QPainterPath QAbstractGraphicsShapeItem::opaqueArea();
extern void* C_ZNK26QAbstractGraphicsShapeItem10opaqueAreaEv(void* qthis); // 4
  // proto:  bool QAbstractGraphicsShapeItem::isObscuredBy(const QGraphicsItem * item);
extern bool C_ZNK26QAbstractGraphicsShapeItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  QPen QAbstractGraphicsShapeItem::pen();
extern void* C_ZNK26QAbstractGraphicsShapeItem3penEv(void* qthis); // 4
  // proto:  QBrush QAbstractGraphicsShapeItem::brush();
extern void* C_ZNK26QAbstractGraphicsShapeItem5brushEv(void* qthis); // 4
  // proto:  void QAbstractGraphicsShapeItem::setPen(const QPen & pen);
extern void C_ZN26QAbstractGraphicsShapeItem6setPenERK4QPen(void* qthis, void* arg0); // 4
  // proto:  void QAbstractGraphicsShapeItem::setBrush(const QBrush & brush);
extern void C_ZN26QAbstractGraphicsShapeItem8setBrushERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  void QAbstractGraphicsShapeItem::~QAbstractGraphicsShapeItem();
extern void C_ZN26QAbstractGraphicsShapeItemD2Ev(void* qthis); // 4
  // proto:  void QAbstractGraphicsShapeItem::QAbstractGraphicsShapeItem(QGraphicsItem * parent);
extern void* C_ZN26QAbstractGraphicsShapeItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsItem::show();
extern void C_ZN13QGraphicsItem4showEv(void* qthis); // 2
  // proto:  bool QGraphicsItem::hasFocus();
extern bool C_ZNK13QGraphicsItem8hasFocusEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setGroup(QGraphicsItemGroup * group);
extern void C_ZN13QGraphicsItem8setGroupEP18QGraphicsItemGroup(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsItem::handlesChildEvents();
extern bool C_ZNK13QGraphicsItem18handlesChildEventsEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsItem::shape();
extern void* C_ZNK13QGraphicsItem5shapeEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isPanel();
extern bool C_ZNK13QGraphicsItem7isPanelEv(void* qthis); // 4
  // proto:  QVariant QGraphicsItem::data(int key);
extern void* C_ZNK13QGraphicsItem4dataEi(void* qthis, int32_t arg0); // 4
  // proto:  QGraphicsItem * QGraphicsItem::parentItem();
extern void C_ZNK13QGraphicsItem10parentItemEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::childrenBoundingRect();
extern void* C_ZNK13QGraphicsItem20childrenBoundingRectEv(void* qthis); // 4
  // proto:  QGraphicsItemGroup * QGraphicsItem::group();
extern void C_ZNK13QGraphicsItem5groupEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isEnabled();
extern bool C_ZNK13QGraphicsItem9isEnabledEv(void* qthis); // 4
  // proto:  QGraphicsWidget * QGraphicsItem::window();
extern void C_ZNK13QGraphicsItem6windowEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setBoundingRegionGranularity(qreal granularity);
extern void C_ZN13QGraphicsItem28setBoundingRegionGranularityEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItem::setRotation(qreal angle);
extern void C_ZN13QGraphicsItem11setRotationEd(void* qthis, double arg0); // 4
  // proto:  QGraphicsItem * QGraphicsItem::commonAncestorItem(const QGraphicsItem * other);
extern void C_ZNK13QGraphicsItem18commonAncestorItemEPKS_(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToItem(const QGraphicsItem * item, const QPolygonF & polygon);
extern void* C_ZNK13QGraphicsItem9mapToItemEPKS_RK9QPolygonF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPolygonF QGraphicsItem::mapToItem(const QGraphicsItem * item, const QRectF & rect);
extern void* C_ZNK13QGraphicsItem9mapToItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPainterPath QGraphicsItem::mapToItem(const QGraphicsItem * item, const QPainterPath & path);
extern void* C_ZNK13QGraphicsItem9mapToItemEPKS_RK12QPainterPath(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPointF QGraphicsItem::mapToItem(const QGraphicsItem * item, const QPointF & point);
extern void* C_ZNK13QGraphicsItem9mapToItemEPKS_RK7QPointF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPolygonF QGraphicsItem::mapToItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem9mapToItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4); // 2
  // proto:  QPointF QGraphicsItem::mapToItem(const QGraphicsItem * item, qreal x, qreal y);
extern void* C_ZNK13QGraphicsItem9mapToItemEPKS_dd(void* qthis, void* arg0, double arg1, double arg2); // 2
  // proto:  QGraphicsObject * QGraphicsItem::toGraphicsObject();
extern void C_ZN13QGraphicsItem16toGraphicsObjectEv(void* qthis); // 4
  // proto:  Qt::MouseButtons QGraphicsItem::acceptedMouseButtons();
extern void C_ZNK13QGraphicsItem20acceptedMouseButtonsEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setGraphicsEffect(QGraphicsEffect * effect);
extern void C_ZN13QGraphicsItem17setGraphicsEffectEP15QGraphicsEffect(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToParent(qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem11mapToParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QPointF QGraphicsItem::mapToParent(const QPointF & point);
extern void* C_ZNK13QGraphicsItem11mapToParentERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsItem::mapToParent(qreal x, qreal y);
extern void* C_ZNK13QGraphicsItem11mapToParentEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QPainterPath QGraphicsItem::mapToParent(const QPainterPath & path);
extern void* C_ZNK13QGraphicsItem11mapToParentERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToParent(const QRectF & rect);
extern void* C_ZNK13QGraphicsItem11mapToParentERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToParent(const QPolygonF & polygon);
extern void* C_ZNK13QGraphicsItem11mapToParentERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::setTransformOriginPoint(const QPointF & origin);
extern void C_ZN13QGraphicsItem23setTransformOriginPointERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::setTransformOriginPoint(qreal ax, qreal ay);
extern void C_ZN13QGraphicsItem23setTransformOriginPointEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QTransform QGraphicsItem::deviceTransform(const QTransform & viewportTransform);
extern void* C_ZNK13QGraphicsItem15deviceTransformERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsItem::rotation();
extern double C_ZNK13QGraphicsItem8rotationEv(void* qthis); // 4
  // proto:  qreal QGraphicsItem::effectiveOpacity();
extern double C_ZNK13QGraphicsItem16effectiveOpacityEv(void* qthis); // 4
  // proto:  QGraphicsItem * QGraphicsItem::focusScopeItem();
extern void C_ZNK13QGraphicsItem14focusScopeItemEv(void* qthis); // 4
  // proto:  Qt::InputMethodHints QGraphicsItem::inputMethodHints();
extern void C_ZNK13QGraphicsItem16inputMethodHintsEv(void* qthis); // 4
  // proto:  QGraphicsItem::PanelModality QGraphicsItem::panelModality();
extern void C_ZNK13QGraphicsItem13panelModalityEv(void* qthis); // 4
  // proto:  QGraphicsWidget * QGraphicsItem::topLevelWidget();
extern void C_ZNK13QGraphicsItem14topLevelWidgetEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isVisibleTo(const QGraphicsItem * parent);
extern bool C_ZNK13QGraphicsItem11isVisibleToEPKS_(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::setAcceptDrops(bool on);
extern void C_ZN13QGraphicsItem14setAcceptDropsEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsItem::clearFocus();
extern void C_ZN13QGraphicsItem10clearFocusEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::filtersChildEvents();
extern bool C_ZNK13QGraphicsItem18filtersChildEventsEv(void* qthis); // 4
  // proto:  qreal QGraphicsItem::x();
extern void C_ZNK13QGraphicsItem1xEv(void* qthis); // 2
  // proto:  void QGraphicsItem::setTransform(const QTransform & matrix, bool combine);
extern void C_ZN13QGraphicsItem12setTransformERK10QTransformb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QGraphicsItem::stackBefore(const QGraphicsItem * sibling);
extern void C_ZN13QGraphicsItem11stackBeforeEPKS_(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::setAcceptHoverEvents(bool enabled);
extern void C_ZN13QGraphicsItem20setAcceptHoverEventsEb(void* qthis, bool arg0); // 4
  // proto:  QGraphicsItem * QGraphicsItem::focusProxy();
extern void C_ZNK13QGraphicsItem10focusProxyEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsItem::mapFromParent(const QPainterPath & path);
extern void* C_ZNK13QGraphicsItem13mapFromParentERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromParent(const QRectF & rect);
extern void* C_ZNK13QGraphicsItem13mapFromParentERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromParent(qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem13mapFromParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QPointF QGraphicsItem::mapFromParent(qreal x, qreal y);
extern void* C_ZNK13QGraphicsItem13mapFromParentEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QPointF QGraphicsItem::mapFromParent(const QPointF & point);
extern void* C_ZNK13QGraphicsItem13mapFromParentERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromParent(const QPolygonF & polygon);
extern void* C_ZNK13QGraphicsItem13mapFromParentERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsItem::scenePos();
extern void* C_ZNK13QGraphicsItem8scenePosEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setPos(qreal x, qreal y);
extern void C_ZN13QGraphicsItem6setPosEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsItem::setPos(const QPointF & pos);
extern void C_ZN13QGraphicsItem6setPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsItem::scale();
extern double C_ZNK13QGraphicsItem5scaleEv(void* qthis); // 4
  // proto:  void QGraphicsItem::hide();
extern void C_ZN13QGraphicsItem4hideEv(void* qthis); // 2
  // proto:  QMatrix QGraphicsItem::matrix();
extern void* C_ZNK13QGraphicsItem6matrixEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isSelected();
extern bool C_ZNK13QGraphicsItem10isSelectedEv(void* qthis); // 4
  // proto:  void QGraphicsItem::moveBy(qreal dx, qreal dy);
extern void C_ZN13QGraphicsItem6moveByEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsItem::setFocusProxy(QGraphicsItem * item);
extern void C_ZN13QGraphicsItem13setFocusProxyEPS_(void* qthis, void* arg0); // 4
  // proto:  QTransform QGraphicsItem::itemTransform(const QGraphicsItem * other, bool * ok);
extern void* C_ZNK13QGraphicsItem13itemTransformEPKS_Pb(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QRectF QGraphicsItem::mapRectToParent(const QRectF & rect);
extern void* C_ZNK13QGraphicsItem15mapRectToParentERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsItem::mapRectToParent(qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem15mapRectToParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  bool QGraphicsItem::isBlockedByModalPanel(QGraphicsItem ** blockingPanel);
extern bool C_ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_(void* qthis, void* arg0); // 4
  // proto:  QGraphicsEffect * QGraphicsItem::graphicsEffect();
extern void C_ZNK13QGraphicsItem14graphicsEffectEv(void* qthis); // 4
  // proto:  QList<QGraphicsItem *> QGraphicsItem::childItems();
extern void C_ZNK13QGraphicsItem10childItemsEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isClipped();
extern bool C_ZNK13QGraphicsItem9isClippedEv(void* qthis); // 4
  // proto:  QString QGraphicsItem::toolTip();
extern void* C_ZNK13QGraphicsItem7toolTipEv(void* qthis); // 4
  // proto:  void QGraphicsItem::~QGraphicsItem();
extern void C_ZN13QGraphicsItemD2Ev(void* qthis); // 4
  // proto:  bool QGraphicsItem::isActive();
extern bool C_ZNK13QGraphicsItem8isActiveEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setFiltersChildEvents(bool enabled);
extern void C_ZN13QGraphicsItem21setFiltersChildEventsEb(void* qthis, bool arg0); // 4
  // proto:  QPainterPath QGraphicsItem::opaqueArea();
extern void* C_ZNK13QGraphicsItem10opaqueAreaEv(void* qthis); // 4
  // proto:  void QGraphicsItem::advance(int phase);
extern void C_ZN13QGraphicsItem7advanceEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGraphicsItem::setToolTip(const QString & toolTip);
extern void C_ZN13QGraphicsItem10setToolTipERK7QString(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToScene(const QPolygonF & polygon);
extern void* C_ZNK13QGraphicsItem10mapToSceneERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToScene(const QRectF & rect);
extern void* C_ZNK13QGraphicsItem10mapToSceneERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsItem::mapToScene(const QPointF & point);
extern void* C_ZNK13QGraphicsItem10mapToSceneERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsItem::mapToScene(const QPainterPath & path);
extern void* C_ZNK13QGraphicsItem10mapToSceneERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapToScene(qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem10mapToSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QPointF QGraphicsItem::mapToScene(qreal x, qreal y);
extern void* C_ZNK13QGraphicsItem10mapToSceneEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsItem::setHandlesChildEvents(bool enabled);
extern void C_ZN13QGraphicsItem21setHandlesChildEventsEb(void* qthis, bool arg0); // 4
  // proto:  QPointF QGraphicsItem::transformOriginPoint();
extern void* C_ZNK13QGraphicsItem20transformOriginPointEv(void* qthis); // 4
  // proto:  QPointF QGraphicsItem::pos();
extern void* C_ZNK13QGraphicsItem3posEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::mapRectToScene(const QRectF & rect);
extern void* C_ZNK13QGraphicsItem14mapRectToSceneERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsItem::mapRectToScene(qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem14mapRectToSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QGraphicsItem::ungrabMouse();
extern void C_ZN13QGraphicsItem11ungrabMouseEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isObscured(qreal x, qreal y, qreal w, qreal h);
extern bool C_ZNK13QGraphicsItem10isObscuredEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  bool QGraphicsItem::isObscured(const QRectF & rect);
extern bool C_ZNK13QGraphicsItem10isObscuredERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::setSelected(bool selected);
extern void C_ZN13QGraphicsItem11setSelectedEb(void* qthis, bool arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromScene(const QPolygonF & polygon);
extern void* C_ZNK13QGraphicsItem12mapFromSceneERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromScene(const QRectF & rect);
extern void* C_ZNK13QGraphicsItem12mapFromSceneERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromScene(qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem12mapFromSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QPointF QGraphicsItem::mapFromScene(qreal x, qreal y);
extern void* C_ZNK13QGraphicsItem12mapFromSceneEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QPainterPath QGraphicsItem::mapFromScene(const QPainterPath & path);
extern void* C_ZNK13QGraphicsItem12mapFromSceneERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsItem::mapFromScene(const QPointF & point);
extern void* C_ZNK13QGraphicsItem12mapFromSceneERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QGraphicsScene * QGraphicsItem::scene();
extern void C_ZNK13QGraphicsItem5sceneEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isAncestorOf(const QGraphicsItem * child);
extern bool C_ZNK13QGraphicsItem12isAncestorOfEPKS_(void* qthis, void* arg0); // 4
  // proto:  QGraphicsItem * QGraphicsItem::topLevelItem();
extern void C_ZNK13QGraphicsItem12topLevelItemEv(void* qthis); // 4
  // proto:  QGraphicsItem * QGraphicsItem::panel();
extern void C_ZNK13QGraphicsItem5panelEv(void* qthis); // 4
  // proto:  void QGraphicsItem::unsetCursor();
extern void C_ZN13QGraphicsItem11unsetCursorEv(void* qthis); // 4
  // proto:  QList<QGraphicsTransform *> QGraphicsItem::transformations();
extern void C_ZNK13QGraphicsItem15transformationsEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromScene(const QRectF & rect);
extern void* C_ZNK13QGraphicsItem16mapRectFromSceneERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromScene(qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem16mapRectFromSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  bool QGraphicsItem::contains(const QPointF & point);
extern bool C_ZNK13QGraphicsItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QTransform QGraphicsItem::transform();
extern void* C_ZNK13QGraphicsItem9transformEv(void* qthis); // 4
  // proto:  void QGraphicsItem::ungrabKeyboard();
extern void C_ZN13QGraphicsItem14ungrabKeyboardEv(void* qthis); // 4
  // proto:  void QGraphicsItem::installSceneEventFilter(QGraphicsItem * filterItem);
extern void C_ZN13QGraphicsItem23installSceneEventFilterEPS_(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsItem::acceptHoverEvents();
extern bool C_ZNK13QGraphicsItem17acceptHoverEventsEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::mapRectToItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem13mapRectToItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4); // 2
  // proto:  QRectF QGraphicsItem::mapRectToItem(const QGraphicsItem * item, const QRectF & rect);
extern void* C_ZNK13QGraphicsItem13mapRectToItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPointF QGraphicsItem::mapFromItem(const QGraphicsItem * item, qreal x, qreal y);
extern void* C_ZNK13QGraphicsItem11mapFromItemEPKS_dd(void* qthis, void* arg0, double arg1, double arg2); // 2
  // proto:  QPointF QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QPointF & point);
extern void* C_ZNK13QGraphicsItem11mapFromItemEPKS_RK7QPointF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem11mapFromItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4); // 2
  // proto:  QPolygonF QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QRectF & rect);
extern void* C_ZNK13QGraphicsItem11mapFromItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPolygonF QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QPolygonF & polygon);
extern void* C_ZNK13QGraphicsItem11mapFromItemEPKS_RK9QPolygonF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QPainterPath QGraphicsItem::mapFromItem(const QGraphicsItem * item, const QPainterPath & path);
extern void* C_ZNK13QGraphicsItem11mapFromItemEPKS_RK12QPainterPath(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QGraphicsItem::QGraphicsItem(QGraphicsItem * parent);
extern void* C_ZN13QGraphicsItemC2EPS_(void* arg0); // 3
  // proto:  int QGraphicsItem::type();
extern int32_t C_ZNK13QGraphicsItem4typeEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setEnabled(bool enabled);
extern void C_ZN13QGraphicsItem10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsItem::grabMouse();
extern void C_ZN13QGraphicsItem9grabMouseEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setActive(bool active);
extern void C_ZN13QGraphicsItem9setActiveEb(void* qthis, bool arg0); // 4
  // proto:  bool QGraphicsItem::isObscuredBy(const QGraphicsItem * item);
extern bool C_ZNK13QGraphicsItem12isObscuredByEPKS_(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsItem::acceptDrops();
extern bool C_ZNK13QGraphicsItem11acceptDropsEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setParentItem(QGraphicsItem * parent);
extern void C_ZN13QGraphicsItem13setParentItemEPS_(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItem::ensureVisible(qreal x, qreal y, qreal w, qreal h, int xmargin, int ymargin);
extern void C_ZN13QGraphicsItem13ensureVisibleEddddii(void* qthis, double arg0, double arg1, double arg2, double arg3, int32_t arg4, int32_t arg5); // 2
  // proto:  void QGraphicsItem::ensureVisible(const QRectF & rect, int xmargin, int ymargin);
extern void C_ZN13QGraphicsItem13ensureVisibleERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QGraphicsItem::grabKeyboard();
extern void C_ZN13QGraphicsItem12grabKeyboardEv(void* qthis); // 4
  // proto:  QGraphicsObject * QGraphicsItem::parentObject();
extern void C_ZNK13QGraphicsItem12parentObjectEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isUnderMouse();
extern bool C_ZNK13QGraphicsItem12isUnderMouseEv(void* qthis); // 4
  // proto:  QGraphicsWidget * QGraphicsItem::parentWidget();
extern void C_ZNK13QGraphicsItem12parentWidgetEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setZValue(qreal z);
extern void C_ZN13QGraphicsItem9setZValueEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItem::setMatrix(const QMatrix & matrix, bool combine);
extern void C_ZN13QGraphicsItem9setMatrixERK7QMatrixb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  QCursor QGraphicsItem::cursor();
extern void* C_ZNK13QGraphicsItem6cursorEv(void* qthis); // 4
  // proto:  qreal QGraphicsItem::zValue();
extern double C_ZNK13QGraphicsItem6zValueEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setVisible(bool visible);
extern void C_ZN13QGraphicsItem10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsItem::resetMatrix();
extern void C_ZN13QGraphicsItem11resetMatrixEv(void* qthis); // 4
  // proto:  QTransform QGraphicsItem::sceneTransform();
extern void* C_ZNK13QGraphicsItem14sceneTransformEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isWindow();
extern bool C_ZNK13QGraphicsItem8isWindowEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromParent(const QRectF & rect);
extern void* C_ZNK13QGraphicsItem17mapRectFromParentERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromParent(qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem17mapRectFromParentEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QGraphicsItem::setScale(qreal scale);
extern void C_ZN13QGraphicsItem8setScaleEd(void* qthis, double arg0); // 4
  // proto:  bool QGraphicsItem::isWidget();
extern bool C_ZNK13QGraphicsItem8isWidgetEv(void* qthis); // 4
  // proto:  void QGraphicsItem::resetTransform();
extern void C_ZN13QGraphicsItem14resetTransformEv(void* qthis); // 4
  // proto:  QRectF QGraphicsItem::sceneBoundingRect();
extern void* C_ZNK13QGraphicsItem17sceneBoundingRectEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setOpacity(qreal opacity);
extern void C_ZN13QGraphicsItem10setOpacityEd(void* qthis, double arg0); // 4
  // proto:  QGraphicsItem * QGraphicsItem::focusItem();
extern void C_ZNK13QGraphicsItem9focusItemEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::hasCursor();
extern bool C_ZNK13QGraphicsItem9hasCursorEv(void* qthis); // 4
  // proto:  void QGraphicsItem::setAcceptTouchEvents(bool enabled);
extern void C_ZN13QGraphicsItem20setAcceptTouchEventsEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsItem::setData(int key, const QVariant & value);
extern void C_ZN13QGraphicsItem7setDataEiRK8QVariant(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  qreal QGraphicsItem::opacity();
extern double C_ZNK13QGraphicsItem7opacityEv(void* qthis); // 4
  // proto:  bool QGraphicsItem::isVisible();
extern bool C_ZNK13QGraphicsItem9isVisibleEv(void* qthis); // 4
  // proto:  void QGraphicsItem::update(qreal x, qreal y, qreal width, qreal height);
extern void C_ZN13QGraphicsItem6updateEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QGraphicsItem::update(const QRectF & rect);
extern void C_ZN13QGraphicsItem6updateERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromItem(const QGraphicsItem * item, const QRectF & rect);
extern void* C_ZNK13QGraphicsItem15mapRectFromItemEPKS_RK6QRectF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QRectF QGraphicsItem::mapRectFromItem(const QGraphicsItem * item, qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsItem15mapRectFromItemEPKS_dddd(void* qthis, void* arg0, double arg1, double arg2, double arg3, double arg4); // 2
  // proto:  void QGraphicsItem::setX(qreal x);
extern void C_ZN13QGraphicsItem4setXEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItem::setY(qreal y);
extern void C_ZN13QGraphicsItem4setYEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItem::setCursor(const QCursor & cursor);
extern void C_ZN13QGraphicsItem9setCursorERK7QCursor(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsItem::boundingRegionGranularity();
extern double C_ZNK13QGraphicsItem25boundingRegionGranularityEv(void* qthis); // 4
  // proto:  void QGraphicsItem::removeSceneEventFilter(QGraphicsItem * filterItem);
extern void C_ZN13QGraphicsItem22removeSceneEventFilterEPS_(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsItem::acceptTouchEvents();
extern bool C_ZNK13QGraphicsItem17acceptTouchEventsEv(void* qthis); // 4
  // proto:  QMatrix QGraphicsItem::sceneMatrix();
extern void* C_ZNK13QGraphicsItem11sceneMatrixEv(void* qthis); // 4
  // proto:  QGraphicsItem::CacheMode QGraphicsItem::cacheMode();
extern void C_ZNK13QGraphicsItem9cacheModeEv(void* qthis); // 4
  // proto:  GraphicsItemFlags QGraphicsItem::flags();
extern void C_ZNK13QGraphicsItem5flagsEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsItem::clipPath();
extern void* C_ZNK13QGraphicsItem8clipPathEv(void* qthis); // 4
  // proto:  qreal QGraphicsItem::y();
extern void C_ZNK13QGraphicsItem1yEv(void* qthis); // 2
  // proto:  void QGraphicsItem::scroll(qreal dx, qreal dy, const QRectF & rect);
extern void C_ZN13QGraphicsItem6scrollEddRK6QRectF(void* qthis, double arg0, double arg1, void* arg2); // 4
  // proto:  QRegion QGraphicsItem::boundingRegion(const QTransform & itemToDeviceTransform);
extern void* C_ZNK13QGraphicsItem14boundingRegionERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsObject::~QGraphicsObject();
extern void C_ZN15QGraphicsObjectD2Ev(void* qthis); // 4
  // proto:  void QGraphicsObject::QGraphicsObject(QGraphicsItem * parent);
extern void* C_ZN15QGraphicsObjectC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  const QMetaObject * QGraphicsObject::metaObject();
extern void C_ZNK15QGraphicsObject10metaObjectEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsSimpleTextItem::opaqueArea();
extern void* C_ZNK23QGraphicsSimpleTextItem10opaqueAreaEv(void* qthis); // 4
  // proto:  void QGraphicsSimpleTextItem::~QGraphicsSimpleTextItem();
extern void C_ZN23QGraphicsSimpleTextItemD2Ev(void* qthis); // 4
  // proto:  QRectF QGraphicsSimpleTextItem::boundingRect();
extern void* C_ZNK23QGraphicsSimpleTextItem12boundingRectEv(void* qthis); // 4
  // proto:  QString QGraphicsSimpleTextItem::text();
extern void* C_ZNK23QGraphicsSimpleTextItem4textEv(void* qthis); // 4
  // proto:  void QGraphicsSimpleTextItem::setText(const QString & text);
extern void C_ZN23QGraphicsSimpleTextItem7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsSimpleTextItem::contains(const QPointF & point);
extern bool C_ZNK23QGraphicsSimpleTextItem8containsERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  bool QGraphicsSimpleTextItem::isObscuredBy(const QGraphicsItem * item);
extern bool C_ZNK23QGraphicsSimpleTextItem12isObscuredByEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsSimpleTextItem::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void C_ZN23QGraphicsSimpleTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  QPainterPath QGraphicsSimpleTextItem::shape();
extern void* C_ZNK23QGraphicsSimpleTextItem5shapeEv(void* qthis); // 4
  // proto:  void QGraphicsSimpleTextItem::QGraphicsSimpleTextItem(QGraphicsItem * parent);
extern void* C_ZN23QGraphicsSimpleTextItemC2EP13QGraphicsItem(void* arg0); // 3
  // proto:  void QGraphicsSimpleTextItem::QGraphicsSimpleTextItem(const QString & text, QGraphicsItem * parent);
extern void* C_ZN23QGraphicsSimpleTextItemC2ERK7QStringP13QGraphicsItem(void* arg0, void* arg1); // 3
  // proto:  void QGraphicsSimpleTextItem::setFont(const QFont & font);
extern void C_ZN23QGraphicsSimpleTextItem7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  QFont QGraphicsSimpleTextItem::font();
extern void* C_ZNK23QGraphicsSimpleTextItem4fontEv(void* qthis); // 4
  // proto:  int QGraphicsSimpleTextItem::type();
extern int32_t C_ZNK23QGraphicsSimpleTextItem4typeEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
//  _linkActivated QGraphicsTextItem_linkActivated_signal;
//  _linkHovered QGraphicsTextItem_linkHovered_signal;
}

// class sizeof(QGraphicsPixmapItem)=1
type QGraphicsPixmapItem struct {
  /*qbase*/ QGraphicsItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsRectItem)=1
type QGraphicsRectItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsEllipseItem)=1
type QGraphicsEllipseItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsPolygonItem)=1
type QGraphicsPolygonItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsPathItem)=1
type QGraphicsPathItem struct {
  /*qbase*/ QAbstractGraphicsShapeItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsLineItem)=1
type QGraphicsLineItem struct {
  /*qbase*/ QGraphicsItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsItemGroup)=1
type QGraphicsItemGroup struct {
  /*qbase*/ QGraphicsItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAbstractGraphicsShapeItem)=1
type QAbstractGraphicsShapeItem struct {
  /*qbase*/ QGraphicsItem;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsItem)=1
type QGraphicsItem struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGraphicsObject)=1
type QGraphicsObject struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// contains(const class QPointF &)
func (this *QGraphicsTextItem) Contains(args ...interface{}) (ret interface{}) {
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
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QGraphicsTextItem8containsERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "contains", args)
  }

  return
}

// boundingRect()
func (this *QGraphicsTextItem) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK17QGraphicsTextItem12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "boundingRect", args)
  }

  return
}

// setPlainText(const class QString &)
func (this *QGraphicsTextItem) Setplaintext(args ...interface{}) () {
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
    // invoke: void setPlainText(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem12setPlainTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setPlainText", args)
  }

  return
}

// setHtml(const class QString &)
func (this *QGraphicsTextItem) Sethtml(args ...interface{}) () {
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
    // invoke: void setHtml(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem7setHtmlERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setHtml", args)
  }

  return
}

// setTabChangesFocus(_Bool)
func (this *QGraphicsTextItem) Settabchangesfocus(args ...interface{}) () {
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
    // invoke: void setTabChangesFocus(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem18setTabChangesFocusEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setTabChangesFocus", args)
  }

  return
}

// setDocument(class QTextDocument *)
func (this *QGraphicsTextItem) Setdocument(args ...interface{}) () {
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
    // invoke: void setDocument(class QTextDocument *)
    var arg0 = args[0].(*QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem11setDocumentEP13QTextDocument(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setDocument", args)
  }

  return
}

// setOpenExternalLinks(_Bool)
func (this *QGraphicsTextItem) Setopenexternallinks(args ...interface{}) () {
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
    // invoke: void setOpenExternalLinks(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem20setOpenExternalLinksEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setOpenExternalLinks", args)
  }

  return
}

// font()
func (this *QGraphicsTextItem) Font(args ...interface{}) (ret interface{}) {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem4fontEv
    // invoke: QFont font()
    var ret0 = C.C_ZNK17QGraphicsTextItem4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "font", args)
  }

  return
}

// setDefaultTextColor(const class QColor &)
func (this *QGraphicsTextItem) Setdefaulttextcolor(args ...interface{}) () {
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
    // invoke: void setDefaultTextColor(const class QColor &)
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem19setDefaultTextColorERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setDefaultTextColor", args)
  }

  return
}

// openExternalLinks()
func (this *QGraphicsTextItem) Openexternallinks(args ...interface{}) (ret interface{}) {
  // openExternalLinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem17openExternalLinksEv
    // invoke: bool openExternalLinks()
    var ret0 = C.C_ZNK17QGraphicsTextItem17openExternalLinksEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "openExternalLinks", args)
  }

  return
}

// setTextCursor(const class QTextCursor &)
func (this *QGraphicsTextItem) Settextcursor(args ...interface{}) () {
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
    // invoke: void setTextCursor(const class QTextCursor &)
    var arg0 = args[0].(*QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem13setTextCursorERK11QTextCursor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setTextCursor", args)
  }

  return
}

// toPlainText()
func (this *QGraphicsTextItem) Toplaintext(args ...interface{}) (ret interface{}) {
  // toPlainText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem11toPlainTextEv
    // invoke: QString toPlainText()
    var ret0 = C.C_ZNK17QGraphicsTextItem11toPlainTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "toPlainText", args)
  }

  return
}

// tabChangesFocus()
func (this *QGraphicsTextItem) Tabchangesfocus(args ...interface{}) (ret interface{}) {
  // tabChangesFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem15tabChangesFocusEv
    // invoke: bool tabChangesFocus()
    var ret0 = C.C_ZNK17QGraphicsTextItem15tabChangesFocusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "tabChangesFocus", args)
  }

  return
}

// textCursor()
func (this *QGraphicsTextItem) Textcursor(args ...interface{}) (ret interface{}) {
  // textCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem10textCursorEv
    // invoke: QTextCursor textCursor()
    var ret0 = C.C_ZNK17QGraphicsTextItem10textCursorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCursor{}) // "QTextCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "textCursor", args)
  }

  return
}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsTextItem) Paint(args ...interface{}) () {
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
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN17QGraphicsTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "paint", args)
  }

  return
}

// QGraphicsTextItem(class QGraphicsItem *)
func NewQGraphicsTextItem(args ...interface{}) *QGraphicsTextItem {
  // QGraphicsTextItem(class QGraphicsItem *)
  // QGraphicsTextItem(const class QString &, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItemC1EP13QGraphicsItem
    // invoke: void QGraphicsTextItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsTextItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsTextItem{Qclsinst:qthis}
  case 1:
    // invoke: _ZN17QGraphicsTextItemC1ERK7QStringP13QGraphicsItem
    // invoke: void QGraphicsTextItem(const class QString &, class QGraphicsItem *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsTextItemC2ERK7QStringP13QGraphicsItem(arg0, arg1)
    return &QGraphicsTextItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "QGraphicsTextItem", args)
  }

  return nil // QGraphicsTextItem{Qclsinst:qthis}
}

// type()
func (this *QGraphicsTextItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK17QGraphicsTextItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "type", args)
  }

  return
}

// textInteractionFlags()
func (this *QGraphicsTextItem) Textinteractionflags(args ...interface{}) () {
  // textInteractionFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem20textInteractionFlagsEv
    // invoke: Qt::TextInteractionFlags textInteractionFlags()
    C.C_ZNK17QGraphicsTextItem20textInteractionFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "textInteractionFlags", args)
  }

  return
}

// toHtml()
func (this *QGraphicsTextItem) Tohtml(args ...interface{}) (ret interface{}) {
  // toHtml()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem6toHtmlEv
    // invoke: QString toHtml()
    var ret0 = C.C_ZNK17QGraphicsTextItem6toHtmlEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "toHtml", args)
  }

  return
}

// adjustSize()
func (this *QGraphicsTextItem) Adjustsize(args ...interface{}) () {
  // adjustSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItem10adjustSizeEv
    // invoke: void adjustSize()
    C.C_ZN17QGraphicsTextItem10adjustSizeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "adjustSize", args)
  }

  return
}

// ~QGraphicsTextItem()
func (this *QGraphicsTextItem) Freeqgraphicstextitem(args ...interface{}) () {
  // ~QGraphicsTextItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsTextItemD0Ev
    // invoke: void ~QGraphicsTextItem()
    C.C_ZN17QGraphicsTextItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "~QGraphicsTextItem", args)
  }

  return
}

// setTextWidth(qreal)
func (this *QGraphicsTextItem) Settextwidth(args ...interface{}) () {
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
    // invoke: void setTextWidth(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem12setTextWidthEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setTextWidth", args)
  }

  return
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsTextItem) Isobscuredby(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QGraphicsTextItem12isObscuredByEPK13QGraphicsItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "isObscuredBy", args)
  }

  return
}

// document()
func (this *QGraphicsTextItem) Document(args ...interface{}) (ret interface{}) {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem8documentEv
    // invoke: QTextDocument * document()
    var ret0 = C.C_ZNK17QGraphicsTextItem8documentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "document", args)
  }

  return
}

// opaqueArea()
func (this *QGraphicsTextItem) Opaquearea(args ...interface{}) (ret interface{}) {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret0 = C.C_ZNK17QGraphicsTextItem10opaqueAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "opaqueArea", args)
  }

  return
}

// defaultTextColor()
func (this *QGraphicsTextItem) Defaulttextcolor(args ...interface{}) (ret interface{}) {
  // defaultTextColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem16defaultTextColorEv
    // invoke: QColor defaultTextColor()
    var ret0 = C.C_ZNK17QGraphicsTextItem16defaultTextColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "defaultTextColor", args)
  }

  return
}

// metaObject()
func (this *QGraphicsTextItem) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK17QGraphicsTextItem10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "metaObject", args)
  }

  return
}

// shape()
func (this *QGraphicsTextItem) Shape(args ...interface{}) (ret interface{}) {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem5shapeEv
    // invoke: QPainterPath shape()
    var ret0 = C.C_ZNK17QGraphicsTextItem5shapeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "shape", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QGraphicsTextItem) Setfont(args ...interface{}) () {
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
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsTextItem7setFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "setFont", args)
  }

  return
}

// textWidth()
func (this *QGraphicsTextItem) Textwidth(args ...interface{}) (ret interface{}) {
  // textWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsTextItem9textWidthEv
    // invoke: qreal textWidth()
    var ret0 = C.C_ZNK17QGraphicsTextItem9textWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsTextItem", "textWidth", args)
  }

  return
}

// boundingRect()
func (this *QGraphicsPixmapItem) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK19QGraphicsPixmapItem12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "boundingRect", args)
  }

  return
}

// setPixmap(const class QPixmap &)
func (this *QGraphicsPixmapItem) Setpixmap(args ...interface{}) () {
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
    // invoke: void setPixmap(const class QPixmap &)
    var arg0 = args[0].(*QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsPixmapItem9setPixmapERK7QPixmap(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "setPixmap", args)
  }

  return
}

// shape()
func (this *QGraphicsPixmapItem) Shape(args ...interface{}) (ret interface{}) {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem5shapeEv
    // invoke: QPainterPath shape()
    var ret0 = C.C_ZNK19QGraphicsPixmapItem5shapeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "shape", args)
  }

  return
}

// transformationMode()
func (this *QGraphicsPixmapItem) Transformationmode(args ...interface{}) () {
  // transformationMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem18transformationModeEv
    // invoke: Qt::TransformationMode transformationMode()
    C.C_ZNK19QGraphicsPixmapItem18transformationModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "transformationMode", args)
  }

  return
}

// pixmap()
func (this *QGraphicsPixmapItem) Pixmap(args ...interface{}) (ret interface{}) {
  // pixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem6pixmapEv
    // invoke: QPixmap pixmap()
    var ret0 = C.C_ZNK19QGraphicsPixmapItem6pixmapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixmap{}) // "QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "pixmap", args)
  }

  return
}

// contains(const class QPointF &)
func (this *QGraphicsPixmapItem) Contains(args ...interface{}) (ret interface{}) {
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
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsPixmapItem8containsERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "contains", args)
  }

  return
}

// shapeMode()
func (this *QGraphicsPixmapItem) Shapemode(args ...interface{}) () {
  // shapeMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem9shapeModeEv
    // invoke: QGraphicsPixmapItem::ShapeMode shapeMode()
    C.C_ZNK19QGraphicsPixmapItem9shapeModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "shapeMode", args)
  }

  return
}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsPixmapItem) Paint(args ...interface{}) () {
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
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN19QGraphicsPixmapItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "paint", args)
  }

  return
}

// setOffset(const class QPointF &)
func (this *QGraphicsPixmapItem) Setoffset(args ...interface{}) () {
  // setOffset(const class QPointF &)
  // setOffset(qreal, qreal)
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
    // invoke: _ZN19QGraphicsPixmapItem9setOffsetERK7QPointF
    // invoke: void setOffset(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsPixmapItem9setOffsetERK7QPointF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN19QGraphicsPixmapItem9setOffsetEdd
    // invoke: void setOffset(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN19QGraphicsPixmapItem9setOffsetEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "setOffset", args)
  }

  return
}

// type()
func (this *QGraphicsPixmapItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK19QGraphicsPixmapItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "type", args)
  }

  return
}

// ~QGraphicsPixmapItem()
func (this *QGraphicsPixmapItem) Freeqgraphicspixmapitem(args ...interface{}) () {
  // ~QGraphicsPixmapItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsPixmapItemD0Ev
    // invoke: void ~QGraphicsPixmapItem()
    C.C_ZN19QGraphicsPixmapItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "~QGraphicsPixmapItem", args)
  }

  return
}

// offset()
func (this *QGraphicsPixmapItem) Offset(args ...interface{}) (ret interface{}) {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem6offsetEv
    // invoke: QPointF offset()
    var ret0 = C.C_ZNK19QGraphicsPixmapItem6offsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "offset", args)
  }

  return
}

// QGraphicsPixmapItem(class QGraphicsItem *)
func NewQGraphicsPixmapItem(args ...interface{}) *QGraphicsPixmapItem {
  // QGraphicsPixmapItem(class QGraphicsItem *)
  // QGraphicsPixmapItem(const class QPixmap &, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[1][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsPixmapItemC1EP13QGraphicsItem
    // invoke: void QGraphicsPixmapItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QGraphicsPixmapItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsPixmapItem{Qclsinst:qthis}
  case 1:
    // invoke: _ZN19QGraphicsPixmapItemC1ERK7QPixmapP13QGraphicsItem
    // invoke: void QGraphicsPixmapItem(const class QPixmap &, class QGraphicsItem *)
    var arg0 = args[0].(*QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QGraphicsPixmapItemC2ERK7QPixmapP13QGraphicsItem(arg0, arg1)
    return &QGraphicsPixmapItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "QGraphicsPixmapItem", args)
  }

  return nil // QGraphicsPixmapItem{Qclsinst:qthis}
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsPixmapItem) Isobscuredby(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsPixmapItem12isObscuredByEPK13QGraphicsItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "isObscuredBy", args)
  }

  return
}

// opaqueArea()
func (this *QGraphicsPixmapItem) Opaquearea(args ...interface{}) (ret interface{}) {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsPixmapItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret0 = C.C_ZNK19QGraphicsPixmapItem10opaqueAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPixmapItem", "opaqueArea", args)
  }

  return
}

// opaqueArea()
func (this *QGraphicsRectItem) Opaquearea(args ...interface{}) (ret interface{}) {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret0 = C.C_ZNK17QGraphicsRectItem10opaqueAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "opaqueArea", args)
  }

  return
}

// boundingRect()
func (this *QGraphicsRectItem) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK17QGraphicsRectItem12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "boundingRect", args)
  }

  return
}

// type()
func (this *QGraphicsRectItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK17QGraphicsRectItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "type", args)
  }

  return
}

// contains(const class QPointF &)
func (this *QGraphicsRectItem) Contains(args ...interface{}) (ret interface{}) {
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
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QGraphicsRectItem8containsERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "contains", args)
  }

  return
}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsRectItem) Paint(args ...interface{}) () {
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
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "paint", args)
  }

  return
}

// shape()
func (this *QGraphicsRectItem) Shape(args ...interface{}) (ret interface{}) {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem5shapeEv
    // invoke: QPainterPath shape()
    var ret0 = C.C_ZNK17QGraphicsRectItem5shapeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "shape", args)
  }

  return
}

// QGraphicsRectItem(class QGraphicsItem *)
func NewQGraphicsRectItem(args ...interface{}) *QGraphicsRectItem {
  // QGraphicsRectItem(class QGraphicsItem *)
  // QGraphicsRectItem(const class QRectF &, class QGraphicsItem *)
  // QGraphicsRectItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRectItemC1EP13QGraphicsItem
    // invoke: void QGraphicsRectItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsRectItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsRectItem{Qclsinst:qthis}
  case 1:
    // invoke: _ZN17QGraphicsRectItemC1ERK6QRectFP13QGraphicsItem
    // invoke: void QGraphicsRectItem(const class QRectF &, class QGraphicsItem *)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsRectItemC2ERK6QRectFP13QGraphicsItem(arg0, arg1)
    return &QGraphicsRectItem{Qclsinst:qthis}
  case 2:
    // invoke: _ZN17QGraphicsRectItemC1EddddP13QGraphicsItem
    // invoke: void QGraphicsRectItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg4)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsRectItemC2EddddP13QGraphicsItem(arg0, arg1, arg2, arg3, arg4)
    return &QGraphicsRectItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "QGraphicsRectItem", args)
  }

  return nil // QGraphicsRectItem{Qclsinst:qthis}
}

// ~QGraphicsRectItem()
func (this *QGraphicsRectItem) Freeqgraphicsrectitem(args ...interface{}) () {
  // ~QGraphicsRectItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRectItemD0Ev
    // invoke: void ~QGraphicsRectItem()
    C.C_ZN17QGraphicsRectItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "~QGraphicsRectItem", args)
  }

  return
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsRectItem) Isobscuredby(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QGraphicsRectItem12isObscuredByEPK13QGraphicsItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "isObscuredBy", args)
  }

  return
}

// setRect(const class QRectF &)
func (this *QGraphicsRectItem) Setrect(args ...interface{}) () {
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
    // invoke: void setRect(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsRectItem7setRectERK6QRectF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN17QGraphicsRectItem7setRectEdddd
    // invoke: void setRect(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN17QGraphicsRectItem7setRectEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "setRect", args)
  }

  return
}

// rect()
func (this *QGraphicsRectItem) Rect(args ...interface{}) (ret interface{}) {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRectItem4rectEv
    // invoke: QRectF rect()
    var ret0 = C.C_ZNK17QGraphicsRectItem4rectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsRectItem", "rect", args)
  }

  return
}

// opaqueArea()
func (this *QGraphicsEllipseItem) Opaquearea(args ...interface{}) (ret interface{}) {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret0 = C.C_ZNK20QGraphicsEllipseItem10opaqueAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "opaqueArea", args)
  }

  return
}

// boundingRect()
func (this *QGraphicsEllipseItem) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK20QGraphicsEllipseItem12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "boundingRect", args)
  }

  return
}

// spanAngle()
func (this *QGraphicsEllipseItem) Spanangle(args ...interface{}) (ret interface{}) {
  // spanAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem9spanAngleEv
    // invoke: int spanAngle()
    var ret0 = C.C_ZNK20QGraphicsEllipseItem9spanAngleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "spanAngle", args)
  }

  return
}

// type()
func (this *QGraphicsEllipseItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK20QGraphicsEllipseItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "type", args)
  }

  return
}

// contains(const class QPointF &)
func (this *QGraphicsEllipseItem) Contains(args ...interface{}) (ret interface{}) {
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
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK20QGraphicsEllipseItem8containsERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "contains", args)
  }

  return
}

// setSpanAngle(int)
func (this *QGraphicsEllipseItem) Setspanangle(args ...interface{}) () {
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
    // invoke: void setSpanAngle(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN20QGraphicsEllipseItem12setSpanAngleEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "setSpanAngle", args)
  }

  return
}

// setRect(const class QRectF &)
func (this *QGraphicsEllipseItem) Setrect(args ...interface{}) () {
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
    // invoke: void setRect(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN20QGraphicsEllipseItem7setRectERK6QRectF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN20QGraphicsEllipseItem7setRectEdddd
    // invoke: void setRect(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN20QGraphicsEllipseItem7setRectEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "setRect", args)
  }

  return
}

// setStartAngle(int)
func (this *QGraphicsEllipseItem) Setstartangle(args ...interface{}) () {
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
    // invoke: void setStartAngle(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN20QGraphicsEllipseItem13setStartAngleEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "setStartAngle", args)
  }

  return
}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsEllipseItem) Paint(args ...interface{}) () {
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
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN20QGraphicsEllipseItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "paint", args)
  }

  return
}

// shape()
func (this *QGraphicsEllipseItem) Shape(args ...interface{}) (ret interface{}) {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem5shapeEv
    // invoke: QPainterPath shape()
    var ret0 = C.C_ZNK20QGraphicsEllipseItem5shapeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "shape", args)
  }

  return
}

// ~QGraphicsEllipseItem()
func (this *QGraphicsEllipseItem) Freeqgraphicsellipseitem(args ...interface{}) () {
  // ~QGraphicsEllipseItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItemD0Ev
    // invoke: void ~QGraphicsEllipseItem()
    C.C_ZN20QGraphicsEllipseItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "~QGraphicsEllipseItem", args)
  }

  return
}

// startAngle()
func (this *QGraphicsEllipseItem) Startangle(args ...interface{}) (ret interface{}) {
  // startAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem10startAngleEv
    // invoke: int startAngle()
    var ret0 = C.C_ZNK20QGraphicsEllipseItem10startAngleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "startAngle", args)
  }

  return
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsEllipseItem) Isobscuredby(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK20QGraphicsEllipseItem12isObscuredByEPK13QGraphicsItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "isObscuredBy", args)
  }

  return
}

// QGraphicsEllipseItem(class QGraphicsItem *)
func NewQGraphicsEllipseItem(args ...interface{}) *QGraphicsEllipseItem {
  // QGraphicsEllipseItem(class QGraphicsItem *)
  // QGraphicsEllipseItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
  // QGraphicsEllipseItem(const class QRectF &, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsEllipseItemC1EP13QGraphicsItem
    // invoke: void QGraphicsEllipseItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QGraphicsEllipseItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsEllipseItem{Qclsinst:qthis}
  case 1:
    // invoke: _ZN20QGraphicsEllipseItemC1EddddP13QGraphicsItem
    // invoke: void QGraphicsEllipseItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg4)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QGraphicsEllipseItemC2EddddP13QGraphicsItem(arg0, arg1, arg2, arg3, arg4)
    return &QGraphicsEllipseItem{Qclsinst:qthis}
  case 2:
    // invoke: _ZN20QGraphicsEllipseItemC1ERK6QRectFP13QGraphicsItem
    // invoke: void QGraphicsEllipseItem(const class QRectF &, class QGraphicsItem *)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QGraphicsEllipseItemC2ERK6QRectFP13QGraphicsItem(arg0, arg1)
    return &QGraphicsEllipseItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "QGraphicsEllipseItem", args)
  }

  return nil // QGraphicsEllipseItem{Qclsinst:qthis}
}

// rect()
func (this *QGraphicsEllipseItem) Rect(args ...interface{}) (ret interface{}) {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsEllipseItem4rectEv
    // invoke: QRectF rect()
    var ret0 = C.C_ZNK20QGraphicsEllipseItem4rectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsEllipseItem", "rect", args)
  }

  return
}

// opaqueArea()
func (this *QGraphicsPolygonItem) Opaquearea(args ...interface{}) (ret interface{}) {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret0 = C.C_ZNK20QGraphicsPolygonItem10opaqueAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "opaqueArea", args)
  }

  return
}

// boundingRect()
func (this *QGraphicsPolygonItem) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK20QGraphicsPolygonItem12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "boundingRect", args)
  }

  return
}

// setPolygon(const class QPolygonF &)
func (this *QGraphicsPolygonItem) Setpolygon(args ...interface{}) () {
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
    // invoke: void setPolygon(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN20QGraphicsPolygonItem10setPolygonERK9QPolygonF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "setPolygon", args)
  }

  return
}

// type()
func (this *QGraphicsPolygonItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK20QGraphicsPolygonItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "type", args)
  }

  return
}

// contains(const class QPointF &)
func (this *QGraphicsPolygonItem) Contains(args ...interface{}) (ret interface{}) {
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
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK20QGraphicsPolygonItem8containsERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "contains", args)
  }

  return
}

// polygon()
func (this *QGraphicsPolygonItem) Polygon(args ...interface{}) (ret interface{}) {
  // polygon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem7polygonEv
    // invoke: QPolygonF polygon()
    var ret0 = C.C_ZNK20QGraphicsPolygonItem7polygonEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "polygon", args)
  }

  return
}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsPolygonItem) Paint(args ...interface{}) () {
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
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN20QGraphicsPolygonItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "paint", args)
  }

  return
}

// shape()
func (this *QGraphicsPolygonItem) Shape(args ...interface{}) (ret interface{}) {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem5shapeEv
    // invoke: QPainterPath shape()
    var ret0 = C.C_ZNK20QGraphicsPolygonItem5shapeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "shape", args)
  }

  return
}

// QGraphicsPolygonItem(const class QPolygonF &, class QGraphicsItem *)
func NewQGraphicsPolygonItem(args ...interface{}) *QGraphicsPolygonItem {
  // QGraphicsPolygonItem(const class QPolygonF &, class QGraphicsItem *)
  // QGraphicsPolygonItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsPolygonItemC1ERK9QPolygonFP13QGraphicsItem
    // invoke: void QGraphicsPolygonItem(const class QPolygonF &, class QGraphicsItem *)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QGraphicsPolygonItemC2ERK9QPolygonFP13QGraphicsItem(arg0, arg1)
    return &QGraphicsPolygonItem{Qclsinst:qthis}
  case 1:
    // invoke: _ZN20QGraphicsPolygonItemC1EP13QGraphicsItem
    // invoke: void QGraphicsPolygonItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QGraphicsPolygonItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsPolygonItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "QGraphicsPolygonItem", args)
  }

  return nil // QGraphicsPolygonItem{Qclsinst:qthis}
}

// ~QGraphicsPolygonItem()
func (this *QGraphicsPolygonItem) Freeqgraphicspolygonitem(args ...interface{}) () {
  // ~QGraphicsPolygonItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsPolygonItemD0Ev
    // invoke: void ~QGraphicsPolygonItem()
    C.C_ZN20QGraphicsPolygonItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "~QGraphicsPolygonItem", args)
  }

  return
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsPolygonItem) Isobscuredby(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK20QGraphicsPolygonItem12isObscuredByEPK13QGraphicsItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "isObscuredBy", args)
  }

  return
}

// fillRule()
func (this *QGraphicsPolygonItem) Fillrule(args ...interface{}) () {
  // fillRule()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsPolygonItem8fillRuleEv
    // invoke: Qt::FillRule fillRule()
    C.C_ZNK20QGraphicsPolygonItem8fillRuleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPolygonItem", "fillRule", args)
  }

  return
}

// opaqueArea()
func (this *QGraphicsPathItem) Opaquearea(args ...interface{}) (ret interface{}) {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret0 = C.C_ZNK17QGraphicsPathItem10opaqueAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "opaqueArea", args)
  }

  return
}

// ~QGraphicsPathItem()
func (this *QGraphicsPathItem) Freeqgraphicspathitem(args ...interface{}) () {
  // ~QGraphicsPathItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsPathItemD0Ev
    // invoke: void ~QGraphicsPathItem()
    C.C_ZN17QGraphicsPathItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "~QGraphicsPathItem", args)
  }

  return
}

// boundingRect()
func (this *QGraphicsPathItem) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK17QGraphicsPathItem12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "boundingRect", args)
  }

  return
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsPathItem) Isobscuredby(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QGraphicsPathItem12isObscuredByEPK13QGraphicsItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "isObscuredBy", args)
  }

  return
}

// contains(const class QPointF &)
func (this *QGraphicsPathItem) Contains(args ...interface{}) (ret interface{}) {
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
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QGraphicsPathItem8containsERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "contains", args)
  }

  return
}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsPathItem) Paint(args ...interface{}) () {
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
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN17QGraphicsPathItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "paint", args)
  }

  return
}

// shape()
func (this *QGraphicsPathItem) Shape(args ...interface{}) (ret interface{}) {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem5shapeEv
    // invoke: QPainterPath shape()
    var ret0 = C.C_ZNK17QGraphicsPathItem5shapeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "shape", args)
  }

  return
}

// path()
func (this *QGraphicsPathItem) Path(args ...interface{}) (ret interface{}) {
  // path()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem4pathEv
    // invoke: QPainterPath path()
    var ret0 = C.C_ZNK17QGraphicsPathItem4pathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "path", args)
  }

  return
}

// setPath(const class QPainterPath &)
func (this *QGraphicsPathItem) Setpath(args ...interface{}) () {
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
    // invoke: void setPath(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsPathItem7setPathERK12QPainterPath(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "setPath", args)
  }

  return
}

// QGraphicsPathItem(const class QPainterPath &, class QGraphicsItem *)
func NewQGraphicsPathItem(args ...interface{}) *QGraphicsPathItem {
  // QGraphicsPathItem(const class QPainterPath &, class QGraphicsItem *)
  // QGraphicsPathItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsPathItemC1ERK12QPainterPathP13QGraphicsItem
    // invoke: void QGraphicsPathItem(const class QPainterPath &, class QGraphicsItem *)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsPathItemC2ERK12QPainterPathP13QGraphicsItem(arg0, arg1)
    return &QGraphicsPathItem{Qclsinst:qthis}
  case 1:
    // invoke: _ZN17QGraphicsPathItemC1EP13QGraphicsItem
    // invoke: void QGraphicsPathItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsPathItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsPathItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "QGraphicsPathItem", args)
  }

  return nil // QGraphicsPathItem{Qclsinst:qthis}
}

// type()
func (this *QGraphicsPathItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsPathItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK17QGraphicsPathItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsPathItem", "type", args)
  }

  return
}

// opaqueArea()
func (this *QGraphicsLineItem) Opaquearea(args ...interface{}) (ret interface{}) {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret0 = C.C_ZNK17QGraphicsLineItem10opaqueAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "opaqueArea", args)
  }

  return
}

// boundingRect()
func (this *QGraphicsLineItem) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK17QGraphicsLineItem12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "boundingRect", args)
  }

  return
}

// type()
func (this *QGraphicsLineItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK17QGraphicsLineItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "type", args)
  }

  return
}

// contains(const class QPointF &)
func (this *QGraphicsLineItem) Contains(args ...interface{}) (ret interface{}) {
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
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QGraphicsLineItem8containsERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "contains", args)
  }

  return
}

// shape()
func (this *QGraphicsLineItem) Shape(args ...interface{}) (ret interface{}) {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem5shapeEv
    // invoke: QPainterPath shape()
    var ret0 = C.C_ZNK17QGraphicsLineItem5shapeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "shape", args)
  }

  return
}

// ~QGraphicsLineItem()
func (this *QGraphicsLineItem) Freeqgraphicslineitem(args ...interface{}) () {
  // ~QGraphicsLineItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsLineItemD0Ev
    // invoke: void ~QGraphicsLineItem()
    C.C_ZN17QGraphicsLineItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "~QGraphicsLineItem", args)
  }

  return
}

// QGraphicsLineItem(class QGraphicsItem *)
func NewQGraphicsLineItem(args ...interface{}) *QGraphicsLineItem {
  // QGraphicsLineItem(class QGraphicsItem *)
  // QGraphicsLineItem(const class QLineF &, class QGraphicsItem *)
  // QGraphicsLineItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[1][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsLineItemC1EP13QGraphicsItem
    // invoke: void QGraphicsLineItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsLineItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsLineItem{Qclsinst:qthis}
  case 1:
    // invoke: _ZN17QGraphicsLineItemC1ERK6QLineFP13QGraphicsItem
    // invoke: void QGraphicsLineItem(const class QLineF &, class QGraphicsItem *)
    var arg0 = args[0].(*QLineF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsLineItemC2ERK6QLineFP13QGraphicsItem(arg0, arg1)
    return &QGraphicsLineItem{Qclsinst:qthis}
  case 2:
    // invoke: _ZN17QGraphicsLineItemC1EddddP13QGraphicsItem
    // invoke: void QGraphicsLineItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg4)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QGraphicsLineItemC2EddddP13QGraphicsItem(arg0, arg1, arg2, arg3, arg4)
    return &QGraphicsLineItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "QGraphicsLineItem", args)
  }

  return nil // QGraphicsLineItem{Qclsinst:qthis}
}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsLineItem) Paint(args ...interface{}) () {
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
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN17QGraphicsLineItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "paint", args)
  }

  return
}

// pen()
func (this *QGraphicsLineItem) Pen(args ...interface{}) (ret interface{}) {
  // pen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem3penEv
    // invoke: QPen pen()
    var ret0 = C.C_ZNK17QGraphicsLineItem3penEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPen{}) // "QPen"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "pen", args)
  }

  return
}

// setPen(const class QPen &)
func (this *QGraphicsLineItem) Setpen(args ...interface{}) () {
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
    // invoke: void setPen(const class QPen &)
    var arg0 = args[0].(*QPen).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsLineItem6setPenERK4QPen(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "setPen", args)
  }

  return
}

// setLine(const class QLineF &)
func (this *QGraphicsLineItem) Setline(args ...interface{}) () {
  // setLine(const class QLineF &)
  // setLine(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsLineItem7setLineERK6QLineF
    // invoke: void setLine(const class QLineF &)
    var arg0 = args[0].(*QLineF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QGraphicsLineItem7setLineERK6QLineF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN17QGraphicsLineItem7setLineEdddd
    // invoke: void setLine(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN17QGraphicsLineItem7setLineEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "setLine", args)
  }

  return
}

// line()
func (this *QGraphicsLineItem) Line(args ...interface{}) (ret interface{}) {
  // line()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsLineItem4lineEv
    // invoke: QLineF line()
    var ret0 = C.C_ZNK17QGraphicsLineItem4lineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLineF{}) // "QLineF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "line", args)
  }

  return
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsLineItem) Isobscuredby(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QGraphicsLineItem12isObscuredByEPK13QGraphicsItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsLineItem", "isObscuredBy", args)
  }

  return
}

// opaqueArea()
func (this *QGraphicsItemGroup) Opaquearea(args ...interface{}) (ret interface{}) {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsItemGroup10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret0 = C.C_ZNK18QGraphicsItemGroup10opaqueAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "opaqueArea", args)
  }

  return
}

// addToGroup(class QGraphicsItem *)
func (this *QGraphicsItemGroup) Addtogroup(args ...interface{}) () {
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
    // invoke: void addToGroup(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QGraphicsItemGroup10addToGroupEP13QGraphicsItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "addToGroup", args)
  }

  return
}

// QGraphicsItemGroup(class QGraphicsItem *)
func NewQGraphicsItemGroup(args ...interface{}) *QGraphicsItemGroup {
  // QGraphicsItemGroup(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsItemGroupC1EP13QGraphicsItem
    // invoke: void QGraphicsItemGroup(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QGraphicsItemGroupC2EP13QGraphicsItem(arg0)
    return &QGraphicsItemGroup{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "QGraphicsItemGroup", args)
  }

  return nil // QGraphicsItemGroup{Qclsinst:qthis}
}

// boundingRect()
func (this *QGraphicsItemGroup) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsItemGroup12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK18QGraphicsItemGroup12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "boundingRect", args)
  }

  return
}

// removeFromGroup(class QGraphicsItem *)
func (this *QGraphicsItemGroup) Removefromgroup(args ...interface{}) () {
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
    // invoke: void removeFromGroup(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QGraphicsItemGroup15removeFromGroupEP13QGraphicsItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "removeFromGroup", args)
  }

  return
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsItemGroup) Isobscuredby(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QGraphicsItemGroup12isObscuredByEPK13QGraphicsItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "isObscuredBy", args)
  }

  return
}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsItemGroup) Paint(args ...interface{}) () {
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
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "paint", args)
  }

  return
}

// type()
func (this *QGraphicsItemGroup) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsItemGroup4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK18QGraphicsItemGroup4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "type", args)
  }

  return
}

// ~QGraphicsItemGroup()
func (this *QGraphicsItemGroup) Freeqgraphicsitemgroup(args ...interface{}) () {
  // ~QGraphicsItemGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsItemGroupD0Ev
    // invoke: void ~QGraphicsItemGroup()
    C.C_ZN18QGraphicsItemGroupD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemGroup", "~QGraphicsItemGroup", args)
  }

  return
}

// opaqueArea()
func (this *QAbstractGraphicsShapeItem) Opaquearea(args ...interface{}) (ret interface{}) {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAbstractGraphicsShapeItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret0 = C.C_ZNK26QAbstractGraphicsShapeItem10opaqueAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "opaqueArea", args)
  }

  return
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QAbstractGraphicsShapeItem) Isobscuredby(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK26QAbstractGraphicsShapeItem12isObscuredByEPK13QGraphicsItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "isObscuredBy", args)
  }

  return
}

// pen()
func (this *QAbstractGraphicsShapeItem) Pen(args ...interface{}) (ret interface{}) {
  // pen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAbstractGraphicsShapeItem3penEv
    // invoke: QPen pen()
    var ret0 = C.C_ZNK26QAbstractGraphicsShapeItem3penEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPen{}) // "QPen"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "pen", args)
  }

  return
}

// brush()
func (this *QAbstractGraphicsShapeItem) Brush(args ...interface{}) (ret interface{}) {
  // brush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAbstractGraphicsShapeItem5brushEv
    // invoke: QBrush brush()
    var ret0 = C.C_ZNK26QAbstractGraphicsShapeItem5brushEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "brush", args)
  }

  return
}

// setPen(const class QPen &)
func (this *QAbstractGraphicsShapeItem) Setpen(args ...interface{}) () {
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
    // invoke: void setPen(const class QPen &)
    var arg0 = args[0].(*QPen).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN26QAbstractGraphicsShapeItem6setPenERK4QPen(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "setPen", args)
  }

  return
}

// setBrush(const class QBrush &)
func (this *QAbstractGraphicsShapeItem) Setbrush(args ...interface{}) () {
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
    // invoke: void setBrush(const class QBrush &)
    var arg0 = args[0].(*QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN26QAbstractGraphicsShapeItem8setBrushERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "setBrush", args)
  }

  return
}

// ~QAbstractGraphicsShapeItem()
func (this *QAbstractGraphicsShapeItem) Freeqabstractgraphicsshapeitem(args ...interface{}) () {
  // ~QAbstractGraphicsShapeItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractGraphicsShapeItemD0Ev
    // invoke: void ~QAbstractGraphicsShapeItem()
    C.C_ZN26QAbstractGraphicsShapeItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "~QAbstractGraphicsShapeItem", args)
  }

  return
}

// QAbstractGraphicsShapeItem(class QGraphicsItem *)
func NewQAbstractGraphicsShapeItem(args ...interface{}) *QAbstractGraphicsShapeItem {
  // QAbstractGraphicsShapeItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractGraphicsShapeItemC1EP13QGraphicsItem
    // invoke: void QAbstractGraphicsShapeItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QAbstractGraphicsShapeItemC2EP13QGraphicsItem(arg0)
    return &QAbstractGraphicsShapeItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractGraphicsShapeItem", "QAbstractGraphicsShapeItem", args)
  }

  return nil // QAbstractGraphicsShapeItem{Qclsinst:qthis}
}

// show()
func (this *QGraphicsItem) Show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem4showEv
    // invoke: void show()
    C.C_ZN13QGraphicsItem4showEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "show", args)
  }

  return
}

// hasFocus()
func (this *QGraphicsItem) Hasfocus(args ...interface{}) (ret interface{}) {
  // hasFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8hasFocusEv
    // invoke: bool hasFocus()
    var ret0 = C.C_ZNK13QGraphicsItem8hasFocusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "hasFocus", args)
  }

  return
}

// setGroup(class QGraphicsItemGroup *)
func (this *QGraphicsItem) Setgroup(args ...interface{}) () {
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
    // invoke: void setGroup(class QGraphicsItemGroup *)
    var arg0 = args[0].(*QGraphicsItemGroup).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem8setGroupEP18QGraphicsItemGroup(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setGroup", args)
  }

  return
}

// handlesChildEvents()
func (this *QGraphicsItem) Handleschildevents(args ...interface{}) (ret interface{}) {
  // handlesChildEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem18handlesChildEventsEv
    // invoke: bool handlesChildEvents()
    var ret0 = C.C_ZNK13QGraphicsItem18handlesChildEventsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "handlesChildEvents", args)
  }

  return
}

// shape()
func (this *QGraphicsItem) Shape(args ...interface{}) (ret interface{}) {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5shapeEv
    // invoke: QPainterPath shape()
    var ret0 = C.C_ZNK13QGraphicsItem5shapeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "shape", args)
  }

  return
}

// isPanel()
func (this *QGraphicsItem) Ispanel(args ...interface{}) (ret interface{}) {
  // isPanel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem7isPanelEv
    // invoke: bool isPanel()
    var ret0 = C.C_ZNK13QGraphicsItem7isPanelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isPanel", args)
  }

  return
}

// data(int)
func (this *QGraphicsItem) Data(args ...interface{}) (ret interface{}) {
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
    // invoke: QVariant data(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem4dataEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "data", args)
  }

  return
}

// parentItem()
func (this *QGraphicsItem) Parentitem(args ...interface{}) () {
  // parentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10parentItemEv
    // invoke: QGraphicsItem * parentItem()
    C.C_ZNK13QGraphicsItem10parentItemEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "parentItem", args)
  }

  return
}

// childrenBoundingRect()
func (this *QGraphicsItem) Childrenboundingrect(args ...interface{}) (ret interface{}) {
  // childrenBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem20childrenBoundingRectEv
    // invoke: QRectF childrenBoundingRect()
    var ret0 = C.C_ZNK13QGraphicsItem20childrenBoundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "childrenBoundingRect", args)
  }

  return
}

// group()
func (this *QGraphicsItem) Group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5groupEv
    // invoke: QGraphicsItemGroup * group()
    C.C_ZNK13QGraphicsItem5groupEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "group", args)
  }

  return
}

// isEnabled()
func (this *QGraphicsItem) Isenabled(args ...interface{}) (ret interface{}) {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9isEnabledEv
    // invoke: bool isEnabled()
    var ret0 = C.C_ZNK13QGraphicsItem9isEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isEnabled", args)
  }

  return
}

// window()
func (this *QGraphicsItem) Window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6windowEv
    // invoke: QGraphicsWidget * window()
    C.C_ZNK13QGraphicsItem6windowEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "window", args)
  }

  return
}

// setBoundingRegionGranularity(qreal)
func (this *QGraphicsItem) Setboundingregiongranularity(args ...interface{}) () {
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
    // invoke: void setBoundingRegionGranularity(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem28setBoundingRegionGranularityEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setBoundingRegionGranularity", args)
  }

  return
}

// setRotation(qreal)
func (this *QGraphicsItem) Setrotation(args ...interface{}) () {
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
    // invoke: void setRotation(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem11setRotationEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setRotation", args)
  }

  return
}

// commonAncestorItem(const class QGraphicsItem *)
func (this *QGraphicsItem) Commonancestoritem(args ...interface{}) () {
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
    // invoke: QGraphicsItem * commonAncestorItem(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QGraphicsItem18commonAncestorItemEPKS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "commonAncestorItem", args)
  }

  return
}

// mapToItem(const class QGraphicsItem *, const class QPolygonF &)
func (this *QGraphicsItem) Maptoitem(args ...interface{}) (ret interface{}) {
  // mapToItem(const class QGraphicsItem *, const class QPolygonF &)
  // mapToItem(const class QGraphicsItem *, const class QRectF &)
  // mapToItem(const class QGraphicsItem *, const class QPainterPath &)
  // mapToItem(const class QGraphicsItem *, const class QPointF &)
  // mapToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
  // mapToItem(const class QGraphicsItem *, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[1][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[2][1] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[3][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[4][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[5][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[5][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK9QPolygonF
    // invoke: QPolygonF mapToItem(const class QGraphicsItem *, const class QPolygonF &)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem9mapToItemEPKS_RK9QPolygonF(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK6QRectF
    // invoke: QPolygonF mapToItem(const class QGraphicsItem *, const class QRectF &)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QRectF).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem9mapToItemEPKS_RK6QRectF(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK12QPainterPath
    // invoke: QPainterPath mapToItem(const class QGraphicsItem *, const class QPainterPath &)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem9mapToItemEPKS_RK12QPainterPath(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_RK7QPointF
    // invoke: QPointF mapToItem(const class QGraphicsItem *, const class QPointF &)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem9mapToItemEPKS_RK7QPointF(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_dddd
    // invoke: QPolygonF mapToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(qtrt.PrimConv(args[4], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZNK13QGraphicsItem9mapToItemEPKS_dddd(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZNK13QGraphicsItem9mapToItemEPKS_dd
    // invoke: QPointF mapToItem(const class QGraphicsItem *, qreal, qreal)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK13QGraphicsItem9mapToItemEPKS_dd(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapToItem", args)
  }

  return
}

// toGraphicsObject()
func (this *QGraphicsItem) Tographicsobject(args ...interface{}) () {
  // toGraphicsObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem16toGraphicsObjectEv
    // invoke: QGraphicsObject * toGraphicsObject()
    C.C_ZN13QGraphicsItem16toGraphicsObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "toGraphicsObject", args)
  }

  return
}

// acceptedMouseButtons()
func (this *QGraphicsItem) Acceptedmousebuttons(args ...interface{}) () {
  // acceptedMouseButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem20acceptedMouseButtonsEv
    // invoke: Qt::MouseButtons acceptedMouseButtons()
    C.C_ZNK13QGraphicsItem20acceptedMouseButtonsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "acceptedMouseButtons", args)
  }

  return
}

// setGraphicsEffect(class QGraphicsEffect *)
func (this *QGraphicsItem) Setgraphicseffect(args ...interface{}) () {
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
    // invoke: void setGraphicsEffect(class QGraphicsEffect *)
    var arg0 = args[0].(*QGraphicsEffect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem17setGraphicsEffectEP15QGraphicsEffect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setGraphicsEffect", args)
  }

  return
}

// mapToParent(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) Maptoparent(args ...interface{}) (ret interface{}) {
  // mapToParent(qreal, qreal, qreal, qreal)
  // mapToParent(const class QPointF &)
  // mapToParent(qreal, qreal)
  // mapToParent(const class QPainterPath &)
  // mapToParent(const class QRectF &)
  // mapToParent(const class QPolygonF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11mapToParentEdddd
    // invoke: QPolygonF mapToParent(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapToParentEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK7QPointF
    // invoke: QPointF mapToParent(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapToParentERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK13QGraphicsItem11mapToParentEdd
    // invoke: QPointF mapToParent(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapToParentEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK12QPainterPath
    // invoke: QPainterPath mapToParent(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapToParentERK12QPainterPath(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK6QRectF
    // invoke: QPolygonF mapToParent(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapToParentERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZNK13QGraphicsItem11mapToParentERK9QPolygonF
    // invoke: QPolygonF mapToParent(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapToParentERK9QPolygonF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapToParent", args)
  }

  return
}

// setTransformOriginPoint(const class QPointF &)
func (this *QGraphicsItem) Settransformoriginpoint(args ...interface{}) () {
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
    // invoke: void setTransformOriginPoint(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem23setTransformOriginPointERK7QPointF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN13QGraphicsItem23setTransformOriginPointEdd
    // invoke: void setTransformOriginPoint(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem23setTransformOriginPointEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setTransformOriginPoint", args)
  }

  return
}

// deviceTransform(const class QTransform &)
func (this *QGraphicsItem) Devicetransform(args ...interface{}) (ret interface{}) {
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
    // invoke: QTransform deviceTransform(const class QTransform &)
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem15deviceTransformERK10QTransform(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "deviceTransform", args)
  }

  return
}

// rotation()
func (this *QGraphicsItem) Rotation(args ...interface{}) (ret interface{}) {
  // rotation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8rotationEv
    // invoke: qreal rotation()
    var ret0 = C.C_ZNK13QGraphicsItem8rotationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "rotation", args)
  }

  return
}

// effectiveOpacity()
func (this *QGraphicsItem) Effectiveopacity(args ...interface{}) (ret interface{}) {
  // effectiveOpacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem16effectiveOpacityEv
    // invoke: qreal effectiveOpacity()
    var ret0 = C.C_ZNK13QGraphicsItem16effectiveOpacityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "effectiveOpacity", args)
  }

  return
}

// focusScopeItem()
func (this *QGraphicsItem) Focusscopeitem(args ...interface{}) () {
  // focusScopeItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14focusScopeItemEv
    // invoke: QGraphicsItem * focusScopeItem()
    C.C_ZNK13QGraphicsItem14focusScopeItemEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "focusScopeItem", args)
  }

  return
}

// inputMethodHints()
func (this *QGraphicsItem) Inputmethodhints(args ...interface{}) () {
  // inputMethodHints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem16inputMethodHintsEv
    // invoke: Qt::InputMethodHints inputMethodHints()
    C.C_ZNK13QGraphicsItem16inputMethodHintsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "inputMethodHints", args)
  }

  return
}

// panelModality()
func (this *QGraphicsItem) Panelmodality(args ...interface{}) () {
  // panelModality()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem13panelModalityEv
    // invoke: QGraphicsItem::PanelModality panelModality()
    C.C_ZNK13QGraphicsItem13panelModalityEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "panelModality", args)
  }

  return
}

// topLevelWidget()
func (this *QGraphicsItem) Toplevelwidget(args ...interface{}) () {
  // topLevelWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14topLevelWidgetEv
    // invoke: QGraphicsWidget * topLevelWidget()
    C.C_ZNK13QGraphicsItem14topLevelWidgetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "topLevelWidget", args)
  }

  return
}

// isVisibleTo(const class QGraphicsItem *)
func (this *QGraphicsItem) Isvisibleto(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isVisibleTo(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem11isVisibleToEPKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isVisibleTo", args)
  }

  return
}

// setAcceptDrops(_Bool)
func (this *QGraphicsItem) Setacceptdrops(args ...interface{}) () {
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
    // invoke: void setAcceptDrops(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem14setAcceptDropsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setAcceptDrops", args)
  }

  return
}

// clearFocus()
func (this *QGraphicsItem) Clearfocus(args ...interface{}) () {
  // clearFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem10clearFocusEv
    // invoke: void clearFocus()
    C.C_ZN13QGraphicsItem10clearFocusEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "clearFocus", args)
  }

  return
}

// filtersChildEvents()
func (this *QGraphicsItem) Filterschildevents(args ...interface{}) (ret interface{}) {
  // filtersChildEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem18filtersChildEventsEv
    // invoke: bool filtersChildEvents()
    var ret0 = C.C_ZNK13QGraphicsItem18filtersChildEventsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "filtersChildEvents", args)
  }

  return
}

// x()
func (this *QGraphicsItem) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem1xEv
    // invoke: qreal x()
    C.C_ZNK13QGraphicsItem1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "x", args)
  }

  return
}

// setTransform(const class QTransform &, _Bool)
func (this *QGraphicsItem) Settransform(args ...interface{}) () {
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
    // invoke: void setTransform(const class QTransform &, _Bool)
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem12setTransformERK10QTransformb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setTransform", args)
  }

  return
}

// stackBefore(const class QGraphicsItem *)
func (this *QGraphicsItem) Stackbefore(args ...interface{}) () {
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
    // invoke: void stackBefore(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem11stackBeforeEPKS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "stackBefore", args)
  }

  return
}

// setAcceptHoverEvents(_Bool)
func (this *QGraphicsItem) Setaccepthoverevents(args ...interface{}) () {
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
    // invoke: void setAcceptHoverEvents(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem20setAcceptHoverEventsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setAcceptHoverEvents", args)
  }

  return
}

// focusProxy()
func (this *QGraphicsItem) Focusproxy(args ...interface{}) () {
  // focusProxy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10focusProxyEv
    // invoke: QGraphicsItem * focusProxy()
    C.C_ZNK13QGraphicsItem10focusProxyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "focusProxy", args)
  }

  return
}

// mapFromParent(const class QPainterPath &)
func (this *QGraphicsItem) Mapfromparent(args ...interface{}) (ret interface{}) {
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
    // invoke: QPainterPath mapFromParent(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem13mapFromParentERK12QPainterPath(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem13mapFromParentERK6QRectF
    // invoke: QPolygonF mapFromParent(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem13mapFromParentERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK13QGraphicsItem13mapFromParentEdddd
    // invoke: QPolygonF mapFromParent(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QGraphicsItem13mapFromParentEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK13QGraphicsItem13mapFromParentEdd
    // invoke: QPointF mapFromParent(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem13mapFromParentEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZNK13QGraphicsItem13mapFromParentERK7QPointF
    // invoke: QPointF mapFromParent(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem13mapFromParentERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZNK13QGraphicsItem13mapFromParentERK9QPolygonF
    // invoke: QPolygonF mapFromParent(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem13mapFromParentERK9QPolygonF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapFromParent", args)
  }

  return
}

// scenePos()
func (this *QGraphicsItem) Scenepos(args ...interface{}) (ret interface{}) {
  // scenePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8scenePosEv
    // invoke: QPointF scenePos()
    var ret0 = C.C_ZNK13QGraphicsItem8scenePosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scenePos", args)
  }

  return
}

// setPos(qreal, qreal)
func (this *QGraphicsItem) Setpos(args ...interface{}) () {
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
    // invoke: void setPos(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem6setPosEdd(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN13QGraphicsItem6setPosERK7QPointF
    // invoke: void setPos(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem6setPosERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setPos", args)
  }

  return
}

// scale()
func (this *QGraphicsItem) Scale(args ...interface{}) (ret interface{}) {
  // scale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5scaleEv
    // invoke: qreal scale()
    var ret0 = C.C_ZNK13QGraphicsItem5scaleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scale", args)
  }

  return
}

// hide()
func (this *QGraphicsItem) Hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem4hideEv
    // invoke: void hide()
    C.C_ZN13QGraphicsItem4hideEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "hide", args)
  }

  return
}

// matrix()
func (this *QGraphicsItem) Matrix(args ...interface{}) (ret interface{}) {
  // matrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6matrixEv
    // invoke: QMatrix matrix()
    var ret0 = C.C_ZNK13QGraphicsItem6matrixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "matrix", args)
  }

  return
}

// isSelected()
func (this *QGraphicsItem) Isselected(args ...interface{}) (ret interface{}) {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10isSelectedEv
    // invoke: bool isSelected()
    var ret0 = C.C_ZNK13QGraphicsItem10isSelectedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isSelected", args)
  }

  return
}

// moveBy(qreal, qreal)
func (this *QGraphicsItem) Moveby(args ...interface{}) () {
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
    // invoke: void moveBy(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem6moveByEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "moveBy", args)
  }

  return
}

// setFocusProxy(class QGraphicsItem *)
func (this *QGraphicsItem) Setfocusproxy(args ...interface{}) () {
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
    // invoke: void setFocusProxy(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem13setFocusProxyEPS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setFocusProxy", args)
  }

  return
}

// itemTransform(const class QGraphicsItem *, _Bool *)
func (this *QGraphicsItem) Itemtransform(args ...interface{}) (ret interface{}) {
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
    // invoke: QTransform itemTransform(const class QGraphicsItem *, _Bool *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem13itemTransformEPKS_Pb(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "itemTransform", args)
  }

  return
}

// mapRectToParent(const class QRectF &)
func (this *QGraphicsItem) Maprecttoparent(args ...interface{}) (ret interface{}) {
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
    // invoke: QRectF mapRectToParent(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem15mapRectToParentERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem15mapRectToParentEdddd
    // invoke: QRectF mapRectToParent(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QGraphicsItem15mapRectToParentEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectToParent", args)
  }

  return
}

// isBlockedByModalPanel(class QGraphicsItem **)
func (this *QGraphicsItem) Isblockedbymodalpanel(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isBlockedByModalPanel(class QGraphicsItem **)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isBlockedByModalPanel", args)
  }

  return
}

// graphicsEffect()
func (this *QGraphicsItem) Graphicseffect(args ...interface{}) () {
  // graphicsEffect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14graphicsEffectEv
    // invoke: QGraphicsEffect * graphicsEffect()
    C.C_ZNK13QGraphicsItem14graphicsEffectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "graphicsEffect", args)
  }

  return
}

// childItems()
func (this *QGraphicsItem) Childitems(args ...interface{}) () {
  // childItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10childItemsEv
    // invoke: QList<QGraphicsItem *> childItems()
    C.C_ZNK13QGraphicsItem10childItemsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "childItems", args)
  }

  return
}

// isClipped()
func (this *QGraphicsItem) Isclipped(args ...interface{}) (ret interface{}) {
  // isClipped()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9isClippedEv
    // invoke: bool isClipped()
    var ret0 = C.C_ZNK13QGraphicsItem9isClippedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isClipped", args)
  }

  return
}

// toolTip()
func (this *QGraphicsItem) Tooltip(args ...interface{}) (ret interface{}) {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem7toolTipEv
    // invoke: QString toolTip()
    var ret0 = C.C_ZNK13QGraphicsItem7toolTipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "toolTip", args)
  }

  return
}

// ~QGraphicsItem()
func (this *QGraphicsItem) Freeqgraphicsitem(args ...interface{}) () {
  // ~QGraphicsItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItemD0Ev
    // invoke: void ~QGraphicsItem()
    C.C_ZN13QGraphicsItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "~QGraphicsItem", args)
  }

  return
}

// isActive()
func (this *QGraphicsItem) Isactive(args ...interface{}) (ret interface{}) {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8isActiveEv
    // invoke: bool isActive()
    var ret0 = C.C_ZNK13QGraphicsItem8isActiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isActive", args)
  }

  return
}

// setFiltersChildEvents(_Bool)
func (this *QGraphicsItem) Setfilterschildevents(args ...interface{}) () {
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
    // invoke: void setFiltersChildEvents(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem21setFiltersChildEventsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setFiltersChildEvents", args)
  }

  return
}

// opaqueArea()
func (this *QGraphicsItem) Opaquearea(args ...interface{}) (ret interface{}) {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret0 = C.C_ZNK13QGraphicsItem10opaqueAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "opaqueArea", args)
  }

  return
}

// advance(int)
func (this *QGraphicsItem) Advance(args ...interface{}) () {
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
    // invoke: void advance(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem7advanceEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "advance", args)
  }

  return
}

// setToolTip(const class QString &)
func (this *QGraphicsItem) Settooltip(args ...interface{}) () {
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
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem10setToolTipERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setToolTip", args)
  }

  return
}

// mapToScene(const class QPolygonF &)
func (this *QGraphicsItem) Maptoscene(args ...interface{}) (ret interface{}) {
  // mapToScene(const class QPolygonF &)
  // mapToScene(const class QRectF &)
  // mapToScene(const class QPointF &)
  // mapToScene(const class QPainterPath &)
  // mapToScene(qreal, qreal, qreal, qreal)
  // mapToScene(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[4][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[5][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK9QPolygonF
    // invoke: QPolygonF mapToScene(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem10mapToSceneERK9QPolygonF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK6QRectF
    // invoke: QPolygonF mapToScene(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem10mapToSceneERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK7QPointF
    // invoke: QPointF mapToScene(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem10mapToSceneERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK13QGraphicsItem10mapToSceneERK12QPainterPath
    // invoke: QPainterPath mapToScene(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem10mapToSceneERK12QPainterPath(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZNK13QGraphicsItem10mapToSceneEdddd
    // invoke: QPolygonF mapToScene(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QGraphicsItem10mapToSceneEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZNK13QGraphicsItem10mapToSceneEdd
    // invoke: QPointF mapToScene(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem10mapToSceneEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapToScene", args)
  }

  return
}

// setHandlesChildEvents(_Bool)
func (this *QGraphicsItem) Sethandleschildevents(args ...interface{}) () {
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
    // invoke: void setHandlesChildEvents(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem21setHandlesChildEventsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setHandlesChildEvents", args)
  }

  return
}

// transformOriginPoint()
func (this *QGraphicsItem) Transformoriginpoint(args ...interface{}) (ret interface{}) {
  // transformOriginPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem20transformOriginPointEv
    // invoke: QPointF transformOriginPoint()
    var ret0 = C.C_ZNK13QGraphicsItem20transformOriginPointEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "transformOriginPoint", args)
  }

  return
}

// pos()
func (this *QGraphicsItem) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem3posEv
    // invoke: QPointF pos()
    var ret0 = C.C_ZNK13QGraphicsItem3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "pos", args)
  }

  return
}

// mapRectToScene(const class QRectF &)
func (this *QGraphicsItem) Maprecttoscene(args ...interface{}) (ret interface{}) {
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
    // invoke: QRectF mapRectToScene(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem14mapRectToSceneERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem14mapRectToSceneEdddd
    // invoke: QRectF mapRectToScene(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QGraphicsItem14mapRectToSceneEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectToScene", args)
  }

  return
}

// ungrabMouse()
func (this *QGraphicsItem) Ungrabmouse(args ...interface{}) () {
  // ungrabMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11ungrabMouseEv
    // invoke: void ungrabMouse()
    C.C_ZN13QGraphicsItem11ungrabMouseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "ungrabMouse", args)
  }

  return
}

// isObscured(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) Isobscured(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscured(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QGraphicsItem10isObscuredEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem10isObscuredERK6QRectF
    // invoke: bool isObscured(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem10isObscuredERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isObscured", args)
  }

  return
}

// setSelected(_Bool)
func (this *QGraphicsItem) Setselected(args ...interface{}) () {
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
    // invoke: void setSelected(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem11setSelectedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setSelected", args)
  }

  return
}

// mapFromScene(const class QPolygonF &)
func (this *QGraphicsItem) Mapfromscene(args ...interface{}) (ret interface{}) {
  // mapFromScene(const class QPolygonF &)
  // mapFromScene(const class QRectF &)
  // mapFromScene(qreal, qreal, qreal, qreal)
  // mapFromScene(qreal, qreal)
  // mapFromScene(const class QPainterPath &)
  // mapFromScene(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
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
  vtys[4][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK9QPolygonF
    // invoke: QPolygonF mapFromScene(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem12mapFromSceneERK9QPolygonF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK6QRectF
    // invoke: QPolygonF mapFromScene(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem12mapFromSceneERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneEdddd
    // invoke: QPolygonF mapFromScene(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QGraphicsItem12mapFromSceneEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneEdd
    // invoke: QPointF mapFromScene(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem12mapFromSceneEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK12QPainterPath
    // invoke: QPainterPath mapFromScene(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem12mapFromSceneERK12QPainterPath(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZNK13QGraphicsItem12mapFromSceneERK7QPointF
    // invoke: QPointF mapFromScene(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem12mapFromSceneERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapFromScene", args)
  }

  return
}

// scene()
func (this *QGraphicsItem) Scene(args ...interface{}) () {
  // scene()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5sceneEv
    // invoke: QGraphicsScene * scene()
    C.C_ZNK13QGraphicsItem5sceneEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scene", args)
  }

  return
}

// isAncestorOf(const class QGraphicsItem *)
func (this *QGraphicsItem) Isancestorof(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isAncestorOf(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem12isAncestorOfEPKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isAncestorOf", args)
  }

  return
}

// topLevelItem()
func (this *QGraphicsItem) Toplevelitem(args ...interface{}) () {
  // topLevelItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12topLevelItemEv
    // invoke: QGraphicsItem * topLevelItem()
    C.C_ZNK13QGraphicsItem12topLevelItemEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "topLevelItem", args)
  }

  return
}

// panel()
func (this *QGraphicsItem) Panel(args ...interface{}) () {
  // panel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5panelEv
    // invoke: QGraphicsItem * panel()
    C.C_ZNK13QGraphicsItem5panelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "panel", args)
  }

  return
}

// unsetCursor()
func (this *QGraphicsItem) Unsetcursor(args ...interface{}) () {
  // unsetCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11unsetCursorEv
    // invoke: void unsetCursor()
    C.C_ZN13QGraphicsItem11unsetCursorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "unsetCursor", args)
  }

  return
}

// transformations()
func (this *QGraphicsItem) Transformations(args ...interface{}) () {
  // transformations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem15transformationsEv
    // invoke: QList<QGraphicsTransform *> transformations()
    C.C_ZNK13QGraphicsItem15transformationsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "transformations", args)
  }

  return
}

// mapRectFromScene(const class QRectF &)
func (this *QGraphicsItem) Maprectfromscene(args ...interface{}) (ret interface{}) {
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
    // invoke: QRectF mapRectFromScene(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem16mapRectFromSceneERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem16mapRectFromSceneEdddd
    // invoke: QRectF mapRectFromScene(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QGraphicsItem16mapRectFromSceneEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectFromScene", args)
  }

  return
}

// contains(const class QPointF &)
func (this *QGraphicsItem) Contains(args ...interface{}) (ret interface{}) {
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
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem8containsERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "contains", args)
  }

  return
}

// transform()
func (this *QGraphicsItem) Transform(args ...interface{}) (ret interface{}) {
  // transform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9transformEv
    // invoke: QTransform transform()
    var ret0 = C.C_ZNK13QGraphicsItem9transformEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "transform", args)
  }

  return
}

// ungrabKeyboard()
func (this *QGraphicsItem) Ungrabkeyboard(args ...interface{}) () {
  // ungrabKeyboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem14ungrabKeyboardEv
    // invoke: void ungrabKeyboard()
    C.C_ZN13QGraphicsItem14ungrabKeyboardEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "ungrabKeyboard", args)
  }

  return
}

// installSceneEventFilter(class QGraphicsItem *)
func (this *QGraphicsItem) Installsceneeventfilter(args ...interface{}) () {
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
    // invoke: void installSceneEventFilter(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem23installSceneEventFilterEPS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "installSceneEventFilter", args)
  }

  return
}

// acceptHoverEvents()
func (this *QGraphicsItem) Accepthoverevents(args ...interface{}) (ret interface{}) {
  // acceptHoverEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem17acceptHoverEventsEv
    // invoke: bool acceptHoverEvents()
    var ret0 = C.C_ZNK13QGraphicsItem17acceptHoverEventsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "acceptHoverEvents", args)
  }

  return
}

// mapRectToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) Maprecttoitem(args ...interface{}) (ret interface{}) {
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
    // invoke: QRectF mapRectToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(qtrt.PrimConv(args[4], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZNK13QGraphicsItem13mapRectToItemEPKS_dddd(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem13mapRectToItemEPKS_RK6QRectF
    // invoke: QRectF mapRectToItem(const class QGraphicsItem *, const class QRectF &)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QRectF).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem13mapRectToItemEPKS_RK6QRectF(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectToItem", args)
  }

  return
}

// mapFromItem(const class QGraphicsItem *, qreal, qreal)
func (this *QGraphicsItem) Mapfromitem(args ...interface{}) (ret interface{}) {
  // mapFromItem(const class QGraphicsItem *, qreal, qreal)
  // mapFromItem(const class QGraphicsItem *, const class QPointF &)
  // mapFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
  // mapFromItem(const class QGraphicsItem *, const class QRectF &)
  // mapFromItem(const class QGraphicsItem *, const class QPolygonF &)
  // mapFromItem(const class QGraphicsItem *, const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[1][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[3][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[4][1] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[5][1] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_dd
    // invoke: QPointF mapFromItem(const class QGraphicsItem *, qreal, qreal)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_dd(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK7QPointF
    // invoke: QPointF mapFromItem(const class QGraphicsItem *, const class QPointF &)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_RK7QPointF(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_dddd
    // invoke: QPolygonF mapFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(qtrt.PrimConv(args[4], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_dddd(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK6QRectF
    // invoke: QPolygonF mapFromItem(const class QGraphicsItem *, const class QRectF &)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QRectF).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_RK6QRectF(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK9QPolygonF
    // invoke: QPolygonF mapFromItem(const class QGraphicsItem *, const class QPolygonF &)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_RK9QPolygonF(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZNK13QGraphicsItem11mapFromItemEPKS_RK12QPainterPath
    // invoke: QPainterPath mapFromItem(const class QGraphicsItem *, const class QPainterPath &)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem11mapFromItemEPKS_RK12QPainterPath(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapFromItem", args)
  }

  return
}

// QGraphicsItem(class QGraphicsItem *)
func NewQGraphicsItem(args ...interface{}) *QGraphicsItem {
  // QGraphicsItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItemC1EPS_
    // invoke: void QGraphicsItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QGraphicsItemC2EPS_(arg0)
    return &QGraphicsItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsItem", "QGraphicsItem", args)
  }

  return nil // QGraphicsItem{Qclsinst:qthis}
}

// type()
func (this *QGraphicsItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK13QGraphicsItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "type", args)
  }

  return
}

// setEnabled(_Bool)
func (this *QGraphicsItem) Setenabled(args ...interface{}) () {
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
    // invoke: void setEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem10setEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setEnabled", args)
  }

  return
}

// grabMouse()
func (this *QGraphicsItem) Grabmouse(args ...interface{}) () {
  // grabMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem9grabMouseEv
    // invoke: void grabMouse()
    C.C_ZN13QGraphicsItem9grabMouseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "grabMouse", args)
  }

  return
}

// setActive(_Bool)
func (this *QGraphicsItem) Setactive(args ...interface{}) () {
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
    // invoke: void setActive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem9setActiveEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setActive", args)
  }

  return
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsItem) Isobscuredby(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem12isObscuredByEPKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isObscuredBy", args)
  }

  return
}

// acceptDrops()
func (this *QGraphicsItem) Acceptdrops(args ...interface{}) (ret interface{}) {
  // acceptDrops()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11acceptDropsEv
    // invoke: bool acceptDrops()
    var ret0 = C.C_ZNK13QGraphicsItem11acceptDropsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "acceptDrops", args)
  }

  return
}

// setParentItem(class QGraphicsItem *)
func (this *QGraphicsItem) Setparentitem(args ...interface{}) () {
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
    // invoke: void setParentItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem13setParentItemEPS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setParentItem", args)
  }

  return
}

// ensureVisible(qreal, qreal, qreal, qreal, int, int)
func (this *QGraphicsItem) Ensurevisible(args ...interface{}) () {
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
    // invoke: void ensureVisible(qreal, qreal, qreal, qreal, int, int)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    C.C_ZN13QGraphicsItem13ensureVisibleEddddii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN13QGraphicsItem13ensureVisibleERK6QRectFii
    // invoke: void ensureVisible(const class QRectF &, int, int)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN13QGraphicsItem13ensureVisibleERK6QRectFii(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "ensureVisible", args)
  }

  return
}

// grabKeyboard()
func (this *QGraphicsItem) Grabkeyboard(args ...interface{}) () {
  // grabKeyboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem12grabKeyboardEv
    // invoke: void grabKeyboard()
    C.C_ZN13QGraphicsItem12grabKeyboardEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "grabKeyboard", args)
  }

  return
}

// parentObject()
func (this *QGraphicsItem) Parentobject(args ...interface{}) () {
  // parentObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12parentObjectEv
    // invoke: QGraphicsObject * parentObject()
    C.C_ZNK13QGraphicsItem12parentObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "parentObject", args)
  }

  return
}

// isUnderMouse()
func (this *QGraphicsItem) Isundermouse(args ...interface{}) (ret interface{}) {
  // isUnderMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12isUnderMouseEv
    // invoke: bool isUnderMouse()
    var ret0 = C.C_ZNK13QGraphicsItem12isUnderMouseEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isUnderMouse", args)
  }

  return
}

// parentWidget()
func (this *QGraphicsItem) Parentwidget(args ...interface{}) () {
  // parentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem12parentWidgetEv
    // invoke: QGraphicsWidget * parentWidget()
    C.C_ZNK13QGraphicsItem12parentWidgetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "parentWidget", args)
  }

  return
}

// setZValue(qreal)
func (this *QGraphicsItem) Setzvalue(args ...interface{}) () {
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
    // invoke: void setZValue(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem9setZValueEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setZValue", args)
  }

  return
}

// setMatrix(const class QMatrix &, _Bool)
func (this *QGraphicsItem) Setmatrix(args ...interface{}) () {
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
    // invoke: void setMatrix(const class QMatrix &, _Bool)
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem9setMatrixERK7QMatrixb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setMatrix", args)
  }

  return
}

// cursor()
func (this *QGraphicsItem) Cursor(args ...interface{}) (ret interface{}) {
  // cursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6cursorEv
    // invoke: QCursor cursor()
    var ret0 = C.C_ZNK13QGraphicsItem6cursorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCursor{}) // "QCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "cursor", args)
  }

  return
}

// zValue()
func (this *QGraphicsItem) Zvalue(args ...interface{}) (ret interface{}) {
  // zValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem6zValueEv
    // invoke: qreal zValue()
    var ret0 = C.C_ZNK13QGraphicsItem6zValueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "zValue", args)
  }

  return
}

// setVisible(_Bool)
func (this *QGraphicsItem) Setvisible(args ...interface{}) () {
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
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem10setVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setVisible", args)
  }

  return
}

// resetMatrix()
func (this *QGraphicsItem) Resetmatrix(args ...interface{}) () {
  // resetMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem11resetMatrixEv
    // invoke: void resetMatrix()
    C.C_ZN13QGraphicsItem11resetMatrixEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "resetMatrix", args)
  }

  return
}

// sceneTransform()
func (this *QGraphicsItem) Scenetransform(args ...interface{}) (ret interface{}) {
  // sceneTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem14sceneTransformEv
    // invoke: QTransform sceneTransform()
    var ret0 = C.C_ZNK13QGraphicsItem14sceneTransformEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "sceneTransform", args)
  }

  return
}

// isWindow()
func (this *QGraphicsItem) Iswindow(args ...interface{}) (ret interface{}) {
  // isWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8isWindowEv
    // invoke: bool isWindow()
    var ret0 = C.C_ZNK13QGraphicsItem8isWindowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isWindow", args)
  }

  return
}

// mapRectFromParent(const class QRectF &)
func (this *QGraphicsItem) Maprectfromparent(args ...interface{}) (ret interface{}) {
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
    // invoke: QRectF mapRectFromParent(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem17mapRectFromParentERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem17mapRectFromParentEdddd
    // invoke: QRectF mapRectFromParent(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QGraphicsItem17mapRectFromParentEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectFromParent", args)
  }

  return
}

// setScale(qreal)
func (this *QGraphicsItem) Setscale(args ...interface{}) () {
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
    // invoke: void setScale(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem8setScaleEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setScale", args)
  }

  return
}

// isWidget()
func (this *QGraphicsItem) Iswidget(args ...interface{}) (ret interface{}) {
  // isWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8isWidgetEv
    // invoke: bool isWidget()
    var ret0 = C.C_ZNK13QGraphicsItem8isWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isWidget", args)
  }

  return
}

// resetTransform()
func (this *QGraphicsItem) Resettransform(args ...interface{}) () {
  // resetTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsItem14resetTransformEv
    // invoke: void resetTransform()
    C.C_ZN13QGraphicsItem14resetTransformEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "resetTransform", args)
  }

  return
}

// sceneBoundingRect()
func (this *QGraphicsItem) Sceneboundingrect(args ...interface{}) (ret interface{}) {
  // sceneBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem17sceneBoundingRectEv
    // invoke: QRectF sceneBoundingRect()
    var ret0 = C.C_ZNK13QGraphicsItem17sceneBoundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "sceneBoundingRect", args)
  }

  return
}

// setOpacity(qreal)
func (this *QGraphicsItem) Setopacity(args ...interface{}) () {
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
    // invoke: void setOpacity(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem10setOpacityEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setOpacity", args)
  }

  return
}

// focusItem()
func (this *QGraphicsItem) Focusitem(args ...interface{}) () {
  // focusItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9focusItemEv
    // invoke: QGraphicsItem * focusItem()
    C.C_ZNK13QGraphicsItem9focusItemEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "focusItem", args)
  }

  return
}

// hasCursor()
func (this *QGraphicsItem) Hascursor(args ...interface{}) (ret interface{}) {
  // hasCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9hasCursorEv
    // invoke: bool hasCursor()
    var ret0 = C.C_ZNK13QGraphicsItem9hasCursorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "hasCursor", args)
  }

  return
}

// setAcceptTouchEvents(_Bool)
func (this *QGraphicsItem) Setaccepttouchevents(args ...interface{}) () {
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
    // invoke: void setAcceptTouchEvents(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem20setAcceptTouchEventsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setAcceptTouchEvents", args)
  }

  return
}

// setData(int, const class QVariant &)
func (this *QGraphicsItem) Setdata(args ...interface{}) () {
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
    // invoke: void setData(int, const class QVariant &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVariant).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsItem7setDataEiRK8QVariant(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setData", args)
  }

  return
}

// opacity()
func (this *QGraphicsItem) Opacity(args ...interface{}) (ret interface{}) {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem7opacityEv
    // invoke: qreal opacity()
    var ret0 = C.C_ZNK13QGraphicsItem7opacityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "opacity", args)
  }

  return
}

// isVisible()
func (this *QGraphicsItem) Isvisible(args ...interface{}) (ret interface{}) {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9isVisibleEv
    // invoke: bool isVisible()
    var ret0 = C.C_ZNK13QGraphicsItem9isVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "isVisible", args)
  }

  return
}

// update(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) Update(args ...interface{}) () {
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
    // invoke: void update(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN13QGraphicsItem6updateEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN13QGraphicsItem6updateERK6QRectF
    // invoke: void update(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem6updateERK6QRectF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "update", args)
  }

  return
}

// mapRectFromItem(const class QGraphicsItem *, const class QRectF &)
func (this *QGraphicsItem) Maprectfromitem(args ...interface{}) (ret interface{}) {
  // mapRectFromItem(const class QGraphicsItem *, const class QRectF &)
  // mapRectFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem15mapRectFromItemEPKS_RK6QRectF
    // invoke: QRectF mapRectFromItem(const class QGraphicsItem *, const class QRectF &)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QRectF).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsItem15mapRectFromItemEPKS_RK6QRectF(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsItem15mapRectFromItemEPKS_dddd
    // invoke: QRectF mapRectFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(qtrt.PrimConv(args[4], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZNK13QGraphicsItem15mapRectFromItemEPKS_dddd(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "mapRectFromItem", args)
  }

  return
}

// setX(qreal)
func (this *QGraphicsItem) Setx(args ...interface{}) () {
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
    // invoke: void setX(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem4setXEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setX", args)
  }

  return
}

// setY(qreal)
func (this *QGraphicsItem) Sety(args ...interface{}) () {
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
    // invoke: void setY(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem4setYEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setY", args)
  }

  return
}

// setCursor(const class QCursor &)
func (this *QGraphicsItem) Setcursor(args ...interface{}) () {
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
    // invoke: void setCursor(const class QCursor &)
    var arg0 = args[0].(*QCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem9setCursorERK7QCursor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "setCursor", args)
  }

  return
}

// boundingRegionGranularity()
func (this *QGraphicsItem) Boundingregiongranularity(args ...interface{}) (ret interface{}) {
  // boundingRegionGranularity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem25boundingRegionGranularityEv
    // invoke: qreal boundingRegionGranularity()
    var ret0 = C.C_ZNK13QGraphicsItem25boundingRegionGranularityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "boundingRegionGranularity", args)
  }

  return
}

// removeSceneEventFilter(class QGraphicsItem *)
func (this *QGraphicsItem) Removesceneeventfilter(args ...interface{}) () {
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
    // invoke: void removeSceneEventFilter(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsItem22removeSceneEventFilterEPS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "removeSceneEventFilter", args)
  }

  return
}

// acceptTouchEvents()
func (this *QGraphicsItem) Accepttouchevents(args ...interface{}) (ret interface{}) {
  // acceptTouchEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem17acceptTouchEventsEv
    // invoke: bool acceptTouchEvents()
    var ret0 = C.C_ZNK13QGraphicsItem17acceptTouchEventsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "acceptTouchEvents", args)
  }

  return
}

// sceneMatrix()
func (this *QGraphicsItem) Scenematrix(args ...interface{}) (ret interface{}) {
  // sceneMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem11sceneMatrixEv
    // invoke: QMatrix sceneMatrix()
    var ret0 = C.C_ZNK13QGraphicsItem11sceneMatrixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "sceneMatrix", args)
  }

  return
}

// cacheMode()
func (this *QGraphicsItem) Cachemode(args ...interface{}) () {
  // cacheMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem9cacheModeEv
    // invoke: QGraphicsItem::CacheMode cacheMode()
    C.C_ZNK13QGraphicsItem9cacheModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "cacheMode", args)
  }

  return
}

// flags()
func (this *QGraphicsItem) Flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem5flagsEv
    // invoke: GraphicsItemFlags flags()
    C.C_ZNK13QGraphicsItem5flagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "flags", args)
  }

  return
}

// clipPath()
func (this *QGraphicsItem) Clippath(args ...interface{}) (ret interface{}) {
  // clipPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem8clipPathEv
    // invoke: QPainterPath clipPath()
    var ret0 = C.C_ZNK13QGraphicsItem8clipPathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "clipPath", args)
  }

  return
}

// y()
func (this *QGraphicsItem) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsItem1yEv
    // invoke: qreal y()
    C.C_ZNK13QGraphicsItem1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "y", args)
  }

  return
}

// scroll(qreal, qreal, const class QRectF &)
func (this *QGraphicsItem) Scroll(args ...interface{}) () {
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
    // invoke: void scroll(qreal, qreal, const class QRectF &)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QRectF).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN13QGraphicsItem6scrollEddRK6QRectF(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsItem", "scroll", args)
  }

  return
}

// boundingRegion(const class QTransform &)
func (this *QGraphicsItem) Boundingregion(args ...interface{}) (ret interface{}) {
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
    // invoke: QRegion boundingRegion(const class QTransform &)
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsItem14boundingRegionERK10QTransform(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItem", "boundingRegion", args)
  }

  return
}

// ~QGraphicsObject()
func (this *QGraphicsObject) Freeqgraphicsobject(args ...interface{}) () {
  // ~QGraphicsObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsObjectD0Ev
    // invoke: void ~QGraphicsObject()
    C.C_ZN15QGraphicsObjectD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsObject", "~QGraphicsObject", args)
  }

  return
}

// QGraphicsObject(class QGraphicsItem *)
func NewQGraphicsObject(args ...interface{}) *QGraphicsObject {
  // QGraphicsObject(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsObjectC1EP13QGraphicsItem
    // invoke: void QGraphicsObject(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QGraphicsObjectC2EP13QGraphicsItem(arg0)
    return &QGraphicsObject{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsObject", "QGraphicsObject", args)
  }

  return nil // QGraphicsObject{Qclsinst:qthis}
}

// metaObject()
func (this *QGraphicsObject) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsObject10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QGraphicsObject10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsObject", "metaObject", args)
  }

  return
}

// opaqueArea()
func (this *QGraphicsSimpleTextItem) Opaquearea(args ...interface{}) (ret interface{}) {
  // opaqueArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem10opaqueAreaEv
    // invoke: QPainterPath opaqueArea()
    var ret0 = C.C_ZNK23QGraphicsSimpleTextItem10opaqueAreaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "opaqueArea", args)
  }

  return
}

// ~QGraphicsSimpleTextItem()
func (this *QGraphicsSimpleTextItem) Freeqgraphicssimpletextitem(args ...interface{}) () {
  // ~QGraphicsSimpleTextItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSimpleTextItemD0Ev
    // invoke: void ~QGraphicsSimpleTextItem()
    C.C_ZN23QGraphicsSimpleTextItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "~QGraphicsSimpleTextItem", args)
  }

  return
}

// boundingRect()
func (this *QGraphicsSimpleTextItem) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK23QGraphicsSimpleTextItem12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "boundingRect", args)
  }

  return
}

// text()
func (this *QGraphicsSimpleTextItem) Text(args ...interface{}) (ret interface{}) {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK23QGraphicsSimpleTextItem4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "text", args)
  }

  return
}

// setText(const class QString &)
func (this *QGraphicsSimpleTextItem) Settext(args ...interface{}) () {
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
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN23QGraphicsSimpleTextItem7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "setText", args)
  }

  return
}

// contains(const class QPointF &)
func (this *QGraphicsSimpleTextItem) Contains(args ...interface{}) (ret interface{}) {
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
    // invoke: bool contains(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK23QGraphicsSimpleTextItem8containsERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "contains", args)
  }

  return
}

// isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsSimpleTextItem) Isobscuredby(args ...interface{}) (ret interface{}) {
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
    // invoke: bool isObscuredBy(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK23QGraphicsSimpleTextItem12isObscuredByEPK13QGraphicsItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "isObscuredBy", args)
  }

  return
}

// paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsSimpleTextItem) Paint(args ...interface{}) () {
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
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStyleOptionGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN23QGraphicsSimpleTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "paint", args)
  }

  return
}

// shape()
func (this *QGraphicsSimpleTextItem) Shape(args ...interface{}) (ret interface{}) {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem5shapeEv
    // invoke: QPainterPath shape()
    var ret0 = C.C_ZNK23QGraphicsSimpleTextItem5shapeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "shape", args)
  }

  return
}

// QGraphicsSimpleTextItem(class QGraphicsItem *)
func NewQGraphicsSimpleTextItem(args ...interface{}) *QGraphicsSimpleTextItem {
  // QGraphicsSimpleTextItem(class QGraphicsItem *)
  // QGraphicsSimpleTextItem(const class QString &, class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsSimpleTextItemC1EP13QGraphicsItem
    // invoke: void QGraphicsSimpleTextItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QGraphicsSimpleTextItemC2EP13QGraphicsItem(arg0)
    return &QGraphicsSimpleTextItem{Qclsinst:qthis}
  case 1:
    // invoke: _ZN23QGraphicsSimpleTextItemC1ERK7QStringP13QGraphicsItem
    // invoke: void QGraphicsSimpleTextItem(const class QString &, class QGraphicsItem *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QGraphicsSimpleTextItemC2ERK7QStringP13QGraphicsItem(arg0, arg1)
    return &QGraphicsSimpleTextItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "QGraphicsSimpleTextItem", args)
  }

  return nil // QGraphicsSimpleTextItem{Qclsinst:qthis}
}

// setFont(const class QFont &)
func (this *QGraphicsSimpleTextItem) Setfont(args ...interface{}) () {
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
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN23QGraphicsSimpleTextItem7setFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "setFont", args)
  }

  return
}

// font()
func (this *QGraphicsSimpleTextItem) Font(args ...interface{}) (ret interface{}) {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem4fontEv
    // invoke: QFont font()
    var ret0 = C.C_ZNK23QGraphicsSimpleTextItem4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "font", args)
  }

  return
}

// type()
func (this *QGraphicsSimpleTextItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsSimpleTextItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK23QGraphicsSimpleTextItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsSimpleTextItem", "type", args)
  }

  return
}

// <= body block end

