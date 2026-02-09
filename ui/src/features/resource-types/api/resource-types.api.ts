import { api, buildPaginationParams } from '@/lib/api'
import type { PaginatedResponse } from '@/lib/api'
import type { ResourceType, ResourceTypeListParams } from '../types'

export const resourceTypesApi = {
  /**
   * List resource types with pagination and filtering
   */
  list: (params: ResourceTypeListParams = {}) => {
    const queryParams: Record<string, string | number | boolean | undefined> = {
      ...buildPaginationParams(params),
      resourceType: params.resourceType,
      resourceTypeName: params.resourceTypeName,
      resourceCategory: params.resourceCategory,
      resourceGroup: params.resourceGroup,
      enterable: params.enterable,
      containerType: params.containerType,
    }
    return api.get<PaginatedResponse<ResourceType>>('/resource-types', queryParams)
  },

  /**
   * Get a single resource type by key
   */
  get: (resourceType: string) =>
    api.get<ResourceType>(`/resource-types/${resourceType}`),
}
