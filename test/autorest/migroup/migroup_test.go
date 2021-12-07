// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package migroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newMultipleInheritanceServiceClient() *MultipleInheritanceServiceClient {
	return NewMultipleInheritanceServiceClient(nil)
}

// GetCat - Get a cat with name 'Whiskers' where likesMilk, meows, and hisses is true
func TestGetCat(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetCat(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Cat, Cat{
		Hisses:    to.BoolPtr(true),
		Meows:     to.BoolPtr(true),
		Name:      to.StringPtr("Whiskers"),
		LikesMilk: to.BoolPtr(true),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetFeline - Get a feline where meows and hisses are true
func TestGetFeline(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetFeline(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Feline, Feline{
		Hisses: to.BoolPtr(true),
		Meows:  to.BoolPtr(true),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetHorse - Get a horse with name 'Fred' and isAShowHorse true
func TestGetHorse(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetHorse(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Horse, Horse{
		Name:         to.StringPtr("Fred"),
		IsAShowHorse: to.BoolPtr(true),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetKitten - Get a kitten with name 'Gatito' where likesMilk and meows is true, and hisses and eatsMiceYet is false
func TestGetKitten(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetKitten(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Kitten, Kitten{
		Hisses:      to.BoolPtr(true),
		Meows:       to.BoolPtr(true),
		Name:        to.StringPtr("Gatito"),
		LikesMilk:   to.BoolPtr(true),
		EatsMiceYet: to.BoolPtr(false),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetPet - Get a pet with name 'Peanut'
func TestGetPet(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetPet(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Pet, Pet{
		Name: to.StringPtr("Peanut"),
	}); r != "" {
		t.Fatal(r)
	}
}

// PutCat - Put a cat with name 'Boots' where likesMilk and hisses is false, meows is true
func TestPutCat(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutCat(context.Background(), Cat{
		Hisses:    to.BoolPtr(false),
		Meows:     to.BoolPtr(true),
		Name:      to.StringPtr("Boots"),
		LikesMilk: to.BoolPtr(false),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Value, to.StringPtr("Cat was correct!")); r != "" {
		t.Fatal(r)
	}
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
	if r := cmp.Diff(result.Value, to.StringPtr("Feline was correct!")); r != "" {
		t.Fatal(r)
	}
}

// PutHorse - Put a horse with name 'General' and isAShowHorse false
func TestPutHorse(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutHorse(context.Background(), Horse{
		Name:         to.StringPtr("General"),
		IsAShowHorse: to.BoolPtr(false),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Value, to.StringPtr("Horse was correct!")); r != "" {
		t.Fatal(r)
	}
}

// PutKitten - Put a kitten with name 'Kitty' where likesMilk and hisses is false, meows and eatsMiceYet is true
func TestPutKitten(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutKitten(context.Background(), Kitten{
		Hisses:      to.BoolPtr(false),
		Meows:       to.BoolPtr(true),
		Name:        to.StringPtr("Kitty"),
		LikesMilk:   to.BoolPtr(false),
		EatsMiceYet: to.BoolPtr(true),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Value, to.StringPtr("Kitten was correct!")); r != "" {
		t.Fatal(r)
	}
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
	if r := cmp.Diff(result.Value, to.StringPtr("Pet was correct!")); r != "" {
		t.Fatal(r)
	}
}
