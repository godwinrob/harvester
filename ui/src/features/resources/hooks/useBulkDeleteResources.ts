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

      queryClient
        .getQueriesData<PaginatedResponse<Resource>>({ queryKey: queryKeys.resources.lists() })
        .forEach(([key, data]) => {
          if (data) {
            previousLists.set(key, data)
            queryClient.setQueryData<PaginatedResponse<Resource>>(key, {
              ...data,
              items: data.items.filter((resource) => !ids.includes(resource.id)),
              total: data.total - ids.filter((id) => data.items.some((r) => r.id === id)).length,
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
