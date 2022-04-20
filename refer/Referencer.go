package refer

// Referencer Helper class that sets and unsets references to components.
var Referencer = &_TReferencer{}

type _TReferencer struct{}

// SetReferencesForOne sets references to specific component.
// To set references components must implement IReferenceable interface.
// If they don't the call to this method has no effect.
//	see IReferenceable
//	Parameters:
//		- references IReferences the references to be set.
//		- component any the component to set references to.
func (c *_TReferencer) SetReferencesForOne(references IReferences, component any) {
	if v, ok := component.(IReferenceable); ok {
		v.SetReferences(references)
	}
}

// SetReferences sets references to multiple components.
// To set references components must implement IReferenceable interface.
// If they don't the call to this method has no effect.
//	see IReferenceable
//	Parameters:
//		- references IReferences the references to be set.
//		- components []any a list of components to set the references to.
func (c *_TReferencer) SetReferences(references IReferences, components []any) {
	for _, component := range components {
		c.SetReferencesForOne(references, component)
	}
}

// UnsetReferencesForOne unsets references in specific component.
// To unset references components must implement IUnreferenceable interface.
// If they don't the call to this method has no effect.
//	see IUnreferenceable
//	Parameters: component any the component to unset references.
func (c *_TReferencer) UnsetReferencesForOne(component any) {
	v, ok := component.(IUnreferenceable)
	if ok {
		v.UnsetReferences()
	}
}

// UnsetReferences unsets references in multiple components.
// To unset references components must implement IUnreferenceable interface.
// If they don't the call to this method has no effect.
//	see IUnreferenceable
//	Parameters: components [] any the list of components, whose references must be cleared.
func (c *_TReferencer) UnsetReferences(components []any) {
	for _, component := range components {
		c.UnsetReferencesForOne(component)
	}
}
