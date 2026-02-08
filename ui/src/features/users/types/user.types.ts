// User types matching the backend API

export type UserRole = 'ADMIN' | 'USER'

export interface User {
  id: string
  name: string
  email: string
  roles: UserRole[]
  guild: string | null
  enabled: boolean
  dateCreated: string
  dateUpdated: string
}

export interface CreateUserInput {
  name: string
  email: string
  password: string
  passwordConfirm: string
  roles: UserRole[]
  guild?: string
}

export interface UpdateUserInput {
  name?: string
  email?: string
  roles?: UserRole[]
  guild?: string
  enabled?: boolean
}

export interface UpdateUserRoleInput {
  roles: UserRole[]
}

// Bulk operation types
export interface BulkCreateUserInput {
  items: CreateUserInput[]
}

export interface BulkUpdateUserInput {
  items: Array<{
    id: string
    data: UpdateUserInput
  }>
}

export interface BulkDeleteUserInput {
  ids: string[]
}

// Query params
export interface UserListParams {
  page?: number
  rows?: number
  orderBy?: {
    field: string
    direction: 'ASC' | 'DESC'
  }
  name?: string
  email?: string
  startCreatedDate?: string
  endCreatedDate?: string
}
