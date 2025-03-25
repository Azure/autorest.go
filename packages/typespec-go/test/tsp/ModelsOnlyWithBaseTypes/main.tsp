import "@azure-tools/typespec-client-generator-core";

using Azure.ClientGenerator.Core;

@service(#{
  title: "ModelsOnly",
})
namespace ModelsOnly;

model BaseClassThatsPruned {
    id: int32;
}

model DerivedOne extends BaseClassThatsPruned {
    name: string;
}

model PreservedBaseClass {
    id: int32;
}

model DerivedTwo extends PreservedBaseClass {
    name: string;
}

@@access(ModelsOnly.PreservedBaseClass, Access.public);
@@access(ModelsOnly.DerivedOne, Access.public);
@@access(ModelsOnly.DerivedTwo, Access.public);
@@usage(ModelsOnly.PreservedBaseClass, Usage.output);
@@usage(ModelsOnly.DerivedOne, Usage.output);
@@usage(ModelsOnly.DerivedTwo, Usage.output);
