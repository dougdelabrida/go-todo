import { useCallback, useEffect, useState } from 'react'

import { Container, Text } from '@chakra-ui/react'

import AddTodo from '../components/AddTodo'
import TodoList from '../components/TodoList'
import { fetchTodos } from '../data'
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

  useEffect(() => {
    fetchData()
  }, [])

  return (
    <Container>
      <AddTodo />
      <TodoList data={todos} />
      {isLoading && <Text>Loading...</Text>}
    </Container>
  )
}
