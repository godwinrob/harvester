import { useQuery } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { resourcesApi } from '../api'

/**
 * Hook to fetch a single resource by ID
 */
export function useResource(id: string | undefined) {
  return useQuery({
    queryKey: queryKeys.resources.detail(id!),
    queryFn: () => resourcesApi.get(id!),
    enabled: !!id,
  })
}
