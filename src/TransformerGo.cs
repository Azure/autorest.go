// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core;
using AutoRest.Core.Logging;
using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Extensions;
using AutoRest.Go.Model;
using AutoRest.Go.Properties;
using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using AutoRest.Extensions.Azure;
using static AutoRest.Core.Utilities.DependencyInjection;

namespace AutoRest.Go
{
    public class TransformerGo : CodeModelTransformer<CodeModelGo>
    {
        private readonly Dictionary<IModelType, IModelType> _normalizedTypes;

        public TransformerGo()
        {
            _normalizedTypes = new Dictionary<IModelType, IModelType>();
        }

        public override CodeModelGo TransformCodeModel(CodeModel cm)
        {
            var cmg = cm as CodeModelGo;

            SwaggerExtensions.ProcessGlobalParameters(cmg);
            // Add the current package name as a reserved keyword
            CodeNamerGo.Instance.ReserveNamespace(cm.Namespace);
            TransformEnumTypes(cmg);
            TransformModelTypes(cmg);
            TransformMethods(cmg);
            AzureExtensions.ProcessParameterizedHost(cmg);
            FixStutteringTypeNames(cmg);
            AssureUniqueNames(cmg);
            TransformPropertyTypes(cmg);

            return cmg;
        }

        private static void TransformPropertyTypes(CodeModelGo cmg)
        {
            foreach (var model in cmg.ModelTypes)
            {
                foreach (var property in model.Properties)
                {
                    // Flattened fields are referred to with their type name,
                    // this name change generates correct custom unmarshalers and validation code,
                    // plus flattening does not need to be checked that often
                    if (property.ShouldBeFlattened() && property.ModelType is CompositeTypeGo)
                    {
                        property.Name = property.ModelType.HasInterface() ? property.ModelType.GetInterfaceName() : property.ModelType.Name.Value;
                    }
                }
            }
        }

