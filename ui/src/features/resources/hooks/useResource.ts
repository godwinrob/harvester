import { useQuery } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { resourcesApi } from '../api'

/**
 * Hook to fetch a single resource by ID
 */
export function useResource(id: string | undefined) {
  // Use a stable fallback key when no ID is provided to avoid cache pollution
  if (!id) {
    return useQuery({
      queryKey: ['resources', 'detail', '__no-id__'] as const,
      queryFn: () => Promise.reject(new Error('Resource ID is required')),
      enabled: false,
    })
  }

  return useQuery({
    queryKey: queryKeys.resources.detail(id),
    queryFn: () => resourcesApi.get(id),
  })
}
