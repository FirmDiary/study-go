package v1

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

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
