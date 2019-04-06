package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QGradient struct {
	*qtrt.CObject
}
type QGradient_ITF interface {
	QGradient_PTR() *QGradient
}

func (ptr *QGradient) QGradient_PTR() *QGradient { return ptr }

func (this *QGradient) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGradient) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGradientFromPointer(cthis unsafe.Pointer) *QGradient {
	return &QGradient{&qtrt.CObject{cthis}}
}
func (*QGradient) NewFromPointer(cthis unsafe.Pointer) *QGradient {
	return NewQGradientFromPointer(cthis)
}

// /usr/include/qt/QtGui/qbrush.h:379
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGradient()

/*

 */
func (*QGradient) NewForInherit() *QGradient {
	return NewQGradient()
}
func NewQGradient() *QGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradientC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:380
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGradient(QGradient::Preset)

/*

 */
func (*QGradient) NewForInherit1(arg0 int) *QGradient {
	return NewQGradient1(arg0)
}
func NewQGradient1(arg0 int) *QGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradientC2ENS_6PresetE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:382
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QGradient::Type type() const

/*

 */
func (this *QGradient) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradient4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:384
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSpread(QGradient::Spread)

/*

 */
func (this *QGradient) SetSpread(spread int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradient9setSpreadENS_6SpreadE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spread)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:385
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QGradient::Spread spread() const

/*

 */
func (this *QGradient) Spread() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradient6spreadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:387
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColorAt(qreal, const QColor &)

/*

 */
func (this *QGradient) SetColorAt(pos float64, color QColor_ITF) {
	var convArg1 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg1 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradient10setColorAtEdRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:392
// index:0
// Public Visibility=Default Availability=Available
// [4] QGradient::CoordinateMode coordinateMode() const

/*

 */
func (this *QGradient) CoordinateMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradient14coordinateModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:393
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCoordinateMode(QGradient::CoordinateMode)

/*

 */
func (this *QGradient) SetCoordinateMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradient17setCoordinateModeENS_14CoordinateModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:395
// index:0
// Public Visibility=Default Availability=Available
// [4] QGradient::InterpolationMode interpolationMode() const

/*

 */
func (this *QGradient) InterpolationMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradient17interpolationModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:396
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInterpolationMode(QGradient::InterpolationMode)

/*

 */
func (this *QGradient) SetInterpolationMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradient20setInterpolationModeENS_17InterpolationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:398
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGradient &) const

/*

 */
