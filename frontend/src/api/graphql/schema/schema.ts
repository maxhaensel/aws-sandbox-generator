import gql from 'graphql-tag'
import { loader } from 'graphql.macro'

const typeDefs = loader('../schema.graphql')

const schema = [typeDefs]

export default schema
