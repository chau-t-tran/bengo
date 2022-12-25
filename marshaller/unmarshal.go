package marshaller

import (
	"errors"
	"reflect"

	"github.com/chau-t-tran/bengo/ast"
	"github.com/chau-t-tran/bengo/parser"
)

func Unmarshal(data string, v any) error {
	node, err := parser.Parse(data)
	if err != nil {
		return err
	}
	dictNode := node.(*ast.DictNode)

	types := reflect.TypeOf(v).Elem()
	values := reflect.ValueOf(v).Elem()

	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		tag := field.Tag.Get("bencode")
		if tag == "" {
			return errors.New("bencode tag not found")
		}

		fieldNode, ok := dictNode.Value()[tag]
		if !ok {
			return errors.New(tag + " not found")
		}

		var valueNode *ast.IntNode = fieldNode.(*ast.IntNode)
		newValue := int64(valueNode.Value())
		values.Field(i).SetInt(newValue)
	}

	return nil
}
