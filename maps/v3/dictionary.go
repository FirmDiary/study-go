package v1

const (
	ErrNotFound  = DictionaryErr("could not find the word you were looking for")
	ErrKeyExists = DictionaryErr("cannot add word because it already exists")
)

//我们将错误声明为常量，这需要我们创建自己的 DictionaryErr 类型来实现 error 接口
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

//我们现在有很好的方法来搜索字典。但是，我们无法在字典中添加新单词。
func (d Dictionary) Search(word string) (string, error) {
	// map 查找的有趣特性。它可以返回两个值。第二个值是一个布尔值，表示是否成功找到 key
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key string, word string) error {
	//Map 有一个有趣的特性，不使用指针传递你就可以修改它们。这是因为 map 是引用类型
	//它底层的数据结构是 hash table 或 hash map

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = word
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}
