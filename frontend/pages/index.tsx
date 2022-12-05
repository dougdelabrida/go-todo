import { useCallback, useEffect, useState } from 'react'

import { Container, Text } from '@chakra-ui/react'

import AddTodo, { NewTodo } from '../components/AddTodo'
import TodoList from '../components/TodoList'
import { fetchTodos, createTodo, updateTodo, removeTodo } from '../data'
import { Status, Todo } from '../types'

export default function Home() {
  const [isLoading, setIsLoading] = useState(true)
  const [todos, setTodos] = useState<Todo[]>([])

  const fetchData = useCallback(async () => {
    setIsLoading(true)
    const data = await fetchTodos()
    setTodos(data)
    setIsLoading(false)
  }, [])

  const handleAddTodo = async (todo: NewTodo) => {
    setIsLoading(true)
    try {
      const newTodo = await createTodo(todo)
      setTodos((todos) => [newTodo, ...todos])
    } catch {}
    setIsLoading(false)
  }

  const handleToggleStatus = async (todo: Todo, status: Status) => {
    setIsLoading(true)
    const updatedTodo = await updateTodo({ ...todo, status })
    setTodos((todos) => {
      const newTodos = [...todos]
      const indexToUpdate = newTodos.findIndex((todo) => todo._id === updatedTodo._id)
      newTodos[indexToUpdate] = updatedTodo
      return newTodos
    })
    setIsLoading(false)
  }

  const handleRemove = async (_id: string) => {
    setIsLoading(true)
    await removeTodo(_id)
    setTodos((todos) => todos.filter((todo) => todo._id !== _id))
    setIsLoading(false)
  }

  useEffect(() => {
    fetchData()
  }, [])

  return (
    <Container>
      <AddTodo onAdd={handleAddTodo} />
      <TodoList data={todos} onToggle={handleToggleStatus} onRemove={handleRemove} />
      {isLoading && <Text>Loading...</Text>}
    </Container>
  )
}
