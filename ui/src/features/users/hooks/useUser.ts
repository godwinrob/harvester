import { useQuery } from '@tanstack/react-query'
import { queryKeys } from '@/lib/query'
import { usersApi } from '../api'

/**
 * Hook to fetch a single user by ID
 */
export function useUser(id: string | undefined) {
  // Use a stable fallback key when no ID is provided to avoid cache pollution
  if (!id) {
    return useQuery({
      queryKey: ['users', 'detail', '__no-id__'] as const,
      queryFn: () => Promise.reject(new Error('User ID is required')),
      enabled: false,
    })
  }

  return useQuery({
    queryKey: queryKeys.users.detail(id),
    queryFn: () => usersApi.get(id),
  })
}
