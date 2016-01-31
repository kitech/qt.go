package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtGui/qtextformat.h
// dst-file: /src/gui/qtextformat.go
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
  // proto:  qreal QTextLength::rawValue();
extern double C_ZNK11QTextLength8rawValueEv(void* qthis); // 2
  // proto:  qreal QTextLength::value(qreal maximumLength);
extern double C_ZNK11QTextLength5valueEd(void* qthis, double arg0); // 2
  // proto:  void QTextLength::QTextLength();
extern void* C_ZN11QTextLengthC2Ev(); // 1
  // proto:  QTextLength::Type QTextLength::type();
extern void C_ZNK11QTextLength4typeEv(void* qthis); // 2
  // proto:  void QTextImageFormat::setHeight(qreal height);
extern void C_ZN16QTextImageFormat9setHeightEd(void* qthis, double arg0); // 2
  // proto:  QString QTextImageFormat::name();
extern void* C_ZNK16QTextImageFormat4nameEv(void* qthis); // 2
  // proto:  void QTextImageFormat::setName(const QString & name);
extern void C_ZN16QTextImageFormat7setNameERK7QString(void* qthis, void* arg0); // 2
  // proto:  bool QTextImageFormat::isValid();
extern bool C_ZNK16QTextImageFormat7isValidEv(void* qthis); // 2
  // proto:  qreal QTextImageFormat::height();
extern double C_ZNK16QTextImageFormat6heightEv(void* qthis); // 2
  // proto:  qreal QTextImageFormat::width();
extern double C_ZNK16QTextImageFormat5widthEv(void* qthis); // 2
  // proto:  void QTextImageFormat::QTextImageFormat();
extern void* C_ZN16QTextImageFormatC2Ev(); // 3
  // proto:  void QTextImageFormat::setWidth(qreal width);
extern void C_ZN16QTextImageFormat8setWidthEd(void* qthis, double arg0); // 2
  // proto:  QTextBlockFormat QTextFormat::toBlockFormat();
extern void* C_ZNK11QTextFormat13toBlockFormatEv(void* qthis); // 4
  // proto:  int QTextFormat::intProperty(int propertyId);
extern int32_t C_ZNK11QTextFormat11intPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTextFormat::isFrameFormat();
extern bool C_ZNK11QTextFormat13isFrameFormatEv(void* qthis); // 2
  // proto:  QTextImageFormat QTextFormat::toImageFormat();
extern void* C_ZNK11QTextFormat13toImageFormatEv(void* qthis); // 4
  // proto:  QColor QTextFormat::colorProperty(int propertyId);
extern void* C_ZNK11QTextFormat13colorPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QTextFormat::stringProperty(int propertyId);
extern void* C_ZNK11QTextFormat14stringPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QTextFormat::doubleProperty(int propertyId);
extern double C_ZNK11QTextFormat14doublePropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextFormat::clearForeground();
extern void C_ZN11QTextFormat15clearForegroundEv(void* qthis); // 2
  // proto:  int QTextFormat::objectIndex();
extern int32_t C_ZNK11QTextFormat11objectIndexEv(void* qthis); // 4
  // proto:  bool QTextFormat::isImageFormat();
extern bool C_ZNK11QTextFormat13isImageFormatEv(void* qthis); // 2
  // proto:  int QTextFormat::objectType();
extern int32_t C_ZNK11QTextFormat10objectTypeEv(void* qthis); // 2
  // proto:  bool QTextFormat::hasProperty(int propertyId);
extern bool C_ZNK11QTextFormat11hasPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextFormat::setForeground(const QBrush & brush);
extern void C_ZN11QTextFormat13setForegroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  bool QTextFormat::isCharFormat();
extern bool C_ZNK11QTextFormat12isCharFormatEv(void* qthis); // 2
  // proto:  QTextTableCellFormat QTextFormat::toTableCellFormat();
extern void* C_ZNK11QTextFormat17toTableCellFormatEv(void* qthis); // 4
  // proto:  void QTextFormat::setBackground(const QBrush & brush);
extern void C_ZN11QTextFormat13setBackgroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  bool QTextFormat::isEmpty();
extern bool C_ZNK11QTextFormat7isEmptyEv(void* qthis); // 2
  // proto:  void QTextFormat::swap(QTextFormat & other);
extern void C_ZN11QTextFormat4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QTextFormat::QTextFormat();
extern void* C_ZN11QTextFormatC2Ev(); // 3
  // proto:  void QTextFormat::QTextFormat(int type);
extern void* C_ZN11QTextFormatC2Ei(int32_t arg0); // 3
  // proto:  void QTextFormat::QTextFormat(const QTextFormat & rhs);
extern void* C_ZN11QTextFormatC2ERKS_(void* arg0); // 3
  // proto:  bool QTextFormat::isBlockFormat();
extern bool C_ZNK11QTextFormat13isBlockFormatEv(void* qthis); // 2
  // proto:  void QTextFormat::clearBackground();
extern void C_ZN11QTextFormat15clearBackgroundEv(void* qthis); // 2
  // proto:  int QTextFormat::type();
extern int32_t C_ZNK11QTextFormat4typeEv(void* qthis); // 4
  // proto:  QBrush QTextFormat::foreground();
extern void* C_ZNK11QTextFormat10foregroundEv(void* qthis); // 2
  // proto:  QTextCharFormat QTextFormat::toCharFormat();
extern void* C_ZNK11QTextFormat12toCharFormatEv(void* qthis); // 4
  // proto:  void QTextFormat::~QTextFormat();
extern void C_ZN11QTextFormatD2Ev(void* qthis); // 4
  // proto:  bool QTextFormat::isValid();
extern bool C_ZNK11QTextFormat7isValidEv(void* qthis); // 2
  // proto:  bool QTextFormat::boolProperty(int propertyId);
extern bool C_ZNK11QTextFormat12boolPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextLength QTextFormat::lengthProperty(int propertyId);
extern void* C_ZNK11QTextFormat14lengthPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  QBrush QTextFormat::background();
extern void* C_ZNK11QTextFormat10backgroundEv(void* qthis); // 2
  // proto:  void QTextFormat::setProperty(int propertyId, const QVariant & value);
