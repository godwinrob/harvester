import { api, buildPaginationParams } from '@/lib/api'
import type { PaginatedResponse, BulkCreateResponse, BulkUpdateResponse } from '@/lib/api'
import type {
  Resource,
  CreateResourceInput,
  UpdateResourceInput,
  ResourceListParams,
} from '../types'

export const resourcesApi = {
  /**
   * List resources with pagination and filtering
   */
  list: (params: ResourceListParams = {}) => {
    const queryParams: Record<string, string | number | undefined> = {
      ...buildPaginationParams(params),
      name: params.name,
      resource_type: params.resourceType,
      resource_group: params.resourceGroup,
      galaxyID: params.galaxyID,
    }
    return api.get<PaginatedResponse<Resource>>('/resources', queryParams)
  },

  /**
   * Get a single resource by ID
   */
  get: (id: string) => api.get<Resource>(`/resources/${id}`),

  /**
   * Get a resource by name
   */
  getByName: (name: string) => api.get<Resource>(`/resources/name/${name}`),

  /**
   * Create a new resource
   */
  create: (data: CreateResourceInput) => api.post<Resource>('/resources', data),

  /**
   * Update an existing resource
   */
  update: (id: string, data: UpdateResourceInput) => api.put<Resource>(`/resources/${id}`, data),

  /**
   * Delete a resource
   */
  delete: (id: string) => api.delete<void>(`/resources/${id}`),

  // Bulk operations

  /**
   * Bulk create resources (max 100)
   */
  bulkCreate: (items: CreateResourceInput[]) =>
    api.post<BulkCreateResponse<Resource>>('/resources/bulk', { items }),

  /**
   * Bulk update resources (max 100)
   */
  bulkUpdate: (items: Array<{ id: string; data: UpdateResourceInput }>) =>
    api.put<BulkUpdateResponse<Resource>>('/resources/bulk', { items }),

  /**
   * Bulk delete resources (max 100)
   */
  bulkDelete: (ids: string[]) => api.delete<{ deleted: number }>('/resources/bulk', { ids }),
}
