/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: GetSandboxes
// ====================================================

export interface GetSandboxes_listSandboxes_sandboxes {
  __typename: "Sandbox";
  account_id: string;
  account_name: string;
  assigned_until: string | null;
  assigned_since: string | null;
  assigned_to: string | null;
}

export interface GetSandboxes_listSandboxes {
  __typename: "ListSandboxesResponse";
  sandboxes: (GetSandboxes_listSandboxes_sandboxes | null)[] | null;
}

export interface GetSandboxes {
  listSandboxes: GetSandboxes_listSandboxes | null;
}
