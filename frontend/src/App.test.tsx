import React from 'react'
import { render } from '@testing-library/react'
import App from './App'
import { MockedProvider } from '@apollo/client/testing'

const mocks: any = [] // We'll fill this in next

test('renders learn react link', () => {
  render(
    <MockedProvider mocks={mocks} addTypename={false}>
      <App />
    </MockedProvider>,
  )
})

test('renders learn react link2', () => {
  render(<div>test</div>)
})
