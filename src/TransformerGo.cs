// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core;
using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Extensions;
using AutoRest.Go.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using static AutoRest.Core.Utilities.DependencyInjection;

namespace AutoRest.Go
{
    public class TransformerGo : CodeModelTransformer<CodeModelGo>
    {
        public override CodeModelGo TransformCodeModel(CodeModel cm)
        {
            var cmg = cm as CodeModelGo;

            SwaggerExtensions.ProcessGlobalParameters(cmg);
            TransformEnumTypes(cmg);
            TransformModelTypes(cmg);
            TransformMethods(cmg);
            SwaggerExtensions.ProcessParameterizedHost(cmg);
            FixStutteringTypeNames(cmg);

            return cmg;
        }

        private static void TransformEnumTypes(CodeModelGo cmg)
        {
            // fix up any enum types that are missing a name.
            // NOTE: this must be done before the next code block
            foreach (var mt in cmg.ModelTypes)
            {
                foreach (var property in mt.Properties)
                {
                    // gosdk: For now, inherit Enumerated type names from the composite type field name
                    if (property.ModelType is EnumTypeGo enumType)
                    {
                        if (!enumType.IsNamed)
                        {
                            enumType.SetName(property.Name);
                        }
                    }
                }
            }

            // And add any others with a defined name and value list (but not already located)
            foreach (var mt in cmg.ModelTypes)
            {
                var namedEnums = mt.Properties.Where(p => p.ModelType is EnumTypeGo enumType && enumType.IsNamed);
                foreach (var p in namedEnums)
                {
                    if (!cmg.EnumTypes.Any(etm => etm.Equals(p.ModelType)))
                    {
                        cmg.Add(p.ModelType as EnumType);
                    }
                };
            }

            // Add discriminators
            // For all polymorphic types we need an enum with values for all the implementing types.
            // To do this:
            // 1. Create an enum with the basic type and all derived types as values.
            // 2. Check if there is any enum already present that has the same name. If there is, check if it contains all the values of current enum.
            // 3. If it has the same name and all values of current enum, use it. Otherwise create a new one and use it.
            foreach (var mt in cmg.ModelTypes.Cast<CompositeTypeGo>())
            {
                if (!mt.IsPolymorphic)
                {
                    continue;
                }

                var enumValues = new List<EnumValue>();

                var baseTypeEnumValue = new EnumValue
                {
                    Name = $"{CodeNamerGo.Instance.GetTypeName(mt.PolymorphicDiscriminator)}{CodeNamerGo.Instance.GetTypeName(mt.SerializedName)}",
                    SerializedName = mt.SerializedName
                };

                enumValues.Add(baseTypeEnumValue);

                foreach (var dt in mt.DerivedTypes)
                {
                    var ev = new EnumValue
                    {
                        Name = $"{CodeNamerGo.Instance.GetTypeName(mt.PolymorphicDiscriminator)}{CodeNamerGo.Instance.GetTypeName(dt.SerializedName)}",
                        SerializedName = dt.SerializedName
                    };
                    enumValues.Add(ev);
                }

                var enumAlreadyExists = false;
                var enumWithSameName = (EnumTypeGo)cmg.EnumTypes.FirstOrDefault(et => et.Name.EqualsIgnoreCase(CodeNamerGo.Instance.GetTypeName(mt.PolymorphicDiscriminator)));

                if (enumWithSameName != null)
                {
                    enumAlreadyExists = !enumValues.Select(value => value.SerializedName).Except(enumWithSameName.Values.Select(value => value.SerializedName)).Any();
                }

                if (enumAlreadyExists)
                {
                    mt.DiscriminatorEnum = enumWithSameName;
                }
                else
                {
                    mt.DiscriminatorEnum = cmg.Add(New<EnumType>(new
                    {
                        Name = enumWithSameName == null ? mt.PolymorphicDiscriminator : $"{mt.PolymorphicDiscriminator}{mt.GetInterfaceName()}",
                        Values = enumValues,
                    })) as EnumTypeGo;
                }
            }

            // now normalize the names
            // NOTE: this must be done after all enum types have been accounted for
            foreach (var enumType in cmg.EnumTypes)
            {
                enumType.SetName(CodeNamerGo.Instance.GetTypeName(enumType.Name.FixedValue));
                foreach (var v in enumType.Values)
                {
                    v.Name = CodeNamerGo.Instance.GetEnumMemberName(v.Name);
                }
            }

            // Ensure all enumerated type values have the simplest possible unique names
            // -- The code assumes that all public type names are unique within the client and that the values
            //    of an enumerated type are unique within that type. To safely promote the enumerated value name
            //    to the top-level, it must not conflict with other existing types. If it does, prepending the
            //    value name with the (assumed to be unique) enumerated type name will make it unique.

            // First, collect all type names (since these cannot change)
            var topLevelNames = new HashSet<string>();
            foreach (var mt in cmg.ModelTypes)
            {
                topLevelNames.Add(mt.Name);
            }

            foreach (var em in cmg.EnumTypes)
            {
                topLevelNames.Add(em.Name);
            }

            // Then, note each enumerated type with one or more conflicting values and collect the values from
            // those enumerated types without conflicts.  do this on a sorted list to ensure consistent naming
            foreach (var em in cmg.EnumTypes.Cast<EnumTypeGo>().OrderBy(etg => etg.Name.Value))
            {
                if (em.Values.Any(v => topLevelNames.Contains(v.Name) || CodeNamerGo.Instance.UserDefinedNames.Contains(v.Name)))
                {
                    foreach (var v in em.Values)
                    {
                        v.Name = em.Name + v.Name;
                    }
                }
                else
                {
                    topLevelNames.UnionWith(em.Values.Select(ev => ev.Name));
                }
            }

            var modelList = new List<EnumValue>();
            foreach (var em in cmg.EnumTypes.OrderBy(etg => etg.Name.Value))
            {
                foreach (var v in em.Values)
                {
                    v.Name = CodeNamerGo.Instance.GetUnique(v.Name, v, cmg.ModelTypes, modelList);
                    modelList.Add(v);
                }
            }
        }

