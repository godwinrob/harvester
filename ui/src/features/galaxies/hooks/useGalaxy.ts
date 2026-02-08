import { useQuery } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { galaxiesApi } from '../api'

/**
 * Hook to fetch a single galaxy by ID
 */
export function useGalaxy(id: string | undefined) {
  return useQuery({
    queryKey: queryKeys.galaxies.detail(id!),
    queryFn: () => galaxiesApi.get(id!),
    enabled: !!id,
  })
}
