import "@azure-tools/typespec-client-generator-core";
import "./main.tsp";

using Azure.ClientGenerator.Core;
using Emitter.Regressions;

namespace Customizations;

model SignedIdentifiersGo is Array<SignedIdentifier>;

@@alternateType(SignedIdentifiers, SignedIdentifiersGo, "go");
@@clientName(SignedIdentifiersGo, "SignedIdentifiers", "go");

@clientName("Messages", "go")
model ReceivedMessagesGo is Array<ReceivedMessage>;

@@alternateType(ReceivedMessages, ReceivedMessagesGo, "go");
@@clientName(ReceivedMessage, "DequeuedMessage", "go");
