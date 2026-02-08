import { useQuery, keepPreviousData } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { galaxiesApi } from '../api'
import type { GalaxyListParams } from '../types'

/**
 * Hook to fetch paginated list of galaxies
 */
export function useGalaxies(params: GalaxyListParams = {}) {
  return useQuery({
    queryKey: queryKeys.galaxies.list(params),
    queryFn: () => galaxiesApi.list(params),
    placeholderData: keepPreviousData,
  })
}