extern void C_ZN11QTextFormat11setPropertyEiRK8QVariant(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QTextListFormat QTextFormat::toListFormat();
extern void* C_ZNK11QTextFormat12toListFormatEv(void* qthis); // 4
  // proto:  QMap<int, QVariant> QTextFormat::properties();
extern void C_ZNK11QTextFormat10propertiesEv(void* qthis); // 4
  // proto:  QTextFrameFormat QTextFormat::toFrameFormat();
extern void* C_ZNK11QTextFormat13toFrameFormatEv(void* qthis); // 4
  // proto:  QPen QTextFormat::penProperty(int propertyId);
extern void* C_ZNK11QTextFormat11penPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTextFormat::isTableCellFormat();
extern bool C_ZNK11QTextFormat17isTableCellFormatEv(void* qthis); // 2
  // proto:  void QTextFormat::setObjectIndex(int object);
extern void C_ZN11QTextFormat14setObjectIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::LayoutDirection QTextFormat::layoutDirection();
extern void C_ZNK11QTextFormat15layoutDirectionEv(void* qthis); // 2
  // proto:  void QTextFormat::clearProperty(int propertyId);
extern void C_ZN11QTextFormat13clearPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  QBrush QTextFormat::brushProperty(int propertyId);
extern void* C_ZNK11QTextFormat13brushPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTextFormat::propertyCount();
extern int32_t C_ZNK11QTextFormat13propertyCountEv(void* qthis); // 4
  // proto:  void QTextFormat::merge(const QTextFormat & other);
extern void C_ZN11QTextFormat5mergeERKS_(void* qthis, void* arg0); // 4
  // proto:  QVector<QTextLength> QTextFormat::lengthVectorProperty(int propertyId);
extern void C_ZNK11QTextFormat20lengthVectorPropertyEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTextFormat::isTableFormat();
extern bool C_ZNK11QTextFormat13isTableFormatEv(void* qthis); // 2
  // proto:  QTextTableFormat QTextFormat::toTableFormat();
extern void* C_ZNK11QTextFormat13toTableFormatEv(void* qthis); // 4
  // proto:  QVariant QTextFormat::property(int propertyId);
extern void* C_ZNK11QTextFormat8propertyEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTextFormat::isListFormat();
extern bool C_ZNK11QTextFormat12isListFormatEv(void* qthis); // 2
  // proto:  void QTextFormat::setObjectType(int type);
extern void C_ZN11QTextFormat13setObjectTypeEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTextBlockFormat::setTextIndent(qreal aindent);
extern void C_ZN16QTextBlockFormat13setTextIndentEd(void* qthis, double arg0); // 2
  // proto:  int QTextBlockFormat::lineHeightType();
extern int32_t C_ZNK16QTextBlockFormat14lineHeightTypeEv(void* qthis); // 2
  // proto:  Qt::Alignment QTextBlockFormat::alignment();
extern void C_ZNK16QTextBlockFormat9alignmentEv(void* qthis); // 2
  // proto:  qreal QTextBlockFormat::leftMargin();
extern double C_ZNK16QTextBlockFormat10leftMarginEv(void* qthis); // 2
  // proto:  qreal QTextBlockFormat::topMargin();
extern double C_ZNK16QTextBlockFormat9topMarginEv(void* qthis); // 2
  // proto:  void QTextBlockFormat::setLeftMargin(qreal margin);
extern void C_ZN16QTextBlockFormat13setLeftMarginEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextBlockFormat::lineHeight();
extern double C_ZNK16QTextBlockFormat10lineHeightEv(void* qthis); // 2
  // proto:  qreal QTextBlockFormat::lineHeight(qreal scriptLineHeight, qreal scaling);
extern double C_ZNK16QTextBlockFormat10lineHeightEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QTextBlockFormat::setTopMargin(qreal margin);
extern void C_ZN16QTextBlockFormat12setTopMarginEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextBlockFormat::rightMargin();
extern double C_ZNK16QTextBlockFormat11rightMarginEv(void* qthis); // 2
  // proto:  bool QTextBlockFormat::isValid();
extern bool C_ZNK16QTextBlockFormat7isValidEv(void* qthis); // 2
  // proto:  void QTextBlockFormat::setIndent(int indent);
extern void C_ZN16QTextBlockFormat9setIndentEi(void* qthis, int32_t arg0); // 2
  // proto:  PageBreakFlags QTextBlockFormat::pageBreakPolicy();
extern void C_ZNK16QTextBlockFormat15pageBreakPolicyEv(void* qthis); // 2
  // proto:  qreal QTextBlockFormat::bottomMargin();
extern double C_ZNK16QTextBlockFormat12bottomMarginEv(void* qthis); // 2
  // proto:  void QTextBlockFormat::setLineHeight(qreal height, int heightType);
extern void C_ZN16QTextBlockFormat13setLineHeightEdi(void* qthis, double arg0, int32_t arg1); // 2
  // proto:  void QTextBlockFormat::setNonBreakableLines(bool b);
extern void C_ZN16QTextBlockFormat20setNonBreakableLinesEb(void* qthis, bool arg0); // 2
  // proto:  void QTextBlockFormat::setRightMargin(qreal margin);
extern void C_ZN16QTextBlockFormat14setRightMarginEd(void* qthis, double arg0); // 2
  // proto:  int QTextBlockFormat::indent();
extern int32_t C_ZNK16QTextBlockFormat6indentEv(void* qthis); // 2
  // proto:  QList<QTextOption::Tab> QTextBlockFormat::tabPositions();
extern void C_ZNK16QTextBlockFormat12tabPositionsEv(void* qthis); // 4
  // proto:  qreal QTextBlockFormat::textIndent();
extern double C_ZNK16QTextBlockFormat10textIndentEv(void* qthis); // 2
  // proto:  void QTextBlockFormat::setBottomMargin(qreal margin);
extern void C_ZN16QTextBlockFormat15setBottomMarginEd(void* qthis, double arg0); // 2
  // proto:  bool QTextBlockFormat::nonBreakableLines();
extern bool C_ZNK16QTextBlockFormat17nonBreakableLinesEv(void* qthis); // 2
  // proto:  void QTextBlockFormat::QTextBlockFormat();
extern void* C_ZN16QTextBlockFormatC2Ev(); // 3
  // proto:  void QTextCharFormat::setTableCellColumnSpan(int tableCellColumnSpan);
extern void C_ZN15QTextCharFormat22setTableCellColumnSpanEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTextCharFormat::setAnchorHref(const QString & value);
extern void C_ZN15QTextCharFormat13setAnchorHrefERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QTextCharFormat::setFontOverline(bool overline);
extern void C_ZN15QTextCharFormat15setFontOverlineEb(void* qthis, bool arg0); // 2
  // proto:  void QTextCharFormat::setTextOutline(const QPen & pen);
extern void C_ZN15QTextCharFormat14setTextOutlineERK4QPen(void* qthis, void* arg0); // 2
  // proto:  QString QTextCharFormat::fontFamily();
extern void* C_ZNK15QTextCharFormat10fontFamilyEv(void* qthis); // 2
  // proto:  QStringList QTextCharFormat::anchorNames();
extern void C_ZNK15QTextCharFormat11anchorNamesEv(void* qthis); // 4
  // proto:  QFont::StyleHint QTextCharFormat::fontStyleHint();
extern void C_ZNK15QTextCharFormat13fontStyleHintEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontStretch(int factor);
extern void C_ZN15QTextCharFormat14setFontStretchEi(void* qthis, int32_t arg0); // 2
  // proto:  QFont QTextCharFormat::font();
extern void* C_ZNK15QTextCharFormat4fontEv(void* qthis); // 4
  // proto:  bool QTextCharFormat::fontItalic();
extern bool C_ZNK15QTextCharFormat10fontItalicEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::fontKerning();
extern bool C_ZNK15QTextCharFormat11fontKerningEv(void* qthis); // 2
  // proto:  qreal QTextCharFormat::fontWordSpacing();
extern double C_ZNK15QTextCharFormat15fontWordSpacingEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::fontOverline();
extern bool C_ZNK15QTextCharFormat12fontOverlineEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontUnderline(bool underline);
extern void C_ZN15QTextCharFormat16setFontUnderlineEb(void* qthis, bool arg0); // 2
  // proto:  QTextCharFormat::UnderlineStyle QTextCharFormat::underlineStyle();
extern void C_ZNK15QTextCharFormat14underlineStyleEv(void* qthis); // 2
  // proto:  QFont::StyleStrategy QTextCharFormat::fontStyleStrategy();
extern void C_ZNK15QTextCharFormat17fontStyleStrategyEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontWordSpacing(qreal spacing);
extern void C_ZN15QTextCharFormat18setFontWordSpacingEd(void* qthis, double arg0); // 2
  // proto:  void QTextCharFormat::setFontWeight(int weight);
extern void C_ZN15QTextCharFormat13setFontWeightEi(void* qthis, int32_t arg0); // 2
  // proto:  QTextCharFormat::VerticalAlignment QTextCharFormat::verticalAlignment();
extern void C_ZNK15QTextCharFormat17verticalAlignmentEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontFamily(const QString & family);
extern void C_ZN15QTextCharFormat13setFontFamilyERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QTextCharFormat::setAnchorNames(const QStringList & names);
extern void C_ZN15QTextCharFormat14setAnchorNamesERK11QStringList(void* qthis, void* arg0); // 2
  // proto:  void QTextCharFormat::setTableCellRowSpan(int tableCellRowSpan);
extern void C_ZN15QTextCharFormat19setTableCellRowSpanEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTextCharFormat::setFontKerning(bool enable);
extern void C_ZN15QTextCharFormat14setFontKerningEb(void* qthis, bool arg0); // 2
  // proto:  void QTextCharFormat::setFontLetterSpacing(qreal spacing);
extern void C_ZN15QTextCharFormat20setFontLetterSpacingEd(void* qthis, double arg0); // 2
  // proto:  QFont::SpacingType QTextCharFormat::fontLetterSpacingType();
extern void C_ZNK15QTextCharFormat21fontLetterSpacingTypeEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::fontUnderline();
extern bool C_ZNK15QTextCharFormat13fontUnderlineEv(void* qthis); // 4
  // proto:  QFont::HintingPreference QTextCharFormat::fontHintingPreference();
extern void C_ZNK15QTextCharFormat21fontHintingPreferenceEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontStrikeOut(bool strikeOut);
extern void C_ZN15QTextCharFormat16setFontStrikeOutEb(void* qthis, bool arg0); // 2
  // proto:  QString QTextCharFormat::anchorName();
extern void* C_ZNK15QTextCharFormat10anchorNameEv(void* qthis); // 4
  // proto:  qreal QTextCharFormat::fontPointSize();
extern double C_ZNK15QTextCharFormat13fontPointSizeEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::fontFixedPitch();
extern bool C_ZNK15QTextCharFormat14fontFixedPitchEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::isValid();
extern bool C_ZNK15QTextCharFormat7isValidEv(void* qthis); // 2
  // proto:  QString QTextCharFormat::toolTip();
extern void* C_ZNK15QTextCharFormat7toolTipEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontPointSize(qreal size);
extern void C_ZN15QTextCharFormat16setFontPointSizeEd(void* qthis, double arg0); // 2
  // proto:  void QTextCharFormat::setAnchor(bool anchor);
extern void C_ZN15QTextCharFormat9setAnchorEb(void* qthis, bool arg0); // 2
  // proto:  void QTextCharFormat::setAnchorName(const QString & name);
extern void C_ZN15QTextCharFormat13setAnchorNameERK7QString(void* qthis, void* arg0); // 2
  // proto:  qreal QTextCharFormat::fontLetterSpacing();
extern double C_ZNK15QTextCharFormat17fontLetterSpacingEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::isAnchor();
extern bool C_ZNK15QTextCharFormat8isAnchorEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontItalic(bool italic);
extern void C_ZN15QTextCharFormat13setFontItalicEb(void* qthis, bool arg0); // 2
  // proto:  void QTextCharFormat::QTextCharFormat();
extern void* C_ZN15QTextCharFormatC2Ev(); // 3
  // proto:  QColor QTextCharFormat::underlineColor();
extern void* C_ZNK15QTextCharFormat14underlineColorEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setUnderlineColor(const QColor & color);
extern void C_ZN15QTextCharFormat17setUnderlineColorERK6QColor(void* qthis, void* arg0); // 2
  // proto:  int QTextCharFormat::fontWeight();
extern int32_t C_ZNK15QTextCharFormat10fontWeightEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setToolTip(const QString & tip);
extern void C_ZN15QTextCharFormat10setToolTipERK7QString(void* qthis, void* arg0); // 2
  // proto:  int QTextCharFormat::tableCellColumnSpan();
extern int32_t C_ZNK15QTextCharFormat19tableCellColumnSpanEv(void* qthis); // 2
  // proto:  QString QTextCharFormat::anchorHref();
extern void* C_ZNK15QTextCharFormat10anchorHrefEv(void* qthis); // 2
  // proto:  QFont::Capitalization QTextCharFormat::fontCapitalization();
extern void C_ZNK15QTextCharFormat18fontCapitalizationEv(void* qthis); // 2
  // proto:  QPen QTextCharFormat::textOutline();
extern void* C_ZNK15QTextCharFormat11textOutlineEv(void* qthis); // 2
  // proto:  int QTextCharFormat::tableCellRowSpan();
extern int32_t C_ZNK15QTextCharFormat16tableCellRowSpanEv(void* qthis); // 2
  // proto:  void QTextCharFormat::setFontFixedPitch(bool fixedPitch);
extern void C_ZN15QTextCharFormat17setFontFixedPitchEb(void* qthis, bool arg0); // 2
  // proto:  void QTextCharFormat::setFont(const QFont & font);
extern void C_ZN15QTextCharFormat7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  int QTextCharFormat::fontStretch();
extern int32_t C_ZNK15QTextCharFormat11fontStretchEv(void* qthis); // 2
  // proto:  bool QTextCharFormat::fontStrikeOut();
extern bool C_ZNK15QTextCharFormat13fontStrikeOutEv(void* qthis); // 2
  // proto:  void QTextTableFormat::clearColumnWidthConstraints();
extern void C_ZN16QTextTableFormat27clearColumnWidthConstraintsEv(void* qthis); // 2
  // proto:  QVector<QTextLength> QTextTableFormat::columnWidthConstraints();
extern void C_ZNK16QTextTableFormat22columnWidthConstraintsEv(void* qthis); // 2
  // proto:  void QTextTableFormat::QTextTableFormat();
extern void* C_ZN16QTextTableFormatC2Ev(); // 3
  // proto:  void QTextTableFormat::setHeaderRowCount(int count);
extern void C_ZN16QTextTableFormat17setHeaderRowCountEi(void* qthis, int32_t arg0); // 2
  // proto:  int QTextTableFormat::headerRowCount();
extern int32_t C_ZNK16QTextTableFormat14headerRowCountEv(void* qthis); // 2
  // proto:  void QTextTableFormat::setColumns(int columns);
extern void C_ZN16QTextTableFormat10setColumnsEi(void* qthis, int32_t arg0); // 2
  // proto:  Qt::Alignment QTextTableFormat::alignment();
extern void C_ZNK16QTextTableFormat9alignmentEv(void* qthis); // 2
  // proto:  bool QTextTableFormat::isValid();
extern bool C_ZNK16QTextTableFormat7isValidEv(void* qthis); // 2
  // proto:  qreal QTextTableFormat::cellPadding();
extern double C_ZNK16QTextTableFormat11cellPaddingEv(void* qthis); // 2
  // proto:  void QTextTableFormat::setCellSpacing(qreal spacing);
extern void C_ZN16QTextTableFormat14setCellSpacingEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextTableFormat::cellSpacing();
extern double C_ZNK16QTextTableFormat11cellSpacingEv(void* qthis); // 2
  // proto:  void QTextTableFormat::setCellPadding(qreal padding);
extern void C_ZN16QTextTableFormat14setCellPaddingEd(void* qthis, double arg0); // 2
  // proto:  int QTextTableFormat::columns();
extern int32_t C_ZNK16QTextTableFormat7columnsEv(void* qthis); // 2
  // proto:  void QTextTableCellFormat::setLeftPadding(qreal padding);
extern void C_ZN20QTextTableCellFormat14setLeftPaddingEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextTableCellFormat::leftPadding();
extern double C_ZNK20QTextTableCellFormat11leftPaddingEv(void* qthis); // 2
  // proto:  void QTextTableCellFormat::setBottomPadding(qreal padding);
extern void C_ZN20QTextTableCellFormat16setBottomPaddingEd(void* qthis, double arg0); // 2
  // proto:  void QTextTableCellFormat::QTextTableCellFormat();
extern void* C_ZN20QTextTableCellFormatC2Ev(); // 3
  // proto:  void QTextTableCellFormat::setTopPadding(qreal padding);
extern void C_ZN20QTextTableCellFormat13setTopPaddingEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextTableCellFormat::rightPadding();
extern double C_ZNK20QTextTableCellFormat12rightPaddingEv(void* qthis); // 2
  // proto:  void QTextTableCellFormat::setRightPadding(qreal padding);
extern void C_ZN20QTextTableCellFormat15setRightPaddingEd(void* qthis, double arg0); // 2
  // proto:  bool QTextTableCellFormat::isValid();
extern bool C_ZNK20QTextTableCellFormat7isValidEv(void* qthis); // 2
  // proto:  qreal QTextTableCellFormat::topPadding();
extern double C_ZNK20QTextTableCellFormat10topPaddingEv(void* qthis); // 2
  // proto:  void QTextTableCellFormat::setPadding(qreal padding);
extern void C_ZN20QTextTableCellFormat10setPaddingEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextTableCellFormat::bottomPadding();
extern double C_ZNK20QTextTableCellFormat13bottomPaddingEv(void* qthis); // 2
  // proto:  QTextListFormat::Style QTextListFormat::style();
extern void C_ZNK15QTextListFormat5styleEv(void* qthis); // 2
  // proto:  QString QTextListFormat::numberPrefix();
extern void* C_ZNK15QTextListFormat12numberPrefixEv(void* qthis); // 2
  // proto:  int QTextListFormat::indent();
extern int32_t C_ZNK15QTextListFormat6indentEv(void* qthis); // 2
  // proto:  void QTextListFormat::setNumberPrefix(const QString & numberPrefix);
extern void C_ZN15QTextListFormat15setNumberPrefixERK7QString(void* qthis, void* arg0); // 2
  // proto:  bool QTextListFormat::isValid();
extern bool C_ZNK15QTextListFormat7isValidEv(void* qthis); // 2
  // proto:  void QTextListFormat::QTextListFormat();
extern void* C_ZN15QTextListFormatC2Ev(); // 3
  // proto:  void QTextListFormat::setIndent(int indent);
extern void C_ZN15QTextListFormat9setIndentEi(void* qthis, int32_t arg0); // 2
  // proto:  QString QTextListFormat::numberSuffix();
extern void* C_ZNK15QTextListFormat12numberSuffixEv(void* qthis); // 2
  // proto:  void QTextListFormat::setNumberSuffix(const QString & numberSuffix);
extern void C_ZN15QTextListFormat15setNumberSuffixERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QTextFrameFormat::setBorderBrush(const QBrush & brush);
extern void C_ZN16QTextFrameFormat14setBorderBrushERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  QTextLength QTextFrameFormat::height();
extern void* C_ZNK16QTextFrameFormat6heightEv(void* qthis); // 2
  // proto:  void QTextFrameFormat::setWidth(qreal width);
extern void C_ZN16QTextFrameFormat8setWidthEd(void* qthis, double arg0); // 2
  // proto:  void QTextFrameFormat::setWidth(const QTextLength & length);
extern void C_ZN16QTextFrameFormat8setWidthERK11QTextLength(void* qthis, void* arg0); // 2
  // proto:  qreal QTextFrameFormat::topMargin();
extern double C_ZNK16QTextFrameFormat9topMarginEv(void* qthis); // 4
  // proto:  qreal QTextFrameFormat::border();
extern double C_ZNK16QTextFrameFormat6borderEv(void* qthis); // 2
  // proto:  QTextLength QTextFrameFormat::width();
extern void* C_ZNK16QTextFrameFormat5widthEv(void* qthis); // 2
  // proto:  void QTextFrameFormat::setLeftMargin(qreal margin);
extern void C_ZN16QTextFrameFormat13setLeftMarginEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextFrameFormat::leftMargin();
extern double C_ZNK16QTextFrameFormat10leftMarginEv(void* qthis); // 4
  // proto:  void QTextFrameFormat::setPadding(qreal padding);
extern void C_ZN16QTextFrameFormat10setPaddingEd(void* qthis, double arg0); // 2
  // proto:  void QTextFrameFormat::setTopMargin(qreal margin);
extern void C_ZN16QTextFrameFormat12setTopMarginEd(void* qthis, double arg0); // 2
  // proto:  qreal QTextFrameFormat::rightMargin();
extern double C_ZNK16QTextFrameFormat11rightMarginEv(void* qthis); // 4
  // proto:  bool QTextFrameFormat::isValid();
extern bool C_ZNK16QTextFrameFormat7isValidEv(void* qthis); // 2
  // proto:  PageBreakFlags QTextFrameFormat::pageBreakPolicy();
extern void C_ZNK16QTextFrameFormat15pageBreakPolicyEv(void* qthis); // 2
  // proto:  qreal QTextFrameFormat::padding();
extern double C_ZNK16QTextFrameFormat7paddingEv(void* qthis); // 2
  // proto:  void QTextFrameFormat::setMargin(qreal margin);
extern void C_ZN16QTextFrameFormat9setMarginEd(void* qthis, double arg0); // 4
  // proto:  qreal QTextFrameFormat::bottomMargin();
extern double C_ZNK16QTextFrameFormat12bottomMarginEv(void* qthis); // 4
  // proto:  void QTextFrameFormat::QTextFrameFormat();
extern void* C_ZN16QTextFrameFormatC2Ev(); // 3
  // proto:  void QTextFrameFormat::setHeight(qreal height);
extern void C_ZN16QTextFrameFormat9setHeightEd(void* qthis, double arg0); // 2
  // proto:  void QTextFrameFormat::setHeight(const QTextLength & height);
extern void C_ZN16QTextFrameFormat9setHeightERK11QTextLength(void* qthis, void* arg0); // 2
  // proto:  QBrush QTextFrameFormat::borderBrush();
extern void* C_ZNK16QTextFrameFormat11borderBrushEv(void* qthis); // 2
  // proto:  void QTextFrameFormat::setRightMargin(qreal margin);
extern void C_ZN16QTextFrameFormat14setRightMarginEd(void* qthis, double arg0); // 2
  // proto:  void QTextFrameFormat::setBorder(qreal border);
extern void C_ZN16QTextFrameFormat9setBorderEd(void* qthis, double arg0); // 2
  // proto:  void QTextFrameFormat::setBottomMargin(qreal margin);
extern void C_ZN16QTextFrameFormat15setBottomMarginEd(void* qthis, double arg0); // 2
  // proto:  QTextFrameFormat::Position QTextFrameFormat::position();
extern void C_ZNK16QTextFrameFormat8positionEv(void* qthis); // 2
  // proto:  QTextFrameFormat::BorderStyle QTextFrameFormat::borderStyle();
extern void C_ZNK16QTextFrameFormat11borderStyleEv(void* qthis); // 2
  // proto:  qreal QTextFrameFormat::margin();
extern double C_ZNK16QTextFrameFormat6marginEv(void* qthis); // 2
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

// class sizeof(QTextLength)=16
type QTextLength struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextImageFormat)=1
type QTextImageFormat struct {
  /*qbase*/ QTextCharFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextFormat)=1
