import { useState } from 'react'

import { Button, HStack, Input, Select } from '@chakra-ui/react'

import { Priority, Status, Todo } from '../types'

export type NewTodo = Omit<Todo, '_id'>

export type AddTodoProps = {
  onAdd?: (todo: NewTodo) => any
}

export default function AddTodo(props: AddTodoProps) {
  const [text, setText] = useState('')
  const [priority, setPriority] = useState(Priority.Normal)
  const priorities = Object.entries(Priority).filter(([key]) => isNaN(Number(key)))

  const handleAdd = () => {
    if (!text) return

    props.onAdd?.({
      text,
      status: Status.Incomplete,
      priority,
    })
  }

  return (
    <HStack paddingY={5}>
      <Input placeholder="Todo text" width="full" value={text} onChange={(e) => setText(e.currentTarget.value)} />
      <Select width={44} value={priority} onChange={(e) => setPriority(Number(e.currentTarget.value))}>
        {priorities.map(([label, value]) => (
          <option key={label} value={value}>
            {label}
          </option>
        ))}
      </Select>
      <Button onClick={handleAdd}>Create</Button>
    </HStack>
  )
}
