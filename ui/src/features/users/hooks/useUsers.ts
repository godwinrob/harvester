import { useQuery, keepPreviousData } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { usersApi } from '../api'
import type { UserListParams } from '../types'

/**
 * Hook to fetch paginated list of users
 */
export function useUsers(params: UserListParams = {}) {
  return useQuery({
    queryKey: queryKeys.users.list(params),
    queryFn: () => usersApi.list(params),
    placeholderData: keepPreviousData, // Smooth pagination transitions
  })
}
