const SERVER_URL = 'http://localhost:3000/server'

export const fetchTodos = async () => {
  const response = await fetch(`${SERVER_URL}/todos`)
  return await response.json()
}
