import { NewTodo } from './components/AddTodo'
import { Todo } from './types'

const SERVER_URL = 'http://localhost:3000/server'

export const fetchTodos = async () => {
  const response = await fetch(`${SERVER_URL}/todos`)
  return await response.json()
}

export const createTodo = async (todo: NewTodo) => {
  const response = await fetch(`${SERVER_URL}/todos`, {
    method: 'POST',
    body: JSON.stringify(todo),
  })
  return await response.json()
}

export const updateTodo = async (todo: Todo) => {
  const response = await fetch(`${SERVER_URL}/todos/${todo._id}`, {
    method: 'PUT',
    body: JSON.stringify(todo),
  })
  return await response.json()
}
