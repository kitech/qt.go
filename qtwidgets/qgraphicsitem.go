// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*
 */
// size 16
type QGraphicsItem struct {
	*qtrt.CObject
}
type QGraphicsItem_ITF interface {
	QGraphicsItem_PTR() *QGraphicsItem
}

func (ptr *QGraphicsItem) QGraphicsItem_PTR() *QGraphicsItem { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
func QGraphicsItemFromptr(cthis unsafe.Pointer) *QGraphicsItem {
	return &QGraphicsItem{&qtrt.CObject{cthis}}
}
func (*QGraphicsItem) Fromptr(cthis unsafe.Pointer) *QGraphicsItem {
	return QGraphicsItemFromptr(cthis)
}

func DeleteQGraphicsItem(this *QGraphicsItem) {
	rv, err := qtrt.Qtcc1(0, "_ZN13QGraphicsItemD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QGraphicsItem__GraphicsItemFlag = int

//
const QGraphicsItem__ItemIsMovable QGraphicsItem__GraphicsItemFlag = 1

//
const QGraphicsItem__ItemIsSelectable QGraphicsItem__GraphicsItemFlag = 2

//
const QGraphicsItem__ItemIsFocusable QGraphicsItem__GraphicsItemFlag = 4

//
const QGraphicsItem__ItemClipsToShape QGraphicsItem__GraphicsItemFlag = 8

//
const QGraphicsItem__ItemClipsChildrenToShape QGraphicsItem__GraphicsItemFlag = 16

//
const QGraphicsItem__ItemIgnoresTransformations QGraphicsItem__GraphicsItemFlag = 32

//
const QGraphicsItem__ItemIgnoresParentOpacity QGraphicsItem__GraphicsItemFlag = 64

//
const QGraphicsItem__ItemDoesntPropagateOpacityToChildren QGraphicsItem__GraphicsItemFlag = 128

//
const QGraphicsItem__ItemStacksBehindParent QGraphicsItem__GraphicsItemFlag = 256

//
const QGraphicsItem__ItemUsesExtendedStyleOption QGraphicsItem__GraphicsItemFlag = 512

//
const QGraphicsItem__ItemHasNoContents QGraphicsItem__GraphicsItemFlag = 1024

//
const QGraphicsItem__ItemSendsGeometryChanges QGraphicsItem__GraphicsItemFlag = 2048

//
const QGraphicsItem__ItemAcceptsInputMethod QGraphicsItem__GraphicsItemFlag = 4096

//
const QGraphicsItem__ItemNegativeZStacksBehindParent QGraphicsItem__GraphicsItemFlag = 8192

//
const QGraphicsItem__ItemIsPanel QGraphicsItem__GraphicsItemFlag = 16384

//
const QGraphicsItem__ItemIsFocusScope QGraphicsItem__GraphicsItemFlag = 32768

//
const QGraphicsItem__ItemSendsScenePositionChanges QGraphicsItem__GraphicsItemFlag = 65536

//
const QGraphicsItem__ItemStopsClickFocusPropagation QGraphicsItem__GraphicsItemFlag = 131072

//
const QGraphicsItem__ItemStopsFocusHandling QGraphicsItem__GraphicsItemFlag = 262144

//
const QGraphicsItem__ItemContainsChildrenInShape QGraphicsItem__GraphicsItemFlag = 524288

func (this *QGraphicsItem) GraphicsItemFlagItemName(val int) string {
	switch val {
	case QGraphicsItem__ItemIsMovable: // 1
		return "ItemIsMovable"
	case QGraphicsItem__ItemIsSelectable: // 2
		return "ItemIsSelectable"
	case QGraphicsItem__ItemIsFocusable: // 4
		return "ItemIsFocusable"
	case QGraphicsItem__ItemClipsToShape: // 8
		return "ItemClipsToShape"
	case QGraphicsItem__ItemClipsChildrenToShape: // 16
		return "ItemClipsChildrenToShape"
	case QGraphicsItem__ItemIgnoresTransformations: // 32
		return "ItemIgnoresTransformations"
	case QGraphicsItem__ItemIgnoresParentOpacity: // 64
		return "ItemIgnoresParentOpacity"
	case QGraphicsItem__ItemDoesntPropagateOpacityToChildren: // 128
		return "ItemDoesntPropagateOpacityToChildren"
	case QGraphicsItem__ItemStacksBehindParent: // 256
		return "ItemStacksBehindParent"
	case QGraphicsItem__ItemUsesExtendedStyleOption: // 512
		return "ItemUsesExtendedStyleOption"
	case QGraphicsItem__ItemHasNoContents: // 1024
		return "ItemHasNoContents"
	case QGraphicsItem__ItemSendsGeometryChanges: // 2048
		return "ItemSendsGeometryChanges"
	case QGraphicsItem__ItemAcceptsInputMethod: // 4096
		return "ItemAcceptsInputMethod"
	case QGraphicsItem__ItemNegativeZStacksBehindParent: // 8192
		return "ItemNegativeZStacksBehindParent"
	case QGraphicsItem__ItemIsPanel: // 16384
		return "ItemIsPanel"
	case QGraphicsItem__ItemIsFocusScope: // 32768
		return "ItemIsFocusScope"
	case QGraphicsItem__ItemSendsScenePositionChanges: // 65536
		return "ItemSendsScenePositionChanges"
	case QGraphicsItem__ItemStopsClickFocusPropagation: // 131072
		return "ItemStopsClickFocusPropagation"
	case QGraphicsItem__ItemStopsFocusHandling: // 262144
		return "ItemStopsFocusHandling"
	case QGraphicsItem__ItemContainsChildrenInShape: // 524288
		return "ItemContainsChildrenInShape"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_GraphicsItemFlagItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.GraphicsItemFlagItemName(val)
}

/*


 */
type QGraphicsItem__GraphicsItemChange = int

//
const QGraphicsItem__ItemPositionChange QGraphicsItem__GraphicsItemChange = 0

//
const QGraphicsItem__ItemMatrixChange QGraphicsItem__GraphicsItemChange = 1

//
const QGraphicsItem__ItemVisibleChange QGraphicsItem__GraphicsItemChange = 2

//
const QGraphicsItem__ItemEnabledChange QGraphicsItem__GraphicsItemChange = 3

//
const QGraphicsItem__ItemSelectedChange QGraphicsItem__GraphicsItemChange = 4

//
const QGraphicsItem__ItemParentChange QGraphicsItem__GraphicsItemChange = 5

//
const QGraphicsItem__ItemChildAddedChange QGraphicsItem__GraphicsItemChange = 6

//
const QGraphicsItem__ItemChildRemovedChange QGraphicsItem__GraphicsItemChange = 7

//
const QGraphicsItem__ItemTransformChange QGraphicsItem__GraphicsItemChange = 8

//
const QGraphicsItem__ItemPositionHasChanged QGraphicsItem__GraphicsItemChange = 9

//
const QGraphicsItem__ItemTransformHasChanged QGraphicsItem__GraphicsItemChange = 10

//
const QGraphicsItem__ItemSceneChange QGraphicsItem__GraphicsItemChange = 11

//
const QGraphicsItem__ItemVisibleHasChanged QGraphicsItem__GraphicsItemChange = 12

//
const QGraphicsItem__ItemEnabledHasChanged QGraphicsItem__GraphicsItemChange = 13

//
const QGraphicsItem__ItemSelectedHasChanged QGraphicsItem__GraphicsItemChange = 14

//
const QGraphicsItem__ItemParentHasChanged QGraphicsItem__GraphicsItemChange = 15

//
const QGraphicsItem__ItemSceneHasChanged QGraphicsItem__GraphicsItemChange = 16

//
const QGraphicsItem__ItemCursorChange QGraphicsItem__GraphicsItemChange = 17

//
const QGraphicsItem__ItemCursorHasChanged QGraphicsItem__GraphicsItemChange = 18

//
const QGraphicsItem__ItemToolTipChange QGraphicsItem__GraphicsItemChange = 19

//
const QGraphicsItem__ItemToolTipHasChanged QGraphicsItem__GraphicsItemChange = 20

//
const QGraphicsItem__ItemFlagsChange QGraphicsItem__GraphicsItemChange = 21

//
const QGraphicsItem__ItemFlagsHaveChanged QGraphicsItem__GraphicsItemChange = 22

//
const QGraphicsItem__ItemZValueChange QGraphicsItem__GraphicsItemChange = 23

//
const QGraphicsItem__ItemZValueHasChanged QGraphicsItem__GraphicsItemChange = 24

//
const QGraphicsItem__ItemOpacityChange QGraphicsItem__GraphicsItemChange = 25

//
const QGraphicsItem__ItemOpacityHasChanged QGraphicsItem__GraphicsItemChange = 26

//
const QGraphicsItem__ItemScenePositionHasChanged QGraphicsItem__GraphicsItemChange = 27

//
const QGraphicsItem__ItemRotationChange QGraphicsItem__GraphicsItemChange = 28

//
const QGraphicsItem__ItemRotationHasChanged QGraphicsItem__GraphicsItemChange = 29

//
const QGraphicsItem__ItemScaleChange QGraphicsItem__GraphicsItemChange = 30

//
const QGraphicsItem__ItemScaleHasChanged QGraphicsItem__GraphicsItemChange = 31

//
const QGraphicsItem__ItemTransformOriginPointChange QGraphicsItem__GraphicsItemChange = 32

//
const QGraphicsItem__ItemTransformOriginPointHasChanged QGraphicsItem__GraphicsItemChange = 33

func (this *QGraphicsItem) GraphicsItemChangeItemName(val int) string {
	switch val {
	case QGraphicsItem__ItemPositionChange: // 0
		return "ItemPositionChange"
	case QGraphicsItem__ItemMatrixChange: // 1
		return "ItemMatrixChange"
	case QGraphicsItem__ItemVisibleChange: // 2
		return "ItemVisibleChange"
	case QGraphicsItem__ItemEnabledChange: // 3
		return "ItemEnabledChange"
	case QGraphicsItem__ItemSelectedChange: // 4
		return "ItemSelectedChange"
	case QGraphicsItem__ItemParentChange: // 5
		return "ItemParentChange"
	case QGraphicsItem__ItemChildAddedChange: // 6
		return "ItemChildAddedChange"
	case QGraphicsItem__ItemChildRemovedChange: // 7
		return "ItemChildRemovedChange"
	case QGraphicsItem__ItemTransformChange: // 8
		return "ItemTransformChange"
	case QGraphicsItem__ItemPositionHasChanged: // 9
		return "ItemPositionHasChanged"
	case QGraphicsItem__ItemTransformHasChanged: // 10
		return "ItemTransformHasChanged"
	case QGraphicsItem__ItemSceneChange: // 11
		return "ItemSceneChange"
	case QGraphicsItem__ItemVisibleHasChanged: // 12
		return "ItemVisibleHasChanged"
	case QGraphicsItem__ItemEnabledHasChanged: // 13
		return "ItemEnabledHasChanged"
	case QGraphicsItem__ItemSelectedHasChanged: // 14
		return "ItemSelectedHasChanged"
	case QGraphicsItem__ItemParentHasChanged: // 15
		return "ItemParentHasChanged"
	case QGraphicsItem__ItemSceneHasChanged: // 16
		return "ItemSceneHasChanged"
	case QGraphicsItem__ItemCursorChange: // 17
		return "ItemCursorChange"
	case QGraphicsItem__ItemCursorHasChanged: // 18
		return "ItemCursorHasChanged"
	case QGraphicsItem__ItemToolTipChange: // 19
		return "ItemToolTipChange"
	case QGraphicsItem__ItemToolTipHasChanged: // 20
		return "ItemToolTipHasChanged"
	case QGraphicsItem__ItemFlagsChange: // 21
		return "ItemFlagsChange"
	case QGraphicsItem__ItemFlagsHaveChanged: // 22
		return "ItemFlagsHaveChanged"
	case QGraphicsItem__ItemZValueChange: // 23
		return "ItemZValueChange"
	case QGraphicsItem__ItemZValueHasChanged: // 24
		return "ItemZValueHasChanged"
	case QGraphicsItem__ItemOpacityChange: // 25
		return "ItemOpacityChange"
	case QGraphicsItem__ItemOpacityHasChanged: // 26
		return "ItemOpacityHasChanged"
	case QGraphicsItem__ItemScenePositionHasChanged: // 27
		return "ItemScenePositionHasChanged"
	case QGraphicsItem__ItemRotationChange: // 28
		return "ItemRotationChange"
	case QGraphicsItem__ItemRotationHasChanged: // 29
		return "ItemRotationHasChanged"
	case QGraphicsItem__ItemScaleChange: // 30
		return "ItemScaleChange"
	case QGraphicsItem__ItemScaleHasChanged: // 31
		return "ItemScaleHasChanged"
	case QGraphicsItem__ItemTransformOriginPointChange: // 32
		return "ItemTransformOriginPointChange"
	case QGraphicsItem__ItemTransformOriginPointHasChanged: // 33
		return "ItemTransformOriginPointHasChanged"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_GraphicsItemChangeItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.GraphicsItemChangeItemName(val)
}

/*


 */
type QGraphicsItem__CacheMode = int

//
const QGraphicsItem__NoCache QGraphicsItem__CacheMode = 0

//
const QGraphicsItem__ItemCoordinateCache QGraphicsItem__CacheMode = 1

//
const QGraphicsItem__DeviceCoordinateCache QGraphicsItem__CacheMode = 2

func (this *QGraphicsItem) CacheModeItemName(val int) string {
	switch val {
	case QGraphicsItem__NoCache: // 0
		return "NoCache"
	case QGraphicsItem__ItemCoordinateCache: // 1
		return "ItemCoordinateCache"
	case QGraphicsItem__DeviceCoordinateCache: // 2
		return "DeviceCoordinateCache"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_CacheModeItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.CacheModeItemName(val)
}

/*


 */
type QGraphicsItem__PanelModality = int

//
const QGraphicsItem__NonModal QGraphicsItem__PanelModality = 0

//
const QGraphicsItem__PanelModal QGraphicsItem__PanelModality = 1

//
const QGraphicsItem__SceneModal QGraphicsItem__PanelModality = 2

func (this *QGraphicsItem) PanelModalityItemName(val int) string {
	switch val {
	case QGraphicsItem__NonModal: // 0
		return "NonModal"
	case QGraphicsItem__PanelModal: // 1
		return "PanelModal"
	case QGraphicsItem__SceneModal: // 2
		return "SceneModal"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_PanelModalityItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.PanelModalityItemName(val)
}

/*


 */
type QGraphicsItem__ = int

//
const QGraphicsItem__Type QGraphicsItem__ = 1

//
const QGraphicsItem__UserType QGraphicsItem__ = 65536

func (this *QGraphicsItem) ItemName(val int) string {
	switch val {
	case QGraphicsItem__Type: // 1
		return "Type"
	case QGraphicsItem__UserType: // 65536
		return "UserType"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_ItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.ItemName(val)
}

/*


 */
type QGraphicsItem__Extension = int

//
const QGraphicsItem__UserExtension QGraphicsItem__Extension = -2147483648

func (this *QGraphicsItem) ExtensionItemName(val int) string {
	switch val {
	case QGraphicsItem__UserExtension: // -2147483648
		return "UserExtension"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_ExtensionItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.ExtensionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10127() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
