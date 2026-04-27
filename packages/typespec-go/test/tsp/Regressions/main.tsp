import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";
import "@typespec/xml";

using Azure.ClientGenerator.Core;
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

@doc("A discriminated base type with no subtypes.")
@discriminator("type")
model DiscriminatedBaseNoSubTypes {
  @doc("The discriminator type.")
  type: string;
  @doc("Some property.")
  name: string;
}

@get
@route("/discriminated-no-subtypes")
op getDiscriminatedNoSubTypes(): DiscriminatedBaseNoSubTypes;

alias ContentXml = {
  @header("Content-Type")
  contentType: "application/xml";
};

model SignedIdentifiers {
  @Xml.unwrapped
  @Xml.name("SignedIdentifier")
  items: Array<SignedIdentifier>;
}

@Xml.name("SignedIdentifier")
model SignedIdentifier {
  @Xml.name("Id") id: string;
}

/**
 * verify we can swap the default definition of SignedIdentifiers with a custom one
 * that maintains the correct XML structure and tags.
 */
@get
@route("/xml-alternate-type-1")
op getXmlOne(): ContentXml & SignedIdentifiers;

@Xml.name("QueueMessagesList")
model ReceivedMessages {
  @Xml.unwrapped
  items: Array<ReceivedMessage>;
}

@Xml.name("QueueMessage")
model ReceivedMessage {
  @Xml.name("MessageId")
  messageId: string;
}

@get
@route("/xml-alternate-type-2")
op getXmlTwo(): ContentXml & ReceivedMessages;

@Xml.name("Queue")
model QueueItem {
  @Xml.name("Name") name: string;
  @Xml.name("Metadata") metadata?: Record<string>;
}

@get
@route("/xml-get-queue")
op getQueue(): ContentXml & QueueItem;

/**
 * ensure that parameter names that start with invalid Go identifiers are properly cased
 */
@post
@route("/fixed-param-name")
op withExpandParam(@query("$expand") $expand: string, @query("$top") $top?: int32): void;

/**
 * ensure that path and query parameters aren't double-decoded by the fake server
 */
@post
@route("/double-decode/{path}")
op doubleDecode(@path path: string, @query query: string): void;

union Color {
  red: "red",
  green: "green",
  blue: "blue",
}

/**
 * extensible enum (i.e. isFixed === false) used to verify that a
 * client-side default value not present in the pre-defined values
 * is still accepted (regression test for issue with property names
 * like `networkApiVersion` whose default is the service API version).
 */
union NetworkVersion {
  string,
  v1: "2020-11-01",
  v2: "2022-11-01",
}

/**
 * path params that aren't string types with client-side defaults
 */
@post
@route("/non-string-path-params/{color}/{count}")
op withClientDefaultValues(
  @Legacy.clientDefaultValue(Color.blue)
  @path
  color: Color,

  @Legacy.clientDefaultValue(123)
  @path
  count: int32,
):void;

model Widget {
  name: string;
  @Legacy.clientDefaultValue(Color.green)
  color: Color;

  /**
   * default value is not one of the pre-defined NetworkVersion values,
   * which is allowed for extensible enums.
   */
  @Legacy.clientDefaultValue("2026-04-15-preview")
  networkVersion?: NetworkVersion;
}

/**
 * model with a client-side default that's an enum
 */
@put
@route("/model-with-default-const")
op withClientDefaultModelField(@body widget: Widget): void;
