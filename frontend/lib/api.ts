import { fetcher, mutateFetch } from '@/lib/api/fetcher'
import type { Todo, CreateTodoRequest } from '../types'
import type { ListResponse } from '@/modules/admin-module/frontend/types'

export const todoAPI = {
  // Get all todos
  getTodos: async (): Promise<Todo[]> => {
    const data = await fetcher<ListResponse<Todo>>('/todos')
    return data.items
  },

  // Create new todo
  createTodo: async (data: CreateTodoRequest): Promise<Todo> => {
    return mutateFetch<Todo>('/todos', { method: 'POST', body: data })
  },

  // Delete todo
  deleteTodo: async (id: number): Promise<void> => {
    await mutateFetch<void>(`/todos/${id}`, { method: 'DELETE' })
  },
}
