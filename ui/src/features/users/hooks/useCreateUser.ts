import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { usersApi } from '../api'
import type { CreateUserInput } from '../types'

/**
 * Hook to create a new user
 */
export function useCreateUser() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: CreateUserInput) => usersApi.create(data),
    onSuccess: () => {
      // Invalidate user lists to refetch
      queryClient.invalidateQueries({ queryKey: queryKeys.users.lists() })
    },
  })
}
