import { render } from '@testing-library/react'

import Home from '../../pages'

describe('Index page', () => {
  it('render properly', () => {
    const tree = render(<Home />)

    expect(tree.getByText('Chakra UI'))
  })
})
