/** @format */

import { ApolloLink } from '@apollo/client'
import { HttpLink } from '@apollo/client/link/http'
import { onErrorLink } from './errorLink'

const httpLink = new HttpLink({
  uri: process.env.REACT_APP_GRAPHQL_ENDPOINT,
})
// expose Client
export const link = ApolloLink.from([onErrorLink, httpLink])
