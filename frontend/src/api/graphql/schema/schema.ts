import gql from 'graphql-tag'
import { loader } from 'graphql.macro'

const typeDefs = loader('../schema.graphql')

const extend = gql`
  extend type Mutation {
    deallocateSandbox(Account_id: String!): String
  }
`

const schema = [typeDefs, extend]

export default schema
