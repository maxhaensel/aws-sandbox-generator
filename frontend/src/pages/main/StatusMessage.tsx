import React from "react"
import { LeaseSandBox_leaseSandBox } from './__generated__/LeaseSandBox'

type A = {
  status: {
    successfully: boolean
    message: string
    sandbox?: LeaseSandBox_leaseSandBox
    display: boolean
  }
}

function StatusMessage(props: A) {
  console.log(props.status.sandbox?.__typename)
  const getSandboxName = () => props.status.sandbox?.__typename === "AwsSandbox" ? props.status.sandbox.accountName : props.status.sandbox?.sandboxName
  return (
    <div
      data-cy="alert"
      className={`
            mt-4 border-2 px-4 py-3 rounded relative
            ${props.status.successfully
          ? 'border-green-500 text-green-800'
          : 'border-rose-500 text-rose-800'
        }
            `}
      role="alert"
    >
      {!props.status.successfully &&
        <>
          <strong className="font-bold">
            Fehler
          </strong>{' '}
          <span className="block sm:inline">{props.status.message}</span>
        </>
      }
      {props.status.successfully &&
        <>
          <strong className="font-bold">
            Erfolgreich!
          </strong>{' '}
          <p>Name: {getSandboxName()}</p>
          <p>Zugewiesen an: {props.status.sandbox?.assignedTo}</p>
          <p>Gültig bis: {props.status.sandbox?.assignedUntil}</p>
          <p>ID: {props.status.sandbox?.id}</p>
        </>
      }
    </div>
  )
}

export { StatusMessage }