// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package migroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newMultipleInheritanceServiceClient() *MultipleInheritanceServiceClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewMultipleInheritanceServiceClient(pl)
}

// GetCat - Get a cat with name 'Whiskers' where likesMilk, meows, and hisses is true
func TestGetCat(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetCat(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Cat, Cat{
		Hisses:    to.Ptr(true),
		Meows:     to.Ptr(true),
		Name:      to.Ptr("Whiskers"),
		LikesMilk: to.Ptr(true),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetFeline - Get a feline where meows and hisses are true
func TestGetFeline(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetFeline(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Feline, Feline{
		Hisses: to.Ptr(true),
		Meows:  to.Ptr(true),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetHorse - Get a horse with name 'Fred' and isAShowHorse true
func TestGetHorse(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetHorse(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Horse, Horse{
		Name:         to.Ptr("Fred"),
		IsAShowHorse: to.Ptr(true),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetKitten - Get a kitten with name 'Gatito' where likesMilk and meows is true, and hisses and eatsMiceYet is false
func TestGetKitten(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetKitten(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Kitten, Kitten{
		Hisses:      to.Ptr(true),
		Meows:       to.Ptr(true),
		Name:        to.Ptr("Gatito"),
		LikesMilk:   to.Ptr(true),
		EatsMiceYet: to.Ptr(false),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetPet - Get a pet with name 'Peanut'
func TestGetPet(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.GetPet(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Pet, Pet{
		Name: to.Ptr("Peanut"),
	}); r != "" {
		t.Fatal(r)
	}
}

// PutCat - Put a cat with name 'Boots' where likesMilk and hisses is false, meows is true
func TestPutCat(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutCat(context.Background(), Cat{
		Hisses:    to.Ptr(false),
		Meows:     to.Ptr(true),
		Name:      to.Ptr("Boots"),
		LikesMilk: to.Ptr(false),
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr("Cat was correct!")); r != "" {
		t.Fatal(r)
	}
}

// PutFeline - Put a feline who hisses and doesn't meow
func TestPutFeline(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutFeline(context.Background(), Feline{
		Hisses: to.Ptr(true),
		Meows:  to.Ptr(false),
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr("Feline was correct!")); r != "" {
		t.Fatal(r)
	}
}

// PutHorse - Put a horse with name 'General' and isAShowHorse false
func TestPutHorse(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutHorse(context.Background(), Horse{
		Name:         to.Ptr("General"),
		IsAShowHorse: to.Ptr(false),
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr("Horse was correct!")); r != "" {
		t.Fatal(r)
	}
}

// PutKitten - Put a kitten with name 'Kitty' where likesMilk and hisses is false, meows and eatsMiceYet is true
func TestPutKitten(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutKitten(context.Background(), Kitten{
		Hisses:      to.Ptr(false),
		Meows:       to.Ptr(true),
		Name:        to.Ptr("Kitty"),
		LikesMilk:   to.Ptr(false),
		EatsMiceYet: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr("Kitten was correct!")); r != "" {
		t.Fatal(r)
	}
}

// PutPet - Put a pet with name 'Butter'
func TestPutPet(t *testing.T) {
	client := newMultipleInheritanceServiceClient()
	result, err := client.PutPet(context.Background(), Pet{
		Name: to.Ptr("Butter"),
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr("Pet was correct!")); r != "" {
		t.Fatal(r)
	}
}
