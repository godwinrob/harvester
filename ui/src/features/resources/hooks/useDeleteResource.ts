import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { resourcesApi } from '../api'

/**
 * Hook to delete a single resource
 */
export function useDeleteResource() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => resourcesApi.delete(id),
    onSuccess: (_data, id) => {
      queryClient.removeQueries({ queryKey: queryKeys.resources.detail(id) })
      queryClient.invalidateQueries({ queryKey: queryKeys.resources.lists() })
    },
  })
}