type QTextFormat struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextBlockFormat)=1
type QTextBlockFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextCharFormat)=1
type QTextCharFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextTableFormat)=1
type QTextTableFormat struct {
  /*qbase*/ QTextFrameFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextTableCellFormat)=1
type QTextTableCellFormat struct {
  /*qbase*/ QTextCharFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextListFormat)=1
type QTextListFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextFrameFormat)=1
type QTextFrameFormat struct {
  /*qbase*/ QTextFormat;
  qclsinst unsafe.Pointer /* *C.void */;
}

// rawValue()
func (this *QTextLength) Rawvalue(args ...interface{}) (ret interface{}) {
  // rawValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLength8rawValueEv
    // invoke: qreal rawValue()
    var ret0 = C.C_ZNK11QTextLength8rawValueEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextLength", "rawValue", args)
  }

  return
}

// value(qreal)
func (this *QTextLength) Value(args ...interface{}) (ret interface{}) {
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
    // invoke: qreal value(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextLength5valueEd(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextLength", "value", args)
  }

  return
}

// QTextLength()
func NewQTextLength(args ...interface{}) *QTextLength {
  // QTextLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLengthC1Ev
    // invoke: void QTextLength()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextLengthC2Ev()
    return &QTextLength{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextLength", "QTextLength", args)
  }

  return nil // QTextLength{qclsinst:qthis}
}

// type()
func (this *QTextLength) Type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLength4typeEv
    // invoke: QTextLength::Type type()
    C.C_ZNK11QTextLength4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLength", "type", args)
  }

  return
}

