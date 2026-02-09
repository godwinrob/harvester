import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { usersApi } from '../api'
import type { User, UpdateUserInput } from '../types'

/**
 * Hook to update an existing user with optimistic updates
 */
export function useUpdateUser() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateUserInput }) =>
      usersApi.update(id, data),

    // Optimistic update
    onMutate: async ({ id, data }) => {
      // Cancel outgoing refetches
      await queryClient.cancelQueries({ queryKey: queryKeys.users.detail(id) })

      // Snapshot current value
      const previousUser = queryClient.getQueryData<User>(queryKeys.users.detail(id))

      // Optimistically update the cache
      if (previousUser) {
        queryClient.setQueryData<User>(queryKeys.users.detail(id), {
          ...previousUser,
          ...data,
          dateUpdated: new Date().toISOString(),
        })
      }

      return { previousUser }
    },

    onError: (_error, { id }, context) => {
      // Rollback on error
      if (context?.previousUser) {
        queryClient.setQueryData(queryKeys.users.detail(id), context.previousUser)
      }
    },

    onSettled: (_data, _error, { id }) => {
      // Refetch to ensure consistency
      queryClient.invalidateQueries({ queryKey: queryKeys.users.detail(id) })
      queryClient.invalidateQueries({ queryKey: queryKeys.users.lists() })
    },
  })
}
