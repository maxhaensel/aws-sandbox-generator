import React from 'react'
import pexonLogo from './assets/pexon.webp'

function App() {
  const [user, setUser] = React.useState({
    mail: '',
    valid_mail: false,
    lease_time: new Date(new Date().setMonth(new Date().getMonth() + 1))
      .toISOString()
      .slice(0, 10),
    valid_lease_time: true,
  })
  const [status, setStatus] = React.useState({ message: '', display: false })

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
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
  }

  const resetStatus = () => {
    setStatus({ message: '', display: false })
  }

  const submitRequest = () => {
    const URL = process.env.REACT_APP_SANDBOX_DOMAIN_URI

    if (URL === undefined) {
      setStatus({
        message: 'Da ist etwas schief gegangen!',
        display: true,
      })
      return
    }

    const response = fetch(
      `${URL}sandbox?name=${user.mail}&lease_time=${user.lease_time}`,
      {
        method: 'POST',
        mode: 'cors',
        headers: { 'Content-Type': 'application/json' },
      },
    )
    response
      .then(data => data.json())
      .then(data => {
        setStatus({
          message: `${data.message}, ${data.sandbox}`,
          display: true,
        })
      })
      .catch(e => {
        setStatus({
          message: 'Da ist etwas schief gegangen!',
          display: true,
        })
      })

    setUser({
      mail: '',
      valid_mail: false,
      lease_time: '',
      valid_lease_time: false,
    })
  }

  React.useEffect(() => {
    if (status.display === true) {
      const timer = setTimeout(() => {
        resetStatus()
      }, 5000)
      return () => clearTimeout(timer)
    }
  }, [status, setStatus])

  const renderMailInput = () => (
    <>
      <label htmlFor="mail" className="mt-2 text-sm text-gray-500">
        Gib hier deine Pexon-Mail-Adresse ein um eine Sandbox zu provisonieren
      </label>
      <input
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
    <div className="flex justify-center mt-32">
      <div className="m-16">
        <img src={pexonLogo} alt="pexon-logo" width={300}></img>
        {renderMailInput()}
        {renderLeaseInput()}
        <br />
        <button
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
            className="mt-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative"
            role="alert"
          >
            <strong className="font-bold">Holy smokes!</strong>
            <span className="block sm:inline">{status.message}</span>
          </div>
        )}
      </div>
    </div>
  )
}

export default App