// setHeight(qreal)
func (this *QTextImageFormat) Setheight(args ...interface{}) () {
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
    // invoke: void setHeight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextImageFormat9setHeightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setHeight", args)
  }

  return
}

// name()
func (this *QTextImageFormat) Name(args ...interface{}) (ret interface{}) {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat4nameEv
    // invoke: QString name()
    var ret0 = C.C_ZNK16QTextImageFormat4nameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextImageFormat", "name", args)
  }

  return
}

// setName(const class QString &)
func (this *QTextImageFormat) Setname(args ...interface{}) () {
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
    // invoke: void setName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextImageFormat7setNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setName", args)
  }

  return
}

// isValid()
func (this *QTextImageFormat) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK16QTextImageFormat7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextImageFormat", "isValid", args)
  }

  return
}

// height()
func (this *QTextImageFormat) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat6heightEv
    // invoke: qreal height()
    var ret0 = C.C_ZNK16QTextImageFormat6heightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextImageFormat", "height", args)
  }

  return
}

// width()
func (this *QTextImageFormat) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextImageFormat5widthEv
    // invoke: qreal width()
    var ret0 = C.C_ZNK16QTextImageFormat5widthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextImageFormat", "width", args)
  }

  return
}

// QTextImageFormat()
func NewQTextImageFormat(args ...interface{}) *QTextImageFormat {
  // QTextImageFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextImageFormatC1Ev
    // invoke: void QTextImageFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTextImageFormatC2Ev()
    return &QTextImageFormat{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextImageFormat", "QTextImageFormat", args)
  }

  return nil // QTextImageFormat{qclsinst:qthis}
}

// setWidth(qreal)
func (this *QTextImageFormat) Setwidth(args ...interface{}) () {
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
    // invoke: void setWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextImageFormat8setWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextImageFormat", "setWidth", args)
  }

  return
}

// toBlockFormat()
func (this *QTextFormat) Toblockformat(args ...interface{}) (ret interface{}) {
  // toBlockFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toBlockFormatEv
    // invoke: QTextBlockFormat toBlockFormat()
    var ret0 = C.C_ZNK11QTextFormat13toBlockFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlockFormat{}) // "QTextBlockFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "toBlockFormat", args)
  }

  return
}

// intProperty(int)
func (this *QTextFormat) Intproperty(args ...interface{}) (ret interface{}) {
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
    // invoke: int intProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextFormat11intPropertyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "intProperty", args)
  }

  return
}

// isFrameFormat()
func (this *QTextFormat) Isframeformat(args ...interface{}) (ret interface{}) {
  // isFrameFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isFrameFormatEv
    // invoke: bool isFrameFormat()
    var ret0 = C.C_ZNK11QTextFormat13isFrameFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "isFrameFormat", args)
  }

  return
}

// toImageFormat()
func (this *QTextFormat) Toimageformat(args ...interface{}) (ret interface{}) {
  // toImageFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toImageFormatEv
    // invoke: QTextImageFormat toImageFormat()
    var ret0 = C.C_ZNK11QTextFormat13toImageFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextImageFormat{}) // "QTextImageFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "toImageFormat", args)
  }

  return
}

// colorProperty(int)
func (this *QTextFormat) Colorproperty(args ...interface{}) (ret interface{}) {
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
    // invoke: QColor colorProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextFormat13colorPropertyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "colorProperty", args)
  }

  return
}

// stringProperty(int)
func (this *QTextFormat) Stringproperty(args ...interface{}) (ret interface{}) {
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
    // invoke: QString stringProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextFormat14stringPropertyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "stringProperty", args)
  }

  return
}

// doubleProperty(int)
func (this *QTextFormat) Doubleproperty(args ...interface{}) (ret interface{}) {
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
    // invoke: qreal doubleProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextFormat14doublePropertyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "doubleProperty", args)
  }

  return
}

// clearForeground()
func (this *QTextFormat) Clearforeground(args ...interface{}) () {
  // clearForeground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat15clearForegroundEv
    // invoke: void clearForeground()
    C.C_ZN11QTextFormat15clearForegroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "clearForeground", args)
  }

  return
}

// objectIndex()
func (this *QTextFormat) Objectindex(args ...interface{}) (ret interface{}) {
  // objectIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat11objectIndexEv
    // invoke: int objectIndex()
    var ret0 = C.C_ZNK11QTextFormat11objectIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "objectIndex", args)
  }

  return
}

// isImageFormat()
func (this *QTextFormat) Isimageformat(args ...interface{}) (ret interface{}) {
  // isImageFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isImageFormatEv
    // invoke: bool isImageFormat()
    var ret0 = C.C_ZNK11QTextFormat13isImageFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "isImageFormat", args)
  }

  return
}

// objectType()
func (this *QTextFormat) Objecttype(args ...interface{}) (ret interface{}) {
  // objectType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10objectTypeEv
    // invoke: int objectType()
    var ret0 = C.C_ZNK11QTextFormat10objectTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "objectType", args)
  }

  return
}

// hasProperty(int)
func (this *QTextFormat) Hasproperty(args ...interface{}) (ret interface{}) {
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
    // invoke: bool hasProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextFormat11hasPropertyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "hasProperty", args)
  }

  return
}

// setForeground(const class QBrush &)
func (this *QTextFormat) Setforeground(args ...interface{}) () {
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
    // invoke: void setForeground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat13setForegroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "setForeground", args)
  }

  return
}

// isCharFormat()
func (this *QTextFormat) Ischarformat(args ...interface{}) (ret interface{}) {
  // isCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12isCharFormatEv
    // invoke: bool isCharFormat()
    var ret0 = C.C_ZNK11QTextFormat12isCharFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "isCharFormat", args)
  }

  return
}

// toTableCellFormat()
func (this *QTextFormat) Totablecellformat(args ...interface{}) (ret interface{}) {
  // toTableCellFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat17toTableCellFormatEv
    // invoke: QTextTableCellFormat toTableCellFormat()
    var ret0 = C.C_ZNK11QTextFormat17toTableCellFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextTableCellFormat{}) // "QTextTableCellFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "toTableCellFormat", args)
  }

  return
}

// setBackground(const class QBrush &)
func (this *QTextFormat) Setbackground(args ...interface{}) () {
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
    // invoke: void setBackground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat13setBackgroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "setBackground", args)
  }

  return
}

// isEmpty()
func (this *QTextFormat) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK11QTextFormat7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "isEmpty", args)
  }

  return
}

// swap(class QTextFormat &)
func (this *QTextFormat) Swap(args ...interface{}) () {
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
    // invoke: void swap(class QTextFormat &)
    var arg0 = args[0].(QTextFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "swap", args)
  }

  return
}

// QTextFormat()
func NewQTextFormat(args ...interface{}) *QTextFormat {
  // QTextFormat()
  // QTextFormat(int)
  // QTextFormat(const class QTextFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QTextFormat{}) // "const QTextFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormatC1Ev
    // invoke: void QTextFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextFormatC2Ev()
    return &QTextFormat{qclsinst:qthis}
  case 1:
    // invoke: _ZN11QTextFormatC1Ei
    // invoke: void QTextFormat(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextFormatC2Ei(arg0)
    return &QTextFormat{qclsinst:qthis}
  case 2:
    // invoke: _ZN11QTextFormatC1ERKS_
    // invoke: void QTextFormat(const class QTextFormat &)
    var arg0 = args[0].(QTextFormat).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextFormatC2ERKS_(arg0)
    return &QTextFormat{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextFormat", "QTextFormat", args)
  }

  return nil // QTextFormat{qclsinst:qthis}
}

// isBlockFormat()
func (this *QTextFormat) Isblockformat(args ...interface{}) (ret interface{}) {
  // isBlockFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isBlockFormatEv
    // invoke: bool isBlockFormat()
    var ret0 = C.C_ZNK11QTextFormat13isBlockFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "isBlockFormat", args)
  }

  return
}

// clearBackground()
func (this *QTextFormat) Clearbackground(args ...interface{}) () {
  // clearBackground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat15clearBackgroundEv
    // invoke: void clearBackground()
    C.C_ZN11QTextFormat15clearBackgroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "clearBackground", args)
  }

  return
}

// type()
func (this *QTextFormat) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK11QTextFormat4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "type", args)
  }

  return
}

// foreground()
func (this *QTextFormat) Foreground(args ...interface{}) (ret interface{}) {
  // foreground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10foregroundEv
    // invoke: QBrush foreground()
    var ret0 = C.C_ZNK11QTextFormat10foregroundEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "foreground", args)
  }

  return
}

// toCharFormat()
func (this *QTextFormat) Tocharformat(args ...interface{}) (ret interface{}) {
  // toCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12toCharFormatEv
    // invoke: QTextCharFormat toCharFormat()
    var ret0 = C.C_ZNK11QTextFormat12toCharFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCharFormat{}) // "QTextCharFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "toCharFormat", args)
  }

  return
}

// ~QTextFormat()
func (this *QTextFormat) Freeqtextformat(args ...interface{}) () {
  // ~QTextFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormatD0Ev
    // invoke: void ~QTextFormat()
    C.C_ZN11QTextFormatD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "~QTextFormat", args)
  }

  return
}

// isValid()
func (this *QTextFormat) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK11QTextFormat7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "isValid", args)
  }

  return
}

// boolProperty(int)
func (this *QTextFormat) Boolproperty(args ...interface{}) (ret interface{}) {
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
    // invoke: bool boolProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextFormat12boolPropertyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "boolProperty", args)
  }

  return
}

// lengthProperty(int)
func (this *QTextFormat) Lengthproperty(args ...interface{}) (ret interface{}) {
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
    // invoke: QTextLength lengthProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextFormat14lengthPropertyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextLength{}) // "QTextLength"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "lengthProperty", args)
  }

  return
}

// background()
func (this *QTextFormat) Background(args ...interface{}) (ret interface{}) {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10backgroundEv
    // invoke: QBrush background()
    var ret0 = C.C_ZNK11QTextFormat10backgroundEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "background", args)
  }

  return
}

// setProperty(int, const class QVariant &)
func (this *QTextFormat) Setproperty(args ...interface{}) () {
  // setProperty(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextFormat11setPropertyEiRK8QVariant
    // invoke: void setProperty(int, const class QVariant &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextFormat11setPropertyEiRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextFormat", "setProperty", args)
  }

  return
}

