type input = { Email: string; Lease_time: string }

function leaseASandBox(parent: any, { Email, Lease_time }: input) {
  console.log(Email, Lease_time)

  if (Email === 'successful.mail@pexon-consulting.de') {
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
    message: 'no Sandbox Available',
    sandbox: {
      account_id: '',
      account_name: '',
      available: '',
      assigned_until: '',
      assigned_since: '',
      assigned_to: '',
    },
  }
}
export { leaseASandBox }