func (this *QGradient) Operator_equal_equal(gradient QGradient_ITF) bool {
	var convArg0 unsafe.Pointer
	if gradient != nil && gradient.QGradient_PTR() != nil {
		convArg0 = gradient.QGradient_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradienteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qbrush.h:399
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QGradient &) const

/*

 */
func (this *QGradient) Operator_not_equal(other QGradient_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGradient_PTR() != nil {
		convArg0 = other.QGradient_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradientneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQGradient(this *QGradient) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradientD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QGradient__Type = int

//
const QGradient__LinearGradient QGradient__Type = 0

//
const QGradient__RadialGradient QGradient__Type = 1

//
const QGradient__ConicalGradient QGradient__Type = 2

//
const QGradient__NoGradient QGradient__Type = 3

func (this *QGradient) TypeItemName(val int) string {
	switch val {
	case QGradient__LinearGradient: // 0
		return "LinearGradient"
	case QGradient__RadialGradient: // 1
		return "RadialGradient"
	case QGradient__ConicalGradient: // 2
		return "ConicalGradient"
	case QGradient__NoGradient: // 3
		return "NoGradient"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGradient_TypeItemName(val int) string {
	var nilthis *QGradient
	return nilthis.TypeItemName(val)
}

/*


 */
type QGradient__Spread = int

//
const QGradient__PadSpread QGradient__Spread = 0

//
const QGradient__ReflectSpread QGradient__Spread = 1

//
const QGradient__RepeatSpread QGradient__Spread = 2

func (this *QGradient) SpreadItemName(val int) string {
	switch val {
	case QGradient__PadSpread: // 0
		return "PadSpread"
	case QGradient__ReflectSpread: // 1
		return "ReflectSpread"
	case QGradient__RepeatSpread: // 2
		return "RepeatSpread"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGradient_SpreadItemName(val int) string {
	var nilthis *QGradient
	return nilthis.SpreadItemName(val)
}

/*


 */
type QGradient__CoordinateMode = int

//
const QGradient__LogicalMode QGradient__CoordinateMode = 0

//
const QGradient__StretchToDeviceMode QGradient__CoordinateMode = 1

//
const QGradient__ObjectBoundingMode QGradient__CoordinateMode = 2

//
const QGradient__ObjectMode QGradient__CoordinateMode = 3

func (this *QGradient) CoordinateModeItemName(val int) string {
	switch val {
	case QGradient__LogicalMode: // 0
		return "LogicalMode"
	case QGradient__StretchToDeviceMode: // 1
		return "StretchToDeviceMode"
	case QGradient__ObjectBoundingMode: // 2
		return "ObjectBoundingMode"
	case QGradient__ObjectMode: // 3
		return "ObjectMode"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGradient_CoordinateModeItemName(val int) string {
	var nilthis *QGradient
	return nilthis.CoordinateModeItemName(val)
}

/*


 */
type QGradient__InterpolationMode = int

//
const QGradient__ColorInterpolation QGradient__InterpolationMode = 0

//
const QGradient__ComponentInterpolation QGradient__InterpolationMode = 1

func (this *QGradient) InterpolationModeItemName(val int) string {
	switch val {
	case QGradient__ColorInterpolation: // 0
		return "ColorInterpolation"
	case QGradient__ComponentInterpolation: // 1
		return "ComponentInterpolation"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGradient_InterpolationModeItemName(val int) string {
	var nilthis *QGradient
	return nilthis.InterpolationModeItemName(val)
}

/*


 */
type QGradient__Preset = int

//
const QGradient__WarmFlame QGradient__Preset = 1

//
const QGradient__NightFade QGradient__Preset = 2

//
const QGradient__SpringWarmth QGradient__Preset = 3

//
const QGradient__JuicyPeach QGradient__Preset = 4

//
const QGradient__YoungPassion QGradient__Preset = 5

//
const QGradient__LadyLips QGradient__Preset = 6

//
const QGradient__SunnyMorning QGradient__Preset = 7

//
const QGradient__RainyAshville QGradient__Preset = 8

//
const QGradient__FrozenDreams QGradient__Preset = 9

//
const QGradient__WinterNeva QGradient__Preset = 10

//
const QGradient__DustyGrass QGradient__Preset = 11

//
const QGradient__TemptingAzure QGradient__Preset = 12

//
const QGradient__HeavyRain QGradient__Preset = 13

//
const QGradient__AmyCrisp QGradient__Preset = 14

//
const QGradient__MeanFruit QGradient__Preset = 15

//
const QGradient__DeepBlue QGradient__Preset = 16

//
const QGradient__RipeMalinka QGradient__Preset = 17

//
const QGradient__CloudyKnoxville QGradient__Preset = 18

//
const QGradient__MalibuBeach QGradient__Preset = 19

//
const QGradient__NewLife QGradient__Preset = 20

//
const QGradient__TrueSunset QGradient__Preset = 21

//
const QGradient__MorpheusDen QGradient__Preset = 22

//
const QGradient__RareWind QGradient__Preset = 23

//
const QGradient__NearMoon QGradient__Preset = 24

//
const QGradient__WildApple QGradient__Preset = 25

//
const QGradient__SaintPetersburg QGradient__Preset = 26

//
const QGradient__PlumPlate QGradient__Preset = 28

//
const QGradient__EverlastingSky QGradient__Preset = 29

//
const QGradient__HappyFisher QGradient__Preset = 30

//
const QGradient__Blessing QGradient__Preset = 31

//
const QGradient__SharpeyeEagle QGradient__Preset = 32

//
const QGradient__LadogaBottom QGradient__Preset = 33

//
const QGradient__LemonGate QGradient__Preset = 34

//
const QGradient__ItmeoBranding QGradient__Preset = 35

//
const QGradient__ZeusMiracle QGradient__Preset = 36

//
const QGradient__OldHat QGradient__Preset = 37

//
const QGradient__StarWine QGradient__Preset = 38

//
const QGradient__HappyAcid QGradient__Preset = 41

//
const QGradient__AwesomePine QGradient__Preset = 42

//
const QGradient__NewYork QGradient__Preset = 43

//
const QGradient__ShyRainbow QGradient__Preset = 44

//
const QGradient__MixedHopes QGradient__Preset = 46

//
const QGradient__FlyHigh QGradient__Preset = 47

//
const QGradient__StrongBliss QGradient__Preset = 48

//
const QGradient__FreshMilk QGradient__Preset = 49

//
const QGradient__SnowAgain QGradient__Preset = 50

//
const QGradient__FebruaryInk QGradient__Preset = 51

//
const QGradient__KindSteel QGradient__Preset = 52

//
const QGradient__SoftGrass QGradient__Preset = 53

//
const QGradient__GrownEarly QGradient__Preset = 54

//
const QGradient__SharpBlues QGradient__Preset = 55

//
const QGradient__ShadyWater QGradient__Preset = 56

//
const QGradient__DirtyBeauty QGradient__Preset = 57

//
const QGradient__GreatWhale QGradient__Preset = 58

//
const QGradient__TeenNotebook QGradient__Preset = 59

//
const QGradient__PoliteRumors QGradient__Preset = 60

//
const QGradient__SweetPeriod QGradient__Preset = 61

//
const QGradient__WideMatrix QGradient__Preset = 62

//
const QGradient__SoftCherish QGradient__Preset = 63

//
const QGradient__RedSalvation QGradient__Preset = 64

//
const QGradient__BurningSpring QGradient__Preset = 65

//
const QGradient__NightParty QGradient__Preset = 66

//
const QGradient__SkyGlider QGradient__Preset = 67

//
const QGradient__HeavenPeach QGradient__Preset = 68

//
const QGradient__PurpleDivision QGradient__Preset = 69

//
const QGradient__AquaSplash QGradient__Preset = 70

//
const QGradient__SpikyNaga QGradient__Preset = 72

//
const QGradient__LoveKiss QGradient__Preset = 73

//
const QGradient__CleanMirror QGradient__Preset = 75

//
const QGradient__PremiumDark QGradient__Preset = 76

//
const QGradient__ColdEvening QGradient__Preset = 77

//
const QGradient__CochitiLake QGradient__Preset = 78

//
const QGradient__SummerGames QGradient__Preset = 79

//
const QGradient__PassionateBed QGradient__Preset = 80

//
const QGradient__MountainRock QGradient__Preset = 81

//
const QGradient__DesertHump QGradient__Preset = 82

//
const QGradient__JungleDay QGradient__Preset = 83

//
const QGradient__PhoenixStart QGradient__Preset = 84

//
const QGradient__OctoberSilence QGradient__Preset = 85

//
const QGradient__FarawayRiver QGradient__Preset = 86

//
const QGradient__AlchemistLab QGradient__Preset = 87

//
const QGradient__OverSun QGradient__Preset = 88

//
const QGradient__PremiumWhite QGradient__Preset = 89

//
const QGradient__MarsParty QGradient__Preset = 90

//
const QGradient__EternalConstance QGradient__Preset = 91

//
const QGradient__JapanBlush QGradient__Preset = 92

//
const QGradient__SmilingRain QGradient__Preset = 93

//
const QGradient__CloudyApple QGradient__Preset = 94

//
const QGradient__BigMango QGradient__Preset = 95

//
const QGradient__HealthyWater QGradient__Preset = 96

//
const QGradient__AmourAmour QGradient__Preset = 97

//
const QGradient__RiskyConcrete QGradient__Preset = 98

//
const QGradient__StrongStick QGradient__Preset = 99

//
const QGradient__ViciousStance QGradient__Preset = 100

//
const QGradient__PaloAlto QGradient__Preset = 101

//
const QGradient__HappyMemories QGradient__Preset = 102

//
const QGradient__MidnightBloom QGradient__Preset = 103

//
const QGradient__Crystalline QGradient__Preset = 104

//
const QGradient__PartyBliss QGradient__Preset = 106

//
const QGradient__ConfidentCloud QGradient__Preset = 107

//
const QGradient__LeCocktail QGradient__Preset = 108

//
const QGradient__RiverCity QGradient__Preset = 109

//
const QGradient__FrozenBerry QGradient__Preset = 110

//
const QGradient__ChildCare QGradient__Preset = 112

//
const QGradient__FlyingLemon QGradient__Preset = 113

//
const QGradient__NewRetrowave QGradient__Preset = 114

//
const QGradient__HiddenJaguar QGradient__Preset = 115

//
const QGradient__AboveTheSky QGradient__Preset = 116

//
const QGradient__Nega QGradient__Preset = 117

//
const QGradient__DenseWater QGradient__Preset = 118

//
const QGradient__Seashore QGradient__Preset = 120

//
const QGradient__MarbleWall QGradient__Preset = 121

//
const QGradient__CheerfulCaramel QGradient__Preset = 122

//
const QGradient__NightSky QGradient__Preset = 123

//
const QGradient__MagicLake QGradient__Preset = 124

//
const QGradient__YoungGrass QGradient__Preset = 125

//
const QGradient__ColorfulPeach QGradient__Preset = 126

//
const QGradient__GentleCare QGradient__Preset = 127

//
const QGradient__PlumBath QGradient__Preset = 128

//
const QGradient__HappyUnicorn QGradient__Preset = 129

//
const QGradient__AfricanField QGradient__Preset = 131

//
const QGradient__SolidStone QGradient__Preset = 132

//
const QGradient__OrangeJuice QGradient__Preset = 133

//
const QGradient__GlassWater QGradient__Preset = 134

//
const QGradient__NorthMiracle QGradient__Preset = 136

//
const QGradient__FruitBlend QGradient__Preset = 137

//
const QGradient__MillenniumPine QGradient__Preset = 138

//
const QGradient__HighFlight QGradient__Preset = 139

//
const QGradient__MoleHall QGradient__Preset = 140

//
const QGradient__SpaceShift QGradient__Preset = 142

//
const QGradient__ForestInei QGradient__Preset = 143

//
const QGradient__RoyalGarden QGradient__Preset = 144

//
const QGradient__RichMetal QGradient__Preset = 145

//
const QGradient__JuicyCake QGradient__Preset = 146

//
const QGradient__SmartIndigo QGradient__Preset = 147

//
const QGradient__SandStrike QGradient__Preset = 148

//
const QGradient__NorseBeauty QGradient__Preset = 149

//
const QGradient__AquaGuidance QGradient__Preset = 150

//
const QGradient__SunVeggie QGradient__Preset = 151

//
const QGradient__SeaLord QGradient__Preset = 152

//
const QGradient__BlackSea QGradient__Preset = 153

//
const QGradient__GrassShampoo QGradient__Preset = 154

//
const QGradient__LandingAircraft QGradient__Preset = 155

//
const QGradient__WitchDance QGradient__Preset = 156

//
const QGradient__SleeplessNight QGradient__Preset = 157

//
const QGradient__AngelCare QGradient__Preset = 158

//
const QGradient__CrystalRiver QGradient__Preset = 159

//
const QGradient__SoftLipstick QGradient__Preset = 160

//
const QGradient__SaltMountain QGradient__Preset = 161

//
const QGradient__PerfectWhite QGradient__Preset = 162

//
const QGradient__FreshOasis QGradient__Preset = 163

//
const QGradient__StrictNovember QGradient__Preset = 164

//
const QGradient__MorningSalad QGradient__Preset = 165

//
const QGradient__DeepRelief QGradient__Preset = 166

//
const QGradient__SeaStrike QGradient__Preset = 167

//
const QGradient__NightCall QGradient__Preset = 168

//
const QGradient__SupremeSky QGradient__Preset = 169

//
const QGradient__LightBlue QGradient__Preset = 170

//
const QGradient__MindCrawl QGradient__Preset = 171

//
const QGradient__LilyMeadow QGradient__Preset = 172

//
const QGradient__SugarLollipop QGradient__Preset = 173

//
const QGradient__SweetDessert QGradient__Preset = 174

//
const QGradient__MagicRay QGradient__Preset = 175

//
const QGradient__TeenParty QGradient__Preset = 176

//
const QGradient__FrozenHeat QGradient__Preset = 177

//
const QGradient__GagarinView QGradient__Preset = 178

//
const QGradient__FabledSunset QGradient__Preset = 179

//
const QGradient__PerfectBlue QGradient__Preset = 180

func (this *QGradient) PresetItemName(val int) string {
	switch val {
	case QGradient__WarmFlame: // 1
		return "WarmFlame"
	case QGradient__NightFade: // 2
		return "NightFade"
	case QGradient__SpringWarmth: // 3
		return "SpringWarmth"
	case QGradient__JuicyPeach: // 4
		return "JuicyPeach"
	case QGradient__YoungPassion: // 5
		return "YoungPassion"
	case QGradient__LadyLips: // 6
		return "LadyLips"
	case QGradient__SunnyMorning: // 7
		return "SunnyMorning"
	case QGradient__RainyAshville: // 8
		return "RainyAshville"
	case QGradient__FrozenDreams: // 9
		return "FrozenDreams"
	case QGradient__WinterNeva: // 10
		return "WinterNeva"
	case QGradient__DustyGrass: // 11
		return "DustyGrass"
	case QGradient__TemptingAzure: // 12
		return "TemptingAzure"
	case QGradient__HeavyRain: // 13
		return "HeavyRain"
	case QGradient__AmyCrisp: // 14
		return "AmyCrisp"
	case QGradient__MeanFruit: // 15
		return "MeanFruit"
	case QGradient__DeepBlue: // 16
		return "DeepBlue"
	case QGradient__RipeMalinka: // 17
		return "RipeMalinka"
	case QGradient__CloudyKnoxville: // 18
		return "CloudyKnoxville"
	case QGradient__MalibuBeach: // 19
		return "MalibuBeach"
	case QGradient__NewLife: // 20
		return "NewLife"
	case QGradient__TrueSunset: // 21
		return "TrueSunset"
	case QGradient__MorpheusDen: // 22
		return "MorpheusDen"
	case QGradient__RareWind: // 23
		return "RareWind"
	case QGradient__NearMoon: // 24
		return "NearMoon"
	case QGradient__WildApple: // 25
		return "WildApple"
	case QGradient__SaintPetersburg: // 26
		return "SaintPetersburg"
	case QGradient__PlumPlate: // 28
		return "PlumPlate"
	case QGradient__EverlastingSky: // 29
		return "EverlastingSky"
	case QGradient__HappyFisher: // 30
		return "HappyFisher"
	case QGradient__Blessing: // 31
		return "Blessing"
	case QGradient__SharpeyeEagle: // 32
		return "SharpeyeEagle"
	case QGradient__LadogaBottom: // 33
		return "LadogaBottom"
	case QGradient__LemonGate: // 34
		return "LemonGate"
	case QGradient__ItmeoBranding: // 35
		return "ItmeoBranding"
	case QGradient__ZeusMiracle: // 36
		return "ZeusMiracle"
	case QGradient__OldHat: // 37
		return "OldHat"
	case QGradient__StarWine: // 38
		return "StarWine"
	case QGradient__HappyAcid: // 41
		return "HappyAcid"
	case QGradient__AwesomePine: // 42
		return "AwesomePine"
	case QGradient__NewYork: // 43
		return "NewYork"
	case QGradient__ShyRainbow: // 44
		return "ShyRainbow"
	case QGradient__MixedHopes: // 46
		return "MixedHopes"
	case QGradient__FlyHigh: // 47
		return "FlyHigh"
	case QGradient__StrongBliss: // 48
		return "StrongBliss"
	case QGradient__FreshMilk: // 49
		return "FreshMilk"
	case QGradient__SnowAgain: // 50
		return "SnowAgain"
	case QGradient__FebruaryInk: // 51
		return "FebruaryInk"
	case QGradient__KindSteel: // 52
		return "KindSteel"
	case QGradient__SoftGrass: // 53
		return "SoftGrass"
	case QGradient__GrownEarly: // 54
		return "GrownEarly"
	case QGradient__SharpBlues: // 55
		return "SharpBlues"
	case QGradient__ShadyWater: // 56
		return "ShadyWater"
	case QGradient__DirtyBeauty: // 57
		return "DirtyBeauty"
	case QGradient__GreatWhale: // 58
		return "GreatWhale"
	case QGradient__TeenNotebook: // 59
		return "TeenNotebook"
	case QGradient__PoliteRumors: // 60
		return "PoliteRumors"
	case QGradient__SweetPeriod: // 61
		return "SweetPeriod"
	case QGradient__WideMatrix: // 62
		return "WideMatrix"
	case QGradient__SoftCherish: // 63
		return "SoftCherish"
	case QGradient__RedSalvation: // 64
		return "RedSalvation"
	case QGradient__BurningSpring: // 65
		return "BurningSpring"
	case QGradient__NightParty: // 66
		return "NightParty"
	case QGradient__SkyGlider: // 67
		return "SkyGlider"
	case QGradient__HeavenPeach: // 68
		return "HeavenPeach"
	case QGradient__PurpleDivision: // 69
		return "PurpleDivision"
	case QGradient__AquaSplash: // 70
		return "AquaSplash"
	case QGradient__SpikyNaga: // 72
		return "SpikyNaga"
	case QGradient__LoveKiss: // 73
		return "LoveKiss"
	case QGradient__CleanMirror: // 75
		return "CleanMirror"
	case QGradient__PremiumDark: // 76
		return "PremiumDark"
	case QGradient__ColdEvening: // 77
		return "ColdEvening"
	case QGradient__CochitiLake: // 78
		return "CochitiLake"
	case QGradient__SummerGames: // 79
		return "SummerGames"
	case QGradient__PassionateBed: // 80
		return "PassionateBed"
	case QGradient__MountainRock: // 81
		return "MountainRock"
	case QGradient__DesertHump: // 82
		return "DesertHump"
	case QGradient__JungleDay: // 83
		return "JungleDay"
	case QGradient__PhoenixStart: // 84
		return "PhoenixStart"
	case QGradient__OctoberSilence: // 85
		return "OctoberSilence"
	case QGradient__FarawayRiver: // 86
		return "FarawayRiver"
	case QGradient__AlchemistLab: // 87
		return "AlchemistLab"
	case QGradient__OverSun: // 88
		return "OverSun"
	case QGradient__PremiumWhite: // 89
		return "PremiumWhite"
	case QGradient__MarsParty: // 90
		return "MarsParty"
	case QGradient__EternalConstance: // 91
		return "EternalConstance"
	case QGradient__JapanBlush: // 92
		return "JapanBlush"
	case QGradient__SmilingRain: // 93
		return "SmilingRain"
	case QGradient__CloudyApple: // 94
		return "CloudyApple"
	case QGradient__BigMango: // 95
		return "BigMango"
	case QGradient__HealthyWater: // 96
		return "HealthyWater"
	case QGradient__AmourAmour: // 97
		return "AmourAmour"
	case QGradient__RiskyConcrete: // 98
		return "RiskyConcrete"
	case QGradient__StrongStick: // 99
		return "StrongStick"
	case QGradient__ViciousStance: // 100
		return "ViciousStance"
	case QGradient__PaloAlto: // 101
		return "PaloAlto"
	case QGradient__HappyMemories: // 102
		return "HappyMemories"
	case QGradient__MidnightBloom: // 103
		return "MidnightBloom"
	case QGradient__Crystalline: // 104
		return "Crystalline"
	case QGradient__PartyBliss: // 106
		return "PartyBliss"
	case QGradient__ConfidentCloud: // 107
		return "ConfidentCloud"
	case QGradient__LeCocktail: // 108
		return "LeCocktail"
	case QGradient__RiverCity: // 109
		return "RiverCity"
	case QGradient__FrozenBerry: // 110
		return "FrozenBerry"
	case QGradient__ChildCare: // 112
		return "ChildCare"
	case QGradient__FlyingLemon: // 113
		return "FlyingLemon"
	case QGradient__NewRetrowave: // 114
		return "NewRetrowave"
	case QGradient__HiddenJaguar: // 115
		return "HiddenJaguar"
	case QGradient__AboveTheSky: // 116
		return "AboveTheSky"
	case QGradient__Nega: // 117
		return "Nega"
	case QGradient__DenseWater: // 118
		return "DenseWater"
	case QGradient__Seashore: // 120
		return "Seashore"
	case QGradient__MarbleWall: // 121
		return "MarbleWall"
	case QGradient__CheerfulCaramel: // 122
		return "CheerfulCaramel"
	case QGradient__NightSky: // 123
		return "NightSky"
	case QGradient__MagicLake: // 124
		return "MagicLake"
	case QGradient__YoungGrass: // 125
		return "YoungGrass"
	case QGradient__ColorfulPeach: // 126
		return "ColorfulPeach"
	case QGradient__GentleCare: // 127
		return "GentleCare"
	case QGradient__PlumBath: // 128
		return "PlumBath"
	case QGradient__HappyUnicorn: // 129
		return "HappyUnicorn"
	case QGradient__AfricanField: // 131
		return "AfricanField"
	case QGradient__SolidStone: // 132
		return "SolidStone"
	case QGradient__OrangeJuice: // 133
		return "OrangeJuice"
	case QGradient__GlassWater: // 134
		return "GlassWater"
	case QGradient__NorthMiracle: // 136
		return "NorthMiracle"
	case QGradient__FruitBlend: // 137
		return "FruitBlend"
	case QGradient__MillenniumPine: // 138
		return "MillenniumPine"
	case QGradient__HighFlight: // 139
		return "HighFlight"
	case QGradient__MoleHall: // 140
		return "MoleHall"
	case QGradient__SpaceShift: // 142
		return "SpaceShift"
	case QGradient__ForestInei: // 143
		return "ForestInei"
	case QGradient__RoyalGarden: // 144
		return "RoyalGarden"
	case QGradient__RichMetal: // 145
		return "RichMetal"
	case QGradient__JuicyCake: // 146
		return "JuicyCake"
	case QGradient__SmartIndigo: // 147
		return "SmartIndigo"
	case QGradient__SandStrike: // 148
		return "SandStrike"
	case QGradient__NorseBeauty: // 149
		return "NorseBeauty"
	case QGradient__AquaGuidance: // 150
		return "AquaGuidance"
	case QGradient__SunVeggie: // 151
		return "SunVeggie"
	case QGradient__SeaLord: // 152
		return "SeaLord"
	case QGradient__BlackSea: // 153
		return "BlackSea"
	case QGradient__GrassShampoo: // 154
		return "GrassShampoo"
	case QGradient__LandingAircraft: // 155
		return "LandingAircraft"
	case QGradient__WitchDance: // 156
		return "WitchDance"
	case QGradient__SleeplessNight: // 157
		return "SleeplessNight"
	case QGradient__AngelCare: // 158
		return "AngelCare"
	case QGradient__CrystalRiver: // 159
		return "CrystalRiver"
	case QGradient__SoftLipstick: // 160
		return "SoftLipstick"
	case QGradient__SaltMountain: // 161
		return "SaltMountain"
	case QGradient__PerfectWhite: // 162
		return "PerfectWhite"
	case QGradient__FreshOasis: // 163
		return "FreshOasis"
	case QGradient__StrictNovember: // 164
		return "StrictNovember"
	case QGradient__MorningSalad: // 165
		return "MorningSalad"
	case QGradient__DeepRelief: // 166
		return "DeepRelief"
	case QGradient__SeaStrike: // 167
		return "SeaStrike"
	case QGradient__NightCall: // 168
		return "NightCall"
	case QGradient__SupremeSky: // 169
		return "SupremeSky"
	case QGradient__LightBlue: // 170
		return "LightBlue"
	case QGradient__MindCrawl: // 171
		return "MindCrawl"
	case QGradient__LilyMeadow: // 172
		return "LilyMeadow"
	case QGradient__SugarLollipop: // 173
		return "SugarLollipop"
	case QGradient__SweetDessert: // 174
		return "SweetDessert"
	case QGradient__MagicRay: // 175
		return "MagicRay"
	case QGradient__TeenParty: // 176
		return "TeenParty"
	case QGradient__FrozenHeat: // 177
		return "FrozenHeat"
	case QGradient__GagarinView: // 178
		return "GagarinView"
	case QGradient__FabledSunset: // 179
		return "FabledSunset"
	case QGradient__PerfectBlue: // 180
		return "PerfectBlue"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGradient_PresetItemName(val int) string {
	var nilthis *QGradient
	return nilthis.PresetItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10747() {
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
}

//  keep block end
