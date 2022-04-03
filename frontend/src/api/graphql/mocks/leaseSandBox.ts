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
      id: 'Sandbox ID',
      accountName: 'AWS Account Name',
      assignedUntil: '10.05.2022',
      assignedSince: 'string',
      assignedTo: 'aws.mail@pexon-consulting.de',
    }
  }

  if (email === 'azure.mail@pexon-consulting.de') {
    return {
      __typename: 'AzureSandbox',
      id: 'Sandbox ID',
      assignedUntil: '10.05.2022',
      assignedSince: 'string',
      assignedTo: 'azure.mail@pexon-consulting.de',
      sandboxName: 'Azure Subscription Name',
    }
  }

  return new ApolloError({
    errorMessage: 'internal Server Error',
  })
}
export { leaseSandBox }
