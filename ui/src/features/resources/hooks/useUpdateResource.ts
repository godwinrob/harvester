import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { resourcesApi } from '../api'
import type { Resource, UpdateResourceInput } from '../types'

/**
 * Hook to update an existing resource with optimistic updates
 */
export function useUpdateResource() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateResourceInput }) =>
      resourcesApi.update(id, data),

    onMutate: async ({ id, data }) => {
      await queryClient.cancelQueries({ queryKey: queryKeys.resources.detail(id) })

      const previousResource = queryClient.getQueryData<Resource>(queryKeys.resources.detail(id))

      if (previousResource) {
        queryClient.setQueryData<Resource>(queryKeys.resources.detail(id), {
          ...previousResource,
          ...data,
          updatedAtDate: new Date().toISOString(),
        })
      }

      return { previousResource }
    },

    onError: (_error, { id }, context) => {
      if (context?.previousResource) {
        queryClient.setQueryData(queryKeys.resources.detail(id), context.previousResource)
      }
    },

    onSettled: (_data, _error, { id }) => {
      queryClient.invalidateQueries({ queryKey: queryKeys.resources.detail(id) })
      queryClient.invalidateQueries({ queryKey: queryKeys.resources.lists() })
    },
  })
}
