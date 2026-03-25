import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";
import "@typespec/xml";

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
