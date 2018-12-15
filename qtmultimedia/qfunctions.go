package qtmultimedia

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
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
	if false {
		qtnetwork.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtMultimedia/qmediatimerange.h:122
// index:49
// Invalid Visibility=Default Availability=Available
// [8] QMediaTimeRange operator+(const QMediaTimeRange &, const QMediaTimeRange &)

/*

 */
func Operator_add49(arg0 QMediaTimeRange_ITF, arg1 QMediaTimeRange_ITF) *QMediaTimeRange /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeRange_PTR() != nil {
		convArg0 = arg0.QMediaTimeRange_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QMediaTimeRange_PTR() != nil {
		convArg1 = arg1.QMediaTimeRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZplRK15QMediaTimeRangeS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaTimeRange)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamera.h:258
// index:72
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QCamera::LockTypes::enum_type, int)

/*

 */
func Operator_or72(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN7QCamera8LockTypeEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:105
// index:73
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QMediaServiceProviderHint::Features::enum_type, int)

/*

 */
func Operator_or73(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN25QMediaServiceProviderHint7FeatureEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h:63
// index:74
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QVideoFilterRunnable::RunFlags::enum_type, int)

/*

 */
func Operator_or74(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN20QVideoFilterRunnable7RunFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraimagecapture.h:149
// index:75
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QCameraImageCapture::CaptureDestinations::enum_type, int)

/*

 */
func Operator_or75(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN19QCameraImageCapture18CaptureDestinationEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:181
// index:76
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QCameraExposure::FlashModes::enum_type, int)

/*

 */
func Operator_or76(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN15QCameraExposure9FlashModeEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamerafocus.h:166
// index:77
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QCameraFocus::FocusModes::enum_type, int)

/*

 */
func Operator_or77(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN12QCameraFocus9FocusModeEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamera.h:272
// index:60
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator!=(const QCamera::FrameRateRange &, const QCamera::FrameRateRange &)

/*

 */
func Operator_not_equal60(r1 int, r2 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZneRKN7QCamera14FrameRateRangeES2_", qtrt.FFI_TYPE_POINTER, &r1, &r2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:96
// index:61
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator!=(const QCameraViewfinderSettings &, const QCameraViewfinderSettings &)

/*

 */
func Operator_not_equal61(lhs QCameraViewfinderSettings_ITF, rhs QCameraViewfinderSettings_ITF) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QCameraViewfinderSettings_PTR() != nil {
		convArg0 = lhs.QCameraViewfinderSettings_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QCameraViewfinderSettings_PTR() != nil {
		convArg1 = rhs.QCameraViewfinderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneRK25QCameraViewfinderSettingsS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:77
// index:62
// Invalid Visibility=Default Availability=Available
// [1] bool operator!=(const QMediaTimeInterval &, const QMediaTimeInterval &)

/*

 */
func Operator_not_equal62(arg0 QMediaTimeInterval_ITF, arg1 QMediaTimeInterval_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeInterval_PTR() != nil {
		convArg0 = arg0.QMediaTimeInterval_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QMediaTimeInterval_PTR() != nil {
		convArg1 = arg1.QMediaTimeInterval_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneRK18QMediaTimeIntervalS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:121
// index:63
// Invalid Visibility=Default Availability=Available
// [1] bool operator!=(const QMediaTimeRange &, const QMediaTimeRange &)

/*

 */
func Operator_not_equal63(arg0 QMediaTimeRange_ITF, arg1 QMediaTimeRange_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeRange_PTR() != nil {
		convArg0 = arg0.QMediaTimeRange_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QMediaTimeRange_PTR() != nil {
		convArg1 = arg1.QMediaTimeRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneRK15QMediaTimeRangeS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:123
// index:26
// Invalid Visibility=Default Availability=Available
// [8] QMediaTimeRange operator-(const QMediaTimeRange &, const QMediaTimeRange &)

/*
Returns a time range containing r2 subtracted from r1.
*/
func Operator_minus26(arg0 QMediaTimeRange_ITF, arg1 QMediaTimeRange_ITF) *QMediaTimeRange /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeRange_PTR() != nil {
		convArg0 = arg0.QMediaTimeRange_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QMediaTimeRange_PTR() != nil {
		convArg1 = arg1.QMediaTimeRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZmiRK15QMediaTimeRangeS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaTimeRange)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamera.h:264
// index:62
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator==(const QCamera::FrameRateRange &, const QCamera::FrameRateRange &)

/*

 */
func Operator_equal_equal62(r1 int, r2 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZeqRKN7QCamera14FrameRateRangeES2_", qtrt.FFI_TYPE_POINTER, &r1, &r2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:76
// index:63
// Invalid Visibility=Default Availability=Available
// [1] bool operator==(const QMediaTimeInterval &, const QMediaTimeInterval &)

/*

 */
func Operator_equal_equal63(arg0 QMediaTimeInterval_ITF, arg1 QMediaTimeInterval_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeInterval_PTR() != nil {
		convArg0 = arg0.QMediaTimeInterval_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QMediaTimeInterval_PTR() != nil {
		convArg1 = arg1.QMediaTimeInterval_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZeqRK18QMediaTimeIntervalS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:120
// index:64
// Invalid Visibility=Default Availability=Available
// [1] bool operator==(const QMediaTimeRange &, const QMediaTimeRange &)

/*

 */
func Operator_equal_equal64(arg0 QMediaTimeRange_ITF, arg1 QMediaTimeRange_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeRange_PTR() != nil {
		convArg0 = arg0.QMediaTimeRange_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QMediaTimeRange_PTR() != nil {
		convArg1 = arg1.QMediaTimeRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZeqRK15QMediaTimeRangeS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudio.h:82
// index:0
// Invalid Visibility=Default Availability=Available
// [8] qreal convertVolume(qreal, QAudio::VolumeScale, QAudio::VolumeScale)

/*
Converts an audio volume from a volume scale to another, and returns the result.

Depending on the context, different scales are used to represent audio volume. All Qt Multimedia classes that have an audio volume use a linear scale, the reason is that the loudness of a speaker is controlled by modulating its voltage on a linear scale. The human ear on the other hand, perceives loudness in a logarithmic way. Using a logarithmic scale for volume controls is therefore appropriate in most applications. The decibel scale is logarithmic by nature and is commonly used to define sound levels, it is usually used for UI volume controls in professional audio applications. The cubic scale is a computationally cheap approximation of a logarithmic scale, it provides more control over lower volume levels.

The following example shows how to convert the volume value from a slider control before passing it to a QMediaPlayer. As a result, the perceived increase in volume is the same when increasing the volume slider from 20 to 30 as it is from 50 to 60:


  void applyVolume(int volumeSliderValue)
  {
      // volumeSliderValue is in the range [0..100]

      qreal linearVolume = QAudio::convertVolume(volumeSliderValue / qreal(100.0),
                                                 QAudio::LogarithmicVolumeScale,
                                                 QAudio::LinearVolumeScale);

      player.setVolume(qRound(linearVolume * 100));
  }



This function was introduced in  Qt 5.8.

See also VolumeScale, QMediaPlayer::setVolume(), QAudioOutput::setVolume(), QAudioInput::setVolume(), QSoundEffect::setVolume(), and QMediaRecorder::setVolume().
*/
func ConvertVolume(volume float64, from int, to int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QAudio13convertVolumeEdNS_11VolumeScaleES0_", qtrt.FFI_TYPE_POINTER, volume, from, to)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h:94
// index:61
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QCameraViewfinderSettings &, QCameraViewfinderSettings &)

/*
Swaps this viewfinder settings object with other. This function is very fast and never fails.
*/
func Swap61(value1 QCameraViewfinderSettings_ITF, value2 QCameraViewfinderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QCameraViewfinderSettings_PTR() != nil {
		convArg0 = value1.QCameraViewfinderSettings_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QCameraViewfinderSettings_PTR() != nil {
		convArg1 = value2.QCameraViewfinderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR25QCameraViewfinderSettingsS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

//  body block end
