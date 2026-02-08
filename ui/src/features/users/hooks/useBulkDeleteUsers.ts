import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { usersApi } from '../api'
import type { User } from '../types'
import type { PaginatedResponse } from '@/lib/api'

/**
 * Hook to bulk delete users (max 100) with optimistic updates
 */
export function useBulkDeleteUsers() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (ids: string[]) => usersApi.bulkDelete(ids),

    onMutate: async (ids) => {
      // Cancel outgoing refetches
      await queryClient.cancelQueries({ queryKey: queryKeys.users.lists() })

      // Snapshot and optimistically remove from all list caches
      const previousLists = new Map<readonly unknown[], PaginatedResponse<User>>()

      const idsToDelete = new Set(ids)

      queryClient
        .getQueriesData<PaginatedResponse<User>>({ queryKey: queryKeys.users.lists() })
        .forEach(([key, data]) => {
          if (data) {
            previousLists.set(key, data)

            // Single-pass filter: remove deleted items and count removals
            let removedCount = 0
            const filteredItems = data.items.filter((user) => {
              if (idsToDelete.has(user.id)) {
                removedCount++
                return false
              }
              return true
            })

            queryClient.setQueryData<PaginatedResponse<User>>(key, {
              ...data,
              items: filteredItems,
              total: data.total - removedCount,
            })
          }
        })

      return { previousLists }
    },

    onError: (_error, _ids, context) => {
      // Restore all previous list states
      context?.previousLists.forEach((data, key) => {
        queryClient.setQueryData(key, data)
      })
    },

    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.all })
    },
  })
}
