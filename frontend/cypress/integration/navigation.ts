describe('navigation', () => {
  beforeEach(() => {
    cy.visit('localhost:3000')
  })

  Cypress._.each(
    [
      { domElement: 'root', urlPath: '/' },
      { domElement: 'sandboxes', urlPath: '/sandboxes' },
      { domElement: 'manual', urlPath: '/manual' },
    ],
    ({ domElement, urlPath }) => {
      it(`click on navigation ${domElement} expect urlPath to be "${urlPath}"`, () => {
        cy.get(`[data-cy=${domElement}]`).click()
        cy.url().should('include', urlPath)
      })
    },
  )
})
