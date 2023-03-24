// Main DID Document Constructor Methods
// I.e. Document allows for Reconstruction from Storage of DID Document and Wallet
package types

import "strings"

func (d *DidDocument) ResolveRelationships(vms []VerificationRelationship) *ResolvedDidDocument {
	resolved := &ResolvedDidDocument{
		Id:                   d.Id,
		Context:              d.Context,
		Controller:           d.Controller,
		AlsoKnownAs:          d.AlsoKnownAs,
		VerificationMethod:   d.VerificationMethod,
		Authentication:       []*VerificationRelationship{},
		AssertionMethod:      []*VerificationRelationship{},
		CapabilityDelegation: []*VerificationRelationship{},
		CapabilityInvocation: []*VerificationRelationship{},
		KeyAgreement:         []*VerificationRelationship{},
		Service:              []*Service{},
		Metadata:             d.Metadata,
	}
	return resolved.AddVerificationRelationship(vms)
}

func (d *DidDocument) ResolveMethods(vms []VerificationMethod) *ResolvedDidDocument {
	resolved := &ResolvedDidDocument{
		Id:                   d.Id,
		Context:              d.Context,
		Controller:           d.Controller,
		AlsoKnownAs:          d.AlsoKnownAs,
		VerificationMethod:   d.VerificationMethod,
		Authentication:       []*VerificationRelationship{},
		AssertionMethod:      []*VerificationRelationship{},
		CapabilityDelegation: []*VerificationRelationship{},
		CapabilityInvocation: []*VerificationRelationship{},
		KeyAgreement:         []*VerificationRelationship{},
		Service:              []*Service{},
		Metadata:             d.Metadata,
	}

	// Iterate through VerificationMethod and create a VerificationRelationship for each
	vrs := []VerificationRelationship{}
	for _, vm := range vms {
		vr := VerificationRelationship{
			Reference:          vm.Id,
			VerificationMethod: &vm,
		}
		vrs = append(vrs, vr)
	}
	return resolved.AddVerificationRelationship(vrs)
}

func (r *ResolvedDidDocument) AddVerificationRelationship(vms []VerificationRelationship) *ResolvedDidDocument {
	for _, vm := range vms {
		switch {
		case r.containsRelationship(r.Authentication, vm.Reference):
			r.Authentication = append(r.Authentication, &vm)
		case r.containsRelationship(r.AssertionMethod, vm.Reference):
			r.AssertionMethod = append(r.AssertionMethod, &vm)
		case r.containsRelationship(r.CapabilityDelegation, vm.Reference):
			r.CapabilityDelegation = append(r.CapabilityDelegation, &vm)
		case r.containsRelationship(r.CapabilityInvocation, vm.Reference):
			r.CapabilityInvocation = append(r.CapabilityInvocation, &vm)
		case r.containsRelationship(r.KeyAgreement, vm.Reference):
			r.KeyAgreement = append(r.KeyAgreement, &vm)
		}
	}
	return r
}

func (r *ResolvedDidDocument) containsRelationship(relations []*VerificationRelationship, ref string) bool {
	for _, relation := range relations {
		if relation.Reference == ref {
			return true
		}
	}
	return false
}

// Method returns the DID method of the document
func (d *ResolvedDidDocument) DIDMethod() string {
	return strings.Split(d.Id, ":")[1]
}

// Identifier returns the DID identifier of the document
func (d *ResolvedDidDocument) DIDIdentifier() string {
	return strings.Split(d.Id, ":")[2]
}

// Fragment returns the DID fragment of the document
func (d *ResolvedDidDocument) DIDFragment() string {
	return strings.Split(d.Id, "#")[1]
}