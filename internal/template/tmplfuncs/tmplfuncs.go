package tmplfuncs

import (
	"github.com/go-sprout/sprout"
)

// OwnRegistry struct implements the Registry interface, embedding the Handler to access shared functionalities.
type TmplFuncsRegistry struct {
	handler *sprout.Handler
}

// NewRegistry initializes and returns a new instance of your registry.
func NewRegistry() *TmplFuncsRegistry {
	return &TmplFuncsRegistry{}
}

// Uid provides a unique identifier for your registry.
func (or *TmplFuncsRegistry) Uid() string {
	return "tmplFuncs" // Ensure this identifier is unique and uses camelCase
}

// LinkHandler connects the Handler to your registry, enabling runtime functionalities.
func (tfr *TmplFuncsRegistry) LinkHandler(fh sprout.Handler) error {
	tfr.handler = &fh
	return nil
}

// RegisterFunctions adds the provided functions into the given function map.
// This method is called by an Handler to register all functions of a registry.
func (tfr *TmplFuncsRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	// Example of registering a function
	sprout.AddFunction(funcsMap, "hasCommand", tfr.hasCommand)
	sprout.AddFunction(funcsMap, "setVar", tfr.setVar)
	sprout.AddFunction(funcsMap, "getVar", tfr.getVar)
	return nil
}
