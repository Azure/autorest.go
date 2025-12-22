import "@azure-tools/typespec-client-generator-core";
import "./main.tsp";
using Azure.ClientGenerator.Core;
using Azure.ClientGenerator.Core.Legacy;
using Microsoft.Web;

@@clientName(MSDeployProperties, "MSDeployCore", "go");
@@clientName(Certificate, "AppCertificate", "go");
@@clientName(CertificatePatchResourceProperties, "Certificate", "go");
@@clientName(CertificatePatchResource, "AppCertificatePatchResource", "go");
@@clientName(StaticSiteCustomDomainOverviewARMResource, "StaticSiteCustomDomainRequestPropertiesARMResource", "go");