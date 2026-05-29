import React from 'react'
import Actual from './Actual'

describe('<Actual />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<Actual />)
  })
})