// toListFormat()
func (this *QTextFormat) Tolistformat(args ...interface{}) (ret interface{}) {
  // toListFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12toListFormatEv
    // invoke: QTextListFormat toListFormat()
    var ret0 = C.C_ZNK11QTextFormat12toListFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextListFormat{}) // "QTextListFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "toListFormat", args)
  }

  return
}

// properties()
func (this *QTextFormat) Properties(args ...interface{}) () {
  // properties()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat10propertiesEv
    // invoke: QMap<int, QVariant> properties()
    C.C_ZNK11QTextFormat10propertiesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "properties", args)
  }

  return
}

// toFrameFormat()
func (this *QTextFormat) Toframeformat(args ...interface{}) (ret interface{}) {
  // toFrameFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toFrameFormatEv
    // invoke: QTextFrameFormat toFrameFormat()
    var ret0 = C.C_ZNK11QTextFormat13toFrameFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextFrameFormat{}) // "QTextFrameFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "toFrameFormat", args)
  }

  return
}

// penProperty(int)
func (this *QTextFormat) Penproperty(args ...interface{}) (ret interface{}) {
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
    // invoke: QPen penProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextFormat11penPropertyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPen{}) // "QPen"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "penProperty", args)
  }

  return
}

// isTableCellFormat()
func (this *QTextFormat) Istablecellformat(args ...interface{}) (ret interface{}) {
  // isTableCellFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat17isTableCellFormatEv
    // invoke: bool isTableCellFormat()
    var ret0 = C.C_ZNK11QTextFormat17isTableCellFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "isTableCellFormat", args)
  }

  return
}

// setObjectIndex(int)
func (this *QTextFormat) Setobjectindex(args ...interface{}) () {
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
    // invoke: void setObjectIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat14setObjectIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "setObjectIndex", args)
  }

  return
}

// layoutDirection()
func (this *QTextFormat) Layoutdirection(args ...interface{}) () {
  // layoutDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat15layoutDirectionEv
    // invoke: Qt::LayoutDirection layoutDirection()
    C.C_ZNK11QTextFormat15layoutDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFormat", "layoutDirection", args)
  }

  return
}

// clearProperty(int)
func (this *QTextFormat) Clearproperty(args ...interface{}) () {
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
    // invoke: void clearProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat13clearPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "clearProperty", args)
  }

  return
}

// brushProperty(int)
func (this *QTextFormat) Brushproperty(args ...interface{}) (ret interface{}) {
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
    // invoke: QBrush brushProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextFormat13brushPropertyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "brushProperty", args)
  }

  return
}

// propertyCount()
func (this *QTextFormat) Propertycount(args ...interface{}) (ret interface{}) {
  // propertyCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13propertyCountEv
    // invoke: int propertyCount()
    var ret0 = C.C_ZNK11QTextFormat13propertyCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "propertyCount", args)
  }

  return
}

// merge(const class QTextFormat &)
func (this *QTextFormat) Merge(args ...interface{}) () {
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
    // invoke: void merge(const class QTextFormat &)
    var arg0 = args[0].(QTextFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat5mergeERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "merge", args)
  }

  return
}

// lengthVectorProperty(int)
func (this *QTextFormat) Lengthvectorproperty(args ...interface{}) () {
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
    // invoke: QVector<QTextLength> lengthVectorProperty(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextFormat20lengthVectorPropertyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "lengthVectorProperty", args)
  }

  return
}

// isTableFormat()
func (this *QTextFormat) Istableformat(args ...interface{}) (ret interface{}) {
  // isTableFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13isTableFormatEv
    // invoke: bool isTableFormat()
    var ret0 = C.C_ZNK11QTextFormat13isTableFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "isTableFormat", args)
  }

  return
}

// toTableFormat()
func (this *QTextFormat) Totableformat(args ...interface{}) (ret interface{}) {
  // toTableFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat13toTableFormatEv
    // invoke: QTextTableFormat toTableFormat()
    var ret0 = C.C_ZNK11QTextFormat13toTableFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextTableFormat{}) // "QTextTableFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "toTableFormat", args)
  }

  return
}

// property(int)
func (this *QTextFormat) Property(args ...interface{}) (ret interface{}) {
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
    // invoke: QVariant property(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextFormat8propertyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "property", args)
  }

  return
}

// isListFormat()
func (this *QTextFormat) Islistformat(args ...interface{}) (ret interface{}) {
  // isListFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextFormat12isListFormatEv
    // invoke: bool isListFormat()
    var ret0 = C.C_ZNK11QTextFormat12isListFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFormat", "isListFormat", args)
  }

  return
}

// setObjectType(int)
func (this *QTextFormat) Setobjecttype(args ...interface{}) () {
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
    // invoke: void setObjectType(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextFormat13setObjectTypeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFormat", "setObjectType", args)
  }

  return
}

// setTextIndent(qreal)
func (this *QTextBlockFormat) Settextindent(args ...interface{}) () {
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
    // invoke: void setTextIndent(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat13setTextIndentEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setTextIndent", args)
  }

  return
}

// lineHeightType()
func (this *QTextBlockFormat) Lineheighttype(args ...interface{}) (ret interface{}) {
  // lineHeightType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat14lineHeightTypeEv
    // invoke: int lineHeightType()
    var ret0 = C.C_ZNK16QTextBlockFormat14lineHeightTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "lineHeightType", args)
  }

  return
}

// alignment()
func (this *QTextBlockFormat) Alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK16QTextBlockFormat9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "alignment", args)
  }

  return
}

// leftMargin()
func (this *QTextBlockFormat) Leftmargin(args ...interface{}) (ret interface{}) {
  // leftMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat10leftMarginEv
    // invoke: qreal leftMargin()
    var ret0 = C.C_ZNK16QTextBlockFormat10leftMarginEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "leftMargin", args)
  }

  return
}

// topMargin()
func (this *QTextBlockFormat) Topmargin(args ...interface{}) (ret interface{}) {
  // topMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat9topMarginEv
    // invoke: qreal topMargin()
    var ret0 = C.C_ZNK16QTextBlockFormat9topMarginEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "topMargin", args)
  }

  return
}

// setLeftMargin(qreal)
func (this *QTextBlockFormat) Setleftmargin(args ...interface{}) () {
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
    // invoke: void setLeftMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat13setLeftMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setLeftMargin", args)
  }

  return
}

// lineHeight()
func (this *QTextBlockFormat) Lineheight(args ...interface{}) (ret interface{}) {
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
    // invoke: qreal lineHeight()
    var ret0 = C.C_ZNK16QTextBlockFormat10lineHeightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK16QTextBlockFormat10lineHeightEdd
    // invoke: qreal lineHeight(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK16QTextBlockFormat10lineHeightEdd(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "lineHeight", args)
  }

  return
}

// setTopMargin(qreal)
func (this *QTextBlockFormat) Settopmargin(args ...interface{}) () {
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
    // invoke: void setTopMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat12setTopMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setTopMargin", args)
  }

  return
}

// rightMargin()
func (this *QTextBlockFormat) Rightmargin(args ...interface{}) (ret interface{}) {
  // rightMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat11rightMarginEv
    // invoke: qreal rightMargin()
    var ret0 = C.C_ZNK16QTextBlockFormat11rightMarginEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "rightMargin", args)
  }

  return
}

// isValid()
func (this *QTextBlockFormat) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK16QTextBlockFormat7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "isValid", args)
  }

  return
}

// setIndent(int)
func (this *QTextBlockFormat) Setindent(args ...interface{}) () {
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
    // invoke: void setIndent(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat9setIndentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setIndent", args)
  }

  return
}

// pageBreakPolicy()
func (this *QTextBlockFormat) Pagebreakpolicy(args ...interface{}) () {
  // pageBreakPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat15pageBreakPolicyEv
    // invoke: PageBreakFlags pageBreakPolicy()
    C.C_ZNK16QTextBlockFormat15pageBreakPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "pageBreakPolicy", args)
  }

  return
}

// bottomMargin()
func (this *QTextBlockFormat) Bottommargin(args ...interface{}) (ret interface{}) {
  // bottomMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat12bottomMarginEv
    // invoke: qreal bottomMargin()
    var ret0 = C.C_ZNK16QTextBlockFormat12bottomMarginEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "bottomMargin", args)
  }

  return
}

// setLineHeight(qreal, int)
func (this *QTextBlockFormat) Setlineheight(args ...interface{}) () {
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
    // invoke: void setLineHeight(qreal, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN16QTextBlockFormat13setLineHeightEdi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setLineHeight", args)
  }

  return
}

// setNonBreakableLines(_Bool)
func (this *QTextBlockFormat) Setnonbreakablelines(args ...interface{}) () {
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
    // invoke: void setNonBreakableLines(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat20setNonBreakableLinesEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setNonBreakableLines", args)
  }

  return
}

// setRightMargin(qreal)
func (this *QTextBlockFormat) Setrightmargin(args ...interface{}) () {
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
    // invoke: void setRightMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat14setRightMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setRightMargin", args)
  }

  return
}

// indent()
func (this *QTextBlockFormat) Indent(args ...interface{}) (ret interface{}) {
  // indent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat6indentEv
    // invoke: int indent()
    var ret0 = C.C_ZNK16QTextBlockFormat6indentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "indent", args)
  }

  return
}

// tabPositions()
func (this *QTextBlockFormat) Tabpositions(args ...interface{}) () {
  // tabPositions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat12tabPositionsEv
    // invoke: QList<QTextOption::Tab> tabPositions()
    C.C_ZNK16QTextBlockFormat12tabPositionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "tabPositions", args)
  }

  return
}

// textIndent()
func (this *QTextBlockFormat) Textindent(args ...interface{}) (ret interface{}) {
  // textIndent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat10textIndentEv
    // invoke: qreal textIndent()
    var ret0 = C.C_ZNK16QTextBlockFormat10textIndentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "textIndent", args)
  }

  return
}

// setBottomMargin(qreal)
func (this *QTextBlockFormat) Setbottommargin(args ...interface{}) () {
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
    // invoke: void setBottomMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextBlockFormat15setBottomMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "setBottomMargin", args)
  }

  return
}

// nonBreakableLines()
func (this *QTextBlockFormat) Nonbreakablelines(args ...interface{}) (ret interface{}) {
  // nonBreakableLines()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextBlockFormat17nonBreakableLinesEv
    // invoke: bool nonBreakableLines()
    var ret0 = C.C_ZNK16QTextBlockFormat17nonBreakableLinesEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "nonBreakableLines", args)
  }

  return
}

