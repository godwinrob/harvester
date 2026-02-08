import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { resourcesApi } from '../api'
import type { Resource } from '../types'
import type { PaginatedResponse } from '@/lib/api'

/**
 * Hook to bulk delete resources (max 100) with optimistic updates
 */
export function useBulkDeleteResources() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (ids: string[]) => resourcesApi.bulkDelete(ids),

    onMutate: async (ids) => {
      await queryClient.cancelQueries({ queryKey: queryKeys.resources.lists() })

      const previousLists = new Map<readonly unknown[], PaginatedResponse<Resource>>()

      const idsToDelete = new Set(ids)

      queryClient
        .getQueriesData<PaginatedResponse<Resource>>({ queryKey: queryKeys.resources.lists() })
        .forEach(([key, data]) => {
          if (data) {
            previousLists.set(key, data)

            // Single-pass filter: remove deleted items and count removals
            let removedCount = 0
            const filteredItems = data.items.filter((resource) => {
              if (idsToDelete.has(resource.id)) {
                removedCount++
                return false
              }
              return true
            })

            queryClient.setQueryData<PaginatedResponse<Resource>>(key, {
              ...data,
              items: filteredItems,
              total: data.total - removedCount,
            })
          }
        })

      return { previousLists }
    },

    onError: (_error, _ids, context) => {
      context?.previousLists.forEach((data, key) => {
        queryClient.setQueryData(key, data)
      })
    },

    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.resources.all })
    },
  })
}
