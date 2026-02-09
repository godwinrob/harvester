import type { PaginationParams } from '@/lib/api'

/**
 * Centralized query keys for TanStack Query
 * Using factory pattern for type-safe and consistent cache invalidation
 */
export const queryKeys = {
  // Users
  users: {
    all: ['users'] as const,
    lists: () => [...queryKeys.users.all, 'list'] as const,
    list: (params: PaginationParams & { search?: string }) =>
      [...queryKeys.users.lists(), params] as const,
    details: () => [...queryKeys.users.all, 'detail'] as const,
    detail: (id: string) => [...queryKeys.users.details(), id] as const,
  },

  // Galaxies
  galaxies: {
    all: ['galaxies'] as const,
    lists: () => [...queryKeys.galaxies.all, 'list'] as const,
    list: (params: PaginationParams & { search?: string }) =>
      [...queryKeys.galaxies.lists(), params] as const,
    details: () => [...queryKeys.galaxies.all, 'detail'] as const,
    detail: (id: string) => [...queryKeys.galaxies.details(), id] as const,
    byName: (name: string) => [...queryKeys.galaxies.all, 'name', name] as const,
  },

  // Resources
  resources: {
    all: ['resources'] as const,
    lists: () => [...queryKeys.resources.all, 'list'] as const,
    list: (params: PaginationParams & { galaxyId?: string; search?: string }) =>
      [...queryKeys.resources.lists(), params] as const,
    details: () => [...queryKeys.resources.all, 'detail'] as const,
    detail: (id: string) => [...queryKeys.resources.details(), id] as const,
    byName: (name: string) => [...queryKeys.resources.all, 'name', name] as const,
    byGalaxy: (galaxyId: string) =>
      [...queryKeys.resources.all, 'galaxy', galaxyId] as const,
  },
} as const
