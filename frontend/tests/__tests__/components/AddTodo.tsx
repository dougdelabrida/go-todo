import { fireEvent } from '@testing-library/react'

import AddTodo from '../../../components/AddTodo'
import { renderWithChakra } from '../../testUtils'

describe('AddTodo', () => {
  it('should render properly', () => {
    const tree = renderWithChakra(<AddTodo />)

    expect(tree.getByPlaceholderText('Todo text'))
    expect(tree.getByText('Normal'))
    expect(tree.getByText('Create'))
  })

  it('should fill the form and call onAdd', () => {
    const onAdd = jest.fn()
    const tree = renderWithChakra(<AddTodo onAdd={onAdd} />)

    fireEvent.change(tree.getByPlaceholderText('Todo text'), { target: { value: 'My first todo' } })

    fireEvent.click(tree.getByText('Create'))

    expect(onAdd).toHaveBeenCalledWith({ priority: 2, status: 0, text: 'My first todo' })
  })

  it('should not call onAdd if text is empty', () => {
    const onAdd = jest.fn()
    const tree = renderWithChakra(<AddTodo onAdd={onAdd} />)

    fireEvent.click(tree.getByText('Create'))

    expect(onAdd).not.toHaveBeenCalled()
  })
})