// QTextBlockFormat()
func NewQTextBlockFormat(args ...interface{}) *QTextBlockFormat {
  // QTextBlockFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextBlockFormatC1Ev
    // invoke: void QTextBlockFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTextBlockFormatC2Ev()
    return &QTextBlockFormat{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextBlockFormat", "QTextBlockFormat", args)
  }

  return nil // QTextBlockFormat{qclsinst:qthis}
}

// setTableCellColumnSpan(int)
func (this *QTextCharFormat) Settablecellcolumnspan(args ...interface{}) () {
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
    // invoke: void setTableCellColumnSpan(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat22setTableCellColumnSpanEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTableCellColumnSpan", args)
  }

  return
}

// setAnchorHref(const class QString &)
func (this *QTextCharFormat) Setanchorhref(args ...interface{}) () {
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
    // invoke: void setAnchorHref(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat13setAnchorHrefERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorHref", args)
  }

  return
}

// setFontOverline(_Bool)
func (this *QTextCharFormat) Setfontoverline(args ...interface{}) () {
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
    // invoke: void setFontOverline(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat15setFontOverlineEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontOverline", args)
  }

  return
}

// setTextOutline(const class QPen &)
func (this *QTextCharFormat) Settextoutline(args ...interface{}) () {
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
    // invoke: void setTextOutline(const class QPen &)
    var arg0 = args[0].(QPen).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat14setTextOutlineERK4QPen(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTextOutline", args)
  }

  return
}

// fontFamily()
func (this *QTextCharFormat) Fontfamily(args ...interface{}) (ret interface{}) {
  // fontFamily()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10fontFamilyEv
    // invoke: QString fontFamily()
    var ret0 = C.C_ZNK15QTextCharFormat10fontFamilyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontFamily", args)
  }

  return
}

// anchorNames()
func (this *QTextCharFormat) Anchornames(args ...interface{}) () {
  // anchorNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11anchorNamesEv
    // invoke: QStringList anchorNames()
    C.C_ZNK15QTextCharFormat11anchorNamesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorNames", args)
  }

  return
}

// fontStyleHint()
func (this *QTextCharFormat) Fontstylehint(args ...interface{}) () {
  // fontStyleHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat13fontStyleHintEv
    // invoke: QFont::StyleHint fontStyleHint()
    C.C_ZNK15QTextCharFormat13fontStyleHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStyleHint", args)
  }

  return
}

// setFontStretch(int)
func (this *QTextCharFormat) Setfontstretch(args ...interface{}) () {
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
    // invoke: void setFontStretch(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat14setFontStretchEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontStretch", args)
  }

  return
}

// font()
func (this *QTextCharFormat) Font(args ...interface{}) (ret interface{}) {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat4fontEv
    // invoke: QFont font()
    var ret0 = C.C_ZNK15QTextCharFormat4fontEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "font", args)
  }

  return
}

// fontItalic()
func (this *QTextCharFormat) Fontitalic(args ...interface{}) (ret interface{}) {
  // fontItalic()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10fontItalicEv
    // invoke: bool fontItalic()
    var ret0 = C.C_ZNK15QTextCharFormat10fontItalicEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontItalic", args)
  }

  return
}

// fontKerning()
func (this *QTextCharFormat) Fontkerning(args ...interface{}) (ret interface{}) {
  // fontKerning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11fontKerningEv
    // invoke: bool fontKerning()
    var ret0 = C.C_ZNK15QTextCharFormat11fontKerningEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontKerning", args)
  }

  return
}

// fontWordSpacing()
func (this *QTextCharFormat) Fontwordspacing(args ...interface{}) (ret interface{}) {
  // fontWordSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat15fontWordSpacingEv
    // invoke: qreal fontWordSpacing()
    var ret0 = C.C_ZNK15QTextCharFormat15fontWordSpacingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontWordSpacing", args)
  }

  return
}

// fontOverline()
func (this *QTextCharFormat) Fontoverline(args ...interface{}) (ret interface{}) {
  // fontOverline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat12fontOverlineEv
    // invoke: bool fontOverline()
    var ret0 = C.C_ZNK15QTextCharFormat12fontOverlineEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontOverline", args)
  }

  return
}

// setFontUnderline(_Bool)
func (this *QTextCharFormat) Setfontunderline(args ...interface{}) () {
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
    // invoke: void setFontUnderline(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat16setFontUnderlineEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontUnderline", args)
  }

  return
}

// underlineStyle()
func (this *QTextCharFormat) Underlinestyle(args ...interface{}) () {
  // underlineStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat14underlineStyleEv
    // invoke: QTextCharFormat::UnderlineStyle underlineStyle()
    C.C_ZNK15QTextCharFormat14underlineStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "underlineStyle", args)
  }

  return
}

// fontStyleStrategy()
func (this *QTextCharFormat) Fontstylestrategy(args ...interface{}) () {
  // fontStyleStrategy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat17fontStyleStrategyEv
    // invoke: QFont::StyleStrategy fontStyleStrategy()
    C.C_ZNK15QTextCharFormat17fontStyleStrategyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStyleStrategy", args)
  }

  return
}

// setFontWordSpacing(qreal)
func (this *QTextCharFormat) Setfontwordspacing(args ...interface{}) () {
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
    // invoke: void setFontWordSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat18setFontWordSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontWordSpacing", args)
  }

  return
}

// setFontWeight(int)
func (this *QTextCharFormat) Setfontweight(args ...interface{}) () {
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
    // invoke: void setFontWeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat13setFontWeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontWeight", args)
  }

  return
}

// verticalAlignment()
func (this *QTextCharFormat) Verticalalignment(args ...interface{}) () {
  // verticalAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat17verticalAlignmentEv
    // invoke: QTextCharFormat::VerticalAlignment verticalAlignment()
    C.C_ZNK15QTextCharFormat17verticalAlignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "verticalAlignment", args)
  }

  return
}

// setFontFamily(const class QString &)
func (this *QTextCharFormat) Setfontfamily(args ...interface{}) () {
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
    // invoke: void setFontFamily(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat13setFontFamilyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontFamily", args)
  }

  return
}

// setAnchorNames(const class QStringList &)
func (this *QTextCharFormat) Setanchornames(args ...interface{}) () {
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
    // invoke: void setAnchorNames(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat14setAnchorNamesERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorNames", args)
  }

  return
}

// setTableCellRowSpan(int)
func (this *QTextCharFormat) Settablecellrowspan(args ...interface{}) () {
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
    // invoke: void setTableCellRowSpan(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat19setTableCellRowSpanEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setTableCellRowSpan", args)
  }

  return
}

// setFontKerning(_Bool)
func (this *QTextCharFormat) Setfontkerning(args ...interface{}) () {
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
    // invoke: void setFontKerning(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat14setFontKerningEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontKerning", args)
  }

  return
}

// setFontLetterSpacing(qreal)
func (this *QTextCharFormat) Setfontletterspacing(args ...interface{}) () {
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
    // invoke: void setFontLetterSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat20setFontLetterSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontLetterSpacing", args)
  }

  return
}

// fontLetterSpacingType()
func (this *QTextCharFormat) Fontletterspacingtype(args ...interface{}) () {
  // fontLetterSpacingType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat21fontLetterSpacingTypeEv
    // invoke: QFont::SpacingType fontLetterSpacingType()
    C.C_ZNK15QTextCharFormat21fontLetterSpacingTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontLetterSpacingType", args)
  }

  return
}

// fontUnderline()
func (this *QTextCharFormat) Fontunderline(args ...interface{}) (ret interface{}) {
  // fontUnderline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat13fontUnderlineEv
    // invoke: bool fontUnderline()
    var ret0 = C.C_ZNK15QTextCharFormat13fontUnderlineEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontUnderline", args)
  }

  return
}

// fontHintingPreference()
func (this *QTextCharFormat) Fonthintingpreference(args ...interface{}) () {
  // fontHintingPreference()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat21fontHintingPreferenceEv
    // invoke: QFont::HintingPreference fontHintingPreference()
    C.C_ZNK15QTextCharFormat21fontHintingPreferenceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontHintingPreference", args)
  }

  return
}

// setFontStrikeOut(_Bool)
func (this *QTextCharFormat) Setfontstrikeout(args ...interface{}) () {
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
    // invoke: void setFontStrikeOut(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat16setFontStrikeOutEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontStrikeOut", args)
  }

  return
}

// anchorName()
func (this *QTextCharFormat) Anchorname(args ...interface{}) (ret interface{}) {
  // anchorName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10anchorNameEv
    // invoke: QString anchorName()
    var ret0 = C.C_ZNK15QTextCharFormat10anchorNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorName", args)
  }

  return
}

// fontPointSize()
func (this *QTextCharFormat) Fontpointsize(args ...interface{}) (ret interface{}) {
  // fontPointSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat13fontPointSizeEv
    // invoke: qreal fontPointSize()
    var ret0 = C.C_ZNK15QTextCharFormat13fontPointSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontPointSize", args)
  }

  return
}

// fontFixedPitch()
func (this *QTextCharFormat) Fontfixedpitch(args ...interface{}) (ret interface{}) {
  // fontFixedPitch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat14fontFixedPitchEv
    // invoke: bool fontFixedPitch()
    var ret0 = C.C_ZNK15QTextCharFormat14fontFixedPitchEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontFixedPitch", args)
  }

  return
}

// isValid()
func (this *QTextCharFormat) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK15QTextCharFormat7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "isValid", args)
  }

  return
}

// toolTip()
func (this *QTextCharFormat) Tooltip(args ...interface{}) (ret interface{}) {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat7toolTipEv
    // invoke: QString toolTip()
    var ret0 = C.C_ZNK15QTextCharFormat7toolTipEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "toolTip", args)
  }

  return
}

// setFontPointSize(qreal)
func (this *QTextCharFormat) Setfontpointsize(args ...interface{}) () {
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
    // invoke: void setFontPointSize(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat16setFontPointSizeEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontPointSize", args)
  }

  return
}

// setAnchor(_Bool)
func (this *QTextCharFormat) Setanchor(args ...interface{}) () {
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
    // invoke: void setAnchor(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat9setAnchorEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchor", args)
  }

  return
}

// setAnchorName(const class QString &)
func (this *QTextCharFormat) Setanchorname(args ...interface{}) () {
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
    // invoke: void setAnchorName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat13setAnchorNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setAnchorName", args)
  }

  return
}

// fontLetterSpacing()
func (this *QTextCharFormat) Fontletterspacing(args ...interface{}) (ret interface{}) {
  // fontLetterSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat17fontLetterSpacingEv
    // invoke: qreal fontLetterSpacing()
    var ret0 = C.C_ZNK15QTextCharFormat17fontLetterSpacingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontLetterSpacing", args)
  }

  return
}

