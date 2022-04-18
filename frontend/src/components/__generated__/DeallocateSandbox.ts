/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: DeallocateSandbox
// ====================================================

export interface DeallocateSandbox_deallocateSandbox {
  __typename: "LeaseASandBoxResponse";
  message: string;
}

export interface DeallocateSandbox {
  deallocateSandbox: DeallocateSandbox_deallocateSandbox | null;
}

export interface DeallocateSandboxVariables {
  Account_id: string;
}
