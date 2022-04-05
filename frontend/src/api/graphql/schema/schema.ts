import { loader } from 'graphql.macro'

const typeDefs = loader('../schema.graphql')
const typeDefsLocal = loader('../schema_local.graphql')

const schema = [typeDefs, typeDefsLocal]

export default schema
