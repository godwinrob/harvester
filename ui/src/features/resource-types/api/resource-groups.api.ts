import { api, buildPaginationParams } from '@/lib/api'
import type { PaginatedResponse } from '@/lib/api'
import type { ResourceGroup, ResourceGroupListParams } from '../types'

export const resourceGroupsApi = {
  /**
   * List resource groups with pagination and filtering
   */
  list: (params: ResourceGroupListParams = {}) => {
    const queryParams: Record<string, string | number | undefined> = {
      ...buildPaginationParams(params),
      resourceGroup: params.resourceGroup,
      groupName: params.groupName,
      groupLevel: params.groupLevel,
      containerType: params.containerType,
    }
    return api.get<PaginatedResponse<ResourceGroup>>('/resource-groups', queryParams)
  },

  /**
   * Get a single resource group by key
   */
  get: (resourceGroup: string) =>
    api.get<ResourceGroup>(`/resource-groups/${resourceGroup}`),
}
