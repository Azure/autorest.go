import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";

using TypeSpec.Http;

@service(#{title: "Pet Store"})
namespace PetStore;

@route("/feed")
@post
op feed(@body food: Food): void;

model Food {
  name: string;
  quantity: int32;
}