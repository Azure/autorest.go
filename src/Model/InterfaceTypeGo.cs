// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using System;
using System.Collections.Generic;
using AutoRest.Core.Model;
using AutoRest.Core.Utilities;

namespace AutoRest.Go.Model
{
    /// <summary>
    /// Represents the empty interface type.  It uses the singleton
    /// pattern as there's no need to create more than one of these.
    /// Most of the IModelType methods aren't required which is why
    /// they throw, so in the event that any of them do become
    /// required it should be obvious which one(s).
    /// </summary>
    internal class InterfaceTypeGo : IModelType
    {
        private static InterfaceTypeGo s_Instance = new InterfaceTypeGo();

        public static InterfaceTypeGo Instance  => s_Instance;

        private InterfaceTypeGo() { }

        public Fixable<string> Name => "interface{}";

        public string ExtendedDocumentation => throw new NotImplementedException();

        public string DefaultValue => null;

        public bool IsConstant => false;

        public string DeclarationName => Name.Value;

        public string ClassName => Name.Value;

        public XmlProperties XmlProperties { get => throw new NotImplementedException(); set => throw new NotImplementedException(); }

        public string XmlName => throw new NotImplementedException();

        public string XmlNamespace => throw new NotImplementedException();

        public string XmlPrefix => throw new NotImplementedException();

        public bool XmlIsWrapped => throw new NotImplementedException();

        public bool XmlIsAttribute => throw new NotImplementedException();

        public IEnumerable<IChild> Children => null;

        public CodeModel CodeModel => null;

        public IEnumerable<IIdentifier> IdentifiersInScope => null;

        public IParent Parent => null;

        public HashSet<string> LocallyUsedNames => null;

        public IEnumerable<string> MyReservedNames => null;

        public string Qualifier => "Interface";

        public void Disambiguate()
        {
            // empty
        }

        public bool StructurallyEquals(IModelType other)
        {
            return other is InterfaceTypeGo;
        }
    }
}
