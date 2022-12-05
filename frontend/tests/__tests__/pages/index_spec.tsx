import { fireEvent, render, waitFor } from '@testing-library/react'
import fetchMock from 'jest-fetch-mock'

import Home from '../../../pages'
import { renderWithChakra } from '../../testUtils'

fetchMock.enableMocks()
fetchMock.dontMock()

jest.mock('../../../data', () => ({
  fetchTodos: jest.fn().mockImplementation(() => {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve(dataMock)
      }, 0)
    })
  }),
  createTodo: jest.fn().mockImplementation((todo) => {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve({ ...todo, _id: (Math.random() + 1).toString(36).substring(7) })
      }, 0)
    })
  }),
  updateTodo: jest.fn().mockImplementation((todo) => {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve({ ...todo })
      }, 0)
    })
  }),
  removeTodo: jest.fn().mockImplementation(() => {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve(null)
      }, 0)
    })
  }),
}))

describe('Index page', () => {
  it('render properly', () => {
    const tree = render(<Home />)
    expect(tree).toBeTruthy()
  })

  it('should fetch data and render todos', async () => {
    const tree = renderWithChakra(<Home />)

    await waitFor(() => {
      expect(tree.getByText('Loading...')).toBeTruthy()
    })

    await waitFor(() => {
      expect(tree.getByText(dataMock[0].text)).toBeTruthy()
      expect(tree.getByText(dataMock[1].text)).toBeTruthy()
      expect(tree.getByText(dataMock[2].text)).toBeTruthy()
      expect(tree.queryByText('Loading...')).toBeFalsy()
    })
  })

  it('should add new todo', async () => {
    const tree = renderWithChakra(<Home />)

    fireEvent.change(tree.getByPlaceholderText('Todo text'), { target: { value: 'New todo' } })

    fireEvent.click(tree.getByText('Create'))

    await waitFor(() => {
      expect(tree.getByText('Loading...')).toBeTruthy()
    })

    await waitFor(() => {
      expect(tree.getByText('New todo')).toBeTruthy()
      expect(tree.queryByText('Loading...')).toBeFalsy()
    })
  })

  it('should mark the first todo as done', async () => {
    const tree = renderWithChakra(<Home />)

    await waitFor(() => {
      expect(tree.getAllByText('Done')[0])
    })

    fireEvent.click(tree.getAllByText('Done')[0])

    await waitFor(() => {
      expect(tree.queryByText('Done')).toBeFalsy()
    })
  })

  it('should delete a todo when click on Remove', async () => {
    const tree = renderWithChakra(<Home />)

    await waitFor(() => {
      expect(tree.getByText(dataMock[0].text))
    })

    fireEvent.click(tree.getAllByText('Remove')[0])

    await waitFor(() => {
      expect(tree.queryByText(dataMock[0].text)).toBeFalsy()
      expect(tree.queryByText('Loading...')).toBeFalsy()
    })
  })
})

const dataMock = [
  { _id: '638c2ac852b3b9e763b7f5f6', text: 'Say what', status: 0, priority: 1 },
  { _id: '638c2b0135fcde43f3a22000', text: 'Looks like it worked', status: 1, priority: 1 },
  { _id: '638ca7a1e9265c7255f4c063', text: 'Check me out', status: 1, priority: 1 },
]
