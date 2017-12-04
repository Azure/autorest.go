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
            FixStutteringTypeNames(cmg);
            TransformEnumTypes(cmg);
            TransformModelTypes(cmg);
            TransformMethods(cmg);
            AzureExtensions.ProcessParameterizedHost(cmg);

            return cmg;
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
            foreach (var mt in cmg.ModelTypes)
            {
                if (mt.IsPolymorphic)
                {
                    var values = new List<EnumValue>();

                    var baseTypeEnumValue = new EnumValue
                    {
                        Name = $"{CodeNamerGo.Instance.GetTypeName(mt.PolymorphicDiscriminator)}{CodeNamerGo.Instance.GetTypeName(mt.SerializedName)}",
                        SerializedName = mt.SerializedName
                    };

                    values.Add(baseTypeEnumValue);

                    foreach (var dt in (mt as CompositeTypeGo).DerivedTypes)
                    {
                        var ev = new EnumValue
                        {
                            Name = $"{CodeNamerGo.Instance.GetTypeName(mt.PolymorphicDiscriminator)}{CodeNamerGo.Instance.GetTypeName(dt.SerializedName)}",
                            SerializedName = dt.SerializedName
                        };
                        values.Add(ev);
                    }
                    bool nameAlreadyExists = cmg.EnumTypes.Any(et => et.Name.EqualsIgnoreCase(mt.PolymorphicDiscriminator));
                    bool alreadyExists = nameAlreadyExists;
                    if (nameAlreadyExists)
                    {
                        (mt as CompositeTypeGo).DiscriminatorEnum = (cmg.EnumTypes.First(et => et.Name.EqualsIgnoreCase(mt.PolymorphicDiscriminator)) as EnumTypeGo);
                        var existingValues = new List<string>();
                        foreach (var v in cmg.EnumTypes.First(et => et.Name.EqualsIgnoreCase(mt.PolymorphicDiscriminator)).Values)
                        {
                            existingValues.Add(v.SerializedName);
                        }
                        foreach (var v in values)
                        {
                            if (!existingValues.Any(ev => ev.Equals(v.SerializedName)))
                            {
                                alreadyExists = false;
                            }
                        }
                    }
                    if (!alreadyExists)
                    {
                        (mt as CompositeTypeGo).DiscriminatorEnum = cmg.Add(New<EnumType>(new{
                            Name = nameAlreadyExists ? $"{mt.PolymorphicDiscriminator}{mt.Name}" :  mt.PolymorphicDiscriminator,
                            Values = values,
                        })) as EnumTypeGo;
                    }
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

            // Then, note each enumerated type with one or more conflicting values and collect the values from
            // those enumerated types without conflicts.  do this on a sorted list to ensure consistent naming
            foreach (var em in cmg.EnumTypes.Cast<EnumTypeGo>().OrderBy(etg => etg.Name.Value))
            {
                if (em.Values.Where(v => topLevelNames.Contains(v.Name) || CodeNamerGo.Instance.UserDefinedNames.Contains(v.Name)).Any())
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

            foreach (var mtm in cmg.ModelTypes)
            {
                if (mtm.IsPolymorphic)
                {
                    foreach (var dt in (mtm as CompositeTypeGo).DerivedTypes)
                    {
                        (dt as CompositeTypeGo).DiscriminatorEnum = (mtm as CompositeTypeGo).DiscriminatorEnum;
                    }
                }
                foreach (var p in mtm.Properties)
                {
                    // Flattened fields are referred to with their type name,
                    // this name change generates correct custom unmarshalers and validation code,
                    // plus flattening does not need to be checked that often
                    if (p.ShouldBeFlattened() && p.ModelType is CompositeTypeGo)
                    {
                        p.Name = p.ModelType.Name;
                    }
                }
            }
        }

        private void TransformMethods(CodeModelGo cmg)
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
                stutteringTypes.ForEach(exported =>
                    {
                        var name = exported is IModelType
                                        ? (exported as IModelType).Name
                                        : (exported as Method).Name;

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
