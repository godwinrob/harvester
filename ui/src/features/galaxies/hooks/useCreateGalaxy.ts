import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { galaxiesApi } from '../api'
import type { CreateGalaxyInput } from '../types'

/**
 * Hook to create a new galaxy
 */
export function useCreateGalaxy() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: CreateGalaxyInput) => galaxiesApi.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.galaxies.lists() })
    },
  })
}
