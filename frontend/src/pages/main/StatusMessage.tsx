import { LeaseSandBox_leaseSandBox } from './__generated__/LeaseSandBox'

type Props = {
  status: {
    successfully: boolean
    message: string
    sandbox?: LeaseSandBox_leaseSandBox
    display: boolean
  }
}

function StatusMessage({ status: { successfully, message, sandbox } }: Props) {
  const getSandboxName = () =>
    sandbox?.__typename === 'AwsSandbox'
      ? sandbox.accountName
      : sandbox?.sandboxName
  return (
    <div
      data-cy="alert"
      className={`
            mt-4 border-2 px-4 py-3 rounded relative
            ${
              successfully
                ? 'border-green-500 text-green-800'
                : 'border-rose-500 text-rose-800'
            }
            `}
      role="alert"
    >
      {!successfully && (
        <>
          <strong className="font-bold">Fehler</strong>{' '}
          <span className="block sm:inline">{message}</span>
        </>
      )}
      {successfully && (
        <>
          <strong className="font-bold">Erfolgreich!</strong>{' '}
          <p>Name: {getSandboxName()}</p>
          <p>Zugewiesen an: {sandbox?.assignedTo}</p>
          <p>Gültig bis: {sandbox?.assignedUntil}</p>
          <p>ID: {sandbox?.id}</p>
        </>
      )}
    </div>
  )
}

export { StatusMessage }
