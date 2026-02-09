import { useMutation, useQueryClient } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { usersApi } from '../api'
import type { CreateUserInput } from '../types'

/**
 * Hook to bulk create users (max 100)
 */
export function useBulkCreateUsers() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (items: CreateUserInput[]) => usersApi.bulkCreate(items),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.lists() })
    },
  })
}
