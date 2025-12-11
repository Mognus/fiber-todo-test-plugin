import { api } from '@/lib/api'
import type { Todo, CreateTodoRequest } from '../types'

export const todoAPI = {
  // Get all todos
  getTodos: async (): Promise<Todo[]> => {
    const response = await api.get('/todos')
    return response.data
  },

  // Create new todo
  createTodo: async (data: CreateTodoRequest): Promise<Todo> => {
    const response = await api.post('/todos', data)
    return response.data
  },

  // Delete todo
  deleteTodo: async (id: number): Promise<void> => {
    await api.delete(`/todos/${id}`)
  },
}
