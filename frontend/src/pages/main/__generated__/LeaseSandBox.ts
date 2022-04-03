/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { Cloud } from "./../../../types/globalTypes";

// ====================================================
// GraphQL mutation operation: LeaseSandBox
// ====================================================

export interface LeaseSandBox_leaseSandBox_AzureSandbox {
  __typename: "AzureSandbox";
  id: string;
  assignedTo: string;
  assignedUntil: string;
  assignedSince: string;
  sandboxName: string;
}

export interface LeaseSandBox_leaseSandBox_AwsSandbox {
  __typename: "AwsSandbox";
  id: string;
  assignedTo: string;
  assignedUntil: string;
  assignedSince: string;
  accountName: string;
}

export type LeaseSandBox_leaseSandBox = LeaseSandBox_leaseSandBox_AzureSandbox | LeaseSandBox_leaseSandBox_AwsSandbox;

export interface LeaseSandBox {
  leaseSandBox: LeaseSandBox_leaseSandBox;
}

export interface LeaseSandBoxVariables {
  email: string;
  leaseTime: string;
  sandbox_type: Cloud;
}
