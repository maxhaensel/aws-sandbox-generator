import { Link } from 'react-router-dom'

const navBar = 'bg-cyan-900'
const textColor = 'text-teal-100 hover:text-teal-300'

function Navigation() {
  return (
    <ul className={`flex p-4 ${navBar}`}>
      <li className="mr-6">
        <Link data-cy={'root'} className={textColor} to="/">
          Sandbox Provisionieren
        </Link>
      </li>
      <li className="mr-6">
        <Link data-cy={'sandboxes'} className={textColor} to="/sandboxes">
          Sandboxes verwalten
        </Link>
      </li>
      <li className="mr-6">
        <Link data-cy={'manual'} className={textColor} to="/manual">
          Anleitung
        </Link>
      </li>
    </ul>
  )
}

export { Navigation }
