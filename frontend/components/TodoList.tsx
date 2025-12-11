'use client'

import { useState, useEffect } from 'react'
import { todoAPI } from '../lib/api'
import type { Todo } from '../types'

export function TodoList() {
  const [todos, setTodos] = useState<Todo[]>([])
  const [newTodoTitle, setNewTodoTitle] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  // Load todos on mount
  useEffect(() => {
    loadTodos()
  }, [])

  const loadTodos = async () => {
    try {
      setLoading(true)
      setError(null)
      const data = await todoAPI.getTodos()
      setTodos(data)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load todos')
    } finally {
      setLoading(false)
    }
  }

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!newTodoTitle.trim()) return

    try {
      setError(null)
      const newTodo = await todoAPI.createTodo({ title: newTodoTitle })
      setTodos([...todos, newTodo])
      setNewTodoTitle('')
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to create todo')
    }
  }

  const handleDelete = async (id: number) => {
    try {
      setError(null)
      await todoAPI.deleteTodo(id)
      setTodos(todos.filter((todo) => todo.id !== id))
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to delete todo')
    }
  }

  return (
    <div className="w-full max-w-2xl mx-auto p-6">
      <h1 className="text-3xl font-bold mb-6">Todos</h1>

      {/* Error message */}
      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
          {error}
        </div>
      )}

      {/* Create form */}
      <form onSubmit={handleCreate} className="mb-6">
        <div className="flex gap-2">
          <input
            type="text"
            value={newTodoTitle}
            onChange={(e) => setNewTodoTitle(e.target.value)}
            placeholder="Enter todo title..."
            className="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <button
            type="submit"
            disabled={loading}
            className="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:bg-gray-300 disabled:cursor-not-allowed"
          >
            Add
          </button>
        </div>
      </form>

      {/* Loading state */}
      {loading && todos.length === 0 && (
        <div className="text-center py-8 text-gray-500">Loading todos...</div>
      )}

      {/* Empty state */}
      {!loading && todos.length === 0 && (
        <div className="text-center py-8 text-gray-500">
          No todos yet. Create one above!
        </div>
      )}

      {/* Todo list */}
      <div className="space-y-2">
        {todos.map((todo) => (
          <div
            key={todo.id}
            className="flex items-center justify-between p-4 bg-white border border-gray-200 rounded-lg hover:shadow-md transition-shadow"
          >
            <div className="flex-1">
              <h3 className="font-medium text-gray-900">{todo.title}</h3>
              {todo.description && (
                <p className="text-sm text-gray-600 mt-1">{todo.description}</p>
              )}
              <p className="text-xs text-gray-400 mt-1">
                Created: {new Date(todo.created_at).toLocaleString()}
              </p>
            </div>
            <button
              onClick={() => handleDelete(todo.id)}
              className="ml-4 px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 transition-colors"
            >
              Delete
            </button>
          </div>
        ))}
      </div>
    </div>
  )
}
