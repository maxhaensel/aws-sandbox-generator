import { gql, useMutation } from '@apollo/client'
import React from 'react'
import { SandboxForm } from './SandboxForm'
import { StatusMessage } from './StatusMessage'
import { LeaseSandBox, LeaseSandBox_leaseSandBox, LeaseSandBoxVariables } from './__generated__/LeaseSandBox'

type State = {
  message: string
  display: boolean
  successfully: boolean
  sandbox?: LeaseSandBox_leaseSandBox
}

function Main() {
  const [status, setStatus] = React.useState<State>({
    message: '',
    display: false,
    successfully: false,
  })

  const [leaseSandBoxRequest, { data, loading, error }] = useMutation<LeaseSandBox>(gql`
    mutation LeaseSandBox(
      $email: String!
      $leaseTime: String!
      $sandbox_type: Cloud!
    ) {
      leaseSandBox(email: $email, leaseTime: $leaseTime, cloud: $sandbox_type) {
        __typename
        ... on CloudSandbox {
          id
          assignedTo
          assignedUntil
          assignedSince
        }
        ... on AzureSandbox {
          sandboxName
        }
        ... on AwsSandbox {
          accountName
        }
      }
    }
  `)

  const submitRequest = async (requestData: LeaseSandBoxVariables) => {
    try {
      await leaseSandBoxRequest({
        variables: requestData,
      })
    } catch (error) {}
  }

  React.useEffect(() => {
    if (loading === false && data) {
      setStatus({
        sandbox: data.leaseSandBox,
        display: true,
        successfully: true,
        message: '',
      })
    }
    if (error) {
      setStatus({
        message: error.message,
        display: true,
        successfully: false,
        sandbox: undefined,
      })
    }
  }, [data, loading, error])
  console.log(error)

  return (
    <div className="flex justify-center mt-32">
      <div className="m-16">
        <SandboxForm submitRequest={submitRequest} />
        {status.display && (
          <StatusMessage status={status} />
        )}
        <div className="mt-8"></div>
      </div>
    </div>
  )
}

export { Main }
