type input = { email: string; leaseTime: string; cloud: string }

function leaseSandBox(parent: any, { email, leaseTime, cloud }: input) {
  if (email === 'successful.mail@pexon-consulting.de') {
    return {
      message: 'Sandbox is provided',
      sandbox: {
        account_id: 'Sandbox 3',
        account_name: 'String',
        available: 'String',
        assigned_until: 'String',
        assigned_since: 'String',
        assigned_to: 'String',
      },
    }
  }
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
export { leaseSandBox }
