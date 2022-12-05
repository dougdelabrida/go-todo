import { Button, HStack, Tag, Text, VStack } from '@chakra-ui/react'

import { Priority, Status, Todo } from '../types'

export type TodoListProps = {
  onToggle?: (todo: Todo, status: Status) => any
  onRemove?: (_id: string) => any
  data: Todo[]
}

const priorityScheme = {
  [Priority.Urgent]: 'red',
  [Priority.Normal]: 'blue',
  [Priority.Low]: 'green',
}

export default function TodoList(props: TodoListProps) {
  return (
    <VStack>
      {props.data.map((todo) => (
        <HStack key={todo._id} width="full" justifyContent="space-between">
          <Text textDecoration={todo.status === Status.Complete ? 'line-through' : 'unset'}>{todo.text}</Text>
          <HStack>
            <Tag colorScheme={priorityScheme[todo.priority]}>{Priority[todo.priority]}</Tag>
            <Button
              onClick={() =>
                props.onToggle?.(todo, todo.status === Status.Incomplete ? Status.Complete : Status.Incomplete)
              }>
              {todo.status === Status.Incomplete ? 'Done' : 'Undo'}
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
