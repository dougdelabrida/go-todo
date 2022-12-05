import { useCallback, useEffect, useState } from 'react'

import { Container, Text } from '@chakra-ui/react'

import AddTodo, { NewTodo } from '../components/AddTodo'
import TodoList from '../components/TodoList'
import { fetchTodos, createTodo } from '../data'
import { Todo } from '../types'

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

  useEffect(() => {
    fetchData()
  }, [])

  return (
    <Container>
      <AddTodo onAdd={handleAddTodo} />
      <TodoList data={todos} />
      {isLoading && <Text>Loading...</Text>}
    </Container>
  )
}
