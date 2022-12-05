export enum Status {
  Complete = 1,
  Incomplete = 0,
}

export enum Priority {
  Urgent = 1,
  Normal = 2,
  Low = 3,
}

export type Todo = {
  _id: string
  text: string
  status: Status
  priority: Priority
}
