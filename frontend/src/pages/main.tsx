import { gql, useMutation } from '@apollo/client'
import React from 'react'
import pexonLogo from '../assets/pexon.webp'

function Main() {
  const sandboxType = {
    Azure: 'Azure',
    AWS: 'AWS',
    GCP: 'GCP',
  }
  const [user, setUser] = React.useState({
    mail: '',
    valid_mail: false,
    sandbox_type: sandboxType.Azure,
    lease_time: new Date(new Date().setMonth(new Date().getMonth() + 1))
      .toISOString()
      .slice(0, 10),
    valid_lease_time: true,
  })
  const [status, setStatus] = React.useState({
    message: '',
    display: false,
    successfully: false,
  })

  const [leaseASandBoxRequest, { data, loading, error }] = useMutation(gql`
    mutation LeaseASandBox($email: String!, $lease_time: String!) {
      leaseASandBox(Email: $email, Lease_time: $lease_time) {
        message
        sandbox {
          account_id
          account_name
          assigned_until
          assigned_since
          assigned_to
          available
        }
      }
    }
  `)

  const onChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>,
  ) => {
    if (e.target.id === 'mail') {
      const mail = e.currentTarget.value
      const reg = new RegExp(/\w+\.\w+@pexon-consulting\.de/gm)
      const valid_mail = reg.test(mail)
      setUser(user => ({ ...user, mail, valid_mail }))
    }

    if (e.target.id === 'lease_time') {
      const lease_time = e.currentTarget.value
      const valid_lease_time = lease_time !== '' ? true : false
      setUser(user => ({ ...user, lease_time, valid_lease_time }))
    }

    if (e.target.id === 'sandbox_type') {
      console.log(e.target.value)
    }
  }

  const resetStatus = () => {
    setStatus({ message: '', display: false, successfully: false })
  }

  const submitRequest = () => {
    leaseASandBoxRequest({
      variables: {
        email: user.mail,
        lease_time: user.lease_time,
      },
    })

    setUser({
      mail: '',
      valid_mail: false,
      sandbox_type: sandboxType.Azure,
      lease_time: new Date(new Date().setMonth(new Date().getMonth() + 1))
        .toISOString()
        .slice(0, 10),
      valid_lease_time: false,
    })
  }

  const renderSandboxInput = () => (
    <>
      <form>
        <label htmlFor="sandbox_type" className="mt-2 text-sm text-gray-500">
          Wähle die Art deiner Sandbox aus
          <select
            id="sandbox_type"
            className="shadow appearance-none border rounded w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            onChange={e => onChange(e)}
          >
            <option selected value="Azure">
              Azure
            </option>
            <option value="AWS">AWS</option>
            <option value="GCP">GCP</option>
          </select>
        </label>
      </form>
    </>
  )

  const renderMailInput = () => (
    <>
      <label htmlFor="mail" className="mt-2 text-sm text-gray-500">
        Gib hier deine Pexon-Mail-Adresse ein um eine Sandbox zu provisonieren
      </label>
      <input
        data-cy={'sandbox-mail-input'}
        id="mail"
        className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        onChange={onChange}
        value={user.mail}
      ></input>
    </>
  )

  const renderLeaseInput = () => (
    <>
      <label htmlFor="lease_time" className="mt-2 text-sm text-gray-500">
        Wähle den Zeitraum, maximal 3 Monate
      </label>
      <input
        id="lease_time"
        data-cy="lease_time_input"
        className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        onChange={onChange}
        value={user.lease_time}
        type={'date'}
        min={new Date().toISOString().slice(0, 10)}
        max={new Date(new Date().setMonth(new Date().getMonth() + 3))
          .toISOString()
          .slice(0, 10)}
      ></input>
    </>
  )

  React.useEffect(() => {
    if (status.display === true) {
      const timer = setTimeout(() => {
        resetStatus()
      }, 5000)
      return () => clearTimeout(timer)
    }
  }, [status, setStatus])

  React.useEffect(() => {
    if (loading === false && data) {
      const successfully =
        data.leaseASandBox.message !== 'no Sandbox Available' ? true : false
      setStatus({
        message: `${data.leaseASandBox.message}`,
        display: true,
        successfully,
      })
    }
    if (error) {
      setStatus({
        message: 'Da ist etwas schief gegangen!',
        display: true,
        successfully: false,
      })
    }
  }, [data, loading, error])

  return (
    <div className="flex justify-center mt-32">
      <div className="m-16">
        <img src={pexonLogo} alt="pexon-logo" width={300}></img>
        {renderSandboxInput()}
        {renderMailInput()}
        {renderLeaseInput()}
        <br />
        <button
          data-cy={'submit-request'}
          disabled={!(user.valid_mail && user.valid_lease_time)}
          className={`mt-4 ${
            user.valid_mail && user.valid_lease_time
              ? 'bg-blue-500 hover:bg-blue-700'
              : 'bg-gray-300'
          } text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline`}
          onClick={submitRequest}
        >
          Reqeust Sandbox
        </button>
        {status.display && (
          <div
            data-cy="alert"
            className={`
            mt-4 border-2 px-4 py-3 rounded relative
            ${
              status.successfully
                ? 'border-green-500 text-green-800'
                : 'border-rose-500 text-rose-800'
            }
            `}
            role="alert"
          >
            <strong className="font-bold">
              {status.successfully ? 'Erfolgreich' : 'Fehler'}
            </strong>{' '}
            <span className="block sm:inline">{status.message}</span>
          </div>
        )}
        <div className="mt-8"></div>
      </div>
    </div>
  )
}

export { Main }
