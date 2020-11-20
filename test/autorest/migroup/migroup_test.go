// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package migroup

import (
	"context"
	"generatortests/helpers"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newMultipleInheritanceServiceClient() MultipleInheritanceServiceClient {
	return NewMultipleInheritanceServiceClient(NewDefaultConnection(nil))
}

// GetCat - Get a cat with name 'Whiskers' where likesMilk, meows, and hisses is true
func TestGetCat(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetCat(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Cat, &Cat{
		Feline: Feline{
			Hisses: to.BoolPtr(true),
			Meows:  to.BoolPtr(true),
		},
		Pet: Pet{
			Name: to.StringPtr("Whiskers"),
		},
		LikesMilk: to.BoolPtr(true),
	})
}

// GetFeline - Get a feline where meows and hisses are true
func TestGetFeline(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetFeline(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Feline, &Feline{
		Hisses: to.BoolPtr(true),
		Meows:  to.BoolPtr(true),
	})
}

// GetHorse - Get a horse with name 'Fred' and isAShowHorse true
func TestGetHorse(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetHorse(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Horse, &Horse{
		Pet: Pet{
			Name: to.StringPtr("Fred"),
		},
		IsAShowHorse: to.BoolPtr(true),
	})
}

// GetKitten - Get a kitten with name 'Gatito' where likesMilk and meows is true, and hisses and eatsMiceYet is false
func TestGetKitten(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetKitten(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Kitten, &Kitten{
		Cat: Cat{
			Feline: Feline{
				Hisses: to.BoolPtr(true),
				Meows:  to.BoolPtr(true),
			},
			Pet: Pet{
				Name: to.StringPtr("Gatito"),
			},
			LikesMilk: to.BoolPtr(true),
		},
		EatsMiceYet: to.BoolPtr(false),
	})
}

// GetPet - Get a pet with name 'Peanut'
func TestGetPet(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetPet(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Pet, &Pet{
		Name: to.StringPtr("Peanut"),
	})
}

// PutCat - Put a cat with name 'Boots' where likesMilk and hisses is false, meows is true
func TestPutCat(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutCat(context.Background(), Cat{
		Feline: Feline{
			Hisses: to.BoolPtr(false),
			Meows:  to.BoolPtr(true),
		},
		Pet: Pet{
			Name: to.StringPtr("Boots"),
		},
		LikesMilk: to.BoolPtr(false),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr("Cat was correct!"))
}

// PutFeline - Put a feline who hisses and doesn't meow
func TestPutFeline(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutFeline(context.Background(), Feline{
		Hisses: to.BoolPtr(true),
		Meows:  to.BoolPtr(false),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr("Feline was correct!"))
}

// PutHorse - Put a horse with name 'General' and isAShowHorse false
func TestPutHorse(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutHorse(context.Background(), Horse{
		Pet: Pet{
			Name: to.StringPtr("General"),
		},
		IsAShowHorse: to.BoolPtr(false),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr("Horse was correct!"))
}

// PutKitten - Put a kitten with name 'Kitty' where likesMilk and hisses is false, meows and eatsMiceYet is true
func TestPutKitten(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutKitten(context.Background(), Kitten{
		Cat: Cat{
			Feline: Feline{
				Hisses: to.BoolPtr(false),
				Meows:  to.BoolPtr(true),
			},
			Pet: Pet{
				Name: to.StringPtr("Kitty"),
			},
			LikesMilk: to.BoolPtr(false),
		},
		EatsMiceYet: to.BoolPtr(true),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr("Kitten was correct!"))
}

// PutPet - Put a pet with name 'Butter'
func TestPutPet(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutPet(context.Background(), Pet{
		Name: to.StringPtr("Butter"),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr("Pet was correct!"))
}
