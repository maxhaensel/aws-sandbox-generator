import { ApolloError } from '@apollo/client/errors'

type input = { email: string; leaseTime: string; cloud: string }

function leaseSandBox(parent: any, { email, leaseTime, cloud }: input) {
  if (email === 'error.mail@pexon-consulting.de') {
    return new ApolloError({
      errorMessage: 'no valid Pexon-Mail',
    })
  }

  if (email === 'aws.mail@pexon-consulting.de') {
    return {
      __typename: 'AwsSandbox',
      id: 'Sandbox 3',
      accountName: 'String',
      assignedUntil: 'string',
      assignedSince: 'string',
      assignedTo: 'string',
    }
  }

  if (email === 'azure.mail@pexon-consulting.de') {
    return {
      __typename: 'AzureSandbox',
      id: 'this is a great id',
      pipelineId: 'string',
      assignedUntil: 'string',
      assignedSince: 'string',
      assignedTo: 'string',
      status: 'string',
      projectId: 'string',
      webUrl: 'string',
    }
  }

  return new ApolloError({
    errorMessage: 'internal Server Error',
  })
}
export { leaseSandBox }