        private static void TransformModelTypes(CodeModelGo cmg)
        {
            foreach (var ctg in cmg.ModelTypes.Cast<CompositeTypeGo>())
            {
                var name = ctg.Name.FixedValue.TrimPackageName(cmg.Namespace);

                // ensure that the candidate name isn't already in use
                if (name != ctg.Name && cmg.ModelTypes.Any(mt => mt.Name == name))
                {
                    name = $"{name}Type";
                }

                if (CodeNamerGo.Instance.UserDefinedNames.Contains(name))
                {
                    name = $"{name}{cmg.Namespace.Capitalize()}";
                }

                ctg.SetName(name);
            }

            // Mark all models returned by one or more methods
            cmg.ModelTypes.Cast<CompositeTypeGo>()
                .Where(mtm =>
                {
                    return cmg.Methods.Cast<MethodGo>().Any(m => m.HasReturnValue() && m.ReturnValue().Body.Equals(mtm));
                })
                .ForEach(mtm =>
                {
                    mtm.IsResponseType = true;
                });

            foreach (var mtm in cmg.ModelTypes.Cast<CompositeTypeGo>())
            {
                if (mtm.IsPolymorphic)
                {
                    foreach (var dt in mtm.DerivedTypes)
                    {
                        ((CompositeTypeGo)dt).DiscriminatorEnum = mtm.DiscriminatorEnum;
                    }
                }
            }
        }

        private static void TransformMethods(CodeModelGo cmg)
        {
            foreach (var mg in cmg.MethodGroups)
            {
                mg.Transform(cmg);
            }

            var wrapperTypes = new Dictionary<string, CompositeTypeGo>();
            foreach (var method in cmg.Methods.Cast<MethodGo>())
            {
                method.Transform(cmg);

                var scope = new VariableScopeProvider();
                foreach (var parameter in method.Parameters)
                {
                    parameter.Name = scope.GetVariableName(parameter.Name);
                }

                // fix up method return types
                if (method.ReturnType.Body.ShouldBeSyntheticType())
                {
                    var ctg = new CompositeTypeGo(method.ReturnType.Body);
                    if (wrapperTypes.ContainsKey(ctg.Name))
                    {
                        method.ReturnType = new Response(wrapperTypes[ctg.Name], method.ReturnType.Headers);
                    }
                    else
                    {
                        wrapperTypes.Add(ctg.Name, ctg);
                        cmg.Add(ctg);
                        method.ReturnType = new Response(ctg, method.ReturnType.Headers);
                    }
                }

                if (method.IsPageable && !method.IsNextMethod)
                {
                    // for pageable methods replace the return type with a page iterator.
                    // do this before LROs as you can have pageable operations that are
                    // long-running and for this case we want the future to return a paged type.
                    // note that we don't want to do this for the "next methods" that are
                    // defined explicitly in swagger as they will be used in lieu of
                    // generating a custom preparer so they must return the underlying type.
                    cmg.CreatePageableTypeForMethod(method);
                }

                if (method.IsLongRunningOperation())
                {
                    // for LROs we replace the return type with a future that
                    // knows how to poll for the operation's status and result
                    cmg.CreateFutureTypeForMethod(method);
                }
            }
        }

        private static void FixStutteringTypeNames(CodeModelGo cmg)
        {
            // Trim the package name from exported types; append a suitable qualifier, if needed, to avoid conflicts.
            var exportedTypes = new HashSet<object>();
            exportedTypes.UnionWith(cmg.EnumTypes);
            exportedTypes.UnionWith(cmg.Methods);
            exportedTypes.UnionWith(cmg.ModelTypes);

            var stutteringTypes = exportedTypes
                                    .Where(exported =>
                                        (exported is IModelType modelType && modelType.Name.FixedValue.StartsWith(cmg.Namespace, StringComparison.OrdinalIgnoreCase)) ||
                                        (exported is Method method && method.Name.FixedValue.StartsWith(cmg.Namespace, StringComparison.OrdinalIgnoreCase)));

            if (stutteringTypes.Any())
            {
                stutteringTypes.ForEach(exported =>
                    {
                        var name = exported is IModelType type
                                        ? type.Name
                                        : ((Method)exported).Name;

                        name = name.Value.TrimPackageName(cmg.Namespace);

                        var nameInUse = exportedTypes
                                            .Any(et => (et is IModelType modeltype && modeltype.Name.Equals(name)) || (et is Method methodType && methodType.Name.Equals(name)));
                        if (exported is EnumType enumType)
                        {
                            enumType.Name.Value = CodeNamerGo.AttachTypeName(name, cmg.Namespace, nameInUse, "Enum");
                        }
                        else if (exported is CompositeType compositeType)
                        {
                            compositeType.Name.Value = CodeNamerGo.AttachTypeName(name, cmg.Namespace, nameInUse, "Type");
                        }
                        else if (exported is Method methodType)
                        {
                            methodType.Name.Value = CodeNamerGo.AttachTypeName(name, cmg.Namespace, nameInUse, "Method");
                        }
                    });
            }
        }
    }
}
