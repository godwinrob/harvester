import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { usersApi } from '../api'

/**
 * Hook to delete a single user
 */
export function useDeleteUser() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => usersApi.delete(id),
    onSuccess: (_data, id) => {
      // Remove from cache
      queryClient.removeQueries({ queryKey: queryKeys.users.detail(id) })
      // Invalidate lists
      queryClient.invalidateQueries({ queryKey: queryKeys.users.lists() })
    },
  })
}
