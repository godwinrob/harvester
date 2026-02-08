import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { resourcesApi } from '../api'
import type { CreateResourceInput } from '../types'

/**
 * Hook to create a new resource
 */
export function useCreateResource() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: CreateResourceInput) => resourcesApi.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.resources.lists() })
    },
  })
}
