import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { galaxiesApi } from '../api'
import type { Galaxy } from '../types'
import type { PaginatedResponse } from '@/lib/api'

/**
 * Hook to bulk delete galaxies (max 100) with optimistic updates
 */
export function useBulkDeleteGalaxies() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (ids: string[]) => galaxiesApi.bulkDelete(ids),

    onMutate: async (ids) => {
      await queryClient.cancelQueries({ queryKey: queryKeys.galaxies.lists() })

      const previousLists = new Map<readonly unknown[], PaginatedResponse<Galaxy>>()

      const idsToDelete = new Set(ids)

      queryClient
        .getQueriesData<PaginatedResponse<Galaxy>>({ queryKey: queryKeys.galaxies.lists() })
        .forEach(([key, data]) => {
          if (data) {
            previousLists.set(key, data)

            // Single-pass filter: remove deleted items and count removals
            let removedCount = 0
            const filteredItems = data.items.filter((galaxy) => {
              if (idsToDelete.has(galaxy.id)) {
                removedCount++
                return false
              }
              return true
            })

            queryClient.setQueryData<PaginatedResponse<Galaxy>>(key, {
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
      queryClient.invalidateQueries({ queryKey: queryKeys.galaxies.all })
    },
  })
}
