import { Button, HStack, Text, VStack } from '@chakra-ui/react'

import { Status, Todo } from '../types'

export type TodoListProps = {
  onToggle?: (todo: Todo, status: Status) => void
  onRemove?: (_id: string) => void
  data: Todo[]
}

export default function TodoList(props: TodoListProps) {
  return (
    <VStack>
      {props.data.map((todo) => (
        <HStack key={todo._id} width="full" justifyContent="space-between">
          <Text textDecoration={todo.status === Status.Complete ? 'line-through' : 'unset'}>{todo.text}</Text>
          <HStack>
            <Button
              onClick={() =>
                props.onToggle?.(todo, todo.status === Status.Incomplete ? Status.Complete : Status.Incomplete)
              }>
              {todo.status === Status.Incomplete ? 'Mark as Done' : 'Mark as Undone'}
            </Button>
            <Button colorScheme="red" onClick={() => props.onRemove?.(todo._id)}>
              Remove
            </Button>
          </HStack>
        </HStack>
      ))}
    </VStack>
  )
}
