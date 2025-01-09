// VM for Reylang
package reylang

import (
	"fmt"
	"strings"

	"github.com/ramonberrutti/reylang/protogen/reylang"
	reylangpb "github.com/ramonberrutti/reylang/protogen/reylang"
)

type VM struct {
	stack     []any
	constants []any
	symbols   map[string]any
	iterator  *iterator
}

type iterator struct {
	collection any
	current    int
	key        string
	value      string
}

func NewVM(constants []any) *VM {
	return &VM{
		stack:     make([]any, 0),
		constants: constants,
		symbols:   make(map[string]any),
	}
}

func (vm *VM) SetSymbol(name string, value any) {
	vm.symbols[name] = value
}

func (vm *VM) Run(bytecode []Instruction) {
	pc := 0

	for pc < len(bytecode) {
		i := bytecode[pc]
		fmt.Printf("Executing %s with Operand: %+v\n", i.Code, i.Operand)
		switch i.Code {
		case reylang.OpCode_HALT:
			return
		case reylang.OpCode_LOAD_CONST:
			vm.push(vm.constants[i.Operand.(int)])
		case reylang.OpCode_LOAD_VAR:
			// if variable is not found, panic
			operand := i.Operand.(string)

			parts := strings.Split(operand, ".")
			if len(parts) > 1 {
				// handle nested variables
				value := vm.symbols[parts[0]]
				for _, part := range parts[1:] {
					switch value.(type) {
					case map[string]any:
						value = value.(map[string]any)[part]
					}
				}
				vm.push(value)
			} else {
				vm.push(vm.symbols[operand])
			}
		case reylang.OpCode_CALL:
			// grab from the symbols table
			fmt.Printf("Calling %+v\n", i.Operand)
			funcName := i.Operand.(string)

			f, ok := vm.symbols[funcName].(func(...any) any)
			if !ok {
				panic("function not found")
			}

			arg := vm.pop()
			f(arg)

		case reylang.OpCode_COMPARE:
			right := vm.pop()
			left := vm.pop()

			switch i.Operand.(reylangpb.ComparisonNode_Operator) {
			case reylangpb.ComparisonNode_OPERATOR_EQ:
				fmt.Printf("Comparing %v == %v\n", left, right)
				vm.push(left == right)
			case reylangpb.ComparisonNode_OPERATOR_NE:
				vm.push(left != right)
			}

		case reylang.OpCode_JUMP:
			pc = i.Operand.(int)
			continue
		case reylang.OpCode_JUMP_IF_FALSE:
			if !vm.pop().(bool) {
				pc = i.Operand.(int)
				continue
			}
		case reylang.OpCode_ITER_START:
			operand := i.Operand.([]string)
			key := operand[0]
			value := operand[1]
			collectionName := operand[2]
			collection := vm.symbols[collectionName]
			switch collection.(type) {
			case []any:
				vm.iterator = &iterator{
					collection: collection,
					current:    -1,
					key:        key,
					value:      value,
				}
			case map[string]any:
				panic("not implemented")
			}

		case reylang.OpCode_ITER_NEXT:
			vm.iterator.current++
			if vm.iterator.current < len(vm.iterator.collection.([]any)) {
				index := vm.iterator.current
				vm.symbols[vm.iterator.key] = index
				vm.symbols[vm.iterator.value] = vm.iterator.collection.([]any)[index]
			} else {
				pc = i.Operand.(int)
			}

		case reylang.OpCode_ITER_END:
			vm.iterator = nil
		}
		pc++
	}
}

func (vm *VM) push(v any) {
	vm.stack = append(vm.stack, v)
}

func (vm *VM) pop() any {
	if len(vm.stack) == 0 {
		return nil
	}

	v := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return v
}
