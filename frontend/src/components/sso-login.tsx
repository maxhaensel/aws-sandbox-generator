import ssoConsole from '../assets/sso-console.png'

function Component() {
  return (
    <>
      <h1 className="text-2xl">Anleitung</h1>
      <a
        className={`
        inline-flex 
        items-center 
        h-8 
        px-4 
        m-2 
        text-sm 
        text-amber-100 
        transition-colors 
        duration-150 
        bg-amber-700 
        rounded-lg 
        focus:shadow-outline 
        hover:bg-amber-800
        `}
        href="https://pexon.awsapps.com/start#/"
        target="_blank"
        rel="noreferrer"
      >
        Navigate To AWS-SSO-Console
      </a>
      <p>Login with your Pexon-Google-Account and Select provisioned Sandbox</p>
      <img src={ssoConsole} alt="ssoConsole" width={600}></img>
    </>
  )
}

export default Component
