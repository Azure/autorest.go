using System;
using System.Collections.Generic;
using System.Text;

namespace AutoRest.Go.Model
{
    public class MetadataGo
    {
        public string[] InputFiles { get; }

        public string OutputFolder { get; }

        public MetadataGo(string[] inputFiles, string outputFolder)
        {
            this.InputFiles = inputFiles;
            this.OutputFolder = outputFolder;
        }
    }
}
