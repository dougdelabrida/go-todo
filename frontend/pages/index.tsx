import { Container } from '@chakra-ui/react'

import AddTodo from '../components/AddTodo'
import TodoList from '../components/TodoList'

export default function Home() {
  return (
    <Container>
      <AddTodo />
      <TodoList
        data={[
          {
            _id: 'test1',
            text: 'First Todo',
            status: 0,
            priority: 4,
          },
          {
            _id: 'test2',
            text: 'Second Todo',
            status: 1,
            priority: 4,
          },
        ]}
      />
    </Container>
  )
}
