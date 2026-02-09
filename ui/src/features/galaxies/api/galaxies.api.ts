import { api, buildPaginationParams } from '@/lib/api'
import type { PaginatedResponse, BulkCreateResponse, BulkUpdateResponse } from '@/lib/api'
import type {
  Galaxy,
  CreateGalaxyInput,
  UpdateGalaxyInput,
  GalaxyListParams,
} from '../types'

export const galaxiesApi = {
  /**
   * List galaxies with pagination and filtering
   */
  list: (params: GalaxyListParams = {}) => {
    const queryParams: Record<string, string | number | undefined> = {
      ...buildPaginationParams(params),
      name: params.name,
    }
    return api.get<PaginatedResponse<Galaxy>>('/galaxies', queryParams)
  },

  /**
   * Get a single galaxy by ID
   */
  get: (id: string) => api.get<Galaxy>(`/galaxies/${id}`),

  /**
   * Get a galaxy by name
   */
  getByName: (name: string) => api.get<Galaxy>(`/galaxies/name/${name}`),

  /**
   * Create a new galaxy
   */
  create: (data: CreateGalaxyInput) => api.post<Galaxy>('/galaxies', data),

  /**
   * Update an existing galaxy
   */
  update: (id: string, data: UpdateGalaxyInput) => api.put<Galaxy>(`/galaxies/${id}`, data),

  /**
   * Delete a galaxy
   */
  delete: (id: string) => api.delete<void>(`/galaxies/${id}`),

  // Bulk operations

  /**
   * Bulk create galaxies (max 100)
   */
  bulkCreate: (items: CreateGalaxyInput[]) =>
    api.post<BulkCreateResponse<Galaxy>>('/galaxies/bulk', { items }),

  /**
   * Bulk update galaxies (max 100)
   */
  bulkUpdate: (items: Array<{ id: string; data: UpdateGalaxyInput }>) =>
    api.put<BulkUpdateResponse<Galaxy>>('/galaxies/bulk', { items }),

  /**
   * Bulk delete galaxies (max 100)
   */
  bulkDelete: (ids: string[]) => api.delete<{ deleted: number }>('/galaxies/bulk', { ids }),
}
