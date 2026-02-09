import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { galaxiesApi } from '../api'

/**
 * Hook to delete a single galaxy
 */
export function useDeleteGalaxy() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => galaxiesApi.delete(id),
    onSuccess: (_data, id) => {
      queryClient.removeQueries({ queryKey: queryKeys.galaxies.detail(id) })
      queryClient.invalidateQueries({ queryKey: queryKeys.galaxies.lists() })
    },
  })
}
