/** @format */

import { ApolloLink, Observable } from '@apollo/client'
import { SchemaLink } from '@apollo/client/link/schema'
import { makeExecutableSchema } from '@graphql-tools/schema'

import typeDefs from '../schema/schema'

import { listSandboxes, leaseASandBox, deallocateSandbox } from '../mocks'

function getRandomArbitrary(min: number, max: number, biased: number) {
  const result = Math.random() >= biased
  const normal = Math.random() * (max - min) + min
  return result ? min : normal
}

const delayLink = new ApolloLink((operation, forward) => {
  const chainObservable = forward(operation) // observable for remaining link chain
  return new Observable(observer => {
    setTimeout(() => {
      chainObservable.subscribe({
        next: observer.next.bind(observer),
      })
    }, getRandomArbitrary(200, 3000, 0.2))
  })
})

const resolvers = {
  Query: { listSandboxes },
  Mutation: { leaseASandBox, deallocateSandbox },
}

const executableSchema = makeExecutableSchema({
  typeDefs,
  resolvers,
})

export const link = ApolloLink.from([
  // delayLink,
  new SchemaLink({ schema: executableSchema }),
])
