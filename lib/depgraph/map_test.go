package depgraph

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Helcaraxan/gomod/lib/modules"
)

func TestMapNew(t *testing.T) {
	newMap := NewModuleDependencies()
	assert.NotNil(t, newMap.dependencyMap)
	assert.NotNil(t, newMap.dependencyList)
}

func TestMapCopy(t *testing.T) {
	dependencyA := &Module{Info: &modules.ModuleInfo{Path: "dependency_a"}}
	dependencyB := &Module{Info: &modules.ModuleInfo{Path: "dependency_b"}}

	originalMap := NewModuleDependencies()
	originalMap.Add(&ModuleReference{Module: dependencyA})
	copiedMap := originalMap.Copy()
	originalMap.Add(&ModuleReference{Module: dependencyB})

	_, okA := originalMap.Get("dependency_a")
	_, okB := originalMap.Get("dependency_b")
	assert.True(t, okA)
	assert.True(t, okB)

	_, okA = copiedMap.Get("dependency_a")
	_, okB = copiedMap.Get("dependency_b")
	assert.True(t, okA)
	assert.False(t, okB)
}

func TestMapAdd(t *testing.T) {
	dependencyA := &Module{Info: &modules.ModuleInfo{Path: "dependency_a"}}

	newMap := NewModuleDependencies()
	newMap.Add(&ModuleReference{Module: dependencyA})
	_, ok := newMap.Get("dependency_a")
	assert.True(t, ok)
	newMap.Add(&ModuleReference{Module: dependencyA})
	_, ok = newMap.Get("dependency_a")
	assert.True(t, ok)
}

func TestMapDelete(t *testing.T) {
	dependencyA := &Module{Info: &modules.ModuleInfo{Path: "dependency_a"}}

	newMap := NewModuleDependencies()

	newMap.Delete("dependency_a")

	newMap.Add(&ModuleReference{Module: dependencyA})
	newMap.Delete("dependency_a")
	assert.NotContains(t, newMap.List(), &Module{Info: &modules.ModuleInfo{Path: "dependency_a"}})
}

func TestMapLen(t *testing.T) {
	dependencyA := &Module{Info: &modules.ModuleInfo{Path: "dependency_a"}}
	dependencyB := &Module{Info: &modules.ModuleInfo{Path: "dependency_b"}}

	newMap := NewModuleDependencies()
	assert.Equal(t, 0, newMap.Len())

	newMap.Add(&ModuleReference{Module: dependencyA})
	assert.Equal(t, 1, newMap.Len())

	newMap.Add(&ModuleReference{Module: dependencyA})
	assert.Equal(t, 1, newMap.Len())

	newMap.Add(&ModuleReference{Module: dependencyB})
	assert.Equal(t, 2, newMap.Len())

	newMap.Delete("dependency_a")
	assert.Equal(t, 1, newMap.Len())
}

func TestMapList(t *testing.T) {
	dependencyA := &Module{Info: &modules.ModuleInfo{Path: "dependency_a"}}
	dependencyB := &Module{Info: &modules.ModuleInfo{Path: "dependency_b"}}

	newMap := NewModuleDependencies()
	newMap.Add(&ModuleReference{Module: dependencyB})
	newMap.Add(&ModuleReference{Module: dependencyA})

	list := newMap.List()
	isSorted := sort.SliceIsSorted(list, func(i int, j int) bool { return list[i].Name() < list[j].Name() })
	assert.True(t, isSorted)
}
