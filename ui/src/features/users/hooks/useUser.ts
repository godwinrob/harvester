import { useQuery } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { usersApi } from '../api'

/**
 * Hook to fetch a single user by ID
 */
export function useUser(id: string | undefined) {
  return useQuery({
    queryKey: queryKeys.users.detail(id!),
    queryFn: () => usersApi.get(id!),
    enabled: !!id,
  })
}
