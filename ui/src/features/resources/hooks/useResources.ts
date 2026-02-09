import { useQuery, keepPreviousData } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { resourcesApi } from '../api'
import type { ResourceListParams } from '../types'

/**
 * Hook to fetch paginated list of resources
 */
export function useResources(params: ResourceListParams = {}) {
  return useQuery({
    queryKey: queryKeys.resources.list(params),
    queryFn: () => resourcesApi.list(params),
    placeholderData: keepPreviousData,
  })
}
