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
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"

func init() {
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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QVariantMap struct {
	*qtrt.CObject
}

// QMap<Key, T> & operator=(const QMap<Key, T> &)
func (this *QVariantMap) Operator_equal_0() {
	// QVariantMap_operator_equal_0()
}

// QMap<Key, T> & operator=(QMap<Key, T> &&)
func (this *QVariantMap) Operator_equal_1() {
	// QVariantMap_operator_equal_1()
}

// void swap(QMap<Key, T> &)
func (this *QVariantMap) Swap_0() {
	// QVariantMap_swap_0()
}

// std::map<Key, T> toStdMap()
func (this *QVariantMap) ToStdMap_0() {
	// QVariantMap_toStdMap_0()
}

// bool operator==(const QMap<Key, T> &)
func (this *QVariantMap) Operator_equal_equal_0() bool {
	// QVariantMap_operator_equal_equal_0()
	return 0 == 0
}

// bool operator!=(const QMap<Key, T> &)
func (this *QVariantMap) Operator_not_equal_0() bool {
	// QVariantMap_operator_not_equal_0()
	return 0 == 0
}

// int size()
func (this *QVariantMap) Size_0() int {
	// QVariantMap_size_0()
	return 0
}

// bool isEmpty()
func (this *QVariantMap) IsEmpty_0() bool {
	// QVariantMap_isEmpty_0()
	return 0 == 0
}

// void detach()
func (this *QVariantMap) Detach_0() {
	// QVariantMap_detach_0()
}

// bool isDetached()
func (this *QVariantMap) IsDetached_0() bool {
	// QVariantMap_isDetached_0()
	return 0 == 0
}

// void setSharable(_Bool)
func (this *QVariantMap) SetSharable_0() {
	// QVariantMap_setSharable_0()
}

// bool isSharedWith(const QMap<Key, T> &)
func (this *QVariantMap) IsSharedWith_0() bool {
	// QVariantMap_isSharedWith_0()
	return 0 == 0
}

// void clear()
func (this *QVariantMap) Clear_0() {
	// QVariantMap_clear_0()
}

// int remove(const Key &)
func (this *QVariantMap) Remove_0() int {
	// QVariantMap_remove_0()
	return 0
}

// T take(const Key &)
func (this *QVariantMap) Take_0() *QVariant {
	// QVariantMap_take_0()
	return &QVariant{}
}

// bool contains(const Key &)
func (this *QVariantMap) Contains_0() bool {
	// QVariantMap_contains_0()
	return 0 == 0
}

// const Key key(const T &, const Key &)
func (this *QVariantMap) Key_0() {
	// QVariantMap_key_0()
}

// const T value(const Key &, const T &)
func (this *QVariantMap) Value_0() *QVariant {
	// QVariantMap_value_0()
	return &QVariant{}
}

// T & operator[](const Key &)
func (this *QVariantMap) Operator_get_index_0() *QVariant {
	// QVariantMap_operator_get_index_0()
	return &QVariant{}
}

// const T operator[](const Key &)
func (this *QVariantMap) Operator_get_index_1() *QVariant {
	// QVariantMap_operator_get_index_1()
	return &QVariant{}
}

// QList<Key> uniqueKeys()
func (this *QVariantMap) UniqueKeys_0() {
	// QVariantMap_uniqueKeys_0()
}

// QList<Key> keys()
func (this *QVariantMap) Keys_0() {
	// QVariantMap_keys_0()
}

// QList<Key> keys(const T &)
func (this *QVariantMap) Keys_1() {
	// QVariantMap_keys_1()
}

// QList<T> values()
func (this *QVariantMap) Values_0() {
	// QVariantMap_values_0()
}

// QList<T> values(const Key &)
func (this *QVariantMap) Values_1() {
	// QVariantMap_values_1()
}

// int count(const Key &)
func (this *QVariantMap) Count_0() int {
	// QVariantMap_count_0()
	return 0
}

// const Key & firstKey()
func (this *QVariantMap) FirstKey_0() {
	// QVariantMap_firstKey_0()
}

// const Key & lastKey()
func (this *QVariantMap) LastKey_0() {
	// QVariantMap_lastKey_0()
}

// T & first()
func (this *QVariantMap) First_0() *QVariant {
	// QVariantMap_first_0()
	return &QVariant{}
}

// const T & first()
func (this *QVariantMap) First_1() *QVariant {
	// QVariantMap_first_1()
	return &QVariant{}
}

// T & last()
func (this *QVariantMap) Last_0() *QVariant {
	// QVariantMap_last_0()
	return &QVariant{}
}

// const T & last()
func (this *QVariantMap) Last_1() *QVariant {
	// QVariantMap_last_1()
	return &QVariant{}
}

// QMap::iterator begin()
func (this *QVariantMap) Begin_0() {
	// QVariantMap_begin_0()
}

// QMap::const_iterator begin()
func (this *QVariantMap) Begin_1() {
	// QVariantMap_begin_1()
}

// QMap::const_iterator constBegin()
func (this *QVariantMap) ConstBegin_0() {
	// QVariantMap_constBegin_0()
}

// QMap::const_iterator cbegin()
func (this *QVariantMap) Cbegin_0() {
	// QVariantMap_cbegin_0()
}

// QMap::iterator end()
func (this *QVariantMap) End_0() {
	// QVariantMap_end_0()
}

// QMap::const_iterator end()
func (this *QVariantMap) End_1() {
	// QVariantMap_end_1()
}

// QMap::const_iterator constEnd()
func (this *QVariantMap) ConstEnd_0() {
	// QVariantMap_constEnd_0()
}

// QMap::const_iterator cend()
func (this *QVariantMap) Cend_0() {
	// QVariantMap_cend_0()
}

// QMap::key_iterator keyBegin()
func (this *QVariantMap) KeyBegin_0() {
	// QVariantMap_keyBegin_0()
}

// QMap::key_iterator keyEnd()
func (this *QVariantMap) KeyEnd_0() {
	// QVariantMap_keyEnd_0()
}

// QMap::key_value_iterator keyValueBegin()
func (this *QVariantMap) KeyValueBegin_0() {
	// QVariantMap_keyValueBegin_0()
}

// QMap::key_value_iterator keyValueEnd()
func (this *QVariantMap) KeyValueEnd_0() {
	// QVariantMap_keyValueEnd_0()
}

// QMap::const_key_value_iterator keyValueBegin()
func (this *QVariantMap) KeyValueBegin_1() {
	// QVariantMap_keyValueBegin_1()
}

// QMap::const_key_value_iterator constKeyValueBegin()
func (this *QVariantMap) ConstKeyValueBegin_0() {
	// QVariantMap_constKeyValueBegin_0()
}

// QMap::const_key_value_iterator keyValueEnd()
func (this *QVariantMap) KeyValueEnd_1() {
	// QVariantMap_keyValueEnd_1()
}

// QMap::const_key_value_iterator constKeyValueEnd()
func (this *QVariantMap) ConstKeyValueEnd_0() {
	// QVariantMap_constKeyValueEnd_0()
}

// QMap::iterator erase(class QMap::iterator)
func (this *QVariantMap) Erase_0() {
	// QVariantMap_erase_0()
}

// int count()
func (this *QVariantMap) Count_1() int {
	// QVariantMap_count_1()
	return 0
}

// QMap::iterator find(const Key &)
func (this *QVariantMap) Find_0() {
	// QVariantMap_find_0()
}

// QMap::const_iterator find(const Key &)
func (this *QVariantMap) Find_1() {
	// QVariantMap_find_1()
}

// QMap::const_iterator constFind(const Key &)
func (this *QVariantMap) ConstFind_0() {
	// QVariantMap_constFind_0()
}

// QMap::iterator lowerBound(const Key &)
func (this *QVariantMap) LowerBound_0() {
	// QVariantMap_lowerBound_0()
}

// QMap::const_iterator lowerBound(const Key &)
func (this *QVariantMap) LowerBound_1() {
	// QVariantMap_lowerBound_1()
}

// QMap::iterator upperBound(const Key &)
func (this *QVariantMap) UpperBound_0() {
	// QVariantMap_upperBound_0()
}

// QMap::const_iterator upperBound(const Key &)
func (this *QVariantMap) UpperBound_1() {
	// QVariantMap_upperBound_1()
}

// QMap::iterator insert(const Key &, const T &)
func (this *QVariantMap) Insert_0() {
	// QVariantMap_insert_0()
}

// QMap::iterator insert(class QMap::const_iterator, const Key &, const T &)
func (this *QVariantMap) Insert_1() {
	// QVariantMap_insert_1()
}

// QMap::iterator insertMulti(const Key &, const T &)
func (this *QVariantMap) InsertMulti_0() {
	// QVariantMap_insertMulti_0()
}

// QMap::iterator insertMulti(class QMap::const_iterator, const Key &, const T &)
func (this *QVariantMap) InsertMulti_1() {
	// QVariantMap_insertMulti_1()
}

// QMap<Key, T> & unite(const QMap<Key, T> &)
func (this *QVariantMap) Unite_0() {
	// QVariantMap_unite_0()
}

// bool empty()
func (this *QVariantMap) Empty_0() bool {
	// QVariantMap_empty_0()
	return 0 == 0
}

// QPair<QMap::iterator, QMap::iterator> equal_range(const Key &)
func (this *QVariantMap) Equal_range_0() {
	// QVariantMap_equal_range_0()
}

// QPair<QMap::const_iterator, QMap::const_iterator> equal_range(const Key &)
func (this *QVariantMap) Equal_range_1() {
	// QVariantMap_equal_range_1()
}

// void detach_helper()
func (this *QVariantMap) Detach_helper_0() {
	// QVariantMap_detach_helper_0()
}

// bool isValidIterator(const class QMap::const_iterator &)
func (this *QVariantMap) IsValidIterator_0() bool {
	// QVariantMap_isValidIterator_0()
	return 0 == 0
}

//  body block end
