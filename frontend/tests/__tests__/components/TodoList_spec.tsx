import { fireEvent, render } from '@testing-library/react'

import TodoList from '../../../components/TodoList'
import { Status } from '../../../types'
import { renderWithChakra } from '../../testUtils'

describe('TodoList', () => {
  it('render properly', () => {
    const tree = render(<TodoList data={data} />)

    expect(tree.getByText('First Todo')).toBeTruthy()
    expect(tree.getByText('Second Todo')).toBeTruthy()
  })

  it('should be able to mark as done when status is 0', () => {
    const onDone = jest.fn()

    const tree = renderWithChakra(<TodoList data={data} onToggle={onDone} />)

    fireEvent.click(tree.getAllByText('Done')[0])

    expect(onDone).toHaveBeenCalledWith(data[0], Status.Complete)
  })

  it('should be able to mark as undone when status is 1', () => {
    const onDone = jest.fn()

    const tree = renderWithChakra(<TodoList data={data} onToggle={onDone} />)

    fireEvent.click(tree.getAllByText('Undo')[0])

    expect(onDone).toHaveBeenCalledWith(data[1], Status.Incomplete)
  })

  it('should call onRemove when tap on remove button', () => {
    const onRemove = jest.fn()
    const tree = renderWithChakra(<TodoList data={data} onToggle={() => null} onRemove={onRemove} />)

    fireEvent.click(tree.getAllByText('Remove')[0])

    expect(onRemove).toHaveBeenCalledWith(data[0]._id)
  })
})

const data = [
  {
    _id: 'test1',
    text: 'First Todo',
    status: Status.Incomplete,
    priority: 4,
  },
  {
    _id: 'test2',
    text: 'Second Todo',
    status: Status.Complete,
    priority: 4,
  },
]
