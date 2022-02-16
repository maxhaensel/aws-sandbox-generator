/** @format */

import { ApolloClient, InMemoryCache } from '@apollo/client'
import { link } from './links/link-production'

const cache = new InMemoryCache({
  typePolicies: {
    Query: {
      fields: {
        listSandboxes: {
          merge: true,
        },
      },
    },
  },
})

const client = new ApolloClient({
  link,
  cache,
})

export { client, cache }
