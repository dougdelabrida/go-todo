import { Container, Text } from '@chakra-ui/react'

import TodoList from '../components/TodoList'

export default function Home() {
  return (
    <Container>
      <Text>Chakra UI</Text>
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
