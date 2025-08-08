import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";

using TypeSpec.Http;

@service(#{
  title: "Pet Store Test",
})
namespace PetStore;

model Food {
    name: string;
    quantity: int32;
}

interface Operations {
  @post
  @route("/feed")
  op feed(@body food: Food): void;
}