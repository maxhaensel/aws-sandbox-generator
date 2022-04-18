import React from 'react'
import pexonLogo from 'assets/pexon.webp'
import { LeaseSandBoxVariables } from './__generated__/LeaseSandBox'
import { Cloud } from 'types/globalTypes'

function SandboxForm(props: {
  submitRequest: (requestData: LeaseSandBoxVariables) => void
}) {
  const [user, setUser] = React.useState({
    mail: '',
    valid_mail: false,
    sandbox_type: Cloud.AZURE,
    lease_time: new Date(new Date().setMonth(new Date().getMonth() + 1))
      .toISOString()
      .slice(0, 10),
    valid_lease_time: true,
  })

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
      if (e.currentTarget.value === 'AWS') {
        const sandbox_type = Cloud.AWS
        setUser(user => ({ ...user, sandbox_type }))
      }

      if (e.currentTarget.value === 'AZURE') {
        const sandbox_type = Cloud.AZURE
        setUser(user => ({ ...user, sandbox_type }))
      }
    }
  }

  const submitRequest = () => {
    props.submitRequest({
      email: user.mail,
      leaseTime: user.lease_time,
      sandbox_type: user.sandbox_type,
    })
    setUser({
      mail: '',
      valid_mail: false,
      sandbox_type: Cloud.AZURE,
      lease_time: new Date(new Date().setMonth(new Date().getMonth() + 1))
        .toISOString()
        .slice(0, 10),
      valid_lease_time: true,
    })
  }

  const renderSandboxInput = () => (
    <form>
      <label htmlFor="sandbox_type" className="mt-2 text-sm text-gray-500">
        Wähle die Art deiner Sandbox aus
        <select
          data-cy={'sandbox-cloud-input'}
          id="sandbox_type"
          className="shadow appearance-none border rounded w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          defaultValue={Cloud.AZURE}
          onChange={onChange}
        >
          <option value={Cloud.AZURE}>Azure</option>
          <option value={Cloud.AWS}>AWS</option>
        </select>
      </label>
    </form>
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

  return (
    <>
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
        Request Sandbox
      </button>
    </>
  )
}

export { SandboxForm }
