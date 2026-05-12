import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";
import "@typespec/rest";
import "@typespec/http";

using Azure.ResourceManager;
using TypeSpec.Rest;
using TypeSpec.Http;

namespace Microsoft.Test;

@doc("The kind of the pet.")
union PetKind {
  string,

  @doc("A cat.")
  Cat: "Cat",

  @doc("A dog.")
  Dog: "Dog",
}

#suppress "@azure-tools/typespec-azure-core/no-string-discriminator" "Existing"
@doc("Base pet model.")
@discriminator("kind")
model Pet {
  @doc("The kind of the pet.")
  kind: PetKind;

  @doc("The name of the pet.")
  name: string;
}

@doc("A cat.")
model Cat extends Pet {
  @doc("The kind of the pet.")
  kind: PetKind.Cat;

  @doc("Whether the cat purrs.")
  purrs: boolean;
}

@doc("A dog.")
model Dog extends Pet {
  @doc("The kind of the pet.")
  kind: PetKind.Dog;

  @doc("The breed of the dog.")
  breed: string;
}

@doc("Pet resource.")
model PetResource is TrackedResource<PetResourceProperties> {
  @doc("Pet resource name.")
  @path
  @key("petResourceName")
  @segment("petResources")
  name: string;
}

@doc("Properties of the pet resource.")
model PetResourceProperties {
  @doc("Status of the current operation.")
  @visibility(Lifecycle.Read)
  provisioningState?: ProvisioningState;
}

@armResourceOperations
interface Pets {
  @doc("Get a pet resource.")
  get is ArmResourceRead<PetResource>;

  @doc("Adopt a pet using a concrete polymorphic model (Cat) as parameter.")
  @post
  @armResourceAction(PetResource)
  adoptCat(...ResourceInstanceParameters<PetResource>, @bodyRoot cat: Cat): ArmResponse<Cat> | ErrorResponse;

  @doc("Adopt a pet using the interface (discriminated base) as parameter.")
  @post
  @armResourceAction(PetResource)
  adoptPet(...ResourceInstanceParameters<PetResource>, @bodyRoot pet: Pet): ArmResponse<Pet> | ErrorResponse;
}
