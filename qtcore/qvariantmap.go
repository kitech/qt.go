package qtcore

// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>

//  header block end

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

//  ext block end

//  keep block begin

func init_unused_10321() {
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
}

//  keep block end

//  body block begin
type QVariantMap struct {
	*qtrt.CObject
}

// QMap<Key, T> & operator=(const QMap<Key, T> &)
func (this *QVariantMap) Operator_equal0() {
	// QVariantMap_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap<Key, T> & operator=(QMap<Key, T> &&)
func (this *QVariantMap) Operator_equal1() {
	// QVariantMap_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(QMap<Key, T> &)
func (this *QVariantMap) Swap0() {
	// QVariantMap_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// std::map<Key, T> toStdMap()
func (this *QVariantMap) ToStdMap0() {
	// QVariantMap_toStdMap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_toStdMap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QMap<Key, T> &)
func (this *QVariantMap) Operator_equal_equal0() bool {
	// QVariantMap_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QMap<Key, T> &)
func (this *QVariantMap) Operator_not_equal0() bool {
	// QVariantMap_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QVariantMap) Size0() int {
	// QVariantMap_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool isEmpty()
func (this *QVariantMap) IsEmpty0() bool {
	// QVariantMap_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void detach()
func (this *QVariantMap) Detach0() {
	// QVariantMap_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QVariantMap) IsDetached0() bool {
	// QVariantMap_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QVariantMap) SetSharable0() {
	// QVariantMap_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QMap<Key, T> &)
func (this *QVariantMap) IsSharedWith0() bool {
	// QVariantMap_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QVariantMap) Clear0() {
	// QVariantMap_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int remove(const Key &)
func (this *QVariantMap) Remove0() int {
	// QVariantMap_remove_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_remove_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T take(const Key &)
func (this *QVariantMap) Take0() *QVariant {
	// QVariantMap_take_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_take_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// bool contains(const Key &)
func (this *QVariantMap) Contains0() bool {
	// QVariantMap_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// const Key key(const T &, const Key &)
func (this *QVariantMap) Key0() {
	// QVariantMap_key_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_key_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T value(const Key &, const T &)
func (this *QVariantMap) Value0() *QVariant {
	// QVariantMap_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T & operator[](const Key &)
func (this *QVariantMap) Operator_get_index0() *QVariant {
	// QVariantMap_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T operator[](const Key &)
func (this *QVariantMap) Operator_get_index1() *QVariant {
	// QVariantMap_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// QList<Key> keys()
func (this *QVariantMap) Keys0() {
	// QVariantMap_keys_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_keys_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<Key> keys(const T &)
func (this *QVariantMap) Keys1() {
	// QVariantMap_keys_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_keys_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> values()
func (this *QVariantMap) Values0() {
	// QVariantMap_values_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_values_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<Key> uniqueKeys()
func (this *QVariantMap) UniqueKeys0() {
	// QVariantMap_uniqueKeys_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_uniqueKeys_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> values(const Key &)
func (this *QVariantMap) Values1() {
	// QVariantMap_values_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_values_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count(const Key &)
func (this *QVariantMap) Count0() int {
	// QVariantMap_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// const Key & firstKey()
func (this *QVariantMap) FirstKey0() {
	// QVariantMap_firstKey_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_firstKey_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const Key & lastKey()
func (this *QVariantMap) LastKey0() {
	// QVariantMap_lastKey_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_lastKey_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & first()
func (this *QVariantMap) First0() *QVariant {
	// QVariantMap_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & first()
func (this *QVariantMap) First1() *QVariant {
	// QVariantMap_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T & last()
func (this *QVariantMap) Last0() *QVariant {
	// QVariantMap_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & last()
func (this *QVariantMap) Last1() *QVariant {
	// QVariantMap_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// QMap::iterator begin()
func (this *QVariantMap) Begin0() {
	// QVariantMap_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_iterator begin()
func (this *QVariantMap) Begin1() {
	// QVariantMap_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_iterator constBegin()
func (this *QVariantMap) ConstBegin0() {
	// QVariantMap_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_iterator cbegin()
func (this *QVariantMap) Cbegin0() {
	// QVariantMap_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::iterator end()
func (this *QVariantMap) End0() {
	// QVariantMap_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_iterator end()
func (this *QVariantMap) End1() {
	// QVariantMap_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_iterator constEnd()
func (this *QVariantMap) ConstEnd0() {
	// QVariantMap_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_iterator cend()
func (this *QVariantMap) Cend0() {
	// QVariantMap_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::key_iterator keyBegin()
func (this *QVariantMap) KeyBegin0() {
	// QVariantMap_keyBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_keyBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::key_iterator keyEnd()
func (this *QVariantMap) KeyEnd0() {
	// QVariantMap_keyEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_keyEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::key_value_iterator keyValueBegin()
func (this *QVariantMap) KeyValueBegin0() {
	// QVariantMap_keyValueBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_keyValueBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::key_value_iterator keyValueEnd()
func (this *QVariantMap) KeyValueEnd0() {
	// QVariantMap_keyValueEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_keyValueEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_key_value_iterator keyValueBegin()
func (this *QVariantMap) KeyValueBegin1() {
	// QVariantMap_keyValueBegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_keyValueBegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_key_value_iterator constKeyValueBegin()
func (this *QVariantMap) ConstKeyValueBegin0() {
	// QVariantMap_constKeyValueBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_constKeyValueBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_key_value_iterator keyValueEnd()
func (this *QVariantMap) KeyValueEnd1() {
	// QVariantMap_keyValueEnd_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_keyValueEnd_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_key_value_iterator constKeyValueEnd()
func (this *QVariantMap) ConstKeyValueEnd0() {
	// QVariantMap_constKeyValueEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_constKeyValueEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::iterator erase(QMap::iterator)
func (this *QVariantMap) Erase0() {
	// QVariantMap_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QVariantMap) Count1() int {
	// QVariantMap_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QMap::iterator find(const Key &)
func (this *QVariantMap) Find0() {
	// QVariantMap_find_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_find_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_iterator find(const Key &)
func (this *QVariantMap) Find1() {
	// QVariantMap_find_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_find_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_iterator constFind(const Key &)
func (this *QVariantMap) ConstFind0() {
	// QVariantMap_constFind_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_constFind_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::iterator lowerBound(const Key &)
func (this *QVariantMap) LowerBound0() {
	// QVariantMap_lowerBound_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_lowerBound_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_iterator lowerBound(const Key &)
func (this *QVariantMap) LowerBound1() {
	// QVariantMap_lowerBound_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_lowerBound_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::iterator upperBound(const Key &)
func (this *QVariantMap) UpperBound0() {
	// QVariantMap_upperBound_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_upperBound_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::const_iterator upperBound(const Key &)
func (this *QVariantMap) UpperBound1() {
	// QVariantMap_upperBound_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_upperBound_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::iterator insert(const Key &, const T &)
func (this *QVariantMap) Insert0() {
	// QVariantMap_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::iterator insert(QMap::const_iterator, const Key &, const T &)
func (this *QVariantMap) Insert1() {
	// QVariantMap_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(const QMap<Key, T> &)
func (this *QVariantMap) Insert2() {
	// QVariantMap_insert_2()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_insert_2", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::iterator insertMulti(const Key &, const T &)
func (this *QVariantMap) InsertMulti0() {
	// QVariantMap_insertMulti_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_insertMulti_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap::iterator insertMulti(QMap::const_iterator, const Key &, const T &)
func (this *QVariantMap) InsertMulti1() {
	// QVariantMap_insertMulti_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_insertMulti_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QMap<Key, T> & unite(const QMap<Key, T> &)
func (this *QVariantMap) Unite0() {
	// QVariantMap_unite_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_unite_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QVariantMap) Empty0() bool {
	// QVariantMap_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QPair<QMap::iterator, QMap::iterator> equal_range(const Key &)
func (this *QVariantMap) Equal_range0() {
	// QVariantMap_equal_range_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_equal_range_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QPair<QMap::const_iterator, QMap::const_iterator> equal_range(const Key &)
func (this *QVariantMap) Equal_range1() {
	// QVariantMap_equal_range_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_equal_range_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QVariantMap) Detach_helper0() {
	// QVariantMap_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QMap::const_iterator &)
func (this *QVariantMap) IsValidIterator0() bool {
	// QVariantMap_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantMap_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

//  body block end
