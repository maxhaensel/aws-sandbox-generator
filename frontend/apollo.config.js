// apollo.config.js
module.exports = {
  client: {
    service: {
      name: 'my-service-name',
      localSchemaFile: ['./src/api/graphql/schema.graphql'],
    },
  },
}