// isAnchor()
func (this *QTextCharFormat) Isanchor(args ...interface{}) (ret interface{}) {
  // isAnchor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat8isAnchorEv
    // invoke: bool isAnchor()
    var ret0 = C.C_ZNK15QTextCharFormat8isAnchorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "isAnchor", args)
  }

  return
}

// setFontItalic(_Bool)
func (this *QTextCharFormat) Setfontitalic(args ...interface{}) () {
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
    // invoke: void setFontItalic(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat13setFontItalicEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontItalic", args)
  }

  return
}

// QTextCharFormat()
func NewQTextCharFormat(args ...interface{}) *QTextCharFormat {
  // QTextCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormatC1Ev
    // invoke: void QTextCharFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QTextCharFormatC2Ev()
    return &QTextCharFormat{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextCharFormat", "QTextCharFormat", args)
  }

  return nil // QTextCharFormat{qclsinst:qthis}
}

// underlineColor()
func (this *QTextCharFormat) Underlinecolor(args ...interface{}) (ret interface{}) {
  // underlineColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat14underlineColorEv
    // invoke: QColor underlineColor()
    var ret0 = C.C_ZNK15QTextCharFormat14underlineColorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "underlineColor", args)
  }

  return
}

// setUnderlineColor(const class QColor &)
func (this *QTextCharFormat) Setunderlinecolor(args ...interface{}) () {
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
    // invoke: void setUnderlineColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat17setUnderlineColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setUnderlineColor", args)
  }

  return
}

// fontWeight()
func (this *QTextCharFormat) Fontweight(args ...interface{}) (ret interface{}) {
  // fontWeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10fontWeightEv
    // invoke: int fontWeight()
    var ret0 = C.C_ZNK15QTextCharFormat10fontWeightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontWeight", args)
  }

  return
}

// setToolTip(const class QString &)
func (this *QTextCharFormat) Settooltip(args ...interface{}) () {
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
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat10setToolTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setToolTip", args)
  }

  return
}

// tableCellColumnSpan()
func (this *QTextCharFormat) Tablecellcolumnspan(args ...interface{}) (ret interface{}) {
  // tableCellColumnSpan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat19tableCellColumnSpanEv
    // invoke: int tableCellColumnSpan()
    var ret0 = C.C_ZNK15QTextCharFormat19tableCellColumnSpanEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "tableCellColumnSpan", args)
  }

  return
}

// anchorHref()
func (this *QTextCharFormat) Anchorhref(args ...interface{}) (ret interface{}) {
  // anchorHref()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat10anchorHrefEv
    // invoke: QString anchorHref()
    var ret0 = C.C_ZNK15QTextCharFormat10anchorHrefEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "anchorHref", args)
  }

  return
}

// fontCapitalization()
func (this *QTextCharFormat) Fontcapitalization(args ...interface{}) () {
  // fontCapitalization()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat18fontCapitalizationEv
    // invoke: QFont::Capitalization fontCapitalization()
    C.C_ZNK15QTextCharFormat18fontCapitalizationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontCapitalization", args)
  }

  return
}

// textOutline()
func (this *QTextCharFormat) Textoutline(args ...interface{}) (ret interface{}) {
  // textOutline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11textOutlineEv
    // invoke: QPen textOutline()
    var ret0 = C.C_ZNK15QTextCharFormat11textOutlineEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPen{}) // "QPen"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "textOutline", args)
  }

  return
}

// tableCellRowSpan()
func (this *QTextCharFormat) Tablecellrowspan(args ...interface{}) (ret interface{}) {
  // tableCellRowSpan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat16tableCellRowSpanEv
    // invoke: int tableCellRowSpan()
    var ret0 = C.C_ZNK15QTextCharFormat16tableCellRowSpanEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "tableCellRowSpan", args)
  }

  return
}

// setFontFixedPitch(_Bool)
func (this *QTextCharFormat) Setfontfixedpitch(args ...interface{}) () {
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
    // invoke: void setFontFixedPitch(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat17setFontFixedPitchEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFontFixedPitch", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QTextCharFormat) Setfont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextCharFormat7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextCharFormat7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCharFormat", "setFont", args)
  }

  return
}

// fontStretch()
func (this *QTextCharFormat) Fontstretch(args ...interface{}) (ret interface{}) {
  // fontStretch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat11fontStretchEv
    // invoke: int fontStretch()
    var ret0 = C.C_ZNK15QTextCharFormat11fontStretchEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStretch", args)
  }

  return
}

// fontStrikeOut()
func (this *QTextCharFormat) Fontstrikeout(args ...interface{}) (ret interface{}) {
  // fontStrikeOut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextCharFormat13fontStrikeOutEv
    // invoke: bool fontStrikeOut()
    var ret0 = C.C_ZNK15QTextCharFormat13fontStrikeOutEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextCharFormat", "fontStrikeOut", args)
  }

  return
}

// clearColumnWidthConstraints()
func (this *QTextTableFormat) Clearcolumnwidthconstraints(args ...interface{}) () {
  // clearColumnWidthConstraints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormat27clearColumnWidthConstraintsEv
    // invoke: void clearColumnWidthConstraints()
    C.C_ZN16QTextTableFormat27clearColumnWidthConstraintsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "clearColumnWidthConstraints", args)
  }

  return
}

// columnWidthConstraints()
func (this *QTextTableFormat) Columnwidthconstraints(args ...interface{}) () {
  // columnWidthConstraints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat22columnWidthConstraintsEv
    // invoke: QVector<QTextLength> columnWidthConstraints()
    C.C_ZNK16QTextTableFormat22columnWidthConstraintsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "columnWidthConstraints", args)
  }

  return
}

// QTextTableFormat()
func NewQTextTableFormat(args ...interface{}) *QTextTableFormat {
  // QTextTableFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextTableFormatC1Ev
    // invoke: void QTextTableFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTextTableFormatC2Ev()
    return &QTextTableFormat{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextTableFormat", "QTextTableFormat", args)
  }

  return nil // QTextTableFormat{qclsinst:qthis}
}

// setHeaderRowCount(int)
func (this *QTextTableFormat) Setheaderrowcount(args ...interface{}) () {
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
    // invoke: void setHeaderRowCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextTableFormat17setHeaderRowCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setHeaderRowCount", args)
  }

  return
}

// headerRowCount()
func (this *QTextTableFormat) Headerrowcount(args ...interface{}) (ret interface{}) {
  // headerRowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat14headerRowCountEv
    // invoke: int headerRowCount()
    var ret0 = C.C_ZNK16QTextTableFormat14headerRowCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextTableFormat", "headerRowCount", args)
  }

  return
}

// setColumns(int)
func (this *QTextTableFormat) Setcolumns(args ...interface{}) () {
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
    // invoke: void setColumns(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextTableFormat10setColumnsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setColumns", args)
  }

  return
}

// alignment()
func (this *QTextTableFormat) Alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK16QTextTableFormat9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "alignment", args)
  }

  return
}

// isValid()
func (this *QTextTableFormat) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK16QTextTableFormat7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextTableFormat", "isValid", args)
  }

  return
}

// cellPadding()
func (this *QTextTableFormat) Cellpadding(args ...interface{}) (ret interface{}) {
  // cellPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat11cellPaddingEv
    // invoke: qreal cellPadding()
    var ret0 = C.C_ZNK16QTextTableFormat11cellPaddingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextTableFormat", "cellPadding", args)
  }

  return
}

// setCellSpacing(qreal)
func (this *QTextTableFormat) Setcellspacing(args ...interface{}) () {
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
    // invoke: void setCellSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextTableFormat14setCellSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setCellSpacing", args)
  }

  return
}

// cellSpacing()
func (this *QTextTableFormat) Cellspacing(args ...interface{}) (ret interface{}) {
  // cellSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat11cellSpacingEv
    // invoke: qreal cellSpacing()
    var ret0 = C.C_ZNK16QTextTableFormat11cellSpacingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextTableFormat", "cellSpacing", args)
  }

  return
}

// setCellPadding(qreal)
func (this *QTextTableFormat) Setcellpadding(args ...interface{}) () {
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
    // invoke: void setCellPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextTableFormat14setCellPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableFormat", "setCellPadding", args)
  }

  return
}

// columns()
func (this *QTextTableFormat) Columns(args ...interface{}) (ret interface{}) {
  // columns()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextTableFormat7columnsEv
    // invoke: int columns()
    var ret0 = C.C_ZNK16QTextTableFormat7columnsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextTableFormat", "columns", args)
  }

  return
}

// setLeftPadding(qreal)
func (this *QTextTableCellFormat) Setleftpadding(args ...interface{}) () {
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
    // invoke: void setLeftPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN20QTextTableCellFormat14setLeftPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setLeftPadding", args)
  }

  return
}

// leftPadding()
func (this *QTextTableCellFormat) Leftpadding(args ...interface{}) (ret interface{}) {
  // leftPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat11leftPaddingEv
    // invoke: qreal leftPadding()
    var ret0 = C.C_ZNK20QTextTableCellFormat11leftPaddingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "leftPadding", args)
  }

  return
}

// setBottomPadding(qreal)
func (this *QTextTableCellFormat) Setbottompadding(args ...interface{}) () {
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
    // invoke: void setBottomPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN20QTextTableCellFormat16setBottomPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setBottomPadding", args)
  }

  return
}

// QTextTableCellFormat()
func NewQTextTableCellFormat(args ...interface{}) *QTextTableCellFormat {
  // QTextTableCellFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QTextTableCellFormatC1Ev
    // invoke: void QTextTableCellFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QTextTableCellFormatC2Ev()
    return &QTextTableCellFormat{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "QTextTableCellFormat", args)
  }

  return nil // QTextTableCellFormat{qclsinst:qthis}
}

// setTopPadding(qreal)
func (this *QTextTableCellFormat) Settoppadding(args ...interface{}) () {
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
    // invoke: void setTopPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN20QTextTableCellFormat13setTopPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setTopPadding", args)
  }

  return
}

// rightPadding()
func (this *QTextTableCellFormat) Rightpadding(args ...interface{}) (ret interface{}) {
  // rightPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat12rightPaddingEv
    // invoke: qreal rightPadding()
    var ret0 = C.C_ZNK20QTextTableCellFormat12rightPaddingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "rightPadding", args)
  }

  return
}

// setRightPadding(qreal)
func (this *QTextTableCellFormat) Setrightpadding(args ...interface{}) () {
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
    // invoke: void setRightPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN20QTextTableCellFormat15setRightPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setRightPadding", args)
  }

  return
}

