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

      queryClient
        .getQueriesData<PaginatedResponse<Galaxy>>({ queryKey: queryKeys.galaxies.lists() })
        .forEach(([key, data]) => {
          if (data) {
            previousLists.set(key, data)
            queryClient.setQueryData<PaginatedResponse<Galaxy>>(key, {
              ...data,
              items: data.items.filter((galaxy) => !ids.includes(galaxy.id)),
              total: data.total - ids.filter((id) => data.items.some((g) => g.id === id)).length,
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
