@doc("Properties of the trusted signing account.")
model BodyRootProperties {
  @doc("The URI of the trusted signing account which is used during signing files.")
  description?: string;

  @doc("Status of the current operation on trusted signing account.")
  @visibility(Lifecycle.Read)
  provisioningState?: ProvisioningState;
}

@doc("The status of the current operation.")
union ProvisioningState {
  string,

  @doc("Resource has been created.")
  Succeeded: "Succeeded",

  @doc("Resource creation failed.")
  Failed: "Failed",

  @doc("Resource creation was canceled.")
  Canceled: "Canceled",

  @doc("Updating in progress.")
  Updating: "Updating",

  @doc("Deletion in progress.")
  Deleting: "Deleting",

  @doc("Resource creation started.")
  Accepted: "Accepted",
}

@doc("The request of the action.")
model ActionRequest {
  @doc("The prop of the action.")
  prop: string;
}
