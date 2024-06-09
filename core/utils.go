package core

import (
	"github.com/deckarep/golang-set/v2"
	"reflect"
	"ubuntuInit/enum/InstallOptions"
)

// #region InstallChoices struct
type InstallChoices struct {
	AptPreInstallMap map[string]mapset.Set[string]
	Options          []string
	CurrentChoices   []string
	MapChoices       map[string]bool
}

func SliceContains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func MapContains(sliceMap interface{}, item interface{}) bool {
	switch reflect.TypeOf(sliceMap).Kind() {
	case reflect.Map:
		s := reflect.ValueOf(sliceMap)
		return s.MapIndex(reflect.ValueOf(item)).IsValid()
	}
	return false
}

func (ic *InstallChoices) CompareChoices() map[string]bool {
	for _, option := range ic.Options {
		if SliceContains(ic.CurrentChoices, option) {
			ic.MapChoices[option] = true
		} else {
			ic.MapChoices[option] = false
		}
	}
	return ic.MapChoices
}

func NewInstallChoices() *InstallChoices {
	options := InstallOptions.Values()
	return &InstallChoices{
		AptPreInstallMap: map[string]mapset.Set[string]{},
		Options:          options,
		CurrentChoices:   []string{},
		MapChoices:       map[string]bool{},
	}
}

// #endregion

func InstallAptCollection(option InstallOptions.Options, libs []string) {
	installSet := mapset.NewSet[string]()
	for _, lib := range libs {
		installSet.Add(lib)
	}
	Ic.AptPreInstallMap[option.String()] = installSet
}
