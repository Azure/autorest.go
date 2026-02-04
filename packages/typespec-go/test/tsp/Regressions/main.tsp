import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";

using TypeSpec.Http;

@service(#{
  title: "Various tests for emitter regressions",
})
@server(
    "{endpoint}/supplemental-path",
    "Contains supplemental path with no params",
    {
        endpoint: url,
    }
)
namespace Emitter.Regressions;

model SpreadParams {
  name: string;
  color?: string;
  inner?: InnerSpreadParam;
}

model InnerSpreadParam {
  name: string;
  code: int32;
}

model SomeModel {
  name: string;
}

@route("/spread-with-model")
op spreadWithModel(...SpreadParams): void;

@route("/optional-binary-payload")
op optionalBinaryBody(@body payload?: bytes): void;

@patch
@route("/force-required-body-patch")
op forceRequiredBodyPatch(@body body?: SomeModel): void;

@put
@route("/force-required-body-put")
op forceRequiredBodyPut(@body body?: SomeModel): void;

@post
@route("/optional-body-post")
op optionalBodyPost(@body body?: SomeModel): void;

@get
@route("/binary-response-with-xml-content-type")
op binaryResponseWithXmlContentType():     {
  @header("Content-Type")
  contentType: "application/xml";

  @bodyRoot
  @encode("bytes")
  body: bytes;
};
