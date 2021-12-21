import React from "react";
import pexonLogo from "./assets/pexon.webp";

function App() {
  const [user, setUser] = React.useState({ mail: "", valid: false });
  const [status, setStatus] = React.useState({ message: "", display: false });

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const mail = e.currentTarget.value;
    const reg = new RegExp(/\w+\.\w+@pexon-consulting\.de/gm);

    const valid = reg.test(mail);

    setUser({ mail, valid });
  };

  const resetStatus = () => {
    setStatus({ message: "", display: false });
  };

  const submitRequest = () => {
    const URL =
      "https://i7bgppkdt9.execute-api.eu-central-1.amazonaws.com/prod/sandbox";
    const response = fetch(`${URL}?name=${user.mail}`, {
      method: "POST",
      mode: "cors",
      headers: { "Content-Type": "application/json" },
    });
    response
      .then((data) => data.json())
      .then((data) => {
        setStatus({
          message: `${data.message}, ${data.sandbox}`,
          display: true,
        });
      })
      .catch((e) => {
        setStatus({
          message: "Da ist etwas schief gegangen!",
          display: true,
        });
      });

    setUser({ mail: "", valid: false });
  };

  React.useEffect(() => {
    if (status.display === true) {
      const timer = setTimeout(() => {
        resetStatus();
      }, 5000);
      return () => clearTimeout(timer);
    }
  }, [status, setStatus]);

  return (
    <div className="flex justify-center mt-32">
      <div className="m-16">
        <img src={pexonLogo} alt="pexon-logo" width={300}></img>

        <input
          id="mail"
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          onChange={onChange}
          value={user.mail}
        ></input>
        <label htmlFor="mail" className="mt-2 text-sm text-gray-500">
          Gebe hier deine Pexon-Mail-Adresse ein um eine Sandbox zu
          provisonieren
        </label>
        <br />
        <button
          disabled={!user.valid}
          className={`mt-4 ${
            user.valid ? "bg-blue-500 hover:bg-blue-700" : "bg-gray-300"
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
  );
}

export default App;
