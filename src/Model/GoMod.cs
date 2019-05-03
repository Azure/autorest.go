// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

namespace AutoRest.Go.Model
{
    public class GoMod
    {
        public GoMod(string goMod)
        {
            Module = goMod;
        }

        public string Module { get; }
    }
}
