import React from 'react'
import { render } from '@testing-library/react'
import App from './App'
import { MockedProvider } from '@apollo/client/testing'

import './index.css'

const mocks: any = [] // We'll fill this in next

it('renders learn react link', () => {
  render(
    <MockedProvider mocks={mocks} addTypename={false}>
      <App />
    </MockedProvider>,
  )
})
