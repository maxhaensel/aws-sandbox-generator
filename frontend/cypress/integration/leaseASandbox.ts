describe('check leaseASandbox Request Sandbox Button', () => {
  beforeEach(() => {
    cy.visit('localhost:3000')
  })

  it('button should be disabled', () => {
    cy.get('[data-cy=submit-request]').should('be.disabled')
  })

  Cypress._.each(
    [
      { input: '.', disabled: true },
      { input: 'max.muster', disabled: true },
      { input: 'max.muster@some-odd.de', disabled: true },
      { input: 'max.muster@pexon-consulting.de', disabled: false },
    ],
    ({ input, disabled }) => {
      it(`type ${input} as input expect button to be ${
        disabled ? '' : 'not'
      } disabled`, () => {
        cy.get('[data-cy=sandbox-mail-input]').type(input)
        cy.get('[data-cy=submit-request]').should(
          `${disabled ? '' : 'not.'}be.disabled`,
        )
      })
    },
  )
})
describe('validate alert displays shows the correct state', () => {
  beforeEach(() => {
    cy.clock()
    cy.visit('localhost:3000')
  })

  Cypress._.each(
    [
      {
        input: { mail: 'error.mail@pexon-consulting.de', cloud: 'AWS' },
        display: ['Fehler'],
      },
      {
        input: { mail: 'max.mustermann@pexon-consulting.de', cloud: 'AZURE' },
        display: ['Erfolgreich', 'Azure'],
      },
      {
        input: { mail: 'max.mustermann@pexon-consulting.de', cloud: 'AWS' },
        display: ['Erfolgreich', 'AWS'],
      },
    ],
    ({ input, display }) => {
      it(`make request with mail:${input.mail} and cloud:${input.cloud} expect alert to be ${display}`, () => {
        cy.get('[data-cy=sandbox-mail-input]').type(input.mail)

        cy.get('[data-cy=sandbox-cloud-input]').select(input.cloud)

        cy.get('[data-cy=alert]').should('not.exist')
        cy.get('[data-cy=submit-request]').click()
        cy.get('[data-cy=alert]').should('exist')
        display.forEach(element => {
          cy.get('[data-cy=alert]').contains(element)
        })
      })
    },
  )
})

describe('after Submit LeaseTime should be defined', () => {
  const date = new Date('2020-10-20')
  beforeEach(() => {
    cy.clock(date)
    cy.visit('localhost:3000')
  })

  it('button should be disabled', () => {
    cy.get('[data-cy=sandbox-mail-input]').type('test.test@pexon-consulting.de')
    cy.get('[data-cy=lease_time_input]').type('2020-11-21')
    cy.get('[data-cy=submit-request]').click()
    cy.get('[data-cy=lease_time_input]').should('have.value', '2020-11-20')
    cy.tick(6000)
  })
})
