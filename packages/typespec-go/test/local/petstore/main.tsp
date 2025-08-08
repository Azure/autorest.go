@service(#{title: "Pet Store"})
namespace PetStore;

@route("/feed")
op feed(food: Food): void;

model Food {
  name: string;
  quantity: int32;
}