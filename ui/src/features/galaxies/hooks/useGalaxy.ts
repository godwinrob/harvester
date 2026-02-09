import { useQuery } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { galaxiesApi } from '../api'

/**
 * Hook to fetch a single galaxy by ID
 */
export function useGalaxy(id: string | undefined) {
  // Use a stable fallback key when no ID is provided to avoid cache pollution
  if (!id) {
    return useQuery({
      queryKey: ['galaxies', 'detail', '__no-id__'] as const,
      queryFn: () => Promise.reject(new Error('Galaxy ID is required')),
      enabled: false,
    })
  }

  return useQuery({
    queryKey: queryKeys.galaxies.detail(id),
    queryFn: () => galaxiesApi.get(id),
  })
}
