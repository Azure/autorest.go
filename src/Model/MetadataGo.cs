using System;
using System.Collections.Generic;
using System.Text;

namespace AutoRest.Go.Model
{
    public class MetadataGo
    {
        public string[] InputFiles { get; }

        public string OutputFolder { get; }

        public string Namespace { get; }

        public MetadataGo(string[] inputFiles, string outputFolder, string ns)
        {
            this.Namespace = ns;
            this.InputFiles = inputFiles;
            this.OutputFolder = outputFolder;
        }
    }
}
