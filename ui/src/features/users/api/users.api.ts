import { api, buildPaginationParams } from '@/lib/api'
import type { PaginatedResponse, BulkCreateResponse, BulkUpdateResponse } from '@/lib/api'
import type {
  User,
  CreateUserInput,
  UpdateUserInput,
  UpdateUserRoleInput,
  UserListParams,
} from '../types'

export const usersApi = {
  /**
   * List users with pagination and filtering
   */
  list: (params: UserListParams = {}) => {
    const queryParams: Record<string, string | number | undefined> = {
      ...buildPaginationParams(params),
      name: params.name,
      email: params.email,
      startCreatedDate: params.startCreatedDate,
      endCreatedDate: params.endCreatedDate,
    }
    return api.get<PaginatedResponse<User>>('/users', queryParams)
  },

  /**
   * Get a single user by ID
   */
  get: (id: string) => api.get<User>(`/users/${id}`),

  /**
   * Create a new user
   */
  create: (data: CreateUserInput) => api.post<User>('/users', data),

  /**
   * Update an existing user
   */
  update: (id: string, data: UpdateUserInput) => api.put<User>(`/users/${id}`, data),

  /**
   * Update user role
   */
  updateRole: (id: string, data: UpdateUserRoleInput) =>
    api.put<User>(`/users/role/${id}`, data),

  /**
   * Delete a user
   */
  delete: (id: string) => api.delete<void>(`/users/${id}`),

  // Bulk operations

  /**
   * Bulk create users (max 100)
   */
  bulkCreate: (items: CreateUserInput[]) =>
    api.post<BulkCreateResponse<User>>('/users/bulk', { items }),

  /**
   * Bulk update users (max 100)
   */
  bulkUpdate: (items: Array<{ id: string; data: UpdateUserInput }>) =>
    api.put<BulkUpdateResponse<User>>('/users/bulk', { items }),

  /**
   * Bulk delete users (max 100)
   */
  bulkDelete: (ids: string[]) => api.delete<{ deleted: number }>('/users/bulk', { ids }),
}
