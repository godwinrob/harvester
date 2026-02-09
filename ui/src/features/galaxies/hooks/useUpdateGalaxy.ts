import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { galaxiesApi } from '../api'
import type { Galaxy, UpdateGalaxyInput } from '../types'

/**
 * Hook to update an existing galaxy with optimistic updates
 */
export function useUpdateGalaxy() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateGalaxyInput }) =>
      galaxiesApi.update(id, data),

    onMutate: async ({ id, data }) => {
      await queryClient.cancelQueries({ queryKey: queryKeys.galaxies.detail(id) })

      const previousGalaxy = queryClient.getQueryData<Galaxy>(queryKeys.galaxies.detail(id))

      if (previousGalaxy) {
        queryClient.setQueryData<Galaxy>(queryKeys.galaxies.detail(id), {
          ...previousGalaxy,
          ...data,
          dateUpdated: new Date().toISOString(),
        })
      }

      return { previousGalaxy }
    },

    onError: (_error, { id }, context) => {
      if (context?.previousGalaxy) {
        queryClient.setQueryData(queryKeys.galaxies.detail(id), context.previousGalaxy)
      }
    },

    onSettled: (_data, _error, { id }) => {
      queryClient.invalidateQueries({ queryKey: queryKeys.galaxies.detail(id) })
      queryClient.invalidateQueries({ queryKey: queryKeys.galaxies.lists() })
    },
  })
}