// isValid()
func (this *QTextTableCellFormat) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK20QTextTableCellFormat7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "isValid", args)
  }

  return
}

// topPadding()
func (this *QTextTableCellFormat) Toppadding(args ...interface{}) (ret interface{}) {
  // topPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat10topPaddingEv
    // invoke: qreal topPadding()
    var ret0 = C.C_ZNK20QTextTableCellFormat10topPaddingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "topPadding", args)
  }

  return
}

// setPadding(qreal)
func (this *QTextTableCellFormat) Setpadding(args ...interface{}) () {
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
    // invoke: void setPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN20QTextTableCellFormat10setPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "setPadding", args)
  }

  return
}

// bottomPadding()
func (this *QTextTableCellFormat) Bottompadding(args ...interface{}) (ret interface{}) {
  // bottomPadding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QTextTableCellFormat13bottomPaddingEv
    // invoke: qreal bottomPadding()
    var ret0 = C.C_ZNK20QTextTableCellFormat13bottomPaddingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextTableCellFormat", "bottomPadding", args)
  }

  return
}

// style()
func (this *QTextListFormat) Style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat5styleEv
    // invoke: QTextListFormat::Style style()
    C.C_ZNK15QTextListFormat5styleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextListFormat", "style", args)
  }

  return
}

// numberPrefix()
func (this *QTextListFormat) Numberprefix(args ...interface{}) (ret interface{}) {
  // numberPrefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat12numberPrefixEv
    // invoke: QString numberPrefix()
    var ret0 = C.C_ZNK15QTextListFormat12numberPrefixEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextListFormat", "numberPrefix", args)
  }

  return
}

// indent()
func (this *QTextListFormat) Indent(args ...interface{}) (ret interface{}) {
  // indent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat6indentEv
    // invoke: int indent()
    var ret0 = C.C_ZNK15QTextListFormat6indentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextListFormat", "indent", args)
  }

  return
}

// setNumberPrefix(const class QString &)
func (this *QTextListFormat) Setnumberprefix(args ...interface{}) () {
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
    // invoke: void setNumberPrefix(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextListFormat15setNumberPrefixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextListFormat", "setNumberPrefix", args)
  }

  return
}

// isValid()
func (this *QTextListFormat) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK15QTextListFormat7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextListFormat", "isValid", args)
  }

  return
}

// QTextListFormat()
func NewQTextListFormat(args ...interface{}) *QTextListFormat {
  // QTextListFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTextListFormatC1Ev
    // invoke: void QTextListFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QTextListFormatC2Ev()
    return &QTextListFormat{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextListFormat", "QTextListFormat", args)
  }

  return nil // QTextListFormat{qclsinst:qthis}
}

// setIndent(int)
func (this *QTextListFormat) Setindent(args ...interface{}) () {
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
    // invoke: void setIndent(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextListFormat9setIndentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextListFormat", "setIndent", args)
  }

  return
}

// numberSuffix()
func (this *QTextListFormat) Numbersuffix(args ...interface{}) (ret interface{}) {
  // numberSuffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTextListFormat12numberSuffixEv
    // invoke: QString numberSuffix()
    var ret0 = C.C_ZNK15QTextListFormat12numberSuffixEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextListFormat", "numberSuffix", args)
  }

  return
}

// setNumberSuffix(const class QString &)
func (this *QTextListFormat) Setnumbersuffix(args ...interface{}) () {
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
    // invoke: void setNumberSuffix(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTextListFormat15setNumberSuffixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextListFormat", "setNumberSuffix", args)
  }

  return
}

// setBorderBrush(const class QBrush &)
func (this *QTextFrameFormat) Setborderbrush(args ...interface{}) () {
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
    // invoke: void setBorderBrush(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat14setBorderBrushERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBorderBrush", args)
  }

  return
}

// height()
func (this *QTextFrameFormat) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat6heightEv
    // invoke: QTextLength height()
    var ret0 = C.C_ZNK16QTextFrameFormat6heightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextLength{}) // "QTextLength"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "height", args)
  }

  return
}

// setWidth(qreal)
func (this *QTextFrameFormat) Setwidth(args ...interface{}) () {
  // setWidth(qreal)
  // setWidth(const class QTextLength &)
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
    // invoke: _ZN16QTextFrameFormat8setWidthEd
    // invoke: void setWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat8setWidthEd(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN16QTextFrameFormat8setWidthERK11QTextLength
    // invoke: void setWidth(const class QTextLength &)
    var arg0 = args[0].(QTextLength).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat8setWidthERK11QTextLength(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setWidth", args)
  }

  return
}

// topMargin()
func (this *QTextFrameFormat) Topmargin(args ...interface{}) (ret interface{}) {
  // topMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat9topMarginEv
    // invoke: qreal topMargin()
    var ret0 = C.C_ZNK16QTextFrameFormat9topMarginEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "topMargin", args)
  }

  return
}

// border()
func (this *QTextFrameFormat) Border(args ...interface{}) (ret interface{}) {
  // border()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat6borderEv
    // invoke: qreal border()
    var ret0 = C.C_ZNK16QTextFrameFormat6borderEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "border", args)
  }

  return
}

// width()
func (this *QTextFrameFormat) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat5widthEv
    // invoke: QTextLength width()
    var ret0 = C.C_ZNK16QTextFrameFormat5widthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextLength{}) // "QTextLength"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "width", args)
  }

  return
}

// setLeftMargin(qreal)
func (this *QTextFrameFormat) Setleftmargin(args ...interface{}) () {
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
    // invoke: void setLeftMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat13setLeftMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setLeftMargin", args)
  }

  return
}

// leftMargin()
func (this *QTextFrameFormat) Leftmargin(args ...interface{}) (ret interface{}) {
  // leftMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat10leftMarginEv
    // invoke: qreal leftMargin()
    var ret0 = C.C_ZNK16QTextFrameFormat10leftMarginEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "leftMargin", args)
  }

  return
}

// setPadding(qreal)
func (this *QTextFrameFormat) Setpadding(args ...interface{}) () {
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
    // invoke: void setPadding(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat10setPaddingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setPadding", args)
  }

  return
}

// setTopMargin(qreal)
func (this *QTextFrameFormat) Settopmargin(args ...interface{}) () {
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
    // invoke: void setTopMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat12setTopMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setTopMargin", args)
  }

  return
}

// rightMargin()
func (this *QTextFrameFormat) Rightmargin(args ...interface{}) (ret interface{}) {
  // rightMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat11rightMarginEv
    // invoke: qreal rightMargin()
    var ret0 = C.C_ZNK16QTextFrameFormat11rightMarginEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "rightMargin", args)
  }

  return
}

// isValid()
func (this *QTextFrameFormat) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK16QTextFrameFormat7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "isValid", args)
  }

  return
}

// pageBreakPolicy()
func (this *QTextFrameFormat) Pagebreakpolicy(args ...interface{}) () {
  // pageBreakPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat15pageBreakPolicyEv
    // invoke: PageBreakFlags pageBreakPolicy()
    C.C_ZNK16QTextFrameFormat15pageBreakPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "pageBreakPolicy", args)
  }

  return
}

// padding()
func (this *QTextFrameFormat) Padding(args ...interface{}) (ret interface{}) {
  // padding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat7paddingEv
    // invoke: qreal padding()
    var ret0 = C.C_ZNK16QTextFrameFormat7paddingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "padding", args)
  }

  return
}

// setMargin(qreal)
func (this *QTextFrameFormat) Setmargin(args ...interface{}) () {
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
    // invoke: void setMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat9setMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setMargin", args)
  }

  return
}

// bottomMargin()
func (this *QTextFrameFormat) Bottommargin(args ...interface{}) (ret interface{}) {
  // bottomMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat12bottomMarginEv
    // invoke: qreal bottomMargin()
    var ret0 = C.C_ZNK16QTextFrameFormat12bottomMarginEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "bottomMargin", args)
  }

  return
}

// QTextFrameFormat()
func NewQTextFrameFormat(args ...interface{}) *QTextFrameFormat {
  // QTextFrameFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTextFrameFormatC1Ev
    // invoke: void QTextFrameFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTextFrameFormatC2Ev()
    return &QTextFrameFormat{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "QTextFrameFormat", args)
  }

  return nil // QTextFrameFormat{qclsinst:qthis}
}

// setHeight(qreal)
func (this *QTextFrameFormat) Setheight(args ...interface{}) () {
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
    // invoke: void setHeight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat9setHeightEd(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN16QTextFrameFormat9setHeightERK11QTextLength
    // invoke: void setHeight(const class QTextLength &)
    var arg0 = args[0].(QTextLength).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat9setHeightERK11QTextLength(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setHeight", args)
  }

  return
}

// borderBrush()
func (this *QTextFrameFormat) Borderbrush(args ...interface{}) (ret interface{}) {
  // borderBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat11borderBrushEv
    // invoke: QBrush borderBrush()
    var ret0 = C.C_ZNK16QTextFrameFormat11borderBrushEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "borderBrush", args)
  }

  return
}

// setRightMargin(qreal)
func (this *QTextFrameFormat) Setrightmargin(args ...interface{}) () {
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
    // invoke: void setRightMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat14setRightMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setRightMargin", args)
  }

  return
}

// setBorder(qreal)
func (this *QTextFrameFormat) Setborder(args ...interface{}) () {
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
    // invoke: void setBorder(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat9setBorderEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBorder", args)
  }

  return
}

// setBottomMargin(qreal)
func (this *QTextFrameFormat) Setbottommargin(args ...interface{}) () {
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
    // invoke: void setBottomMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTextFrameFormat15setBottomMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "setBottomMargin", args)
  }

  return
}

// position()
func (this *QTextFrameFormat) Position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat8positionEv
    // invoke: QTextFrameFormat::Position position()
    C.C_ZNK16QTextFrameFormat8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "position", args)
  }

  return
}

// borderStyle()
func (this *QTextFrameFormat) Borderstyle(args ...interface{}) () {
  // borderStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat11borderStyleEv
    // invoke: QTextFrameFormat::BorderStyle borderStyle()
    C.C_ZNK16QTextFrameFormat11borderStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "borderStyle", args)
  }

  return
}

// margin()
func (this *QTextFrameFormat) Margin(args ...interface{}) (ret interface{}) {
  // margin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTextFrameFormat6marginEv
    // invoke: qreal margin()
    var ret0 = C.C_ZNK16QTextFrameFormat6marginEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextFrameFormat", "margin", args)
  }

  return
}

// <= body block end

