# Getting Started with Create React App

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `yarn start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

### `codegen:gql`

Jedes Mal nachdem man eine GraphQL Query oder Mutation geschrieben hat, sollte man `yarn codegen:gql` ausführen. Dieses Skript erzeugt automatisch TypeScript-Types basierend aus dem Schema.

**Achtung:** Es wird auch `schema_local.graphql` eingelesen, sollte das File Leer sein, sollte es gelöscht werden. Ein leeres File ist nicht möglich.