        private void TransformEnumTypes(CodeModelGo cmg)
        {
            // fix up any enum types that are missing a name.
            // NOTE: this must be done before the next code block
            foreach (var mt in cmg.ModelTypes)
            {
                foreach (var property in mt.Properties)
                {
                    // gosdk: For now, inherit Enumerated type names from the composite type field name
                    if (property.ModelType is EnumTypeGo)
                    {
                        var enumType = property.ModelType as EnumTypeGo;

                        if (!enumType.IsNamed)
                        {
                            enumType.SetName(property.Name);
                        }
                    }
                }
            }

            // create a "none" enum value for all enum types
            foreach (var et in cmg.EnumTypes)
            {
                var e = et as EnumTypeGo;
                var ev = new EnumValueGo();
                ev.Name = "None";
                ev.Description = $"{EnumValueGo.FormatName(e, ev)} represents an empty {e.Name}.";
                e.Values.Add(ev);
            }

            // And add any others with a defined name and value list (but not already located)
            foreach (var mt in cmg.ModelTypes)
            {
                var namedEnums = mt.Properties.Where(p => p.ModelType is EnumTypeGo && (p.ModelType as EnumTypeGo).IsNamed);
                foreach (var p in namedEnums)
                {
                    if (!cmg.EnumTypes.Any(etm => etm.Equals(p.ModelType)))
                    {
                        cmg.Add(new EnumTypeGo(p.ModelType as EnumType));
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
        }

        private static void AssureUniqueNames(CodeModelGo cmg)
        {
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

        private void TransformModelTypes(CodeModelGo cmg)
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

            // Find all methods that returned paged results

            cmg.Methods.Cast<MethodGo>()
                .Where(m => m.IsPageable).ToList()
                .ForEach(m =>
                {
                    if (!cmg.PagedTypes.ContainsKey(m.ReturnValue().Body))
                    {
                        cmg.PagedTypes.Add(m.ReturnValue().Body, m.NextLink);
                    }

                    if (!m.NextMethodExists(cmg.Methods.Cast<MethodGo>()))
                    {
                        cmg.NextMethodUndefined.Add(m.ReturnValue().Body);
                    }
                });

            // Mark all models returned by one or more methods and note any "next link" fields used with paged data
            cmg.ModelTypes.Cast<CompositeTypeGo>()
                .Where(mtm =>
                {
                    return cmg.Methods.Cast<MethodGo>().Any(m => m.HasReturnValue() && m.ReturnValue().Body.Equals(mtm));
                }).ToList()
                .ForEach(mtm =>
                {
                    mtm.IsResponseType = true;
                    if (cmg.PagedTypes.ContainsKey(mtm))
                    {
                        mtm.NextLink = CodeNamerGo.Instance.GetPropertyName(cmg.PagedTypes[mtm]);
                        mtm.PreparerNeeded = cmg.NextMethodUndefined.Contains(mtm);
                    }
                });

            foreach (var mtm in cmg.ModelTypes)
            {
                if (mtm.IsPolymorphic)
                {
                    foreach (var dt in (mtm as CompositeTypeGo).DerivedTypes)
                    {
                        (dt as CompositeTypeGo).DiscriminatorEnum = (mtm as CompositeTypeGo).DiscriminatorEnum;
                    }

                }
            }
        }

        private void TransformMethods(CodeModelGo cmg)
        {
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
                    // method returns a primitive type, wrap it in a composite type
                    var ctg = new CompositeTypeGo(method);
                    ctg.IsResponseType = true;
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
                else if (!method.HasReturnValue() && method.ReturnType.Headers != null)
                {
                    // method has no return body but does return values via headers.  generate a
                    // wrapper type for it so we'll get convenience methods for the header values
                    var ctg = new CompositeTypeGo($"{method.MethodGroup.Name}{method.Name}Response");
                    ctg.IsResponseType = true;
                    cmg.Add(ctg);
                    method.ReturnType = new Response(ctg, method.ReturnType.Headers);
                }
            }

            // do this after transforming methods as the creation of synthetic
            // types can have an impact on the group transformations.
            foreach (var mg in cmg.MethodGroups)
            {
                mg.Transform(cmg);
            }
        }

        private void FixStutteringTypeNames(CodeModelGo cmg)
        {
            // Trim the package name from exported types; append a suitable qualifier, if needed, to avoid conflicts.
            var exportedTypes = new HashSet<object>();
            exportedTypes.UnionWith(cmg.EnumTypes);
            exportedTypes.UnionWith(cmg.Methods);
            exportedTypes.UnionWith(cmg.ModelTypes);

            var stutteringTypes = exportedTypes
                                    .Where(exported =>
                                        (exported is IModelType && (exported as IModelType).Name.FixedValue.StartsWith(cmg.Namespace, StringComparison.OrdinalIgnoreCase)) ||
                                        (exported is Method && (exported as Method).Name.FixedValue.StartsWith(cmg.Namespace, StringComparison.OrdinalIgnoreCase)));

            if (stutteringTypes.Any())
            {
                Logger.Instance.Log(Category.Warning, string.Format(CultureInfo.InvariantCulture, Resources.NamesStutter, stutteringTypes.Count()));
                stutteringTypes.ForEach(exported =>
                    {
                        var name = exported is IModelType
                                        ? (exported as IModelType).Name
                                        : (exported as Method).Name;

                        Logger.Instance.Log(Category.Warning, string.Format(CultureInfo.InvariantCulture, Resources.StutteringName, name));

                        name = name.Value.TrimPackageName(cmg.Namespace);

                        var nameInUse = exportedTypes
                                            .Any(et => (et is IModelType && (et as IModelType).Name.Equals(name)) || (et is Method && (et as Method).Name.Equals(name)));
                        if (exported is EnumType)
                        {
                            (exported as EnumType).Name.FixedValue = CodeNamerGo.AttachTypeName(name, cmg.Namespace, nameInUse, "Enum");
                        }
                        else if (exported is CompositeType)
                        {
                            (exported as CompositeType).Name.FixedValue = CodeNamerGo.AttachTypeName(name, cmg.Namespace, nameInUse, "Type");
                        }
                        else if (exported is Method)
                        {
                            (exported as Method).Name.FixedValue = CodeNamerGo.AttachTypeName(name, cmg.Namespace, nameInUse, "Method");
                        }
                    });
            }
        }
    }
}
