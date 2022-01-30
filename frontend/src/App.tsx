import { Route, Routes } from 'react-router-dom'
import { Layout, Navigation } from './components'
import { Main, Manual, MySandboxes } from './pages'

function App() {
  return (
    <>
      <Navigation></Navigation>
      <Layout>
        <Routes>
          <Route path="/" element={<Main />} />
          <Route path="manual" element={<Manual />} />
          <Route path="sandboxes" element={<MySandboxes />} />
        </Routes>
      </Layout>
    </>
  )
}

export default App
